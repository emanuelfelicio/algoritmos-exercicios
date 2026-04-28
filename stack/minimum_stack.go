package algoritmosexercicios

type MinStack struct {
	a        []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{a: make([]int, 0), minStack: make([]int, 0)}
}

func (this *MinStack) Push(val int) {
	this.a = append(this.a, val)
	if len(this.minStack) > 0 {
		lastMin := this.minStack[len(this.minStack)-1]
		if val < lastMin {
			this.minStack = append(this.minStack, val)
		} else {
			this.minStack = append(this.minStack, lastMin)
		}
	} else {
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack) Pop() {
	this.a = this.a[:len(this.a)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.a[len(this.a)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
