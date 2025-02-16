package hello

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"
const norwegian = "Norwegian"
const helloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const norwegianHelloPrefix = "Hei, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case norwegian:
		prefix = norwegianHelloPrefix
	default:
		prefix = helloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Chriss", ""))
}
