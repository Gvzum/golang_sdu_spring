package main

import (
	"strconv"

	diasmath "github.com/Gvzum/dimath/dimath"
	"github.com/Gvzum/dimath/printer"
)

func main() {
	message := "Math operations\n" +
		"Add 1 + 2 = " + strconv.Itoa(diasmath.Add(1, 2)) + "\n" +
		"Minus 2 - 3 = " + strconv.Itoa(diasmath.Subs(2, 3)) + "\n" +
		"Multi 2 * 3 = " + strconv.Itoa(diasmath.Multi(2, 3)) + "\n" +
		"Divide 10 / 5 = " + strconv.Itoa(diasmath.Divi(10, 5))

	printer.Printer(message)
}
