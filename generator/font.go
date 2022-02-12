package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type font map[rune][][]int

func decodeCharacter(data string) (chr rune, bitmap [][]int) {
	unicodeExp := regexp.MustCompile(`Unicode: \[([^\]]*)\];`)
	subgroups := unicodeExp.FindStringSubmatch(data)
	if len(subgroups) < 2 {
		return -1, nil
	}

	chrInt32, _ := strconv.ParseInt(subgroups[1], 16, 32)
	chr = rune(chrInt32)

	lines := strings.Split(data, "\n")

	for _, line := range lines {
		if len(line) == 0 || (line[0] != ' ' && line[0] != 'B') {
			continue
		}
		bitmap = append(bitmap, []int{})
		for _, chr := range line {
			switch(chr) {
			case '-':
				bitmap[len(bitmap)-1] = append(bitmap[len(bitmap)-1], 0)
				break
			case '#':
				bitmap[len(bitmap)-1] = append(bitmap[len(bitmap)-1], 1)
				break
			}
		}
	}

	return chr, bitmap
}

func fromTXT(data string) font {
	characters := strings.Split(data, "\n%\n")
	font := map[rune][][]int{}
	for _, characterData := range characters {
		chr, bitmap := decodeCharacter(characterData)
		font[chr] = bitmap
	}
	return font
}

var example = map[rune][][]int {
	1: {
		{0, 1, 0, 0,},
		{0, 1, 0, 0,},
	},
	2: {
		{0, 1, 0, 0,},
		{0, 1, 0, 0,},
	},
}

func serialize(name string, font font) {
	fmt.Printf("var %v = map[rune][][]int {\n", name)
	for chr, bitmap := range font {
		fmt.Printf("	%v: {\n", chr)
		for _, row := range bitmap {
			fmt.Print("		{")
			for _, bit := range row {
				fmt.Printf("%v, ", bit)
			}
			fmt.Print("},\n")
		}
		fmt.Print("	},\n")
	}
	fmt.Print("}\n")
}

func main() {
	if len(os.Args) < 3 {
		panic("Usage: go run . FONT.PSF NAME")
	}

	command := exec.Command("psf2txt", "/dev/stdin", "/dev/stdout")
	reader, err := os.Open(os.Args[1])
	command.Stdin = reader

	if err != nil {
		panic("Could not read file")
	}
	
	output, err := command.Output()

	if err != nil {
		panic("psf2txt command failed")
	}

	font := fromTXT(string(output))

	fmt.Print("package main\n\n")
	serialize(os.Args[2], font)
}
