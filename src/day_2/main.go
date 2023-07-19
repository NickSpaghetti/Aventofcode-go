package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	rpcMap = map[string]RPC{
		"A": RPC{name: "Rock", value: 1, beats: "Scissors"},
		"B": RPC{name: "Paper", value: 2, beats: "Rock"},
		"C": RPC{name: "Scissors", value: 3, beats: "Paper"},
		"X": RPC{name: "Rock", value: 1, beats: "Scissors"},
		"Y": RPC{name: "Paper", value: 2, beats: "Rock"},
		"Z": RPC{name: "Scissors", value: 3, beats: "Paper"},
	}
)

type RPC struct {
	value int
	name  string
	beats string
}

const winningPoints = 6
const drawPoints = 3

func main() {
	fmt.Println("Starting Question 1")
	//question 1
	turnRpcMap1 := parseStrategyGuide()
	playGame(turnRpcMap1)
	fmt.Println("Starting Question 2")
	//question 2
	turnRpcMap2 := parseStrategyGuide2()
	playGame(turnRpcMap2)
}

func playGame(turnRpcMap map[int][2]RPC) {
	p1Score, p2Score := 0, 0
	for i := 0; i < len(turnRpcMap); i++ {
		r, found := turnRpcMap[i]
		if !found {
			fmt.Println("Turn was not found:", i)
			os.Exit(1)
		}
		p1, p2 := r[0], r[1]
		if p1.beats == p2.name {
			p1Score += winningPoints + p1.value
			p2Score += p2.value
		} else if p1.name == p2.name {
			p1Score += p1.value + drawPoints
			p2Score += p2.value + drawPoints
		} else {
			p2Score += winningPoints + p2.value
			p1Score += p1.value
		}
	}

	if p1Score > p2Score {
		fmt.Println("the elf wins")
		fmt.Println(p1Score)
		fmt.Println("you lose")
		fmt.Println(p2Score)
	} else if p1Score == p2Score {
		fmt.Println("it is a draw")
		fmt.Println(p1Score)
	} else {
		fmt.Println("you win")
		fmt.Println(p2Score)
		fmt.Println("the elf losses")
		fmt.Println(p1Score)
	}
}

func readStrategyGuide() string {
	byteSlice, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	strategyGuide := string(byteSlice)
	return strategyGuide
}

func parseStrategyGuide() map[int][2]RPC {
	decodedTurnMap := map[int][2]RPC{}
	strategyGuideStr := readStrategyGuide()
	splitOnNewLines := strings.Split(strategyGuideStr, "\n")
	for turn, nls := range splitOnNewLines {
		rpcs := strings.Split(nls, " ")
		p1, p1Err := toRpc(rpcs[0])
		if p1Err != nil {
			fmt.Println("Error:", p1Err)
		}
		p2, p2Err := toRpc(rpcs[1])
		if p2Err != nil {
			fmt.Println("Error:", p1Err)
		}
		playerMoves := [2]RPC{*p1, *p2}
		decodedTurnMap[turn] = playerMoves
	}
	return decodedTurnMap
}

func parseStrategyGuide2() map[int][2]RPC {
	decodedTurnMap := map[int][2]RPC{}
	strategyGuideStr := readStrategyGuide()
	splitOnNewLines := strings.Split(strategyGuideStr, "\n")
	for turn, nls := range splitOnNewLines {
		rpcs := strings.Split(nls, " ")
		p1, p1Err := toRpc(rpcs[0])
		if p1Err != nil {
			fmt.Println("Error:", p1Err)
		}
		p2, p2Err := toRpc2(*p1, rpcs[1])
		if p2Err != nil {
			fmt.Println("Error:", p1Err)
		}
		playerMoves := [2]RPC{*p1, *p2}
		decodedTurnMap[turn] = playerMoves
	}
	return decodedTurnMap
}

func toRpc(encodedLetter string) (*RPC, error) {
	decodedLetter, ok := rpcMap[encodedLetter]
	if !ok {
		return nil, errors.New(encodedLetter + "is not a valid encodedLetter found in the strategy guide")
	}
	return &decodedLetter, nil
}

func toRpc2(opponent RPC, encodedLetter string) (*RPC, error) {
	var decodedRpc RPC
	switch encodedLetter {
	case "X": //lose
		for _, r := range rpcMap {
			if r.name == opponent.beats {
				decodedRpc = r
				break
			}
		}
		break
	case "Y": //draw
		decodedRpc = opponent
		break
	case "Z": //win
		for _, r := range rpcMap {
			if r.name != opponent.beats && r.name != opponent.name {
				decodedRpc = r
				break
			}
		}
	default:
		return nil, errors.New(encodedLetter + "is not a valid encoded letter in the strategy guide")
	}
	return &decodedRpc, nil
}
