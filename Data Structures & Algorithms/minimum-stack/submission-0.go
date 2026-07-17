type MinStack struct {
    stack []int
}


func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(value int)  {
    this.stack = append(this.stack, value)
}


func (this *MinStack) Pop()  {
    if len(this.stack) == 0 {
        return
    }
    this.stack = this.stack[:len(this.stack)-1]
}


func (this *MinStack) Top() int {
    if len(this.stack) == 0 {
        return 0
    }
    return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
    if len(this.stack) == 0 {
        return 0
    }
    var min = this.stack[0]
    for _,v:= range this.stack {
        if v < min {
            min = v
            continue
        }
    }
    return min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(value);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */