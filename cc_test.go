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

	fmt.Println(geocc.Geo2Country([]float32{100.5018, 13.7563}))

	start := time.Now()
	n := 50000
	for i := 0; i < n; i++ {
		geocc.Geo2Country([]float32{100.5018, 13.7563})
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
