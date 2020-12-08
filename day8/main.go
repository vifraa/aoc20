package main

import (
	"fmt"
	"github.com/vifraa/aoc20/util"
	"strconv"
	"strings"
)

var accumulator = 0

func main() {
	input, _ := util.ReadInputAsString("./day8/input.txt")

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	commands := parseCommands(input)
	acc, ok := isCorrectExec(commands)
	if !ok {
		return acc
	}
	return -1
}


func part2(input []string) int {
	baseCommands := parseCommands(input)
	allCommandPosibilities := createCmdSequences(baseCommands)

	for _, cmdSeq := range allCommandPosibilities {
		acc, ok := isCorrectExec(cmdSeq)
		if ok {
			return acc
		}
	}

	return -1
}

func isCorrectExec(cmds []Command) (int, bool) {
	lineExecutionCounter := make(map[int]int)
	localAcc := 0
	for i := 0; i < len(cmds);{
		cmd := cmds[i]

		if lineExecutionCounter[i] == 1 {
			return localAcc, false
		}

		if cmd.name == "acc" {
			localAcc += cmd.operand
			lineExecutionCounter[i]++
			i++
		} else if cmd.name == "jmp" {
			lineExecutionCounter[i]++
			i += cmd.operand
		} else if cmd.name == "nop" {
			lineExecutionCounter[i]++
			i++
		}
	}

	return localAcc, true
}

func createCmdSequences(cmds []Command) [][]Command {
	res := make([][]Command, 0)
	for i, c := range cmds {
		if c.name == "jmp" {
			tmp := make([]Command, len(cmds))
			copy(tmp, cmds)
			tmp[i].name = "nop"
			res = append(res, tmp)
		} else if c.name == "nop" {
			tmp := make([]Command, len(cmds))
			copy(tmp, cmds)
			tmp[i].name = "jmp"
			res = append(res, tmp)
		}
	}
	return res
}


func parseCommands(input []string) []Command {
	result := make([]Command, 0)
	for _, i := range input {
		ss := strings.Split(i, " ")
		c, _ := strconv.Atoi(ss[1])

		result = append(result, Command{
			ss[0],
			c,
		})
	}
	return result
}

type Command struct {
	name string
	operand int
}
