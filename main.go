package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

type Message struct {
	Domain  string          `json:"domain"`
	Command string          `json:"command"`
	Content json.RawMessage `json:"content,omitempty"`
}

type CtCommentGot struct {
	No      int       `json:"no"`
	Date    time.Time `json:"date"`
	Raw     string    `json:"raw"`
	Comment string    `json:"comment"`

	UserID           string `json:"user_id"`
	UserName         string `json:"user_name"`
	UserThumbnailURL string `json:"user_thumbnail_url,omitempty"`
	Score            int    `json:"score,omitempty"`
	IsPremium        bool   `json:"is_premium"`
	IsBroadcaster    bool   `json:"is_broadcaster"`
	IsStaff          bool   `json:"is_staff"`
	IsAnonymity      bool   `json:"is_anonymity"`
}

type SayOption struct {
	Voice    string
	Volume   float64
	Rate     int
	Duration time.Duration
}

func say(text string, o SayOption) {
	input := fmt.Sprintf("[[%f]][[%d]]%s", o.Volume, o.Rate, text)
	err := exec.Command("say", "-v", o.Voice, input).Run()
	if err != nil {
		log.Println("error occurred")
	}
}

func handleRawMessage(data string, o SayOption) error {
	var m Message

	err := json.Unmarshal([]byte(data), &m)
	if err != nil {
		return err
	}

	switch m.Command {
	case "Got":
		return handleCommentGot([]byte(m.Content), o)
	default:
		return fmt.Errorf("unexpected command: %s", m.Command)
	}
}

func handleCommentGot(content []byte, o SayOption) error {
	var ccg CtCommentGot

	err := json.Unmarshal(content, &ccg)
	if err != nil {
		return err
	}

	if time.Now().Add(-o.Duration).Before(ccg.Date) {
		say(ccg.Comment, o)
	}

	return nil
}

func main() {
	option := SayOption{
		Voice:    "Kyoko",
		Volume:   0.9,
		Rate:     200,
		Duration: 20 * time.Second,
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := scanner.Text()
		err := handleRawMessage(data, option)
		if err != nil {
			log.Println(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println("error occurred while reding stdin:", err)
	}
}
