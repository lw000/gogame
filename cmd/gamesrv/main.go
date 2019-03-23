package main

import "log"

type Game struct {
	Name string
	Id   int64
}

func (g *Game) Start() error {

	return nil
}

func (g *Game) Stop() error {

	return nil
}

func main() {
	log.Println("game")
}
