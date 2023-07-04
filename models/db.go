/*
 * Copyright (c) 2023 Laura Kalb <dev@lauka.net>
 * The code of this project is available under the MIT license. See the LICENSE file for more info.
 *
 */

package models

import (
	"github.com/uptrace/bun"
	"time"
)

var DB *bun.DB

type QueueItem struct {
	bun.BaseModel `bun:"table:queue"`
	ID            int64     `bun:"id,pk,autoincrement"`
	Name          string    `bun:"name"`
	Headline      string    `bun:"headline"`
	Body          string    `bun:"body"`
	URL           string    `bun:"url"`
	Hashtags      string    `bun:"hashtags"`
	Created       time.Time `bun:"created"`
}

type SeenBundle struct {
	bun.BaseModel `bun:"table:seen"`
	ID            int64     `bun:"id,pk,autoincrement"`
	Name          string    `bun:"name,notnull"`
	URL           string    `bun:"url,notnull"`
	SeenAt        time.Time `bun:"seenat"`
}
