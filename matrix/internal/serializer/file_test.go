package serializer

import (
	"os"
	"testing"

	"matrix/internal/fns"

	"github.com/stretchr/testify/require"
)

func TestProtoToBinaryFile(t *testing.T) {
	bfile := "../tmp/laptop.bin"

	laptop := fns.NewLaptop()

	err := ProtoToBinaryFile(laptop, bfile)
	require.NoError(t, err, "ProtoToBinaryFile should not return an error")

	_, err = os.Stat(bfile)
	require.NoError(t, err, "File should be created")

	defer os.Remove(bfile) // to delete the file after test pass
}

func TestProtoToJSONFile(t *testing.T) {
	jfile := "../tmp/laptop.json"

	laptop := fns.NewLaptop()

	err := ProtoToJSONFile(laptop, jfile)
	require.NoError(t, err, "ProtoToJSONFile should not return an error")

	_, err = os.Stat(jfile)
	require.NoError(t, err, "File should be created")

	defer os.Remove(jfile) // to delete the file after test pass
}
