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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
)

// createBranchCmd represents the createBranch command
var createBranchCmd = &cobra.Command{
	Use:   "createBranch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// check for bad configuration
		token := viper.GetString(GITLAB_TOKEN_KEY)
		if token == "" {
			cmd.Printf("No gitlab token found. Please run `%s %s`\n", rootCmd.Use, configureCmd.Use)
			return
		}

		git := gitlab.NewClient(nil, token)
		project, _, err := git.Projects.GetProject("simplenexus-engineering/simplenexus.com", nil)
		if err != nil {
			panic(err)
		}
		branch_name := args[0]
		_, _, err = git.Branches.CreateBranch(project.ID, &gitlab.CreateBranchOptions{Branch: &branch_name, Ref: &project.DefaultBranch})
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	gitlabCmd.AddCommand(createBranchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createBranchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createBranchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
