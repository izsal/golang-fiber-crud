package gormschema

import (
	"io"
	"my-fiber-app/entities"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
)

func LoadGormSchema() (string, error) {
	stmts, err := gormschema.New("mysql").Load(&entities.Task{})
	if err != nil {
		return "", err
	}

	return stmts, nil
}

func PrintSchema(stmts string) {
	io.WriteString(os.Stdout, stmts)
}
