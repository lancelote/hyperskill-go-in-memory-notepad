package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const PromptData = "Enter a command and data: "
const PromptCap = "Enter the maximum number of notes: "

const ErrorFull = "[Error] Notepad is full"
const ErrorUnknownCommand = "[Error] Unknown command"
const ErrorMissingData = "[Error] Missing note argument"

const OkCreated = "[OK] The note was successfully created"
const OkCleared = "[OK] All notes were successfully deleted"

const InfoList = "[Info] %d: %s\n"
const InfoExit = "[Info] Bye!"
const InfoEmpty = "[Info] Notepad is empty"

type NotesList struct {
	notes []string
	cap   int
}

func NewNotesList(cap int) *NotesList {
	return &NotesList{
		notes: make([]string, 0, cap),
		cap:   cap,
	}
}

func (n *NotesList) ExecuteCommand(command, data string) {
	switch command {
	case "exit":
		fmt.Println(InfoExit)
		os.Exit(0)
	case "create":
		n.Create(data)
	case "list":
		n.List()
	case "clear":
		n.Clear()
	default:
		fmt.Println(ErrorUnknownCommand)
	}
}

func (n *NotesList) List() {
	if len(n.notes) == 0 {
		fmt.Println(InfoEmpty)
	} else {
		for i, x := range n.notes {
			fmt.Printf(InfoList, i+1, x)
		}
	}
}

func (n *NotesList) Create(note string) {
	if len(n.notes) >= n.cap {
		fmt.Println(ErrorFull)
	} else if strings.TrimSpace(note) == "" {
		fmt.Println(ErrorMissingData)
	} else {
		n.notes = append(n.notes, note)
		fmt.Println(OkCreated)
	}
}

func (n *NotesList) Clear() {
	n.notes = n.notes[:0]
	fmt.Println(OkCleared)
}

func GetCap() int {
	fmt.Print(PromptCap)

	var cap int
	fmt.Scan(&cap)
	return cap
}

func GetCommandData() (string, string) {
	fmt.Print(PromptData)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")
	command, data := input[0], strings.Join(input[1:], " ")

	return command, data
}

func main() {
	cap := GetCap()
	notesList := NewNotesList(cap)

	for {
		command, data := GetCommandData()
		notesList.ExecuteCommand(command, data)
	}
}
