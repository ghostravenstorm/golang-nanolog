
package nanolog

import (
   "log"
   // "fmt"
)

var DebugLevel int = 2

///
/// 0 = No Logging
/// 1 = Errors only.
/// 2 = Warnings and Errors.
/// 3 = All loging.
///

const (
   NoLogging = iota
   Error
   Warning
   Info
)

func init () () {
   log.SetFlags(log.Lshortfile)
}

func LogError (v ...interface{}) () {
   if !(DebugLevel >= Error) { return }
   log.Print(v)
}

func LogWarning () () {

}

func Log (v ...interface{}) () {
   log.Print(v)
}

func Logln () () {

}
