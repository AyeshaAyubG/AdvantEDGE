/*
 * Copyright (c) 2022  The AdvantEDGE Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"errors"
	"fmt"
	"sort"

	"github.com/roymx/viper"
	"github.com/spf13/cobra"

	sandbox "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-sandbox-ctrl-client"
)

// replayCmd represents the replay command
var replayCmd = &cobra.Command{
	Use:   "replay -s <sandbox> <action>",
	Short: "Use and manage auto-replay feature",
	Long: `AdvantEDGE supports creation and usage of auto-replay files.

Replay files are 'script-like' files that are maintained in AdvantEDGE document store.
Replay file contain sequence of events that can be automatically replayed following the specific time sequence of events.

Multiple actions can be performed on replay files.`,

	Run: func(cmd *cobra.Command, args []string) {
		keys := viper.AllKeys()
		sort.Strings(keys)

		_ = cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(replayCmd)
}

func createClient(path string) (*sandbox.APIClient, error) {
	// Create & store client for App REST API
	ceClientCfg := sandbox.NewConfiguration()
	ceClientCfg.BasePath = path
	ceClient := sandbox.NewAPIClient(ceClientCfg)
	if ceClient == nil {
		err := errors.New("Failed to create REST API client")
		return nil, err
	}
	return ceClient, nil
}

func printError(errorString string, err error, verbose bool) {
	if verbose {
		fmt.Println(errorString, err)
	} else {
		fmt.Println("Command failed, use -v for more details")
	}
}

func getBasePath(cmd *cobra.Command) string {
	host := viper.GetString("node.ip")
	sandbox, _ := cmd.Flags().GetString("sandbox")
	reqString := "http://" + host + "/" + sandbox + "/sandbox-ctrl/v1"
	return reqString
}

func setSandboxFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("sandbox", "s", "", "Sandbox to send request to")
	_ = cmd.MarkFlagRequired("sandbox")
}
