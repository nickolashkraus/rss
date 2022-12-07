// Copyright 2022 Nickolas Kraus. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rss

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

//	func TestPubDate(t *testing.T) {
//		t.Run("test <pubDate> - ok", func(t *testing.T) {
//			var p PubDate
//			// RFC822
//			p = "01 Jan 70 00:00 GMT"
//			ret, err := p.IsValid()
//			assert.True(t, ret)
//			assert.Nil(t, err)
//			// RFC1123
//			p = "Thu, 01 Jan 1970 00:00:00 GMT"
//			ret, err = p.IsValid()
//			assert.True(t, ret)
//			assert.Nil(t, err)
//		})
//		t.Run("test <pubDate> - fail - empty", func(t *testing.T) {
//			var p PubDate = ""
//			ret, err := p.IsValid()
//			assert.False(t, ret)
//			assert.ErrorIs(t, err, ErrEmptyValue)
//		})
//		t.Run("test <pubDate> - fail - invalid date", func(t *testing.T) {
//			var p PubDate = "bad date"
//			ret, err := p.IsValid()
//			assert.False(t, ret)
//			assert.ErrorIs(t, err, ErrInvalidDate)
//		})
//		t.Run("test <pubDate> - unmarshal", func(t *testing.T) {
//			var p PubDate
//			s := []byte(`<lastBuildDate>Thu, 01 Jan 1970 00:00:00 GMT</lastBuildDate>`)
//			err := xml.Unmarshal(s, &p)
//			assert.Equal(t, "Thu, 01 Jan 1970 00:00:00 GMT", string(p))
//			assert.Nil(t, err)
//		})
//	}
//
//	func TestLastBuildDate(t *testing.T) {
//		t.Run("test <lastBuildDate> - ok", func(t *testing.T) {
//			var l LastBuildDate
//			// RFC822
//			l = "01 Jan 70 00:00 GMT"
//			ret, err := l.IsValid()
//			assert.True(t, ret)
//			assert.Nil(t, err)
//			// RFC1123
//			l = "Thu, 01 Jan 1970 00:00:00 GMT"
//			ret, err = l.IsValid()
//			assert.True(t, ret)
//			assert.Nil(t, err)
//		})
//		t.Run("test <lastBuildDate> - fail - empty", func(t *testing.T) {
//			var l LastBuildDate = ""
//			ret, err := l.IsValid()
//			assert.False(t, ret)
//			assert.ErrorIs(t, err, ErrEmptyValue)
//		})
//		t.Run("test <lastBuildDate> - fail - invalid date", func(t *testing.T) {
//			var l LastBuildDate = "bad date"
//			ret, err := l.IsValid()
//			assert.False(t, ret)
//			assert.ErrorIs(t, err, ErrInvalidDate)
//		})
//		t.Run("test <lastBuildDate> - unmarshal", func(t *testing.T) {
//			var l LastBuildDate
//			s := []byte(`<lastBuildDate>Thu, 01 Jan 1970 00:00:00 GMT</lastBuildDate>`)
//			err := xml.Unmarshal(s, &l)
//			assert.Equal(t, "Thu, 01 Jan 1970 00:00:00 GMT", string(l))
//			assert.Nil(t, err)
//		})
//	}
//
//	func TestCategory(t *testing.T) {
//		t.Run("test <category> - ok", func(t *testing.T) {
//			var g GUID = GUID{
//				XMLName:     xml.Name{Space: "", Local: "guid"},
//				CharData:    []byte("https://example.com/guid"),
//				IsPermaLink: "true",
//			}
//			ret, err := g.IsValid()
//			assert.True(t, ret)
//			assert.Nil(t, err)
//		})
//		t.Run("test <category> - fail - empty", func(t *testing.T) {
//			var g GUID = GUID{}
//			ret, err := g.IsValid()
//			assert.False(t, ret)
//			assert.ErrorIs(t, err, ErrEmptyValue)
//		})
//		t.Run("test <category> - unmarshal", func(t *testing.T) {
//			var g GUID
//			s := []byte(`<guid isPermaLink="true">https://example.com/guid</guid>`)
//			err := xml.Unmarshal(s, &g)
//			assert.Equal(t, "true", string(g.IsPermaLink))
//			assert.Equal(t, "https://example.com/guid", string(g.CharData))
//			assert.Nil(t, err)
//		})
//	}

