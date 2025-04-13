package optional

import "encoding/xml"

// MarshalXML encodes the value if present, otherwise it does nothing (Currently
// always omitemtpy behaviour).
func (o Optional[T]) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if !o.present {
		return nil
	}

	return e.EncodeElement(o.value, start)
}

// UnmarshalXML creates a present Optional if the value is not empty, otherwise
// an empty.
func (o *Optional[T]) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	err := d.DecodeElement(&o.value, &start)
	if err != nil {
		return err
	}

	o.present = true

	return nil
}
