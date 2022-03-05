package hello

import "fmt"

func Hello(name string) string {
	// return "Hello " + name + "!"
	return fmt.Sprintf("Hello %s!", name)
}
