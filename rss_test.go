// Copyright 2022 Nickolas Kraus. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rss

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Use table drive tests.

// Returns a pointer to the passed value.
func Ptr[T any](v T) *T {
	return &v
}

// func TestTitle(t *testing.T) {
// 	t.Run("test <title> - ok", func(t *testing.T) {
// 		var r Title = Title{
// 			XMLName:  xml.Name{Space: "", Local: "title"},
// 			CharData: []byte("Title"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.True(t, ret)
// 		assert.Empty(t, errs)
// 	})
// 	t.Run("test <title> - fail - empty", func(t *testing.T) {
// 		var r Title = Title{
// 			XMLName:  xml.Name{Space: "", Local: "title"},
// 			CharData: []byte(""),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 1, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrEmptyValue)
// 		assert.ErrorContains(t, errs[0], "Element <title> value '' is invalid")
// 	})
// 	t.Run("test <title> - unmarshal", func(t *testing.T) {
// 		var r Title
// 		s := []byte(`<title>Title</title>`)
// 		err := xml.Unmarshal(s, &r)
// 		assert.Equal(t, "Title", string(r.CharData))
// 		assert.Nil(t, err)
// 	})
// 	t.Run("test <title> - marshal", func(t *testing.T) {
// 		var r Title = Title{
// 			XMLName:  xml.Name{Space: "", Local: "title"},
// 			CharData: []byte("Title"),
// 		}
// 		exp := []byte(`<title>Title</title>`)
// 		s, err := xml.Marshal(r)
// 		assert.Equal(t, exp, s)
// 		assert.Nil(t, err)
// 		ret, errs := Validate(r)
// 		assert.True(t, ret)
// 		assert.Empty(t, errs)
// 	})
// }
//
// func TestLink(t *testing.T) {
// 	t.Run("test <link> - ok", func(t *testing.T) {
// 		var r Link = Link{
// 			XMLName:  xml.Name{Space: "", Local: "link"},
// 			CharData: []byte("https://example.com"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.True(t, ret)
// 		assert.Empty(t, errs)
// 	})
// 	t.Run("test <link> - fail - empty", func(t *testing.T) {
// 		var r Link = Link{
// 			XMLName:  xml.Name{Space: "", Local: "link"},
// 			CharData: []byte(""),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 2, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrEmptyValue)
// 		assert.ErrorContains(t, errs[0], "Element <link> value '' is invalid")
// 		assert.ErrorIs(t, errs[1], ErrInvalidURI)
// 		assert.ErrorContains(t, errs[1], "Element <link> value '' is invalid")
// 	})
// 	t.Run("test <link> - fail - invalid uri", func(t *testing.T) {
// 		var r Link = Link{
// 			XMLName:  xml.Name{Space: "", Local: "link"},
// 			CharData: []byte("bad uri"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 1, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrInvalidURI)
// 		assert.ErrorContains(t, errs[0], "Element <link> value 'bad uri' is invalid")
// 	})
// 	t.Run("test <link> - unmarshal", func(t *testing.T) {
// 		var r Link
// 		s := []byte(`<link>https://example.com</link>`)
// 		err := xml.Unmarshal(s, &r)
// 		assert.Equal(t, "https://example.com", string(r.CharData))
// 		assert.Nil(t, err)
// 	})
// 	t.Run("test <link> - marshal", func(t *testing.T) {
// 		var r Link = Link{
// 			XMLName:  xml.Name{Space: "", Local: "link"},
// 			CharData: []byte("https://example.com"),
// 		}
// 		exp := []byte(`<link>https://example.com</link>`)
// 		s, err := xml.Marshal(r)
// 		assert.Equal(t, exp, s)
// 		assert.Nil(t, err)
// 		ret, errs := Validate(r)
// 		assert.True(t, ret)
// 		assert.Empty(t, errs)
// 	})
// }
//
// func TestDescription(t *testing.T) {
// 	t.Run("test <description> - ok", func(t *testing.T) {
// 		var r Description = Description{
// 			XMLName:  xml.Name{Space: "", Local: "description"},
// 			CharData: []byte("Description"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.True(t, ret)
// 		assert.Empty(t, errs)
// 	})
// 	t.Run("test <description> - fail - empty", func(t *testing.T) {
// 		var r Description = Description{
// 			XMLName:  xml.Name{Space: "", Local: "description"},
// 			CharData: []byte(""),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 1, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrEmptyValue)
// 		assert.ErrorContains(t, errs[0], "Element <description> value '' is invalid")
// 	})
// 	t.Run("test <description> - unmarshal", func(t *testing.T) {
// 		var r Description
// 		s := []byte(`<description>Description</description>`)
// 		err := xml.Unmarshal(s, &r)
// 		assert.Equal(t, "Description", string(r.CharData))
// 		assert.Nil(t, err)
// 	})
// 	t.Run("test <description> - marshal", func(t *testing.T) {
// 		var r Description = Description{
// 			XMLName:  xml.Name{Space: "", Local: "description"},
// 			CharData: []byte("Description"),
// 		}
// 		exp := []byte(`<description>Description</description>`)
// 		s, err := xml.Marshal(r)
// 		assert.Equal(t, exp, s)
// 		assert.Nil(t, err)
// 		ret, errs := Validate(r)
// 		assert.True(t, ret)
// 		assert.Empty(t, errs)
// 	})
// }
//
// func TestCloud(t *testing.T) {
// 	t.Run("test <cloud> - ok", func(t *testing.T) {
// 		var r Cloud = Cloud{
// 			XMLName:           xml.Name{Space: "", Local: "cloud"},
// 			Domain:            Ptr("rpc.example.com"),
// 			Port:              Ptr("1337"),
// 			Path:              Ptr("/path"),
// 			RegisterProcedure: Ptr("procedure"),
// 			Protocol:          Ptr("xml-rpc"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.True(t, ret)
// 		assert.Empty(t, errs)
// 	})
// 	t.Run("test <cloud> - fail - not empty", func(t *testing.T) {
// 		var r Cloud = Cloud{
// 			XMLName:           xml.Name{Space: "", Local: "cloud"},
// 			CharData:          []byte("not empty"),
// 			Domain:            Ptr("rpc.example.com"),
// 			Port:              Ptr("1337"),
// 			Path:              Ptr("/path"),
// 			RegisterProcedure: Ptr("procedure"),
// 			Protocol:          Ptr("xml-rpc"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 1, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrNonEmptyValue)
// 		assert.ErrorContains(t, errs[0], "Element <cloud> is invalid")
// 	})
// 	t.Run("test <cloud domain=\"...\"> - fail - nil", func(t *testing.T) {
// 		var r Cloud = Cloud{
// 			XMLName:           xml.Name{Space: "", Local: "cloud"},
// 			Domain:            nil,
// 			Port:              Ptr("1337"),
// 			Path:              Ptr("/path"),
// 			RegisterProcedure: Ptr("procedure"),
// 			Protocol:          Ptr("xml-rpc"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 1, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrInvalidElement)
// 		assert.ErrorContains(t, errs[0], "Attribute 'domain' of <cloud> is required")
// 	})
// 	t.Run("test <cloud domain=\"...\"> - fail - empty value", func(t *testing.T) {
// 		var r Cloud = Cloud{
// 			XMLName:           xml.Name{Space: "", Local: "cloud"},
// 			Domain:            Ptr(""),
// 			Port:              Ptr("1337"),
// 			Path:              Ptr("/path"),
// 			RegisterProcedure: Ptr("procedure"),
// 			Protocol:          Ptr("xml-rpc"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 1, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrEmptyValue)
// 		assert.ErrorContains(t, errs[0], "Attribute 'domain' of <cloud> value '' is invalid")
// 	})
// 	t.Run("test <cloud port=\"...\"> - fail - nil", func(t *testing.T) {
// 		var r Cloud = Cloud{
// 			XMLName:           xml.Name{Space: "", Local: "cloud"},
// 			Domain:            Ptr("rpc.example.com"),
// 			Port:              nil,
// 			Path:              Ptr("/path"),
// 			RegisterProcedure: Ptr("procedure"),
// 			Protocol:          Ptr("xml-rpc"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 1, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrInvalidElement)
// 		assert.ErrorContains(t, errs[0], "Attribute 'port' of <cloud> is required")
// 	})
// 	t.Run("test <cloud port=\"...\"> - fail - invalid value", func(t *testing.T) {
// 		var r Cloud = Cloud{
// 			XMLName:           xml.Name{Space: "", Local: "cloud"},
// 			Domain:            Ptr("rpc.example.com"),
// 			Port:              Ptr("-1"),
// 			Path:              Ptr("/path"),
// 			RegisterProcedure: Ptr("procedure"),
// 			Protocol:          Ptr("xml-rpc"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 1, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrInvalidValue)
// 		assert.ErrorContains(t, errs[0], "Attribute 'port' of <cloud> value '-1' is invalid")
// 	})
// 	t.Run("test <cloud path=\"...\"> - fail - nil", func(t *testing.T) {
// 		var r Cloud = Cloud{
// 			XMLName:           xml.Name{Space: "", Local: "cloud"},
// 			Domain:            Ptr("rpc.example.com"),
// 			Port:              Ptr("1337"),
// 			Path:              nil,
// 			RegisterProcedure: Ptr("procedure"),
// 			Protocol:          Ptr("xml-rpc"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 1, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrInvalidElement)
// 		assert.ErrorContains(t, errs[0], "Attribute 'path' of <cloud> is required")
// 	})
// 	t.Run("test <cloud path=\"...\"> - fail - empty value", func(t *testing.T) {
// 		var r Cloud = Cloud{
// 			XMLName:           xml.Name{Space: "", Local: "cloud"},
// 			Domain:            Ptr("rpc.example.com"),
// 			Port:              Ptr("1337"),
// 			Path:              Ptr(""),
// 			RegisterProcedure: Ptr("procedure"),
// 			Protocol:          Ptr("xml-rpc"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 1, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrEmptyValue)
// 		assert.ErrorContains(t, errs[0], "Attribute 'path' of <cloud> value '' is invalid")
// 	})
// 	t.Run("test <cloud registerProcedure=\"...\"> - fail - nil", func(t *testing.T) {
// 		var r Cloud = Cloud{
// 			XMLName:           xml.Name{Space: "", Local: "cloud"},
// 			Domain:            Ptr("rpc.example.com"),
// 			Port:              Ptr("1337"),
// 			Path:              Ptr("/path"),
// 			RegisterProcedure: nil,
// 			Protocol:          Ptr("xml-rpc"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 1, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrInvalidElement)
// 		assert.ErrorContains(t, errs[0], "Attribute 'registerProcedure' of <cloud> is required")
// 	})
// 	t.Run("test <cloud registerProcedure=\"...\"> - fail - empty value", func(t *testing.T) {
// 		var r Cloud = Cloud{
// 			XMLName:           xml.Name{Space: "", Local: "cloud"},
// 			Domain:            Ptr("rpc.example.com"),
// 			Port:              Ptr("1337"),
// 			Path:              Ptr("/path"),
// 			RegisterProcedure: Ptr(""),
// 			Protocol:          Ptr("xml-rpc"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 1, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrEmptyValue)
// 		assert.ErrorContains(t, errs[0], "Attribute 'registerProcedure' of <cloud> value '' is invalid")
// 	})
// 	t.Run("test <cloud protocol=\"...\"> - fail - nil", func(t *testing.T) {
// 		var r Cloud = Cloud{
// 			XMLName:           xml.Name{Space: "", Local: "cloud"},
// 			Domain:            Ptr("rpc.example.com"),
// 			Port:              Ptr("1337"),
// 			Path:              Ptr("/path"),
// 			RegisterProcedure: Ptr("procedure"),
// 			Protocol:          nil,
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 1, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrInvalidElement)
// 		assert.ErrorContains(t, errs[0], "Attribute 'protocol' of <cloud> is required")
// 	})
// 	t.Run("test <cloud port=\"...\"> - fail - invalid value", func(t *testing.T) {
// 		var r Cloud = Cloud{
// 			XMLName:           xml.Name{Space: "", Local: "cloud"},
// 			Domain:            Ptr("rpc.example.com"),
// 			Port:              Ptr("1337"),
// 			Path:              Ptr("/path"),
// 			RegisterProcedure: Ptr("procedure"),
// 			Protocol:          Ptr("bad value"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 1, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrInvalidValue)
// 		assert.ErrorContains(t, errs[0], "Attribute 'protocol' of <cloud> value 'bad value' is invalid")
// 	})
// 	t.Run("test <cloud> - unmarshal", func(t *testing.T) {
// 		var r Cloud
// 		s := []byte(`<cloud domain="rpc.example.com" port="1337" path="/path" registerProcedure="procedure" protocol="xml-rpc" />`)
// 		err := xml.Unmarshal(s, &r)
// 		assert.Equal(t, "", string(r.CharData))
// 		assert.Equal(t, "rpc.example.com", *r.Domain)
// 		assert.Equal(t, "1337", *r.Port)
// 		assert.Equal(t, "/path", *r.Path)
// 		assert.Equal(t, "procedure", *r.RegisterProcedure)
// 		assert.Equal(t, "xml-rpc", *r.Protocol)
// 		assert.Nil(t, err)
// 	})
// 	t.Run("test <cloud> - marshal", func(t *testing.T) {
// 		var r Cloud = Cloud{
// 			XMLName:           xml.Name{Space: "", Local: "cloud"},
// 			Domain:            Ptr("rpc.example.com"),
// 			Port:              Ptr("1337"),
// 			Path:              Ptr("/path"),
// 			RegisterProcedure: Ptr("procedure"),
// 			Protocol:          Ptr("xml-rpc"),
// 		}
// 		// NOTE: In XML and XHTML, a self-closing tag is a shorthand notation for
// 		// an opening and closing tag in one. It's used to communicate lack of
// 		// content in between the opening and closing tags. So, rather than typing
// 		// <cloud></cloud> (with no space at all in between), you'd be able write
// 		// <cloud/>.
// 		exp := []byte(`<cloud domain="rpc.example.com" port="1337" path="/path" registerProcedure="procedure" protocol="xml-rpc"></cloud>`)
// 		s, err := xml.Marshal(r)
// 		assert.Equal(t, exp, s)
// 		assert.Nil(t, err)
// 		ret, errs := Validate(r)
// 		assert.True(t, ret)
// 		assert.Empty(t, errs)
// 	})
// }
//
// func TestTTL(t *testing.T) {
// 	t.Run("test <ttl> - ok", func(t *testing.T) {
// 		var r TTL = TTL{
// 			XMLName:  xml.Name{Space: "", Local: "ttl"},
// 			CharData: []byte("60"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.True(t, ret)
// 		assert.Empty(t, errs)
// 	})
// 	t.Run("test <ttl> - fail - empty", func(t *testing.T) {
// 		var r TTL = TTL{
// 			XMLName:  xml.Name{Space: "", Local: "ttl"},
// 			CharData: []byte(""),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 2, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrEmptyValue)
// 		assert.ErrorContains(t, errs[0], "Element <ttl> value '' is invalid")
// 		assert.ErrorIs(t, errs[1], ErrInvalidValue)
// 		assert.ErrorContains(t, errs[1], "Element <ttl> value '' is invalid")
// 	})
// 	t.Run("test <ttl> - fail - invalid value", func(t *testing.T) {
// 		var r TTL = TTL{
// 			XMLName:  xml.Name{Space: "", Local: "ttl"},
// 			CharData: []byte("-1"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 1, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrInvalidValue)
// 		assert.ErrorContains(t, errs[0], "Element <ttl> value '-1' is invalid")
// 	})
// 	t.Run("test <ttl> - unmarshal", func(t *testing.T) {
// 		var r TTL
// 		s := []byte(`<ttl>60</ttl>`)
// 		err := xml.Unmarshal(s, &r)
// 		assert.Equal(t, "60", string(r.CharData))
// 		assert.Nil(t, err)
// 	})
// 	t.Run("test <ttl> - marshal", func(t *testing.T) {
// 		var r TTL = TTL{
// 			XMLName:  xml.Name{Space: "", Local: "ttl"},
// 			CharData: []byte("60"),
// 		}
// 		exp := []byte(`<ttl>60</ttl>`)
// 		s, err := xml.Marshal(r)
// 		assert.Equal(t, exp, s)
// 		assert.Nil(t, err)
// 		ret, errs := Validate(r)
// 		assert.True(t, ret)
// 		assert.Empty(t, errs)
// 	})
// }
//
// func TestTextInput(t *testing.T) {
// 	t.Run("test <textInput> - ok", func(t *testing.T) {
// 		var r TextInput = TextInput{
// 			XMLName: xml.Name{Space: "", Local: "textInput"},
// 			Title: &Title{
// 				XMLName:  xml.Name{Space: "", Local: "title"},
// 				CharData: []byte("Title"),
// 			},
// 			Description: &Description{
// 				XMLName:  xml.Name{Space: "", Local: "description"},
// 				CharData: []byte("Description"),
// 			},
// 			Name: &Name{
// 				XMLName:  xml.Name{Space: "", Local: "name"},
// 				CharData: []byte("Name"),
// 			},
// 			Link: &Link{
// 				XMLName:  xml.Name{Space: "", Local: "link"},
// 				CharData: []byte("https://example.com/search"),
// 			},
// 		}
// 		ret, errs := r.IsValid()
// 		assert.True(t, ret)
// 		assert.Empty(t, errs)
// 	})
// 	t.Run("test <textInput> - fail - nil", func(t *testing.T) {
// 		var r TextInput = TextInput{
// 			XMLName:     xml.Name{Space: "", Local: "textInput"},
// 			Title:       nil,
// 			Description: nil,
// 			Name:        nil,
// 			Link:        nil,
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 1, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrInvalidElement)
// 		assert.ErrorContains(t, errs[0], "Element <textInput> is invalid")
// 	})
// 	t.Run("test <textInput> - unmarshal", func(t *testing.T) {
// 		var r TextInput
// 		s := []byte(`<textInput><title>Title</title><description>Description</description><name>Name</name><link>https://example.com/search</link></textInput>`)
// 		err := xml.Unmarshal(s, &r)
// 		assert.Equal(t, "Title", string(r.Title.CharData))
// 		assert.Equal(t, "Description", string(r.Description.CharData))
// 		assert.Equal(t, "Name", string(r.Name.CharData))
// 		assert.Equal(t, "https://example.com/search", string(r.Link.CharData))
// 		assert.Nil(t, err)
// 	})
// 	t.Run("test <textInput> - marshal", func(t *testing.T) {
// 		var r TextInput = TextInput{
// 			XMLName: xml.Name{Space: "", Local: "textInput"},
// 			Title: &Title{
// 				XMLName:  xml.Name{Space: "", Local: "title"},
// 				CharData: []byte("Title"),
// 			},
// 			Description: &Description{
// 				XMLName:  xml.Name{Space: "", Local: "description"},
// 				CharData: []byte("Description"),
// 			},
// 			Name: &Name{
// 				XMLName:  xml.Name{Space: "", Local: "name"},
// 				CharData: []byte("Name"),
// 			},
// 			Link: &Link{
// 				XMLName:  xml.Name{Space: "", Local: "link"},
// 				CharData: []byte("https://example.com/search"),
// 			},
// 		}
// 		exp := []byte(`<textInput><title>Title</title><description>Description</description><name>Name</name><link>https://example.com/search</link></textInput>`)
// 		s, err := xml.Marshal(r)
// 		assert.Equal(t, exp, s)
// 		assert.Nil(t, err)
// 		ret, errs := Validate(r)
// 		assert.True(t, ret)
// 		assert.Empty(t, errs)
// 	})
// }
//
// func TestName(t *testing.T) {
// 	t.Run("test <name> - ok", func(t *testing.T) {
// 		var r Name = Name{
// 			XMLName:  xml.Name{Space: "", Local: "name"},
// 			CharData: []byte("Name"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.True(t, ret)
// 		assert.Empty(t, errs)
// 	})
// 	t.Run("test <name> - fail - empty", func(t *testing.T) {
// 		var r Name = Name{
// 			XMLName:  xml.Name{Space: "", Local: "name"},
// 			CharData: []byte(""),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 1, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrEmptyValue)
// 		assert.ErrorContains(t, errs[0], "Element <name> value '' is invalid")
// 	})
// 	t.Run("test <name> - unmarshal", func(t *testing.T) {
// 		var r Name
// 		s := []byte(`<name>Name</name>`)
// 		err := xml.Unmarshal(s, &r)
// 		assert.Equal(t, "Name", string(r.CharData))
// 		assert.Nil(t, err)
// 	})
// 	t.Run("test <name> - marshal", func(t *testing.T) {
// 		var r Name = Name{
// 			XMLName:  xml.Name{Space: "", Local: "name"},
// 			CharData: []byte("Name"),
// 		}
// 		exp := []byte(`<name>Name</name>`)
// 		s, err := xml.Marshal(r)
// 		assert.Equal(t, exp, s)
// 		assert.Nil(t, err)
// 		ret, errs := Validate(r)
// 		assert.True(t, ret)
// 		assert.Empty(t, errs)
// 	})
// }
//
// func TestPubDate(t *testing.T) {
// 	t.Run("test <pubDate> - ok", func(t *testing.T) {
// 		var r PubDate
// 		r = PubDate{
// 			XMLName:  xml.Name{Space: "", Local: "pubDate"},
// 			CharData: []byte("01 Jan 70 00:00 GMT"), // RFC822
// 		}
// 		ret, errs := r.IsValid()
// 		assert.True(t, ret)
// 		assert.Empty(t, errs)
// 		r = PubDate{
// 			XMLName:  xml.Name{Space: "", Local: "pubDate"},
// 			CharData: []byte("Thu, 01 Jan 1970 00:00:00 GMT"), // RFC1123
// 		}
// 		ret, errs = r.IsValid()
// 		assert.True(t, ret)
// 		assert.Empty(t, errs)
// 	})
// 	t.Run("test <pubDate> - fail - empty", func(t *testing.T) {
// 		var r = PubDate{
// 			XMLName:  xml.Name{Space: "", Local: "pubDate"},
// 			CharData: []byte(""),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.Equal(t, 2, len(errs))
// 		assert.False(t, ret)
// 		assert.ErrorIs(t, errs[0], ErrEmptyValue)
// 		assert.ErrorContains(t, errs[0], "Element <pubDate> value '' is invalid")
// 		assert.ErrorIs(t, errs[1], ErrInvalidDate)
// 		assert.ErrorContains(t, errs[1], "Element <pubDate> value '' is invalid")
// 	})
// 	t.Run("test <pubDate> - fail - invalid date", func(t *testing.T) {
// 		var r = PubDate{
// 			XMLName:  xml.Name{Space: "", Local: "pubDate"},
// 			CharData: []byte("bad date"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.Equal(t, 1, len(errs))
// 		assert.False(t, ret)
// 		assert.ErrorIs(t, errs[0], ErrInvalidDate)
// 		assert.ErrorContains(t, errs[0], "Element <pubDate> value 'bad date' is invalid")
// 	})
// 	t.Run("test <pubDate> - unmarshal", func(t *testing.T) {
// 		var r PubDate
// 		s := []byte(`<pubDate>Thu, 01 Jan 1970 00:00:00 GMT</pubDate>`)
// 		err := xml.Unmarshal(s, &r)
// 		assert.Equal(t, "Thu, 01 Jan 1970 00:00:00 GMT", string(r.CharData))
// 		assert.Nil(t, err)
// 	})
// 	t.Run("test <pubDate> - marshal", func(t *testing.T) {
// 		var r PubDate = PubDate{
// 			XMLName:  xml.Name{Space: "", Local: "pubDate"},
// 			CharData: []byte("Thu, 01 Jan 1970 00:00:00 GMT"),
// 		}
// 		exp := []byte(`<pubDate>Thu, 01 Jan 1970 00:00:00 GMT</pubDate>`)
// 		s, err := xml.Marshal(r)
// 		assert.Equal(t, exp, s)
// 		assert.Nil(t, err)
// 		ret, errs := Validate(r)
// 		assert.True(t, ret)
// 		assert.Empty(t, errs)
// 	})
// }
//
// func TestLastBuildDate(t *testing.T) {
// 	t.Run("test <lastBuildDate> - ok", func(t *testing.T) {
// 		var r LastBuildDate
// 		r = LastBuildDate{
// 			XMLName:  xml.Name{Space: "", Local: "lastBuildDate"},
// 			CharData: []byte("01 Jan 70 00:00 GMT"), // RFC822
// 		}
// 		ret, errs := r.IsValid()
// 		assert.True(t, ret)
// 		assert.Empty(t, errs)
// 		r = LastBuildDate{
// 			XMLName:  xml.Name{Space: "", Local: "lastBuildDate"},
// 			CharData: []byte("Thu, 01 Jan 1970 00:00:00 GMT"), // RFC1123
// 		}
// 		ret, errs = r.IsValid()
// 		assert.True(t, ret)
// 		assert.Empty(t, errs)
// 	})
// 	t.Run("test <lastBuildDate> - fail - empty", func(t *testing.T) {
// 		var r = LastBuildDate{
// 			XMLName:  xml.Name{Space: "", Local: "lastBuildDate"},
// 			CharData: []byte(""),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.Equal(t, 2, len(errs))
// 		assert.False(t, ret)
// 		assert.ErrorIs(t, errs[0], ErrEmptyValue)
// 		assert.ErrorContains(t, errs[0], "Element <lastBuildDate> value '' is invalid")
// 		assert.ErrorIs(t, errs[1], ErrInvalidDate)
// 		assert.ErrorContains(t, errs[1], "Element <lastBuildDate> value '' is invalid")
// 	})
// 	t.Run("test <lastBuildDate> - fail - invalid date", func(t *testing.T) {
// 		var r = LastBuildDate{
// 			XMLName:  xml.Name{Space: "", Local: "lastBuildDate"},
// 			CharData: []byte("bad date"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.Equal(t, 1, len(errs))
// 		assert.False(t, ret)
// 		assert.ErrorIs(t, errs[0], ErrInvalidDate)
// 		assert.ErrorContains(t, errs[0], "Element <lastBuildDate> value 'bad date' is invalid")
// 	})
// 	t.Run("test <lastBuildDate> - unmarshal", func(t *testing.T) {
// 		var r LastBuildDate
// 		s := []byte(`<lastBuildDate>Thu, 01 Jan 1970 00:00:00 GMT</lastBuildDate>`)
// 		err := xml.Unmarshal(s, &r)
// 		assert.Equal(t, "Thu, 01 Jan 1970 00:00:00 GMT", string(r.CharData))
// 		assert.Nil(t, err)
// 	})
// 	t.Run("test <lastBuildDate> - marshal", func(t *testing.T) {
// 		var r LastBuildDate = LastBuildDate{
// 			XMLName:  xml.Name{Space: "", Local: "lastBuildDate"},
// 			CharData: []byte("Thu, 01 Jan 1970 00:00:00 GMT"),
// 		}
// 		exp := []byte(`<lastBuildDate>Thu, 01 Jan 1970 00:00:00 GMT</lastBuildDate>`)
// 		s, err := xml.Marshal(r)
// 		assert.Equal(t, exp, s)
// 		assert.Nil(t, err)
// 		ret, errs := Validate(r)
// 		assert.True(t, ret)
// 		assert.Empty(t, errs)
// 	})
// }
//
// func TestCategory(t *testing.T) {
// 	t.Run("test <category> - ok", func(t *testing.T) {
// 		var r Category = Category{
// 			XMLName:  xml.Name{Space: "", Local: "category"},
// 			CharData: []byte(`Category`),
// 			Domain:   Ptr("https://example.com/category"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.True(t, ret)
// 		assert.Empty(t, errs)
// 	})
// 	t.Run("test <category> - fail - empty", func(t *testing.T) {
// 		var r Category = Category{
// 			XMLName:  xml.Name{Space: "", Local: "category"},
// 			CharData: []byte(``),
// 			Domain:   Ptr("https://example.com/category"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 1, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrEmptyValue)
// 		assert.ErrorContains(t, errs[0], "Element <category> value '' is invalid")
// 	})
// 	t.Run("test <category domain=\"...\"> - fail - empty", func(t *testing.T) {
// 		var r Category = Category{
// 			XMLName:  xml.Name{Space: "", Local: "category"},
// 			CharData: []byte(`Category`),
// 			Domain:   Ptr(""),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 1, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrEmptyValue)
// 		assert.ErrorContains(t, errs[0], "Attribute 'domain' of <category> value '' is invalid")
// 	})
// 	t.Run("test <category> - unmarshal", func(t *testing.T) {
// 		var r Category
// 		s := []byte(`<category domain="https://example.com/category">Category</category>`)
// 		err := xml.Unmarshal(s, &r)
// 		assert.Equal(t, "Category", string(r.CharData))
// 		assert.Equal(t, "https://example.com/category", *r.Domain)
// 		assert.Nil(t, err)
// 	})
// 	t.Run("test <category> - marshal", func(t *testing.T) {
// 		var r Category = Category{
// 			XMLName:  xml.Name{Space: "", Local: "category"},
// 			CharData: []byte(`Category`),
// 			Domain:   Ptr("https://example.com/category"),
// 		}
// 		exp := []byte(`<category domain="https://example.com/category">Category</category>`)
// 		s, err := xml.Marshal(r)
// 		assert.Equal(t, exp, s)
// 		assert.Nil(t, err)
// 		ret, errs := Validate(r)
// 		assert.True(t, ret)
// 		assert.Empty(t, errs)
// 	})
// }
//
// func TestItem(t *testing.T) {
// 	t.Run("test <item> - ok", func(t *testing.T) {
// 		var r Item
// 		r = Item{
// 			Title: &Title{
// 				XMLName:  xml.Name{Space: "", Local: "title"},
// 				CharData: []byte("Title"),
// 			},
// 			Link: &Link{
// 				XMLName:  xml.Name{Space: "", Local: "link"},
// 				CharData: []byte("https://example.com"),
// 			},
// 			Description: &Description{
// 				XMLName:  xml.Name{Space: "", Local: "description"},
// 				CharData: []byte("Description"),
// 			},
// 			Source: &Source{
// 				XMLName:  xml.Name{Space: "", Local: "source"},
// 				CharData: []byte("Title"),
// 				URL:      Ptr("https://example.com/source.xml"),
// 			},
// 			Enclosure: &Enclosure{
// 				XMLName:  xml.Name{Space: "", Local: "enclosure"},
// 				CharData: []byte(""),
// 				URL:      Ptr("https://example.com/audio.mp3"),
// 				Length:   Ptr("1337"),
// 				Type:     Ptr("audio/mpeg"),
// 			},
// 			Category: &Category{
// 				XMLName:  xml.Name{Space: "", Local: "category"},
// 				CharData: []byte("Category"),
// 				Domain:   Ptr("https://example.com/category"),
// 			},
// 			PubDate: &PubDate{
// 				XMLName:  xml.Name{Space: "", Local: "pubDate"},
// 				CharData: []byte("Thu, 01 Jan 1970 00:00:00 GMT"),
// 			},
// 			GUID: &GUID{
// 				XMLName:     xml.Name{Space: "", Local: "guid"},
// 				CharData:    []byte("https://example.com/guid"),
// 				IsPermaLink: Ptr("true"),
// 			},
// 			Comments: &Comments{
// 				XMLName:  xml.Name{Space: "", Local: "comments"},
// 				CharData: []byte("https://example.com/comments"),
// 			},
// 			Author: &Author{
// 				XMLName:  xml.Name{Space: "", Local: "author"},
// 				CharData: []byte("first.last@example.com"),
// 			},
// 		}
// 		ret, errs := r.IsValid()
// 		assert.True(t, ret)
// 		assert.Empty(t, errs)
// 	})
// 	t.Run("test <item> - fail - invalid element", func(t *testing.T) {
// 		var r Item
// 		r = Item{
// 			XMLName: xml.Name{Space: "", Local: "item"},
// 			Title: &Title{
// 				XMLName:  xml.Name{Space: "", Local: "title"},
// 				CharData: []byte(""),
// 			},
// 			Link: &Link{
// 				XMLName:  xml.Name{Space: "", Local: "link"},
// 				CharData: []byte("https://example.com"),
// 			},
// 			Description: &Description{
// 				XMLName:  xml.Name{Space: "", Local: "description"},
// 				CharData: []byte(""),
// 			},
// 			Source: &Source{
// 				XMLName:  xml.Name{Space: "", Local: "source"},
// 				CharData: []byte("Title"),
// 				URL:      Ptr("https://example.com/source.xml"),
// 			},
// 			Enclosure: &Enclosure{
// 				XMLName:  xml.Name{Space: "", Local: "enclosure"},
// 				CharData: []byte(""),
// 				URL:      Ptr("https://example.com/audio.mp3"),
// 				Length:   Ptr("1337"),
// 				Type:     Ptr("audio/mpeg"),
// 			},
// 			Category: &Category{
// 				XMLName:  xml.Name{Space: "", Local: "category"},
// 				CharData: []byte("Category"),
// 				Domain:   Ptr("https://example.com/category"),
// 			},
// 			PubDate: &PubDate{
// 				XMLName:  xml.Name{Space: "", Local: "pubDate"},
// 				CharData: []byte("Thu, 01 Jan 1970 00:00:00 GMT"),
// 			},
// 			GUID: &GUID{
// 				XMLName:     xml.Name{Space: "", Local: "guid"},
// 				CharData:    []byte("https://example.com/guid"),
// 				IsPermaLink: Ptr("true"),
// 			},
// 			Comments: &Comments{
// 				XMLName:  xml.Name{Space: "", Local: "comments"},
// 				CharData: []byte("https://example.com/comments"),
// 			},
// 			Author: &Author{
// 				XMLName:  xml.Name{Space: "", Local: "author"},
// 				CharData: []byte("first.last@example.com"),
// 			},
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 3, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrInvalidElement)
// 		assert.ErrorContains(t, errs[0], "Element <item> is invalid")
// 		assert.ErrorIs(t, errs[1], ErrEmptyValue)
// 		assert.ErrorContains(t, errs[1], "Element <title> value '' is invalid")
// 		assert.ErrorIs(t, errs[2], ErrEmptyValue)
// 		assert.ErrorContains(t, errs[2], "Element <description> value '' is invalid")
// 	})
// 	t.Run("test <item> - unmarshal", func(t *testing.T) {
// 		var r Item
// 		s := []byte(`<item>
// 	<title>Title</title>
// 	<link>https://example.com</link>
// 	<description>Description</description>
// 	<source url="https://example.com/source.xml">Title</source>
//   <enclosure url="https://example.com/audio.mp3" length="1337" type="audio/mpeg"/>
// 	<category domain="https://example.com/category">Category</category>
// 	<pubDate>Thu, 01 Jan 1970 00:00:00 GMT</pubDate>
// 	<guid isPermaLink="true">https://example.com/guid</guid>
// 	<comments>https://example.com/comments</comments>
// 	<author>first.last@example.com</author>
// </item>
// `)
// 		err := xml.Unmarshal(s, &r)
// 		assert.Equal(t, "Title", string(r.Title.CharData))
// 		assert.Equal(t, "https://example.com", string(r.Link.CharData))
// 		assert.Equal(t, "Description", string(r.Description.CharData))
// 		assert.Equal(t, "https://example.com/source.xml", *r.Source.URL)
// 		assert.Equal(t, "", string(r.Enclosure.CharData))
// 		assert.Equal(t, "https://example.com/audio.mp3", *r.Enclosure.URL)
// 		assert.Equal(t, "1337", *r.Enclosure.Length)
// 		assert.Equal(t, "audio/mpeg", *r.Enclosure.Type)
// 		assert.Equal(t, "Category", string(r.Category.CharData))
// 		assert.Equal(t, "https://example.com/category", *r.Category.Domain)
// 		assert.Equal(t, "Thu, 01 Jan 1970 00:00:00 GMT", string(r.PubDate.CharData))
// 		assert.Equal(t, "https://example.com/guid", string(r.GUID.CharData))
// 		assert.Equal(t, "true", *r.GUID.IsPermaLink)
// 		assert.Equal(t, "https://example.com/comments", string(r.Comments.CharData))
// 		assert.Equal(t, "first.last@example.com", string(r.Author.CharData))
// 		assert.Nil(t, err)
// 		ret, errs := Validate(r)
// 		assert.True(t, ret)
// 		assert.Empty(t, errs)
// 	})
// 	t.Run("test <item> - unmarshal", func(t *testing.T) {
// 		var r Item
// 		s := []byte(`<item>
// 	<title>Title</title>
// 	<description>Description</description>
// </item>
// `)
// 		err := xml.Unmarshal(s, &r)
// 		assert.Equal(t, "Title", string(r.Title.CharData))
// 		assert.Equal(t, "Description", string(r.Description.CharData))
// 		assert.Nil(t, err)
// 		ret, errs := Validate(r)
// 		assert.True(t, ret)
// 		assert.Empty(t, errs)
// 	})
// 	t.Run("test <item> - marshal", func(t *testing.T) {
// 		var r = Item{
// 			Title: &Title{
// 				XMLName:  xml.Name{Space: "", Local: "title"},
// 				CharData: []byte("Title"),
// 			},
// 			Link: &Link{
// 				XMLName:  xml.Name{Space: "", Local: "link"},
// 				CharData: []byte("https://example.com"),
// 			},
// 			Description: &Description{
// 				XMLName:  xml.Name{Space: "", Local: "description"},
// 				CharData: []byte("Description"),
// 			},
// 			Source: &Source{
// 				XMLName:  xml.Name{Space: "", Local: "source"},
// 				CharData: []byte("Title"),
// 				URL:      Ptr("https://example.com/source.xml"),
// 			},
// 			Enclosure: &Enclosure{
// 				XMLName:  xml.Name{Space: "", Local: "enclosure"},
// 				CharData: []byte(""),
// 				URL:      Ptr("https://example.com/audio.mp3"),
// 				Length:   Ptr("1337"),
// 				Type:     Ptr("audio/mpeg"),
// 			},
// 			Category: &Category{
// 				XMLName:  xml.Name{Space: "", Local: "category"},
// 				CharData: []byte("Category"),
// 				Domain:   Ptr("https://example.com/category"),
// 			},
// 			PubDate: &PubDate{
// 				XMLName:  xml.Name{Space: "", Local: "pubDate"},
// 				CharData: []byte("Thu, 01 Jan 1970 00:00:00 GMT"),
// 			},
// 			GUID: &GUID{
// 				XMLName:     xml.Name{Space: "", Local: "guid"},
// 				CharData:    []byte("https://example.com/guid"),
// 				IsPermaLink: Ptr("true"),
// 			},
// 			Comments: &Comments{
// 				XMLName:  xml.Name{Space: "", Local: "comments"},
// 				CharData: []byte("https://example.com/comments"),
// 			},
// 			Author: &Author{
// 				XMLName:  xml.Name{Space: "", Local: "author"},
// 				CharData: []byte("first.last@example.com"),
// 			},
// 		}
// 		exp := []byte(`<item><title>Title</title><link>https://example.com</link><description>Description</description><source url="https://example.com/source.xml">Title</source><enclosure url="https://example.com/audio.mp3" length="1337" type="audio/mpeg"></enclosure><category domain="https://example.com/category">Category</category><pubDate>Thu, 01 Jan 1970 00:00:00 GMT</pubDate><guid isPermaLink="true">https://example.com/guid</guid><comments>https://example.com/comments</comments><author>first.last@example.com</author></item>`)
// 		s, err := xml.Marshal(r)
// 		assert.Equal(t, exp, s)
// 		assert.Nil(t, err)
// 		ret, errs := Validate(r)
// 		assert.True(t, ret)
// 		assert.Empty(t, errs)
// 	})
// }
//
// func TestSource(t *testing.T) {
// 	t.Run("test <source> - ok", func(t *testing.T) {
// 		var r = Source{
// 			XMLName:  xml.Name{Space: "", Local: "source"},
// 			CharData: []byte("Title"),
// 			URL:      Ptr("https://example.com/source.xml"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.True(t, ret)
// 		assert.Empty(t, errs)
// 	})
// 	t.Run("test <source> - ok - empty", func(t *testing.T) {
// 		// NOTE: <source> can be empty.
// 		var r = Source{
// 			XMLName:  xml.Name{Space: "", Local: "source"},
// 			CharData: []byte(""),
// 			URL:      Ptr("https://example.com/source.xml"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.True(t, ret)
// 		assert.Empty(t, errs)
// 	})
// 	t.Run("test <source url=\"...\"> - fail - nil", func(t *testing.T) {
// 		var r = Source{
// 			XMLName:  xml.Name{Space: "", Local: "source"},
// 			CharData: []byte("Title"),
// 			URL:      nil,
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 1, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrInvalidElement)
// 		assert.ErrorContains(t, errs[0], "Attribute 'url' of <source> is required")
// 	})
// 	t.Run("test <source url=\"...\"> - fail - empty", func(t *testing.T) {
// 		var r = Source{
// 			XMLName:  xml.Name{Space: "", Local: "source"},
// 			CharData: []byte("Title"),
// 			URL:      Ptr(""),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 2, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrEmptyValue)
// 		assert.ErrorContains(t, errs[0], "Attribute 'url' of <source> value '' is invalid")
// 	})
// 	t.Run("test <source url=\"...\"> - fail - invalid uri", func(t *testing.T) {
// 		var r = Source{
// 			XMLName:  xml.Name{Space: "", Local: "source"},
// 			CharData: []byte("Title"),
// 			URL:      Ptr("bad uri"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 1, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrInvalidURI)
// 		assert.ErrorContains(t, errs[0], "Attribute 'url' of <source> value 'bad uri' is invalid")
// 	})
// 	t.Run("test <source> - unmarshal", func(t *testing.T) {
// 		var r Source
// 		s := []byte(`<source url="https://example.com/source.xml">Title</source>`)
// 		err := xml.Unmarshal(s, &r)
// 		assert.Equal(t, "Title", string(r.CharData))
// 		assert.Equal(t, "https://example.com/source.xml", *r.URL)
// 		assert.Nil(t, err)
// 	})
// 	t.Run("test <source> - marshal", func(t *testing.T) {
// 		var r = Source{
// 			XMLName:  xml.Name{Space: "", Local: "source"},
// 			CharData: []byte("Title"),
// 			URL:      Ptr("https://example.com/source.xml"),
// 		}
// 		exp := []byte(`<source url="https://example.com/source.xml">Title</source>`)
// 		s, err := xml.Marshal(r)
// 		assert.Equal(t, exp, s)
// 		assert.Nil(t, err)
// 		ret, errs := Validate(r)
// 		assert.True(t, ret)
// 		assert.Empty(t, errs)
// 	})
// }
//
// func TestEnclosure(t *testing.T) {
// 	t.Run("test <enclosure> - ok", func(t *testing.T) {
// 		var r Enclosure = Enclosure{
// 			XMLName:  xml.Name{Space: "", Local: "enclosure"},
// 			CharData: []byte(""),
// 			URL:      Ptr("https://example.com/audio.mp3"),
// 			Length:   Ptr("1337"),
// 			Type:     Ptr("audio/mpeg"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.True(t, ret)
// 		assert.Empty(t, errs)
// 	})
// 	t.Run("test <enclosure> - fail - not empty", func(t *testing.T) {
// 		var r Enclosure = Enclosure{
// 			XMLName:  xml.Name{Space: "", Local: "enclosure"},
// 			CharData: []byte("not empty"),
// 			URL:      Ptr("https://example.com/audio.mp3"),
// 			Length:   Ptr("1337"),
// 			Type:     Ptr("audio/mpeg"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 1, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrNonEmptyValue)
// 		assert.ErrorContains(t, errs[0], "Element <enclosure> is invalid")
// 	})
// 	t.Run("test <enclosure url=\"...\"> - fail - nil", func(t *testing.T) {
// 		var r Enclosure = Enclosure{
// 			XMLName:  xml.Name{Space: "", Local: "enclosure"},
// 			CharData: []byte(""),
// 			URL:      nil,
// 			Length:   Ptr("1337"),
// 			Type:     Ptr("audio/mpeg"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 1, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrInvalidElement)
// 		assert.ErrorContains(t, errs[0], "Attribute 'url' of <enclosure> is required")
// 	})
// 	t.Run("test <enclosure url=\"...\"> - fail - empty", func(t *testing.T) {
// 		var r Enclosure = Enclosure{
// 			XMLName:  xml.Name{Space: "", Local: "enclosure"},
// 			CharData: []byte(""),
// 			URL:      Ptr(""),
// 			Length:   Ptr("1337"),
// 			Type:     Ptr("audio/mpeg"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 2, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrEmptyValue)
// 		assert.ErrorContains(t, errs[0], "Attribute 'url' of <enclosure> value '' is invalid")
// 	})
// 	t.Run("test <enclosure url=\"...\"> - fail - invalid uri", func(t *testing.T) {
// 		var r Enclosure = Enclosure{
// 			XMLName:  xml.Name{Space: "", Local: "enclosure"},
// 			CharData: []byte(""),
// 			URL:      Ptr("bad uri"),
// 			Length:   Ptr("1337"),
// 			Type:     Ptr("audio/mpeg"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 1, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrInvalidURI)
// 		assert.ErrorContains(t, errs[0], "Attribute 'url' of <enclosure> value 'bad uri' is invalid")
// 	})
// 	t.Run("test <enclosure length=\"...\"> - fail - nil", func(t *testing.T) {
// 		var r Enclosure = Enclosure{
// 			XMLName:  xml.Name{Space: "", Local: "enclosure"},
// 			CharData: []byte(""),
// 			URL:      Ptr("https://example.com/audio.mp3"),
// 			Length:   nil,
// 			Type:     Ptr("audio/mpeg"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 1, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrInvalidElement)
// 		assert.ErrorContains(t, errs[0], "Attribute 'length' of <enclosure> is required")
// 	})
// 	t.Run("test <enclosure length=\"...\"> - fail - invalid value", func(t *testing.T) {
// 		var r Enclosure = Enclosure{
// 			XMLName:  xml.Name{Space: "", Local: "enclosure"},
// 			CharData: []byte(""),
// 			URL:      Ptr("https://example.com/audio.mp3"),
// 			Length:   Ptr("-1"),
// 			Type:     Ptr("audio/mpeg"),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 1, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrInvalidValue)
// 		assert.ErrorContains(t, errs[0], "Attribute 'length' of <enclosure> value '-1' is invalid")
// 	})
// 	t.Run("test <enclosure type=\"...\"> - fail - nil", func(t *testing.T) {
// 		var r Enclosure = Enclosure{
// 			XMLName:  xml.Name{Space: "", Local: "enclosure"},
// 			CharData: []byte(""),
// 			URL:      Ptr("https://example.com/audio.mp3"),
// 			Length:   Ptr("1337"),
// 			Type:     nil,
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 1, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrInvalidElement)
// 		assert.ErrorContains(t, errs[0], "Attribute 'type' of <enclosure> is required")
// 	})
// 	t.Run("test <enclosure type=\"...\"> - fail - empty", func(t *testing.T) {
// 		var r Enclosure = Enclosure{
// 			XMLName:  xml.Name{Space: "", Local: "enclosure"},
// 			CharData: []byte(""),
// 			URL:      Ptr("https://example.com/audio.mp3"),
// 			Length:   Ptr("1337"),
// 			Type:     Ptr(""),
// 		}
// 		ret, errs := r.IsValid()
// 		assert.False(t, ret)
// 		assert.Equal(t, 1, len(errs))
// 		assert.ErrorIs(t, errs[0], ErrEmptyValue)
// 		assert.ErrorContains(t, errs[0], "Attribute 'type' of <enclosure> value '' is invalid")
// 	})
// 	t.Run("test <enclosure> - unmarshal", func(t *testing.T) {
// 		var r Enclosure
// 		s := []byte(`<enclosure url="https://example.com/audio.mp3" length="1337" type="audio/mpeg" />`)
// 		err := xml.Unmarshal(s, &r)
// 		assert.Equal(t, "", string(r.CharData))
// 		assert.Equal(t, "https://example.com/audio.mp3", *r.URL)
// 		assert.Equal(t, "1337", *r.Length)
// 		assert.Equal(t, "audio/mpeg", *r.Type)
// 		assert.Nil(t, err)
// 	})
// 	t.Run("test <enclosure> - marshal", func(t *testing.T) {
// 		var r Enclosure = Enclosure{
// 			XMLName:  xml.Name{Space: "", Local: "enclosure"},
// 			CharData: []byte(""),
// 			URL:      Ptr("https://example.com/audio.mp3"),
// 			Length:   Ptr("1337"),
// 			Type:     Ptr("audio/mpeg"),
// 		}
// 		// NOTE: In XML and XHTML, a self-closing tag is a shorthand notation for
// 		// an opening and closing tag in one. It's used to communicate lack of
// 		// content in between the opening and closing tags. So, rather than typing
// 		// <enclosure></enclosure> (with no space at all in between), you'd be able
// 		// write <enclosure/>.
// 		exp := []byte(`<enclosure url="https://example.com/audio.mp3" length="1337" type="audio/mpeg"></enclosure>`)
// 		s, err := xml.Marshal(r)
// 		assert.Equal(t, exp, s)
// 		assert.Nil(t, err)
// 		ret, errs := Validate(r)
// 		assert.True(t, ret)
// 		assert.Empty(t, errs)
// 	})
// }

