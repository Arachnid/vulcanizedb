// VulcanizeDB
// Copyright © 2018 Vulcanize

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package documentation

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/vulcanize/pkg/filters"
	"math/big"
)

// Entities are created for unpacking ethereum logs
// Struct field names and types MUST match the argument names and types in the ABI

// ENS

// event NewOwner(bytes32 indexed node, bytes32 indexed label, address owner);
type NewOwnerEntity struct {
	Node  []byte
	Label []byte
	Owner common.Address
}

// event Transfer(bytes32 indexed node, address owner);
type TransferEntity struct {
	Node  []byte
	Owner common.Address
}

// event NewResolver(bytes32 indexed node, address resolver);
type NewResolverEntity struct {
	Node     []byte
	Resolver common.Address
}

// event NewTTL(bytes32 indexed node, uint64 ttl);
type NewTTLEntity struct {
	Node []byte
	Ttl  *big.Int
}

// Models are create for persisting unpacked ethereum logs into postgres
// Field names are labeled with their corresponding db column names and custom ethereum types are resolved to those that are handled by postgres

type NewOwnerModel struct {
	Node  string `db:"node"`
	Label string `db:"label"`
	Owner string `db:"owner"`
}

type TransferModel struct {
	Node  string `db:"node"`
	Owner string `db:"owner"`
}

type NewResolverModel struct {
	Node     string `db:"node"`
	Resolver string `db:"resolver"`
}

type NewTTLModel struct {
	Node string `db:"node"`
	Ttl  string `db:"ttl"`
}

// Filters are created to request specific event logs (Full mode only)
var ENSFilters = []filters.LogFilter{
	{
		Name:      "NewOwner",
		FromBlock: 0,  // Block to begin search
		ToBlock:   -1, // Block to end search
		Address:   "", // Filter for ENS contract using address
		Topics: core.Topics{ // Filters for NewOrder event using signature
			helpers.GenerateSignature("NewOwner(bytes32,bytes32,address)")},
	},
	{
		Name:      "Transfer",
		FromBlock: 0,
		ToBlock:   -1,
		Address:   "",
		Topics: core.Topics{
			helpers.GenerateSignature("Transfer(bytes32,address)")},
	},
	{
		Name:      "NewResolver",
		FromBlock: 0,
		ToBlock:   -1,
		Address:   "",
		Topics: core.Topics{
			helpers.GenerateSignature("NewResolver(bytes32,address)")},
	},
	{
		Name:      "NewTTL",
		FromBlock: 0,
		ToBlock:   -1,
		Address:   "",
		Topics: core.Topics{
			helpers.GenerateSignature("NewTTL(bytes32,uint64)")},
	},
}

// PublicResolver

// event AddrChanged(bytes32 indexed node, address a);
type AddrChangedEntity struct {
	Node []byte
	A    common.Address
}

type AddrChangedModel struct {
	Node    string `db:"node"`
	Address string `db:"address"`
}

/*
	rest of resolver events:
 	event AddrChanged(bytes32 indexed node, address a);
    event ContentChanged(bytes32 indexed node, bytes32 hash);
    event NameChanged(bytes32 indexed node, string name);
    event ABIChanged(bytes32 indexed node, uint256 indexed contentType);
    event PubkeyChanged(bytes32 indexed node, bytes32 x, bytes32 y);
    event TextChanged(bytes32 indexed node, string indexedKey, string key);
    event MultihashChanged(bytes32 indexed node, bytes hash);
*/

var PublicResolverFilters = []filters.LogFilter{
	{
		Name:      "AddrChanged",
		FromBlock: 0,
		ToBlock:   -1,
		Address:   "",
		Topics: core.Topics{
			helpers.GenerateSignature("AddrChanged(bytes32,address)")},
	},
}
