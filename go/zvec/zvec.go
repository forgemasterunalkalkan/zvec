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

// Package zvec provides Go bindings for the zvec vector database library.
//
// Zvec is an in-process vector database — lightweight, fast, and designed to
// embed directly into applications. It supports dense and sparse vectors,
// hybrid search with structured filters, and multiple index types.
//
// # Quick Start
//
// To use zvec in a Go application, first initialize the library, then create
// or open a collection:
//
//	import "github.com/alibaba/zvec/go/zvec"
//
//	// Initialize with defaults
//	if err := zvec.Init(nil); err != nil {
//	    log.Fatal(err)
//	}
//
//	// Define schema
//	schema := &zvec.CollectionSchema{
//	    Name: "example",
//	    Vectors: []zvec.VectorSchema{
//	        {
//	            Name:      "embedding",
//	            DataType:  zvec.DataTypeVectorFP32,
//	            Dimension: 4,
//	        },
//	    },
//	}
//
//	// Create and open collection
//	collection, err := zvec.CreateAndOpen("./zvec_example", schema, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Insert documents
//	docs := []*zvec.Doc{
//	    zvec.NewDoc("doc_1").SetVector("embedding", []float32{0.1, 0.2, 0.3, 0.4}),
//	    zvec.NewDoc("doc_2").SetVector("embedding", []float32{0.2, 0.3, 0.4, 0.1}),
//	}
//	statuses, err := collection.Insert(docs)
//
//	// Search
//	results, err := collection.Query(zvec.QueryOption{
//	    Vectors: []zvec.VectorQuery{
//	        {FieldName: "embedding", Vector: []float32{0.4, 0.3, 0.3, 0.1}},
//	    },
//	    TopK: 10,
//	})
//
// # Architecture
//
// The Go client communicates with the zvec C++ library through a C API layer.
// The library must be built and linked for the target platform before the Go
// client can be used. See the project documentation for build instructions.
//
// # Thread Safety
//
// Collection instances are safe for concurrent use from multiple goroutines.
package zvec
