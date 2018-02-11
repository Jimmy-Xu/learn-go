package main

import (
	"fmt"
	"log"
	"encoding/json"
	"io/ioutil"
	"flag"

	"github.com/jimmy-xu/learn-go/gojsonschema/config"
	"github.com/ghodss/yaml"
	"github.com/xeipuuv/gojsonschema"

	"k8s.io/api/core/v1"
)

func main() {

	file := flag.String("file","", "pod yaml file name")
	flag.Parse()

	//schemaLoader := gojsonschema.NewReferenceLoader("file:///Users/xjimmy/gopath/src/github.com/jimmy-xu/learn-go/gojsonschema/pod-schema.json")

	var (
		schemaRaw interface{}
		//schema    map[string]interface{}
	)
	err := json.Unmarshal([]byte(config.SchemaV1), &schemaRaw)
	if err != nil {
		log.Fatalf("convert schemaV1 to json failed, error:%v",err)
	}

	//schema = schemaRaw.(map[string]interface{})
	var schemaLoader gojsonschema.JSONLoader
	schemaLoader = gojsonschema.NewGoLoader(schemaRaw)

	//documentLoader := gojsonschema.NewReferenceLoader("file:///Users/xjimmy/gopath/src/github.com/jimmy-xu/learn-go/gojsonschema/pod.json")
    documentLoader := gojsonschema.NewStringLoader(`{
  "kind": "Pod",
  "apiVersion": "v1",
  "metadata": {
    "name": "test-nginx",
    "creationTimestamp": null
  },
  "spec": {
    "containers": [
      {
        "name": "test-nginx",
        "image": "nginx:1.7.9",
        "ports": [
          {
            "containerPort": 80
          }
        ],
        "resources": {}
      }
    ]
  },
  "status": {}
}`)

	//prop1 := schema["properties"].(map[string]interface{})
	//metadata := prop1["metadata"].(map[string]interface{})
	//prop2 := metadata["properties"].(map[string]interface{})
	//name := prop2["name"].(map[string]interface{})
	//nameType := name["type"].(string)
	//fmt.Printf("name type:%v\n",  nameType)

	fmt.Println("validate documentLoader")
	validate(schemaLoader, documentLoader)


	/////////////////////////////////////////////////
	schemaLoader1 := gojsonschema.NewStringLoader(config.SchemaV1)


	//read yaml
	ymlData, err := ioutil.ReadFile(*file)
	if err != nil {
		log.Fatalf("read yaml file %v error: %v", *file, err)
	}

	//conver to json
	jsonData, err := yaml.YAMLToJSON(ymlData)
	if err != nil {
		log.Fatalf("convert yaml to json error: %v", err)
	}

	var pod v1.Pod
	json.Unmarshal([]byte(jsonData), &pod)

	buf, err := json.MarshalIndent(pod, "" ,"  ")
	fmt.Printf("jsonData:\n%v", string(buf))



	documentLoader2 := gojsonschema.NewGoLoader(pod)
	fmt.Println("validate documentLoader2")
	validate(schemaLoader1, documentLoader2)


	fmt.Println("validate documentLoader3")
	documentLoader3 := gojsonschema.NewStringLoader(`{
"kind": "Pod",
"apiVersion": "v1",
"metadata": {
"name": "test-nginx",
"creationTimestamp": null
},
"spec": {
"containers": [
  {
	"name": "test-nginx",
	"image": "nginx:1.7.9",
	"ports": [
	  {
		"containerPort": 80
	  }
	],
	"resources": {}
  }
]
},
"status": {}
}`)

	validate(schemaLoader1, documentLoader3)
}
func validate(schemaLoader gojsonschema.JSONLoader, documentLoader gojsonschema.JSONLoader) {
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		log.Fatalf("validate error:%v", err)
	}
	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}
