# Zap Fluent Encoder

A fluent interface encoder for Uber's Zap logging library that provides a more intuitive way to add structured logging
fields.

## ⚠️ Work in Progress

This project is currently under active development. APIs and functionality may change without notice.

## Development

The rules presented here are a framework for software craftsmanship, informed by the principles of Clean Code. They are not meant to be followed dogmatically, but rather to guide us in a continuous effort to improve the quality of our codebase.

Our responsibility is to leave the code cleaner than we found it—to fix "code smells" and apply these principles incrementally with every change. The ultimate measure of success is code that is readable and tells a clear story, even if it means making a pragmatic exception to a specific rule.

### **General Principles**

These principles are language-agnostic and form the core philosophy of writing clean, maintainable code.

#### **Naming**

* **Use Intention-Revealing Names:** Names should answer the big questions: why it exists, what it does, and how it's used. Avoid names that need a comment to explain their purpose.
* **Avoid Disinformation:** Don't use names that mislead. For example, `hp` shouldn't be used for a variable representing `horsepower` if it actually represents something else, like `health points`.
* **Make Meaningful Distinctions:** When two things are different, their names should be different. Avoid using arbitrary number series (`a1`, `a2`, ...) or noise words (`the`, `data`, `info`, `variable`) that don't add meaning.
* **Use Pronounceable Names:** Names should be easy to say and remember.
* **Use Searchable Names:** Use names that are easy to find in a codebase. A single letter like `i` is a poor choice for anything other than a loop index.
* **Class and Object Names:** Use nouns or noun phrases (**Customer**, **Account**, **WikiPage**). Avoid verbs.
* **Method Names:** Use verbs or verb phrases (**save**, **deletePage**, **postComment**).
* **One Word Per Concept:** Choose one word for a concept and stick with it. For example, don't mix `fetch`, `retrieve`, and `get` for the same kind of action.
* **Add Context to Names:** A variable named `state` is only meaningful if it has a context, like `addrState`.
* **Don't Add a Prefix to Interface Names:** It's generally not necessary to prefix interfaces with an `I` or similar. The name of the interface should be sufficient.

#### **Functions**

* **Functions Should Be Small:** Functions should be very short, typically no more than a few lines.
* **Do One Thing:** A function should do one thing, and do it well. If a function can be described by more than one verb, it's likely doing too much.
* **Function Parameters:** The ideal number of arguments for a function is zero (niladic). Next comes one (monadic), followed closely by two (dyadic). Three arguments (triadic) should be avoided where possible. More than three requires very special justification — and then shouldn't be used anyway.
* **Function Return Values:** Functions should either be commands that perform an action and do not return a value, or queries that return a single, meaningful value without performing any side effects. Avoid using output arguments.
* **Use Descriptive Names:** Prefer self-documenting code over comments. The name of the function should clearly describe what it does.
* **Avoid Side Effects:** A function shouldn't do something that isn't explicitly stated in its name. For example, a function named `isPasswordValid()` shouldn't also initialize a session.
* **Prefer Command-Query Separation:** Functions should either do something or answer something, but not both. For example, a `set()` function shouldn't also return a status code.
* **Don't Repeat Yourself (DRY):** Avoid duplicating code. If you find yourself writing the same logic in multiple places, extract it into a function.

#### **Comments**

* **Prefer Self-Documenting Code:** The best code is its own documentation. Comments should be a last resort.
* **Explain "Why," Not "What":** Use comments to explain **why** a particular decision was made, not to restate what the code is doing.
* **Do Not Comment Bad Code:** If code is unclear, refactor it. Don't add a comment to explain it.
* **Avoid Redundant Comments:** Don't comment on something that's already obvious from the code.
* **Do Not Use Commented-Out Code:** Delete commented-out code. Use version control to retrieve old code if needed.

#### **Class and Object Design**

