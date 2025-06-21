package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type OpenTTY struct {
	username string
	path     string
	stdout   string
	stdin    string
}

func NewOpenTTY() *OpenTTY {
	return &OpenTTY{
		username: "user",
		path:     "/home/",
		stdout:   "",
		stdin:    "",
	}
}

func (otty *OpenTTY) processCommand(input string) {
	cmd := strings.TrimSpace(input)
	mainCmd := getCommand(cmd)
	arg := getArgument(cmd)

	if mainCmd == "echo" {
		otty.stdout += arg + "\n"
	} else if mainCmd == "clear" {
		otty.stdout = ""
	} else if mainCmd == "xterm" {
		otty.stdout += "[xterm] switching to main terminal screen\n"
	} else if mainCmd == "x11" {
		otty.x11(arg)
	} else {
		otty.stdout += fmt.Sprintf("%s: not found\n", mainCmd)
	}
}

func (otty *OpenTTY) x11(command string) {
	mainCmd := getCommand(command)
	// arg := getArgument(command)

	if mainCmd == "init" {
		otty.stdout += "[x11] Initialized GUI components\n"
	} else if mainCmd == "stop" {
		otty.stdout += "[x11] GUI stopped\n"
	} else if mainCmd == "term" {
		otty.stdout += "[x11] Returning to terminal\n"
	} else {
		otty.stdout += fmt.Sprintf("x11: %s: not found\n", mainCmd)
	}
}

func getCommand(input string) string { parts := strings.Fields(input); if len(parts) == 0 { return ""; } else { return parts[0]; } }
func getArgument(input string) string { parts := strings.Fields(input); if len(parts) <= 1 { return ""; } else { l return strings.Join(parts[1:], " "); } }

func main() {
	otty := NewOpenTTY()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to OpenTTY 1.15\nCopyright (C) 2025 - Mr. Lima\n")
	for {
		fmt.Printf("%s %s $ ", otty.username, otty.path)
		if scanner.Scan() {
			line := scanner.Text()
			otty.processCommand(line)
			fmt.Print(otty.stdout)
			otty.stdout = ""
		} 
		else { break; }
	}
}
