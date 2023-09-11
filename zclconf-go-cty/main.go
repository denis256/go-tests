package main

import (
	"encoding/json"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
	ctyjson "github.com/zclconf/go-cty/cty/json"

	"log"
)

func main() {

	valMap := map[string]cty.Value{}
	valMap["test"] = cty.NullVal(cty.DynamicPseudoType)
	valMap["test2"] = cty.UnknownVal(cty.DynamicPseudoType)

	valMapAsCty, err := gocty.ToCtyValue(valMap, generateTypeFromValuesMap(valMap))
	if err != nil {
		log.Fatal(err)
	}

	jsonBytes, err := ctyjson.Marshal(valMapAsCty, cty.DynamicPseudoType)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(jsonBytes))

	var result map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &result); err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v\n", result)
}

func generateTypeFromValuesMap(valMap map[string]cty.Value) cty.Type {
	outType := map[string]cty.Type{}
	for k, v := range valMap {
		outType[k] = v.Type()
	}
	return cty.Object(outType)
}
