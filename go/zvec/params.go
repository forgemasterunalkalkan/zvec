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

// IndexParams is the interface for all index parameter types.
type IndexParams interface {
	// Type returns the index type.
	Type() IndexType
}

// InvertIndexParams configures an inverted index for scalar fields.
type InvertIndexParams struct {
	// EnableRangeOptimization enables range query optimization.
	EnableRangeOptimization bool

	// EnableExtendedWildcard enables suffix and infix wildcard search.
	// Note: prefix search is always enabled regardless of this setting.
	EnableExtendedWildcard bool
}

// Type returns IndexTypeInvert.
func (p *InvertIndexParams) Type() IndexType {
	return IndexTypeInvert
}

// VectorIndexParams is the base interface for vector index parameters.
type VectorIndexParams interface {
	IndexParams

	// MetricType returns the distance metric type.
	GetMetricType() MetricType

	// QuantizeType returns the quantization type.
	GetQuantizeType() QuantizeType
}

// HnswIndexParams configures an HNSW (Hierarchical Navigable Small World) index.
type HnswIndexParams struct {
	// Metric is the distance metric for similarity computation.
	Metric MetricType

	// M is the number of bi-directional links per element during construction.
	// Higher values improve accuracy but increase memory and build time. Default: 50.
	M int

	// EfConstruction is the size of the dynamic candidate list during construction.
	// Larger values yield better graph quality at the cost of slower build time. Default: 500.
	EfConstruction int

	// Quantize is the optional quantization type for vector compression.
	Quantize QuantizeType
}

// NewHnswIndexParams creates HnswIndexParams with default values.
func NewHnswIndexParams(metric MetricType) *HnswIndexParams {
	return &HnswIndexParams{
		Metric:         metric,
		M:              50,
		EfConstruction: 500,
		Quantize:       QuantizeTypeUndefined,
	}
}

// Type returns IndexTypeHNSW.
func (p *HnswIndexParams) Type() IndexType {
	return IndexTypeHNSW
}

// GetMetricType returns the metric type.
func (p *HnswIndexParams) GetMetricType() MetricType {
	return p.Metric
}

// GetQuantizeType returns the quantize type.
func (p *HnswIndexParams) GetQuantizeType() QuantizeType {
	return p.Quantize
}

// FlatIndexParams configures a flat (brute-force) index.
type FlatIndexParams struct {
	// Metric is the distance metric for similarity computation.
	Metric MetricType

	// Quantize is the optional quantization type for vector compression.
	Quantize QuantizeType
}

// NewFlatIndexParams creates FlatIndexParams with default values.
func NewFlatIndexParams(metric MetricType) *FlatIndexParams {
	return &FlatIndexParams{
		Metric:   metric,
		Quantize: QuantizeTypeUndefined,
	}
}

// Type returns IndexTypeFlat.
func (p *FlatIndexParams) Type() IndexType {
	return IndexTypeFlat
}

// GetMetricType returns the metric type.
func (p *FlatIndexParams) GetMetricType() MetricType {
	return p.Metric
}

// GetQuantizeType returns the quantize type.
func (p *FlatIndexParams) GetQuantizeType() QuantizeType {
	return p.Quantize
}

// IVFIndexParams configures an IVF (Inverted File Index) index.
type IVFIndexParams struct {
	// Metric is the distance metric for similarity computation.
	Metric MetricType

	// NList is the number of clusters (inverted lists). 0 means auto-select.
	NList int

	// NIters is the number of k-means clustering iterations during training. Default: 10.
	NIters int

	// UseSoar enables SOAR for improved IVF search performance.
	UseSoar bool

	// Quantize is the optional quantization type for vector compression.
	Quantize QuantizeType
}

// NewIVFIndexParams creates IVFIndexParams with default values.
func NewIVFIndexParams(metric MetricType) *IVFIndexParams {
	return &IVFIndexParams{
		Metric:   metric,
		NList:    0,
		NIters:   10,
		UseSoar:  false,
		Quantize: QuantizeTypeUndefined,
	}
}

// Type returns IndexTypeIVF.
func (p *IVFIndexParams) Type() IndexType {
	return IndexTypeIVF
}

// GetMetricType returns the metric type.
func (p *IVFIndexParams) GetMetricType() MetricType {
	return p.Metric
}

// GetQuantizeType returns the quantize type.
func (p *IVFIndexParams) GetQuantizeType() QuantizeType {
	return p.Quantize
}

// QueryParams is the interface for all query parameter types.
type QueryParams interface {
	// Type returns the index type this query is configured for.
	Type() IndexType
}

// HnswQueryParams configures query parameters for HNSW index search.
type HnswQueryParams struct {
	// Ef is the size of the dynamic candidate list during search.
	// Higher values improve recall but slow down search. Default: 300.
	Ef int

	// Radius is the search radius for range queries. Default: 0.0.
	Radius float32

	// IsLinear forces brute-force linear search. Default: false.
	IsLinear bool

	// IsUsingRefiner enables the refiner for the query. Default: false.
	IsUsingRefiner bool
}

