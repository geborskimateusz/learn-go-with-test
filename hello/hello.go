package hello

const pl = "PL"
const fr = "FR"
const greetingFrenchPrefix = "Bonjour "
const greetingEnglishPrefix = "Hello "
const greetingPolishPrefix = "Witaj "

// Hello greets user
func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefixCreator(language) + name
}

func greetingPrefixCreator(language string) (greetingPrefix string) {

	switch language {
	case pl:
		greetingPrefix = greetingPolishPrefix
	case fr:
		greetingPrefix = greetingFrenchPrefix
	default:
		greetingPrefix = greetingEnglishPrefix
	}
	return
}
