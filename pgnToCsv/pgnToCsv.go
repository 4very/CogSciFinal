package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

var header = []string{"Event", "Site", "White", "Black", "Result", "UTCDate", "WhiteElo", "BlackElo", "WhiteRatingDiff", "BlackRatingDiff", "ECO", "Opening", "TimeControl", "Termination", "game"}
var defMap = map[string]string{
	"Event":           "",
	"Site":            "",
	"White":           "",
	"Black":           "",
	"Result":          "",
	"UTCDate":         "",
	"WhiteElo":        "",
	"BlackElo":        "",
	"WhiteRatingDiff": "",
	"BlackRatingDiff": "",
	"ECO":             "",
	"Opening":         "",
	"TimeControl":     "",
	"Termination":     "",
	"game":            "",
}

func mapToList(x map[string]string) []string {
	var a []string
	i := 0
	for _, element := range header {
		a = append(a, x[element])
		i++
	}
	return a

}

// func multRep(str string, rep []string) string {
// 	var ret string
// 	ret = str
// 	for _, elt := range rep {
// 		ret = strings.ReplaceAll(ret, elt, "")
// 	}
// 	return ret
// }

func main() {
	godotenv.Load("../.env")

	var folder = os.Getenv("DATA_SUBFOLDER")
	var filename = os.Getenv("FILE_NAME")

	// open up the in and out files
	file, _ := os.Open("../" + folder + "/" + filename + ".pgn")
	scanner := bufio.NewScanner(file)

	ofile, _ := os.Create("../" + folder + "/" + filename + ".csv")
	// convert the output to be a CSV writer
	ocsv := csv.NewWriter(ofile)
	ocsv.Write(header)

	// setup variables for read loop
	realMap := defMap
	count, writ := 0, 0
	totStart := time.Now()
	now := time.Now()

	for scanner.Scan() {
		// get rid of all newline characters
		text := strings.ReplaceAll(scanner.Text(), "\n", "")

		// if the text is blank we want to skip this itr
		if text == "" || len(text) <= 1 {
			continue
		} else if text[0] == '1' || text[0] == ' ' { // if we're on the line that should contain the game

			if text[0] == '1' { // the game is not blank
				realMap["game"] = text
			} else { // the game is blank (due to forfeit before a move)
				realMap["game"] = ""
			}

			if strings.Contains(realMap["Event"], "Classical") {
				// write to csv and inc writ counter
				writ++
				ocsv.Write(mapToList(realMap))
			}
			realMap = defMap // reset map
			count++          // inc line counter

			// used for logging how much is done and and estimate on remaining time
			if count%100000 == 0 {
				fmt.Println("P:", fmt.Sprintf("%.1fm", float64(count)/1000000.0),
					"\tW:", writ,
					"\tL:", time.Since(now),
					"\tE:", time.Since(totStart),
					"\tR:", time.Since(now)*time.Duration(680-(count/100000)))
				now = time.Now()
			}

		} else if text[0] == '[' { // if the line is metadata
			var key, elt string
			key = text[1:strings.Index(text, " ")]
			elt = text[strings.Index(text, "\"")+1 : strings.LastIndex(text, "\"")]
			realMap[key] = elt
		} else { // catch if there was a line where data didnt match the above logic
			fmt.Println("malformed:", text)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
