package errors

import "github.com/gin-gonic/gin"

func GinMiddleware(config *Config) gin.HandlerFunc {
	cfg := config
	if cfg == nil {
		cfg = newDefaultConfig()
	}

	return func(context *gin.Context) {
		context.Next()

		if context.Errors == nil {
			return
		}

		var code int
		err, ok := context.Errors.Last().Err.(*Error)
		if !ok {
			code = cfg.DefaultError
		}

		code = err.code
		msg, ok := cfg.Messages[code]
		if !ok {
			msg = cfg.MessageNotDefined
		}

		context.AbortWithStatusJSON(code, gin.H{cfg.MessageKey: msg})
	}
}
