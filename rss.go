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
			// Indirect returns the value that v points to.
			// If v is a nil pointer, Indirect returns a zero Value.
			// If v is not a pointer, Indirect returns v.
			v := reflect.Indirect(reflect.ValueOf(t))
			// Kind returns v's Kind.
			// If v is the zero Value (IsValid returns false), Kind returns Invalid.
			if v.Kind() == reflect.Ptr {
				// Check whether v is nil before calling IsValid.
				if v.IsNil() {
					isValid = false
					errs = append(errs, fmt.Errorf("Element cannot be empty"))
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
	XMLName        xml.Name       `xml:"channel"`        // required
	Title          Title          `xml:"title"`          // required
	Link           Link           `xml:"link"`           // required
	Description    Description    `xml:"description"`    // required
	Language       Language       `xml:"language"`       // optional
	Copyright      Copyright      `xml:"copyright"`      // optional
	ManagingEditor ManagingEditor `xml:"managingEditor"` // optional
	WebMaster      WebMaster      `xml:"webMaster"`      // optional
	PubDate        PubDate        `xml:"pubDate"`        // optional
	LastBuildDate  LastBuildDate  `xml:"lastBuildDate"`  // optional
	Category       Category       `xml:"category"`       // optional
	Generator      Generator      `xml:"generator"`      // optional
	Docs           Docs           `xml:"docs"`           // optional
	Cloud          Cloud          `xml:"cloud"`          // optional
	TTL            TTL            `xml:"ttl"`            // optional
	Image          Image          `xml:"image"`          // optional
	Rating         Rating         `xml:"rating"`         // optional
	TextInput      TextInput      `xml:"textInput"`      // optional
	SkipHours      SkipHours      `xml:"skipHours"`      // optional
	SkipDays       SkipDays       `xml:"skipDays"`       // optional
	Item           []*Item        `xml:"item"`           // optional
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
type Title string

// Whether <title> is valid.
func (r Title) IsValid() bool {
	return r != ""
}

// <link> is a required sub-element of <channel>, <image>, <textInput>, and
// <item>.
//
// See:
//   - https://validator.w3.org/feed/docs/rss2.html#requiredChannelElements
//   - https://validator.w3.org/feed/docs/rss2.html#ltimagegtSubelementOfLtchannelgt
//   - https://validator.w3.org/feed/docs/rss2.html#lttextinputgtSubelementOfLtchannelgt
//   - https://validator.w3.org/feed/docs/rss2.html#hrelementsOfLtitemgt
type Link string

// Whether <link> is valid.
//
// <link> must be a valid URL.
func (r Link) IsValid() bool {
	// return IsValidURI(string(r))
	return true
}

// <description> is a required sub-element of <channel> and <textInput> and an
// optional sub-element of <image> and <item>
//
// See:
//   - https://validator.w3.org/feed/docs/rss2.html#requiredChannelElements
//   - https://validator.w3.org/feed/docs/rss2.html#ltimagegtSubelementOfLtchannelgt
//   - https://validator.w3.org/feed/docs/rss2.html#lttextinputgtSubelementOfLtchannelgt
//   - https://validator.w3.org/feed/docs/rss2.html#hrelementsOfLtitemgt
type Description string

// Whether <descripton> is valid.
func (r Description) IsValid() bool {
	return r != ""
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
type PubDate string

// Whether <pubDate> is valid.
//
// <pubDate> must conform to the Date and Time Specification of RFC822, with
// the exception that the year may be expressed with two characters or four
// characters (four preferred).
//
// See: http://asg.web.cmu.edu/rfc/rfc822.html
func (r PubDate) IsValid() (bool, error) {
	msg := fmt.Sprintf("Element <pubDate> value '%s' is invalid", r)
	if ok, err := IsNotEmpty(string(r)); !ok {
		return false, fmt.Errorf("%s: %w", msg, err)
	}
	if ok, err := IsValidDate(string(r)); !ok {
		return false, fmt.Errorf("%s: %w", msg, err)
	}
	return true, nil
}

// <lastBuildDate> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#optionalChannelElements
type LastBuildDate string

// Whether <lastBuildDate> is valid.
//
// <lastBuildDate> must conform to the Date and Time Specification of RFC822,
// with the exception that the year may be expressed with two characters or
// four characters (four preferred).
//
// See: http://asg.web.cmu.edu/rfc/rfc822.html
func (r LastBuildDate) IsValid() (bool, error) {
	msg := fmt.Sprintf("Element <lastBuildDate> value '%s' is invalid", r)
	if ok, err := IsNotEmpty(string(r)); !ok {
		return false, fmt.Errorf("%s: %w", msg, err)
	}
	if ok, err := IsValidDate(string(r)); !ok {
		return false, fmt.Errorf("%s: %w", msg, err)
	}
	return true, nil
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
	XMLName        xml.Name       `xml:"category"`    // required
	CharData       []byte         `xml:",chardata"`   // required
	CategoryDomain CategoryDomain `xml:"domain,attr"` // optional
}

// Whether <category> is valid.
func (r Category) IsValid() bool { return true }

// domain (<category>) is an optional attribute of <category>. It differs from
// the required domain attribute of <cloud>.
//
// See:
//   - https://validator.w3.org/feed/docs/rss2.html#ltcategorygtSubelementOfLtitemgtv
type CategoryDomain string

// Whether domain (<category>) is valid.
func (r CategoryDomain) IsValid() (bool, error) {
	msg := fmt.Sprintf("Attribute domain value '%s' is invalid", r)
	if ok, err := IsNotEmpty(string(r)); !ok {
		return false, fmt.Errorf("%s: %w", msg, err)
	}
	if ok, err := IsValidURI(string(r)); !ok {
		return false, fmt.Errorf("%s: %w", msg, err)
	}
	return true, nil
}

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

// TODO: Convert struct fields to types.
// <cloud> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltcloudgtSubelementOfLtchannelgt
type Cloud struct {
	XMLName           xml.Name    `xml:"cloud"`                  // required
	CloudDomain       CloudDomain `xml:"domain,attr"`            // required
	Port              string      `xml:"port,attr"`              // required
	Path              string      `xml:"path,attr"`              // required
	RegisterProcedure string      `xml:"registerProcedure,attr"` // required
	Protocol          string      `xml:"protocol,attr"`          // required
}

// Whether <cloud> is valid.
//
// TODO: https://www.rssboard.org/rsscloud-interface
func (r Cloud) IsValid() bool { return true }

// domain (<cloud>) is a required attribute of <cloud>. It differs from the
// optional domain attribute of <category>.
//
// See:
//   - https://validator.w3.org/feed/docs/rss2.html#ltcloudgtSubelementOfLtchannelgt
type CloudDomain string

// Whether domain (<cloud>) is valid.
func (r CloudDomain) IsValid() bool { return true }

// <ttl> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltttlgtSubelementOfLtchannelgt
type TTL string

// Whether <ttl> is valid.
//
// <ttl> must be a positive integer.
func (r TTL) IsValid() bool {
	if i, err := strconv.ParseUint(string(r), 10, 0); err != nil || i < 0 {
		return false
	}
	return true
}

// <image> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltimagegtSubelementOfLtchannelgt
//
// TODO: Set default values for width and height.
type Image struct {
	XMLName     xml.Name    `xml:"image"`       // required
	URL         URL         `xml:"url"`         // required
	Title       Title       `xml:"title"`       // required
	Link        Link        `xml:"link"`        // required
	Width       Width       `xml:"width"`       // optional
	Height      Height      `xml:"height"`      // optional
	Description Description `xml:"description"` // optional
}

// Whether <image> is valid.
func (r Image) IsValid() bool {
	// Required sub-elements: <url>, <title>, <link>
	//
	// NOTE: In practice the image <title> and <link> should have the same value
	// as the channel's <title> and <link>.
	if r.URL.IsValid() || r.Title.IsValid() || r.Link.IsValid() {
		return false
	}
	// Optional sub-elements: <width>, <height>, <description>
	return r.Width.IsValid() && r.Height.IsValid()
}

// <url> is a required sub-element of <image> and a required attribute of
// <enclosure>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltimagegtSubelementOfLtchannelgt
type URL string

// Whether <url> is valid.
func (r URL) IsValid() bool {
	// return IsValidURI(string(r))
	return true
}

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
	XMLName     xml.Name    `xml:"textInput"`   // required
	Title       Title       `xml:"title"`       // required
	Description Description `xml:"description"` // required
	Name        Name        `xml:"name"`        // required
	Link        Link        `xml:"link"`        // required
}

// Whether <textInput> is valid.
func (r TextInput) IsValid() bool {
	return r.Title != "" && r.Description != "" && r.Name != "" && r.Link != ""
}

// <name> is a required sub-element of <textInput>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#lttextinputgtSubelementOfLtchannelgt
type Name string

// Whether <name> is valid.
func (r Name) IsValid() bool { return true }

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
	XMLName     xml.Name    `xml:"item"`             // required
	Title       Title       `xml:"title"`            // conditionally required
	Link        Link        `xml:"link"`             // optional
	Description Description `xml:"description"`      // conditionally required
	Source      Source      `xml:"source"`           // optional
	Enclosure   Enclosure   `xml:"enclosure"`        // optional
	Category    Category    `xml:"category"`         // optional
	PubDate     PubDate     `xml:"pubDate"`          // optional
	GUID        GUID        `xml:"guid"`             // optional
	Comments    Comments    `xml:"comments"`         // optional
	Author      Author      `xml:"author,omitempty"` // optional
}

// Whether <item> is valid.
func (r Item) IsValid() bool {
	return r.Title.IsValid() || r.Description.IsValid()
}

// <source> is an optional sub-element of <item>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltsourcegtSubelementOfLtitemgt
type Source struct {
	XMLName   xml.Name  `xml:"source"`    // required
	CharData  []byte    `xml:",chardata"` // optional
	SourceURL SourceURL `xml:"url,attr"`  // required
}

// Returns whether <source> is valid and a slice containing any errors.
func (r Source) IsValid() (bool, []error) {
	isValid, errs := true, []error{}
	if ok, e := Validate(r); !ok {
		isValid = false
		errs = append(errs, e...)
	}
	return isValid, errs
}

// 'url' is a required attribute of <source>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltsourcegtSubelementOfLtitemgt
type SourceURL string

// Returns whether 'url' is valid and a slice containing any errors.
func (r SourceURL) IsValid() (bool, []error) {
	isValid, errs := true, []error{}
	msg := fmt.Sprintf("Attribute 'url' of <source> value '%s' is invalid", r)
	if ok, err := IsNotEmpty(string(r)); !ok {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, err))
	}
	if ok, err := IsValidURI(string(r)); !ok {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, err))
	}
	return isValid, errs
}

