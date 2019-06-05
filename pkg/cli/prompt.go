package cli

import (
	"fmt"
	"os"

	"gopkg.in/AlecAivazis/survey.v1"
)

type PromptHandler interface {
	Input(label string) string
	Select(label string, items []string) string
	MultiSelect(label string, items []string) []string
	Confirm(label string) bool
}

type promptHandler struct {}

func NewPromptHandler() PromptHandler {
	return &promptHandler{}
}

func (r *promptHandler) Input(label string) string {
	response := ""
	prompt := &survey.Input{
		Message: label,
	}
	err := survey.AskOne(prompt, &response, nil)

	if err != nil {
		fmt.Println("Canceled")
		os.Exit(1)
		return ""
	}

	return response
}

func (r *promptHandler) Select(label string, items []string) string {
	response := ""
	prompt := &survey.Select{
		Message: label,
		Options: items,
	}
	err := survey.AskOne(prompt, &response, nil)

	if err != nil {
		fmt.Println("Canceled")
		os.Exit(1)
		return ""
	}

	return response
}

func (r *promptHandler) MultiSelect(label string, items []string) []string {
	response := []string{}
	prompt := &survey.MultiSelect{
		Message: label,
		Options: items,
	}
	err := survey.AskOne(prompt, &response, nil)

	if err != nil {
		fmt.Println("Canceled")
		os.Exit(1)
		return []string{}
	}

	return response
}

func (r *promptHandler) Confirm(label string) bool {
	response := false
	prompt := &survey.Confirm{
		Message: label,
	}
	err :=survey.AskOne(prompt, &response, nil)

	if err != nil {
		fmt.Println("Canceled")
		os.Exit(1)
		return false
	}

	return response
}

