package packing

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

/*
	BestFit
	Алгоритм наилучший подходящий (BF)
	входные данные:
		weights - веса предметов
		capacity - вместимость контейнеров
	выходные данные:
		заполненные предметами контейнеры
*/
func BestFit(weights []int, capacity int) []Container {
	containers := []Container{New()}

	// количество предметов
	n := len(weights)
	if n > 0 {
		// помещаем 1-й предмет в 1-й контейнер
		containers[0].append(weights[0])
	}

	for k := 1; k < n; k++ {
		// вычисляем минимальный размер пустого
		// пространства с учётом текущего веса
		// размер = вместимость - (сумма весов + текущий вес)
		minDelta, minI := capacity+1, -1

		// текущее количество контейнеров
		m := len(containers)
		for i := 0; i < m; i++ {
			sum := containers[i].getSum()
			delta := capacity - (sum + weights[k])
			if delta >= 0 && delta < minDelta {
				minDelta = delta
				minI = i
			}
		}
		// если подходящий вес был найден помещаем
		// в соответствующий контейнер
		// иначе создаём новый и помещаем уже туда
		if minI != -1 {
			containers[minI].append(weights[k])
		} else {
			containers = append(containers, New())
			containers[m].append(weights[k])
		}
	}
	return containers
}

// intUniform - Генерирует случайное целое число,
// равномерно распределенное в полуинтервале [a, b)
func intUniform(a, b int) int {
	rand.Seed(time.Now().UnixNano())
	return a + rand.Intn(b-a)
}

// floatUniform - Генерирует случайное вещественное число,
//равномерно распределенное в полуинтервале [a, b)
func floatUniform(a, b float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return a + (b-a)*rand.Float64()
}

func createCopy(containers []Container) []Container {
	copy := make([]Container, len(containers))
	for i, container := range containers {
		for _, weight := range container.weights {
			copy[i].weights = append(copy[i].weights, weight)
		}
	}
	return copy
}

/*
	Функция нахождения нового решения
	входные данные:
		containers - текущее решение (заполненные предметами контейнеры)
		capacity - вместимость контейнеров
	выходные данные:
		новое решение
*/
func NewSolution(containers []Container, capacity int) []Container {
	// случайно выбираем либо перемещение, либо обмен предметов
	// между контейнерами
	methodID := intUniform(0, 2)
	methods := []func([]Container, int) []Container{moveRandWeights, swapRandWeights}
	return methods[methodID](containers, capacity)
}