* **Classes Should Be Small:** A class should have a single, well-defined responsibility. As a general rule, a class with more than a few methods or instance variables is likely doing too much.
* **Encapsulation of Data:** Classes should hide their internal data structures and expose behavior through well-defined methods. Avoid creating classes that are just collections of public variables with no meaningful functions (known as Data Transfer Objects, which are a different concept).
* **Tell, Don't Ask:** Prefer telling objects what to do rather than asking them for their internal data and then acting on that data. This helps preserve encapsulation. For example, instead of `if (car.getSpeed() > speedLimit) { car.slowDown(); }`, prefer `car.checkAndRegulateSpeed(speedLimit);`.
* **Objects vs. Data Structures:** A core distinction exists between an object and a data structure. An **object** hides its data and exposes functions. A **data structure** exposes its data and has no meaningful functions. Clean code understands this difference and chooses the right approach for the problem at hand.
* **The Law of Demeter:** A class should only talk to its immediate friends. Avoid long chains of method calls like `a.getB().getC().doSomething()`. This indicates that your classes are overly coupled and lack proper encapsulation.

#### **System Boundaries**

* **Wrap Third-Party Code:** Isolate third-party APIs and libraries behind an interface or a wrapper class that you control. This prevents changes in an external library from propagating throughout your codebase.

---

### **Go-Specific Guidelines**

These guidelines are tailored to the idioms and conventions of the Go programming language.

#### **Functions**

* **Max 2 Parameters:** Functions should have at most two parameters.
    * **The Context Corner Case:** A `context.Context` is often a necessary first parameter, but it does not count against the parameter limit, as it is a common and predictable part of the function signature.
* **No Unnecessary Pointers:** Avoid using pointers for optional values or to pass arguments. Pointers should only be used when strictly necessary (e.g., for channels, or for large objects where copying is a performance concern). Prefer a value-based approach or an `Optional` monad for optional values.
* **The Error Return Corner Case:** A function may return two values (e.g., `(result, error)`). This is a common and valid pattern, as the second value (the error) is conceptually part of a single result: a successful value or a failed state.
* **Don't Return `nil`:** Avoid returning `nil` from functions that could fail. Prefer returning a zero value or an empty slice/map.

#### **Formatting and Organization**

* **Code Locality and Cohesion:** Group related code together. Avoid mixing unrelated concepts in the same file or package. For example, a generic `constants.go` file is discouraged as it tends to aggregate values that have little in common besides being constants.
* **Standard Import Formatting:** All import blocks must be grouped into four categories in a specific order: standard library, third-party, shared internal modules, and intra-module dependencies.

#### **Tests**

* **AAA Unit Test Structure:** All unit tests must be implicitly structured following the **Arrange-Act-Assert** (AAA) pattern.
* **Complete Test Coverage:** All production code must be covered by at least one unit test.
* **Test One Concept:** A test should verify a single, specific behavior or concept. This makes the test's purpose immediately clear and aids in diagnosing failures.
    * **One Assert Guideline:** While a test should focus on one concept, it's a good practice to use only one assertion to enforce this. However, multiple assertions may be acceptable if they all verify different aspects of that *single* concept (e.g., asserting different fields on a returned object). The primary goal is diagnostic clarity.
    * **Custom Assertions:** Consider creating custom assertion functions if they are reusable across the project, simple to define (e.g., a low number of parameters), and have a very clear, unambiguous name and purpose. These can help enforce the "one concept" rule while simplifying the test code.
* **First-Class Tests:** Tests are as important as production code. They should be clean, readable, and well-organized.
* **No Magic Values:** Avoid magic strings and numbers. If the same value is used more than once, introduce a constant for it.

#### **Principles and General Rules**

* **SOLID Principles:** Code should adhere to **SOLID principles**, with a strong emphasis on the **Single Responsibility Principle**.
* **Composition over Inheritance:** Prefer composition to achieve code reuse over inheritance.
* **Dependency Injection:** Use dependency injection to decouple components and make them more testable.
