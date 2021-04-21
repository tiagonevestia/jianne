package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Deletar todas as tarefas que estão concluídas.",
	Args:  cobra.NoArgs,
	Run: func(_ *cobra.Command, _ []string) {
		const prompt = "Tem certeza que deseja excluir todas as tarefas? [Sim/Não] (padrão \"Não\")"
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print(prompt)
		for scanner.Scan() {
			answer := scanner.Text()
			if contains([]string{"Sim", "sim", "S", "s"}, answer) {
				break
			}
			if contains([]string{"Não", "não", "nao", "n", "N", ""}, answer) {
				return
			}
			fmt.Print("\n" + prompt)
		}
		if err := writeToFile([]*Todo{}, filepath); err != nil {
			printErrorAndExit(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
