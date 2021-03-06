package vast

// AdParameters type represents an <AdParameters> element that defines arbitrary ad parameters.
type AdParameters struct {
	Parameters TrimmedData `xml:",cdata"`

	IsXmlEncoded bool `xml:"xmlEncoded,attr,omitempty"` // VAST3.0.
}
