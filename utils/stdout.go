package utils

import "os"

// charBell is the ASCII code for the bell character (\a)
const charBell = 7

// noBellStdout is a wrapper around os.Stdout that suppresses the bell character (\a)
type noBellStdout struct{}

func (n *noBellStdout) Write(p []byte) (int, error) {
	if len(p) == 1 && p[0] == charBell {
		return 0, nil
	}
	return os.Stdout.Write(p)
}

func (n *noBellStdout) Close() error {
	return os.Stdout.Close()
}

// NoBellStdout returns a stdout wrapper that doesn't ring the terminal bell
// https://github.com/manifoldco/promptui/issues/49 & https://github.com/manifoldco/promptui/issues/49#issuecomment-1012640880
var NoBellStdout = &noBellStdout{}
