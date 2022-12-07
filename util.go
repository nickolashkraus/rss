// Copyright 2022 Nickolas Kraus. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Utility functions for the rss package.
package rss

import (
	"fmt"
	"net/mail"
	"net/url"
	"time"
)

// Whether 's' is not an empty string.
func IsNotEmpty(s string) (bool, error) {
	if s == "" {
		return false, fmt.Errorf("%w", ErrEmptyValue)
	}
	return true, nil
}

// Whether 's' is an empty string.
func IsEmpty(s string) (bool, error) {
	if s != "" {
		return false, fmt.Errorf("%w", ErrNonEmptyValue)
	}
	return true, nil
}

// Whether 's' is a valid date (RFC822).
//
// TODO: Valiate day of week.
func IsValidDate(s string) (bool, error) {
	if _, err := time.Parse(time.RFC822, s); err != nil {
		if _, err := time.Parse(time.RFC1123, s); err != nil {
			return false, fmt.Errorf("%w: %v", ErrInvalidDate, err)
		}
	}
	return true, nil
}

// Whether 's' is a valid mail address (RFC5322).
func IsValidMailAddress(s string) (bool, error) {
	if _, err := mail.ParseAddress(s); err != nil {
		return false, fmt.Errorf("%w: %v", ErrInvalidMailAddress, err)
	}
	return true, nil
}

// Whether 's' is a valid URI (RFC3986).
func IsValidURI(s string) (bool, error) {
	if _, err := url.ParseRequestURI(s); err != nil {
		return false, fmt.Errorf("%w: %v", ErrInvalidURI, err)
	}
	return true, nil
}
