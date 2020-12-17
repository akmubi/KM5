package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"./packing"
)

// Вспомогательная функция для чтения целого числа
func readInt(reader *bufio.Reader) (int, error) {
	bytes, _, err := reader.ReadLine()
	if err != nil {
		return -1, err
	}
	integer, err := strconv.ParseInt(string(bytes), 10, 64)
	if err != nil {
		return -1, err
	}
	return int(integer), nil
}

// Выполняет чтение вместимости контейнеров, количества
// предметов и их весов
func readData(filename string) ([]int, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, -1, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	// вместимость
	capacity, err := readInt(reader)
	if err != nil {
		return nil, -1, err
	}

	// количество предметов
	n, err := readInt(reader)
	if err != nil {
		return nil, -1, err
	}

	weights := make([]int, n)

	for i := 0; i < n; i++ {
		weight, err := readInt(reader)
		if err != nil {
			return nil, -1, err
		}
		weights[i] = weight
	}
	return weights, capacity, nil
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type parameters struct {
	T float64
	r float64
	L int
	E int
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Необходимо указание входного файла!")
		return
	}

	weights, capacity, err := readData(args[0])
	check(err)

	params := []parameters{
		parameters{T: 1000.0, r: 0.8, L: 100, E: 5},
		parameters{T: 2000.0, r: 0.83, L: 200, E: 10},
		parameters{T: 3000.0, r: 0.87, L: 300, E: 15},
		parameters{T: 4000.0, r: 0.92, L: 400, E: 20},
		parameters{T: 5000.0, r: 0.99, L: 500, E: 25},
	}

	for _, param := range params {
		fmt.Println("Вместимость (W) =", capacity)
		fmt.Println("Количество предметов (n) =", len(weights))
		sum := 0
		for _, weight := range weights {
			sum += weight
		}
		fmt.Println("Сумма предметов - ", sum)
		fmt.Println()
		fmt.Println("Температура (T) =", param.T)
		fmt.Println("Коэффициент охлаждения (r) =", param.r)
		fmt.Println("Число шагов алгоритма (L) =", param.L)
		fmt.Println("Число смен температуры без изменения текущего решения (E) =", param.E)

		containers := packing.SimulatedAnnealing(weights, capacity, param.T, param.r, param.L, param.E)
		fmt.Println("Общее количество контейнеров:", len(containers))
		for i, container := range containers {
			paddingPercentage := float32(container.GetPadding(capacity)) / float32(capacity) * 100.0
			fmt.Printf("%5d:\t%6.2f%%\n", i, 100.0-paddingPercentage)
		}

		// вычисление процента заполненности
		var fillPercentage float64
		for _, container := range containers {
			padding := container.GetPadding(capacity)
			fillPercentage += float64(capacity-padding) / float64(capacity) * 100.0
		}
		fillPercentage /= float64(len(containers))

		fmt.Printf("Процент заполненности контейнеров: %.2f%%\n\n\n", fillPercentage)
	}
}
