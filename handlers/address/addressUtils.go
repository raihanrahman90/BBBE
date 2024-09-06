package address

import "bbbe/models"

type responseAddress struct {
	Address  string
	City     string
	Province string
	Username string
}

func responseConvert(data models.Address) responseAddress{
	var address responseAddress
	address.Address = data.Address
	address.City = data.City
	address.Province = data.Province
	address.Username = data.User.Username
	return address
}