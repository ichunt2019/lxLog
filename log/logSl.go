package log

import (
	"fmt"
	"path"
)

var LogIchunt map[string]*Logger



func createLogs(logPath string,wfLogPath string,subdirectories string){
	nlog:= NewLogger()
	logConf:= LogConfig{
		Level:"trace",
		FW: ConfFileWriter{
			On:true,
			LogPath:logPath,
			RotateLogPath:logPath,
			WfLogPath:wfLogPath,
			RotateWfLogPath:wfLogPath,
		},
		CW: ConfConsoleWriter{
			On:true,
			Color:true,
		},
	}
	SetupLogInstanceWithConf(logConf,nlog)
	if LogIchunt == nil{
		LogIchunt = make(map[string]*Logger,0)
	}
	LogIchunt[subdirectories] = nlog
}

func Init(logRootDIR string,subdirectories ...string){
	for i:=0;i<len(subdirectories);i++{
		logPath := path.Join(logRootDIR,subdirectories[i],"/nolog.log")
		_subdirectories := subdirectories[i]
		wflogPath := path.Join(logRootDIR,_subdirectories,"/wflog.log")
		fmt.Println(logPath)
		fmt.Println(wflogPath)
		if len(subdirectories) == 2 {
			logPath = subdirectories[0]
			wflogPath = subdirectories[1]
		}
		createLogs(logPath,wflogPath,_subdirectories)
	}
}

func Instance(name string) *Logger {
	return LogIchunt[name]
}


