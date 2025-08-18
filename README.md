# Zap Fluent Encoder

This library provides a fluent interface encoder for Uber's Zap logging library. Instead of making multiple calls to add fields to a log entry (e.g., encoder.AddString, encoder.AddInt), you can use a chained approach. This can improve readability, especially for logs with many fields.

## ⚠️ Work in Progress

This project is currently under active development. APIs and functionality may change without notice.

## Development

This library is being developed with the assistance of Jules by Google, an AI-powered software engineer. The following rules are being followed during development:

- **No Unnecessary Comments**: The code must be self-documenting. Comments are to be avoided unless they are 100% necessary. This includes not using comments to delineate the sections of a test.
- **AAA Unit Test Structure**: All unit tests must be implicitly structured following the Arrange-Act-Assert (AAA) pattern.
- **Complete Test Coverage**: All production code must be covered by at least one unit test.
- **Standard Import Formatting**: All import blocks must be grouped into four categories in a specific order: standard library, third-party, shared internal modules, and intra-module dependencies.
