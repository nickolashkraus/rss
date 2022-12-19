// Copyright 2022 Nickolas Kraus. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// The rss package provides utilities for handling RSS documents. It leverages
// the RSS 2.0 Specification for determining struct fields used during
// marshaling, unmarshaling, and validation.
//
// See:
//   - https://validator.w3.org/feed/docs/rss2.html
//
// This package relies heavily on the encoding/xml package of the Go standard
// library:
//
// See:
//   - https://pkg.go.dev/encoding/xml
//
// Some notes on parsing XML:
//   - The XMLName field name dictates the name of the XML element representing
//     this struct.
//   - If the XML element contains character data, that data is accumulated in
//     the first struct field that has tag ",chardata". The struct field may
//     have type []byte or string. If there is no such field, the character
//     data is discarded.
//   - A field with a tag including the "omitempty" option is omitted if the
//     field value is empty. The empty values are false, 0, any nil pointer or
//     interface value, and any array, slice, map, or string of length zero.
package rss

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"strconv"
)

const RSSVERSION = "2.0"

// The RSSElement interface specifies a single method, IsValid. IsValid checks
// whether the element conforms to the RSS 2.0 Specification.
type RSSElement interface {
	IsValid() (bool, []error)
}

// In order for an RSS element (RSSElement) to be valid, it must comprise all
// required elements and sub-elements.
//
// If the RSS element contains optional sub-elements with required elements,
// these too must be valid.
//
// To accomplish this, we recurse through all struct fields of the struct of
// interface type RSSElement. If the struct field is of interface type
// RSSElement, the IsValid method is called. Each RSSElement is responsible for
// implementing its IsValid method in accordance with the RSS 2.0
// Specification.
func Validate(r RSSElement) (bool, []error) {
	isValid, errs := true, []error{}
	// ValueOf returns a new Value initialized to the concrete value
	// stored in the interface i. ValueOf(nil) returns the zero Value.
	v := reflect.ValueOf(r)
	// NumField returns the number of fields in the struct v.
	// It panics if v's Kind is not Struct.
	for i := 0; i < v.NumField(); i++ {
		// Field returns the i'th field of the struct v.
		// It panics if v's Kind is not Struct or i is out of range.
		//
		// Interface returns v's current value as an interface{}.
		// It is equivalent to:
		//
		//	var i interface{} = (v's underlying value)
		//
		// It panics if the Value was obtained by accessing
		// unexported struct fields.
		//
		// To test whether an interface value holds a specific type, a type
		// assertion can return two values: the underlying value and a boolean
		// value that reports whether the assertion succeeded.
		//
		//  t, ok := i.(T)
		//
		// If i holds a T, then t will be the underlying value and ok will be true.
		//
		// If not, ok will be false and t will be the zero value of type T, and no
		// panic occurs.
		if t, ok := v.Field(i).Interface().(RSSElement); ok {
			// ValueOf returns a new Value initialized to the concrete value
			// stored in the interface i. ValueOf(nil) returns the zero Value.
			v := reflect.ValueOf(t)
			// Kind returns v's Kind.
			// If v is the zero Value (IsValid returns false), Kind returns Invalid.
			if v.Kind() == reflect.Pointer {
				// Check whether v is nil before calling IsValid.
				if v.IsNil() {
					continue
				}
			}
			if ok, e := t.IsValid(); !ok {
				isValid = false
				errs = append(errs, e...)
			}
		}
	}
	return isValid, errs
}

// At the top level, a RSS document is a <rss> element, with a mandatory
// attribute called version, that specifies the version of RSS that the
// document conforms to. If it conforms to this specification, the version
// attribute must be 2.0.
//
// Subordinate to the <rss> element is a single <channel> element, which
// contains information about the channel (metadata) and its contents.
//
// See:
//   - https://validator.w3.org/feed/docs/rss2.html#whatIsRss
type RSS struct {
	XMLName xml.Name `xml:"rss"`          // required
	Version Version  `xml:"version,attr"` // required
	Channel *Channel `xml:"channel"`      // required
}

// Whether <rss> is valid.
func (r RSS) IsValid() (bool, []error) { return Validate(r) }

// version is a required attribute of <rss>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#whatIsRss
type Version string

// Whether version is valid.
//
// <rss> must contain "version" attribute with value "2.0".
//
// NOTE: A version 0.91 or 0.92 file is also a valid 2.0 file.
func (r Version) IsValid() bool {
	return r == RSSVERSION
}

