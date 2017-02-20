package main

import "github.com/veandco/go-sdl2/sdl"

const (
	FIRST = '!'
	LAST  = '~'
)

type Text struct {
	w, h  int
	chars [LAST - FIRST]*sdl.Texture
}

func (t *Text) Init(gd GameData, font string, size int, color sdl.Color) {
	f := LoadFont(gd, font, size)
	for i := 0; i < LAST-FIRST; i++ {
		c := i + FIRST
		s, err := f.RenderUTF8_Solid(string(c), color)
		if err != nil {
			gd.Log.Fatal(err)
		}
		tex, err := gd.Render.CreateTextureFromSurface(s)
		if err != nil {
			gd.Log.Fatal(err)
		}
		s.Free()
		t.chars[i] = tex
	}
	w, h, _ := f.SizeUTF8(" ")
	t.w = w
	t.h = h
	f.Close()
}

func (t *Text) Render(gd GameData, x, y int, s string) {
	src := sdl.Rect{0, 0, 0, 0}
	dst := sdl.Rect{int32(x), int32(y), 0, 0}
	for _, c := range []byte(s) {
		if c >= FIRST && c <= LAST {
			tex := t.chars[c-FIRST]
			_, _, w, h, _ := tex.Query()
			dst.W = int32(w)
			dst.H = int32(h)
			src.W = int32(w)
			src.H = int32(h)
			gd.Render.Copy(tex, &src, &dst)
			dst.X += int32(w)
		} else if c == '\n' {
			dst.Y += int32(t.h)
			dst.X = int32(x)
		} else {
			dst.X += int32(t.w)
		}
	}
}

func (t *Text) Destroy() {
	for i, _ := range t.chars {
		t.chars[i].Destroy()
	}
}
