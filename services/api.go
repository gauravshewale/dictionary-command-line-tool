package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gauravshewale/dictionary-command-line-tool/config"
	"github.com/gauravshewale/dictionary-command-line-tool/model"
)

func SearchDictionary(searchString string) ([]model.SearchResults, error) {
	fmt.Println("Searching... Please wait...")
	reqUrl := config.Manager.DictionaryBaseUrl + "/" + searchString
	parsedUrl, err := url.Parse(reqUrl)
	if err != nil {
		log.Fatal("Error parsing URL::", err)
		return []model.SearchResults{}, err
	}
	q, _ := url.ParseQuery(parsedUrl.RawQuery)
	q.Add("key", config.Manager.ApiKey)
	parsedUrl.RawQuery = q.Encode()
	start := time.Now()
	res, err := http.Get(parsedUrl.String())
	fmt.Printf("Time Taken to search results: %v\n\n", time.Since(start))
	if err != nil {
		log.Fatalln("Error occurred while looking information::", err)
		return []model.SearchResults{}, err
	}
	defer res.Body.Close()
	rs, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("Error occurred while reading response::", err)
		return []model.SearchResults{}, err
	}
	var ob []model.SearchResults
	err = json.Unmarshal(rs, &ob)
	if err != nil {
		log.Fatalln("Error occurred while reading response::", err)
		return []model.SearchResults{}, err
	}
	return ob, nil
}
