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

func TestNewHnswIndexParams(t *testing.T) {
	p := NewHnswIndexParams(MetricTypeCosine)
	if p.Type() != IndexTypeHNSW {
		t.Errorf("Type() = %v, want %v", p.Type(), IndexTypeHNSW)
	}
	if p.GetMetricType() != MetricTypeCosine {
		t.Errorf("MetricType() = %v, want %v", p.GetMetricType(), MetricTypeCosine)
	}
	if p.M != 50 {
		t.Errorf("M = %d, want 50", p.M)
	}
	if p.EfConstruction != 500 {
		t.Errorf("EfConstruction = %d, want 500", p.EfConstruction)
	}
	if p.GetQuantizeType() != QuantizeTypeUndefined {
		t.Errorf("QuantizeType() = %v, want %v", p.GetQuantizeType(), QuantizeTypeUndefined)
	}
}

func TestNewFlatIndexParams(t *testing.T) {
	p := NewFlatIndexParams(MetricTypeIP)
	if p.Type() != IndexTypeFlat {
		t.Errorf("Type() = %v, want %v", p.Type(), IndexTypeFlat)
	}
	if p.GetMetricType() != MetricTypeIP {
		t.Errorf("MetricType() = %v, want %v", p.GetMetricType(), MetricTypeIP)
	}
}

func TestNewIVFIndexParams(t *testing.T) {
	p := NewIVFIndexParams(MetricTypeL2)
	if p.Type() != IndexTypeIVF {
		t.Errorf("Type() = %v, want %v", p.Type(), IndexTypeIVF)
	}
	if p.GetMetricType() != MetricTypeL2 {
		t.Errorf("MetricType() = %v, want %v", p.GetMetricType(), MetricTypeL2)
	}
	if p.NList != 0 {
		t.Errorf("NList = %d, want 0 (auto)", p.NList)
	}
	if p.NIters != 10 {
		t.Errorf("NIters = %d, want 10", p.NIters)
	}
}

func TestInvertIndexParams(t *testing.T) {
	p := &InvertIndexParams{
		EnableRangeOptimization: true,
		EnableExtendedWildcard:  false,
	}
	if p.Type() != IndexTypeInvert {
		t.Errorf("Type() = %v, want %v", p.Type(), IndexTypeInvert)
	}
}

func TestNewHnswQueryParams(t *testing.T) {
	p := NewHnswQueryParams()
	if p.Type() != IndexTypeHNSW {
		t.Errorf("Type() = %v, want %v", p.Type(), IndexTypeHNSW)
	}
	if p.Ef != 300 {
		t.Errorf("Ef = %d, want 300", p.Ef)
	}
}

func TestNewIVFQueryParams(t *testing.T) {
	p := NewIVFQueryParams()
	if p.Type() != IndexTypeIVF {
		t.Errorf("Type() = %v, want %v", p.Type(), IndexTypeIVF)
	}
	if p.NProbe != 10 {
		t.Errorf("NProbe = %d, want 10", p.NProbe)
	}
}

func TestDefaultCollectionOption(t *testing.T) {
	opt := DefaultCollectionOption()
	if opt.ReadOnly {
		t.Error("ReadOnly should be false by default")
	}
	if !opt.EnableMmap {
		t.Error("EnableMmap should be true by default")
	}
	if opt.MaxBufferSize != 64*1024*1024 {
		t.Errorf("MaxBufferSize = %d, want %d", opt.MaxBufferSize, 64*1024*1024)
	}
}

func TestInitConfigValidate(t *testing.T) {
	tests := []struct {
		name    string
		config  InitConfig
		wantErr bool
	}{
		{
			name:    "empty config",
			config:  InitConfig{},
			wantErr: false,
		},
		{
			name:    "valid config",
			config:  InitConfig{QueryThreads: intPtr(4)},
			wantErr: false,
		},
		{
			name:    "invalid query threads",
			config:  InitConfig{QueryThreads: intPtr(0)},
			wantErr: true,
		},
		{
			name:    "invalid memory limit",
			config:  InitConfig{MemoryLimitMB: intPtr(-1)},
			wantErr: true,
		},
		{
			name:    "valid ratio",
			config:  InitConfig{InvertToForwardScanRatio: float32Ptr(0.5)},
			wantErr: false,
		},
		{
			name:    "invalid ratio too high",
			config:  InitConfig{InvertToForwardScanRatio: float32Ptr(1.5)},
			wantErr: true,
		},
		{
			name:    "invalid ratio negative",
			config:  InitConfig{BruteForceByKeysRatio: float32Ptr(-0.1)},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("InitConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDefaultQueryOption(t *testing.T) {
	opt := DefaultQueryOption()
	if opt.TopK != 10 {
		t.Errorf("TopK = %d, want 10", opt.TopK)
	}
}

func TestVectorQueryHasID(t *testing.T) {
	q := VectorQuery{FieldName: "emb", ID: "doc_1"}
	if !q.HasID() {
		t.Error("HasID() should return true when ID is set")
	}
	if q.HasVector() {
		t.Error("HasVector() should return false when Vector is nil")
	}

	q2 := VectorQuery{FieldName: "emb", Vector: []float32{0.1, 0.2}}
	if q2.HasID() {
		t.Error("HasID() should return false when ID is empty")
	}
	if !q2.HasVector() {
		t.Error("HasVector() should return true when Vector is set")
	}
}

func intPtr(i int) *int {
	return &i
}

func float32Ptr(f float32) *float32 {
	return &f
}