// <channel> is a required sub-element of <rss>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#requiredChannelElements
type Channel struct {
	XMLName        xml.Name       `xml:"channel"`                  // required
	Title          Title          `xml:"title"`                    // required
	Link           Link           `xml:"link"`                     // required
	Description    Description    `xml:"description"`              // required
	Language       Language       `xml:"language,omitempty"`       // optional
	Copyright      Copyright      `xml:"copyright,omitempty"`      // optional
	ManagingEditor ManagingEditor `xml:"managingEditor,omitempty"` // optional
	WebMaster      WebMaster      `xml:"webMaster,omitempty"`      // optional
	PubDate        PubDate        `xml:"pubDate,omitempty"`        // optional
	LastBuildDate  LastBuildDate  `xml:"lastBuildDate,omitempty"`  // optional
	Category       Category       `xml:"category,omitempty"`       // optional
	Generator      Generator      `xml:"generator,omitempty"`      // optional
	Docs           Docs           `xml:"docs,omitempty"`           // optional
	Cloud          Cloud          `xml:"cloud,omitempty"`          // optional
	TTL            TTL            `xml:"ttl,omitempty"`            // optional
	Image          Image          `xml:"image,omitempty"`          // optional
	Rating         Rating         `xml:"rating,omitempty"`         // optional
	TextInput      TextInput      `xml:"textInput,omitempty"`      // optional
	SkipHours      SkipHours      `xml:"skipHours,omitempty"`      // optional
	SkipDays       SkipDays       `xml:"skipDays,omitempty"`       // optional
	Item           []*Item        `xml:"item,omitempty"`           // optional
}

// Whether <channel> is valid.
//
// In order for <channel> to be valid, it must comprise all required elements
// and sub-elements.
//
// If <channel> contains optional sub-elements with required elements, these
// too must be valid.
//
// To accomplish this, we recurse through all struct fields. If the struct
// field is of interface type RSSElement, the IsValid method is called. Each
// RSSElement is responsible for implementing its IsValid method in accordance
// with the RSS 2.0 Specification.
func (r Channel) IsValid() bool {
	// ValueOf returns a new Value initialized to the concrete value
	// stored in the interface i. ValueOf(nil) returns the zero Value.
	v := reflect.ValueOf(r)
	// NumField returns the number of fields in the struct v.
	// It panics if v's Kind is not Struct.
	for i := 0; i < v.NumField(); i++ {
		// Field returns the i'th field of the struct v.
		// It panics if v's Kind is not Struct or i is out of range.
		//
		// Interface returns v's current value as an interface{}.
		// It is equivalent to:
		//
		//	var i interface{} = (v's underlying value)
		//
		// It panics if the Value was obtained by accessing
		// unexported struct fields.
		//
		// To test whether an interface value holds a specific type, a type
		// assertion can return two values: the underlying value and a boolean
		// value that reports whether the assertion succeeded.
		//
		//  t, ok := i.(T)
		//
		// If i holds a T, then t will be the underlying value and ok will be true.
		//
		// If not, ok will be false and t will be the zero value of type T, and no
		// panic occurs.
		if t, ok := v.Field(i).Interface().(RSSElement); ok {
			// Indirect returns the value that v points to.
			// If v is a nil pointer, Indirect returns a zero Value.
			// If v is not a pointer, Indirect returns v.
			v := reflect.Indirect(reflect.ValueOf(t))
			if v.IsNil() || !v.IsValid() {
				return false
			}
		}
	}
	return true
}

// <title> is a required sub-element of <channel>, <textInput>, and <item>.
//
// See:
//   - https://validator.w3.org/feed/docs/rss2.html#requiredChannelElements
//   - https://validator.w3.org/feed/docs/rss2.html#lttextinputgtSubelementOfLtchannelgt
//   - https://validator.w3.org/feed/docs/rss2.html#hrelementsOfLtitemgt
type Title struct {
	XMLName  xml.Name `xml:"title"`     // required
	CharData []byte   `xml:",chardata"` // required
}

// Returns whether <title> is valid and a slice containing any errors.
func (r Title) IsValid() (bool, []error) {
	isValid, errs := true, []error{}
	msg := fmt.Sprintf("Element <%s> value '%s' is invalid", r.XMLName.Local, r.CharData)
	if ok, err := IsNotEmpty(string(r.CharData)); !ok {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, err))
	}
	return isValid, errs
}

