/**
* Created by chaolinding on 2018/4/23
*/

package config

import (
	"os"
	"runtime"
)

var env string

func init()  {
	runtime.GOMAXPROCS(runtime.NumCPU())
	sysEnv := os.Getenv("GOENV")
	if len(sysEnv) > 0 && (sysEnv == "production" || sysEnv == "testing" || sysEnv == "development"){
		env =  sysEnv
	}else{
		env = "development"
	}
}

var GetEnv = func() string {
	return env
}