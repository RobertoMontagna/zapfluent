# Go-Specific Guidelines

These guidelines are tailored to the idioms and conventions of the Go programming language.

### Package and Project Organization

* **Package Naming:** Choose package names that are short, concise, and all lowercase. Avoid underscores and mixed case. The name should be descriptive of the package's purpose. Avoid generic names like `util`, `common`, or `lib`.
* **Package Size and Cohesion:** A package should have a single, well-defined responsibility. Group related types and functions together. If two packages are so tightly coupled that they are almost always used together, consider merging them.
* **Project Layout:** For consistency, consider following the [Standard Go Project Layout](https://github.com/golang-standards/project-layout). This is not an official standard, but it is a community-accepted convention.

### Error Handling

Proper error handling is critical for writing robust Go code.

* **Use the `(result, error)` Pattern:** This is the idiomatic Go way to handle fallible operations. It cleanly separates the expected return value from the error status, aligning with the Command-Query Separation principle.
* **Error Types:**
    * For simple, static error messages, use `errors.New`.
    * For dynamic error messages, use `fmt.Errorf`.
    * For errors that need to be handled programmatically by callers, define a custom error type that implements the `error` interface. This allows callers to use `errors.As` to inspect the error's details.
    * For sentinel errors that callers can check with `errors.Is`, declare them as exported variables (e.g., `var ErrNotFound = errors.New("not found")`).
* **Error Wrapping:**
    * When propagating an error from a downstream function, add context using `fmt.Errorf` with the `%w` verb. This allows callers to inspect the underlying error chain with `errors.Is` and `errors.As`.
    * Keep the context concise. Avoid phrases like "failed to", which are redundant.
* **Error Naming:**
    * Error variables should be prefixed with `Err` (e.g., `ErrNotFound`).
    * Custom error types should be suffixed with `Error` (e.g., `NotFoundError`).
* **Handling Errors Once:** An error should be handled only once. If you log an error, don't also return it up the call stack, as this can lead to duplicate logging. Decide at each level whether to handle the error (e.g., by logging it and returning a default value) or to propagate it.

### Concurrency

* **Avoid Shared Mutable State:** Prefer sharing memory by communicating, rather than communicating by sharing memory. This is a core Go philosophy.
* **Immutability and State Management:** Favor immutable data structures to reduce the surface area for bugs in concurrent code. If state must be shared, manage it explicitly with channels or mutexes.
* **Goroutines in Loops:** When launching a goroutine in a loop, be careful about capturing the loop variable. The goroutine may not execute until the loop has finished, causing all goroutines to use the final value of the variable. Pass the variable's value as an explicit parameter or reassign it to a local variable inside the loop.
* **Know Your Concurrency Primitives:** Understand the Go concurrency model, including goroutines, channels, and the `sync` package primitives (Mutexes, WaitGroups).
* **Isolate Concurrent Code:** The code that uses concurrency should be isolated from the rest of the application. Separate code that handles goroutines, channels, or locks into its own logical unit to prevent the rest of the system from having to know about it.
* **Goroutine Lifetime:** Never start a goroutine without knowing when it will exit. Leaked goroutines can lead to memory leaks and other issues. Use a `sync.WaitGroup` to wait for goroutines to finish, or use a `context` for cancellation.
* **Channel Usage:** Channels should usually have a size of one or be unbuffered (size zero). Any other size should be carefully considered to avoid deadlocks and other issues.
* **Mutexes:** Use the zero-value of a `sync.Mutex` or `sync.RWMutex`. Do not use a pointer to a mutex. Do not embed mutexes in public structs, as this exposes implementation details.

### Context

The `context` package is essential for managing deadlines, cancellation, and request-scoped data in concurrent programs.

* **Pass `context.Context` as the First Argument:** For functions that may block (e.g., I/O, API calls), accept a `context.Context` as the first argument. This is an idiomatic exception to the "keep argument count low" principle, making three-argument functions common.
* **Use Context for Cancellation, Not Parameters:** The primary purpose of a `context` is to signal cancellation or deadlines across API boundaries. Do not use it as a general-purpose property bag to pass optional parameters; use a dedicated struct for that.
* **Use `context.WithValue` Sparingly:** Only use `context.WithValue` for request-scoped data that is necessary for downstream logic, such as a request ID or trace information. The keys used should be of an unexported type to avoid collisions.

### Interfaces and Types

* **Accept Interfaces, Return Structs:** Functions and methods should be designed to accept a general interface as input to be more flexible and testable. However, they should return a concrete struct to give callers a clear understanding of the return type and its full capabilities.
* **Don't Prefix Interface Names:** It is not idiomatic in Go to prefix interfaces with an `I` (e.g., use `io.Reader`, not `IReader`). The name of the interface should describe its behavior.
* **Strive for Useful Zero Values:** A powerful Go idiom is to make the zero value of a struct useful. This means a newly declared variable (e.g., `var s MyStruct`) should be in a valid, ready-to-use state without requiring explicit initialization.

### Generics (Go 1.18+)

* **When to Use Generics:** Use generics for functions and data structures that work with a collection of some type, where the logic is identical for all supported types (e.g., a function that works on a slice of any type, or a generic graph data structure).
* **When to Prefer Interfaces:** Prefer interfaces when you need to abstract behavior. If different types share a common method set (e.g., they all have a `Read()` method), an interface is usually the cleaner, more idiomatic solution. Do not use generics simply to avoid a small amount of boilerplate.

### Testing

* **Table-Driven Tests:** Use table-driven tests to avoid duplicating code when testing multiple scenarios. This makes it easy to add new test cases and improves readability.
    * **Test Data in Tables:** For data specific to a test case, define it directly within the table's struct literal. This keeps the test cohesive and the data visible. For data that would still make the table unreadable (e.g., large JSON payloads), use an external file from the `testdata/` directory.
* **Test Helpers:** Distinguish between test setup helpers and assertion helpers. It is idiomatic in Go to avoid assertion helpers that call `t.Fatal` or `t.Error`. Instead, return an error from your validation function and let the test function decide whether to fail the test.
* **`t.Fatal` vs. `t.Error`:** Use `t.Fatal` when a test cannot continue because a setup step has failed. Use `t.Error` when a test case has failed but other test cases can still be run.

### Performance

* **Use `strconv`:** When converting primitives to and from strings, `strconv` is generally faster than `fmt`.
* **Pre-allocate Slices and Maps:** When you know the size of a slice or map in advance, pre-allocate it using `make` with a capacity hint. This can significantly reduce the number of allocations.
* **Avoid Repeated String-to-Byte Conversions:** If you need to use the byte representation of a fixed string multiple times, convert it to a `[]byte` once and reuse it.

### Documentation (godoc)

* **Document All Exported Identifiers:** All exported packages, types, functions, constants, and variables must have a clear `godoc` comment.
* **Explain What, Not How:** The documentation should explain what the code does from a caller's perspective and how to use it correctly. Implementation details should be left to inline comments if necessary.

### Linting

To ensure code quality and consistency, use a standard set of linters. We recommend using `golangci-lint` as a lint runner with the following linters enabled at a minimum:
* `errcheck`: Checks for unhandled errors.
* `goimports`: Formats code and manages imports.
* `golint`: Points out common style mistakes.
* `govet`: Analyzes code for common mistakes.
* `staticcheck`: Provides a wide range of static analysis checks.

### Formatting and Organization

* **Gofmt:** All Go code in the repository must be formatted with `gofmt`.
* **Import Grouping:** All import blocks must be grouped into four categories in a specific order: standard library, third-party, shared internal modules, and intra-module dependencies.
* **Reduce Nesting:** Avoid deep nesting by handling error cases and special conditions first and returning early.
* **Function Grouping:** Functions in a file should be grouped by receiver. Exported functions should appear first, after type, const, and var definitions.
* **Variable Scope:** Reduce the scope of variables as much as possible. If a variable is only used inside an `if` block, declare it inside the `if`.
* **Raw String Literals:** Use raw string literals (backticks) to avoid escaping quotes and to write multi-line strings.
* **Initializing Structs:**
    * Always use field names when initializing structs.
    * Omit zero-value fields unless they provide meaningful context.
    * Use the `var` form to declare zero-value structs (e.g., `var u User`).
    * Use `&T{}` instead of `new(T)` to initialize struct references.
* **Initializing Maps:**
    * Use `make` for empty maps and maps that are populated programmatically.
    * Use map literals for maps with a fixed set of elements.
* **Printf-style Functions:** If you declare a `Printf`-style function, name it with an `f` suffix (e.g., `Wrapf`) so that `go vet` can check the format string.
