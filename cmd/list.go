package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var allFlag bool

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "Lista de tarefas",
	Args:    cobra.NoArgs,
	Run: func(_ *cobra.Command, _ []string) {
		todos, err := readFromFile(filepath)
		if err != nil {
			printErrorAndExit(err)
		}
		sortTodos(todos)
		printTodos(todos, allFlag)

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&allFlag, "all", "a", false, "Lista todas as tarefas, incluindo as concluídas.")
}

func printTodos(todos []*Todo, all bool) {
	if len(todos) == 0 {
		fmt.Println("Não há tarefas 😊")
		return
	}
	undones := filter(todos, func(t *Todo) bool {
		return !t.Done
	})

	if len(undones) == 0 {
		fmt.Println("Parabéns, as tarefas estão concluídas 🎉")
		if !all {
			return
		}
		fmt.Println()
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Título", "Criado", "Prioridade", "Status"})
	table.SetBorder(false)
	table.SetColumnColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor})

	for _, t := range todos {
		if !all && t.Done {
			continue
		}
		status := ""

		if t.Done {
			status = "   \u2714   "
		}
		table.Append([]string{
			t.Title,
			t.CreatedTimeInWords(),
			strconv.Itoa(t.Priority),
			status})
	}

	table.Render()
}
