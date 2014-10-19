package nearest

import (
	"github.com/gansidui/geohash"
	"math"
)

const EARTH_RADIUS float64 = 6378.137 // 地球半径，单位为 km

type CoordNode struct {
	Latitude  float64 // 纬度
	Longitude float64 // 经度
	Geohash   string  // 该点的geohash值
}

type Nearest struct {
	// geohash算法的精度，默认为5
	precision int

	// 保存所有节点, key --> CoordNode
	allCoordNodes map[string]*CoordNode

	// 每个geohash值表示一个区域，一个区域包含多个节点
	// geohash指向一个保存若干个key的map: map[geohash]map[key]bool
	geohashMapkeys map[string]map[string]bool
}

func NewCoordNode(latitude, longitude float64, precision int) *CoordNode {
	ghash, _ := geohash.Encode(latitude, longitude, precision)
	return &CoordNode{
		Latitude:  latitude,
		Longitude: longitude,
		Geohash:   ghash,
	}
}

func NewNearest() *Nearest {
	return &Nearest{
		precision:      5,
		allCoordNodes:  make(map[string]*CoordNode),
		geohashMapkeys: make(map[string]map[string]bool),
	}
}

func (this *Nearest) GetPrecision() int {
	return this.precision
}

// 设置精度
func (this *Nearest) SetPrecision(precision int) {
	this.precision = precision
}

func (this *Nearest) GetAllCoordNodes() map[string]*CoordNode {
	return this.allCoordNodes
}

// 增加key的坐标节点
func (this *Nearest) AddCoordNode(key string, coordNode *CoordNode) {
	this.allCoordNodes[key] = coordNode

	if this.geohashMapkeys[coordNode.Geohash] == nil {
		this.geohashMapkeys[coordNode.Geohash] = make(map[string]bool)
	}
	this.geohashMapkeys[coordNode.Geohash][key] = true
}

func (this *Nearest) AddCoord(key string, latitude, longitude float64) {
	ghash, _ := geohash.Encode(latitude, longitude, this.precision)
	coordNode := &CoordNode{
		Latitude:  latitude,
		Longitude: longitude,
		Geohash:   ghash,
	}
	this.AddCoordNode(key, coordNode)
}

// 删除key的坐标节点
func (this *Nearest) DeleteCoordNode(key string) bool {
	if _, ok := this.allCoordNodes[key]; !ok {
		return false
	}

	ghash := this.allCoordNodes[key].Geohash
	delete(this.geohashMapkeys[ghash], key)
	delete(this.allCoordNodes, key)

	if len(this.geohashMapkeys[ghash]) == 0 {
		this.geohashMapkeys[ghash] = nil
	}

	return true
}

// 更新key的坐标节点
func (this *Nearest) UpdateCoordNode(key string, coordNode *CoordNode) bool {
	if !this.DeleteCoordNode(key) {
		return false
	}
	this.AddCoordNode(key, coordNode)
	return true
}

func (this *Nearest) UpdateCoord(key string, newLatitude, newLongitude float64) bool {
	if !this.DeleteCoordNode(key) {
		return false
	}
	this.AddCoord(key, newLatitude, newLongitude)
	return true
}

// 得到key的坐标节点
func (this *Nearest) GetCoordNode(key string) (*CoordNode, bool) {
	coordNode, ok := this.allCoordNodes[key]
	return coordNode, ok
}

// 查找key附近(九宫格内)的节点，返回他们的key
func (this *Nearest) QueryNearestSquareFromKey(key string) []string {
	if coordNode, ok := this.GetCoordNode(key); ok {
		return this.QueryNearestSquare(coordNode.Latitude, coordNode.Longitude)
	}
	return []string{}
}

// 查找(latitude, longitude)附近(九宫格内)的节点,返回它们的key
func (this *Nearest) QueryNearestSquare(latitude, longitude float64) []string {
	keys := make([]string, 0)
	neighbors := geohash.GetNeighbors(latitude, longitude, this.precision)
	for _, ghash := range neighbors {
		if this.geohashMapkeys[ghash] != nil {
			for key, _ := range this.geohashMapkeys[ghash] {
				keys = append(keys, key)
			}
		}
	}
	return keys
}

// 计算弧度
func rad(d float64) float64 {
	return d * math.Pi / 180.0
}

// 计算两坐标点的地球面距离, 单位为 km
func Distance(latitude1, longitude1, latitude2, longitude2 float64) float64 {
	radLat1 := rad(latitude1)
	radLat2 := rad(latitude2)
	a := radLat1 - radLat2
	b := rad(longitude1) - rad(longitude2)

	s := 2 * math.Asin(math.Sqrt(math.Pow(math.Sin(a/2), 2)+
		math.Cos(radLat1)*math.Cos(radLat2)*math.Pow(math.Sin(b/2), 2)))
	s = s * EARTH_RADIUS

	return s
}

// 计算两坐标点的地球面距离, 单位为 km
func DistanceCoordNode(coordNode1, coordNode2 *CoordNode) float64 {
	return Distance(coordNode1.Latitude, coordNode1.Longitude, coordNode2.Latitude, coordNode2.Longitude)
}
