// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: alarm/v1/alarm.proto

package v1

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

// Validate checks the field values on StatisticsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *StatisticsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StatisticsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StatisticsRequestMultiError, or nil if none found.
func (m *StatisticsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StatisticsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return StatisticsRequestMultiError(errors)
	}

	return nil
}

// StatisticsRequestMultiError is an error wrapping multiple validation errors
// returned by StatisticsRequest.ValidateAll() if the designated constraints
// aren't met.
type StatisticsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StatisticsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StatisticsRequestMultiError) AllErrors() []error { return m }

// StatisticsRequestValidationError is the validation error returned by
// StatisticsRequest.Validate if the designated constraints aren't met.
type StatisticsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StatisticsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StatisticsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StatisticsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StatisticsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StatisticsRequestValidationError) ErrorName() string {
	return "StatisticsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StatisticsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStatisticsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StatisticsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StatisticsRequestValidationError{}

// Validate checks the field values on StatisticsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StatisticsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StatisticsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StatisticsResponseMultiError, or nil if none found.
func (m *StatisticsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *StatisticsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	// no validation rules for Pages

	// no validation rules for Levels

	// no validation rules for Business

	// no validation rules for Groups

	if len(errors) > 0 {
		return StatisticsResponseMultiError(errors)
	}

	return nil
}

// StatisticsResponseMultiError is an error wrapping multiple validation errors
// returned by StatisticsResponse.ValidateAll() if the designated constraints
// aren't met.
type StatisticsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StatisticsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StatisticsResponseMultiError) AllErrors() []error { return m }

// StatisticsResponseValidationError is the validation error returned by
// StatisticsResponse.Validate if the designated constraints aren't met.
type StatisticsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StatisticsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StatisticsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StatisticsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StatisticsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StatisticsResponseValidationError) ErrorName() string {
	return "StatisticsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e StatisticsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStatisticsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StatisticsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StatisticsResponseValidationError{}

// Validate checks the field values on Filter with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Filter) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Filter with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in FilterMultiError, or nil if none found.
func (m *Filter) ValidateAll() error {
	return m.validate(true)
}

func (m *Filter) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	// no validation rules for Level

	// no validation rules for Business

	// no validation rules for Group

	// no validation rules for Status

	// no validation rules for StartAt

	// no validation rules for EndAt

	if len(errors) > 0 {
		return FilterMultiError(errors)
	}

	return nil
}

// FilterMultiError is an error wrapping multiple validation errors returned by
// Filter.ValidateAll() if the designated constraints aren't met.
type FilterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FilterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FilterMultiError) AllErrors() []error { return m }

// FilterValidationError is the validation error returned by Filter.Validate if
// the designated constraints aren't met.
type FilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FilterValidationError) ErrorName() string { return "FilterValidationError" }

// Error satisfies the builtin error interface
func (e FilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FilterValidationError{}

// Validate checks the field values on ListRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListRequestMultiError, or
// nil if none found.
func (m *ListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPage()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListRequestValidationError{
					field:  "Page",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListRequestValidationError{
					field:  "Page",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListRequestValidationError{
				field:  "Page",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetFilter()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListRequestValidationError{
					field:  "Filter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListRequestValidationError{
					field:  "Filter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFilter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListRequestValidationError{
				field:  "Filter",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ListRequestMultiError(errors)
	}

	return nil
}

// ListRequestMultiError is an error wrapping multiple validation errors
// returned by ListRequest.ValidateAll() if the designated constraints aren't met.
type ListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListRequestMultiError) AllErrors() []error { return m }

// ListRequestValidationError is the validation error returned by
// ListRequest.Validate if the designated constraints aren't met.
type ListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRequestValidationError) ErrorName() string { return "ListRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRequestValidationError{}

// Validate checks the field values on AlarmItem with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AlarmItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AlarmItem with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AlarmItemMultiError, or nil
// if none found.
func (m *AlarmItem) ValidateAll() error {
	return m.validate(true)
}

func (m *AlarmItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return AlarmItemMultiError(errors)
	}

	return nil
}

// AlarmItemMultiError is an error wrapping multiple validation errors returned
// by AlarmItem.ValidateAll() if the designated constraints aren't met.
type AlarmItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AlarmItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AlarmItemMultiError) AllErrors() []error { return m }

// AlarmItemValidationError is the validation error returned by
// AlarmItem.Validate if the designated constraints aren't met.
type AlarmItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AlarmItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AlarmItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AlarmItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AlarmItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AlarmItemValidationError) ErrorName() string { return "AlarmItemValidationError" }

// Error satisfies the builtin error interface
func (e AlarmItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAlarmItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AlarmItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AlarmItemValidationError{}

// Validate checks the field values on ListResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListResponseMultiError, or
// nil if none found.
func (m *ListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetAlarms() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListResponseValidationError{
						field:  fmt.Sprintf("Alarms[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListResponseValidationError{
						field:  fmt.Sprintf("Alarms[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListResponseValidationError{
					field:  fmt.Sprintf("Alarms[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetPage()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListResponseValidationError{
					field:  "Page",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListResponseValidationError{
					field:  "Page",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListResponseValidationError{
				field:  "Page",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetResponse()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListResponseValidationError{
					field:  "Response",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListResponseValidationError{
					field:  "Response",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResponse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListResponseValidationError{
				field:  "Response",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ListResponseMultiError(errors)
	}

	return nil
}

// ListResponseMultiError is an error wrapping multiple validation errors
// returned by ListResponse.ValidateAll() if the designated constraints aren't met.
type ListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListResponseMultiError) AllErrors() []error { return m }

// ListResponseValidationError is the validation error returned by
// ListResponse.Validate if the designated constraints aren't met.
type ListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListResponseValidationError) ErrorName() string { return "ListResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListResponseValidationError{}