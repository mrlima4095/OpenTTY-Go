package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// AppState holds the main application data and state.
type AppState struct {
	Username   string
	Path       string
	History    []string
	StdoutText string
	Build      string
}

// LoadRMS simulates reading persistent data (replace with file or database as needed)
func LoadRMS(key string) string {
	// Example: Read from a file; you can enhance this as needed
	filename := filepath.Join(".", key+".rms")
	data, err := os.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(data)
}

// WriteRMS simulates writing persistent data
func WriteRMS(key, value string) {
	filename := filepath.Join(".", key+".rms")
	_ = os.WriteFile(filename, []byte(value), 0644)
}

// ProcessCommand handles a command string; stub for now
func (state *AppState) ProcessCommand(cmd string) string {
	cmd = strings.TrimSpace(cmd)
	if cmd == "" {
		return ""
	}
	state.History = append(state.History, cmd)
	switch cmd {
	case "help":
		return "Available commands: help, history, clear, nano, exit"
	case "history":
		return strings.Join(state.History, "\n")
	case "clear":
		state.StdoutText = ""
		return ""
	case "nano":
		return "(Nano editor not implemented yet)"
	case "exit":
		os.Exit(0)
	}
	return fmt.Sprintf("You entered: %s", cmd)
}

func main() {
	a := app.New()
	w := a.NewWindow("OpenTTY 0.6.2")

	// App state
	state := &AppState{
		Username: LoadRMS("OpenRMS"),
		Path:     "/home/",
		History:  []string{},
		Build:    "2025-1.15-02x06",
	}

	// UI widgets
	stdin := widget.NewEntry()
	stdin.SetPlaceHolder("Command")
	stdout := widget.NewMultiLineEntry()
	stdout.SetText(fmt.Sprintf("Welcome to OpenTTY 0.6.2\nCopyright (C) 2025 - Mr. Lima\n"))

	// Set up buttons
	executeBtn := widget.NewButton("Send", func() {
		command := stdin.Text
		output := state.ProcessCommand(command)
		if output != "" {
			stdout.SetText(stdout.Text + "\n" + output)
		}
		stdin.SetText("")
	})
	helpBtn := widget.NewButton("Help", func() {
		output := state.ProcessCommand("help")
		stdout.SetText(stdout.Text + "\n" + output)
	})
	nanoBtn := widget.NewButton("Nano", func() {
		output := state.ProcessCommand("nano")
		stdout.SetText(stdout.Text + "\n" + output)
	})
	clearBtn := widget.NewButton("Clear", func() {
		state.StdoutText = ""
		stdout.SetText("")
	})
	historyBtn := widget.NewButton("History", func() {
		output := state.ProcessCommand("history")
		stdout.SetText(stdout.Text + "\n" + output)
	})

	btns := container.NewHBox(executeBtn, helpBtn, nanoBtn, clearBtn, historyBtn)
	content := container.NewVBox(stdout, stdin, btns)

	w.SetContent(content)
	w.Resize(fyne.NewSize(520, 360))

	// Save nano content and exit
	w.SetCloseIntercept(func() {
		WriteRMS("nano", "") // Save nano content here, implement editor later
		a.Quit()
	})

	w.ShowAndRun()
}