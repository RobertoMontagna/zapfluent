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

*   Shared naming conventions allow teams to collaborate efficiently.
*   This rule raises an issue when a function name does not match the provided regular expression.
*   The required format is `camelCase` for unexported functions and `PascalCase` for exported functions.
*   The regular expression used to enforce this is: `^(_|[a-zA-Z0-9]+)$`

---

## `go:S1186`: Functions should not be empty

*   **Rule ID:** `go:S1186`
*   **Software qualities impacted:** Maintainability
*   **Severity:** Critical
*   **Effort:** 5 min
*   **Tags:** `suspicious`
*   **Tool:** Sonar (Go)
*   **Type:** Code Smell

### Description

*   An empty function is generally considered bad practice and can lead to confusion, readability, and maintenance issues.
*   Empty functions bring no functionality and are misleading to others as they might think the function implementation fulfills a specific and identified requirement.
*   There are several reasons for a function not to have a body:
    *   It is an unintentional omission, and should be fixed to prevent an unexpected behavior in production.
    *   It is not yet, or never will be, supported. In this case an exception should be thrown.
    *   The method is an intentionally-blank override. In this case a nested comment should explain the reason for the blank override.
