package main

import (
	"errors"
	"fmt"
	"reflect"
)

type IGame interface {
	getPlayerById(id int) (*IPlayer, error)
}

type IPlayer interface {
	ShowData()
}

type Game struct {
	Id  int
	Ptr interface{}
	// 游戏中有玩家
	Players map[int]*Player
}

// 获取玩家 希望每个子游戏类型都可以继承这个方法
func (g *Game) getPlayerById(id int) (IPlayer, error) {
	pdata, ok := g.Players[id]
	if ok {
		pdata.ShowData()
		return pdata, nil
	} else {
		return nil, errors.New("this id not exist")
	}
}

// 相当于基类
type Player struct {
	Id   int
	Name string
}

func (p *Player) ShowData() {
	fmt.Println("Player Data :Id = ", p.Id, "| Name = ", p.Name)
}

type Game1 struct {
	Name   string
	Age    int
	Weight int
}

func (g *Game1) SayHello() {
	fmt.Println("111222:", g.Name, g.Age, g.Weight)
}

func main() {
	t := reflect.ValueOf(Game1{}).Type()
	//    h := reflect.New(t).Elem()
	// new return address pointer
	h := reflect.New(t).Interface()
	fmt.Println(h)
	hh := h.(*Game1)
	fmt.Println(hh)
	hh.SayHello()
	hh.Age = 123
	hh.Name = "abc"
	hh.Weight = 345
	hh.SayHello()
	//-------------

}
