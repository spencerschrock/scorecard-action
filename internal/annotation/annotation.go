// Package annotation implements GitHub workflow run annotations.
//
// https://docs.github.com/actions/writing-workflows/choosing-what-your-workflow-does/workflow-commands-for-github-actions
package annotation

import "fmt"

// Notice creates a notice annotation.
func Notice(msg string) {
	message("notice", msg)
}

// Warning creates a warning annotation.
func Warning(msg string) {
	message("warning", msg)
}

// Error creates an error annotation.
func Error(msg string) {
	message("error", msg)
}

// we can support file level details later if needed.
func message(level, msg string) {
	fmt.Printf("::%s ::%s\n", level, msg)
}
