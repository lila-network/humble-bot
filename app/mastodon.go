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
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"io"
	"net/url"
	"strings"
)

func postQueueItemToMastodon(bundle models.QueueItem) error {
	urlstring := viper.GetString("mastodon.url") + "/api/v1/statuses"
	endpoint, _ := url.Parse(urlstring)

	token := viper.GetString("mastodon.token")
	idemkey := uuid.New().String()

	builder := new(strings.Builder)

	builder.WriteString(bundle.Headline)
	builder.WriteString("\n\n")

	builder.WriteString(bundle.Body)
	builder.WriteString("\n\n")

	builder.WriteString(bundle.URL)

	builder.WriteString("\n\n")
	builder.WriteString(bundle.Hashtags)

	payload := url.Values{}
	payload.Add("status", builder.String())
	payload.Add("visibility", "private")
	payload.Add("language", "en")

	client := misc.CustomHttpClient()
	req := misc.CustomHttpRequest()
	req.URL = endpoint
	req.Method = "POST"
	req.Body = io.NopCloser(bytes.NewBuffer([]byte(payload.Encode())))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Idempotency-Key", idemkey)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		log.Error().Int("code", res.StatusCode).Str("status", res.Status).
			Str("url", urlstring).Msg("Error with http request")
		return errors.New("couldn't POST message to Mastodon")
	}

	log.Info().
		Str("url", bundle.URL).Msg("posted new Bundle to Mastodon =)")
	return nil
}
