package sprites

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/m110/moonshot-rts/internal/assets/fonts"
	"golang.org/x/image/colornames"
)

type Showcase struct {
	atlas Atlas

	width  int
	height int
}

func NewShowcase(atlas Atlas) Showcase {
	var maxWidth, maxHeight int

	for _, t := range atlas.tiles {
		if t.Width > maxWidth {
			maxWidth = t.Width
		}
		if t.Height > maxHeight {
			maxHeight = t.Height
		}
	}

	return Showcase{
		atlas:  atlas,
		width:  maxWidth * 2,
		height: int(float64(maxHeight) * 1.25),
	}
}

func (s Showcase) Draw(canvas *ebiten.Image) {
	i := 0
	for dy := 0; dy < canvas.Bounds().Max.Y-s.height; dy += s.height {
		for dx := 0; dx < canvas.Bounds().Max.X-s.width; dx += s.width {
			if i == len(s.atlas.Tiles()) {
				return
			}

			tile := s.atlas.Tiles()[i]

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(dx+s.width/2), float64(dy+s.height/2))
			canvas.DrawImage(tile.Image, op)

			text.Draw(
				canvas,
				tile.Name,
				fonts.OpenSansRegular,
				dx+s.width/4,
				dy+int(float64(s.height)*1.25),
				colornames.White,
			)

			i++
		}
	}
}
