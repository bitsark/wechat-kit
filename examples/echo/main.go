package main

import (
	"context"
	"encoding/xml"
	"log"
	"net/http"
	"time"

	"github.com/bitsark/wechat-kit/offiaccount"
)

func echo(ctx context.Context, inbound []byte) (ret interface{}, err error) {
	req := &offiaccount.TextReqMessage{}
	if err = xml.Unmarshal(inbound, req); err != nil {
		return
	}
	ret = &offiaccount.TextRespMessage{
		RespMessageBase: offiaccount.RespMessageBase{
			ToUserName:   offiaccount.ToUserName{Text: req.FromUserName.Text},
			FromUserName: offiaccount.FromUserName{Text: req.ToUserName.Text},
			CreateTime:   time.Now().Unix(),
			MessageType:  offiaccount.MessageType{Text: offiaccount.MessageTypeText},
		},
		MessageContent: offiaccount.MessageContent{Text: "wechat-kit: " + req.MessageContent.Text},
	}
	return
}

func main() {
	addr := ":8080"
	token := "w5Lip5ZN5LuB5aSy"

	wexinMux := offiaccount.NewMessageMux()

	wexinMux.HandlerFunc(offiaccount.MessageTypeText, echo)

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/offiaccount/",
		offiaccount.Authentication(token,
			offiaccount.EchoStr(wexinMux.StdHandlerFunc())))

	server := &http.Server{
		Handler:           serveMux,
		Addr:              addr,
		ReadHeaderTimeout: 100 * time.Millisecond,
		WriteTimeout:      500 * time.Millisecond,
	}

	log.Println("wechat offiaccount server listening on " + addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
