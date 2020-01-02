package main

//Time：O(n)，Space：O(n)
func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) == 0 && len(pushed) == 0 { //注意如果两个输入数组都为空，是合法情况
		return true
	}
	//任意一个数组为空或两个数组长度不相同，均不是合法情况
	if len(pushed) == 0 || len(popped) == 0 || len(pushed) != len(popped) {
		return false
	}

	stack := make([]int, len(pushed))
	p, top :=0,  -1
	for i := 0; i < len(pushed); i++ {
		top++
		stack[top] = pushed[i]
		for ; top != -1 && stack[top] == popped[p]; {
			top--
			p++
		}
	}
	return top == -1
}
