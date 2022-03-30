package encoder

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/jjtny1/protoeth/go/gen/solidity"
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
				&solidity.Uint32{
					Uint: 69,
				},
				&solidity.Bool{
					Boolean: true,
				},
			},
			expectedEncodingHex: "00000000000000000000000000000000000000000000000000000000000000450000000000000000000000000000000000000000000000000000000000000001",
		},
		{
			// Test case created on http://remix.ethereum.org/
			name: "store(uint256 num, address f)",
			parameters: []EncodableProto{
				&solidity.Uint32{
					Uint: 350,
				},
				&solidity.Address{
					Address: hexStringToBytes("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
				},
			},
			expectedEncodingHex: "000000000000000000000000000000000000000000000000000000000000015e000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
		},
		{
			// Test case created on http://remix.ethereum.org/
			name: "store(uint256 num, bytes17 b, address f)",
			parameters: []EncodableProto{
				&solidity.Uint256{
					UintBytes: big.NewInt(350).Bytes(),
				},
				&solidity.Bytes17{
					Bytes: hexStringToBytes("7468697320697320612074657374206161"),
				},
				&solidity.Address{
					Address: hexStringToBytes("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
				},
			},
			expectedEncodingHex: "000000000000000000000000000000000000000000000000000000000000015e7468697320697320612074657374206161000000000000000000000000000000000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
		},
		{
			// Test case created on http://remix.ethereum.org/
			name: "store(uint256 num, bytes memory b)",
			parameters: []EncodableProto{
				&solidity.Uint256{
					UintBytes: big.NewInt(350).Bytes(),
				},
				&solidity.Bytes{
					Bytes: hexStringToBytes("48657920746869732069732061207465737420746861742073686f756c64206265206c6f6e676572207468616e203332206279746573"),
				},
			},
			expectedEncodingHex: "000000000000000000000000000000000000000000000000000000000000015e0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000003648657920746869732069732061207465737420746861742073686f756c64206265206c6f6e676572207468616e20333220627974657300000000000000000000",
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
