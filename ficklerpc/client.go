package ficklerpc

import (
	"errors"

	config "github.com/abhishekkr/ficklepickle/config"
	pickle "github.com/abhishekkr/ficklepickle/pickle"
	unpickle "github.com/abhishekkr/ficklepickle/unpickle"

	golnet "github.com/abhishekkr/gol/golnet"
)

// Write returns error state for writing over a TCP server with 'blob' mapped to 'id'.
func Write(id string, blob []byte) error {
	client := golnet.CreateTCPClient(config.RpcServerPort)
	defer client.Connection.Close()

	args := Args{
		Action: "write",
		Id:     id,
		Blob:   blob,
	}
	request, err := pickle.Gob(args)
	if err != nil {
		return errors.New("ficklepickle: rpc write failed while encoding request")
	}

	reply := Args{}
	bufReply := client.Request(request)
	if err := unpickle.Gob(bufReply, &reply); err != nil {
		return errors.New("ficklepickle: rpc write failed while unpickling reply")
	}
	return reply.Err
}

// Read returns byte array and error state for reading over TCP server mapped to 'id'.
func Read(id string) ([]byte, error) {
	client := golnet.CreateTCPClient(config.RpcServerPort)
	defer client.Connection.Close()

	args := Args{
		Action: "read",
		Id:     id,
	}
	request, err := pickle.Gob(args)
	if err != nil {
		return []byte{}, errors.New("ficklepickle: rpc read failed while encoding request")
	}

	reply := Args{}
	bufReply := client.Request(request)
	if err := unpickle.Gob(bufReply, &reply); err != nil {
		return []byte{}, errors.New("ficklepickle: rpc read failed while unpickling reply")
	}
	return reply.Blob, reply.Err
}

// Delete returns error state for deleting entry mapped to 'id' over a TCP server.
func Delete(id string) error {
	client := golnet.CreateTCPClient(config.RpcServerPort)
	defer client.Connection.Close()

	args := Args{
		Action: "delete",
		Id:     id,
	}
	request, err := pickle.Gob(args)
	if err != nil {
		return errors.New("ficklepickle: rpc delete failed while encoding request")
	}

	reply := Args{}
	bufReply := client.Request(request)
	if err := unpickle.Gob(bufReply, &reply); err != nil {
		return errors.New("ficklepickle: rpc delete failed while unpickling reply")
	}
	return reply.Err
}
