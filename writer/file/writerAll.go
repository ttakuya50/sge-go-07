package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	// ファイル作成
	file, err := os.Create("writerAll.csv")
	if err != nil {
		log.Fatal(err)
	}

	records := [][]string{
		{"田中", "23歳", "ラーメン"},
		{"佐々木", "25歳", "肉"},
		{"佐藤", "21歳", "ケーキ"},
	}

	// 出力先を作成したファイルに指定
	w := csv.NewWriter(file)
	if err := w.WriteAll(records); err != nil {
		log.Fatal(err)
	}
}
