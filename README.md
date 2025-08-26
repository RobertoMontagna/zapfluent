# ✨ Zap Fluent Encoder

[![CI](https://github.com/RobertoMontagna/zapfluent/actions/workflows/ci.yml/badge.svg)](https://github.com/RobertoMontagna/zapfluent/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/RobertoMontagna/zapfluent/graph/badge.svg)](https://codecov.io/gh/RobertoMontagna/zapfluent)
[![SonarCloud](https://sonarcloud.io/api/project_badges/measure?project=RobertoMontagna_zapfluent&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=RobertoMontagna_zapfluent)
[![Known Vulnerabilities](https://snyk.io/test/github/RobertoMontagna/ad606eb3-40a4-4157-9a53-77cb70724132/badge.svg)](https://app.snyk.io/org/robertomontagna/project/ad606eb3-40a4-4157-9a53-77cb70724132)
[![Developed with Google Jules](https://img.shields.io/badge/Developed%20with-Google%20Jules-blue?logo=google)](https://jules.google/)

<div align="center">

![Zap Fluent Encoder Logo](assets/images/fluentzap_logo.png)

A fluent interface encoder for [Uber's Zap logging library](https://github.com/uber-go/zap) that provides a more
intuitive and expressive way to add structured logging fields. Say goodbye to endless `zap.String("key", "value")` calls
and hello to a cleaner, more readable logging style!

</div>

## ⚠️ Work in Progress

This project is currently under active development. APIs and functionality may change without notice.

## 🚀 Usage

Here's a quick example of how to use `zapfluent` to create a `zapcore.ObjectMarshaler` with a fluent API.

First, define your struct:

```go
package main

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
package main

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

## 🚧 TODO

This is a non-exhaustive list of planned features and improvements, in no particular order.

- [ ] Add a configuration option to automatically sort logging fields in lexicographical order.
- [ ] Implement an error handling configuration to control the placement of field-specific error messages (e.g., in-place or at the end of the field list).
- [ ] Add a global configuration setting to automatically omit all zero-value fields from the log output.
- [ ] Expand field support to include all remaining primitive types (non-pointer, non-slice, non-array, non-map).
- [ ] Add support for logging pointer fields, automatically handling `nil` values.
- [ ] Implement support for logging slice fields.
- [ ] [Optional] Implement support for logging map fields.
- [X] Introduce Snyk or a similar service for dependency scanning.
- [ ] Introduce Renovate or a similar service for automated dependency updates.

## 🤝 Contributing

Contributions are welcome! This project is open to improvements and new features.

To get started, please check out the [Development Guidelines](assets/DEVELOPMENT.md).

### Local Development

The `Makefile` provides several useful targets for local development:

- `make fmt`: Format all Go files.
- `make lint`: Run the linter to check for code style issues.
- `make test`: Run all tests.
- `make coverage`: Generate a test coverage report.

Please run `make fmt`, `make lint`, and `make test` before submitting a pull request to ensure your changes meet the project's quality standards.

---
*This project is actively developed with the assistance of Google Jules, an AI-powered software engineer.*
