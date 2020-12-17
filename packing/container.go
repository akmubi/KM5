package packing

// Container - Представляет собой контейнер с предметами
type Container struct {
	weights []int
}

// New - Возвращает новый контейнер
func New() Container {
	return Container{weights: nil}
}

// getSum - Вычисляет сумму весов контейнера
func (container Container) getSum() int {
	sum := 0
	for _, weight := range container.weights {
		sum += weight
	}
	return sum
}

func (container Container) isEqual(anotherContainer Container) bool {
	if len(container.weights) != len(anotherContainer.weights) {
		return false
	}

	for i, weight := range container.weights {
		if weight != anotherContainer.weights[i] {
			return false
		}
	}
	return true
}

func areEqual(first, second []Container) bool {
	if len(first) != len(second) {
		return false
	}

	for i, container := range first {
		if !container.isEqual(second[i]) {
			return false
		}
	}
	return true
}

// Добавляет вес в контейнер
func (container *Container) append(weight int) {
	container.weights = append(container.weights, weight)
}

// GetPadding - Вычисляет размер оставшегося места в контейнере
func (container Container) GetPadding(capacity int) int {
	return capacity - container.getSum()
}
