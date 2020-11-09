package sprites

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
)

var (
	Grass1 engine.Sprite
	Sand1  engine.Sprite
	Water1 engine.Sprite

	FieldSmallEmpty engine.Sprite
	FieldSmallWheat engine.Sprite
	FieldBigEmpty   engine.Sprite
	FieldBigWheat   engine.Sprite

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

	Chapel    engine.Sprite
	ChapelTop engine.Sprite

	Tower    engine.Sprite
	TowerTop engine.Sprite

	Forge engine.Sprite

	Units = map[components.Team]map[components.Class]engine.Sprites{}

	PanelBrown engine.Sprite

	ButtonBeige        engine.Sprite
	ButtonBeigePressed engine.Sprite

	ButtonBrown        engine.Sprite
	ButtonBrownPressed engine.Sprite

	Hammer engine.Sprite

	BarBackHorizontalLeft  engine.Sprite
	BarBackHorizontalMid   engine.Sprite
	BarBackHorizontalRight engine.Sprite

	BarGreenHorizontalLeft  engine.Sprite
	BarGreenHorizontalMid   engine.Sprite
	BarGreenHorizontalRight engine.Sprite
)

type TeamSprites struct {
	Blue  engine.Sprite
	Red   engine.Sprite
	Green engine.Sprite
	Gray  engine.Sprite
}

func LoadSprites(rtsPath string, uiPath string) error {
	rtsAtlas, err := NewAtlas(rtsPath)
	if err != nil {
		return err
	}

	uiAtlas, err := NewAtlas(uiPath)
	if err != nil {
		return err
	}

	Grass1 = engine.NewSpriteFromImage(rtsAtlas.ImageByName("grass_1"))
	Sand1 = engine.NewSpriteFromImage(rtsAtlas.ImageByName("sand_1"))
	Water1 = engine.NewSpriteFromImage(rtsAtlas.ImageByName("water_1"))

	FieldSmallEmpty = engine.NewSpriteFromImage(rtsAtlas.ImageByName("grass_field_small_empty"))
	FieldSmallWheat = engine.NewSpriteFromImage(rtsAtlas.ImageByName("grass_field_small_wheat"))
	FieldBigEmpty = engine.NewSpriteFromImage(rtsAtlas.ImageByName("grass_field_big_empty"))
	FieldBigWheat = engine.NewSpriteFromImage(rtsAtlas.ImageByName("grass_field_big_wheat"))

	loadSpritePivotBottom := func(name string) engine.Sprite {
		return engine.NewSpriteFromImageWithPivotType(rtsAtlas.ImageByName(name), engine.PivotBottom)
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

	Chapel = loadSpritePivotBottom("chapel")
	ChapelTop = loadSpritePivotBottom("chapel_top")

	Tower = loadSpritePivotBottom("tower")
	TowerTop = loadSpritePivotBottom("tower_top")

	Forge = loadSpritePivotBottom("forge")

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

	PanelBrown = engine.NewSpriteFromImage(uiAtlas.ImageByName("panel_brown"))

	ButtonBeige = engine.NewSpriteFromImage(uiAtlas.ImageByName("buttonSquare_beige"))
	ButtonBeigePressed = engine.NewSpriteFromImage(uiAtlas.ImageByName("buttonSquare_beige_pressed"))

	ButtonBrown = engine.NewSpriteFromImage(uiAtlas.ImageByName("buttonSquare_brown"))
	ButtonBrownPressed = engine.NewSpriteFromImage(uiAtlas.ImageByName("buttonSquare_brown_pressed"))

	hammerImg, _, err := ebitenutil.NewImageFromFile("assets/hammer.png")
	if err != nil {
		return err
	}
	Hammer = engine.NewSpriteFromImage(hammerImg)

	BarBackHorizontalLeft = engine.NewSpriteFromImage(uiAtlas.ImageByName("barBack_horizontalLeft"))
	BarBackHorizontalMid = engine.NewSpriteFromImage(uiAtlas.ImageByName("barBack_horizontalMid"))
	BarBackHorizontalRight = engine.NewSpriteFromImage(uiAtlas.ImageByName("barBack_horizontalRight"))

	BarGreenHorizontalLeft = engine.NewSpriteFromImage(uiAtlas.ImageByName("barGreen_horizontalLeft"))
	BarGreenHorizontalMid = engine.NewSpriteFromImage(uiAtlas.ImageByName("barGreen_horizontalMid"))
	BarGreenHorizontalRight = engine.NewSpriteFromImage(uiAtlas.ImageByName("barGreen_horizontalRight"))

	return nil
}
