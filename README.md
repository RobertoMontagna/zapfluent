# ✨ Zap Fluent Encoder

[![CI](https://github.com/RobertoMontagna/zapfluent/actions/workflows/ci.yml/badge.svg)](https://github.com/RobertoMontagna/zapfluent/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/RobertoMontagna/zapfluent/graph/badge.svg)](https://codecov.io/gh/RobertoMontagna/zapfluent)
[![Developed with Google Jules](https://img.shields.io/badge/Developed%20with-Google%20Jules-blue?logo=google&style=for-the-badge)](https://jules.google/)

A fluent interface encoder for Uber's Zap logging library that provides a more intuitive and expressive way to add structured logging fields. Say goodbye to endless `zap.String("key", "value")` calls and hello to a cleaner, more readable logging style!

## ⚠️ Work in Progress

This project is currently under active development. APIs and functionality may change without notice.

## 🚀 Usage

Here's a quick example of how to use `zapfluent` to create a `zapcore.ObjectMarshaler` with a fluent API.

First, define your struct:

```go
import (
	"go.uber.org/zap/zapcore"
	"go.robertomontagna.dev/zapfluent"
)

// User represents a user with some data.
type User struct {
	ID       int
	Username string
	Email    string // This field will be logged only if it's not empty.
}

// MarshalLogObject implements zapcore.ObjectMarshaler using the fluent API.
func (u User) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	return zapfluent.AsFluent(enc).
		Add(zapfluent.Int("id", u.ID)).
		Add(zapfluent.String("username", u.Username)).
		Add(zapfluent.String("email", u.Email).NonZero()). // Only log email if not empty!
		Done()
}
```

Now, you can use it with your Zap logger. `zapfluent` makes it easy to control which fields get logged based on their value.

```go
import "go.uber.org/zap"

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// This user has an email, so it will be included in the log.
	userWithEmail := User{
		ID:       123,
		Username: "jules",
		Email:    "jules@example.com",
	}
	logger.Info("User with full details", zap.Object("user", userWithEmail))
	// Log output: {"level":"info","ts":...,"caller":"...", "msg":"User with full details","user":{"id":123,"username":"jules","email":"jules@example.com"}}

	// This user has no email, so the email field will be omitted from the log.
	userWithoutEmail := User{
		ID:       456,
		Username: "jane",
	}
	logger.Info("User with partial details", zap.Object("user", userWithoutEmail))
	// Log output: {"level":"info","ts":...,"caller":"...", "msg":"User with partial details","user":{"id":456,"username":"jane"}}
}
```

## 🤝 Contributing

Contributions are welcome! This project is open to improvements and new features.

To get started, please check out the [Development Guidelines](docs/DEVELOPMENT.md).

### Local Development

The `Makefile` provides several useful targets for local development:

- `make fmt`: Format all Go files.
- `make lint`: Run the linter to check for code style issues.
- `make test`: Run all tests.
- `make coverage`: Generate a test coverage report.

Please run `make fmt`, `make lint`, and `make test` before submitting a pull request to ensure your changes meet the project's quality standards.

---
*This project is actively developed with the assistance of Google Jules, an AI-powered software engineer.*
