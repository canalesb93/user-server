package main

import (
    "fmt"
    "os"
    "os/signal"
    "log"

    "github.com/canalesb93/user-server/internal/user"
    "github.com/canalesb93/user-server/internal/database"
    "github.com/canalesb93/user-server/internal/server"
)

func main() {
    db, err := database.New()
    if err != nil {
        panic(fmt.Errorf("failed to connect to database: %w", err))
    }
    defer db.Close()

    // Do something with the database connection
    repo := user.NewUserRepository(db.Conn)

    // Create a dummy user
    u := user.DummyUser()
    err = u.Validate()
    if err != nil {
        log.Fatalf("error validating user: %v", err)
    }

    // Save the user to the database
    err = repo.SaveUser(u)
    fmt.Println("User saved successfully")
    if err != nil {
        log.Fatalf("could not save user: %v", err)
    }

    // Start server in separate goroutine
    srv := server.NewServer(db)
    go func() {
        srv.Start()
    }()

    // Listen for interrupt signal and gracefully shut down server
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    <-c

    err = srv.Shutdown()
    if err != nil {
        log.Fatal("Failed to shutdown server: ", err)
    }

    // Get all users from the database
    users, err := repo.GetAllUsers()
    if err != nil {
        log.Fatalf("could not get users: %v", err)
    }

    user.PrintUsers(users)
}
