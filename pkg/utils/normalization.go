/*
 * Copyright (c) 2023 Zander Schwid & Co. LLC.
 * SPDX-License-Identifier: BUSL-1.1
 */

package utils

import (
	"strings"
	"unicode"
)

func NormalizeUserId(userId string) string {

	var out strings.Builder

	for _, ch := range []byte(userId) {

		if ch >= '0' && ch <= '9' {
			out.WriteByte(ch)
			continue
		}

		if ch >= 'a' && ch <= 'z' {
			out.WriteByte(ch)
			continue
		}

		if ch >= 'A' && ch <= 'Z' {
			out.WriteByte(ch)
			continue
		}

	}

	return out.String()
}

func NormalizeEmail(email string) string {
	s := strings.TrimSpace(email)
	s = strings.ToLower(s)
	return strings.ReplaceAll(s, ":", "")
}

func NormalizeField(email string) string {
	s := strings.TrimSpace(email)
	return strings.ReplaceAll(s, ":", "")
}

func NormalizeCode(email string) string {
	s := strings.TrimSpace(email)
	return strings.ReplaceAll(s, ":", "")
}

func NormalizePageId(pageId string) string {
	return NormalizeIdentityField(pageId)
}

func NormalizeIdentityField(email string) string {

	var out strings.Builder

	for _, ch := range email {
		low := unicode.ToLower(ch)
		if low >= 'a' && low <= 'z' {
			out.WriteRune(low)
			continue
		}
		if low >= '0' && low <= '9' {
			out.WriteRune(low)
			continue
		}
		if low == '-' || low == '_' {
			out.WriteRune(low)
			continue
		}
	}

	return out.String()
}




