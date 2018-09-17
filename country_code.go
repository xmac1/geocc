package geocc

import (
	"io/ioutil"
	"os"

	"fmt"

	"github.com/json-iterator/go"
	"github.com/xmac1/example/kdtree"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Data struct {
	Countries []Country `json:"countries"`
}

type Country struct {
	Name   string `json:"name"`
	Geo    Region `json:"geo"`
	Bounds Bounds `json:"bounds"`
}

var regionMap = make(map[int]*Country)

type PolygonCountry struct {
	Name  string
	Areas []Area
}

type Area struct {
	Region Region // Region is a Polygon of a country
	Bounds Bounds // Bounds mean AABB in legacy quadtree
}

type Region [][]float32
type Point []float32 // Point is Long/Lat multiply 1e5 for efficient

type CountryPoint struct {
	Point
	Country string
}

var tree = &Quadtree{MaxLevels: 2, MaxObjects: 5}

// Initialize Country Code Map
func InitCountryMap(filename string) (err error) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	bts, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	data := &Data{}

	if err = json.Unmarshal(bts, data); err != nil {
		panic(err)
	}

	i := 1
	for _, country := range data.Countries {
		country.Bounds.ID = int32(i)
		tree.Insert(country.Bounds)
		regionMap[i] = &Country{
			Name:   country.Name,
			Geo:    country.Geo,
			Bounds: country.Bounds,
		}
		i++
	}
	return
}

// Geo Point to Country Code
func Geo2Country(point Point) string {
	bs := tree.RetrieveIntersections(point[0], point[1])

	for _, b := range bs {
		region := regionMap[int(b)]
		if pnpoly(region.Geo, point) {
			return region.Name
		}

	}
	return "LostCity"
}

// check if  point inside polygon
func pnpoly(polygon Region, point []float32) bool {
	//fmt.Println(polygon)
	i, j := 0, 0
	c := false
	size := len(polygon)
	for j = size - 1; i < size; i++ {
		if ((polygon[i][1] > point[1]) != (polygon[j][1] > point[1])) &&
			(point[0] < (polygon[j][0]-polygon[i][0])*(point[1]-polygon[i][1])/(polygon[j][1]-polygon[i][1])+polygon[i][0]) {
			c = !c
		}

		j = i
	}
	return c
}

var kdTree *kdtree.KDNode

//
func InitKDTree(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	bts, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	data := &Data{}

	if err = json.Unmarshal(bts, data); err != nil {
		panic(err)
	}

	nodeList := make([]*kdtree.KDNode, 0, 10000)

	for _, country := range data.Countries {
		for _, point := range country.Geo {
			node := &kdtree.KDNode{
				X:       point[0],
				Y:       point[1],
				Country: country.Name,
			}
			nodeList = append(nodeList, node)
		}
	}

	fmt.Println(len(nodeList))

	kdTree = kdtree.BuildKDTree(nodeList, kdtree.AxisX)
	return
}

func SearchCountry(long, lat float32) string {
	node := kdTree.SearchDepthFirst(long, lat)
	fmt.Println(node)
	return node.Country
}
