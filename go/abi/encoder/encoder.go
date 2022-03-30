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
		Encode(*bytes.Buffer, EncodableProto) error
	}

	staticTypeEncoder struct{}
)

func Encode(parameters []EncodableProto) ([]byte, error) {
	var buffer bytes.Buffer

	for _, parameter := range parameters {
		if err := EncodeParameter(&buffer, parameter); err != nil {
			return nil, err
		}
	}

	return buffer.Bytes(), nil
}

// func encode(buffer *bytes.Buffer, parameter EncodableProto) error {
// 	encoder, err := getEncoder(parameter)
// 	if err != nil {
// 		return fmt.Errorf("unable to encode parameter: %v : %w", parameter.ProtoReflect().Descriptor().Name(), err)
// 	}

// 	return encoder.Encode(buffer, parameter)
// }

// func getEncoder(parameter EncodableProto) (Encoder, error) {
// 	switch parameter.(type) {
// 	// Uint cases
// 	case *solidity.Uint8:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Uint16:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Uint24:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Uint32:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Uint40:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Uint48:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Uint56:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Uint64:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Uint72:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Uint80:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Uint88:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Uint96:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Uint104:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Uint112:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Uint120:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Uint128:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Uint136:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Uint144:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Uint152:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Uint160:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Uint168:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Uint176:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Uint184:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Uint192:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Uint200:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Uint208:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Uint216:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Uint224:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Uint232:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Uint240:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Uint248:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Uint256:
// 		return &staticTypeEncoder{}, nil

// 	// Int cases
// 	case *solidity.Int8:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Int16:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Int24:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Int32:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Int40:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Int48:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Int56:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Int64:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Int72:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Int80:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Int88:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Int96:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Int104:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Int112:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Int120:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Int128:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Int136:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Int144:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Int152:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Int160:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Int168:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Int176:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Int184:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Int192:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Int200:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Int208:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Int216:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Int224:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Int232:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Int240:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Int248:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Int256:
// 		return &staticTypeEncoder{}, nil

// 	//Bytes cases
// 	case *solidity.Bytes1:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bytes2:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bytes3:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bytes4:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bytes5:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bytes6:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bytes7:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bytes8:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bytes9:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bytes10:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bytes11:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bytes12:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bytes13:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bytes14:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bytes15:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bytes16:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bytes17:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bytes18:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bytes19:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bytes20:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bytes21:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bytes22:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bytes23:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bytes24:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bytes25:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bytes26:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bytes27:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bytes28:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bytes29:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bytes30:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bytes31:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bytes32:
// 		return &staticTypeEncoder{}, nil

// 	case *solidity.Address:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Bool:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Fixed:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Ufixed:
// 		return &staticTypeEncoder{}, nil
// 	case *solidity.Function:
// 		return &staticTypeEncoder{}, nil
// 	default:
// 		return nil, fmt.Errorf("no encoder found")
// 	}
// }

