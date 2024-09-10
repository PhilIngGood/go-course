package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"course.go/note/note"
	"course.go/note/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()

	todoText := getUserInput("Todo text: ")
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func outputData(output outputtable) error {
	output.Display()
	return output.Save()
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the data failed")
		return err
	}

	fmt.Println("Saving the data succeeded")
	return nil
}

func getTodo() string {
	return getUserInput("Todo text: ")
}

func getNoteData() (string, string) {
	titleInput := getUserInput("Note title:")

	contentInput := getUserInput("Note content:")

	return titleInput, contentInput
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
