package systems

import (
	"github.com/m110/moonshot-rts/internal/objects"
	"github.com/m110/moonshot-rts/internal/tiles"
	"github.com/m110/moonshot-rts/internal/units"
)

type spawner interface {
	SpawnUnit(unit units.Unit)
	SpawnTile(tile tiles.Tile)
	SpawnBuilding(building objects.Building)
	SpawnObject(object objects.Object)
	SpawnDrawingEntity(entity DrawingEntity)
}
