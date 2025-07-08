package main

import "fmt"

type OpenTTY struct {
	*MIDlet
	cursorX     int
	cursorY     int
	player      *Player
	random      *Random
	runtime     *Runtime
	attributes  *Hashtable
	aliases     *Hashtable
	shell       *Hashtable
	functions   *Hashtable
	paths       *Hashtable
	desktops    *Hashtable
	trace       *Hashtable
	stack       *Vector
	history     *Vector
	sessions    *Vector
	username    string
	nanoContent string
	logs        string
	path        string
	build       string
	display     *Display
	form        *Form
	stdin       *TextField
	stdout      *StringItem
	eXECUTE     *Command
	hELP        *Command
	nANO        *Command
	cLEAR       *Command
	hISTORY     *Command
}

func (rcvr *OpenTTY) getAddress(command string) (int) {
	command = env(command.trim())
	mainCommand := rcvr.getCommand(command)
	argument := rcvr.getArgument(command)
	if mainCommand.equals("") {
		return rcvr.processCommand("ifconfig")
	} else {
		if try() {
			CONN, ok := Connector.open(fmt.Sprintf("%v%v", "datagram://", <<unimp_expr[*grammar.JConditionalExpr]>>)).(*DatagramConnection)
			if !ok {
				panic("XXX Cast fail for *parser.GoCastType")
			}
			OUT := NewByteArrayOutputStream()
			OUT.write(0x12)
			OUT.write(0x34)
			OUT.write(0x01)
			OUT.write(0x00)
			OUT.write(0x00)
			OUT.write(0x01)
			OUT.write(0x00)
			OUT.write(0x00)
			OUT.write(0x00)
			OUT.write(0x00)
			OUT.write(0x00)
			OUT.write(0x00)
			parts := rcvr.split(mainCommand, '.')
			for i := 0; i < len(parts); i++ {
				OUT.write(parts[i].length())
				OUT.write(parts[i].getBytes())
			}
			OUT.write(0x00)
			OUT.write(0x00)
			OUT.write(0x01)
			OUT.write(0x00)
			OUT.write(0x01)
			query := OUT.toByteArray()
			REQUEST := CONN.newDatagram(query, len(query))
			CONN.send(REQUEST)
			RESPONSE := CONN.newDatagram(512)
			CONN.receive(RESPONSE)
			CONN.close()
			data := RESPONSE.getData()
			if data[3]&0x0F != 0 {
				rcvr.echoCommand("not found")
				return 127
			}
			offset := 12
			for data[offset] != 0 {
				offset++
			}
			offset += 5
			if data[offset+2] == 0x00 && data[offset+3] == 0x01 {
				BUFFER := NewStringBuffer()
				for i := offset + 12; i < offset+16; i++ {
					BUFFER.append(data[i] & 0xFF)
					if i < offset+15 {
						BUFFER.append(".")
					}
				}
				echoCommand(BUFFER.toString())
			} else {
				rcvr.echoCommand("not found")
				return 127
			}
		} else if catch_IOException(e) {
			echoCommand(e.getMessage())
			return 1
		}
	}
	return 0
}
func (rcvr *OpenTTY) mIDletLogs(command string) (int) {
	command = env(command.trim())
	mainCommand := rcvr.getCommand(command)
	argument := rcvr.getArgument(command)
	if mainCommand.equals("") {
	} else if mainCommand.equals("clear") {
		rcvr.logs = ""
	} else if mainCommand.equals("swap") {
		writeRMS(<<unimp_expr[*grammar.JConditionalExpr]>>, rcvr.logs)
	} else if mainCommand.equals("view") {
		viewer(rcvr.form.getTitle(), rcvr.logs)
	} else if mainCommand.equals("add") {
		LEVEL := rcvr.getCommand(argument).toLowerCase()
		MESSAGE := rcvr.getArgument(argument)
		if !MESSAGE.equals("") {
			if LEVEL.equals("info") || LEVEL.equals("warn") || LEVEL.equals("debug") || LEVEL.equals("error") {
				rcvr.logs += fmt.Sprintf("%v%v%v%v%v%v%v", "[", LEVEL.toUpperCase(), "] ", split(Newjava.util.Date().toString(), ' ')[3], " ", MESSAGE, "\n")
			} else {
				rcvr.echoCommand(fmt.Sprintf("%v%v%v", "log: add: ", LEVEL, ": not found"))
				return 127
			}
		}
	} else {
		rcvr.echoCommand(fmt.Sprintf("%v%v%v", "log: ", mainCommand, ": not found"))
		return 127
	}
	return 0
}
func NewOpenTTY() (rcvr *OpenTTY) {
	rcvr = &OpenTTY{}
	rcvr.cursorX = 10
	rcvr.cursorY = 10
	rcvr.player = nil
	rcvr.random = NewRandom()
	rcvr.runtime = Runtime.getRuntime()
	rcvr.attributes = NewHashtable()
	rcvr.aliases = NewHashtable()
	rcvr.shell = NewHashtable()
	rcvr.functions = NewHashtable()
	rcvr.paths = NewHashtable()
	rcvr.desktops = NewHashtable()
	rcvr.trace = NewHashtable()
	rcvr.stack = NewVector()
	rcvr.history = NewVector()
	rcvr.sessions = NewVector()
	rcvr.username = rcvr.loadRMS("OpenRMS")
	rcvr.nanoContent = rcvr.loadRMS("nano")
	rcvr.logs = ""
	rcvr.path = "/home/"
	rcvr.build = "2025-1.15-02x14"
	rcvr.display = Display.getDisplay(rcvr)
	rcvr.form = NewForm(fmt.Sprintf("%v%v", "OpenTTY ", getAppProperty("MIDlet-Version")))
	rcvr.stdin = NewTextField("Command", "", 256, TextField.ANY)
	rcvr.stdout = NewStringItem("", fmt.Sprintf("%v%v%v", "Welcome to OpenTTY ", getAppProperty("MIDlet-Version"), "\nCopyright (C) 2025 - Mr. Lima\n"))
	rcvr.eXECUTE = NewCommand("Send", Command.OK, 1)
	rcvr.hELP = NewCommand("Help", Command.SCREEN, 2)
	rcvr.nANO = NewCommand("Nano", Command.SCREEN, 3)
	rcvr.cLEAR = NewCommand("Clear", Command.SCREEN, 4)
	rcvr.hISTORY = NewCommand("History", Command.SCREEN, 5)
	return
}
func (rcvr *OpenTTY) stringEditor(command string) (int) {
	command = env(command.trim())
	mainCommand := rcvr.getCommand(command)
	argument := rcvr.getArgument(command)
	if mainCommand.equals("") {
	} else if mainCommand.equals("-2u") {
		rcvr.nanoContent = rcvr.nanoContent.toUpperCase()
	} else if mainCommand.equals("-2l") {
		rcvr.nanoContent = rcvr.nanoContent.toLowerCase()
	} else if mainCommand.equals("-d") {
		rcvr.nanoContent = replace(rcvr.nanoContent, rcvr.split(argument, ' ')[0], "")
	} else if mainCommand.equals("-a") {
		rcvr.nanoContent = <<unimp_expr[*grammar.JConditionalExpr]>>
	} else if mainCommand.equals("-r") {
		rcvr.nanoContent = replace(rcvr.nanoContent, rcvr.split(argument, ' ')[0], rcvr.split(argument, ' ')[1])
	} else if mainCommand.equals("-l") {
		i := 0
		if try() {
			i = Integer.parseInt(argument)
		} else if catch_NumberFormatException(e) {
			echoCommand(e.getMessage())
			return 2
		}
		echoCommand(rcvr.split(rcvr.nanoContent, '\n')[i])
	} else if mainCommand.equals("-s") {
		i := 0
		if try() {
			i = Integer.parseInt(rcvr.getCommand(argument))
		} else if catch_NumberFormatException(e) {
			echoCommand(e.getMessage())
			return 2
		}
		lines := NewVector()
		div := rcvr.getArgument(argument)
		start := 0
		var index int
		for (index = rcvr.nanoContent.indexOf(div, start)) != -1 {
			lines.addElement(rcvr.nanoContent.substring(start, index))
			start = index + div.length()
		}
		if start < rcvr.nanoContent.length() {
			lines.addElement(rcvr.nanoContent.substring(start))
		}
		result := make([]string, len(lines))
		lines.copyInto(result)
		if i >= 0 && i < len(result) {
			echoCommand(result[i])
		} else {
			rcvr.echoCommand("null")
			return 1
		}
	} else if mainCommand.equals("-p") {
		contentLines := rcvr.split(rcvr.nanoContent, '\n')
		updatedContent := NewStringBuffer()
		for i := 0; i < len(contentLines); i++ {
			updatedContent.append(argument).append(contentLines[i]).append("\n")
		}
		rcvr.nanoContent = updatedContent.toString().trim()
	} else if mainCommand.equals("-v") {
		lines := rcvr.split(rcvr.nanoContent, '\n')
		reversed := NewStringBuffer()
		for i := len(lines) - 1; i >= 0; i-- {
			reversed.append(lines[i]).append("\n")
		}
		rcvr.nanoContent = reversed.toString().trim()
	} else {
		return 127
	}
	return 0
}
func (rcvr *OpenTTY) about(script string) {
	if script == nil || script.length() == 0 {
		rcvr.warnCommand("About", rcvr.env("OpenTTY $VERSION\n(C) 2025 - Mr. Lima"))
		return
	}
	PKG := rcvr.parseProperties(rcvr.getcontent(script))
	if PKG.containsKey("name") {
		rcvr.echoCommand(fmt.Sprintf("%v%v%v", PKG.get("name").(string), " ", PKG.get("version").(string)))
	}
	if PKG.containsKey("description") {
		rcvr.echoCommand(PKG.get("description").(string))
	}
}
func (rcvr *OpenTTY) addDirectory(fullPath string) {
	isDirectory := fullPath.endsWith("/")
	if !rcvr.paths.containsKey(fullPath) {
		if isDirectory {
			rcvr.paths.put(fullPath, []string{".."})
			parentPath := fullPath.substring(0, fullPath.lastIndexOf('/', fullPath.length()-2)+1)
			parentContents, ok := rcvr.paths.get(parentPath).([]string)
			if !ok {
				panic("XXX Cast fail for *parser.GoCastType")
			}
			updatedContents := NewVector()
			if parentContents != nil {
				for k := 0; k < len(parentContents); k++ {
					updatedContents.addElement(parentContents[k])
				}
			}
			dirName := fullPath.substring(parentPath.length(), fullPath.length()-1)
			updatedContents.addElement(fmt.Sprintf("%v%v", dirName, "/"))
			newContents := make([]string, len(updatedContents))
			updatedContents.copyInto(newContents)
			rcvr.paths.put(parentPath, newContents)
		} else {
			parentPath := fullPath.substring(0, fullPath.lastIndexOf('/')+1)
			fileName := fullPath.substring(fullPath.lastIndexOf('/') + 1)
			parentContents, ok := rcvr.paths.get(parentPath).([]string)
			if !ok {
				panic("XXX Cast fail for *parser.GoCastType")
			}
			updatedContents := NewVector()
			if parentContents != nil {
				for k := 0; k < len(parentContents); k++ {
					updatedContents.addElement(parentContents[k])
				}
			}
			updatedContents.addElement(fileName)
			newContents := make([]string, len(updatedContents))
			updatedContents.copyInto(newContents)
			rcvr.paths.put(parentPath, newContents)
		}
	}
}
func (rcvr *OpenTTY) applyOpSimple(op char, a float64, b float64) (float64) {
	if op == '+' {
		return a + b
	}
	if op == '-' {
		return a - b
	}
	if op == '*' {
		return a * b
	}
	if op == '/' {
		return <<unimp_expr[*grammar.JConditionalExpr]>>
	}
	return 0
}
func (rcvr *OpenTTY) audio(command string) (int) {
	command = env(command.trim())
	mainCommand := rcvr.getCommand(command)
	argument := rcvr.getArgument(command)
	if mainCommand.equals("") {
	} else if mainCommand.equals("volume") {
		if rcvr.player != nil {
			vc, ok := rcvr.player.getControl("VolumeControl").(*VolumeControl)
			if !ok {
				panic("XXX Cast fail for *parser.GoCastType")
			}
			if argument.equals("") {
				rcvr.echoCommand(fmt.Sprintf("%v%v", "", vc.getLevel()))
			} else {
				if try() {
					vc.setLevel(Integer.parseInt(argument))
				} else if catch_Exception(e) {
					echoCommand(e.getMessage())
					return 2
				}
			}
		} else {
			rcvr.echoCommand("audio: not running.")
			return 69
		}
	} else if mainCommand.equals("play") {
		if argument.equals("") {
		} else {
			if argument.startsWith("/mnt/") {
				argument = argument.substring(5)
			} else if argument.startsWith("/") {
				rcvr.echoCommand("audio: invalid source.")
				return 1
			} else {
				return rcvr.audio(fmt.Sprintf("%v%v%v", "play ", rcvr.path, argument))
			}
			if try() {
				CONN, ok := Connector.open(fmt.Sprintf("%v%v", "file:///", argument), Connector.READ).(*FileConnection)
				if !ok {
					panic("XXX Cast fail for *parser.GoCastType")
				}
				if !CONN.exists() {
					rcvr.echoCommand(fmt.Sprintf("%v%v%v", "audio: ", rcvr.basename(argument), ": not found"))
					return 127
				}
				IN := CONN.openInputStream()
				CONN.close()
				rcvr.player = Manager.createPlayer(IN, rcvr.getMimeType(argument))
				rcvr.player.prefetch()
				rcvr.player.start()
				rcvr.start("audio")
			} else if catch_Exception(e) {
				echoCommand(e.getMessage())
				return 1
			}
		}
	} else if mainCommand.equals("pause") {
		if try() {
			if rcvr.player != nil {
				rcvr.player.stop()
			} else {
				rcvr.echoCommand("audio: not running.")
				return 69
			}
		} else if catch_Exception(e) {
			echoCommand(e.getMessage())
			return 1
		}
	} else if mainCommand.equals("resume") {
		if try() {
			if rcvr.player != nil {
				rcvr.player.start()
			} else {
				rcvr.echoCommand("audio: not running.")
				return 69
			}
		} else if catch_Exception(e) {
			echoCommand(e.getMessage())
			return 1
		}
	} else if mainCommand.equals("stop") {
		if try() {
			if rcvr.player != nil {
				rcvr.player.stop()
				rcvr.player.close()
				rcvr.player = nil
				rcvr.stop("audio")
			} else {
				rcvr.echoCommand("audio: not running.")
				return 69
			}
		} else if catch_Exception(e) {
			echoCommand(e.getMessage())
			return 1
		}
	} else if mainCommand.equals("status") {
		echoCommand(<<unimp_expr[*grammar.JConditionalExpr]>>)
	} else {
		rcvr.echoCommand(fmt.Sprintf("%v%v%v", "audio: ", mainCommand, ": not found"))
		return 127
	}
	return 0
}
func (rcvr *OpenTTY) basename(path string) (string) {
	if path == nil || path.length() == 0 {
		return ""
	}
	if path.endsWith("/") {
		path = path.substring(0, path.length()-1)
	}
	lastSlashIndex := path.lastIndexOf('/')
	if lastSlashIndex == -1 {
		return path
	}
	return path.substring(lastSlashIndex + 1)
}
func (rcvr *OpenTTY) caseCommand(argument string) (int) {
	argument = argument.trim()
	firstParenthesis := argument.indexOf('(')
	lastParenthesis := argument.indexOf(')')
	if firstParenthesis == -1 || lastParenthesis == -1 || firstParenthesis > lastParenthesis {
		return 2
	}
	METHOD := rcvr.getCommand(argument)
	EXPR := argument.substring(firstParenthesis+1, lastParenthesis).trim()
	CMD := argument.substring(lastParenthesis + 1).trim()
	CONDITION := false
	NEGATED := METHOD.startsWith("!")
	if NEGATED {
		METHOD = METHOD.substring(1)
	}
	if METHOD.equals("file") {
		recordStores := RecordStore.listRecordStores()
		if recordStores != nil {
			for i := 0; i < len(recordStores); i++ {
				if recordStores[i].equals(EXPR) {
					CONDITION = true
					break
				}
			}
		}
	} else if METHOD.equals("root") {
		roots := FileSystemRegistry.listRoots()
		for roots.hasMoreElements() {
			if roots.nextElement().(string).equals(EXPR) {
				CONDITION = true
				break
			}
		}
	} else if METHOD.equals("thread") {
		CONDITION = replace(replace(Thread.currentThread().getName(), "MIDletEventQueue", "MIDlet"), "Thread-1", "MIDlet").equals(EXPR)
	} else if METHOD.equals("screen") {
		CONDITION = rcvr.desktops.containsKey(EXPR)
	} else if METHOD.equals("key") {
		CONDITION = rcvr.attributes.containsKey(EXPR)
	} else if METHOD.equals("alias") {
		CONDITION = rcvr.aliases.containsKey(EXPR)
	} else if METHOD.equals("trace") {
		CONDITION = rcvr.trace.containsKey(EXPR)
	}
	if CONDITION != NEGATED {
		return rcvr.processCommand(CMD)
	}
	return 0
}
func (rcvr *OpenTTY) CommandAction(c *Command, d *Displayable) {
	if c == rcvr.eXECUTE {
		command := rcvr.stdin.getString().trim()
		if !command.equals("") && !command.startsWith("!!") {
			rcvr.history.addElement(command.trim())
		}
		rcvr.stdin.setString("")
		rcvr.processCommand(command)
		rcvr.stdin.setLabel(fmt.Sprintf("%v%v%v%v", rcvr.username, " ", rcvr.path, " $"))
	} else if c == rcvr.hELP {
		rcvr.processCommand("help")
	} else if c == rcvr.nANO {
		NewNanoEditor("")
	} else if c == rcvr.cLEAR {
		rcvr.stdout.setText("")
	} else if c == rcvr.hISTORY {
		NewHistory()
	} else if c.getCommandType() == Command.BACK {
		rcvr.processCommand("xterm")
	} else if c.getCommandType() == Command.EXIT {
		rcvr.processCommand("exit")
	}
}
func (rcvr *OpenTTY) deleteFile(filename string) (int) {
	if filename == nil || filename.length() == 0 {
		return 2
	} else if filename.startsWith("/mnt/") {
		if try() {
			CONN, ok := Connector.open(fmt.Sprintf("%v%v", "file:///", filename.substring(5)), Connector.READ_WRITE).(*FileConnection)
			if !ok {
				panic("XXX Cast fail for *parser.GoCastType")
			}
			if CONN.exists() {
				CONN.delete()
			} else {
				rcvr.echoCommand(fmt.Sprintf("%v%v%v", "rm: ", rcvr.basename(filename), ": not found"))
				return 127
			}
			CONN.close()
		} else if catch_Exception(e) {
			echoCommand(e.getMessage())
		}
	} else if filename.startsWith("/home/") {
		if try() {
			RecordStore.deleteRecordStore(filename.substring(6))
		} else if catch_RecordStoreNotFoundException(e) {
			rcvr.echoCommand(fmt.Sprintf("%v%v%v", "rm: ", filename.substring(6), ": not found"))
			return 127
		} else if catch_RecordStoreNotFoundException(e) {
			rcvr.echoCommand(fmt.Sprintf("%v%v", "rm: ", e.getMessage()))
			return 1
		}
	} else if filename.startsWith("/") {
		rcvr.echoCommand("read-only storage")
		return 5
	} else {
		return rcvr.deleteFile(fmt.Sprintf("%v%v", rcvr.path, filename))
	}
	return 0
}
func (rcvr *OpenTTY) DestroyApp(unconditional bool) {
	rcvr.writeRMS("/home/nano", rcvr.nanoContent)
}
func (rcvr *OpenTTY) echoCommand(message string) {
	echoCommand(message, rcvr.stdout)
	rcvr.attributes.put("OUTPUT", message)
}
func (rcvr *OpenTTY) echoCommand2(message string, console *StringItem) {
	console.setText(<<unimp_expr[*grammar.JConditionalExpr]>>)
}
func (rcvr *OpenTTY) env(text string) (string) {
	text = rcvr.replace(text, "$PATH", rcvr.path)
	text = rcvr.replace(text, "$USERNAME", rcvr.username)
	text = replace(text, "$TITLE", rcvr.form.getTitle())
	text = replace(text, "$PROMPT", rcvr.stdin.getString())
	text = rcvr.replace(text, "\\n", "\n")
	text = rcvr.replace(text, "\\r", "\r")
	text = rcvr.replace(text, "\\t", "\t")
	e := rcvr.attributes.keys()
	for e.hasMoreElements() {
		key, ok := e.nextElement().(string)
		if !ok {
			panic("XXX Cast fail for *parser.GoCastType")
		}
		value, ok := rcvr.attributes.get(key).(string)
		if !ok {
			panic("XXX Cast fail for *parser.GoCastType")
		}
		text = rcvr.replace(text, fmt.Sprintf("%v%v", "$", key), value)
	}
	text = rcvr.replace(text, "$.", "$")
	text = rcvr.replace(text, "\\.", "\\")
	return text
}
func (rcvr *OpenTTY) exprCommand(expr string) (string) {
	tokens := expr.toCharArray()
	vals := make([]float64, 32)
	ops := make([]char, 32)
	valTop := -1
	opTop := -1
	i := 0
	len := len(tokens)
	for i < len {
		c := tokens[i]
		if c == ' ' {
			i++
			continue
		}
		if c >= '0' && c <= '9' {
			num := 0
			for i < len && tokens[i] >= '0' && tokens[i] <= '9' {
				num = num*10 + (tokens[++i] - '0')
			}
			if i < len && tokens[i] == '.' {
				i++
				frac := 0
				div := 10
				for i < len && tokens[i] >= '0' && tokens[i] <= '9' {
					frac += (tokens[++i] - '0') / div
					div *= 10
				}
				num += frac
			}
			vals[++valTop] = num
		} else if c == '(' {
			ops[++opTop] = c
			i++
		} else if c == ')' {
			for opTop >= 0 && ops[opTop] != '(' {
				b := vals[--valTop]
				a := vals[--valTop]
				op := ops[--opTop]
				vals[++valTop] = rcvr.applyOpSimple(op, a, b)
			}
			opTop--
			i++
		} else if c == '+' || c == '-' || c == '*' || c == '/' {
			for opTop >= 0 && rcvr.prec(ops[opTop]) >= prec(c) {
				b := vals[--valTop]
				a := vals[--valTop]
				op := ops[--opTop]
				vals[++valTop] = rcvr.applyOpSimple(op, a, b)
			}
			ops[++opTop] = c
			i++
		} else {
			return fmt.Sprintf("%v%v%v", "expr: invalid char '", c, "'")
		}
	}
	for opTop >= 0 {
		b := vals[--valTop]
		a := vals[--valTop]
		op := ops[--opTop]
		vals[++valTop] = rcvr.applyOpSimple(op, a, b)
	}
	result := vals[valTop]
	return <<unimp_expr[*grammar.JConditionalExpr]>>
}
func (rcvr *OpenTTY) extractClassName(code string) (string) {
	idx := code.indexOf("class ")
	if idx == -1 {
		return "Unnamed"
	}
	idx += 6
	end := code.indexOf(' ', idx)
	if end == -1 {
		end = code.indexOf('{', idx)
	}
	return code.substring(idx, end).trim()
}
func (rcvr *OpenTTY) extractImports(code string) (<<array>>) {
	imports := NewVector()
	start := 0
	for start < code.length() {
		end := code.indexOf('\n', start)
		if end == -1 {
			end = code.length()
		}
		line := code.substring(start, end).trim()
		if line.startsWith("import ") {
			semi := line.indexOf(';')
			if semi != -1 {
				imp := line.substring(7, semi).trim()
				imports.addElement(rcvr.replace(imp, ".", "/"))
			}
		}
		start = end + 1
	}
	result := make([]string, len(imports))
	for i := 0; i < len(result); i++ {
		result[i] = imports.elementAt(i).(string)
	}
	return result
}
func (rcvr *OpenTTY) extractMnemonics(code string) (string) {
	idx := code.indexOf("main")
	if idx == -1 {
		return ""
	}
	braceStart := code.indexOf('{', idx)
	if braceStart == -1 {
		return ""
	}
	braceCount := 1
	i := braceStart + 1
	for i < code.length() && braceCount > 0 {
		c := code.charAt(i)
		if c == '{' {
			braceCount++
		} else if c == '}' {
			braceCount--
		}
		i++
	}
	if braceCount != 0 {
		return ""
	}
	return code.substring(braceStart+1, i-1).trim()
}
func (rcvr *OpenTTY) extractTag(htmlContent string, tag string, fallback string) (string) {
	startTag := fmt.Sprintf("%v%v%v", "<", tag, ">")
	endTag := fmt.Sprintf("%v%v%v", "</", tag, ">")
	start := htmlContent.indexOf(startTag)
	end := htmlContent.indexOf(endTag)
	if start != -1 && end != -1 && end > start {
		return htmlContent.substring(start+startTag.length(), end).trim()
	} else {
		return fallback
	}
}
func (rcvr *OpenTTY) extractTitle(htmlContent string) (string) {
	return extractTag(htmlContent, "title", "HTML Viewer")
}
func (rcvr *OpenTTY) forCommand(argument string) (int) {
	argument = argument.trim()
	firstParenthesis := argument.indexOf('(')
	lastParenthesis := argument.indexOf(')')
	if firstParenthesis == -1 || lastParenthesis == -1 || firstParenthesis > lastParenthesis {
		return 2
	}
	KEY := rcvr.getCommand(argument)
	FILE := getcontent(argument.substring(firstParenthesis+1, lastParenthesis).trim())
	CMD := argument.substring(lastParenthesis + 1).trim()
	if KEY.startsWith("(") {
		return 2
	}
	if KEY.startsWith("$") {
		KEY = rcvr.replace(KEY, "$", "")
	}
	LINES := rcvr.split(FILE, '\n')
	for i := 0; i < len(LINES); i++ {
		if LINES[i] != nil || LINES[i].length() != 0 {
			rcvr.processCommand2(fmt.Sprintf("%v%v%v%v", "set ", KEY, "=", LINES[i]), false)
			STATUS := rcvr.processCommand(CMD)
			rcvr.processCommand2(fmt.Sprintf("%v%v", "unset ", KEY), false)
			if STATUS != 0 {
				return STATUS
			}
		}
	}
	return 0
}
func (rcvr *OpenTTY) generateClass(code string) (<<array>>) {
	className := rcvr.extractClassName(code)
	imports := rcvr.extractImports(code)
	mnemonics := rcvr.extractMnemonics(code)
	out := NewByteArrayOutputStream()
	if try() {
		nameBytes := className.getBytes()
		nameLen := len(nameBytes)
		codeBytes := rcvr.mnemonicsToBytes(mnemonics)
		codeLen := len(codeBytes)
		codeAttrLen := 12 + codeLen
		constantPoolSize := 11 + len(imports)*2
		cpCount := constantPoolSize
		out.write(0xCA)
		out.write(0xFE)
		out.write(0xBA)
		out.write(0xBE)
		out.write(0x00)
		out.write(0x00)
		out.write(0x00)
		out.write(0x2E)
		out.write(0x00)
		out.write(cpCount)
		out.write(0x07)
		out.write(0x00)
		out.write(0x02)
		out.write(0x01)
		out.write(uint32(nameLen) >> 8 & 0xFF)
		out.write(nameLen & 0xFF)
		for i := 0; i < nameLen; i++ {
			out.write(nameBytes[i])
		}
		out.write(0x07)
		out.write(0x00)
		out.write(0x04)
		obj := "java/lang/Object".getBytes()
		out.write(0x01)
		out.write(0x00)
		out.write(len(obj))
		for i := 0; i < len(obj); i++ {
			out.write(obj[i])
		}
		init := "<init>".getBytes()
		out.write(0x01)
		out.write(0x00)
		out.write(len(init))
		for i := 0; i < len(init); i++ {
			out.write(init[i])
		}
		desc := "()V".getBytes()
		out.write(0x01)
		out.write(0x00)
		out.write(len(desc))
		for i := 0; i < len(desc); i++ {
			out.write(desc[i])
		}
		codeStr := "Code".getBytes()
		out.write(0x01)
		out.write(0x00)
		out.write(len(codeStr))
		for i := 0; i < len(codeStr); i++ {
			out.write(codeStr[i])
		}
		out.write(0x0A)
		out.write(0x00)
		out.write(0x03)
		out.write(0x00)
		out.write(0x09)
		out.write(0x0C)
		out.write(0x00)
		out.write(0x05)
		out.write(0x00)
		out.write(0x06)
		main := "main".getBytes()
		out.write(0x01)
		out.write(0x00)
		out.write(len(main))
		for i := 0; i < len(main); i++ {
			out.write(main[i])
		}
		for i := 0; i < len(imports); i++ {
			b := imports[i].getBytes()
			out.write(0x01)
			out.write(uint32(len(b)) >> 8 & 0xFF)
			out.write(len(b) & 0xFF)
			for j := 0; j < len(b); j++ {
				out.write(b[j])
			}
		}
		out.write(0x00)
		out.write(0x21)
		out.write(0x00)
		out.write(0x01)
		out.write(0x00)
		out.write(0x03)
		out.write(0x00)
		out.write(0x00)
		out.write(0x00)
		out.write(0x00)
		out.write(0x00)
		out.write(0x02)
		out.write(0x00)
		out.write(0x01)
		out.write(0x00)
		out.write(0x05)
		out.write(0x00)
		out.write(0x06)
		out.write(0x00)
		out.write(0x01)
		out.write(0x00)
		out.write(0x07)
		out.write(0x00)
		out.write(0x00)
		out.write(0x00)
		out.write(0x11)
		out.write(0x00)
		out.write(0x01)
		out.write(0x00)
		out.write(0x01)
		out.write(0x00)
		out.write(0x00)
		out.write(0x00)
		out.write(0x05)
		out.write(0x2A)
		out.write(0xB7)
		out.write(0x00)
		out.write(0x08)
		out.write(0xB1)
		out.write(0x00)
		out.write(0x00)
		out.write(0x00)
		out.write(0x00)
		out.write(0x00)
		out.write(0x09)
		out.write(0x00)
		out.write(0x0A)
		out.write(0x00)
		out.write(0x06)
		out.write(0x00)
		out.write(0x01)
		out.write(0x00)
		out.write(0x07)
		out.write(uint32(codeAttrLen) >> 24 & 0xFF)
		out.write(uint32(codeAttrLen) >> 16 & 0xFF)
		out.write(uint32(codeAttrLen) >> 8 & 0xFF)
		out.write(codeAttrLen & 0xFF)
		out.write(0x00)
		out.write(0x02)
		out.write(0x00)
		out.write(0x01)
		out.write(uint32(codeLen) >> 24 & 0xFF)
		out.write(uint32(codeLen) >> 16 & 0xFF)
		out.write(uint32(codeLen) >> 8 & 0xFF)
		out.write(codeLen & 0xFF)
		for i := 0; i < codeLen; i++ {
			out.write(codeBytes[i])
		}
		out.write(0x00)
		out.write(0x00)
		out.write(0x00)
		out.write(0x00)
		out.write(0x00)
		out.write(0x00)
	} else if catch_Exception(e) {
		rcvr.echoCommand(fmt.Sprintf("%v%v", "Build failed: ", e.getMessage()))
		return nil
	}
	return out.toByteArray()
}
func (rcvr *OpenTTY) generateUUID() (string) {
	chars := "0123456789abcdef"
	uuid := NewStringBuffer()
	for i := 0; i < 36; i++ {
		if i == 8 || i == 13 || i == 18 || i == 23 {
			uuid.append('-')
		} else if i == 14 {
			uuid.append('4')
		} else if i == 19 {
			uuid.append(chars.charAt(8 + rcvr.random.nextInt(4)))
		} else {
			uuid.append(chars.charAt(rcvr.random.nextInt(16)))
		}
	}
	return uuid.toString()
}
func (rcvr *OpenTTY) getArgument(input string) (string) {
	spaceIndex := input.indexOf(' ')
	if spaceIndex == -1 {
		return ""
	} else {
		return rcvr.getpattern(input.substring(spaceIndex + 1).trim())
	}
}
func (rcvr *OpenTTY) getCommand(input string) (string) {
	spaceIndex := input.indexOf(' ')
	if spaceIndex == -1 {
		return input
	} else {
		return input.substring(0, spaceIndex)
	}
}
func (rcvr *OpenTTY) getMimeType(filename string) (string) {
	filename = filename.toLowerCase()
	if filename.endsWith(".amr") {
		return "audio/amr"
	} else if filename.endsWith(".wav") {
		return "audio/x-wav"
	} else {
		return "audio/mpeg"
	}
}
func (rcvr *OpenTTY) getNumber(s string) (*Double) {
	if try() {
		return Integer.valueOf(s)
	} else if catch_NumberFormatException(e) {
		return nil
	}
}
func (rcvr *OpenTTY) getcontent(file string) (string) {
	return <<unimp_expr[*grammar.JConditionalExpr]>>
}
func (rcvr *OpenTTY) getpattern(text string) (string) {
	return <<unimp_expr[*grammar.JConditionalExpr]>>
}
func (rcvr *OpenTTY) html2text(htmlContent string) (string) {
	text := NewStringBuffer()
	inTag := false
	inStyle := false
	inScript := false
	inTitle := false
	for i := 0; i < htmlContent.length(); i++ {
		c := htmlContent.charAt(i)
		if c == '<' {
			inTag = true
			if htmlContent.regionMatches(true, i, "<title>", 0, 7) {
				inTitle = true
			} else if htmlContent.regionMatches(true, i, "<style>", 0, 7) {
				inStyle = true
			} else if htmlContent.regionMatches(true, i, "<script>", 0, 8) {
				inScript = true
			} else if htmlContent.regionMatches(true, i, "</title>", 0, 8) {
				inTitle = false
			} else if htmlContent.regionMatches(true, i, "</style>", 0, 8) {
				inStyle = false
			} else if htmlContent.regionMatches(true, i, "</script>", 0, 9) {
				inScript = false
			}
		} else if c == '>' {
			inTag = false
		} else if !inTag && !inStyle && !inScript && !inTitle {
			text.append(c)
		}
	}
	return text.toString().trim()
}
func (rcvr *OpenTTY) ifCommand(argument string) (int) {
	argument = argument.trim()
	firstParenthesis := argument.indexOf('(')
	lastParenthesis := argument.indexOf(')')
	if firstParenthesis == -1 || lastParenthesis == -1 || firstParenthesis > lastParenthesis {
		rcvr.echoCommand("if (expr) [command]")
		return 2
	}
	EXPR := argument.substring(firstParenthesis+1, lastParenthesis).trim()
	CMD := argument.substring(lastParenthesis + 1).trim()
	PARTS := rcvr.split(EXPR, ' ')
	if len(PARTS) == 3 {
		CONDITION := false
		NEGATED := PARTS[1].startsWith("!") && !PARTS[1].equals("!=")
		if NEGATED {
			PARTS[1] = PARTS[1].substring(1)
		}
		N1 := getNumber(PARTS[0])
		N2 := getNumber(PARTS[2])
		if N1 != nil && N2 != nil {
			if PARTS[1].equals("==") {
				CONDITION = N1.doubleValue() == N2.doubleValue()
			} else if PARTS[1].equals("!=") {
				CONDITION = N1.doubleValue() != N2.doubleValue()
			} else if PARTS[1].equals(">") {
				CONDITION = N1.doubleValue() > N2.doubleValue()
			} else if PARTS[1].equals("<") {
				CONDITION = N1.doubleValue() < N2.doubleValue()
			} else if PARTS[1].equals(">=") {
				CONDITION = N1.doubleValue() >= N2.doubleValue()
			} else if PARTS[1].equals("<=") {
				CONDITION = N1.doubleValue() <= N2.doubleValue()
			}
		} else {
			if PARTS[1].equals("startswith") {
				CONDITION = PARTS[0].startsWith(PARTS[2])
			} else if PARTS[1].equals("endswith") {
				CONDITION = PARTS[0].endsWith(PARTS[2])
			} else if PARTS[1].equals("contains") {
				CONDITION = PARTS[0].indexOf(PARTS[2]) != -1
			} else if PARTS[1].equals("==") {
				CONDITION = PARTS[0].equals(PARTS[2])
			} else if PARTS[1].equals("!=") {
				CONDITION = !PARTS[0].equals(PARTS[2])
			}
		}
		if CONDITION != NEGATED {
			return rcvr.processCommand(CMD)
		}
	} else if len(PARTS) == 2 {
		if PARTS[0].equals(PARTS[1]) {
			return rcvr.processCommand(CMD)
		}
	} else if len(PARTS) == 1 {
		if !PARTS[0].equals("") {
			return rcvr.processCommand(CMD)
		}
	}
	return 0
}
func (rcvr *OpenTTY) importScript(script string) (int) {
	if script == nil || script.length() == 0 {
		return 2
	}
	PKG := rcvr.parseProperties(rcvr.getcontent(script))
	if PKG.containsKey("api.version") {
		if !rcvr.env("$VERSION").startsWith(PKG.get("api.version").(string)) {
			processCommand(<<unimp_expr[*grammar.JConditionalExpr]>>)
			return 3
		}
	}
	if PKG.containsKey("process.name") {
		rcvr.start(PKG.get("process.name").(string))
	}
	if PKG.containsKey("process.type") {
		TYPE, ok := PKG.get("process.type").(string)
		if !ok {
			panic("XXX Cast fail for *parser.GoCastType")
		}
		if TYPE.equals("server") {
		} else if TYPE.equals("bind") {
			NewBind(rcvr.env(fmt.Sprintf("%v%v%v", PKG.get("process.port").(string), " ", PKG.get("process.db").(string))))
		} else {
			rcvr.mIDletLogs(fmt.Sprintf("%v%v%v", "add warn '", TYPE.toUpperCase(), "' is a invalid value for 'process.type'"))
		}
	}
	if PKG.containsKey("process.host") && PKG.containsKey("process.port") {
		NewServer(rcvr.env(fmt.Sprintf("%v%v%v", PKG.get("process.port").(string), " ", PKG.get("process.host").(string))))
	}
	if PKG.containsKey("include") {
		include := rcvr.split(PKG.get("include").(string), ',')
		for i := 0; i < len(include); i++ {
			STATUS := importScript(include[i])
			if STATUS != 0 {
				return STATUS
			}
		}
	}
	if PKG.containsKey("config") {
		rcvr.processCommand(PKG.get("config").(string))
	}
	if PKG.containsKey("mod") && PKG.containsKey("process.name") {
		PROCESS, ok := PKG.get("process.name").(string)
		if !ok {
			panic("XXX Cast fail for *parser.GoCastType")
		}
		MOD, ok := PKG.get("mod").(string)
		if !ok {
			panic("XXX Cast fail for *parser.GoCastType")
		}
		NewAnonymous_Thread_0("MIDlet-Mod").start()
	}
	if PKG.containsKey("command") {
		commands := rcvr.split(PKG.get("command").(string), ',')
		for i := 0; i < len(commands); i++ {
			if PKG.containsKey(commands[i]) {
				rcvr.aliases.put(commands[i], rcvr.env(PKG.get(commands[i]).(string)))
			} else {
				rcvr.mIDletLogs(fmt.Sprintf("%v%v%v", "add error Failed to create command '", commands[i], "' content not found"))
			}
		}
	}
	if PKG.containsKey("file") {
		files := rcvr.split(PKG.get("file").(string), ',')
		for i := 0; i < len(files); i++ {
			if PKG.containsKey(files[i]) {
				STATUS := rcvr.writeRMS2(fmt.Sprintf("%v%v", "/home/", files[i]), rcvr.env(PKG.get(files[i]).(string)))
			} else {
				rcvr.mIDletLogs(fmt.Sprintf("%v%v%v", "add error Failed to create file '", files[i], "' content not found"))
			}
		}
	}
	if PKG.containsKey("shell.name") && PKG.containsKey("shell.args") {
		args := rcvr.split(PKG.get("shell.args").(string), ',')
		TABLE := NewHashtable()
		for i := 0; i < len(args); i++ {
			NAME := args[i].trim()
			VALUE, ok := PKG.get(NAME).(string)
			if !ok {
				panic("XXX Cast fail for *parser.GoCastType")
			}
			TABLE.put(NAME, <<unimp_expr[*grammar.JConditionalExpr]>>)
		}
		if PKG.containsKey("shell.unknown") {
			TABLE.put("shell.unknown", PKG.get("shell.unknown").(string))
		}
		rcvr.shell.put(PKG.get("shell.name").(string).trim(), TABLE)
	}
	return 0
}
func (rcvr *OpenTTY) java(command string) (int) {
	command = env(command.trim())
	mainCommand := rcvr.getCommand(command)
	argument := rcvr.getArgument(command)
	if mainCommand.equals("") {
		rcvr.viewer("Java ME", rcvr.env("Java 1.2 (OpenTTY Edition)\n\nMicroEdition-Config: $CONFIG\nMicroEdition-Profile: $PROFILE"))
	} else if mainCommand.equals("-class") {
		if argument.equals("") {
		} else {
			STATUS := rcvr.javaClass(argument)
			echoCommand(<<unimp_expr[*grammar.JConditionalExpr]>>)
			return STATUS
		}
	} else if mainCommand.equals("--version") {
		var s string
		BUFFER := NewStringBuffer()
		if (s = System.getProperty("java.vm.name")) != nil {
			BUFFER.append(s).append(", ").append(System.getProperty("java.vm.vendor"))
			if (s = System.getProperty("java.vm.version")) != nil {
				BUFFER.append('\n').append(s)
			}
			if (s = System.getProperty("java.vm.specification.name")) != nil {
				BUFFER.append('\n').append(s)
			}
		} else if (s = System.getProperty("com.ibm.oti.configuration")) != nil {
			BUFFER.append("J9 VM, IBM (").append(s).append(')')
			if (s = System.getProperty("java.fullversion")) != nil {
				BUFFER.append("\n\n").append(s)
			}
		} else if (s = System.getProperty("com.oracle.jwc.version")) != nil {
			BUFFER.append("OJWC v").append(s).append(", Oracle")
		} else if javaClass([]string{"com.sun.cldchi.io.ConsoleOutputStream", "com.sun.cldchi.jvm.JVM"}) {
			BUFFER.append("CLDC Hotspot Implementation, Sun")
		} else if javaClass([]string{"com.sun.midp.io.InternalConnector", "com.sun.midp.io.ConnectionBaseAdapter", "com.sun.midp.Main"}) {
			BUFFER.append("KVM, Sun (MIDP)")
		} else if javaClass([]string{"com.sun.cldc.util.j2me.CalendarImpl", "com.sun.cldc.i18n.Helper", "com.sun.cldc.io.ConsoleOutputStream", "com.sun.cldc.i18n.uclc.DefaultCaseConverter"}) {
			BUFFER.append("KVM, Sun (CLDC)")
		} else if javaClass([]string{"com.jblend.util.SortedVector", "com.jblend.tck.socket2http.Protocol", "com.jblend.io.j2me.resource.Protocol", "com.jblend.security.midp20.SecurityManagerImpl", "com.jblend.security.midp20.UserConfirmDialogImpl", "jp.co.aplix.cldc.io.MIDPURLChecker", "jp.co.aplix.cldc.io.j2me.http.HttpConnectionImpl"}) {
			BUFFER.append("JBlend, Aplix")
		} else if javaClass([]string{"com.jbed.io.CharConvUTF8", "com.jbed.runtime.MemSupport", "com.jbed.midp.lcdui.GameCanvasPeer", "com.jbed.microedition.media.CoreManager", "com.jbed.runtime.Mem", "com.jbed.midp.lcdui.GameCanvas", "com.jbed.microedition.media.Core"}) {
			BUFFER.append("Jbed, Esmertec/Myriad Group")
		} else if javaClass([]string{"MahoTrans.IJavaObject"}) {
			BUFFER.append("MahoTrans")
		} else {
			BUFFER.append("Unknown")
		}
		echoCommand(BUFFER.append('\n').toString())
	} else {
	}
	return 0
}
func (rcvr *OpenTTY) javaClass(argument string) (int) {
	if try() {
		Class.forName(argument)
		return 0
	} else if catch_ClassNotFoundException(e) {
		return 3
	}
}
func (rcvr *OpenTTY) javaClass2(classes []string) (bool) {
	for i := 0; i < len(classes); i++ {
		if javaClass(classes[i]) != 0 {
			return false
		}
	}
	return true
}
func (rcvr *OpenTTY) kill(pid string) (int) {
	if pid == nil || pid.length() == 0 {
		return 2
	}
	KEYS := rcvr.trace.keys()
	for KEYS.hasMoreElements() {
		KEY, ok := KEYS.nextElement().(string)
		if !ok {
			panic("XXX Cast fail for *parser.GoCastType")
		}
		if pid.equals(rcvr.trace.get(KEY)) {
			rcvr.trace.remove(KEY)
			rcvr.echoCommand(fmt.Sprintf("%v%v%v", "Process with PID ", pid, " terminated"))
			if KEY.equals("sh") {
				rcvr.processCommand("exit")
			}
			return 0
		}
	}
	rcvr.echoCommand(fmt.Sprintf("%v%v%v", "PID '", pid, "' not found"))
	return 127
}
func (rcvr *OpenTTY) loadRMS(filename string) (string) {
	return rcvr.read(fmt.Sprintf("%v%v", "/home/", filename))
}
func (rcvr *OpenTTY) mnemonicsToBytes(mnemonics string) (<<array>>) {
	out := NewByteArrayOutputStream()
	opcodes := NewHashtable()
	opcodes.put("nop", NewInteger(0x00))
	opcodes.put("aconst_null", NewInteger(0x01))
	opcodes.put("iconst_0", NewInteger(0x03))
	opcodes.put("iconst_1", NewInteger(0x04))
	opcodes.put("iconst_2", NewInteger(0x05))
	opcodes.put("iload_0", NewInteger(0x1A))
	opcodes.put("aload_0", NewInteger(0x2A))
	opcodes.put("istore_0", NewInteger(0x3B))
	opcodes.put("astore_0", NewInteger(0x4B))
	opcodes.put("pop", NewInteger(0x57))
	opcodes.put("iadd", NewInteger(0x60))
	opcodes.put("return", NewInteger(0xB1))
	opcodes.put("invokespecial", NewInteger(0xB7))
	start := 0
	length := mnemonics.length()
	for start < length {
		end := mnemonics.indexOf('\n', start)
		if end == -1 {
			end = length
		}
		line := mnemonics.substring(start, end).trim()
		start = end + 1
		if line.length() == 0 {
			continue
		}
		spaceIndex := line.indexOf(' ')
		instr := line
		arg := nil
		if spaceIndex != -1 {
			instr = line.substring(0, spaceIndex)
			arg = line.substring(spaceIndex + 1).trim()
		}
		opcodeInt, ok := opcodes.get(instr).(*Integer)
		if !ok {
			panic("XXX Cast fail for *parser.GoCastType")
		}
		if opcodeInt == nil {
			throw(NewException(fmt.Sprintf("%v%v", "Opcode desconhecido: ", instr)))
		}
		out.write(opcodeInt.intValue())
		if arg != nil {
			args := rcvr.split(arg, ' ')
			for i := 0; i < len(args); i++ {
				val := Integer.parseInt(args[i])
				out.write(val & 0xFF)
			}
		}
	}
	return out.toByteArray()
}
func (rcvr *OpenTTY) mount(script string) {
	lines := rcvr.split(script, '\n')
	for i := 0; i < len(lines); i++ {
		line := ""
		if lines[i] != nil {
			line = lines[i].trim()
		}
		if line.length() == 0 || line.startsWith("#") {
			continue
		}
		if line.startsWith("/") {
			fullPath := ""
			start := 0
			for j := 1; j < line.length(); j++ {
				if line.charAt(j) == '/' {
					dir := line.substring(start+1, j)
					fullPath += fmt.Sprintf("%v%v", "/", dir)
					rcvr.addDirectory(fmt.Sprintf("%v%v", fullPath, "/"))
					start = j
				}
			}
			finalPart := line.substring(start + 1)
			fullPath += fmt.Sprintf("%v%v", "/", finalPart)
			if line.endsWith("/") {
				rcvr.addDirectory(fmt.Sprintf("%v%v", fullPath, "/"))
			} else {
				rcvr.addDirectory(fullPath)
			}
		}
	}
}
func (rcvr *OpenTTY) newFont(argument string) (*Font) {
	if argument == nil || argument.length() == 0 || argument.equals("default") {
		return Font.getDefaultFont()
	}
	style := Font.STYLE_PLAIN
	size := Font.SIZE_MEDIUM
	if argument.equals("bold") {
		style = Font.STYLE_BOLD
	} else if argument.equals("italic") {
		style = Font.STYLE_ITALIC
	} else if argument.equals("ul") {
		style = Font.STYLE_UNDERLINED
	} else if argument.equals("small") {
		size = Font.SIZE_SMALL
	} else if argument.equals("large") {
		size = Font.SIZE_LARGE
	} else {
		return rcvr.newFont("default")
	}
	return Font.getFont(Font.FACE_SYSTEM, style, size)
}
func (rcvr *OpenTTY) parseConf(text string) (string) {
	iniBuffer := NewStringBuffer()
	text = text.trim()
	if text.startsWith("{") && text.endsWith("}") {
		text = text.substring(1, text.length()-1)
	}
	pairs := rcvr.split(text, ',')
	for i := 0; i < len(pairs); i++ {
		pair := pairs[i].trim()
		keyValue := rcvr.split(pair, ':')
		if len(keyValue) == 2 {
			key := getpattern(keyValue[0].trim())
			value := getpattern(keyValue[1].trim())
			iniBuffer.append(key).append("=").append(value).append("\n")
		}
	}
	return iniBuffer.toString()
}
func (rcvr *OpenTTY) parseJson(text string) (string) {
	properties := rcvr.parseProperties(text)
	if properties.isEmpty() {
		return "{}"
	}
	keys := properties.keys()
	jsonBuffer := NewStringBuffer()
	jsonBuffer.append("{")
	for keys.hasMoreElements() {
		key, ok := keys.nextElement().(string)
		if !ok {
			panic("XXX Cast fail for *parser.GoCastType")
		}
		value, ok := properties.get(key).(string)
		if !ok {
			panic("XXX Cast fail for *parser.GoCastType")
		}
		jsonBuffer.append("\n  \"").append(key).append("\": ")
		jsonBuffer.append("\"").append(value).append("\"")
		if keys.hasMoreElements() {
			jsonBuffer.append(",")
		}
	}
	jsonBuffer.append("\n}")
	return jsonBuffer.toString()
}
func (rcvr *OpenTTY) parseProperties(text string) (*Hashtable) {
	properties := NewHashtable()
	lines := rcvr.split(text, '\n')
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if !line.startsWith("#") {
			equalIndex := line.indexOf('=')
			if equalIndex > 0 && equalIndex < line.length()-1 {
				key := line.substring(0, equalIndex).trim()
				value := line.substring(equalIndex + 1).trim()
				properties.put(key, value)
			}
		}
	}
	return properties
}
func (rcvr *OpenTTY) PauseApp() {
}
func (rcvr *OpenTTY) prec(op char) (int) {
	if op == '+' || op == '-' {
		return 1
	}
	if op == '*' || op == '/' {
		return 2
	}
	return 0
}
func (rcvr *OpenTTY) processCommand(command string) (int) {
	return processCommand(command, true)
}
func (rcvr *OpenTTY) processCommand2(command string, ignore bool) (int) {
	command = <<unimp_expr[*grammar.JConditionalExpr]>>
	mainCommand := rcvr.getCommand(command)
	argument := rcvr.getArgument(command)
	if rcvr.shell.containsKey(mainCommand) && ignore {
		args, ok := rcvr.shell.get(mainCommand).(*Hashtable)
		if !ok {
			panic("XXX Cast fail for *parser.GoCastType")
		}
		if argument.equals("") {
			return processCommand(<<unimp_expr[*grammar.JConditionalExpr]>>)
		} else if args.containsKey(rcvr.getCommand(argument).toLowerCase()) {
			return rcvr.processCommand(fmt.Sprintf("%v%v%v", args.get(rcvr.getCommand(argument)).(string), " ", rcvr.getArgument(argument)))
		} else {
			return processCommand(<<unimp_expr[*grammar.JConditionalExpr]>>, <<unimp_expr[*grammar.JConditionalExpr]>>)
		}
	}
	if rcvr.aliases.containsKey(mainCommand) && ignore {
		return rcvr.processCommand(fmt.Sprintf("%v%v%v", rcvr.aliases.get(mainCommand).(string), " ", argument))
	}
	if rcvr.functions.containsKey(mainCommand) && ignore {
		return runScript(rcvr.functions.get(mainCommand).(string))
	}
	if mainCommand.equals("") {
	} else if mainCommand.equals("alias") {
		if argument.equals("") {
			KEYS := rcvr.aliases.keys()
			for KEYS.hasMoreElements() {
				KEY, ok := KEYS.nextElement().(string)
				if !ok {
					panic("XXX Cast fail for *parser.GoCastType")
				}
				VALUE, ok := rcvr.aliases.get(KEY).(string)
				if !ok {
					panic("XXX Cast fail for *parser.GoCastType")
				}
				if !KEY.equals("xterm") && !VALUE.equals("") {
					rcvr.echoCommand(fmt.Sprintf("%v%v%v%v%v", "alias ", KEY, "='", VALUE.trim(), "'"))
				}
			}
		} else {
			INDEX := argument.indexOf('=')
			if INDEX == -1 {
				if rcvr.aliases.containsKey(argument) {
					rcvr.echoCommand(fmt.Sprintf("%v%v%v%v%v", "alias ", argument, "='", rcvr.aliases.get(argument).(string), "'"))
				} else {
					rcvr.echoCommand(fmt.Sprintf("%v%v%v", "alias: ", argument, ": not found"))
					return 127
				}
			} else {
				rcvr.aliases.put(argument.substring(0, INDEX).trim(), argument.substring(INDEX+1).trim())
			}
		}
	} else if mainCommand.equals("unalias") {
		if argument.equals("") {
		} else if rcvr.aliases.containsKey(argument) {
			rcvr.aliases.remove(argument)
		} else {
			rcvr.echoCommand(fmt.Sprintf("%v%v%v", "unalias: ", argument, ": not found"))
			return 127
		}
	} else if mainCommand.equals("set") {
		if argument.equals("") {
		} else {
			INDEX := argument.indexOf('=')
			if INDEX == -1 {
				rcvr.attributes.put(argument, "")
			} else {
				rcvr.attributes.put(argument.substring(0, INDEX).trim(), argument.substring(INDEX+1).trim())
			}
		}
	} else if mainCommand.equals("unset") {
		if argument.equals("") {
		} else if rcvr.attributes.containsKey(argument) {
			rcvr.attributes.remove(argument)
		} else {
		}
	} else if mainCommand.equals("export") {
		return processCommand(<<unimp_expr[*grammar.JConditionalExpr]>>, false)
	} else if mainCommand.equals("env") {
		if argument.equals("") {
			KEYS := rcvr.attributes.keys()
			for KEYS.hasMoreElements() {
				KEY, ok := KEYS.nextElement().(string)
				if !ok {
					panic("XXX Cast fail for *parser.GoCastType")
				}
				VALUE, ok := rcvr.aliases.get(KEY).(string)
				if !ok {
					panic("XXX Cast fail for *parser.GoCastType")
				}
				if !KEY.equals("OUTPUT") && !VALUE.equals("") {
					rcvr.echoCommand(fmt.Sprintf("%v%v%v", KEY, "=", VALUE.trim()))
				}
			}
		} else if rcvr.attributes.containsKey(argument) {
			rcvr.echoCommand(fmt.Sprintf("%v%v%v", argument, "=", rcvr.attributes.get(argument).(string)))
		} else {
			rcvr.echoCommand(fmt.Sprintf("%v%v%v", "env: ", argument, ": not found"))
			return 127
		}
	} else if mainCommand.equals("log") {
		return rcvr.mIDletLogs(argument)
	} else if mainCommand.equals("logcat") {
		rcvr.echoCommand(rcvr.logs)
	} else if mainCommand.equals("logout") {
		rcvr.writeRMS("/home/OpenRMS", "")
		rcvr.processCommand("exit")
	} else if mainCommand.equals("whoami") || mainCommand.equals("logname") {
		rcvr.echoCommand(rcvr.username)
	} else if mainCommand.equals("sh") || mainCommand.equals("login") {
		return processCommand("import /java/bin/sh", false)
	} else if mainCommand.equals("x11") {
		return rcvr.xserver(argument)
	} else if mainCommand.equals("xterm") {
		rcvr.display.setCurrent(rcvr.form)
	} else if mainCommand.equals("gauge") {
		return rcvr.xserver(fmt.Sprintf("%v%v", "gauge ", argument))
	} else if mainCommand.equals("warn") {
		return rcvr.warnCommand(rcvr.form.getTitle(), argument)
	} else if mainCommand.equals("title") {
		rcvr.form.setTitle(<<unimp_expr[*grammar.JConditionalExpr]>>)
	} else if mainCommand.equals("tick") {
		if argument.equals("label") {
			echoCommand(rcvr.display.getCurrent().getTicker().getString())
		} else {
			return rcvr.xserver(fmt.Sprintf("%v%v", "tick ", argument))
		}
	} else if mainCommand.equals("for") {
		return rcvr.forCommand(argument)
	} else if mainCommand.equals("if") {
		return rcvr.ifCommand(argument)
	} else if mainCommand.equals("case") {
		return rcvr.caseCommand(argument)
	} else if mainCommand.equals("builtin") || mainCommand.equals("command") {
		return processCommand(argument, false)
	} else if mainCommand.equals("bruteforce") {
		rcvr.start("bruteforce")
		for rcvr.trace.containsKey("bruteforce") {
			STATUS := processCommand(argument, ignore)
			if STATUS != 0 {
				rcvr.stop("bruteforce")
				return STATUS
			}
		}
	} else if mainCommand.equals("cron") {
		if argument.equals("") {
		} else {
			return rcvr.processCommand(fmt.Sprintf("%v%v%v%v", "execute sleep ", rcvr.getCommand(argument), "; ", rcvr.getArgument(argument)))
		}
	} else if mainCommand.equals("sleep") {
		if argument.equals("") {
		} else {
			if try() {
				Thread.sleep(Integer.parseInt(argument) * 1000)
			} else if catch_InterruptedException(e) {
			} else if catch_InterruptedException(e) {
				echoCommand(e.getMessage())
				return 2
			}
		}
	} else if mainCommand.equals("time") {
		if argument.equals("") {
		} else {
			START := System.currentTimeMillis()
			STATUS := processCommand(argument, ignore)
			rcvr.echoCommand(fmt.Sprintf("%v%v%v", "at ", System.currentTimeMillis()-START, "ms"))
			return STATUS
		}
	} else if mainCommand.equals("exec") {
		CMDS := rcvr.split(argument, '&')
		for i := 0; i < len(CMDS); i++ {
			processCommand(CMDS[i].trim(), ignore)
		}
	} else if mainCommand.equals("execute") {
		CMDS := rcvr.split(argument, ';')
		for i := 0; i < len(CMDS); i++ {
			processCommand(CMDS[i].trim(), ignore)
		}
	} else if mainCommand.equals("gc") {
		System.gc()
	} else if mainCommand.equals("htop") {
		NewHTopViewer(argument)
	} else if mainCommand.equals("top") {
		if argument.equals("") {
			NewHTopViewer("monitor")
		} else if argument.equals("used") {
			rcvr.echoCommand(fmt.Sprintf("%v%v", "", (rcvr.runtime.totalMemory()-rcvr.runtime.freeMemory())/1024))
		} else if argument.equals("free") {
			rcvr.echoCommand(fmt.Sprintf("%v%v", "", rcvr.runtime.freeMemory()/1024))
		} else if argument.equals("total") {
			rcvr.echoCommand(fmt.Sprintf("%v%v", "", rcvr.runtime.totalMemory()/1024))
		} else {
			rcvr.echoCommand(fmt.Sprintf("%v%v%v", "top: ", rcvr.getCommand(argument), ": not found"))
			return 127
		}
	} else if mainCommand.equals("start") {
		rcvr.start(argument)
	} else if mainCommand.equals("kill") {
		rcvr.kill(argument)
	} else if mainCommand.equals("stop") {
		rcvr.stop(argument)
	} else if mainCommand.equals("ps") {
		rcvr.echoCommand("PID\tPROCESS")
		KEYS := rcvr.trace.keys()
		for KEYS.hasMoreElements() {
			KEY, ok := KEYS.nextElement().(string)
			if !ok {
				panic("XXX Cast fail for *parser.GoCastType")
			}
			PID, ok := rcvr.trace.get(KEY).(string)
			if !ok {
				panic("XXX Cast fail for *parser.GoCastType")
			}
			rcvr.echoCommand(fmt.Sprintf("%v%v%v", PID, "\t", KEY))
		}
	} else if mainCommand.equals("trace") {
		if argument.equals("") {
		} else if rcvr.getCommand(argument).equals("pid") {
			echoCommand(<<unimp_expr[*grammar.JConditionalExpr]>>)
		} else if rcvr.getCommand(argument).equals("check") {
			echoCommand(<<unimp_expr[*grammar.JConditionalExpr]>>)
		} else {
			rcvr.echoCommand(fmt.Sprintf("%v%v%v", "trace: ", rcvr.getCommand(argument), ": not found"))
			return 127
		}
	} else if mainCommand.equals("pkg") {
		echoCommand(<<unimp_expr[*grammar.JConditionalExpr]>>)
	} else if mainCommand.equals("uname") {
		INFO := ""
		if argument.equals("") || argument.equals("-i") {
			INFO = "$TYPE"
		} else if argument.equals("-a") || argument.equals("--all") {
			INFO = fmt.Sprintf("%v%v%v", "$TYPE (OpenTTY $VERSION) main/$RELEASE ", rcvr.build, " - $CONFIG $PROFILE")
		} else if argument.equals("-r") || argument.equals("--release") {
			INFO = "$VERSION"
		} else if argument.equals("-v") || argument.equals("--build") {
			INFO = rcvr.build
		} else if argument.equals("-s") {
			INFO = "J2ME"
		} else if argument.equals("-m") {
			INFO = "$PROFILE"
		} else if argument.equals("-p") {
			INFO = "$CONFIG"
		} else if argument.equals("-n") {
			INFO = "$HOSTNAME"
		} else {
			rcvr.echoCommand(fmt.Sprintf("%v%v%v", "uname: ", argument, ": not found"))
			return 127
		}
		echoCommand(rcvr.env(INFO))
	} else if mainCommand.equals("hostname") {
		processCommand(<<unimp_expr[*grammar.JConditionalExpr]>>, false)
	} else if mainCommand.equals("hostid") {
		DATA := System.getProperty("microedition.platform") + System.getProperty("microedition.configuration") + System.getProperty("microedition.profiles")
		HASH := 7
		for i := 0; i < DATA.length(); i++ {
			HASH = HASH*31 + DATA.charAt(i)
		}
		echoCommand(Integer.toHexString(HASH).toLowerCase())
	} else if mainCommand.equals("tty") {
		echoCommand(rcvr.env("$TTY"))
	} else if mainCommand.equals("ttysize") {
		rcvr.echoCommand(rcvr.stdout.getText().length() + " B")
	} else if mainCommand.equals("echo") {
		rcvr.echoCommand(argument)
	} else if mainCommand.equals("buff") {
		rcvr.stdin.setString(argument)
	} else if mainCommand.equals("uuid") {
		echoCommand(rcvr.generateUUID())
	} else if mainCommand.equals("locale") {
		echoCommand(rcvr.env("$LOCALE"))
	} else if mainCommand.equals("expr") {
		echoCommand(rcvr.exprCommand(argument))
	} else if mainCommand.equals("basename") {
		echoCommand(rcvr.basename(argument))
	} else if mainCommand.equals("getopt") {
		echoCommand(rcvr.getArgument(argument))
	} else if mainCommand.equals("trim") {
		rcvr.stdout.setText(rcvr.stdout.getText().trim())
	} else if mainCommand.equals("date") {
		echoCommand(Newjava.util.Date().toString())
	} else if mainCommand.equals("clear") {
		if argument.equals("") || argument.equals("stdout") {
			rcvr.stdout.setText("")
		} else if argument.equals("stdin") {
			rcvr.stdin.setString("")
		} else if argument.equals("history") {
			rcvr.history = NewVector()
		} else if argument.equals("logs") {
			rcvr.logs = ""
		} else {
			rcvr.echoCommand(fmt.Sprintf("%v%v%v", "clear: ", argument, ": not found"))
			return 127
		}
	} else if mainCommand.equals("seed") {
		if try() {
			rcvr.echoCommand(fmt.Sprintf("%v%v%v", "", rcvr.random.nextInt(Integer.parseInt(argument)), ""))
		} else if catch_NumberFormatException(e) {
			echoCommand(e.getMessage())
			return 2
		}
	} else if mainCommand.equals("throw") {
		Thread.currentThread().interrupt()
	} else if mainCommand.equals("mmspt") {
		echoCommand(replace(rcvr.replace(Thread.currentThread().getName(), "MIDletEventQueue", "MIDlet"), "Thread-1", "MIDlet"))
	} else if mainCommand.equals("bg") {
		bgCommand := argument
		NewAnonymous_Thread_0("Background").start()
	} else if mainCommand.equals("call") {
		if argument.equals("") {
		} else {
			if try() {
				platformRequest(fmt.Sprintf("%v%v", "tel:", argument))
			} else if catch_Exception(e) {
			}
		}
	} else if mainCommand.equals("open") {
		if argument.equals("") {
		} else {
			if try() {
				platformRequest(argument)
			} else if catch_Exception(e) {
				rcvr.echoCommand(fmt.Sprintf("%v%v%v", "open: ", argument, ": not found"))
				return 127
			}
		}
	} else if mainCommand.equals("prg") {
		if argument.equals("") {
			argument = "5"
		}
		if try() {
			PushRegistry.registerAlarm(<<unimp_expr[*grammar.JConditionalExpr]>>, System.currentTimeMillis()+Integer.parseInt(rcvr.getCommand(argument))*1000)
		} else if catch_NumberFormatException(e) {
			echoCommand(e.getMessage())
			return 2
		} else if catch_NumberFormatException(e) {
			rcvr.echoCommand(fmt.Sprintf("%v%v%v", "prg: ", rcvr.getArgument(argument), ": not found"))
			return 127
		} else if catch_NumberFormatException(e) {
			rcvr.echoCommand(fmt.Sprintf("%v%v", "AutoRunError: ", e.getMessage()))
			return 3
		}
	} else if mainCommand.equals("bind") {
		NewBind(<<unimp_expr[*grammar.JConditionalExpr]>>)
	} else if mainCommand.equals("server") {
		NewServer(rcvr.env("$PORT $RESPONSE"))
	} else if mainCommand.equals("gobuster") {
		NewGoBuster(argument)
	} else if mainCommand.equals("pong") {
		if argument.equals("") {
		} else {
			START := System.currentTimeMillis()
			if try() {
				CONN, ok := Connector.open(fmt.Sprintf("%v%v", "socket://", argument)).(*SocketConnection)
				if !ok {
					panic("XXX Cast fail for *parser.GoCastType")
				}
				rcvr.echoCommand(fmt.Sprintf("%v%v%v%v%v", "Pong to ", argument, " successful, time=", System.currentTimeMillis()-START, "ms"))
				CONN.close()
			} else if catch_IOException(e) {
				rcvr.echoCommand(fmt.Sprintf("%v%v%v%v", "Pong to ", argument, " failed: ", e.getMessage()))
				return 101
			}
		}
	} else if mainCommand.equals("ping") {
		if argument.equals("") {
		} else {
			START := System.currentTimeMillis()
			if try() {
				CONN, ok := Connector.open(<<unimp_expr[*grammar.JConditionalExpr]>>).(*HttpConnection)
				if !ok {
					panic("XXX Cast fail for *parser.GoCastType")
				}
				CONN.setRequestMethod(HttpConnection.GET)
				responseCode := CONN.getResponseCode()
				CONN.close()
				rcvr.echoCommand(fmt.Sprintf("%v%v%v%v%v", "Ping to ", argument, " successful, time=", System.currentTimeMillis()-START, "ms"))
			} else if catch_IOException(e) {
				rcvr.echoCommand(fmt.Sprintf("%v%v%v%v", "Ping to ", argument, " failed: ", e.getMessage()))
				return 101
			}
		}
	} else if mainCommand.equals("curl") || mainCommand.equals("wget") || mainCommand.equals("clone") || mainCommand.equals("proxy") {
		if argument.equals("") {
		} else {
			URL := rcvr.getCommand(argument)
			if mainCommand.equals("clone") || mainCommand.equals("proxy") {
				URL = getAppProperty("MIDlet-Proxy") + URL
			}
			HEADERS := <<unimp_expr[*grammar.JConditionalExpr]>>
			RESPONSE := rcvr.request(URL, HEADERS)
			if mainCommand.equals("curl") {
				rcvr.echoCommand(RESPONSE)
			} else if mainCommand.equals("wget") || mainCommand.equals("proxy") {
				rcvr.nanoContent = RESPONSE
			} else if mainCommand.equals("clone") {
				return runScript(RESPONSE)
			}
		}
	} else if mainCommand.equals("query") {
		return rcvr.query(argument)
	} else if mainCommand.equals("prscan") {
		NewPortScanner(argument)
	} else if mainCommand.equals("gaddr") {
		return rcvr.getAddress(argument)
	} else if mainCommand.equals("nc") {
		NewRemoteConnection(argument)
	} else if mainCommand.equals("wrl") {
		echoCommand(System.getProperty("wireless.messaging.sms.smsc"))
	} else if mainCommand.equals("who") {
		SESSIONS := NewStringBuffer()
		for i := 0; i < len(rcvr.sessions); i++ {
			SESSIONS.append(rcvr.sessions.elementAt(i).(string)).append("\n")
		}
		echoCommand(SESSIONS.toString().trim())
	} else if mainCommand.equals("fw") {
		echoCommand(request(fmt.Sprintf("%v%v", "http://ipinfo.io/", <<unimp_expr[*grammar.JConditionalExpr]>>)))
	} else if mainCommand.equals("genip") {
		rcvr.echoCommand(fmt.Sprintf("%v%v%v%v%v%v", rcvr.random.nextInt(256)+".", rcvr.random.nextInt(256), ".", rcvr.random.nextInt(256), ".", rcvr.random.nextInt(256)))
	} else if mainCommand.equals("ifconfig") {
		if argument.equals("") {
			argument = "1.1.1.1:53"
		}
		if try() {
			CONN, ok := Connector.open(fmt.Sprintf("%v%v", "socket://", argument)).(*SocketConnection)
			if !ok {
				panic("XXX Cast fail for *parser.GoCastType")
			}
			echoCommand(CONN.getLocalAddress())
			CONN.close()
		} else if catch_IOException(e) {
			rcvr.echoCommand("null")
			return 101
		}
	} else if mainCommand.equals("report") {
		rcvr.processCommand("open mailto:felipebr4095@gmail.com")
	} else if mainCommand.equals("mail") {
		echoCommand(request(getAppProperty("MIDlet-Proxy") + "raw.githubusercontent.com/mrlima4095/OpenTTY-J2ME/main/assets/root/mail.txt"))
	} else if mainCommand.equals("netstat") {
		STATUS := 0
		if try() {
			CONN, ok := Connector.open("http://ipinfo.io/ip").(*HttpConnection)
			if !ok {
				panic("XXX Cast fail for *parser.GoCastType")
			}
			CONN.setRequestMethod(HttpConnection.GET)
			if CONN.getResponseCode() == HttpConnection.HTTP_OK {
			} else {
				STATUS = 101
			}
			CONN.close()
		} else if catch_Exception(e) {
			STATUS = 101
		}
		echoCommand(<<unimp_expr[*grammar.JConditionalExpr]>>)
		return STATUS
	} else if mainCommand.equals("dir") {
		NewExplorer()
	} else if mainCommand.equals("pwd") {
		rcvr.echoCommand(rcvr.path)
	} else if mainCommand.equals("umount") {
		rcvr.paths = NewHashtable()
	} else if mainCommand.equals("mount") {
		if argument.equals("") {
		} else {
			rcvr.mount(rcvr.getcontent(argument))
		}
	} else if mainCommand.equals("cd") {
		if argument.equals("") {
			rcvr.path = "/home/"
		} else if argument.equals("..") {
			if rcvr.path.equals("/") {
				return 0
			}
			lastSlashIndex := rcvr.path.lastIndexOf('/', <<unimp_expr[*grammar.JConditionalExpr]>>)
			rcvr.path = <<unimp_expr[*grammar.JConditionalExpr]>>
		} else {
			TARGET := <<unimp_expr[*grammar.JConditionalExpr]>>
			if !TARGET.endsWith("/") {
				TARGET += "/"
			}
			if rcvr.paths.containsKey(TARGET) {
				rcvr.path = TARGET
			} else if TARGET.startsWith("/mnt/") {
				if try() {
					REALPWD := fmt.Sprintf("%v%v", "file:///", TARGET.substring(5))
					if !REALPWD.endsWith("/") {
						REALPWD += "/"
					}
					fc, ok := Connector.open(REALPWD, Connector.READ).(*FileConnection)
					if !ok {
						panic("XXX Cast fail for *parser.GoCastType")
					}
					if fc.exists() && fc.isDirectory() {
						rcvr.path = TARGET
					} else {
						rcvr.echoCommand(fmt.Sprintf("%v%v%v%v", "cd: ", rcvr.basename(TARGET), ": not ", <<unimp_expr[*grammar.JConditionalExpr]>>))
						return 127
					}
					fc.close()
				} else if catch_IOException(e) {
					rcvr.echoCommand(fmt.Sprintf("%v%v%v%v", "cd: ", rcvr.basename(TARGET), ": ", e.getMessage()))
					return 1
				}
			} else {
				rcvr.echoCommand(fmt.Sprintf("%v%v%v", "cd: ", rcvr.basename(TARGET), ": not accessible"))
				return 127
			}
		}
	} else if mainCommand.equals("pushd") {
		if argument.equals("") {
			echoCommand(<<unimp_expr[*grammar.JConditionalExpr]>>)
		} else {
			STATUS := processCommand(fmt.Sprintf("%v%v", "cd ", argument), false)
			if STATUS == 0 {
				rcvr.stack.addElement(rcvr.path)
				echoCommand(rcvr.readStack())
			}
			return STATUS
		}
	} else if mainCommand.equals("popd") {
		if len(rcvr.stack) == 0 {
			rcvr.echoCommand("popd: empty stack")
		} else {
			rcvr.path = rcvr.stack.lastElement().(string)
			rcvr.stack.removeElementAt(len(rcvr.stack) - 1)
			echoCommand(rcvr.readStack())
		}
	} else if mainCommand.equals("ls") {
		BUFFER := NewVector()
		if rcvr.path.equals("/mnt/") {
			if try() {
				ROOTS := FileSystemRegistry.listRoots()
				for ROOTS.hasMoreElements() {
					ROOT, ok := ROOTS.nextElement().(string)
					if !ok {
						panic("XXX Cast fail for *parser.GoCastType")
					}
					if !BUFFER.contains(ROOT) {
						BUFFER.addElement(ROOT)
					}
				}
			} else if catch_Exception(e) {
			}
		} else if rcvr.path.startsWith("/mnt/") {
			if try() {
				REALPWD := fmt.Sprintf("%v%v", "file:///", rcvr.path.substring(5))
				if !REALPWD.endsWith("/") {
					REALPWD += "/"
				}
				CONN, ok := Connector.open(REALPWD, Connector.READ).(*FileConnection)
				if !ok {
					panic("XXX Cast fail for *parser.GoCastType")
				}
				CONTENT := CONN.list()
				for CONTENT.hasMoreElements() {
					ITEM, ok := CONTENT.nextElement().(string)
					if !ok {
						panic("XXX Cast fail for *parser.GoCastType")
					}
					BUFFER.addElement(ITEM)
				}
				CONN.close()
			} else if catch_Exception(e) {
			}
		} else if rcvr.path.equals("/home/") && argument.indexOf("-v") != -1 {
			if try() {
				FILES := RecordStore.listRecordStores()
				if FILES != nil {
					for i := 0; i < len(FILES); i++ {
						NAME := FILES[i]
						if (argument.indexOf("-a") != -1 || !NAME.startsWith(".")) && !BUFFER.contains(NAME) {
							BUFFER.addElement(NAME)
						}
					}
				}
			} else if catch_RecordStoreException(e) {
			}
		} else if rcvr.path.equals("/home/") {
			NewExplorer()
			return 0
		}
		FILES, ok := rcvr.paths.get(rcvr.path).([]string)
		if !ok {
			panic("XXX Cast fail for *parser.GoCastType")
		}
		if FILES != nil {
			for i := 0; i < len(FILES); i++ {
				f := FILES[i].trim()
				if f == nil || f.equals("..") || f.equals("/") {
					continue
				}
				if !BUFFER.contains(f) && !BUFFER.contains(fmt.Sprintf("%v%v", f, "/")) {
					BUFFER.addElement(f)
				}
			}
		}
		if !(len(BUFFER) == 0) {
			FORMATTED := NewStringBuffer()
			for i := 0; i < len(BUFFER); i++ {
				ITEM, ok := BUFFER.elementAt(i).(string)
				if !ok {
					panic("XXX Cast fail for *parser.GoCastType")
				}
				if !ITEM.equals("/") {
					FORMATTED.append(ITEM).append(<<unimp_expr[*grammar.JConditionalExpr]>>)
				}
			}
			echoCommand(FORMATTED.toString().trim())
		}
	} else if mainCommand.equals("fdisk") {
		return processCommand("lsblk", false)
	} else if mainCommand.equals("lsblk") {
		if argument.equals("") || argument.equals("-x") {
			echoCommand(replace("MIDlet.RMS.Storage", ".", <<unimp_expr[*grammar.JConditionalExpr]>>))
		} else {
			rcvr.echoCommand(fmt.Sprintf("%v%v%v", "lsblk: ", argument, ": not found"))
			return 127
		}
	} else if mainCommand.equals("rm") {
		if argument.equals("") {
		} else {
			rcvr.deleteFile(argument)
		}
	} else if mainCommand.equals("install") {
		if argument.equals("") {
		} else {
			rcvr.writeRMS(argument, rcvr.nanoContent)
		}
	} else if mainCommand.equals("touch") {
		if argument.equals("") {
			rcvr.nanoContent = ""
		} else {
			rcvr.writeRMS(argument, "")
		}
	} else if mainCommand.equals("mkdir") {
		if argument.equals("") {
		} else {
			argument = <<unimp_expr[*grammar.JConditionalExpr]>>
			argument = <<unimp_expr[*grammar.JConditionalExpr]>>
			if argument.startsWith("/mnt/") {
				if try() {
					CONN, ok := Connector.open(fmt.Sprintf("%v%v", "file:///", argument.substring(5)), Connector.READ_WRITE).(*FileConnection)
					if !ok {
						panic("XXX Cast fail for *parser.GoCastType")
					}
					if !CONN.exists() {
						CONN.mkdir()
						CONN.close()
					} else {
						rcvr.echoCommand(fmt.Sprintf("%v%v%v", "mkdir: ", rcvr.basename(argument), ": found"))
					}
					CONN.close()
				} else if catch_SecurityException(e) {
					echoCommand(e.getMessage())
					return 13
				} else if catch_SecurityException(e) {
					echoCommand(e.getMessage())
					return 1
				}
			} else if argument.startsWith("/home/") {
				rcvr.echoCommand("mkdir: 405 Method not allowed")
				return 3
			} else if argument.startsWith("/") {
				rcvr.echoCommand("read-only storage")
				return 5
			}
		}
	} else if mainCommand.equals("cp") {
		if argument.equals("") {
			rcvr.echoCommand("cp: missing [origin]")
		} else {
			ORIGIN := rcvr.getCommand(argument)
			TARGET := rcvr.getArgument(argument)
			writeRMS(<<unimp_expr[*grammar.JConditionalExpr]>>, rcvr.getcontent(ORIGIN))
		}
	} else if mainCommand.equals("raw") {
		rcvr.echoCommand(rcvr.nanoContent)
	} else if mainCommand.equals("rraw") {
		rcvr.stdout.setText(rcvr.nanoContent)
	} else if mainCommand.equals("sed") {
		return rcvr.stringEditor(argument)
	} else if mainCommand.equals("getty") {
		rcvr.nanoContent = rcvr.stdout.getText()
	} else if mainCommand.equals("add") {
		rcvr.nanoContent = <<unimp_expr[*grammar.JConditionalExpr]>>
	} else if mainCommand.equals("du") {
		if argument.equals("") {
		} else {
			processCommand(fmt.Sprintf("%v%v", "wc -c ", argument), false)
		}
	} else if mainCommand.equals("hash") {
		if argument.equals("") {
		} else {
			rcvr.echoCommand(fmt.Sprintf("%v%v", "", rcvr.getcontent(argument).hashCode()))
		}
	} else if mainCommand.equals("cat") {
		if argument.equals("") {
		} else {
			echoCommand(rcvr.getcontent(argument))
		}
	} else if mainCommand.equals("get") {
		if argument.equals("") || argument.equals("nano") {
			rcvr.nanoContent = rcvr.loadRMS("nano")
		} else {
			rcvr.nanoContent = rcvr.getcontent(argument)
		}
	} else if mainCommand.equals("read") {
		if argument.equals("") || <<unimp_obj.nm_*parser.GoMethodAccess>> < 2 {
			return 2
		} else {
			ARGS := rcvr.split(argument, ' ')
			rcvr.attributes.put(ARGS[0], getcontent(ARGS[1]))
		}
	} else if mainCommand.equals("grep") {
		if argument.equals("") || <<unimp_obj.nm_*parser.GoMethodAccess>> < 2 {
			return 2
		} else {
			ARGS := rcvr.split(argument, ' ')
			echoCommand(<<unimp_expr[*grammar.JConditionalExpr]>>)
		}
	} else if mainCommand.equals("find") {
		if argument.equals("") || <<unimp_obj.nm_*parser.GoMethodAccess>> < 2 {
			return 2
		} else {
			ARGS := rcvr.split(argument, ' ')
			CONTENT := getcontent(ARGS[1])
			VALUE, ok := rcvr.parseProperties(CONTENT).get(ARGS[0]).(string)
			if !ok {
				panic("XXX Cast fail for *parser.GoCastType")
			}
			echoCommand(<<unimp_expr[*grammar.JConditionalExpr]>>)
		}
	} else if mainCommand.equals("head") {
		if argument.equals("") {
		} else {
			CONTENT := rcvr.getcontent(argument)
			LINES := rcvr.split(CONTENT, '\n')
			COUNT := Math.min(10, len(LINES))
			for i := 0; i < COUNT; i++ {
				echoCommand(LINES[i])
			}
		}
	} else if mainCommand.equals("tail") {
		if argument.equals("") {
		} else {
			CONTENT := rcvr.getcontent(argument)
			LINES := rcvr.split(CONTENT, '\n')
			COUNT := Math.max(0, len(LINES)-10)
			for i := COUNT; i < len(LINES); i++ {
				echoCommand(LINES[i])
			}
		}
	} else if mainCommand.equals("diff") {
		if argument.equals("") || <<unimp_obj.nm_*parser.GoMethodAccess>> < 2 {
			return 2
		} else {
			FILES := rcvr.split(argument, ' ')
			LINES1 := split(getcontent(FILES[0]), '\n')
			LINES2 := split(getcontent(FILES[1]), '\n')
			MAX_RANGE := Math.max(len(LINES1), len(LINES2))
			for i := 0; i < MAX_RANGE; i++ {
				LINE1 := <<unimp_expr[*grammar.JConditionalExpr]>>
				LINE2 := <<unimp_expr[*grammar.JConditionalExpr]>>
				if !LINE1.equals(LINE2) {
					rcvr.echoCommand(fmt.Sprintf("%v%v%v%v%v%v%v", "--- Line ", i+1, " ---\n< ", LINE1, "\n", "> ", LINE2))
				}
				if i > len(LINES1) || i > len(LINES2) {
					break
				}
			}
		}
	} else if mainCommand.equals("wc") {
		if argument.equals("") {
		} else {
			SHOW_LINES := false
			SHOW_WORDS := false
			SHOW_BYTES := false
			if argument.indexOf("-c") != -1 {
				SHOW_BYTES = true
			} else if argument.indexOf("-w") != -1 {
				SHOW_WORDS = true
			} else if argument.indexOf("-l") != -1 {
				SHOW_LINES = true
			}
			argument = replace(argument, "-w", "")
			argument = replace(argument, "-c", "")
			argument = replace(argument, "-l", "").trim()
			CONTENT := rcvr.getcontent(argument)
			LINES := 0
			WORDS := 0
			CHARS := CONTENT.length()
			LINE_ARRAY := rcvr.split(CONTENT, '\n')
			LINES = len(LINE_ARRAY)
			for i := 0; i < len(LINE_ARRAY); i++ {
				WORD_ARRAY := split(LINE_ARRAY[i], ' ')
				for j := 0; j < len(WORD_ARRAY); j++ {
					if !WORD_ARRAY[j].trim().equals("") {
						WORDS++
					}
				}
			}
			FILENAME := rcvr.basename(argument)
			if SHOW_LINES {
				echoCommand(LINES + "\t" + FILENAME)
			} else if SHOW_WORDS {
				echoCommand(WORDS + "\t" + FILENAME)
			} else if SHOW_BYTES {
				echoCommand(CHARS + "\t" + FILENAME)
			} else {
				echoCommand(LINES + "\t" + WORDS + "\t" + CHARS + "\t" + FILENAME)
			}
		}
	} else if mainCommand.equals("pjnc") {
		rcvr.nanoContent = rcvr.parseJson(rcvr.nanoContent)
	} else if mainCommand.equals("pinc") {
		rcvr.nanoContent = rcvr.parseConf(rcvr.nanoContent)
	} else if mainCommand.equals("conf") {
		echoCommand(parseConf(<<unimp_expr[*grammar.JConditionalExpr]>>))
	} else if mainCommand.equals("json") {
		echoCommand(parseJson(<<unimp_expr[*grammar.JConditionalExpr]>>))
	} else if mainCommand.equals("vnt") {
		if argument.equals("") {
		} else {
			IN := getcontent(rcvr.getCommand(argument))
			OUT := rcvr.getArgument(argument)
			if OUT.equals("") {
				rcvr.nanoContent = rcvr.text2note(IN)
			} else {
				writeRMS(OUT, rcvr.text2note(IN))
			}
		}
	} else if mainCommand.equals("ph2s") {
		BUFFER := NewStringBuffer()
		for i := 0; i < len(rcvr.history)-1; i++ {
			BUFFER.append(rcvr.history.elementAt(i))
			if i < len(rcvr.history)-1 {
				BUFFER.append("\n")
			}
		}
		script := fmt.Sprintf("%v%v", "#!/java/bin/sh\n\n", BUFFER.toString())
		if argument.equals("") || argument.equals("nano") {
			rcvr.nanoContent = script
		} else {
			rcvr.writeRMS(argument, script)
		}
	} else if mainCommand.equals("nano") {
		NewNanoEditor(argument)
	} else if mainCommand.equals("html") {
		rcvr.viewer(rcvr.extractTitle(rcvr.env(rcvr.nanoContent)), rcvr.html2text(rcvr.env(rcvr.nanoContent)))
	} else if mainCommand.equals("view") {
		if argument.equals("") {
		} else {
			viewer(extractTitle(rcvr.env(argument)), html2text(rcvr.env(argument)))
		}
	} else if mainCommand.equals("audio") {
		return rcvr.audio(argument)
	} else if mainCommand.equals("java") {
		return rcvr.java(argument)
	} else if mainCommand.equals("javac") {
		return writeRMS(rcvr.getCommand(argument), rcvr.generateClass(getcontent(rcvr.getArgument(argument))))
	} else if mainCommand.equals("chmod") {
		if argument.equals("") {
		} else {
			NODES := rcvr.parseProperties("http=javax.microedition.io.Connector.http\nsocket=javax.microedition.io.Connector.socket\nfile=javax.microedition.io.Connector.file\nprg=javax.microedition.io.PushRegistry")
			STATUS := 0
			if NODES.containsKey(argument) {
				if try() {
					if argument.equals("http") {
						Connector.open("http://google.com").(*HttpConnection).close()
					} else if argument.equals("socket") {
						Connector.open(rcvr.env("socket://127.0.0.1:1")).(*SocketConnection).close()
					} else if argument.equals("file") {
						FileSystemRegistry.listRoots()
					} else if argument.equals("prg") {
						PushRegistry.registerAlarm(getClass().getName(), System.currentTimeMillis()+1000)
					}
				} else if catch_SecurityException(e) {
					STATUS = 13
				} else if catch_SecurityException(e) {
					STATUS = 1
				}
			} else if argument.equals("*") {
				KEYS := NODES.keys()
				for KEYS.hasMoreElements() {
					processCommand(fmt.Sprintf("%v%v", "chmod ", KEYS.nextElement().(string)), false)
				}
			} else {
				rcvr.echoCommand(fmt.Sprintf("%v%v%v", "chmod: ", argument, ": not found"))
				return 127
			}
			if STATUS == 0 {
				rcvr.mIDletLogs(fmt.Sprintf("%v%v%v", "add info Permission '", NODES.get(argument).(string), "' granted"))
			} else if STATUS == 1 {
				rcvr.mIDletLogs(fmt.Sprintf("%v%v%v", "add debug Permission '", NODES.get(argument).(string), "' granted with exceptions"))
			} else if STATUS == 13 {
				rcvr.mIDletLogs(fmt.Sprintf("%v%v%v", "add error Permission '", NODES.get(argument).(string), "' denied"))
			} else if STATUS == 3 {
				rcvr.mIDletLogs(fmt.Sprintf("%v%v%v", "add warn Unsupported API '", NODES.get(argument).(string), "'"))
			}
			return STATUS
		}
	} else if mainCommand.equals("history") {
		NewHistory()
	} else if mainCommand.equals("debug") {
		return runScript(rcvr.read("/scripts/debug.sh"))
	} else if mainCommand.equals("help") {
		viewer(rcvr.form.getTitle(), rcvr.read("/java/help.txt"))
	} else if mainCommand.equals("man") {
		verbose := argument.indexOf("-v") != -1
		if verbose {
			argument = replace(argument, "-v", "").trim()
		}
		if argument.equals("") {
			argument = "sh"
		}
		content := rcvr.read("/home/man.html")
		if content.equals("") || argument.equals("--update") {
			STATUS := rcvr.processCommand("netstat")
			if STATUS == 0 {
				STATUS = processCommand("execute install /home/nano; tick Downloading...; proxy proxy github.com/mrlima4095/OpenTTY-J2ME/raw/refs/heads/main/assets/root/man.html; install /home/man.html; get; tick;", false)
				if STATUS == 0 {
					content = rcvr.read("/home/man.html")
				} else {
					return STATUS
				}
			} else {
				rcvr.echoCommand("man: download error")
				return STATUS
			}
		}
		content = rcvr.extractTag(content, argument.toLowerCase(), "")
		if content.equals("") {
			rcvr.echoCommand(fmt.Sprintf("%v%v%v", "man: ", argument, ": not found"))
			return 127
		} else {
			if verbose {
				rcvr.echoCommand(content)
			} else {
				viewer(rcvr.form.getTitle(), content)
			}
		}
	} else if mainCommand.equals("true") || mainCommand.equals("false") || mainCommand.startsWith("#") {
	} else if mainCommand.equals("exit") || mainCommand.equals("quit") {
		rcvr.writeRMS("/home/nano", rcvr.nanoContent)
		notifyDestroyed()
	} else if mainCommand.equals("eval") {
		if argument.equals("") {
		} else {
			rcvr.echoCommand(fmt.Sprintf("%v%v", "", rcvr.processCommand(argument)))
		}
	} else if mainCommand.equals("return") {
		if try() {
			return Integer.valueOf(argument)
		} else if catch_NumberFormatException(e) {
			return 128
		}
	} else if mainCommand.equals("@exec") {
		commandAction(rcvr.eXECUTE, rcvr.display.getCurrent())
	} else if mainCommand.equals("@login") {
		if argument.equals("") {
			rcvr.username = rcvr.loadRMS("OpenRMS")
		} else {
			rcvr.username = argument
		}
	} else if mainCommand.equals("@alert") {
		if try() {
			rcvr.display.vibrate(<<unimp_expr[*grammar.JConditionalExpr]>>)
		} else if catch_NumberFormatException(e) {
			echoCommand(e.getMessage())
			return 2
		}
	} else if mainCommand.equals("@reload") {
		rcvr.shell = NewHashtable()
		rcvr.aliases = NewHashtable()
		rcvr.username = rcvr.loadRMS("OpenRMS")
		rcvr.mIDletLogs("add debug API reloaded")
		rcvr.processCommand("execute x11 stop; x11 init; x11 term; run initd; sh;")
	} else if mainCommand.startsWith("@") {
	} else if mainCommand.equals("about") {
		rcvr.about(argument)
	} else if mainCommand.equals("import") {
		return rcvr.importScript(argument)
	} else if mainCommand.equals("run") {
		return processCommand(fmt.Sprintf("%v%v", ". ", argument), false)
	} else if mainCommand.equals("function") {
		if argument.equals("") {
		} else {
			braceIndex := argument.indexOf('{')
			braceEnd := argument.lastIndexOf('}')
			if braceIndex != -1 && braceEnd != -1 && braceEnd > braceIndex {
				name := rcvr.getCommand(argument).trim()
				body := replace(argument.substring(braceIndex+1, braceEnd).trim(), ";", "\n")
				rcvr.functions.put(name, body)
			} else {
				rcvr.echoCommand("invalid syntax")
				return 2
			}
		}
	} else if mainCommand.equals("!") {
		echoCommand(rcvr.env("main/$RELEASE"))
	} else if mainCommand.equals("!!") {
		if len(rcvr.history) > 0 {
			rcvr.stdin.setString(rcvr.history.elementAt(len(rcvr.history) - 1).(string))
		}
	} else if mainCommand.equals(".") {
		if argument.equals("") {
			return runScript(rcvr.nanoContent)
		} else {
			return runScript(rcvr.getcontent(argument))
		}
	} else {
		rcvr.echoCommand(fmt.Sprintf("%v%v", mainCommand, ": not found"))
		return 127
	}
	return 0
}
func (rcvr *OpenTTY) query(command string) (int) {
	command = env(command.trim())
	mainCommand := rcvr.getCommand(command)
	argument := rcvr.getArgument(command)
	if mainCommand.equals("") {
		rcvr.echoCommand("query: missing [address]")
	} else {
		if try() {
			CONN, ok := Connector.open(mainCommand).(*StreamConnection)
			if !ok {
				panic("XXX Cast fail for *parser.GoCastType")
			}
			IN := CONN.openInputStream()
			OUT := CONN.openOutputStream()
			if !argument.equals("") {
				OUT.write(fmt.Sprintf("%v%v", argument, "\r\n").getBytes())
				OUT.flush()
			}
			BAOS := NewByteArrayOutputStream()
			BUFFER := make([]byte, 1024)
			var LENGTH int
			for (LENGTH = IN.read(BUFFER)) != -1 {
				BAOS.write(BUFFER, 0, LENGTH)
			}
			DATA := NewString(BAOS.toByteArray(), "UTF-8")
			if rcvr.env("$QUERY").equals("$QUERY") || rcvr.env("$QUERY").equals("") {
				rcvr.echoCommand(DATA)
				rcvr.mIDletLogs("add warn Query storage setting not found")
			} else if rcvr.env("$QUERY").toLowerCase().equals("show") {
				rcvr.echoCommand(DATA)
			} else if rcvr.env("$QUERY").toLowerCase().equals("nano") {
				rcvr.nanoContent = DATA
				rcvr.echoCommand("query: data retrieved")
			} else {
				rcvr.writeRMS(rcvr.env("$QUERY"), DATA)
			}
			IN.close()
			OUT.close()
			CONN.close()
		} else if catch_Exception(e) {
			echoCommand(e.getMessage())
			return 1
		}
	}
	return 0
}
func (rcvr *OpenTTY) read(filename string) (string) {
	if try() {
		if filename.startsWith("/mnt/") {
			fileConn, ok := Connector.open(fmt.Sprintf("%v%v", "file:///", filename.substring(5)), Connector.READ).(*FileConnection)
			if !ok {
				panic("XXX Cast fail for *parser.GoCastType")
			}
			is := fileConn.openInputStream()
			content := NewStringBuffer()
			var ch int
			for (ch = is.read()) != -1 {
				content.append(ch.(char))
			}
			is.close()
			fileConn.close()
			return env(content.toString())
		} else if filename.startsWith("/home/") {
			recordStore := nil
			content := ""
			if try() {
				recordStore = RecordStore.openRecordStore(filename.substring(6), true)
				if recordStore.getNumRecords() >= 1 {
					data := recordStore.getRecord(1)
					if data != nil {
						content = NewString(data)
					}
				}
			} else if catch_RecordStoreException(e) {
				content = ""
			} else if finally() {
				if recordStore != nil {
					if try() {
						recordStore.closeRecordStore()
					} else if catch_RecordStoreException(e) {
					}
				}
			}
			return content
		} else {
			content := NewStringBuffer()
			is := getClass().getResourceAsStream(filename)
			if is == nil {
				return ""
			}
			isr := NewInputStreamReader(is, "UTF-8")
			var ch int
			for (ch = isr.read()) != -1 {
				content.append(ch.(char))
			}
			isr.close()
			return env(content.toString())
		}
	} else if catch_IOException(e) {
		return ""
	}
}
func (rcvr *OpenTTY) readStack() (string) {
	sb := NewStringBuffer()
	sb.append(rcvr.path)
	for i := 0; i < len(rcvr.stack); i++ {
		sb.append(" ").append(rcvr.stack.elementAt(i).(string))
	}
	return sb.toString()
}
func (rcvr *OpenTTY) replace(source string, target string, replacement string) (string) {
	result := NewStringBuffer()
	start := 0
	var end int
	for (end = source.indexOf(target, start)) >= 0 {
		result.append(source.substring(start, end))
		result.append(replacement)
		start = end + target.length()
	}
	result.append(source.substring(start))
	return result.toString()
}
func (rcvr *OpenTTY) request(url string, headers *Hashtable) (string) {
	if url == nil || url.length() == 0 {
		return ""
	}
	if !url.startsWith("http://") && !url.startsWith("https://") {
		url = fmt.Sprintf("%v%v", "http://", url)
	}
	if try() {
		conn, ok := Connector.open(url).(*HttpConnection)
		if !ok {
			panic("XXX Cast fail for *parser.GoCastType")
		}
		conn.setRequestMethod(HttpConnection.GET)
		if headers != nil {
			keys := headers.keys()
			for keys.hasMoreElements() {
				key, ok := keys.nextElement().(string)
				if !ok {
					panic("XXX Cast fail for *parser.GoCastType")
				}
				value, ok := headers.get(key).(string)
				if !ok {
					panic("XXX Cast fail for *parser.GoCastType")
				}
				conn.setRequestProperty(key, value)
			}
		}
		is := conn.openInputStream()
		baos := NewByteArrayOutputStream()
		var ch int
		for (ch = is.read()) != -1 {
			baos.write(ch)
		}
		is.close()
		conn.close()
		return NewString(baos.toByteArray(), "UTF-8")
	} else if catch_IOException(e) {
		return e.getMessage()
	}
}
func (rcvr *OpenTTY) request2(url string) (string) {
	return request(url, nil)
}
func (rcvr *OpenTTY) runScript(script string) (int) {
	CMDS := rcvr.split(script, '\n')
	for i := 0; i < len(CMDS); i++ {
		STATUS := processCommand(CMDS[i].trim())
		if STATUS != 0 {
			return STATUS
		}
	}
	return 0
}
func (rcvr *OpenTTY) split(content string, div char) (<<array>>) {
	lines := NewVector()
	start := 0
	for i := 0; i < content.length(); i++ {
		if content.charAt(i) == div {
			lines.addElement(content.substring(start, i))
			start = i + 1
		}
	}
	if start < content.length() {
		lines.addElement(content.substring(start))
	}
	result := make([]string, len(lines))
	lines.copyInto(result)
	return result
}
func (rcvr *OpenTTY) start(app string) (int) {
	if app == nil || app.length() == 0 || rcvr.trace.containsKey(app) {
		return 2
	}
	rcvr.trace.put(app, String.valueOf(1000+rcvr.random.nextInt(9000)))
	if app.equals("sh") {
		rcvr.sessions.addElement("127.0.0.1")
	}
	return 0
}
func (rcvr *OpenTTY) StartApp() {
	if !rcvr.trace.containsKey("sh") {
		rcvr.attributes.put("PATCH", "Hidden Void")
		rcvr.attributes.put("VERSION", getAppProperty("MIDlet-Version"))
		rcvr.attributes.put("RELEASE", "stable")
		rcvr.attributes.put("XVERSION", "0.6.2")
		rcvr.attributes.put("TYPE", System.getProperty("microedition.platform"))
		rcvr.attributes.put("CONFIG", System.getProperty("microedition.configuration"))
		rcvr.attributes.put("PROFILE", System.getProperty("microedition.profiles"))
		rcvr.attributes.put("LOCALE", System.getProperty("microedition.locale"))
		rcvr.runScript(rcvr.read("/java/etc/initd.sh"))
		rcvr.stdin.setLabel(fmt.Sprintf("%v%v%v%v", rcvr.username, " ", rcvr.path, " $"))
		if rcvr.username.equals("") {
			NewLogin()
		} else {
			runScript(rcvr.read("/home/initd"))
		}
	}
}
func (rcvr *OpenTTY) stop(app string) (int) {
	if app == nil || app.length() == 0 {
		return 2
	}
	rcvr.trace.remove(app)
	if app.equals("sh") {
		rcvr.processCommand("exit")
	}
	return 0
}
func (rcvr *OpenTTY) text2note(content string) (string) {
	if content == nil || content.length() == 0 {
		return "BEGIN:VNOTE\nVERSION:1.1\nBODY;ENCODING=QUOTED-PRINTABLE;CHARSET=UTF-8:\nEND:VNOTE"
	}
	content = rcvr.replace(content, "=", "=3D")
	content = rcvr.replace(content, "\n", "=0A")
	vnote := NewStringBuffer()
	vnote.append(fmt.Sprintf("%v%v%v", "BEGIN:VNOTE\nVERSION:1.1\nBODY;ENCODING=QUOTED-PRINTABLE;CHARSET=UTF-8:", content, "\nEND:VNOTE"))
	return vnote.toString()
}
func (rcvr *OpenTTY) viewer(title string, text string) (int) {
	viewer := NewForm(rcvr.env(title))
	viewer.append(NewStringItem(nil, rcvr.env(text)))
	viewer.addCommand(NewCommand("Back", Command.BACK, 1))
	viewer.setCommandListener(rcvr)
	rcvr.display.setCurrent(viewer)
	return 0
}
func (rcvr *OpenTTY) warnCommand(title string, message string) (int) {
	if message == nil || message.length() == 0 {
		return 2
	}
	alert := NewAlert(title, message, nil, AlertType.WARNING)
	alert.setTimeout(Alert.FOREVER)
	rcvr.display.setCurrent(alert)
	return 0
}
func (rcvr *OpenTTY) writeRMS(filename string, data []byte) (int) {
	if filename == nil || filename.length() == 0 {
		return 2
	} else if filename.startsWith("/mnt/") {
		if try() {
			CONN, ok := Connector.open(fmt.Sprintf("%v%v", "file:///", filename.substring(5)), Connector.READ_WRITE).(*FileConnection)
			if !ok {
				panic("XXX Cast fail for *parser.GoCastType")
			}
			if !CONN.exists() {
				CONN.create()
			}
			OUT := CONN.openOutputStream()
			OUT.write(data)
			OUT.flush()
		} else if catch_SecurityException(e) {
			echoCommand(e.getMessage())
			return 13
		} else if catch_SecurityException(e) {
			echoCommand(e.getMessage())
			return 1
		}
	} else if filename.startsWith("/home/") {
		CONN := nil
		if try() {
			CONN = RecordStore.openRecordStore(filename.substring(6), true)
			if CONN.getNumRecords() > 0 {
				CONN.setRecord(1, data, 0, len(data))
			} else {
				CONN.addRecord(data, 0, len(data))
			}
		} else if catch_RecordStoreException(e) {
		} else if finally() {
			if CONN != nil {
				if try() {
					CONN.closeRecordStore()
				} else if catch_RecordStoreException(e) {
				}
			}
		}
	} else if filename.startsWith("/") {
		rcvr.echoCommand("read-only storage")
		return 5
	} else {
		return writeRMS(fmt.Sprintf("%v%v", rcvr.path, filename), data)
	}
	return 0
}
func (rcvr *OpenTTY) writeRMS2(filename string, data string) (int) {
	return writeRMS(filename, data.getBytes())
}
func (rcvr *OpenTTY) xserver(command string) (int) {
	command = env(command.trim())
	mainCommand := rcvr.getCommand(command)
	argument := rcvr.getArgument(command)
	if mainCommand.equals("") {
		viewer("OpenTTY X.Org", rcvr.env("OpenTTY X.Org - X Server $XVERSION\nRelease Date: 2025-05-04\nX Protocol Version 1, Revision 3\nBuild OS: $TYPE"))
	} else if mainCommand.equals("version") {
		rcvr.echoCommand(rcvr.env("X Server $XVERSION"))
	} else if mainCommand.equals("buffer") {
		rcvr.echoCommand(fmt.Sprintf("%v%v%v%v%v", "", rcvr.display.getCurrent().getWidth(), "x", rcvr.display.getCurrent().getHeight(), ""))
	} else if mainCommand.equals("term") {
		rcvr.display.setCurrent(rcvr.form)
	} else if mainCommand.equals("stop") {
		rcvr.form.setTitle("")
		rcvr.form.setTicker(nil)
		rcvr.form.deleteAll()
		rcvr.xserver("cmd hide")
		rcvr.xserver("font")
		rcvr.form.removeCommand(rcvr.eXECUTE)
	} else if mainCommand.equals("init") {
		rcvr.form.setTitle(rcvr.env("OpenTTY $VERSION"))
		rcvr.form.append(rcvr.stdout)
		rcvr.form.append(rcvr.stdin)
		rcvr.form.addCommand(rcvr.eXECUTE)
		rcvr.xserver("cmd")
		rcvr.form.setCommandListener(rcvr)
	} else if mainCommand.equals("xfinit") {
		if argument.equals("") {
			rcvr.xserver("init")
		}
		if argument.equals("stdin") {
			rcvr.form.append(rcvr.stdin)
		} else if argument.equals("stdout") {
			rcvr.form.append(rcvr.stdout)
		}
	} else if mainCommand.equals("cmd") {
		CMDS := []*Command{rcvr.hELP, rcvr.nANO, rcvr.cLEAR, rcvr.hISTORY}
		for i := 0; i < len(CMDS); i++ {
			if argument.equals("hide") {
				rcvr.form.removeCommand(CMDS[i])
			} else {
				rcvr.form.addCommand(CMDS[i])
			}
		}
	} else if mainCommand.equals("title") {
		rcvr.display.getCurrent().setTitle(argument)
	} else if mainCommand.equals("font") {
		if argument.equals("") {
			rcvr.xserver("font default")
		} else {
			rcvr.stdout.setFont(rcvr.newFont(argument))
		}
	} else if mainCommand.equals("tick") {
		current := rcvr.display.getCurrent()
		current.setTicker(<<unimp_expr[*grammar.JConditionalExpr]>>)
	} else if mainCommand.equals("gauge") {
		alert := NewAlert(rcvr.form.getTitle(), argument, nil, AlertType.WARNING)
		alert.setTimeout(Alert.FOREVER)
		alert.setIndicator(NewGauge(nil, false, Gauge.INDEFINITE, Gauge.CONTINUOUS_RUNNING))
		rcvr.display.setCurrent(alert)
	} else if mainCommand.equals("set") {
		if argument.equals("") {
		} else {
			rcvr.desktops.put(argument, rcvr.display.getCurrent())
		}
	} else if mainCommand.equals("load") {
		if argument.equals("") {
		} else {
			if rcvr.desktops.containsKey(argument) {
				rcvr.display.setCurrent(rcvr.desktops.get(argument).(*Displayable))
			} else {
				rcvr.echoCommand(fmt.Sprintf("%v%v%v", "x11: load: ", argument, ": not found"))
				return 127
			}
		}
	} else if mainCommand.equals("canvas") {
		rcvr.display.setCurrent(NewMyCanvas(<<unimp_expr[*grammar.JConditionalExpr]>>))
	} else if mainCommand.equals("make") || mainCommand.equals("list") || mainCommand.equals("quest") {
		NewScreen(mainCommand, argument)
	} else if mainCommand.equals("item") {
		NewItemLoader(rcvr.form, "item", argument)
	} else {
		rcvr.echoCommand(fmt.Sprintf("%v%v%v", "x11: ", mainCommand, ": not found"))
		return 127
	}
	return 0
}
