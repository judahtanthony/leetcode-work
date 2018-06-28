package main

import (
	"bytes"
	"fmt"
)

func intToRoman(num int) string {
	var buffer bytes.Buffer
	for num > 0 {
		switch {
		case num >= 1000:
			buffer.WriteByte('M')
			num -= 1000
		case num >= 900:
			buffer.WriteByte('C')
			buffer.WriteByte('M')
			num -= 900
		case num >= 500:
			buffer.WriteByte('D')
			num -= 500
		case num >= 400:
			buffer.WriteByte('C')
			buffer.WriteByte('D')
			num -= 400
		case num >= 100:
			buffer.WriteByte('C')
			num -= 100
		case num >= 90:
			buffer.WriteByte('X')
			buffer.WriteByte('C')
			num -= 90
		case num >= 50:
			buffer.WriteByte('L')
			num -= 50
		case num >= 40:
			buffer.WriteByte('X')
			buffer.WriteByte('L')
			num -= 40
		case num >= 10:
			buffer.WriteByte('X')
			num -= 10
		case num >= 9:
			buffer.WriteByte('I')
			buffer.WriteByte('X')
			num -= 9
		case num >= 5:
			buffer.WriteByte('V')
			num -= 5
		case num >= 4:
			buffer.WriteByte('I')
			buffer.WriteByte('V')
			num -= 4
		default:
			buffer.WriteByte('I')
			num--
		}
	}
	return buffer.String()
}

func main() {
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(9))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1994))
}
