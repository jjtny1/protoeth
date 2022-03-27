package encoder

import (
	"fmt"
	"testing"

	"github.com/jjtny1/protoeth/go/gen/solidity"
)

func TestStuff(t *testing.T) {
	uint := solidity.Uint{
		Type:      solidity.UintType_UINT256,
		UintBytes: []byte{0x01},
	}

	for i := 0; i < uint.ProtoReflect().Descriptor().Fields().Len(); i = i + 1 {
		kind := uint.ProtoReflect().Descriptor().Fields().Get(i).Kind()
		if kind == 14 {
			fmt.Println("enum is", uint.ProtoReflect().Descriptor().Fields().Get(i).Enum().Values().Get(0).Name())
		}
		fmt.Println("here", uint.ProtoReflect().Descriptor().Fields().Get(i).Kind())
		fmt.Println("here", uint.ProtoReflect().Descriptor().Name())
	}
}
