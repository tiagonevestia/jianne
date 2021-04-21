package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var undoneCmd = &cobra.Command{
	Use:   "undone",
	Short: "Refazer as tarefas",
	Run: func(_ *cobra.Command, _ []string) {
		todos, err := readFromFile(filepath)
		if err != nil {
			printErrorAndExit(err)
		}

		filtered := filter(todos, func(t *Todo) bool { return t.Done })
		if len(filtered) == 0 {
			fmt.Println("Você ainda não completou as suas tarefas.")
			return
		}

		sortTodos(filtered)
		prompt := promptui.Select{
			Label:    "Selecione uma tarefa para colocar como não concluída",
			Items:    titles(filtered),
			HideHelp: true,
		}

		_, selected, err := prompt.Run()
		if err != nil {
			printErrorAndExit(err)
		}

		for _, t := range todos {
			if t.Title == selected {
				t.Done = false
			}
		}
		if err := writeToFile(todos, filepath); err != nil {
			printErrorAndExit(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(undoneCmd)
}
