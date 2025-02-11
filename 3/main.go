package main

import "fmt"

func main() {
	dayName, err := predictWeek(31)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dayName)
}

func predictWeek(dayNumber int) (string, error) {
	switch dayNumber {
	case 1:
		return "Monday", nil
	case 2:
		return "Tuesday", nil
	case 3:
		return "Wednesday", nil
	case 4:
		return "Thursday", nil
	case 5:
		return "Friday", nil
	case 6:
		return "Saturday", nil
	case 7:
		return "Sunday", nil
	default:
		return "", fmt.Errorf("Invalid day number")
	}
}
