package units

import "github.com/m110/moonshot-rts/internal/components"

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

func (u Unit) GetBuilder() *components.Builder {
	return u.Builder
}

func (u Unit) GetTimeActions() *components.TimeActions {
	return u.TimeActions
}
