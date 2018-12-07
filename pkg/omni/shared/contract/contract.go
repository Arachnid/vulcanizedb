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

package contract

import (
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/filters"
	"github.com/vulcanize/vulcanizedb/pkg/omni/shared/helpers"
	"github.com/vulcanize/vulcanizedb/pkg/omni/shared/types"
)

// Contract object to hold our contract data
type Contract struct {
	Name           string                       // Name of the contract
	Address        string                       // Address of the contract
	Network        string                       // Network on which the contract is deployed; default empty "" is Ethereum mainnet
	StartingBlock  int64                        // Starting block of the contract
	LastBlock      int64                        // Most recent block on the network
	Abi            string                       // Abi string
	ParsedAbi      abi.ABI                      // Parsed abi
	Events         map[string]types.Event       // Map of events to their names
	Methods        map[string]types.Method      // Map of methods to their names
	Filters        map[string]filters.LogFilter // Map of event filters to their names; used only for full sync watcher
	FilterArgs     map[string]bool              // User-input list of values to filter event logs for
	MethodArgs     map[string]bool              // User-input list of values to limit method polling to
	EmittedAddrs   map[string]bool              // List of all addresses collected from converted event logs
	EmittedBytes   map[string]bool              // List of all bytes collected from converted event logs
	EmittedHashes  map[string]bool              // List of all hashes collected from converted event logs
	CreateAddrList bool 						// Whether or not to persist address list to postgres
}

// If we are watching events that emit addr, hash, or byte arrays
// then we initialize map to hold the emitted values
func (c Contract) Init() *Contract {
	for _, event := range c.Events {
		for _, field := range event.Fields {
			switch field.Type.T {
			case abi.AddressTy:
				c.EmittedAddrs = map[string]bool{}
			case abi.HashTy:
				c.EmittedHashes = map[string]bool{}
			case abi.BytesTy, abi.FixedBytesTy:
				c.EmittedBytes = map[string]bool{}
			default:
			}
		}
	}

	return &c
}

// Use contract info to generate event filters - full sync omni watcher only
func (c *Contract) GenerateFilters() error {
	c.Filters = map[string]filters.LogFilter{}

	for name, event := range c.Events {
		c.Filters[name] = filters.LogFilter{
			Name:      name,
			FromBlock: c.StartingBlock,
			ToBlock:   -1,
			Address:   c.Address,
			Topics:    core.Topics{helpers.GenerateSignature(event.Sig())}, // move generate signatrue to pkg
		}
	}
	// If no filters were generated, throw an error (no point in continuing with this contract)
	if len(c.Filters) == 0 {
		return errors.New("error: no filters created")
	}

	return nil
}

// Returns true if address is in list of arguments to
// filter events for or if no filtering is specified
func (c *Contract) WantedEventArg(arg string) bool {
	if c.FilterArgs == nil {
		return false
	} else if len(c.FilterArgs) == 0 {
		return true
	} else if a, ok := c.FilterArgs[arg]; ok {
		return a
	}

	return false
}

// Returns true if address is in list of arguments to
// poll methods with or if no filtering is specified
func (c *Contract) WantedMethodArg(arg string) bool {
	if c.MethodArgs == nil {
		return false
	} else if len(c.MethodArgs) == 0 {
		return true
	} else if a, ok := c.MethodArgs[arg]; ok {
		return a
	}

	return false
}

// Returns true if any mapping value matches filtered for address or if no filter exists
// Used to check if an event log name-value mapping should be filtered or not
func (c *Contract) PassesEventFilter(args map[string]string) bool {
	for _, arg := range args {
		if c.WantedEventArg(arg) {
			return true
		}
	}

	return false
}

// Add event emitted address to our list if it passes filter and method polling is on
func (c *Contract) AddEmittedAddr(addresses ...string) {
	for _, addr := range addresses {
		if c.WantedMethodArg(addr) && c.Methods != nil {
			c.EmittedAddrs[addr] = true
		}
	}
}

// Add event emitted hash to our list if it passes filter and method polling is on
func (c *Contract) AddEmittedHash(hashes ...string) {
	for _, hash := range hashes {
		if c.WantedMethodArg(hash) && c.Methods != nil {
			c.EmittedHashes[hash] = true
		}
	}
}


// Add event emitted bytes to our list if it passes filter and method polling is on
func (c *Contract) AddEmittedBytes(byteArrays ...string) {
	for _, bytes := range byteArrays {
		if c.WantedMethodArg(bytes) && c.Methods != nil {
			c.EmittedBytes[bytes] = true
		}
	}
}

