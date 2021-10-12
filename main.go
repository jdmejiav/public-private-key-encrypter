package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
	"unicode"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func randomInt(c rune) int {
	r := rand.NewSource(time.Now().UnixNano())
	s := rand.New(r)
	num := s.Intn(1024)
	for {

		sign := rand.Intn(2)
		if sign == 1 {
			num = num * (-1)
		}
		if num+int(c) > 0 {
			break
		}
	}
	time.Sleep(1 * time.Microsecond)
	return num

}

func main() {

	file, err := os.ReadFile("./data.txt")
	check(err)
	data := string(file)
	fOut1, err := os.Create("./dataEnc.out")
	fOut2, err := os.Create("./key.out")
	check(err)
	var out string = ""
	var changes string = ""
	for _, c := range data {

		if !unicode.IsSpace(c) {
			randInt := randomInt(c)
			newRune := rune(randInt + int(c))
			changes += strconv.Itoa(randInt) + " "
			out += string(newRune)
		} else {
			out += string(c)
		}

	}

	dataEnc := []byte(out)
	fOut1.Write(dataEnc)

	dataKey := []byte(changes)
	fOut2.Write(dataKey)

	fmt.Println(out)
	fmt.Println(changes)

}
