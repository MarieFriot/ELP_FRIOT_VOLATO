package main

import (
	"Golang_project/quicksort"
	"encoding/binary"
	"fmt"
	"net"
	"runtime"
)

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Read the length of the data
	buffLen := make([]byte, 4)
	_, err := conn.Read(buffLen)
	if err != nil {
		fmt.Println("Error reading length: ", err)
		return
	}

	sizeData := int(binary.BigEndian.Uint32(buffLen))
	fmt.Print("Length read:", sizeData)

	// Determining amount of buffers of x ints to send (here x = 100)
	intInBuff := 100
	var iteration int
	iteration = sizeData / intInBuff

	// Doing the first partition while reading the data from client.
	var l1 []int
	var l2 []int
	var pivot int
	for i := 0; i <= iteration; i++ {

		buffRecep := make([]byte, 4)

		//case of the last loop, int in buffer needs to be set as the amount of int left
		if i == iteration && (sizeData%intInBuff) != 0 {
			buffRecep = make([]byte, 4*(sizeData-i*intInBuff))
			_, err := conn.Read(buffRecep)
			if err != nil {
				fmt.Println("Error reading data:", err)
				return
			}
			// Convert bytes into int then partition based on a pivot
			for j := 0; j < 4*(sizeData-i*intInBuff); j += 4 {
				val := int(binary.BigEndian.Uint32(buffRecep[j : j+4]))

				if (i == 0) && (j == 0) {
					pivot = val
				}

				if val <= pivot {
					l1 = append(l1, val)
				} else {
					l2 = append(l2, val)
				}
			}
		  //in last loop still, if we had a multiple of intInBuff, then there is no data left so we do nothing
		} else if i == iteration && (sizeData%intInBuff) == 0 {
			fmt.Println("Nothing to be done on final iteration.")
		} else {
			buffRecep = make([]byte, 4*intInBuff)
			_, err := conn.Read(buffRecep)
			if err != nil {
				fmt.Println("Error reading data:", err)
				return
			}

			for j := 0; j < 4*intInBuff; j += 4 {
				val := int(binary.BigEndian.Uint32(buffRecep[j : j+4]))

				if (i == 0) && (j == 0) {
					pivot = val
				}

				if val <= pivot {
					l1 = append(l1, val)
				} else {
					l2 = append(l2, val)
				}
			}
		}
	}

	fmt.Println("Data fully received, initiating treatment.")

	// Perform treatment on data
	listnumber := runtime.NumCPU()/4
	result1 := quicksort.QuicksortParallel(l1,listnumber)
	result2 := quicksort.QuicksortParallel(l2,listnumber)

	var data []int
	data = append(data, result1...)
	data = append(data, result2...)

	fmt.Println("Data fully treated, initiating sending.")
	// Sending segments of the data back to the client
	for i := 0; i <= iteration; i++ {

		buffAns := make([]byte, 0)
		// on last loop we follow the same logic as when receiving for the size of the buffer
		if i == iteration && (sizeData%intInBuff) != 0 {
			buffAns = make([]byte, 4*(sizeData-i*intInBuff))
			for j, v := range data[i*intInBuff : sizeData] {
				binary.BigEndian.PutUint32(buffAns[j*4:], uint32(v))
			}
		} else if i == iteration && (sizeData%intInBuff) == 0 {
			fmt.Println("Nothing to be done on final iteration.")
		} else {
			// Conversion int into bytes for buffer
			buffAns = make([]byte, 4*intInBuff)
			for j, v := range data[i*intInBuff : (i+1)*intInBuff] {
				binary.BigEndian.PutUint32(buffAns[j*4:], uint32(v))
			}
		}

		// Send data to the server
		_, err = conn.Write(buffAns)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	fmt.Println("Data treated and response sent successfully")
}

func test_server() {
	// Listen for incoming connections
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080.")

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		} else {
			fmt.Println("Client accepted.")
		}

		// Handle client connection in a goroutine
		go handleClient(conn)
	}
}

func main() {

	test_server()

}