func TestSource(t *testing.T) {
	t.Run("test <source> - ok", func(t *testing.T) {
		var r Source
		r = Source{
			XMLName:   xml.Name{Space: "", Local: "source"},
			CharData:  []byte("Title"),
			SourceURL: "https://example.com/source.xml",
		}
		ret, errs := r.IsValid()
		assert.True(t, ret)
		assert.Empty(t, errs)
		// NOTE: <source> can be empty.
		r = Source{
			XMLName:   xml.Name{Space: "", Local: "source"},
			CharData:  []byte(""),
			SourceURL: "https://example.com/source.xml",
		}
		ret, errs = r.IsValid()
		assert.True(t, ret)
		assert.Empty(t, errs)
	})
	t.Run("test <source> - fail - all", func(t *testing.T) {
		var r Source = Source{
			XMLName:   xml.Name{Space: "", Local: "source"},
			CharData:  []byte(""),
			SourceURL: "",
		}
		ret, errs := r.IsValid()
		assert.False(t, ret)
		assert.Equal(t, 2, len(errs))
		assert.ErrorIs(t, errs[0], ErrEmptyValue)
		assert.ErrorIs(t, errs[1], ErrInvalidURI)
	})
	t.Run("test <source> - unmarshal", func(t *testing.T) {
		var r Source
		s := []byte(`<source url="https://example.com/source.xml">Title</source>`)
		err := xml.Unmarshal(s, &r)
		assert.Equal(t, "Title", string(r.CharData))
		assert.Equal(t, "https://example.com/source.xml", string(r.SourceURL))
		assert.Nil(t, err)
	})
}

func TestSourceURL(t *testing.T) {
	t.Run("test <source url=\"...\"> - ok", func(t *testing.T) {
		var r SourceURL = "https://example.com/source.xml"
		ret, errs := r.IsValid()
		assert.True(t, ret)
		assert.Empty(t, errs)
	})
	t.Run("test <source url=\"...\"> - fail - empty", func(t *testing.T) {
		var r SourceURL = ""
		ret, errs := r.IsValid()
		assert.False(t, ret)
		assert.Equal(t, 2, len(errs))
		assert.ErrorIs(t, errs[0], ErrEmptyValue)
	})
	t.Run("test <source url=\"...\"> - fail - invalid uri", func(t *testing.T) {
		var r SourceURL = "bad uri"
		ret, errs := r.IsValid()
		assert.False(t, ret)
		assert.Equal(t, 1, len(errs))
		assert.ErrorIs(t, errs[0], ErrInvalidURI)
	})
	t.Run("test <source url=\"...\"> - fail - all", func(t *testing.T) {
		var r SourceURL = ""
		ret, errs := r.IsValid()
		assert.False(t, ret)
		assert.Equal(t, 2, len(errs))
		assert.ErrorIs(t, errs[0], ErrEmptyValue)
		assert.ErrorIs(t, errs[1], ErrInvalidURI)
	})
}

func TestEnclosure(t *testing.T) {
	t.Run("test <enclosure> - ok", func(t *testing.T) {
		var r Enclosure = Enclosure{
			XMLName:      xml.Name{Space: "", Local: "enclosure"},
			CharData:     []byte(""),
			EnclosureURL: "https://example.com/audio.mp3",
			Length:       "1337",
			Type:         "audio/mpeg",
		}
		ret, errs := r.IsValid()
		assert.True(t, ret)
		assert.Empty(t, errs)
	})
	t.Run("test <enclosure> - fail - not empty", func(t *testing.T) {
		var r Enclosure = Enclosure{
			XMLName:      xml.Name{Space: "", Local: "enclosure"},
			CharData:     []byte("not empty"),
			EnclosureURL: "https://example.com/audio.mp3",
			Length:       "1337",
			Type:         "audio/mpeg",
		}
		ret, errs := r.IsValid()
		assert.False(t, ret)
		assert.Equal(t, 1, len(errs))
		assert.ErrorIs(t, errs[0], ErrNonEmptyValue)
	})
	t.Run("test <enclosure> - fail - all", func(t *testing.T) {
		var r Enclosure = Enclosure{
			XMLName:      xml.Name{Space: "", Local: "enclosure"},
			CharData:     []byte("not empty"),
			EnclosureURL: "",
			Length:       "-1",
			Type:         "",
		}
		ret, errs := r.IsValid()
		assert.False(t, ret)
		assert.Equal(t, 5, len(errs))
		assert.ErrorIs(t, errs[0], ErrNonEmptyValue)
		assert.ErrorIs(t, errs[1], ErrEmptyValue)
		assert.ErrorIs(t, errs[2], ErrInvalidURI)
		assert.ErrorIs(t, errs[3], ErrInvalidValue)
		assert.ErrorIs(t, errs[4], ErrEmptyValue)
	})
	t.Run("test <enclosure> - unmarshal", func(t *testing.T) {
		var r Enclosure
		s := []byte(`<enclosure url="https://example.com/audio.mp3" length="1337" type="audio/mpeg" />`)
		err := xml.Unmarshal(s, &r)
		assert.Equal(t, "", string(r.CharData))
		assert.Equal(t, "https://example.com/audio.mp3", string(r.EnclosureURL))
		assert.Equal(t, "1337", string(r.Length))
		assert.Equal(t, "audio/mpeg", string(r.Type))
		assert.Nil(t, err)
	})
}

