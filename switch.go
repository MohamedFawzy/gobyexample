package main

import "fmt"
import "time"

func main() {

	i:=2
	fmt.Println("Write ", i , " as")

	// input : 2017-10-17T09:29:58+02:00
	// output: 2017-10-19T11:40:37+0000
	fmt.Println("time", time.Now().UTC().Format("2006-01-02T15:04:05-0700"))

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's weekend")
	default:
		fmt.Println("it's a weekday")
	}

	t := time.Now()

	switch  {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}


	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("i'm bool")
		case int:
			fmt.Println("i'm int")

		default:
			fmt.Println("Don't know type  ", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")


}
