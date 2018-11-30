# Light Transformer Guide 

This is a guide to creating transformers that use a lightSynced vulcanizeDB <br/> 
to capture contract event logs and transform them into custom postgres tables  <br/>

Given the event NewOwner(bytes32 indexed node, bytes32 indexed label, address owner) <br/>
from the ENS-Registrar contract 0x6090A6e47849629b7245Dfa1Ca21D94cd15878Ef  <br/>
How do we capture pull this out from a lightSynced vDB? <br/>

To start, we create custom entity and model structures to transform the log data into: <br/>

```go
type NewOwnerEntity struct {
	Node  []byte
	Label []byte
	Owner common.Address
}

type NewOwnerModel struct {
	Node  string `db:"node"`
	Label string `db:"label"`
	Owner string `db:"owner"`
}
```

And custom event filters that are used to fetch event logs of interest <br/>

```go
var ENSFilters = []filters.LogFilter{
	{
		Name:      "NewOwner",
		FromBlock: 3605331,                                       // Block to begin search e.g. height the contract was published
		ToBlock:   -1,                                            // Block to end search
		Address:   "0x6090A6e47849629b7245Dfa1Ca21D94cd15878Ef",  // Contract address to filter for
		Topics:    core.Topics{                                   // Event signature(s) to filter for
			helpers.GenerateSignature("NewOwner(bytes32,bytes32,address)")},
	},
}

```

The entity is used to unpack a raw geth/core/types.Log fetched using ... <br/>
As such, the field names and types need to match those expected for the event <br/>

The model is used to persist the log into postgres <br/>
As such, the fields need to be labelled with their corresponding database column id </br>
and data types need to be resolved to types that postgres can handle </br>

This is why we need both a entity and model- the entity conforms to the abi and as such </br>
allows for direct unpacking of the log into itself, whereas the model conforms to our postgres </br>
schema and as such allows for persistence of the data into our tables </br>

To mediate the conversion from out entity to our model, we create a converter </br>
This converter is responsible for unpacking the raw eth logs fetched from geth </br>
into our entity and then converting the entity into our model </br>

Once the entity has been created 

