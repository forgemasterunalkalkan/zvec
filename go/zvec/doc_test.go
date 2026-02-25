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

func TestNewDoc(t *testing.T) {
	doc := NewDoc("doc_1")
	if doc.ID != "doc_1" {
		t.Errorf("NewDoc ID = %q, want %q", doc.ID, "doc_1")
	}
	if doc.Score != 0 {
		t.Errorf("NewDoc Score = %f, want 0", doc.Score)
	}
	if doc.Fields == nil {
		t.Error("NewDoc Fields should not be nil")
	}
	if doc.Vectors == nil {
		t.Error("NewDoc Vectors should not be nil")
	}
}

func TestDocSetField(t *testing.T) {
	doc := NewDoc("doc_1")
	doc.SetField("title", "Hello World")
	doc.SetField("age", 42)

	if !doc.HasField("title") {
		t.Error("HasField('title') should return true")
	}
	if doc.HasField("nonexistent") {
		t.Error("HasField('nonexistent') should return false")
	}

	if got := doc.GetField("title"); got != "Hello World" {
		t.Errorf("GetField('title') = %v, want %q", got, "Hello World")
	}
	if got := doc.GetField("age"); got != 42 {
		t.Errorf("GetField('age') = %v, want 42", got)
	}
	if got := doc.GetField("nonexistent"); got != nil {
		t.Errorf("GetField('nonexistent') = %v, want nil", got)
	}
}

func TestDocSetVector(t *testing.T) {
	doc := NewDoc("doc_1")
	vec := []float32{0.1, 0.2, 0.3, 0.4}
	doc.SetVector("embedding", vec)

	if !doc.HasVector("embedding") {
		t.Error("HasVector('embedding') should return true")
	}
	if doc.HasVector("nonexistent") {
		t.Error("HasVector('nonexistent') should return false")
	}

	got, ok := doc.GetVector("embedding").([]float32)
	if !ok {
		t.Fatal("GetVector('embedding') should return []float32")
	}
	if len(got) != 4 {
		t.Errorf("GetVector('embedding') length = %d, want 4", len(got))
	}
}

func TestDocFieldNames(t *testing.T) {
	doc := NewDoc("doc_1")
	doc.SetField("a", 1)
	doc.SetField("b", 2)
	doc.SetVector("vec1", []float32{1.0})
	doc.SetVector("vec2", []float32{2.0})

	fieldNames := doc.FieldNames()
	if len(fieldNames) != 2 {
		t.Errorf("FieldNames() length = %d, want 2", len(fieldNames))
	}

	vectorNames := doc.VectorNames()
	if len(vectorNames) != 2 {
		t.Errorf("VectorNames() length = %d, want 2", len(vectorNames))
	}
}

func TestDocChaining(t *testing.T) {
	doc := NewDoc("doc_1").
		SetField("title", "test").
		SetVector("emb", []float32{0.1, 0.2})

	if doc.ID != "doc_1" {
		t.Errorf("doc.ID = %q, want %q", doc.ID, "doc_1")
	}
	if !doc.HasField("title") {
		t.Error("should have field 'title' after chained SetField")
	}
	if !doc.HasVector("emb") {
		t.Error("should have vector 'emb' after chained SetVector")
	}
}

func TestDocNilMaps(t *testing.T) {
	doc := &Doc{ID: "test"}
	if doc.GetField("x") != nil {
		t.Error("GetField on nil Fields should return nil")
	}
	if doc.GetVector("x") != nil {
		t.Error("GetVector on nil Vectors should return nil")
	}
	if doc.HasField("x") {
		t.Error("HasField on nil Fields should return false")
	}
	if doc.HasVector("x") {
		t.Error("HasVector on nil Vectors should return false")
	}
}
