package auth

import "fmt"

// if authLoginWithCredentials (like private func) => a small then it means this function is use only in the packages or file
// if AuthLoginWithCredentials (like public func) => A caps then it means this funcion is use outoff package or files like export the scope

// this is public function
func AuthLoginWithCredentials(userName string, password string) {
	fmt.Println("Login Sucessfully", userName, password)
}

// go mod tidy

// The go mod tidy command is used in Go modules to ensure that your module's dependencies are clean, correct, and minimal. It is a maintenance command that helps manage the go.mod and go.sum files in your project.

// Remove Unused Dependencies
// Add Missing Dependencies:
// Update go.sum:
