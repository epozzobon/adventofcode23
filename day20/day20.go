package main

import (
	"fmt"
	"regexp"

	"epozzobon.it/adventofcode23/utils"
)

type module struct {
	typ   byte
	name  string
	dst   []*module
	src   []*module
	state int
}

type pulse struct {
	src   *module
	dst   *module
	level int
}

func Day20(filename string, problemPart int) int {

	var modules = make(map[string]*module)
	lines, err := utils.ReadLines(filename)
	if err != nil {
		panic(err)
	}
	re1 := regexp.MustCompile(`^([%&]?)([a-z]+) -> ([a-z, ]+)$`)
	re2 := regexp.MustCompile(`([a-z])+`)
	for _, txt := range lines {
		if txt == "" {
			break
		}
		pieces := re1.FindSubmatch([]byte(txt))
		moduleType := string(pieces[1])
		moduleName := string(pieces[2])
		spieces := re2.FindAll(pieces[3], -1)

		moduleDestinations := make([]*module, len(spieces))
		for i, spiece := range spieces {
			moduleDestinationName := string(spiece)
			if modules[moduleDestinationName] == nil {
				modules[moduleDestinationName] = &module{name: moduleDestinationName}
			}
			moduleDestinations[i] = modules[moduleDestinationName]
		}
		fmt.Println(spieces)

		if modules[moduleName] == nil {
			modules[moduleName] = &module{name: moduleName}
		}
		m := modules[moduleName]
		m.typ = 0
		if len(moduleType) > 0 {
			m.typ = moduleType[0]
		}
		m.dst = moduleDestinations
	}

	for _, m := range modules {
		for _, s := range m.dst {
			s.src = append(s.src, m)
		}
	}

	periods := make(map[string]int)
	focused := make(map[string]bool)
	rx := modules["rx"]
	if rx != nil {
		for _, s := range rx.src {
			for _, s := range s.src {
				focused[s.name] = true
				fmt.Println("Focusing on", s.name)
			}
		}
	}

	rxPulses := 0
	buttonPresses := 1
	doButton := func() (int, int) {
		//fmt.Println("button -low-> broadcaster")
		highPulses := 0
		lowPulses := 1
		pulseBackLog := []pulse{{nil, modules["broadcaster"], 0}}
		doPulse := func(m *module) {
			if m.state == 1 {
				highPulses += len(m.dst)
			} else if m.state == 0 {
				lowPulses += len(m.dst)
			}
			for _, d := range m.dst {
				pulseBackLog = append(pulseBackLog, pulse{m, d, m.state})
				if d.name == "rx" && m.state == 0 {
					rxPulses++
				}
				if focused[d.name] && m.state == 0 {
					periods[d.name] = buttonPresses
				}
			}
			stateStr := "low"
			if m.state == 1 {
				stateStr = "high"
			}
			for _, d := range m.dst {
				fmt.Printf("%s -%s-> %s\n", m.name, stateStr, d.name)
				if d.name == "rx" && m.state == 0 {
					fmt.Printf("rx pulsed %d\n", m.state)
				}
				if focused[d.name] && m.state == 0 {
					fmt.Println(d.name, "pulsed low at button press", buttonPresses)
				}
			}
		}
		for len(pulseBackLog) > 0 {
			p := pulseBackLog[0]
			pulseBackLog = pulseBackLog[1:]
			m := p.dst
			if m.name == "" {
				panic("Module not found")
			}
			if m.typ == 0 {
				m.state = p.level
				doPulse(m)
			} else if m.typ == '&' {
				m.state = 0
				for _, k := range m.src {
					s := k.state
					if s == 0 {
						m.state = 1
					}
				}
				doPulse(m)
			} else if m.typ == '%' {
				if p.level == 0 {
					m.state = 1 - m.state
					doPulse(m)
				}
			} else {
				fmt.Println(m)
				panic("unknown module")
			}
		}
		return lowPulses, highPulses
	}
	sumh := 0
	suml := 0
	for buttonPresses = 1; len(periods) < 4 || buttonPresses < 1000; buttonPresses++ {
		h, l := doButton()
		sumh += h
		suml += l
		if buttonPresses == 1000 {
			fmt.Println(buttonPresses, suml, sumh, sumh*suml, rxPulses)
			if problemPart == 1 {
				return sumh * suml
			}
		}
	}

	m := 1
	for _, v := range periods {
		m *= v
	}
	fmt.Println(m)
	if problemPart == 2 {
		return m
	}
	panic("wut")
}

func main() {
	utils.CheckSolution(Day20, "example1.txt", 1, 32000000)
	utils.CheckSolution(Day20, "example2.txt", 1, 11687500)
	utils.CheckSolution(Day20, "input.txt", 1, 763500168)
	utils.CheckSolution(Day20, "input.txt", 2, 207652583562007)
}
