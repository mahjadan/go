package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	generate()
}

func generate() {
	file, _ := os.Create("input-gen.csv")
	writer := csv.NewWriter(file)

	for i := 0; i < 20; i++ {
		line := []string{
			fmt.Sprintf("44-%d", i),
			fmt.Sprintf("mmll%d", i),
			fmt.Sprintf("2020/10/%d", i),
		}
		writer.Write(line)
		writer.Flush()
	}
}
