package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

const fileName = "test.csv"

func main() {
	log.Println("Start")

	start := time.Now()

	hashTable := countDuplicates()

	for key, value := range hashTable {
		fmt.Println("Msisdn:", key, "Count rpt:", value)
	}

	elapsed := time.Since(start)
	log.Printf("Took: %v", elapsed.Seconds())
	log.Println("End")
}

func countDuplicates() map[string]int {
	hashTable := make(map[string]int)

	f, err := os.Open(fileName)
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(bufio.NewReader(f))

	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
		}

		result := strings.Split(record[0], ";")

		if err != nil {
			log.Println(err)
			os.Exit(0)
		}

		hashTable[result[1]]++
	}

	return hashTable
}
