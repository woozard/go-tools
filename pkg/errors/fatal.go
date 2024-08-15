package errors

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"os"
)

// FatalWithCode will print the error and exit
// Application with given code
func FatalWithCode(err string, code int) {
	if err != "" {
		fmt.Println(aurora.Red(err))
		os.Exit(code)
	}
}

// FatalErrorWithCode will print the error message and exit
// Application with given code
func FatalErrorWithCode(err error, code int) {
	if err != nil {
		FatalWithCode(err.Error(), code)
	}
}

// FatalError Will print error message and will exit
// app with code 1
func Fatal(message string) {
	FatalWithCode(message, 1)
}

// FatalError Will print error message and will exit
// app with code 1
func FatalError(err error) {
	FatalErrorWithCode(err, 1)
}

// FatalError Will print error message and will exit
// app with code 1
func FatalIfError(err error) {
	FatalErrorWithCode(err, 1)
}

// FatalOrValue
func FatalOrValue(value interface{}, err error) interface{} {
	FatalError(err)
	return value
}

// FatalOrString
func FatalOrString(value string, err error) string {
	return FatalOrValue(value, err).(string)
}

// FatalOrString
func FatalOrBool(value bool, err error) bool {
	return FatalOrValue(value, err).(bool)
}

// FatalOrString
func FatalOrInt(value int, err error) int {
	return FatalOrValue(value, err).(int)
}
