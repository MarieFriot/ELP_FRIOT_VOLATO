package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
)

func bigListeGen(size int) []int {
	liste := []int{}
	for i := 0; i < size; i++ {
		liste = append(liste, rand.Intn(1000))
	}
	return liste
}

func clientV2(size int) []int {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer conn.Close()

	// Data to send
	data := bigListeGen(size)
	fmt.Println("Data sent :", data)
	sizeData := len(data)

	// Send data size
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
	intInBuff := 100
	var iteration int
	iteration = sizeData / intInBuff

	// Sending segments of the data to the server
	for i := 0; i <= iteration; i++ {

		buffSend := make([]byte, 0)

		if i == iteration && (sizeData%intInBuff) != 0 {
			for j, v := range data[i*intInBuff : sizeData] {
				binary.BigEndian.PutUint32(buffSend[j*4:], uint32(v))
			}
		} else if i == iteration && (sizeData%intInBuff) == 0 {
			fmt.Println("Nothing to be done on final iteration.")
		} else {
			// Conversion into bytes for buffer
			buffSend = make([]byte, 4*intInBuff)
			for j, v := range data[i*intInBuff : (i+1)*intInBuff] {
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
	for i := 0; i <= iteration; i++ {

		buffRecep := make([]byte, 0)

		if i == iteration && (sizeData%intInBuff) != 0 {
			buffRecep = make([]byte, 4*(sizeData-i*intInBuff))
			_, err := conn.Read(buffRecep)
			if err != nil {
				fmt.Println("Error reading data:", err)
				return nil
			}

			for j := 0; j < 4*(sizeData-i*intInBuff); j += 4 {
				val := int(binary.BigEndian.Uint32(buffRecep[j : j+4]))
				dataTreated = append(dataTreated, val)
			}

		} else if i == iteration && (sizeData%intInBuff) == 0 {
			fmt.Println("Nothing to be done on final iteration.")
		} else {
			buffRecep = make([]byte, 4*intInBuff)
			_, err := conn.Read(buffRecep)
			if err != nil {
				fmt.Println("Error reading data:", err)
				return nil
			}

			for j := 0; j < 4*intInBuff; j += 4 {
				val := int(binary.BigEndian.Uint32(buffRecep[j : j+4]))
				dataTreated = append(dataTreated, val)
			}
		}
	}

	return dataTreated
}

func main() {
	var n int

	fmt.Print("Type an integer: ")
	fmt.Scan(&n)

	dataTreated := clientV2(n)
	fmt.Println("Data received from the server:", dataTreated)
}
