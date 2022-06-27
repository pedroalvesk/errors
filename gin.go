package GoErrors

import "github.com/gin-gonic/gin"

func GinMiddleware(messages map[int]string) gin.HandlerFunc {
	cfg := newDefaultConfig()
	if messages != nil {
		cfg.messages = messages
	}

	return func(context *gin.Context) {
		context.Next()

		if context.Errors == nil {
			return
		}

		code := getCode(context.Errors.Last())
		if code == -1 {
			code = cfg.defaultError
		}

		context.AbortWithStatusJSON(code, gin.H{cfg.returnKey: cfg.messages[code]})
	}
}
