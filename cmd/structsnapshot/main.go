package main

import (
	"fmt"
	"go/types"
	"os"

	"github.com/gustavosbarreto/structsnapshot"
	"golang.org/x/tools/go/packages"
)

func getFields(structType *types.Struct) []structsnapshot.FieldSnapshot {
	var fields []structsnapshot.FieldSnapshot

	for i := 0; i < structType.NumFields(); i++ {
		if structType.Field(i).Embedded() {
			namedType, ok := structType.Field(i).Type().(*types.Named)
			if !ok {
				continue
			}

			structType, ok := namedType.Underlying().(*types.Struct)
			if !ok {
				continue
			}

			fields = append(fields, getFields(structType)...)
			continue
		}

		field := structsnapshot.FieldSnapshot{
			Name: structType.Field(i).Name(),
			Type: structType.Field(i).Type().String(),
			Tag:  structType.Tag(i),
		}

		fields = append(fields, field)
	}

	return fields
}

func TakeSnapshot(obj types.Object) (*structsnapshot.Snapshot, error) {
	snapshot := &structsnapshot.Snapshot{
		Name: obj.Name(),
	}

	namedType, ok := obj.Type().(*types.Named)
	if !ok {
		return nil, fmt.Errorf("%s is not a defined type", obj.Type())
	}

	structType, ok := namedType.Underlying().(*types.Struct)
	if !ok {
		return nil, fmt.Errorf("%s is not a struct type", obj.Type())
	}

	snapshot.Fields = getFields(structType)

	return snapshot, nil
}

func main() {
	if len(os.Args) < 1 {
		panic("Usage: structsnapshot [struct]")
	}

	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedFiles |
			packages.NeedSyntax |
			packages.NeedTypes |
			packages.NeedTypesInfo},
		"")
	if err != nil {
		panic(err)
	}

	snapshot, err := TakeSnapshot(pkgs[0].Types.Scope().Lookup(os.Args[1]))
	if err != nil {
		panic(err)
	}

	if err := snapshot.SaveToFile(); err != nil {
		panic(err)
	}
}
