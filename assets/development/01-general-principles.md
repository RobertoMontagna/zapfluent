# General Principles

## Naming

* **Use Intention-Revealing Names**
    * Names should answer the big questions: why it exists, what it does, and how it's used. Avoid names that need a comment to explain their purpose.
* **Avoid Disinformation**
    * Don't use names that mislead. For example, `hp` shouldn't be used for a variable representing `horsepower` if it actually represents something else, like `health points`.
* **Make Meaningful Distinctions**
    * When two things are different, their names should be different. Avoid using arbitrary number series (`a1`, `a2`, ...) or noise words (`the`, `data`, `info`, `variable`) that don't add meaning.
* **Use Pronounceable Names**
    * Names should be easy to say and remember.
* **Use Searchable Names**
    * Use names that are easy to find in a codebase. A single letter like `i` is a poor choice for anything other than a loop index.
* **Class and Object Names**
    * Use nouns or noun phrases (**Customer**, **Account**, **WikiPage**). Avoid verbs.
* **Method Names**
    * Use verbs or verb phrases (**save**, **deletePage**, **postComment**).
* **One Word Per Concept**
    * Choose one word for a concept and stick with it. For example, don't mix `fetch`, `retrieve`, and `get` for the same kind of action.
* **Add Context to Names**
    * A variable named `state` is only meaningful if it has a context, like `addrState`.

## Functions

* **Functions Should Be Small**
    * Functions should be very short, typically no more than a few lines.
* **Do One Thing**
    * A function should do one thing, and do it well. If a function can be described by more than one verb, it's likely doing too much.
* **Function Parameters**
    * The ideal number of arguments for a function is zero (niladic). Next comes one (monadic), followed closely by two (dyadic). Three arguments (triadic) should be avoided where possible. More than three requires very special justification. For functions requiring many parameters, consider bundling them into a single parameter object.
* **Function Return Values**
    * Functions should either be commands that perform an action and do not return a value, or queries that return a single, meaningful value without performing any side effects. Avoid using output arguments (passing pointers or references to be modified by the function).
* **Use Descriptive Names**
    * Prefer self-documenting code over comments. The name of the function should clearly describe what it does.
* **Avoid Side Effects**
    * A function shouldn't do something that isn't explicitly stated in its name. For example, a function named `isPasswordValid()` shouldn't also initialize a session.
* **Prefer Command-Query Separation**
    * Functions should either do something or answer something, but not both. For example, a `set()` function shouldn't also return a status code; its success should be implied unless an error is reported.
* **Don't Repeat Yourself (DRY)**
    * Avoid duplicating code. If you find yourself writing the same logic in multiple places, extract it into a function.
* **Avoid Temporary Variables for Arguments**
    * If a variable is created only to be immediately passed to a function, and its creation is simple, inline it directly into the function call. This reduces vertical space and keeps the logic concise.
    * **Exception**: A variable is acceptable if its creation is complex or if assigning it a descriptive name significantly improves the readability of the code.

## Comments

* **Prefer Self-Documenting Code**
    * The best code is its own documentation. Comments should be a last resort.
* **Explain "Why," Not "What"**
    * Use comments to explain **why** a particular decision was made, not to restate what the code is doing.
* **Do Not Comment Bad Code**
    * If code is unclear, refactor it. Don't add a comment to explain it.
* **Avoid Redundant Comments**
    * Don't comment on something that's already obvious from the code.
* **Do Not Use Commented-Out Code**
    * Delete commented-out code. Use version control to retrieve old code if needed.

## Object and Data Structure Design

* **Classes Should Be Small**
    * A class should have a single, well-defined responsibility. As a general rule, a class with more than a few methods or instance variables is likely doing too much.
* **Encapsulation of Data**
    * Classes should hide their internal data structures and expose behavior through well-defined methods. Avoid creating classes that are just collections of public variables with no meaningful functions (known as Data Transfer Objects, which are a different concept).
* **Tell, Don't Ask**
    * Prefer telling objects what to do rather than asking them for their internal data and then acting on that data. This helps preserve encapsulation. For example, instead of `if (car.getSpeed() > speedLimit) { car.slowDown(); }`, prefer `car.checkAndRegulateSpeed(speedLimit);`.