// moveRandWeights - Перемещает случайный предмет из одного случайного
// контейнера в другой случайный контейнер
func moveRandWeights(containers []Container, capacity int) []Container {
	newSolution := createCopy(containers)

	// индексы незаполненных до конца контейнеров
	unfilledIndices := []int{}
	for i, container := range containers {
		if container.getSum() < capacity {
			unfilledIndices = append(unfilledIndices, i)
		}
	}

	// если все контейнеры заполнены,
	// значит не осталось контейнеров,
	// куда можно было бы переместить предмет
	if len(unfilledIndices) == 0 {
		return newSolution
	}

	// возможные кандидаты, откуда можно переместить
	// предмет в выбранный контейнер
	var appropriateContainers [][]int

	// массив с настоящими индексами контейнеров
	var containerIndices []int

	// количество подходящих контейнеров
	var l int

	// индекс случайно выбранного контейнера,
	// куда будет перемещен предмет
	var destinationIndex int

	// пока не найдутся контейнеры, с которыми можно было бы
	// произвести обмен предметами
	for l == 0 {
		// выбираем случайный незаполненный контейнер
		u1 := intUniform(0, len(unfilledIndices))
		destinationIndex = unfilledIndices[u1]

		padding := containers[destinationIndex].GetPadding(capacity)

		// для каждого нового рассматриваемого контейнера с индексом U1
		// нужно заново рассматривать возможных кандидатов, с которыми
		// можно было бы произвести обмен
		l = 0
		containerIndices = []int{}
		appropriateContainers = [][]int{}

		for i, container := range containers {
			if i != destinationIndex {
				k := 0
				for j, weight := range container.weights {
					// если один предмет из другого контейнера в влезает в выбранный
					// контейнер, то сохраняем индеск другого предмета
					if weight <= padding {
						// если количество контейнеров меньше, чтобы
						// создать новый контейнер и добавить туда новый
						// предмет, то удлиняем массив
						if len(appropriateContainers) < l+1 {
							appropriateContainers = append(appropriateContainers, []int{})
							containerIndices = append(containerIndices, i)
						}
						// если количество индексов предметов в контейнере меньше, чтобы
						// добавить ещё один, то удлиняем массив
						if len(appropriateContainers[l]) < k+1 {
							appropriateContainers[l] = append(appropriateContainers[l], j)
						}
						// вес сохрананен, увеличивам количество сохраненных весов
						k = k + 1
					}
				}
				// только в случае, когда контейнер не пустой
				// мы "создаём новый контейнер
				if k > 0 {
					l = l + 1
				}
			}
		}
	}
	// случайные числа
	// u2 - индекс случайного контейнера
	// u3 - индекс случайного индекса предмета
	var u2, u3 int
	u2 = intUniform(0, l)

	// количество подходящих предметов
	weightCount := len(appropriateContainers[u2])
	u3 = intUniform(0, weightCount)

	// индекс случайно выбранного контейнера
	containerIndex := containerIndices[u2]

	// индекс случайно выбранного предмета
	weightIndex := appropriateContainers[u2][u3]

	weightToMove := newSolution[containerIndex].weights[weightIndex]

	newSolution[destinationIndex].weights = append(newSolution[destinationIndex].weights, weightToMove)

	// удаляем предмет из контейнера, откуда он был взят
	weightCount = len(newSolution[containerIndex].weights)

	// если до перемещения оставался только 1 предмет, то
	// удаляем контейнер, иначе удаляем перемещенный предмет
	if weightCount == 1 {
		newSolution[containerIndex] = newSolution[len(newSolution)-1]
		newSolution[len(newSolution)-1] = Container{}
		newSolution = newSolution[:len(newSolution)-1]
	} else {
		// перезаписываем перемещенный предмет последним элементом
		newSolution[containerIndex].weights[weightIndex] = newSolution[containerIndex].weights[weightCount-1]

		// "стираем" последний элемент
		newSolution[containerIndex].weights[weightCount-1] = 0

		// укорачиваем длину среза
		newSolution[containerIndex].weights = newSolution[containerIndex].weights[:weightCount-1]
	}
	return newSolution
}

