package packing

import (
	"testing"
)

func TestBestFit(t *testing.T) {
	samples := []struct {
		weights []int
		capacity int
		containers []Container
	}{
		{
			nil,
			0,
			[]Container{
				Container{ weights: nil },
			},
		},
		{
			[]int{},
			0,
			[]Container{
				Container{ weights: []int{} },
			},
		},
		{
			[]int{ 1, 2, 3, 4 },
			5,
			[]Container{
				Container{ weights: []int{ 1, 2 } },
				Container{ weights: []int{ 3 } },
				Container{ weights: []int{ 4 } },
			},
		},
	}

	for _, sample := range samples {
		expected := sample.containers
		result := BestFit(sample.weights, sample.capacity)
		if !areEqual(result, expected) {
			t.Error("result: ", result, "| expected: ", expected)
		}
	}
}

func TestNewSolution(t *testing.T) {
	samples := []struct {
		containers []Container
		capacity int
		newContainers []Container
	}{
		{
			[]Container{},
			10,
			[]Container{},
		}, {
			nil,
			0,
			nil,
		},
	}

	for _, sample := range samples {
		expected := sample.newContainers
		result := NewSolution(sample.containers, sample.capacity)
		if !areEqual(result, expected) {
			t.Error("result:", result, "| expected:", expected)
		}
	}
}

func TestRandNewSolution(t *testing.T) {
	samples := []struct {
		containers []Container
		capacity int
	} {
		{
			[]Container{
				Container{ weights: []int{ 1, 2, 3, 7 } },
				Container{ weights: []int{ 6, 8 } },
				Container{ weights: []int{ 1, 1, 10 } },
			},
			15,
		}, {
			[]Container{
				Container{ weights: []int{ 1, 2, 3 } },
				Container{ weights: []int{ 6 } },
			},
			10,
		},
	}

	for _, sample := range samples {
		initial := sample.containers
		capacity := sample.capacity
		solution := NewSolution(sample.containers, sample.capacity)
		t.Log("containers:", initial, "| capacity:", capacity, "| new solution:", solution)
	}
}

func TestCalculatePadding(t *testing.T) {
	samples := []struct {
		containers []Container
		capacity int
		padding int
	} {
		{
			[]Container{
				Container{ weights: []int{ 1, 2, 3, 7 } },
				Container{ weights: []int{ 6, 8 } },
				Container{ weights: []int{ 1, 1, 10 } },
			},
			15,
			6,
		}, {
			[]Container{
				Container{ weights: []int{ 1, 2, 3 } },
				Container{ weights: []int{ 6 } },
			},
			10,
			8,
		}, {
			[]Container{
				Container{ weights: []int{} },
				Container{ weights: []int{} },
			},
			10,
			20,
		}, {
			[]Container{},
			10,
			0,
		},
	}

	for _, sample := range samples {
		expected := sample.padding
		result := calculatePadding(sample.containers, sample.capacity)
		if result != expected {
			t.Error("result:", result, "| expected:", expected)
		}
	}
}