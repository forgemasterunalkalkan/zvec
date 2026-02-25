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

//go:build cgo && zvec_cgo

package zvec

// #cgo CFLAGS: -I${SRCDIR}/internal/capi
// #cgo LDFLAGS: -lzvec
// #include "zvec_c_api.h"
// #include <stdlib.h>
import "C"

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

// cgoCollection implements the Collection interface using the C API.
type cgoCollection struct {
	handle C.zvec_collection_t
	schema *CollectionSchema
	path   string
}

func statusFromC(s C.zvec_status_t) Status {
	if s.code == 0 {
		return StatusOK()
	}
	msg := ""
	if s.message != nil {
		msg = C.GoString(s.message)
	}
	return NewStatus(StatusCode(s.code), msg)
}

func statusToError(s C.zvec_status_t) error {
	return statusFromC(s).Err()
}

// Init initializes the zvec library with optional configuration.
// Pass nil for default configuration.
// This function can only be called once.
func Init(config *InitConfig) error {
	var cJSON *C.char
	if config != nil {
		if err := config.Validate(); err != nil {
			return err
		}
		data, err := json.Marshal(config)
		if err != nil {
			return fmt.Errorf("failed to marshal init config: %w", err)
		}
		cJSON = C.CString(string(data))
		defer C.free(unsafe.Pointer(cJSON))
	}
	return statusToError(C.zvec_init(cJSON))
}

// CreateAndOpen creates a new collection and opens it for use.
func CreateAndOpen(path string, schema *CollectionSchema, option *CollectionOption) (Collection, error) {
	if schema == nil {
		return nil, fmt.Errorf("schema must not be nil")
	}
	if err := schema.Validate(); err != nil {
		return nil, fmt.Errorf("invalid schema: %w", err)
	}

	schemaJSON, err := json.Marshal(schema)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal schema: %w", err)
	}

	var optionJSON []byte
	if option != nil {
		optionJSON, err = json.Marshal(option)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal options: %w", err)
		}
	}

	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	cSchema := C.CString(string(schemaJSON))
	defer C.free(unsafe.Pointer(cSchema))

	var cOption *C.char
	if optionJSON != nil {
		cOption = C.CString(string(optionJSON))
		defer C.free(unsafe.Pointer(cOption))
	}

	var handle C.zvec_collection_t
	status := C.zvec_collection_create_and_open(cPath, cSchema, cOption, &handle)
	if err := statusToError(status); err != nil {
		return nil, err
	}

	return &cgoCollection{
		handle: handle,
		schema: schema,
		path:   path,
	}, nil
}

// Open opens an existing collection from disk.
func Open(path string, option *CollectionOption) (Collection, error) {
	var optionJSON []byte
	var err error
	if option != nil {
		optionJSON, err = json.Marshal(option)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal options: %w", err)
		}
	}

	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	var cOption *C.char
	if optionJSON != nil {
		cOption = C.CString(string(optionJSON))
		defer C.free(unsafe.Pointer(cOption))
	}

	var handle C.zvec_collection_t
	status := C.zvec_collection_open(cPath, cOption, &handle)
	if err := statusToError(status); err != nil {
		return nil, err
	}

	// Retrieve schema from the opened collection
	var cSchemaJSON *C.char
	status = C.zvec_collection_schema(handle, &cSchemaJSON)
	if err := statusToError(status); err != nil {
		C.zvec_collection_close(handle)
		return nil, fmt.Errorf("failed to get schema: %w", err)
	}
	defer C.zvec_free_string(cSchemaJSON)

	var schema CollectionSchema
	if err := json.Unmarshal([]byte(C.GoString(cSchemaJSON)), &schema); err != nil {
		C.zvec_collection_close(handle)
		return nil, fmt.Errorf("failed to parse schema: %w", err)
	}

	return &cgoCollection{
		handle: handle,
		schema: &schema,
		path:   path,
	}, nil
}

func (c *cgoCollection) Path() string {
	return c.path
}

