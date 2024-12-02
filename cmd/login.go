package cmd

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"tikoflano/aoc/lib/constants"
	"tikoflano/aoc/lib/utils"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Opens the login page so you can get the session cookie and save it",
	Long:  `Sets up the cookie used in the requests to get data from the Advent of Code page.`,
	Run: func(cmd *cobra.Command, args []string) {
		if noOpen, _ := cmd.Flags().GetBool("no-open"); !noOpen {
			cmnd := exec.Command("xdg-open", fmt.Sprintf("%s/auth/login", constants.BaseURL))
			cmnd.Start()
		}

		validate := func(input string) error {
			re := regexp.MustCompile("^[a-f0-9]{128}$")

			if match := re.MatchString(input); !match {
				return errors.New("invalid cookie")
			}

			return nil
		}

		prompt := promptui.Prompt{
			Label:    "Cookie",
			Validate: validate,
		}

		result, err := prompt.Run()
		utils.CheckError(err, "Invalid input")

		sessionFile := utils.CreateFileIfNotExists(utils.SessionFilePath)
		utils.CheckError(err, "Failed to create problem file")

		defer func() {
			err := sessionFile.Close()
			utils.CheckError(err, "Failed to close problem file")
		}()

		_, err = sessionFile.Write([]byte(result))
		utils.CheckError(err, "Failed to write to session file")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().BoolP("no-open", "n", false, "do not open the browser")
}
