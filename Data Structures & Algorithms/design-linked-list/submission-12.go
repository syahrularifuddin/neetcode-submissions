type MyLinkedList struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	var i int
	var curr = this.Head
	for {
		if i == index {
			return curr.Val
		}
		if curr.Next == nil {
			break
		}
		curr = curr.Next
		i++
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	fmt.Println("add head", val)
	if this.Head == nil {
		this.Head = &Node{
			Val: val,
		}
		this.Tail = this.Head
		return
	}
	node := &Node{
		Val: val,
	}
	node.Next = this.Head
	this.Head.Prev = node
	this.Head = node
}

func (this *MyLinkedList) AddAtTail(val int) {
	fmt.Println("add tail", val)
	node := &Node{
		Val: val,
	}
    if this.Tail == nil || this.Head == nil {
        this.Tail = node
        this.Head = node
        return
    }
	node.Prev = this.Tail
	this.Tail.Next = node
	this.Tail = node
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	fmt.Println("add at index", index, val)

	if index == 0 {
		this.AddAtHead(val)
		return
	}

	var i int
	curr := this.Head
	for {
		if index == i {
			node := &Node{
				Val: val,
			}
			node.Next = curr
			node.Prev = curr.Prev
			curr.Prev.Next = node
			curr.Prev = node
			break
		}
		if curr.Next == nil {
			if index == i+1 {
				this.AddAtTail(val)
			}
			break
		}
		i++
		curr = curr.Next
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	fmt.Println("delete at index", index)
	// if index == 0 {
	// 	this.Head = this.Head.Next
	// 	this.Head.Prev = nil
	// 	return
	// }
	var i int
	curr := this.Head
    if curr == nil {
        return
    }
	for {
		// 1 -> 2 -> 3
		if i == index {
			if curr.Next == nil {
				this.Tail = curr.Prev
                if curr.Prev != nil {
				    curr.Prev.Next = nil
                }
				return
			}
            if curr.Prev == nil {
                this.Head = curr.Next
                this.Head.Prev = nil
                return
            }
			curr.Prev.Next = curr.Next
			curr.Next.Prev = curr.Prev
			return
		}
		if curr.Next == nil {
			break
		}
		i++
		curr = curr.Next
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
