package tenniskata

type tennisGame1 struct {
	m_score1    int
	m_score2    int
	player1Name string
	player2Name string
}

func TennisGame1(player1Name string, player2Name string) TennisGame {
	game := &tennisGame1{
		player1Name: player1Name,
		player2Name: player2Name}

	return game
}

func (game *tennisGame1) WonPoint(playerName string) {
	if playerName == game.player1Name {
		game.m_score1 += 1
	} else {
		game.m_score2 += 1
	}
}

func (game *tennisGame1) GetScore() string {
	score := ""
	if game.m_score1 == game.m_score2 {
		switch game.m_score1 {
		case 0:
			score = "Love-All"
		case 1:
			score = "Fifteen-All"
		case 2:
			score = "Thirty-All"
		default:
			score = "Deuce"
		}
	} else if game.m_score1 >= 4 || game.m_score2 >= 4 {
		minusResult := game.m_score1 - game.m_score2
		if minusResult == 1 {
			score = "Advantage " + game.player1Name
		} else if minusResult == -1 {
			score = "Advantage " + game.player2Name
		} else if minusResult >= 2 {
			score = "Win for " + game.player1Name
		} else {
			score = "Win for " + game.player2Name
		}
	} else {
		score = getScoreName(game.m_score1) + "-" + getScoreName(game.m_score2)
	}
	return score
}

func getScoreName(tempScore int) string {
	switch tempScore {
	case 0:
		return "Love"
	case 1:
		return "Fifteen"
	case 2:
		return "Thirty"
	case 3:
		return "Forty"
	}
	return ""
}
