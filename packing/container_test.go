package packing

import (
	"testing"
)

func TestIsEqual(t *testing.T) {
	samples := []struct {
		first Container
		second Container
		areEqual bool
	}{
		{
			Container{ weights: nil },
			Container{ weights: nil },
			true,
		}, {
			Container{ weights: []int{ 1 } },
			Container{ weights: []int{} },
			false,
		}, {
			Container{ weights: []int{ 1, 2, 3, 4 } },
			Container{ weights: []int{ 4, 3, 2, 1 } },
			false,
		}, {
			Container{ weights: []int{} },
			Container{ weights: []int{} },
			true,
		},
	}

	for _, sample := range samples {
		expected := sample.areEqual
		result := sample.first.isEqual(sample.second)
		if result != expected {
			t.Error("result:", result, "| expected:", expected)
		}
	}
}

func TestAreEqual(t *testing.T) {
	samples := []struct {
		first []Container
		second []Container
		areEqual bool
	}{
		{
			[]Container{},
			[]Container{},
			true,
		}, {
			[]Container{
				Container{ weights: nil },
				Container{ weights: nil }, 
			},
			[]Container{
				Container{ weights: nil },
				Container{ weights: nil },
			},
			true,
		}, {
			[]Container{},
			[]Container{
				Container{ weights: nil },
			},
			false,
		}, {
			[]Container{
				Container{ weights: []int{ 1, 2, 3, 4 } },
				Container{ weights: []int{ 5, 6, 7, 8 } }, 
			},
			[]Container{
				Container{ weights: []int{ 1, 2, 3, 4 } },
				Container{ weights: []int{ 5, 6, 7, 8 } }, 
			},
			true,
		}, {
			[]Container{
				Container{ weights: []int{ 1 } },
			},
			[]Container{
				Container{ weights: []int{ 2 } },
			},
			false,
		},
	}
	for _, sample := range samples {
		expected := sample.areEqual
		result := areEqual(sample.first, sample.second)
		if result != expected {
			t.Error("result:", result, "| expected:", expected)
		}
	}
}

func TestGetSum(t *testing.T) {
	samples := []struct {
		container Container
		sum int
	}{
		{
			Container{ weights: nil },
			0,
		}, {
			Container{ weights: []int{} },
			0,
		}, {
			Container{ weights: []int{ 1, 2, 3, 4 } },
			10,
		}, {
			Container{ weights: []int{ 0, 0, 0, 0, 0 } },
			0,
		},
	}

	for _, sample := range samples {
		expected := sample.sum
		result := sample.container.getSum()
		if result != expected {
			t.Error("result:", result, "| expected:", expected)
		}
	}
}

func TestGetPadding(t *testing.T) {
	samples := []struct {
		container Container
		capacity int
		padding int
	} {
		{
			Container{ weights: []int{ 1, 2, 3, 7 } },
			15,
			2,
		}, {
			Container{ weights: []int{ 1, 2, 3 } },
			10,
			4,
		}, {
			Container{ weights: []int{} },
			10,
			10,
		}, {
			Container{ weights: []int{ 1, 2, 3 } },
			6,
			0,
		}, {
			Container{ weights: []int{ 1, 2, 3 } },
			1,
			-5,
		},
	}

	for _, sample := range samples {
		expected := sample.padding
		result := sample.container.GetPadding(sample.capacity)
		if result != expected {
			t.Error("result:", result, "| expected:", expected)
		}
	}
}