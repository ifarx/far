package util

import "./i18n"

type ErrorCode struct {
	_c uint32 /*code*/
	_m string /*messgae*/
	_l string /*language*/
}

func (e *ErrorCode) Code() uint32 {
	return e._c
}

func (e *ErrorCode) Message() string {
	if e._l != "" {
		return i18n.Message(e._l, e._m)
	}
	return e._m
}

func (e *ErrorCode) Succ() bool {
	return e._c == 0
}

func CreateErrorCode(c uint32, m string) *ErrorCode {
	return &ErrorCode{c, m, ""}
}

func CreateI18nErrorCode(c uint32, m, l string) *ErrorCode {
	return &ErrorCode{c, m, l}
}

var ERROR_CODE_SUCC = CreateErrorCode(0, "SUCC")

var ERROR_CODE_FAIL = CreateErrorCode(1, "FAIL")
