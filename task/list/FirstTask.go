package list

import (
	"fmt"
	"scheduled/common"
)

type FirstTask struct{
	r common.R
}

func (this FirstTask) Todo(){
	fmt.Println("测试第一个任务")
}
