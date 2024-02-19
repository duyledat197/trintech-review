// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: coupon.proto

package coupon

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

// Validate checks the field values on CreateCouponRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateCouponRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateCouponRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateCouponRequestMultiError, or nil if none found.
func (m *CreateCouponRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateCouponRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := CouponType_name[int32(m.GetType())]; !ok {
		err := CreateCouponRequestValidationError{
			field:  "Type",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Amount

	if all {
		switch v := interface{}(m.GetFrom()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateCouponRequestValidationError{
					field:  "From",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateCouponRequestValidationError{
					field:  "From",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFrom()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateCouponRequestValidationError{
				field:  "From",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateCouponRequestValidationError{
					field:  "To",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateCouponRequestValidationError{
					field:  "To",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateCouponRequestValidationError{
				field:  "To",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if _, err := url.Parse(m.GetImageUrl()); err != nil {
		err = CreateCouponRequestValidationError{
			field:  "ImageUrl",
			reason: "value must be a valid URI",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Description

	// no validation rules for Rules

	switch v := m.ApplyId.(type) {
	case *CreateCouponRequest_UserId:
		if v == nil {
			err := CreateCouponRequestValidationError{
				field:  "ApplyId",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetUserId()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CreateCouponRequestValidationError{
						field:  "UserId",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CreateCouponRequestValidationError{
						field:  "UserId",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetUserId()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateCouponRequestValidationError{
					field:  "UserId",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *CreateCouponRequest_ProductId:
		if v == nil {
			err := CreateCouponRequestValidationError{
				field:  "ApplyId",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetProductId()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CreateCouponRequestValidationError{
						field:  "ProductId",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CreateCouponRequestValidationError{
						field:  "ProductId",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetProductId()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateCouponRequestValidationError{
					field:  "ProductId",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return CreateCouponRequestMultiError(errors)
	}

	return nil
}

// CreateCouponRequestMultiError is an error wrapping multiple validation
// errors returned by CreateCouponRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateCouponRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateCouponRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateCouponRequestMultiError) AllErrors() []error { return m }

// CreateCouponRequestValidationError is the validation error returned by
// CreateCouponRequest.Validate if the designated constraints aren't met.
type CreateCouponRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateCouponRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateCouponRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateCouponRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateCouponRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateCouponRequestValidationError) ErrorName() string {
	return "CreateCouponRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateCouponRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateCouponRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateCouponRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateCouponRequestValidationError{}

// Validate checks the field values on CreateCouponResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateCouponResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateCouponResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateCouponResponseMultiError, or nil if none found.
func (m *CreateCouponResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateCouponResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return CreateCouponResponseMultiError(errors)
	}

	return nil
}

// CreateCouponResponseMultiError is an error wrapping multiple validation
// errors returned by CreateCouponResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateCouponResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateCouponResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateCouponResponseMultiError) AllErrors() []error { return m }

// CreateCouponResponseValidationError is the validation error returned by
// CreateCouponResponse.Validate if the designated constraints aren't met.
type CreateCouponResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateCouponResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateCouponResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateCouponResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateCouponResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateCouponResponseValidationError) ErrorName() string {
	return "CreateCouponResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateCouponResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateCouponResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateCouponResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateCouponResponseValidationError{}

// Validate checks the field values on DeleteCouponByIDRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteCouponByIDRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteCouponByIDRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteCouponByIDRequestMultiError, or nil if none found.
func (m *DeleteCouponByIDRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteCouponByIDRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteCouponByIDRequestMultiError(errors)
	}

	return nil
}

// DeleteCouponByIDRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteCouponByIDRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteCouponByIDRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteCouponByIDRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteCouponByIDRequestMultiError) AllErrors() []error { return m }

// DeleteCouponByIDRequestValidationError is the validation error returned by
// DeleteCouponByIDRequest.Validate if the designated constraints aren't met.
type DeleteCouponByIDRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteCouponByIDRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteCouponByIDRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteCouponByIDRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteCouponByIDRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteCouponByIDRequestValidationError) ErrorName() string {
	return "DeleteCouponByIDRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteCouponByIDRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteCouponByIDRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteCouponByIDRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteCouponByIDRequestValidationError{}

// Validate checks the field values on DeleteCouponByIDResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteCouponByIDResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteCouponByIDResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteCouponByIDResponseMultiError, or nil if none found.
func (m *DeleteCouponByIDResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteCouponByIDResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteCouponByIDResponseMultiError(errors)
	}

	return nil
}

// DeleteCouponByIDResponseMultiError is an error wrapping multiple validation
// errors returned by DeleteCouponByIDResponse.ValidateAll() if the designated
// constraints aren't met.
type DeleteCouponByIDResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteCouponByIDResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteCouponByIDResponseMultiError) AllErrors() []error { return m }

// DeleteCouponByIDResponseValidationError is the validation error returned by
// DeleteCouponByIDResponse.Validate if the designated constraints aren't met.
type DeleteCouponByIDResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteCouponByIDResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteCouponByIDResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteCouponByIDResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteCouponByIDResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteCouponByIDResponseValidationError) ErrorName() string {
	return "DeleteCouponByIDResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteCouponByIDResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteCouponByIDResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteCouponByIDResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteCouponByIDResponseValidationError{}
