library(corrplot)
library(ggplot2)
library(ggcorrplot)

save.plot <- function(data, path, fname, cap){
  plot <- ggcorrplot(cor(data), title = cap,
                     type = "lower",
                     lab = TRUE,
                     ggtheme = ggplot2::theme_minimal(), tl.col="white")
  ggsave(paste(path,fname,".png", sep=""), plot = plot)
}


corrsubfolder <- paste(plotsfolder,"/corrplot/",sep = "")
fexist(corrsubfolder)

games.num <- games[,2:9]
games.num$ELO <- as.numeric(games.num$ELO)


save.plot(games.num, corrsubfolder, "total", "All Games")

games.tophalf <- games.num[games.num$ELO > quantile(games.num$ELO,prob=1-50/100),]
games.bottomhalf <- games.num[games.num$ELO <= quantile(games.num$ELO,prob=1-50/100),]

corrsubfolder.split <- paste(corrsubfolder, "/split/",sep = "")
fexist(corrsubfolder.split)

save.plot(games.tophalf, corrsubfolder.split, "thalf", paste("Top half of ELO; ", min(games.tophalf$ELO), "-", max(games.tophalf$ELO), sep=""))
save.plot(games.bottomhalf, corrsubfolder.split, "bhalf", paste("Bottom half of ELO; ", min(games.bottomhalf$ELO), "-", max(games.bottomhalf$ELO), sep=""))

plot <- ggcorrplot(cor(games.bottomhalf)-cor(games.tophalf), type = "lower",
                   title = "Bottom half - Top half",
                   lab = TRUE,
                   ggtheme = ggplot2::theme_minimal())
ggsave(paste(corrsubfolder.split,"/diff.png",sep=""), plot=plot)

max <- 5

corrsubfolder.n <- paste(corrsubfolder, "/split_n/",sep = "")
fexist(corrsubfolder.n)

for (x in 1:max){
  val.top <- 1-((x-1)*(100/max))/100
  val.bottom <- 1-(x*(100/max))/100
  data.var <- games.num[games.num$ELO <= quantile(games.num$ELO, prob=val.top) & 
                        games.num$ELO >  quantile(games.num$ELO, prob=val.bottom),]
  title <- paste(val.bottom*100,"% - ", val.top*100, "%; ", min(data.var$ELO), "-", max(data.var$ELO), "; n=", nrow(data.var), sep="")
  
  save.plot(data.var, corrsubfolder.n, (x)*(100/max), paste(min(data.var$ELO), "-", max(data.var$ELO), "; ", 
                                                                     val.bottom*100, "%-", val.top*100, "%", sep=""))
  
}

corrsubfolder.inc200 <- paste(corrsubfolder, "/inc200/",sep = "")
fexist(corrsubfolder.inc200)

for (x in 3:15){
  val <- x * 200
  data.use <- games.num[games.num$ELO >= val & games.num$ELO < val+199,]
  title <- paste(val, "-", val+199, ": n = ", nrow(data.use),sep="")
  
  save.plot(data.use, corrsubfolder.inc200, val, paste(min(data.use$ELO), "-", max(data.use$ELO)))
  
  
}

# all variables
remove(max, title, val, val.bottom, val.top, x, plot)

# all df
remove(games.bottomhalf, games.tophalf, games.num, data.use, data.var)


