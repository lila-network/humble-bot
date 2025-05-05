/*
Copyright © 2023 Laura Kalb <dev@lauka.net>
The code of this project is available under the MIT license. See the LICENSE file for more info.
*/

package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.lila.network/adora-kalb/humble-bot/app"
)

// queueshowCmd represents the queue command
var queueshowCmd = &cobra.Command{
	Use:     "show",
	Short:   "Show items in queue",
	Run:     app.RunQueueShow,
	Args:    cobra.ExactArgs(0),
	Aliases: []string{"list"},
}

func init() {
	queueCmd.AddCommand(queueshowCmd)
}
