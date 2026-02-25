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

func TestDataTypeString(t *testing.T) {
	tests := []struct {
		dt   DataType
		want string
	}{
		{DataTypeUndefined, "UNDEFINED"},
		{DataTypeBinary, "BINARY"},
		{DataTypeString, "STRING"},
		{DataTypeBool, "BOOL"},
		{DataTypeInt32, "INT32"},
		{DataTypeInt64, "INT64"},
		{DataTypeUint32, "UINT32"},
		{DataTypeUint64, "UINT64"},
		{DataTypeFloat, "FLOAT"},
		{DataTypeDouble, "DOUBLE"},
		{DataTypeVectorFP32, "VECTOR_FP32"},
		{DataTypeVectorFP64, "VECTOR_FP64"},
		{DataTypeSparseVectorFP32, "SPARSE_VECTOR_FP32"},
		{DataTypeArrayInt32, "ARRAY_INT32"},
	}
	for _, tt := range tests {
		if got := tt.dt.String(); got != tt.want {
			t.Errorf("DataType(%d).String() = %q, want %q", tt.dt, got, tt.want)
		}
	}
}

func TestDataTypeIsVector(t *testing.T) {
	if !DataTypeVectorFP32.IsVector() {
		t.Error("VectorFP32 should be a vector type")
	}
	if !DataTypeVectorFP32.IsDenseVector() {
		t.Error("VectorFP32 should be a dense vector type")
	}
	if DataTypeVectorFP32.IsSparseVector() {
		t.Error("VectorFP32 should not be a sparse vector type")
	}
	if !DataTypeSparseVectorFP32.IsSparseVector() {
		t.Error("SparseVectorFP32 should be a sparse vector type")
	}
	if DataTypeString.IsVector() {
		t.Error("String should not be a vector type")
	}
}

func TestDataTypeIsArray(t *testing.T) {
	if !DataTypeArrayInt32.IsArray() {
		t.Error("ArrayInt32 should be an array type")
	}
	if DataTypeInt32.IsArray() {
		t.Error("Int32 should not be an array type")
	}
}

func TestIndexTypeString(t *testing.T) {
	tests := []struct {
		it   IndexType
		want string
	}{
		{IndexTypeUndefined, "UNDEFINED"},
		{IndexTypeHNSW, "HNSW"},
		{IndexTypeIVF, "IVF"},
		{IndexTypeFlat, "FLAT"},
		{IndexTypeInvert, "INVERT"},
	}
	for _, tt := range tests {
		if got := tt.it.String(); got != tt.want {
			t.Errorf("IndexType(%d).String() = %q, want %q", tt.it, got, tt.want)
		}
	}
}

func TestMetricTypeString(t *testing.T) {
	tests := []struct {
		mt   MetricType
		want string
	}{
		{MetricTypeUndefined, "UNDEFINED"},
		{MetricTypeL2, "L2"},
		{MetricTypeIP, "IP"},
		{MetricTypeCosine, "COSINE"},
		{MetricTypeMIPSL2, "MIPSL2"},
	}
	for _, tt := range tests {
		if got := tt.mt.String(); got != tt.want {
			t.Errorf("MetricType(%d).String() = %q, want %q", tt.mt, got, tt.want)
		}
	}
}

func TestQuantizeTypeString(t *testing.T) {
	tests := []struct {
		qt   QuantizeType
		want string
	}{
		{QuantizeTypeUndefined, "UNDEFINED"},
		{QuantizeTypeFP16, "FP16"},
		{QuantizeTypeInt8, "INT8"},
		{QuantizeTypeInt4, "INT4"},
	}
	for _, tt := range tests {
		if got := tt.qt.String(); got != tt.want {
			t.Errorf("QuantizeType(%d).String() = %q, want %q", tt.qt, got, tt.want)
		}
	}
}

func TestStatusCodeString(t *testing.T) {
	tests := []struct {
		sc   StatusCode
		want string
	}{
		{StatusCodeOK, "OK"},
		{StatusCodeNotFound, "NOT_FOUND"},
		{StatusCodeAlreadyExists, "ALREADY_EXISTS"},
		{StatusCodeInvalidArgument, "INVALID_ARGUMENT"},
		{StatusCodeInternalError, "INTERNAL_ERROR"},
	}
	for _, tt := range tests {
		if got := tt.sc.String(); got != tt.want {
			t.Errorf("StatusCode(%d).String() = %q, want %q", tt.sc, got, tt.want)
		}
	}
}

func TestLogLevelString(t *testing.T) {
	if got := LogLevelDebug.String(); got != "DEBUG" {
		t.Errorf("LogLevelDebug.String() = %q, want %q", got, "DEBUG")
	}
	if got := LogLevelWarn.String(); got != "WARN" {
		t.Errorf("LogLevelWarn.String() = %q, want %q", got, "WARN")
	}
}

func TestLogTypeString(t *testing.T) {
	if got := LogTypeConsole.String(); got != "CONSOLE" {
		t.Errorf("LogTypeConsole.String() = %q, want %q", got, "CONSOLE")
	}
	if got := LogTypeFile.String(); got != "FILE" {
		t.Errorf("LogTypeFile.String() = %q, want %q", got, "FILE")
	}
}
