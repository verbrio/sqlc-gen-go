package golang

import "github.com/sqlc-dev/sqlc-gen-go/internal/opts"

func parseDriver(sqlPackage string) opts.SQLDriver {
	switch sqlPackage {
	case opts.SQLPackagePGXV4:
		return opts.SQLDriverPGXV4
	case opts.SQLPackagePGXV5:
		return opts.SQLDriverPGXV5
	case opts.SQLPackageGoFr:
		return opts.SQLDriverGoFr
	default:
		return opts.SQLDriverLibPQ
	}
}
