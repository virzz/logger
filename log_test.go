package logger

import (
	"testing"
)

func TestDebug(t *testing.T) {
	Debug("test")
	DebugF("test Debug")
}

func TestSuccess(t *testing.T) {
	Success("test")
	Success("test", "aaa")
	SuccessF("test Success")
}

func TestError(t *testing.T) {
	Error("test")
	ErrorF("test Error")
}

func TestWarn(t *testing.T) {
	Warn("test")
	WarnF("test Warn")
}

func TestInfo(t *testing.T) {
	Info("test")
	InfoF("test Info")
}

func TestNormal(t *testing.T) {
	Normal("test")
	NormalF("test Normal")
}

func TestFatal(t *testing.T) {
	t.Skip()
	Fatal("test")
	FatalF("test Normal")
}

func TestPrefix(t *testing.T) {
	SetPrefix("PREFIX")
	Normal("test")
	NormalF("test Prefix")
	Success("test Prefix Success")
	Debug("Debug ttttttbikasebd")
	SetPrefix("")
}

func TestFlags(t *testing.T) {
	SetDevFlags()
	SetPrefix("PREFIX")
	Normal("test")
	NormalF("test Prefix")
	Success("test Prefix Success")
	Debug("Debug ttttttbikasebd")
	SetPrefix("")
}
func TestAll(t *testing.T) {
	Debug("Debug")
	Success("Success")
	Error("Error")
	Warn("Warn")
	Info("Info")
	Normal("Normal")
	DebugF("DebugF %s", "test")
	SuccessF("SuccessF %s", "test")
	ErrorF("ErrorF %s", "test")
	WarnF("WarnF %s", "test")
	InfoF("InfoF %s", "test")
	NormalF("NormalF %s", "test")
}

func TestFile(t *testing.T) {
	l := NewLogger(LoggerFile, "test.log")
	l.Error("NewLogger(LoggerPrint|LoggerFile testtttttttttttt")
	l.Success("TestFile Success")
	l.Error("TestFile Error")
	l.Warn("TestFile Warn")
	l.Info("TestFile Info")
	l.Normal("TestFile Normal")
	l.Debug("TestFile Debug")
}

func TestBothFile(t *testing.T) {
	l := NewLogger(LoggerPrint|LoggerFile, "both_test.log")
	l.Success("TestBothFile Success")
	l.Error("TestBothFile Error")
	l.Warn("TestBothFile Warn")
	l.Info("TestBothFile Info")
	l.Normal("TestBothFile Normal")
	l.Debug("TestBothFile Debug")
}

func TestSetFile(t *testing.T) {
	SetFile("set_test.log")
	Success("TestSetFile Success")
	Error("TestSetFile Error")
	Warn("TestSetFile Warn")
	Info("TestSetFile Info")
	Normal("TestSetFile Normal")
	Debug("TestSetFile Debug")
	NoFile()
	Success("NoFile Success")
	Error("NoFile Error")
	Warn("NoFile Warn")
	Info("NoFile Info")
	Normal("NoFile Normal")
	Debug("NoFile Debug")
}
