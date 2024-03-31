package main

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
	"github.com/go-flexi/codegenerator/generator/backend"
	"github.com/go-flexi/codegenerator/openai"
	"github.com/go-flexi/codegenerator/ui/core"
)

func main() {
	textToCopy := "Hello, clipboard! test"
	err := clipboard.WriteAll(textToCopy)
	if err != nil {
		fmt.Println("Failed to copy text to the clipboard:", err)
		return
	}

	generator := backend.NewGenerator(openai.NewAPI(
		"api token",
		openai.DefaultConfig(),
	), "go-flexi", "ecom-backend")

	switch os.Args[1] {
	case "core":
		core := core.NewCore(generator)
		core.View()
	}
}
