package main

import (
	"fmt"

	"github.com/chyroc/tui/components/tui_select"
)

func main() {
	fmt.Printf("===== 分割线 =====\n\n")

	// 有标题
	if true {
		r := tui_select.NewStringSelect()
		r.SetOptions([]string{"齐木楠雄", "铁臂阿童木", "孙悟空"})
		r.SetTitle("选择你喜爱的人物")
		idx, option, err := r.Select()
		if err != nil {
			panic(err)
		}

		fmt.Printf("idx: %d, option: %s\n", idx, option)
	}

	fmt.Printf("===== 分割线 =====\n\n")

	// 无标题
	if true {
		r := tui_select.NewStringSelect()
		r.SetOptions([]string{"齐木楠雄", "铁臂阿童木", "孙悟空"})
		idx, option, err := r.Select()
		if err != nil {
			panic(err)
		}

		fmt.Printf("idx: %d, option: %s\n", idx, option)
	}

	fmt.Printf("===== 分割线 =====\n\n")

	// 有窗口
	if true {
		r := tui_select.NewStringSelect()
		r.SetOptions([]string{"齐木楠雄", "铁臂阿童木", "孙悟空"})
		r.SetTitle("选择你喜爱的人物")
		r.SetSize(2)
		idx, option, err := r.Select()
		if err != nil {
			panic(err)
		}

		fmt.Printf("idx: %d, option: %s\n", idx, option)
	}
}
