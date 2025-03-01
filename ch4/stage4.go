package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const PromptData = "Enter a command and data: "
const PromptCap = "Enter the maximum number of notes: "

const ErrorFull = "[Error] Notepad is full"
const ErrorUnknownCommand = "[Error] Unknown command"
const ErrorMissingData = "[Error] Missing note argument"
const ErrorMissingPos = "[Error] Missing position argument"
const ErrorInvalidPos = "[Error] Invalid position: %s"
const ErrorNothingUpdate = "[Error] There is nothing to update"
const ErrorNothingDelete = "[Error] There is nothing to delete"
const ErrorOutOfBounds = "[Error] Position %d is out of the boundaries [1, %d]\n"

const OkCreated = "[OK] The note was successfully created"
const OkCleared = "[OK] All notes were successfully deleted"
const OkUpdated = "[OK] The note at position %d was successfully updated\n"
const OkDeleted = "[OK] The note at position %d was successfully deleted\n"

const InfoList = "[Info] %d: %s\n"
const InfoExit = "[Info] Bye!"
const InfoEmpty = "[Info] Notepad is empty"

type NotesList struct {
	notes    []string
	capacity int
}

func NewNotesList(cap int) *NotesList {
	return &NotesList{
		notes:    make([]string, 0, cap),
		capacity: cap,
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
	case "update":
		n.Update(data)
	case "delete":
		n.Delete(data)
	default:
		fmt.Println(ErrorUnknownCommand)
	}
}

func GetPosNote(data string) (int, string, error) {
	input := strings.Split(data, " ")
	inputPos, note := input[0], strings.Join(input[1:], " ")

	if inputPos == "" {
		return 0, "", errors.New(ErrorMissingPos)
	}

	pos, err := strconv.Atoi(inputPos)
	if err != nil {
		return 0, "", fmt.Errorf(ErrorInvalidPos, inputPos)
	}

	return pos, note, nil
}

func (n *NotesList) Update(data string) {
	pos, note, err := GetPosNote(data)

	if err != nil {
		fmt.Println(err)
	} else if pos < 0 || pos > n.capacity {
		fmt.Printf(ErrorOutOfBounds, pos, n.capacity)
	} else if strings.TrimSpace(note) == "" {
		fmt.Println(ErrorMissingData)
	} else if pos > len(n.notes) {
		fmt.Println(ErrorNothingUpdate)
	} else {
		n.notes[pos-1] = note
		fmt.Printf(OkUpdated, pos)
	}
}

func (n *NotesList) Delete(data string) {
	pos, _, err := GetPosNote(data)
	if err != nil {
		fmt.Println(err)
	} else if pos < 0 || pos > n.capacity {
		fmt.Printf(ErrorOutOfBounds, pos, n.capacity)
	} else if pos > len(n.notes) {
		fmt.Println(ErrorNothingDelete)
	} else {
		n.notes = append(n.notes[:pos-1], n.notes[pos:]...)
		fmt.Printf(OkDeleted, pos)
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
	if len(n.notes) >= n.capacity {
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

func GetCapacity() int {
	fmt.Print(PromptCap)

	var capacity int
	_, err := fmt.Scan(&capacity)
	if err != nil {
		panic("capacity not provided")
	}
	return capacity
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
	capacity := GetCapacity()
	notesList := NewNotesList(capacity)

	for {
		command, data := GetCommandData()
		notesList.ExecuteCommand(command, data)
	}
}
