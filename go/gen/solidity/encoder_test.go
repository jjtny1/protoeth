package solidity

import (
	"encoding/hex"
	"math/big"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStaticEncoding(t *testing.T) {
	tests := []struct {
		name                string
		parameters          []EncodableProto
		expectedEncodingHex string
	}{
		{
			// Test case from https://docs.soliditylang.org/en/v0.5.3/abi-spec.html#formal-specification-of-the-encoding
			name: "baz(uint32 x, bool y)",
			parameters: []EncodableProto{
				&Type_Uint32{
					Uint32: &Uint32{
						Uint: 69,
					},
				},
				&Type_Bool{
					Bool: &Bool{
						Boolean: true,
					},
				},
			},
			expectedEncodingHex: "00000000000000000000000000000000000000000000000000000000000000450000000000000000000000000000000000000000000000000000000000000001",
		},
		{
			// Test case created on http://remix.ethereum.org/
			name: "store(uint256 num, address f)",
			parameters: []EncodableProto{
				&Type_Uint256{
					Uint256: &Uint256{
						UintBytes: big.NewInt(int64(350)).Bytes(),
					},
				},
				&Type_Address{
					Address: &Address{
						Address: hexStringToBytes("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
					},
				},
			},
			expectedEncodingHex: "000000000000000000000000000000000000000000000000000000000000015e000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
		},
		{
			// Test case created on http://remix.ethereum.org/
			name: "store(uint256 num, bytes17 b, address f)",
			parameters: []EncodableProto{
				&Type_Uint256{
					Uint256: &Uint256{
						UintBytes: big.NewInt(int64(350)).Bytes(),
					},
				},
				&Type_Bytes17{
					Bytes17: &Bytes17{
						Bytes: hexStringToBytes("7468697320697320612074657374206161"),
					},
				},
				&Type_Address{
					Address: &Address{
						Address: hexStringToBytes("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
					},
				},
			},
			expectedEncodingHex: "000000000000000000000000000000000000000000000000000000000000015e7468697320697320612074657374206161000000000000000000000000000000000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
		},
		{
			// Test case created on http://remix.ethereum.org/
			name: "store(uint256 num, bytes memory b)",
			parameters: []EncodableProto{
				&Type_Uint256{
					Uint256: &Uint256{
						UintBytes: big.NewInt(int64(350)).Bytes(),
					},
				},
				&Type_Bytes{
					Bytes: &Bytes{
						Bytes: hexStringToBytes("48657920746869732069732061207465737420746861742073686f756c64206265206c6f6e676572207468616e203332206279746573"),
					},
				},
			},
			expectedEncodingHex: "000000000000000000000000000000000000000000000000000000000000015e0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000003648657920746869732069732061207465737420746861742073686f756c64206265206c6f6e676572207468616e20333220627974657300000000000000000000",
		},
		{
			// Test case created on http://remix.ethereum.org/
			name: "f(uint,uint32[],bytes10,bytes)",
			parameters: []EncodableProto{
				&Type_Uint256{
					Uint256: &Uint256{
						UintBytes: []byte{0x01, 0x23},
					},
				},
				&Type_DynamicArray{
					DynamicArray: &TypeDynamicArray{
						TypeList: []*Type{
							&Type{
								SolidityType: &Type_Uint32{
									Uint32: &Uint32{
										Uint: 0x456,
									},
								},
							},
							&Type{
								SolidityType: &Type_Uint32{
									Uint32: &Uint32{
										Uint: 0x789,
									},
								},
							},
						},
					},
				},
				&Type_Bytes10{
					Bytes10: &Bytes10{
						Bytes: []byte("1234567890"),
					},
				},
				&Type_Bytes{
					Bytes: &Bytes{
						Bytes: []byte("Hello, world!"),
					},
				},
			},
			expectedEncodingHex: "00000000000000000000000000000000000000000000000000000000000001230000000000000000000000000000000000000000000000000000000000000080313233343536373839300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000004560000000000000000000000000000000000000000000000000000000000000789000000000000000000000000000000000000000000000000000000000000000d48656c6c6f2c20776f726c642100000000000000000000000000000000000000",
		},
		{
			// Test case created on http://remix.ethereum.org/
			name: "g(uint[][],string[])",
			parameters: []EncodableProto{
				&Type_DynamicArray{
					DynamicArray: &TypeDynamicArray{
						TypeList: []*Type{
							&Type{
								SolidityType: &Type_DynamicArray{
									DynamicArray: &TypeDynamicArray{
										TypeList: []*Type{
											&Type{
												SolidityType: &Type_Uint256{
													Uint256: &Uint256{
														UintBytes: []byte{0x01},
													},
												},
											},
											&Type{
												SolidityType: &Type_Uint256{
													Uint256: &Uint256{
														UintBytes: []byte{0x02},
													},
												},
											},
										},
									},
								},
							},
							&Type{
								SolidityType: &Type_DynamicArray{
									DynamicArray: &TypeDynamicArray{
										TypeList: []*Type{
											&Type{
												SolidityType: &Type_Uint256{
													Uint256: &Uint256{
														UintBytes: []byte{0x03},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
				&Type_DynamicArray{
					DynamicArray: &TypeDynamicArray{
						TypeList: []*Type{
							&Type{
								SolidityType: &Type_String_{
									String_: &String{
										String_: "one",
									},
								},
							},
							&Type{
								SolidityType: &Type_String_{
									String_: &String{
										String_: "two",
									},
								},
							},
							&Type{
								SolidityType: &Type_String_{
									String_: &String{
										String_: "three",
									},
								},
							},
						},
					},
				},
			},
			expectedEncodingHex: "000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000030000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000000036f6e650000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000374776f000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000057468726565000000000000000000000000000000000000000000000000000000",
		},
	}

	for _, test := range tests {
		boundTest := test
		Convey("Given a list of solidity parameters", t, func() {
			parameters := boundTest.parameters

			Convey("When the parameters are encoded", func() {
				encoding, err := Encode(parameters)
				So(err, ShouldBeNil)

				Convey("Then the correct encoding should be returned", func() {
					So(hex.EncodeToString(encoding), ShouldEqual, boundTest.expectedEncodingHex)
				})
			})
		})
	}
}

func hexStringToBytes(hexString string) []byte {
	if len(hexString) > 2 && hexString[0:2] == "0x" {
		hexString = hexString[2:]
	}

	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		panic("invalid hexString: " + hexString)
	}

	return bytes
}
