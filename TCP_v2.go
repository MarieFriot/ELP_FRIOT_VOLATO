package main

import (
	"Golang_project/quicksort"
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
	"time"
)

func bigListeGen(size int) []int {
	liste := []int{}
	for i := 0; i < size; i++ {
		liste = append(liste, rand.Intn(1000))
	}
	return liste
}

func test_client() []int {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer conn.Close()

	// Data to send
	data := bigListeGen(250)
	print(data)
	sizeData := len(data)

	// 1st step: Send data size
	// Conversion into bytes for buffer
	buffLen := make([]byte, 4)
	binary.BigEndian.PutUint32(buffLen, uint32(sizeData))

	// Send data to the server
	_, err = conn.Write(buffLen)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	// Determining amount of buffers of size x to send (here x = 100)
	sizeBuff := 100
	var iteration int
	iteration = (sizeData / sizeBuff) + 1

	// 2nd step: Sending segments of the data to the server
	for i := 0; i < iteration; i++ {

		buffSend := make([]byte, 0)

		if i == (iteration - 1) {
			buffSend = make([]byte, 4*(sizeData-i*sizeBuff))
			for j, v := range data[i*sizeBuff : sizeData] {
				binary.BigEndian.PutUint32(buffSend[j*4:], uint32(v))
			}
		} else {
			// Conversion into bytes for buffer
			buffSend = make([]byte, 4*sizeBuff)
			for j, v := range data[i*sizeBuff : (i+1)*sizeBuff] {
				binary.BigEndian.PutUint32(buffSend[j*4:], uint32(v))
			}
		}

		// Send data to the server
		_, err = conn.Write(buffSend)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
	}

	// Read and process data from the server and convert the binary data back to integersm [TBD]
	var dataTreated []int
	for i := 0; i < iteration; i++ {

		buffRecep := make([]byte, 0)

		if i == (iteration - 1) {
			buffRecep = make([]byte, 4*(sizeData-i*sizeBuff))
			_, err := conn.Read(buffRecep)
			if err != nil {
				fmt.Println("Error reading data:", err)
				return nil
			}

			for j := 0; j < 4*(sizeData-i*sizeBuff); j += 4 {
				val := int(binary.BigEndian.Uint32(buffRecep[j : j+4]))
				dataTreated = append(dataTreated, val)
			}

		} else {
			buffRecep = make([]byte, 4*sizeBuff)
			_, err := conn.Read(buffRecep)
			if err != nil {
				fmt.Println("Error reading data:", err)
				return nil
			}

			for j := 0; j < 4*sizeBuff; j += 4 {
				val := int(binary.BigEndian.Uint32(buffRecep[j : j+4]))
				dataTreated = append(dataTreated, val)
			}
		}
	}

	return dataTreated
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// 1st step : Read the length of the data (to optimize buffer size if possible)
	buffLen := make([]byte, 4)
	_, err := conn.Read(buffLen)
	if err != nil {
		fmt.Println("Error reading length:", err)
		return
	}

	sizeData := int(binary.BigEndian.Uint32(buffLen))

	// Determining amount of buffers of size x to send (here x = 100)
	sizeBuff := 100
	var iteration int
	iteration = (sizeData / sizeBuff) + 1

	// Doing the partition step and read data from client. !TBD!
	var l1 []int
	var l2 []int
	var pivot int
	for i := 0; i < iteration; i++ {

		buffRecep := make([]byte, 4)

		if i == (iteration - 1) {
			buffRecep = make([]byte, 4*(sizeData-i*sizeBuff))
			_, err := conn.Read(buffRecep)
			if err != nil {
				fmt.Println("Error reading data:", err)
				return
			}

			for j := 0; j < 4*(sizeData-i*sizeBuff); j += 4 {
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

		} else {
			buffRecep = make([]byte, 4*sizeBuff)
			_, err := conn.Read(buffRecep)
			if err != nil {
				fmt.Println("Error reading data:", err)
				return
			}

			for j := 0; j < 4*sizeBuff; j += 4 {
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

	// Perform treatment on data
	result1 := quicksort.QuicksortParallel(l1)
	result2 := quicksort.QuicksortParallel(l2)

	var data []int
	data = append(data, result1...)
	data = append(data, result2...)

	// Sending segments of the data back to the client
	for i := 0; i < iteration; i++ {

		buffAns := make([]byte, 0)

		if i == (iteration - 1) {
			buffAns = make([]byte, 4*(sizeData-i*sizeBuff))
			for j, v := range data[i*sizeBuff : sizeData] {
				binary.BigEndian.PutUint32(buffAns[j*4:], uint32(v))
			}
		} else {
			// Conversion into bytes for buffer
			buffAns = make([]byte, 4*sizeBuff)
			for j, v := range data[i*sizeBuff : (i+1)*sizeBuff] {
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
		}

		// Handle client connection in a goroutine
		go handleClient(conn)
	}
}

func main() {

	go test_server()
	time.Sleep(2000)

	dataTreated := test_client()
	fmt.Println("Data received from the server:", dataTreated)

}
