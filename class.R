
library(tidyverse)
library(car)
library(aod)




lm  <- lm(ELO ~ P + R, games.wins)

games.wins.sample <- sample_n(games.wins, 50000)
games.sample <- sample_n(games, 25000)

glm <- glm(RESULT ~ ELO + B + K + Q + R + P + KN, games.sample, family =binomial())




print(summary(glm))
print(confint(glm))
print(wald.test(b = coef(glm), Sigma = vcov(glm)))
#avPlots(glm)

remove(lm,glm, games.wins.sample, games.sample)