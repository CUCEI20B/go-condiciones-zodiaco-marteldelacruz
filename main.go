package main

import "fmt"

func main() {
	var day, month int
	var zodiacSign string

	// Read day from user input
	fmt.Scanf("%d ", &day)

	// Read month from user input
	fmt.Scanf("%d", &month)

	// Switch for each month
	switch month {
	case 1:
		if day < 20 {
			zodiacSign = "capricornio"
		} else {
			zodiacSign = "acuario"
		}
		break
	case 2:
		if day < 18 {
			zodiacSign = "acuario"
		} else {
			zodiacSign = "piscis"
		}
		break
	case 3:
		if day < 21 {
			zodiacSign = "piscis"
		} else {
			zodiacSign = "aries"
		}
		break
	case 4:
		if day < 20 {
			zodiacSign = "aries"
		} else {
			zodiacSign = "tauro"
		}
		break
	case 5:
		if day < 21 {
			zodiacSign = "tauro"
		} else {
			zodiacSign = "geminis"
		}
		break
	case 6:
		if day < 21 {
			zodiacSign = "geminis"
		} else {
			zodiacSign = "cancer"
		}
		break
	case 7:
		if day < 23 {
			zodiacSign = "cancer"
		} else {
			zodiacSign = "leo"
		}
		break
	case 8:
		if day < 23 {
			zodiacSign = "leo"
		} else {
			zodiacSign = "virgo"
		}
		break
	case 9:
		if day < 23 {
			zodiacSign = "virgo"
		} else {
			zodiacSign = "libra"
		}
		break
	case 10:
		if day < 23 {
			zodiacSign = "libra"
		} else {
			zodiacSign = "escorpio"
		}
		break
	case 11:
		if day < 22 {
			zodiacSign = "escorpio"
		} else {
			zodiacSign = "sagitario"
		}
		break
	case 12:
		if day < 22 {
			zodiacSign = "sagitario"
		} else {
			zodiacSign = "capricornio"
		}
		break
	default:
		zodiacSign = "Unknown"
		break
	}
	fmt.Print(zodiacSign)
}
