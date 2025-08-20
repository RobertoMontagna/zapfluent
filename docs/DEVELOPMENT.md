# Development

The rules presented here are a framework for software craftsmanship, informed by the principles of Clean Code and Go community best practices. They are not meant to be followed dogmatically, but rather to guide us in a continuous effort to improve the quality of our codebase.

Our responsibility is to leave the code cleaner than we found it—to fix "code smells" and apply these principles incrementally with every change. The ultimate measure of success is code that is readable and tells a clear story, even if it means making a pragmatic exception to a specific rule.

## General Principles

These principles are language-agnostic and form the core philosophy of writing clean, maintainable code.

#### Naming

* **Use Intention-Revealing Names:** Names should answer the big questions: why it exists, what it does, and how it's used. Avoid names that need a comment to explain their purpose.
* **Avoid Disinformation:** Don't use names that mislead. For example, `hp` shouldn't be used for a variable representing `horsepower` if it actually represents something else, like `health points`.
* **Make Meaningful Distinctions:** When two things are different, their names should be different. Avoid using arbitrary number series (`a1`, `a2`, ...) or noise words (`the`, `data`, `info`, `variable`) that don't add meaning.
* **Use Pronounceable Names:** Names should be easy to say and remember.
* **Use Searchable Names:** Use names that are easy to find in a codebase. A single letter like `i` is a poor choice for anything other than a loop index.
* **Class and Object Names:** Use nouns or noun phrases (**Customer**, **Account**, **WikiPage**). Avoid verbs.
* **Method Names:** Use verbs or verb phrases (**save**, **deletePage**, **postComment**).
* **One Word Per Concept:** Choose one word for a concept and stick with it. For example, don't mix `fetch`, `retrieve`, and `get` for the same kind of action.
* **Add Context to Names:** A variable named `state` is only meaningful if it has a context, like `addrState`.

#### Functions

* **Functions Should Be Small:** Functions should be very short, typically no more than a few lines.
* **Do One Thing:** A function should do one thing, and do it well. If a function can be described by more than one verb, it's likely doing too much.
* **Function Parameters:** The ideal number of arguments for a function is zero (niladic). Next comes one (monadic), followed closely by two (dyadic). Three arguments (triadic) should be avoided where possible. More than three requires very special justification. For functions requiring many parameters, consider bundling them into a single parameter object.
* **Function Return Values:** Functions should either be commands that perform an action and do not return a value, or queries that return a single, meaningful value without performing any side effects. Avoid using output arguments (passing pointers or references to be modified by the function).
* **Use Descriptive Names:** Prefer self-documenting code over comments. The name of the function should clearly describe what it does.
* **Avoid Side Effects:** A function shouldn't do something that isn't explicitly stated in its name. For example, a function named `isPasswordValid()` shouldn't also initialize a session.
* **Prefer Command-Query Separation:** Functions should either do something or answer something, but not both. For example, a `set()` function shouldn't also return a status code; its success should be implied unless an error is reported.
* **Don't Repeat Yourself (DRY):** Avoid duplicating code. If you find yourself writing the same logic in multiple places, extract it into a function.

#### Comments

* **Prefer Self-Documenting Code:** The best code is its own documentation. Comments should be a last resort.
* **Explain "Why," Not "What":** Use comments to explain **why** a particular decision was made, not to restate what the code is doing.
* **Do Not Comment Bad Code:** If code is unclear, refactor it. Don't add a comment to explain it.
* **Avoid Redundant Comments:** Don't comment on something that's already obvious from the code.
* **Do Not Use Commented-Out Code:** Delete commented-out code. Use version control to retrieve old code if needed.

#### Object and Data Structure Design

* **Classes Should Be Small:** A class should have a single, well-defined responsibility. As a general rule, a class with more than a few methods or instance variables is likely doing too much.
* **Encapsulation of Data:** Classes should hide their internal data structures and expose behavior through well-defined methods. Avoid creating classes that are just collections of public variables with no meaningful functions (known as Data Transfer Objects, which are a different concept).
* **Tell, Don't Ask:** Prefer telling objects what to do rather than asking them for their internal data and then acting on that data. This helps preserve encapsulation. For example, instead of `if (car.getSpeed() > speedLimit) { car.slowDown(); }`, prefer `car.checkAndRegulateSpeed(speedLimit);`.
* **Objects vs. Data Structures:** A core distinction exists between an object and a data structure. An **object** hides its data and exposes functions. A **data structure** exposes its data and has no meaningful functions. Clean code understands this difference and chooses the right approach for the problem at hand.
* **The Law of Demeter:** A class should only talk to its immediate friends. Avoid long chains of method calls like `a.getB().getC().doSomething()`. This indicates that your classes are overly coupled and lack proper encapsulation.
* **SOLID Principles:** Code should adhere to **SOLID principles**, with a strong emphasis on the **Single Responsibility Principle**.
* **Composition over Inheritance:** Prefer composition to achieve code reuse over inheritance.
* **Dependency Injection:** Use dependency injection to decouple components and make them more testable.

#### **Testing**

* **AAA Unit Test Structure:** All unit tests must be implicitly structured following the **Arrange-Act-Assert** (AAA) pattern.
* **Complete Test Coverage:** All production code must be covered by at least one unit test.
* **First-Class Tests:** Tests are as important as production code. They should be clean, readable, and well-organized.
* **Test One Concept:** A test should verify a single, specific behavior or concept. This makes the test's purpose immediately clear and aids in diagnosing failures.
    * **One Assert Guideline:** While a test should focus on one concept, it's a good practice to use only one assertion to enforce this. However, multiple assertions may be acceptable if they all verify different aspects of that *single* concept (e.g., asserting different fields on a returned object). The primary goal is diagnostic clarity.
    * **Custom Assertions:** Consider creating custom assertion functions if they are reusable across the project, simple to define (e.g., a low number of parameters), and have a very clear, unambiguous name and purpose. These can help enforce the "one concept" rule while simplifying the test code.
* **No Magic Values:** Avoid magic strings and numbers. If the same value is used more than once, introduce a constant for it.
* **Using Test Doubles:** Use test doubles (Fakes, Stubs, Mocks) to isolate the code under test and break external dependencies. Prefer using Fakes and Stubs to manage test state, as they are often more resilient to refactoring. Use Mocks judiciously, only when the primary goal is to verify a critical interaction.
* **Test Data Locality:** Place test data directly inside the test file itself. This improves readability and maintainability by keeping the input and expected output visible alongside the test logic. As a rule of thumb, it is fine to embed small payloads (e.g., JSON/XML under 50 lines or simple strings) directly. Consider externalizing for:
    * **Large Payloads:** JSON or XML files over 50 lines, or multi-line strings that exceed the screen height.
    * **Binary Data:** Images, videos, or other non-text files.
    * **Extensive Scripts:** SQL scripts or other large text dumps.
    * **External Test Data:** When externalizing, store the data in a `testdata/` directory within the same package.

#### System Boundaries

* **Validate Inputs and Boundaries:** All data entering the system from an external source (e.g., user input, API requests, configuration files) must be validated at the system boundary before it is used. This prevents bugs and security vulnerabilities.
* **Wrap Third-Party Code:** Isolate third-party APIs and libraries behind an interface or a wrapper class that you control. This prevents changes in an external library from propagating throughout your codebase.

---

## Go-Specific Guidelines

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