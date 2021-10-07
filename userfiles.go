package main

import (
	"flag"
	"fmt"
	"github.com/ashishkujoy/go-tools/utils"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"syscall"
)

func main() {
	directory := flag.String("d", ".", "directory to search for")
	userId := flag.Int("uid", -1, "user id for which files need to be looked up")
	username := flag.String("user", "", "username for which files need to be looked up")
	flag.Parse()
	var userIdToLookFor int32

	if *userId < 0 {
		if *username == "" {
			fmt.Println("Usage: userfiles < -uid 501 | -user 'some user' > ")
			os.Exit(1)
		} else {
			userDetails, err := user.Lookup(*username)
			utils.ExitIfErrNotNil(err)
			parseInt, err := strconv.ParseInt(userDetails.Uid, 10, 32)
			utils.ExitIfErrNotNil(err)
			userIdToLookFor = int32(parseInt)
		}

	} else {
		userIdToLookFor = int32(*userId)
	}

	err := filepath.Walk(*directory, walkFor(userIdToLookFor))
	utils.ExitIfErrNotNil(err)
}

func userIdOfFileOwner(filename string) int32 {
	fileStat, err := os.Stat(filename)
	utils.ExitIfErrNotNil(err)
	uid := fileStat.Sys().(*syscall.Stat_t).Uid
	return int32(uid)
}

func walkFor(userId int32) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		owner := userIdOfFileOwner(path)
		if owner == userId {
			fmt.Println(path)
		}
		return nil
	}
}
