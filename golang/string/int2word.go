package string

import (
	"fmt"
)

var Max = 1<<31 - 1
var Billion = 1000000000
var Million = 1000000
var Thousand = 1000

func int2word(n int) string {
	n20 := []string{
		"",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"fifteen",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen",
	}

	n100 := []string{
		"",
		"",
		"twenty",
		"thirty",
		"fourty",
		"fifty",
		"sixty",
		"seventy",
		"eighty",
		"ninety",
	}

	i2w := func(n int) string {
		if n < 20 {
			return n20[n]
		}
		if n < 100 {
			s := n100[n/10]
			if n == 20 {
				return s
			}
			return s + " " + n20[n%10]
		}
		return ""
	}

	i1002w := func(n int) string {
		s := i2w(n)
		if n < 100 {
			return s
		}
		return n20[n/100] + " hundred " + i2w(n%100)
	}

	if n > Max {
		return "not supported"
	}

	res := ""
	x := n / Billion
	if x > 0 {
		res += i1002w(x) + " billion "
	}
	x = n % Billion
	x = x / Million
	if x > 0 {
		res += i1002w(x) + " million "
	}

	x = n % Billion % Million / Thousand
	if x > 0 {
		res += i1002w(x) + " thousand "
	}

	x = n % Billion % Million % Thousand
	fmt.Println("x", x)

	if x > 0 {
		res += i1002w(x)
	}

	return res
}
