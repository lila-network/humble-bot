/*
 * Copyright (c) 2023 Laura Kalb <dev@lauka.net>
 * The code of this project is available under the MIT license. See the LICENSE file for more info.
 *
 */

package app

import (
	"codeberg.org/lauralani/humble-bot/db"
	"codeberg.org/lauralani/humble-bot/models"
	"context"
	"database/sql"
	"github.com/rs/zerolog/log"
)

func RunSingleQueueItem() {
	log.Debug().Str("func", "RunSingleQueueItem").Msg("starting single queue item run")
	var item = models.QueueItem{}

	err := models.DB.NewSelect().Model(&item).Order("created ASC").Limit(1).Scan(context.Background())
	if err != nil {
		if err == sql.ErrNoRows {
			log.Debug().Str("func", "RunSingleQueueItem").
				Str("table", "queue").Msg("no item in queue")
			return
		} else {
			log.Error().Str("func", "RunSingleQueueItem").
				Str("table", "queue").Msgf("DB Error: SELECT failed: %v", err)
			return
		}
	}
	log.Debug().Str("func", "RunSingleQueueItem").Str("item", item.Name).
		Msg("got new item from db")

	err = postQueueItemToMastodon(item)
	if err != nil {
		return
	}
	log.Debug().Str("func", "RunSingleQueueItem").Str("item", item.Name).
		Msg("posted to mastodon")

	err = db.Dequeue(item)
	if err != nil {
		log.Error().Str("func", "RunSingleQueueItem").Str("item", item.Name).
			Msg("Error: couldnt dequeue item")
		return
	}
}