// <link> is a required sub-element of <channel>, <image>, <textInput>, and
// <item>.
//
// See:
//   - https://validator.w3.org/feed/docs/rss2.html#requiredChannelElements
//   - https://validator.w3.org/feed/docs/rss2.html#ltimagegtSubelementOfLtchannelgt
//   - https://validator.w3.org/feed/docs/rss2.html#lttextinputgtSubelementOfLtchannelgt
//   - https://validator.w3.org/feed/docs/rss2.html#hrelementsOfLtitemgt
type Link struct {
	XMLName  xml.Name `xml:"link"`      // required
	CharData []byte   `xml:",chardata"` // required
}

// Returns whether <link> is valid and a slice containing any errors.
func (r Link) IsValid() (bool, []error) {
	isValid, errs := true, []error{}
	msg := fmt.Sprintf("Element <%s> value '%s' is invalid", r.XMLName.Local, r.CharData)
	if ok, err := IsNotEmpty(string(r.CharData)); !ok {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, err))
	}
	if ok, err := IsValidURI(string(r.CharData)); !ok {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, err))
	}
	return isValid, errs
}

// <description> is a required sub-element of <channel> and <textInput> and an
// optional sub-element of <image> and <item>
//
// See:
//   - https://validator.w3.org/feed/docs/rss2.html#requiredChannelElements
//   - https://validator.w3.org/feed/docs/rss2.html#ltimagegtSubelementOfLtchannelgt
//   - https://validator.w3.org/feed/docs/rss2.html#lttextinputgtSubelementOfLtchannelgt
//   - https://validator.w3.org/feed/docs/rss2.html#hrelementsOfLtitemgt
type Description struct {
	XMLName  xml.Name `xml:"description"` // required
	CharData []byte   `xml:",chardata"`   // required
}

// Returns whether <description> is valid and a slice containing any errors.
func (r Description) IsValid() (bool, []error) {
	isValid, errs := true, []error{}
	msg := fmt.Sprintf("Element <%s> value '%s' is invalid", r.XMLName.Local, r.CharData)
	if ok, err := IsNotEmpty(string(r.CharData)); !ok {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, err))
	}
	return isValid, errs
}

// <language> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#optionalChannelElements
type Language string

// TODO
// Whether <language> is valid.
//
// The <language> element must be one of the identifiers specified in the
// current list of ISO 639 language codes:
//
// See:
//   - https://www.rssboard.org/rss-language-codes
//   - https://www.loc.gov/standards/iso639-2
func (r Language) IsValid() bool { return true }

// <copyright> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#optionalChannelElements
type Copyright string

// Whether <copyright> is valid.
func (r Copyright) IsValid() bool { return true }

// <managingEditor> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#optionalChannelElements
type ManagingEditor string

// Whether <managingEditor> is valid.
func (r ManagingEditor) IsValid() bool { return true }

// <webMaster> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#optionalChannelElements
type WebMaster string

// Whether <webMaster> is valid.
func (r WebMaster) IsValid() bool { return true }

// <pubDate> is an optional sub-element of <channel> and <item>.
//
// See:
//   - https://validator.w3.org/feed/docs/rss2.html#optionalChannelElements
//   - https://validator.w3.org/feed/docs/rss2.html#ltpubdategtSubelementOfLtitemgt
type PubDate struct {
	XMLName  xml.Name `xml:"pubDate"`   // required
	CharData []byte   `xml:",chardata"` // required
}

// Returns whether <pubDate> is valid and a slice containing any errors.
//
// <pubDate> must conform to the Date and Time Specification of RFC822, with
// the exception that the year may be expressed with two characters or four
// characters (four preferred).
//
// See: http://asg.web.cmu.edu/rfc/rfc822.html
func (r PubDate) IsValid() (bool, []error) {
	isValid, errs := true, []error{}
	msg := fmt.Sprintf("Element <%s> value '%s' is invalid", r.XMLName.Local, r.CharData)
	if ok, err := IsNotEmpty(string(r.CharData)); !ok {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, err))
	}
	if ok, err := IsValidDate(string(r.CharData)); !ok {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, err))
	}
	return isValid, errs
}

// <lastBuildDate> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#optionalChannelElements
type LastBuildDate struct {
	XMLName  xml.Name `xml:"lastBuildDate"` // required
	CharData []byte   `xml:",chardata"`     // required
}

