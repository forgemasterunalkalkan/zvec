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

// DataType represents column data types.
type DataType uint32

const (
	DataTypeUndefined DataType = 0

	DataTypeBinary DataType = 1
	DataTypeString DataType = 2
	DataTypeBool   DataType = 3
	DataTypeInt32  DataType = 4
	DataTypeInt64  DataType = 5
	DataTypeUint32 DataType = 6
	DataTypeUint64 DataType = 7
	DataTypeFloat  DataType = 8
	DataTypeDouble DataType = 9

	DataTypeVectorBinary32 DataType = 20
	DataTypeVectorBinary64 DataType = 21
	DataTypeVectorFP16     DataType = 22
	DataTypeVectorFP32     DataType = 23
	DataTypeVectorFP64     DataType = 24
	DataTypeVectorInt4     DataType = 25
	DataTypeVectorInt8     DataType = 26
	DataTypeVectorInt16    DataType = 27

	DataTypeSparseVectorFP16 DataType = 30
	DataTypeSparseVectorFP32 DataType = 31

	DataTypeArrayBinary DataType = 40
	DataTypeArrayString DataType = 41
	DataTypeArrayBool   DataType = 42
	DataTypeArrayInt32  DataType = 43
	DataTypeArrayInt64  DataType = 44
	DataTypeArrayUint32 DataType = 45
	DataTypeArrayUint64 DataType = 46
	DataTypeArrayFloat  DataType = 47
	DataTypeArrayDouble DataType = 48
)

// IsDenseVector reports whether the data type is a dense vector type.
func (d DataType) IsDenseVector() bool {
	return d >= DataTypeVectorBinary32 && d <= DataTypeVectorInt16
}

// IsSparseVector reports whether the data type is a sparse vector type.
func (d DataType) IsSparseVector() bool {
	return d >= DataTypeSparseVectorFP16 && d <= DataTypeSparseVectorFP32
}

// IsVector reports whether the data type is a vector type (dense or sparse).
func (d DataType) IsVector() bool {
	return d.IsDenseVector() || d.IsSparseVector()
}

// IsArray reports whether the data type is an array type.
func (d DataType) IsArray() bool {
	return d >= DataTypeArrayBinary && d <= DataTypeArrayDouble
}

// String returns the string representation of the DataType.
func (d DataType) String() string {
	switch d {
	case DataTypeUndefined:
		return "UNDEFINED"
	case DataTypeBinary:
		return "BINARY"
	case DataTypeString:
		return "STRING"
	case DataTypeBool:
		return "BOOL"
	case DataTypeInt32:
		return "INT32"
	case DataTypeInt64:
		return "INT64"
	case DataTypeUint32:
		return "UINT32"
	case DataTypeUint64:
		return "UINT64"
	case DataTypeFloat:
		return "FLOAT"
	case DataTypeDouble:
		return "DOUBLE"
	case DataTypeVectorBinary32:
		return "VECTOR_BINARY32"
	case DataTypeVectorBinary64:
		return "VECTOR_BINARY64"
	case DataTypeVectorFP16:
		return "VECTOR_FP16"
	case DataTypeVectorFP32:
		return "VECTOR_FP32"
	case DataTypeVectorFP64:
		return "VECTOR_FP64"
	case DataTypeVectorInt4:
		return "VECTOR_INT4"
	case DataTypeVectorInt8:
		return "VECTOR_INT8"
	case DataTypeVectorInt16:
		return "VECTOR_INT16"
	case DataTypeSparseVectorFP16:
		return "SPARSE_VECTOR_FP16"
	case DataTypeSparseVectorFP32:
		return "SPARSE_VECTOR_FP32"
	case DataTypeArrayBinary:
		return "ARRAY_BINARY"
	case DataTypeArrayString:
		return "ARRAY_STRING"
	case DataTypeArrayBool:
		return "ARRAY_BOOL"
	case DataTypeArrayInt32:
		return "ARRAY_INT32"
	case DataTypeArrayInt64:
		return "ARRAY_INT64"
	case DataTypeArrayUint32:
		return "ARRAY_UINT32"
	case DataTypeArrayUint64:
		return "ARRAY_UINT64"
	case DataTypeArrayFloat:
		return "ARRAY_FLOAT"
	case DataTypeArrayDouble:
		return "ARRAY_DOUBLE"
	default:
		return "UNDEFINED"
	}
}

// IndexType represents column index types.
type IndexType uint32

const (
	IndexTypeUndefined IndexType = 0
	IndexTypeHNSW     IndexType = 1
	IndexTypeIVF      IndexType = 3
	IndexTypeFlat     IndexType = 4
	IndexTypeInvert   IndexType = 10
)

