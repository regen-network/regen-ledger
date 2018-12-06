package data

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

//type MsgRegisterSchema struct {
//	// TODO figure out how to do naming and version - essentially schemas which are upgrades of previous versions and maintain backwards compatibility
//	// Name string
//	// Version string
//	Schema string
//}

type MsgStoreData struct {
	//SchemaRef string
	Data []byte
	Signer sdk.AccAddress
}

//type MsgStoreDataPointer struct {
//	Hash string
//	HashAlg string
//	Url string
//	//SchemaRef string
//	PartialData string
//}


//func NewMsgRegisterSchema(schema string) MsgRegisterSchema {
//	return MsgRegisterSchema {
//		Schema: schema,
//	}
//}
//
//func (msg MsgRegisterSchema) Route() string { return "data" }
//
//func (msg MsgRegisterSchema) Type() string { return "register_schema" }
//
//func (msg MsgRegisterSchema) ValidateBasic() sdk.Error {
//	if msg.Schema == ""  {
//		return sdk.ErrUnknownRequest("Schema cannot be empty")
//	}
//
//	loader := gojsonschema.NewStringLoader(msg.Schema)
//	sl := gojsonschema.NewSchemaLoader()
//	_, err := sl.Compile(loader)
//
//	if err != nil {
//		return sdk.ErrUnknownRequest(err.Error())
//	}
//	return nil
//}
//
//
//func (msg MsgRegisterSchema) GetSignBytes() []byte {
//	b, err := json.Marshal(msg)
//	if err != nil {
//		panic(err)
//	}
//	return sdk.MustSortJSON(b)
//}
//
//func (msg MsgRegisterSchema) GetSigners() []sdk.AccAddress {
//	return nil
//}

func NewMsgStoreData(data []byte, signer sdk.AccAddress) MsgStoreData {
	return MsgStoreData{
        Data:data,
	}
}

func (msg MsgStoreData) Route() string { return "data" }

func (msg MsgStoreData) Type() string { return "store_data" }

func (msg MsgStoreData) ValidateBasic() sdk.Error {
	//if msg.SchemaRef == "" {
	//	return sdk.ErrUnknownRequest("SchemaRef cannot be empty")
	//}
	if len(msg.Data) == 0 {
		return sdk.ErrUnknownRequest("Data cannot be empty")
	}
	// Schema gets looked up and validated on the blockchain
	return nil
}


func (msg MsgStoreData) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgStoreData) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}

//func (msg MsgStoreDataPointer) Route() string { return "data" }
//
//func (msg MsgStoreDataPointer) Type() string { return "store_data_pointer" }
//
//func (msg MsgStoreDataPointer) ValidateBasic() sdk.Error {
//	if msg.Hash == "" {
//		return sdk.ErrUnknownRequest("Hash cannot be empty")
//	}
//	if msg.HashAlg == "" {
//		return sdk.ErrUnknownRequest("HashAlg cannot be empty")
//	}
//	if msg.Url == "" {
//		return sdk.ErrUnknownRequest("Url cannot be empty")
//	}
//	// Schema and PartialData are optional and not validated
//	return nil
//}
//
//func (msg MsgStoreDataPointer) GetSignBytes() []byte {
//	b, err := json.Marshal(msg)
//	if err != nil {
//		panic(err)
//	}
//	return sdk.MustSortJSON(b)
//}
//
//func (msg MsgStoreDataPointer) GetSigners() []sdk.AccAddress {
//}