// Returns whether <lastBuildDate> is valid and a slice containing any errors.
//
// <lastBuildDate> must conform to the Date and Time Specification of RFC822,
// with the exception that the year may be expressed with two characters or
// four characters (four preferred).
//
// See: http://asg.web.cmu.edu/rfc/rfc822.html
func (r LastBuildDate) IsValid() (bool, []error) {
	isValid, errs := true, []error{}
	msg := fmt.Sprintf("Element <%s> value '%s' is invalid", r.XMLName.Local, r.CharData)
	if ok, err := IsNotEmpty(string(r.CharData)); !ok {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, err))
	}
	if ok, err := IsValidDate(string(r.CharData)); !ok {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, err))
	}
	return isValid, errs
}

// <category> is an optional sub-element of <channel> and <item>.
//
// The <channel>-level category element follows the same rules as the
// <item>-level category element.
//
// See:
//   - https://validator.w3.org/feed/docs/rss2.html#optionalChannelElements
//   - https://validator.w3.org/feed/docs/rss2.html#ltcategorygtSubelementOfLtitemgt
type Category struct {
	XMLName  xml.Name `xml:"category"`              // required
	CharData []byte   `xml:",chardata"`             // required
	Domain   Domain   `xml:"domain,attr,omitempty"` // optional
}

// Returns whether <category> is valid and a slice containing any errors.
//
// TODO: Check that the value of the element is a forward-slash-separated
// string.
func (r Category) IsValid() (bool, []error) {
	isValid, errs := true, []error{}
	msg := fmt.Sprintf("Element <%s> value '%s' is invalid", r.XMLName.Local, r.CharData)
	if ok, err := IsNotEmpty(string(r.CharData)); !ok {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, err))
	}
	if r.Domain != nil {
		msg := fmt.Sprintf("Attribute 'domain' of <%s> value '%s' is invalid", r.XMLName.Local, *r.Domain)
		if ok, err := IsNotEmpty(*r.Domain); !ok {
			isValid = false
			errs = append(errs, fmt.Errorf("%s: %w", msg, err))
		}
	}
	return isValid, errs
}

// 'domain' is an optional attribute of <category> and a required attribute of
// <cloud>.
//
// See:
//   - https://validator.w3.org/feed/docs/rss2.html#ltcategorygtSubelementOfLtitemgtv
//   - https://validator.w3.org/feed/docs/rss2.html#ltcloudgtSubelementOfLtchannelgt
type Domain *string

// <generator> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#optionalChannelElements
type Generator string

// Whether <generator> is valid.
func (r Generator) IsValid() bool { return true }

// <docs> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#optionalChannelElements
type Docs string

// Whether <docs> is valid.
func (r Docs) IsValid() bool { return true }

// <cloud> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltcloudgtSubelementOfLtchannelgt
type Cloud struct {
	XMLName           xml.Name          `xml:"cloud"`                  // required
	CharData          []byte            `xml:",chardata"`              // prohibited
	Domain            Domain            `xml:"domain,attr"`            // required
	Port              Port              `xml:"port,attr"`              // required
	Path              Path              `xml:"path,attr"`              // required
	RegisterProcedure RegisterProcedure `xml:"registerProcedure,attr"` // required
	Protocol          Protocol          `xml:"protocol,attr"`          // required
}

