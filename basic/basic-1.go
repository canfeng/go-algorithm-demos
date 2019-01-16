package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

func main() {
	s := IntToString(143)
	log.Printf("res=>%s", s)
	s1 := IntToString(-143)
	log.Printf("res=>%s", s1)
	s2 := IntToString(0)
	log.Printf("res=>%s", s2)
}

func Run() {
	for n := 1; n < 11; n++ {
		start := time.Now().UnixNano()
		num := int(math.Pow10(n))

		var total int
		for i := 0; i < num; i++ {
			total += i
		}
		end := time.Now().UnixNano()
		log.Printf("n=%d - total=%d - total used %f s", n, total, float64(end-start)/1000000000)
	}
}

//O(logN)
func IntToString(num int) string {
	var s string
	var prefix = ""
	if num == 0 {
		return "0"
	}
	if num < 0 {
		prefix = "-"
		num = int(math.Abs(float64(num)))
	}
	for num > 0 {
		s += fmt.Sprintf("%d", num%10)
		num /= 10
	}
	return prefix + Reverse(s)
}

//反转字符串
//1/2n * O(n)
func Reverse(s string) string {
	var n = len(s)

	var arr = []byte(s)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-1-i] = s[n-1-i], s[i]
	}
	return string(arr[:])
}

//判断是否是素数（质数）
//O(sqrt(n))
func IsPrime(num int) bool {
	if num == 0 || num == 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
