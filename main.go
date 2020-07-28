package main

import (
	"fmt"
	"github.com/swgloomy/gutil"
	"io/ioutil"
	"os/exec"
)

var (
	rootDir    = "../github.com"
	gitDirName = ".git"
	number     = 0
	gitDirAll  = []string{}
)

func main() {
	gitPull(rootDir)
}

func gitPull(dirPath string) {
	dirAllIn, err := gutil.GetMyAllDirByDir(dirPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bo := false
	for _, item := range *dirAllIn {
		bo = item == gitDirName
		if bo {
			break
		}
	}
	if bo {
		gitPullCmd(dirPath)
	} else {
		for _, item := range *dirAllIn {
			gitPull(fmt.Sprintf("%s/%s", dirPath, item))
		}
	}
}

func gitPullCmd(dirPath string) {
	fmt.Println(dirPath)

	cmd := exec.Command("git", "-C", dirPath, "pull")

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer stdout.Close()

	if err := cmd.Start(); err != nil {
		fmt.Println(err.Error())
		return
	}
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(opBytes))
}
