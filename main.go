package main

import "fmt"

/*
	func main() {
		var a int
		fmt.Scan(&a)

		if a%10 == 1 && a%100 != 11 {
			fmt.Printf("%d korova", a)
		} else if (a%10 == 2 || a%10 == 3 || a%10 == 4) && !(a%100 == 12 || a%100 == 13 || a%100 == 14) {
			fmt.Printf("%d korovy", a)
		} else {
			fmt.Printf("%d korov", a)
		}

}

	func main() {
		var n int
		fmt.Scan(&n)
		for i := 1; i < n; {
			fmt.Printf("%d ", i)
			i = i * 2
		}
	}

	func main() {
		var n, a, sum, numb, min int
		fmt.Scan(&numb)
		min = -1
		a, n = 1, 1
		for i := 1; i <= numb; i++ {
			sum = a + n
			a = n
			n = sum
			if sum == numb {
				fmt.Println(i + 2)
				break
			} else if numb == 1 {
				fmt.Println("1")
				break
			} else if sum > numb {
				fmt.Println(min)
				break
			}

		}
	}

	func main() {
		var n int
		fmt.Scan(&n)
		fmt.Printf("%b", n)

}
func main() {
	var num string
	var n string
	fmt.Scan(&num, &n)
	result := strings.ReplaceAll(num, n, "")
	fmt.Println(result)

}
func main() {
	minimumFromFour()

}
func minimumFromFour() int {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	return min(a, b, c, d)
}
func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	fmt.Println(vote(x, y, z))

}
func vote(x int, y int, z int) int {
	var res int
	if x == y || x == z {
		res = x
	} else if y == z {
		res = y
	}
	return res
}*/
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(fibonacci(n))
}
func fibonacci(n int) int {
	var sum = 1
	var a = 1
	var j = 1
	for i := 1; i < n-1; i++ {
		sum = a + j
		a = j
		j = sum
	}
	return sum
}
