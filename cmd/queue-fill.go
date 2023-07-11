/*
Copyright © 2023 Laura Kalb <dev@lauka.net>
The code of this project is available under the MIT license. See the LICENSE file for more info.
*/

package cmd

import (
	"codeberg.org/lauralani/humble-bot/app"
	"github.com/spf13/cobra"
)

// queueshowCmd represents the queue command
var queuefillCmd = &cobra.Command{
	Use:   "fill",
	Short: "Fill the queue with new Bundles",
	Run:   app.RunQueueFill,
	Args:  cobra.ExactArgs(0),
}

func init() {
	queueCmd.AddCommand(queuefillCmd)
}
