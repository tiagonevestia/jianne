package cmd

import (
	"errors"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var priorityFlag int

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adiciona uma nova tarefa",
	Run: func(_ *cobra.Command, _ []string) {
		todos, err := readFromFile(filepath)
		if err != nil {
			printErrorAndExit(err)
		}

		titles := make(map[string]bool)
		for _, t := range todos {
			titles[t.Title] = true
		}

		templates := &promptui.PromptTemplates{
			Prompt:  "{{ . }}",
			Valid:   "{{ . | green }}",
			Invalid: "{{ . | red }}",
			Success: "{{ . | bold }}",
		}

		titlePrompt := promptui.Prompt{
			Label:     "Adicione o nome da sua tarefa: ",
			Templates: templates,
			Validate: func(input string) error {
				if len(strings.TrimSpace(input)) == 0 {
					return errors.New("string vazia")
				}
				if titles[input] {
					return errors.New("já existe")
				}
				return nil
			},
		}

		priorityPrompt := promptui.Prompt{
			Label:     "Adicione a prioridade [padrão = 2]: ",
			Templates: templates,
			Validate: func(input string) error {
				if input == "" {
					return nil
				}
				if _, err := strconv.Atoi(input); err != nil {
					return err
				}
				return nil
			},
		}

		title, err := titlePrompt.Run()
		if err != nil {
			printErrorAndExit(err)
		}

		priority, err := priorityPrompt.Run()
		if err != nil {
			printErrorAndExit(err)
		}

		p := 2
		if priority != "" {
			p, _ = strconv.Atoi(priority)
		}

		t := NewTodo(title, p)
		todos = append(todos, t)

		if err := writeToFile(todos, filepath); err != nil {
			printErrorAndExit(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
