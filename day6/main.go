package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type orange struct {
	from int
	to   int
}

type traslation struct {
	src orange
	dst orange
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

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		panic("missing first line")
	}

	line := scanner.Text()
	if line[:6] != "Time: " {
		panic("expecting times")
	}
	times, err := makeIntList(line[6:])
	if err != nil {
		panic("expected int list")
	}
	if !scanner.Scan() {
		panic("missing second line")
	}

	line = scanner.Text()
	if line[:10] != "Distance: " {
		panic("expecting distances")
	}
	records, err := makeIntList(line[10:])
	if err != nil {
		panic("expected int list")
	}

	err1 := scanner.Err()
	if err1 != nil {
		panic(err1)
	}

	if len(records) != len(times) {
		panic("expected same number of records and times")
	}

	output := 1
	for i := 0; i < len(times); i++ {
		raceDuration := times[i]
		r := records[i]

		fmt.Println(raceDuration, r)
		waysToBeat := 0
		for speed := 0; speed < raceDuration; speed++ {
			travelTime := raceDuration - speed
			travelDistance := travelTime * speed
			if travelDistance > r {
				// beat the record
				waysToBeat++
			}
		}
		fmt.Println(waysToBeat)
		output *= waysToBeat
	}
	fmt.Println(output)

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
}
