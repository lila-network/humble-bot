/*
 * Copyright (c) 2023 Laura Kalb <dev@lauka.net>
 * The code of this project is available under the MIT license. See the LICENSE file for more info.
 *
 */

package misc

import (
	"net/http"
	"time"

	"gitlab.lila.network/adora-kalb/humble-bot/constants"
)

func CustomHttpClient() *http.Client {
	c := http.Client{}
	c.Timeout = time.Duration(60) * time.Second
	return &c
}

func CustomHttpRequest() *http.Request {
	h := map[string][]string{
		"User-Agent": {constants.UserAgent},
	}
	r := http.Request{Header: h}
	return &r
}
