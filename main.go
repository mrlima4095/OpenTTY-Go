package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
)

const (
	DEFAULT_PORT   = "31522"
	END_OF_OUTPUT  = "<<<END>>>"
)

var (
	username string
	hostname string
	path     string
	conn     net.Conn
)

func parseAddress(arg string) string {
	if strings.Contains(arg, ":") {
		return arg
	}
	return arg + ":" + DEFAULT_PORT
}

func clearTerminal() {
	fmt.Print("\033[2J\033[H")
}

func sendAndReceive(cmd string) string {
	_, err := fmt.Fprintln(conn, cmd)
	if err != nil {
		return "[erro ao enviar comando]"
	}

	var output strings.Builder
	reader := bufio.NewReader(conn)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			output.WriteString("[erro na leitura da resposta]\n")
			break
		}
		line = strings.TrimRight(line, "\r\n")
		if line == END_OF_OUTPUT {
			break
		}
		output.WriteString(line + "\n")
	}
	return output.String()
}

func updatePromptInfo() {
	username = sendAndReceive("whoami")
	hostname = sendAndReceive("hostname")
	path = sendAndReceive("pwd")
}

func prompt() string {
	return fmt.Sprintf("%s@%s:%s$ ", username, hostname, path)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: optty <host[:port]>")
		return
	}

	address := parseAddress(os.Args[1])
	var err error
	conn, err = net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Erro ao conectar:", err)
		return
	}
	defer conn.Close()

	clearTerminal()
	sendAndReceive("execute install nano; touch; add screen.title=OpenTTY GO; add screen.title=OpenTTY Go; add screen.back.label=Exit; add screen.back=execute exit; add screen.button=Return; add screen.button.cmd=exec unalias xterm & xterm & stop bind; add screen.fields=notes; add screen.notes.type=text; add screen.notes.value=You're accessing this Device via OpenTTY Golang.; install go-term; alias xterm=x11 make go-term; xterm;")
	updatePromptInfo()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		for range sig {
			clearTerminal()
			updatePromptInfo()
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt())

		if !scanner.Scan() {
			fmt.Println("\nSaindo...")
			return
		}

		cmd := scanner.Text()
		if cmd == "" {
			continue
		}

		output := sendAndReceive(cmd)
		fmt.Print(output)
	}
}
