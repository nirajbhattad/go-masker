# go-masker
A Go library for masking fields of a struct or nested struct based on a specific struct tag using reflection.

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
go get github.com/[username]/masker
```

### Contributing
If you want to contribute to the project, please open a pull request or an issue on GitHub.

### License
The project is open-sourced under the [MIT License](LICENSE). See the LICENSE file for more information.