// String returns the string representation of the IndexType.
func (i IndexType) String() string {
	switch i {
	case IndexTypeUndefined:
		return "UNDEFINED"
	case IndexTypeHNSW:
		return "HNSW"
	case IndexTypeIVF:
		return "IVF"
	case IndexTypeFlat:
		return "FLAT"
	case IndexTypeInvert:
		return "INVERT"
	default:
		return "UNDEFINED"
	}
}

// MetricType represents distance metric types for vector search.
type MetricType uint32

const (
	MetricTypeUndefined MetricType = 0
	MetricTypeL2       MetricType = 1
	MetricTypeIP       MetricType = 2
	MetricTypeCosine   MetricType = 3
	MetricTypeMIPSL2   MetricType = 4
)

// String returns the string representation of the MetricType.
func (m MetricType) String() string {
	switch m {
	case MetricTypeUndefined:
		return "UNDEFINED"
	case MetricTypeL2:
		return "L2"
	case MetricTypeIP:
		return "IP"
	case MetricTypeCosine:
		return "COSINE"
	case MetricTypeMIPSL2:
		return "MIPSL2"
	default:
		return "UNDEFINED"
	}
}

// QuantizeType represents quantization types for vectors.
type QuantizeType uint32

const (
	QuantizeTypeUndefined QuantizeType = 0
	QuantizeTypeFP16     QuantizeType = 1
	QuantizeTypeInt8     QuantizeType = 2
	QuantizeTypeInt4     QuantizeType = 3
)

// String returns the string representation of the QuantizeType.
func (q QuantizeType) String() string {
	switch q {
	case QuantizeTypeUndefined:
		return "UNDEFINED"
	case QuantizeTypeFP16:
		return "FP16"
	case QuantizeTypeInt8:
		return "INT8"
	case QuantizeTypeInt4:
		return "INT4"
	default:
		return "UNDEFINED"
	}
}

// StatusCode represents the status code of an operation.
type StatusCode uint32

const (
	StatusCodeOK                StatusCode = 0
	StatusCodeNotFound          StatusCode = 1
	StatusCodeAlreadyExists     StatusCode = 2
	StatusCodeInvalidArgument   StatusCode = 3
	StatusCodePermissionDenied  StatusCode = 4
	StatusCodeFailedPrecondition StatusCode = 5
	StatusCodeResourceExhausted StatusCode = 6
	StatusCodeUnavailable       StatusCode = 7
	StatusCodeInternalError     StatusCode = 8
	StatusCodeNotSupported      StatusCode = 9
	StatusCodeUnknown           StatusCode = 10
)

// String returns the string representation of the StatusCode.
func (s StatusCode) String() string {
	switch s {
	case StatusCodeOK:
		return "OK"
	case StatusCodeNotFound:
		return "NOT_FOUND"
	case StatusCodeAlreadyExists:
		return "ALREADY_EXISTS"
	case StatusCodeInvalidArgument:
		return "INVALID_ARGUMENT"
	case StatusCodePermissionDenied:
		return "PERMISSION_DENIED"
	case StatusCodeFailedPrecondition:
		return "FAILED_PRECONDITION"
	case StatusCodeResourceExhausted:
		return "RESOURCE_EXHAUSTED"
	case StatusCodeUnavailable:
		return "UNAVAILABLE"
	case StatusCodeInternalError:
		return "INTERNAL_ERROR"
	case StatusCodeNotSupported:
		return "NOT_SUPPORTED"
	case StatusCodeUnknown:
		return "UNKNOWN"
	default:
		return "UNKNOWN"
	}
}

// LogLevel represents logging severity levels.
type LogLevel uint8

const (
	LogLevelDebug LogLevel = 0
	LogLevelInfo  LogLevel = 1
	LogLevelWarn  LogLevel = 2
	LogLevelError LogLevel = 3
	LogLevelFatal LogLevel = 4
)

// String returns the string representation of the LogLevel.
func (l LogLevel) String() string {
	switch l {
	case LogLevelDebug:
		return "DEBUG"
	case LogLevelInfo:
		return "INFO"
	case LogLevelWarn:
		return "WARN"
	case LogLevelError:
		return "ERROR"
	case LogLevelFatal:
		return "FATAL"
	default:
		return "WARN"
	}
}

// LogType represents log output destinations.
type LogType uint32

const (
	LogTypeConsole LogType = 0
	LogTypeFile    LogType = 1
)

// String returns the string representation of the LogType.
func (l LogType) String() string {
	switch l {
	case LogTypeConsole:
		return "CONSOLE"
	case LogTypeFile:
		return "FILE"
	default:
		return "CONSOLE"
	}
}
