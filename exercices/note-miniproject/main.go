package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"course.go/note/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving data to file failed")
		return
	}

	fmt.Println("Saving the note succeeded")
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
