# Google Jules Directives

This document outlines the workflow for developing and maintaining this project with the assistance of Google Jules.

## 1. Initial Assessment

Upon starting a new task, perform the following steps:

1.  **Read Development Guidelines**: Thoroughly read all documents in the `assets/development/` directory to understand the project's coding standards.
2.  **Full Repository Scan**: Perform a full and detailed scan of the repository to assess compliance with the development guidelines.
3.  **Ask me**: How I want to proceed? {New development, Assess current state}
    **New Development**: Wait for my input and then start the normal development cycle.
    **Assess Current State**: Assess the current state of the project and generate an initial report [see point 4].
4.  **Generate Initial Report**:
    *   Generate an initial response listing all possible improvements and enhancements to make the project compliant with the guidelines.
    *   **Do not** generate or modify any files at this stage.
    *   Present the entirety of report to the me.
    *   Ask me if I want a full scan or I want to focus on a single point and in case which one?
    *   If I want a full scan, present one point at the time and for each point:
      * Ask if I want to **S**kip, **D**iscuss or **Apply**.
        * if Skip, move to the next point.
        * if Discuss, generate a report with the proposed changes and ask me if I want to **A**pply **C**hange or **R**eject.
        * If Apply, start applying the change.

## 2. Development Cycle

For each set of approved plan, follow this cycle:

1.  **Implement Changes**: Make the required code modifications.
2.  **Run Quality Checks**: After completing a set of changes, you **must** run the following commands using the Makefile:
    *   `make test`: Run all tests to ensure no regressions were introduced.
    *   `make lint-fix`: Automatically fix any linting issues.
    *   `make fmt`: Format the code according to the project style.
3.  **Generate Commit**: Once all checks pass, generate a single, cohesive commit with a clear and descriptive message.
