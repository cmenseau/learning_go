package tenniskata

type tennisGame2 struct {
	P1point int
	P2point int

	player1Name string
	player2Name string
}

func TennisGame2(player1Name string, player2Name string) TennisGame {
	game := &tennisGame2{
		player1Name: player1Name,
		player2Name: player2Name}

	return game
}

func (game *tennisGame2) GetScore() string {
	score := ""
	if game.P1point == game.P2point {
		if game.P1point < 3 {
			if game.P1point == 0 {
				score = "Love"
			}
			if game.P1point == 1 {
				score = "Fifteen"
			}
			if game.P1point == 2 {
				score = "Thirty"
			}
			score += "-All"
		} else {
			score = "Deuce"
		}
	} else {

		score = getScore(game.P1point) + "-" + getScore(game.P2point)

		if game.P1point > game.P2point && game.P2point >= 3 {
			score = "Advantage player1"
		}

		if game.P2point > game.P1point && game.P1point >= 3 {
			score = "Advantage player2"
		}

		if game.P1point >= 4 && (game.P1point-game.P2point) >= 2 {
			score = "Win for player1"
		}
		if game.P2point >= 4 && (game.P2point-game.P1point) >= 2 {
			score = "Win for player2"
		}
	}

	return score
}

func getScore(point int) string {
	if point == 0 {
		return "Love"
	}
	if point == 1 {
		return "Fifteen"
	}
	if point == 2 {
		return "Thirty"
	}
	if point == 3 {
		return "Forty"
	}
	return ""
}

func (game *tennisGame2) WonPoint(player string) {
	if player == "player1" {
		game.P1point++
	} else {
		game.P2point++
	}
}
