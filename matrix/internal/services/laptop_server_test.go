package services_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"matrix/internal/fns"
	"matrix/internal/pb"
	"matrix/internal/services"
)

func TestCreateLaptop(t *testing.T) {
	t.Parallel()

	no_id := fns.NewLaptop()
	no_id.Id = ""

	invalid := fns.NewLaptop()
	invalid.Id = "invalid-uuid"

	duplicate := fns.NewLaptop()
	dupid := services.NewInMemoryLaptopStore()
	err := dupid.Save(duplicate)
	require.Nil(t, err)

	cases := []struct {
		name   string
		laptop *pb.Laptop
		store  services.LaptopStore
		code   codes.Code
	}{
		{
			name:   "success_with_id",
			laptop: fns.NewLaptop(),
			store:  services.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "success_no_id",
			laptop: no_id,
			store:  services.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "failure_invalid_id",
			laptop: invalid,
			store:  services.NewInMemoryLaptopStore(),
			code:   codes.InvalidArgument,
		},
		{
			name:   "failure_duplicate_id",
			laptop: duplicate,
			store:  dupid,
			code:   codes.AlreadyExists,
		},
	}

	for i := range cases {
		tc := cases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			req := &pb.CreateLaptopRequest{
				Laptop: tc.laptop,
			}

			server := services.NewLaptopServer(tc.store)

			res, err := server.CreateLaptop(context.Background(), req)
			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.Id)
				if len(tc.laptop.Id) > 0 {
					require.Equal(t, tc.laptop.Id, res.Id)
				}
			} else {
				require.Error(t, err) // Ensure error is returned.
				require.Nil(t, res)   // Response should be nil on error.
				st, ok := status.FromError(err)
				require.True(t, ok) // Verify it's a gRPC status error.
				require.Equal(t, tc.code, st.Code())
			}
		})
	}
}
