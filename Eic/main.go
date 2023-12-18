package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	out, err := os.Create("C:/Users/achon/onedrive/Desktop/diplomna/rra/Eic/out1.txt")
	if err != nil {
		os.Exit(1)
		fmt.Println(err)
	}
	defer out.Close()

	resp, err := http.Get("https://www.eicar.org/download/eicar-com-2/?wpdmdl=8842&refresh=656a05f4d5e5d1701447156")
	if err != nil {
		os.Exit(2)
		fmt.Println(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		os.Exit(3)
		fmt.Println(err)
	}

	_, err = os.Open("C:/Users/achon/onedrive/Desktop/diplomna/rra/Eic/out1.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("eho")
}

//aaaaaaa
