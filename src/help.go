package src

import (
	"math/rand"
	"time"
)

func RandSign(len int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func RandByte(min int, max int) byte {
	rand.Seed(time.Now().UnixNano())
	return byte(min + rand.Intn(max-min))
}

func RandCode(len int) []byte {
	var code = make([]byte, len)
	for i := 0; i < len; i++ {
		code[i] = byte(48 + rand.Intn(9))
	}
	return code
}

func Combine(arrA []byte, arrs ...[]byte) []byte {
	Alen := len(arrA)
	Blen := 0
	for _, item := range arrs {
		Blen += len(item)
	}
	carrier := make([]byte, Alen+Blen)
	i := 0
	for _, item := range arrA {
		carrier[i] = item
		i++
	}
	for _, data := range arrs {
		for _, d := range data {
			carrier[i] = d
			i++
		}
	}
	return carrier
}