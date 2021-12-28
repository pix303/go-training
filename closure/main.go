package main

import (
	"fmt"
)

func main(){
	names := []string{
		"Sergio",
		"Lorenzo",
		"Alessandro",
	}

	morning := welcome("buongiorno")
	afternoon := welcome("buon pomerggio")
	evening := welcome("buonasera")

	for _,n := range names{
		fmt.Println(morning(n))
		fmt.Println(afternoon(n))
		fmt.Println(evening(n))
		fmt.Println("-----------------")
	}
}


func welcome(msg string) func (string) string{
	welcomeMsg := msg
	return func (name string) string{
		return fmt.Sprintf("%s - %s -", welcomeMsg, name)
	}
}