package selector

import (
	"golang.org/x/text/language"
	"play-with-go-lang/package_structure/package2/internal/greeting"
	"time"
)

func GreetInSelectedLanguage(locale language.Tag) string {
	now := time.Now()
	if locale == language.English {
		return greeting.GreetInEnglish(now)
	}
	if locale == language.Japanese {
		return greeting.GreetInJapanese(now)
	}
	if locale == language.German {
		return greeting.GreetInGerman(now)
	}
	return "Hello"
}
