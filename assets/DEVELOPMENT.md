# Development Guidelines

> This project is actively developed with the assistance of Google Jules, an AI-powered software engineer.

These principles are language-agnostic and form the core philosophy of writing clean, maintainable code.

This document serves as a table of contents for our development guidelines. Please read these documents to understand the coding standards and best practices for this project.

### How the Rules Are Organized

The development guidelines are split into several files, each with a specific purpose. They are designed to be read in order and build upon one another.

1.  **[General Principles](./development/01-general-principles.md)**: These are the foundational, language-agnostic software engineering principles (e.g., Clean Code) that apply to all code in this repository.
2.  **[Go-Specific Guidelines](./development/02-go-guidelines.md)**: This document provides specific guidelines and best practices for writing idiomatic Go code. These rules build on or specialize the general principles for the Go ecosystem.
3.  **[SonarQube Rules](./development/03-sonar-qube-rules.md)**: This lists specific, non-negotiable rules that are actively enforced by our SonarQube static analysis pipeline.
4.  **[Repository-Specific Guidelines](./development/04-repository-specific.md)**: This contains rules, conventions, or exceptions unique to this project.

In case of conflict, the rules have a clear order of precedence: **Repository-Specific > Go-Specific > General Principles**.
