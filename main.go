package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type AppState struct {
	Username     string
	Path         string
	History      []string
	StdoutText   string
	Build        string
	nanoContent  string
}

func LoadRMS(key string) string {
	filename := filepath.Join(".", key+".rms")
	data, err := os.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(data)
}

func WriteRMS(key, value string) {
	filename := filepath.Join(".", key+".rms")
	_ = os.WriteFile(filename, []byte(value), 0644)
}

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
		OpenHistoryViewer(state)
		return "[history viewer opened]"
	case "clear":
		state.StdoutText = ""
		return ""
	case "nano":
		OpenNanoEditor(state)
		return "[nano editor opened]"
	case "exit":
		os.Exit(0)
	}
	return fmt.Sprintf("You entered: %s", cmd)
}

func OpenNanoEditor(state *AppState) {
	nanoWin := fyne.CurrentApp().NewWindow("Nano")

	editor := widget.NewMultiLineEntry()
	editor.SetText(LoadRMS("nano"))

	saveBtn := widget.NewButton("Save", func() {
		text := editor.Text
		state.nanoContent = text
		WriteRMS("nano", text)
		nanoWin.Close()
	})

	cancelBtn := widget.NewButton("Cancel", func() {
		nanoWin.Close()
	})

	buttons := container.NewHBox(saveBtn, cancelBtn)
	layout := container.NewBorder(nil, buttons, nil, nil, editor)

	nanoWin.SetContent(layout)
	nanoWin.Resize(fyne.NewSize(500, 400))
	nanoWin.Show()
}

func OpenHistoryViewer(state *AppState) {
	historyWin := fyne.CurrentApp().NewWindow("Command History")

	historyList := widget.NewList(
		func() int {
			return len(state.History)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i int, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(state.History[i])
		},
	)

	closeBtn := widget.NewButton("Close", func() {
		historyWin.Close()
	})

	layout := container.NewBorder(nil, closeBtn, nil, nil, historyList)
	
	historyWin.SetContent(layout)
	historyWin.Resize(fyne.NewSize(400, 300))
	historyWin.Show()
}

func main() {
	a := app.New()
	w := a.NewWindow("OpenTTY 0.6.2")

	state := &AppState{
		Username:    LoadRMS("OpenRMS"),
		Path:        "/home/",
		History:     []string{},
		Build:       "2025-1.15-02x06",
		nanoContent: LoadRMS("nano"),
	}

	stdin := widget.NewEntry()
	stdin.SetPlaceHolder("Command")

	stdout := widget.NewMultiLineEntry()
	stdout.Wrapping = fyne.TextWrapWord
	stdout.SetText("Welcome to OpenTTY 0.6.2\nCopyright (C) 2025 - Mr. Lima\n")
	stdout.SetMinRowsVisible(25)

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
	term := container.NewBorder(nil, container.NewVBox(stdin, btns), nil, nil, stdout)

	w.SetContent(container.New(layout.NewMaxLayout(), term))
	w.Resize(fyne.NewSize(800, 600))

	w.SetCloseIntercept(func() {
		WriteRMS("nano", state.nanoContent)
		a.Quit()
	})

	w.ShowAndRun()
}