// Returns whether <cloud> is valid and a slice containing any errors.
func (r Cloud) IsValid() (bool, []error) {
	isValid, errs := true, []error{}
	msg := fmt.Sprintf("Element <%s> is invalid", r.XMLName.Local)
	if ok, err := IsEmpty(string(r.CharData)); !ok {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, err))
	}
	// <cloud> contains four required attributes: domain, port,
	// registerProcedure, protocol
	if r.Domain == nil {
		msg := fmt.Sprintf("Attribute 'domain' of <%s> is required", r.XMLName.Local)
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, ErrInvalidElement))
	} else {
		msg := fmt.Sprintf("Attribute 'domain' of <%s> value '%s' is invalid", r.XMLName.Local, *r.Domain)
		if ok, err := IsNotEmpty(*r.Domain); !ok {
			isValid = false
			errs = append(errs, fmt.Errorf("%s: %w", msg, err))
		}
	}
	if r.Port == nil {
		msg := fmt.Sprintf("Attribute 'port' of <%s> is required", r.XMLName.Local)
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, ErrInvalidElement))
	} else {
		msg := fmt.Sprintf("Attribute 'port' of <%s> value '%s' is invalid", r.XMLName.Local, *r.Port)
		// 'port' must be a positive integer.
		if i, err := strconv.ParseUint(*r.Port, 10, 0); err != nil || (i < 1 && i > 65535) {
			isValid = false
			errs = append(errs, fmt.Errorf("%s: %w: must be a valid port (1-65535)", msg, ErrInvalidValue))
		}
	}
	if r.Path == nil {
		msg := fmt.Sprintf("Attribute 'path' of <%s> is required", r.XMLName.Local)
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, ErrInvalidElement))
	} else {
		msg := fmt.Sprintf("Attribute 'path' of <%s> value '%s' is invalid", r.XMLName.Local, *r.Path)
		if ok, err := IsNotEmpty(*r.Path); !ok {
			isValid = false
			errs = append(errs, fmt.Errorf("%s: %w", msg, err))
		}
	}
	if r.RegisterProcedure == nil {
		msg := fmt.Sprintf("Attribute 'registerProcedure' of <%s> is required", r.XMLName.Local)
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, ErrInvalidElement))
	} else {
		msg := fmt.Sprintf("Attribute 'registerProcedure' of <%s> value '%s' is invalid", r.XMLName.Local, *r.RegisterProcedure)
		if ok, err := IsNotEmpty(*r.RegisterProcedure); !ok {
			isValid = false
			errs = append(errs, fmt.Errorf("%s: %w", msg, err))
		}
	}
	if r.Protocol == nil {
		msg := fmt.Sprintf("Attribute 'protocol' of <%s> is required", r.XMLName.Local)
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, ErrInvalidElement))
	} else {
		msg := fmt.Sprintf("Attribute 'protocol' of <%s> value '%s' is invalid", r.XMLName.Local, *r.Protocol)
		if *r.Protocol != "xml-rpc" && *r.Protocol != "soap" && *r.Protocol != "http-post" {
			isValid = false
			errs = append(errs, fmt.Errorf("%s: %w: must be one of \"xml-rpc\", \"soap\", or \"http-post\"", msg, ErrInvalidValue))
		}
	}
	return isValid, errs
}

// 'port' is required attribute of <cloud>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltcloudgtSubelementOfLtchannelgt
type Port *string

// 'path' is required attribute of <cloud>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltcloudgtSubelementOfLtchannelgt
type Path *string

// 'registerProcedure' is required attribute of <cloud>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltcloudgtSubelementOfLtchannelgt
type RegisterProcedure *string

// 'protocol' is required attribute of <cloud>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltcloudgtSubelementOfLtchannelgt
type Protocol *string

// <ttl> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltttlgtSubelementOfLtchannelgt
type TTL struct {
	XMLName  xml.Name `xml:"ttl"`       // required
	CharData []byte   `xml:",chardata"` // required
}

// Returns whether <ttl> is valid and a slice containing any errors.
func (r TTL) IsValid() (bool, []error) {
	isValid, errs := true, []error{}
	msg := fmt.Sprintf("Element <%s> value '%s' is invalid", r.XMLName.Local, r.CharData)
	if ok, err := IsNotEmpty(string(r.CharData)); !ok {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, err))
	}
	// '<ttl>' must be a positive integer.
	if i, err := strconv.ParseUint(string(r.CharData), 10, 0); err != nil || i < 0 {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w: must be a positive integer", msg, ErrInvalidValue))
	}
	return isValid, errs
}

// <image> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltimagegtSubelementOfLtchannelgt
//
// TODO: Set default values for width and height.
type Image struct {
	XMLName     xml.Name    `xml:"image"`                 // required
	URL         URL         `xml:"url"`                   // required
	Title       Title       `xml:"title"`                 // required
	Link        Link        `xml:"link"`                  // required
	Width       Width       `xml:"width,omitempty"`       // optional
	Height      Height      `xml:"height,omitempty"`      // optional
	Description Description `xml:"description,omitempty"` // optional
}

// Whether <image> is valid.
func (r Image) IsValid() (bool, []error) {
	return true, nil
	// // Required sub-elements: <url>, <title>, <link>
	// //
	// // NOTE: In practice the image <title> and <link> should have the same value
	// // as the channel's <title> and <link>.
	//
	//	if r.URL.IsValid() || r.Title.IsValid() || r.Link.IsValid() {
	//		return false
	//	}
	//
	// // Optional sub-elements: <width>, <height>, <description>
	// return r.Width.IsValid() && r.Height.IsValid()
}

// <url> is a required sub-element of <image>.
//
// It is also a required attribute of <source> and <enclosure>.
//
// See:
//   - https://validator.w3.org/feed/docs/rss2.html#ltimagegtSubelementOfLtchannelgt
//   - https://validator.w3.org/feed/docs/rss2.html#ltsourcegtSubelementOfLtitemgt
//   - https://validator.w3.org/feed/docs/rss2.html#ltenclosuregtSubelementOfLtitemgt
type URL *string

