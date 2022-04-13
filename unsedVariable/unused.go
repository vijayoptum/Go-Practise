package main

func mul(abc int, xyz int) (int, int) {
	return (abc * xyz), (abc / xyz)

}

func main() {

	abc := 10
	xyz := 5
	mu, _ := mul(abc, xyz)
	println(mu)

}
