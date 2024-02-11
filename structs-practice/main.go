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

type outputtable interface {
	saver
	Display()
}

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("1")

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

	err = outputData(userNote)
	if err != nil {
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}
}

func printSomething(value any) {
	typedValue, ok := value.(int)

	if ok {
		fmt.Printf("I can do anything with type integer: %v\n", typedValue)
	}

	switch value.(type) {
	case int:
		fmt.Println("integer: ", value)
	case float64:
		fmt.Println("float:", value)
	}
}

func outputData(o outputtable) error {
	o.Display()
	return saveData(o)
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
