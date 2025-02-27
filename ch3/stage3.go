package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const DataPrompt = "Enter a command and data: "
const CapPrompt = "Enter the maximum number of notes: "

const ErrorFull = "[Error] Notepad is full"

const OkCreated = "[OK] The note was successfully created"
const OkCleared = "[OK] All notes were successfully deleted"

const InfoList = "[Info] %d: %s\n"
const InfoExit = "[Info] Bye!"

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

func (n *NotesList) List() {
	for i, x := range n.notes {
		fmt.Printf(InfoList, i+1, x)
	}
}

func (n *NotesList) Create(note string) error {
	if len(n.notes) >= n.cap {
		return errors.New(ErrorFull)
	}
	n.notes = append(n.notes, note)
	return nil
}

func (n *NotesList) Clear() {
	n.notes = n.notes[:0]
}

func getCap() int {
	fmt.Print(CapPrompt)

	val cap int
	fmt.Scan(&cap)
	return cap
}

func getCommandData() (string, string) {
	fmt.Print(DataPrompt)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")
	command, data := input[0], strings.Join(input[1:], " ")
	
	return command, data
}

func main() {
	cap := getCap()
	notesList := NewNotesList(cap)

	for {
		command, data := getCommandData()

		switch command {
		case "exit":
			fmt.Println(InfoExit)
			return
		case "create":
			err := notesList.Create(data)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(OkCreated)
			}
		case "list":
			notesList.List()
		case "clear":
			notesList.Clear()
			fmt.Println(OkCleared)
		default:
			fmt.Println(command + " " + data)
		}
	}
}
