package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

var keyword = flag.String("k", "gopher", "The twitter search key.")

func configure() (TwitterConfiguration, error) {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	configuration := TwitterConfiguration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		return configuration, err
	}
	return configuration, nil
}

func TweetSplitter(item interface{}, c chan interface{}) {
	c <- strings.Split(item.(string), " ")
}

func TermFrequency(in chan interface{}, out chan interface{}) {
	tweets := 0
	TweetWordCounter := map[string]int{}

	for tweet := range in {
		t := tweet.([]string)
		tweets++
		for _, word := range t {

			TweetWordCounter[word]++
		}
	}
	out <- TweetWordCounter
}

func main() {
	flag.Parse()

	conf, err := configure()
	if err != nil {
		panic("Error configuration: Create your 'conf.json' file")
	}

	twitter := newTwitter(conf)
	TermFrequency := MapReduce(TweetSplitter, TermFrequency, twitter.find(*keyword, 15), 15).(map[string]int)

	for word, counter := range TermFrequency {
		fmt.Println(fmt.Sprintf("[  %v  ] \t of (%v)", counter, word))
	}

	fmt.Println(*keyword)
}
