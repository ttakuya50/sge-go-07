package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// csvファイルを読み込み
	file, err := os.Open(wd + "/shiftJis.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// transform,japaneseを使用しShiftJIS変換
	reader := transform.NewReader(file, japanese.ShiftJIS.NewDecoder())
	r := csv.NewReader(reader)

	// 文字コード変換を行わない場合はこっち↓
	//r := csv.NewReader(file)

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
