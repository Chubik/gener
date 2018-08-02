package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type DataSentence struct {
	First  []string
	Second []string
	Third  []string
}

var DataSet DataSentence

func main() {
	rand.Seed(time.Now().UnixNano())

	file, err := ioutil.ReadFile("./data.json")
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	err = json.Unmarshal(file, &DataSet)
	if err != nil {
		fmt.Printf("File format error: %v\n", err)
		os.Exit(1)
	}

	http.HandleFunc("/", sayGen)
	if err := http.ListenAndServe(":8085", nil); err != nil {
		panic(err)
	}
}

func sayGen(w http.ResponseWriter, r *http.Request) {
	first := DataSet.First[randomInt(len(DataSet.First))]
	second := DataSet.Second[randomInt(len(DataSet.Second))]
	third := DataSet.Third[randomInt(len(DataSet.Third))]
	message := fmt.Sprintf("%v %v %v\n", first, second, third)

	w.Write([]byte(message))
}

func randomInt(i int) int {
	return rand.Intn(i)
}
