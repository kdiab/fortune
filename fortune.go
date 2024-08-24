package main

import (
	"fmt"
	"math/rand"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func filetostring(file []byte) {
	var arr []string
	var s string
	for i := 0; i < len(file); i++ {
		if file[i] == '%' {
			arr = append(arr, s)
			s = ""
		} else {
			s = s + string(file[i])
		}
	}
	fmt.Printf("%s\n", arr[rand.Intn(len(arr))])
}

func main() {
	fs, err := os.ReadDir("./data/")
	check(err)
	filename := fmt.Sprintf("%s", fs[rand.Intn(len(fs))].Name())

	data, err := os.ReadFile("./data/" + filename)
	check(err)

	filetostring(data)
}
