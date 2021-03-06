## Census Protocol

Census is a stable fork of Ethereum for storing public cryptographic data and address 
mappings.   The API is Ethereum compatible as of today with extensions - so you can use any non-contract code on this chain.
Census was originally a custom blockchain - but has moved to an Ethereum base to use the numerous tools available for it.   Documentation updates are behind development - so you will see references to geth and Ethereum in the docs.  We try to not change base code if possible to reduce work when Ethereum releases are merged back into this fork,  so obsolete modules like legacy contracts are still in this repo. 

Where there is a large difference between then - the terminology has been changed.  In particular Ethereum's gas policy has been normalized with the introduction of max 
tariffs to make the process more predictable, and generic contract functionality has 
been mostly ripped out in for specific cargos to be included in the chain. 

*'Tariffs'* is used to describe the max charges and QoS that is required of carriers in a regulated market - this addition makes it possible to compute exactly the correct amount of gas needed to store the cargo on the client - although overpaying causes priority mining.

Cargo is a series of defined contracts to storage and retrieve select data.  Nodes validate the cargo as part of the transaction checks - users are not allowed to store arbitrary data on the system to prevent storage of illegal files and things like revenge porn and embarrassing photos.   This code is a port from Census 1.0 and does not match the original ABI.  The initial byte of each cargo is 0x54 followed by a cargo type.  

To prevent 51% attacks since the network is small - new miners need to be approved by the current group.  Add an issue to the code repo.

Current Cargo

