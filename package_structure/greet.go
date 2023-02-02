package packagestructure

import (
	"fmt"
	"play-with-go-lang/package_structure/package1"
	"play-with-go-lang/package_structure/package2"
	"play-with-go-lang/package_structure/package2/selector"

	"golang.org/x/text/language"
	// could not import play-with-go-lang/package_structure/package2/internal/greeting (invalid use of internal package play-with-go-lang/package_structure/package2/internal/greeting)
	// "play-with-go-lang/package_structure/package2/internal/greeting"
)

func PrintGreeting() {
	fmt.Println("--- package1 ---")
	package1.Greet()
	fmt.Println("--- package2 ---")
	package2.Greet()
	fmt.Println("--- selector ---")
	greetingString := selector.GreetInSelectedLanguage(language.Georgian)
	fmt.Println(greetingString)
}
