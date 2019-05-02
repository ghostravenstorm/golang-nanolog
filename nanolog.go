/// Author:  GhostRavenstorm
/// Project: Nanolog
/// Date:    2019-03-06
/// Summary:
///   Go package with the intent to provide a logging class that allows for togglable
///   error, warning, info, and debug statments wihtout needing to cleanup log statements
///   that do not wish to be shown at a given instance.
///
///   Set LogLevel accordingly for which statements are shown at runtime.
///   LogLevel INFO (3) is the default for all general purpose statements.
///
///   This package can be used as a replacement for fmt.Print() and log.Print().
///
///   More functionality and polish to come as the need arises.

package nanolog

import (
   "log"
	"fmt"
	"os"
)

var LogLevel int = 3

///
/// All log levels equal to or less than the set log level will
/// print to console.
/// 0  = No Logging
/// 1  = Errors only.
/// 2  = Warnings and Errors.
/// 3  = All info statments.
/// 4  = Debugging & trace.
/// 5+ = User defined priority.
///

const (
   LnoLogging = iota
   Lerror
   Lwarning
   Linfo
   Ldebug
)

var (
	error   *log.Logger
	warning *log.Logger
	info    *log.Logger
	debug   *log.Logger
)

func init () () {
	error   = log.New(os.Stderr, "[ERR] ", log.Lshortfile)
	warning = log.New(os.Stdout, "[WARN] ", 0)
	info    = log.New(os.Stdout, "[INFO] ", 0)
	debug   = log.New(os.Stdout, "[DEBUG] ", log.Lshortfile)
}

func AppendLine(lines int) () {
	if lines == 0 { return }
	fmt.Print("\n")
	lines--
	AppendLine(lines)
}

func Error (v ...interface{}) () {
	if LogLevel < Lerror { return }
	error.Output(2, fmt.Sprint(v...))
}

func Errorln (v ...interface{}) () {
	if LogLevel < Lerror { return }
	error.Output(2, fmt.Sprintln(v...))
}

/// Outputs standard error message with a set calldepth to log line position.
/// Calldepth greater than 2 is not recommended if called directly from main().
func ErrorlnCalldepth (calldepth int, v ...interface{}) {
	if LogLevel < Lerror { return }
	error.Output(calldepth, fmt.Sprintln(v...))
}

func Errorf (format string, v ...interface{}) () {
	if LogLevel < Lerror { return }
	error.Output(2, fmt.Sprintf(format, v...))
}

func Warning (v ...interface{}) () {
	if LogLevel < Lwarning { return }
	warning.Output(2, fmt.Sprint(v...))
}

func Warningln (v ...interface{}) () {
	if LogLevel < Lwarning { return }
	warning.Output(2, fmt.Sprintln(v...))
}

func Warningf (format string, v ...interface{}) () {
	if LogLevel < Lwarning { return }
	warning.Output(2, fmt.Sprintf(format, v...))
}

func Info (v ...interface{}) () {
	if LogLevel < Linfo { return }
	info.Output(2, fmt.Sprint(v...))
}

func Infoln (v ...interface{}) () {
	if LogLevel < Linfo { return }
	info.Output(2, fmt.Sprintln(v...))
}

func Infof (format string, v ...interface{}) () {
	if LogLevel < Linfo { return }
	info.Output(2, fmt.Sprintf(format, v...))
}

/// Prints a debug statement with call depth of 2.
func Debug (v ...interface{}) () {
	if LogLevel < Ldebug { return }
	debug.Output(2, fmt.Sprint(v...))
}

/// Prints a new line debug statement with call depth of 2.
func Debugln (v ...interface{}) () {
	if LogLevel < Ldebug { return }
	debug.Output(2, fmt.Sprintln(v...))
}

/// Prints a formatted debug statement with call depth of 2.
func Debugf (format string, v ...interface{}) () {
	if LogLevel < Ldebug { return }
	debug.Output(2, fmt.Sprintf(format, v...))
}

/// Debug statements with adjustible priority.
/// Loglevel must be set equal to or greater than the set priority for these statements to print.

func Debugp (p int, v ...interface{}) () {
	if LogLevel < p { return }
	debug.Output(2, fmt.Sprint(v...))
}

func Debuglnp (p int, v ...interface{}) () {
	if LogLevel < p { return }
	debug.Output(2, fmt.Sprintln(v...))
}

func Debugfp (p int, format string, v ...interface{}) () {
	if LogLevel < p { return }
	debug.Output(2, fmt.Sprintf(format, v...))
}