// <enclosure> is an optional sub-element of <item>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltenclosuregtSubelementOfLtitemgt
type Enclosure struct {
	XMLName      xml.Name     `xml:"enclosure"`   // required
	CharData     []byte       `xml:",chardata"`   // prohibited
	EnclosureURL EnclosureURL `xml:"url,attr"`    // required
	Length       Length       `xml:"length,attr"` // required
	Type         Type         `xml:"type,attr"`   // required
}

// Returns whether <enclosure> is valid and a slice containing any errors.
func (r Enclosure) IsValid() (bool, []error) {
	isValid, errs := true, []error{}
	msg := fmt.Sprintf("Element <enclosure> value '%s' is invalid", r)
	if ok, err := IsEmpty(string(r.CharData)); !ok {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, err))
	}
	if ok, e := Validate(r); !ok {
		isValid = false
		errs = append(errs, e...)
	}
	return isValid, errs
}

// 'url' is a required attribute of <enclosure>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltenclosuregtSubelementOfLtitemgt
type EnclosureURL string

// Returns whether 'url' is valid and a slice containing any errors.
func (r EnclosureURL) IsValid() (bool, []error) {
	isValid, errs := true, []error{}
	msg := fmt.Sprintf("Attribute 'url' of <enclosure> value '%s' is invalid", r)
	if ok, err := IsNotEmpty(string(r)); !ok {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, err))
	}
	if ok, err := IsValidURI(string(r)); !ok {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, err))
	}
	return isValid, errs
}

