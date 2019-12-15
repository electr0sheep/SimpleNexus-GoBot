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
	"os"

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
		gitlabToken := viper.GetString(GITLAB_TOKEN_KEY)
		cmd.Println(`Go to https://gitlab.com/profile/personal_access_tokens, create a token with an
api scope, and copy the token here`)
		cmd.Printf("Gitlab Token [%s]: ", getCurrentConfig(gitlabToken, true))
		gitlabToken = getNewTokenValue()
		verifyTokenLength(gitlabToken, GITLAB_TOKEN_LENGTH, GITLAB_TOKEN_KEY)
		cmd.Println("")

		// configure slack token
		slackToken := viper.GetString(SLACK_TOKEN_KEY)
		cmd.Println(`Get the token from Michael DeGraw[https://simplenexus.slack.com/team/U9NPLPKHQ],
and copy the token here`)
		cmd.Printf("Slack Token %s]: ", getCurrentConfig(slackToken, true))
		slackToken = getNewTokenValue()
		verifyTokenLength(slackToken, SLACK_TOKEN_LENGTH, SLACK_TOKEN_KEY)
		cmd.Println("")

		// configure atlassian email
		atlassianEmail := viper.GetString(ATLASSIAN_EMAIL_KEY)
		cmd.Println(`This should be the email you use to log into Atlassian`)
		cmd.Printf("Atlassian Email [%s]: ", getCurrentConfig(atlassianEmail, false))
		atlassianEmail = getNewTokenValue()
		if len(atlassianEmail) > 0 {
			viper.Set(ATLASSIAN_EMAIL_KEY, atlassianEmail)
		}
		cmd.Println("")

		// configure atlassian token
		atlassianToken := viper.GetString("atlassian-token")
		cmd.Println(`Go to https://id.atlassian.com/manage/api-tokens, create a token, and copy the
token here`)
		cmd.Printf("Atlassian Token [%s]: ", getCurrentConfig(atlassianToken, true))
		atlassianToken = getNewTokenValue()
		verifyTokenLength(atlassianToken, ATLASSIAN_TOKEN_LENGTH, ATLASSIAN_TOKEN_KEY)
		cmd.Println("")

		// create config file if it doesn't exist
		if _, err := os.Stat("./config.yaml"); os.IsNotExist(err) {
			os.Create("config.yaml")
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

func getNewTokenValue() string {
	var token string
	if _, err := fmt.Scanln(&token); err != nil && err.Error() != "unexpected newline" {
		fmt.Printf("\n")
		panic(err)
	}
	return token
}

func verifyTokenLength(token string, length int, key string) {
	if token != "" {
		if len(token) == length {
			viper.Set(key, token)
		} else {
			fmt.Printf("Unable to save %s, length wasn't %d.\nToken: %s\n", key, ATLASSIAN_TOKEN_LENGTH, token)
		}
	}
}

func getCurrentConfig(config string, obfuscate bool) string {
	if config == "" {
		config = "None"
	} else if obfuscate {
		config = fmt.Sprintf("****************%s", config[len(config)-LAST_N_CHARACTERS:])
	}
	return config
}
