/**
@author: Administrator
@Date: 2022-11-10-0010
@Note
**/
package main

import (
	"github.com/urfave/cli"
	"go_test/findPath"
	"math/rand"
	"os"
	"sync"
)

func main() {
	app := cli.NewApp()

	app.Name = "game server"
	app.Author = "ltx"
	app.Version = "0.0.1"
	app.Usage = "game server"

	app.Action = serve
	app.Run(os.Args)
}

func serve(c *cli.Context) error {

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()

	}()

	wg.Wait()
	return nil
}

func find() {
	start := &findPath.Cell{X: rand.Int63n(99) + 1, Y: rand.Int63n(99) + 1}
	end := &findPath.Cell{X: rand.Int63n(99) + 1, Y: rand.Int63n(99) + 1}
	findPath.QuadFindPath(*start, *end)
}
