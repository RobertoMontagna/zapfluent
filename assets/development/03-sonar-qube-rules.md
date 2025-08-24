# SonarQube Rules

> **Disclaimer:** The rules listed in this document are reverse-engineered from observations of SonarQube's analysis. They are intended to serve as a guideline for maintaining code quality and may not be an exhaustive or official list.

---

## `go:S100`: Function names should comply with a naming convention

*   **Rule ID:** `go:S100`
*   **Software qualities impacted:** Maintainability
*   **Severity:** Minor
*   **Effort:** 5 min
*   **Tags:** `convention`
*   **Tool:** Sonar (Go)
*   **Type:** Code Smell

### Description

Shared naming conventions allow teams to collaborate efficiently. This rule raises an issue when a function name does not match the provided regular expression.

The required format is `camelCase` for unexported functions and `PascalCase` for exported functions. The regular expression used to enforce this is:

`^(_|[a-zA-Z0-9]+)$`

### Example

With the default regular expression, a function like this:

```go
func execute_all() { // Noncompliant
...
}
```

should be renamed to:

```go
func executeAll() { // Compliant
...
}
```
