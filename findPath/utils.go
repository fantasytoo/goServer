package findPath

import "math"

/**
格子
*/
type Cell struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
}

type Block struct {
	Cell     Cell
	Id       int64
	Walkable bool
}

type HexGrid struct {
	TagIDs []int16
	ID     int64
}

type MapTerrainData struct {
	Name        string
	Description string
	GridType    int8
	MapTags     []int8
	Maps        []string
	Size        Cell
	Grids       []HexGrid
}

func GetCoord(index int64) *Cell {
	if index == 0 {
		return &Cell{X: 0, Y: 0}
	}
	sqrtX := math.Sqrt(float64(index))
	k := int64((sqrtX + 1) / 2)
	if sqrtX > (float64(2 * k)) {
		xAy := index - (4*k*k + 1) - 2*k
		if xAy < 0 {
			z := -1 * k
			x := xAy - z
			return &Cell{X: x, Y: z}
		} else {
			x := k
			z := xAy - x
			return &Cell{X: x, Y: z}
		}
	} else {
		xAy := (4*k*k + 1) - index - 2*k
		if xAy < 0 {
			x := -k
			z := xAy - x
			return &Cell{X: x, Y: z}
		} else {
			z := k
			x := xAy - z
			return &Cell{X: x, Y: z}
		}
	}
}
