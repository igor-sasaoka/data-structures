package list

// Package list implements a doubly linked list.

type Element struct {
    prev, next *Element
    list *List
    Value any
}

func (e *Element) Next() *Element { 
    if n := e.next; n != e && n != nil && n != &e.list.root {
        return n
    }
     return nil
}

func (e *Element) Prev() *Element {
    if p := e.prev; p != e && p != nil &&  p != &e.list.root {
        return p
    }

    return nil
}

// Here root is not a pointer, bc we should be able to start a empty list with zero value
// and it must be ready to use.
type List struct {
    root Element
    len int
}

// Initializes or clear a list, then returns it
func (l *List) Init() *List {
    l.root.next = &l.root
    l.root.prev = &l.root
    l.len = 0

    return l
}

// Returns a new initialized list
func New() *List { return new(List).Init() }

// Returns the lenght of the list
func (l *List) Len() int { return l.len }

// Returns the first element of the list
func (l *List) Front() *Element {
    if l.len == 0 {
        return nil
    }

    return l.root.next
}

// Returns the last element of the list
func (l *List) Back() *Element {
    if l.len == 0 {
        return nil
    }

    return l.root.prev
}

func (l *List) initIfEmpty() {
    if l.len == 0 {
        l.Init()
    }
}

// Inserts e after at
func (l *List) insertAfter (e, at *Element) *Element {
    e.prev = at
    e.next = at.next

    e.prev.next = e
    e.next.prev = e

    e.list = l
    l.len++

    return e
}

// Instantiates a new Element from the given value and inserts it after given Element
func (l *List) insertValueAfter(v any, at *Element) *Element {
    return l.insertAfter(&Element{Value: v}, at) 
}

func (l *List) remove(e *Element) {
    e.prev.next = e.next
    e.next.prev = e.prev

    e.next = nil // avoid memory leak 
    e.prev = nil // avoid memory leak
    e.list = nil
    l.len--
}

// Moves e to after at
func (l *List) move (e, at *Element) {
    if e == at {
        return
    }

    e.prev.next = e.next
    e.next.prev = e.prev

    e.prev = at
    e.next = at.next
    e.prev.next = e
    e.next.prev = e
}

func (l *List) Remove (e *Element) {
    if e.list == l {
        l.remove(e)
    }
}

func (l *List) PushFront (v any) *Element {
    l.initIfEmpty()
    return l.insertValueAfter(v, &l.root)
}

func (l *List) PushBack (v any) *Element {
    l.initIfEmpty()
    return l.insertValueAfter(v, l.root.prev)
}

func (l *List) Swap (a, b *Element) {
    if a.list != l || b.list != l {
        return
    }

    aPrev := a.prev
    aNext := a.next

    a.prev = b.prev
    a.next = b.next
    a.prev.next = a
    a.next.prev = a

    b.prev = aPrev
    b.next = aNext
    b.prev.next = b
    b.next.prev = b
}

func (l *List) MoveToFront (e *Element) {
    if e.list != l {
        return 
    }
    
    l.move(e, &l.root)
}

func (l *List) MoveToBack (e *Element) {
    if e.list != l {
        return
    }

    l.move(e, l.root.prev)
}

func (l *List) MoveAfter (a, b *Element) {
    if a.list != l || b.list != l || a == b {
        return
    }

    l.move(a, b)
}

func (l *List) MoveBefore (a, b *Element) {
    if a.list != l || b.list != l || a == b {
        return
    }

    l.move(a, b.prev)
}

func (l *List) PushListBack (list *List) {
    l.initIfEmpty()

    for i := list.Front(); i != nil; i = i.Next() {
        l.insertValueAfter(i.Value, l.root.prev)
    }
}

func (l *List) PushListFront (list *List) {
    l.initIfEmpty()

    for i := list.Back(); i != nil; i = i.Prev() {
        l.insertValueAfter(i.Value, &l.root)
    }
}
