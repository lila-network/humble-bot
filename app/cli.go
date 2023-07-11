/*
 * Copyright (c) 2023 Laura Kalb <dev@lauka.net>
 * The code of this project is available under the MIT license. See the LICENSE file for more info.
 *
 */

package app

import (
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"time"
)

func RunDaemon(cmd *cobra.Command, args []string) {
	StartCronJobs()
}

func StartCronJobs() {
	s := gocron.NewScheduler(time.UTC)

	_, err := s.Every(viper.GetString("humblebundle.pollinterval")).Do(UpdateBundles)
	if err != nil {
		log.Error().Str("func", "StartCronJobs").Str("job", "UpdateBundlesJob").
			Msgf("Scheduler Error: %q", err)
		return
	}

	_, err = s.Every(viper.GetString("mastodon.postinterval")).WaitForSchedule().Do(RunSingleQueueItem)
	if err != nil {
		log.Error().Str("func", "StartCronJobs").Str("job", "RunSingleQueueItem").
			Msgf("Scheduler Error: %q", err)
		return
	}
	s.StartBlocking()
}
