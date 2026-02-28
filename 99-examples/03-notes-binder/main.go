package main

import (
	"fmt"
	notes "notes-binder/notes"
	userevents "notes-binder/utils/user-events"
)

func main() {
	title, content := getNotesData()

	note, err := notes.New(title, content)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	note.Display()
	err = note.Save()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Note saved successfully
	fmt.Println("\n\nNote saved successfully!")
}

func getNotesData() (string, string) {
	title := userevents.GetUserData("Enter the note title: ")

	content := userevents.GetUserData("Enter the note content: ")

	return title, content
}
