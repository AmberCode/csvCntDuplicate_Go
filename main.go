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
	log.Println("Start msisdn count")

	start := time.Now()

	hashTable := CountRepeats()

	for key, value := range hashTable {
		fmt.Println("Msisdn:", key, "Count rpt:", value)
	}

	elapsed := time.Since(start)
	log.Printf("Took: %s", elapsed)
	log.Println("End")
}

func CountRepeats() map[string]int {
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
