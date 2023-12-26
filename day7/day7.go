package day7

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"epozzobon.it/adventofcode23/utils"
)

type play struct {
	bid  int
	hand [5]rune
}

func makeIntList(txt string) ([]int, error) {
	re := regexp.MustCompile(`\d+`)
	vs := re.FindAll([]byte(txt), -1)
	vsm := make([]int, len(vs))
	for i, v := range vs {
		n, err := strconv.Atoi(string(v))
		if err != nil {
			return nil, err
		}
		vsm[i] = n
	}
	return vsm, nil
}

func valueOfCard(element rune) int {
	var cards = []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2', '*'}
	for k, v := range cards {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func findHandType(hand [5]rune) int {
	vhand := []rune{}
	jokers := 0
	for _, c := range hand {
		if c == '*' {
			jokers++
		} else {
			vhand = append(vhand, c)
		}
	}

	dupesCount := [5]int{}
	for i := 0; i < len(vhand); i++ {
		for j := i + 1; j < len(vhand); j++ {
			if dupesCount[j] != -1 { // Make sure to only count characters that we didn't see dupes for yet
				if vhand[i] == vhand[j] {
					dupesCount[i]++
					dupesCount[j] = -1 // Mark this position as already counted
				}
			}
		}
	}

	if jokers == 5 || jokers == 4 {
		return 6 // Five of a kind
	}

	pairs := 0
	triplets := 0
	for _, d := range dupesCount {
		if d == 1 {
			pairs++
		} else if d == 2 {
			triplets++
		} else if d == 3 {
			if jokers == 0 {
				return 5 // Four of a kind
			} else if jokers == 1 {
				return 6 // Five of a kind
			} else {
				panic("impossible number of jokers")
			}
		} else if d == 4 {
			return 6 // Five of a kind
		} else if d != -1 && d != 0 {
			panic("impossible number of dupes")
		}
	}

	if triplets == 1 && pairs == 1 {
		if jokers == 0 {
			return 4 // Full house
		}
	} else if triplets == 1 && pairs == 0 {
		if jokers == 0 {
			return 3 // Three of a kind
		} else if jokers == 1 {
			return 5 // Four of a kind
		} else if jokers == 2 {
			return 6 // Five of a kind
		}
	} else if pairs == 2 && triplets == 0 {
		if jokers == 0 {
			return 2 // Two pair
		} else if jokers == 1 {
			return 4 // Full house
		}
	} else if pairs == 1 && triplets == 0 {
		if jokers == 0 {
			return 1 // One pair
		} else if jokers == 1 {
			return 3 // Three of a kind
		} else if jokers == 2 {
			return 5 // Four of a kind
		} else if jokers == 3 {
			return 6 // Five of a kind
		}
	} else if pairs == 0 && triplets == 0 {
		if jokers == 0 {
			return 0
		} else if jokers == 1 {
			return 1 // One pair
		} else if jokers == 2 {
			return 3 // Three of a kind
		} else if jokers == 3 {
			return 5 // Four of a kind
		}
	}

	panic("impossible number of pairs or triplets")
}

func Day7(filepath string, problemPart int) int {

	texts, err := utils.ReadLines(filepath)
	if err != nil {
		panic(err)
	}

	plays := []play{}
	for _, txt := range texts {
		fmt.Println(txt)

		gameParts := strings.Split(txt, " ")
		if len(gameParts) != 2 {
			panic("bad line")
		}

		hand := gameParts[0]
		bid, err := strconv.Atoi(gameParts[1])
		if err != nil {
			panic("bad bid")
		}

		if len(hand) != 5 {
			panic("bad hand")
		}

		handRunes := [5]rune{}
		for i, c := range hand {
			if problemPart == 2 && c == 'J' {
				c = '*'
			}
			handRunes[i] = c
		}

		plays = append(plays, play{hand: handRunes, bid: bid})
	}

	sort.SliceStable(plays, func(i int, j int) bool {
		p1 := plays[i]
		p2 := plays[j]
		t1 := findHandType(p1.hand)
		t2 := findHandType(p2.hand)
		if t1 != t2 {
			return t1 > t2
		}
		for i := range p1.hand {
			if p1.hand[i] != p2.hand[i] {
				return valueOfCard(p1.hand[i]) < valueOfCard(p2.hand[i])
			}
		}
		panic("Same card")
	})

	sum := 0
	for i, p := range plays {
		rank := len(plays) - i
		sum += p.bid * rank
		fmt.Println(p.bid, "*", rank)
	}

	fmt.Println(sum)
	return sum
}
