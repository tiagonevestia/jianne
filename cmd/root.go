package cmd

import (
	"os"
	"os/user"
	"path"

	"github.com/spf13/cobra"
)

var filepath string

var rootCmd = &cobra.Command{
	Use:   "jianne",
	Short: "jianne é um aplicativo de linha de comando 'CLI' para rastrear suas tarefas diárias.",
	Long:  "Um simples aplicativo de ToDo para linha de comando.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		usr, err := user.Current()
		if err != nil {
			printErrorAndExit(err)
		}

		filepath = path.Join(usr.HomeDir, "jianne.json")

		// verifica se o arquivo existe
		if _, err := os.Stat(filepath); err == nil {
			return // como o arquivo existe, o GO executa a função Run do comando
		} else if !os.IsNotExist(err) {
			printErrorAndExit(err)
		}

		f, err := os.Create(filepath)
		if err != nil {
			printErrorAndExit(err)
		}

		if _, err := f.WriteString("[]"); err != nil {
			printErrorAndExit(err)
		}

		if err := f.Sync(); err != nil {
			printErrorAndExit(err)
		}
	},
}

func Execute() {
	rootCmd.Execute()
}
