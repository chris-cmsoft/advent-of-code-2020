package main

import (
	"bufio"
	"errors"
	"flag"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
)

func SumInputs(inputs []int64) int64 {
	total := int64(0)
	for _, input := range inputs {
		total += input
	}
	return total
}

func GetCombosMatching(total int64, inputs []int64) ([]int64, error) {
	for key1, input1 := range inputs {
		// We want to only loop items after the first, as items before it would already have been compared before
		for key2, input2 := range inputs[key1+1:] {
			key2calculated := key1 + key2 + 1
			for _, input3 := range inputs[key2calculated+1:] {
				if SumInputs([]int64{
					input1,
					input2,
					input3,
				}) == total {
					return []int64{
						input1,
						input2,
						input3,
					}, nil
				}
			}
		}
	}
	return []int64{}, nil
}

func GetInput(filename string) ([]int64, error) {
	inputs := make([]int64, 0)
	file, err := os.Open(filename)
	if err != nil {
		return inputs, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parsedInteger, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			return inputs, err
		}
		inputs = append(inputs, parsedInteger)
	}
	return inputs, nil
}

func main() {

	log.Info().Msg("Day 1")

	log.Info().Msg("Parsing input")

	var inputFile string
	flag.StringVar(&inputFile, "input", "", "Input file to use for puzzle")
	flag.Parse()

	if inputFile == "" {
		flag.Usage()
		log.Panic().Err(errors.New("input file is required")).Msg("input file is required")
	}

	inputs, err := GetInput(inputFile)
	if err != nil {
		panic(err)
	}

	matches, err := GetCombosMatching(2020, inputs)
	puzzleOutput := matches[0] * matches[1] * matches[2]
	log.Info().Msgf("Answer: %d", puzzleOutput)
}
