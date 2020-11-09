package sprites

import (
	"encoding/xml"
	"image"
	_ "image/png"
	"io/ioutil"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Atlas struct {
	tiles    []Tile
	tilesMap map[string]Tile
}

type Tile struct {
	Name string

	X      int
	Y      int
	Width  int
	Height int

	Image *ebiten.Image
}

type textureAtlas struct {
	TextureAtlas xml.Name     `xml:"TextureAtlas"`
	ImagePath    string       `xml:"imagePath,attr"`
	SubTextures  []subTexture `xml:"SubTexture"`
}

type subTexture struct {
	Name   string `xml:"name,attr"`
	X      int    `xml:"x,attr"`
	Y      int    `xml:"y,attr"`
	Width  int    `xml:"width,attr"`
	Height int    `xml:"height,attr"`
}

func NewAtlas(descriptorPath string) (Atlas, error) {
	descriptorFile, err := os.Open(descriptorPath)
	if err != nil {
		return Atlas{}, err
	}
	descriptorContent, err := ioutil.ReadAll(descriptorFile)
	if err != nil {
		return Atlas{}, err
	}

	var texAtlas textureAtlas
	err = xml.Unmarshal(descriptorContent, &texAtlas)
	if err != nil {
		return Atlas{}, err
	}

	tilesFile, err := os.Open(texAtlas.ImagePath)
	if err != nil {
		return Atlas{}, err
	}
	imgContent, _, err := image.Decode(tilesFile)
	if err != nil {
		return Atlas{}, err
	}
	tilesImage := ebiten.NewImageFromImage(imgContent)

	var tiles []Tile
	tilesMap := map[string]Tile{}

	for _, st := range texAtlas.SubTextures {
		subImage := tilesImage.SubImage(image.Rect(st.X, st.Y, st.X+st.Width, st.Y+st.Height)).(*ebiten.Image)
		tile := Tile{
			Name:   st.Name,
			X:      st.X,
			Y:      st.Y,
			Width:  st.Width,
			Height: st.Height,
			Image:  subImage,
		}

		tiles = append(tiles, tile)
		tilesMap[st.Name] = tile
	}

	return Atlas{
		tiles:    tiles,
		tilesMap: tilesMap,
	}, nil
}

func (a Atlas) Tiles() []Tile {
	return a.tiles
}

func (a Atlas) TileByName(name string) Tile {
	tile, ok := a.tilesMap[name]
	if !ok {
		panic("tile not found: " + name)
	}
	return tile
}

func (a Atlas) ImageByName(name string) *ebiten.Image {
	return a.TileByName(name).Image
}
