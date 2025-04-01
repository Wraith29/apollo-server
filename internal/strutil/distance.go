package strutil

import (
	"strings"
)

func calculateDistance(left, right string) int {
	previousDistances := make([]int, len(right)+1)
	currentDistances := make([]int, len(right)+1)

	for i := range previousDistances {
		previousDistances[i] = i
	}

	for i := range len(left) {
		currentDistances[0] = i + 1

		for j := range len(right) {
			deletionCost := previousDistances[j+1] + 1
			insertionCost := currentDistances[j] + 1

			substitutionCost := previousDistances[j]
			if left[i] != right[j] {
				substitutionCost += 1
			}

			currentDistances[j+1] = min(deletionCost, insertionCost, substitutionCost)
		}

		previousDistances = currentDistances
	}

	return previousDistances[len(right)-1]
}

func Distance(left, right string) int {
	if len(right) > len(left) {
		left, right = right, left
	}

	return calculateDistance(strings.ToLower(left), strings.ToLower(right))
}
