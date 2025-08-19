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
* **Don't Add a Prefix to Interface Names:** It's generally not necessary to prefix interfaces with an `I` or similar. The name of the interface should be sufficient.

#### Functions

* **Functions Should Be Small:** Functions should be very short, typically no more than a few lines.
* **Do One Thing:** A function should do one thing, and do it well. If a function can be described by more than one verb, it's likely doing too much.
* **Function Parameters:** The ideal number of arguments for a function is zero (niladic). Next comes one (monadic), followed closely by two (dyadic). Three arguments (triadic) should be avoided where possible. More than three requires very special justification — and then shouldn't be used anyway.
* **Function Return Values:** Functions should either be commands that perform an action and do not return a value, or queries that return a single, meaningful value without performing any side effects. Avoid using output arguments.
* **Use Descriptive Names:** Prefer self-documenting code over comments. The name of the function should clearly describe what it does.
* **Avoid Side Effects:** A function shouldn't do something that isn't explicitly stated in its name. For example, a function named `isPasswordValid()` shouldn't also initialize a session.
* **Prefer Command-Query Separation:** Functions should either do something or answer something, but not both. For example, a `set()` function shouldn't also return a status code.
* **Don't Repeat Yourself (DRY):** Avoid duplicating code. If you find yourself writing the same logic in multiple places, extract it into a function.

#### Comments

* **Prefer Self-Documenting Code:** The best code is its own documentation. Comments should be a last resort.
* **Explain "Why," Not "What":** Use comments to explain **why** a particular decision was made, not to restate what the code is doing.
* **Do Not Comment Bad Code:** If code is unclear, refactor it. Don't add a comment to explain it.
* **Avoid Redundant Comments:** Don't comment on something that's already obvious from the code.
* **Do Not Use Commented-Out Code:** Delete commented-out code. Use version control to retrieve old code if needed.

#### Class and Object Design

* **Classes Should Be Small:** A class should have a single, well-defined responsibility. As a general rule, a class with more than a few methods or instance variables is likely doing too much.
* **Encapsulation of Data:** Classes should hide their internal data structures and expose behavior through well-defined methods. Avoid creating classes that are just collections of public variables with no meaningful functions (known as Data Transfer Objects, which are a different concept).
* **Tell, Don't Ask:** Prefer telling objects what to do rather than asking them for their internal data and then acting on that data. This helps preserve encapsulation. For example, instead of `if (car.getSpeed() > speedLimit) { car.slowDown(); }`, prefer `car.checkAndRegulateSpeed(speedLimit);`.
* **Objects vs. Data Structures:** A core distinction exists between an object and a data structure. An **object** hides its data and exposes functions. A **data structure** exposes its data and has no meaningful functions. Clean code understands this difference and chooses the right approach for the problem at hand.
* **The Law of Demeter:** A class should only talk to its immediate friends. Avoid long chains of method calls like `a.getB().getC().doSomething()`. This indicates that your classes are overly coupled and lack proper encapsulation.
* **SOLID Principles:** Code should adhere to **SOLID principles**, with a strong emphasis on the **Single Responsibility Principle**.
* **Composition over Inheritance:** Prefer composition to achieve code reuse over inheritance.
* **Dependency Injection:** Use dependency injection to decouple components and make them more testable.

#### System Boundaries

* **Wrap Third-Party Code:** Isolate third-party APIs and libraries behind an interface or a wrapper class that you control. This prevents changes in an external library from propagating throughout your codebase.

---

## Go-Specific Guidelines

These guidelines are tailored to the idioms and conventions of the Go programming language.

### Package and Project Organization

