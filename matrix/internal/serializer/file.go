package serializer

import (
	"fmt"
	"io"
	"os"

	"google.golang.org/protobuf/proto"
)

func ProtoToJSONFile(message proto.Message, filename string) error {
	data, err := ProtoToJSON(message)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	_, err = io.WriteString(file, string(data))
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	return nil
}

func ProtoToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	_, err = io.WriteString(file, string(data))
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	return nil
}
