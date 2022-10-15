package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "abm",
	Short: "Bienvenido al ABM de SoccerBot",
	Long: `Este programa permite obtener el resultado del versus entre dos equipos
	Para obtener el resultado usa ./abm versus equipo1 equipo2
    el log se guarda en out/log.txt
`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
