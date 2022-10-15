package cmd

import (
	"abm_go/match"
	"github.com/spf13/cobra"
)

var versusCmd = &cobra.Command{
	Use:   "versus",
	Short: "abm versus equipo1 equipo2",
	Long: `Versus entre dos equipos de SoccerBot
    Para obtener el resultado usa ./abm equipo1 equipo2
    Ejemplo:
  	./abm versus AIKHomoG BrianTeam
`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 2 {
			cmd.Help()
			return
		} else {
			equipo1 := args[0]
			equipo2 := args[1]
			match.Versus(equipo1, equipo2)
		}
	},
}

func init() {
	rootCmd.AddCommand(versusCmd)
}
