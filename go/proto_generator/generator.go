package proto_generator

import (
	"fmt"
)

type ABI []FunctionDescription

type FunctionDescription struct {
	Type            string           `json:"type"`
	Name            string           `json:"name"`
	Inputs          []FunctionInput  `json:"inputs"`
	Outputs         []FunctionOutput `json:"outputs"`
	StateMutability string           `json:"stateMutability"`
	Payable         bool             `json:"payabale"`
	Constant        bool             `json:"constant"`
}

type FunctionInput struct {
	ParameterName string          `json:"name"`
	ParameterType string          `json:"type"`
	Components    []FunctionInput `json:"components"`
}

type FunctionOutput FunctionInput

type ProtobufService struct {
	Name           string
	InputParameter ProtobufMessage
	ReturnValue    string
}

type ProtobufMessage struct {
	Name   string
	Fields []ProtobufField
}

type ProtobufField struct {
	// a nil value for ArrayInfo indicates the field
	// is not an array
	ArrayInfo *ArrayInfo
	TypeName  string
	FieldName string
}

type ArrayInfo struct {
	// a valued of < 0 indicates that it is a
	// dynamic array
	SolidityArray string
}

const (
	functionParametersName  = "%s_Parameters"
	functionReturnValueName = "EncodedFunctionCall"
	protoTemplate           = `
syntax = "proto3";

package solidity;

option go_package = "go/gen/solidity_contracts";

import "solidity_types.proto";
import "google/protobuf/descriptor.proto";

extend google.protobuf.FieldOptions {
  optional string array = 50001;
}

service ContractService {
%s
}

%s
`
)

func GenerateSolidityContractProtoFileFromABIFile(abi ABI) (string, error) {
	var services []ProtobufService
	for _, fd := range abi {
		// In solidity, an empty type field defaults to function.
		// https://docs.soliditylang.org/en/v0.5.3/abi-spec.html#json
		if fd.Type == "function" || fd.Type == "" {
			service := ProtobufService{
				Name: fd.Name,
				InputParameter: ProtobufMessage{
					Name: fmt.Sprintf(functionParametersName, fd.Name),
				},
				ReturnValue: functionReturnValueName,
			}

			for _, parameter := range fd.Inputs {
				protoFieldType, err := solidityTypeToProtoType(parameter.ParameterType)
				if err != nil {
					return "", err
				}

				var arrayInfo *ArrayInfo
				if isArray(parameter.ParameterType) {
					arrayInfo = &ArrayInfo{
						SolidityArray: parameter.ParameterType,
					}
				}

				service.InputParameter.Fields = append(service.InputParameter.Fields, ProtobufField{
					ArrayInfo: arrayInfo,
					TypeName:  protoFieldType,
					FieldName: parameter.ParameterName,
				})
			}

			services = append(services, service)
		}
	}

	return fmt.Sprintf(protoTemplate, getRPCDefinitions(services), getMessageDefinitions(services)), nil
}

func getRPCDefinitions(services []ProtobufService) string {
	s := ""
	for i, service := range services {
		s = s + "\t" + service.RPCDefinition()
		if i < len(services)-1 {
			s = s + "\n"
		}
	}

	return s
}

func getMessageDefinitions(services []ProtobufService) string {
	s := ""
	for i, service := range services {
		s = s + service.InputParameterMessageDefinition()
		if i < len(services)-1 {
			s = s + "\n\n"
		}
	}

	return s
}

func (ps ProtobufService) RPCDefinition() string {
	return fmt.Sprintf("rpc %s (%s) returns (%s);", ps.Name, ps.InputParameter.Name, ps.ReturnValue)
}

func (ps ProtobufService) InputParameterMessageDefinition() string {
	template := `message %s {
%s
}`

	fieldsDefinition := ""
	for i, field := range ps.InputParameter.Fields {
		fieldsDefinition = fieldsDefinition + "\t" + field.Definition(i+1)
		if i < len(ps.InputParameter.Fields)-1 {
			fieldsDefinition = fieldsDefinition + "\n"
		}
	}

	return fmt.Sprintf(template, ps.InputParameter.Name, fieldsDefinition)
}

func (pf ProtobufField) Definition(protoFieldIndex int) string {
	template := "%s %s = %d"
	if pf.ArrayInfo != nil {
		template = fmt.Sprintf("repeated %v [(array) = \"%s\"]", template, pf.ArrayInfo.SolidityArray)
	}
	template = template + ";"

	return fmt.Sprintf(template, pf.TypeName, pf.FieldName, protoFieldIndex)
}