func TestEnclosureURL(t *testing.T) {
	t.Run("test <enclosure url=\"...\"> - ok", func(t *testing.T) {
		var r EnclosureURL = "https://example.com/audio.mp3"
		ret, errs := r.IsValid()
		assert.True(t, ret)
		assert.Empty(t, errs)
	})
	t.Run("test <enclosure url=\"...\"> - fail - empty", func(t *testing.T) {
		var r EnclosureURL = ""
		ret, errs := r.IsValid()
		assert.False(t, ret)
		assert.Equal(t, 2, len(errs))
		assert.ErrorIs(t, errs[0], ErrEmptyValue)
	})
	t.Run("test <enclosure url=\"...\"> - fail - invalid uri", func(t *testing.T) {
		var r EnclosureURL = "bad uri"
		ret, errs := r.IsValid()
		assert.False(t, ret)
		assert.Equal(t, 1, len(errs))
		assert.ErrorIs(t, errs[0], ErrInvalidURI)
	})
	t.Run("test <enclosure url=\"...\"> - fail - all", func(t *testing.T) {
		var r EnclosureURL = ""
		ret, errs := r.IsValid()
		assert.False(t, ret)
		assert.Equal(t, 2, len(errs))
		assert.ErrorIs(t, errs[0], ErrEmptyValue)
		assert.ErrorIs(t, errs[1], ErrInvalidURI)
	})
}

func TestLength(t *testing.T) {
	t.Run("test <enclosure length=\"...\"> - ok", func(t *testing.T) {
		var r Length = "1337"
		ret, errs := r.IsValid()
		assert.True(t, ret)
		assert.Empty(t, errs)
	})
	t.Run("test <enclosure length=\"...\"> - fail - invalid value", func(t *testing.T) {
		var r Length = "-1"
		ret, errs := r.IsValid()
		assert.False(t, ret)
		assert.Equal(t, 1, len(errs))
		assert.ErrorIs(t, errs[0], ErrInvalidValue)
	})
}

func TestType(t *testing.T) {
	t.Run("test <enclosure type=\"...\"> - ok", func(t *testing.T) {
		var r Type = "audio/mpeg"
		ret, errs := r.IsValid()
		assert.True(t, ret)
		assert.Empty(t, errs)
	})
	t.Run("test <enclosure type=\"...\"> - fail - empty", func(t *testing.T) {
		var r Type = ""
		ret, errs := r.IsValid()
		assert.False(t, ret)
		assert.Equal(t, 1, len(errs))
		assert.ErrorIs(t, errs[0], ErrEmptyValue)
	})
}

func TestGUID(t *testing.T) {
	t.Run("test <guid> - ok", func(t *testing.T) {
		var r GUID = GUID{
			XMLName:     xml.Name{Space: "", Local: "guid"},
			CharData:    []byte("https://example.com/guid"),
			IsPermaLink: "true",
		}
		ret, errs := r.IsValid()
		assert.True(t, ret)
		assert.Empty(t, errs)
	})
	t.Run("test <guid> - fail - empty", func(t *testing.T) {
		var r GUID = GUID{}
		ret, errs := r.IsValid()
		assert.False(t, ret)
		assert.Equal(t, 2, len(errs))
		assert.ErrorIs(t, errs[0], ErrEmptyValue)
	})
	t.Run("test <guid> - fail - invalid uri", func(t *testing.T) {
		var r GUID = GUID{
			XMLName:     xml.Name{Space: "", Local: "guid"},
			CharData:    []byte("bad uri"),
			IsPermaLink: "true",
		}
		ret, errs := r.IsValid()
		assert.False(t, ret)
		assert.Equal(t, 1, len(errs))
		assert.ErrorIs(t, errs[0], ErrInvalidURI)
	})
	t.Run("test <guid> - fail - all", func(t *testing.T) {
		var r GUID = GUID{
			XMLName:     xml.Name{Space: "", Local: "guid"},
			CharData:    []byte(""),
			IsPermaLink: "true",
		}
		ret, errs := r.IsValid()
		assert.False(t, ret)
		assert.Equal(t, 2, len(errs))
		assert.ErrorIs(t, errs[0], ErrEmptyValue)
		assert.ErrorIs(t, errs[1], ErrInvalidURI)
	})
	t.Run("test <guid> - fail - invalid isPermaLink", func(t *testing.T) {
		var r GUID = GUID{
			XMLName:     xml.Name{Space: "", Local: "guid"},
			CharData:    []byte("https://example.com/guid"),
			IsPermaLink: "bad value",
		}
		ret, errs := r.IsValid()
		assert.False(t, ret)
		assert.Equal(t, 1, len(errs))
		assert.ErrorIs(t, errs[0], ErrInvalidValue)
	})
	t.Run("test <guid> - unmarshal", func(t *testing.T) {
		var r GUID
		s := []byte(`<guid isPermaLink="true">https://example.com/guid</guid>`)
		err := xml.Unmarshal(s, &r)
		assert.Equal(t, "https://example.com/guid", string(r.CharData))
		assert.Equal(t, "true", string(r.IsPermaLink))
		assert.Nil(t, err)
	})
}

