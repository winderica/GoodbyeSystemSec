KVERS = $(shell uname -r)

# Kernel modules
obj-m += MeltdownKernel.o

build: kernel_modules

kernel_modules:
	make -C /home/winderica/kernel M=$(CURDIR) modules

clean:
	make -C /home/winderica/kernel M=$(CURDIR) clean
