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
	DebugF("DebugF %s", "test")
	Success("Success")
	SuccessF("SuccessF %s", "test")
	Error("Error")
	ErrorF("ErrorF %s", "test")
	Warn("Warn")
	WarnF("WarnF %s", "test")
	Info("Info")
	InfoF("InfoF %s", "test")
	Normal("Normal")
	NormalF("NormalF %s", "test")
	Fatal("Fatal")
}

func TestLogToFile(t *testing.T) {
	std.SetFileLog("test.log")
	std.Info("test")
}
