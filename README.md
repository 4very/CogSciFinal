# Chess Piece Effectiveness Explored
###### Avery Sommer

### About this Project
This project is meant to explore the effectiveness of different pieces through different elo rankings. As someone who hasn't played any chess in his life I was interested in learning something about the game, in an unusual way. Since chess games contain a lot of data in which I can analyze I chose to hyper-focus on piece importance. More specifically, the strategic attack importance each piece has over different skill levels, measured by ELO rankings. 

I ran into a lot of trouble with this project and did not get as much done as I would've liked. First, I chose to use a PGN file with ~80 million games. This slowed my testing down to a crawl and meant I had to account for a lot of edge cases. I solved this by just used a small sample and then scaling up whenever the programs were stable. Second, there is a package to import the PGN file directly into R but it is incredibly slow so I had to externally convert to a CSV and derive the data I wanted from there. 

This project contains programs written in both R and Go. The Go programs parse and derive that data which I wanted to explore. R is used exclusively as a data analysis language due to it's lack of scalability. The first program we will run is `pgnToCsv.go` which just simply converts a PGN file into a CSV, since PGN is a multiline entry file format it is not easy to work with (see Appendix A for an example of what a chess game in PGN looks like). The second and final Go program is called `takes.go` which gathers all of the data from the chess games. This program iterates through ever single move in every single chess game and counts each take from each piece per side. The usage of each R file is self explanatory, `corrplot.r` creates corrplots of the data, `tree.r` creates trees, etc.

### How to run the project
1. Clone the repo locally
2. Download a database of games in the pgn format
    - I used [the lichess database](https://database.lichess.org/). WARNING they contain a bunch of games and will be very slow to download and process!
3. Create a subfolder for the all of the data files, I recommend `data`
4. Place the .pgn file into the folder you created. You may need to unzip first
5. Create a subfolder for all of the plots, I recommend `plots`
6. Edit the `.env` file in the root folder to the name of the subfolders created and the name of the file downloaded
7. Setup go environment
    - Download go from https://golang.org/
    - After downloaded navigate to project folder and run command 
    >`go install`
8. Run go programs
   - To convert the pgn file to a CSV run command 
   >`go run ./pgnToCsv/pgntocsv.go`
   - To gather the takes data from the pgn file run command 
   >`go run ./takes/takes.go`

There should not be a file in your data folder called `[base file name]takes.csv`. This contains all of the derived data on the number of takes per piece. You can use the CSV file to analyze the data in any way you see fit. 

9. Download and install R studio
    - https://www.rstudio.com/products/rstudio/#rstudio-desktop
10. Setup RStudio
    - Open RStudio and create a new project from and existing folder (`File > New Project`)
    - Select the root directory as your project folder
11. Load Data
    - First run the R file `read_data.R` by opening it up and clicking "Source" in the top right of the editor
    - Depending on the length of the pgn file downloaded this may take a little while
12. Create Corrplots
    - Run file `corrplot.r`, same as above
    - This will create a folder in your plots folder titled corrplots
13. Create Decision Trees
    - Run file `trees.r`, same as above
    - This will create a folder in your plots folder titled trees
14. Look at pretty plots (important!!)

### Appendix
#### Appendix A
```
[Event "Rated Classical game"]
[Site "https://lichess.org/j1dkb5dw"]
[White "BFG9k"]
[Black "mamalak"]
[Result "1-0"]
[UTCDate "2012.12.31"]
[UTCTime "23:01:03"]
[WhiteElo "1639"]
[BlackElo "1403"]
[WhiteRatingDiff "+5"]
[BlackRatingDiff "-8"]
[ECO "C00"]
[Opening "French Defense: Normal Variation"]
[TimeControl "600+8"]
[Termination "Normal"]

1. e4 e6 2. d4 b6 3. a3 Bb7 4. Nc3 Nh6 5. Bxh6 gxh6 6. Be2 Qg5 7. Bg4 h5 8. Nf3 Qg6 9. Nh4 Qg5 10. Bxh5 Qxh4 11. Qf3 Kd8 12. Qxf7 Nc6 13. Qe8# 1-0

[Event "Rated Classical game"]
[Site "https://lichess.org/a9tcp02g"]
[White "Desmond_Wilson"]
[Black "savinka59"]
[Result "1-0"]
[UTCDate "2012.12.31"]
[UTCTime "23:04:12"]
[WhiteElo "1654"]
[BlackElo "1919"]
[WhiteRatingDiff "+19"]
[BlackRatingDiff "-22"]
[ECO "D04"]
[Opening "Queen's Pawn Game: Colle System, Anti-Colle"]
[TimeControl "480+2"]
[Termination "Normal"]

1. d4 d5 2. Nf3 Nf6 3. e3 Bf5 4. Nh4 Bg6 5. Nxg6 hxg6 6. Nd2 e6 7. Bd3 Bd6 8. e4 dxe4 9. Nxe4 Rxh2 10. Ke2 Rxh1 11. Qxh1 Nc6 12. Bg5 Ke7 13. Qh7 Nxd4+ 14. Kd2 Qe8 15. Qxg7 Qh8 16. Bxf6+ Kd7 17. Qxh8 Rxh8 18. Bxh8 1-0
```