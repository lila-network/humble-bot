/*
 * Copyright (c) 2023 Laura Kalb <dev@lauka.net>
 * The code of this project is available under the MIT license. See the LICENSE file for more info.
 *
 */

package app

import (
	"fmt"
	"github.com/spf13/cobra"
)

func RunRoot(cmd *cobra.Command, args []string) {
	for _, category := range HumbleCategories {
		UpdateCategory(category)
	}
}

func RunDaemon(cmd *cobra.Command, args []string) {
	fmt.Println("daemon called")
}
