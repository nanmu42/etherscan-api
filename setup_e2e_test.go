/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan

import (
	"testing"
	"time"
)

var (
	// api test client for many test cases
	api *Client
	// bucket default rate limiter
	bucket *Bucket
)

func init() {
	bucket = NewBucket(200 * time.Millisecond)

	api = New(Mainnet, "etherscan-api-e2e-test")
	api.Verbose = true
	api.BeforeRequest = func(module string, action string, param map[string]interface{}) error {
		bucket.Take()
		return nil
	}
}

// Bucket is a simple and easy rate limiter
// Use NewBucket() to construct one.
type Bucket struct {
	bucket     chan bool
	refillTime time.Duration
}

// NewBucket factory of Bucket
func NewBucket(refillTime time.Duration) (b *Bucket) {
	b = &Bucket{
		bucket:     make(chan bool),
		refillTime: refillTime,
	}

	go b.fillRoutine()

	return
}

// Take a action token from bucket,
// blocks if there is currently no token left.
func (b *Bucket) Take() {
	<-b.bucket
}

// fill a action token into bucket,
// no-op if the bucket is currently full
func (b *Bucket) fill() {
	b.bucket <- true
}

func (b *Bucket) fillRoutine() {
	ticker := time.NewTicker(b.refillTime)

	for range ticker.C {
		b.fill()
	}
}

// noError checks for testing error
func noError(t *testing.T, err error, msg string) {
	if err != nil {
		t.Fatalf("%s: %v", msg, err)
	}
}
