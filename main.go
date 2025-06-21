package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type OpenTTY struct {
	username string
	path     string
	stdout   string
	stdin    string
	rmsDir   string
}

func NewOpenTTY() *OpenTTY {
	rmsPath := filepath.Join(os.TempDir(), "opentty_rms")
	os.MkdirAll(rmsPath, 0755)
	return &OpenTTY{
		username: loadRMS("OpenRMS", rmsPath),
		path:     "/home/",
		stdout:   "",
		stdin:    "",
		rmsDir:   rmsPath,
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

// loadRMS simula a leitura de dados persistentes (RMS)
func loadRMS(name, basePath string) string {
	path := filepath.Join(basePath, name)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(data)
}

// writeRMS simula a escrita em um RMS (persistência simples por arquivo)
func writeRMS(name, content, basePath string) {
	path := filepath.Join(basePath, name)
	_ = ioutil.WriteFile(path, []byte(content), 0644)
}

// getcontent retorna o conteúdo de um pseudo-arquivo ou string literal
func (otty *OpenTTY) getcontent(file string) string {
	if strings.HasPrefix(file, "/") {
		return otty.read(file)
	}
	return file
}

// read faz leitura condicional conforme a origem do caminho (/home/, /mnt/, ou asset)
func (otty *OpenTTY) read(filename string) string {
	if strings.HasPrefix(filename, "/home/") {
		return loadRMS(strings.TrimPrefix(filename, "/home/"), otty.rmsDir)
	} else if strings.HasPrefix(filename, "/mnt/") {
		sysPath := filepath.FromSlash(strings.TrimPrefix(filename, "/mnt/"))
		data, err := ioutil.ReadFile(sysPath)
		if err != nil {
			return fmt.Sprintf("read error: %v", err)
		}
		return string(data)
	} else {
		assetPath := filepath.Join("assets", strings.TrimPrefix(filename, "/"))
		data, err := ioutil.ReadFile(assetPath)
		if err != nil {
			return fmt.Sprintf("asset error: %v", err)
		}
		return string(data)
	}
}

func getCommand(input string) string {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return ""
	}
	return parts[0]
}

func getArgument(input string) string {
	parts := strings.Fields(input)
	if len(parts) <= 1 {
		return ""
	}
	return strings.Join(parts[1:], " ")
}

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
		} else {
			break
		}
	}
}
