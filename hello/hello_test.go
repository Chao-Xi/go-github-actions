package hello

import "testing"

// [Creating packages as function libraries]
// (https://github.com/Chao-Xi/JacobTechBlog/blob/master/golang/7Go_Struc_code.md#3creating-packages-as-function-libraries)

func TestGreet(t *testing.T) {
	result := Greet()
	if result != "Hello Github Action, Jacob is awesome" {
		t.Errorf("Greet() = %s; Expected Hello Github Action, Jacob is awesome", result)
	}
}
