// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package util

import (
	"fmt"
	"io"
	"math"
	"sort"

	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

func (t *basicSchema) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)
	fieldCount := 8

	if t.Absent == nil {
		fieldCount--
	}

	if _, err := cw.Write(cbg.CborEncodeMajorType(cbg.MajMap, uint64(fieldCount))); err != nil {
		return err
	}

	// t.Bool (bool) (bool)
	if uint64(len("bool")) > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"bool\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("bool"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("bool")); err != nil {
		return err
	}

	if err := cbg.WriteBool(w, t.Bool); err != nil {
		return err
	}

	// t.Null (string) (string)
	if uint64(len("null")) > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"null\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("null"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("null")); err != nil {
		return err
	}

	if t.Null == nil {
		if _, err := cw.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if uint64(len(*t.Null)) > cbg.MaxLength {
			return xerrors.Errorf("Value in field t.Null was too long")
		}

		if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(*t.Null))); err != nil {
			return err
		}
		if _, err := cw.WriteString(string(*t.Null)); err != nil {
			return err
		}
	}

	// t.Array ([]string) (slice)
	if uint64(len("array")) > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"array\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("array"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("array")); err != nil {
		return err
	}

	if uint64(len(t.Array)) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Array was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.Array))); err != nil {
		return err
	}
	for _, v := range t.Array {
		if uint64(len(v)) > cbg.MaxLength {
			return xerrors.Errorf("Value in field v was too long")
		}

		if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(v))); err != nil {
			return err
		}
		if _, err := cw.WriteString(string(v)); err != nil {
			return err
		}

	}

	// t.Absent (string) (string)
	if t.Absent != nil {

		if uint64(len("absent")) > cbg.MaxLength {
			return xerrors.Errorf("Value in field \"absent\" was too long")
		}

		if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("absent"))); err != nil {
			return err
		}
		if _, err := cw.WriteString(string("absent")); err != nil {
			return err
		}

		if t.Absent == nil {
			if _, err := cw.Write(cbg.CborNull); err != nil {
				return err
			}
		} else {
			if uint64(len(*t.Absent)) > cbg.MaxLength {
				return xerrors.Errorf("Value in field t.Absent was too long")
			}

			if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(*t.Absent))); err != nil {
				return err
			}
			if _, err := cw.WriteString(string(*t.Absent)); err != nil {
				return err
			}
		}
	}

	// t.Object (util.basicSchemaInner) (struct)
	if uint64(len("object")) > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"object\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("object"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("object")); err != nil {
		return err
	}

	if err := t.Object.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.String (string) (string)
	if uint64(len("string")) > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"string\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("string"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("string")); err != nil {
		return err
	}

	if uint64(len(t.String)) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.String was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.String))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.String)); err != nil {
		return err
	}

	// t.Integer (int64) (int64)
	if uint64(len("integer")) > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"integer\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("integer"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("integer")); err != nil {
		return err
	}

	if t.Integer >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Integer)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.Integer-1)); err != nil {
			return err
		}
	}

	// t.Unicode (string) (string)
	if uint64(len("unicode")) > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"unicode\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("unicode"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("unicode")); err != nil {
		return err
	}

	if uint64(len(t.Unicode)) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Unicode was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.Unicode))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.Unicode)); err != nil {
		return err
	}
	return nil
}

