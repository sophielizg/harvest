package request

import "errors"

var InvalidCookieItemFormatError = errors.New("Item returned from cookies cache has invalid format, cannot parse")

var InvalidCookieJarTypeError = errors.New("Cannot save cookies from non remote cookie jar")

var BadStatusCodeError = errors.New("Recieved non-200s status code from server")
