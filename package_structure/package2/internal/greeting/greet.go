package greeting

import "time"

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
