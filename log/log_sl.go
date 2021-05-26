package log

import (
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
			//RotateLogPath:logPath,
			RotateLogPath:logPath+".%Y%M%D",
			WfLogPath:wfLogPath,
			//RotateWfLogPath:wfLogPath,
			RotateWfLogPath:wfLogPath+".%Y%M%D",
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
		logPath := path.Join(logRootDIR,subdirectories[i],"/access.log")
		_subdirectories := subdirectories[i]
		wflogPath := path.Join(logRootDIR,_subdirectories,"/error.log")
		createLogs(logPath,wflogPath,_subdirectories)
	}
}

func Instance(name string) *Logger {
	return LogIchunt[name]
}


