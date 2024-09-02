package user

// 1. Email (Capitalized)
// Exported Identifier: In Go, any identifier (function, type, variable, etc.) that begins with a capital letter is exported. This means it is accessible from other packages.

// 2. email (Lowercase)
// Unexported Identifier: An identifier that starts with a lowercase letter is unexported. This means it is only accessible within the package where it is defined.

type User struct {
	Email string
	Name  string
}