|    Cargo   | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| :-----------: | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|  `CargoPublicKey`   | 0x01 - announces a public key on the network.  Cargo type tag must be followed by a public key in PEM format.   The chain will parse and reject the PEM if invalid.    |
|   `CargoClaim`    | 0x02 - Claim an email or social media address for this account - will appear in indices.  Type tag is followed immediately by a subtype and up to 64 bytes of data.   Subtypes are 0x00 (email), 0x01 Twitter, 0x02 Hayachat. |
|  `CargoAttest`   | 0x03 - Attest to an address being owned by an account (improves it's score).  Followed by a claim subtype, an address, and the tag itself (up to 64 bytes)         |

- CargoPublicKey

In raw format, this is:

```bigquery
0x5401001677696c6c69616d6c6f70657a407961686f6f2e636f6d2d2d2d2d2d424547494e205055424c4943204b45592d2d2d2d2d0a4d494943496a414e42676b71686b6947397730424151454641414f43416738414d49494343674b43416745416b3070766b362f355351707a4151756c4e6874750a4d785a72634b5a6f37723368742f585578494d5132686935596663514b2b7059526d44744e646232493770316661672b723656664a66736a344d36656b5432330a656c674859733077355a357961683848654b4e7a69614948587935396d47567a342b36796150765035527a4c7841394c714777592b36644c6664577a6d4c71330a765a3357496b53497452316a50544642796c584b5965574f7a446a49415a59586f3973484f2f7657506972662f6750324539504f424363454e5a763965786b6b0a356939506d6d7a2f636c2b4b71586945385768624956567642786d7070514541363153373737584b786878375736485152596b4c3263466c54397a39747434680a554d4e716d4a476475556d4d4b646931797872735a306e6146316839526b61506c3047465666694264326a7573494267484241527059625757373341786265720a49393858694143743862664d596a39674351506d2b6a6955667758357379476657626b4c4a636d745953754e59465937332f43776f4d6b564e2f4a37394c59370a77786755344c4a5259684e384275674d4345595751772f4948496833516e682f37305176536a6e446f6e336f69565a66544f754631734347487370716375522f0a524d30692f4a66476430526948645250576f4e433836565637764866664570564f514c416e48416b4f6153616336756d2f687074752f646e6f42514b4b4535650a77523176747a4a6e4d6734436c30654b6b6e2b6d2f62674548797a77696a676269353437344e4955453470494854394b5159354c6d6f4661456573776c6d72780a506e564d59585273386a307553686257636651595a5373627a6751426a457352585739304272626e6e4b61744f7169412b3764756f4a75515a647941464850590a444a6d71424563625357767945434b6563517371663173434177454141513d3d0a2d2d2d2d2d454e44205055424c4943204b45592d2d2d2d2d0a
```
Note that organizations may use a single rotating key for records preservation (i.e. - the message may be decrypted by the SMTP or XMPP gateway).
For this case, the fingerprint is at an org level (i.e. - 'example.com'). To parse the data - the fields are

|    Byte   | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| :-----------: | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|  0x00   | tanker tag (0x54) |
|  0x01   | Cargo type (0x01 - public key) |
|  0x02   | Key type hint (0x00 is RSA) |
|  0x03   | Fingerprint length - tag to identify such as an email address, GPG label, or social network address |
|  0x04   | The fingerprint itself, limited to 64 characters |
|  0x04+fingerprint length   | The public key, this is is parsed for validity by nodes |

|   `CargoClaim`    | 0x02 - Claim an email or social media address for this account - will appear in indices.  Type tag is followed immediately by a subtype and up to 64 bytes of data.   Subtypes are 0x00 (email), 0x01 Twitter, 0x02 Hayachat. |
|  `CargoAttest`   | 0x03 - Attest to an address being owned by an account (improves it's score).  Followed by a claim subtype, an address, and the tag itself (up to 64 bytes)         |

![img.png](img.png)

[![API Reference](
https://camo.githubusercontent.com/915b7be44ada53c290eb157634330494ebe3e30a/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f676f6c616e672f6764646f3f7374617475732e737667
)](https://pkg.go.dev/github.com/ethereum/go-ethereum?tab=doc)


## Building the source

For prerequisites and detailed build instructions please read the [Installation Instructions](https://github.com/ethereum/go-ethereum/wiki/Building-Ethereum) on the wiki.

Building `tankerd` requires both a Go (version 1.13 or later) and a C compiler. You can install
them using your favourite package manager. Once the dependencies are installed, run

```shell
make tankerd
```

or, to build the full suite of utilities:

```shell
make all
```

## Executables

The go-ethereum project comes with several wrappers/executables found in the `cmd`
directory.

|    Command    | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| :-----------: | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|  **`tankerd`**   | Our main Ethereum CLI client. It is the entry point into the Ethereum network (main-, test- or private net), capable of running as a full node (default), archive node (retaining all historical state) or a light node (retrieving data live). It can be used by other processes as a gateway into the Ethereum network via JSON RPC endpoints exposed on top of HTTP, WebSocket and/or IPC transports. `tankerd --help` and the [CLI Wiki page](https://github.com/ethereum/go-ethereum/wiki/Command-Line-Options) for command line options.          |
|   `abigen`    | Source code generator to convert Ethereum contract definitions into easy to use, compile-time type-safe Go packages. It operates on plain [Ethereum contract ABIs](https://github.com/ethereum/wiki/wiki/Ethereum-Contract-ABI) with expanded functionality if the contract bytecode is also available. However, it also accepts Solidity source files, making development much more streamlined. Please see our [Native DApps](https://github.com/ethereum/go-ethereum/wiki/Native-DApps:-Go-bindings-to-Ethereum-contracts) wiki page for details. |
|  `bootnode`   | Stripped down version of our Ethereum client implementation that only takes part in the network node discovery protocol, but does not run any of the higher level application protocols. It can be used as a lightweight bootstrap node to aid in finding peers in private networks.                                                                                                                                                                                                                                                                 |
|     `evm`     | Developer utility version of the EVM (Ethereum Virtual Machine) that is capable of running bytecode snippets within a configurable environment and execution mode. Its purpose is to allow isolated, fine-grained debugging of EVM opcodes (e.g. `evm --code 60ff60ff --debug run`).                                                                                                                                                                                                                                                                     |
| `gethrpctest` | Developer utility tool to support our [ethereum/rpc-test](https://github.com/ethereum/rpc-tests) test suite which validates baseline conformity to the [Ethereum JSON RPC](https://github.com/ethereum/wiki/wiki/JSON-RPC) specs. Please see the [test suite's readme](https://github.com/ethereum/rpc-tests/blob/master/README.md) for details.                                                                                                                                                                                                     |
|   `rlpdump`   | Developer utility tool to convert binary RLP ([Recursive Length Prefix](https://github.com/ethereum/wiki/wiki/RLP)) dumps (data encoding used by the Ethereum protocol both network as well as consensus wise) to user-friendlier hierarchical representation (e.g. `rlpdump --hex CE0183FFFFFFC4C304050583616263`).                                                                                                                                                                                                                                 |
|   `puppeth`   | a CLI wizard that aids in creating a new Ethereum network.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |

## Running `geth`

Going through all the possible command line flags is out of scope here (please consult our
[CLI Wiki page](https://github.com/ethereum/go-ethereum/wiki/Command-Line-Options)),
but we've enumerated a few common parameter combos to get you up to speed quickly
on how you can run your own `tankerd` instance.

### Full node on the main Ethereum network

By far the most common scenario is people wanting to simply interact with the Ethereum
network: create accounts; transfer funds; deploy and interact with contracts. For this
particular use-case the user doesn't care about years-old historical data, so we can
fast-sync quickly to the current state of the network. To do so:

```shell
$ tankerd console
```

This command will:
 * Start `tankerd` in fast sync mode (default, can be changed with the `--syncmode` flag),
   causing it to download more data in exchange for avoiding processing the entire history
   of the Ethereum network, which is very CPU intensive.
 * Start up `tankerd`'s built-in interactive [JavaScript console](https://github.com/ethereum/go-ethereum/wiki/JavaScript-Console),
   (via the trailing `console` subcommand) through which you can invoke all official [`web3` methods](https://github.com/ethereum/wiki/wiki/JavaScript-API)
   as well as `tankerd`'s own [management APIs](https://github.com/ethereum/go-ethereum/wiki/Management-APIs).
   This tool is optional and if you leave it out you can always attach to an already running
   `tankerd` instance with `tankerd attach`.

### A Full node on the Görli test network

Transitioning towards developers, if you'd like to play around with creating Ethereum
contracts, you almost certainly would like to do that without any real money involved until
you get the hang of the entire system. In other words, instead of attaching to the main
network, you want to join the **test** network with your node, which is fully equivalent to
the main network, but with play-Ether only.

```shell
$ tankerd --goerli console
```

The `console` subcommand has the exact same meaning as above and they are equally
useful on the testnet too. Please, see above for their explanations if you've skipped here.

Specifying the `--goerli` flag, however, will reconfigure your `tankerd` instance a bit:

 * Instead of connecting the main Ethereum network, the client will connect to the Görli
   test network, which uses different P2P bootnodes, different network IDs and genesis
   states.
 * Instead of using the default data directory (`~/.ethereum` on Linux for example), `tankerd`
   will nest itself one level deeper into a `goerli` subfolder (`~/.ethereum/goerli` on
   Linux). Note, on OSX and Linux this also means that attaching to a running testnet node
   requires the use of a custom endpoint since `tankerd attach` will try to attach to a
   production node endpoint by default, e.g.,
   `tankerd attach <datadir>/goerli/tankerd.ipc`. Windows users are not affected by
   this.

*Note: Although there are some internal protective measures to prevent transactions from
crossing over between the main network and test network, you should make sure to always
use separate accounts for play-money and real-money. Unless you manually move
accounts, `tankerd` will by default correctly separate the two networks and will not make any
accounts available between them.*



### Configuration

As an alternative to passing the numerous flags to the `tankerd` binary, you can also pass a
configuration file via:

```shell
$ tankerd --config /path/to/your_config.toml
```

To get an idea how the file should look like you can use the `dumpconfig` subcommand to
export your existing configuration:

```shell
$ tankerd --your-favourite-flags dumpconfig
```

*Note: This works only with `tankerd` v1.6.0 and above.*

### Programmatically interfacing `tankerd` nodes

As a developer, sooner rather than later you'll want to start interacting with `tankerd` and the
Ethereum network via your own programs and not manually through the console. To aid
this, `tankerd` has built-in support for a JSON-RPC based APIs ([standard APIs](https://github.com/ethereum/wiki/wiki/JSON-RPC)
and [`tankerd` specific APIs](https://github.com/ethereum/go-ethereum/wiki/Management-APIs)).
These can be exposed via HTTP, WebSockets and IPC (UNIX sockets on UNIX based
platforms, and named pipes on Windows).

The IPC interface is enabled by default and exposes all the APIs supported by `tankerd`,
whereas the HTTP and WS interfaces need to manually be enabled and only expose a
subset of APIs due to security reasons. These can be turned on/off and configured as
you'd expect.

HTTP based JSON-RPC API options:

  * `--http` Enable the HTTP-RPC server
  * `--http.addr` HTTP-RPC server listening interface (default: `localhost`)
  * `--http.port` HTTP-RPC server listening port (default: `8545`)
  * `--http.api` API's offered over the HTTP-RPC interface (default: `eth,net,web3`)
  * `--http.corsdomain` Comma separated list of domains from which to accept cross origin requests (browser enforced)
  * `--ws` Enable the WS-RPC server
  * `--ws.addr` WS-RPC server listening interface (default: `localhost`)
  * `--ws.port` WS-RPC server listening port (default: `8546`)
  * `--ws.api` API's offered over the WS-RPC interface (default: `eth,net,web3`)
  * `--ws.origins` Origins from which to accept websockets requests
  * `--ipcdisable` Disable the IPC-RPC server
  * `--ipcapi` API's offered over the IPC-RPC interface (default: `admin,debug,eth,miner,net,personal,shh,txpool,web3`)
  * `--ipcpath` Filename for IPC socket/pipe within the datadir (explicit paths escape it)

You'll need to use your own programming environments' capabilities (libraries, tools, etc) to
connect via HTTP, WS or IPC to a `tankerd` node configured with the above flags and you'll
need to speak [JSON-RPC](https://www.jsonrpc.org/specification) on all transports. You
can reuse the same connection for multiple requests!

**Note: Please understand the security implications of opening up an HTTP/WS based
transport before doing so! Hackers on the internet are actively trying to subvert
Ethereum nodes with exposed APIs! Further, all browser tabs can access locally
running web servers, so malicious web pages could try to subvert locally available
APIs!**

### Operating a private network

Maintaining your own private network is more involved as a lot of configurations taken for
granted in the official networks need to be manually set up.

#### Defining the private genesis state

First, you'll need to create the genesis state of your networks, which all nodes need to be
aware of and agree upon. This consists of a small JSON file (e.g. call it `genesis.json`):

```json
{
  "config": {
    "chainId": <arbitrary positive integer>,
    "homesteadBlock": 0,
    "eip150Block": 0,
    "eip155Block": 0,
    "eip158Block": 0,
    "byzantiumBlock": 0,
    "constantinopleBlock": 0,
    "petersburgBlock": 0,
    "istanbulBlock": 0
  },
  "alloc": {},
  "coinbase": "0x0000000000000000000000000000000000000000",
  "difficulty": "0x20000",
  "extraData": "",
  "gasLimit": "0x2fefd8",
  "nonce": "0x0000000000000042",
  "mixhash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "timestamp": "0x00"
}
```

The above fields should be fine for most purposes, although we'd recommend changing
the `nonce` to some random value so you prevent unknown remote nodes from being able
to connect to you. If you'd like to pre-fund some accounts for easier testing, create
the accounts and populate the `alloc` field with their addresses.

```json
"alloc": {
  "0x0000000000000000000000000000000000000001": {
    "balance": "111111111"
  },
  "0x0000000000000000000000000000000000000002": {
    "balance": "222222222"
  }
}
```

With the genesis state defined in the above JSON file, you'll need to initialize **every**
`tankerd` node with it prior to starting it up to ensure all blockchain parameters are correctly
set:

```shell
$ tankerd init path/to/genesis.json
```

#### Creating the rendezvous point

With all nodes that you want to run initialized to the desired genesis state, you'll need to
start a bootstrap node that others can use to find each other in your network and/or over
the internet. The clean way is to configure and run a dedicated bootnode:

```shell
$ bootnode --genkey=boot.key
$ bootnode --nodekey=boot.key
```

With the bootnode online, it will display an [`enode` URL](https://github.com/ethereum/wiki/wiki/enode-url-format)
that other nodes can use to connect to it and exchange peer information. Make sure to
replace the displayed IP address information (most probably `[::]`) with your externally
accessible IP to get the actual `enode` URL.

*Note: You could also use a full-fledged `tankerd` node as a bootnode, but it's the less
recommended way.*

#### Starting up your member nodes

With the bootnode operational and externally reachable (you can try
`telnet <ip> <port>` to ensure it's indeed reachable), start every subsequent `tankerd`
node pointed to the bootnode for peer discovery via the `--bootnodes` flag. It will
probably also be desirable to keep the data directory of your private network separated, so
do also specify a custom `--datadir` flag.

```shell
$ tankerd --datadir=path/to/custom/data/folder --bootnodes=<bootnode-enode-url-from-above>
```

*Note: Since your network will be completely cut off from the main and test networks, you'll
also need to configure a miner to process transactions and create new blocks for you.*

#### Running a private miner

Mining on the public Ethereum network is a complex task as it's only feasible using GPUs,
requiring an OpenCL or CUDA enabled `ethminer` instance. For information on such a
setup, please consult the [EtherMining subreddit](https://www.reddit.com/r/EtherMining/)
and the [ethminer](https://github.com/ethereum-mining/ethminer) repository.

In a private network setting, however a single CPU miner instance is more than enough for
practical purposes as it can produce a stable stream of blocks at the correct intervals
without needing heavy resources (consider running on a single thread, no need for multiple
ones either). To start a `tankerd` instance for mining, run it with all your usual flags, extended
by:

```shell
$ tankerd <usual-flags> --mine --miner.threads=1 --etherbase=0x0000000000000000000000000000000000000000
```

Which will start mining blocks and transactions on a single CPU thread, crediting all
proceedings to the account specified by `--etherbase`. You can further tune the mining
by changing the default gas limit blocks converge to (`--targetgaslimit`) and the price
transactions are accepted at (`--gasprice`).

## Contribution

Thank you for considering to help out with the source code! We welcome contributions
from anyone on the internet, and are grateful for even the smallest of fixes!

If you'd like to contribute to go-ethereum, please fork, fix, commit and send a pull request
for the maintainers to review and merge into the main code base. If you wish to submit
more complex changes though, please check up with the core devs first on [our gitter channel](https://gitter.im/ethereum/go-ethereum)
to ensure those changes are in line with the general philosophy of the project and/or get
some early feedback which can make both your efforts much lighter as well as our review
and merge procedures quick and simple.

Please make sure your contributions adhere to our coding guidelines:

 * Code must adhere to the official Go [formatting](https://golang.org/doc/effective_go.html#formatting)
   guidelines (i.e. uses [gofmt](https://golang.org/cmd/gofmt/)).
 * Code must be documented adhering to the official Go [commentary](https://golang.org/doc/effective_go.html#commentary)
   guidelines.
 * Pull requests need to be based on and opened against the `master` branch.
 * Commit messages should be prefixed with the package(s) they modify.
   * E.g. "eth, rpc: make trace configs optional"

Please see the [Developers' Guide](https://github.com/ethereum/go-ethereum/wiki/Developers'-Guide)
for more details on configuring your environment, managing project dependencies, and
testing procedures.

## License

The go-ethereum library (i.e. all code outside of the `cmd` directory) is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html),
also included in our repository in the `COPYING.LESSER` file.

The go-ethereum binaries (i.e. all code inside of the `cmd` directory) is licensed under the
[GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also
included in our repository in the `COPYING` file.
