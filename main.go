package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const (
	commentLifetime = 20 * time.Second
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

type SayArgs struct {
	Text   string
	Voice  string
	Volume float64
	Rate   int
}

func say(sa SayArgs) {
	input := fmt.Sprintf("[[%f]][[%d]]%s", sa.Volume, sa.Rate, sa.Text)
	err := exec.Command("say", "-v", sa.Voice, input).Run()
	if err != nil {
		log.Println("error occurred")
	}
}

func handleRawMessage(data string, config SayConfig) error {
	var m Message

	err := json.Unmarshal([]byte(data), &m)
	if err != nil {
		return err
	}

	switch m.Command {
	case "Got":
		return handleCommentGot([]byte(m.Content), config)
	default:
		return fmt.Errorf("unexpected command: %s", m.Command)
	}
}

func handleCommentGot(content []byte, config SayConfig) error {
	var ccg CtCommentGot

	err := json.Unmarshal(content, &ccg)
	if err != nil {
		return err
	}

	if time.Now().Add(-commentLifetime).Before(ccg.Date) {
		sa := SayArgs{
			Text:   ccg.Comment,
			Voice:  config.Voice.Ja,
			Volume: config.Volume,
			Rate:   config.Rate,
		}
		say(sa)
	}

	return nil
}

func loadConfig() (SayConfig, error) {
	dir := filepath.Dir(os.Args[0])
	path := filepath.Join(dir, "nagome-mac-say.yml")

	file, err := os.Open(path)
	if err != nil {
		return SayConfig{}, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return SayConfig{}, err
	}

	return parseConfig(data)
}

func main() {
	config, err := loadConfig()
	if err != nil {
		log.Println("error occurred while reding config:", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := scanner.Text()
		err = handleRawMessage(data, config)
		if err != nil {
			log.Println(err)
		}
	}

	if err = scanner.Err(); err != nil {
		log.Println("error occurred while reding stdin:", err)
	}
}
