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

// getMembersCmd represents the getMembers command
var getMembersCmd = &cobra.Command{
	Use:   "getMembers",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		git := gitlab.NewClient(nil, viper.GetString("gitlab-token"))
		project, _, err := git.Projects.GetProject("simplenexus-engineering/simplenexus.com", nil)
		if err != nil {
			panic(err)
		}
		options := &gitlab.ListProjectMembersOptions{
			ListOptions: gitlab.ListOptions{
				Page:    1,
				PerPage: 100,
			},
		}
		project_members, _, err := git.ProjectMembers.ListAllProjectMembers(project.ID, options)
		for _, project_member := range project_members {
			cmd.Printf("Name: %-30s Username: %-20s Access Level: %d\n", project_member.Name, project_member.Username, project_member.AccessLevel)
		}
	},
}

func init() {
	gitlabCmd.AddCommand(getMembersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getMembersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getMembersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
