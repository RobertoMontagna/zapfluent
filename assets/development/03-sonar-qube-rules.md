# SonarQube Rules

> **Disclaimer:** The rules listed in this document are reverse-engineered from observations of SonarQube's analysis. They are intended to serve as a guideline for maintaining code quality and may not be an exhaustive or official list.

---

## `go:S100`: Function names should comply with a naming convention

*   **Quality:** Maintainability
*   **Severity:** Minor
*   **Type:** Code Smell
*   **Effort:** 5 min
*   **Tags:** `convention`

#### Description

*   Shared naming conventions are crucial for team collaboration and code readability.
*   This rule flags function names that do not adhere to a specified regular expression.
*   By default, the convention enforced is `^(_|[a-zA-Z0-9]+)$`, which supports `mixedCase` and `snake_case`.
