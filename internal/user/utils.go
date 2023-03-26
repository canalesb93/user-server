package user

import (
    "fmt"
    "strings"
    "math/rand"
    "time"
)

// PrintUsers prints the given slice of users in a formatted manner
func PrintUsers(users []*User) {
    // Define column sizes
    idColSize := 8
    colSize := 25
    createdDateColSize := 15
    dateColSize := 25

    // Print header row
    fmt.Printf("|%s|%s|%s|%s|%s|\n",
        pad("ID", idColSize),
        pad("Name", colSize),
        pad("Email", colSize),
        pad("Created At", createdDateColSize),
        pad("Updated At", dateColSize),
    )

    // Print separator row
    fmt.Printf("|%s|%s|%s|%s|%s|\n",
        strings.Repeat("-", idColSize),
        strings.Repeat("-", colSize),
        strings.Repeat("-", colSize),
        strings.Repeat("-", createdDateColSize),
        strings.Repeat("-", dateColSize),
    )

    // Print each user
    for _, u := range users {
        fmt.Printf("|%s|%s|%s|%s|%s|\n",
            pad(fmt.Sprintf("%d", u.ID), idColSize),
            pad(u.Name, colSize),
            pad(u.Email, colSize),
            pad(u.CreatedAt.Format("01/02/2006"), createdDateColSize),
            pad(u.UpdatedAt.Format("01/02/2006 03:04:05 PM"), dateColSize),
        )
    }
}

func pad(s string, size int) string {
    // If s is longer than size, truncate it
    if len(s) > size {
        return s[:size-3] + "..."
    }

    // If s is shorter than size, pad it with spaces
    return s + strings.Repeat(" ", size-len(s))
}

// DummyUser generates a random User struct with dummy data
func DummyUser() *User {
    rand.Seed(time.Now().UnixNano())

    name := "User " + randomString(10)
    email := randomString(10) + "@example.com"

    return &User{
        Name:  name,
        Email: email,
    }
}

func randomString(length int) string {
    const charset = "abcdefghijklmnopqrstuvwxyz" +
        "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    b := make([]byte, length)
    for i := range b {
        b[i] = charset[rand.Intn(len(charset))]
    }
    return string(b)
}