func TestIsPermaLink(t *testing.T) {
	t.Run("test <guid isPermaLink=\"...\"> - ok", func(t *testing.T) {
		var r IsPermaLink
		r = "true"
		ret, errs := r.IsValid()
		assert.True(t, ret)
		assert.Empty(t, errs)
		r = "false"
		ret, errs = r.IsValid()
		assert.True(t, ret)
		assert.Empty(t, errs)
	})
	t.Run("test <guid isPermaLink=\"...\"> - fail - invalid value", func(t *testing.T) {
		var r IsPermaLink = "bad value"
		ret, errs := r.IsValid()
		assert.False(t, ret)
		assert.Equal(t, 1, len(errs))
		assert.ErrorIs(t, errs[0], ErrInvalidValue)
	})
}

func TestComments(t *testing.T) {
	t.Run("test <comments> - ok", func(t *testing.T) {
		var r Comments = "https://example.com/comments"
		ret, errs := r.IsValid()
		assert.True(t, ret)
		assert.Empty(t, errs)
	})
	t.Run("test <comments> - fail - empty", func(t *testing.T) {
		var r Comments = ""
		ret, errs := r.IsValid()
		assert.False(t, ret)
		assert.Equal(t, 2, len(errs))
		assert.ErrorIs(t, errs[0], ErrEmptyValue)
	})
	t.Run("test <comments> - fail - invalid uri", func(t *testing.T) {
		var r Comments = "bad uri"
		ret, errs := r.IsValid()
		assert.False(t, ret)
		assert.Equal(t, 1, len(errs))
		assert.ErrorIs(t, errs[0], ErrInvalidURI)
	})
	t.Run("test <comments> - fail - all", func(t *testing.T) {
		var r Comments = ""
		ret, errs := r.IsValid()
		assert.False(t, ret)
		assert.Equal(t, 2, len(errs))
		assert.ErrorIs(t, errs[0], ErrEmptyValue)
		assert.ErrorIs(t, errs[1], ErrInvalidURI)
	})
	t.Run("test <comments> - unmarshal", func(t *testing.T) {
		var r Comments
		s := []byte(`<comments>https://example.com/comments</comments>`)
		err := xml.Unmarshal(s, &r)
		assert.Equal(t, "https://example.com/comments", string(r))
		assert.Nil(t, err)
	})
}

func TestAuthor(t *testing.T) {
	t.Run("test <author> - ok", func(t *testing.T) {
		var r Author = "first.last@example.com"
		ret, errs := r.IsValid()
		assert.True(t, ret)
		assert.Empty(t, errs)
	})
	t.Run("test <author> - fail - empty", func(t *testing.T) {
		var r Author = ""
		ret, errs := r.IsValid()
		assert.False(t, ret)
		assert.Equal(t, 2, len(errs))
		assert.ErrorIs(t, errs[0], ErrEmptyValue)
	})
	t.Run("test <author> - fail - invalid mail address", func(t *testing.T) {
		var r Author = "bad mail address"
		ret, errs := r.IsValid()
		assert.False(t, ret)
		assert.Equal(t, 1, len(errs))
		assert.ErrorIs(t, errs[0], ErrInvalidMailAddress)
	})
	t.Run("test <author> - fail - all", func(t *testing.T) {
		var r Author = ""
		ret, errs := r.IsValid()
		assert.False(t, ret)
		assert.Equal(t, 2, len(errs))
		assert.ErrorIs(t, errs[0], ErrEmptyValue)
		assert.ErrorIs(t, errs[1], ErrInvalidMailAddress)
	})
	t.Run("test <author> - unmarshal", func(t *testing.T) {
		var r Author
		s := []byte(`<author>first.last@example.com</author>`)
		err := xml.Unmarshal(s, &r)
		assert.Equal(t, "first.last@example.com", string(r))
		assert.Nil(t, err)
	})
}
