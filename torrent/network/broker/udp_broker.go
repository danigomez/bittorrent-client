package broker

import (
	"fmt"
	"net"
)

type UDPBroker struct {
}

func (broker *UDPBroker) SendRequest(request Request) ([]byte, error) {

	connection, err := net.Dial("udp", request.address)

	if err != nil {
		return nil, fmt.Errorf("error: there was an error while creating connection to %s, \n%s", request.address, err)
	}

	defer connection.Close()

	_, err = connection.Write(request.data)

	if err != nil {
		return nil, fmt.Errorf("error: there was an error sending UDP data to tracker \n%s", err)
	}

	buffer := make([]byte, 2048)

	read, err := connection.Read(buffer)

	if err != nil {
		return nil, fmt.Errorf("error: there was an error while reading UDP data from %s, \n%s", request.address, err)
	}

	// Truncate the buffer to the read size
	return buffer[:read], nil
}
