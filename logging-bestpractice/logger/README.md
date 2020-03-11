# golibs-logger

プラットフォーム共通ログライブラリ.

## Usage
仕組みに関しての詳細は以下を参照

https://rarejob.atlassian.net/wiki/spaces/RP/pages/724633353


### Simple Console Log
標準エラーに出力される。  
何も設定しない。

```go
package main
    
import (
	"errors"
    
	"git.rarejob.com/rarejob-platform/golibs/logger"
)
    
func main() {
    // debug
    logger.Debug("debug message")
    // info
    logger.Info("info message")
    // warning
    logger.Warn("warning message")
    // error
    logger.Error("error message",
        logger.E(errors.New("validation error")),
    )
}
```

### Auto Rotate File Log
指定したファイルに出力される。

```go
package main
    
import (
	"errors"
    
	"git.rarejob.com/rarejob-platform/golibs/logger"
	"git.rarejob.com/rarejob-platform/golibs/logger/fio"
)
    
func main() {
    logger.SetupRootLogger(logger.NewConfig("app",
              logger.WithLevel(logger.Levels.Debug),
              logger.WithOut(fio.NewFileWriter(
                  fio.NewFileConfig("app.log",      // ログファイル名
                      fio.WithLogDirectory("logs"), // ログ出力ディレクトリ
                      fio.WithMaxSize(5),           // 最大ログサイズ(mb)
                      fio.WithMaxBackups(10),       // 最大バックアップ数
                      fio.WithMaxAge(30),           // 最大保持期間(day)
                  ),
              )),
          ))
    defer logger.Close()
    
    // debug
    logger.Debug("debug message")
    // info
    logger.Info("info message")
    // warning
    logger.Warn("warning message")
    // error
    logger.Error("error message",
        logger.E(errors.New("validation error")),
    )
}
```

### Auto Daily Rotate File Log
指定したファイルに出力される。  
cronで毎日ファイル削除する。

```go
package main
    
import (
	"errors"
    
	"git.rarejob.com/rarejob-platform/golibs/logger"
	"git.rarejob.com/rarejob-platform/golibs/logger/fio"
)
    
func main() {
    logger.SetupRootLogger(logger.NewConfig("app",
        logger.WithLevel(logger.Levels.Debug),
        logger.WithMaxLevel(logger.Levels.Fatal),
        logger.WithFormatter(logger.Formatters.Text),
        logger.WithOut(fio.NewFileWriter(
            fio.NewFileConfig("app.log",
                fio.WithLogDirectory("logs"),
                fio.WithMaxSize(5),
                fio.WithMaxBackups(10),
                fio.WithMaxAge(30),
            ),
        )),
    ))
    logger.StartCronTimer("0 0 0 * * *") // daily log rotation
    defer logger.StopCronTimer()
    
    // debug
    logger.Debug("debug message")
    // info
    logger.Info("info message")
    // warning
    logger.Warn("warning message")
    // error
    logger.Error("error message",
        logger.E(errors.New("validation error")),
    )
}
```

### Separate File Log
エラーの種類毎にファイル出力される。

```go
package main
    
import (
	"errors"
    
	"git.rarejob.com/rarejob-platform/golibs/logger"
	"git.rarejob.com/rarejob-platform/golibs/logger/fio"
)
    
func main() {
	// debug-warning をapp.logに出力
    appLogConfig := logger.NewConfig("app",
        logger.WithLevel(logger.Levels.Debug),
        logger.WithMaxLevel(logger.Levels.Warn),
        logger.WithOut(fio.NewFileWriter(
            fio.NewFileConfig("app.log",
                fio.WithLogDirectory("logs"),
                fio.WithMaxSize(5),
                fio.WithMaxBackups(10),
                fio.WithMaxAge(30),
            ),
        )),
    )
    // error-fatal をerror.logに出力
    errorLogConfig := logger.NewConfig("error",
        logger.WithLevel(logger.Levels.Error),
        logger.WithMaxLevel(logger.Levels.Fatal),
        logger.WithOut(fio.NewFileWriter(
            fio.NewFileConfig("error.log",
                fio.WithLogDirectory("logs"),
                fio.WithMaxSize(5),
                fio.WithMaxBackups(10),
                fio.WithMaxAge(30),
            ),
        )),
    )
    logger.SetupRootLogger(appLogConfig, errorLogConfig)
    defer logger.Close()
    
    // debug   (app.logのみ出力)
    logger.Debug("debug message")
    // info    (app.logのみ出力)
    logger.Info("info message")
    // warning (app.logのみ出力)
    logger.Warn("warning message")
    // error   (error.logのみ出力)
    logger.Error("error message",
        logger.E(errors.New("validation error")),
    )
}
```

### Buffering Log For Lambda Function
バッファを確保して標準エラー出力する。  
バイナリの実行終了後に出力される。  
lambdaで使う。

```go
package main
    
import (
	"errors"
    
	"git.rarejob.com/rarejob-platform/golibs/logger"
	"git.rarejob.com/rarejob-platform/golibs/logger/fio"
)
    
func main() {
    logger.SetupRootLogger(logger.NewConfig("lambda-log",
        logger.WithLevel(logger.Levels.Debug),
        logger.WithOut(fio.NewBufferedWriter()),
    ))
    // closeの時点で自動Flush
    defer logger.Close()
    
    // debug
    logger.Debug("debug message")
    // info
    logger.Info("info message")
    // warning
    logger.Warn("warning message")
    // error
    logger.Error("error message",
        logger.E(errors.New("validation error")),
    )
}
```