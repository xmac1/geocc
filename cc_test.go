package geocc_test

import (
	"fmt"
	"testing"

	"aladinfun.com/Koalas/server/srv/recsrv/internal/collect/engine/geocc"
)

func BenchmarkGeo2Country(b *testing.B) {
	if err := geocc.InitCountryMap("countries.json"); err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		fmt.Println(geocc.Geo2Country([]float32{116.4074, 39.9042}))
	}

}

func TestGeo2Country(t *testing.T) {
	if err := geocc.InitCountryMap("countries.json"); err != nil {
		panic(err)
	}

	fmt.Println(geocc.Geo2Country([]float32{100.5018, 13.7563}))
}
