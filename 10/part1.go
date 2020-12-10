package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/diamondburned/aoc-2020/aocutil"
)

// This code to remain largely unaltered from the state it was used to solve,
// mostly for historical purposes.

const maxJoltDiff = 3

func main() {
	lines := aocutil.SplitFile("input", "\n")

	var jolts = make([]int, 0, len(lines))

	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			continue
		}
		jolts = append(jolts, n)
	}

	_, maxJolt := aocutil.MinMaxes(jolts)
	maxJolt += maxJoltDiff

	// var arrangements int

	var outletJolts = append([]int{0}, jolts...)
	sort.Ints(outletJolts)

	// var permus = permute(outletJolts)
	// var checked = map[[100]int]struct{}{}

	// // somehow find the combinations shorter than outletJolts while removing
	// // some?
	// for _, permu := range permus {
	// 	var nums [100]int
	// 	copy(nums[:], permu)
	// 	if _, ok := checked[nums]; ok {
	// 		continue
	// 	}
	// 	checked[nums] = struct{}{}
	// 	// Check if this is a valid combination.
	// 	if isValidCombo(permu, maxJolt) {
	// 		log.Println("valid:", permu)
	// 		arrangements++
	// 	}
	// }

	// fmt.Println(arrangements)
	// fmt.Println(nPermute1(outletJolts[1 : len(outletJolts)-1]))
	// fmt.Println(nPermute2(outletJolts[1 : len(outletJolts)-1]))

	fmt.Println(shortestCombo(outletJolts))

	permute(outletJolts)
	// for _, permu := range permute(outletJolts) {
	// 	fmt.Println(permu)
	// }
}

func nPermute(jolts []int) int {
	var nPermute int

	for i := 0; i < len(jolts); {
		next := nextCombos(jolts, i)
		if len(next) == 0 {
			break
		}

		i = next[len(next)-1]
		nPermute += len(next)
	}

	return nPermute
}

func permute(jolts []int) [][]int {
	var permus [][]int

	for i := 0; i < len(jolts); {
		next := nextCombos(jolts, i)
		if len(next) == 0 {
			break
		}

		permus = append(permus, permuteNext(permus, jolts, next)...)
		i = next[len(next)-1]
	}

	return permus
}

func permuteNext(pastPerms [][]int, jolts []int, nextIxs []int) [][]int {
	var permus [][]int

	log.Println("-->", batchAccess(jolts, nextIxs))

	for i := 0; i < len(nextIxs); i++ {
		var head = jolts[:nextIxs[i]]

		for k := 0; k < nextIxs[i]; {
			var next = nextCombos(head, k)

			log.Println("i    -->", nextIxs[i])
			log.Println("next -->", next)
			var heads = permuteNextNoRec(head, next)

			log.Println("head -->", head, heads)

			for _, head := range heads {
				for j := i; j < len(nextIxs); j++ {
					perm := append(append([]int(nil), head...), jolts[nextIxs[j]:]...)
					if !hasJoltsEq(pastPerms, perm) {
						permus = append(permus, perm)
					}
				}
			}

			k = nextIxs[len(nextIxs)-1]
		}
	}

	return permus
}

func permuteNextNoRec(jolts []int, nextIxs []int) [][]int {
	var permus = [][]int{jolts}
	if len(nextIxs) == 0 {
		return permus
	}

	var head = jolts[:nextIxs[0]]

	for _, nextIx := range nextIxs {
		perm := append(append([]int(nil), head...), jolts[nextIx+1:]...)
		permus = append(permus, perm)
	}
	return permus
}

func batchAccess(jolts []int, ixs []int) []int {
	var access = make([]int, len(ixs))
	for i, ix := range ixs {
		access[i] = jolts[ix]
	}
	return access
}

func hasJoltsEq(allJolts [][]int, jolts []int) bool {
	for i := len(allJolts) - 1; i >= 0; i-- {
		if joltsEq(allJolts[i], jolts) {
			return true
		}
	}
	return false
}

func joltsEq(jolts1, jolts2 []int) bool {
	if len(jolts1) != len(jolts2) {
		return false
	}
	for i := range jolts1 {
		if jolts1[i] != jolts2[i] {
			return false
		}
	}
	return true
}

func shortestCombo(jolts []int) (combo []int) {
	for i := 0; i < len(jolts); {
		next := nextCombos(jolts, i)
		if len(next) == 0 {
			break
		}

		i = next[len(next)-1]
		combo = append(combo, jolts[i])
	}

	return combo
}

func nextCombos(jolts []int, currentIx int) (nextIxs []int) {
	for i := currentIx + 1; i < len(jolts); i++ {
		if jolts[i]-jolts[currentIx] > maxJoltDiff {
			return
		}
		nextIxs = append(nextIxs, i)
	}
	return
}

func isValidCombo(jolts []int, max int) bool {
	if jolts[len(jolts)-1]+3 < max {
		return false
	}

	var lastJolt = jolts[0]

	for i := 1; i < len(jolts); i++ {
		var difference = jolts[i] - lastJolt
		if difference > maxJoltDiff {
			return false
		}

		lastJolt = jolts[i]
	}

	return true
}
