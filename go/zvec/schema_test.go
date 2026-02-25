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

import "testing"

func TestFieldSchemaValidate(t *testing.T) {
	tests := []struct {
		name    string
		schema  FieldSchema
		wantErr bool
	}{
		{
			name:    "valid scalar field",
			schema:  FieldSchema{Name: "age", DataType: DataTypeInt32},
			wantErr: false,
		},
		{
			name:    "valid string field",
			schema:  FieldSchema{Name: "title", DataType: DataTypeString, Nullable: true},
			wantErr: false,
		},
		{
			name:    "empty name",
			schema:  FieldSchema{Name: "", DataType: DataTypeInt32},
			wantErr: true,
		},
		{
			name:    "vector data type in field schema",
			schema:  FieldSchema{Name: "vec", DataType: DataTypeVectorFP32},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.schema.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("FieldSchema.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVectorSchemaValidate(t *testing.T) {
	tests := []struct {
		name    string
		schema  VectorSchema
		wantErr bool
	}{
		{
			name:    "valid vector field",
			schema:  VectorSchema{Name: "embedding", DataType: DataTypeVectorFP32, Dimension: 128},
			wantErr: false,
		},
		{
			name:    "empty name",
			schema:  VectorSchema{Name: "", DataType: DataTypeVectorFP32, Dimension: 128},
			wantErr: true,
		},
		{
			name:    "non-vector data type",
			schema:  VectorSchema{Name: "vec", DataType: DataTypeString, Dimension: 128},
			wantErr: true,
		},
		{
			name:    "negative dimension",
			schema:  VectorSchema{Name: "vec", DataType: DataTypeVectorFP32, Dimension: -1},
			wantErr: true,
		},
		{
			name:    "zero dimension (valid for sparse)",
			schema:  VectorSchema{Name: "sparse", DataType: DataTypeSparseVectorFP32, Dimension: 0},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.schema.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("VectorSchema.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCollectionSchemaValidate(t *testing.T) {
	tests := []struct {
		name    string
		schema  CollectionSchema
		wantErr bool
	}{
		{
			name: "valid schema",
			schema: CollectionSchema{
				Name: "test_collection",
				Fields: []FieldSchema{
					{Name: "age", DataType: DataTypeInt32},
				},
				Vectors: []VectorSchema{
					{Name: "embedding", DataType: DataTypeVectorFP32, Dimension: 4},
				},
			},
			wantErr: false,
		},
		{
			name:    "empty name",
			schema:  CollectionSchema{Name: ""},
			wantErr: true,
		},
		{
			name: "duplicate field names",
			schema: CollectionSchema{
				Name: "test",
				Fields: []FieldSchema{
					{Name: "field1", DataType: DataTypeInt32},
					{Name: "field1", DataType: DataTypeString},
				},
			},
			wantErr: true,
		},
		{
			name: "duplicate field and vector names",
			schema: CollectionSchema{
				Name: "test",
				Fields: []FieldSchema{
					{Name: "data", DataType: DataTypeInt32},
				},
				Vectors: []VectorSchema{
					{Name: "data", DataType: DataTypeVectorFP32, Dimension: 4},
				},
			},
			wantErr: true,
		},
		{
			name: "vectors only schema",
			schema: CollectionSchema{
				Name: "vectors_only",
				Vectors: []VectorSchema{
					{Name: "emb", DataType: DataTypeVectorFP32, Dimension: 128},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.schema.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("CollectionSchema.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCollectionSchemaGetField(t *testing.T) {
	schema := CollectionSchema{
		Name: "test",
		Fields: []FieldSchema{
			{Name: "age", DataType: DataTypeInt32},
			{Name: "title", DataType: DataTypeString},
		},
		Vectors: []VectorSchema{
			{Name: "embedding", DataType: DataTypeVectorFP32, Dimension: 4},
		},
	}

	if f := schema.GetField("age"); f == nil {
		t.Error("GetField('age') should return non-nil")
	} else if f.DataType != DataTypeInt32 {
		t.Errorf("GetField('age').DataType = %v, want %v", f.DataType, DataTypeInt32)
	}

	if f := schema.GetField("nonexistent"); f != nil {
		t.Error("GetField('nonexistent') should return nil")
	}

	if v := schema.GetVector("embedding"); v == nil {
		t.Error("GetVector('embedding') should return non-nil")
	} else if v.Dimension != 4 {
		t.Errorf("GetVector('embedding').Dimension = %d, want 4", v.Dimension)
	}

	if v := schema.GetVector("nonexistent"); v != nil {
		t.Error("GetVector('nonexistent') should return nil")
	}
}

func TestCollectionSchemaAllFieldNames(t *testing.T) {
	schema := CollectionSchema{
		Name: "test",
		Fields: []FieldSchema{
			{Name: "age", DataType: DataTypeInt32},
		},
		Vectors: []VectorSchema{
			{Name: "emb", DataType: DataTypeVectorFP32, Dimension: 4},
		},
	}

	names := schema.AllFieldNames()
	if len(names) != 2 {
		t.Errorf("AllFieldNames() length = %d, want 2", len(names))
	}
}

func TestCollectionStats(t *testing.T) {
	stats := CollectionStats{
		DocCount: 1000,
		IndexCompleteness: map[string]float32{
			"embedding": 0.95,
		},
	}
	if stats.DocCount != 1000 {
		t.Errorf("DocCount = %d, want 1000", stats.DocCount)
	}
	if stats.IndexCompleteness["embedding"] != 0.95 {
		t.Errorf("IndexCompleteness['embedding'] = %f, want 0.95", stats.IndexCompleteness["embedding"])
	}
}
