package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/slack-go/slack"
	"gopkg.in/yaml.v2"
)

type Config struct {
	WebhookURL string `yaml:"webhookURL"`
}

func main() {
	executablePath, err := os.Executable()
	if err != nil {
		log.Fatal("Error getting executable path:", err)
	}

	configPath := filepath.Join(filepath.Dir(executablePath), "config.yaml")

	var config Config
	if err := loadConfig(configPath, &config); err != nil {
		log.Fatalf("Error loading configuration from %s: %v", configPath, err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading standard input:", err)
	}

	if len(lines) > 0 {
		err := sendSlackNotification(config.WebhookURL, strings.Join(lines, "\n"))
		if err != nil {
			log.Fatalf("Error sending Slack notification: %v", err)
		} else {
			log.Println("Notification sent to Slack successfully.")
		}
	} else {
		log.Println("No output to send to Slack.")
	}
}

func loadConfig(filename string, config *Config) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return fmt.Errorf("failed to decode config file: %v", err)
	}
	return nil
}

func sendSlackNotification(webhookURL, message string) error {
	msg := &slack.WebhookMessage{
		Text: message,
	}

	err := slack.PostWebhook(webhookURL, msg)
	if err != nil {
		return fmt.Errorf("failed to send Slack notification: %v", err)
	}
	return nil
}
