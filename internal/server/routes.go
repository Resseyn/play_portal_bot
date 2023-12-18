package server

import "net/http"

func CreateMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomePage)
	mux.HandleFunc("/success", PayPalychSuccessPaymentHandler)
	mux.HandleFunc("/fail", PayPalychFailPaymentHandler)
	mux.HandleFunc("/result", PayPalychPaymentHandler)

	return mux
}
