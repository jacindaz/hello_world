package main
import (
    "testing"
)

func TestFail(t *testing.T) {
    t.Log("logging something cool")
    t.FailNow()
}
