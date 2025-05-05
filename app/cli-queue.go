/*
 * Copyright (c) 2023 Laura Kalb <dev@lauka.net>
 * The code of this project is available under the MIT license. See the LICENSE file for more info.
 *
 */

package app

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gitlab.lila.network/adora-kalb/humble-bot/db"
	"gitlab.lila.network/adora-kalb/humble-bot/models"
)

func RunQueueShow(cmd *cobra.Command, args []string) {
	items := db.GetAllQueuedItems()
	if len(items) == 0 {
		fmt.Println("The queue is empty.")
		return
	}

	t := table.NewWriter()
	t.SetStyle(table.StyleRounded)
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Short Name", "URL"})

	for _, item := range items {
		t.AppendRow(table.Row{item.ID, item.Name, item.URL})
	}
	t.Render()
}

func RunQueueRemove(cmd *cobra.Command, args []string) {
	id, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		log.Error().Str("parameter", args[0]).Msg("Parameter is not a valid number!")
		os.Exit(1)
	}
	var bundle models.QueueItem
	bundle.ID = id

	err = db.Dequeue(bundle)
	if err != nil {
		log.Error().Int64("parameter", id).Msgf("Couldn't remove item from queue: %q", err)
		os.Exit(1)
	}

	log.Info().Msgf("Deleted item with ID %v from queue", id)
}

func RunQueueFill(cmd *cobra.Command, args []string) {
	UpdateBundles()
}
