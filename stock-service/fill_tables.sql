insert into stock.public.algorithm (name, task, solution, price) values ('hash_map', 'Дан неотсортированный массив из N чисел от 1 до N,
при этом несколько чисел из диапазона [1, N] пропущено,
а некоторые присутствуют дважды.

Найти все пропущенные числа.', '
func Solver(s []int) []int {

	output := make([]int, 0)
	m := make(map[int]struct{}, len(s))
	for _, v := range s {
		m[v] = struct{}{}
	}
	for i := 1; i <= len(s); i++ {
		_, ok := m[i]
		if !ok {
			output = append(output, i)
		}
	}
	return output
}', '100');