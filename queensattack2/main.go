package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func min(x, y int32) int32 {
	if x < y {
		return x
	}
	return y
}

func max(x, y int32) int32 {
	if x > y {
		return x
	}
	return y
}

// Complete the queensAttack function below.
func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	r1 := r_q - 1                   // [|] Top
	r2 := n - r_q                   // [|] Bottom
	c1 := c_q - 1                   // [-] Left
	c2 := n - c_q                   // [-] Right
	maxD11 := min(r_q, c_q) - 1     // [/] Bottom Left
	maxD22 := n - max(r_q, c_q)     // [/] Top Right
	maxD21 := n - max(r_q, n+1-c_q) // [\] Top Left
	maxD12 := n - max(n+1-r_q, c_q) // [\] Bottom Right
	d11 := maxD11
	d22 := maxD22
	d21 := maxD21
	d12 := maxD12

	for _, ob := range obstacles {
		r_o, c_o := ob[0], ob[1]
		if c_o == c_q && r_o < r_q && r_q-r_o-1 < r1 {
			r1 = r_q - r_o - 1
		} else if c_o == c_q && r_o > r_q && r_o-r_q < r2 {
			r2 = r_o - r_q - 1
		} else if r_o == r_q && c_o < c_q && c_q-c_o-1 < c1 {
			c1 = c_q - c_o - 1
		} else if r_o == r_q && c_o > c_q && c_o-c_q < c2 {
			c2 = c_o - c_q - 1
		} else if r_o-r_q == c_o-c_q && r_o < r_q && maxD11-min(r_o, c_o) < d11 {
			d11 = maxD11 - min(r_o, c_o)
		} else if r_o-r_q == c_o-c_q && r_o > r_q && maxD22-(n-max(r_o, c_o)+1) < d22 {
			d22 = maxD22 - (n - max(r_o, c_o) + 1)
		} else if r_o-r_q == (c_o-c_q)*-1 && r_o > r_q && maxD21-(n-max(r_o, n+1-c_o)+1) < d21 {
			d21 = maxD21 - (n - max(r_o, n+1-c_o) + 1)
		} else if r_o-r_q == (c_o-c_q)*-1 && r_o < r_q && maxD12-(n-max(n+1-r_o, c_o)+1) < d12 {
			d12 = maxD12 - (n - max(n+1-r_o, c_o) + 1)
		}
	}

	return r1 + r2 + c1 + c2 + d11 + d22 + d21 + d12
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	r_qC_q := strings.Split(readLine(reader), " ")

	r_qTemp, err := strconv.ParseInt(r_qC_q[0], 10, 64)
	checkError(err)
	r_q := int32(r_qTemp)

	c_qTemp, err := strconv.ParseInt(r_qC_q[1], 10, 64)
	checkError(err)
	c_q := int32(c_qTemp)

	var obstacles [][]int32
	for i := 0; i < int(k); i++ {
		obstaclesRowTemp := strings.Split(readLine(reader), " ")

		var obstaclesRow []int32
		for _, obstaclesRowItem := range obstaclesRowTemp {
			obstaclesItemTemp, err := strconv.ParseInt(obstaclesRowItem, 10, 64)
			checkError(err)
			obstaclesItem := int32(obstaclesItemTemp)
			obstaclesRow = append(obstaclesRow, obstaclesItem)
		}

		if len(obstaclesRow) != int(2) {
			panic("Bad input")
		}

		obstacles = append(obstacles, obstaclesRow)
	}

	result := queensAttack(n, k, r_q, c_q, obstacles)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
