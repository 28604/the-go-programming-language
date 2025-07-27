// The run command compiles the source code from many .go files,
// links the files with libraries, and runs the executable file.
// $ go run main.go

// The build command compiles the source code from many .go files,
// links the files with libraries, and compiles it to an executable file.
// $ go build main.go
// $ ./main

package main

// fmt package contains scanning input and printing formatted output.
import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
