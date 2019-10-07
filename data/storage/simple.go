package storage

const (
	dataSliceSizeBytes = 1024
)

type (
	Simple interface{
		Get(from uint64, to uint64) [][]byte
		GetPages(from int, to int) [][]byte
		Put([]byte) error
	}

	simple struct {
		data [][]byte
	}
)

func NewSimpleStorage() Simple {
	return &simple{[][]byte{}}
}

/*
	return pages from the gap
 */
func (s *simple) GetPages(from int, to int) [][]byte {
	if from > to || to >= len(s.data) {
		return [][]byte{}
	}
	return s.data[from:to]
}

/*
	returns data from the specified byte gap
 */
func (s *simple) Get(from uint64, to uint64) [][]byte {
	if to <= from {
		return [][]byte{}
	}

	pageIdxFrom := int(from / dataSliceSizeBytes)
	byteIdxFrom := int(from % dataSliceSizeBytes)
	if pageIdxFrom >= len(s.data) || byteIdxFrom >= len(s.data[pageIdxFrom]) {
		return [][]byte{}
	}

	pageIdxTo := int(from / dataSliceSizeBytes)
	byteIdxTo := int(from % dataSliceSizeBytes)
	if pageIdxTo >= len(s.data) || byteIdxTo >= len(s.data[pageIdxTo]) {
		pageIdxTo = len(s.data)-1
		byteIdxTo = len(s.data[pageIdxTo])-1
	}

	ret := make([][]byte, 0, pageIdxTo - pageIdxFrom)
	byteIdxFrom-- // crutch to negotiate first increment of nextIdx function
	nextIdx := func() (int, int) {
		byteIdxFrom++
		if byteIdxFrom == len(s.data[pageIdxFrom]) {
			 pageIdxFrom++
			 byteIdxFrom = 0
		}
		return pageIdxFrom, byteIdxFrom
	}

	retPageIdx, retByteIdx := 0, 0
	for pageIdxFrom != pageIdxTo && byteIdxFrom != byteIdxFrom {
		if retByteIdx == 0 {
			s.data = append(s.data, make([]byte, dataSliceSizeBytes))
		}

		p, b := nextIdx()
		ret[retPageIdx][retByteIdx] = s.data[p][b]
		retByteIdx++
		if retByteIdx == dataSliceSizeBytes {
			retByteIdx = 0
			retByteIdx++
		}
	}

	return ret
}

func (s *simple) Put(data []byte) error {
	curPage := len(s.data)-1
	if len(s.data) == 0 {
		// appending the first page
		curPage = 0
		s.data = append(s.data, make([]byte, 0, dataSliceSizeBytes))
	}

	curByte := len(s.data[curPage]) - 1
	if len(s.data[curPage]) == 0 {
		curByte = 0
	}

	for dataSliceSizeBytes - curByte < len(data) {
		s.data[curPage] = append(s.data[curPage], data[:dataSliceSizeBytes - curByte]...)
		s.data = append(s.data, make([]byte, 0, dataSliceSizeBytes))
		data = data[dataSliceSizeBytes - curByte + 1:]
		curPage++
		curByte = 0
	}
	s.data[curPage] = append(s.data[curPage], data[:]...)

	return nil
}