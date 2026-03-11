package eelog

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoopLogger(t *testing.T) {
	ctx := context.Background()
	lvl := DebugLvl
	msg := "TestNoopLogger"

	noop := NoopLogger{}
	noop.Log(ctx, lvl, msg)
	noop.Debug(ctx, msg)
	noop.Info(ctx, msg)
	noop.Warn(ctx, msg)
	noop.Error(ctx, msg)
	assert.False(t, noop.Enabled(ctx, lvl))
	assert.Equal(t, noop, noop.With())
	assert.NotNil(t, noop.InjectCtx(ctx))
}

func TestOutDump(t *testing.T) {
	dump := NewOutDump()
	writeByte := []byte("hello world")
	writeByteLen := len(writeByte)

	n, err := dump.Write(writeByte)
	assert.Equal(t, writeByteLen, n)
	assert.NoError(t, err)

	assert.Equal(t, writeByte, dump.Dump)
}

func TestOutMultiDump(t *testing.T) {
	dump := NewOutMultiDump()
	writeByte := []byte("hello world")
	writeByteLen := len(writeByte)
	writeByte2 := []byte("how are you?")
	writeByteLen2 := len(writeByte2)

	n, err := dump.Write(writeByte)
	assert.Equal(t, writeByteLen, n)
	assert.NoError(t, err)
	n, err = dump.Write(writeByte2)
	assert.Equal(t, writeByteLen2, n)
	assert.NoError(t, err)

	assert.Equal(t, [][]byte{writeByte, writeByte2}, dump.Dumps)

	n, err = dump.Write(writeByte2)
	assert.Equal(t, writeByteLen2, n)
	assert.NoError(t, err)

	assert.Equal(t, [][]byte{writeByte, writeByte2, writeByte2}, dump.Dumps)
}
