
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "regexp"
)

func maskPII(text string) string {
    // Your code here
    return text
}

func main() {
    // Compiler Boilerplate
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    text = strings.TrimSpace(text)
    result := maskPII(text)
    fmt.Println(result)
}
