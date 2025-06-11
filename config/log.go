package config

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Log struct {
	CallerDeep int           `toml:"caller_deep" json:"caller_deep" yaml:"caller_deep" env:"CALLER_DEEP"`
	Level      zerolog.Level `toml:"level" json:"level" yaml:"level" env:"LEVEL"`
	TraceFiled string        `toml:"trace_filed" json:"trace_filed" yaml:"trace_filed" env:"TRACE_FILED"`
	Console    Console       `toml:"console" json:"console" yaml:"console" envPrefix:"CONSOLE"`
	File       File          `toml:"file" json:"file" yaml:"file" envPrefix:"FILE_"`

	lock sync.Mutex
	root *zerolog.Logger
}

type Console struct {
	Enable  bool `toml:"enable" json:"enable" yaml:"enable"`
	NoColor bool `toml:"no_color" json:"no_color" yaml:"no_color" env:"NO_COLOR"`
}

type File struct {
	Enable     bool   `toml:"enable" json:"enable" yaml:"enable"  env:"ENABLE"`
	FilePath   string `toml:"file_path" json:"file_path" yaml:"file_path"  env:"PATH"`
	MaxSize    int    `toml:"max_size" json:"max_size" yaml:"max_size"  env:"MAX_SIZE"`
	MaxBackups int    `toml:"max_backups" json:"max_backups" yaml:"max_backups"  env:"MAX_BACKUPS"`
	MaxAge     int    `toml:"max_age" json:"max_age" yaml:"max_age"  env:"MAX_AGE"`
	Compress   bool   `toml:"compress" json:"compress" yaml:"compress"  env:"COMPRESS"`
}

func (c *Console) ConsoleWriter() io.Writer {
	output := zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.NoColor = c.NoColor
		w.TimeFormat = time.RFC3339
	})

	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%-6s", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}
	return output
}

func (f *File) FileWriter() io.Writer {
	return &lumberjack.Logger{
		Filename:   f.FilePath,
		MaxSize:    f.MaxSize,
		MaxAge:     f.MaxAge,
		MaxBackups: f.MaxBackups,
		Compress:   f.Compress,
	}
}

func (l *Log) Logger() *zerolog.Logger {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.root == nil {
		var writers []io.Writer
		if l.Console.Enable {
			writers = append(writers, l.Console.ConsoleWriter())
		}
		if l.File.Enable {
			if l.File.FilePath == "" {
				l.File.FilePath = "logs/app.log"
			}
			writers = append(writers, l.File.FileWriter())
		}

		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

		root := zerolog.New(io.MultiWriter(writers...)).With().Timestamp().Any("Service", "book-api")
		if l.CallerDeep > 0 {
			root = root.Caller()
			zerolog.CallerMarshalFunc = l.CallerMarshalFunc
		}
		l.SetRoot(root.Logger().Level(l.Level))
	}

	return l.root
}

func (m *Log) SetRoot(v zerolog.Logger) {
	m.root = &v
}

func (m *Log) CallerMarshalFunc(pc uintptr, file string, line int) string {
	if m.CallerDeep == 0 {
		return file
	}

	short := file
	count := 0
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			count++
		}

		if count >= m.CallerDeep {
			break
		}
	}
	file = short
	return file + ":" + strconv.Itoa(line)
}
