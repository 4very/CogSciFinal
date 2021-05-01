package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/notnil/chess"
)

// func sconv(str string) int64 {
// 	return (((int64(rune(str[1])) - 49) * 8) + (int64(rune(str[0])) - 97))
// }

// func iconv(i int64) string {
// 	return string((i%8)+97) + string(((i-(i%8))/8)+49)
// }

func mapToList(x map[string]int) []string {
	var a []string
	i := 0
	for _, element := range header2 {
		a = append(a, strconv.Itoa(x[element]))
		i++
	}
	return a

}

var gameex = `1. e4 { [%clk 0:30:00] } e5 { [%clk 0:30:00] } 2. Nf3 { [%clk 0:29:56] } Nc6 { [%clk 0:29:57] } 3. Bb5 { [%clk 0:29:53] } a6 { [%clk 0:29:51] } 4. Bxc6 { [%clk 0:29:52] } bxc6 { [%clk 0:29:50] } 5. Nxe5 { [%clk 0:29:51] } f6 { [%clk 0:29:41] } 6. Nf3 { [%clk 0:29:50] } Bd6 { [%clk 0:29:26] } 7. d3 { [%clk 0:29:47] } Ne7 { [%clk 0:29:21] } 8. O-O { [%clk 0:29:46] } O-O { [%clk 0:29:19] } 9. Nc3 { [%clk 0:29:42] } c5 { [%clk 0:29:12] } 10. Be3 { [%clk 0:29:40] } Nc6 { [%clk 0:28:45] } 11. a4 { [%clk 0:29:33] } a5 { [%clk 0:28:37] } 12. Nd5 { [%clk 0:29:25] } Nb4 { [%clk 0:28:18] } 13. c3 { [%clk 0:29:16] } Nxd3 { [%clk 0:28:11] } 14. Qxd3 { [%clk 0:29:06] } c6 { [%clk 0:28:04] } 15. Nf4 { [%clk 0:28:55] } Ba6 { [%clk 0:28:00] } 16. Qd2 { [%clk 0:28:29] } Bxf4 { [%clk 0:27:55] } 17. Bxf4 { [%clk 0:28:25] } Bxf1 { [%clk 0:27:48] } 18. Rxf1 { [%clk 0:28:21] } h6 { [%clk 0:27:40] } 19. e5 { [%clk 0:28:02] } fxe5 { [%clk 0:27:17] } 20. Bxe5 { [%clk 0:27:58] } Qe7 { [%clk 0:26:58] } 21. Qe3 { [%clk 0:27:53] } d6 { [%clk 0:26:54] } 22. Bg3 { [%clk 0:27:34] } Qxe3 { [%clk 0:26:50] } 23. fxe3 { [%clk 0:27:33] } Rae8 { [%clk 0:26:41] } 24. Bf4 { [%clk 0:27:21] } d5 { [%clk 0:26:32] } 25. Nd2 { [%clk 0:27:14] } d4 { [%clk 0:26:22] } 26. cxd4 { [%clk 0:27:08] } cxd4 { [%clk 0:26:14] } 27. exd4 { [%clk 0:27:03] } Re2 { [%clk 0:26:08] } 28. g3 { [%clk 0:26:50] } Rxd2 { [%clk 0:26:05] } 29. Bxd2 { [%clk 0:26:47] } Rxf1+ { [%clk 0:26:03] } 30. Kxf1 { [%clk 0:26:45] } g6 { [%clk 0:25:47] } 31. Bxa5 { [%clk 0:26:44] } Kg7 { [%clk 0:25:46] } 32. b4 { [%clk 0:26:41] } Kf6 { [%clk 0:25:28] } 33. b5 { [%clk 0:26:36] } cxb5 { [%clk 0:25:23] } 34. axb5 { [%clk 0:26:35] } Kg5 { [%clk 0:25:20] } 35. b6 { [%clk 0:26:33] } Kg4 { [%clk 0:25:15] } 36. Kg2 { [%clk 0:26:30] } h5 { [%clk 0:25:13] } 37. b7 { [%clk 0:26:29] } Kf5 { [%clk 0:25:10] } 38. d5 { [%clk 0:26:24] } g5 { [%clk 0:25:10] } 39. d6 { [%clk 0:26:23] } Ke6 { [%clk 0:25:10] } 40. Bc7 { [%clk 0:26:21] } Kd7 { [%clk 0:25:07] } 41. b8=Q { [%clk 0:26:18] } Kc6 { [%clk 0:24:59] } 42. Qb6+ { [%clk 0:26:16] } Kd7 { [%clk 0:24:55] } 43. h4 { [%clk 0:26:06] } gxh4 { [%clk 0:24:53] } 44. gxh4 { [%clk 0:26:05] } Ke6 { [%clk 0:24:50] } 45. Qe3+ { [%clk 0:25:58] } Kd7 { [%clk 0:24:48] } 46. Qe7+ { [%clk 0:25:56] } Kc6 { [%clk 0:24:47] } 47. Qd8 { [%clk 0:25:50] } Kb7 { [%clk 0:24:44] } 48. d7 { [%clk 0:25:48] } Kc6 { [%clk 0:24:43] } 49. Qb8 { [%clk 0:25:45] } Kxd7 { [%clk 0:24:41] } 50. Bg3 { [%clk 0:25:41] } Kc6 { [%clk 0:24:33] } 51. Kf3 { [%clk 0:25:40] } Kd5 { [%clk 0:24:31] } 52. Qe5+ { [%clk 0:25:38] } Kc4 { [%clk 0:24:28] } 53. Ke3 { [%clk 0:25:34] } Kb4 { [%clk 0:24:24] } 54. Qd5 { [%clk 0:25:32] } Kc3 { [%clk 0:24:22] } 55. Qd3+ { [%clk 0:25:30] } Kb4 { [%clk 0:24:20] } 56. Ke4 { [%clk 0:25:25] } Kc5 { [%clk 0:24:18] } 57. Qd5+ { [%clk 0:25:23] } Kb4 { [%clk 0:24:16] } 58. Kd3 { [%clk 0:25:18] } Ka4 { [%clk 0:24:09] } 59. Qb7 { [%clk 0:25:16] } Ka3 { [%clk 0:24:08] } 60. Kc4 { [%clk 0:25:14] } Ka4 { [%clk 0:24:06] } 61. Qb4# { [%clk 0:25:13] } 1-0`

