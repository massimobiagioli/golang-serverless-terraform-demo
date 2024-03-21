package models

type Device struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func CreateDevice(id string, name string, address string) Device {
	return Device{Id: id, Name: name, Address: address}
}
