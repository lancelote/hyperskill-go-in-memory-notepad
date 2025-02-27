package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const Prompt = "Enter a command and data: "
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

func (n NotesList) List() {
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	notesList := NewNotesList(5)

	for {
		fmt.Print(Prompt)

		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		command, data := input[0], strings.Join(input[1:], " ")

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
