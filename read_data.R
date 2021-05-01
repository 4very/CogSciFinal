library(dotenv)
load_dot_env()

fexist <- function(foldern){
  if (!dir.exists(foldern)){
    dir.create(foldern)
  }
}

filename <- Sys.getenv("FILE_NAME")
foldername <- Sys.getenv("PLOT_SUBFOLDER")
datafolder <- Sys.getenv("DATA_SUBFOLDER")

plotsfolder <- paste("./",foldername,"/",filename,sep="")

fexist(plotsfolder)

data <- read.csv(paste("./",datafolder,"/",filename,"takes.csv",sep=""))
# remove all ties and non-started games
fdata <- data[data$Result != "1/2-1/2" &  data$Result != "*",]

attach(fdata)

wgames <- data.frame("COLOR"="W", "ELO" = WhiteElo, "RESULT" = as.integer(substr(Result,1,1)), "P"=wp, "KN"=wkn, "B"=wb, "K"=wk, "Q"=wq, "R"=wr, "Opening"=Opening)
bgames <- data.frame("COLOR"="B", "ELO" = BlackElo, "RESULT" = as.integer(substr(Result,3,3)), "P"=bp, "KN"=bkn, "B"=bb, "K"=bk, "Q"=bq, "R"=br, "Opening"=Opening)
games <- rbind(wgames,bgames)

for (name in names(games[1,9])){
  games[[name]] = as.numeric(games[[name]])
}
games$ELO = as.numeric(games$ELO)
games <- games[complete.cases(games),]

remove(wgames,bgames)
detach(fdata)

games.wins <- games[games$RESULT==1,]

games.losses <- games[games$RESULT==0,]

