package main

import (
	"path/filepath"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
	"github.com/veandco/go-sdl2/sdl_mixer"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

const (
	RES_DIR = "res/"
	IMG_DIR = "img/"
	SND_DIR = "snd/"
	FNT_DIR = "fnt/"
	MAP_DIR = "map/"
)

// TODO cache textures
func LoadTexture(gd GameData, res string) *sdl.Texture {
	gd.Log.Println("loading texture", res)
	return loadTexture(gd, filepath.Join(RES_DIR, IMG_DIR, res))
}

func loadTexture(gd GameData, file string) *sdl.Texture {
	tex, err := img.LoadTexture(gd.Render, file)
	if err != nil {
		gd.Log.Fatal(err)
	}
	return tex
}

func LoadSound(gd GameData, res string) *mix.Chunk {
	gd.Log.Println("loading sound", res)
	snd, err := mix.LoadWAV(filepath.Join(RES_DIR, SND_DIR, res))
	if err != nil {
		gd.Log.Fatal(err)
	}
	return snd
}

func LoadFont(gd GameData, res string, size int) *ttf.Font {
	gd.Log.Println("loading font", res)
	fnt, err := ttf.OpenFont(filepath.Join(RES_DIR, FNT_DIR, res), size)
	if err != nil {
		gd.Log.Fatal(err)
	}
	return fnt
}
