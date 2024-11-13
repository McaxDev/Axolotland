package backup

import "github.com/mholt/archiver/v3"

var Compressor *archiver.TarGz

func Init() {
	Compressor = archiver.NewTarGz()
}
