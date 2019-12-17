// Copyright 2019 Alberto Bregliano. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package uddi creates pseudo UDDI tokens.
package uddi

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	rand.Seed(int64(99))
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{name: "primo", want: "75ed1842-49e9-bc19-675e-4d1f766213da"},
		{name: "secondo", want: "71b64278-fed2-cad5-f18a-247619205e30"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Create(); got != tt.want {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateCtx(t *testing.T) {
	rand.Seed(int64(99))
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "primo",
			args:    args{ctx: context.TODO()},
			want:    "75ed1842-49e9-bc19-675e-4d1f766213da",
			wantErr: false},
		{name: "primo",
			args:    args{ctx: context.TODO()},
			want:    "71b64278-fed2-cad5-f18a-247619205e30",
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateCtx(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleCreate() {
	rand.Seed(int64(99))
	t := Create()
	fmt.Println(t)
	// Output:
	// 75ed1842-49e9-bc19-675e-4d1f766213da
}

func ExampleCreateCtx() {
	rand.Seed(int64(99))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	time.Sleep(time.Second)
	t, err := CreateCtx(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)
	// Output:
	// impossible to generate UDDI: context deadline exceeded
}
