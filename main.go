package main

import "fmt"

/*
	func main() {
		var workArray [10]uint8
		var x, y, a uint8
		for i := 0; i < len(workArray); i++ {
			fmt.Scan(&workArray[i])

		}
		for i := 0; i < 3; i++ {
			for i=0;i<3;i++ {
	    fmt.Scan(&a, &b)				//Менять места в массиве можно так!!!!
	    w[a], w[b] = w[b], w[a]
		fmt.Scan(&x, &y)
		a = workArray[y]
		workArray[y] = workArray[x]
		workArray[x] = a
		}
		for i := 0; i < len(workArray); i++ {
			fmt.Print(workArray[i], " ")
		}
	}
	func main() {
		var num1, num2 string
		var result []string
		fmt.Scan(&num1, &num2)
		for i := 0; i < len(num1); i++ {
			digit1 := num1[i]
			found := false
			for j := 0; j < len(num2); j++ {
				if digit1 == num2[j] {
					found = true
					break
				}
			}
			if found {
				result = append(result, string(digit1))
			}
		}
		if len(result) > 0 {
			fmt.Printf(strings.Join(result, " "))
		} else {
			fmt.Println()
		}

}

	func main() {
		var n int
		fmt.Scan(&n)

		slice := make([]int, n)

		for i := 0; i < n; i++ {
			fmt.Scan(&slice[i])
		}

		fmt.Println(slice[3])

}

	func main() {
		array := [5]int{}
		var a int
		for i := 0; i < 5; i++ {
			fmt.Scan(&a)
			array[i] = a
		}
		var max int = array[0]
		for i := 0; i < len(array); i++ {
			if max < array[i] {
				max = array[i]
			}
		}
		fmt.Println(max)
	}

	func main() {
		var n int
		fmt.Scan(&n)

		slice := make([]int, n)

		for i := 0; i < n; i++ {
			fmt.Scan(&slice[i])
		}
		for i := 0; i < len(slice); i++ {
			if i%2 == 0 {
				fmt.Print(slice[i], " ")
			}

		}
	}

	func main() {
		var n int
		fmt.Scan(&n)
		var result int = 0

		arr := make([]int, n)

		for i := 0; i < n; i++ {
			fmt.Scan(&arr[i])
		}
		for i := 0; i < len(arr); i++ {
			if arr[i] > 0 {
				result++
			}
		}
		fmt.Println(result)

}

	func main() {
		var a, b, c, d int
		fmt.Scan(&a)
		b = a / 100
		c = (a / 10) % 10
		d = a % 10
		fmt.Println(b + c + d)

}

	func main() {
		var a, b, c int
		Scanf("%1d%1d%1d", &a, &b, &c)
		Printf("%d%d%d", c, b, a)
	}

	func main() {
		var a, b, c int
		Scan(&a, &b, &c)
		if a*a+b*b == c*c {
			Print("Прямоугольный")
		} else {
			Print("Непрямоугольный")
		}

}

	func main() {
		var a, b int
		Scan(&a, &b)
		Println(float64(float64(a+b)/2))

}

	func main() {
		var n, min, c, num int
		fmt.Scan(&n)
		fmt.Scan(&min)
		c = 1
		for i := 0; i < n-1; i++ {
			fmt.Scan(&num)
			if num == min {
				c++
			} else if num < min {
				min = num
				c = 1
			}
		}
		fmt.Println(c)

}

	func main() {
		var a, b int
		fmt.Scan(&a, &b)

		var result int

		if b >= 0 {

			result = (b / 7) * 7
		} else {
			result = (b / 7) * 7
			if result > b {
				result -= 7
			}
		}

		if result >= a {
			fmt.Println(result)
		} else {
			fmt.Println("NO")
		}
	}
func main() {
	var a int
	fmt.Scan(&a)

	if 5 > a && a < 20 {
		fmt.Printf("%d korov", a)
	} else if a%10 == 1 {
		fmt.Printf("%d korova", a)
	} else if a%10 == 2 || a%10 == 3 || a%10 == 4 {
		fmt.Printf("%d korovy", a)
	} else {
		fmt.Printf("%d korov", a)
	}

}
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
	}
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
*/
func main() {
	second := 8520
	var x, y, z int
	x = second / 3600
	y = (second % 3600) / 60
	z = second % 60
	fmt.Printf("%d часа %d минут %d секунд", x, y, z)
}
