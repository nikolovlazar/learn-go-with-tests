package helloworld

import "fmt"

const (
	englishHelloPrefix    = "Hello, "
	spanish               = "Spanish"
	spanishHelloPrefix    = "Hola, "
	french                = "French"
	frenchHelloPrefix     = "Bonjour, "
	macedonian            = "Macedonian"
	macedonianHelloPrefix = "Здраво, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case macedonian:
		prefix = macedonianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", "English"))
}
