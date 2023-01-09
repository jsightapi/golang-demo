package main

import (
	"plugin"
)

//------------------------- interfaces --------------------------------

func NewJSight(pluginPath string) JSight {
	return newjsightPlugin(pluginPath)
}

type JSight interface {
	ValidateHTTPRequest(apiSpecFilePath, requestMethod, requestURI string, requestHeaders map[string][]string, requestBody []byte) JSightValidationError
}

type JSightValidationError interface {
	ReportedBy() string
	Position() JSightPosition
}

type JSightPosition interface {
	Index() int
}

//-------------------------- implementation ---------------------------

type jsightPlugin struct {
	validateHTTPRequestSymbol func(
		apiSpecFilePath, requestMethod, requestURI string,
		requestHeaders map[string][]string,
		requestBody []byte) error
}

type jsightValidationErrorStruct struct {
	e jsightPluginValidationError
	p JSightPosition
}

type jsightPositionStruct struct {
	e jsightPluginValidationError
}

type jsightPluginValidationError interface {
	ReportedBy() string
	PositionIndex() int
}

func newjsightPlugin(pluginPath string) JSight {

	j := jsightPlugin{}

	p, err := plugin.Open(pluginPath)
	if err != nil {
		panic(err)
	}
	s, err := p.Lookup("JSightValidateHTTPRequest")
	if err != nil {
		panic(err)
	}
	j.validateHTTPRequestSymbol = s.(func(string, string, string, map[string][]string, []byte) error)

	return j
}

func (j jsightPlugin) ValidateHTTPRequest(
	apiSpecFilePath, requestMethod, requestURI string,
	requestHeaders map[string][]string,
	requestBody []byte) JSightValidationError {
	return j.validateHTTPRequestSymbol(apiSpecFilePath, requestMethod, requestURI, requestHeaders, requestBody).(JSightValidationError)
}

func newjsightValidationErrorStruct(e jsightPluginValidationError) JSightValidationError {
	return jsightValidationErrorStruct{
		e: e,
		p: newjsightPositionStruct(e),
	}
}

func (j jsightValidationErrorStruct) ReportedBy() string {
	return j.e.ReportedBy()
}

func (j jsightValidationErrorStruct) Position() JSightPosition {
	return j.p
}

func newjsightPositionStruct(e jsightPluginValidationError) JSightPosition {
	return jsightPositionStruct{e: e}
}

func (p jsightPositionStruct) Index() int {
	return p.e.PositionIndex()
}
