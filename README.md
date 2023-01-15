# go-masker
A Go library for masking fields of a struct or nested struct based on a specific struct tag using reflection.

With Masker, you can easily and securely mask sensitive data in your structs, such as credit card numbers and personal information. Using struct tags, you can specify which fields should be masked.

Features:

- Mask fields of a struct or nested struct
- Use struct tags to specify which fields to mask and how to mask them
- Mask data using various masking methods for various data types such as string, int, float. 
- Use reflection to mask fields dynamically

# Description
This library is a simple utility for masking sensitive data (for example, personal information) in structs in Go. it uses reflection to find the fields with a specific struct tag, such as "mask", and it will mask its value if the struct tag value is true. It also provides the option to specify a custom masking character.

This library is useful for masking sensitive data in logs, APIs or for debugging purposes.

To use this library, import it into your project, and call the Mask function passing in a pointer to the struct you wish to mask, as well as the struct tag you're using for masking.

### Usage

To use the library, you first need to import it in your code:
```go
import "github.com/nirajbhattad/masker"
```

### Installation
To install the library, use go get:
```go
go get github.com/nirajbhattad/masker
```

### Contributing
If you want to contribute to the project, please open a pull request or an issue on GitHub.

### License
The project is open-sourced under the [MIT License](LICENSE). See the LICENSE file for more information.
