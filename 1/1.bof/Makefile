all:
	gcc -o stack -z execstack -fno-stack-protector stack.c -m32 -no-pie -g;
	sudo chown root:root stack;
	sudo chmod 4755 stack;
canary:
	gcc -o stack -z execstack stack.c -m32 -no-pie -g;
	sudo chown root:root stack;
	sudo chmod 4755 stack;
