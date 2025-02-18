package fns

import (
	"math/rand"
	"time"

	"matrix/internal/pb"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func rand_layout() pb.Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Layout_LAYOUT_AZERTY
	case 2:
		return pb.Layout_LAYOUT_QWERTY
	case 3:
		return pb.Layout_LAYOUT_QWERTZ
	default:
		return pb.Layout_LAYOUT_UNSPECIFIED
	}
}

func rand_bool() bool {
	return rand.Intn(2) == 1
}

func rand_string(s ...string) string {
	if len(s) == 0 {
		return ""
	}

	return s[rand.Intn(len(s))]
}

func rand_int32(min, max int32) int32 {
	return min + rand.Int31n(max-min+1)
}

func rand_uint32(min, max uint32) uint32 {
	return min + rand.Uint32()%(max-min+1)
}

func rand_int64(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func rand_uint64(min, max uint64) uint64 {
	return min + rand.Uint64()%(max-min+1)
}

func rand_float64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func rand_float32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func rand_cpu_brand() string {
	return rand_string(CpuBrands...)
}

func rand_cpu_name(brand string) string {
	if brand == "Intel" {
		return rand_string(IntelCPUs...)
	}

	return rand_string(AmdCPUs...)
}

func rand_gpu_brand() string {
	return rand_string(GpuBrands...)
}

func rand_gpu_name(brand string) string {
	if brand == "Nvdia" {
		return rand_string(NvidiaGPUs...)
	}

	return rand_string(AmdGPUs...)
}

func rand_panel() pb.Panel {
	if rand.Intn(2) == 1 {
		return pb.Panel(pb.Panel_PANEL_IPS)
	}

	return pb.Panel(pb.Panel_PANEL_OLED)
}

func rand_screen_resolution() *pb.Resolution {
	h := rand_uint32(1000, 4320)
	w := h * 16 / 9

	return &pb.Resolution{
		Height: h,
		Width:  w,
	}
}

func rand_id() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	id := make([]byte, length)
	for i := range id {
		id[i] = charset[rand.Intn(len(charset))]
	}

	return string(id)
}

func rand_laptop_brand() string {
	return rand_string(LaptopBrands...)
}

func rand_laptop_name(brand string) string {
	switch brand {
	case "Apple":
		return rand_string(AppleLaptops...)
	case "Dell":
		return rand_string(DellLaptops...)
	case "HP":
		return rand_string(HpLaptops...)
	default:
		return ""
	}
}

func set_memory(unit pb.Unit) *pb.Memory {
	return &pb.Memory{
		Value: rand_uint64(2, 6),
		Unit:  unit,
	}
}
