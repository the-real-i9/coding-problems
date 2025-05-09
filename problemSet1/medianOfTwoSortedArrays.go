package problemSet1

func MedianOfTwoSortedArrays(numsArr1 []int, numsArr1Len int, numsArr2 []int, numsArr2Len int) float64 {
	var mergedArr []int

	numsArrLenSum := numsArr1Len + numsArr2Len

	numsArr1RearIndex, numsArr2RearIndex := 0, 0

	for i := 0; i < numsArrLenSum; i++ {
		var lowestRearValue int

		if numsArr1RearIndex != numsArr1Len && numsArr2RearIndex != numsArr2Len {
			numsArr1RearValue := numsArr1[numsArr1RearIndex]
			numsArr2RearValue := numsArr2[numsArr2RearIndex]

			if numsArr1RearValue <= numsArr2RearValue {
				lowestRearValue = numsArr1RearValue
				numsArr1RearIndex++
			} else {
				lowestRearValue = numsArr2RearValue
				numsArr2RearIndex++
			}
		} else if numsArr1RearIndex != numsArr1Len && numsArr2RearIndex == numsArr2Len {
			numsArr1RearValue := numsArr1[numsArr1RearIndex]
			lowestRearValue = numsArr1RearValue
			numsArr1RearIndex++
		} else if numsArr1RearIndex == numsArr1Len && numsArr2RearIndex != numsArr2Len {
			numsArr2RearValue := numsArr2[numsArr2RearIndex]
			lowestRearValue = numsArr2RearValue
			numsArr2RearIndex++
		} else {
			return 0
		}

		mergedArr = append(mergedArr, lowestRearValue)

		midIndex := numsArrLenSum / 2

		if (len(mergedArr) - 1) == midIndex {
			if (numsArrLenSum % 2) == 0 {
				finalAnswer := float64((mergedArr[midIndex-1] + mergedArr[midIndex])) / 2.0

				return finalAnswer
			} else {
				finalAnswer := float64(mergedArr[midIndex])

				return finalAnswer
			}
		}

		continue
	}

	return 0
}
