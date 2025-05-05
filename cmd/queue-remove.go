/*
Copyright © 2023 Laura Kalb <dev@lauka.net>
The code of this project is available under the MIT license. See the LICENSE file for more info.
*/

package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.lila.network/adora-kalb/humble-bot/app"
)

// queueCmd represents the queue command
var queueremoveCmd = &cobra.Command{
	Use:     "remove QUEUEID",
	Short:   "Delete single queue item",
	Run:     app.RunQueueRemove,
	Args:    cobra.ExactArgs(1),
	Example: "humble-bot queue remove 25",
	Aliases: []string{"delete"},
}

func init() {
	queueCmd.AddCommand(queueremoveCmd)
}
