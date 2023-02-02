package package2

import (
	"fmt"
	"math/rand"
	"play-with-go-lang/package_structure/package2/selector"
	"time"

	"golang.org/x/text/language"
)

var supportedLanguages = []language.Tag{
	language.English,
	language.German,
	language.Japanese,
}

func Greet() {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(3)

	selectedLanguage := supportedLanguages[random]
	fmt.Printf("Selected Language: %s\n", selectedLanguage)

	greetStr := selector.GreetInSelectedLanguage(selectedLanguage)
	fmt.Println(greetStr)
}
