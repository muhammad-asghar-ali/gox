package fns

const (
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length  = 16
)

var (
	CpuBrands = []string{
		"Intel", "AMD",
	}

	GpuBrands = []string{
		"NVIDIA", "AMD",
	}

	LaptopBrands = []string{
		"Apple", "Dell", "HP",
	}
)

var (
	NvidiaGPUs = []string{
		"RTX 4090", "RTX 4080", "RTX 4070 Ti", "RTX 4070", "RTX 4060 Ti", "RTX 4060",
		"RTX 4050", "RTX 3090 Ti", "RTX 3090", "RTX 3080 Ti", "RTX 3080", "RTX 3070 Ti",
		"RTX 3070", "RTX 3060 Ti", "RTX 3060", "RTX 3050", "Titan RTX", "Quadro RTX 8000",
		"Quadro RTX 6000", "Quadro RTX 5000", "GeForce GTX 1660 Ti", "GeForce GTX 1660 Super",
		"GeForce GTX 1660", "GeForce GTX 1650 Super", "GeForce GTX 1650", "GeForce GTX 1080 Ti",
		"GeForce GTX 1080", "GeForce GTX 1070 Ti", "GeForce GTX 1070", "GeForce GTX 1060",
		"GeForce GTX 1050 Ti", "GeForce GTX 1050", "GTX Titan X", "GTX 980 Ti", "GTX 980", "GTX 970",
		"GTX 960", "GTX 950",
	}

	AmdGPUs = []string{
		"Radeon RX 7900 XTX", "Radeon RX 7900 XT", "Radeon RX 7800 XT", "Radeon RX 7700 XT",
		"Radeon RX 7600", "Radeon RX 6950 XT", "Radeon RX 6900 XT", "Radeon RX 6800 XT",
		"Radeon RX 6800", "Radeon RX 6700 XT", "Radeon RX 6700", "Radeon RX 6600 XT", "Radeon RX 6600",
		"Radeon RX 6500 XT", "Radeon RX 6400", "Radeon VII", "Radeon RX Vega 64", "Radeon RX Vega 56",
		"Radeon RX 590", "Radeon RX 580", "Radeon RX 570", "Radeon RX 560", "Radeon RX 550", "Radeon R9 Fury X",
		"Radeon R9 Fury", "Radeon R9 390X", "Radeon R9 390", "Radeon R9 380X", "Radeon R9 380", "Radeon R9 370",
		"Radeon R7 370", "Radeon R7 360", "Radeon R9 290X", "Radeon R9 290", "Radeon HD 7970", "Radeon HD 7950",
	}
)

var (
	IntelCPUs = []string{
		"Intel Core i9-13900K", "Intel Core i9-12900K", "Intel Core i7-13700K", "Intel Core i7-12700K",
		"Intel Core i5-13600K", "Intel Core i5-12600K", "Intel Core i3-12100", "Intel Core i9-11900K",
		"Intel Core i7-11700K", "Intel Core i5-11600K", "Intel Core i3-10100", "Intel Core i5-10400",
		"Intel Xeon W-3175X", "Intel Core i9-10980XE", "Intel Core i7-9700K", "Intel Core i5-9600K",
		"Intel Core i3-9100", "Intel Pentium Gold G5400", "Intel Celeron G4930",
	}

	AmdCPUs = []string{
		"AMD Ryzen 9 7950X", "AMD Ryzen 9 7900X", "AMD Ryzen 7 7700X", "AMD Ryzen 7 7800X3D", "AMD Ryzen 5 7600X",
		"AMD Ryzen 5 7600", "AMD Ryzen 5 5500", "AMD Ryzen 5 3400G", "AMD Ryzen 5 3600", "AMD Ryzen 7 3700X",
		"AMD Ryzen 9 3950X", "AMD Ryzen 9 5900X", "AMD Ryzen 9 5950X", "AMD Ryzen 7 5800X", "AMD Ryzen 5 5600X",
		"AMD Ryzen 3 3200G", "AMD Ryzen 3 2200G", "AMD Ryzen 3 1200", "AMD Ryzen Threadripper 3990X",
		"AMD Ryzen Threadripper 3960X", "AMD Ryzen Threadripper 2950X", "AMD Athlon 3000G",
	}
)

var (
	AppleLaptops = []string{
		"Apple MacBook Air",
		"Apple MacBook Pro",
		"Dell XPS 13",
		"Dell XPS 15",
		"HP Spectre x360",
		"HP Envy 13",
	}

	DellLaptops = []string{
		"Dell XPS 13",
		"Dell XPS 15",
		"Dell Inspiron 14",
	}

	HpLaptops = []string{
		"HP Spectre x360",
		"HP Envy 13",
		"HP Pavilion 15",
	}
)
