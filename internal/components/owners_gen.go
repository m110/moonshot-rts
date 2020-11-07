package components

type BuilderOwner interface {
	GetBuilder() *Builder
}

type ButtonOwner interface {
	GetButton() *Button
}

type ClickableOwner interface {
	GetClickable() *Clickable
}

type DrawableOwner interface {
	GetDrawable() *Drawable
}

type MovableOwner interface {
	GetMovable() *Movable
}

type ProgressBarOwner interface {
	GetProgressBar() *ProgressBar
}

type SelectableOwner interface {
	GetSelectable() *Selectable
}

type SizeOwner interface {
	GetSize() *Size
}

type TimeActionsOwner interface {
	GetTimeActions() *TimeActions
}

type UnitSpawnerOwner interface {
	GetUnitSpawner() *UnitSpawner
}

type WorldSpaceOwner interface {
	GetWorldSpace() *WorldSpace
}
