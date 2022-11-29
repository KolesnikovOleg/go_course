package solver

import "encoding/json"

func NewCheckSeqSolver(desc string) CheckSeqSolver {
	return CheckSeqSolver{
		description: desc,
	}
}

type CheckSeqSolver struct {
	description string
}

func (c CheckSeqSolver) Solution(A []int) int {
	hashMap := make(map[int]int)
	for id, value := range A {
		hashMap[value] = id
	}
	for i := 0; i < len(A); i++ {
		_, find := hashMap[i+1]
		if !find {
			return 0
		}
	}

	return 1
}

func (c CheckSeqSolver) Solve(data interface{}) interface{} /*[][]int */ {
	res := []int{}
	for _, value := range data.([][][]int) {
		res = append(res, c.Solution(value[0]))
	}

	return res
}

func (c CheckSeqSolver) Description() string {
	return c.description
}

func (c CheckSeqSolver) Decode(rawData []byte) (interface{}, error) {
	var data [][][]int
	if err := json.Unmarshal(rawData, &data); err != nil {
		//fmt.Println(err)
		return nil, err
	}
	return data, nil
}
