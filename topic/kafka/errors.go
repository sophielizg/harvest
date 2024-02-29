package kafka

import "errors"

var ErrNilOption = errors.New("Missing required options")

var ErrNilContextMetadata = errors.New("Got nil metadata value from context")

var ErrWrongContextMetadataType = errors.New("Got wrong type for metadata value from context")
