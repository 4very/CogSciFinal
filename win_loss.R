
library(ggcorrplot)
library(ggplot2)

games.wins.data <- games.wins[,2:9]
games.losses.data <- games.losses[,2:9]

corrsubfolder <- paste(plotsfolder,"/corrplot/",sep = "")
fexist(corrsubfolder)

corrsubfolder.wl <- paste(corrsubfolder, "wl/",sep = "")
fexist(corrsubfolder.wl)

plot <- ggcorrplot(cor(games.wins.data), title="Wins",
                   type = "lower",
                   lab = TRUE,
                   ggtheme = ggplot2::theme_minimal())
ggsave(paste(corrsubfolder.wl,"wins.png",sep=""), plot=plot)


plot <- ggcorrplot(cor(games.losses.data), title="losses",
                   type = "lower",
                   lab = TRUE,
                   ggtheme = ggplot2::theme_minimal())
ggsave(paste(corrsubfolder.wl,"losses.png",sep=""), plot=plot)



remove(games.losses.data, games.wins.data, plot)