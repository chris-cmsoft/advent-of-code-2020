package main

import (
	"bufio"
	"errors"
	"flag"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
	"strings"
)

type PasswordPolicy struct {
	min    int
	max    int
	letter string
}

func (policy PasswordPolicy) ValidatePassword(password string) bool {
	if strings.Count(password, policy.letter) < policy.min {
		return false
	}
	if strings.Count(password, policy.letter) > policy.max {
		return false
	}
	return true
}

func ParsePolicy(policyString string) PasswordPolicy {
	policyParts := strings.Split(policyString, " ")
	minMax := strings.Split(policyParts[0], "-")
	min, err := strconv.ParseInt(minMax[0], 10, 64)
	if err != nil {
		panic(err)
	}
	max, err := strconv.ParseInt(minMax[1], 10, 64)
	if err != nil {
		panic(err)
	}
	return PasswordPolicy{
		min:    int(min),
		max:    int(max),
		letter: policyParts[1],
	}
}

func GetInputScanner(file *os.File) (chan string, error) {
	ch := make(chan string)
	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		close(ch)
	}()
	return ch, nil
}

//GetInput is a generator using Go Channels in order to read the input line by line and not load everything into memory
// Why ?! Because I want to learn how channels work :)
func GetInput(filename string) chan string {
	ch := make(chan string)

	go func() {

		close(ch)
	}()
	return ch
}

func main() {

	log.Info().Msg("Day 2")

	var inputFile string
	flag.StringVar(&inputFile, "input", "", "Input file to use for puzzle")
	flag.Parse()

	if inputFile == "" {
		flag.Usage()
		log.Panic().Err(errors.New("input file is required")).Msg("input file is required")
	}

	log.Info().Msg("Parsing input")

	//inputGenerator, err := GetInput(inputFile)
	file, err := os.Open(inputFile)
	if err != nil {
		log.Panic().Err(err).Msgf("Failed to load inputFile")
	}
	defer file.Close()

	scanner, err := GetInputScanner(file)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to load input")
	}

	validPasswords := 0
	for passwordRecord := range scanner {
		passwordParts := strings.Split(passwordRecord, ":")
		policy := ParsePolicy(passwordParts[0])
		password := passwordParts[1]
		if policy.ValidatePassword(password) {
			validPasswords += 1
		}
	}
	log.Info().Msgf("Total correct passwords: %d", validPasswords)

}
