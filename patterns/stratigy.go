package patterns

type SortStrategy interface {
	Sort(data []int) []int
}

type BubbleSortStrategy struct{}

func (b *BubbleSortStrategy) Sort(data []int) []int {
	// Реализация сортировки пузырьком
	return nil
}

type InsertionSortStrategy struct{}

func (i *InsertionSortStrategy) Sort(data []int) []int {
	// Реализация сортировки вставками
	return nil
}

type MergeSortStrategy struct{}

func (m *MergeSortStrategy) Sort(data []int) []int {
	// Реализация сортировки слиянием
	return nil
}

type Sorter struct {
	strategy SortStrategy
}

func (s *Sorter) SetStrategy(strategy SortStrategy) {
	s.strategy = strategy
}

func (s *Sorter) Sort(data []int) []int {
	if s.strategy == nil {
		// Обработка ошибки: стратегия не установлена
	}
	return s.strategy.Sort(data)
}
