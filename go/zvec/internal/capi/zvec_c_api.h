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

#ifndef ZVEC_C_API_H
#define ZVEC_C_API_H

#include <stddef.h>
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Opaque handle types */
typedef void* zvec_collection_t;
typedef void* zvec_doc_t;

/* Status codes matching zvec::StatusCode */
typedef uint32_t zvec_status_code_t;

/* Status structure */
typedef struct {
    zvec_status_code_t code;
    const char* message; /* Owned by callee; caller must copy if needed */
} zvec_status_t;

/* Data types matching zvec::DataType */
typedef uint32_t zvec_data_type_t;

/* Index types matching zvec::IndexType */
typedef uint32_t zvec_index_type_t;

/* Metric types matching zvec::MetricType */
typedef uint32_t zvec_metric_type_t;

/* Quantize types matching zvec::QuantizeType */
typedef uint32_t zvec_quantize_type_t;

/* ====== Initialization ====== */

/**
 * Initialize the zvec library with configuration.
 * Can only be called once.
 *
 * @param config_json JSON-encoded configuration string (may be NULL for defaults)
 * @return Status of the operation
 */
zvec_status_t zvec_init(const char* config_json);

/* ====== Collection Lifecycle ====== */

/**
 * Create a new collection and open it.
 *
 * @param path Filesystem path for the collection
 * @param schema_json JSON-encoded collection schema
 * @param option_json JSON-encoded collection options (may be NULL for defaults)
 * @param out_collection Output: opaque collection handle
 * @return Status of the operation
 */
zvec_status_t zvec_collection_create_and_open(
    const char* path,
    const char* schema_json,
    const char* option_json,
    zvec_collection_t* out_collection);

/**
 * Open an existing collection.
 *
 * @param path Filesystem path of the collection
 * @param option_json JSON-encoded collection options (may be NULL for defaults)
 * @param out_collection Output: opaque collection handle
 * @return Status of the operation
 */
zvec_status_t zvec_collection_open(
    const char* path,
    const char* option_json,
    zvec_collection_t* out_collection);

/**
 * Close and release a collection handle.
 *
 * @param collection The collection handle to close
 */
void zvec_collection_close(zvec_collection_t collection);

/**
 * Destroy a collection (permanently delete from disk).
 *
 * @param collection The collection handle
 * @return Status of the operation
 */
zvec_status_t zvec_collection_destroy(zvec_collection_t collection);

/**
 * Flush pending writes to disk.
 *
 * @param collection The collection handle
 * @return Status of the operation
 */
zvec_status_t zvec_collection_flush(zvec_collection_t collection);

/**
 * Get the filesystem path of a collection.
 *
 * @param collection The collection handle
 * @return The path string (owned by the collection; do not free)
 */
const char* zvec_collection_path(zvec_collection_t collection);

/**
 * Get collection statistics as JSON.
 *
 * @param collection The collection handle
 * @param out_json Output: JSON-encoded statistics (caller must free with zvec_free_string)
 * @return Status of the operation
 */
zvec_status_t zvec_collection_stats(zvec_collection_t collection, char** out_json);

/**
 * Get collection schema as JSON.
 *
 * @param collection The collection handle
 * @param out_json Output: JSON-encoded schema (caller must free with zvec_free_string)
 * @return Status of the operation
 */
zvec_status_t zvec_collection_schema(zvec_collection_t collection, char** out_json);

/* ====== Index Operations ====== */

/**
 * Create an index on a field.
 *
 * @param collection The collection handle
 * @param field_name Name of the field to index
 * @param params_json JSON-encoded index parameters
 * @param concurrency Number of threads (0 for auto)
 * @return Status of the operation
 */
zvec_status_t zvec_collection_create_index(
    zvec_collection_t collection,
    const char* field_name,
    const char* params_json,
    int concurrency);

/**
 * Drop an index from a field.
 *
 * @param collection The collection handle
 * @param field_name Name of the indexed field
 * @return Status of the operation
 */
zvec_status_t zvec_collection_drop_index(
    zvec_collection_t collection,
    const char* field_name);

/**
 * Optimize a collection.
 *
 * @param collection The collection handle
 * @param concurrency Number of threads (0 for auto)
 * @return Status of the operation
 */
