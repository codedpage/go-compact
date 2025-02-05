// You can edit this code!
// Click here and start typing.
package main

import "fmt"

var fact int

func main() {

	n := 5
	fact = cal(n)
	fmt.Println(fact)

}

// without recursion
func cal(n int) int {

	fact := 1
	for i := 1; i <= n; i++ {
		fact = fact * i
	}
	return fact
}

// with recursion
func cal2(n int) int {

	if n == 1 {
		fmt.Println("........", n)
		return 1
	}

	fmt.Println("--------------------", n)
	n = n * cal(n-1)
	return n

}

/*
$n = 5;
echo fact($n);

function fact($n) {
	if ($n == 1) {
		return 1;
	} else {
		return $n = $n * fact($n-1);
	}
}
*/

/*
-------------------- 5
-------------------- 4
-------------------- 3
-------------------- 2
........ 1
120
*/
