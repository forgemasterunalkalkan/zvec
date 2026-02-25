// Copyright 2025-present the zvec project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build ignore

// Example demonstrates basic usage of the zvec Go client.
//
// This example creates a collection with a vector field, inserts documents,
// and performs a similarity search.
//
// Prerequisites:
//   - zvec C library must be built and installed
//   - Build with: go build -tags zvec_cgo example.go
package main

import (
	"fmt"
	"log"

	"github.com/alibaba/zvec/go/zvec"
)

func main() {
	// Initialize zvec with defaults
	if err := zvec.Init(nil); err != nil {
		log.Fatal(err)
	}

	// Define collection schema
	schema := &zvec.CollectionSchema{
		Name: "example",
		Vectors: []zvec.VectorSchema{
			{
				Name:      "embedding",
				DataType:  zvec.DataTypeVectorFP32,
				Dimension: 4,
			},
		},
	}

	// Create and open collection
	collection, err := zvec.CreateAndOpen("./zvec_example", schema, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Insert documents
	docs := []*zvec.Doc{
		zvec.NewDoc("doc_1").SetVector("embedding", []float32{0.1, 0.2, 0.3, 0.4}),
		zvec.NewDoc("doc_2").SetVector("embedding", []float32{0.2, 0.3, 0.4, 0.1}),
	}

	statuses, err := collection.Insert(docs)
	if err != nil {
		log.Fatal(err)
	}
	for i, s := range statuses {
		if !s.OK() {
			log.Printf("Insert doc %d failed: %s", i, s.Error())
		}
	}

	// Search by vector similarity
	results, err := collection.Query(zvec.QueryOption{
		Vectors: []zvec.VectorQuery{
			{
				FieldName: "embedding",
				Vector:    []float32{0.4, 0.3, 0.3, 0.1},
			},
		},
		TopK: 10,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Print results
	for _, doc := range results {
		fmt.Printf("ID: %s, Score: %f\n", doc.ID, doc.Score)
	}
}
