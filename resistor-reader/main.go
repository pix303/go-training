package main

import (
	"fmt"
	"math"
)

var resistorColor []string = []string{
	 "black",
  "brown",
  "red",
  "orange",
  "yellow",
  "green",
  "blue",
  "violet",
  "grey",
  "white",
} 

func main(){
	fmt.Println(decodeResistorColor([]string{"red","black","red"}))
	fmt.Println(decodeResistorColor([]string{"red","black","brown"}))
	fmt.Println(decodeResistorColor([]string{"red","black","black"}))
	fmt.Println(decodeResistorColor([]string{"green","brown","orange"}))
}

func decodeResistorColor( colors []string) string{
	// first digit
	c1 := getValueByColor(colors[0])
	// second digit
	c2 := getValueByColor(colors[1])
	// ten multiplier
	c3 := getValueByColor(colors[2])

	value := (c1 * 10 + c2) * int(math.Pow10(c3))

	if (value >= 1000){
		return fmt.Sprintf("%d kiloohms",value/1000)
	}
	return fmt.Sprintf("%d ohms",value)
}

func getValueByColor(color string) int{
	for i,c := range resistorColor{
		if c == color {
			return i
		}
	}
	return -1
}