package archetypes

import "github.com/m110/moonshot-rts/internal/components"

func (b Building) GetSelectable() *components.Selectable {
	return b.Selectable
}

func (b Building) GetClickable() *components.Clickable {
	return b.Clickable
}

func (b Building) GetCollider() *components.Collider {
	return b.Collider
}

func (b Building) GetUnitSpawner() *components.UnitSpawner {
	return b.UnitSpawner
}

func (b Building) GetTimeActions() *components.TimeActions {
	return b.TimeActions
}

func (o Object) GetWorldSpace() *components.WorldSpace {
	return o.WorldSpace
}

func (o Object) GetDrawable() *components.Drawable {
	return o.Drawable
}

func (o Overlay) GetWorldSpace() *components.WorldSpace {
	return o.WorldSpace
}

func (o Overlay) GetDrawable() *components.Drawable {
	return o.Drawable
}

func (o Overlay) GetSize() *components.Size {
	return o.Size
}

func (p Panel) GetClickable() *components.Clickable {
	return p.Clickable
}

func (p PanelButton) GetClickable() *components.Clickable {
	return p.Clickable
}

func (p PanelButton) GetButton() *components.Button {
	return p.Button
}

func (p ProgressBar) GetProgressBar() *components.ProgressBar {
	return p.ProgressBar
}

func (u Unit) GetWorldSpace() *components.WorldSpace {
	return u.WorldSpace
}

func (u Unit) GetCitizen() *components.Citizen {
	return u.Citizen
}

func (u Unit) GetDrawable() *components.Drawable {
	return u.Drawable
}

func (u Unit) GetMovable() *components.Movable {
	return u.Movable
}

func (u Unit) GetSelectable() *components.Selectable {
	return u.Selectable
}

func (u Unit) GetClickable() *components.Clickable {
	return u.Clickable
}

func (u Unit) GetCollider() *components.Collider {
	return u.Collider
}

func (u Unit) GetAreaOccupant() *components.AreaOccupant {
	return u.AreaOccupant
}

func (w Worker) GetBuilder() *components.Builder {
	return w.Builder
}

func (w Worker) GetResourcesCollector() *components.ResourcesCollector {
	return w.ResourcesCollector
}
