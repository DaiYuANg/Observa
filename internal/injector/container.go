package injector

import (
	"github/DaiYuANg/Observa/modules/logger"
	"github/DaiYuANg/Observa/modules/scheduler"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var commonOption = fx.Options(
	logger.Module,
	scheduler.Module,
	fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
		fxLogger := &fxevent.ZapLogger{Logger: log}
		fxLogger.UseLogLevel(zapcore.DebugLevel)
		return fxLogger
	}),
)

func CreateContainer(modules ...fx.Option) (*fx.App, []fx.Option, error) {
	// 合并 commonOption + modules
	options := append([]fx.Option{commonOption}, modules...)

	// 先校验
	if err := fx.ValidateApp(options...); err != nil {
		return nil, nil, err
	}

	// 创建 App
	app := fx.New(options...)

	return app, options, nil
}
