package drawable

import (
	"log"

	"github.com/tgascoigne/xdr2obj/resource"
	"github.com/tgascoigne/xdr2obj/resource/types"
)

type GeometryHeader struct {
	_             uint32 /* vtable */
	_             uint32
	_             uint32
	VertexBuffer  types.Ptr32
	_             uint32
	_             uint32
	_             uint32
	IndexBuffer   types.Ptr32
	_             uint32
	_             uint32
	_             uint32
	IndexCount    uint32
	FaceCount     uint32
	VertexCount   uint16
	PrimitiveType uint16
}

type Geometry struct {
	GeometryHeader
	Vertices VertexBuffer
	Indices  IndexBuffer
}

func (geom *Geometry) Unpack(res *resource.Container) error {
	if err := res.Parse(&geom.GeometryHeader); err != nil {
		return err
	}

	if err := res.Detour(geom.VertexBuffer, func() error {
		return geom.Vertices.Unpack(res)
	}); err != nil {
		log.Printf("error parsing vertex buffer")
	}

	if err := res.Detour(geom.IndexBuffer, func() error {
		return geom.Indices.Unpack(res)
	}); err != nil {
		log.Printf("error parsing index buffer")
	}
	return nil
}
