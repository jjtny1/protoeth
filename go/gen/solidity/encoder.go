package solidity

import (
	"bytes"
	"fmt"
	"math/big"
	"reflect"
)

type (
	EncodableProto interface {
		isType_SolidityType
	}

	Encoder interface {
		Encode(*bytes.Buffer, EncodableProto) error
	}

	staticTypeEncoder struct{}
)

func Encode(parameters []EncodableProto) ([]byte, error) {
	var buffer bytes.Buffer

	typeList := make([]*Type, 0, len(parameters))
	for _, parameter := range parameters {
		typeList = append(typeList, &Type{
			SolidityType: parameter,
		})
	}

	parametersTuple := &Type_Tuple{
		Tuple: &Tuple{
			TupleList: typeList,
		},
	}

	if err := EncodeSolidityType(&buffer, parametersTuple); err != nil {
		return nil, fmt.Errorf("error encoding parameters: %w", err)
	}

	return buffer.Bytes(), nil
}

func EncodeSolidityType(buffer *bytes.Buffer, solidityType EncodableProto) error {
	switch n := solidityType.(type) {
	case *Type_Tuple:
		return encodeTuple(buffer, n.Tuple)
	case *Type_Uint8:
		return encodeUint(buffer, new(big.Int).SetUint64(n.Uint8.Uint).Bytes(), 8)
	case *Type_Uint16:
		return encodeUint(buffer, new(big.Int).SetUint64(n.Uint16.Uint).Bytes(), 16)
	case *Type_Uint24:
		return encodeUint(buffer, new(big.Int).SetUint64(n.Uint24.Uint).Bytes(), 24)
	case *Type_Uint32:
		return encodeUint(buffer, new(big.Int).SetUint64(n.Uint32.Uint).Bytes(), 32)
	case *Type_Uint40:
		return encodeUint(buffer, new(big.Int).SetUint64(n.Uint40.Uint).Bytes(), 40)
	case *Type_Uint48:
		return encodeUint(buffer, new(big.Int).SetUint64(n.Uint48.Uint).Bytes(), 48)
	case *Type_Uint56:
		return encodeUint(buffer, new(big.Int).SetUint64(n.Uint56.Uint).Bytes(), 56)
	case *Type_Uint64:
		return encodeUint(buffer, new(big.Int).SetUint64(n.Uint64.Uint).Bytes(), 64)
	case *Type_Uint72:
		return encodeUint(buffer, n.Uint72.UintBytes, 72)
	case *Type_Uint80:
		return encodeUint(buffer, n.Uint80.UintBytes, 80)
	case *Type_Uint88:
		return encodeUint(buffer, n.Uint88.UintBytes, 88)
	case *Type_Uint96:
		return encodeUint(buffer, n.Uint96.UintBytes, 96)
	case *Type_Uint104:
		return encodeUint(buffer, n.Uint104.UintBytes, 104)
	case *Type_Uint112:
		return encodeUint(buffer, n.Uint112.UintBytes, 112)
	case *Type_Uint120:
		return encodeUint(buffer, n.Uint120.UintBytes, 120)
	case *Type_Uint128:
		return encodeUint(buffer, n.Uint128.UintBytes, 128)
	case *Type_Uint136:
		return encodeUint(buffer, n.Uint136.UintBytes, 136)
	case *Type_Uint144:
		return encodeUint(buffer, n.Uint144.UintBytes, 144)
	case *Type_Uint152:
		return encodeUint(buffer, n.Uint152.UintBytes, 152)
	case *Type_Uint160:
		return encodeUint(buffer, n.Uint160.UintBytes, 160)
	case *Type_Uint168:
		return encodeUint(buffer, n.Uint168.UintBytes, 168)
	case *Type_Uint176:
		return encodeUint(buffer, n.Uint176.UintBytes, 176)
	case *Type_Uint184:
		return encodeUint(buffer, n.Uint184.UintBytes, 184)
	case *Type_Uint192:
		return encodeUint(buffer, n.Uint192.UintBytes, 192)
	case *Type_Uint200:
		return encodeUint(buffer, n.Uint200.UintBytes, 200)
	case *Type_Uint208:
		return encodeUint(buffer, n.Uint208.UintBytes, 208)
	case *Type_Uint216:
		return encodeUint(buffer, n.Uint216.UintBytes, 216)
	case *Type_Uint224:
		return encodeUint(buffer, n.Uint224.UintBytes, 224)
	case *Type_Uint232:
		return encodeUint(buffer, n.Uint232.UintBytes, 232)
	case *Type_Uint240:
		return encodeUint(buffer, n.Uint240.UintBytes, 240)
	case *Type_Uint248:
		return encodeUint(buffer, n.Uint248.UintBytes, 248)
	case *Type_Uint256:
		return encodeUint(buffer, n.Uint256.UintBytes, 256)

	case *Type_Int8:
		return encodeInt(buffer, big.NewInt(n.Int8.Int).Bytes(), 8)
	case *Type_Int16:
		return encodeInt(buffer, big.NewInt(n.Int16.Int).Bytes(), 16)
	case *Type_Int24:
		return encodeInt(buffer, big.NewInt(n.Int24.Int).Bytes(), 24)
	case *Type_Int32:
		return encodeInt(buffer, big.NewInt(n.Int32.Int).Bytes(), 32)
	case *Type_Int40:
		return encodeInt(buffer, big.NewInt(n.Int40.Int).Bytes(), 40)
	case *Type_Int48:
		return encodeInt(buffer, big.NewInt(n.Int48.Int).Bytes(), 48)
	case *Type_Int56:
		return encodeInt(buffer, big.NewInt(n.Int56.Int).Bytes(), 56)
	case *Type_Int64:
		return encodeInt(buffer, big.NewInt(n.Int64.Int).Bytes(), 64)
	case *Type_Int72:
		return encodeInt(buffer, n.Int72.IntBytes, 72)
	case *Type_Int80:
		return encodeInt(buffer, n.Int80.IntBytes, 80)
	case *Type_Int88:
		return encodeInt(buffer, n.Int88.IntBytes, 88)
	case *Type_Int96:
		return encodeInt(buffer, n.Int96.IntBytes, 96)
	case *Type_Int104:
		return encodeInt(buffer, n.Int104.IntBytes, 104)
	case *Type_Int112:
		return encodeInt(buffer, n.Int112.IntBytes, 112)
	case *Type_Int120:
		return encodeInt(buffer, n.Int120.IntBytes, 120)
	case *Type_Int128:
		return encodeInt(buffer, n.Int128.IntBytes, 128)
	case *Type_Int136:
		return encodeInt(buffer, n.Int136.IntBytes, 136)
	case *Type_Int144:
		return encodeInt(buffer, n.Int144.IntBytes, 144)
	case *Type_Int152:
		return encodeInt(buffer, n.Int152.IntBytes, 152)
	case *Type_Int160:
		return encodeInt(buffer, n.Int160.IntBytes, 160)
	case *Type_Int168:
		return encodeInt(buffer, n.Int168.IntBytes, 168)
	case *Type_Int176:
		return encodeInt(buffer, n.Int176.IntBytes, 176)
	case *Type_Int184:
		return encodeInt(buffer, n.Int184.IntBytes, 184)
	case *Type_Int192:
		return encodeInt(buffer, n.Int192.IntBytes, 192)
	case *Type_Int200:
		return encodeInt(buffer, n.Int200.IntBytes, 200)
	case *Type_Int208:
		return encodeInt(buffer, n.Int208.IntBytes, 208)
	case *Type_Int216:
		return encodeInt(buffer, n.Int216.IntBytes, 216)
	case *Type_Int224:
		return encodeInt(buffer, n.Int224.IntBytes, 224)
	case *Type_Int232:
		return encodeInt(buffer, n.Int232.IntBytes, 232)
	case *Type_Int240:
		return encodeInt(buffer, n.Int240.IntBytes, 240)
	case *Type_Int248:
		return encodeInt(buffer, n.Int248.IntBytes, 248)
	case *Type_Int256:
		return encodeInt(buffer, n.Int256.IntBytes, 256)

	case *Type_Bytes1:
		return encodeStaticBytes(buffer, n.Bytes1.Bytes, 1)
	case *Type_Bytes2:
		return encodeStaticBytes(buffer, n.Bytes2.Bytes, 2)
	case *Type_Bytes3:
		return encodeStaticBytes(buffer, n.Bytes3.Bytes, 3)
	case *Type_Bytes4:
		return encodeStaticBytes(buffer, n.Bytes4.Bytes, 4)
	case *Type_Bytes5:
		return encodeStaticBytes(buffer, n.Bytes5.Bytes, 5)
	case *Type_Bytes6:
		return encodeStaticBytes(buffer, n.Bytes6.Bytes, 6)
	case *Type_Bytes7:
		return encodeStaticBytes(buffer, n.Bytes7.Bytes, 7)
	case *Type_Bytes8:
		return encodeStaticBytes(buffer, n.Bytes8.Bytes, 8)
	case *Type_Bytes9:
		return encodeStaticBytes(buffer, n.Bytes9.Bytes, 9)
	case *Type_Bytes10:
		return encodeStaticBytes(buffer, n.Bytes10.Bytes, 10)
	case *Type_Bytes11:
		return encodeStaticBytes(buffer, n.Bytes11.Bytes, 11)
	case *Type_Bytes12:
		return encodeStaticBytes(buffer, n.Bytes12.Bytes, 12)
	case *Type_Bytes13:
		return encodeStaticBytes(buffer, n.Bytes13.Bytes, 13)
	case *Type_Bytes14:
		return encodeStaticBytes(buffer, n.Bytes14.Bytes, 14)
	case *Type_Bytes15:
		return encodeStaticBytes(buffer, n.Bytes15.Bytes, 15)
	case *Type_Bytes16:
		return encodeStaticBytes(buffer, n.Bytes16.Bytes, 16)
	case *Type_Bytes17:
		return encodeStaticBytes(buffer, n.Bytes17.Bytes, 17)
	case *Type_Bytes18:
		return encodeStaticBytes(buffer, n.Bytes18.Bytes, 18)
	case *Type_Bytes19:
		return encodeStaticBytes(buffer, n.Bytes19.Bytes, 19)
	case *Type_Bytes20:
		return encodeStaticBytes(buffer, n.Bytes20.Bytes, 20)
	case *Type_Bytes21:
		return encodeStaticBytes(buffer, n.Bytes21.Bytes, 21)
	case *Type_Bytes22:
		return encodeStaticBytes(buffer, n.Bytes22.Bytes, 22)
	case *Type_Bytes23:
		return encodeStaticBytes(buffer, n.Bytes23.Bytes, 23)
	case *Type_Bytes24:
		return encodeStaticBytes(buffer, n.Bytes24.Bytes, 24)
	case *Type_Bytes25:
		return encodeStaticBytes(buffer, n.Bytes25.Bytes, 25)
	case *Type_Bytes26:
		return encodeStaticBytes(buffer, n.Bytes26.Bytes, 26)
	case *Type_Bytes27:
		return encodeStaticBytes(buffer, n.Bytes27.Bytes, 27)
	case *Type_Bytes28:
		return encodeStaticBytes(buffer, n.Bytes28.Bytes, 28)
	case *Type_Bytes29:
		return encodeStaticBytes(buffer, n.Bytes29.Bytes, 29)
	case *Type_Bytes30:
		return encodeStaticBytes(buffer, n.Bytes30.Bytes, 30)
	case *Type_Bytes31:
		return encodeStaticBytes(buffer, n.Bytes31.Bytes, 31)
	case *Type_Bytes32:
		return encodeStaticBytes(buffer, n.Bytes32.Bytes, 32)

	case *Type_Address:
		return encodeAddress(buffer, n)
	case *Type_Bool:
		return encodeBool(buffer, n)
	// case *Type_Fixed:
	// 	return encodeFixed(buffer, n)
	// case *Type_Ufixed:
	// 	return encodeUfixed(buffer, n)
	case *Type_Bytes:
		return encodeBytes(buffer, n)
	case *Type_String_:
		return encodeString(buffer, n)
	case *Type_StaticArray:
		return encodeStaticArray(buffer, n)
	case *Type_DynamicArray:
		return encodeDynamicArray(buffer, n)
	default:
		return fmt.Errorf("unsupported solidity type: %v", fmt.Sprintf("%T", n))
	}

	return nil
}

