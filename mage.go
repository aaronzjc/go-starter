//go:build mage

package main

import (
	"github.com/magefile/mage/sh"
)

// run build
func Build() error {
	err := sh.RunV("go", "version")
	if err != nil {
		return err
	}
	return nil
}

// run gen mock test data
func Gen_mock() error {
	// mockery --dir ./internal/domain/repo --name [a-z]+Repo --output ./internal/mocks --outpkg mocks --case underscore --with-expecter
	mocks := map[string]string{
		"Repo":    "./internal/domain/repo",
		"Service": "./internal/application/service",
	}
	for name, dir := range mocks {
		err := sh.RunV(
			"mockery",
			"--dir", dir,
			"--name", "[a-z]+"+name,
			"--output", "./internal/mocks",
			"--outpkg", "mocks",
			"--case", "underscore",
			"--with-expecter",
		)
		if err != nil {
			return err
		}
	}
	return nil
}

// run gen proto files
func Gen_proto() error {
	inputs := []string{"hello.proto"} // os.Exec不支持通配符，只能手动了
	for _, v := range inputs {
		err := sh.RunV(
			"protoc",
			"--go_out=.",
			"--go_opt=paths=source_relative",
			"--go-grpc_out=.",
			"--go-grpc_opt=paths=source_relative",
			"./internal/pb/"+v,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

// run tests
func Test() error {
	err := sh.RunV(
		"go", "test", "-cover", "-coverprofile=coverage.out", "./...",
	)
	if err != nil {
		return err
	}
	return nil
}
