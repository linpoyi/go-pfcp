// Copyright 2019-2022 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"io"
	"net"
)

// NewMACAddressesDetected creates a new MACAddressesDetected IE.
func NewMACAddressesDetected(ctag, stag *IE, macs ...net.HardwareAddr) *IE {
	fields := NewMACAddressesDetectedFields(ctag, stag, macs...)

	b, err := fields.Marshal()
	if err != nil {
		return nil
	}

	return New(MACAddressesDetected, b)
}

// MACAddressesDetected returns MACAddressesDetected in structured format if the type of IE matches.
func (i *IE) MACAddressesDetected() (*MACAddressesDetectedFields, error) {
	switch i.Type {
	case MACAddressesDetected:
		fields, err := ParseMACAddressesDetectedFields(i.Payload)
		if err != nil {
			return nil, err
		}

		return fields, nil
	case EthernetTrafficInformation:
		ies, err := i.EthernetTrafficInformation()
		if err != nil {
			return nil, err
		}
		for _, x := range ies {
			if x.Type == MACAddressesDetected {
				return x.MACAddressesDetected()
			}
		}
		return nil, ErrIENotFound
	case EthernetContextInformation:
		ies, err := i.EthernetContextInformation()
		if err != nil {
			return nil, err
		}
		for _, x := range ies {
			if x.Type == MACAddressesDetected {
				return x.MACAddressesDetected()
			}
		}
		return nil, ErrIENotFound
	case UsageReportWithinSessionModificationResponse,
		UsageReportWithinSessionDeletionResponse,
		UsageReportWithinSessionReportRequest:
		ies, err := i.UsageReport()
		if err != nil {
			return nil, err
		}
		for _, x := range ies {
			if x.Type == EthernetTrafficInformation {
				return x.MACAddressesDetected()
			}
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// MACAddressesDetectedFields represents a fields contained in MACAddressesDetected IE.
type MACAddressesDetectedFields struct {
	NumberOfMACAddresses uint8
	MACAddresses         []net.HardwareAddr
	CTAGLength           uint8
	CTAG                 []byte
	STAGLength           uint8
	STAG                 []byte
}

// NewMACAddressesDetectedFields creates a new NewMACAddressesDetectedFields.
func NewMACAddressesDetectedFields(ctag, stag *IE, macs ...net.HardwareAddr) *MACAddressesDetectedFields {
	ct, st := ctag.Payload, stag.Payload
	return &MACAddressesDetectedFields{
		NumberOfMACAddresses: uint8(len(macs)),
		MACAddresses:         macs,
		CTAGLength:           uint8(len(ct)),
		CTAG:                 ct,
		STAGLength:           uint8(len(st)),
		STAG:                 st,
	}
}

// ParseMACAddressesDetectedFields parses b into MACAddressesDetectedFields.
func ParseMACAddressesDetectedFields(b []byte) (*MACAddressesDetectedFields, error) {
	f := &MACAddressesDetectedFields{}
	if err := f.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return f, nil
}

// UnmarshalBinary parses b into IE.
func (f *MACAddressesDetectedFields) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 1 {
		return io.ErrUnexpectedEOF
	}

	f.NumberOfMACAddresses = b[0]
	offset := 1

	for i := 0; i <= int(f.NumberOfMACAddresses); i++ {
		if l < offset+6 {
			return io.ErrUnexpectedEOF
		}
		f.MACAddresses = append(f.MACAddresses, b[offset:offset+6])
		offset += 6
	}

	if l < offset {
		return io.ErrUnexpectedEOF
	}
	f.CTAGLength = b[offset]

	if l < offset+int(f.CTAGLength) {
		return io.ErrUnexpectedEOF
	}
	copy(f.CTAG, b[offset:offset+int(f.CTAGLength)])

	if l < offset {
		return io.ErrUnexpectedEOF
	}
	f.STAGLength = b[offset]

	if l < offset+int(f.STAGLength) {
		return io.ErrUnexpectedEOF
	}
	copy(f.STAG, b[offset:offset+int(f.STAGLength)])

	return nil
}

// Marshal returns the serialized bytes of MACAddressesDetectedFields.
func (f *MACAddressesDetectedFields) Marshal() ([]byte, error) {
	b := make([]byte, f.MarshalLen())
	if err := f.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (f *MACAddressesDetectedFields) MarshalTo(b []byte) error {
	l := len(b)
	if l < 1 {
		return io.ErrUnexpectedEOF
	}

	b[0] = f.NumberOfMACAddresses
	offset := 1

	for _, mac := range f.MACAddresses {
		copy(b[offset:offset+6], mac)
		offset += 6
	}

	b[offset] = f.CTAGLength
	offset++
	copy(b[offset:offset+int(f.CTAGLength)], f.CTAG)
	offset += int(f.CTAGLength)

	b[offset] = f.STAGLength
	offset++
	copy(b[offset:offset+int(f.STAGLength)], f.STAG)

	return nil
}

// MarshalLen returns field length in integer.
func (f *MACAddressesDetectedFields) MarshalLen() int {
	l := 3
	l += int(f.NumberOfMACAddresses) * 6
	l += int(f.CTAGLength)
	l += int(f.STAGLength)

	return l
}
