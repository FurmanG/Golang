package utile

import (
	"fmt"
	"strconv"
	"strings"
)

var WinPerMonth [12][69]int
var PowerballPerMoth [12][40]int

type WinNum struct {
	// The 5 Powerball wining numbers
	n1        int
	n2        int
	n3        int
	n4        int
	n5        int
	powerball int // The poer ball wining number
	count10   int // All games with 5 two-digit numbers
	count9    int // All games with 4 two-digit numbers and 1 one-digit
	count8    int // All games with 3 two-digit numbers and 2 one-digit
	count7    int // All games with 2 two-digit numbers and 3 one-digit
	count6    int // All games with 1 two-digit numbers and 4 one-digit
	count5    int // All games with 0 two-digit numbers and 5 one-digit
	date      string
}

// Extracting the winning numbers (int) from the total string of numbers.
func (n *WinNum) WinNumbers(winNumsStr string) {
	fmt.Println("Num")
	winNumsStrLength := len(winNumsStr)
	// 5 winning numbers of 2 digits each
	if winNumsStrLength == 10 {
		n.n1, _ = strconv.Atoi(string([]rune(winNumsStr)[0:2]))
		n.n2, _ = strconv.Atoi(string([]rune(winNumsStr)[2:4]))
		n.n3, _ = strconv.Atoi(string([]rune(winNumsStr)[4:6]))
		n.n4, _ = strconv.Atoi(string([]rune(winNumsStr)[6:8]))
		n.n5, _ = strconv.Atoi(string([]rune(winNumsStr)[8:10]))
		fmt.Println("The wining numbers 10; ", n.n1, n.n2, n.n3, n.n4, n.n5)
		n.count10++

	}

	// 4 winning numbers of 2 digits, and 1 number of 1 digit.
	if winNumsStrLength == 9 {
		n.n1, _ = strconv.Atoi(string([]rune(winNumsStr)[0:1]))
		n.n2, _ = strconv.Atoi(string([]rune(winNumsStr)[1:3]))
		n.n3, _ = strconv.Atoi(string([]rune(winNumsStr)[3:5]))
		n.n4, _ = strconv.Atoi(string([]rune(winNumsStr)[5:7]))
		n.n5, _ = strconv.Atoi(string([]rune(winNumsStr)[7:9]))
		fmt.Println("The wining numbers 9; ", n.n1, n.n2, n.n3, n.n4, n.n5)
		n.count9++

	}

	// 3 winning numbers of 2 digits, and 2 numbers of 1 digit.
	if winNumsStrLength == 8 {
		n.n1, _ = strconv.Atoi(string([]rune(winNumsStr)[0:1]))
		n.n2, _ = strconv.Atoi(string([]rune(winNumsStr)[1:2]))
		n.n3, _ = strconv.Atoi(string([]rune(winNumsStr)[2:4]))
		n.n4, _ = strconv.Atoi(string([]rune(winNumsStr)[4:6]))
		n.n5, _ = strconv.Atoi(string([]rune(winNumsStr)[6:8]))
		fmt.Println("The wining numbers 8; ", n.n1, n.n2, n.n3, n.n4, n.n5)
		n.count8++

	}

	// 2 winning numbers of 2 digits, and 3 numbers of 1 digit.
	if winNumsStrLength == 7 {
		n.n1, _ = strconv.Atoi(string([]rune(winNumsStr)[0:1]))
		n.n2, _ = strconv.Atoi(string([]rune(winNumsStr)[1:2]))
		n.n3, _ = strconv.Atoi(string([]rune(winNumsStr)[2:3]))
		n.n4, _ = strconv.Atoi(string([]rune(winNumsStr)[3:5]))
		n.n5, _ = strconv.Atoi(string([]rune(winNumsStr)[5:7]))
		fmt.Println("The wining numbers 7; ", n.n1, n.n2, n.n3, n.n4, n.n5)
		n.count7++

	}

	// 1 winning numbers of 2 digits, and 4 numbers of 1 digit.
	if winNumsStrLength == 6 {
		n.n1, _ = strconv.Atoi(string([]rune(winNumsStr)[0:1]))
		n.n2, _ = strconv.Atoi(string([]rune(winNumsStr)[1:2]))
		n.n3, _ = strconv.Atoi(string([]rune(winNumsStr)[2:3]))
		n.n4, _ = strconv.Atoi(string([]rune(winNumsStr)[3:4]))
		n.n5, _ = strconv.Atoi(string([]rune(winNumsStr)[4:6]))
		fmt.Println("The wining numbers 6; ", n.n1, n.n2, n.n3, n.n4, n.n5)
		n.count6++
	}

	// 5 numbers of 1 digit.
	if winNumsStrLength == 5 {
		n.n1, _ = strconv.Atoi(string([]rune(winNumsStr)[0:1]))
		n.n2, _ = strconv.Atoi(string([]rune(winNumsStr)[1:2]))
		n.n3, _ = strconv.Atoi(string([]rune(winNumsStr)[2:3]))
		n.n4, _ = strconv.Atoi(string([]rune(winNumsStr)[3:4]))
		n.n5, _ = strconv.Atoi(string([]rune(winNumsStr)[4:5]))
		fmt.Println("The wining numbers 5; ", n.n1, n.n2, n.n3, n.n4, n.n5)
		n.count5++

	}

}

