package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

const (
	DEBUG_LEVEL = "debug"
	ERROR_LEVEL = "error"
	WARN_LEVEL  = "warn"

	FILE    = "file"
	CONSOLE = "console"

	FILESINK = "wf"
)

var (
	myLogger *Logger
	encoderConfig zapcore.EncoderConfig
	logCfg        LogConfig
)

type Logger struct {
	name        string
	development bool
	core        zapcore.Core
	errorOutput zapcore.WriteSyncer
	addCaller   bool
	addStack    zapcore.LevelEnabler
	callerSkip  int
}

type LogConfig struct {
	Dev      bool              `json:"dev" yaml:"dev"`
	Level    string            `json:"level" yaml:"level"`
	Encoding string            `json:"encoding" yaml:"encoding"`
	Encode   map[string]string `json:"encode" yaml:"encode"`
	Key      map[string]string `json:"key" yaml:"key"`
	OutPuts  []string          `json:"outputs" yaml:"outputs"`
	LogPath  string            `json:"path" yaml:"path"`
	LogFile  string            `json:"file" yaml:"file"`
	Format   string            `json:"format" yaml:"format"`
}

func initmyLogger(logConf *LogConfig) {
	encoderConfig.NameKey = logConf.Key["name"]
	encoderConfig.TimeKey = logConf.Key["time"]
	encoderConfig.LevelKey = logConf.Key["level"]
	encoderConfig.CallerKey = logConf.Key["caller"]
	encoderConfig.MessageKey = logConf.Key["message"]
	encoderConfig.StacktraceKey = logConf.Key["stacktrace"]

	encoderConfig.LineEnding = zapcore.DefaultLineEnding

	logConf.timeEncoder()
	logConf.lvlEncoder()
	logConf.durEncoder()
	logConf.callerEncoder()

	var lvl zap.AtomicLevel
	switch logConf.Level {
	case DEBUG_LEVEL:
		lvl = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case WARN_LEVEL:
		lvl = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case ERROR_LEVEL:
		lvl = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	default:
		lvl = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}

	var outputs []string
	for _, p := range logConf.OutPuts {
		if p == FILE {
			_ = zap.RegisterSink(FILESINK, func(u *url.URL) (zap.Sink, error) {
				return os.OpenFile(u.Path[1:], os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
			})
			filename := fmt.Sprintf("%s_%s.log", logConf.LogFile, time.Now().Format(logConf.Format))
			//增加 如果日志目录不存在则建立该目录
			_, err := os.Stat(logConf.LogPath)
			if err != nil {
				err := os.Mkdir(logConf.LogPath, os.ModePerm)
				if err != nil {
					fmt.Printf("mkdir failed![%v]\n", err)
				} else {
					fmt.Printf("mkdir success!\n")
				}
			}

			logFile := fmt.Sprintf("%s:///%s", FILESINK, filepath.Join(logConf.LogPath, filename))
			outputs = append(outputs, logFile)
		} else {
			outputs = append(outputs, p)
		}
	}

	sink, close, err := zap.Open(outputs...)
	if err != nil {
		close()
	}

	var cores []zapcore.Core
	switch logConf.Encoding {
	case CONSOLE:
		consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
		cores = append(cores, zapcore.NewCore(consoleEncoder, sink, lvl))
	case "json":
		jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)
		cores = append(cores, zapcore.NewCore(jsonEncoder, sink, lvl))
	}

	myLogger = &Logger{
		name:        logConf.LogFile,
		core:        zapcore.NewTee(cores...),
		development: logConf.Dev,
		errorOutput: zapcore.Lock(os.Stderr),
		addStack:    zapcore.FatalLevel + 1,
		addCaller:   true,
	}
}

