package solver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"my_service/util"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

const username string = "Олег Колесников"
const serviceSolutionAddr string = "https://kuvaev-ituniversity.vps.elewise.com/tasks/"

func NewService() solveService {
	return solveService{
		solvers: []ISolverInterface{
			NewCycleShiftSolver("Циклическая ротация"),
			NewCheckSeqSolver("Проверка последовательности"),
			NewCheckPareSolver("Чудные вхождения в массив"),
			NewFindMissSolver("Поиск отсутствующего элемента"),
		},
	}
}

type solveService struct {
	solvers []ISolverInterface
}

func (s solveService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("new request | uri:", r.RequestURI, ", content-length:", r.ContentLength)

	if r.RequestURI == "/tasks" {

		responseArr := [][]byte{}
		var wg sync.WaitGroup
		for id, solver := range s.solvers {

			responseArr = append(responseArr, []byte(""))
			wg.Add(1)
			go func(idx int, solver ISolverInterface) {
				defer wg.Done()
				resp, err := s.requestSolutionAndCheck(solver)
				if err != nil {
					responseArr[idx] = []byte(err.Error())
					return
				}
				responseArr[idx] = append([]byte(solver.Description()), resp...)
			}(id, solver)
		}

		wg.Wait()
		_, err := w.Write(bytes.Join(responseArr, []byte("\n")))
		util.CheckIfResponseWrited(err)

		return
	}

	if strings.Contains(r.RequestURI, "/task/") {
		sections := strings.Split(r.RequestURI, "/")
		taskDesc := sections[len(strings.Split(r.RequestURI, "/"))-1]
		taskDesc, err := url.QueryUnescape(taskDesc)
		if err != nil {
			util.LogErrorAsResponse(w, err)
			return
		}
		fmt.Println("TASK: ", taskDesc)

		solver, ok := s.findSolver(taskDesc)
		if !ok {
			_, err = w.Write([]byte("UNDEFINED TASK!!!"))
			util.CheckIfResponseWrited(err)

			return
		}

		response, err := s.requestSolutionAndCheck(solver)
		if err != nil {
			util.LogErrorAsResponse(w, err)
			return
		}

		_, err = w.Write(response)
		util.CheckIfResponseWrited(err)

		return
	}

	_, err := w.Write([]byte("not found"))
	util.CheckIfResponseWrited(err)
	w.WriteHeader(http.StatusNotFound)
}

func (s solveService) requestSolutionAndCheck(solver ISolverInterface) ([]byte, error) {
	taskDesc := solver.Description()

	rawData, err := s.requestTaskData(taskDesc)
	if err != nil {
		return nil, err
	}

	decode, err := solver.Decode(rawData)
	if err != nil {
		return nil, err
	}

	response, err := s.requestCheckSolution(taskDesc, decode, solver.Solve(decode))

	return response, err
}

func (s solveService) requestTaskData(taskDesc string) ([]byte, error) {
	res, err := http.Get(serviceSolutionAddr + taskDesc)
	if err != nil {
		fmt.Println("error making http request task data:", err)
		return nil, err
	}

	resRaw, err := io.ReadAll(res.Body)

	return resRaw, err
}

func (s solveService) requestCheckSolution(desc string, payload interface{}, results interface{}) ([]byte, error) {

	values := map[string]interface{}{
		"user_name": username,
		"task":      desc,
		"results": map[string]interface{}{
			"payload": payload,
			"results": results,
		},
	}
	jsonData, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("json data: %s\n", jsonData)

	resp, err := http.Post(serviceSolutionAddr+"solution", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	respRaw, err := io.ReadAll(resp.Body)

	return respRaw, err

}

func (s *solveService) findSolver(desc string) (ISolverInterface, bool) {
	ok := false
	var solver ISolverInterface
	for _, v := range s.solvers {
		if v.Description() == desc {
			ok = true
			solver = v
			break
		}
	}

	return solver, ok
}
