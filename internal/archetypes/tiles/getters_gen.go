package tiles

import "github.com/m110/moonshot-rts/internal/components"

func (t Tile) GetWorldSpace() *components.WorldSpace {
	return t.WorldSpace
}

func (t Tile) GetDrawable() *components.Drawable {
	return t.Drawable
}

func (t Tile) GetClickable() *components.Clickable {
	return t.Clickable
}

func (t Tile) GetCollider() *components.Collider {
	return t.Collider
}

func (t Tile) GetResourcesSource() *components.ResourcesSource {
	return t.ResourcesSource
}
