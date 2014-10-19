##nearest 

k-NN algorithm

基于geohash算法，给出一个坐标点，查找附近的其他点

~~~ go
package main

import (
	"fmt"
	"github.com/gansidui/nearest"
)

func main() {
	near := nearest.NewNearest()
	near.SetPrecision(5)

	near.AddCoord("A", 40.92424, 116.3906)
	near.AddCoord("B", 39.93224, 116.3927)
	near.AddCoord("C", 39.92484, 116.3916)
	near.AddCoord("D", 39.92494, 116.3923)
	near.AddCoord("E", 39.92220, 116.3915)
	near.AddCoord("F", 39.92424, 117.3906)

	keys := near.QueryNearestSquareFromKey("C")

	coordNode1, ok := near.GetCoordNode("C")
	if !ok {
		return
	}

	for _, key := range keys {
		coordNode2, _ := near.GetCoordNode(key)
		fmt.Println(key, nearest.DistanceCoordNode(coordNode1, coordNode2))
	}
}

~~~



##LICENSE

MIT