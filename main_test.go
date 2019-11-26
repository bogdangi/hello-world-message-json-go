package main

import (
    "fmt"
    "net/http"
    "net/http/httptest"
    "testing"
)

func Test_HelloWorld(t *testing.T) {
    req, err := http.NewRequest("GET", "http://localhost:3000/", nil)
    if err != nil {
        t.Fatal(err)
    }

    res := httptest.NewRecorder()
    HelloWorld(res, req)

    exp := fmt.Sprintf("{\"id\":\"1\",\"message\":\"Hello world\"}")
    act := res.Body.String()
    if exp != act {
        t.Fatalf("Expected %s got %s", exp, act)
    }
}

func Test_Healthz(t *testing.T) {
    req, err := http.NewRequest("GET", "http://localhost:3000/healthz", nil)
    if err != nil {
        t.Fatal(err)
    }

    res := httptest.NewRecorder()
    Healthz(res, req)

    exp := fmt.Sprintf("OK")
    act := res.Body.String()
    if exp != act {
        t.Fatalf("Expected %s got %s", exp, act)
    }
}