func EncodeParameter(buffer *bytes.Buffer, parameter EncodableProto) error {
	switch n := parameter.(type) {
	case *solidity.Uint8:
		return encodeUint(buffer, new(big.Int).SetUint64(n.Uint).Bytes(), 8)
	case *solidity.Uint16:
		return encodeUint(buffer, new(big.Int).SetUint64(n.Uint).Bytes(), 16)
	case *solidity.Uint24:
		return encodeUint(buffer, new(big.Int).SetUint64(n.Uint).Bytes(), 24)
	case *solidity.Uint32:
		return encodeUint(buffer, new(big.Int).SetUint64(n.Uint).Bytes(), 32)
	case *solidity.Uint40:
		return encodeUint(buffer, new(big.Int).SetUint64(n.Uint).Bytes(), 40)
	case *solidity.Uint48:
		return encodeUint(buffer, new(big.Int).SetUint64(n.Uint).Bytes(), 48)
	case *solidity.Uint56:
		return encodeUint(buffer, new(big.Int).SetUint64(n.Uint).Bytes(), 56)
	case *solidity.Uint64:
		return encodeUint(buffer, new(big.Int).SetUint64(n.Uint).Bytes(), 64)
	case *solidity.Uint72:
		return encodeUint(buffer, n.UintBytes, 72)
	case *solidity.Uint80:
		return encodeUint(buffer, n.UintBytes, 80)
	case *solidity.Uint88:
		return encodeUint(buffer, n.UintBytes, 88)
	case *solidity.Uint96:
		return encodeUint(buffer, n.UintBytes, 96)
	case *solidity.Uint104:
		return encodeUint(buffer, n.UintBytes, 104)
	case *solidity.Uint112:
		return encodeUint(buffer, n.UintBytes, 112)
	case *solidity.Uint120:
		return encodeUint(buffer, n.UintBytes, 120)
	case *solidity.Uint128:
		return encodeUint(buffer, n.UintBytes, 128)
	case *solidity.Uint136:
		return encodeUint(buffer, n.UintBytes, 136)
	case *solidity.Uint144:
		return encodeUint(buffer, n.UintBytes, 144)
	case *solidity.Uint152:
		return encodeUint(buffer, n.UintBytes, 152)
	case *solidity.Uint160:
		return encodeUint(buffer, n.UintBytes, 160)
	case *solidity.Uint168:
		return encodeUint(buffer, n.UintBytes, 168)
	case *solidity.Uint176:
		return encodeUint(buffer, n.UintBytes, 176)
	case *solidity.Uint184:
		return encodeUint(buffer, n.UintBytes, 184)
	case *solidity.Uint192:
		return encodeUint(buffer, n.UintBytes, 192)
	case *solidity.Uint200:
		return encodeUint(buffer, n.UintBytes, 200)
	case *solidity.Uint208:
		return encodeUint(buffer, n.UintBytes, 208)
	case *solidity.Uint216:
		return encodeUint(buffer, n.UintBytes, 216)
	case *solidity.Uint224:
		return encodeUint(buffer, n.UintBytes, 224)
	case *solidity.Uint232:
		return encodeUint(buffer, n.UintBytes, 232)
	case *solidity.Uint240:
		return encodeUint(buffer, n.UintBytes, 240)
	case *solidity.Uint248:
		return encodeUint(buffer, n.UintBytes, 248)
	case *solidity.Uint256:
		return encodeUint(buffer, n.UintBytes, 256)

	case *solidity.Int8:
		return encodeInt(buffer, big.NewInt(n.Int).Bytes(), 8)
	case *solidity.Int16:
		return encodeInt(buffer, big.NewInt(n.Int).Bytes(), 16)
	case *solidity.Int24:
		return encodeInt(buffer, big.NewInt(n.Int).Bytes(), 24)
	case *solidity.Int32:
		return encodeInt(buffer, big.NewInt(n.Int).Bytes(), 32)
	case *solidity.Int40:
		return encodeInt(buffer, big.NewInt(n.Int).Bytes(), 40)
	case *solidity.Int48:
		return encodeInt(buffer, big.NewInt(n.Int).Bytes(), 48)
	case *solidity.Int56:
		return encodeInt(buffer, big.NewInt(n.Int).Bytes(), 56)
	case *solidity.Int64:
		return encodeInt(buffer, big.NewInt(n.Int).Bytes(), 64)
	case *solidity.Int72:
		return encodeInt(buffer, n.IntBytes, 72)
	case *solidity.Int80:
		return encodeInt(buffer, n.IntBytes, 80)
	case *solidity.Int88:
		return encodeInt(buffer, n.IntBytes, 88)
	case *solidity.Int96:
		return encodeInt(buffer, n.IntBytes, 96)
	case *solidity.Int104:
		return encodeInt(buffer, n.IntBytes, 104)
	case *solidity.Int112:
		return encodeInt(buffer, n.IntBytes, 112)
	case *solidity.Int120:
		return encodeInt(buffer, n.IntBytes, 120)
	case *solidity.Int128:
		return encodeInt(buffer, n.IntBytes, 128)
	case *solidity.Int136:
		return encodeInt(buffer, n.IntBytes, 136)
	case *solidity.Int144:
		return encodeInt(buffer, n.IntBytes, 144)
	case *solidity.Int152:
		return encodeInt(buffer, n.IntBytes, 152)
	case *solidity.Int160:
		return encodeInt(buffer, n.IntBytes, 160)
	case *solidity.Int168:
		return encodeInt(buffer, n.IntBytes, 168)
	case *solidity.Int176:
		return encodeInt(buffer, n.IntBytes, 176)
	case *solidity.Int184:
		return encodeInt(buffer, n.IntBytes, 184)
	case *solidity.Int192:
		return encodeInt(buffer, n.IntBytes, 192)
	case *solidity.Int200:
		return encodeInt(buffer, n.IntBytes, 200)
	case *solidity.Int208:
		return encodeInt(buffer, n.IntBytes, 208)
	case *solidity.Int216:
		return encodeInt(buffer, n.IntBytes, 216)
	case *solidity.Int224:
		return encodeInt(buffer, n.IntBytes, 224)
	case *solidity.Int232:
		return encodeInt(buffer, n.IntBytes, 232)
	case *solidity.Int240:
		return encodeInt(buffer, n.IntBytes, 240)
	case *solidity.Int248:
		return encodeInt(buffer, n.IntBytes, 248)
	case *solidity.Int256:
		return encodeInt(buffer, n.IntBytes, 256)

	case *solidity.Bytes1:
		return encodeStaticBytes(buffer, n.Bytes, 1)
	case *solidity.Bytes2:
		return encodeStaticBytes(buffer, n.Bytes, 2)
	case *solidity.Bytes3:
		return encodeStaticBytes(buffer, n.Bytes, 3)
	case *solidity.Bytes4:
		return encodeStaticBytes(buffer, n.Bytes, 4)
	case *solidity.Bytes5:
		return encodeStaticBytes(buffer, n.Bytes, 5)
	case *solidity.Bytes6:
		return encodeStaticBytes(buffer, n.Bytes, 6)
	case *solidity.Bytes7:
		return encodeStaticBytes(buffer, n.Bytes, 7)
	case *solidity.Bytes8:
		return encodeStaticBytes(buffer, n.Bytes, 8)
	case *solidity.Bytes9:
		return encodeStaticBytes(buffer, n.Bytes, 9)
	case *solidity.Bytes10:
		return encodeStaticBytes(buffer, n.Bytes, 10)
	case *solidity.Bytes11:
		return encodeStaticBytes(buffer, n.Bytes, 11)
	case *solidity.Bytes12:
		return encodeStaticBytes(buffer, n.Bytes, 12)
	case *solidity.Bytes13:
		return encodeStaticBytes(buffer, n.Bytes, 13)
	case *solidity.Bytes14:
		return encodeStaticBytes(buffer, n.Bytes, 14)
	case *solidity.Bytes15:
		return encodeStaticBytes(buffer, n.Bytes, 15)
	case *solidity.Bytes16:
		return encodeStaticBytes(buffer, n.Bytes, 16)
	case *solidity.Bytes17:
		return encodeStaticBytes(buffer, n.Bytes, 17)
	case *solidity.Bytes18:
		return encodeStaticBytes(buffer, n.Bytes, 18)
	case *solidity.Bytes19:
		return encodeStaticBytes(buffer, n.Bytes, 19)
	case *solidity.Bytes20:
		return encodeStaticBytes(buffer, n.Bytes, 20)
	case *solidity.Bytes21:
		return encodeStaticBytes(buffer, n.Bytes, 21)
	case *solidity.Bytes22:
		return encodeStaticBytes(buffer, n.Bytes, 22)
	case *solidity.Bytes23:
		return encodeStaticBytes(buffer, n.Bytes, 23)
	case *solidity.Bytes24:
		return encodeStaticBytes(buffer, n.Bytes, 24)
	case *solidity.Bytes25:
		return encodeStaticBytes(buffer, n.Bytes, 25)
	case *solidity.Bytes26:
		return encodeStaticBytes(buffer, n.Bytes, 26)
	case *solidity.Bytes27:
		return encodeStaticBytes(buffer, n.Bytes, 27)
	case *solidity.Bytes28:
		return encodeStaticBytes(buffer, n.Bytes, 28)
	case *solidity.Bytes29:
		return encodeStaticBytes(buffer, n.Bytes, 29)
	case *solidity.Bytes30:
		return encodeStaticBytes(buffer, n.Bytes, 30)
	case *solidity.Bytes31:
		return encodeStaticBytes(buffer, n.Bytes, 31)
	case *solidity.Bytes32:
		return encodeStaticBytes(buffer, n.Bytes, 32)

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
		return fmt.Errorf("invalid static type: %v", parameter.ProtoReflect().Descriptor().Name())
	}

	return nil
}

