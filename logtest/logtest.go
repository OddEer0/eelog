package logtest

import (
	"context"
	"sync"

	"github.com/OddEer0/eelog"
)

var _ eelog.Logger = (*LogTest)(nil)

const (
	DefaultCap = 8
)

type logKey struct{}
type LogTest struct {
	mu         *sync.Mutex
	level      eelog.Level
	withFields []eelog.Field
	levels     []eelog.Level
	messages   []string
	fields     [][]eelog.Field
}

func NewLogTest(level eelog.Level) *LogTest {
	return &LogTest{
		mu:         &sync.Mutex{},
		level:      level,
		withFields: make([]eelog.Field, 0),
		levels:     make([]eelog.Level, 0, DefaultCap),
		messages:   make([]string, 0, DefaultCap),
		fields:     make([][]eelog.Field, 0, DefaultCap),
	}
}

func (l *LogTest) Level() eelog.Level {
	return l.level
}
func (l *LogTest) Levels() []eelog.Level {
	return l.levels
}

func (l *LogTest) Messages() []string {
	return l.messages
}

func (l *LogTest) Fields() [][]eelog.Field {
	return l.fields
}

func (l *LogTest) WithFields() []eelog.Field {
	return l.withFields
}

func (l *LogTest) Log(_ context.Context, level eelog.Level, message string, fields ...eelog.Field) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.levels = append(l.levels, level)
	l.messages = append(l.messages, message)
	l.fields = append(l.fields, fields)
}

func (l *LogTest) Debug(ctx context.Context, message string, fields ...eelog.Field) {
	l.Log(ctx, eelog.DebugLvl, message, fields...)
}

func (l *LogTest) Info(ctx context.Context, message string, fields ...eelog.Field) {
	l.Log(ctx, eelog.InfoLvl, message, fields...)
}

func (l *LogTest) Warn(ctx context.Context, message string, fields ...eelog.Field) {
	l.Log(ctx, eelog.WarnLvl, message, fields...)
}

func (l *LogTest) Error(ctx context.Context, message string, fields ...eelog.Field) {
	l.Log(ctx, eelog.ErrorLvl, message, fields...)
}

func (l *LogTest) With(fields ...eelog.Field) eelog.Logger {
	l.mu.Lock()
	defer l.mu.Unlock()

	withFields := make([]eelog.Field, 0, len(l.withFields)+len(fields))
	withFields = append(withFields, l.withFields...)
	withFields = append(withFields, fields...)
	return &LogTest{
		mu:         &sync.Mutex{},
		withFields: withFields,
	}
}

func (l *LogTest) InjectCtx(ctx context.Context) context.Context {
	return context.WithValue(ctx, logKey{}, l)
}

func (l *LogTest) Enabled(_ context.Context, level eelog.Level) bool {
	return l.level >= level
}
