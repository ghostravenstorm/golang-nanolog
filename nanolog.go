
package nanolog

import (
   "log"
)

var DebugLevel int = 3

///
/// 0 = No Logging
/// 1 = Errors only.
/// 2 = Warnings and Errors.
/// 3 = All info statments.
/// 4 = Debugging.
///

const (
   LnoLogging = iota
   Lerror
   Lwarning
   Linfo
   Ldebug
)

func init () () {
   log.SetFlags(log.Lshortfile)
}

func Error (v ...interface{}) () {
   if !(DebugLevel >= Lerror) { return }
	b := []interface{}{"[ERR] "}
	b = append(b, v...)
   log.Print(b...)
}

func Errorln (v ...interface{}) () {
   if !(DebugLevel >= Lerror) { return }
	b := []interface{}{"[ERR]"}
	b = append(b, v...)
   log.Println(b...)
}

func Errorf (format string, v ...interface{}) () {
   if !(DebugLevel >= Lerror) { return }
	b := []interface{}{"[ERR]"}
	b = append(b, v...)
   log.Printf(format, b...)
}

func Warning (v ...interface{}) () {
   if !(DebugLevel >= Lwarning) { return }
	b := []interface{}{"[WARN] "}
	b = append(b, v...)
   log.Print(b...)
}

func Warningln (v ...interface{}) () {
	if !(DebugLevel >= Lwarning) { return }
	b := []interface{}{"[WARN]"}
	b = append(b, v...)
   log.Println(b...)
}

func Warningf (format string, v ...interface{}) () {
	if !(DebugLevel >= Lwarning) { return }
	b := []interface{}{"[WARN]"}
	b = append(b, v...)
   log.Printf(format, b...)
}

func Info (v ...interface{}) () {
   if !(DebugLevel >= Linfo) { return }
	b := []interface{}{"[INFO] "}
	b = append(b, v...)
   log.Print(b...)
}

func Infoln (v ...interface{}) () {
   if !(DebugLevel >= Linfo) { return }
	b := []interface{}{"[INFO]"}
	b = append(b, v...)
   log.Println(b...)
}

func Infof (format string, v ...interface{}) () {
   if !(DebugLevel >= Linfo) { return }
	b := []interface{}{"[INFO]"}
	b = append(b, v...)
   log.Printf(format, b...)
}

func Debug (v ...interface{}) () {
   if !(DebugLevel >= Ldebug) { return }
	b := []interface{}{"[DEBUG] "}
	b = append(b, v...)
   log.Print(b...)
}

func Debugln (v ...interface{}) () {
   if !(DebugLevel >= Ldebug) { return }
	b := []interface{}{"[DEBUG]"}
	b = append(b, v...)
   log.Println(b...)
}

func Debugf (format string, v ...interface{}) () {
   if !(DebugLevel >= Ldebug) { return }
	b := []interface{}{"[DEBUG]"}
	b = append(b, v...)
   log.Printf(format, b...)
}

func Log (v ...interface{}) () {
   log.Print(v...)
}

func Logln (v ...interface{}) () {
   log.Println(v...)
}

func Logf (format string, v ...interface{}) () {
   log.Printf(format, v...)
}
