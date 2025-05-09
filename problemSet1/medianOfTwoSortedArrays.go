package problemSet1

// numsArr1 and numsArr2 are sorted
func medianOfTwoSortedArrays(numsArr1 []int, numsArr1Len int, numsArr2 []int, numsArr2Len int) float64 {
	// container for sorted placements
	var mergedArrSorted []int

	// sum of the length of both numsArr
	numsArrLenSum := numsArr1Len + numsArr2Len

	// instead of removing the current rear element after each iteration,
	// in order to have a new rear element for the target numsArr,
	// just point to (keep the index of) the current rear element for target numsArr
	// and move the pointer to the new rear element after each iteration
	// by increamenting the index
	numsArr1RearIndex, numsArr2RearIndex := 0, 0

	for range numsArrLenSum {
		var lowestRearValue int

		if numsArr1RearIndex != numsArr1Len && numsArr2RearIndex != numsArr2Len {
			numsArr1RearValue := numsArr1[numsArr1RearIndex]
			numsArr2RearValue := numsArr2[numsArr2RearIndex]

			// compare the current rear elements of both numsArr:
			// to find the one with the lowest value.
			//
			// Whichever target numsArr has the lowest value at its rear,
			// move the pointer to the rear futher by one
			// (making it seem like we removed it OR we shift the array one element backwards)
			if numsArr1RearValue <= numsArr2RearValue {
				lowestRearValue = numsArr1RearValue
				numsArr1RearIndex++
			} else {
				lowestRearValue = numsArr2RearValue
				numsArr2RearIndex++
			}

			// If numsArr2 has reached its end or is empty, no comparison is needed,
			// numsArr1's rear element is the lowest value
		} else if numsArr1RearIndex != numsArr1Len && numsArr2RearIndex == numsArr2Len {
			numsArr1RearValue := numsArr1[numsArr1RearIndex]
			lowestRearValue = numsArr1RearValue
			numsArr1RearIndex++

			// If numsArr1 has reached its end or is empty, no comparison is needed,
			// numsArr2's rear element is the lowest value
		} else if numsArr1RearIndex == numsArr1Len && numsArr2RearIndex != numsArr2Len {
			numsArr2RearValue := numsArr2[numsArr2RearIndex]
			lowestRearValue = numsArr2RearValue
			numsArr2RearIndex++

			// Practically, we can only get here if both numsArr are empty
			// Otherwise, we can't reach the ends of both numsArr
			// before we arrive at our final answer
		} else {
			return 0
		}

		// lowest-to-highest value placement over multiple iterations
		mergedArrSorted = append(mergedArrSorted, lowestRearValue)

		midIndex := numsArrLenSum / 2

		// this loop will repeat until we arrive at half of (numsArr1Len + numsArr2Len)
		// this is what makes the O(log (m + n)) time complexity happen
		// since, for any input sizes of both numsArr,
		// the iteration will always have its final answer
		// at the half of the sum of both input sizes
		//
		// When we're at the midIndex, just do the median logic
		if (len(mergedArrSorted) - 1) == midIndex {
			if (numsArrLenSum % 2) == 0 {
				finalAnswer := float64((mergedArrSorted[midIndex-1] + mergedArrSorted[midIndex])) / 2.0

				return finalAnswer
			} else {
				finalAnswer := float64(mergedArrSorted[midIndex])

				return finalAnswer
			}
		}

		continue
	}

	return 0
}
