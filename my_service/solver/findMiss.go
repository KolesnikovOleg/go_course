package solver

import "encoding/json"

func NewFindMissSolver(desc string) FindMissSolver {
	return FindMissSolver{
		description: desc,
	}
}

type FindMissSolver struct {
	description string
}

func (c FindMissSolver) Solution(A []int) int {
	counter := len(A) + 1
	sum := 0
	for i := 0; i < len(A); i++ {
		counter += i + 1
		sum += A[i]
	}
	return counter - sum
}

func (c FindMissSolver) Solve(data interface{}) interface{} /*[][]int */ {
	res := []int{}
	for _, value := range data.([][][]int) {
		res = append(res, c.Solution(value[0]))
	}

	return res
}

func (c FindMissSolver) Description() string {
	return c.description
}

func (c FindMissSolver) Decode(rawData []byte) (interface{}, error) {
	var data [][][]int
	if err := json.Unmarshal(rawData, &data); err != nil {
		//fmt.Println(err)
		return nil, err
	}
	return data, nil
}
