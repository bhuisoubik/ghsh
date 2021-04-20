package auth

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func printGenTokSteps() {
	fmt.Println("Steps to Generate Personal Access Token\n\t1. Go to your Github Profile\n\t2. Go to Settings -> Developer Settings -> Personal access tokens -> Generate new token\n\t3. Copy the Token and paste it in the below prompt")
}

func Login() {
	printGenTokSteps()
	fmt.Print("GitHub Token: ")
	byteToken, _ := terminal.ReadPassword(int(syscall.Stdin))

	os.Setenv("GHVE_USER_AUTH_TOKEN", string(byteToken))
}

func Logout() {
	os.Setenv("GHVE_USER_AUTH_TOKEN", "")
}