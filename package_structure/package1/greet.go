package package1

import (
	"fmt"
	"math/rand"
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

	greetStr := greetInSelectedLanguage(selectedLanguage)
	fmt.Println(greetStr)
}

func greetInSelectedLanguage(locale language.Tag) string {
	now := time.Now()
	if locale == language.English {
		return GreetInEnglish(now)
	}
	if locale == language.Japanese {
		return GreetInJapanese(now)
	}
	if locale == language.German {
		return GreetInGerman(now)
	}
	return "Hello"
}

func GreetInEnglish(now time.Time) string {
	hour := now.Hour()
	if hour >= 6 && hour < 11 {
		return "Good morning"
	} else if hour >= 11 && hour < 16 {
		return "Good afternoon"
	} else if hour >= 16 && hour < 22 {
		return "Good evening"
	}
	return "Good night"
}

func GreetInJapanese(now time.Time) string {
	hour := now.Hour()
	if hour >= 6 && hour < 11 {
		return "おはよう"
	} else if hour >= 11 && hour < 16 {
		return "こんにちは"
	} else if hour >= 16 && hour < 22 {
		return "こんばんは"
	}
	return "おやすみ"
}

func GreetInGerman(now time.Time) string {
	hour := now.Hour()
	if hour >= 6 && hour < 11 {
		return "Guten Morgen"
	} else if hour >= 11 && hour < 16 {
		return "Guten Tag"
	} else if hour >= 16 && hour < 22 {
		return "Guten Abend"
	}
	return "Gute Nacht"
}
