package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	records := [][]string{
		{"田中", "23歳", "ラーメン"},
		{"佐々木", "25歳", "肉"},
		{"佐藤", "21歳", "ケーキ"},
	}

	w := csv.NewWriter(os.Stdout)
	if err := w.WriteAll(records); err != nil {
		log.Fatal(err)
	}
}
