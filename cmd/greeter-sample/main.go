package main

import (
	"github.com/spf13/cobra"
	"m-tools/cmd/greeter-sample/command"
	"math/rand"
	"runtime"
	"time"
)

var (
	GitCommit = "n/a"
	GitTag    = "0"
	BuildTime = "n/a"
	GitAuthor = "none"
)

func main() {
	//  设置并发度
	CORE_NUM := runtime.NumCPU()
	runtime.GOMAXPROCS(CORE_NUM * 4)
	// debug.SetGCPercent(200)
	rand.Seed(time.Now().UnixNano())

	var rootCmd = &cobra.Command{Use: "dsp_retarget_server"}

	rootCmd.AddCommand(command.NewClientCmd())
	rootCmd.AddCommand(command.NewServerCmd())
	rootCmd.AddCommand(command.NewVersionCmd(GitTag, BuildTime, GitCommit, GitAuthor))

	_ = rootCmd.Execute()
}
