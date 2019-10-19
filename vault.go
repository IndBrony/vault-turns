package vaultturns

import (
	"strconv"
	"strings"
)

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

func VaultTurns(nums []int) int {
	vaultNums := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	counter := 0
	prevTurn := "left"
	halfTurn := len(vaultNums) / 2

	for destIndex := 1; destIndex < len(nums); destIndex++ {
		currIndex := destIndex - 1
		gap := nums[destIndex] - nums[currIndex]
		var turn string

		if gap%halfTurn == 0 {
			continue
		}
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
