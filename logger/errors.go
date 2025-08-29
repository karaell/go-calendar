package logger

import "errors"

var ErrInitLogger = errors.New("error init logger")
var ErrSaveInfoLog = errors.New("error info log save")
var ErrSaveErrorLog = errors.New("error error log save")
var ErrSaveWarningLog = errors.New("error warning log save")
