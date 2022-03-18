package logs

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "time",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg: "message",
		},
	})

	if isDevEnv, _ := strconv.ParseBool(os.Getenv("LOCAL_ENV")); isDevEnv {
		logrus.SetFormatter(&prefixed.TextFormatter{
			ForceFormatting: true,
		})
	}
}