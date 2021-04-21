package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Marcar tarefas como concluídas.",
	Run: func(_ *cobra.Command, _ []string) {
		todos, err := readFromFile(filepath)
		if err != nil {
			printErrorAndExit(err)
		}

		filtered := filter(todos, func(t *Todo) bool { return !t.Done })
		if len(filtered) == 0 {
			fmt.Println("Parabéns, você não tem mais tarefas 🎉")
			return
		}

		sortTodos(filtered)
		prompt := promptui.Select{
			Label:    "Selecione uma tarefa para marcar como concluída.",
			Items:    titles(filtered),
			HideHelp: true,
		}

		_, selected, err := prompt.Run()
		if err != nil {
			printErrorAndExit(err)
		}

		for _, t := range todos {
			if t.Title == selected {
				t.Done = true
			}
		}

		if err := writeToFile(todos, filepath); err != nil {
			printErrorAndExit(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
