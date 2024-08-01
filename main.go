package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response struct {
    DayOrder string `json:"day_order"`
}

func main() {
    payload := []byte("{}")

    res, err := http.Post("https://academia-s-2.azurewebsites.net/do", "application/json", bytes.NewBuffer(payload))
    if err != nil {
        panic(err)
    }
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        fmt.Printf("Unexpected status code: %d\n", res.StatusCode)
        panic("Request Unsuccessful")
    }

    body, err := io.ReadAll(res.Body)
    if err != nil {
        panic(err)
    }

    var response Response
    err = json.Unmarshal(body, &response)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Day order today: %s\n", response.DayOrder)
}