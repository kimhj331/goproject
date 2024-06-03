package main

import (
	"fmt"              //표준 패키지(Go설치시 함께 설치)
	"usepkg/custompkg" // 현재 모듈에 속한 패키지

	"github.com/guptarohit/asciigraph"           //외부 저장소 패키지
	"github.com/tuckersGo/musthaveGo/ch16/expkg" // 외부 저장소 패키지
)

func main() {
	custompkg.PrintCustom()
	expkg.PrintSample()

	data := []float64{3, 4, 5, 6, 9, 6, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}
