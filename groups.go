package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	var u *user.User
	if len(os.Args) == 1 {
		currentUsers, err := user.Current()
		exitErrNotNil(err)
		u = currentUsers
	} else {
		requiredUser, err := user.Lookup(os.Args[1])
		exitErrNotNil(err)
		u = requiredUser
	}
	displayUserGroupsFor(u)
}

func displayUserGroupsFor(selectedUser *user.User) {
	ids, err := selectedUser.GroupIds()
	exitErrNotNil(err)
	fmt.Printf("The current selectedUser %s, belongs to below groups\n", selectedUser.Name)
	for _, id := range ids {
		group, err := user.LookupGroupId(id)
		exitErrNotNil(err)
		fmt.Printf("| %s | %s |\n", group.Name, id)
	}
	os.Exit(0)
}

func exitErrNotNil(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
