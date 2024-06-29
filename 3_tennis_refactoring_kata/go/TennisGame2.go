package tenniskata

type Player struct {
	point int
	name  string
}

func (p Player) getScoreName() string {
	if p.point == 0 {
		return "Love"
	}
	if p.point == 1 {
		return "Fifteen"
	}
	if p.point == 2 {
		return "Thirty"
	}
	if p.point == 3 {
		return "Forty"
	}
	return ""
}

func (p Player) hasAvantageAgain(p2 Player) bool {
	return p.point > p2.point && p2.point >= 3
}

func (p Player) hasWonAgain(p2 Player) bool {
	return p.point >= 4 && (p.point-p2.point) >= 2
}

func NewPlayer(name string) Player {
	return Player{name: name, point: 0}
}

type tennisGame2 struct {
	player1 Player
	player2 Player
}

func TennisGame2(player1Name string, player2Name string) TennisGame {
	game := &tennisGame2{
		player1: NewPlayer(player1Name),
		player2: NewPlayer(player2Name),
	}

	return game
}

func (game *tennisGame2) GetScore() string {
	score := ""
	if game.isTie() {
		if game.player1.point < 3 {
			score = game.player1.getScoreName() + "-All"
		} else {
			score = "Deuce"
		}
	} else if game.isDone() {
		if game.player1.hasWonAgain(game.player2) {
			score = "Win for " + game.player1.name
		} else {
			score = "Win for " + game.player2.name
		}
	} else if game.isAvantaging() {
		if game.player1.hasAvantageAgain(game.player2) {
			score = "Advantage " + game.player1.name
		} else {
			score = "Advantage " + game.player2.name
		}
	} else {
		score = game.player1.getScoreName() + "-" + game.player2.getScoreName()
	}

	return score
}

func (game *tennisGame2) isTie() bool {
	return game.player1.point == game.player2.point
}

func (game *tennisGame2) isDone() bool {
	return game.player2.hasWonAgain(game.player1) || game.player1.hasWonAgain(game.player2)
}

func (game *tennisGame2) isAvantaging() bool {
	return game.player2.hasAvantageAgain(game.player1) || game.player1.hasAvantageAgain(game.player2)
}

func (game *tennisGame2) WonPoint(player string) {
	if player == game.player1.name {
		game.player1.point++
	} else {
		game.player2.point++
	}
}
