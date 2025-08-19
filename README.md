# Zap Fluent Encoder

A fluent interface encoder for Uber's Zap logging library that provides a more intuitive way to add structured logging
fields.

## ⚠️ Work in Progress

This project is currently under active development. APIs and functionality may change without notice.

## Development

This library is being developed with the assistance of Jules by Google, an AI-powered software engineer. The following rules are being followed during development:

These rules serve as a strict guideline to ensure code quality, readability, and maintainability. However, they are not unbreakable laws; exceptions can be made when adhering to a rule would detract from the overall quality of the code.


### All Code

#### **Naming**

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

#### **Functions**

* **Functions Should Be Small:** Functions should be very short, typically no more than a few lines.
* **Do One Thing:** A function should do one thing, and do it well. If a function can be described by more than one verb, it's likely doing too much.
* **Max 2 Parameters:** Functions and methods should have at most two parameters. If a third is required, it should typically be a **`context.Context`** or a struct for more complex arguments.
* **Use Descriptive Names:** Prefer self-documenting code over comments. The name of the function should clearly describe what it does.
* **No Unnecessary Pointers:** Avoid using pointers for optional values or to pass arguments. Pointers should only be used when strictly necessary (e.g., for channels, or for large objects where copying is a performance concern). Prefer a value-based approach or an `Optional` monad for optional values.
* **Avoid Side Effects:** A function shouldn't do something that isn't explicitly stated in its name. For example, a function named `isPasswordValid()` shouldn't also initialize a session.
* **Prefer Command-Query Separation:** Functions should either do something or answer something, but not both. For example, a `set()` function shouldn't also return a status code.
* **Don't Repeat Yourself (DRY):** Avoid duplicating code. If you find yourself writing the same logic in multiple places, extract it into a function.

#### **Comments**

* **Prefer Self-Documenting Code:** The best code is its own documentation. Comments should be a last resort.
* **Explain "Why," Not "What":** Use comments to explain **why** a particular decision was made, not to restate what the code is doing.
* **Do Not Comment Bad Code:** If code is unclear, refactor it. Don't add a comment to explain it.
* **Avoid Redundant Comments:** Don't comment on something that's already obvious from the code.
* **Do Not Use Commented-Out Code:** Delete commented-out code. Use version control to retrieve old code if needed.

#### **Formatting and Organization**

* **Vertical Formatting:** Strive to keep files short, with a strong preference for high cohesion by grouping related code together.
* **Code Locality and Cohesion:** Group related code together. Avoid mixing unrelated concepts in the same file or package. For example, a generic `constants.go` file is discouraged as it tends to aggregate values that have little in common besides being constants.
* **Standard Import Formatting:** All import blocks must be grouped into four categories in a specific order: standard library, third-party, shared internal modules, and intra-module dependencies.

#### **Error Handling**

* **Don't Return `nil`:** Avoid returning `nil` from functions that could fail. Prefer returning a zero value or an empty slice/map.
* **Use Exceptions (or Error Types) over Error Codes:** When possible, use specific error types rather than returning generic error codes or boolean values.

---

### **Test Code Rules**

* **AAA Unit Test Structure:** All unit tests must be implicitly structured following the **Arrange-Act-Assert** (AAA) pattern.
* **Complete Test Coverage:** All production code must be covered by at least one unit test.
* **One Assert per Test:** Each test function should test one concept and have a single assertion. This makes it clear what the test is for and what failed if it breaks.
* **First-Class Tests:** Tests are as important as production code. They should be clean, readable, and well-organized.
* **No Magic Values:** Avoid magic strings and numbers. If the same value is used more than once, introduce a constant for it.

---

### **Principles and General Rules**

* **SOLID Principles:** Code should adhere to **SOLID principles**, with a strong emphasis on the **Single Responsibility Principle**.
* **Composition over Inheritance:** Prefer composition to achieve code reuse over inheritance.
* **Dependency Injection:** Use dependency injection to decouple components and make them more testable.
