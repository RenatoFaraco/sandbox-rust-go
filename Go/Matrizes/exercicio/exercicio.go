package main

import (
	"bufio"
	"encoding/base64"
	"image"
	"image/png"
	"io"
	"os"
)

func Show(f func(dx, dy int) [][]uint8) {
	const (
		dx = 256
		dy = 256
	)
	data := f(dx, dy)
	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = v
			m.Pix[i+1] = v
			m.Pix[i+2] = 255
			m.Pix[i+3] = 255
		}
	}
	ShowImagePng(m)
}

func ShowImage(m image.Image) {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	io.WriteString(w, "IMAGE:")
	b64 := base64.NewEncoder(base64.StdEncoding, w)
	err := (&png.Encoder{CompressionLevel: png.BestCompression}).Encode(b64, m)
	if err != nil {
		panic(err)
	}
	b64.Close()
	io.WriteString(w, "\n")
}

func ShowImagePng(m image.Image) {
	file, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = (&png.Encoder{CompressionLevel: png.BestCompression}).Encode(file, m)
	if err != nil {
		panic(err)
	}
}

func main() {
	Show(func(dx, dy int) [][]uint8 {
		data := make([][]uint8, dy)
		for y := 0; y < dy; y++ {
			data[y] = make([]uint8, dx)
			for x := 0; x < dx; x++ {
				data[y][x] = uint8(x ^ y)
			}
		}
		return data
	})
}
