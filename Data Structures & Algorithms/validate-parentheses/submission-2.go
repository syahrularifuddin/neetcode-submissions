func isValid(s string) bool {
    var stack []rune
    pairs := map[rune]rune{
        ')': '(',
        '}': '{',
        ']': '[',
    }
    for _,v := range s {
        if v == '(' || v == '{' || v == '[' {
            stack = append(stack, v)
            continue
        }
        if len(stack) == 0 {
            return false
        }
        popped := stack[len(stack)-1]
        if popped != pairs[v] {
            return false
        }
        stack = stack[:len(stack)-1]
    }
    return len(stack) == 0
}