* **Objects vs. Data Structures**
    * A core distinction exists between an object and a data structure. An **object** hides its data and exposes functions. A **data structure** exposes its data and has no meaningful functions. Clean code understands this difference and chooses the right approach for the problem at hand.
* **The Law of Demeter**
    * A class should only talk to its immediate friends. Avoid long chains of method calls like `a.getB().getC().doSomething()`. This indicates that your classes are overly coupled and lack proper encapsulation.
* **SOLID Principles**
    * Code should adhere to **SOLID principles**, with a strong emphasis on the **Single Responsibility Principle**.
* **Composition over Inheritance**
    * Prefer composition to achieve code reuse over inheritance.
* **Dependency Injection**
    * Use dependency injection to decouple components and make them more testable.

## Testing

*   **Use Semantic Assertions**: Strive to use the most semantically expressive assertion or matcher available for any given check. This makes the test's intent clearer at a glance and often leads to more robust tests. For example, when checking an error, `Expect(err).To(MatchError("..."))` is semantically richer than `Expect(err).To(Equal(someErrVar))`.
* **AAA Unit Test Structure**
    * All unit tests must be structured following the **Arrange-Act-Assert** (AAA) pattern.
    * **Visual Separation**: The three sections of the test (Arrange, Act, and Assert) should be visually separated by a blank line. This is the preferred method for making the structure of the test immediately obvious to the reader.
    * **Use of Comments**: In most cases, the blank line is sufficient. However, for more complex tests where the separation might not be clear enough, it is acceptable to use comments (`// Arrange`, `// Act`, `// Assert`) to explicitly label the sections. They are not forbidden, but should be used judiciously when they add significant clarity.
* **Complete Test Coverage**
    * All production code must be covered by at least one unit test.
* **First-Class Tests**
    * Tests are as important as production code. They should be clean, readable, and well-organized.
* **Test One Concept**
    * A test should verify a single, specific behavior or concept. This makes the test's purpose immediately clear and aids in diagnosing failures.
    * **One Assert Guideline**
        * While a test should focus on one concept, it's a good practice to use only one assertion to enforce this. However, multiple assertions may be acceptable if they all verify different aspects of that *single* concept (e.g., asserting different fields on a returned object). The primary goal is diagnostic clarity.
    * **Custom Assertions**
        * Consider creating custom assertion functions if they are reusable across the project, simple to define (e.g., a low number of parameters), and have a very clear, unambiguous name and purpose. These can help enforce the "one concept" rule while simplifying the test code.
* **No Magic Values**
    * Avoid magic strings and numbers. If the same value is used more than once, introduce a constant for it.
* **Using Test Doubles**
    * Use test doubles (Fakes, Stubs, Mocks) to isolate the code under test and break external dependencies. Prefer using Fakes and Stubs to manage test state, as they are often more resilient to refactoring. Use Mocks judiciously, only when the primary goal is to verify a critical interaction.
* **Test Data Locality**
    * Place test data directly inside the test file itself. This improves readability and maintainability by keeping the input and expected output visible alongside the test logic. As a rule of thumb, it is fine to embed small payloads (e.g., JSON/XML under 50 lines or simple strings) directly. Consider externalizing for:
        * **Large Payloads**
            * JSON or XML files over 50 lines, or multi-line strings that exceed the screen height.
        * **Binary Data**
            * Images, videos, or other non-text files.
        * **Extensive Scripts**
            * SQL scripts or other large text dumps.
        * **External Test Data**
            * When externalizing, store the data in a `testdata/` directory within the same package.
* **Testing Strategy**
    * Black-box testing is preferred except in very extreme corner cases. Tests should be written in a separate `_test` package and should only interact with the public API of the package under test.

## System Boundaries

* **Validate Inputs and Boundaries**
    * All data entering the system from an external source (e.g., user input, API requests, configuration files) must be validated at the system boundary before it is used. This prevents bugs and security vulnerabilities.
* **Wrap Third-Party Code**
    * Isolate third-party APIs and libraries behind an interface or a wrapper class that you control. This prevents changes in an external library from propagating throughout your codebase.
