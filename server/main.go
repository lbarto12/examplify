package main

import (
	"fmt"
	"net/http"
	"server/api/server"

	"github.com/go-chi/chi/v5"
)

// type Res struct {
// 	Description string `json:"description"`
// }

// func (Res) Describe() string {
// 	return `
// 		{
// 			'description': a description of the image
// 		}
// 		`
// }

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

	// core, err := core.NewCore(&server.Services)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// img, err := imageanalysis.NewImageAnalyzer[Res](imageanalysis.NewImageAnalyzerParams{
	// 	ObjectStore: server.Services.Minio,
	// 	AI:          server.Services.OpenAI,
	// 	Postgres:    server.Services.Postgres,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// read, err := url.Parse("https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSgVfHORQFLyUf_rNove-xUmxIskDeMJ63REz_YIMQ6S0vCyQdkBvJos4igKspvCgpqnpy8h0xM--1uckzZIxDgyoHy37-MowkF-YzvVx8&s=10")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// res, err := img.AnalyzeURL(context.Background(), read)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(res.Description)
}
