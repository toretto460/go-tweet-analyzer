package main

import (
	"github.com/ChimeraCoder/anaconda"
	"net/url"
)

type TwitterConfiguration struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

type Twitter struct {
	api *anaconda.TwitterApi
}

func newTwitter(conf TwitterConfiguration) Twitter {
	anaconda.SetConsumerKey(conf.ConsumerKey)
	anaconda.SetConsumerSecret(conf.ConsumerSecret)
	t := Twitter{}
	t.api = anaconda.NewTwitterApi(conf.AccessToken, conf.AccessTokenSecret)
	return t
}

func (t Twitter) find(searchKey string, counter int) chan interface{} {
	stream := make(chan interface{})

	go func() {
		v := url.Values{}
		v.Set("count", string(counter))
		searchResult, _ := t.api.GetSearch(searchKey, v)
		for _, tweet := range searchResult {
			stream <- tweet.Text
		}
		defer close(stream)
	}()

	return stream
}
