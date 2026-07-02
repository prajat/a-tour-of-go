package main

import (
	"fmt"
	"runtime"
	"time"
)

func operatingSystem() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macos")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s. \n", os)
	}
}

// Switch cases evaluate cases from top to bottom, stopping when a case succeeds.

func whenSaturday() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Printf("Today is %s. \n", today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("In one Day or Tomorrow.")
	case today + 2:
		fmt.Println("In two Days.")
	default:
		fmt.Println("Too far away.")
	}
}

func greet() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning.")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon.")
	default:
		fmt.Println("Good Evening.")
	}
}
