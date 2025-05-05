/*
 * Copyright (c) 2023 Laura Kalb <dev@lauka.net>
 * The code of this project is available under the MIT license. See the LICENSE file for more info.
 *
 */

package app

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gitlab.lila.network/adora-kalb/humble-bot/misc"
	"gitlab.lila.network/adora-kalb/humble-bot/models"
)

func postQueueItemToMastodon(bundle models.QueueItem) error {
	urlstring := viper.GetString("mastodon.url") + "/api/v1/statuses"
	endpoint, _ := url.Parse(urlstring)
	visibility := viper.GetString("mastodon.visibility")
	var post models.MastodonPost

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
	payload.Add("visibility", visibility)
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

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error().Str("func", "postQueueItemToMastodon.body.Close()").Msg(err.Error())
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error().Str("func", "postQueueItemToMastodon.body.ReadAll()").Msg(err.Error())
	}

	err = json.Unmarshal(body, &post)
	if err != nil {
		log.Error().Str("func", "postQueueItemToMastodon.body.Unmarshal()").
			Msgf("can't unmarshal body: %q", err)
	}

	log.Info().
		Str("bundle-url", bundle.URL).Str("post-url", post.URL).Msg("posted new Bundle =)")
	return nil
}
