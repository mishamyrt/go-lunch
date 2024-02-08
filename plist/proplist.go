// Package plist provides tools for rendering macOS plist files.
package plist

import "strings"

const (
	xmlHeader    = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>"
	specURL      = "http://www.apple.com/DTDs/PropertyList-1.0.dtd"
	typePropList = "<!DOCTYPE plist PUBLIC \"-//Apple//DTD PLIST 1.0//EN\" \"" + specURL + "\">"
)

// PropList represents Apple property list file.
type PropList struct {
	props []string
}

// AddBool tag to PropList.
func (d *PropList) AddBool(key string, value bool) *PropList {
	d.appendRaw("<key>" + key + "</key>")
	if value {
		d.appendRaw("<true/>")
	} else {
		d.appendRaw("<false/>")
	}
	return d
}

// AddString tag to PropList.
func (d *PropList) AddString(key string, value string) *PropList {
	d.append("key", key)
	d.append("string", value)
	return d
}

// AddStringArray adds multiple literal tags to PropList.
func (d *PropList) AddStringArray(key string, values []string) *PropList {
	d.append("key", key)
	d.appendRaw("<array>")
	for _, value := range values {
		d.append("string", value)
	}
	d.appendRaw("</array>")
	return d
}

// String joins props to string.
func (d *PropList) String() string {
	result := xmlHeader + "\n" + typePropList + "\n"
	result += "<plist version=\"1.0\">" + "\n"
	result += "<dict>" + "\n"
	result += strings.Join(d.props, "\n") + "\n"
	result += "</dict>" + "\n"
	result += "</plist>" + "\n"
	return result
}

// Bytes returns file bytes content.
func (d *PropList) Bytes() []byte {
	return []byte(d.String())
}

func (d *PropList) appendRaw(value string) {
	d.props = append(d.props, value)
}

func (d *PropList) append(tag, value string) {
	d.appendRaw("<" + tag + ">" + value + "</" + tag + ">")
}

// New creates new property list.
func New() *PropList {
	builder := PropList{
		props: make([]string, 0),
	}
	return &builder
}
