# gopherkit

![Github Actions](https://github.com/ferdiebergado/gopherkit/actions/workflows/go.yml/badge.svg?event=push) [![Go Report Card](https://goreportcard.com/badge/github.com/ferdiebergado/gopherkit)](https://goreportcard.com/report/github.com/ferdiebergado/gopherkit) [![Go Reference](https://pkg.go.dev/badge/github.com/ferdiebergado/gopherkit@v0.0.3.svg)](https://pkg.go.dev/github.com/ferdiebergado/gopherkit@v0.0.3)

GopherKit is a versatile and lightweight utility library designed for Go developers.

## Installation

```sh
go get github.com/ferdiebergado/gopherkit
```

## Documentation

### env

`env` is a package that provides utility functions for managing environment variables. These functions help you load, retrieve, and validate environment variables with ease, ensuring robust and consistent behavior in your application.

### Functions

#### `Load(envFile string) error`

**Description**: Loads environment variables from a specified file.

- **Parameters**:
  - `envFile`: The path to the environment file (e.g., `.env`).
- **Returns**:
  - `error`: An error if the file cannot be read or if an environment variable cannot be set.

**Usage**:

```go
err := env.Load(".env")
if err != nil {
    log.Fatalf("Error loading .env file: %v", err)
}
```

---

#### `MustGet(envVar string) string`

**Description**: Retrieves the value of an environment variable. If the variable is not set, the program panics.

- **Parameters**:
  - `envVar`: The name of the environment variable.
- **Returns**:
  - `string`: The value of the environment variable.

**Usage**:

```go
dbUser := env.MustGet("DB_USER")
```

**Notes**: Use this function for mandatory environment variables to ensure the application cannot start without them.

---

#### `Get(envVar string, fallback string) string`

**Description**: Retrieves the value of an environment variable. If the variable is not set, a fallback value is returned.

- **Parameters**:
  - `envVar`: The name of the environment variable.
  - `fallback`: The fallback value to use if the variable is not set.
- **Returns**:
  - `string`: The value of the environment variable or the fallback.

**Usage**:

```go
dbHost := env.Get("DB_HOST", "localhost")
```

---

#### `GetInt(envVar string, fallback int) int`

**Description**: Retrieves the value of an environment variable as an integer. If the variable is not set or the value is invalid, a fallback value is returned.

- **Parameters**:
  - `envVar`: The name of the environment variable.
  - `fallback`: The fallback value to use if the variable is not set or invalid.
- **Returns**:
  - `int`: The value of the environment variable or the fallback.

**Usage**:

```go
pingTimeout := env.GetInt("DB_PING_TIMEOUT", 10)
```

---

#### `GetBool(envVar string, fallback bool) bool`

**Description**: Retrieves the value of an environment variable as a boolean. If the variable is not set or the value is invalid, a fallback value is returned.

- **Parameters**:
  - `envVar`: The name of the environment variable.
  - `fallback`: The fallback value to use if the variable is not set or invalid.
- **Returns**:
  - `bool`: The value of the environment variable or the fallback.

**Usage**:

```go
isDebug := env.GetBool("DEBUG", false)
```

---

### Example Usage

```go
package main

import (
	"log"

    "github.com/ferdiebergado/gopherkit/env"
)

func main() {
  // Logger initialization with programLevel here

	// Load environment variables from a file
	if err := env.Load(".env"); err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}

	// Get mandatory environment variable
	dbUser := env.MustGet("DB_USER")

	// Get environment variable with fallback
	dbHost := env.Get("DB_HOST", "localhost")

	// Get integer environment variable with fallback
	port := env.GetInt("PORT", 8080)

  isDebug := env.GetBool("DEBUG", false)

  if isDebug {
    programLevel.Set(slog.LevelDebug)
  }

	log.Printf("Database User: %s, Host: %s, Port: %d", dbUser, dbHost, port)
}
```

---

### assert

`assert` is a package that provides utility functions for simplifying common test assertions. These functions help to assert conditions in tests, such as equality, errors, and string containment, with customizable messages for better test failure reports.

### Functions

#### `Equal(t *testing.T, expected, actual any)`

**Description**: Asserts that two values are equal. If the values are not equal, an error is reported with the provided message.

- **Parameters**:
  - `t`: The test context (from `testing.T`).
  - `expected`: The expected value.
  - `actual`: The actual value to compare against.

**Usage**:

```go
assert.Equal(t, expectedValue, actualValue)
```

---

#### `NotEqual(t *testing.T, expected, actual any)`

**Description**: Asserts that two values are not equal. If the values are equal, an error is reported with the provided message.

- **Parameters**:
  - `t`: The test context (from testing.T).
  - `expected`: The value that should not be equal to actual.
  - `actual`: The actual value to compare.

**Usage**:

```go
assert.NotEqual(t, expectedValue, actualValue)
```

---

#### `NoError(t *testing.T, err error)`

**Description**: Asserts that an error is nil. If the error is not nil, an error is reported with the provided message.

- **Parameters**:
  - `t`: The test context (from testing.T).
  - `err`: The error to check.

**Usage**:

```go
assert.NoError(t, err)
```

---

#### `Error(t *testing.T, err error)`

**Description**: Asserts that an error is not nil. If the error is nil, an error is reported with the provided message.

- **Parameters**:
  - `t`: The test context (from testing.T).
  - `err`: The error to check.

**Usage**:

```go
assert.Error(t, err)
```

---

#### `Contains(t *testing.T, s, substr string)`

**Description**: Asserts that a string contains a substring. If the substring is not found, an error is reported with the provided message.

- `Parameters`:
  - `t`: The test context (from testing.T).
  - `s`: The string to check.
  - `substr`: The substring to check for.

**Usage**:

```go
assert.Contains(t, str, "substring")
```

---

#### `Len(t *testing.T, collection any, length int)`

**Description**: Asserts that a collection (e.g., slice, array, map) has the expected length. If the lengths do not match, an error is reported with the provided message.

- **Parameters**:
  - `t`: The test context (from testing.T).
  - `collection`: The collection whose length is to be checked.
  - `length`: The expected length.

**Usage**:

```go
assert.Len(t, collection, expectedLength)
```

---

## Example Usage

```go
package main

import (
	"testing"
	"github.com/ferdiebergado/gopherkit/assert"
)

func TestExample(t *testing.T) {
	// Assert equal values
	assert.Equal(t, 5, 5)

	// Assert not equal values
	assert.NotEqual(t, 5, 10)

	// Assert no error
	assert.NoError(t, nil)

	// Assert error
	assert.Error(t, fmt.Errorf("some error"))

	// Assert string contains substring
	assert.Contains(t, "Hello World", "World")

	// Assert collection length
	assert.Len(t, []int{1, 2, 3}, 3)
}
```

### Miscellaneous Helpers

#### `Sum[T Number](values ...any) Number`

**Description**: Calculates the sum of the given numbers or slice of numbers.

- **Parameters**:
  - `values`: Variadic or slice of numbers.
- **Returns**:
  - `Number`: The sum of the given numbers.

**Usage**:

```go
totalItems := gopherkit.Sum[int](10675, 8050, 2503)
totalRate := gopherkit.Sum[float64]([]float64{1.5, 0.3, 0.1})
```

---

#### `GetIPAddress(r *http.Request) string`

**Description**: Extracts the client's IP address from the request.

- **Parameters**:
  - `r`: The http request.
- **Returns**:
  - `String`: The client ip address.

**Usage**:

```go
ip := GetIPAddress(r)
```

---

### Logging

The package uses [slog](https://pkg.go.dev/log/slog) for logging. Make sure to set up your logging configuration to capture relevant logs for debugging or monitoring purposes.

## License

This package is distributed under the MIT License. See [LICENSE](https://github.com/ferdiebergado/gopherkit/blob/main/LICENSE) for more details.