*   **Package Naming:** Choose package names that are short, concise, and all lowercase. Avoid underscores and mixed case. The name should be descriptive of the package's purpose. Avoid generic names like `util`, `common`, or `lib`.
*   **Package Size and Cohesion:** A package should have a single, well-defined responsibility. Group related types and functions together. If two packages are so tightly coupled that they are almost always used together, consider merging them.
*   **Project Layout:** For consistency, consider following the [Standard Go Project Layout](https://github.com/golang-standards/project-layout). This is not an official standard, but it is a community-accepted convention.

### Error Handling

Proper error handling is critical for writing robust Go code.

*   **Error Types:**
    *   For simple, static error messages, use `errors.New`.
    *   For dynamic error messages, use `fmt.Errorf`.
    *   For errors that need to be handled programmatically by callers, define a custom error type that implements the `error` interface. This allows callers to use `errors.As` to inspect the error's details.
    *   For sentinel errors that callers can check with `errors.Is`, declare them as exported variables (e.g., `var ErrNotFound = errors.New("not found")`).
*   **Error Wrapping:**
    *   When propagating an error from a downstream function, add context using `fmt.Errorf` with the `%w` verb. This allows callers to inspect the underlying error chain with `errors.Is` and `errors.As`.
    *   Keep the context concise. Avoid phrases like "failed to", which are redundant.
*   **Error Naming:**
    *   Error variables should be prefixed with `Err` (e.g., `ErrNotFound`).
    *   Custom error types should be suffixed with `Error` (e.g., `NotFoundError`).
*   **Handling Errors Once:** An error should be handled only once. If you log an error, don't also return it up the call stack, as this can lead to duplicate logging. Decide at each level whether to handle the error (e.g., by logging it and returning a default value) or to propagate it.

### Concurrency

*   **Goroutine Lifetime:** Never start a goroutine without knowing when it will exit. Leaked goroutines can lead to memory leaks and other issues. Use a `sync.WaitGroup` to wait for goroutines to finish.
*   **Channel Usage:** Channels should usually have a size of one or be unbuffered (size zero). Any other size should be carefully considered to avoid deadlocks and other issues.
*   **Mutexes:** Use the zero-value of a `sync.Mutex` or `sync.RWMutex`. Do not use a pointer to a mutex. Do not embed mutexes in public structs, as this exposes implementation details.

### Testing

*   **Table-Driven Tests:** Use table-driven tests to avoid duplicating code when testing multiple scenarios. This makes it easy to add new test cases and improves readability.
*   **Test Helpers:** Distinguish between test setup helpers and assertion helpers. It is idiomatic in Go to avoid assertion helpers that call `t.Fatal` or `t.Error`. Instead, return an error from your validation function and let the test function decide whether to fail the test.
*   **`t.Fatal` vs. `t.Error`:** Use `t.Fatal` when a test cannot continue because a setup step has failed. Use `t.Error` when a test case has failed but other test cases can still be run.

### Performance

*   **Use `strconv`:** When converting primitives to and from strings, `strconv` is generally faster than `fmt`.
*   **Pre-allocate Slices and Maps:** When you know the size of a slice or map in advance, pre-allocate it using `make` with a capacity hint. This can significantly reduce the number of allocations.
*   **Avoid Repeated String-to-Byte Conversions:** If you need to use the byte representation of a fixed string multiple times, convert it to a `[]byte` once and reuse it.

### Linting

To ensure code quality and consistency, use a standard set of linters. We recommend using `golangci-lint` as a lint runner with the following linters enabled at a minimum:
*   `errcheck`: Checks for unhandled errors.
*   `goimports`: Formats code and manages imports.
*   `golint`: Points out common style mistakes.
*   `govet`: Analyzes code for common mistakes.
*   `staticcheck`: Provides a wide range of static analysis checks.

### Formatting and Organization

*   **Gofmt:** All Go code in the repository must be formatted with `gofmt`.
*   **Import Grouping:** Imports should be grouped into two blocks: standard library and everything else.
*   **Reduce Nesting:** Avoid deep nesting by handling error cases and special conditions first and returning early.
*   **Function Grouping:** Functions in a file should be grouped by receiver. Exported functions should appear first, after type, const, and var definitions.
*   **Variable Scope:** Reduce the scope of variables as much as possible. If a variable is only used inside an `if` block, declare it inside the `if`.
*   **Raw String Literals:** Use raw string literals (backticks) to avoid escaping quotes and to write multi-line strings.
*   **Initializing Structs:**
    *   Always use field names when initializing structs.
    *   Omit zero-value fields unless they provide meaningful context.
    *   Use the `var` form to declare zero-value structs (e.g., `var u User`).
    *   Use `&T{}` instead of `new(T)` to initialize struct references.
*   **Initializing Maps:**
    *   Use `make` for empty maps and maps that are populated programmatically.
    *   Use map literals for maps with a fixed set of elements.
*   **Printf-style Functions:** If you declare a `Printf`-style function, name it with an `f` suffix (e.g., `Wrapf`) so that `go vet` can check the format string.
