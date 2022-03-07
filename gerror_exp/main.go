// (C) Copyright 2021-2022 Hewlett Packard Enterprise Development LP
// Package gerror_exp contains ...
package main

import (
	"fmt"

	"github.com/Naga2HPE/go-examples/gerror_exp/pkg/gerror"
	//"github.com/cf-guardian/guardian/gerror"
)

/*
Author : Nagarjuna S
Date : 07/03/22 10:47 AM
Project : gerror_exp
File : main.go
*/

func main() {

	fmt.Printf("%v", first())
}

func first() error {
	return second()
}

func second() error {
	return gerror.New("ErrExample", "Example error message")
}
