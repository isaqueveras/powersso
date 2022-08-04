// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package i18n

import (
	"github.com/gin-gonic/gin"
	"github.com/isaqueveras/lingo"
)

const (
	// PortugueseBR indicates the abbreviation of the language of the country Brazil
	PortugueseBR string = "pt_BR"
	// EnglishUS indicates the language abbreviation of the country United States of America
	EnglishUS string = "en_US"
	// SpainES indicates the abbreviation of the language of the country Spain
	SpainES string = "es_ES"

	// languageHeader indicates the header name to get the language in the request
	languageHeader string = "lang"
)

var ling lingo.T = lingo.T{}

// Setup add the language to use
func Setup(ctx *gin.Context, l *lingo.L) {
	var lang = ctx.GetHeader(languageHeader)
	if isValid(lang) {
		ling = l.TranslationsForLocale(lang)
	} else {
		ling = l.TranslationsForRequest(ctx.Request)
	}
}

// Value traverses the translations map and finds translation for given key.
func Value(value string, args ...string) string {
	return ling.Value(value, args...)
}

func isValid(value string) bool {
	return value != "" ||
		(value == PortugueseBR ||
			value == EnglishUS ||
			value == SpainES)
}
