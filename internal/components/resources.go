package components

type Resources struct {
	Food  int
	Wood  int
	Stone int
	Gold  int
	Iron  int
}

func (r *Resources) Update(update Resources) {
	r.Food = nonZero(r.Food + update.Food)
	r.Wood = nonZero(r.Wood + update.Wood)
	r.Stone = nonZero(r.Stone + update.Stone)
	r.Gold = nonZero(r.Gold + update.Gold)
	r.Iron = nonZero(r.Iron + update.Iron)
}

func nonZero(v int) int {
	if v < 0 {
		return 0
	}
	return v
}

type ResourcesSource struct {
	Resources Resources
}

type ResourcesCollector struct {
	CurrentResources Resources
	Collecting       bool
}
