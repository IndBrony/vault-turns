package vaultturns

import (
	"strconv"
	"testing"
)

func TestVaultTurns(t *testing.T) {
	templateTestVaultTurns(t, []int{1, 2, 3, 4, 5}, 0)
	templateTestVaultTurns(t, []int{2, 1, 2, 1}, 3)
	templateTestVaultTurns(t, []int{9, 8, 1}, 2)
	templateTestVaultTurns(t, []int{5, 0, 9, 1, 8, 9}, 4)
	templateTestVaultTurns(t, []int{5, 0, 9, 1, 8, 9, 6, 1, 4}, 6)
	templateTestVaultTurns(t, []int{}, 0)
	templateTestVaultTurns(t, []int{1}, 0)
}
func TestVaultTurnsString(t *testing.T) {
	templateTestVaultTurnsString(t, "12345", 0, nil)
	templateTestVaultTurnsString(t, "2121", 3, nil)
	templateTestVaultTurnsString(t, "981", 2, nil)
	templateTestVaultTurnsString(t, "509189", 4, nil)
	templateTestVaultTurnsString(t, "509189614", 6, nil)
	templateTestVaultTurnsString(t, "", 0, nil)
	templateTestVaultTurnsString(t, "a", -1, strconv.ErrSyntax)
	templateTestVaultTurnsString(t, "123567123567a", -1, strconv.ErrSyntax)

}

func templateTestVaultTurns(t *testing.T, nums []int, expectedReturn int) {
	actual := VaultTurns(nums)
	if expectedReturn != actual {
		t.Errorf("testing fail, expected : %v but got %v", expectedReturn, actual)
	}
}

func templateTestVaultTurnsString(t *testing.T, nums string, expectedReturn int, expectedError error) {
	actual, actualErr := VaultTurnsString(nums)
	if expectedReturn != actual {
		t.Errorf("testing fail, expected return : %v and error : %s but got %v and %s", expectedReturn, expectedError, actual, actualErr)
	}
}
