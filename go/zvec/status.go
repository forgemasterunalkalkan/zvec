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

import "fmt"

// Status represents the result of an operation: success or failure with a message.
type Status struct {
	code    StatusCode
	message string
}

// NewStatus creates a new Status with the given code and message.
func NewStatus(code StatusCode, message string) Status {
	return Status{code: code, message: message}
}

// StatusOK returns a successful Status.
func StatusOK() Status {
	return Status{code: StatusCodeOK}
}

// OK reports whether the status indicates success.
func (s Status) OK() bool {
	return s.code == StatusCodeOK
}

// Code returns the status code.
func (s Status) Code() StatusCode {
	return s.code
}

// Message returns the error message (empty if OK).
func (s Status) Message() string {
	return s.message
}

// Error implements the error interface for Status.
func (s Status) Error() string {
	if s.OK() {
		return ""
	}
	return fmt.Sprintf("%s: %s", s.code, s.message)
}

// Err returns nil if the status is OK, otherwise returns the status as an error.
func (s Status) Err() error {
	if s.OK() {
		return nil
	}
	return s
}

// ZvecError represents an error returned by zvec operations.
type ZvecError struct {
	Status Status
}

// Error implements the error interface.
func (e *ZvecError) Error() string {
	return e.Status.Error()
}

// NewError creates a new ZvecError from a StatusCode and message.
func NewError(code StatusCode, message string) *ZvecError {
	return &ZvecError{Status: NewStatus(code, message)}
}
