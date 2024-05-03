package app

import (
	"errors"
	"reflect"
)

func Multiply(x float64, y int) float64 {
	return x * float64(y)
}

func IsInteger(x interface{}) (bool, error) {
	switch reflect.TypeOf(x).Kind() {
	case reflect.Int:
		return true, nil
	default:
		return false, errors.New("parameter bukan int")
	}
}

func IntToDayName(x int) string {
	hari := []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	if x > 0 && x <= len(hari) {
		return hari[x-1]
	} else {
		return "Error, coba cek lagi bang valuenya"
	}
}

func DayNameToInt(x string) int {
	hari := map[string]int{
		"Senin":  1,
		"Selasa": 2,
		"Rabu":   3,
		"Kamis":  4,
		"Jumat":  5,
		"Sabtu":  6,
		"Minggu": 7,
	}

	dayName, ok := hari[x]
	if ok {
		return dayName
	} else {
		return 0
	}
}

func IsHello(kosong *string) {
	if *kosong == "" {
		*kosong = "Hello World"
	}
}
