package GoErrors

import "net/http"

type config struct {
	defaultError int
	returnKey    string
	messages     map[int]string
}

func newDefaultConfig() config {
	return config{
		defaultError: http.StatusInternalServerError,
		returnKey:    "message",
		messages: map[int]string{
			http.StatusBadRequest:           "the server could not understand the request with the data that was given",
			http.StatusUnauthorized:         "unauthorized action",
			http.StatusForbidden:            "forbidden action",
			http.StatusNotFound:             "the requested entity was not found",
			http.StatusConflict:             "the requested entity conflicts with existing one",
			http.StatusUnsupportedMediaType: "the requested media type is not supported",
			http.StatusInternalServerError:  "a error occurred while working on the request",
		},
	}
}
