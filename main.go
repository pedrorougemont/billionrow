package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Measure struct {
	highest float64
	lowest  float64
	sum     float64
	count   float64
}

func (m Measure) String() string {
	return fmt.Sprintf("%.1f/%.1f/%.1f", m.lowest, m.sum/m.count, m.highest)
}

func main() {
	start := time.Now()
	measurementsFile, err := os.Open("measurements.txt")
	if err != nil {
		panic(err)
	}
	defer measurementsFile.Close()

	cityMap := make(map[string]Measure)
	scanner := bufio.NewScanner(measurementsFile)
	for scanner.Scan() {
		var city string
		var strTemperature string
		var temperature float64
		var measure Measure

		city, strTemperature, found := strings.Cut(scanner.Text(), ";")
		if !found {
			panic("malformed input")
		}

		temperature, err = strconv.ParseFloat(strTemperature, 64)
		if err != nil {
			panic("malformed input")
		}

		measure, ok := cityMap[city]
		if !ok {
			measure = Measure{temperature, temperature, temperature, 1}
		} else {
			measure.count++
			measure.sum += temperature
			if measure.highest < temperature {
				measure.highest = temperature
			}
			if measure.lowest > temperature {
				measure.lowest = temperature
			}
		}

		cityMap[city] = measure
	}

	keys := make([]string, 0, len(cityMap))
	for k := range cityMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	fmt.Printf("{")
	for _, city := range keys {
		fmt.Printf("%s=%v, ", city, cityMap[city])
	}
	fmt.Printf("}\n")
	fmt.Println(time.Since(start))
}
