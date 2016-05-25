package frame

import (
	"bytes"
	"errors"
)

type APICSequencer interface {
	Framer

	AddPicture(PictureFramer)
	Picture(picType string) (PictureFramer, error)
}

// APICSequence stores several APICs and implements interface Framer.
// Key for APICSequnce is a key for PictureType array,
// so there is only one picture with the same picture type
type APICSequence struct {
	sequence map[int]PictureFramer
	id       string
}

func NewAPICSequence() *APICSequence {
	return &APICSequence{
		sequence: make(map[int]PictureFramer),
	}
}

func (as APICSequence) Form() []byte {
	b := bytesBufPool.Get().(*bytes.Buffer)
	b.Reset()
	for _, pf := range as.sequence {
		frame := pf.Form()
		b.Write(frame)
	}
	bytesBufPool.Put(b)
	return b.Bytes()
}

func (as APICSequence) ID() string {
	return as.id
}

func (as *APICSequence) SetID(id string) {
	as.id = id
}

//TODO: if PictureType not found
func (as *APICSequence) AddPicture(pic PictureFramer) {
	for k, v := range PictureTypes {
		if v == pic.PictureType() {
			as.sequence[k] = pic
			break
		}
	}
}

func (as APICSequence) Picture(picType string) (PictureFramer, error) {
	for k, v := range PictureTypes {
		if v == picType {
			return as.sequence[k], nil
		}
	}
	return &PictureFrame{}, errors.New("Unsupported picture type")
}
