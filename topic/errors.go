package topic

import "errors"

var ErrNilSessionAck = errors.New("Cannot ack message result due to missing session")

var ErrNilOption = errors.New("Missing required options")
