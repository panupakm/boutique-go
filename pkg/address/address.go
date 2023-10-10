package address

import (
	spb "github.com/panupakm/boutique-go/api/shared"
)

type Address struct {
	StreetAddress string `bson:"street_address" json:"street_address"`
	City          string `bson:"city" json:"city"`
	State         string `bson:"state" json:"state"`
	Country       string `bson:"country" json:"country"`
	ZipCode       int32  `bson:"zip_code" json:"zip_code"`
}

func ToProto(in *Address, out *spb.Address) {
	out.StreetAddress = in.StreetAddress
	out.City = in.City
	out.State = in.State
	out.Country = in.Country
	out.ZipCode = in.ZipCode
}
