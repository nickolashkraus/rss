// Copyright 2022 Nickolas Kraus. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rss

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPubDate(t *testing.T) {
	t.Run("test <pubDate> - ok", func(t *testing.T) {
		var r PubDate
		r = PubDate{
			XMLName:  xml.Name{Space: "", Local: "pubDate"},
			CharData: []byte("01 Jan 70 00:00 GMT"), // RFC822
		}
		ret, errs := r.IsValid()
		assert.True(t, ret)
		assert.Empty(t, errs)
		r = PubDate{
			XMLName:  xml.Name{Space: "", Local: "pubDate"},
			CharData: []byte("Thu, 01 Jan 1970 00:00:00 GMT"), // RFC1123
		}
		ret, errs = r.IsValid()
		assert.True(t, ret)
		assert.Empty(t, errs)
	})
	t.Run("test <pubDate> - fail - empty", func(t *testing.T) {
		var r = PubDate{
			XMLName:  xml.Name{Space: "", Local: "pubDate"},
			CharData: []byte(""),
		}
		ret, errs := r.IsValid()
		assert.Equal(t, 2, len(errs))
		assert.False(t, ret)
		assert.ErrorIs(t, errs[0], ErrEmptyValue)
	})
	t.Run("test <pubDate> - fail - invalid date", func(t *testing.T) {
		var r = PubDate{
			XMLName:  xml.Name{Space: "", Local: "pubDate"},
			CharData: []byte("bad date"),
		}
		ret, errs := r.IsValid()
		assert.Equal(t, 1, len(errs))
		assert.False(t, ret)
		assert.ErrorIs(t, errs[0], ErrInvalidDate)
	})
	t.Run("test <pubDate> - fail - multiple", func(t *testing.T) {
		var r = PubDate{
			XMLName:  xml.Name{Space: "", Local: "pubDate"},
			CharData: []byte(""),
		}
		ret, errs := r.IsValid()
		assert.False(t, ret)
		assert.Equal(t, 2, len(errs))
		assert.ErrorIs(t, errs[0], ErrEmptyValue)
		assert.ErrorIs(t, errs[1], ErrInvalidDate)
	})
	t.Run("test <pubDate> - unmarshal", func(t *testing.T) {
		var r PubDate
		s := []byte(`<pubDate>Thu, 01 Jan 1970 00:00:00 GMT</pubDate>`)
		err := xml.Unmarshal(s, &r)
		assert.Equal(t, "Thu, 01 Jan 1970 00:00:00 GMT", string(r.CharData))
		assert.Nil(t, err)
	})
	t.Run("test <pubDate> - marshal", func(t *testing.T) {
		var r PubDate = PubDate{
			XMLName:  xml.Name{Space: "", Local: "pubDate"},
			CharData: []byte("Thu, 01 Jan 1970 00:00:00 GMT"),
		}
		exp := []byte(`<pubDate>Thu, 01 Jan 1970 00:00:00 GMT</pubDate>`)
		s, err := xml.Marshal(r)
		assert.Equal(t, exp, s)
		assert.Nil(t, err)
		ret, errs := Validate(r)
		assert.True(t, ret)
		assert.Empty(t, errs)
	})
}