func encodeUint(buffer *bytes.Buffer, uintBytes []byte, totalBits int) error {
	numBytesTotalForUint := totalBits / 8

	if len(uintBytes) > numBytesTotalForUint {
		return fmt.Errorf("uint exceeds total number of bits: %v", totalBits)
	}
	if len(uintBytes) > 32 {
		return fmt.Errorf("uint type can only have 32 bytes max")
	}

	buffer.Write(leftPadByteSliceZero32(uintBytes))
	return nil
}

// encodeInt assumes the bytes are in two's-complement bigendian
func encodeInt(buffer *bytes.Buffer, intBytes []byte, totalBits int) error {
	numBytesTotalForInt := totalBits / 8

	if len(intBytes) > numBytesTotalForInt {
		return fmt.Errorf("int exceeds total number of bits: %v", totalBits)
	}
	if len(intBytes) > 32 {
		return fmt.Errorf("int type can only have 32 bytes max")
	}

	intAsBigEndianBytes := leftPadByteSliceToSizeWithZeros(intBytes, numBytesTotalForInt)

	// check if the number is negative (msb is set)
	isNegative := len(intAsBigEndianBytes) > 0 && (intAsBigEndianBytes[0]&1<<7) != 0
	if isNegative {
		buffer.Write(leftPadBytSliceFF32(intAsBigEndianBytes))
	} else {
		buffer.Write(leftPadByteSliceZero32(intAsBigEndianBytes))
	}

	return nil
}

