package database

import (
    "math/rand"
    "time"
)

// GenerateRandomID generates a random integer ID.
func GenerateRandomID() int64 {
    rand.Seed(time.Now().UnixNano())
    min := 10000000
    max := 99999999
    return int64(rand.Intn(max-min+1) + min)
}