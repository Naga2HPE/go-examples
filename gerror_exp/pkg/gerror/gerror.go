// (C) Copyright 2021-2022 Hewlett Packard Enterprise Development LP
/*
Package gerror provides an improved error type which captures an error tag and the stack trace
at construction time.
*/
package gerror

import (
	"fmt"
	"reflect"
	"runtime"
)

const stackSize = 4096

// A Tag represents an error identifier of any type.
type Tag interface{}

// A Gerror is a tagged error with a stack trace embedded in the Error() string.
type Gerror interface {
	// Returns the tag used to create this error.
	Tag() Tag

	// Returns the concrete type of the tag used to create this error.
	TagType() reflect.Type

	// Returns the string form of this error, which includes the tag value, the tag type, the error message, and a stack trace.
	Error() string

	// Test the tag used to create this error for equality with a given tag. Returns `true` if and only if the two are equal.
	EqualTag(Tag) bool

	// Message
	Message() string

	// Cause
	Cause() error
}

// New Returns an error containing the given tag and message and the current stack trace.
func New(tag Tag, message string) Gerror {
	var stack [stackSize]byte
	n := runtime.Stack(stack[:], false)
	return &err{tag, reflect.TypeOf(tag), message, stack[:n], nil}
}

// Newf Returns an error containing the given tag and format string and the current stack trace. The given inserts are applied to the format string to produce an error message.
func Newf(tag Tag, format string, insert ...interface{}) Gerror {
	return New(tag, fmt.Sprintf(format, insert...))
}

// NewFromError Return an error containing the given tag, the cause of the error, and the current stack trace.
func NewFromError(tag Tag, cause error) Gerror {
	if cause != nil {
		var stack [stackSize]byte
		n := runtime.Stack(stack[:], false)
		return &err{tag, reflect.TypeOf(tag), "Error caused by: " + cause.Error(), stack[:n], cause}
	}
	return nil
}

type err struct {
	tag        Tag
	typ        reflect.Type
	message    string
	stackTrace []byte
	cause      error
}

func (e *err) Error() string {
	return fmt.Sprintf("%v %v", e.tag, e.typ) + ": " + e.message + "\n" + string(e.stackTrace)
}

func (e *err) Tag() Tag {
	return e.tag
}

func (e *err) TagType() reflect.Type {
	return e.typ
}

func (e *err) EqualTag(tag Tag) bool {
	return e.typ == reflect.TypeOf(tag) && e.tag == tag
}

func (e *err) Message() string {
	return e.message
}

func (e *err) Cause() error {
	return e.cause
}

func (e *err) StackTrace() string {
	return string(e.stackTrace)
}
