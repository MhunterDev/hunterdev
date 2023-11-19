package main

import (
	fs "github.com/MhunterDev/hunterdev/src/base/initfs"
	db "github.com/MhunterDev/hunterdev/src/db"
)

func main() {

	fs.BuildFS()
	db.AddDefault()
}
