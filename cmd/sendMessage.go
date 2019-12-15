/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/nlopes/slack"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	SLACK_CHANNEL string = "go-test"
)

// sendMessageCmd represents the sendMessage command
var sendMessageCmd = &cobra.Command{
	Use:   "sendMessage",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// check for bad configuration
		token := viper.GetString(SLACK_TOKEN_KEY)
		if token == "" {
			cmd.Printf("No slack token found. Please run `%s %s`\n", rootCmd.Use, configureCmd.Use)
			return
		}

		api := slack.New(token)
		_, _, err := api.PostMessage(SLACK_CHANNEL, slack.MsgOptionText(args[0], false))
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	slackCmd.AddCommand(sendMessageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendMessageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendMessageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
