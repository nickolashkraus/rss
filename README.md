# rss

TODO: Use pointer to string types, not types that are pointers to strings.

A Go package for marshaling, unmarshaling, and validating RSS documents.

This package aims to codify the RSS specification in code, making it easier to generate and validate RSS documents.

## Overview

This package differs from other implementations of the RSS specification, since it distinguishes between non-existent (`nil`) and empty values (`""`).

When marshaling XML using the [encoding/xml](https://pkg.go.dev/encoding/xml) package, non-existent elements or attributes are omitted using the `omitempty` field tag:

>The XML element for a struct contains marshaled elements for each of the exported fields of the struct, with these exceptions:
>  • a field with a tag including the "omitempty" option is omitted if the field value is empty. The empty values are false, 0, any nil pointer or interface value, and any array, slice, map, or string of length zero.

However, when unmarshaling XML a missing element or empty attribute value will be unmarshaled as a zero value. This can cause a valid RSS document in XML form to become invalid when represented as a Go struct.

To handle this issue, pointers are used, such that `nil` corresponds to a non-existent element or attribute and `""` corresponds to an element or attribute with an empty value.

## RSS Validation

rss is used by [RSS Validator](https://github.com/nickolashkraus/rss-validator), a web application and command-line utility for validating RSS documents.