type Testable interface {
	Test(t *testing.T)
}

type ElementTestCase[T RSSElement] struct {
	name              string
	wantIsValid       bool
	wantErrorIs       []error
	wantErrorContains []string
	r                 T
}

func (tc ElementTestCase[T]) Test(t *testing.T) {
	t.Run(tc.name, func(t *testing.T) {
		r := tc.r
		ret, errs := r.IsValid()
		assert.Equal(t, tc.wantIsValid, ret)
		for i, err := range errs {
			assert.ErrorIs(t, err, tc.wantErrorIs[i])
			assert.ErrorContains(t, err, tc.wantErrorContains[i])
		}
	})
}

func TestElement(t *testing.T) {
	cases := []Testable{
		// test <guid>
		ElementTestCase[GUID]{
			name:              "test <guid> - ok",
			wantIsValid:       true,
			wantErrorIs:       []error{},
			wantErrorContains: []string{},
			r: GUID{
				XMLName:     xml.Name{Space: "", Local: "guid"},
				CharData:    []byte("https://example.com/1337"),
				IsPermaLink: nil,
			},
		},
		ElementTestCase[GUID]{
			name:              "test <guid> - ok",
			wantIsValid:       true,
			wantErrorIs:       []error{},
			wantErrorContains: []string{},
			r: GUID{
				XMLName:     xml.Name{Space: "", Local: "guid"},
				CharData:    []byte("https://example.com/1337"),
				IsPermaLink: Ptr(IsPermaLink("true")),
			},
		},
		ElementTestCase[GUID]{
			name:              "test <guid> - ok",
			wantIsValid:       true,
			wantErrorIs:       []error{},
			wantErrorContains: []string{},
			r: GUID{
				XMLName:     xml.Name{Space: "", Local: "guid"},
				CharData:    []byte("1337"),
				IsPermaLink: Ptr(IsPermaLink("false")),
			},
		},
		ElementTestCase[GUID]{
			name:        "test <guid> - fail - empty",
			wantIsValid: false,
			wantErrorIs: []error{ErrEmptyValue, ErrInvalidURI},
			wantErrorContains: []string{
				"Element <guid> value '' is invalid: Element must not have empty " +
					"value",
				"Element <guid> value '' is invalid: Element must contain a valid " +
					"URI (RFC3986): parse \"\": empty url",
			},
			r: GUID{
				XMLName:     xml.Name{Space: "", Local: "guid"},
				CharData:    []byte(""),
				IsPermaLink: nil,
			},
		},
		ElementTestCase[GUID]{
			name:        "test <guid> - fail - invalid uri",
			wantIsValid: false,
			wantErrorIs: []error{ErrInvalidURI},
			wantErrorContains: []string{
				"Element <guid> value 'bad uri' is invalid: Element must contain a " +
					"valid URI (RFC3986): parse \"bad uri\": invalid URI for request",
			},
			r: GUID{
				XMLName:     xml.Name{Space: "", Local: "guid"},
				CharData:    []byte("bad uri"),
				IsPermaLink: nil,
			},
		},
		ElementTestCase[GUID]{
			name:        "test <guid isPermaLink=\"...\"> - fail - empty",
			wantIsValid: false,
			wantErrorIs: []error{ErrInvalidValue},
			wantErrorContains: []string{
				"Attribute 'isPermaLink' of <guid> value '' is invalid: Element or " +
					"attribute must have valid value: must be \"true\" or \"false\"",
			},
			r: GUID{
				XMLName:     xml.Name{Space: "", Local: "guid"},
				CharData:    []byte("https://example.com/1337"),
				IsPermaLink: Ptr(IsPermaLink("")),
			},
		},
		ElementTestCase[GUID]{
			name:        "test <guid isPermaLink=\"...\"> - fail - invalid",
			wantIsValid: false,
			wantErrorIs: []error{ErrInvalidValue},
			wantErrorContains: []string{
				"Attribute 'isPermaLink' of <guid> value 'bad value' is invalid: " +
					"Element or attribute must have valid value: must be \"true\" or " +
					"\"false\"",
			},
			r: GUID{
				XMLName:     xml.Name{Space: "", Local: "guid"},
				CharData:    []byte("https://example.com/1337"),
				IsPermaLink: Ptr(IsPermaLink("bad value")),
			},
		},
		// test <comments>
		ElementTestCase[Comments]{
			name:              "test <comments> - ok",
			wantIsValid:       true,
			wantErrorIs:       []error{},
			wantErrorContains: []string{},
			r: Comments{
				XMLName:  xml.Name{Space: "", Local: "comments"},
				CharData: []byte("https://example.com/comments"),
			},
		},
		ElementTestCase[Comments]{
			name:        "test <comments> - fail - empty",
			wantIsValid: false,
			wantErrorIs: []error{ErrEmptyValue, ErrInvalidURI},
			wantErrorContains: []string{
				"Element <comments> value '' is invalid: Element must not have " +
					"empty value",
				"Element <comments> value '' is invalid: Element must contain a " +
					"valid URI (RFC3986): parse \"\": empty url",
			},
			r: Comments{
				XMLName:  xml.Name{Space: "", Local: "comments"},
				CharData: []byte(""),
			},
		},
		ElementTestCase[Comments]{
			name:        "test <comments> - fail - invalid uri",
			wantIsValid: false,
			wantErrorIs: []error{ErrInvalidURI},
			wantErrorContains: []string{
				"Element <comments> value 'bad uri' is invalid: Element must " +
					"contain a valid URI (RFC3986): parse \"bad uri\": invalid URI " +
					"for request",
			},
			r: Comments{
				XMLName:  xml.Name{Space: "", Local: "comments"},
				CharData: []byte("bad uri"),
			},
		},
		// test <author>
		ElementTestCase[Author]{
			name:              "test <author> - ok",
			wantIsValid:       true,
			wantErrorIs:       []error{},
			wantErrorContains: []string{},
			r: Author{
				XMLName:  xml.Name{Space: "", Local: "author"},
				CharData: []byte("first.last@example.com"),
			},
		},
		ElementTestCase[Author]{
			name:        "test <author> - fail - empty",
			wantIsValid: false,
			wantErrorIs: []error{ErrEmptyValue, ErrInvalidMailAddress},
			wantErrorContains: []string{
				"Element <author> value '' is invalid: Element must not have empty " +
					"value",
				"Element <author> value '' is invalid: Element must contain a valid " +
					"mail address (RFC5322): mail: no address",
			},
			r: Author{
				XMLName:  xml.Name{Space: "", Local: "author"},
				CharData: []byte(""),
			},
		},
		ElementTestCase[Author]{
			name:        "test <author> - fail - invalid mail address",
			wantIsValid: false,
			wantErrorIs: []error{ErrInvalidMailAddress},
			wantErrorContains: []string{
				"Element <author> value 'bad mail address' is invalid: Element must " +
					"contain a valid mail address (RFC5322): mail: no angle-addr",
			},
			r: Author{
				XMLName:  xml.Name{Space: "", Local: "author"},
				CharData: []byte("bad mail address"),
			},
		},
	}
	for _, tc := range cases {
		tc.Test(t)
	}
}

