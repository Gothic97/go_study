package main

import "fmt"

func main() {

	var num1 complex64 = complex(1, 2)

	fmt.Println(num1)

	// Q) 복소수의 용도는?

	// A)  아마도 정해진 실수에서 변수를 늘려야 할 경우 사용할 수 있을 것. 1, 2, 3이라는 수로는 3개의 경우의 수를 표현할 수 있으나
	//    복소수를 이용하면 1 + 2i, 1 + 3i ... 등을 표현할 수 있음. 그러나 1, 2, 3, 4, 5 ... 로도 표현할 수 있는데, '굳이
	//    복소수를 이용해야 하나?'라는 의문이 들기도 함.

	// A)  암호화 할 때 좋을 수도 있겠다는 생각이 듦. 수 연산에서 타입을 숨길 수 있다면, 복소수는 아주 좋은 암호 방식이 될 것임.
	//    특히 곱셈에서 i*i == -1이므로, 함수의 복잡성을 극대화시킬 수 있음.

}