// <width> is an optional sub-element of <image>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltimagegtSubelementOfLtchannelgt
type Width string

// Whether <width> is valid.
//
// The maximum value for width is 144, default value is 88.
func (r Width) IsValid() bool {
	if i, err := strconv.ParseUint(string(r), 10, 0); err != nil || i > 144 {
		return false
	}
	return true
}

// <height> is an optional sub-element of <image>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltimagegtSubelementOfLtchannelgt
type Height string

// Whether <height> is valid.
//
// The maximum value for height is 400, default value is 31.
func (r Height) IsValid() bool {
	if i, err := strconv.ParseUint(string(r), 10, 0); err != nil || i > 400 {
		return false
	}
	return true
}

// <rating> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#optionalChannelElements
type Rating string

// Whether <rating> is valid.
func (r Rating) IsValid() bool { return true }

// <textInput> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#lttextinputgtSubelementOfLtchannelgt
type TextInput struct {
	XMLName     xml.Name     `xml:"textInput"`   // required
	Title       *Title       `xml:"title"`       // required
	Description *Description `xml:"description"` // required
	Name        *Name        `xml:"name"`        // required
	Link        *Link        `xml:"link"`        // required
}

// Returns whether <textInput> is valid and a slice containing any errors.
func (r TextInput) IsValid() (bool, []error) {
	isValid, errs := true, []error{}
	msg := fmt.Sprintf("Element <%s> is invalid", r.XMLName.Local)
	// <textInput> contains four required sub-elements: <title>, <description>,
	// <name>, <link>
	if r.Title == nil || r.Description == nil || r.Name == nil || r.Link == nil {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w: <title>, <description>, <name> and <link> must be present", msg, ErrInvalidElement))
	}
	if ok, e := Validate(r); !ok {
		isValid = false
		errs = append(errs, e...)
	}
	return isValid, errs
}

// <name> is a required sub-element of <textInput>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#lttextinputgtSubelementOfLtchannelgt
type Name struct {
	XMLName  xml.Name `xml:"name"`      // required
	CharData []byte   `xml:",chardata"` // required
}

// Returns whether <name> is valid and a slice containing any errors.
func (r Name) IsValid() (bool, []error) {
	isValid, errs := true, []error{}
	msg := fmt.Sprintf("Element <%s> value '%s' is invalid", r.XMLName.Local, r.CharData)
	if ok, err := IsNotEmpty(string(r.CharData)); !ok {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, err))
	}
	return isValid, errs
}

// <skipHours> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#optionalChannelElements
type SkipHours struct {
	XMLName xml.Name `xml:"skipHours"` // required
	Hour    []*Hour  `xml:"hour"`      // required
}

// Whether <skipHours> is valid.
//
// This element contains up to 24 <hour> sub-elements whose value is a number
// between 0 and 23.
func (r SkipHours) IsValid() bool {
	if len(r.Hour) > 24 {
		return false
	} else {
		for _, h := range r.Hour {
			if !h.IsValid() {
				return false
			}
		}
	}
	return true
}

// <hour> is an optional sub-element of <skipHours>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#optionalChannelElements
type Hour int

// Whether <hour> is valid.
func (r Hour) IsValid() bool {
	if r < 0 || r > 23 {
		return false
	}
	return true
}

// <skipDays> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#optionalChannelElements
type SkipDays struct {
	XMLName xml.Name `xml:"skipDays"` // required
	Day     []*Day   `xml:"hour"`     // required
}

// Whether <skipDays> is valid.
//
// This element contains up to seven <day> sub-elements whose value is
// Monday, Tuesday, Wednesday, Thursday, Friday, Saturday or Sunday.
func (r SkipDays) IsValid() bool {
	if len(r.Day) > 7 {
		return false
	} else {
		for _, d := range r.Day {
			if !d.IsValid() {
				return false
			}
		}
	}
	return true
}

// <day> is an optional sub-element of <skipDays>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#optionalChannelElements
type Day string

// Whether <day> is valid.
//
// TODO: Check if Monday - Sunday.
func (r Day) IsValid() bool { return true }

