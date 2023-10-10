package product

import (
	"testing"

	spb "github.com/panupakm/boutique-go/api/shared"
	"github.com/panupakm/boutique-go/pkg/money"
)

func TestToProto(t *testing.T) {
	in := &Product{
		Id:          "1",
		Name:        "Test Product",
		Description: "This is a test product",
		Picture:     "https://example.com/test.jpg",
		PriceUsd:    money.Money{Units: 1000, CurrencyCode: "USD", Nanos: 500000},
		Categories:  []string{"test", "product"},
	}

	out := &spb.Product{}
	ToProto(in, out)

	if out.Id != in.Id {
		t.Errorf("Expected Id to be %s, but got %s", in.Id, out.Id)
	}

	if out.Name != in.Name {
		t.Errorf("Expected Name to be %s, but got %s", in.Name, out.Name)
	}

	if out.Description != in.Description {
		t.Errorf("Expected Description to be %s, but got %s", in.Description, out.Description)
	}

	if out.Picture != in.Picture {
		t.Errorf("Expected Picture to be %s, but got %s", in.Picture, out.Picture)
	}

	if out.PriceUsd.Units != in.PriceUsd.Units {
		t.Errorf("Expected PriceUsd.Units to be %d, but got %d", in.PriceUsd.Units, out.PriceUsd.Units)
	}

	if out.PriceUsd.CurrencyCode != in.PriceUsd.CurrencyCode {
		t.Errorf("Expected PriceUsd.CurrencyCode to be %s, but got %s", in.PriceUsd.CurrencyCode, out.PriceUsd.CurrencyCode)
	}

	if len(out.Categories) != len(in.Categories) {
		t.Errorf("Expected %d categories, but got %d", len(in.Categories), len(out.Categories))
	}

	for i, cat := range in.Categories {
		if out.Categories[i] != cat {
			t.Errorf("Expected category %d to be %s, but got %s", i, cat, out.Categories[i])
		}
	}
}

func TestToBiz(t *testing.T) {
	in := &spb.Product{
		Id:          "1",
		Name:        "Test Product",
		Description: "This is a test product",
		Picture:     "https://example.com/test.jpg",
		PriceUsd:    &spb.Money{Units: 1000, CurrencyCode: "USD", Nanos: 5000000},
		Categories:  []string{"test", "product"},
	}

	out := &Product{}
	ToBiz(in, out)

	if out.Id != in.Id {
		t.Errorf("Expected Id to be %s, but got %s", in.Id, out.Id)
	}

	if out.Name != in.Name {
		t.Errorf("Expected Name to be %s, but got %s", in.Name, out.Name)
	}

	if out.Description != in.Description {
		t.Errorf("Expected Description to be %s, but got %s", in.Description, out.Description)
	}

	if out.Picture != in.Picture {
		t.Errorf("Expected Picture to be %s, but got %s", in.Picture, out.Picture)
	}

	if out.PriceUsd.Units != in.PriceUsd.Units {
		t.Errorf("Expected PriceUsd.Units to be %d, but got %d", in.PriceUsd.Units, out.PriceUsd.Units)
	}

	if out.PriceUsd.CurrencyCode != in.PriceUsd.CurrencyCode {
		t.Errorf("Expected PriceUsd.CurrencyCode to be %s, but got %s", in.PriceUsd.CurrencyCode, out.PriceUsd.CurrencyCode)
	}

	if len(out.Categories) != len(in.Categories) {
		t.Errorf("Expected %d categories, but got %d", len(in.Categories), len(out.Categories))
	}

	for i, cat := range in.Categories {
		if out.Categories[i] != cat {
			t.Errorf("Expected category %d to be %s, but got %s", i, cat, out.Categories[i])
		}
	}
}
