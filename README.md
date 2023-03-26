# Logger

## Build Tags

```bash
# Debug
    go build -tags debug

# Release
    go build
    # or
    go build -tags !debug
```

## Import 

```golang
import "github.com/virzz/logger"

// Debug just for -tags debug
logger.Debug("Debug")
logger.DebugF("DebugF")

logger.Success("Success")
logger.SuccessF("SuccessF")
logger.Error("Error")
logger.ErrorF("ErrorF")
logger.Warn("Warn")
logger.WarnF("WarnF")
logger.Info("Info")
logger.InfoF("InfoF")
logger.Normal("Normal")
logger.NormalF("NormalF")

```