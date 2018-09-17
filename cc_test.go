package geocc_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/xmac1/geocc"
)

func BenchmarkGeo2Country(b *testing.B) {

	for i := 0; i < b.N; i++ {

	}

}

func TestGeo2Country(t *testing.T) {
	if err := geocc.InitCountryMap("D:/countries.json"); err != nil {
		panic(err)
	}

	fmt.Println(geocc.Geo2Country([]float32{113.93474, 22.525246}))

	start := time.Now()
	n := 10000
	for i := 0; i < n; i++ {
		geocc.Geo2Country([]float32{113.93474, 22.525246})
	}
	dur := time.Since(start)

	fmt.Println(dur.String())

	fmt.Println("per second, ", int64(n*1000)/(dur.Nanoseconds()/1e6))
}

func TestInitKDTree(t *testing.T) {
	geocc.InitKDTree("D:/countries.json")

	cn := geocc.SearchCountry(106.9123, 29.4316)

	fmt.Println(cn)
}
