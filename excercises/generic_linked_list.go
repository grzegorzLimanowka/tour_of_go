package excercises

import "fmt"

// List represents singly-linked list that holds values of any type
type List[T comparable] struct {
	tail *Node[T]
}

// Adds new element to the list
func (l *List[T]) Add(v T) *List[T] {
	node := Node[T]{nil, v}

	if l.tail == nil {
		l.tail = &node

		fmt.Printf("Tail is %v, prev is nil \n", l.tail.val)
		return l
	}

	node.next = l.tail
	l.tail = &node

	fmt.Printf("Tail is %v, prev is %v \n", l.tail.val, l.tail.next.val)
	return l
}

// Show all elements in list
func (l *List[T]) Show() {
	fmt.Print("[ ")

	if l.tail == nil {
		fmt.Println("Empty list ")
		return
	}

	for curr := l.tail; curr != nil; {
		fmt.Printf("%v", curr.val)
		if curr.next != nil {
			fmt.Printf("(-> %v) ", curr.next.val)
		} else {
			fmt.Printf("(-> Nil) ")
		}

		curr = curr.next
	}

	fmt.Println("]")
}

// Remove from list first encountered element with matching value
// [1, 2, 3, 4, 5].remove(2) -> [1, 3, 4, 5].remove(6) -> error // TODO: return proper error
func (l *List[T]) Remove(v T) (T, error) { // TODO: Make returned val generic
	if l.tail == nil {
		fmt.Println("Empty list, no elements to remove")

		return v, nil // TODO: return error
	}

	fmt.Printf("Removing elem: %v \n", v)

	// Case: 1 elem list:
	if l.tail.val == v {
		if l.tail.next == nil {
			l.tail = nil

			return v, nil
		}

		if l.tail.next != nil {
			l.tail = l.tail.next

			return v, nil
		}
	}

	// Case: >1 elem list:
	for curr := l.tail; curr.next != nil; {
		if curr.next.val == v {
			if curr.next.next != nil {
				curr.next = curr.next.next
			} else {
				curr.next = nil
			}

			return v, nil
		}

		curr = curr.next
	}

	return v, nil // TODO: return error, no elems removed
}

// Remove from list element with given index
// [1, 2, 3, 4, 5].remove(3) -> [1, 2, 3, 5].remove(6) -> error // TODO: return proper error
func (l *List[T]) Remove_nth(val int) (int, error) {
	// TODO: impl

	return -1, nil
}

// show elements of list from i idx to j idx, for example
// [1, 2, 3, 4, 5].sublist(1, 3) -> [2, 3, 4]
func (l *List[T]) Sublist(from, to int) (int, error) {
	// TODO: impl

	return -1, nil
}

// remove multiple elements from list with 1 bulk remove
// [1, 2, 3, 4, 5].remove_multiple([1, 3, 5]) -> [2, 4]
func (l *List[T]) Remove_multiple(v []T) (int, error) {
	// TODO: impl

	return -1, nil
}

func (l *List[T]) Len() (len int) {
	if l.tail == nil {
		return 0
	}

	for curr := l.tail; curr != nil; {
		len += 1

		curr = curr.next
	}

	return
}

type Node[T interface{}] struct {
	next *Node[T]
	val  T
}

func main() {
	list := List[int]{nil}

	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Show()
	fmt.Println(list.Len())

	list.Remove(1)
	list.Show()

	list.Remove(3)
	list.Show()

	list.Remove(4)
	list.Show()

	list.Remove(2)
	list.Show()

	// List 2

	list2 := List[string]{nil}
	list2.Add("sdsfd")
	list2.Add("aaa")
	list2.Add("v33v")
	list.Show()
	fmt.Println(list2.Len())

	// List 3

	list3 := List[float64]{nil}
	list3.Add(5)
	list3.Add(4)
	list3.Add(3)
	list3.Add(2)
	list3.Add(1)

	fmt.Println(list3.Len())

	// list3.remove_nth(1)
}

// Remove element with given index (calculated form head)
// It shold return error, when for example idx is out of list bounds
