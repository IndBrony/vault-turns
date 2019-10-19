package vaultturns

import (
	"strconv"
	"strings"
)

//VaultTurnsString is Wrapper to call VaultTurns with string input
func VaultTurnsString(s string) (int, error) {
	stringSlice := strings.Split(s, "")
	nums := make([]int, len(stringSlice))
	for index, char := range stringSlice {
		converted, err := strconv.Atoi(char)
		if err != nil {
			return -1, err
		}
		nums[index] = converted
	}

	return VaultTurns(nums), nil
}

//VaultTurns counts the number of changing turn when opening a circular vault key combination
func VaultTurns(nums []int) int {
	//for future use tobe more scalable
	vaultNums := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	counter := 0
	prevTurn := "left"
	halfTurn := len(vaultNums) / 2

	//comparing each number with it's side using loop
	//loop from index 1 to avoid IndexOutOfRange error
	for destIndex := 1; destIndex < len(nums); destIndex++ {
		currIndex := destIndex - 1
		gap := nums[destIndex] - nums[currIndex]
		var turn string

		//if the gap is 0 or half or the same as the length of vaultNums, it wont increase the counter
		if gap%halfTurn == 0 {
			continue
		}

		//Ithis point forward, the code speaks for itself
		if gap > halfTurn {
			turn = "right"
		} else if gap < -1*halfTurn {
			turn = "left"
		} else if gap < 0 {
			turn = "right"
		} else {
			turn = "left"
		}

		if prevTurn != turn {
			counter++
			prevTurn = turn
		}
	}
	return counter
}