// NewHnswQueryParams creates HnswQueryParams with default values.
func NewHnswQueryParams() *HnswQueryParams {
	return &HnswQueryParams{
		Ef: 300,
	}
}

// Type returns IndexTypeHNSW.
func (p *HnswQueryParams) Type() IndexType {
	return IndexTypeHNSW
}

// IVFQueryParams configures query parameters for IVF index search.
type IVFQueryParams struct {
	// NProbe is the number of closest clusters to search. Default: 10.
	NProbe int
}

// NewIVFQueryParams creates IVFQueryParams with default values.
func NewIVFQueryParams() *IVFQueryParams {
	return &IVFQueryParams{
		NProbe: 10,
	}
}

// Type returns IndexTypeIVF.
func (p *IVFQueryParams) Type() IndexType {
	return IndexTypeIVF
}

// CollectionOption configures options for opening or creating a collection.
type CollectionOption struct {
	// ReadOnly opens the collection in read-only mode. Default: false.
	ReadOnly bool

	// EnableMmap enables memory-mapped I/O for data files. Default: true.
	EnableMmap bool

	// MaxBufferSize is the maximum buffer size in bytes. Default: 64MB.
	// Ignored when ReadOnly is true.
	MaxBufferSize uint32
}

// DefaultCollectionOption returns a CollectionOption with default values.
func DefaultCollectionOption() CollectionOption {
	return CollectionOption{
		ReadOnly:      false,
		EnableMmap:    true,
		MaxBufferSize: 64 * 1024 * 1024,
	}
}

// CreateIndexOption configures options for index creation.
type CreateIndexOption struct {
	// Concurrency is the number of threads for index creation. 0 means auto.
	Concurrency int
}

// OptimizeOption configures options for collection optimization.
type OptimizeOption struct {
	// Concurrency is the number of threads for optimization. 0 means auto.
	Concurrency int
}

// AddColumnOption configures options for adding a column.
type AddColumnOption struct {
	// Concurrency is the number of threads for data backfill. 0 means auto.
	Concurrency int
}

// AlterColumnOption configures options for altering a column.
type AlterColumnOption struct {
	// Concurrency is the number of threads for column alteration. 0 means auto.
	Concurrency int
}

// InitConfig configures the zvec library initialization.
type InitConfig struct {
	// LogType specifies the log output destination. Default: LogTypeConsole.
	LogType *LogType

	// LogLevel specifies the minimum log severity. Default: LogLevelWarn.
	LogLevel *LogLevel

	// LogDir is the directory for log files (only used with LogTypeFile).
	LogDir *string

	// LogBasename is the base name for rotated log files. Default: "zvec.log".
	LogBasename *string

	// LogFileSize is the max size per log file in MB before rotation. Default: 2048.
	LogFileSize *int

	// LogOverdueDays is the number of days to retain rotated log files. Default: 7.
	LogOverdueDays *int

	// QueryThreads is the number of threads for query execution.
	// If nil, auto-detected from available CPU cores.
	QueryThreads *int

	// OptimizeThreads is the number of threads for background tasks.
	// If nil, defaults to same as QueryThreads or CPU count.
	OptimizeThreads *int

	// InvertToForwardScanRatio is the threshold to switch from inverted index
	// to full forward scan. Range: [0.0, 1.0]. Default: 0.9.
	InvertToForwardScanRatio *float32

	// BruteForceByKeysRatio is the threshold to use brute-force key lookup
	// over index. Range: [0.0, 1.0]. Default: 0.1.
	BruteForceByKeysRatio *float32

	// MemoryLimitMB is the soft memory cap in MB.
	// If nil, inferred from system/cgroup memory limits.
	MemoryLimitMB *int
}

// Validate checks the InitConfig for invalid values.
func (c *InitConfig) Validate() error {
	if c.QueryThreads != nil && *c.QueryThreads < 1 {
		return errors.New("query_threads must be >= 1")
	}
	if c.OptimizeThreads != nil && *c.OptimizeThreads < 1 {
		return errors.New("optimize_threads must be >= 1")
	}
	if c.MemoryLimitMB != nil && *c.MemoryLimitMB <= 0 {
		return errors.New("memory_limit_mb must be > 0")
	}
	if c.InvertToForwardScanRatio != nil {
		r := *c.InvertToForwardScanRatio
		if r < 0.0 || r > 1.0 {
			return fmt.Errorf("invert_to_forward_scan_ratio must be in [0.0, 1.0], got %f", r)
		}
	}
	if c.BruteForceByKeysRatio != nil {
		r := *c.BruteForceByKeysRatio
		if r < 0.0 || r > 1.0 {
			return fmt.Errorf("brute_force_by_keys_ratio must be in [0.0, 1.0], got %f", r)
		}
	}
	return nil
}
