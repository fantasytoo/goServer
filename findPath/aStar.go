/**
@author: Administrator
@Date: 2022-11-10-0010
@Note
**/
package findPath

import (
	"container/heap"
	"encoding/json"
	"fmt"
	"go_test/log"
	"math"
	"os"
	"time"
)

var mapData map[int64]*Block

var DIRECT_VALUE float64 = 1

func init() {
	content, err := os.ReadFile("./findPath/config/map_1.json")
	if err != nil {
		log.AppLogger.Error("file read error.")
		return
	}
	var mapDataConfig MapTerrainData
	err = json.Unmarshal(content, &mapDataConfig)
	if err != nil {
		log.AppLogger.Error("data decode json error.")
		return
	}
	mapData = make(map[int64]*Block)
	for i := 0; i < len(mapDataConfig.Grids); i++ {
		cell := GetCoord(mapDataConfig.Grids[i].ID)
		block := GetBlock(cell, mapDataConfig.Grids[i].TagIDs)
		mapData[block.Id] = block
	}
}

func QuadFindPath(start Cell, end Cell) {
	startTime := time.Now().UnixNano() / 1e6
	//fmt.Println("findPath........")
	startBlock := mapData[ComTwoInt16(start.X, start.Y)]
	endBlock := mapData[ComTwoInt16(end.X, end.Y)]
	if startBlock == nil || endBlock == nil || startBlock.Id == endBlock.Id {
		return
	}
	endNode := findParent(startBlock, endBlock)
	if endNode != nil {
		list := getNavigitionPath(endNode)
		fmt.Printf("寻路成功,路径长度=%d,耗时:%d\n", len(list), time.Now().UnixNano()/1e6-startTime)
		return
	}
	fmt.Println("寻路失败")
}

func getNavigitionPath(node *Node) []*Node {
	list := make([]*Node, 0)
	list = append(list, node)
	for {
		if node.parent == nil {
			break
		}
		node = node.parent
		//fmt.Printf("x=%d,y=%d\n", node.block.Cell.X, node.block.Cell.Y)
		list = append(list, node)
	}
	return list
}

func findParent(start *Block, end *Block) *Node {
	openMap := make(map[int64]*Node)
	closeMap := make(map[int64]*Node)
	nodes := make(map[int64]*Node)
	startNode := getNode(nodes, start.Cell.X, start.Cell.Y)
	endNode := getNode(nodes, end.Cell.X, end.Cell.Y)

	minHeap := make(MinHeap, 0)
	heap.Init(&minHeap)
	heap.Push(&minHeap, startNode)
	// 加入起点
	openMap[startNode.block.Id] = startNode
	for {
		if len(openMap) <= 0 || len(minHeap) <= 0 {
			break
		}
		node := heap.Pop(&minHeap).(*Node)
		delete(openMap, node.block.Id)
		closeMap[node.block.Id] = node
		addNodes := make([]*Node, 0)
		addNodes = findNeighborNodes(nodes, node, addNodes)
		for _, s := range addNodes {
			_, ok := closeMap[s.block.Id]
			if !ok {
				_, ok := openMap[s.block.Id]
				if ok {
					foundPoint(node, s)
				} else {
					notFoundPoint(node, endNode, s)
					openMap[s.block.Id] = s
					heap.Push(&minHeap, s)
				}
			}
		}
		eNode, isEndNode := openMap[endNode.block.Id]
		if isEndNode {
			return eNode
		}
	}
	return nil
}

func foundPoint(start *Node, end *Node) {
	g := calcG(start, end)
	if g < end.gCost {
		end.parent = start
		end.gCost = g
	}
}

func notFoundPoint(start *Node, end *Node, current *Node) {
	current.parent = start
	current.gCost = calcG(start, current)
	current.hCost = calcH(current, end)
}

func calcG(start *Node, end *Node) float64 {
	return start.gCost + DIRECT_VALUE
}

func calcH(current *Node, end *Node) float64 {
	step := calcDistance(current, end)
	return step * getHRatio(step)
}

func getHRatio(step float64) float64 {
	if step > 1000 {
		return 10
	} else if step > 500 {
		return 5
	} else if step > 300 {
		return 3
	} else if step > 100 {
		return 2
	} else {
		return 1
	}
}

func calcDistance(current *Node, end *Node) float64 {
	dx := math.Abs((float64)(current.block.Cell.X - end.block.Cell.X))
	dy := math.Abs((float64)(current.block.Cell.Y - end.block.Cell.Y))
	return math.Sqrt(dx*dx + dy*dy)
}

func findNeighborNodes(nodes map[int64]*Node, node *Node, list []*Node) []*Node {
	list = findNode(nodes, list, node, 0, 1)
	list = findNode(nodes, list, node, 0, -1)
	list = findNode(nodes, list, node, -1, 0)
	list = findNode(nodes, list, node, 1, 0)
	return list
}

func findNode(nodes map[int64]*Node, list []*Node, node *Node, x int64, y int64) []*Node {
	addNode := getNode(nodes, node.block.Cell.X+x, node.block.Cell.Y+y)
	if addNode != nil && !addNode.block.Walkable {
		list = append(list, addNode)
	}
	return list
}

func getNode(maps map[int64]*Node, x int64, y int64) *Node {
	id := ComTwoInt16(x, y)
	value, ok := maps[id]
	if ok {
		return value
	} else {
		block, ok1 := mapData[id]
		if ok1 {
			node := &Node{block: block}
			maps[id] = node
			return node
		} else {
			return nil
		}
	}
}

type MinHeap []*Node

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].GetF() < h[j].GetF()
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*Node))
}
func (h *MinHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}
