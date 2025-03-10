// Go Substrate RPC Client (GSRPC) provides APIs and types around Polkadot and any Substrate-based chain RPC calls
//
// Copyright 2019 Centrifuge GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types_test

import (
	"encoding/json"
	"testing"

	. "github.com/snowfork/go-substrate-rpc-client/v4/types"
)

func TestGenerateMMRProofResponse_Unmarshal(t *testing.T) {
	jsonData := map[string]interface{}{"blockHash": "0x62318380e42f2a47f3b74d63bf1d7dc245920766759661f3beb51d35f9f2450e", "leaf": "0xc10198140000264e19fe1a3795cf2aefd6b260261674c6bf642af60eb499f8a4e4d976576375480907734fe2450e32a567628c22c5374c71b69d091e4c0219c571ed932a0e5301000000000000000200000007b13d25743592825cea32c9a24ba67b50b7e90d92cbd1d0f4eab2dc94dba5c6", "proof": "0x98140000000000009a1400000000000018fcb383a3d3f35fc452db71cd423925d8c8f9d0708de71e5a22d18e2a54fea0e98efb55ea18c7bec19de2c247e9bbcd0162c7b8fed10d43d9d88add1e80bcdf35839159300b267069f62c5bc7d3c02cd4f4186eb0355f62acda9ba3a6431f75c1887f4e400eef548eba662c01bf8c7a67608d56fbbc179ef890f2d55184e42b03685b29893fe0ebe38d937027b372e6e984aee88c19c14c80512f48975244beb224f2267daabe69303850a53272e70c00592d663a12d374db22c98a63c670bffb"}

	marshalled, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	expected := GenerateMMRProofResponse{BlockHash: H256{0x62, 0x31, 0x83, 0x80, 0xe4, 0x2f, 0x2a, 0x47, 0xf3, 0xb7, 0x4d, 0x63, 0xbf, 0x1d, 0x7d, 0xc2, 0x45, 0x92, 0x7, 0x66, 0x75, 0x96, 0x61, 0xf3, 0xbe, 0xb5, 0x1d, 0x35, 0xf9, 0xf2, 0x45, 0xe}, Leaf: MMRLeaf{ParentNumberAndHash: ParentNumberAndHash{ParentNumber: 0x1498, Hash: Hash{0x26, 0x4e, 0x19, 0xfe, 0x1a, 0x37, 0x95, 0xcf, 0x2a, 0xef, 0xd6, 0xb2, 0x60, 0x26, 0x16, 0x74, 0xc6, 0xbf, 0x64, 0x2a, 0xf6, 0xe, 0xb4, 0x99, 0xf8, 0xa4, 0xe4, 0xd9, 0x76, 0x57, 0x63, 0x75}}, ParachainHeads: H256{0x48, 0x9, 0x7, 0x73, 0x4f, 0xe2, 0x45, 0xe, 0x32, 0xa5, 0x67, 0x62, 0x8c, 0x22, 0xc5, 0x37, 0x4c, 0x71, 0xb6, 0x9d, 0x9, 0x1e, 0x4c, 0x2, 0x19, 0xc5, 0x71, 0xed, 0x93, 0x2a, 0xe, 0x53}, BeefyNextAuthoritySet: BeefyNextAuthoritySet{ID: 0x1, Len: 0x2, Root: H256{0x7, 0xb1, 0x3d, 0x25, 0x74, 0x35, 0x92, 0x82, 0x5c, 0xea, 0x32, 0xc9, 0xa2, 0x4b, 0xa6, 0x7b, 0x50, 0xb7, 0xe9, 0xd, 0x92, 0xcb, 0xd1, 0xd0, 0xf4, 0xea, 0xb2, 0xdc, 0x94, 0xdb, 0xa5, 0xc6}}}, Proof: MMRProof{LeafIndex: 0x1498, LeafCount: 0x149a, Items: []H256{H256{0xfc, 0xb3, 0x83, 0xa3, 0xd3, 0xf3, 0x5f, 0xc4, 0x52, 0xdb, 0x71, 0xcd, 0x42, 0x39, 0x25, 0xd8, 0xc8, 0xf9, 0xd0, 0x70, 0x8d, 0xe7, 0x1e, 0x5a, 0x22, 0xd1, 0x8e, 0x2a, 0x54, 0xfe, 0xa0, 0xe9}, H256{0x8e, 0xfb, 0x55, 0xea, 0x18, 0xc7, 0xbe, 0xc1, 0x9d, 0xe2, 0xc2, 0x47, 0xe9, 0xbb, 0xcd, 0x1, 0x62, 0xc7, 0xb8, 0xfe, 0xd1, 0xd, 0x43, 0xd9, 0xd8, 0x8a, 0xdd, 0x1e, 0x80, 0xbc, 0xdf, 0x35}, H256{0x83, 0x91, 0x59, 0x30, 0xb, 0x26, 0x70, 0x69, 0xf6, 0x2c, 0x5b, 0xc7, 0xd3, 0xc0, 0x2c, 0xd4, 0xf4, 0x18, 0x6e, 0xb0, 0x35, 0x5f, 0x62, 0xac, 0xda, 0x9b, 0xa3, 0xa6, 0x43, 0x1f, 0x75, 0xc1}, H256{0x88, 0x7f, 0x4e, 0x40, 0xe, 0xef, 0x54, 0x8e, 0xba, 0x66, 0x2c, 0x1, 0xbf, 0x8c, 0x7a, 0x67, 0x60, 0x8d, 0x56, 0xfb, 0xbc, 0x17, 0x9e, 0xf8, 0x90, 0xf2, 0xd5, 0x51, 0x84, 0xe4, 0x2b, 0x3}, H256{0x68, 0x5b, 0x29, 0x89, 0x3f, 0xe0, 0xeb, 0xe3, 0x8d, 0x93, 0x70, 0x27, 0xb3, 0x72, 0xe6, 0xe9, 0x84, 0xae, 0xe8, 0x8c, 0x19, 0xc1, 0x4c, 0x80, 0x51, 0x2f, 0x48, 0x97, 0x52, 0x44, 0xbe, 0xb2}, H256{0x24, 0xf2, 0x26, 0x7d, 0xaa, 0xbe, 0x69, 0x30, 0x38, 0x50, 0xa5, 0x32, 0x72, 0xe7, 0xc, 0x0, 0x59, 0x2d, 0x66, 0x3a, 0x12, 0xd3, 0x74, 0xdb, 0x22, 0xc9, 0x8a, 0x63, 0xc6, 0x70, 0xbf, 0xfb}}}}

	var unmarshalled GenerateMMRProofResponse

	json.Unmarshal(marshalled, &unmarshalled)

	assertEqual(t, unmarshalled, expected)
}
