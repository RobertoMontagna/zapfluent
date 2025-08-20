# Repository-Specific Guidelines

This document contains guidelines and conventions that are specific to this repository.

## Imports

### Dot Imports

Dot imports (`import . "pacakge"`) are strongly discouraged. They can make code confusing by obscuring which package a function or type belongs to, breaking the usual `package.Function` pattern that is standard in Go. Always prefer using a package alias if the package name is too long or conflicts with another name.

An exception is made for certain testing libraries like Gomega, where dot imports are idiomatic and significantly improve test readability. This practice should be used with caution and confined to `_test.go` files.
