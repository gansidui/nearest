package nearest

import (
	// "fmt"
	"math/rand"
	"strconv"
	"testing"
)

// func TestNearest(t *testing.T) {
// 	var latitude, longitude float64 = 39.92324, 116.3906
// 	var precision int = 5

// 	curCoordNode := NewCoordNode(latitude, longitude, precision)

// 	near := NewNearest()
// 	near.SetPrecision(5)

// 	nodes := make([]*CoordNode, 6)
// 	nodes[0] = NewCoordNode(39.92424, 116.3906, precision)
// 	nodes[1] = NewCoordNode(39.91324, 116.3907, precision)
// 	nodes[2] = NewCoordNode(39.92364, 116.3806, precision)
// 	nodes[3] = NewCoordNode(39.92384, 116.3916, precision)
// 	nodes[4] = NewCoordNode(39.96324, 116.3106, precision)
// 	nodes[5] = NewCoordNode(39.92329, 116.3936, precision)

// 	for i, _ := range nodes {
// 		near.AddCoordNode("key["+strconv.Itoa(i)+"]", nodes[i])
// 	}

// 	// query 1
// 	fmt.Println("query 1")
// 	for k, v := range near.GetAllCoordNodes() {
// 		fmt.Printf("curCoordNode -- %v: %v\n", k, DistanceCoordNode(curCoordNode, v))
// 	}

// 	keys := near.QueryNearestSquare(latitude, longitude)
// 	for _, key := range keys {
// 		coordNode, _ := near.GetCoordNode(key)
// 		fmt.Printf("%v -- %v -- %v\n", key, coordNode, DistanceCoordNode(curCoordNode, coordNode))
// 	}
// 	fmt.Println()

// 	near.UpdateCoord("key[0]", 39.97424, 116.4906)
// 	near.UpdateCoord("key[4]", 39.91424, 116.3306)
// 	near.DeleteCoordNode("key[3]")

// 	// query 2
// 	fmt.Println("query 2")
// 	for k, v := range near.GetAllCoordNodes() {
// 		fmt.Printf("curCoordNode -- %v: %v\n", k, DistanceCoordNode(curCoordNode, v))
// 	}

// 	keys = near.QueryNearestSquare(latitude, longitude)
// 	for _, key := range keys {
// 		coordNode, _ := near.GetCoordNode(key)
// 		fmt.Printf("%v -- %v -- %v\n", key, coordNode, DistanceCoordNode(curCoordNode, coordNode))
// 	}
// 	fmt.Println()

// 	// query 3
// 	fmt.Println("query 3")
// 	for k, v := range near.GetAllCoordNodes() {
// 		fmt.Printf("nodes[4] -- %v: %v\n", k, DistanceCoordNode(nodes[4], v))
// 	}

// 	keys = near.QueryNearestSquareFromKey("key[4]")
// 	for _, key := range keys {
// 		coordNode, _ := near.GetCoordNode(key)
// 		fmt.Printf("%v -- %v -- %v\n", key, nodes[4], DistanceCoordNode(nodes[4], coordNode))
// 	}
// 	fmt.Println()
// }

// func TestDistance(t *testing.T) {
// 	latitude1, longitude1, latitude2, longitude2 := 39.92324, 116.3906, 39.92324, 116.3907
// 	fmt.Println(Distance(latitude1, longitude1, latitude2, longitude2))

// 	latitude1, longitude1, latitude2, longitude2 = 39.941, 116.45, 39.94, 116.451
// 	fmt.Println(Distance(latitude1, longitude1, latitude2, longitude2))

// 	latitude1, longitude1, latitude2, longitude2 = 39.96, 116.45, 39.94, 116.40
// 	fmt.Println(Distance(latitude1, longitude1, latitude2, longitude2))

// 	latitude1, longitude1, latitude2, longitude2 = 39.96, 116.45, 39.94, 117.30
// 	fmt.Println(Distance(latitude1, longitude1, latitude2, longitude2))

// 	latitude1, longitude1, latitude2, longitude2 = 39.26, 115.25, 41.04, 117.30
// 	fmt.Println(Distance(latitude1, longitude1, latitude2, longitude2))
// }

func BenchmarkDelete(b *testing.B) {
	near := NewNearest()
	var latitude, longitude float64 = 39.92324, 116.3906

	for i := 0; i < b.N; i++ {
		near.AddCoord("key"+strconv.Itoa(i), latitude, longitude)
		near.DeleteCoordNode("key" + strconv.Itoa(i))
	}
}

func BenchmarkUpdate(b *testing.B) {
	near := NewNearest()
	var latitude, longitude float64 = 39.92324, 116.3906
	var newLatitude, newLongitude float64 = 49.92324, 120.3906

	for i := 0; i < b.N; i++ {
		near.AddCoord("key"+strconv.Itoa(i), latitude, longitude)
		near.UpdateCoord("key"+strconv.Itoa(i), newLatitude, newLongitude)
	}
}

func BenchmarkQuery(b *testing.B) {
	b.StopTimer()
	near := NewNearest()
	var latitude, longitude float64 = 39.92324, 116.3906
	var lat, lon float64

	for i := 0; i < 1000000; i++ {
		lat = float64(rand.Intn(100000)-50000) / 100000
		lon = float64(rand.Intn(10000)-5000) / 10000
		near.AddCoord("key"+strconv.Itoa(i), latitude+lat, longitude+lon)
	}

	b.StartTimer()

	for i := 0; i < 100; i++ {
		lat = float64(rand.Intn(100000)-50000) / 100000
		lon = float64(rand.Intn(10000)-5000) / 10000
		near.QueryNearestSquare(latitude+lat, longitude+lon)
	}
}
