package main

import (
	"encoding/csv"
	"log"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {
	// ファイル作成
	file, err := os.Create("shiftJis.csv")
	if err != nil {
		log.Fatal(err)
	}

	records := [][]string{
		{"田中", "23歳", "ラーメン"},
		{"佐々木", "25歳", "肉"},
		{"佐藤", "21歳", "ケーキ"},
	}

	// transform,japaneseを使用しShiftJIS変換
	writer := transform.NewWriter(file, japanese.ShiftJIS.NewEncoder())
	w := csv.NewWriter(writer)

	for _, r := range records {
		if err := w.Write(r); err != nil {
			log.Fatal(err)
		}
	}
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
