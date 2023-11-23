package main

import "fmt"

func main() {

	//interface-array
	arr := [4]interface{}{1, "aa", 2.5, nil}
	fmt.Println(arr)
	fmt.Println(arr[2])

	//interface-slice
	slice := []interface{}{1, "aa", 2.5, nil, true}
	fmt.Println(slice)
	fmt.Println(slice[2])

	//function
	show(slice)
}

func show(b []interface{}) {
	for i := 0; i < len(b); i++ {
		fmt.Print(b[i], " ")
	}
}

func main1() {

	emp := map[string]interface{}{}
	emp["name"] = "rahul"
	emp["handicap"] = false
	emp["city"] = "delhi"
	emp["zip"] = 110056
	fmt.Println(emp)             //map[city:delhi name:rahul zip:110056]
	fmt.Println(emp["handicap"]) //false
	fmt.Printf("%T", emp)        //map[string]interface {}
	fmt.Println()

	dep := map[string]string{}
	dep["name"] = "rahul"
	dep["handicap"] = "false"
	dep["city"] = "delhi"
	dep["zip"] = "110056"
	fmt.Println(dep)             //map[city:delhi name:rahul zip:110056]
	fmt.Println(dep["handicap"]) //false
	fmt.Printf("%T", dep)        //map[string]string

}

func main2() {

	var val1 interface{} = 6.5
	val11 := val1.(float64)

	var val2 interface{} = "aa"
	val22 := val2.(string)

	fmt.Println("Value: ", val1.(float64), val11)
	fmt.Println("Value: ", val2.(string), val22)
}
