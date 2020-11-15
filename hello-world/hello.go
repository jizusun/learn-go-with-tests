package helloworld

const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

// Hello greet with different languauges
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := englishHelloPrefix

	if language == spanish {
		prefix = spanishHelloPrefix
	}

	return prefix + name
}
