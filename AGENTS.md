# Google Jules Directives

This document outlines the workflow for developing and maintaining this project with the assistance of Google Jules.

## 1. Initial Assessment

Upon starting a new task, perform the following steps:

1.  **Read Development Guidelines**: Thoroughly read all documents in the `assets/development/` directory to understand the project's coding standards.
2.  **Full Repository Scan**: Perform a full scan of the repository to assess compliance with the development guidelines.
3.  **Generate Initial Report**:
    *   Generate an initial response listing all possible improvements and enhancements to make the project compliant with the guidelines.
    *   **Do not** generate or modify any files at this stage.
    *   Present this report to the user for review and wait for approval before proceeding.

## 2. Development Cycle

For each set of approved changes, follow this cycle:

1.  **Implement Changes**: Make the required code modifications.
2.  **Run Quality Checks**: After completing a set of changes, you **must** run the following commands using the Makefile:
    *   `make test`: Run all tests to ensure no regressions were introduced.
    *   `make lint-fix`: Automatically fix any linting issues.
    *   `make fmt`: Format the code according to the project style.
3.  **Generate Commit**: Once all checks pass, generate a single, cohesive commit with a clear and descriptive message.
