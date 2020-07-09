all:
	gcc -o stack -z noexecstack -fno-stack-protector stack.c -m32 -no-pie -g -static;
	sudo chown root:root stack;
	sudo chmod 4755 stack;
