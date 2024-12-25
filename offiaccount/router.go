package offiaccount

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"
	"sync"
	"time"
)

type (
	MessageHandlerFunc func(context.Context, []byte) (interface{}, error)

	MessageMux struct {
		rwlock   *sync.RWMutex
		handlers map[string]MessageHandlerFunc

		Unimplemented MessageHandlerFunc
	}
)

func NewMessageMux() *MessageMux {
	return &MessageMux{
		rwlock:   &sync.RWMutex{},
		handlers: make(map[string]MessageHandlerFunc),
	}
}

func (mux *MessageMux) HandlerFunc(typ string, f MessageHandlerFunc) {
	mux.rwlock.Lock()
	defer mux.rwlock.Unlock()
	mux.handlers[typ] = f //TODO
}

type baseMessageKey struct{}

func GetBaseMessage(ctx context.Context) *ReqMessageBase {
	return ctx.Value(baseMessageKey{}).(*ReqMessageBase)
}

func (mux *MessageMux) StdHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		inbound, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, http.StatusText(http.StatusBadRequest)+err.Error())
		} else {
			if outbound, err := mux.process(r.Context(), inbound); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				io.WriteString(w, http.StatusText(http.StatusInternalServerError))
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write(outbound)
			}
		}
	}
}

func (mux *MessageMux) process(ctx context.Context, inbound []byte) (outbound []byte, err error) {
	req := &ReqMessageBase{}
	if err = xml.Unmarshal(inbound, req); err == nil {
		if f, exists := func() (f MessageHandlerFunc, exists bool) {
			mux.rwlock.RLock()
			defer mux.rwlock.RUnlock()
			f, exists = mux.handlers[req.MessageType.Text]
			return
		}(); exists {
			var ret interface{}
			ctx = context.WithValue(ctx, baseMessageKey{}, req)
			if ret, err = f(ctx, inbound); err == nil {
				outbound, err = xml.Marshal(ret)
			}
		} else {
			var ret interface{}
			ctx = context.WithValue(ctx, baseMessageKey{}, req)
			if ret, err = defaultUnimplemented(ctx, inbound); err == nil {
				outbound, err = xml.Marshal(ret)
			}
		}
	}
	return
}

var defaultUnimplemented = func(ctx context.Context, _ []byte) (ret interface{}, err error) {
	base := GetBaseMessage(ctx)
	ret = &TextRespMessage{
		RespMessageBase: RespMessageBase{
			ToUserName:   ToUserName{Text: base.FromUserName.Text},
			FromUserName: FromUserName{Text: base.ToUserName.Text},
			CreateTime:   time.Now().Unix(),
			MessageType:  MessageType{Text: "text"},
		},
		MessageContent: MessageContent{Text: base.MessageType.Text + " unimplemented"},
	}
	return
}
