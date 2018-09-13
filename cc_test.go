package geocc_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/xmac1/geocc"
)

func init() {
	if err := geocc.InitCountryMap("countries_int32.json"); err != nil {
		panic(err)
	}

}

func BenchmarkGeo2Country(b *testing.B) {

	for i := 0; i < b.N; i++ {

	}

}

func TestGeo2Country(t *testing.T) {
	if err := geocc.InitCountryMap("D:/countries.json"); err != nil {
		panic(err)
	}

	start := time.Now()
	n := 1000000
	for i := 0; i < n; i++ {
		geocc.Geo2Country([]float32{113.93474, 22.525246})
	}
	dur := time.Since(start)

	fmt.Println(dur.String())

	fmt.Println("per second, ", int64(n*1000)/(dur.Nanoseconds()/1e6))
}
