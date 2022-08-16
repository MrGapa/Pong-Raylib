package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Should_Close bool
	player_1     Palette
	player_2     Palette

	ball Ball

	p1_score  int
	p2_score  int
	max_score int

	game_over bool
	start     bool
	winner    string
}

func Init() *Game {
	rl.InitWindow(W_WIDTH, W_HEIGHT, W_TITLE)
	rl.SetTargetFPS(60)

	return &Game{
		Should_Close: false,
		p1_score:     0,
		p2_score:     0,
		max_score:    5,
		player_1: Palette{
			x:      20,
			y:      20,
			width:  25,
			height: 175,
		},
		player_2: Palette{
			x:      760,
			y:      20,
			width:  25,
			height: 175,
		},
		ball: Ball{
			x:       float32(W_WIDTH) / 2,
			y:       float32(W_HEIGHT) / 2,
			size:    25,
			speed_x: BALL_SPEED_X,
			speed_y: BALL_SPEED_y,
		},
		game_over: false,
		start:     false,
		winner:    "",
	}
}

func (g *Game) Update() {
	g.Should_Close = rl.WindowShouldClose()
	delta_time := rl.GetFrameTime()

	if !g.game_over && g.start {
		g.player_movement(delta_time)
		g.ball.Move(delta_time)

		g.ball.check_palette(&g.player_1, &g.player_2)

		g.check_if_score()
	}

	//* Start Game
	if !g.start && rl.IsKeyPressed(rl.KeySpace) {
		g.start = true
		fmt.Println("start GAme")
	}

	//* Restart Game
	if g.game_over && rl.IsKeyPressed(rl.KeyR) {
		g.game_over = false
		g.winner = ""
		g.p1_score = 0
		g.p2_score = 0
	}
}

func (g *Game) player_movement(dt float32) {
	p_movement := PLAYER_MOVEMENT * dt

	// * Player 1
	if rl.IsKeyDown(rl.KeyW) {
		g.player_1.Move(-p_movement)
	}

	if rl.IsKeyDown(rl.KeyS) {
		g.player_1.Move(p_movement)
	}

	//* Player 2
	if rl.IsKeyDown(rl.KeyI) {
		g.player_2.Move(-p_movement)
	}

	if rl.IsKeyDown(rl.KeyK) {
		g.player_2.Move(p_movement)
	}
}

func (g *Game) reset_ball(dir float32) {
	g.ball.x = float32(W_WIDTH) / 2
	g.ball.y = float32(W_HEIGHT) / 2
	g.ball.speed_x = BALL_SPEED_X * dir
	g.ball.speed_y = BALL_SPEED_y * dir
}

func (g *Game) check_if_score() {
	//* Check if ball is out letf screen
	if g.ball.x < -float32(g.ball.size) {
		g.reset_ball(1)

		g.p2_score += 1
	}

	//* Check if ball is out right screen
	if g.ball.x > float32(W_WIDTH) {
		g.reset_ball(-1)

		g.p1_score += 1
	}

	//* P1 Won
	if g.p1_score == g.max_score {
		g.game_over = true
		g.winner = "Player 1 Won"
	}

	//* P2 Won
	if g.p2_score == g.max_score {
		g.game_over = true
		g.winner = "Player 2 Won"
	}
}

func (g *Game) show_score() {
	rl.DrawText(
		fmt.Sprintf("%v", g.p1_score),
		200,
		10,
		20,
		rl.RayWhite,
	)

	rl.DrawText(
		fmt.Sprintf("%v", g.p2_score),
		600,
		10,
		20,
		rl.RayWhite,
	)
}

func draw_lines() {
	rl.DrawLine(W_WIDTH/2, 0, W_WIDTH/2, W_HEIGHT, rl.RayWhite)
}

