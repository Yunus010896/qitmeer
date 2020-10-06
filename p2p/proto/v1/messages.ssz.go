// Code generated by fastssz. DO NOT EDIT.
package qitmeer_p2p_v1

import (
	"fmt"

	ssz "github.com/ferranbt/fastssz"
)

var (
	errDivideInt           = fmt.Errorf("incorrect int divide")
	errListTooBig          = fmt.Errorf("incorrect list size, too big")
	errMarshalDynamicBytes = fmt.Errorf("incorrect dynamic bytes marshalling")
	errMarshalFixedBytes   = fmt.Errorf("incorrect fixed bytes marshalling")
	errMarshalList         = fmt.Errorf("incorrect vector list")
	errMarshalVector       = fmt.Errorf("incorrect vector marshalling")
	errOffset              = fmt.Errorf("incorrect offset")
	errSize                = fmt.Errorf("incorrect size")
)

// MarshalSSZ ssz marshals the ErrorResponse object
func (e *ErrorResponse) MarshalSSZ() ([]byte, error) {
	buf := make([]byte, e.SizeSSZ())
	return e.MarshalSSZTo(buf[:0])
}

// MarshalSSZTo ssz marshals the ErrorResponse object to a target array
func (e *ErrorResponse) MarshalSSZTo(dst []byte) ([]byte, error) {
	var err error
	offset := int(4)

	// Offset (0) 'Message'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(e.Message)

	// Field (0) 'Message'
	if len(e.Message) > 256 {
		return nil, errMarshalDynamicBytes
	}
	dst = append(dst, e.Message...)

	return dst, err
}

// UnmarshalSSZ ssz unmarshals the ErrorResponse object
func (e *ErrorResponse) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 4 {
		return errSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'Message'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return errOffset
	}

	// Field (0) 'Message'
	{
		buf = tail[o0:]
		e.Message = append(e.Message, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ErrorResponse object
func (e *ErrorResponse) SizeSSZ() (size int) {
	size = 4

	// Field (0) 'Message'
	size += len(e.Message)

	return
}

// MarshalSSZ ssz marshals the MetaData object
func (m *MetaData) MarshalSSZ() ([]byte, error) {
	buf := make([]byte, m.SizeSSZ())
	return m.MarshalSSZTo(buf[:0])
}

// MarshalSSZTo ssz marshals the MetaData object to a target array
func (m *MetaData) MarshalSSZTo(dst []byte) ([]byte, error) {
	var err error

	// Field (0) 'SeqNumber'
	dst = ssz.MarshalUint64(dst, m.SeqNumber)

	// Field (1) 'Subnets'
	if dst, err = ssz.MarshalFixedBytes(dst, m.Subnets, 8); err != nil {
		return nil, errMarshalFixedBytes
	}

	return dst, err
}

// UnmarshalSSZ ssz unmarshals the MetaData object
func (m *MetaData) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 16 {
		return errSize
	}

	// Field (0) 'SeqNumber'
	m.SeqNumber = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'Subnets'
	m.Subnets = append(m.Subnets, buf[8:16]...)

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the MetaData object
func (m *MetaData) SizeSSZ() (size int) {
	size = 16
	return
}

// MarshalSSZ ssz marshals the ChainState object
func (c *ChainState) MarshalSSZ() ([]byte, error) {
	buf := make([]byte, c.SizeSSZ())
	return c.MarshalSSZTo(buf[:0])
}

// MarshalSSZTo ssz marshals the ChainState object to a target array
func (c *ChainState) MarshalSSZTo(dst []byte) ([]byte, error) {
	var err error
	offset := int(60)

	// Field (0) 'GenesisHash'
	if dst, err = ssz.MarshalFixedBytes(dst, c.GenesisHash, 32); err != nil {
		return nil, errMarshalFixedBytes
	}

	// Field (1) 'ProtocolVersion'
	dst = ssz.MarshalUint32(dst, c.ProtocolVersion)

	// Field (2) 'Timestamp'
	dst = ssz.MarshalUint64(dst, c.Timestamp)

	// Field (3) 'Services'
	dst = ssz.MarshalUint64(dst, c.Services)

	// Field (4) 'GraphState'
	dst = ssz.MarshalUint32(dst, c.GraphState)

	// Offset (5) 'UserAgent'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(c.UserAgent)

	// Field (5) 'UserAgent'
	if len(c.UserAgent) > 256 {
		return nil, errMarshalDynamicBytes
	}
	dst = append(dst, c.UserAgent...)

	return dst, err
}

// UnmarshalSSZ ssz unmarshals the ChainState object
func (c *ChainState) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 60 {
		return errSize
	}

	tail := buf
	var o5 uint64

	// Field (0) 'GenesisHash'
	c.GenesisHash = append(c.GenesisHash, buf[0:32]...)

	// Field (1) 'ProtocolVersion'
	c.ProtocolVersion = ssz.UnmarshallUint32(buf[32:36])

	// Field (2) 'Timestamp'
	c.Timestamp = ssz.UnmarshallUint64(buf[36:44])

	// Field (3) 'Services'
	c.Services = ssz.UnmarshallUint64(buf[44:52])

	// Field (4) 'GraphState'
	c.GraphState = ssz.UnmarshallUint32(buf[52:56])

	// Offset (5) 'UserAgent'
	if o5 = ssz.ReadOffset(buf[56:60]); o5 > size {
		return errOffset
	}

	// Field (5) 'UserAgent'
	{
		buf = tail[o5:]
		c.UserAgent = append(c.UserAgent, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ChainState object
func (c *ChainState) SizeSSZ() (size int) {
	size = 60

	// Field (5) 'UserAgent'
	size += len(c.UserAgent)

	return
}
