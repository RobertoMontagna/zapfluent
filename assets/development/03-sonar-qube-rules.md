# SonarQube Rules

> **Disclaimer:** The rules listed in this document are reverse-engineered from observations of SonarQube's analysis. They are intended to serve as a guideline for maintaining code quality and may not be an exhaustive or official list.

---

## `go:S100`: Function names should comply with a naming convention

*   **Rule ID:** `go:S100`
*   **Repository:** Sonar (Go)
*   **Type:** Code Smell
*   **Severity:** Minor
*   **Effort:** 5 min
*   **Tags:** `convention`

### Why is this an issue?

Shared naming conventions allow teams to collaborate efficiently. This rule helps to ensure that all function names within the project are consistent by raising an issue when a function name does not match a provided regular expression.

By default, the convention for Go is to use `camelCase` for internal functions and `PascalCase` for exported functions. This rule enforces that standard.

#### Noncompliant Code Example

```go
func My_function() { // Noncompliant; contains an underscore
  // ...
}
```

#### Compliant Code Example

```go
func MyFunction() { // Compliant
  // ...
}

func myPrivateFunction() { // Compliant
  // ...
}
```