// <item> is an optional sub-element of <channel>.
//
// A channel may contain any number of <item>s.
//
// All elements of an item are optional, however at least one of title or
// description must be present.
//
// See: https://validator.w3.org/feed/docs/rss2.html#hrelementsOfLtitemgt
type Item struct {
	XMLName     xml.Name     `xml:"item"`                  // required
	Title       *Title       `xml:"title,omitempty"`       // conditionally required
	Link        *Link        `xml:"link,omitempty"`        // optional
	Description *Description `xml:"description,omitempty"` // conditionally required
	Source      *Source      `xml:"source,omitempty"`      // optional
	Enclosure   *Enclosure   `xml:"enclosure,omitempty"`   // optional
	Category    *Category    `xml:"category,omitempty"`    // optional
	PubDate     *PubDate     `xml:"pubDate,omitempty"`     // optional
	GUID        *GUID        `xml:"guid,omitempty"`        // optional
	Comments    *Comments    `xml:"comments,omitempty"`    // optional
	Author      *Author      `xml:"author,omitempty"`      // optional
}

// Returns whether <item> is valid and a slice containing any errors.
func (r Item) IsValid() (bool, []error) {
	isValid, errs := true, []error{}
	msg := fmt.Sprintf("Element <%s> is invalid", r.XMLName.Local)
	// At least one of title or description must be present.
	if (r.Title == nil || string(r.Title.CharData) == "") && (r.Description == nil || string(r.Description.CharData) == "") {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w: one of <title> or <description> must be present", msg, ErrInvalidElement))
	}
	if ok, e := Validate(r); !ok {
		isValid = false
		errs = append(errs, e...)
	}
	return isValid, errs
}

// <source> is an optional sub-element of <item>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltsourcegtSubelementOfLtitemgt
type Source struct {
	XMLName  xml.Name `xml:"source"`    // required
	CharData []byte   `xml:",chardata"` // optional
	URL      URL      `xml:"url,attr"`  // required
}

// Returns whether <source> is valid and a slice containing any errors.
func (r Source) IsValid() (bool, []error) {
	isValid, errs := true, []error{}
	if r.URL == nil {
		msg := fmt.Sprintf("Attribute 'url' of <%s> is required", r.XMLName.Local)
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, ErrInvalidElement))
	} else {
		msg := fmt.Sprintf("Attribute 'url' of <%s> value '%s' is invalid", r.XMLName.Local, *r.URL)
		if ok, err := IsNotEmpty(*r.URL); !ok {
			isValid = false
			errs = append(errs, fmt.Errorf("%s: %w", msg, err))
		}
		if ok, err := IsValidURI(*r.URL); !ok {
			isValid = false
			errs = append(errs, fmt.Errorf("%s: %w", msg, err))
		}
	}
	return isValid, errs
}

// <enclosure> is an optional sub-element of <item>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltenclosuregtSubelementOfLtitemgt
//
// NOTE: In most cases, the <enclosure> element is represented using a
// self-closing tag:
//
//	<enclosure url="..." length="..." type="..." />
//
// From the XML specification,
//
//	>The representation of an empty element is either a start-tag immediately
//	 followed by an end-tag, or an empty-element tag.
//
// Source: https://www.w3.org/TR/xml/#NT-ETag
//
// Self-closing tags are currently not implemented in encoding/xml:
//   - https://github.com/golang/go/issues/21399
//
// Ideally, both forms (a start-tag immediately followed by an end-tag, or an
// empty-element tag) represent valid XML, and therefore valid RSS.
type Enclosure struct {
	XMLName  xml.Name `xml:"enclosure"`   // required
	CharData []byte   `xml:",chardata"`   // prohibited
	URL      URL      `xml:"url,attr"`    // required
	Length   Length   `xml:"length,attr"` // required
	Type     Type     `xml:"type,attr"`   // required
}

// Returns whether <enclosure> is valid and a slice containing any errors.
func (r Enclosure) IsValid() (bool, []error) {
	isValid, errs := true, []error{}
	msg := fmt.Sprintf("Element <%s> is invalid", r.XMLName.Local)
	if ok, err := IsEmpty(string(r.CharData)); !ok {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, err))
	}
	if r.URL == nil {
		msg := fmt.Sprintf("Attribute 'url' of <%s> is required", r.XMLName.Local)
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, ErrInvalidElement))
	} else {
		msg := fmt.Sprintf("Attribute 'url' of <%s> value '%s' is invalid", r.XMLName.Local, *r.URL)
		if ok, err := IsNotEmpty(*r.URL); !ok {
			isValid = false
			errs = append(errs, fmt.Errorf("%s: %w", msg, err))
		}
		if ok, err := IsValidURI(*r.URL); !ok {
			isValid = false
			errs = append(errs, fmt.Errorf("%s: %w", msg, err))
		}
	}
	if r.Length == nil {
		msg := fmt.Sprintf("Attribute 'length' of <%s> is required", r.XMLName.Local)
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, ErrInvalidElement))
	} else {
		msg := fmt.Sprintf("Attribute 'length' of <%s> value '%s' is invalid", r.XMLName.Local, *r.Length)
		// 'length' must be a positive integer. NOTE: Use zero for unknown length.
		if i, err := strconv.ParseUint(*r.Length, 10, 0); err != nil || i < 0 {
			isValid = false
			errs = append(errs, fmt.Errorf("%s: %w: must be a positive integer", msg, ErrInvalidValue))
		}
	}
	if r.Type == nil {
		msg := fmt.Sprintf("Attribute 'type' of <%s> is required", r.XMLName.Local)
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, ErrInvalidElement))
	} else {
		msg := fmt.Sprintf("Attribute 'type' of <%s> value '%s' is invalid", r.XMLName.Local, *r.Type)
		if ok, err := IsNotEmpty(*r.Type); !ok {
			isValid = false
			errs = append(errs, fmt.Errorf("%s: %w", msg, err))
		}
	}
	return isValid, errs
}

