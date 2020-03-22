package main

import "fmt"

type BitMap struct {
	data []byte
}

/*
	获取在字节数组中下标，num除以8 就得到num在数组中的位置
	num / 8 == num >> 3
*/
func (bm *BitMap) GetIndex(num uint) uint {
	return num >> 3
}

/*
	获取在一个字节中的位的位置，byte[index]中的第几位
	num % 8得到在一个字节中的位置
*/
func (bm *BitMap) GetPosition(num uint) uint {
	return num & (8 - 1)
}

/*
	标记指定数字（num）在bitmap中的值，标记其已经出现过
	将1左移position后，那个位置自然就是1，然后和以前的数据做或运算，这样，那个位置就替换成1了
*/
func (bm *BitMap) Add(num uint) {
	index := bm.GetIndex(num)
	bm.data[index] |= 1 << bm.GetPosition(num)
}

/*
	判断num是否在位图中
	将1左移position位置后，那个位置自然就是1，然后和以前的数据做与运算，判断是否为1，入
	如果结果为1，则以前那个位置就是1，否则以前的那个位置的数是0
*/
func (bm *BitMap) Contains(num uint) bool{
	index := bm.GetIndex(num)
	return bm.data[index] & 1 << bm.GetPosition(num) == 1
}

/*
	打印byte类型变量
	把byte转换为一个长度为8的数组，数组每个值代表bit
*/

func (bm *BitMap) ShowByte(b byte) {
	array := make([]byte, 8)

	for i := 0; i<8; i ++ {
		array[i] = b & 1
		b = b >> 1
	}

	for _, b1 := range array {
		fmt.Println(b1)
	}

}
