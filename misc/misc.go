/*
 * Copyright (c) 2023 Laura Kalb <dev@lauka.net>
 * The code of this project is available under the MIT license. See the LICENSE file for more info.
 *
 */

package misc

import (
	"github.com/microcosm-cc/bluemonday"
	"html"
)

func Sanitize(d string) string {
	p := bluemonday.StrictPolicy()
	return html.UnescapeString(p.Sanitize(d))
}