func (c *cgoCollection) Schema() *CollectionSchema {
	return c.schema
}

func (c *cgoCollection) Stats() (*CollectionStats, error) {
	var cJSON *C.char
	status := C.zvec_collection_stats(c.handle, &cJSON)
	if err := statusToError(status); err != nil {
		return nil, err
	}
	defer C.zvec_free_string(cJSON)

	var stats CollectionStats
	if err := json.Unmarshal([]byte(C.GoString(cJSON)), &stats); err != nil {
		return nil, fmt.Errorf("failed to parse stats: %w", err)
	}
	return &stats, nil
}

func (c *cgoCollection) Options() CollectionOption {
	return DefaultCollectionOption()
}

func (c *cgoCollection) Destroy() error {
	return statusToError(C.zvec_collection_destroy(c.handle))
}

func (c *cgoCollection) Flush() error {
	return statusToError(C.zvec_collection_flush(c.handle))
}

func (c *cgoCollection) CreateIndex(fieldName string, params IndexParams, option ...CreateIndexOption) error {
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		return fmt.Errorf("failed to marshal index params: %w", err)
	}

	cField := C.CString(fieldName)
	defer C.free(unsafe.Pointer(cField))

	cParams := C.CString(string(paramsJSON))
	defer C.free(unsafe.Pointer(cParams))

	concurrency := 0
	if len(option) > 0 {
		concurrency = option[0].Concurrency
	}

	return statusToError(C.zvec_collection_create_index(c.handle, cField, cParams, C.int(concurrency)))
}

func (c *cgoCollection) DropIndex(fieldName string) error {
	cField := C.CString(fieldName)
	defer C.free(unsafe.Pointer(cField))
	return statusToError(C.zvec_collection_drop_index(c.handle, cField))
}

func (c *cgoCollection) Optimize(option ...OptimizeOption) error {
	concurrency := 0
	if len(option) > 0 {
		concurrency = option[0].Concurrency
	}
	return statusToError(C.zvec_collection_optimize(c.handle, C.int(concurrency)))
}

func (c *cgoCollection) AddColumn(schema *FieldSchema, expression string, option ...AddColumnOption) error {
	schemaJSON, err := json.Marshal(schema)
	if err != nil {
		return fmt.Errorf("failed to marshal field schema: %w", err)
	}

	cSchema := C.CString(string(schemaJSON))
	defer C.free(unsafe.Pointer(cSchema))

	cExpr := C.CString(expression)
	defer C.free(unsafe.Pointer(cExpr))

	concurrency := 0
	if len(option) > 0 {
		concurrency = option[0].Concurrency
	}

	return statusToError(C.zvec_collection_add_column(c.handle, cSchema, cExpr, C.int(concurrency)))
}

func (c *cgoCollection) DropColumn(fieldName string) error {
	cField := C.CString(fieldName)
	defer C.free(unsafe.Pointer(cField))
	return statusToError(C.zvec_collection_drop_column(c.handle, cField))
}

func (c *cgoCollection) AlterColumn(oldName string, newName string, newSchema *FieldSchema, option ...AlterColumnOption) error {
	cOldName := C.CString(oldName)
	defer C.free(unsafe.Pointer(cOldName))

	cNewName := C.CString(newName)
	defer C.free(unsafe.Pointer(cNewName))

	var cNewSchema *C.char
	if newSchema != nil {
		schemaJSON, err := json.Marshal(newSchema)
		if err != nil {
			return fmt.Errorf("failed to marshal field schema: %w", err)
		}
		cNewSchema = C.CString(string(schemaJSON))
		defer C.free(unsafe.Pointer(cNewSchema))
	}

	concurrency := 0
	if len(option) > 0 {
		concurrency = option[0].Concurrency
	}

	return statusToError(C.zvec_collection_alter_column(c.handle, cOldName, cNewName, cNewSchema, C.int(concurrency)))
}

