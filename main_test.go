package main

import "testing"

type MathData struct {
	x, y   int
	result int
}
type DivideData struct {
	x, y   int
	result float64
}

func TestAdd(t *testing.T) {
	addData := []MathData{
		{1, 2, 3},
		{5, 5, 10},
		{77, 33, 110},
	}
	for _, datum := range addData {
		result := Add(datum.x, datum.y)
		if result != datum.result {
			t.Errorf("Adding(%d, %d) FAILED. Expected %d got %d\n", datum.x, datum.y, datum.result, result)
		}
	}
}

func TestSubtract(t *testing.T) {
	subData := []MathData{
		{2, 1, 1},
		{5, 5, 0},
		{77, 33, 44},
	}
	for _, datum := range subData {
		result := Subtract(datum.x, datum.y)
		if result != datum.result {
			t.Errorf("Subtracting(%d, %d) FAILED. Expected %d got %d\n", datum.x, datum.y, datum.result, result)
		}
	}
}

func TestMultiply(t *testing.T) {
	mulData := []MathData{
		{1, 2, 2},
		{5, 5, 25},
		{7, 11, 77},
	}
	for _, datum := range mulData {
		result := Multiply(datum.x, datum.y)
		if result != datum.result {
			t.Errorf("Multiplying(%d, %d) FAILED. Expected %d got %d\n", datum.x, datum.y, datum.result, result)
		}
	}
}

func TestDivide(t *testing.T) {
	divData := []DivideData{
		{5, 0, 0},
		{21, 3, 7},
		{110, 11, 10},
	}
	for _, datum := range divData {
		result := Divide(datum.x, datum.y)
		if result != datum.result {
			t.Errorf("Dividing(%d, %d) FAILED. Expected %v got %v\n", datum.x, datum.y, datum.result, result)
		}
	}
}
