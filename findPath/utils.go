package findPath

import "math"

type Node struct {
	block  *Block
	gCost  float64
	hCost  float64
	parent *Node
}

func (n Node) GetF() float64 {
	return n.gCost + n.hCost
}

/**
格子
*/
type Cell struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
}

type Block struct {
	Cell     *Cell
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

func GetBlock(cell *Cell, tagIds []int16) *Block {
	return &Block{Id: ComTwoInt16(cell.X, cell.Y), Cell: cell, Walkable: BlockIsWalkable(tagIds)}
}

func BlockIsWalkable(tagIds []int16) bool {
	if tagIds == nil || len(tagIds) <= 0 {
		return false
	}
	for _, tag := range tagIds {
		if (tag >= 32 && tag <= 35) || (tag >= 51 && tag <= 53) {
			return true
		}
	}
	return false
}

func ComTwoInt16(x int64, y int64) int64 {
	xt := x << 16
	xt = xt | (y & 0xFFFF)
	return xt
}
