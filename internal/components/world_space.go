package components

import "github.com/m110/moonshot-rts/internal/engine"

type WorldSpace struct {
	worldPosition engine.Vector
	localPosition engine.Vector

	Parent   *WorldSpace
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
		w.worldPosition = w.worldPosition.Add(parent.LocalPosition())
		parent = parent.Parent
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

func (w *WorldSpace) AddChild(child WorldSpaceOwner) {
	w.Children = append(w.Children, child)

	child.GetWorldSpace().Parent = w
	child.GetWorldSpace().updateWorldPositionFromParents()
}
