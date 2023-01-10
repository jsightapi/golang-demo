package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test_JSightValidationErrorWithPosition(t *testing.T) {
	jSight := NewJSight("./plugins/jsightplugin.so")
	jsightSpecPath := "./testdata/test.jst"

	err := jSight.ValidateHTTPRequest(
		jsightSpecPath,
		"POST",
		"/users",
		nil,
		[]byte("Bad body"),
	)

	assertString(t, err.ReportedBy(), "HTTP Request validation", "ReportedBy")
	assertString(t, err.Type(), "http_body_error", "Type")
	assertInt(t, err.Code(), 32001, "Code")
	assertString(t, err.Title(), "HTTP body error", "Title")
	assertString(t, err.Detail(), "Invalid character \"B\" looking for beginning of value", "Detail")
	if err.Position() != nil {
		assertString(t, err.Position().Filepath(), "", "Position.FilePath")
		assertInt(t, err.Position().Index(), 0, "Index")
		assertInt(t, err.Position().Line(), 1, "Line")
		assertInt(t, err.Position().Col(), 1, "Col")
	}
}

func Test_JSightValidationErrorWithoutPosition(t *testing.T) {
	jSight := NewJSight("./plugins/jsightplugin.so")
	jsightSpecPath := "./testdata/test.jst"

	err := jSight.ValidateHTTPRequest(
		jsightSpecPath,
		"GET",
		"/users",
		nil,
		[]byte("Bad body"),
	)

	assertString(t, err.ReportedBy(), "HTTP Request validation", "ReportedBy")
	assertString(t, err.Type(), "path_error", "Type")
	assertInt(t, err.Code(), 24001, "Code")
	assertString(t, err.Title(), "Request path is not allowed", "Title")
	assertString(t, err.Detail(), "Request `GET /users` is not allowed.", "Detail")
	if err.Position() != nil {
		t.Error("Position must be nil.")
	}
}

func assertInt(t *testing.T, a int, e int, err string) {
	if a != e {
		t.Errorf("%s\n\nACTUAL:\n\n`%d`\n\nEXPECTED:\n\n`%d`\n", err, a, e)
	}
}

func assertString(t *testing.T, a string, e string, err string) {
	if a != e {
		t.Errorf("%s\n\nACTUAL:\n\n`%s`\n\nEXPECTED:\n\n`%s`\n", err, a, e)
	}
}

func printError(err JSightValidationError) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("ReportedBy = %s\n", err.ReportedBy()))
	sb.WriteString(fmt.Sprintf("Type = %s\n", err.Type()))
	sb.WriteString(fmt.Sprintf("Code = %d\n", err.Code()))
	sb.WriteString(fmt.Sprintf("Title = %s\n", err.Title()))
	sb.WriteString(fmt.Sprintf("Detail = %s\n", err.Detail()))
	if err.Position() != nil {
		sb.WriteString(fmt.Sprintf("Position.FilePath = %s\n", err.Position().Filepath()))
		sb.WriteString(fmt.Sprintf("Position.Index = %d\n", err.Position().Index()))
		sb.WriteString(fmt.Sprintf("Position.Line = %d\n", err.Position().Line()))
		sb.WriteString(fmt.Sprintf("Position.Col = %d\n", err.Position().Col()))
	}
	sb.WriteString(fmt.Sprintf("Trace = %v\n", err.Trace()))
	sb.WriteString(fmt.Sprintf("ToJSON = %s\n", err.ToJSON()))

	s := sb.String()
	return s
}
