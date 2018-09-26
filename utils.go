package main

import (
	"math/rand"
)

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyz_-!.=$@:1234567890")
)

func inArray(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
