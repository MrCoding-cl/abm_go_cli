package match

import (
	"abm_go/helpers"
	"os/exec"
	"path/filepath"
)

func Versus(equipo1 string, equipo2 string) {
	jar_path, _ := filepath.Abs("SoccerBots/teams/LabABM.jar")
	parsing := helpers.Parsing_text(equipo1, equipo2)
	helpers.WriteRobocup(parsing)
	robocup_path, _ := filepath.Abs("SoccerBots/teams/robocup.dsc")
	cmd := exec.Command("java", "-jar", jar_path, robocup_path)
	//save log
	out, _ := cmd.Output()

	//extract numbers form out using regex
	goles := helpers.ExtractNumber(string(out))

	helpers.CreateFile(goles[0] + " " + goles[1])

}
