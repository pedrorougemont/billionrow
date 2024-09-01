package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Measure struct {
	highest float64
	lowest  float64
	sum     float64
	count   float64
}

func (m Measure) String() string {
	return fmt.Sprintf("{highest: %.2f, lowest: %.2f, avg: %.2f}", m.highest, m.lowest, m.sum/m.count)
}

func main() {
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

	fmt.Printf("%v", cityMap)
}