func (c *cgoCollection) doWrite(docs []*Doc, fn func(*C.char, **C.char) C.zvec_status_t) ([]Status, error) {
	docsJSON, err := json.Marshal(docs)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal docs: %w", err)
	}

	cDocs := C.CString(string(docsJSON))
	defer C.free(unsafe.Pointer(cDocs))

	var cStatuses *C.char
	status := fn(cDocs, &cStatuses)
	if err := statusToError(status); err != nil {
		return nil, err
	}
	defer C.zvec_free_string(cStatuses)

	var statuses []Status
	if err := json.Unmarshal([]byte(C.GoString(cStatuses)), &statuses); err != nil {
		return nil, fmt.Errorf("failed to parse write statuses: %w", err)
	}
	return statuses, nil
}

func (c *cgoCollection) Insert(docs []*Doc) ([]Status, error) {
	return c.doWrite(docs, func(cDocs *C.char, out **C.char) C.zvec_status_t {
		return C.zvec_collection_insert(c.handle, cDocs, out)
	})
}

func (c *cgoCollection) Upsert(docs []*Doc) ([]Status, error) {
	return c.doWrite(docs, func(cDocs *C.char, out **C.char) C.zvec_status_t {
		return C.zvec_collection_upsert(c.handle, cDocs, out)
	})
}

func (c *cgoCollection) Update(docs []*Doc) ([]Status, error) {
	return c.doWrite(docs, func(cDocs *C.char, out **C.char) C.zvec_status_t {
		return C.zvec_collection_update(c.handle, cDocs, out)
	})
}

func (c *cgoCollection) Delete(ids []string) ([]Status, error) {
	idsJSON, err := json.Marshal(ids)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal IDs: %w", err)
	}

	cIDs := C.CString(string(idsJSON))
	defer C.free(unsafe.Pointer(cIDs))

	var cStatuses *C.char
	status := C.zvec_collection_delete(c.handle, cIDs, &cStatuses)
	if err := statusToError(status); err != nil {
		return nil, err
	}
	defer C.zvec_free_string(cStatuses)

	var statuses []Status
	if err := json.Unmarshal([]byte(C.GoString(cStatuses)), &statuses); err != nil {
		return nil, fmt.Errorf("failed to parse delete statuses: %w", err)
	}
	return statuses, nil
}

func (c *cgoCollection) DeleteByFilter(filter string) error {
	cFilter := C.CString(filter)
	defer C.free(unsafe.Pointer(cFilter))
	return statusToError(C.zvec_collection_delete_by_filter(c.handle, cFilter))
}

func (c *cgoCollection) Query(option QueryOption) ([]*Doc, error) {
	queryJSON, err := json.Marshal(option)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal query: %w", err)
	}

	cQuery := C.CString(string(queryJSON))
	defer C.free(unsafe.Pointer(cQuery))

	var cResults *C.char
	status := C.zvec_collection_query(c.handle, cQuery, &cResults)
	if err := statusToError(status); err != nil {
		return nil, err
	}
	defer C.zvec_free_string(cResults)

	var docs []*Doc
	if err := json.Unmarshal([]byte(C.GoString(cResults)), &docs); err != nil {
		return nil, fmt.Errorf("failed to parse query results: %w", err)
	}
	return docs, nil
}

func (c *cgoCollection) Fetch(ids []string) (map[string]*Doc, error) {
	idsJSON, err := json.Marshal(ids)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal IDs: %w", err)
	}

	cIDs := C.CString(string(idsJSON))
	defer C.free(unsafe.Pointer(cIDs))

	var cDocs *C.char
	status := C.zvec_collection_fetch(c.handle, cIDs, &cDocs)
	if err := statusToError(status); err != nil {
		return nil, err
	}
	defer C.zvec_free_string(cDocs)

	var docs map[string]*Doc
	if err := json.Unmarshal([]byte(C.GoString(cDocs)), &docs); err != nil {
		return nil, fmt.Errorf("failed to parse fetch results: %w", err)
	}
	return docs, nil
}
