package utils

import (
	"math/rand"
	"strconv"
)

func RandomString(length int) string {
	str := ""
	for i := 0; i < length; i++ {
		str += strconv.Itoa(rand.Intn(10))
	}
	return str
}
