package writer

type Wrapper interface {
	Write(content []byte) (int, error)
}
