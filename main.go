package main

import (
	"github.com/charmbracelet/huh"
	"log"
	"os"
	"time"
)

func main() {
	// https://youtrack.jetbrains.com/issue/GO-16810/
	// https://github.com/charmbracelet/huh/issues/204
	if _, set := os.LookupEnv("IDEA_INITIAL_DIRECTORY"); set {
		time.Sleep(460 * time.Millisecond)
	}

	var userInput bool

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Can we proceed?").
				Affirmative("Yes!").
				Negative("No.").
				Value(&userInput),
		),
	).Run()
	if err != nil {
		log.Fatal(err)
	}

	if !userInput {
		println("No")
		os.Exit(0)
	}

	println("Yes")
}
