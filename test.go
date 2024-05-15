package main

// based on https://stackoverflow.com/questions/38777961/converting-string-input-to-float64-using-parsefloat-in-golang
import (
	"fmt"
	"strconv"
	"strings"
)

const conversionWeight float64 = 453.592
const conversionSize float64 = 0.3048

func feetToMeters(feet float64) (meters float64) {
	return feet * conversionSize
}
func lbsToGrams(lbs float64) (grams float64) {
	return lbs * conversionWeight
}

func main() {
	fmt.Println("hello, this calc converts 15 pounds into grams & 21 feet to meters for you")
	var lbs = " 15.00 "
	lbsFloat, _ := strconv.ParseFloat(strings.TrimSpace(lbs), 64)
	fmt.Println("\nChecking lbs in float format:", lbsFloat)
	grams := lbsToGrams(lbsFloat)
	fmt.Printf("%v lbs converted to grams gives %v gr", lbsFloat, grams)

	/*
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter feet value: \n")
		feet, _ := reader.ReadString('\n')
	*/

	var feet = " 21.00 "
	feetFloat, _ := strconv.ParseFloat(strings.TrimSpace(feet), 64)
	fmt.Println("\nChecking feet in float format:", feetFloat)
	meters := feetToMeters(feetFloat)
	fmt.Printf("%v feet converted to meters give you %v meters", feetFloat, meters)
}
