# Repository-Specific Guidelines

This document contains guidelines and conventions that are specific to this repository.

## General

* **Dot Imports**
    * **Rule**: Dot imports (`import . "pacakge"`) are strongly discouraged. They can make code confusing by obscuring which package a function or type belongs to, breaking the usual `package.Function` pattern that is standard in Go. Always prefer using a package alias if the package name is too long or conflicts with another name.
    * **Exception**: An exception is made for certain testing libraries like Gomega, where dot imports are idiomatic and significantly improve test readability. This practice should be used with caution and confined to `_test.go` files.
* **Functional Programming**
    * **Rule**: While Go is not a functional language, this project encourages the use of functional programming concepts where they improve code clarity and expressiveness. This includes the use of higher-order functions, immutable data structures, and constructs like `Optional` types to handle the absence of a value.
    * **Guideline**: When using these patterns, ensure that the code remains readable and maintainable. The goal is to leverage functional concepts to write more robust and predictable code, not to obscure logic with overly complex abstractions.

## Testing

* **Test Style**
    * **Rule**: All tests in this repository must be written in the standard XUnit style, using the built-in `testing` package. BDD-style frameworks like Ginkgo are not permitted.
    * **Guideline**: Prefer table-driven tests for their clarity and ease of extension. Use Gomega for assertions to improve readability, but do not use the Ginkgo test runner.
    * **Note**: The use of Gomega is a deliberate exception to the general guideline in `02-go-guidelines.md` which discourages assertion helpers that call `t.Fatal`. Gomega's expressive syntax is prioritized for readability in this project's tests.
