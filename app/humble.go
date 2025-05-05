/*
 * Copyright (c) 2023 Laura Kalb <dev@lauka.net>
 * The code of this project is available under the MIT license. See the LICENSE file for more info.
 *
 */

package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/rs/zerolog/log"
	"gitlab.lila.network/adora-kalb/humble-bot/constants"
	"gitlab.lila.network/adora-kalb/humble-bot/db"
	"gitlab.lila.network/adora-kalb/humble-bot/misc"
	"gitlab.lila.network/adora-kalb/humble-bot/models"
)

func UpdateBundles() {
	log.Debug().Str("func", "UpdateBundles").Msg("starting bundle run")

	for _, category := range constants.HumbleCategories {
		log.Debug().Str("category", category).Msg("starting bundle update")
		endpoint, _ := url.Parse("https://www.humblebundle.com/" + category)

		client := misc.CustomHttpClient()
		req := misc.CustomHttpRequest()
		req.URL = endpoint
		req.Method = "GET"

		response, err := client.Do(req)
		if err != nil {
			log.Error().Str("category", category).Str("status", response.Status).
				Str("func", "app.UpdateBundles").Msgf("HTTP Error: %q", err)
		}
		defer func(Body io.ReadCloser) {
			err = Body.Close()
			if err != nil {
				log.Error().Str("func", "app.UpdateBundles").Msg(err.Error())
			}
		}(response.Body)

		document, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil {
			log.Error().Str("func", "app.UpdateBundles").Msg(err.Error())
		}

		document.Find("script#landingPage-json-data").Each(func(idx int, s *goquery.Selection) {
			node := s.Nodes[0]
			data := node.FirstChild.Data
			bundles, err := parseBundles([]byte(data), category)
			if err != nil {
				log.Error().Str("func", "UpdateBundles").Msg(err.Error())
			}

			for _, bundle := range bundles {
				err := handleHumbleBundle(bundle, category)
				if err != nil {
					log.Error().Str("bundle", bundle.MachineName).
						Msgf("Error handling humble bundle: %v", err)
				}
			}

			log.Info().Int("bundles", len(bundles)).Msg("finished queueing " + category + " bundles")
		})
		log.Debug().Str("category", category).Msg("finished bundle update")
	}
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

func handleHumbleBundle(bundle models.Bundle, category string) error {
	isnew, err := db.IsANewBundle(bundle)
	if err != nil {
		return err
	}

	if !isnew {
		log.Debug().Str("bundle", bundle.MachineName).Msg("bundle already exists")
		return nil
	}

	log.Info().Str("bundle", bundle.ProductURL).Msg("found new Bundle")

	err = db.Enqueue(bundle, category)
	return err
}
