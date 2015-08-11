package ykimg

import "testing"
import "fmt"
import "os"

func Test_save_img_channel(t *testing.T) {

	dowpic := "http://p5.yokacdn.com/pic/YOKA/2015-07-24/U10013P1TS1437709837_91385.jpg"

	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	go save_img_channel(dowpic, "4.jpg", c1)
	fmt.Println("doloading...1")
	go save_img_channel(dowpic, "5.jpg", c2)
	fmt.Println("doloading...2")
	s1 := <-c1
	s2 := <-c2
	fmt.Println(s1)
	fmt.Println(s2)

	for _, v := range []string{"4.jpg", "5.jpg"} {
		_, err := os.Stat(v)
		if err != nil {
			t.Error("failed!")
		}
	}
	fmt.Print(".")
}
