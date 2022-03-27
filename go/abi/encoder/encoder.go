package encoder

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/jjtny1/protoeth/go/gen/solidity"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

type (
	EncodableProto interface {
		// Used to get static information about proto types
		ProtoReflect() protoreflect.Message
	}

	Encoder interface {
		Encode(bytes.Buffer, EncodableProto) error
	}

	staticTypeEncoder struct{}
)

func Encode(parameters []EncodableProto) ([]byte, error) {
	var buffer bytes.Buffer

	for _, parameter := range parameters {
		if err := encode(buffer, parameter); err != nil {
			return nil, err
		}
	}

	return buffer.Bytes(), nil
}

func encode(buffer bytes.Buffer, parameter EncodableProto) error {
	encoder, err := getEncoder(parameter)
	if err != nil {
		return fmt.Errorf("unable to encode parameter: %v : %w", parameter.ProtoReflect().Descriptor().Name(), err)
	}

	return encoder.Encode(buffer, parameter)
}

func getEncoder(parameter EncodableProto) (Encoder, error) {
	switch parameter.(type) {
	case *solidity.Uint:
		return &staticTypeEncoder{}, nil
	case *solidity.Int:
		return &staticTypeEncoder{}, nil
	case *solidity.Address:
		return &staticTypeEncoder{}, nil
	case *solidity.Bool:
		return &staticTypeEncoder{}, nil
	case *solidity.Fixed:
		return &staticTypeEncoder{}, nil
	case *solidity.Ufixed:
		return &staticTypeEncoder{}, nil
	case *solidity.Bytes:
		return &staticTypeEncoder{}, nil
	case *solidity.Function:
		return &staticTypeEncoder{}, nil
	default:
		return nil, fmt.Errorf("no encoder found")
	}
}

func (s *staticTypeEncoder) Encode(buffer bytes.Buffer, parameter EncodableProto) error {
	switch n := parameter.(type) {
	case *solidity.Uint:
		return encodeUint(buffer, n)
	case *solidity.Int:
		return encodeInt(buffer, n)
	case *solidity.Address:
		return encodeAddress(buffer, n)
	case *solidity.Bool:
		return encodeBool(buffer, n)
	case *solidity.Fixed:
		return encodeFixed(buffer, n)
	case *solidity.Ufixed:
		return encodeUfixed(buffer, n)
	case *solidity.Bytes:
		return encodeBytes(buffer, n)
	default:
		return fmt.Errorf("invalid static type")
	}

	return nil
}

func encodeUint(buffer bytes.Buffer, uint *solidity.Uint) error {
	if len(uint.UintBytes) > 32 {
		return fmt.Errorf("uint type can only have 32 bytes max")
	}

	buffer.Write(leftPadByteSliceZero32(uint.UintBytes))
	return nil
}

// encodeInt assumes the bytes are in twos-complement big-endian
func encodeInt(buffer bytes.Buffer, intType *solidity.Int) error {
	if len(intType.IntBytes) > 32 {
		return fmt.Errorf("int type can only have 32 bytes max")
	}

	numBytesTotalForInt := int(intType.Type) + 1
	intAsBigEndianBytes := leftPadByteSliceToSizeWithZeros(intType.IntBytes, numBytesTotalForInt)

	// check if the number is negative (msb is set)
	isNegative := len(intAsBigEndianBytes) > 0 && (intAsBigEndianBytes[0]&1<<7) != 0
	if isNegative {
		buffer.Write(leftPadBytSliceFF32(intAsBigEndianBytes))
	} else {
		buffer.Write(leftPadByteSliceZero32(intAsBigEndianBytes))
	}

	return nil
}

func encodeAddress(buffer bytes.Buffer, address *solidity.Address) error {
	if len(address.Address) != 20 {
		return fmt.Errorf("address type must be exactly 20 bytes")
	}

	buffer.Write(leftPadByteSliceZero32(address.Address))
	return nil
}