func (t *basicSchema) UnmarshalCBOR(r io.Reader) (err error) {
	*t = basicSchema{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("basicSchema: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(cr)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Bool (bool) (bool)
		case "bool":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}
			if maj != cbg.MajOther {
				return fmt.Errorf("booleans must be major type 7")
			}
			switch extra {
			case 20:
				t.Bool = false
			case 21:
				t.Bool = true
			default:
				return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
			}
			// t.Null (string) (string)
		case "null":

			{
				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}

					sval, err := cbg.ReadString(cr)
					if err != nil {
						return err
					}

					t.Null = (*string)(&sval)
				}
			}
			// t.Array ([]string) (slice)
		case "array":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}

			if extra > cbg.MaxLength {
				return fmt.Errorf("t.Array: array too large (%d)", extra)
			}

			if maj != cbg.MajArray {
				return fmt.Errorf("expected cbor array")
			}

			if extra > 0 {
				t.Array = make([]string, extra)
			}

			for i := 0; i < int(extra); i++ {
				{
					var maj byte
					var extra uint64
					var err error
					_ = maj
					_ = extra
					_ = err

					{
						sval, err := cbg.ReadString(cr)
						if err != nil {
							return err
						}

						t.Array[i] = string(sval)
					}

				}
			}
			// t.Absent (string) (string)
		case "absent":

			{
				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}

					sval, err := cbg.ReadString(cr)
					if err != nil {
						return err
					}

					t.Absent = (*string)(&sval)
				}
			}
			// t.Object (util.basicSchemaInner) (struct)
		case "object":

			{

				if err := t.Object.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.Object: %w", err)
				}

			}
			// t.String (string) (string)
		case "string":

			{
				sval, err := cbg.ReadString(cr)
				if err != nil {
					return err
				}

				t.String = string(sval)
			}
			// t.Integer (int64) (int64)
		case "integer":
			{
				maj, extra, err := cr.ReadHeader()
				var extraI int64
				if err != nil {
					return err
				}
				switch maj {
				case cbg.MajUnsignedInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 positive overflow")
					}
				case cbg.MajNegativeInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 negative overflow")
					}
					extraI = -1 - extraI
				default:
					return fmt.Errorf("wrong type for int64 field: %d", maj)
				}

				t.Integer = int64(extraI)
			}
			// t.Unicode (string) (string)
		case "unicode":

			{
				sval, err := cbg.ReadString(cr)
				if err != nil {
					return err
				}

				t.Unicode = string(sval)
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *basicSchemaInner) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{164}); err != nil {
		return err
	}

	// t.Arr ([]string) (slice)
	if uint64(len("arr")) > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"arr\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("arr"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("arr")); err != nil {
		return err
	}

	if uint64(len(t.Arr)) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Arr was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.Arr))); err != nil {
		return err
	}
	for _, v := range t.Arr {
		if uint64(len(v)) > cbg.MaxLength {
			return xerrors.Errorf("Value in field v was too long")
		}

		if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(v))); err != nil {
			return err
		}
		if _, err := cw.WriteString(string(v)); err != nil {
			return err
		}

	}

	// t.Bool (bool) (bool)
	if uint64(len("bool")) > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"bool\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("bool"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("bool")); err != nil {
		return err
	}

	if err := cbg.WriteBool(w, t.Bool); err != nil {
		return err
	}

	// t.Number (int64) (int64)
	if uint64(len("number")) > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"number\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("number"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("number")); err != nil {
		return err
	}

	if t.Number >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Number)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.Number-1)); err != nil {
			return err
		}
	}

	// t.String (string) (string)
	if uint64(len("string")) > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"string\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("string"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("string")); err != nil {
		return err
	}

	if uint64(len(t.String)) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.String was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.String))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.String)); err != nil {
		return err
	}
	return nil
}

