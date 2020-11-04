package systems

type Config struct {
	TileMap TilemapConfig
	UI      UIConfig
}

func (c Config) WindowSize() (width int, height int) {
	width = c.TileMap.TotalWidth()
	height = c.UI.Height + c.TileMap.TotalHeight()
	return width, height
}
