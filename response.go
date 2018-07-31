/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan

import "encoding/json"

// Envelope is the carrier of nearly every response
type Envelope struct {
	// 1 for good, 0 for error
	Status int `json:"status,string"`
	// OK for good, other words when Status equals 0
	Message string `json:"message"`
	// where response lies
	Result json.RawMessage `json:"result"`
}
