package atlas

import (
	"github.com/m110/rts/internal/components"
	"github.com/m110/rts/internal/engine"
)

var (
	Grass1 engine.Sprite
	Sand1  engine.Sprite
	Water1 engine.Sprite

	TreeBig   engine.Sprite
	TreeSmall engine.Sprite
	PineBig   engine.Sprite
	PineSmall engine.Sprite

	StoneSmall engine.Sprite
	StoneBig   engine.Sprite
	StoneTwo   engine.Sprite
	StoneThree engine.Sprite

	GoldThree engine.Sprite

	IronThree engine.Sprite

	Castle    engine.Sprite
	CastleTop engine.Sprite

	Barracks    engine.Sprite
	BarracksTop engine.Sprite

	Units = map[components.Team]map[components.Class]engine.Sprites{}
)

type TeamSprites struct {
	Blue  engine.Sprite
	Red   engine.Sprite
	Green engine.Sprite
	Gray  engine.Sprite
}

func LoadSprites(atlasDescriptorPath string) error {
	atlas, err := NewAtlas(atlasDescriptorPath)
	if err != nil {
		return err
	}

	Grass1 = engine.NewSpriteFromImage(atlas.ImageByName("grass_1"))
	Sand1 = engine.NewSpriteFromImage(atlas.ImageByName("sand_1"))
	Water1 = engine.NewSpriteFromImage(atlas.ImageByName("water_1"))

	loadSpritePivotBottom := func(name string) engine.Sprite {
		return engine.NewSpriteFromImageWithPivotType(atlas.ImageByName(name), engine.PivotBottom)
	}

	TreeBig = loadSpritePivotBottom("tree_big")
	TreeSmall = loadSpritePivotBottom("tree_small")
	PineBig = loadSpritePivotBottom("pine_big")
	PineSmall = loadSpritePivotBottom("pine_small")

	StoneSmall = loadSpritePivotBottom("stone_small")
	StoneBig = loadSpritePivotBottom("stone_big")
	StoneTwo = loadSpritePivotBottom("stone_two")
	StoneThree = loadSpritePivotBottom("stone_three")

	GoldThree = loadSpritePivotBottom("gold_three")
	IronThree = loadSpritePivotBottom("iron_three")

	Castle = loadSpritePivotBottom("castle")
	CastleTop = loadSpritePivotBottom("castle_top")

	Barracks = loadSpritePivotBottom("barracks")
	BarracksTop = loadSpritePivotBottom("barracks_top")

	Units = map[components.Team]map[components.Class]engine.Sprites{
		components.TeamBlue: {
			components.ClassWorker: {
				loadSpritePivotBottom("unit_blue_worker_man"),
				loadSpritePivotBottom("unit_blue_worker_woman"),
			},
			components.ClassWarrior: {loadSpritePivotBottom("unit_blue_warrior")},
			components.ClassKnight:  {loadSpritePivotBottom("unit_blue_knight")},
			components.ClassPriest:  {loadSpritePivotBottom("unit_blue_priest")},
			components.ClassKing:    {loadSpritePivotBottom("unit_blue_king")},
		},
		components.TeamRed: {
			components.ClassWorker: {
				loadSpritePivotBottom("unit_red_worker_man"),
				loadSpritePivotBottom("unit_red_worker_woman"),
			},
			components.ClassWarrior: {loadSpritePivotBottom("unit_red_warrior")},
			components.ClassKnight:  {loadSpritePivotBottom("unit_red_knight")},
			components.ClassPriest:  {loadSpritePivotBottom("unit_red_priest")},
			components.ClassKing:    {loadSpritePivotBottom("unit_red_king")},
		},
		components.TeamGreen: {
			components.ClassWorker: {
				loadSpritePivotBottom("unit_green_worker_man"),
				loadSpritePivotBottom("unit_green_worker_woman"),
			},
			components.ClassWarrior: {loadSpritePivotBottom("unit_green_warrior")},
			components.ClassKnight:  {loadSpritePivotBottom("unit_green_knight")},
			components.ClassPriest:  {loadSpritePivotBottom("unit_green_priest")},
			components.ClassKing:    {loadSpritePivotBottom("unit_green_king")},
		},
		components.TeamGray: {
			components.ClassWorker: {
				loadSpritePivotBottom("unit_gray_worker_man"),
				loadSpritePivotBottom("unit_gray_worker_woman"),
			},
			components.ClassWarrior: {loadSpritePivotBottom("unit_gray_warrior")},
			components.ClassKnight:  {loadSpritePivotBottom("unit_gray_knight")},
			components.ClassPriest:  {loadSpritePivotBottom("unit_gray_priest")},
			components.ClassKing:    {loadSpritePivotBottom("unit_gray_king")},
		},
	}

	return nil
}
