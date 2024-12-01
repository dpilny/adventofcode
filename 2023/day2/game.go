package day2

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type game struct {
	id     int
	rounds []round
}

type round struct {
	red   int
	green int
	blue  int
}

var gameIdRegex = regexp.MustCompile("Game (\\d*)")
var gameRegex = regexp.MustCompile("\\s(\\d*)\\s(\\w*)")

func sumValidGames(path string, r, g, b int) (int, error) {
	validGames, err := validateGames(path, r, g, b)
	if err != nil {
		return -1, err
	}
	sum := 0
	for _, validGame := range validGames {
		sum = sum + validGame.id
	}
	return sum, nil
}

func validateGames(path string, r, g, b int) ([]game, error) {
	games, err := parseGames(path)
	if err != nil {
		return nil, err
	}
	var validGames []game
	for _, game := range games {
		if isValidGame(game, r, g, b) {
			validGames = append(validGames, game)
		}
	}
	return validGames, nil
}

func isValidGame(g game, red, green, blue int) bool {
	for _, round := range g.rounds {
		if !isValidRound(round, red, green, blue) {
			return false
		}
	}
	return true
}

func isValidRound(r round, red, green, blue int) bool {
	if r.red > red {
		return false
	}
	if r.green > green {
		return false
	}
	if r.blue > blue {
		return false
	}
	return true
}

func parseGames(path string) ([]game, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var games []game
	raw := strings.Split(string(data), "\n")
	for _, line := range raw {
		game := parseGame(line)
		fmt.Println(game)
		games = append(games, game)
	}
	return games, nil
}

func parseGame(line string) game {
	parts := strings.Split(line, ":")

	idCapture := gameIdRegex.FindStringSubmatch(parts[0])
	id, _ := strconv.Atoi(idCapture[1])

	rawRounds := strings.Split(parts[1], ";")
	var rounds []round
	for _, rawRound := range rawRounds {
		colors := make(map[string]int)
		values := gameRegex.FindAllStringSubmatch(rawRound, -1)
		for _, val := range values {
			count, _ := strconv.Atoi(val[1])
			color := val[2]
			i, _ := colors[color]
			colors[color] = count + i
		}
		rounds = append(rounds, round{
			red:   colors["red"],
			green: colors["green"],
			blue:  colors["blue"],
		})
	}

	return game{
		id:     id,
		rounds: rounds,
	}
}
