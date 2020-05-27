package main

import (
	"encoding/csv"
	"log"
	"os"
)

func _main() {
	records := [][]string{
		{"田中", "23歳", "ラーメン"},
		{"佐々木", "25歳", "肉"},
		{"佐藤", "21歳", "ケーキ"},
	}

	w := csv.NewWriter(os.Stdout)
	for _, r := range records {
		// Writeメソッドでバッファに溜め込む
		if err := w.Write(r); err != nil {
			log.Fatal(err)
		}
	}
	w.Flush() // バッファに溜め込まれたものを出力

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