type ParseTestCase[T RSSElement] struct {
	name string
	data []byte
	want any
	r    T
}

func (tc ParseTestCase[T]) Test(t *testing.T) {
	t.Run(tc.name, func(t *testing.T) {
		r := tc.r
		// test unmarshaling
		err := xml.Unmarshal(tc.data, &r)
		assert.Equal(t, tc.want, r)
		assert.Nil(t, err)
		// test validation
		ret, errs := Validate(r)
		assert.True(t, ret)
		assert.Empty(t, errs)
		// test marshaling
		s, err := xml.Marshal(r)
		assert.Equal(t, tc.data, s)
		assert.Nil(t, err)
	})
}

func TestParse(t *testing.T) {
	cases := []Testable{
		ParseTestCase[GUID]{
			name: "test <guid> - parse",
			data: []byte(`<guid>https://example.com/1337</guid>`),
			want: GUID{
				XMLName:     xml.Name{Space: "", Local: "guid"},
				CharData:    []byte("https://example.com/1337"),
				IsPermaLink: nil,
			},
			r: GUID{},
		},
		ParseTestCase[GUID]{
			name: "test <guid> - parse",
			data: []byte(`<guid isPermaLink="true">https://example.com/1337</guid>`),
			want: GUID{
				XMLName:     xml.Name{Space: "", Local: "guid"},
				CharData:    []byte("https://example.com/1337"),
				IsPermaLink: Ptr(IsPermaLink("true")),
			},
			r: GUID{},
		},
		ParseTestCase[GUID]{
			name: "test <guid> - parse",
			data: []byte(`<guid isPermaLink="false">1337</guid>`),
			want: GUID{
				XMLName:     xml.Name{Space: "", Local: "guid"},
				CharData:    []byte("1337"),
				IsPermaLink: Ptr(IsPermaLink("false")),
			},
			r: GUID{},
		},
		ParseTestCase[Comments]{
			name: "test <comments> - parse",
			data: []byte(`<comments>https://example.com/comments</comments>`),
			want: Comments{
				XMLName:  xml.Name{Space: "", Local: "comments"},
				CharData: []byte("https://example.com/comments"),
			},
			r: Comments{},
		},
		ParseTestCase[Author]{
			name: "test <author> - parse",
			data: []byte(`<author>first.last@example.com</author>`),
			want: Author{
				XMLName:  xml.Name{Space: "", Local: "author"},
				CharData: []byte("first.last@example.com"),
			},
			r: Author{},
		},
	}
	for _, tc := range cases {
		tc.Test(t)
	}
}
