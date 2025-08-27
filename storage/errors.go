package storage

import "errors"

var ErrSaveStorage = errors.New("error saving storage")
var ErrLoadStorage = errors.New("error loading storage")
var ErrCloseFile = errors.New("error closing storage file")
var ErrCloseZip = errors.New("error closing storage zip")
var ErrEmptyZip = errors.New("storage zip is empty")
var ErrEmptyFile = errors.New("storage file is empty")
