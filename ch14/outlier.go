package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func variance(x []float64) float64 {
	mean := meanValue(x)
	sum := float64(0)
	for _, v := range x {
		sum = sum + (v-mean)*(v-mean)
	}
	return sum / float64(len(x))
}

func meanValue(x []float64) float64 {
	sum := float64(0)
	for _, v := range x {
		sum = sum + v
	}
	return sum / float64(len(x))
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("usage: stats filename\n")
		return
	}

	file := flag.Args()[0]
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}
		line = strings.TrimRight(line, "\r\n")
		value, err := strconv.ParseFloat(line, 64)
		if err == nil {
			data = append(data, value)
		}
	}

	sort.Float64s(data)

	fmt.Println("Standard Deviation:", math.Sqrt(variance(data)))

}
