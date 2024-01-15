package fieldslog

import (
	"runtime"

	"github.com/sirupsen/logrus"
)

func log(subject FieldsDefaulter, ff ...map[string]any) map[string]any {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	_, file, _, _ := runtime.Caller(0)
	r := map[string]any{
		"metadata": map[string]any{
			"file": file,
		},
	}

	fields := MergeMaps(r, MergeMaps(ff...))

	if subject != nil {
		fields = MergeMaps(fields, subject.DefaultLogFields())
	}

	return fields
}

func Error(subject FieldsDefaulter, text string, err error, ff ...map[string]any) {
	fields := MergeMaps(MergeMaps(ff...), logrus.Fields{"error": err})

	logrus.WithFields(log(subject, fields)).Error(text)
}

func Info(subject FieldsDefaulter, info string, ff ...map[string]any) {
	logrus.WithFields(log(subject, ff...)).Info(info)
}

func Warning(subject FieldsDefaulter, info string, ff ...map[string]any) {
	logrus.WithFields(log(subject, ff...)).Warning(info)
}

type FieldsDefaulter interface {
	DefaultLogFields() map[string]any
}
