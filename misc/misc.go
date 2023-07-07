/*
 * Copyright (c) 2023 Laura Kalb <dev@lauka.net>
 * The code of this project is available under the MIT license. See the LICENSE file for more info.
 *
 */

package misc

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"html"
)

func Sanitize(d string) string {
	p := bluemonday.StrictPolicy()
	return html.UnescapeString(p.Sanitize(d))
}

func CheckConfig() {
	visibility := viper.GetString("mastodon.visibility")
	visibilityallowed := []string{"public", "unlisted", "private"}

	if !isValueInList(visibility, visibilityallowed) {
		log.Fatal().Str("component", "config").Str("config_item", "mastodon.visibility").
			Msgf("Config value must be one of [public, unlisted, private]: %q", visibility)
	}

}

func isValueInList(value string, list []string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}
