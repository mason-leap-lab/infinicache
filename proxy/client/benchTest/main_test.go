package main

import (
	"fmt"
	"github.com/neboduus/infinicache/proxy/client"
	"math/rand"
	"testing"
)

func BenchmarkMkSet(b *testing.B) {
	sizes := []int{1,160,500,1600}
	cli := client.NewClient(10, 2, 32, 3)
	cli.Dial([]string{"10.4.0.100:6378"})
	for size := range sizes {
		data := cli.GenerateSetData(size)
		b.Run(fmt.Sprintf("/%d B", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_,_,_ = cli.MkSet(fmt.Sprintf("k-%d-%d", i, size), data)
			}
		})
	}
}

func BenchmarkMkGet(b *testing.B) {
	sizes := []int{1,160,500,1600}
	cli := client.NewClient(10, 2, 32, 3)
	cli.Dial([]string{"10.4.0.100:6378"})

	// we first set some data to be sure our GET ops are successfull
	allSets := make(map[int][][3]client.KVSetGroup)
	for size := range sizes {
		dSet := cli.GenerateSetData(size)
		var okSets [][3]client.KVSetGroup
		for i := 0; i <= 500; i++{
			_, _, err := cli.MkSet(fmt.Sprintf("k-%d-%d", i, size), dSet)
			if err != false {
				okSets = append(okSets, dSet)
			}
		}
		allSets[size] = okSets
	}

	for size := range sizes {
		setOps := allSets[size]
		b.Run(fmt.Sprintf("/%d B", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// we randomly choose some data to GET from the previous set ops
				dGet := cli.GenerateSingleRandomGet(setOps[rand.Intn(len(setOps)-1)])
				_,_,_ = cli.MkGet(fmt.Sprintf("k-%d-%d", i, size), dGet)
			}
		})
	}

}