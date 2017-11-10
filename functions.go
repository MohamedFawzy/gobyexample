package main

import "fmt"

func plus(a int, b int)int  {
	return a + b
}

func plusPlus(a, b, c int)int  {
	return a+b+c
}

func vals() (int, int)  {
	return 3, 7
}

func sum(nums...int)  {

	fmt.Println(nums, " ")

	total := 0

	for _, num := range nums{
		total += num
	}

	fmt.Println(total)
}


func main() {

	res := plus(1,5)
	fmt.Println("1+5:=", res)

	res = plusPlus(1,24,4)
	fmt.Println("1+24+4", res)

	a, b := vals()
	_ , c := vals()
	fmt.Println(a,b)
	fmt.Println(c)


	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1,2,3,4}
	sum(nums...)

}