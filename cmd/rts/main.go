package main

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/m110/moonshot-rts/internal/assets/fonts"
	"github.com/m110/moonshot-rts/internal/assets/sprites"
	"github.com/m110/moonshot-rts/internal/game"
	"github.com/m110/moonshot-rts/internal/systems"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	err := fonts.LoadFonts()
	if err != nil {
		panic(err)
	}

	err = sprites.LoadSprites("assets/spritesheet.xml", "assets/ui.xml")
	if err != nil {
		panic(err)
	}

	tileMapConfig := systems.TilemapConfig{
		OffsetX:    0,
		OffsetY:    128,
		Width:      24,
		Height:     12,
		TileWidth:  64,
		TileHeight: 64,
	}

	config := systems.Config{
		TileMap: tileMapConfig,
		UI: systems.UIConfig{
			OffsetX: 0,
			OffsetY: 0,
			Width:   tileMapConfig.TotalWidth(),
			Height:  tileMapConfig.OffsetY,
			Font:    fonts.OpenSansRegular,
		},
	}

	g := game.NewGame(config)

	ebiten.SetWindowSize(config.WindowSize())

	err = ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
