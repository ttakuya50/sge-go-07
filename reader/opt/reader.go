package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// csvファイルを読み込み
	file, err := os.Open(wd + "/reader.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = '|'             // 区切り文字
	r.FieldsPerRecord = 3     // 各行のフィールド数。多くても少なくてもエラーになる
	r.Comment = '#'           // 1フィールド目が「#」からの始まる場合はコメント行として扱う
	r.TrimLeadingSpace = true // true の場合は、各フィールド先頭の空白文字を無視する
	r.ReuseRecord = true      // true の場合は、Readで戻ってくるスライスを次回再利用する。
	r.TrailingComma = true    // 非推奨なので使わない
	//r.LazyQuotes = true     // 使い方が分からないので割愛

	for {
		// 一行ずつ読み込み
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
	}
}
