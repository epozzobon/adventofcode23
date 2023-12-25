package main

import (
	"fmt"
	"strconv"

	"epozzobon.it/adventofcode23/utils"
)

type orange struct {
	from int
	to   int
}

type traslation struct {
	src orange
	dst orange
}

func Day6(filepath string, problemPart int) int {

	file, err := utils.ReadLines(filepath)
	if err != nil {
		panic(err)
	}

	line := file[0]
	if line[:6] != "Time: " {
		panic("expecting times")
	}
	times, err := utils.SpacedIntsList(line[6:])
	if err != nil {
		panic("expected int list")
	}

	line = file[1]
	if line[:10] != "Distance: " {
		panic("expecting distances")
	}
	records, err := utils.SpacedIntsList(line[10:])
	if err != nil {
		panic("expected int list")
	}

	if len(records) != len(times) {
		panic("expected same number of records and times")
	}

	if problemPart == 1 {
		output := 1
		for i := 0; i < len(times); i++ {
			raceDuration := times[i]
			distanceRecord := records[i]

			fmt.Println(raceDuration, distanceRecord)
			waysToBeat := countWaysToBeat(raceDuration, distanceRecord)
			output *= waysToBeat
		}
		return output
	} else if problemPart == 2 {

		timeStr := ""
		recordStr := ""
		for i := 0; i < len(times); i++ {
			timeStr += strconv.Itoa(times[i])
			recordStr += strconv.Itoa(records[i])
		}

		raceDuration, err := strconv.Atoi(timeStr)
		if err != nil {
			panic("time is not a number")
		}
		distanceRecord, err := strconv.Atoi(recordStr)
		if err != nil {
			panic("distance is not a number")
		}

		fmt.Println(raceDuration, distanceRecord)
		return countWaysToBeat(raceDuration, distanceRecord)
	} else {
		panic("Unknown problem part")
	}
}

func countWaysToBeat(raceDuration, distanceRecord int) int {
	minimumSpeed := 0
	for speed := 0; speed < raceDuration; speed++ {
		if raceDuration*speed-speed*speed > distanceRecord {
			minimumSpeed = speed
			break
		}
	}

	maximumSpeed := 0
	for speed := raceDuration - 1; speed >= 0; speed-- {
		if raceDuration*speed-speed*speed > distanceRecord {
			maximumSpeed = speed
			break
		}
	}
	waysToBeat := maximumSpeed - minimumSpeed + 1
	fmt.Println(minimumSpeed, maximumSpeed)
	fmt.Println(waysToBeat)
	return waysToBeat
}

func main() {
	utils.CheckSolution(Day6, "example1.txt", 1, 288)
	utils.CheckSolution(Day6, "example1.txt", 2, 71503)
	utils.CheckSolution(Day6, "input.txt", 1, 227850)
	utils.CheckSolution(Day6, "input.txt", 2, 42948149)
}