// 'length' is a required attribute of <enclosure>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltenclosuregtSubelementOfLtitemgt
type Length *string

// 'type' is a required attribute of <enclosure>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltenclosuregtSubelementOfLtitemgt
type Type *string

// <guid> is an optional sub-element of <item>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltguidgtSubelementOfLtitemgt
type GUID struct {
	XMLName     xml.Name    `xml:"guid"`                       // required
	CharData    []byte      `xml:",chardata"`                  // required
	IsPermaLink IsPermaLink `xml:"isPermaLink,attr,omitempty"` // optional
}

// Returns whether <guid> is valid and a slice containing any errors.
func (r GUID) IsValid() (bool, []error) {
	isValid, errs := true, []error{}
	msg := fmt.Sprintf("Element <%s> value '%s' is invalid", r.XMLName.Local, r.CharData)
	if ok, err := IsNotEmpty(string(r.CharData)); !ok {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, err))
	}
	// 'isPermaLink' attribute must be "true" or "false"
	if r.IsPermaLink != nil {
		msg := fmt.Sprintf("Attribute 'isPermaLink' of <%s> value '%s' is invalid", r.XMLName.Local, *r.IsPermaLink)
		if *r.IsPermaLink != "true" && *r.IsPermaLink != "false" {
			isValid = false
			errs = append(errs, fmt.Errorf("%s: %w: must be \"true\" or \"false\"", msg, ErrInvalidValue))
		}
	}
	// If the guid element has an attribute named isPermaLink with a value of
	// true, the reader may assume that it is a permalink to the item.
	if r.IsPermaLink != nil && *r.IsPermaLink == "true" {
		if ok, err := IsValidURI(string(r.CharData)); !ok {
			isValid = false
			errs = append(errs, fmt.Errorf("%s: %w", msg, err))
		}
	}
	return isValid, errs
}

// 'isPermaLink' is an optional attribute of <guid>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltguidgtSubelementOfLtitemgt
type IsPermaLink *string

// <comments> is an optional sub-element of <item>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltcommentsgtSubelementOfLtitemgt
type Comments struct {
	XMLName  xml.Name `xml:"comments"`  // required
	CharData []byte   `xml:",chardata"` // required
}

// Returns whether <comments> is valid and a slice containing any errors.
func (r Comments) IsValid() (bool, []error) {
	isValid, errs := true, []error{}
	msg := fmt.Sprintf("Element <%s> value '%s' is invalid", r.XMLName.Local, r.CharData)
	if ok, err := IsNotEmpty(string(r.CharData)); !ok {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, err))
	}
	if ok, err := IsValidURI(string(r.CharData)); !ok {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, err))
	}
	return isValid, errs
}

// <author> is an optional sub-element of <item>.
//
// Example:
//
//	<author>first.last@example.com</author>
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltauthorgtSubelementOfLtitemgt
type Author struct {
	XMLName  xml.Name `xml:"author"`    // required
	CharData []byte   `xml:",chardata"` // required
}

// Returns whether <author> is valid and a slice containing any errors.
func (r Author) IsValid() (bool, []error) {
	isValid, errs := true, []error{}
	msg := fmt.Sprintf("Element <%s> value '%s' is invalid", r.XMLName.Local, r.CharData)
	if ok, err := IsNotEmpty(string(r.CharData)); !ok {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, err))
	}
	if ok, err := IsValidMailAddress(string(r.CharData)); !ok {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, err))
	}
	return isValid, errs
}
