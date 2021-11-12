package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

type interceptWriter struct {
	http.ResponseWriter
	status int
	body   *bytes.Buffer
}

func (w *interceptWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *interceptWriter) WriteHeader(statusCode int) {
	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// ロギング用ミドルウェア
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bufOfRequestBody, _ := io.ReadAll(r.Body)
		// [For Request Body] 消費されてしまったRequest Bodyを修復する
		r.Body = io.NopCloser(bytes.NewBuffer(bufOfRequestBody))

		// [For Response Body] ResponseWriterへ割り込む (interceptする)
		iw := &interceptWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: w, // 元のwriterインスタンスを埋め込む
		}
		w = iw

		// アプリケーション処理
		next.ServeHTTP(w, r)

		// レスポンスがエラーの場合、ログ出力する
		if iw.status >= 400 {
			var w io.Writer = os.Stdout
			fmt.Fprintln(w, "error status: ", iw.status)
			fmt.Fprintln(w, "request body: ", bytes.NewBuffer(bufOfRequestBody))
			fmt.Fprintln(w, "response body:", iw.body.String())
		}
	})
}
