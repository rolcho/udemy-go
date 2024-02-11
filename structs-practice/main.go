package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/structs/note"
	"example.com/structs/todo"
)

type saver interface {
	Save() error
}

func main() {
	title, content := getNoteData()
	todoContent := getUserInput("Todo")

	todo, err := todo.New(todoContent)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	err = saveData(userNote)
	if err != nil {
		return
	}

	err = saveData(todo)
	if err != nil {
		return
	}

	userNote.Display()
	todo.Display()
}

func saveData(s saver) error {
	err := s.Save()

	if err != nil {
		fmt.Println("Saving failed")
		return err
	}

	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title")

	content := getUserInput("Note content")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v: ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
