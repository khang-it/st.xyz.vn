package main

import (
	"fmt"
	"st.xyz.vn/lib"
)

func main() {
	// Tạo một hình chữ nhật có chiều dài 10 và chiều rộng 5.
	rectangle := lib.Rectangle{
		Length: 10,
		Width:  5,
	}

	// Tính diện tích của hình chữ nhật.
	area := rectangle.Area()

	// In diện tích của hình chữ nhật.
	fmt.Println("Diện tích hình chữ nhật là:", area)
}