func (g *Game) start_game() {
	//* LIGMA
	w := (W_WIDTH / 12)
	h := (W_HEIGHT / 12)

	rl.DrawRectangle(
		w,
		h,
		w*10,
		h*10,
		rl.RayWhite,
	)

	rl.DrawText(
		"Press 'Space' to start game",
		w*4,
		h*4,
		20,
		rl.Black,
	)

	rl.DrawText(
		"W and S for Left PLayer, I and K for Right Player",
		w*2,
		h*6,
		20,
		rl.Black,
	)
}

func (g *Game) restart_game() {
	//* LIGMA
	w := (W_WIDTH / 12)
	h := (W_HEIGHT / 12)

	rl.DrawRectangle(
		w,
		h,
		w*10,
		h*10,
		rl.RayWhite,
	)

	rl.DrawText(
		g.winner,
		w*5,
		h*4,
		20,
		rl.Black,
	)

	rl.DrawText(
		"Press 'R' to restart game",
		w*4,
		h*6,
		20,
		rl.Black,
	)
}

func (g *Game) Draw() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.GetColor(0x333333FF))

	if !g.game_over && g.start {
		g.ball.Render()

		g.player_1.Render()
		g.player_2.Render()

		draw_lines()

		g.show_score()
	} else {
		g.restart_game()
	}

	if !g.start {
		g.start_game()
	}

	rl.EndDrawing()
}

func (g *Game) Close() {
	rl.CloseWindow()
}

/**
*	———————————No bitches?———————————
*	⠀⣞⢽⢪⢣⢣⢣⢫⡺⡵⣝⡮⣗⢷⢽⢽⢽⣮⡷⡽⣜⣜⢮⢺⣜⢷⢽⢝⡽⣝
*	⠸⡸⠜⠕⠕⠁⢁⢇⢏⢽⢺⣪⡳⡝⣎⣏⢯⢞⡿⣟⣷⣳⢯⡷⣽⢽⢯⣳⣫⠇
*	⠀⠀⢀⢀⢄⢬⢪⡪⡎⣆⡈⠚⠜⠕⠇⠗⠝⢕⢯⢫⣞⣯⣿⣻⡽⣏⢗⣗⠏⠀
*	⠀⠪⡪⡪⣪⢪⢺⢸⢢⢓⢆⢤⢀⠀⠀⠀⠀⠈⢊⢞⡾⣿⡯⣏⢮⠷⠁⠀⠀
*	⠀⠀⠀⠈⠊⠆⡃⠕⢕⢇⢇⢇⢇⢇⢏⢎⢎⢆⢄⠀⢑⣽⣿⢝⠲⠉⠀⠀⠀⠀
*	⠀⠀⠀⠀⠀⡿⠂⠠⠀⡇⢇⠕⢈⣀⠀⠁⠡⠣⡣⡫⣂⣿⠯⢪⠰⠂⠀⠀⠀⠀
*	⠀⠀⠀⠀⡦⡙⡂⢀⢤⢣⠣⡈⣾⡃⠠⠄⠀⡄⢱⣌⣶⢏⢊⠂⠀⠀⠀⠀⠀⠀
*	⠀⠀⠀⠀⢝⡲⣜⡮⡏⢎⢌⢂⠙⠢⠐⢀⢘⢵⣽⣿⡿⠁⠁⠀⠀⠀⠀⠀⠀⠀
*	⠀⠀⠀⠀⠨⣺⡺⡕⡕⡱⡑⡆⡕⡅⡕⡜⡼⢽⡻⠏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
*	⠀⠀⠀⠀⣼⣳⣫⣾⣵⣗⡵⡱⡡⢣⢑⢕⢜⢕⡝⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
*	⠀⠀⠀⣴⣿⣾⣿⣿⣿⡿⡽⡑⢌⠪⡢⡣⣣⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
*	⠀⠀⠀⡟⡾⣿⢿⢿⢵⣽⣾⣼⣘⢸⢸⣞⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
*	⠀⠀⠀⠀⠁⠇⠡⠩⡫⢿⣝⡻⡮⣒⢽⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
*	—————————————————————————————
 */
