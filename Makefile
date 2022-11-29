all :createstaticlib labcompile

createstaticlib:
	@gcc -c library/lib.c
	@ar cr libmath.a lib.o
	@rm *.o

createdynamiclib:
	@gcc -c library/lib.c
	@export LD_LIBRARY_PATH="$$LD_LIBRARY_PATH:$$PWD"
	@gcc -shared -o libmath.so lib.o
	@rm *.o

labcompile:
	@if ([ $(number) != 3 ] && [ $(number) != 4 ] && [ $(number) != 7 ]); then \
		gcc -c $(number)/$(number).c; \
		gcc -o temp $(number).o -L ./ -lmath -lm; \
		./temp; \
	else \
		echo -n "a or b: "; \
		read letter; \
		gcc -c $(number)/$(number)$$letter.c; \
		gcc -o temp $(number)$$letter.o -L ./ -lmath -lm; \
		./temp; \
	fi; \

	@rm -rf *.o
	@rm -rf temp
clean:
	@rm -rf *.o
	@rm -rf *.a
	@rm -rf *.so
	@rm -rf temp	