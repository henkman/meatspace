package main

import "github.com/veandco/go-sdl2/sdl"

type Game struct {
	Text Text
}

func (g *Game) Init(gd GameData) {
	g.Text.Init(gd, "Inconsolata-Regular.ttf", 30, sdl.Color{0, 0xff, 0, 0xff})
}

func (g *Game) Run(gd GameData) Scene {
	for {
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch t := e.(type) {
			case *sdl.QuitEvent:
				return nil
			case *sdl.KeyUpEvent:
				switch t.Keysym.Scancode {
				case sdl.SCANCODE_ESCAPE:
					return nil
				}
			}
		}
		gd.Render.SetDrawColor(0x00, 0x00, 0x00, 0xFF)
		gd.Render.Clear()
		{
			g.Text.Render(gd, 10, 10, `so let's try
some longer text with a few special
characters /()#'!$`)
		}
		gd.Render.Present()
	}
}
