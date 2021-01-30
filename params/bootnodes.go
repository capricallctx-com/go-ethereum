// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// census boot - frankfurt
	"enode://f1c78e9b6f5414d6fdf2d846e386f8f583766799c0f721abcc035c25fd13146f84a8cdbe6b354a385e15972446495ec7acbc675bb223b4fb7f7059aa2b0a0146@138.68.91.168:33333", // de.censusprotocol.net
}

// RopstenBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var RopstenBootnodes = []string{
	"enode://57af55c7ee99e0e3dd3c6f211313840cde1294edc86dffe57670bd500438e57b074abc5c9bd862a832ae75d4b6ca7facc4f084acc1b9ab641d6f11ba422d7425@64.225.56.2:30303?discport=30301", // bootnode-aws-ap-southeast-1-001
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://57af55c7ee99e0e3dd3c6f211313840cde1294edc86dffe57670bd500438e57b074abc5c9bd862a832ae75d4b6ca7facc4f084acc1b9ab641d6f11ba422d7425@64.225.56.2:30303?discport=30301", // bootnode-aws-ap-southeast-1-001
}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{
	// Upstream bootnodes
	"enode://57af55c7ee99e0e3dd3c6f211313840cde1294edc86dffe57670bd500438e57b074abc5c9bd862a832ae75d4b6ca7facc4f084acc1b9ab641d6f11ba422d7425@64.225.56.2:30303?discport=30301", // bootnode-aws-ap-southeast-1-001
}

// YoloV2Bootnodes are the enode URLs of the P2P bootstrap nodes running on the
// YOLOv2 ephemeral test network.
var YoloV2Bootnodes = []string{
	"enode://57af55c7ee99e0e3dd3c6f211313840cde1294edc86dffe57670bd500438e57b074abc5c9bd862a832ae75d4b6ca7facc4f084acc1b9ab641d6f11ba422d7425@64.225.56.2:30303?discport=30301", // bootnode-aws-ap-southeast-1-001
}

const dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
	case MainnetGenesisHash:
		net = "mainnet"
	case RopstenGenesisHash:
		net = "ropsten"
	case RinkebyGenesisHash:
		net = "rinkeby"
	case GoerliGenesisHash:
		net = "goerli"
	default:
		return ""
	}
	return dnsPrefix + protocol + "." + net + ".ethdisco.net"
}
