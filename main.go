package main

import (
	"encoding/json"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"os"
	"strings"
)

type Configuration struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func configure() (Configuration, error) {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		return configuration, err
	}
	return configuration, nil
}

func main() {
	conf, err := configure()
	if err != nil {
		panic("Error configuration")
	}

	anaconda.SetConsumerKey(conf.ConsumerKey)
	anaconda.SetConsumerSecret(conf.ConsumerSecret)
	api := anaconda.NewTwitterApi(conf.AccessToken, conf.AccessTokenSecret)

	mapper := func(item interface{}, c chan interface{}) {
		c <- strings.Split(item.(string), " ")
	}

	reducer := func(in chan interface{}, out chan interface{}) {
		reduced := map[string]int{}

		for tweets := range in {
			for _, word := range tweets.([]string) {
				reduced[word]++
			}
		}
		out <- reduced
	}

	producer := func(searchKey string, counter int) chan interface{} {
		big_string := make(chan interface{})

		go func() {
			v := url.Values{}
			v.Set("count", string(counter))
			searchResult, _ := api.GetSearch(searchKey, v)
			for _, tweet := range searchResult {
				big_string <- tweet.Text
			}
			defer close(big_string)
		}()

		return big_string
	}

	for word, counter := range MapReduce(mapper, reducer, producer("golang", 100), 2).(map[string]int) {
		fmt.Println(word, counter)
	}
}
