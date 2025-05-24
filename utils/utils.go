package utils

import (
	"encoding/json"

	"google.golang.org/protobuf/types/known/structpb"
)

func TransformRawData2StructpbStruct(rawData map[string]interface{}) ( *structpb.Struct,error) {
	
	//   // 1. 构建原始数据结构
	//   rawData := map[string]interface{}{
    //     "tags":     tgSlice,
    //     "real_pgs": int(pgNums),
    // }
	
	// 2. 通过JSON中转
	 jsonBytes, _ := json.Marshal(rawData)
	 var jsonMap map[string]interface{}
	 json.Unmarshal(jsonBytes, &jsonMap)
 
	 // 3. 转换为protobuf结构
	 return structpb.NewStruct(jsonMap)
}

func ConvertInt32ToIntSlice(int32Slice []int32) []int {
    intSlice := make([]int, len(int32Slice))
    for i, v := range int32Slice {
        intSlice[i] = int(v)
    }
    return intSlice
}