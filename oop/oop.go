package main

// 按照传统opp的方式写的代码

// 相当于基类
type Game struct {
	id int

	// 游戏中有玩家
	Players []Player
}

// 获取玩家 希望每个子游戏类型都可以继承这个方法
func (g *Game) getPlayerById(id int) *Player {
	for _, p := range g.Players {
		if p.id == id {
			return &p
		}
	}

	return nil
}

// 相当于基类
type Player struct {
	id int
}

// a游戏类
type AGame struct {
	Game
	Ptr interface{}

	Players []APlayer
}

func NewAGame() *AGame {
	newAGame := &AGame{}
	return newAGame
}

// a游戏玩家类
type APlayer struct {
	Player
}

func main() {
	a := AGame{}

	// 调用这个方法获取的是 Game里面的Players 不是AGame里面的APlayers 如果在BaseGame里面写APlayer 也不合适
	a.getPlayerById(1)
}
