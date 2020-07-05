package clients

import "errors"

var prefix = "len of field "
var suffixMore0 = " must be more 0"
var suffixEq0 = " must be equal 0"

var CountArgsError = errors.New("field countArgs must be equal length of the len(args)")
var EmptyMethodError = errors.New(prefix + "method" + suffixMore0)
var EmptyPathError = errors.New(prefix + "url" + suffixMore0)
var NotEmptyBodyError = errors.New(prefix + "body" + suffixEq0)

var EmptyPageError = errors.New("empty page error")
