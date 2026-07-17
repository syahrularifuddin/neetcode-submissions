type BrowserHistory struct {
    cur *Page
}

type Page struct {
    url string
    next *Page
    prev *Page
}


func Constructor(homepage string) BrowserHistory {
    return BrowserHistory{
        cur: &Page{
            url: homepage,
        },
    }
}


func (this *BrowserHistory) Visit(url string)  {
    cur := this.cur
    cur.next = &Page{
        url: url,
        prev: cur,
    }
    this.cur = cur.next
}


func (this *BrowserHistory) Back(steps int) string {
    cur := this.cur
    for i:=0;i<steps;i++{
        if cur.prev == nil {
            this.cur = cur
            return this.cur.url
        }
        cur = cur.prev
    }
    this.cur = cur
    return this.cur.url
}


func (this *BrowserHistory) Forward(steps int) string {
    cur := this.cur
    for i:=0;i<steps;i++{
        if cur.next == nil {
            this.cur = cur
            return this.cur.url
        }
        cur = cur.next
    }
    this.cur = cur
    return this.cur.url
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */