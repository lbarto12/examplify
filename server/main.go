package main

import (
	"fmt"
	"server/api/server"
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
}