// swapRandWeights - Производит обмен между случайно взятыми предметами
// в случайных контейнерах
func swapRandWeights(containers []Container, capacity int) []Container {
	// количество контейнеров
	m := len(containers)

	newSolution := createCopy(containers)

	// если контейнеров нет, то нечего рассматривать
	if m > 0 {
		// количество подходящих контейнеров
		var l int
		// случайные числа:
		// u1 - индекс случайного контейнера
		// u2 - индекс случайного предмета
		var u1, u2 int

		// массив контейнеров с подходящими весами
		var appropriateContainers [][]int

		// массив с настоящими индексами контейнеров
		var containerIndices []int

		// пока не найдутся контейнеры, с которыми можно было бы
		// произвести обмен предметами
		for l == 0 {
			u1 = intUniform(0, m)

			weightCount1 := len(newSolution[u1].weights)
			u2 = intUniform(0, weightCount1)

			// текущий (рассматриваемый) предмет
			currentWeight := newSolution[u1].weights[u2]
			// вычисляем оставшееся место в контейнере, без учёта
			// текущего предмета (для того, чтобы узнать, может ли другой
			// предмет влезть в этот контейнер, если бы текущего предмета
			// в контейнере не было)
			delta1 := capacity - (newSolution[u1].getSum() - currentWeight)

			// для каждого нового рассматриваемого контейнера с индексом U1
			// нужно заново рассматривать возможных кандидатов, с которыми
			// можно было бы произвести обмен
			l = 0
			containerIndices = []int{}
			appropriateContainers = [][]int{}

			for i, container := range newSolution {
				if i != u1 {
					k := 0
					for j, weight := range container.weights {
						// тажке как и для currentWeight вычисляем разницу
						delta2 := capacity - (container.getSum() - weight)
						// если один предмет из 1-го контейнера в влезает во 2-й
						// и если другой предмет из 2-го контейнера влезает в 1-й,
						// то сохраняем индеск другого предмета в качестве возможной
						// замены первого
						if weight <= delta1 && currentWeight <= delta2 {
							// если количество контейнеров меньше, чтобы
							// создать новый контейнер и добавить туда новый
							// предмет, то удлиняем массив
							if len(appropriateContainers) < l+1 {
								appropriateContainers = append(appropriateContainers, []int{})
								containerIndices = append(containerIndices, i)
							}
							// если количество весов в контейнере меньше, чтобы
							// добавить ещё один, то удлиняем массив
							if len(appropriateContainers[l]) < k+1 {
								appropriateContainers[l] = append(appropriateContainers[l], j)
							}
							// вес сохрананен, увеличивам количество сохраненных весов
							k = k + 1
						}
					}
					// только в случае, когда контейнер не пустой
					// мы "создаём новый контейнер
					// нельзя, чтобы оставались пустые контейнеры
					if k > 0 {
						l = l + 1
					}
				}
			}
		}

		// случайные числа
		// u3 - индекс случайного контейнера из appropriateContainers
		// u4 - индекс случайного индекса предмета
		var u3, u4 int
		u3 = intUniform(0, l)

		// количество подходящих предметов
		weightCount2 := len(appropriateContainers[u3])
		u4 = intUniform(0, weightCount2)

		// индекс случайно выбранного контейнера
		containerIndex := containerIndices[u3]

		// индекс случайно выбранного предмета
		weightIndex := appropriateContainers[u3][u4]

		// обмен предметами
		temp := newSolution[u1].weights[u2]
		newSolution[u1].weights[u2] = newSolution[containerIndex].weights[weightIndex]
		newSolution[containerIndex].weights[weightIndex] = temp
	}
	return newSolution
}

/*
	Вычисление "функции энергии"
	входные данные:
		containers - заполненные контейнеры
		capacity - вместимость контейнеров
	выходные данные:
		максимальное занчение незаполненного пространства в контейнерах
*/
func calculateUnfilledContainers(containers []Container, capacity int) int {
	// количество не заполненных до конца контейнеров
	var unfilledCount int = 0
	for _, container := range containers {
		padding := container.GetPadding(capacity)
		if padding != 0 {
			unfilledCount++
		}
	}
	return unfilledCount
}

/*
	Алгоритм имитации отжига
	входные данные:
		weights - веса предметов
		capacity - вместимость контейнеров
		T - начальная температура
		r - коэффициент охлаждения
		L - число шагов алгоритма
		E - число смен температуры без изменения текущего решения
	выходные данные:
		полученное решение (заполенные контейнеры)
*/
func SimulatedAnnealing(weights []int, capacity int, T, r float64, L, E int) []Container {
	// текущее число смен температуры
	// без изменения текущего решения
	var p int
	solution := BestFit(weights, capacity)
	for p < E {
		// копируем текущее решениея для дальнейшего сравнения
		initialSolution := createCopy(solution)

		for i := 0; i < L; i++ {
			anotherSolution := NewSolution(solution, capacity)
			delta := calculateUnfilledContainers(anotherSolution, capacity) - calculateUnfilledContainers(solution, capacity)
			u := floatUniform(0, 1)
			border := -1.0
			if delta > 0 {
				border = math.Exp(float64(-delta) / T)
			}

			if delta <= 0 || u <= border {
				solution = anotherSolution
			}
		}
		T = T * r

		// если решение не изменилось, то
		// увеличиваем счётчик
		if areEqual(solution, initialSolution) {
			p = p + 1
		}

		fmt.Printf("\rСчётчик неизмененных решений (P) - %2d | Температура (T) - %g", p, T)
	}
	fmt.Println()
	return solution
}
