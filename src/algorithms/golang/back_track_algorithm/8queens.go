package back_track_algorithm

import (
	"fmt"
)

/*
int[] result = new int[8];//全局或成员变量,下标表示行,值表示queen存储在哪一列

    public void cal8queens(int row) { // 调用方式：cal8queens(0);
        if (row == 8) { // 8个棋子都放置好了，打印结果
            printQueens(result);
            return; // 8行棋子都放好了，已经没法再往下递归了，所以就return
        }
        for (int column = 0; column < 8; ++column) { // 每一行都有8中放法
            if (isOk(row, column)) { // 有些放法不满足要求
                result[row] = column; // 第row行的棋子放到了column列
                cal8queens(row + 1); // 考察下一行
            }
        }
    }

    private boolean isOk(int row, int column) {//判断row行column列放置是否合适
        int leftup = column - 1, rightup = column + 1;
        for (int i = row - 1; i >= 0; --i) { // 逐行往上考察每一行
            if (result[i] == column) return false; // 第i行的column列有棋子吗？
            if (leftup >= 0) { // 考察左上对角线：第i行leftup列有棋子吗？
                if (result[i] == leftup) return false;
            }
            if (rightup < 8) { // 考察右上对角线：第i行rightup列有棋子吗？
                if (result[i] == rightup) return false;
            }
            --leftup;
            ++rightup;
        }
        return true;
    }
*/

/*
8皇后问题：
我们有一个 8x8 的棋盘，希望往里放 8 个棋子（皇后），每个棋子所在的行、列、对角线都不能有另一个棋子。

1.用一个列表result来存储8皇后每行所在位置，下标是行号，值存储列号
2.考察每一行应该放在哪个位置，每一行有8种放法，每一行的放置位置递归的考察
3.用一个函数来判断(row,column)这个点是否可以放置，规则：
	上面每一行的column列都没有值；
	上面每一行的左斜对角线都没有值；
	上面每一行的右斜对角线都没有值；
*/
var result = make([]int, 8)
var i int

func Call8queens(row int) {
	if row == 8 {
		i ++
		fmt.Println(fmt.Sprintf("解法:%d", i))
		priQueens(result)
		fmt.Println("")
		return
	}

	for column := 0; column < 8; column++ {
		if isOK(row, column) {
			result[row] = column
			Call8queens(row+1)
		}
	}
}

func isOK(row, column int) bool {
	leftUp := column - 1
	rightUp := column + 1
	for i := row-1; i >= 0; i-- {
		if result[i] == column {
			return false
		}
		if leftUp >= 0 {
			if result[i] == leftUp {
				return false
			}
		}
		if rightUp < 8 {
			if result[i] == rightUp {
				return false
			}
		}
		leftUp --
		rightUp ++
	}
	return true
}

func priQueens(result []int) () {
	for row := 0; row < 8; row++ {
		for column := 0; column < 8; column++ {
			if result[row] == column {
				fmt.Print("Q ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println("")
	}
}
