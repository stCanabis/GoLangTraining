package main

import "fmt"

func main()  {
	slice := []int{1,2,3,4,5}
	var data1 string =  "test string"
	data2 := &data1
	data3 := &slice[3]
	fmt.Println("1",slice[3])
	fmt.Println("2",data1)
	fmt.Println("3",data2)
	fmt.Println("4",*data2)
	fmt.Println("5",data3)
	fmt.Println("6",*data3)

}
