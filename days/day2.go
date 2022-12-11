package days

import (
	"bufio"
	"fmt"
	"os"
)

type gameElem int

const (
	rock gameElem = iota
	paper
	scissors
)

var _charToElem = map[byte]gameElem{
	'A': rock,
	'B': paper,
	'C': scissors,
	'X': rock,
	'Y': paper,
	'Z': scissors,
}

var _scoreShape = map[gameElem]int{
	rock:     1,
	paper:    2,
	scissors: 3,
}

const (
	win  = 6
	draw = 3
	lose = 0
)

var _wins = map[gameElem]gameElem{
	rock:     paper,
	paper:    scissors,
	scissors: rock,
}

var _loses = map[gameElem]gameElem{
	rock:     scissors,
	paper:    rock,
	scissors: paper,
}

type outcomeType int

const (
	winType outcomeType = iota
	loseType
	drawType
)

var _charToOutcome = map[byte]outcomeType{
	'X': loseType,
	'Y': drawType,
	'Z': winType,
}

func Day2() int {
	reader := bufio.NewReader(os.Stdin)
	res := 0

	for {
		var in, out string
		_, err := fmt.Fscanf(reader, "%s %s\n", &in, &out)
		if err != nil {
			break
		}
		cur := outcome(in[0], out[0])
		fmt.Println(in[0], out[0], cur)
		res += cur
	}

	return res
}

func outcome(opponent, youShould byte) int {
	opponentShape := _charToElem[opponent]
	yourOutcome := _charToOutcome[youShould]

	yourShape := rock
	res := 0

	switch yourOutcome {
	case winType:
		yourShape = _wins[opponentShape]
		res = win + _scoreShape[yourShape]
	case loseType:
		yourShape = _loses[opponentShape]
		res = lose + _scoreShape[yourShape]
	case drawType:
		yourShape = opponentShape
		res = draw + _scoreShape[yourShape]
	}

	return res
}