func encodeAddress(buffer *bytes.Buffer, address *solidity.Address) error {
	if len(address.Address) != 20 {
		return fmt.Errorf("address type must be exactly 20 bytes")
	}

	buffer.Write(leftPadByteSliceZero32(address.Address))
	return nil
}

func encodeBool(buffer *bytes.Buffer, bool *solidity.Bool) error {
	if bool.Boolean {
		buffer.Write(leftPadByteSliceZero32([]byte{0x01}))
	} else {
		buffer.Write(leftPadByteSliceZero32([]byte{0x00}))
	}

	return nil
}

func encodeFixed(buffer *bytes.Buffer, fixed *solidity.Fixed) error {
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

	err := encodeInt(buffer, asInt256Bytes, 256)
	if err != nil {
		return fmt.Errorf("error encoding fixed type as int256: %w", err)
	}

	return nil
}

func encodeUfixed(buffer *bytes.Buffer, ufixed *solidity.Ufixed) error {
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

	err := encodeUint(buffer, value.Bytes(), 256)
	if err != nil {
		return fmt.Errorf("error encoding ufixed type as uint256: %w", err)
	}
	return nil
}

func encodeStaticBytes(buffer *bytes.Buffer, bytes []byte, totalBytes int) error {
	if len(bytes) > totalBytes {
		return fmt.Errorf("bytes type exceeds totalBytes: %v", totalBytes)
	}
	if len(bytes) > 32 {
		return fmt.Errorf("bytes type must be less than or equal to 32 bytes")
	}

	buffer.Write(rightPadBytesSliceZero32(bytes))
	return nil
}

func encodeFunction(buffer *bytes.Buffer, function *solidity.Function) error {
	if len(function.Address.Address) != 20 {
		return fmt.Errorf("function type address field must be exactly 20 bytes")
	}

	if len(function.FunctionSelector) != 4 {
		return fmt.Errorf("function type function_selector field must be exactly 4 bytes")
	}

	err := encodeStaticBytes(buffer, append(function.Address.Address, function.FunctionSelector...), 24)
	if err != nil {
		return fmt.Errorf("error encoding function type as bytes24: %w", err)
	}

	return nil
}

func encodeBytes(buffer *bytes.Buffer, bytes *solidity.Bytes) error {
	length := len(bytes.Bytes)

	// treat length as a uint256 number
	lengthUint256 := &solidity.Uint256{
		UintBytes: big.NewInt(int64(length)).Bytes(),
	}

	if err := EncodeParameter(buffer, lengthUint256); err != nil {
		return fmt.Errorf("error encoding length of dynamic bytes: %w", err)
	}

	buffer.Write(padRightMultipleOf32(bytes.Bytes))
	return nil
}
