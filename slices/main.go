package main

import (
	"fmt"
)

func main() {
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	fmt.Println("Items:", items, "\tlen:", len(items), "\tcap:", cap(items))

	// add item
	fmt.Println("\nADD ITEM -------------------------------------------------------------")
	items = append(items, "i")
	fmt.Println("Items:", items, "\tlen:", len(items), "\tcap:", cap(items), "\titems for cap", items[:cap(items)-1])

	fmt.Println("\nSLICE OF SLICE -------------------------------------------------------------")
	bcItems := items[1:3]
	fmt.Println("bcItems:", bcItems, "\tlen:", len(bcItems), "\tcap:", cap(bcItems), bcItems[:cap(bcItems)-1])

	fmt.Println("\nREMOVE ITEM -----------------------------------------------------------")
	items, items[len(items)-1] = append(items[:3], items[3+1:]...), ""
	fmt.Println("Items:", items, "\tlen:", len(items), "\tcap:", cap(items), "\titems for cap", items[:cap(items)-1])

	fmt.Println("\nCHANGE A VALUE IN ARRAY -----------------------------------------------------------")
	bcItems[0] = "X"
	fmt.Println("Items:", items, "bcItems", bcItems)

	fmt.Println("\nCOPY SLICE -----------------------------------------------------------")
	copiedItems := make([]string, 2)
	copy(copiedItems, bcItems)
	fmt.Println("Copied Items:", copiedItems, "\tlen:", len(copiedItems), "\tcap:", cap(copiedItems), "\tbcItems", bcItems)
	fmt.Println("\nCHANGE A VALUE IN COPIED -----------------------------------------------------------")
	copiedItems[0] = "Q"
	fmt.Println("Copied Items:", copiedItems, "bcItems", bcItems)

	fmt.Println("\nMEM LEAK IN SLICE -----------------------------------------------------------")
	retrivedSliceWrong := getSliceWrong()
	fmt.Println("Retrived Slice:", retrivedSliceWrong, "\tlen:", len(retrivedSliceWrong), "\tcap:", cap(retrivedSliceWrong), retrivedSliceWrong[:cap(retrivedSliceWrong)-1])

	fmt.Println("\nAVOID MEM LEAK IN SLICE -----------------------------------------------------------")
	retrivedSlice := getSlice()
	fmt.Println("Retrived Slice:", retrivedSlice, "\tlen:", len(retrivedSlice), "\tcap:", cap(retrivedSlice), retrivedSlice[:cap(retrivedSlice)-1])

}

func getSlice() []int {
	nums := []int{1, 2, 3, 4, 5}
	result := make([]int, 2)
	copy(result, nums[:2])
	return result
}

func getSliceWrong() []int {
	nums := []int{1, 2, 3, 4, 5}
	return nums[:2]
}
