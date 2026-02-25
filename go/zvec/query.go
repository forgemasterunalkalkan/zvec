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

// VectorQuery represents a vector search query for a specific field.
//
// Exactly one of ID or Vector should be provided. If both are given,
// the behavior is implementation-defined.
type VectorQuery struct {
	// FieldName is the name of the vector field to query.
	FieldName string

	// ID is the document ID to fetch the vector from (for query-by-ID).
	// If set, Vector should be nil.
	ID string

	// Vector is the explicit query vector.
	// If set, ID should be empty.
	Vector interface{}

	// Param is the optional index-specific query parameters.
	Param QueryParams
}

// HasID reports whether the query uses a document ID.
func (q *VectorQuery) HasID() bool {
	return q.ID != ""
}

// HasVector reports whether the query contains an explicit vector.
func (q *VectorQuery) HasVector() bool {
	return q.Vector != nil
}

// QueryOption configures a vector similarity search.
type QueryOption struct {
	// Vectors is the list of vector queries to execute.
	Vectors []VectorQuery

	// TopK is the number of nearest neighbors to return. Default: 10.
	TopK int

	// Filter is an optional boolean expression to pre-filter candidates.
	Filter string

	// IncludeVector indicates whether to include vector data in results.
	IncludeVector bool

	// OutputFields specifies which scalar fields to include in results.
	// If nil, all fields are returned.
	OutputFields []string
}

// DefaultQueryOption returns a QueryOption with default values.
func DefaultQueryOption() QueryOption {
	return QueryOption{
		TopK: 10,
	}
}
