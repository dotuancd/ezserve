package fs

type Filesystem struct {
	driver Driver
}

type Driver interface {
	put(path string, content []byte) error
	get(path string) (content []byte, err error)
}

type Local struct {
	root string
}

func NewLocal(root string) *Local {
	return &Local{root: root}
}

func (f *Local) put(path string, content []byte) error {
	panic("implement me")
}

func (f *Local) get(path string) (content []byte, err error) {
	panic("implement me")
}


