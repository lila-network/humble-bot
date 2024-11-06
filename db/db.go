/*
 * Copyright (c) 2023 Laura Kalb <dev@lauka.net>
 * The code of this project is available under the MIT license. See the LICENSE file for more info.
 *
 */

package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"codeberg.org/lauralani/humble-bot/misc"
	"codeberg.org/lauralani/humble-bot/models"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func Initialize() {
	dsn := viper.GetString("database.url")
	if dsn == "" {
		log.Error().Str("config", viper.ConfigFileUsed()).
			Str("func", "db.Initialize").Msg("empty config value vor key database.url")
		os.Exit(1)
	}

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	models.DB = bun.NewDB(sqldb, pgdialect.New())

	_, err := models.DB.NewCreateTable().IfNotExists().Model((*models.QueueItem)(nil)).Exec(context.Background())
	if err != nil {
		log.Error().Str("table", "queue").
			Str("func", "db.Initialize").Msgf("DB Error: CREATE TABLE failed: %q", err)
		os.Exit(1)
	}

	_, err = models.DB.NewCreateTable().IfNotExists().Model((*models.SeenBundle)(nil)).Exec(context.Background())
	if err != nil {
		log.Error().Str("table", "queue").
			Str("func", "db.Initialize").Msgf("DB Error: CREATE TABLE failed: %q", err)
		os.Exit(1)
	}

}

func Enqueue(bundle models.Bundle, category string) error {
	var item = models.QueueItem{}
	var seenitem = models.SeenBundle{}

	item.Name = bundle.MachineName
	item.Headline = bundle.TileName
	item.Body = misc.Sanitize(bundle.MarketingBlurb)
	item.URL = "https://www.humblebundle.com" + bundle.ProductURL
	item.Hashtags = fmt.Sprintf("#humblebundle #humble%sbundle #%s", category, category)
	item.Created = time.Now()

	seenitem.Name = bundle.MachineName
	seenitem.URL = bundle.ProductURL
	seenitem.SeenAt = time.Now()

	log.Debug().Str("bundle", bundle.MachineName).Msg("adding bundle to seen table")
	_, err := models.DB.NewInsert().Model(&seenitem).Exec(context.Background())
	if err != nil {
		log.Debug().Str("func", "db.Enqueue").Str("bundle", bundle.MachineName).
			Str("table", "seen").Msgf("DB Error: INSERT failed: %v", err)
		return err
	}

	_, err = models.DB.NewInsert().Model(&item).Exec(context.Background())
	if err != nil {
		log.Debug().Str("func", "db.Enqueue").Str("bundle", bundle.MachineName).
			Str("table", "queue").Msgf("DB Error: INSERT failed: %v", err)
		return err
	}

	log.Info().Str("bundle-name", bundle.MachineName).Msg("added bundle to queue")
	return nil
}

func Dequeue(item models.QueueItem) error {
	_, err := models.DB.NewDelete().Model((*models.QueueItem)(nil)).
		Where("id = ?", item.ID).Exec(context.Background())
	if err != nil {
		log.Debug().Str("func", "Dequeue").Str("item", item.Name).
			Msgf("DB Error: DELETE failed: %v", err)
		return err
	}
	log.Debug().Str("func", "Dequeue").Str("item", item.Name).
		Msg("dequeued item")
	return nil
}

func IsANewBundle(bundle models.Bundle) (bool, error) {
	count, err := models.DB.NewSelect().Model((*models.SeenBundle)(nil)).Where("name = ?", bundle.MachineName).
		Where("url = ?", bundle.ProductURL).Count(context.Background())
	if err != nil {
		log.Error().Str("bundle", bundle.MachineName).Str("func", "IsANewBundle").
			Msgf("DB error: %v", err.Error())
		return false, err
	}

	if count != 0 {
		return false, nil
	}
	return true, nil
}

func GetAllQueuedItems() []models.QueueItem {
	var items []models.QueueItem
	err := models.DB.NewSelect().Model(&items).Scan(context.Background())
	if err != nil {
		log.Error().Str("func", "GetAllQueuedItems").
			Msgf("DB error: %v", err.Error())
		os.Exit(1)
	}

	return items
}
