package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
)

type DataState int64

const (
	STATE_EXIST   DataState = 1
	STATE_DELETED DataState = 2
)

// ContentType
const (
	TextPlain   = "text/plain" //content
	ImageUrl    = "image/url"
	ImageJpg    = "image/jpg"
	ImageJpeg   = "image/jpeg"
	ImagePng    = "image/png"
	ImageGif    = "image/gif"
	ImageWebp   = "image/webp"
	AudioWav    = "audio/wav"
	Transaction = "transaction"
	BsvProtocol = "bsvProtocol"
)

func ToMapData(data interface{}) interface{} {
	if ValueOf(data).Kind() == reflect.Map {
		return data
	} else if ValueOf(data).Kind() == reflect.Slice {
		switch data.(type) {
		case primitive.D:
			realData := make(map[string]interface{})
			listMap := data.(primitive.D)
			for _, v := range listMap {
				realData[v.Key] = ToMapData(v.Value)
			}
			return realData
		case primitive.A:
			realList := make([]interface{}, 0)
			list := data.(primitive.A)
			for _, v := range list {
				realList = append(realList, ToMapData(v))
			}
			return realList
		default:
			return data
		}
	} else {
		return data
	}
}

func ValueOf(data interface{}) reflect.Value {
	vof := reflect.ValueOf(data)
	if vof.Kind() == reflect.Ptr {
		vof = vof.Elem()
	}
	return vof
}
