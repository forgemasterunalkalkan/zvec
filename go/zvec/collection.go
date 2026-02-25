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

// Collection represents an opened collection in zvec.
//
// A Collection provides methods for data definition (DDL), data manipulation (DML),
// and querying (DQL). It is obtained via CreateAndOpen or Open.
type Collection interface {
	// Path returns the filesystem path of the collection.
	Path() string

	// Schema returns the schema defining the structure of the collection.
	Schema() *CollectionSchema

	// Stats returns runtime statistics about the collection.
	Stats() (*CollectionStats, error)

	// Options returns the options used to open the collection.
	Options() CollectionOption

	// Destroy permanently deletes the collection from disk.
	// This operation is irreversible.
	Destroy() error

	// Flush forces all pending writes to disk.
	Flush() error

	// CreateIndex creates an index on a field.
	//
	// Vector index types (HNSW, IVF, FLAT) can only be applied to vector fields.
	// Inverted index (InvertIndexParams) is for scalar fields.
	CreateIndex(fieldName string, params IndexParams, option ...CreateIndexOption) error

	// DropIndex removes the index from a field.
	DropIndex(fieldName string) error

	// Optimize optimizes the collection (e.g., merge segments, rebuild index).
	Optimize(option ...OptimizeOption) error

	// AddColumn adds a new column to the collection.
	//
	// The column is populated using the provided expression for existing documents.
	AddColumn(schema *FieldSchema, expression string, option ...AddColumnOption) error

	// DropColumn removes a column from the collection.
	DropColumn(fieldName string) error

	// AlterColumn renames a column and/or updates its schema.
	//
	// If newName is empty, no rename occurs. If newSchema is nil, only renaming
	// is performed. Only scalar numeric columns are supported.
	AlterColumn(oldName string, newName string, newSchema *FieldSchema, option ...AlterColumnOption) error

	// Insert inserts new documents into the collection.
	// Documents must have unique IDs and conform to the schema.
	Insert(docs []*Doc) ([]Status, error)

	// Upsert inserts new documents or updates existing ones by ID.
	Upsert(docs []*Doc) ([]Status, error)

	// Update updates existing documents by ID.
	// Only specified fields are updated; others remain unchanged.
	Update(docs []*Doc) ([]Status, error)

	// Delete deletes documents by their IDs.
	Delete(ids []string) ([]Status, error)

	// DeleteByFilter deletes documents matching a filter expression.
	DeleteByFilter(filter string) error

	// Query performs vector similarity search with optional filtering.
	Query(option QueryOption) ([]*Doc, error)

	// Fetch retrieves documents by their IDs.
	// Missing IDs are omitted from the result.
	Fetch(ids []string) (map[string]*Doc, error)
}