var header = []string{"Event", "Site", "White", "Black", "Result", "UTCDate", "WhiteElo", "BlackElo", "WhiteRatingDiff", "BlackRatingDiff", "ECO", "Opening", "TimeControl", "Termination", "game", "wp", "wr", "wb", "wkn", "wq", "wk", "bp", "br", "bb", "bkn", "bq", "bk"}
var header2 = []string{"wp", "wr", "wb", "wkn", "wq", "wk", "bp", "br", "bb", "bkn", "bq", "bk"}

var captures = map[string]int{
	"wp":  0,
	"wr":  0,
	"wb":  0,
	"wkn": 0,
	"wq":  0,
	"wk":  0,
	"bp":  0,
	"br":  0,
	"bb":  0,
	"bkn": 0,
	"bq":  0,
	"bk":  0,
}

func main() {

	godotenv.Load("../.env")
	var folder = os.Getenv("DATA_SUBFOLDER")
	var filename = os.Getenv("FILE_NAME")

	totStart := time.Now()
	now := time.Now()
	ifile, _ := os.Open("../" + folder + "/" + filename + ".csv")
	ofile, _ := os.Create("../" + folder + "/" + filename + "takes.csv")

	// set defaults
	pgn, _ := chess.PGN(strings.NewReader(gameex))
	chess.NewGame(pgn)
	realMap := captures
	count := 0

	icsv, ocsv := csv.NewReader(ifile), csv.NewWriter(ofile)
	ocsv.Write(header)

	for {
		line, err := icsv.Read()
		if err == io.EOF { // end of file
			break
		}
		if line[0] == "Event" {
			continue
		}
		pgn, _ = chess.PGN(strings.NewReader(line[len(line)-1]))
		game := chess.NewGame(pgn)
		// fmt.Println(game.Position().Board().Draw())

		for i := 0; i < len(game.Positions())-1; i++ {
			// fmt.Println(game.Positions()[i].Board().Draw())
			// fmt.Println(game.Moves()[i])
			sq := game.Moves()[i].S2()
			cap := game.Positions()[i].Board().Piece(sq)
			if cap.String() != " " { // if there's a capture
				pi := game.Positions()[i+1].Board().Piece(sq)
				ans := ""
				switch pi.String() {
				case "♟":
					ans = "bp"
				case "♙":
					ans = "wp"
				case "♞":
					ans = "bkn"
				case "♘":
					ans = "wkn"
				case "♜":
					ans = "br"
				case "♖":
					ans = "wr"
				case "♝":
					ans = "bb"
				case "♗":
					ans = "wb"
				case "♛":
					ans = "bq"
				case "♕":
					ans = "wq"
				case "♚":
					ans = "bk"
				case "♔":
					ans = "wk"
				default:
					ans = ""
				}

				realMap[ans]++

				// fmt.Println(strconv.Itoa(i)+":\t", ans)
			}
		}

		ocsv.Write(append(line, mapToList(realMap)...))
		for index := range realMap { //reset map
			realMap[index] = 0
		} // reset
		count++
		if count%1000 == 0 {
			fmt.Println("P:", fmt.Sprintf("%.0fk", float64(count)/1000.0),
				"\tL:", time.Since(now),
				"\tE:", time.Since(totStart),
				"\tR:", time.Since(now)*time.Duration(590-(count/1000)))
			now = time.Now()
		}
		// if count == 5000 {
		// 	break
		// }
	}

	// for square, piece := range game.Positions()[0].Board().SquareMap() {
	// 	fmt.Println(square.String())
	// }

}
