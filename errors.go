// Copyright 2022 Nickolas Kraus. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Errors for the rss package.
package rss

import "errors"

var ErrEmptyValue = errors.New("Element must not have empty value")
var ErrNonEmptyValue = errors.New("Element must not have value")
var ErrInvalidElement = errors.New("Element must contain required sub-elements")
var ErrInvalidValue = errors.New("Element must have valid value")
var ErrInvalidDate = errors.New("Element must contain a valid date (RFC822)")
var ErrInvalidMailAddress = errors.New("Element must contain a valid mail address (RFC5322)")
var ErrInvalidURI = errors.New("Element must contain a valid URI (RFC3986)")
