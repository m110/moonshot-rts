package fonts

import (
	"io/ioutil"
	"os"

	"golang.org/x/image/font"

	"golang.org/x/image/font/opentype"
)

var (
	OpenSansRegular      font.Face
	OpenSansRegularSmall font.Face
)

func LoadFonts() error {
	fontFile, err := os.Open("assets/OpenSans-Regular.ttf")
	if err != nil {
		return err
	}
	fontBytes, err := ioutil.ReadAll(fontFile)
	if err != nil {
		return err
	}
	font, err := opentype.Parse(fontBytes)
	if err != nil {
		return err
	}
	face, err := opentype.NewFace(font, &opentype.FaceOptions{
		Size: 34,
		DPI:  70,
	})
	if err != nil {
		return err
	}

	smallFace, err := opentype.NewFace(font, &opentype.FaceOptions{
		Size: 14,
		DPI:  70,
	})
	if err != nil {
		return err
	}

	OpenSansRegular = face
	OpenSansRegularSmall = smallFace

	return nil
}
