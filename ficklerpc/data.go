package ficklerpc

// Args is struct utilized by TCP Server and Client to exchange data in pickled state.
type Args struct {
	Action string
	Id     string
	Blob   []byte
	Err    error
}
