package log

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

// Normal logging
var Normal = color.New(color.FgWhite).SprintFunc()

// Error logging
var Error = color.New(color.Bold, color.FgRed).SprintFunc()

// Information logging
var Information = color.New(color.Bold, color.FgCyan).SprintFunc()

// InformationPrint logs information
func InformationPrint(text string) {
	fmt.Println(Information(text))
}

// NormalPrint logs information
func NormalPrint(text string) {
	fmt.Println(Normal(text))
}

// ErrorPrint logs information
func ErrorPrint(text string) {
	fmt.Println(Error(text))
}

// Loading provides a loading animation during processes
var Loading = spinner.New(spinner.CharSets[39], 100*time.Millisecond)
