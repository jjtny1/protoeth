package proto_generator

import (
	"fmt"
	"strings"
)

func SolidityTypeToProtoType(solidityType string) (string, bool, error) {
	protoType := ""

	if isUint(solidityType) {
		protoType = "Uint"
	} else if isInt(solidityType) {
		protoType = "Int"
	} else if isAddress(solidityType) {
		protoType = "Address"
	} else if isBool(solidityType) {
		protoType = "Bool"
	} else if isFixed(solidityType) {
		protoType = "Fixed"
	} else if isUFixed(solidityType) {
		protoType = "Ufixed"
	} else if isBytes(solidityType) {
		protoType = "Bytes"
	} else if isFunction(solidityType) {
		protoType = "Function"
	} else {
		return "", false, fmt.Errorf("unsupported solidity type: %v", solidityType)
	}

	if isArray(solidityType) {
		return protoType, true, nil
	}

	return protoType, false, nil
}

func isUint(solidityType string) bool {
	return strings.Contains(solidityType, "uint")
}

func isInt(solidityType string) bool {
	return strings.Contains(solidityType, "int")
}

func isAddress(solidityType string) bool {
	return solidityType == "address"
}

func isBool(solidityType string) bool {
	return solidityType == "bool"
}

func isFixed(solidityType string) bool {
	return strings.Contains(solidityType, "fixed")
}

func isUFixed(solidityType string) bool {
	return strings.Contains(solidityType, "ufixed")
}

func isBytes(solidityType string) bool {
	return strings.Contains(solidityType, "bytes")
}

func isFunction(solidityType string) bool {
	return solidityType == "function"
}

func isArray(solidityType string) bool {
	return strings.Contains(solidityType, "[") && strings.Contains(solidityType, "]")
}
