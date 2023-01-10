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
	ValidateHTTPResponse(apiSpecFilePath, requestMethod, requestURI string, responseStatusCode int, responseHeaders map[string][]string, responseBody []byte) JSightValidationError
	ClearCache()
	Stat() string
}

type JSightValidationError interface {
	ReportedBy() string
	Type() string
	Code() int
	Title() string
	Detail() string
	Position() JSightPosition
	Trace() []string
	ToJSON() string
}

type JSightPosition interface {
	Index() int
}

//-------------------------- implementation ---------------------------

type jsightPlugin struct {
	validateHTTPRequestSymbol func(
		apiSpecFilePath, requestMethod, requestURI string,
		requestHeaders map[string][]string,
		requestBody []byte,
	) error
	validateHTTPResponseSymbol func(
		apiSpecFilePath, requestMethod, requestURI string,
		responseStatusCode int,
		responseHeaders map[string][]string,
		responseBody []byte,
	) error
	clearCacheSymbol func()
	statSymbol       func() string
}

type jsightValidationErrorStruct struct {
	e jsightPluginValidationError
}

type jsightPluginValidationError interface {
	ReportedBy() string
	Type() string
	Code() int
	Title() string
	Detail() string
	Position() any
	Trace() []string
	ToJSON() string
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

	s, err = p.Lookup("JSightValidateHTTPResponse")
	if err != nil {
		panic(err)
	}
	j.validateHTTPResponseSymbol = s.(func(string, string, string, int, map[string][]string, []byte) error)

	s, err = p.Lookup("JSightClearCache")
	if err != nil {
		panic(err)
	}
	j.clearCacheSymbol = s.(func())

	s, err = p.Lookup("JSightStat")
	if err != nil {
		panic(err)
	}
	j.statSymbol = s.(func() string)

	return j
}

func (j jsightPlugin) ValidateHTTPRequest(
	apiSpecFilePath, requestMethod, requestURI string,
	requestHeaders map[string][]string,
	requestBody []byte) JSightValidationError {
	e := j.validateHTTPRequestSymbol(apiSpecFilePath, requestMethod, requestURI, requestHeaders, requestBody)
	if e == nil {
		return nil
	}
	return newjsightValidationErrorStruct(e.(jsightPluginValidationError))
}

func (j jsightPlugin) ValidateHTTPResponse(
	apiSpecFilePath, requestMethod, requestURI string,
	responseCode int,
	responseHeaders map[string][]string,
	responseBody []byte) JSightValidationError {
	e := j.validateHTTPResponseSymbol(apiSpecFilePath, requestMethod, requestURI, responseCode, responseHeaders, responseBody)
	if e == nil {
		return nil
	}
	return newjsightValidationErrorStruct(e.(jsightPluginValidationError))
}

func (j jsightPlugin) ClearCache() {
	j.clearCacheSymbol()
}

func (j jsightPlugin) Stat() string {
	return j.statSymbol()
}

func newjsightValidationErrorStruct(e jsightPluginValidationError) JSightValidationError {
	return jsightValidationErrorStruct{e: e}
}

func (j jsightValidationErrorStruct) ReportedBy() string {
	return j.e.ReportedBy()
}

func (j jsightValidationErrorStruct) Type() string {
	return j.e.Type()
}

func (j jsightValidationErrorStruct) Code() int {
	return j.e.Code()
}

func (j jsightValidationErrorStruct) Title() string {
	return j.e.Title()
}

func (j jsightValidationErrorStruct) Detail() string {
	return j.e.Detail()
}

func (j jsightValidationErrorStruct) Position() JSightPosition {
	if j.e.Position() == nil {
		return nil
	}
	return j.e.Position().(JSightPosition)
}

func (j jsightValidationErrorStruct) Trace() []string {
	return j.e.Trace()
}

func (j jsightValidationErrorStruct) ToJSON() string {
	return j.e.ToJSON()
}
