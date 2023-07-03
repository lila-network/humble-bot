/*
 * Copyright (c) 2023 Laura Kalb <dev@lauka.net>
 * The code of this project is available under the MIT license. See the LICENSE file for more info.
 *
 */

package db

import (
	"codeberg.org/lauralani/humble-bot/models"
	"context"
	"database/sql"
	"github.com/spf13/viper"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"log"
)

func Initialize() {
	dsn := viper.GetString("database.url")
	if dsn == "" {
		log.Fatalln("config item database.url is invalid! Aborting...")
	}

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	models.DB = bun.NewDB(sqldb, pgdialect.New())

	_, err := models.DB.NewCreateTable().IfNotExists().Model((*models.QueueItem)(nil)).Exec(context.Background())
	if err != nil {
		log.Panicf("[DB] couldn't create database: %q", err)
	}

	_, err = models.DB.NewCreateTable().IfNotExists().Model((*models.SeenBundle)(nil)).Exec(context.Background())
	if err != nil {
		log.Panicf("[DB] couldn't create database: %q", err)
	}

}
