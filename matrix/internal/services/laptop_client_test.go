package services_test

import (
	"context"
	"net"
	"net/http"
	"testing"

	"connectrpc.com/connect"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	"matrix/internal/fns"
	"matrix/internal/pb"
	"matrix/internal/pb/pbconnect"
	"matrix/internal/serializer"
	"matrix/internal/services"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	store := services.NewInMemoryLaptopStore()
	addr := _test_server(t)
	client := _new_client(addr)

	laptop := fns.NewLaptop()
	lpq := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	req := connect.NewRequest(lpq)

	res, err := client.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res.Msg)
	require.Equal(t, laptop.Id, res.Msg.GetId())

	other, err := store.Find(res.Msg.Id)
	require.NoError(t, err)
	require.NotNil(t, other)

	require_same_laptop(t, laptop, other)
}

func _test_server(t *testing.T) string {
	server := grpc.NewServer()

	l, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	go server.Serve(l)

	return l.Addr().String()
}

func _new_client(addr string) pbconnect.LaptopServiceClient {
	c := &http.Client{}

	client := pbconnect.NewLaptopServiceClient(c, addr)

	return client
}

func require_same_laptop(t *testing.T, laptop1 *pb.Laptop, laptop2 *pb.Laptop) {
	json1, err := serializer.ProtoToJSON(laptop1)
	require.NoError(t, err)

	json2, err := serializer.ProtoToJSON(laptop2)
	require.NoError(t, err)

	require.Equal(t, json1, json2)
}
