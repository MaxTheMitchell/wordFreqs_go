package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	printtop(sortfreqs(countfreqs(removestopwords(parsetext(readfile(os.Args[1]))))))
}

func readfile(filename string) (textstream string) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	textstream = string(b)
	return
}
func parsetext(text string) (words []string) {
	words = regexp.MustCompile(`(\w|'){3,}`).FindAllString(strings.ToLower(text), -1)
	return
}

func removestopwords(words []string) (fixedwords []string) {
	stopwords := strings.Split(readfile("stop_words.txt"), ",")
	for _, word := range words {
		notstopword := true
		for _, stopword := range stopwords {
			if word == stopword {
				notstopword = false
				break
			}
		}
		if notstopword {
			fixedwords = append(fixedwords, word)
		}
	}
	return
}
func countfreqs(words []string) (newwords []string, freqs []int) {
	for _, word := range words {
		notfoundword := true
		for i, newword := range newwords {
			if word == newword {
				freqs[i]++
				notfoundword = false
				break
			}
		}
		if notfoundword {
			newwords = append(newwords, word)
			freqs = append(freqs, 1)
		}
	}
	return
}

func sortfreqs(words []string, freqs []int) (sortedwords []string, sortedfreqs []int) {
	for len(freqs) != 0 {
		largest, index := freqs[0], 0
		for i, freq := range freqs {
			if freq > largest {
				largest, index = freq, i
			}
		}
		sortedwords = append(sortedwords, words[index])
		sortedfreqs = append(sortedfreqs, freqs[index])
		freqs = freqs[:index+copy(freqs[index:], freqs[index+1:])]
		words = words[:index+copy(words[index:], words[index+1:])]
	}
	return
}
func printtop(words []string, freqs []int) {
	for i := 0; i < 25; i++ {
		if i >= len(freqs) {
			fmt.Println("These are all the words on the file")
			break
		}
		fmt.Println(words[i] + " " + strconv.Itoa(freqs[i]))
	}
}