func encodeTuple(buffer *bytes.Buffer, tuple *Tuple) error {
	// tailLengths is a running sum array that contains the length of all
	// dynamic data up to and including some index (i+1) of tuple.TupleList.
	// Said differently, for any element at index i in tuple.TupleList, the
	// total amount of dynamic data BEFORE i is tailLengths[i] and including
	// i is tailLengths[i+1]. Note: The array is sized one bigger than the actual
	// list so look up for i == 0 will not cause out of bounds errors.
	// See docs for full algorithm: https://docs.soliditylang.org/en/v0.5.3/abi-spec.html#formal-specification-of-the-encoding
	tailLengths := make([]int, len(tuple.TupleList)+1)
	tailData := make([]*bytes.Buffer, 0)

	for i, tupleElement := range tuple.TupleList {
		solidityType := tupleElement.SolidityType

		if isStaticType(solidityType) {
			// static types encode in 32 bytes so they are directly encoded
			if err := EncodeSolidityType(buffer, tupleElement.SolidityType); err != nil {
				return fmt.Errorf("error encoding static tuple element %T: %w", solidityType, err)
			}
			// static types have no dynamic data so tail(i) is defined as "" in the solidity docs
			tailLengths[i+1] = tailLengths[i] + 0
		} else {
			// each dynamic element's data starts after every parameters header (which is 32 bytes each) plus
			// the length of all dynamic data of dynamic parameters before it.
			offsetToDynamicData := (len(tuple.TupleList) * 32) + tailLengths[i]

			tailBuffer := new(bytes.Buffer)
			if err := EncodeSolidityType(tailBuffer, solidityType); err != nil {
				return fmt.Errorf("error encoding dynamic tuple element %T: %w", solidityType, err)
			}

			tailData = append(tailData, tailBuffer)
			tailLengths[i+1] = tailLengths[i] + tailBuffer.Len()

			// encode offsetToDynamicData as a uint256 type
			err := EncodeSolidityType(buffer, &Type_Uint256{
				Uint256: &Uint256{
					UintBytes: big.NewInt(int64(offsetToDynamicData)).Bytes(),
				},
			})
			if err != nil {
				return fmt.Errorf("error encoding dynamic data offset for tuple element %T: %w", solidityType, err)
			}
		}
	}

	for _, dynamicData := range tailData {
		buffer.Write(dynamicData.Bytes())
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

func encodeAddress(buffer *bytes.Buffer, address *Type_Address) error {
	if len(address.Address.Address) != 20 {
		return fmt.Errorf("address type must be exactly 20 bytes")
	}

	buffer.Write(leftPadByteSliceZero32(address.Address.Address))
	return nil
}

func encodeBool(buffer *bytes.Buffer, bool *Type_Bool) error {
	if bool.Bool.Boolean {
		buffer.Write(leftPadByteSliceZero32([]byte{0x01}))
	} else {
		buffer.Write(leftPadByteSliceZero32([]byte{0x00}))
	}

	return nil
}

// func encodeFixed(buffer *bytes.Buffer, fixed *Type_Fixed) error {
// 	if fixed.M < 8 || fixed.M > 256 {
// 		return fmt.Errorf("fixed type M must be in the range 8 <= M <= 256")
// 	}

// 	if fixed.M%8 != 0 {
// 		return fmt.Errorf("fixed type M must be a multiple of 8 (i.e M mod 8 == 0)")
// 	}

// 	if fixed.N <= 0 || fixed.N > 80 {
// 		return fmt.Errorf("fixed type N must be in the range 0 < N <= 80")
// 	}

// 	numBytesTotalForInt := int(fixed.M / 8)
// 	intAsBigEndianBytes := leftPadByteSliceToSizeWithZeros(fixed.IntBytes, numBytesTotalForInt)
// 	// check if the number is negative (msb is set)
// 	isNegative := len(intAsBigEndianBytes) > 0 && (intAsBigEndianBytes[0]&1<<7) != 0

// 	xBytes := new(big.Int).SetBytes(intAsBigEndianBytes)
// 	nPow10 := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(fixed.N)), nil)
// 	value := new(big.Int).Mul(xBytes, nPow10)

// 	asInt256Bytes := value.Bytes()
// 	if isNegative {
// 		// if the number was originally negative we need to make sure to pad
// 		// with 0xff when we cast up to an int256 (otherwise casting up will
// 		// left pad with 0's which resets the msb and makes the number positive)
// 		asInt256Bytes = leftPadBytSliceFF32(asInt256Bytes)
// 	}

// 	err := encodeInt(buffer, asInt256Bytes, 256)
// 	if err != nil {
// 		return fmt.Errorf("error encoding fixed type as int256: %w", err)
// 	}

// 	return nil
// }

// func encodeUfixed(buffer *bytes.Buffer, ufixed *Type_Ufixed) error {
// 	if ufixed.M < 8 || ufixed.M > 256 {
// 		return fmt.Errorf("ufixed type M must be in the range 8 <= M <= 256")
// 	}

// 	if ufixed.M%8 != 0 {
// 		return fmt.Errorf("ufixed type M must be a multiple of 8 (i.e M mod 8 == 0)")
// 	}

// 	if ufixed.N <= 0 || ufixed.N > 80 {
// 		return fmt.Errorf("ufixed type N must be in the range 0 < N <= 80")
// 	}

// 	numBytesTotalForUInt := int(ufixed.M / 8)
// 	uintAsBigEndianBytes := leftPadByteSliceToSizeWithZeros(ufixed.IntBytes, numBytesTotalForUInt)

// 	xBytes := new(big.Int).SetBytes(uintAsBigEndianBytes)
// 	nPow10 := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(ufixed.N)), nil)
// 	value := new(big.Int).Mul(xBytes, nPow10)

// 	err := encodeUint(buffer, value.Bytes(), 256)
// 	if err != nil {
// 		return fmt.Errorf("error encoding ufixed type as uint256: %w", err)
// 	}
// 	return nil
// }

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

func encodeFunction(buffer *bytes.Buffer, function *Type_Function) error {
	if len(function.Function.Address.Address) != 20 {
		return fmt.Errorf("function type address field must be exactly 20 bytes")
	}

	if len(function.Function.FunctionSelector) != 4 {
		return fmt.Errorf("function type function_selector field must be exactly 4 bytes")
	}

	err := encodeStaticBytes(buffer, append(function.Function.Address.Address, function.Function.FunctionSelector...), 24)
	if err != nil {
		return fmt.Errorf("error encoding function type as bytes24: %w", err)
	}

	return nil
}

func encodeBytes(buffer *bytes.Buffer, bytes *Type_Bytes) error {
	length := len(bytes.Bytes.Bytes)

	// treat length as a uint256 number
	lengthUint256 := &Type_Uint256{
		Uint256: &Uint256{
			UintBytes: big.NewInt(int64(length)).Bytes(),
		},
	}

	if err := EncodeSolidityType(buffer, lengthUint256); err != nil {
		return fmt.Errorf("error encoding length of dynamic bytes: %w", err)
	}

	buffer.Write(padRightMultipleOf32(bytes.Bytes.Bytes))
	return nil
}

func encodeString(buffer *bytes.Buffer, str *Type_String_) error {
	// utf-8 encode the string. Protobuf string should already be utf-8 encoded!
	return EncodeSolidityType(buffer, &Type_Bytes{
		Bytes: &Bytes{
			Bytes: []byte(str.String_.String_),
		},
	})
}

func encodeStaticArray(buffer *bytes.Buffer, sa *Type_StaticArray) error {
	// validate size and all elements are the same
	if len(sa.StaticArray.TypeList) != int(sa.StaticArray.Size) {
		return fmt.Errorf("size of static array does not match number of elements")
	}

	if sa.StaticArray.Size == 0 {
		return fmt.Errorf("static array cannot have size 0")
	}

	expectedElementType := reflect.TypeOf(sa.StaticArray.TypeList[0])

	for _, element := range sa.StaticArray.TypeList {
		if reflect.TypeOf(element) != expectedElementType {
			return fmt.Errorf("static arrays cannot have elements of different types")
		}
	}

	// A static array's encoding is the same as a tuple's encoding of
	// the elements
	return EncodeSolidityType(buffer, &Type_Tuple{
		Tuple: &Tuple{
			TupleList: sa.StaticArray.TypeList,
		},
	})
}

func encodeDynamicArray(buffer *bytes.Buffer, da *Type_DynamicArray) error {
	length := len(da.DynamicArray.TypeList)

	// treat length as a uint256 number
	lengthUint256 := &Type_Uint256{
		Uint256: &Uint256{
			UintBytes: big.NewInt(int64(length)).Bytes(),
		},
	}

	if err := EncodeSolidityType(buffer, lengthUint256); err != nil {
		return fmt.Errorf("error encoding length of dynamic array: %w", err)
	}

	if length == 0 {
		return nil
	}

	// treat the dynamic array as a static array of size length
	return EncodeSolidityType(buffer, &Type_StaticArray{
		StaticArray: &TypeStaticArray{
			Size:     int32(length),
			TypeList: da.DynamicArray.TypeList,
		},
	})
}

func isStaticType(solidityType EncodableProto) bool {
	return !isDynamicType(solidityType)
}

func isDynamicType(solidityType EncodableProto) bool {
	switch n := solidityType.(type) {
	case *Type_Bytes:
		return true
	case *Type_String_:
		return true
	case *Type_DynamicArray:
		return true
	case *Type_StaticArray:
		return len(n.StaticArray.TypeList) > 0 && isDynamicType(n.StaticArray.TypeList[0].SolidityType)
	}

	return false
}
