package utils

type Bit uint8

const (
	UP Bit = 1 << iota
	DOWN
	LEFT
	RIGHT
)

func Set(b, flag Bit) Bit    { return b | flag }
func Clear(b, flag Bit) Bit  { return b &^ flag }
func Toggle(b, flag Bit) Bit { return b ^ flag }
func Has(b, flag Bit) bool   { return b&flag != 0 }

//func (b *Bit) Set(flag Bit) *Bit {
//	*b = *b | flag
//	return b
//}
//func (b *Bit) Clear(flag Bit) *Bit {
//	*b = *b &^ flag
//	return b
//}
//func (b *Bit) Toggle(flag Bit) *Bit {
//	*b = *b ^ flag
//	return b
//}

func (b *Bit) Has(flag Bit) bool { return *b&flag != 0 }
