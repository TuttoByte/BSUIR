package savers

type SaverFacade struct {
	cmp  *Compressor
	enc  *Encrypter
	mrsh *Marshaller
}

func NewServerFacade() *SaverFacade {
	return &SaverFacade{
		cmp:  NewCompressor(),
		enc:  NewEncrypter(),
		mrsh: NewMarshaller(),
	}
}
