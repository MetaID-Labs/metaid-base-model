package model

import (
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"metaid-base-model/major"
)

const (
	SizeLimit = 2 * 1024 * 1024

	Unconfirmed int64 = 0
	Confirmed   int64 = 1
)

type MetaIDModel struct {
	Id               int64        `json:"id" bson:"_id"`
	ChainFlag        string       `json:"chainFlag" bson:"chainFlag"`
	NodeId           string       `json:"nodeId" bson:"nodeId"`
	MetanetId        string       `json:"metanetId" bson:"metanetId"`
	RootTxId         string       `json:"rootTxId" bson:"rootTxId"`
	TxId             string       `json:"txId" bson:"txId"`
	OutputIndex      int64        `json:"outputIndex" bson:"outputIndex"`
	TxHex            string       `json:"txHex" bson:"txHex"`
	Size             int64        `json:"size" bson:"size"`
	Vins             []*MetaTxIn  `json:"vins" bson:"vins"`
	Vouts            []*MetaTxOut `json:"vouts" bson:"vouts"`
	ScriptHex        string       `json:"scriptHex" bson:"scriptHex"`
	OverSizeLimit    int64        `json:"overSizeLimit" bson:"overSizeLimit"`
	Parts            []string     `json:"parts" bson:"parts"`
	Address          string       `json:"address" bson:"address"`
	PublicKey        string       `json:"publicKey" bson:"publicKey"`
	ParentChainFlag  string       `json:"parentChainFlag" bson:"parentChainFlag"`
	ParentTxId       string       `json:"parentTxId" bson:"parentTxId"`
	ParentNodeName   string       `json:"parentNodeName" bson:"parentNodeName"`
	MetaId           string       `json:"metaId" bson:"metaId"`
	NodeName         string       `json:"nodeName" bson:"nodeName"`
	Data             interface{}  `json:"data" bson:"data"`
	Encrypt          string       `json:"encrypt" bson:"encrypt"`
	Version          string       `json:"version" bson:"version"`
	DataType         string       `json:"dataType" bson:"dataType"`
	Encoding         string       `json:"encoding" bson:"encoding"`
	Params           []string     `json:"params" bson:"params"`
	BlockHeight      int64        `json:"blockHeight" bson:"blockHeight"`
	MetaBlockHeight  int64        `json:"metaBlockHeight" bson:"metaBlockHeight"`
	ConfirmState     int64        `json:"confirmState" bson:"confirmState"`
	IsValid          bool         `json:"isValid" bson:"isValid"`
	Reason           string       `json:"reason" bson:"reason"`
	IsNew            bool         `json:"isNew" bson:"isNew"`
	Fee              uint64       `json:"fee" bson:"fee"`
	Timestamp        int64        `json:"timestamp" bson:"timestamp"`
	TimestampBlock   int64        `json:"timestampBlock" bson:"timestampBlock"`
	TimestampWitness int64        `json:"timestampWitness" bson:"timestampWitness"`
	CreateTime       int64        `json:"createTime" bson:"createTime"`
	UpdateTime       int64        `json:"updateTime" bson:"updateTime"`
	State            DataState    `json:"state" bson:"state"`
}

type MetaTxIn struct {
	OutTxId     string `json:"outTxId" bson:"outTxId"` //
	Index       uint64 `json:"index" bson:"index"`
	Value       uint64 `json:"value" bson:"value"`
	Address     string `json:"address" bson:"address"`
	PublicKey   string `json:"publicKey" bson:"publicKey"`
	OutTxScript string `json:"outTxScript" bson:"outTxScript"`
}

type MetaTxOut struct {
	TxId         string `json:"txId" bson:"txId"`
	Index        uint64 `json:"index" bson:"index"`
	Address      string `json:"address" bson:"address"`
	PublicKey    string `json:"publicKey" bson:"publicKey"`
	Value        uint64 `json:"value" bson:"value"`
	ScriptPubKey string `json:"scriptPubKey" bson:"scriptPubKey"`
	Type         string `json:"type" bson:"type"`
}

func (s MetaIDModel) getCollection() string {
	return "metaid_data"
}

func (s MetaIDModel) getDB() string {
	return "MetaIDData"
}

// DB - read
func (s MetaIDModel) GetReadDB() (*mongo.Collection, error) {
	mongoDB, err := major.GetMetaIdReadDB()
	if err != nil {
		return nil, err
	}
	collection := mongoDB.Database(s.getDB()).Collection(s.getCollection())
	if collection == nil {
		return nil, errors.New("db connect error")
	}
	return collection, nil
}

// DB - read
func (s MetaIDModel) GetWriteDB() (*mongo.Collection, error) {
	mongoDB, err := major.GetMetaIdReadDB()
	if err != nil {
		return nil, err
	}
	collection := mongoDB.Database(s.getDB()).Collection(s.getCollection())
	if collection == nil {
		return nil, errors.New("db connect error")
	}
	return collection, nil
}
