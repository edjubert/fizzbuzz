package utils

import "github.com/gookit/slog"

func ConfigureLogger() {
	slog.DefaultChannelName = "fizzbuzz"
	slog.DefaultTimeFormat = "02-01-2006 15:04:05"

	template := slog.DefaultTemplate

	slog.Configure(func(logger *slog.SugaredLogger) {
		f := logger.Formatter.(*slog.TextFormatter)
		f.SetTemplate(template)
		f.EnableColor = true
	})

	lvls := []string{"PANIC", "FATAL", "ERROR", "WARN", "NOTICE", "INFO", "DEBUG", "TRACE"}
	envLvl := GetEnvWithDefault("DEBUG_LEVEL", "INFO")
	envIdx := (func(arr []string, str string) int {
		for i, a := range arr {
			if a == str {
				return i
			}
		}
		return -1
	}(lvls, envLvl) + 1) * 100

	slog.SetLogLevel(slog.Level(envIdx))
}
