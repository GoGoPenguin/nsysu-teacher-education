package typecast

import (
	"fmt"
	"strconv"
)

// ToString transform any supported type to string
func ToString(from interface{}) string {
	return fmt.Sprintf("%v", from)
}

// StringToInt convert string to int, e.g string "18" => int 18
func StringToInt(from string) int {
	to, err := strconv.ParseInt(from, 10, 0)
	if err != nil {
		panic(err)
	}
	return int(to)
}

// StringToFloat64 convert string to int, e.g string "3.1415" => float64 3.1415
func StringToFloat64(from string) float64 {
	to, err := strconv.ParseFloat(from, 64)
	if err != nil {
		panic(err)
	}
	return to
}

// StringToBool convert string to boolean
func StringToBool(form string) bool {
	to, err := strconv.ParseBool(form)
	if err != nil {
		panic(err)
	}
	return to
}
