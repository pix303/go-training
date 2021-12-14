package main

import (
	"fmt"
	"math"
	"time"
)


func main(){

	presentTime:=time.Now()
	fmt.Printf("Present time (no format) %v in location %v\n",presentTime,presentTime.Location())

	locEuRome,_ :=time.LoadLocation("Europe/Rome") 
	locUSLA,_ :=time.LoadLocation("America/Los_Angeles") 
	fmt.Printf("Present time (no format) %v in location %v\n",presentTime.In(locEuRome),locEuRome)

	// Jan		month name short
	// January	month name complete
	// Mon		day name short
	// Monday	day name complete
	// 2006		year
    // 01		month
	// 02		day
	// 15		hours
	// 3PM		hours in AM/PM format
	// 04		minutes
	// 05		seconds
	// .000x	milliseconds
	// .999x	milliseconds
	fmt.Println(presentTime.Format("date: 02/01/2006 hour: 15:04:05.999"))
	fmt.Println(presentTime.Format("2006-01-02 Mon 3PM 04min"))
	fmt.Println(presentTime.Format("January 02 2006 Mon 3PM 04min"))
	fmt.Println(presentTime.In(locUSLA).Format("January 02 2006 Mon 3PM 04min"))

	
	hh, mm, ss := presentTime.Clock()
	fmt.Printf("The time is %d:%d:%d\n", hh, mm , ss)

	startTime := time.Now()
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(time.Since(startTime).String())

	bornDate := time.Date(1973, time.March, 7, 9, 30, 0, 0, locEuRome)
	birthdaty2022Date := time.Date(2022, time.March, 7, 9, 30, 0, 0, locEuRome)
	lifeDays := math.Floor(birthdaty2022Date.Sub(bornDate).Hours() / (24 * 365))
	fmt.Println(lifeDays)
}