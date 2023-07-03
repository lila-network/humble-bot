/*
 * Copyright (c) 2023 Laura Kalb <dev@lauka.net>
 * The code of this project is available under the MIT license. See the LICENSE file for more info.
 *
 */

package cmd

import (
	"codeberg.org/lauralani/humble-bot/app"
	"codeberg.org/lauralani/humble-bot/db"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "humble-bot",
	Short: "A simple daemon posting new Humble Bundles on Mastodon",
	Long:  `This app regularly checks for new Humble Bundles and posts updates on Mastodon`,
	Run:   app.RunRoot,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig, db.Initialize)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $workingdir/config.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {

		workingdir, err := os.Getwd()
		cobra.CheckErr(err)

		viper.AddConfigPath(workingdir)
		viper.SetConfigName("config")
		//viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Printf("Using config file: %v\n", viper.ConfigFileUsed())
	} else {
		log.Panicf("Can't find config file: %v\n", viper.ConfigFileUsed())
	}
}
