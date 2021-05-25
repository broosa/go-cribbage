package main

import (
	"fmt"
)

func main() {
	fmt.Println(GetSubsetSums(15, 5, []uint{1, 2, 3, 3, 9}))
	fmt.Println(GetCribbageCardFromString("8S"))
}

//Pass a sorted slice as set
func GetSubsetSums(sum uint, setSize uint, set []uint) [][]uint {
	fmt.Println("Getting subsets", sum, setSize, set)
	sums := make([][]uint, 0)
	for i, n := range set {
		if n > sum {
			continue
		}

		if n == sum {
			sums = append(sums, []uint{n})
			continue
		}

		otherItems := make([]uint, 0)
		copy(otherItems, set[:i])
		if i < len(set)-1 {
			otherItems = append(otherItems, set[i+1:]...)
		}

		if setSize > 2 {
			smallerSums := GetSubsetSums(sum-n, setSize-1, otherItems)
			for _, smallSubset := range smallerSums {
				sums = append(sums, append(smallSubset, n))
			}
			continue
		}

		//Find all 2-ples {a, b} such that a + b == sum
		doubles := make([][]uint, 0)
		for _, d := range set[:len(set)-2] {
			if d+n == sum {
				doubles = append(doubles, []uint{n, d})
			}
		}
		fmt.Println("Doubles", sum, setSize, set, doubles)
		sums = append(sums, doubles...)

	}
	fmt.Println("Returning", sum, setSize, set, sums)
	return sums
}
