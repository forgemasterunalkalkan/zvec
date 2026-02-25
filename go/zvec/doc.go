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

package zvec

// Doc represents a document in a collection.
//
// A Doc has a primary key (ID), an optional relevance score,
// scalar fields, and vector embeddings.
type Doc struct {
	// ID is the unique identifier of the document.
	ID string

	// Score is the relevance score from search (0 if not from a query).
	Score float32

	// Fields contains scalar metadata fields (e.g., title, timestamp).
	Fields map[string]interface{}

	// Vectors contains named vector embeddings associated with the document.
	Vectors map[string]interface{}
}

// NewDoc creates a new Doc with the given ID.
func NewDoc(id string) *Doc {
	return &Doc{
		ID:      id,
		Fields:  make(map[string]interface{}),
		Vectors: make(map[string]interface{}),
	}
}

// SetField sets a scalar field value on the document.
func (d *Doc) SetField(name string, value interface{}) *Doc {
	if d.Fields == nil {
		d.Fields = make(map[string]interface{})
	}
	d.Fields[name] = value
	return d
}

// SetVector sets a vector field value on the document.
// The value should be a slice of float32, float64, int8, etc.
func (d *Doc) SetVector(name string, value interface{}) *Doc {
	if d.Vectors == nil {
		d.Vectors = make(map[string]interface{})
	}
	d.Vectors[name] = value
	return d
}

// HasField reports whether the document contains a scalar field with the given name.
func (d *Doc) HasField(name string) bool {
	_, ok := d.Fields[name]
	return ok
}

// HasVector reports whether the document contains a vector with the given name.
func (d *Doc) HasVector(name string) bool {
	_, ok := d.Vectors[name]
	return ok
}

// GetField returns the value of a scalar field, or nil if not found.
func (d *Doc) GetField(name string) interface{} {
	if d.Fields == nil {
		return nil
	}
	return d.Fields[name]
}

// GetVector returns the value of a vector field, or nil if not found.
func (d *Doc) GetVector(name string) interface{} {
	if d.Vectors == nil {
		return nil
	}
	return d.Vectors[name]
}

// FieldNames returns the names of all scalar fields.
func (d *Doc) FieldNames() []string {
	names := make([]string, 0, len(d.Fields))
	for name := range d.Fields {
		names = append(names, name)
	}
	return names
}

// VectorNames returns the names of all vector fields.
func (d *Doc) VectorNames() []string {
	names := make([]string, 0, len(d.Vectors))
	for name := range d.Vectors {
		names = append(names, name)
	}
	return names
}
