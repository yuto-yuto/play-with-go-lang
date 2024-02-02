package benchmark

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

var covers = []string{
	"AAAAAAAAAAAA",
	"BBBBBBBBBBBB",
	"CCCCCCCCCCCC",
	"DDDDDDDDDDDD",
	"EEEEEEEEEEEE",
	"FFFFFFFFFFFF",
	"GGGGGGGGGGGG",
	"HHHHHHHHHHHH",
	"IIIIIIIIIIII",
	"JJJJJJJJJJJJ",
	"KKKKKKKKKKKK",
	"LLLLLLLLLLLL",
	"MMMMMMMMMMMM",
	"NNNNNNNNNNNN",
	"OOOOOOOOOOOO",
	"PPPPPPPPPPPP",
	"QQQQQQQQQQQQ",
	"RRRRRRRRRRRR",
	"SSSSSSSSSSSS",
	"TTTTTTTTTTTT",
	"UUUUUUUUUUUU",
	"VVVVVVVVVVVV",
	"WWWWWWWWWWWW",
	"XXXXXXXXXXXX",
	"YYYYYYYYYYYY",
	"ZZZZZZZZZZZZ",
}

var targets = []string{
	"AAAAAAAAAAAA.1234567",
	"BBBBBBBBBBBB.1234567",
	"CCCCCCCCCCCC\\1234567",
	"DDDDDDDDDDDD.1234567",
	"EEEEEEEEEEEE.1234567",
	"FFFFFFFFFFFF.1234567",
	"ABDECGRERW3R2.437298",
	"SDFAJKLCXVZ.243798",
	"GGGGGGGGGGGG.1234567",
	"HHHHHHHHHHHH.1234567",
	"IIIIIIIIIIII\\1234567",
	"WWJYEAEQXDRZ.58739",
	"JJJJJJJJJJJJ.1234567",
	"KKKKKKKKKKKK.1234567",
	"LLLLLLLLLLLL.1234567",
	"MMMMMMMMMMMM.1234567",
	"BVWYFUOIERFT.193678",
	"NNNNNNNNNNNN.1234567",
	"OOOOOOOOOOOO.1234567",
	"PPPPPPPPPPPP.1234567",
	"QQQQQQQQQQQQ\\1234567",
	"RRRRRRRRRRRR.1234567",
	"EFQUYGVDSFRE.3524967",
	"SSSSSSSSSSSS.1234567",
	"TTTTTTTTTTTT.1234567",
	"UUUUUUUUUUUU.1234567",
	"VVVVVVVVVVVV.1234567",
	"VDSBHAOVFDSK.0952843",
	"WWWWWWWWWWWW.1234567",
	"XXXXXXXXXXXX.1234567",
	"YYYYYYYYYYYY.1234567",
	"ZZZZZZZZZZZZ.1234567",
}

func funcExample(targetString string) {

}

func BenchmarkRegexLoopMustCompile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, converString := range covers {
			regex := regexp.MustCompile(converString + `[\.\\]`)

			for _, target := range targets {
				if regex.MatchString(target) {
					break
				}
			}
		}
	}
}

func BenchmarkRegexLoopCompile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, converString := range covers {
			regex, err := regexp.Compile(converString + `[\.\\]`)
			if err != nil {
				fmt.Println("failed")

				return
			}

			for _, target := range targets {
				if regex.MatchString(target) {
					break
				}
			}
		}
	}
}

func BenchmarkRegexLoopCompileOnlyWithDot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, converString := range covers {
			regex, err := regexp.Compile(converString + `\.`)
			if err != nil {
				fmt.Println("failed")

				return
			}

			for _, target := range targets {
				if regex.MatchString(target) {
					break
				}
			}
		}
	}
}

func BenchmarkRegexOnce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		regexString := fmt.Sprintf(`(%s)[\.\\]`, strings.Join(covers, "|"))
		regex, err := regexp.Compile(regexString)
		if err != nil {
			fmt.Println("failed")

			return
		}

		for _, target := range targets {
			if regex.MatchString(target) {
				break
			}
		}
	}
}

func BenchmarkRegexOnceOnlyWithDot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		regexString := fmt.Sprintf(`(%s)\.`, strings.Join(covers, "|"))
		regex, err := regexp.Compile(regexString)
		if err != nil {
			fmt.Println("failed")

			return
		}

		for _, target := range targets {
			if regex.MatchString(target) {
				break
			}
		}
	}
}

func BenchmarkRegexContains1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, target := range targets {
			for _, coverString := range covers {
				if strings.Contains(target, coverString) {
					break
				}
			}
		}
	}
}

func BenchmarkRegexContains2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, target := range targets {
			for _, coverString := range covers {
				if strings.Contains(target, fmt.Sprintf("%s.", coverString)) {
					break
				}
			}
		}
	}
}

func BenchmarkRegexContains3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, target := range targets {
			for _, coverString := range covers {
				withDot := fmt.Sprintf("%s.", coverString)
				withBackSlash := fmt.Sprintf("%s\\", coverString)
				if strings.Contains(target, withDot) || strings.Contains(target, withBackSlash) {
					break
				}
			}
		}
	}
}

type Covers struct {
	withDot       string
	withBackSlash string
}

func BenchmarkRegexContains4(b *testing.B) {
	coversStruct := make([]Covers, 0, len(covers))

	for _, cover := range covers {
		coversStruct = append(coversStruct, Covers{
			withDot:       fmt.Sprintf("%s.", cover),
			withBackSlash: fmt.Sprintf("%s\\", cover),
		})
	}

	for i := 0; i < b.N; i++ {
		for _, target := range targets {
			for _, coverStruct := range coversStruct {
				if strings.Contains(target, coverStruct.withDot) || strings.Contains(target, coverStruct.withBackSlash) {
					break
				}
			}
		}
	}
}

// $ go test ./benchmark -bench Regex -benchmem
// goos: linux
// goarch: amd64
// pkg: play-with-go-lang/benchmark
// cpu: Intel(R) Core(TM) i7-9850H CPU @ 2.60GHz
// BenchmarkRegexLoopMustCompile-12                   14544             82590 ns/op           56624 B/op        598 allocs/op
// BenchmarkRegexLoopCompile-12                       14756             79542 ns/op           56656 B/op        598 allocs/op
// BenchmarkRegexLoopCompileOnlyWithDot-12            15729             78042 ns/op           55637 B/op        546 allocs/op
// BenchmarkRegexOnce-12                              34831             35230 ns/op           52094 B/op        139 allocs/op
// BenchmarkRegexOnceOnlyWithDot-12                   34657             33822 ns/op           52078 B/op        138 allocs/op
// BenchmarkRegexContains1-12                        248371              4692 ns/op               0 B/op          0 allocs/op
// BenchmarkRegexContains2-12                         19749             60038 ns/op           17799 B/op       1112 allocs/op
// BenchmarkRegexContains3-12                         10000            111526 ns/op           32461 B/op       2028 allocs/op
// BenchmarkRegexContains4-12                        129026              9165 ns/op               0 B/op          0 allocs/op
// PASS
// ok      play-with-go-lang/benchmark     14.569s
