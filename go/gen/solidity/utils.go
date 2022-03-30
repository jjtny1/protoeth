package solidity

import ()

func leftPadByteSliceZero32(byteSlice []byte) []byte {
	return leftPadByteSliceToSizeWithZeros(byteSlice, 32)
}

func rightPadBytesSliceZero32(byteSlice []byte) []byte {
	if len(byteSlice) >= 32 {
		return byteSlice
	}

	rightPaddedBytes := make([]byte, 32)
	copy(rightPaddedBytes, byteSlice)
	return rightPaddedBytes
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

func padRightMultipleOf32(byteSlice []byte) []byte {
	length := len(byteSlice)

	if length%32 == 0 {
		return byteSlice
	}

	nextMultipleOf32 := ((length / 32) + 1) * 32
	rightPaddedBytes := make([]byte, nextMultipleOf32)
	copy(rightPaddedBytes, byteSlice)
	return rightPaddedBytes
}
