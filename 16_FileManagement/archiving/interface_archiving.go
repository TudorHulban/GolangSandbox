package archiving

// IArchive Interface for working with different type of archiving.
type IArchive interface {
	CompressFile(fileName string) error
}
