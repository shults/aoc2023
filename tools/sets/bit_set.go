package sets

func NewBitSet() {

}

type BitSet interface {
	Clear()
	Set(int)
	Has(int)
	Del(int)
}

type bitSet struct {
}
