package notes

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title string, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid note data")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (n Note) Display() {
	fmt.Printf("The note is titled %v, Reads the following content - \n%v\n\nCreated at: %v\n", n.Title, n.Content, n.CreatedAt)
}

func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	// Marshaler already returns []byte content
	jsonContent, err := json.Marshal(n)
	if err != nil {
		return err
	}

	// os.WriteFile returns an error, so directly returning it
	return os.WriteFile(fileName, jsonContent, 0644)
}
