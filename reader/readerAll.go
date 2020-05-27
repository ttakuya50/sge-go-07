package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func _main() {
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

	// 一度に全て読み込み
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", records)
}
