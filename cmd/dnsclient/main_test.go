package main

import (
	"context"
	"testing"
)

func TestIntegration(t *testing.T) {
	main()
}

func TestHelp(t *testing.T) {
	*flagHelp = true
	err := mainWithContext(context.Background())
	*flagHelp = false
	if err != nil {
		t.Fatal(err)
	}
}

func TestLookupAddr(t *testing.T) {
	*flagType = "Addr"
	*flagName = "8.8.8.8"
	err := mainWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}

func TestLookupCNAME(t *testing.T) {
	*flagType = "CNAME"
	*flagName = "www.ooni.io"
	err := mainWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}

func TestLookupMX(t *testing.T) {
	*flagType = "MX"
	*flagName = "ooni.io"
	err := mainWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}

func TestLookupNS(t *testing.T) {
	*flagType = "NS"
	*flagName = "ooni.io"
	err := mainWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}

func TestLookupInvalid(t *testing.T) {
	*flagType = "Invalid"
	*flagName = "ooni.io"
	err := mainWithContext(context.Background())
	if err == nil {
		t.Fatal("expected an error here")
	}
}
