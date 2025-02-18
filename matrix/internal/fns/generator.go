package fns

import (
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	"matrix/internal/pb"
)

func NewKeyboard() *pb.Keyboard {
	return &pb.Keyboard{
		Layout:  rand_layout(),
		Backlit: rand_bool(),
	}
}

func NewCPU() *pb.CPU {
	b := rand_cpu_brand()
	cores := rand_int32(2, 8)
	min_ghz := rand_float64(2.0, 3.3)

	return &pb.CPU{
		Brand:         b,
		Name:          rand_cpu_name(b),
		NumberCores:   cores,
		NumberThreads: rand_int32(cores, 12),
		MinGhz:        min_ghz,
		MaxGhz:        rand_float64(min_ghz, 5.0),
	}
}

func NewGPU() *pb.GPU {
	b := rand_gpu_brand()
	min_ghz := rand_float64(1.0, 1.5)

	return &pb.GPU{
		Brand:  b,
		Name:   rand_gpu_name(b),
		MinGhz: min_ghz,
		MaxGhz: rand_float64(min_ghz, 2.0),
		Memory: set_memory(pb.Unit_UNIT_GIGABYTE),
	}
}

func NewRAM() *pb.Memory {
	return &pb.Memory{
		Value: rand_uint64(4, 64),
		Unit:  pb.Unit_UNIT_GIGABYTE,
	}
}

func NewSSD() *pb.Stroage {
	return &pb.Stroage{
		Driver: pb.Driver_DRIVER_SSD,
		Memory: set_memory(pb.Unit_UNIT_GIGABYTE),
	}
}

func NewHDD() *pb.Stroage {
	return &pb.Stroage{
		Driver: pb.Driver_DRIVER_SSD,
		Memory: set_memory(pb.Unit_UNIT_TERABYTE),
	}
}

func NewScreen() *pb.Screen {
	return &pb.Screen{
		SizeInch:   rand_float32(13, 17),
		Resolution: rand_screen_resolution(),
		Panel:      rand_panel(),
		Multitouch: rand_bool(),
	}
}

func NewLaptop() *pb.Laptop {
	b := rand_gpu_brand()

	return &pb.Laptop{
		Id:           uuid.NewString(),
		Brand:        b,
		Name:         rand_laptop_name(b),
		Price:        rand_float64(100.0, 2000.0),
		ReleasedYear: rand_uint32(2000, 2024),
		Cpu:          NewCPU(),
		Ram:          NewRAM(),
		Screen:       NewScreen(),
		Keyboard:     NewKeyboard(),
		Gpuses:       []*pb.GPU{NewGPU()},
		Stroages:     []*pb.Stroage{NewHDD(), NewSSD()},
		Weigth: &pb.Laptop_WeigthKg{
			WeigthKg: rand_float64(1.0, 3.0),
		},
		UpdatedAt: timestamppb.Now(),
	}
}
