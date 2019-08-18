/// [Topic1] 関数にインタフェース実装させる
package main

// HWriter はhttp.ResponseWriterの代理
type HWriter struct{}

func (p HWriter) Write(data []byte) (n int, err error) {
	return 0, nil
}

// HRequest はhttp.Requestの代理
type HRequest struct{}

// HHandlerFunc はhttp.HandlerFuncの代わり
// HServerHTTPメソッド(http.ServeHTTP)を実装して
// 関数がこの型へキャストすることでHHander(http.Handler)
// インターフェースを実装できるようにする役割
type HHandlerFunc func(w HWriter, r *HRequest)

// HServeHTTP はHHandlerFunc型が持つメソッドでhttp.ServeHTTPメソッドの代わり
func (f HHandlerFunc) HServeHTTP(w HWriter, r *HRequest) {
	f(w, r)
}

// HHandler はhttp.Handlerの代わり
type HHandler interface {
	HServeHTTP(w HWriter, r *HRequest)
}

// HHandle はhttp.Handleの代わり
func HHandle(pattern string, handler HHandler) {
	// process
}
