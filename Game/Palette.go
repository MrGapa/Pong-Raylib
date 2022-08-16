package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Palette struct {
	width  int32
	height int32

	x float32
	y float32
}

func (p *Palette) Render() {
	rl.DrawRectangle(int32(p.x), int32(p.y), p.width, p.height, rl.RayWhite)
}

func (p *Palette) Move(y float32) {
	pos_y := p.y + y

	if !p.check_border() {
		p.y = pos_y
	} else if p.y <= 0 {
		p.y = 10
	} else if p.y >= float32(W_HEIGHT)-float32(p.height) {
		p.y = (float32(W_HEIGHT) - float32(p.height)) - 10
	}
}

func (p *Palette) check_border() bool {
	if p.y <= 0 || p.y >= float32(W_HEIGHT)-float32(p.height) {
		return true
	}

	return false
}

func (p *Palette) Get_Rec() rl.Rectangle {
	return rl.NewRectangle(
		p.x,
		p.y,
		float32(p.width),
		float32(p.height),
	)
}
