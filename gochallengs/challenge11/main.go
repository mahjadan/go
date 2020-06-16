package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

// Flight struct which contains
// the origin, destination and price of a flight
type Flight struct {
	Origin      string
	Destination string
	Price       int
}

// IsSubset checks to see if the first set of
// flights is a subset of the second set of flights.

func IsSubset(first, second []Flight) bool {
	// implement
	for _, val := range first {
		if !exists(val, second) {
			return false
		}
	}
	return true
}

func exists(elm Flight, list []Flight) bool {
	for _, val := range list {
		if elm == val {
			return true
		}
	}
	return false
}

func main() {

	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file.")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	fmt.Println("Sets and Subsets Challenge")
	firstFlights := []Flight{
		Flight{Origin: "GLA", Destination: "CDG", Price: 1000},
		Flight{Origin: "GLA", Destination: "JFK", Price: 5000},
		Flight{Origin: "GLA", Destination: "SNG", Price: 3000},
	}

	secondFlights := []Flight{
		Flight{Origin: "GLA", Destination: "CDG", Price: 1000},
		Flight{Origin: "GLA", Destination: "JFK", Price: 5000},
		Flight{Origin: "GLA", Destination: "SNG", Price: 3000},
		Flight{Origin: "GLA", Destination: "AMS", Price: 500},
	}

	subset := IsSubsetHash(firstFlights, secondFlights)
	fmt.Println(subset)
}

//todo: depending on the requirement if (there are doubles in the first set is it considered or not as subset of second)
// u may need to add more checking
func IsSubsetHash(first, second []Flight) bool {

	set := make(map[Flight]int)

	for _, val := range second {
		set[val]++
	}
	fmt.Println("set:  ", set)

	for _, val := range first {
		fmt.Println("checking: ", val)
		if count, found := set[val]; !found {
			return false
		} else if count < 1 { //this is the part where extra checking to avoid dubplication
			return false
		} else {
			set[val] = count - 1
		}
	}
	return true

}
