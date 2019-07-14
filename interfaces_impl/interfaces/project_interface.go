package interfaces

type Project interface {
	GetProjectsList(...string)
	GetFilesList(...string)
	GetFilesStats(...string)
	UpdateFileName(...string)
	UpdateFileMetadata(...string)
	DownloadFile(...string)
}