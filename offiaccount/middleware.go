package offiaccount

import (
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
)

// https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Access_Overview.html
func Authentication(token string, next http.HandlerFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		} else {
			array := []string{
				token,
				query.Get("nonce"),
				query.Get("timestamp"),
			}
			sort.Strings(array)
			hash := sha1.New()
			for _, v := range array {
				io.WriteString(hash, v)
			}
			if fmt.Sprintf("%x", hash.Sum(nil)) != query.Get("signature") {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			} else {
				next(w, r)
			}
		}
	}
}

func EchoStr(next http.HandlerFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		} else {
			if echostr := query.Get("echostr"); len(echostr) > 0 {
				w.Header().Set("Content-Type", "text/plain; charset=utf-8")
				w.WriteHeader(http.StatusOK)
				io.WriteString(w, echostr)
			} else {
				next(w, r)
			}
		}
	}
}
