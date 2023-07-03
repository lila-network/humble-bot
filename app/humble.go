/*
 * Copyright (c) 2023 Laura Kalb <dev@lauka.net>
 * The code of this project is available under the MIT license. See the LICENSE file for more info.
 *
 */

package app

import (
	"bytes"
	"codeberg.org/lauralani/humble-bot/misc"
	"codeberg.org/lauralani/humble-bot/models"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func UpdateCategory(category string) {
	response, err := http.Get("https://www.humblebundle.com/" + category)
	if err != nil {
		log.Printf("Status: %v, Error: %q\n", response.Status, err)
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Printf("Error: %q\n", err)
		}
	}(response.Body)

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Printf("Error: %q\n", err)
	}

	document.Find("script#landingPage-json-data").Each(func(idx int, s *goquery.Selection) {
		node := s.Nodes[0]
		data := node.FirstChild.Data
		products, err := parseBundles([]byte(data), category)
		if err != nil {
			log.Printf("Error: %q\n", err)
		}

		for _, product := range products {
			builder := new(strings.Builder)

			builder.WriteString(product.TileName)
			builder.WriteString("\n\n")

			builder.WriteString(misc.Sanitize(product.MarketingBlurb))
			builder.WriteString("\n\n")

			builder.WriteString("https://www.humblebundle.com")
			builder.WriteString(product.ProductURL)

			builder.WriteString("\n\n")
			builder.WriteString("#humblebundle #humble" + category + "bundle #" + category)

			payload := url.Values{}
			payload.Add("status", builder.String())
			payload.Add("visibility", "private")
			payload.Add("language", "en")

			endpoint := viper.GetString("mastodon.url") + "/api/v1/statuses"
			token := viper.GetString("mastodon.token")
			idemkey := uuid.New().String()

			client := &http.Client{}
			req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte(payload.Encode())))
			if err != nil {
				log.Panicln(err)
			}

			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			req.Header.Add("Idempotency-Key", idemkey)
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))

			res, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}

			if res.StatusCode != 200 {
				log.Printf("Error sending Mastodon post: %q\n", res.Status)
			}

			//sum := sha256.Sum256([]byte(product.ProductURL))
			//sumstring := hex.EncodeToString(sum[:])
			log.Printf("%v: %v\n", category, misc.Sanitize(product.MarketingBlurb))
			time.Sleep(5 * time.Second)
		}

	})
}

func parseBundles(data []byte, category string) ([]models.Bundle, error) {
	switch category {
	case "books":
		var books models.Books
		err := json.Unmarshal(data, &books)
		if err != nil {
			return nil, err
		}
		return books.Data.Books.Mosaic[0].Products, nil
	case "games":
		var games models.Games
		err := json.Unmarshal(data, &games)
		if err != nil {
			return nil, err
		}
		return games.Data.Games.Mosaic[0].Products, nil
	case "software":
		var software models.Software
		err := json.Unmarshal(data, &software)
		if err != nil {
			return nil, err
		}
		return software.Data.Software.Mosaic[0].Products, nil
	default:
		return nil, fmt.Errorf("unknown category %s", category)
	}
}
