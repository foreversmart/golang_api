package ktree

import "math/rand"

func RandByte() byte {
	return byte(rand.Intn(256))
}

func RandString(n int) string {
	arr := make([]byte, n)
	for i := 0; i < n; i++ {
		arr[i] = RandByte()
	}

	return string(arr)
}
