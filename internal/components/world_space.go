package components

import "github.com/m110/rts/internal/engine"

type WorldSpaceOwner interface {
	GetWorldSpace() *WorldSpace
}

// TODO Maybe not the best name, as it's local position for children?
type WorldSpace struct {
	worldPosition engine.Vector
	localPosition engine.Vector

	Parent   WorldSpaceOwner
	Children []WorldSpaceOwner
}

func (w *WorldSpace) SetInWorld(x float64, y float64) {
	current := w.worldPosition
	w.worldPosition.Set(x, y)
	w.localPosition.Translate(x-current.X, y-current.Y)

	for _, child := range w.Children {
		child.GetWorldSpace().updateWorldPositionFromParents()
	}
}

func (w *WorldSpace) SetLocal(x float64, y float64) {
	current := w.localPosition
	w.localPosition.Set(x, y)
	w.worldPosition.Translate(x-current.X, y-current.Y)

	for _, child := range w.Children {
		child.GetWorldSpace().updateWorldPositionFromParents()
	}
}

func (w *WorldSpace) Translate(x float64, y float64) {
	if w.HasParent() {
		w.localPosition.Translate(x, y)
		w.worldPosition.Translate(x, y)
	} else {
		w.worldPosition.Translate(x, y)
	}

	for _, child := range w.Children {
		child.GetWorldSpace().updateWorldPositionFromParents()
	}
}

func (w WorldSpace) HasParent() bool {
	return w.Parent != nil
}

func (w *WorldSpace) updateWorldPositionFromParents() {
	if w.Parent == nil {
		return
	}

	w.worldPosition = w.localPosition

	parent := w.Parent
	for parent != nil {
		w.worldPosition = w.worldPosition.Add(parent.GetWorldSpace().LocalPosition())
		parent = parent.GetWorldSpace().Parent
	}

	for _, child := range w.Children {
		child.GetWorldSpace().updateWorldPositionFromParents()
	}
}

func (w WorldSpace) LocalPosition() engine.Vector {
	if w.HasParent() {
		return w.localPosition
	}

	return w.worldPosition
}

func (w WorldSpace) WorldPosition() engine.Vector {
	return w.worldPosition
}

// TODO Oh this owner is so bad
func (w *WorldSpace) AddChild(owner WorldSpaceOwner, child WorldSpaceOwner) {
	w.Children = append(w.Children, child)

	// TODO This seems hacky
	child.GetWorldSpace().Parent = owner
	child.GetWorldSpace().updateWorldPositionFromParents()
}
