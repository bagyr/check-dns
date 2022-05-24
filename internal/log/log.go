package log

import "go.uber.org/zap"

var S *zap.SugaredLogger

func init() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	S = logger.Sugar()
}
