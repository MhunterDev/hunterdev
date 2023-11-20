package main

import (
	api "github.com/MhunterDev/hunterdev/src/web"
	initfs "github.com/Mhunterdev/hunterdev/src/base/initfs"
)

func main() {

	initfs.BuildFS()
	api.Router()
}
