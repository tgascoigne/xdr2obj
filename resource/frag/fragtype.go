package frag

import (
	"github.com/tgascoigne/xdr2obj/resource"
	"github.com/tgascoigne/xdr2obj/resource/drawable"
	"github.com/tgascoigne/xdr2obj/resource/types"
)

type FragTypeHeader struct {
	_        uint32
	BlockMap types.Ptr32
	_        uint32
	_        uint32
	_        uint32
	_        uint32
	_        uint32
	_        uint32
	Drawable types.Ptr32
}

type FragType struct {
	Header FragTypeHeader
	drawable.Drawable
}

func (frag *FragType) Unpack(res *resource.Container) error {
	res.Parse(&frag.Header)

	if err := res.Detour(frag.Header.Drawable, func() error {
		frag.Drawable.Unpack(res)
		return nil
	}); err != nil {
		return err
	}
	return nil
}