func (jlc *LogConfig) lvlEncoder() {
	lvl := jlc.Encode["level"]
	switch lvl {
	case "capital":
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	case "capitalColor":
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	case "color":
		encoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	default:
		encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
}

func (jlc *LogConfig) timeEncoder() {
	encTime := jlc.Encode["time"]
	switch encTime {
	case "iso8601", "ISO8601":
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	case "millis":
		encoderConfig.EncodeTime = zapcore.EpochMillisTimeEncoder
	case "nanos":
		encoderConfig.EncodeTime = zapcore.EpochNanosTimeEncoder
	case "localtime":
		encoderConfig.EncodeTime = zapcore.EpochTimeEncoder
	case "unix":
		encoderConfig.EncodeTime = func(i time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(i.Local().Unix())
		}
	default:
		encoderConfig.EncodeTime = func(i time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(i.Format("2006-01-02 15:04:05"))
		}
	}
}

func (jlc *LogConfig) durEncoder() {
	dur := jlc.Encode["duration"]
	switch dur {
	case "string":
		encoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	case "nanos":
		encoderConfig.EncodeDuration = zapcore.NanosDurationEncoder
	default:
		encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	}
}

func (jlc *LogConfig) callerEncoder() {
	caller := jlc.Encode["caller"]
	switch caller {
	case "full":
		encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	default:
		encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	}
}

func (jl *Logger) check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry {
	const callerSkipOffset = 2
	if myLogger == nil {
		initmyLogger(nil)
	}
	ent := zapcore.Entry{
		LoggerName: myLogger.name,
		Time:       time.Now().Local(),
		Level:      lvl,
		Message:    msg,
	}
	ce := myLogger.core.Check(ent, nil)
	willWrite := ce != nil

	switch ent.Level {
	case zapcore.PanicLevel:
		ce = ce.Should(ent, zapcore.WriteThenPanic)
	case zapcore.FatalLevel:
		ce = ce.Should(ent, zapcore.WriteThenFatal)
	case zapcore.DPanicLevel:
		if myLogger.development {
			ce = ce.Should(ent, zapcore.WriteThenPanic)
		}
	}

	if !willWrite {
		return ce
	}

	ce.ErrorOutput = myLogger.errorOutput
	if myLogger.addCaller {
		ce.Entry.Caller = zapcore.NewEntryCaller(runtime.Caller(myLogger.callerSkip + callerSkipOffset))
		if !ce.Entry.Caller.Defined {
			_,_ = fmt.Fprintf(myLogger.errorOutput, "%v Logger.check error: failed to get caller\n", time.Now().Local())
			_ = myLogger.errorOutput.Sync()
		}
	}
	if myLogger.addStack.Enabled(ce.Entry.Level) {
		ce.Entry.Stack = zap.Stack("").String
	}

	return ce
}

func Debug(details ...interface{}) {
	if ce := myLogger.check(zapcore.DebugLevel, fmt.Sprint(details...)); ce != nil {
		ce.Write()
	}
}
func Info(details ...interface{}) {
	if ce := myLogger.check(zapcore.InfoLevel, fmt.Sprint(details...)); ce != nil {
		ce.Write()
	}
}
func CInfo(msg string, fields map[string]interface{}) {
	if len(fields) > 0 {
		if ce := myLogger.check(zapcore.InfoLevel, msg); ce != nil {
			ce.Write(genFields(fields)...)
		}
	} else {
		Info(msg)
	}
}
func Warn(details ...interface{}) {
	if ce := myLogger.check(zapcore.WarnLevel, fmt.Sprint(details...)); ce != nil {
		ce.Write()
	}
}
func Error(details ...interface{}) {
	if ce := myLogger.check(zapcore.ErrorLevel, fmt.Sprint(details...)); ce != nil {
		ce.Write()
	}
}
func DPanic(details ...interface{}) {
	if ce := myLogger.check(zapcore.DPanicLevel, fmt.Sprint(details...)); ce != nil {
		ce.Write()
	}
}
func Panic(details ...interface{}) {
	if ce := myLogger.check(zapcore.PanicLevel, fmt.Sprint(details...)); ce != nil {
		ce.Write()
	}
}
func Fatal(details ...interface{}) {
	if ce := myLogger.check(zapcore.FatalLevel, fmt.Sprint(details...)); ce != nil {
		ce.Write()
	}
}
func Sync() {
	_ = myLogger.core.Sync()
}

func genFields(details map[string]interface{}) []zapcore.Field {
	var fields = make([]zapcore.Field, 0)
	for k, v := range details {
		switch v.(type) {
		case bool:
			fields = append(fields, zap.Bool(k, v.(bool)))
		case int8:
			fields = append(fields, zap.Int8(k, v.(int8)))
		case int, int32:
			fields = append(fields, zap.Int(k, v.(int)))
		case uint, uint32:
			fields = append(fields, zap.Uint(k, v.(uint)))
		case int64:
			fields = append(fields, zap.Int64(k, v.(int64)))
		case string:
			fields = append(fields, zap.String(k, v.(string)))
		default:
			fields = append(fields, zap.Reflect(k, v))
		}
	}
	return fields
}

func Init(file string) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("文件读取失败")
	}
	err = yaml.Unmarshal(buf, &logCfg)
	if err != nil {
		fmt.Println(file + "解析失败")
	}
	initmyLogger(&logCfg)
}