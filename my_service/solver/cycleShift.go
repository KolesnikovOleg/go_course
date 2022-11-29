package solver

import (
	"encoding/json"
)

type shiftStruct struct {
	array  []int
	number int
}

func (sh *shiftStruct) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{sh.array, sh.number})
}

func (sh *shiftStruct) UnmarshalJSON(p []byte) error {
	var tmp []json.RawMessage
	if err := json.Unmarshal(p, &tmp); err != nil {
		return err
	}
	if err := json.Unmarshal(tmp[0], &sh.array); err != nil {
		return err
	}
	if err := json.Unmarshal(tmp[1], &sh.number); err != nil {
		return err
	}
	return nil
}

type ShifterData []shiftStruct

func NewCycleShiftSolver(desc string) CycleShiftSolver {
	return CycleShiftSolver{
		description: desc,
	}
}

type CycleShiftSolver struct {
	description string
}

func (c CycleShiftSolver) Solution(A []int, K int) []int {
	//Если K превышает длину массива, заменяем K на остаток от деления на длинну
	if K > len(A) {
		K = (K % len(A))
	}
	var startShift int = len(A) - K
	A = append(A[startShift:], A[:startShift]...)
	return A
}

func (c CycleShiftSolver) Solve(data interface{}) interface{} /*[][]int */ {
	res := [][]int{}
	for _, value := range data.(ShifterData) {
		res = append(res, c.Solution(value.array, value.number))
	}

	return res
}

func (c CycleShiftSolver) Description() string {
	return c.description
}

func (c CycleShiftSolver) Decode(rawData []byte) (interface{} /*ShifterData*/, error) {
	var data ShifterData
	if err := json.Unmarshal(rawData, &data); err != nil {
		//fmt.Println(err)
		return nil, err
	}
	return data, nil
}
