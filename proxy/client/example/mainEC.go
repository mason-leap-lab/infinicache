package main

import (
	"fmt"
	"github.com/neboduus/infinicache/proxy/client"
	"log"
	"math/rand"
	"strings"
)

func main() {
	addrList := "10.4.0.100:6378"
	// initial object with random value
	val := make([]byte, 160)
	rand.Read(val)

	// parse server address
	addrArr := strings.Split(addrList, ",")

	// initial new ecRedis client
	cli := client.NewClient(10, 2, 32, 3)
	cli.Dial(addrArr)

	key := fmt.Sprintf("k1")
	if _, stats, ok := cli.EcSet(key, val); !ok {
		log.Println("Failed to SET ", key)
	}else{
		log.Println("Successfull SET ", key, " ", stats)
	}

	if _, _, stats, ok := cli.EcGet(key, 160); !ok {
		log.Println("Failed to GET ", key)
	} else {
		log.Println("Successfull GET ", key, " ", stats)
	}

	return
}