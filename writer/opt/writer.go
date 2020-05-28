package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	// ファイル作成
	file, err := os.Create("writer.csv")
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

	w.Comma = '|'    // 区切り文字を変更
	w.UseCRLF = true // 改行文字を CRLF(\r\n) にする

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
