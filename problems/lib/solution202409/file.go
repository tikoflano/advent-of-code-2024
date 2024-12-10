package solution202409

type File struct {
	Id int
}

func (file *File) IsFree() bool {
	return false
}

func (file *File) GetId() int {
	return file.Id
}
