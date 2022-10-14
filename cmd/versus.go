/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"abm_go/match"
	"github.com/spf13/cobra"
)

// versusCmd represents the versus command
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
