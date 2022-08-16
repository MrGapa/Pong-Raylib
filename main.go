package main

import g "PONG/Game"

func main() {
	game := g.Init()

	for !game.Should_Close {
		game.Update()
		game.Draw()
	}

	game.Close()
}
