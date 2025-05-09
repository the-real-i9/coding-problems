package main

// Arguments: arr1[], arr1Length, arr2[], arr2Length

/*
* Create an array, mergedArr[], to store the merge of the two arrays

* Create a variable, arr1IndexVisited and arr2IndexVisited
* for both arr1[] and arr2[] that stores the index being visited

* Create a for loop that will run for arr1Length + arr2Length

	* Compare the values at the arrNIndexVisited of arr1[] and arr2[],
	* to find the lowest value.
	* However, if the arrNIndexVisited of one arrN[] == arrNLength,
	* don't bother finding the lowest, just choose the value at the
	* arrNIndexVisited of the other arrN[]

	* Push this value to mergedArr[], then for the arrN[] in which the lowest value is found,
	* increment its arrNIndexVisited by 1

	* Calculate, midIndex := math.trucate((arr1Length + arr2Length) / 2)

	* If: the last element of mergedArr[] is at midIndex
	* i.e the (len(mergedArr) - 1) is midIndex

		* If: (arr1Length + arr2Length) is even i.e. (arr1Length + arr2Length) % 2 == 0
			* Then, medianOfTwoSA := (mergedArr[midIndex - 1] + mergedArr[midIndex]) / 2

		* Else If: (arr1Length + arr2Length) is odd i.e. (arr1Length + arr2Length) % 2 == 1
			* Then, medianOfTwoSA := mergedArr[midIndex]

	* Else:
			* Then for the arrN[] in which the lowest value is found,
			* increment its arrNIndexVisited by 1

			* continue the for loop
*/

/*
* To calculate the middle index:

  - midIndex :=
  - If: (arr1Length + arr2Length) is even i.e. (arr1Length + arr2Length) % 2 == 0
  - Then, [midIndex - 1, midIndex] is your middle index
  - Else If: (arr1Length + arr2Length) is odd i.e. (arr1Length + arr2Length) % 2 == 1
  - Then, midIndex is your middle index
*/
func medianOfTwoSortedArrays() {

}
