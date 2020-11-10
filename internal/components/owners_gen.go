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

type ColliderOwner interface {
	GetCollider() *Collider
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

type ResourcesSourceOwner interface {
	GetResourcesSource() *ResourcesSource
}

type ResourcesCollectorOwner interface {
	GetResourcesCollector() *ResourcesCollector
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
