package main

import (
	"fmt"
	"net/http"
	"server/api/server"

	"github.com/go-chi/chi/v5"
)

func Main(server *server.Server) {
	fmt.Println(
		"Hello from main",
	)

	// // AI usage Demo
	// result, err := server.Services.Gemini.Models.GenerateContent(
	// 	context.Background(),
	// 	"gemini-2.5-flash",
	// 	genai.Text("Explain how AI works in a few words"),
	// 	nil,
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("result: ", result.Text())

	chi.Walk(server.Mux, func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Println(method, route)
		return nil
	})

}