// Counting winning numbers and the powerball per month
func (n *WinNum) WinPerMoth(powerballstr string, date string) {
	n.powerball, _ = strconv.Atoi(powerballstr)
	if strings.Contains(date, "January") {
		WinPerMonth[0][n.n1-1]++
		WinPerMonth[0][n.n2-1]++
		WinPerMonth[0][n.n3-1]++
		WinPerMonth[0][n.n4-1]++
		WinPerMonth[0][n.n5-1]++
		PowerballPerMoth[0][n.powerball-1]++
	}
	if strings.Contains(date, "February") {
		WinPerMonth[1][n.n1-1]++
		WinPerMonth[1][n.n2-1]++
		WinPerMonth[1][n.n3-1]++
		WinPerMonth[1][n.n4-1]++
		WinPerMonth[1][n.n5-1]++
		PowerballPerMoth[1][n.powerball-1]++
	}
	if strings.Contains(date, "March") {
		WinPerMonth[2][n.n1-1]++
		WinPerMonth[2][n.n2-1]++
		WinPerMonth[2][n.n3-1]++
		WinPerMonth[2][n.n4-1]++
		WinPerMonth[2][n.n5-1]++
		PowerballPerMoth[2][n.powerball-1]++
	}
	if strings.Contains(date, "April") {
		WinPerMonth[3][n.n1-1]++
		WinPerMonth[3][n.n2-1]++
		WinPerMonth[3][n.n3-1]++
		WinPerMonth[3][n.n4-1]++
		WinPerMonth[3][n.n5-1]++
		PowerballPerMoth[3][n.powerball-1]++
	}
	if strings.Contains(date, "May") {
		WinPerMonth[4][n.n1-1]++
		WinPerMonth[4][n.n2-1]++
		WinPerMonth[4][n.n3-1]++
		WinPerMonth[4][n.n4-1]++
		WinPerMonth[4][n.n5-1]++
		PowerballPerMoth[4][n.powerball-1]++
	}
	if strings.Contains(date, "June") {
		WinPerMonth[5][n.n1-1]++
		WinPerMonth[5][n.n2-1]++
		WinPerMonth[5][n.n3-1]++
		WinPerMonth[5][n.n4-1]++
		WinPerMonth[5][n.n5-1]++
		PowerballPerMoth[5][n.powerball-1]++
	}
	if strings.Contains(date, "July") {
		WinPerMonth[6][n.n1-1]++
		WinPerMonth[6][n.n2-1]++
		WinPerMonth[6][n.n3-1]++
		WinPerMonth[6][n.n4-1]++
		WinPerMonth[6][n.n5-1]++
		PowerballPerMoth[6][n.powerball-1]++
	}
	if strings.Contains(date, "August") {
		WinPerMonth[7][n.n1-1]++
		WinPerMonth[7][n.n2-1]++
		WinPerMonth[7][n.n3-1]++
		WinPerMonth[7][n.n4-1]++
		WinPerMonth[7][n.n5-1]++
		PowerballPerMoth[7][n.powerball-1]++
	}
	if strings.Contains(date, "September") {
		WinPerMonth[8][n.n1-1]++
		WinPerMonth[8][n.n2-1]++
		WinPerMonth[8][n.n3-1]++
		WinPerMonth[8][n.n4-1]++
		WinPerMonth[8][n.n5-1]++
		PowerballPerMoth[8][n.powerball-1]++
	}
	if strings.Contains(date, "October") {
		WinPerMonth[9][n.n1-1]++
		WinPerMonth[9][n.n2-1]++
		WinPerMonth[9][n.n3-1]++
		WinPerMonth[9][n.n4-1]++
		WinPerMonth[9][n.n5-1]++
		PowerballPerMoth[9][n.powerball-1]++
	}
	if strings.Contains(date, "November") {
		WinPerMonth[10][n.n1-1]++
		WinPerMonth[10][n.n2-1]++
		WinPerMonth[10][n.n3-1]++
		WinPerMonth[10][n.n4-1]++
		WinPerMonth[10][n.n5-1]++
		PowerballPerMoth[10][n.powerball-1]++
	}
	if strings.Contains(date, "December") {
		WinPerMonth[11][n.n1-1]++
		WinPerMonth[11][n.n2-1]++
		WinPerMonth[11][n.n3-1]++
		WinPerMonth[11][n.n4-1]++
		WinPerMonth[11][n.n5-1]++
		PowerballPerMoth[11][n.powerball-1]++
	}

}
