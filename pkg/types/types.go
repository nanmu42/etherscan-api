package types

import (
	"math/big"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// BigInt is a wrapper over big.Int to implement only unmarshalText
// for json decoding.
type BigInt big.Int

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (b *BigInt) UnmarshalText(text []byte) error {
	var bigInt = new(big.Int)

	if string(text) == "" {
		bigInt.SetInt64(0)
		*b = BigInt(*bigInt)
		return nil
	}

	if err := bigInt.UnmarshalText(text); err != nil {
		return err
	}

	*b = BigInt(*bigInt)
	return nil
}

// Int returns b's *big.Int form
func (b *BigInt) Int() *big.Int { return (*big.Int)(b) }

// MarshalText implements the encoding.TextMarshaler
func (b *BigInt) MarshalText() ([]byte, error) {
	return []byte(b.Int().String()), nil
}

// Time is a wrapper over big.Int to implement only unmarshalText
// for json decoding.
type Time time.Time

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (t *Time) UnmarshalText(text []byte) error {
	input, err := strconv.ParseInt(string(text), 10, 64)
	if err != nil {
		return errors.Wrap(err, "strconv.ParseInt")
	}

	var timestamp = time.Unix(input, 0)
	*t = Time(timestamp)

	return nil
}

// Time returns t's time.Time form
func (t Time) Time() time.Time { return time.Time(t) }

// MarshalText implements the encoding.TextMarshaler
func (t Time) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatInt(t.Time().Unix(), 10)), nil
}
