package solution202409

type BlockParser func(length int) []Block

func makeFileParser() BlockParser {
	nextId := 0

	return func(length int) (ret []Block) {
		for i := 0; i < length; i++ {
			ret = append(ret, &File{Id: nextId})
		}
		nextId++
		return
	}
}

func makeFreeMemoryParser() BlockParser {
	return func(length int) (ret []Block) {
		for i := 0; i < length; i++ {
			ret = append(ret, &FreeMemory{})
		}
		return
	}
}
