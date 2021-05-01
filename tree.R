

library(rpart)
library(RColorBrewer)
library(rpart.plot)
library(rattle)
library(dplyr)
library(caret) 

# helper function that will save the tree for us
save.tree <- function(plot, path, fname, cap, forpred, length){
  pred <- predict(plot, newdata = forpred, type="class")
  tab <- table(forpred$RESULT, pred)
  conf <- confusionMatrix(tab)
  cap <- paste(cap,"; n=", length, ", acc=", round(conf$overall['Accuracy'], digits=5), sep="")

  png(file=paste(path,fname,".png", sep=""))
  fancyRpartPlot(plot, caption = cap)
  dev.off()

  remove(pred, tab, conf, cap)
}

treesubfolder <- paste(plotsfolder,"/trees/",sep = "")
fexist(treesubfolder)

tcontrol <- rpart.control(
  maxdepth = 3
)

# get the columns we need
games.numcols <- games[,2:9]
# create a generation set and a validation set
sample <- sample(1:nrow(games.numcols),size = ceiling(0.90*nrow(games.numcols)),replace = FALSE)
games.num <- games.numcols[sample,]
games.anum <- games.numcols[-sample,]

# create a tree "all" with all of the data
save.tree(rpart(RESULT~., games.num, method="class", control = tcontrol), treesubfolder, "all", "all", games.anum, nrow(games.num))

# split games into top half and bottom half, with a respective sample
games.tophalf <- games.num[games.num$ELO > quantile(games.num$ELO,prob=1-50/100),]
games.tophalf.samp <- games.anum[games.anum$ELO > quantile(games.anum$ELO,prob=1-50/100),]
games.bottomhalf <- games.num[games.num$ELO <= quantile(games.num$ELO,prob=1-50/100),]
games.bottomhalf.samp <- games.anum[games.anum$ELO <= quantile(games.anum$ELO,prob=1-50/100),]


treesubfolder.split <- paste(treesubfolder,"split/",sep="")
fexist(treesubfolder.split)

# create split trees
save.tree(rpart(RESULT~., games.tophalf, method="class", control = tcontrol), treesubfolder.split, "top", "top half", games.tophalf.samp, nrow(games.tophalf))
save.tree(rpart(RESULT~., games.bottomhalf, method="class", control = tcontrol), treesubfolder.split, "bottom", "bottom half", games.bottomhalf.samp, nrow(games.bottomhalf))

# create trees split into 5 different 
max <- 5
treesubfolder.n <- paste(treesubfolder,"split_n/",sep="")
fexist(treesubfolder.n)

ncontrol <- rpart.control(
  minsplit = 1, 
  minbucket = 1, 
  maxdepth = 3,
  cp = 0.001
)

for (x in 1:max){
  val.top <- 1-((x-1)*(100/max))/100
  val.bottom <- 1-(x*(100/max))/100
  games.split <- games.num[games.num$ELO <= quantile(games.num$ELO, prob=val.top) & 
                          games.num$ELO >  quantile(games.num$ELO, prob=val.bottom),]
  games.split.samp <- games.anum[games.anum$ELO <= quantile(games.anum$ELO, prob=val.top) & 
                                   games.anum$ELO >  quantile(games.anum$ELO, prob=val.bottom),]
  title <- paste(val.bottom*100,"% - ", val.top*100, "%; ", min(games.split$ELO), "-", max(games.split$ELO), sep="")
  
  try(save.tree(rpart(RESULT~., games.split, method="class", control = ncontrol), treesubfolder.n, val.top*100, title, games.split.samp, nrow(games.split)))
}



remove(games.num, games.bottomhalf, games.tophalf, val.top, val.bottom, x, max, title, games.split, treesubfolder, treesubfolder.n, treesubfolder.split)