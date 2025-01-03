package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os/exec"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) RunScript(script string) {

	// Create a command to execute the script
	cmd := exec.Command("sh", "-c", script)

	// Capture the output of the script
	var out bytes.Buffer
	cmd.Stdout = &out

	// Run the command
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Print the output of the script
	fmt.Println(out.String())
}

func (a *App) RunScriptIncrement(script string) {

	// Create a command to execute the script
	cmd := exec.Command("sh", "-c", script)

	// Capture the output of the script
	var out bytes.Buffer
	cmd.Stdout = &out

	// Run the command
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Print the output of the script
	fmt.Println(out.String())
}
