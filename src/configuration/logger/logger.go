package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger

	LOG_OUTPUT = "LOG_OUTPUT"
	LOG_LEVEL  = "LOG_LEVEL"
	INIT       = "init"
	DONE       = "done"
	PROGRESS   = "progress"
)

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{getOutputLogs()},
		Level:       zap.NewAtomicLevelAt(getLevelLogs()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	log, _ = logConfig.Build()
}

func Info(message string, function string, action string, tags ...zap.Field) {
	logging(message, zapcore.InfoLevel, false, nil, zap.String("function", function),
		zap.String("action", action), tags)
}

func Debug(message string, function string, action string, tags ...zap.Field) {
	logging(message, zapcore.DebugLevel, false, nil, zap.String("function", function),
		zap.String("action", action), tags)
}

func Warn(message string, err error, function string, action string, tags ...zap.Field) {
	logging(message, zapcore.WarnLevel, true, err, zap.String("function", function),
		zap.String("action", action), tags)
}

func Error(message string, err error, function string, action string, tags ...zap.Field) {
	logging(message, zapcore.ErrorLevel, true, err, zap.String("function", function),
		zap.String("action", action), tags)
}

func Fatal(message string, err error, function string, action string, tags ...zap.Field) {
	logging(message, zapcore.FatalLevel, true, err, zap.String("function", function),
		zap.String("action", action), tags)
}

func logging(message string, logType zapcore.Level, isError bool, err error, function zapcore.Field, action zapcore.Field,
	tags []zapcore.Field) {
	tags = append(tags, function, action)

	if isError {
		tags = append(tags, zap.NamedError("error", err))
	}

	switch logType {
	case zapcore.InfoLevel:
		log.Info(message, tags...)
	case zapcore.ErrorLevel:
		log.Error(message, tags...)
	case zapcore.DebugLevel:
		log.Debug(message, tags...)
	case zapcore.WarnLevel:
		log.Warn(message, tags...)
	case zapcore.FatalLevel:
		log.Fatal(message, tags...)
	default:
		log.Info(message, tags...)
	}

	errSync := log.Sync()
	if errSync != nil {
		return
	}
}

func getOutputLogs() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_OUTPUT)))
	if output == "" {
		return "stdout"
	}

	return output
}

func getLevelLogs() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(LOG_LEVEL))) {
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	case "debug":
		return zapcore.DebugLevel
	case "warn":
		return zapcore.WarnLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
