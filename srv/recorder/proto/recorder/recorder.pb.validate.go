// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: srv/recorder/proto/recorder/recorder.proto

package recordersrv

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _recorder_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate is disabled for TransactionReadRequest. This method will always
// return nil.
func (m *TransactionReadRequest) Validate() error {
	return nil
}

// TransactionReadRequestValidationError is the validation error returned by
// TransactionReadRequest.Validate if the designated constraints aren't met.
type TransactionReadRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TransactionReadRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TransactionReadRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TransactionReadRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TransactionReadRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TransactionReadRequestValidationError) ErrorName() string {
	return "TransactionReadRequestValidationError"
}

// Error satisfies the builtin error interface
func (e TransactionReadRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTransactionReadRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TransactionReadRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TransactionReadRequestValidationError{}

// Validate is disabled for TransactionWriteRequest. This method will always
// return nil.
func (m *TransactionWriteRequest) Validate() error {
	return nil
}

// TransactionWriteRequestValidationError is the validation error returned by
// TransactionWriteRequest.Validate if the designated constraints aren't met.
type TransactionWriteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TransactionWriteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TransactionWriteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TransactionWriteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TransactionWriteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TransactionWriteRequestValidationError) ErrorName() string {
	return "TransactionWriteRequestValidationError"
}

// Error satisfies the builtin error interface
func (e TransactionWriteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTransactionWriteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TransactionWriteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TransactionWriteRequestValidationError{}

// Validate checks the field values on TransactionEvent with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *TransactionEvent) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Req

	// no validation rules for Rsp

	return nil
}

// TransactionEventValidationError is the validation error returned by
// TransactionEvent.Validate if the designated constraints aren't met.
type TransactionEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TransactionEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TransactionEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TransactionEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TransactionEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TransactionEventValidationError) ErrorName() string { return "TransactionEventValidationError" }

// Error satisfies the builtin error interface
func (e TransactionEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTransactionEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TransactionEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TransactionEventValidationError{}
