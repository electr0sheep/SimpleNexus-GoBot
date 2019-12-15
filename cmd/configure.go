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
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	LAST_N_CHARACTERS      int = 4
	GITLAB_TOKEN_LENGTH    int = 20
	SLACK_TOKEN_LENGTH     int = 53
	ATLASSIAN_TOKEN_LENGTH int = 24
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// configure gitlab token
		gitlabToken := viper.GetString("gitlab-token")
		if gitlabToken == "" {
			gitlabToken = "None"
		}
		cmd.Println(`Go to https://gitlab.com/profile/personal_access_tokens, create a token with an
api scope, and copy the token here`)
		cmd.Printf("Gitlab Token [****************%s]: ", gitlabToken[len(gitlabToken)-LAST_N_CHARACTERS:])
		if _, err := fmt.Scanln(&gitlabToken); err != nil && err.Error() != "unexpected newline" {
			cmd.Printf("\n")
			panic(err)
		}
		if len(gitlabToken) == GITLAB_TOKEN_LENGTH {
			viper.Set("gitlab-token", gitlabToken)
		} else if len(gitlabToken) != 0 {
			cmd.Printf("Unable to save gitlab token, length wasn't %d.\nToken: %s", GITLAB_TOKEN_LENGTH, gitlabToken)
		}

		cmd.Println("")

		// configure slack token
		slackToken := viper.GetString("slack-token")
		if slackToken == "" {
			slackToken = "None"
		}
		cmd.Println(`Get the token from Michael DeGraw[https://simplenexus.slack.com/team/U9NPLPKHQ],
and copy the token here`)
		cmd.Printf("Slack Token [****************%s]: ", slackToken[len(slackToken)-LAST_N_CHARACTERS:])
		if _, err := fmt.Scanln(&slackToken); err != nil && err.Error() != "unexpected newline" {
			cmd.Printf("\n")
			panic(err)
		}
		if len(slackToken) == SLACK_TOKEN_LENGTH {
			viper.Set("slack-token", slackToken)
		} else if len(slackToken) != 0 {
			cmd.Printf("Unable to save slack token, length wasn't %d.\nToken: %s", SLACK_TOKEN_LENGTH, slackToken)
		}

		cmd.Println("")

		// configure atlassian email
		atlassianEmail := viper.GetString("atlassian-email")
		if atlassianEmail == "" {
			atlassianEmail = "None"
		}
		cmd.Println(`This should be the email you use to log into Atlassian`)
		cmd.Printf("Atlassian Email [%s]: ", atlassianEmail)
		if _, err := fmt.Scanln(&atlassianEmail); err != nil && err.Error() != "unexpected newline" {
			cmd.Printf("\n")
			panic(err)
		}
		if len(atlassianEmail) > 0 {
			viper.Set("atlassian-email", atlassianEmail)
		}

		cmd.Println("")

		// configure atlassian token
		atlassianToken := viper.GetString("atlassian-token")
		if atlassianToken == "" {
			atlassianToken = "None"
		}
		cmd.Println(`Go to https://id.atlassian.com/manage/api-tokens, create a token, and copy the
token here`)
		cmd.Printf("Atlassian Token [****************%s]: ", atlassianToken[len(atlassianToken)-LAST_N_CHARACTERS:])
		if _, err := fmt.Scanln(&atlassianToken); err != nil && err.Error() != "unexpected newline" {
			cmd.Printf("\n")
			panic(err)
		}
		if len(atlassianToken) == ATLASSIAN_TOKEN_LENGTH {
			viper.Set("atlassian-token", atlassianToken)
		} else if len(atlassianToken) != 0 {
			cmd.Printf("Unable to save atlassian token, length wasn't %d.\nToken: %s", ATLASSIAN_TOKEN_LENGTH, atlassianToken)
		}

		// write tokens
		err := viper.WriteConfig()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