zvec_status_t zvec_collection_optimize(
    zvec_collection_t collection,
    int concurrency);

/* ====== Document Operations ====== */

/**
 * Insert documents into the collection.
 *
 * @param collection The collection handle
 * @param docs_json JSON-encoded array of documents
 * @param out_statuses_json Output: JSON-encoded array of statuses (caller must free)
 * @return Status of the operation
 */
zvec_status_t zvec_collection_insert(
    zvec_collection_t collection,
    const char* docs_json,
    char** out_statuses_json);

/**
 * Upsert documents into the collection.
 *
 * @param collection The collection handle
 * @param docs_json JSON-encoded array of documents
 * @param out_statuses_json Output: JSON-encoded array of statuses (caller must free)
 * @return Status of the operation
 */
zvec_status_t zvec_collection_upsert(
    zvec_collection_t collection,
    const char* docs_json,
    char** out_statuses_json);

/**
 * Update documents in the collection.
 *
 * @param collection The collection handle
 * @param docs_json JSON-encoded array of documents
 * @param out_statuses_json Output: JSON-encoded array of statuses (caller must free)
 * @return Status of the operation
 */
zvec_status_t zvec_collection_update(
    zvec_collection_t collection,
    const char* docs_json,
    char** out_statuses_json);

/**
 * Delete documents by IDs.
 *
 * @param collection The collection handle
 * @param ids_json JSON-encoded array of document IDs
 * @param out_statuses_json Output: JSON-encoded array of statuses (caller must free)
 * @return Status of the operation
 */
zvec_status_t zvec_collection_delete(
    zvec_collection_t collection,
    const char* ids_json,
    char** out_statuses_json);

/**
 * Delete documents matching a filter expression.
 *
 * @param collection The collection handle
 * @param filter Boolean expression string
 * @return Status of the operation
 */
zvec_status_t zvec_collection_delete_by_filter(
    zvec_collection_t collection,
    const char* filter);

/* ====== Query Operations ====== */

/**
 * Perform vector similarity search.
 *
 * @param collection The collection handle
 * @param query_json JSON-encoded query parameters
 * @param out_results_json Output: JSON-encoded array of result documents (caller must free)
 * @return Status of the operation
 */
zvec_status_t zvec_collection_query(
    zvec_collection_t collection,
    const char* query_json,
    char** out_results_json);

/**
 * Fetch documents by IDs.
 *
 * @param collection The collection handle
 * @param ids_json JSON-encoded array of document IDs
 * @param out_docs_json Output: JSON-encoded map of documents (caller must free)
 * @return Status of the operation
 */
zvec_status_t zvec_collection_fetch(
    zvec_collection_t collection,
    const char* ids_json,
    char** out_docs_json);

/* ====== Column Operations ====== */

/**
 * Add a new column to the collection.
 *
 * @param collection The collection handle
 * @param schema_json JSON-encoded field schema
 * @param expression Expression to compute values for existing documents
 * @param concurrency Number of threads (0 for auto)
 * @return Status of the operation
 */
zvec_status_t zvec_collection_add_column(
    zvec_collection_t collection,
    const char* schema_json,
    const char* expression,
    int concurrency);

/**
 * Drop a column from the collection.
 *
 * @param collection The collection handle
 * @param field_name Name of the column to drop
 * @return Status of the operation
 */
zvec_status_t zvec_collection_drop_column(
    zvec_collection_t collection,
    const char* field_name);

/**
 * Alter a column (rename and/or change schema).
 *
 * @param collection The collection handle
 * @param old_name Current column name
 * @param new_name New column name (empty string for no rename)
 * @param new_schema_json JSON-encoded new field schema (NULL for rename only)
 * @param concurrency Number of threads (0 for auto)
 * @return Status of the operation
 */
zvec_status_t zvec_collection_alter_column(
    zvec_collection_t collection,
    const char* old_name,
    const char* new_name,
    const char* new_schema_json,
    int concurrency);

/* ====== Utility ====== */

/**
 * Free a string allocated by the C API.
 *
 * @param str The string to free
 */
void zvec_free_string(char* str);

#ifdef __cplusplus
}
#endif

#endif /* ZVEC_C_API_H */
