package ficklerpc

import (
	"errors"

	config "github.com/abhishekkr/ficklepickle/config"
	database "github.com/abhishekkr/ficklepickle/database"
	pickle "github.com/abhishekkr/ficklepickle/pickle"
	unpickle "github.com/abhishekkr/ficklepickle/unpickle"

	gollog "github.com/abhishekkr/gol/gollog"
	golnet "github.com/abhishekkr/gol/golnet"
)

var (
	// serverOpened is global flag to manage state of TCP server initialization.
	serverOpened = false
)

// Close closes db if open.
func Close() {
	if serverOpened == true {
		golnet.TcpServerHalt = true
		serverOpened = false
	}
}

func fickleReplies(bufRequest []byte) []byte {
	args := Args{}
	if err := unpickle.Gob(bufRequest, &args); err != nil {
		args.Err = err
	} else {
		switch args.Action {
		case "read":
			args.Blob, args.Err = database.Read(args.Id)
			gollog.Infof("ficklepickle rpc read: %s", args.Id)
		case "write":
			args.Err = database.Write(args.Id, args.Blob)
			gollog.Infof("ficklepickle rpc write: %s", args.Id)
		case "delete":
			args.Err = database.Delete(args.Id)
			gollog.Infof("ficklepickle rpc delete: %s", args.Id)
		default:
			args.Err = errors.New("unidentifed procedure")
			gollog.Errf("ficklepickle rpc error: %s", args.Err.Error())
		}
	}

	reply, err := pickle.Gob(args)
	if err != nil {
		return []byte{}
	}
	return reply
}

// Server starts a TCP server at config.RpcServerPort handling calls via fickleReplies procedure.
func Server() {
	gollog.Infof("FicklePickle TCP server will listen at %s", config.RpcServerPort)
	golnet.TcpServerHalt = false
	serverOpened = true
	golnet.TCPServer(config.RpcServerPort, fickleReplies)
}
