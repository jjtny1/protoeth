package proto_generator

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	arrayRegex = regexp.MustCompile(`([a-z0-9]*)\[([0-9]*)\]`)

	solidityTypeProtoTypesMap = map[string]string{
		"int8":   "Int8",
		"int16":  "Int16",
		"int24":  "Int24",
		"int32":  "Int32",
		"int40":  "Int40",
		"int48":  "Int48",
		"int56":  "Int56",
		"int64":  "Int64",
		"int72":  "Int72",
		"int80":  "Int80",
		"int88":  "Int88",
		"int96":  "Int96",
		"int104": "Int104",
		"int112": "Int112",
		"int120": "Int120",
		"int128": "Int128",
		"int136": "Int136",
		"int144": "Int144",
		"int152": "Int152",
		"int160": "Int160",
		"int168": "Int168",
		"int176": "Int176",
		"int184": "Int184",
		"int192": "Int192",
		"int200": "Int200",
		"int208": "Int208",
		"int216": "Int216",
		"int224": "Int224",
		"int232": "Int232",
		"int240": "Int240",
		"int248": "Int248",
		"int256": "Int256",

		"uint8":   "Uint8",
		"uint16":  "Uint16",
		"uint24":  "Uint24",
		"uint32":  "Uint32",
		"uint40":  "Uint40",
		"uint48":  "Uint48",
		"uint56":  "Uint56",
		"uint64":  "Uint64",
		"uint72":  "Uint72",
		"uint80":  "Uint80",
		"uint88":  "Uint88",
		"uint96":  "Uint96",
		"uint104": "Uint104",
		"uint112": "Uint112",
		"uint120": "Uint120",
		"uint128": "Uint128",
		"uint136": "Uint136",
		"uint144": "Uint144",
		"uint152": "Uint152",
		"uint160": "Uint160",
		"uint168": "Uint168",
		"uint176": "Uint176",
		"uint184": "Uint184",
		"uint192": "Uint192",
		"uint200": "Uint200",
		"uint208": "Uint208",
		"uint216": "Uint216",
		"uint224": "Uint224",
		"uint232": "Uint232",
		"uint240": "Uint240",
		"uint248": "Uint248",
		"uint256": "Uint256",

		"bytes1":  "Bytes1",
		"bytes2":  "Bytes2",
		"bytes3":  "Bytes3",
		"bytes4":  "Bytes4",
		"bytes5":  "Bytes5",
		"bytes6":  "Bytes6",
		"bytes7":  "Bytes7",
		"bytes8":  "Bytes8",
		"bytes9":  "Bytes9",
		"bytes10": "Bytes10",
		"bytes11": "Bytes11",
		"bytes12": "Bytes12",
		"bytes13": "Bytes13",
		"bytes14": "Bytes14",
		"bytes15": "Bytes15",
		"bytes16": "Bytes16",
		"bytes17": "Bytes17",
		"bytes18": "Bytes18",
		"bytes19": "Bytes19",
		"bytes20": "Bytes20",
		"bytes21": "Bytes21",
		"bytes22": "Bytes22",
		"bytes23": "Bytes23",
		"bytes24": "Bytes24",
		"bytes25": "Bytes25",
		"bytes26": "Bytes26",
		"bytes27": "Bytes27",
		"bytes28": "Bytes28",
		"bytes29": "Bytes29",
		"bytes30": "Bytes30",
		"bytes31": "Bytes31",
		"bytes32": "Bytes32",

		"address":  "Address",
		"bool":     "Bool",
		"function": "Function",
	}
)

func solidityTypeToProtoType(solidityType string) (string, error) {
	var err error

	solidityTypeOnly := solidityType
	if isArray(solidityType) {
		if solidityTypeOnly, err = parseArray(solidityType); err != nil {
			return "", fmt.Errorf("error getting prototype from solidity type %v: %w", solidityType, err)
		}
	}

	protoType, ok := solidityTypeProtoTypesMap[solidityTypeOnly]
	if !ok {
		return "", fmt.Errorf("unsupported solidity type: %v", solidityType)
	}

	return protoType, nil
}

func isFixed(solidityType string) bool {
	return strings.Contains(solidityType, "fixed")
}

func isUFixed(solidityType string) bool {
	return strings.Contains(solidityType, "ufixed")
}

func isArray(solidityType string) bool {
	return arrayRegex.MatchString(solidityType)
}

func isDynamicArray(solidityType string) bool {
	matches := arrayRegex.FindStringSubmatch(solidityType)
	if matches == nil || len(matches) != 3 {
		return false
	}

	// matches[2] should be the size of the array.
	// e.g uint256[73]
	// for dynamic arrays the size is omitted so the length
	// should be 0
	return len(matches[2]) == 0
}

func isStaticArray(solidityType string) bool {
	matches := arrayRegex.FindStringSubmatch(solidityType)
	if matches == nil || len(matches) != 3 {
		return false
	}

	return len(matches[2]) != 0
}

func parseArray(solidityType string) (string, error) {
	if !isArray(solidityType) {
		return "", fmt.Errorf("parseArray called for non-array type: %v", solidityType)
	}

	matches := arrayRegex.FindStringSubmatch(solidityType)
	// note: no size checking for matches since isStaticArray should've
	// handled that

	return matches[1], nil
}
