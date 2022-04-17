// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: videos/v1/videos.proto

package gate

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on GetVideoRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetVideoRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetVideoRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetVideoRequestMultiError, or nil if none found.
func (m *GetVideoRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetVideoRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for LastName

	if len(errors) > 0 {
		return GetVideoRequestMultiError(errors)
	}

	return nil
}

// GetVideoRequestMultiError is an error wrapping multiple validation errors
// returned by GetVideoRequest.ValidateAll() if the designated constraints
// aren't met.
type GetVideoRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetVideoRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetVideoRequestMultiError) AllErrors() []error { return m }

// GetVideoRequestValidationError is the validation error returned by
// GetVideoRequest.Validate if the designated constraints aren't met.
type GetVideoRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetVideoRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetVideoRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetVideoRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetVideoRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetVideoRequestValidationError) ErrorName() string { return "GetVideoRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetVideoRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetVideoRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetVideoRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetVideoRequestValidationError{}

// Validate checks the field values on GetVideoResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetVideoResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetVideoResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetVideoResponseMultiError, or nil if none found.
func (m *GetVideoResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetVideoResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if len(errors) > 0 {
		return GetVideoResponseMultiError(errors)
	}

	return nil
}

// GetVideoResponseMultiError is an error wrapping multiple validation errors
// returned by GetVideoResponse.ValidateAll() if the designated constraints
// aren't met.
type GetVideoResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetVideoResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetVideoResponseMultiError) AllErrors() []error { return m }

// GetVideoResponseValidationError is the validation error returned by
// GetVideoResponse.Validate if the designated constraints aren't met.
type GetVideoResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetVideoResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetVideoResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetVideoResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetVideoResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetVideoResponseValidationError) ErrorName() string { return "GetVideoResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetVideoResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetVideoResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetVideoResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetVideoResponseValidationError{}