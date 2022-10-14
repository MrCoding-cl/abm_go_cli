package helpers

import (
	"abm_go/template"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Parsing_text(equipo1, equipo2 string) string {
	texto := template.Text
	var seed int = rand.Intn(10000)
	texto = strings.Replace(texto, "Equipo1", equipo1, -1)
	texto = strings.Replace(texto, "Equipo2", equipo2, -1)
	texto = strings.Replace(texto, "seedex", strconv.Itoa(seed), -1)
	return texto
}

func CreateFile(data string) {
	f, err := os.Create("out/log.txt")

	if err != nil {
		log.Fatal(err)
	}

	_, err2 := f.WriteString(data)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("log partido creado")
}

// extract number from string using regex return array of numbers [1,2]
func ExtractNumber(text string) []string {
	re := regexp.MustCompile("(\\d+)\\s(\\d+)")
	match := re.FindStringSubmatch(text)
	var numbers []string
	for i, n := range match {
		if i != 0 {
			numbers = append(numbers, n)
		}
	}
	return numbers

}

func WriteRobocup(texto string) {
	f, err := os.Create("SoccerBots/teams/robocup.dsc")

	if err != nil {
		log.Fatal(err)
	}

	_, err2 := f.WriteString(texto)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("robocup.dsc creado")

}
