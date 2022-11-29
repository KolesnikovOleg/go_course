package util

import (
    "fmt"
    "net/http"
)

func CheckIfResponseWrited(err error) {
    if err == nil {
        return
    }

    err = fmt.Errorf("can't write response: %s", err.Error())
    fmt.Println(err.Error())
}

func LogErrorAsResponse(w http.ResponseWriter, err error) {
    if err == nil {
        return
    }

    _, _ = w.Write([]byte(fmt.Errorf("request error: %s", err.Error()).Error()))
    w.WriteHeader(http.StatusInternalServerError)
}