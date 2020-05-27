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

	// NewReader関数の引数には読み込みたいファイル情報を指定
	r := csv.NewReader(file)
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