// 'length' is a required attribute of <enclosure>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltenclosuregtSubelementOfLtitemgt
type Length string

// Returns whether 'length' is valid and a slice containing any errors.
func (r Length) IsValid() (bool, []error) {
	isValid, errs := true, []error{}
	msg := fmt.Sprintf("Attribute 'length' of <enclosure> value '%s' is invalid", r)
	// 'length' must be a positive integer. NOTE: Use zero for unknown length.
	if i, err := strconv.ParseUint(string(r), 10, 0); err != nil || i < 0 {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w: must be a positive integer", msg, ErrInvalidValue))
	}
	return isValid, errs
}

// 'type' is a required attribute of <enclosure>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltenclosuregtSubelementOfLtitemgt
type Type string

// Returns whether 'type' is valid and a slice containing any errors.
func (r Type) IsValid() (bool, []error) {
	isValid, errs := true, []error{}
	msg := fmt.Sprintf("Attribute 'type' of <enclosure> value '%s' is invalid", r)
	if ok, err := IsNotEmpty(string(r)); !ok {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, err))
	}
	return isValid, errs
}

// <guid> is an optional sub-element of <item>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltguidgtSubelementOfLtitemgt
type GUID struct {
	XMLName     xml.Name    `xml:"guid"`             // required
	CharData    []byte      `xml:",chardata"`        // required
	IsPermaLink IsPermaLink `xml:"isPermaLink,attr"` // optional
}

