package encoder

func leftPadByteSliceZero32(byteSlice []byte) []byte {
	return leftPadByteSliceToSizeWithZeros(byteSlice, 32)
}

func leftPadBytSliceFF32(byteSlice []byte) []byte {
	leftPaddedBytes := make([]byte, 32)
	for i := 0; i < len(leftPaddedBytes); i = i + 1 {
		leftPaddedBytes[i] = 0xff
	}

	copy(leftPaddedBytes[32-len(byteSlice):], byteSlice)
	return leftPaddedBytes
}

func leftPadByteSliceToSizeWithZeros(byteSlice []byte, size int) []byte {
	if size <= len(byteSlice) {
		return byteSlice
	}

	leftPaddedBytes := make([]byte, size)
	copy(leftPaddedBytes[size-len(byteSlice):], byteSlice)
	return leftPaddedBytes
}
