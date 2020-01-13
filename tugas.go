package tugas4

import "fmt"

func Tugas4(mahasiswa map[string]int) {
	for key, item := range mahasiswa {
		fmt.Println(key, "\t:", item, "cm")
	}
}
