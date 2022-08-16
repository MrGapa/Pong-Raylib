package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	size    int32
	x       float32
	y       float32
	speed_x float32
	speed_y float32
}

func (b *Ball) Move(delta_time float32) {
	b.x += b.speed_x * delta_time
	b.y += b.speed_y * delta_time

	b.check_bounds()
}

func (b *Ball) check_bounds() {
	if b.y < 0 || b.y > (float32(W_HEIGHT)-float32(b.size)) {
		b.speed_y *= -1
	}
}

//! Fix Math
func (b *Ball) check_palette(p1 *Palette, p2 *Palette) {
	if rl.CheckCollisionRecs(b.Get_Rec(), p1.Get_Rec()) {
		if b.speed_x < 0 {
			b.speed_x *= -1.05

			b.speed_y = (b.y - p1.y) / (float32(p1.height) / 2) * b.speed_x
		}
	}

	if rl.CheckCollisionRecs(b.Get_Rec(), p2.Get_Rec()) {
		if b.speed_x > 0 {
			b.speed_x *= -1.05

			b.speed_y = (b.y - p2.y) / (float32(p2.height) / 2) * -b.speed_x
		}
	}
}

func (b *Ball) Get_Rec() rl.Rectangle {
	return rl.NewRectangle(
		b.x,
		b.y,
		float32(b.size),
		float32(b.size),
	)
}

func (b *Ball) Render() {
	rl.DrawRectangle(
		int32(b.x),
		int32(b.y),
		b.size,
		b.size,
		rl.RayWhite,
	)
}