func encodeBool(buffer bytes.Buffer, bool *solidity.Bool) error {
	if bool.Boolean {
		buffer.Write(leftPadByteSliceZero32([]byte{0x01}))
	} else {
		buffer.Write(leftPadByteSliceZero32([]byte{0x00}))
	}

	return nil
}

func encodeFixed(buffer bytes.Buffer, fixed *solidity.Fixed) error {
	if fixed.M < 8 || fixed.M > 256 {
		return fmt.Errorf("fixed type M must be in the range 8 <= M <= 256")
	}

	if fixed.M%8 != 0 {
		return fmt.Errorf("fixed type M must be a multiple of 8 (i.e M mod 8 == 0)")
	}

	if fixed.N <= 0 || fixed.N > 80 {
		return fmt.Errorf("fixed type N must be in the range 0 < N <= 80")
	}

	numBytesTotalForInt := int(fixed.M / 8)
	intAsBigEndianBytes := leftPadByteSliceToSizeWithZeros(fixed.IntBytes, numBytesTotalForInt)
	// check if the number is negative (msb is set)
	isNegative := len(intAsBigEndianBytes) > 0 && (intAsBigEndianBytes[0]&1<<7) != 0

	xBytes := new(big.Int).SetBytes(intAsBigEndianBytes)
	nPow10 := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(fixed.N)), nil)
	value := new(big.Int).Mul(xBytes, nPow10)

	asInt256Bytes := value.Bytes()
	if isNegative {
		// if the number was originally negative we need to make sure to pad
		// with 0xff when we cast up to an int256 (otherwise casting up will
		// left pad with 0's which resets the msb and makes the number positive)
		asInt256Bytes = leftPadBytSliceFF32(asInt256Bytes)
	}

	err := encodeInt(buffer, &solidity.Int{
		Type:     solidity.IntType_INT256,
		IntBytes: asInt256Bytes,
	})
	if err != nil {
		return fmt.Errorf("error encoding fixed type as int256: %w", err)
	}

	return nil
}

func encodeUfixed(buffer bytes.Buffer, ufixed *solidity.Ufixed) error {
	if ufixed.M < 8 || ufixed.M > 256 {
		return fmt.Errorf("ufixed type M must be in the range 8 <= M <= 256")
	}

	if ufixed.M%8 != 0 {
		return fmt.Errorf("ufixed type M must be a multiple of 8 (i.e M mod 8 == 0)")
	}

	if ufixed.N <= 0 || ufixed.N > 80 {
		return fmt.Errorf("ufixed type N must be in the range 0 < N <= 80")
	}

	numBytesTotalForUInt := int(ufixed.M / 8)
	uintAsBigEndianBytes := leftPadByteSliceToSizeWithZeros(ufixed.IntBytes, numBytesTotalForUInt)

	xBytes := new(big.Int).SetBytes(uintAsBigEndianBytes)
	nPow10 := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(ufixed.N)), nil)
	value := new(big.Int).Mul(xBytes, nPow10)

	err := encodeUint(buffer, &solidity.Uint{
		Type:      solidity.UintType_UINT256,
		UintBytes: value.Bytes(),
	})
	if err != nil {
		return fmt.Errorf("error encoding ufixed type as uint256: %w", err)
	}
	return nil
}

func encodeBytes(buffer bytes.Buffer, bytes *solidity.Bytes) error {
	if len(bytes.Bytes) > 32 {
		return fmt.Errorf("bytes type must be less than or equal to 32 bytes")
	}

	buffer.Write(leftPadByteSliceZero32(bytes.Bytes))
	return nil
}

func encodeFunction(buffer bytes.Buffer, function *solidity.Function) error {
	if len(function.Address.Address) != 20 {
		return fmt.Errorf("function type address field must be exactly 20 bytes")
	}

	if len(function.FunctionSelector) != 4 {
		return fmt.Errorf("function type function_selector field must be exactly 4 bytes")
	}

	err := encodeBytes(buffer, &solidity.Bytes{
		Bytes: append(function.Address.Address, function.FunctionSelector...),
	})
	if err != nil {
		return fmt.Errorf("error encoding function type as bytes24: %w", err)
	}

	return nil
}

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
