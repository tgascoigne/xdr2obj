package obj

import (
	"fmt"
	"os"

	"github.com/tgascoigne/xdr2obj/export"
)

type Context struct {
	ObjFile *os.File
	MtlFile *os.File
	NextID  int
}

func Export(object export.Exportable) error {
	var err error
	ctx := Context{}

	if ctx.ObjFile, err = os.Create(fmt.Sprintf("%v.obj", object.GetName())); err != nil {
		return err
	}
	defer ctx.ObjFile.Close()

	if ctx.MtlFile, err = os.Create(fmt.Sprintf("%v.mtl", object.GetName())); err != nil {
		return err
	}
	defer ctx.MtlFile.Close()

	fmt.Fprintf(ctx.ObjFile, "mtllib %v.mtl\n", object.GetName())

	for _, model := range object.GetModels() {
		if err := ExportModel(&ctx, model); err != nil {
			return err
		}
	}
	return nil
}

func (ctx *Context) Unique(name string) string {
	ctx.NextID++
	return fmt.Sprintf("%v_%.4d", name, ctx.NextID-1)
}
