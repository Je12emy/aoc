package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAX_RED int = 12
const MAX_GREEN int = 13
const MAX_BLUE int = 14

type Game struct {
	Id      int
	Series  []map[string]int
	canPlay bool
}

func NewGame(id int, input []string) Game {
	var series []map[string]int
	var isPosible bool
	for _, v := range input {
		match := make(map[string]int)
		// 1 red, 2 blue, 3 green
		entries := strings.Split(v, ",")
		for _, k := range entries {
			// red 1
			parts := strings.Split(strings.TrimSpace(k), " ")
			conv, _ := strconv.Atoi(parts[0])
			// ["red": 1]
			// fmt.Println(parts[1])
			// fmt.Println(conv)
			match[parts[1]] = conv
		}
		series = append(series, match)
	}

	for _, v := range series {
		isPosible = v["red"] <= MAX_RED && v["blue"] <= MAX_BLUE && v["green"] <= MAX_GREEN
		fmt.Println(v, isPosible)
		if !isPosible {
			break
		}
	}

	// fmt.Println("Game ", id, " - " , isPosible)
	return Game{
		Id:      id,
		canPlay: isPosible,
		Series:  series,
	}
}

func sum(series []Game) int {
	sum := 0
	for _, s := range series {
		if s.canPlay {
			sum = sum + s.Id
		}
	}
	return sum
}

func Part1(scanner *bufio.Scanner) {
	gameId := 1
	var series []Game
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		input := strings.Split(strings.Split(line, ":")[1], ";")
		game := NewGame(gameId, input)
		// fmt.Println(line)
		series = append(series, game)
		// take care of when you increase the ID
		gameId = gameId + 1
	}
	result := sum(series)
	fmt.Println(result)
}

func main() {
	buf, _ := os.Open("./input.txt")
	defer buf.Close()
	scanner := bufio.NewScanner(buf)
	Part1(scanner)
}
