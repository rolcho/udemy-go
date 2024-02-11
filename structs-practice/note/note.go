package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	User      User      `json:"user"`
}

func (n Note) Display() {
	fmt.Printf("%v\n%v\n---------\n", n.Title, n.Content)
}

func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.MarshalIndent(n, "", "  ")

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		User: User{
			Name:  "Anonymous",
			Email: "ab@c.de",
		},
	}, nil
}
