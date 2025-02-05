package main

import "fmt"

func main() {
	n := 12345
	fact := cal(n)
	fmt.Println(fact)
}

func cal(n int) int {
	ds := 0
	for n >= 1 {
		digit := n % 10
		ds = ds + digit
		n = n / 10
	}
	return ds
}

/*

echo "<pre>";

$n = 12345;
echo fact($n);

function fact($n) {
$ds = 0;
while ($n >= 1) {

 $digit = $n % 10;
 $ds = $ds + $digit ;

 $n = intval($n / 10);
}

return $ds;
}

*/
