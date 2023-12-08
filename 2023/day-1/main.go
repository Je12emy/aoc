package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var stringNumbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	buf, _ := os.Open("./input.txt")
	defer buf.Close()
	scanner := bufio.NewScanner(buf)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		number := ""
		// fmt.Println("input is: ", line)
		f := get_number(line)
		l := get_number(Reverse(line))
		// fmt.Println("first, ", f)
		// fmt.Println("last, ", l)
		number = fmt.Sprintf("%v%v", f, l)
		// fmt.Println("result is :", number)
		// fmt.Println("---")
		conv, _ := strconv.Atoi(number)
		sum = sum + conv
	}
	fmt.Println(sum)
}

func get_number(str string) int {
	var n int
	var err error
	for pos, c := range str {
		n, err = strconv.Atoi(string(c))
		// fmt.Println(n)
		if err != nil {
			// fmt.Println("Not a number")
			// Is not a number, string lookup
			for key := range stringNumbers {
				// fmt.Println("----")
				// fmt.Println("Key is: ", key)
				// fmt.Println("Reverse key is: ", Reverse(key))
				// fmt.Println("offset is", len(key) - 1)
				substr := Substr(str, pos, len(key)-1)
				// fmt.Println("Substring is: ", substr)
				if substr == key || substr == Reverse(key) {
					// fmt.Println("char was: ",string(c))
					n = stringNumbers[key]
					// fmt.Println("String match: ", key)
					return n
				}
			}
			continue
			// Found nothing here!
		}
		break
	}
	// fmt.Println("int match: ", n)
	return n
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func Substr(s string, start int, offset int) string {
	var result string
	for pos, v := range s {
		end := start + offset
		isInRange := pos >= start && pos <= end
		// fmt.Println("range is :" , start , "-", end)
		if isInRange {
			result = result + string(v)
		}
	}
	return result
}
