package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
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

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
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

func applyTraslations(seeds []orange, traslations []traslation) []orange {
	sort.SliceStable(seeds, func(i, j int) bool {
		return seeds[i].from < seeds[j].from
	})
	sort.SliceStable(traslations, func(i, j int) bool {
		return traslations[i].src.from < traslations[j].src.from
	})

	fmt.Println(traslations)

	i, j := 0, 0
	splitSeeds := []orange{}
	for i < len(seeds) {
		if j >= len(traslations) {
			// no traslations left, don't traslate this segment
			splitSeeds = append(splitSeeds, seeds[i])
			fmt.Println(seeds[i], "->", seeds[i])
			i++
		} else if seeds[i].to <= traslations[j].src.from {
			// non-traslated segment
			splitSeeds = append(splitSeeds, seeds[i])
			fmt.Println(seeds[i], "->", seeds[i])
			i++
		} else if seeds[i].from >= traslations[j].src.to {
			// look at next traslation, this one is "old"
			j++
		} else if seeds[i].from >= traslations[j].src.from && seeds[i].to <= traslations[j].src.to {
			// segment is fully traslated
			t := traslations[j].dst.from - traslations[j].src.from
			src := orange{from: seeds[i].from, to: seeds[i].to}
			dst := orange{from: src.from + t, to: src.to + t}
			fmt.Println(src, "->", dst)
			splitSeeds = append(splitSeeds, orange{from: seeds[i].from + t, to: seeds[i].to + t})
			i++
		} else if seeds[i].from < traslations[j].src.from {
			// seed:                <----|----------~~
			// traslation                <----------~~
			// split                <--->
			// leftover                  <----------~~
			// NOTE: split is NOT traslated
			src := orange{from: seeds[i].from, to: traslations[j].src.from}
			fmt.Println(src, "->", src)
			splitSeeds = append(splitSeeds, src)
			seeds[i].from = traslations[j].src.from
		} else if seeds[i].to > traslations[j].src.to {
			// seed:                <--------------|-->
			// traslation        <----------------->
			// split                <------------->
			// leftover                            <-->
			// NOTE: split IS traslated
			t := traslations[j].dst.from - traslations[j].src.from
			src := orange{from: seeds[i].from, to: traslations[j].src.to}
			dst := orange{from: src.from + t, to: src.to + t}
			fmt.Println(src, "->", dst)
			splitSeeds = append(splitSeeds, dst)
			seeds[i].from = traslations[j].src.to
		} else {
			panic("what?")
		}
	}

	seeds = splitSeeds
	sort.SliceStable(seeds, func(i, j int) bool {
		return seeds[i].from < seeds[j].from
	})

	return seeds
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

	seedsLine := scanner.Text()
	if seedsLine[:7] != "seeds: " {
		panic("expecting seeds")
	}

	numList, err := makeIntList(seedsLine[7:])
	if err != nil {
		panic("expected int list")
	}

	seedsRanges := make([]orange, len(numList)/2)
	traslations := []traslation{}
	for i := 0; i < len(numList)/2; i++ {
		seedsRanges[i].from = numList[i*2]
		seedsRanges[i].to = numList[i*2] + numList[i*2+1]
	}

	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {

		} else if txt[len(txt)-1:] == ":" {
			if len(traslations) > 0 {
				seedsRanges = applyTraslations(seedsRanges, traslations)
				fmt.Println(seedsRanges)
				fmt.Println(txt)
			}
			traslations = []traslation{}
		} else {
			t, err := makeIntList(txt)
			if err != nil {
				panic("expected int list")
			}
			if len(t) != 3 {
				panic("expected 3 integers")
			}
			dstRange := orange{from: t[0], to: t[0] + t[2]}
			srcRange := orange{from: t[1], to: t[1] + t[2]}
			traslations = append(traslations, traslation{src: srcRange, dst: dstRange})

		}
	}
	seedsRanges = applyTraslations(seedsRanges, traslations)
	fmt.Println(seedsRanges)

	err1 := scanner.Err()
	if err1 != nil {
		panic(err1)
	}
}
