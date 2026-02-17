package testutil

import (
	"fmt"
	"reflect"
)

// Run prints a labeled PASS/FAIL line comparing expected and got.
func Run[T comparable](label string, expected, got T) {
	status := "PASS"
	if got != expected {
		status = "FAIL"
	}
	fmt.Printf("[%s] %s | expected=%v got=%v\n", status, label, expected, got)
}

// RunDeep is like Run but uses reflect.DeepEqual — use this for slices and maps.
func RunDeep[T any](label string, expected, got T) {
	status := "PASS"
	if !reflect.DeepEqual(expected, got) {
		status = "FAIL"
	}
	fmt.Printf("[%s] %s | expected=%v got=%v\n", status, label, expected, got)
}
