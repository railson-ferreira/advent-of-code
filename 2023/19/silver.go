package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type part struct {
	x int
	m int
	a int
	s int
}

type inputData struct {
	workflows map[string]func(thePart part, onAccept func(), onReject func())
	parts     []part
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	bytes, err := os.ReadFile("input.txt")
	check(err)
	input := string(bytes)

	theInputData := formatInput(input)

	sum := 0
	for index, thePart := range theInputData.parts {
		theInputData.workflows["in"](thePart, func() {
			sum += thePart.x + thePart.m + thePart.a + thePart.s
		}, func() {
			fmt.Println("The part", index, "was reject")
		})
	}

	fmt.Println(sum)

}

func formatInput(input string) (theInputData *inputData) {
	lines := strings.Split(input, "\n")

	regex1 := regexp.MustCompile("(\\w+){(.+)}")
	regex2 := regexp.MustCompile("^(\\w+)$|([xmas])([<>])(\\d+):(\\w+)")
	regex3 := regexp.MustCompile("{x=(\\d+),m=(\\d+),a=(\\d+),s=(\\d+)}")
	isPartsPart := false

	theInputData = &inputData{
		workflows: make(map[string]func(thePart part, onAccept func(), onReject func())),
	}

	for _, line := range lines {
		if len(line) == 0 {
			isPartsPart = true
			continue
		}
		if !isPartsPart {
			results1 := regex1.FindStringSubmatch(line)
			workflowName := results1[1]
			var rules []func(thePart part) (stop bool, next string)
			for _, rule := range strings.Split(results1[2], ",") {
				results2 := regex2.FindStringSubmatch(rule)
				direct := results2[1]
				category := results2[2]
				operator := results2[3]
				var operatorIsGreaterThan bool
				if operator != "" {
					switch operator {
					case ">":
						operatorIsGreaterThan = true
					case "<":
						operatorIsGreaterThan = false
					default:
						panic(fmt.Sprintf("Invalid operator '%v' in rule '%v'", operator, rule))
					}
				}
				operand := results2[4]
				var operandInt int
				if operand != "" {
					operandInt, _ = strconv.Atoi(operand)
				}
				next := results2[5]
				rules = append(rules, func(thePart part) (_stop bool, _next string) {
					if direct != "" {
						return true, direct
					}
					var operationResult bool
					switch category {
					case "x":
						if operatorIsGreaterThan {
							operationResult = thePart.x > operandInt
						} else {
							operationResult = thePart.x < operandInt
						}
					case "m":
						if operatorIsGreaterThan {
							operationResult = thePart.m > operandInt
						} else {
							operationResult = thePart.m < operandInt
						}
					case "a":
						if operatorIsGreaterThan {
							operationResult = thePart.a > operandInt
						} else {
							operationResult = thePart.a < operandInt
						}
					case "s":
						if operatorIsGreaterThan {
							operationResult = thePart.s > operandInt
						} else {
							operationResult = thePart.s < operandInt
						}
					default:
						panic(fmt.Sprintf("Invalid category '%v' in rule '%v'", category, rule))
					}
					if operationResult {
						return true, next
					}
					// do not stop
					return false, "NEVER_REACHABLE"
				})
			}
			theInputData.workflows[workflowName] = func(thePart part, onAccept func(), onReject func()) {
				for _, rule := range rules {

					stop, next := rule(thePart)
					if !stop {
						if next != "NEVER_REACHABLE" {
							panic("When stop = true, next must be = 'NEVER_REACHABLE'")
						}
						continue
					}
					if next == "A" {
						onAccept()
						break
					} else if next == "R" {
						onReject()
						break
					} else {
						// Is it trampoline?
						theInputData.workflows[next](thePart, onAccept, onReject)
						break
					}
				}
			}
		} else {
			results := regex3.FindStringSubmatch(line)
			x, _ := strconv.Atoi(results[1])
			m, _ := strconv.Atoi(results[2])
			a, _ := strconv.Atoi(results[3])
			s, _ := strconv.Atoi(results[4])
			theInputData.parts = append(theInputData.parts, part{x, m, a, s})
		}

	}

	return
}
