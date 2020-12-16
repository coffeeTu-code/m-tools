package command

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var (
	__VersionMsg *VersionInfo
)

type VersionInfo struct {
	Tag       string `json:"tag"`
	BuildTime string `json:"build_time"`
	Commit    string `json:"commit"`
	Author    string `json:"author"`
}

func (v *VersionInfo) String() string {
	d, _ := json.Marshal(v)
	return string(d)
}

func NewVersionCmd(GitTag, BuildTime, GitCommit, GitAuthor string) *cobra.Command {

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Show version info",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(__VersionMsg.String())
		},
	}

	__VersionMsg = &VersionInfo{
		Tag:       GitTag,
		BuildTime: BuildTime,
		Commit:    GitCommit,
		Author:    GitAuthor,
	}
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	return versionCmd
}
