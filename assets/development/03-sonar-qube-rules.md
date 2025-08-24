# SonarQube Rules

> **Disclaimer:** The rules listed in this document are reverse-engineered from observations of SonarQube's analysis. They are intended to serve as a guideline for maintaining code quality and may not be an exhaustive or official list.

---

## `go:S100`: Function names should comply with a naming convention

*   **Rule ID:** `go:S100`
*   **Software qualities impacted:** Maintainability
*   **Severity:** Minor
*   **Effort:** 5 min
*   **Tags:** `convention`

### Why is this an issue?

Shared naming conventions allow teams to collaborate efficiently. This rule raises an issue when a function name does not match a provided regular expression.

For example, with the default provided regular expression `^(_|[a-zA-Z0-9]+)$`, the function:

```go
func execute_all() {
...
}
```
should be renamed to

```go
func executeAll() {
...
}
```

### Parameters
* **format**: Regular expression used to check the function names against.
* **Default Value**: `^(_|[a-zA-Z0-9]+)$`
