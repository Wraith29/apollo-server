package data

import (
	"strings"
)

// https://en.wikipedia.org/wiki/Levenshtein_distance#Iterative_with_two_matrix_rows
func LevenshteinDistance(left, right string) int {
	left, right = strings.ToLower(left), strings.ToLower(right)

	if left == right {
		return 0
	}

	if len(right) > len(left) {
		left, right = right, left
	}

	prevDist := make([]int, len(right)+1)
	currDist := make([]int, len(right)+1)

	for i := 0; i < len(prevDist); i++ {
		prevDist[i] = i
	}

	for i := 0; i < len(left); i++ {
		currDist[0] = i + 1

		for j := 0; j < len(right); j++ {
			deletionCost := prevDist[j+1] + 1
			insertionCost := currDist[j] + 1
			subCost := prevDist[j]

			if left[i] != right[j] {
				subCost++
			}

			currDist[j+1] = min(deletionCost, insertionCost, subCost)
		}

		copy(prevDist, currDist)
		currDist = make([]int, len(right)+1)
	}

	return prevDist[len(right)]
}
