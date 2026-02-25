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

func TestStatusOK(t *testing.T) {
	s := StatusOK()
	if !s.OK() {
		t.Error("StatusOK() should return an OK status")
	}
	if s.Code() != StatusCodeOK {
		t.Errorf("StatusOK().Code() = %v, want %v", s.Code(), StatusCodeOK)
	}
	if s.Message() != "" {
		t.Errorf("StatusOK().Message() = %q, want empty", s.Message())
	}
	if s.Err() != nil {
		t.Error("StatusOK().Err() should return nil")
	}
}

func TestStatusError(t *testing.T) {
	s := NewStatus(StatusCodeNotFound, "document not found")
	if s.OK() {
		t.Error("error status should not be OK")
	}
	if s.Code() != StatusCodeNotFound {
		t.Errorf("Code() = %v, want %v", s.Code(), StatusCodeNotFound)
	}
	if s.Message() != "document not found" {
		t.Errorf("Message() = %q, want %q", s.Message(), "document not found")
	}
	if s.Err() == nil {
		t.Error("Err() should return non-nil for error status")
	}
	if s.Error() == "" {
		t.Error("Error() should return non-empty string for error status")
	}
}

func TestZvecError(t *testing.T) {
	e := NewError(StatusCodeInternalError, "something went wrong")
	if e.Error() == "" {
		t.Error("ZvecError.Error() should return non-empty string")
	}
	if e.Status.Code() != StatusCodeInternalError {
		t.Errorf("ZvecError status code = %v, want %v", e.Status.Code(), StatusCodeInternalError)
	}
}
