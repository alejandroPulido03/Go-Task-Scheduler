package core

import "time"

func ErrorMessage(err error) map[string]string {
	return map[string]string{
		"error": err.Error(),
		"timestamp": time.Now().String(),
	}
}