func (t *basicSchemaInner) UnmarshalCBOR(r io.Reader) (err error) {
	*t = basicSchemaInner{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("basicSchemaInner: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(cr)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Arr ([]string) (slice)
		case "arr":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}

			if extra > cbg.MaxLength {
				return fmt.Errorf("t.Arr: array too large (%d)", extra)
			}

			if maj != cbg.MajArray {
				return fmt.Errorf("expected cbor array")
			}

			if extra > 0 {
				t.Arr = make([]string, extra)
			}

			for i := 0; i < int(extra); i++ {
				{
					var maj byte
					var extra uint64
					var err error
					_ = maj
					_ = extra
					_ = err

					{
						sval, err := cbg.ReadString(cr)
						if err != nil {
							return err
						}

						t.Arr[i] = string(sval)
					}

				}
			}
			// t.Bool (bool) (bool)
		case "bool":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}
			if maj != cbg.MajOther {
				return fmt.Errorf("booleans must be major type 7")
			}
			switch extra {
			case 20:
				t.Bool = false
			case 21:
				t.Bool = true
			default:
				return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
			}
			// t.Number (int64) (int64)
		case "number":
			{
				maj, extra, err := cr.ReadHeader()
				var extraI int64
				if err != nil {
					return err
				}
				switch maj {
				case cbg.MajUnsignedInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 positive overflow")
					}
				case cbg.MajNegativeInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 negative overflow")
					}
					extraI = -1 - extraI
				default:
					return fmt.Errorf("wrong type for int64 field: %d", maj)
				}

				t.Number = int64(extraI)
			}
			// t.String (string) (string)
		case "string":

			{
				sval, err := cbg.ReadString(cr)
				if err != nil {
					return err
				}

				t.String = string(sval)
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *ipldSchema) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{163}); err != nil {
		return err
	}

	// t.A (util.LexLink) (struct)
	if uint64(len("a")) > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"a\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("a"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("a")); err != nil {
		return err
	}

	if err := t.A.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.B (util.LexBytes) (slice)
	if uint64(len("b")) > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"b\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("b"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("b")); err != nil {
		return err
	}

	if uint64(len(t.B)) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.B was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.B))); err != nil {
		return err
	}

	if _, err := cw.Write(t.B); err != nil {
		return err
	}

	// t.C (util.LexBlob) (struct)
	if uint64(len("c")) > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"c\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("c"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("c")); err != nil {
		return err
	}

	if err := t.C.MarshalCBOR(cw); err != nil {
		return err
	}
	return nil
}

func (t *ipldSchema) UnmarshalCBOR(r io.Reader) (err error) {
	*t = ipldSchema{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("ipldSchema: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(cr)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.A (util.LexLink) (struct)
		case "a":

			{

				if err := t.A.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.A: %w", err)
				}

			}
			// t.B (util.LexBytes) (slice)
		case "b":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}

			if extra > cbg.ByteArrayMaxLen {
				return fmt.Errorf("t.B: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}

			if extra > 0 {
				t.B = make([]uint8, extra)
			}

			if _, err := io.ReadFull(cr, t.B); err != nil {
				return err
			}

			// t.C (util.LexBlob) (struct)
		case "c":

			{

				if err := t.C.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.C: %w", err)
				}

			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *basicOldSchema) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)
	fieldCount := 7

	if t.D == nil {
		fieldCount--
	}

	if _, err := cw.Write(cbg.CborEncodeMajorType(cbg.MajMap, uint64(fieldCount))); err != nil {
		return err
	}

	// t.A (string) (string)
	if uint64(len("a")) > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"a\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("a"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("a")); err != nil {
		return err
	}

	if uint64(len(t.A)) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.A was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.A))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.A)); err != nil {
		return err
	}

	// t.B (int64) (int64)
	if uint64(len("b")) > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"b\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("b"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("b")); err != nil {
		return err
	}

	if t.B >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.B)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.B-1)); err != nil {
			return err
		}
	}

	// t.C (bool) (bool)
	if uint64(len("c")) > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"c\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("c"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("c")); err != nil {
		return err
	}

	if err := cbg.WriteBool(w, t.C); err != nil {
		return err
	}

	// t.D (string) (string)
	if t.D != nil {

		if uint64(len("d")) > cbg.MaxLength {
			return xerrors.Errorf("Value in field \"d\" was too long")
		}

		if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("d"))); err != nil {
			return err
		}
		if _, err := cw.WriteString(string("d")); err != nil {
			return err
		}

		if t.D == nil {
			if _, err := cw.Write(cbg.CborNull); err != nil {
				return err
			}
		} else {
			if uint64(len(*t.D)) > cbg.MaxLength {
				return xerrors.Errorf("Value in field t.D was too long")
			}

			if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(*t.D))); err != nil {
				return err
			}
			if _, err := cw.WriteString(string(*t.D)); err != nil {
				return err
			}
		}
	}

	// t.E (string) (string)
	if uint64(len("e")) > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"e\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("e"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("e")); err != nil {
		return err
	}

	if t.E == nil {
		if _, err := cw.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if uint64(len(*t.E)) > cbg.MaxLength {
			return xerrors.Errorf("Value in field t.E was too long")
		}

		if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(*t.E))); err != nil {
			return err
		}
		if _, err := cw.WriteString(string(*t.E)); err != nil {
			return err
		}
	}

	// t.F ([]string) (slice)
	if uint64(len("f")) > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"f\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("f"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("f")); err != nil {
		return err
	}

	if uint64(len(t.F)) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.F was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.F))); err != nil {
		return err
	}
	for _, v := range t.F {
		if uint64(len(v)) > cbg.MaxLength {
			return xerrors.Errorf("Value in field v was too long")
		}

		if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(v))); err != nil {
			return err
		}
		if _, err := cw.WriteString(string(v)); err != nil {
			return err
		}

	}

	// t.G (util.basicOldSchemaInner) (struct)
	if uint64(len("g")) > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"g\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("g"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("g")); err != nil {
		return err
	}

	if err := t.G.MarshalCBOR(cw); err != nil {
		return err
	}
	return nil
}

