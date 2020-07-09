all:
	gcc -o retlib -z noexecstack -fno-stack-protector retlib.c -m32 -no-pie -g;
	sudo chown root:root retlib;
	sudo chmod 4755 retlib;
canary:
	gcc -o retlib -z noexecstack retlib.c -m32 -no-pie -g;
	sudo chown root:root retlib;
	sudo chmod 4755 retlib;
