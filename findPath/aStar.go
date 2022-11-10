/**
@author: Administrator
@Date: 2022-11-10-0010
@Note
**/
package findPath

import (
	"encoding/json"
	"fmt"
	"go_test/log"
	"os"
)

var mapData map[int64]*Block

func init() {
	content, err := os.ReadFile("config/map_1.json")
	if err != nil {
		log.AppLogger.Error("file read error.")
		return
	}
	var mapData MapTerrainData
	err = json.Unmarshal(content, &mapData)
	if err != nil {
		log.AppLogger.Error("data decode json error.")
		return
	}
	for i := 0; i < len(mapData.Grids); i++ {

	}
}

func QuadFindPath() {
	fmt.Println("findPath........")
}
