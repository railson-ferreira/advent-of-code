package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type moduleType int8
type pulse bool

const (
	flipFlop moduleType = iota + 1
	conjunction
	broadcaster
)
const (
	low  pulse = false
	high pulse = true
)

type module struct {
	modType moduleType
	modName string
	inputs  []*module
	outputs []*module
}

type inputData struct {
	modules map[string]*module
}

type moduleState struct {
	theModule    *module
	inputsStates []pulse
}

type machineState struct {
	broadcaster *module
	states      map[string]*moduleState
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
	machine := bootMachine(theInputData)

	sumLow, sumHigh := 0, 0
	for i := 0; i < 1000; i++ {
		sumLow++
		countLow, countHigh := machine.pressTheButton()
		sumLow += countLow
		sumHigh += countHigh
	}
	fmt.Println("low:", sumLow, "high:", sumHigh, "result:", sumLow*sumHigh)

}

func formatInput(input string) (theInputData *inputData) {
	theInputData = &inputData{
		modules: make(map[string]*module),
	}
	lines := strings.Split(input, "\n")

	isSecondPart := false
	for {

		for _, line := range lines {
			if len(line) == 0 {
				continue
			}
			lineParts := strings.Split(line, "->")
			if !isSecondPart {
				moduleTypeAndName := strings.Replace(lineParts[0], " ", "", -1)
				modType := moduleTypeAndName[0]
				modName := moduleTypeAndName[1:]
				theModule := &module{
					modName: modName,
				}

				switch modType {
				case '%':
					theModule.modType = flipFlop
				case '&':
					theModule.modType = conjunction
				default:
					if moduleTypeAndName == "broadcaster" {
						theModule.modType = broadcaster
						theModule.modName = "broadcaster"
						modName = "broadcaster"
					} else {
						panic("Invalid modules type")
					}
				}
				theInputData.modules[modName] = theModule
			} else {
				moduleTypeAndName := strings.Replace(lineParts[0], " ", "", -1)
				modName := moduleTypeAndName[1:]
				if moduleTypeAndName == "broadcaster" {
					modName = "broadcaster"
				}
				theInput := theInputData.modules[modName]
				for _, outputName := range strings.Split(strings.Replace(lineParts[1], " ", "", -1), ",") {
					output := theInputData.modules[outputName]
					if output == nil {
						fmt.Println("There is no module ", outputName, ". Faking...")
						fake := &module{
							modName: outputName,
							modType: -1,
						}
						theInput.outputs = append(theInput.outputs, fake)
						fake.inputs = append(fake.inputs, theInput)
						continue
					}
					theInput.outputs = append(theInput.outputs, output)
					output.inputs = append(output.inputs, theInput)
				}
			}
		}
		if isSecondPart {
			break
		} else {
			isSecondPart = true
		}
	}

	return
}

func bootMachine(theInput *inputData) (theMachineState *machineState) {
	theMachineState = &machineState{}
	theMachineState.broadcaster = theInput.modules["broadcaster"]
	theMachineState.states = make(map[string]*moduleState)
	for modName, theModule := range theInput.modules {
		theModuleState := &moduleState{
			theModule:    theModule,
			inputsStates: make([]pulse, len(theModule.inputs)),
		}
		for index := range theModuleState.inputsStates {
			// Ensure starts as low
			theModuleState.inputsStates[index] = low
		}
		theMachineState.states[modName] = theModuleState
	}
	return
}

func (theMachine *machineState) pressTheButton() (lowPulses int, highPulses int) {
	nextSteps := make([]pulseStep, 0)
	fmt.Println("Sending broadcast to", len(theMachine.broadcaster.outputs))
	lowPulses += len(theMachine.broadcaster.outputs)
	for _, toModule := range theMachine.broadcaster.outputs {
		nextSteps = append(nextSteps, theMachine.handlePulse(low, theMachine.broadcaster, toModule)...)
	}

	for len(nextSteps) > 0 {
		fmt.Println("Sending more", len(nextSteps), "pulses")
		var stepsCopy = append(nextSteps)
		nextSteps = make([]pulseStep, 0)
		for _, step := range stepsCopy {
			nextSteps = append(nextSteps, step(func(from string, thePulse pulse, to string) {
				fmt.Println(from, thePulse, to)
				if thePulse == high {
					highPulses++
				} else {
					lowPulses++
				}
			})...)
		}
	}
	return
}

type pulseStep func(notify func(string, pulse, string)) []pulseStep

func (theMachine *machineState) handlePulse(thePulse pulse, fromModule *module, toModule *module) (nextSteps []pulseStep) {
	var nextPulse pulse
	switch toModule.modType {
	case broadcaster:
		nextPulse = thePulse
	case flipFlop:
		if thePulse == high {
			return
		}
		nextPulse = !theMachine.states[toModule.modName].inputsStates[0]
		theMachine.states[toModule.modName].inputsStates[0] = nextPulse
	case conjunction:
		fromModuleIndex := slices.Index(toModule.inputs, fromModule)
		theMachine.states[toModule.modName].inputsStates[fromModuleIndex] = thePulse
		nextPulse = low
		for _, toModuleInputPulse := range theMachine.states[toModule.modName].inputsStates {
			if toModuleInputPulse == low {
				nextPulse = high
				break
			}
		}
	case -1:
		return
	default:
		panic("Invalid module type")
	}
	for _, theOutputModuleOfToModule := range toModule.outputs {
		copyRef := theOutputModuleOfToModule
		nextSteps = append(nextSteps, func(notify func(string, pulse, string)) []pulseStep {
			notify(toModule.modName, nextPulse, copyRef.modName)
			return theMachine.handlePulse(nextPulse, toModule, copyRef)
		})
	}
	return
}
