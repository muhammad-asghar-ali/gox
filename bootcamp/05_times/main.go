package main

import (
	"fmt"
	"time"
)

func main() {
	wel := "Welcome to the time"
	fmt.Println(wel)

	curr := time.Now()
	fmt.Println("Current Time:", curr)
	fmt.Println("Current Time Format Only Date:", curr.Format("01-02-2006"))
	fmt.Println("Current Time Format Only Date with Day:", curr.Format("01-02-2006 Monday"))
	fmt.Println("Current Time Format Only Date with Day and time:", curr.Format("01-02-2006 15:04:05 Monday"))

	then := time.Date(2024, time.December, 13, 1, 30, 0, 0, time.UTC)
	fmt.Println("Created Date:", then)

	time_string := "2025-01-21 16:00:00"
	layout := "2006-01-02 15:04:05"
	pasred, err := time.Parse(layout, time_string)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Parsed Time:", pasred)

	later_1_hour := curr.Add(1 * time.Hour)
	fmt.Println("One hour later:", later_1_hour)

	ago_10_min := curr.Add(-10 * time.Minute)
	fmt.Println("Ten minutes ago:", ago_10_min)
}
