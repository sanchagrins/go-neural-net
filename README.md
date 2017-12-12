# go-neural-net

A simple two-layer neural network written in Go that utilizes back-propagation for weight adjustment. The code was adapted from iamtrask's [A Neural Network in 11 Lines of Python](https://iamtrask.github.io/2015/07/12/basic-python-network/). For a more detailed explanation of the neural net please visit his blog post.

The project consists of two source files, `main.go` that serves an entry point for the application, and `utils.go` that contains various Linear Algebra based functions. While some of these functions can be found in the [gonum libraries](https://github.com/gonum/gonum), for this example it may be beneficial to see some of the simple matrix and vector functions to help highlight what's happening behind the scenes. 

## Requirements

This application was built using `go1.8.3`. For installation instructions see the link below:

* [Go](https://golang.org/doc/install)

## Download and Setup
Make sure that you have Go properly installed, then Download or clone the repository. From within the main directory run the application:

    $ go version
    go version go1.8.3 darwin/amd64
    $ git clone https://github.com/sanchagrins/go-neural-net.git
    $ cd go-neural-net

## Run
To run the application simply execute the `go run` command:

    $ go run *.go

The first layer of the network is specified by the input data `var l0`, which is a 2D array representation of the following 4x3 matrix, and the output data `var outData` as the 4x1 matrix:

<center>
Input | Output
---| --- 
0 0 1 | 0
0 1 1 | 0
1 0 1 | 1
1 1 1 | 1
</center>


After 100,000 rounds of training the results are as follows:

    Results after training:  [[0.003017650943251503] [0.0024610878499001974] [0.9979916200803733] [0.9975371819323701]]
    
As can be seen above the results after training are as expected, and are approximately equal to output data.

## Issues

If you find any bugs, feel free to file an issue on the github issue tracker.