func (t *basicOldSchema) UnmarshalCBOR(r io.Reader) (err error) {
	*t = basicOldSchema{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("basicOldSchema: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(cr)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.A (string) (string)
		case "a":

			{
				sval, err := cbg.ReadString(cr)
				if err != nil {
					return err
				}

				t.A = string(sval)
			}
			// t.B (int64) (int64)
		case "b":
			{
				maj, extra, err := cr.ReadHeader()
				var extraI int64
				if err != nil {
					return err
				}
				switch maj {
				case cbg.MajUnsignedInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 positive overflow")
					}
				case cbg.MajNegativeInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 negative overflow")
					}
					extraI = -1 - extraI
				default:
					return fmt.Errorf("wrong type for int64 field: %d", maj)
				}

				t.B = int64(extraI)
			}
			// t.C (bool) (bool)
		case "c":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}
			if maj != cbg.MajOther {
				return fmt.Errorf("booleans must be major type 7")
			}
			switch extra {
			case 20:
				t.C = false
			case 21:
				t.C = true
			default:
				return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
			}
			// t.D (string) (string)
		case "d":

			{
				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}

					sval, err := cbg.ReadString(cr)
					if err != nil {
						return err
					}

					t.D = (*string)(&sval)
				}
			}
			// t.E (string) (string)
		case "e":

			{
				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}

					sval, err := cbg.ReadString(cr)
					if err != nil {
						return err
					}

					t.E = (*string)(&sval)
				}
			}
			// t.F ([]string) (slice)
		case "f":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}

			if extra > cbg.MaxLength {
				return fmt.Errorf("t.F: array too large (%d)", extra)
			}

			if maj != cbg.MajArray {
				return fmt.Errorf("expected cbor array")
			}

			if extra > 0 {
				t.F = make([]string, extra)
			}

			for i := 0; i < int(extra); i++ {
				{
					var maj byte
					var extra uint64
					var err error
					_ = maj
					_ = extra
					_ = err

					{
						sval, err := cbg.ReadString(cr)
						if err != nil {
							return err
						}

						t.F[i] = string(sval)
					}

				}
			}
			// t.G (util.basicOldSchemaInner) (struct)
		case "g":

			{

				if err := t.G.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.G: %w", err)
				}

			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *basicOldSchemaInner) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{164}); err != nil {
		return err
	}

	// t.H (string) (string)
	if uint64(len("h")) > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"h\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("h"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("h")); err != nil {
		return err
	}

	if uint64(len(t.H)) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.H was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.H))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.H)); err != nil {
		return err
	}

	// t.I (int64) (int64)
	if uint64(len("i")) > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"i\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("i"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("i")); err != nil {
		return err
	}

	if t.I >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.I)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.I-1)); err != nil {
			return err
		}
	}

	// t.J (bool) (bool)
	if uint64(len("j")) > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"j\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("j"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("j")); err != nil {
		return err
	}

	if err := cbg.WriteBool(w, t.J); err != nil {
		return err
	}

	// t.K ([]string) (slice)
	if uint64(len("k")) > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"k\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("k"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("k")); err != nil {
		return err
	}

	if uint64(len(t.K)) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.K was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.K))); err != nil {
		return err
	}
	for _, v := range t.K {
		if uint64(len(v)) > cbg.MaxLength {
			return xerrors.Errorf("Value in field v was too long")
		}

		if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(v))); err != nil {
			return err
		}
		if _, err := cw.WriteString(string(v)); err != nil {
			return err
		}

	}
	return nil
}

