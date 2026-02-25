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

import (
	"errors"
	"fmt"
)

// FieldSchema defines a scalar (non-vector) field in a collection schema.
type FieldSchema struct {
	// Name is the field name. Must be unique within the collection.
	Name string

	// DataType is the data type of the field (e.g., DataTypeInt64, DataTypeString).
	DataType DataType

	// Nullable indicates whether the field can contain null values.
	Nullable bool

	// IndexParam is the optional inverted index configuration for this field.
	IndexParam *InvertIndexParams
}

// Validate checks the FieldSchema for invalid values.
func (f *FieldSchema) Validate() error {
	if f.Name == "" {
		return errors.New("field name must not be empty")
	}
	if f.DataType.IsVector() {
		return fmt.Errorf("field %q has vector data type %s; use VectorSchema instead", f.Name, f.DataType)
	}
	return nil
}

// VectorSchema defines a vector field in a collection schema.
type VectorSchema struct {
	// Name is the vector field name. Must be unique within the collection.
	Name string

	// DataType is the vector data type (e.g., DataTypeVectorFP32).
	DataType DataType

	// Dimension is the dimensionality of the vector.
	// Must be > 0 for dense vectors; may be 0 for sparse vectors.
	Dimension int

	// IndexParam is the index configuration for this vector field.
	// If nil, a default FlatIndexParams with MetricTypeIP is used.
	IndexParam VectorIndexParams
}

// Validate checks the VectorSchema for invalid values.
func (v *VectorSchema) Validate() error {
	if v.Name == "" {
		return errors.New("vector field name must not be empty")
	}
	if !v.DataType.IsVector() {
		return fmt.Errorf("vector field %q has non-vector data type %s", v.Name, v.DataType)
	}
	if v.Dimension < 0 {
		return fmt.Errorf("vector field %q dimension must be >= 0, got %d", v.Name, v.Dimension)
	}
	return nil
}

// CollectionSchema defines the structure of a collection.
type CollectionSchema struct {
	// Name is the collection name.
	Name string

	// Fields is the list of scalar field definitions.
	Fields []FieldSchema

	// Vectors is the list of vector field definitions.
	Vectors []VectorSchema
}

// Validate checks the CollectionSchema for invalid values.
func (s *CollectionSchema) Validate() error {
	if s.Name == "" {
		return errors.New("collection name must not be empty")
	}

	names := make(map[string]bool)
	for i := range s.Fields {
		if err := s.Fields[i].Validate(); err != nil {
			return fmt.Errorf("field[%d]: %w", i, err)
		}
		if names[s.Fields[i].Name] {
			return fmt.Errorf("duplicate field name %q", s.Fields[i].Name)
		}
		names[s.Fields[i].Name] = true
	}

	for i := range s.Vectors {
		if err := s.Vectors[i].Validate(); err != nil {
			return fmt.Errorf("vector[%d]: %w", i, err)
		}
		if names[s.Vectors[i].Name] {
			return fmt.Errorf("duplicate field name %q", s.Vectors[i].Name)
		}
		names[s.Vectors[i].Name] = true
	}

	return nil
}

// GetField returns the scalar field with the given name, or nil if not found.
func (s *CollectionSchema) GetField(name string) *FieldSchema {
	for i := range s.Fields {
		if s.Fields[i].Name == name {
			return &s.Fields[i]
		}
	}
	return nil
}

// GetVector returns the vector field with the given name, or nil if not found.
func (s *CollectionSchema) GetVector(name string) *VectorSchema {
	for i := range s.Vectors {
		if s.Vectors[i].Name == name {
			return &s.Vectors[i]
		}
	}
	return nil
}

// AllFieldNames returns the names of all fields (scalar and vector).
func (s *CollectionSchema) AllFieldNames() []string {
	names := make([]string, 0, len(s.Fields)+len(s.Vectors))
	for _, f := range s.Fields {
		names = append(names, f.Name)
	}
	for _, v := range s.Vectors {
		names = append(names, v.Name)
	}
	return names
}

// CollectionStats contains runtime statistics about a collection.
type CollectionStats struct {
	// DocCount is the total number of documents in the collection.
	DocCount uint64

	// IndexCompleteness maps column names to their index build completeness (0.0 to 1.0).
	IndexCompleteness map[string]float32
}
