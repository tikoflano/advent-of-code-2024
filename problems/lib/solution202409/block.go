package solution202409

type Block interface {
	IsFree() bool
	GetId() int
}