// Returns whether <guid> is valid and a slice containing any errors.
func (r GUID) IsValid() (bool, []error) {
	isValid, errs := true, []error{}
	msg := fmt.Sprintf("Element <guid> value '%s' is invalid", r)
	if ok, err := IsNotEmpty(string(r.CharData)); !ok {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, err))
	}
	// If the guid element has an attribute named isPermaLink with a value of
	// true, the reader may assume that it is a permalink to the item.
	if r.IsPermaLink == "true" {
		if ok, err := IsValidURI(string(r.CharData)); !ok {
			isValid = false
			errs = append(errs, fmt.Errorf("%s: %w", msg, err))
		}
	}
	if ok, e := Validate(r); !ok {
		isValid = false
		errs = append(errs, e...)
	}
	return isValid, errs
}

// 'isPermaLink' is an optional attribute of <guid>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltguidgtSubelementOfLtitemgt
type IsPermaLink string

// Returns whether 'isPermaLink' is valid and a slice containing any errors.
func (r IsPermaLink) IsValid() (bool, []error) {
	isValid, errs := true, []error{}
	msg := fmt.Sprintf("Attribute 'isPermaLink' of <guid> value '%s' is invalid", r)
	// 'isPermaLink' attribute must be "true" or "false"
	if r != "true" && r != "false" {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w: must be \"true\" or \"false\"", msg, ErrInvalidValue))
	}
	return isValid, errs
}

// <comments> is an optional sub-element of <item>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltcommentsgtSubelementOfLtitemgt
type Comments string

// Returns whether <comments> is valid and a slice containing any errors.
func (r Comments) IsValid() (bool, []error) {
	isValid, errs := true, []error{}
	msg := fmt.Sprintf("Element <comments> value '%s' is invalid", r)
	if ok, err := IsNotEmpty(string(r)); !ok {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, err))
	}
	if ok, err := IsValidURI(string(r)); !ok {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, err))
	}
	return isValid, errs
}

// <author> is an optional sub-element of <item>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltauthorgtSubelementOfLtitemgt
type Author string

// Returns whether <author> is valid and a slice containing any errors.
func (r Author) IsValid() (bool, []error) {
	isValid, errs := true, []error{}
	msg := fmt.Sprintf("Element <author> value '%s' is invalid", r)
	if ok, err := IsNotEmpty(string(r)); !ok {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, err))
	}
	if ok, err := IsValidMailAddress(string(r)); !ok {
		isValid = false
		errs = append(errs, fmt.Errorf("%s: %w", msg, err))
	}
	return isValid, errs
}
