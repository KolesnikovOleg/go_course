package solver

type ISolverInterface interface {
	Description() string
	Decode(rawData []byte) (interface{}, error)
	Solve(interface{}) interface{}
}
