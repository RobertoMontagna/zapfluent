package zapfluent_test

import (
	"strings"

	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent"
	"go.robertomontagna.dev/zapfluent/testutil"
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

//revive:disable:line-length-limit
func Example_withComplexObject() {
	logger := testutil.StdOutLoggerForTest()

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

	logger.Infow("Logging a complex, nested object", "user", user)

	// Output:
	//{"level":"info","msg":"Logging a complex, nested object","user":{"id":123,"name":"John Doe","isActive":true,"address":{"street":"123 Main St","city":"Anytown","zip":"12345"},"tags":"go,logging,zap"}}
}

//revive:enable:line-length-limit