func TestLastBuildDate(t *testing.T) {
	t.Run("test <lastBuildDate> - ok", func(t *testing.T) {
		var r LastBuildDate
		r = LastBuildDate{
			XMLName:  xml.Name{Space: "", Local: "lastBuildDate"},
			CharData: []byte("01 Jan 70 00:00 GMT"), // RFC822
		}
		ret, errs := r.IsValid()
		assert.True(t, ret)
		assert.Empty(t, errs)
		r = LastBuildDate{
			XMLName:  xml.Name{Space: "", Local: "lastBuildDate"},
			CharData: []byte("Thu, 01 Jan 1970 00:00:00 GMT"), // RFC1123
		}
		ret, errs = r.IsValid()
		assert.True(t, ret)
		assert.Empty(t, errs)
	})
	t.Run("test <lastBuildDate> - fail - empty", func(t *testing.T) {
		var r = LastBuildDate{
			XMLName:  xml.Name{Space: "", Local: "lastBuildDate"},
			CharData: []byte(""),
		}
		ret, errs := r.IsValid()
		assert.Equal(t, 2, len(errs))
		assert.False(t, ret)
		assert.ErrorIs(t, errs[0], ErrEmptyValue)
	})
	t.Run("test <lastBuildDate> - fail - invalid date", func(t *testing.T) {
		var r = LastBuildDate{
			XMLName:  xml.Name{Space: "", Local: "lastBuildDate"},
			CharData: []byte("bad date"),
		}
		ret, errs := r.IsValid()
		assert.Equal(t, 1, len(errs))
		assert.False(t, ret)
		assert.ErrorIs(t, errs[0], ErrInvalidDate)
	})
	t.Run("test <lastBuildDate> - fail - multiple", func(t *testing.T) {
		var r = LastBuildDate{
			XMLName:  xml.Name{Space: "", Local: "lastBuildDate"},
			CharData: []byte(""),
		}
		ret, errs := r.IsValid()
		assert.False(t, ret)
		assert.Equal(t, 2, len(errs))
		assert.ErrorIs(t, errs[0], ErrEmptyValue)
		assert.ErrorIs(t, errs[1], ErrInvalidDate)
	})
	t.Run("test <lastBuildDate> - unmarshal", func(t *testing.T) {
		var r LastBuildDate
		s := []byte(`<lastBuildDate>Thu, 01 Jan 1970 00:00:00 GMT</lastBuildDate>`)
		err := xml.Unmarshal(s, &r)
		assert.Equal(t, "Thu, 01 Jan 1970 00:00:00 GMT", string(r.CharData))
		assert.Nil(t, err)
	})
	t.Run("test <lastBuildDate> - marshal", func(t *testing.T) {
		var r LastBuildDate = LastBuildDate{
			XMLName:  xml.Name{Space: "", Local: "lastBuildDate"},
			CharData: []byte("Thu, 01 Jan 1970 00:00:00 GMT"),
		}
		exp := []byte(`<lastBuildDate>Thu, 01 Jan 1970 00:00:00 GMT</lastBuildDate>`)
		s, err := xml.Marshal(r)
		assert.Equal(t, exp, s)
		assert.Nil(t, err)
		ret, errs := Validate(r)
		assert.True(t, ret)
		assert.Empty(t, errs)
	})
}

func TestCategoryDomain(t *testing.T) {
	t.Run("test <category domain=\"...\"> - ok", func(t *testing.T) {
		var r CategoryDomain = "https://example.com/category"
		ret, errs := r.IsValid()
		assert.True(t, ret)
		assert.Empty(t, errs)
	})
	t.Run("test <category domain=\"...\"> - fail - empty", func(t *testing.T) {
		var r CategoryDomain = ""
		ret, errs := r.IsValid()
		assert.False(t, ret)
		assert.Equal(t, 1, len(errs))
		assert.ErrorIs(t, errs[0], ErrEmptyValue)
	})
}

