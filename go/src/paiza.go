// 超愚直探索
package main

import (
	"bufio"
	"os"
	"strings"
	"log"
	"strconv"
	"sort"
	"math"
	"fmt"
)

type T下請け会社情報 struct {
	x会社ID int
    x下請け会社人数 int
	x下請け会社費用 int
	x一人あたり費用 int
}

type L下請け会社情報 []T下請け会社情報

var l下請け会社情報 L下請け会社情報
var x人数 int

func scan() [][]string {
	var input [][]string
	var arr []string
	reader := bufio.NewReader(os.Stdin);
	for true {
		str, err := reader.ReadString('\n')
		str = strings.TrimRight(str, "\n")
		if err != nil {
			return input
		}
		arr = strings.Split(str," ")
		input = append(input, arr)
	}
	return input;
}
func normalization(input [][]string) (int,L下請け会社情報) {
	if len(input) < 3 {
		log.Fatal("err")
	}
	val,err := strconv.ParseInt(input[0][0],10,32)
	x人数 := int(val)
	if err != nil {
		log.Fatal("err")
	}
	if x人数 < 1 || 200000 < x人数 {
		log.Fatal("err")
	}
	val,err = strconv.ParseInt(input[1][0],10,32)
	x下請け会社数 := int(val)
	if err != nil {
		log.Fatal("err")
	}
	if x下請け会社数 < 1 || 50 < x下請け会社数 {
		log.Fatal("err")
	}
	var i int
	i = 0
	var line []string
	for _,line = range input[2:] {
		i++
		if len(line) != 2 {
			log.Fatal("err")
		}
		val,err = strconv.ParseInt(line[0],10,32)
		x下請け会社人数 := int(val)
		if err != nil {
			log.Fatal("err")
		}
		if x下請け会社人数 < 1 || 10000 < x下請け会社人数 {
			continue
		}
		val,err = strconv.ParseInt(line[1],10,32)
		x下請け会社発注費用 := int(val)
		if err != nil {
			log.Fatal("err")
		}
		if x下請け会社発注費用 < 1 || 5000000 < x下請け会社発注費用 {
			continue
		}
		x一人あたり費用 := x下請け会社発注費用/x下請け会社人数

		l下請け会社情報 = append(l下請け会社情報, T下請け会社情報{i, x下請け会社人数,x下請け会社発注費用,x一人あたり費用})
	}
	if i != x下請け会社数 {
		log.Fatal("err")
	}
}
func calc(l下請け会社情報 L下請け会社情報) (int, []int) {
	return calcInternal(x人数, math.MaxInt64, nil, 0, 0, l下請け会社情報);
}

func calcInternal(x必要人数, x最小コスト int, l選択済み会社 []int, x人数 int, x総コスト int, l下請け会社情報 L下請け会社情報) (int, []int) {
	if (x最小コスト < x総コスト) {
		return x最小コスト, nil
	}
	if (x人数 >= x必要人数) {
		return x総コスト, l選択済み会社
	}
	var x最良選択コスト int = x最小コスト
	var l最良選択会社 []int = nil
	for idx,x下請け会社情報 := range l下請け会社情報 {
		if -1 == pos(l選択済み会社, x下請け会社情報.x会社ID){

			x選択コスト, l選択会社 := calcInternal(x必要人数, x最小コスト, append(l選択済み会社, x下請け会社情報.x会社ID), x人数+x下請け会社情報.x下請け会社人数, x総コスト+x下請け会社情報.x下請け会社費用, l下請け会社情報[idx+1:])
			if x選択コスト < x最良選択コスト {
				x最良選択コスト = x選択コスト
				l最良選択会社 = l選択会社
			}
		}
	}
	if x最良選択コスト == x最小コスト {
		return math.MaxInt64,nil
	}
	return x最良選択コスト,l最良選択会社
}
func pos(slice []int,value int) int {
	for p, v := range slice {
		if (v == value) {
			return p
		}
	}
	return -1
}
func (t L下請け会社情報) Len() int {
	return len(t)
}
func (t L下請け会社情報) Less(i,j int) bool {
	return t[i].x一人あたり費用 < t[j].x一人あたり費用
}
func (t L下請け会社情報) Swap(i,j int) {
	t[i],t[j] = t[j],t[i]
}
func (t L下請け会社情報) Sort() {
	sort.Sort(t)
}

func main() {
	var input [][]string
	input = scan()
	normalization(input)
	xコスト,_ := calc()
	fmt.Println(xコスト)
}

