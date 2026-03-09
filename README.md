# Test-Driven Development with Bazel and Go

This project demonstrates a simple Phone Book application built using **Bazel**, **Gazelle**, and **Go**, following Test-Driven Development (TDD) principles.

## Project Structure

- `phonebook/`: Contains the core logic and tests for the phone book application.
- `BUILD.bazel`: Root build file with Gazelle configuration.
- `MODULE.bazel`: Modern Bazel dependency management (Bzlmod).
- `.bazelrc`: Project-specific Bazel configuration for macOS.

## Getting Started

### Prerequisites

- [Bazel](https://bazel.build/install) (latest version recommended)
- [Go](https://go.dev/doc/install)

### Initializing the Project

If you are starting from scratch:

1. Initialize Go module:
   ```bash
   go mod init github.com/johanastborg/phonebook
   ```

2. Run Gazelle to generate Build files:
   ```bash
   bazel run //:gazelle
   ```

## Test-Driven Development (TDD) Workflow

TDD is a software development process relying on the repetition of a very short development cycle:
1. **Red**: Write a test for a new function and watch it fail.
2. **Green**: Write the minimum amount of code to make the test pass.
3. **Refactor**: Clean up the code while ensuring tests still pass.

### TDD with Bazel

Bazel is exceptionally well-suited for TDD due to its **caching** and **incremental builds**. Only the tests affected by your changes will be re-run, making the feedback loop extremely fast.

#### 1. Running All Tests
To run all tests in the project:
```bash
bazel test //...
```

#### 2. Running Specific Tests
To run tests for a specific package:
```bash
bazel test //phonebook:phonebook_test
```

#### 3. Filtering Tests (The "Red" Phase)
When you are focusing on a single test case, use the `--test_filter` flag:
```bash
bazel test //phonebook:phonebook_test --test_filter=TestFindContact/found
```

#### 4. Seeing Test Output
By default, Bazel hides the output of passing tests. To see logs (useful for debugging):
```bash
bazel test //phonebook:phonebook_test --test_output=all
```

### The Phone Book Example

In this project, we implemented `TestFindContact`. The TDD cycle looked like this:

1. **Red**: We defined `TestFindContact` in `phonebook/phonebook_test.go` before implementing the lookup logic. Running `bazel test` failed.
2. **Green**: We implemented the `FindContact` method in `phonebook/phonebook.go`.
3. **Bazel Sync**: We ran `bazel run //:gazelle` to update the build targets.
4. **Verification**: Running `bazel test //phonebook:phonebook_test` resulted in `PASSED`.

## Useful Bazel Commands

| Command | Description |
|---------|-------------|
| `bazel test //...` | Run all tests in the workspace |
| `bazel run //:gazelle` | Update BUILD files based on Go source code |
| `bazel test --test_output=errors //...` | Show output only for failed tests |
| `bazel test --test_arg=-test.v //...` | Pass verbose flag to Go test runner |
| `bazel clean` | Remove all build artifacts |

## macOS Compatibility

This project uses a `.bazelrc` file optimized for macOS:
- `CGO_ENABLED=0`: Disables CGO for easier cross-compilation and faster builds.
- `--@rules_go//go/config:pure`: Ensures the pure Go compiler is used.

---
Happy Coding with Bazel!