func TestItem(t *testing.T) {
	t.Run("test <item> - ok", func(t *testing.T) {
		var r Item
		r = Item{
			Title:       "Title",
			Link:        "https://example.com",
			Description: "Description",
			Source: &Source{
				XMLName:   xml.Name{Space: "", Local: "source"},
				CharData:  []byte("Title"),
				SourceURL: "https://example.com/source.xml",
			},
			Enclosure: &Enclosure{
				XMLName:      xml.Name{Space: "", Local: "enclosure"},
				CharData:     []byte(""),
				EnclosureURL: "https://example.com/audio.mp3",
				Length:       "1337",
				Type:         "audio/mpeg",
			},
			Category: &Category{
				XMLName:        xml.Name{Space: "", Local: "category"},
				CharData:       []byte("Category"),
				CategoryDomain: "https://example.com/category",
			},
			PubDate: &PubDate{
				XMLName:  xml.Name{Space: "", Local: "pubDate"},
				CharData: []byte("Thu, 01 Jan 1970 00:00:00 GMT"),
			},
			GUID: &GUID{
				XMLName:     xml.Name{Space: "", Local: "guid"},
				CharData:    []byte("https://example.com/guid"),
				IsPermaLink: "true",
			},
			Comments: &Comments{
				XMLName:  xml.Name{Space: "", Local: "comments"},
				CharData: []byte("https://example.com/comments"),
			},
			Author: &Author{
				XMLName:  xml.Name{Space: "", Local: "author"},
				CharData: []byte("first.last@example.com"),
			},
		}
		ret, errs := r.IsValid()
		assert.True(t, ret)
		assert.Empty(t, errs)
	})
	t.Run("test <item> - fail - invalid element", func(t *testing.T) {
		var r Item
		r = Item{
			Title:       "",
			Link:        "https://example.com",
			Description: "",
			Source: &Source{
				XMLName:   xml.Name{Space: "", Local: "source"},
				CharData:  []byte("Title"),
				SourceURL: "https://example.com/source.xml",
			},
			Enclosure: &Enclosure{
				XMLName:      xml.Name{Space: "", Local: "enclosure"},
				CharData:     []byte(""),
				EnclosureURL: "https://example.com/audio.mp3",
				Length:       "1337",
				Type:         "audio/mpeg",
			},
			Category: &Category{
				XMLName:        xml.Name{Space: "", Local: "category"},
				CharData:       []byte("Category"),
				CategoryDomain: "https://example.com/category",
			},
			PubDate: &PubDate{
				XMLName:  xml.Name{Space: "", Local: "pubDate"},
				CharData: []byte("Thu, 01 Jan 1970 00:00:00 GMT"),
			},
			GUID: &GUID{
				XMLName:     xml.Name{Space: "", Local: "guid"},
				CharData:    []byte("https://example.com/guid"),
				IsPermaLink: "true",
			},
			Comments: &Comments{
				XMLName:  xml.Name{Space: "", Local: "comments"},
				CharData: []byte("https://example.com/comments"),
			},
			Author: &Author{
				XMLName:  xml.Name{Space: "", Local: "author"},
				CharData: []byte("first.last@example.com"),
			},
		}
		ret, errs := r.IsValid()
		assert.False(t, ret)
		assert.Equal(t, 1, len(errs))
		assert.ErrorIs(t, errs[0], ErrInvalidElement)
	})
	t.Run("test <item> - unmarshal", func(t *testing.T) {
		var r Item
		s := []byte(`<item>
	<title>Title</title>
	<link>https://example.com</link>
	<description>Description</description>
	<source url="https://example.com/source.xml">Title</source>
  <enclosure url="https://example.com/audio.mp3" length="1337" type="audio/mpeg"/>
	<category domain="https://example.com/category">Category</category>
	<pubDate>Thu, 01 Jan 1970 00:00:00 GMT</pubDate>
	<guid isPermaLink="true">https://example.com/guid</guid>
	<comments>https://example.com/comments</comments>
	<author>first.last@example.com</author>
</item>
`)
		err := xml.Unmarshal(s, &r)
		assert.Equal(t, "Title", string(r.Title))
		assert.Equal(t, "https://example.com", string(r.Link))
		assert.Equal(t, "Description", string(r.Description))
		assert.Equal(t, "https://example.com/source.xml", string(r.Source.SourceURL))
		assert.Equal(t, "", string(r.Enclosure.CharData))
		assert.Equal(t, "https://example.com/audio.mp3", string(r.Enclosure.EnclosureURL))
		assert.Equal(t, "1337", string(r.Enclosure.Length))
		assert.Equal(t, "audio/mpeg", string(r.Enclosure.Type))
		assert.Equal(t, "Category", string(r.Category.CharData))
		assert.Equal(t, "https://example.com/category", string(r.Category.CategoryDomain))
		assert.Equal(t, "Thu, 01 Jan 1970 00:00:00 GMT", string(r.PubDate.CharData))
		assert.Equal(t, "https://example.com/guid", string(r.GUID.CharData))
		assert.Equal(t, "true", string(r.GUID.IsPermaLink))
		assert.Equal(t, "https://example.com/comments", string(r.Comments.CharData))
		assert.Equal(t, "first.last@example.com", string(r.Author.CharData))
		assert.Nil(t, err)
		ret, errs := Validate(r)
		assert.True(t, ret)
		assert.Empty(t, errs)
	})
	t.Run("test <item> - unmarshal", func(t *testing.T) {
		var r Item
		s := []byte(`<item>
	<title>Title</title>
	<description>Description</description>
</item>
`)
		err := xml.Unmarshal(s, &r)
		assert.Equal(t, "Title", string(r.Title))
		assert.Equal(t, "Description", string(r.Description))
		assert.Nil(t, err)
		ret, errs := Validate(r)
		assert.True(t, ret)
		assert.Empty(t, errs)
	})
	t.Run("test <item> - marshal", func(t *testing.T) {
		var r = Item{
			Title:       "Title",
			Link:        "https://example.com",
			Description: "Description",
			Source: &Source{
				XMLName:   xml.Name{Space: "", Local: "source"},
				CharData:  []byte("Title"),
				SourceURL: "https://example.com/source.xml",
			},
			Enclosure: &Enclosure{
				XMLName:      xml.Name{Space: "", Local: "enclosure"},
				CharData:     []byte(""),
				EnclosureURL: "https://example.com/audio.mp3",
				Length:       "1337",
				Type:         "audio/mpeg",
			},
			Category: &Category{
				XMLName:        xml.Name{Space: "", Local: "category"},
				CharData:       []byte("Category"),
				CategoryDomain: "https://example.com/category",
			},
			PubDate: &PubDate{
				XMLName:  xml.Name{Space: "", Local: "pubDate"},
				CharData: []byte("Thu, 01 Jan 1970 00:00:00 GMT"),
			},
			GUID: &GUID{
				XMLName:     xml.Name{Space: "", Local: "guid"},
				CharData:    []byte("https://example.com/guid"),
				IsPermaLink: "true",
			},
			Comments: &Comments{
				XMLName:  xml.Name{Space: "", Local: "comments"},
				CharData: []byte("https://example.com/comments"),
			},
			Author: &Author{
				XMLName:  xml.Name{Space: "", Local: "author"},
				CharData: []byte("first.last@example.com"),
			},
		}
		exp := []byte(`<item><title>Title</title><link>https://example.com</link><description>Description</description><source url="https://example.com/source.xml">Title</source><enclosure url="https://example.com/audio.mp3" length="1337" type="audio/mpeg"></enclosure><category domain="https://example.com/category">Category</category><pubDate>Thu, 01 Jan 1970 00:00:00 GMT</pubDate><guid isPermaLink="true">https://example.com/guid</guid><comments>https://example.com/comments</comments><author>first.last@example.com</author></item>`)
		s, err := xml.Marshal(r)
		assert.Equal(t, exp, s)
		assert.Nil(t, err)
		ret, errs := Validate(r)
		assert.True(t, ret)
		assert.Empty(t, errs)
	})
}

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
	t.Run("test <source> - fail - multiple", func(t *testing.T) {
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
	t.Run("test <source> - marshal", func(t *testing.T) {
		var r Source = Source{
			XMLName:   xml.Name{Space: "", Local: "source"},
			CharData:  []byte("Title"),
			SourceURL: "https://example.com/source.xml",
		}
		exp := []byte(`<source url="https://example.com/source.xml">Title</source>`)
		s, err := xml.Marshal(r)
		assert.Equal(t, exp, s)
		assert.Nil(t, err)
		ret, errs := Validate(r)
		assert.True(t, ret)
		assert.Empty(t, errs)
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
	t.Run("test <source url=\"...\"> - fail - multiple", func(t *testing.T) {
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
	t.Run("test <enclosure> - fail - multiple", func(t *testing.T) {
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
	t.Run("test <enclosure> - marshal", func(t *testing.T) {
		var r Enclosure = Enclosure{
			XMLName:      xml.Name{Space: "", Local: "enclosure"},
			CharData:     []byte(""),
			EnclosureURL: "https://example.com/audio.mp3",
			Length:       "1337",
			Type:         "audio/mpeg",
		}
		// NOTE: In XML and XHTML, a self-closing tag is a shorthand notation for
		// an opening and closing tag in one. It's used to communicate lack of
		// content in between the opening and closing tags. So, rather than typing
		// <enclosure></enclosure> (with no space at all in between), you'd be able
		// write <enclosure/>.
		exp := []byte(`<enclosure url="https://example.com/audio.mp3" length="1337" type="audio/mpeg"></enclosure>`)
		s, err := xml.Marshal(r)
		assert.Equal(t, exp, s)
		assert.Nil(t, err)
		ret, errs := Validate(r)
		assert.True(t, ret)
		assert.Empty(t, errs)
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
	t.Run("test <enclosure url=\"...\"> - fail - multiple", func(t *testing.T) {
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
	t.Run("test <guid> - fail - multiple", func(t *testing.T) {
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
	t.Run("test <guid> - marshal", func(t *testing.T) {
		var r GUID = GUID{
			XMLName:     xml.Name{Space: "", Local: "guid"},
			CharData:    []byte("https://example.com/guid"),
			IsPermaLink: "true",
		}
		exp := []byte(`<guid isPermaLink="true">https://example.com/guid</guid>`)
		s, err := xml.Marshal(r)
		assert.Equal(t, exp, s)
		assert.Nil(t, err)
		ret, errs := Validate(r)
		assert.True(t, ret)
		assert.Empty(t, errs)
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
		var r Comments = Comments{
			XMLName:  xml.Name{Space: "", Local: "comments"},
			CharData: []byte("https://example.com/comments"),
		}
		ret, errs := r.IsValid()
		assert.True(t, ret)
		assert.Empty(t, errs)
	})
	t.Run("test <comments> - fail - empty", func(t *testing.T) {
		var r Comments = Comments{
			XMLName:  xml.Name{Space: "", Local: "comments"},
			CharData: []byte(""),
		}
		ret, errs := r.IsValid()
		assert.False(t, ret)
		assert.Equal(t, 2, len(errs))
		assert.ErrorIs(t, errs[0], ErrEmptyValue)
	})
	t.Run("test <comments> - fail - invalid uri", func(t *testing.T) {
		var r Comments = Comments{
			XMLName:  xml.Name{Space: "", Local: "comments"},
			CharData: []byte("bad uri"),
		}
		ret, errs := r.IsValid()
		assert.False(t, ret)
		assert.Equal(t, 1, len(errs))
		assert.ErrorIs(t, errs[0], ErrInvalidURI)
	})
	t.Run("test <comments> - fail - multiple", func(t *testing.T) {
		var r Comments = Comments{
			XMLName:  xml.Name{Space: "", Local: "comments"},
			CharData: []byte(""),
		}
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
		assert.Equal(t, "https://example.com/comments", string(r.CharData))
		assert.Nil(t, err)
	})
	t.Run("test <comments> - marshal", func(t *testing.T) {
		var r Comments = Comments{
			XMLName:  xml.Name{Space: "", Local: "comments"},
			CharData: []byte("https://example.com/comments"),
		}
		exp := []byte(`<comments>https://example.com/comments</comments>`)
		s, err := xml.Marshal(r)
		assert.Equal(t, exp, s)
		assert.Nil(t, err)
		ret, errs := Validate(r)
		assert.True(t, ret)
		assert.Empty(t, errs)
	})
}

func TestAuthor(t *testing.T) {
	t.Run("test <author> - ok", func(t *testing.T) {
		var r Author = Author{
			XMLName:  xml.Name{Space: "", Local: "author"},
			CharData: []byte("first.last@example.com"),
		}
		ret, errs := r.IsValid()
		assert.True(t, ret)
		assert.Empty(t, errs)
	})
	t.Run("test <author> - fail - empty", func(t *testing.T) {
		var r Author = Author{
			XMLName:  xml.Name{Space: "", Local: "author"},
			CharData: []byte(""),
		}
		ret, errs := r.IsValid()
		assert.False(t, ret)
		assert.Equal(t, 2, len(errs))
		assert.ErrorIs(t, errs[0], ErrEmptyValue)
	})
	t.Run("test <author> - fail - invalid mail address", func(t *testing.T) {
		var r Author = Author{
			XMLName:  xml.Name{Space: "", Local: "author"},
			CharData: []byte("bad mail address"),
		}
		ret, errs := r.IsValid()
		assert.False(t, ret)
		assert.Equal(t, 1, len(errs))
		assert.ErrorIs(t, errs[0], ErrInvalidMailAddress)
	})
	t.Run("test <author> - fail - multiple", func(t *testing.T) {
		var r Author = Author{
			XMLName:  xml.Name{Space: "", Local: "author"},
			CharData: []byte(""),
		}
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
		assert.Equal(t, "first.last@example.com", string(r.CharData))
		assert.Nil(t, err)
	})
	t.Run("test <author> - marshal", func(t *testing.T) {
		var r Author = Author{
			XMLName:  xml.Name{Space: "", Local: "author"},
			CharData: []byte("first.last@example.com"),
		}
		exp := []byte(`<author>first.last@example.com</author>`)
		s, err := xml.Marshal(r)
		assert.Equal(t, exp, s)
		assert.Nil(t, err)
		ret, errs := Validate(r)
		assert.True(t, ret)
		assert.Empty(t, errs)
	})
}
