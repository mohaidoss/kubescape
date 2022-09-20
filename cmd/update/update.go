package update

//This update command updates to the latest kubescape release.
//Example:-
//          kubescape update

import (
	"fmt"
	"os/exec"
	"runtime"

	logger "github.com/kubescape/go-logger"
	"github.com/kubescape/kubescape/v2/core/cautils"
	"github.com/spf13/cobra"
)

func GetUpdateCmd() *cobra.Command {
	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update your version",
		Long:  ``,
		RunE: func(_ *cobra.Command, args []string) error {
			//Checking the user's version of kubescape to the latest release
			if cautils.BuildNumber == cautils.LatestReleaseVersion {
				//your version == latest version
					logger.L().Info( ("You are in the latest version"))
				} else {
				//execute the install.sh if linux, install.ps1 for windows,.....depending on your OS
				const OSTYPE string = runtime.GOOS
				switch OSTYPE{
				case "linux":
					cautils.StartSpinner()
					//run the installation command for linux
					cmd, err := exec.Command("./install.sh").Output()
					if err != nil {
						logger.L().Fatal(err.Error())
					}
					fmt.Println(string(cmd))
					cautils.StopSpinner()

				case "windows":
					cautils.StartSpinner()
					//run the installation command for windows
					_, err := exec.Command("./install.ps1").Output()
					if err != nil {
						logger.L().Fatal(err.Error())
					}
					cautils.StopSpinner()

				default :
					cautils.StartSpinner()
					//run the installation command for macOS
					_, err := exec.Command("./install.sh").Output()
					if err != nil {
						logger.L().Fatal(err.Error())
					}
					cautils.StopSpinner()

				}
			}
			return nil
		},
	}
	return updateCmd
}
