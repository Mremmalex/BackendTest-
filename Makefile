run:
	./main 
    
build:
	go build main.go

compile:
	go build main.go | ./main 
