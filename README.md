# metaid-base-model

What is metaid-base-model
-------------------------
This is a universal model for MetaID data, a comprehensive framework designed to store various MetaID data, encompassing fundamental MetaID details, attribute data related to MetaID, and more.

ShowMAN's querying system primarily relies on this universal model for general inquiries.

How to use metaid-base-model
----------------------------
By utilizing the metaid-base-model as a reference, you can implement fundamental operations on MetaAid data. These operations include on-chain data synchronization, data querying, data modification, and more.

The metaid-base-model is a universal model for MetaID data, and it is not designed to be used directly. Instead, it is designed to be extended by other models. For example, the metaid-buzz-model is a model that extends the metaid-base-model, and it is designed to store buzz data.


base-model-structure
--------------------
```
{
    "ChainFlag": "",//string; This refers to the identifier of the blockchain where the MetaID data is hosted. For example, the identifier for the Microvisionchain is "mvc".
    "nodeId": "",//string; Retrieve it by concatenating the <publicKey> with the <txId> and then applying SHA256.
    "metanetId": "",//string; Retrieve it by concatenating the <publicKey> with the <parentTxId> and then applying SHA256. 
    "rootTxId": "",//string; txId of root node
    "txId": "mvc",//string; txId of current node
    "outputIndex": 0,//int; This refers to the index of the output where the metaData is situated. 
    "txHex": "",//string; tx original data
    "size": 0,//int; The size of tx
    "vins": [],//[]object; The inputs of the tx.
    "vouts": [],//[]object; The outputs of the tx.
    "scriptHex": "",//string; metaData script data
    "overSizeLimit": 0,//int; The size of the tx exceeds the limit
    "parts": [],//[]string; The parts of the tx.
    "address": "",//string; The address corresponding to PuklicKey of current node
    "publicKey": "",//string; PuklicKey used of current node
    "parentChainFlag": "",//string; 
    "parentTxId": "",//string; The txId of parent node
    "parentNodeName": "",//string; 
    "metaId": "",//string; The flag of "metaid"
    "data": nil,//string; Data content of current node
    "encrypt": "",//string; The encryption method used for the data.
    "version": "",//string; The version of the data.
    "dataType": "",//string; The type of the data.
    "encoding": "",//string; The encoding method used for the data.
    "params": "",//string;
    "blockHeight": 0,//int; The height of the block where the tx is situated.
    "metaBlockHeight": 0,//int; 
    "confirmState": 0,//int; The state of the tx's confirmation.
    "isValid": true,//boolean; The validity of the metaid data.
    "reason": "",//string; The reason for the invalidity of the metaid data.
    "isNew": true,//boolean; This determines whether the MetaAid data represents the latest version.
    "fee": 0,//int; The fee of the tx.
    "timestamp": 0,//int; The timestamp of the tx.
    "timestampBlock": 0,//int; The timestamp of the block where the tx is situated.
    "timestampWitness": 0,//int; The timestamp of the witness where the tx is situated.
}
```



