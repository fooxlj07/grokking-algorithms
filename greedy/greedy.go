package greedy

import (
	"fmt"
	"grokking-algorithms/go-helper"
)

/* Class schedule
class		start_time		end_time
  A			    9				10
  B 		    9:30			10:30
  C			    10				11
  D				10:30			11:30
  E				11				12
*/
// Arrange as much class as possible
func arrangeClass(class []string, schedule [][2]float32) map[string][2]float32 {
	classIndex := []int{0}
	for i, j := 0, 1; i < len(class) && j < len(class); {
		// each step, find the next earlest no conflict class
		if schedule[i][1] <= schedule[j][0] {
			classIndex = append(classIndex, j)
			i = j
		}
		j++
	}

	optSchedule := make(map[string][2]float32)
	for _, i := range classIndex {
		optSchedule[class[i]] = schedule[i]
	}
	return optSchedule
}

func generateClassSchedule() ([]string, [][2]float32) {
	// For convinient purpose, using float instead of time
	schedule := [][2]float32{
		[2]float32{9, 10},
		[2]float32{9.30, 10.30},
		[2]float32{10, 11},
		[2]float32{10.30, 11.30},
		[2]float32{11, 12},
	}
	return []string{"A", "B", "C", "D", "E"}, schedule
}

/*
	Find a minumum set of giving radio set to cover all the states listed
*/
func findMinumumRadioSet(radioSet map[string][]string, states []string) []string {
	finalRadioSet := []string{}
	for len(states) > 0 {
		bestStation := ""

		for radioStation, radioCoveredStates := range radioSet {
			covered := helper.IntersectionStrSlice(radioCoveredStates, states)
			if len(covered) > len(radioSet[bestStation]) {
				bestStation = radioStation
			}
		}
		states = helper.RemoveStrSlice1FromSlice2(states, radioSet[bestStation])
		fmt.Println(bestStation, radioSet[bestStation], states)
		finalRadioSet = append(finalRadioSet, bestStation)
		delete(radioSet, bestStation)
	}

	return finalRadioSet

}

func generateRadioAndStateSet() map[string][]string {
	radioSet := map[string][]string{
		"KONE":   []string{"ID", "NV", "UT"},
		"KTWO":   []string{"WA", "ID", "MT"},
		"KTHREE": []string{"OR", "NV", "CA"},
		"KFOUR":  []string{"NV", "UT"},
		"KFIVE":  []string{"CA", "AZ"},
	}
	return radioSet
}

func Run() {
	// fmt.Println(arrangeClass(generateClassSchedule()))
	fmt.Println(findMinumumRadioSet(generateRadioAndStateSet(),
		[]string{"ID", "NV", "UT", "WA", "MT", "OR", "CA", "AZ"}))
}
