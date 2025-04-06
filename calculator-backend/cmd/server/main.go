package main

import (
	"github.com/WUZHUHE38/calculator-fullstack/calculator-backend/gen/go/calculator/v1/calculatorv1connect"
	"github.com/WUZHUHE38/calculator-fullstack/calculator-backend/service"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
)

func main() {
	calculator := &service.CalculatorService{}
	mux := http.NewServeMux()
	path, handler := calculatorv1connect.NewCalculatorServiceHandler(calculator)
	mux.Handle(path, handler)

	// CORS配置
	corsHandler := cors.AllowAll().Handler(h2c.NewHandler(mux, &http2.Server{}))
	log.Println("Server started on :8080")
	err := http.ListenAndServe(
		":8080",
		corsHandler,
	)
	if err != nil {
		return
	}
}
