package zapfluent_test

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent"
)

// Address represents a street address.
type Address struct {
	Street string
	City   string
	Zip    string
}

// MarshalLogObject makes Address implement zapcore.ObjectMarshaler.
func (a Address) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	return zapfluent.AsFluent(enc).
		Add(zapfluent.String("street", a.Street)).
		Add(zapfluent.String("city", a.City)).
		Add(zapfluent.String("zip", a.Zip)).
		Done()
}

// User represents a user with personal information.
type User struct {
	ID       int
	Name     string
	IsActive bool
	Address  Address
	Tags     []string
}

// MarshalLogObject makes User implement zapcore.ObjectMarshaler.
func (u User) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	return zapfluent.AsFluent(enc).
		Add(zapfluent.Int("id", u.ID)).
		Add(zapfluent.String("name", u.Name)).
		Add(zapfluent.Bool("isActive", u.IsActive)).
		Add(zapfluent.Object("address", u.Address, func(a Address) bool { return a != Address{} })).
		Add(zapfluent.String("tags", strings.Join(u.Tags, ","))).
		Done()
}

func Example_withComplexObject() {
	logger, _ := zap.NewProduction()
	_ = logger.Sync()

	user := User{
		ID:       123,
		Name:     "John Doe",
		IsActive: true,
		Address: Address{
			Street: "123 Main St",
			City:   "Anytown",
			Zip:    "12345",
		},
		Tags: []string{"go", "logging", "zap"},
	}

	logger.Info("Logging a complex, nested object", zap.Object("user", user))

	// In a real application, the output would be a JSON log line.
	// For this example, we just demonstrate the usage.
	// Output:
}
