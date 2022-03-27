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
	IsArray   bool
	TypeName  string
	FieldName string
}

const (
	functionParametersName  = "%s_FunctionParameters"
	functionReturnValueName = "EncodedFunctionCall"
	template                = `
syntax = "proto3";

package solidity;

option go_package = "gen/solidity_contracts";

import "solidity_types.proto";

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
				protoFieldType, isArray, err := SolidityTypeToProtoType(parameter.ParameterType)
				if err != nil {
					return "", err
				}

				service.InputParameter.Fields = append(service.InputParameter.Fields, ProtobufField{
					IsArray:   isArray,
					TypeName:  protoFieldType,
					FieldName: parameter.ParameterName,
				})
			}

			services = append(services, service)
		}
	}

	return fmt.Sprintf(template, getRPCDefinitions(services), getMessageDefinitions(services)), nil
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
		s = s + service.InputParameterMessageDefinition() + "\n\n"
		if i < len(services)-1 {
			s = s + "\n"
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
	template := "%s %s = %d;"
	if pf.IsArray {
		template = "repeated " + template
	}
	return fmt.Sprintf(template, pf.TypeName, pf.FieldName, protoFieldIndex)
}
