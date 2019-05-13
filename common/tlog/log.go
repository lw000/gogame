package tlog

import (
	log "github.com/alecthomas/log4go"
)

func init() {

}

func LoadConfiguration(logf string) {
	log.LoadConfiguration(logf)
}

// func Error(arg0 interface{}, args ...interface{})  {
// 	log.Error(arg0, args)
// }
//
// func Info(arg0 interface{}, args ...interface{})  {
// 	log.Info(arg0, args)
// }
//
// func Debug(arg0 interface{}, args ...interface{})  {
// 	log.Debug(arg0, args)
// }
//
// func Trace(arg0 interface{}, args ...interface{})  {
// 	log.Trace(arg0, args)
// }
//
// type Error11 log.struct{
//
// }
