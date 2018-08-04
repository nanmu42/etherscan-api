/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan

import (
	"math/big"
	"testing"
	"time"
)

func TestBigInt(t *testing.T) {
	const ansStr = "255"
	var ans = big.NewInt(255)

	b := new(BigInt)
	err := b.UnmarshalText([]byte(ansStr))
	noError(t, err, "BigInt.UnmarshalText")

	if b.Int().Cmp(ans) != 0 {
		t.Fatalf("BigInt.UnmarshalText not working, got %v, want %v", b.Int(), ans)
	}
	textBytes, err := b.MarshalText()
	noError(t, err, "BigInt.MarshalText")

	if string(textBytes) != ansStr {
		t.Fatalf("BigInt.MarshalText not working, got %s, want %s", textBytes, ansStr)
	}
}

func TestTime(t *testing.T) {
	const ansStr = "1533396289"
	var ans = time.Unix(1533396289, 0)

	b := new(Time)
	err := b.UnmarshalText([]byte(ansStr))
	noError(t, err, "Time.UnmarshalText")

	if !b.Time().Equal(ans) {
		t.Fatalf("Time.UnmarshalText not working, got %v, want %v", b, ans)
	}
	textBytes, err := b.MarshalText()
	noError(t, err, "BigInt.MarshalText")

	if string(textBytes) != ansStr {
		t.Fatalf("Time.MarshalText not working, got %s, want %s", textBytes, ansStr)
	}
}