func (t *basicOldSchemaInner) UnmarshalCBOR(r io.Reader) (err error) {
	*t = basicOldSchemaInner{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("basicOldSchemaInner: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(cr)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.H (string) (string)
		case "h":

			{
				sval, err := cbg.ReadString(cr)
				if err != nil {
					return err
				}

				t.H = string(sval)
			}
			// t.I (int64) (int64)
		case "i":
			{
				maj, extra, err := cr.ReadHeader()
				var extraI int64
				if err != nil {
					return err
				}
				switch maj {
				case cbg.MajUnsignedInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 positive overflow")
					}
				case cbg.MajNegativeInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 negative overflow")
					}
					extraI = -1 - extraI
				default:
					return fmt.Errorf("wrong type for int64 field: %d", maj)
				}

				t.I = int64(extraI)
			}
			// t.J (bool) (bool)
		case "j":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}
			if maj != cbg.MajOther {
				return fmt.Errorf("booleans must be major type 7")
			}
			switch extra {
			case 20:
				t.J = false
			case 21:
				t.J = true
			default:
				return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
			}
			// t.K ([]string) (slice)
		case "k":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}

			if extra > cbg.MaxLength {
				return fmt.Errorf("t.K: array too large (%d)", extra)
			}

			if maj != cbg.MajArray {
				return fmt.Errorf("expected cbor array")
			}

			if extra > 0 {
				t.K = make([]string, extra)
			}

			for i := 0; i < int(extra); i++ {
				{
					var maj byte
					var extra uint64
					var err error
					_ = maj
					_ = extra
					_ = err

					{
						sval, err := cbg.ReadString(cr)
						if err != nil {
							return err
						}

						t.K[i] = string(sval)
					}

				}
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *ipldOldSchema) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{162}); err != nil {
		return err
	}

	// t.A (util.LexLink) (struct)
	if uint64(len("a")) > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"a\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("a"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("a")); err != nil {
		return err
	}

	if err := t.A.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.B (util.LexBytes) (slice)
	if uint64(len("b")) > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"b\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("b"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("b")); err != nil {
		return err
	}

	if uint64(len(t.B)) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.B was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.B))); err != nil {
		return err
	}

	if _, err := cw.Write(t.B); err != nil {
		return err
	}

	return nil
}

func (t *ipldOldSchema) UnmarshalCBOR(r io.Reader) (err error) {
	*t = ipldOldSchema{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("ipldOldSchema: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(cr)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.A (util.LexLink) (struct)
		case "a":

			{

				if err := t.A.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.A: %w", err)
				}

			}
			// t.B (util.LexBytes) (slice)
		case "b":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}

			if extra > cbg.ByteArrayMaxLen {
				return fmt.Errorf("t.B: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}

			if extra > 0 {
				t.B = make([]uint8, extra)
			}

			if _, err := io.ReadFull(cr, t.B); err != nil {
				return err
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
