package exec

import "regexp"

func GetArgs(c string) []string {
	regex := regexp.MustCompile(`[^\s"']+|"([^"]*)"|'([^']*)`) 
	return regex.FindAllString(c, -1)
}