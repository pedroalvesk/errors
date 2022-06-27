package errors

import "net/http"

type Config struct {
	DefaultError      int
	MessageKey        string
	Messages          map[int]string
	MessageNotDefined string
}

func newDefaultConfig() *Config {
	return &Config{
		DefaultError: http.StatusInternalServerError,
		MessageKey:   "message",
		Messages: map[int]string{
			http.StatusBadRequest:           "the server could not understand the request with the data that was given",
			http.StatusUnauthorized:         "unauthorized action",
			http.StatusForbidden:            "forbidden action",
			http.StatusNotFound:             "the requested entity was not found",
			http.StatusConflict:             "the requested entity conflicts with existing one",
			http.StatusUnsupportedMediaType: "the requested media type is not supported",
			http.StatusInternalServerError:  "a error occurred while working on the request",
		},
		MessageNotDefined: "message for this error code is not defined yet",
	}
}
