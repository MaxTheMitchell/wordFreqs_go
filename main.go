package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	printtop(sortfreqs(countfreqs(removestopwords(parsetext(readfile(os.Args[1]))))))
}

func readfile(filename string) (textstream string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, _ := ioutil.ReadAll(file)
	textstream = string(b)
	return
}
func parsetext(text string) (words []string) {
	re := regexp.MustCompile(`\w{3,}`)
	words = re.FindAllString(strings.ToLower(text), -1)
	return
}

func removestopwords(words []string) (fixedwords []string) {
	stopwords := strings.Split(readfile("stop_words.txt"), ",")
	for _, word := range words {
		notstopword := true
		for _, stopword := range stopwords {
			if word == stopword {
				notstopword = false
			}
		}
		if notstopword {
			fixedwords = append(fixedwords, word)
		}
	}
	return
}
func countfreqs(words []string) (wordfreqs map[string]int) {
	wordfreqs = make(map[string]int)
	for _, word := range words {
		notfoundword := true
		for mapword := range wordfreqs {
			if word == mapword {
				wordfreqs[mapword]++
				notfoundword = false
				break
			}
		}
		if notfoundword {
			wordfreqs[word] = 1
		}
	}
	return
}

func sortfreqs(wordfreqs map[string]int) (sortedfreqs map[int]string, topfreqs []int) {
	sortedfreqs = make(map[int]string)
	for word, freq := range wordfreqs {
		sortedfreqs[freq] = word
		topfreqs = append(topfreqs, freq)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(topfreqs)))
	return
}
func printtop(sortedfreqs map[int]string, topfreqs []int) {
	for i, freq := range topfreqs {
		fmt.Println(sortedfreqs[freq] + " " + strconv.Itoa(freq))
		if i >= 25 {
			break
		}
	}
}
