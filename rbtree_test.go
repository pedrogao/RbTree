package rbtree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type key int

func (n key) LessThan(b interface{}) bool {
	value, _ := b.(key)
	return n < value
}

// Preorder prints the tree in pre order
func (t *Tree) Preorder() {
	fmt.Println("preorder begin!")
	if t.root != nil {
		t.root.preorder()
	}
	fmt.Println("preorder end!")
}

func (n *node) preorder() {
	fmt.Printf("(%v %v) ", n.Key, n.Value)
	if n.parent == nil {
		fmt.Printf("nil")
	} else {
		fmt.Printf("whose parent is %v", n.parent.Key)
	}
	if n.color == RED {
		fmt.Println(" and color RED")
	} else {
		fmt.Println(" and color BLACK")
	}
	if n.left != nil {
		fmt.Printf("%v's left child is ", n.Key)
		n.left.preorder()
	}
	if n.right != nil {
		fmt.Printf("%v's right child is ", n.Key)
		n.right.preorder()
	}
}

func TestPreorder(t *testing.T) {
	tree := NewTree()
	if !tree.Empty() {
		t.Error("tree not empty")
	}

	tree.Insert(key(1), "123")
	tree.Insert(key(3), "234")
	tree.Insert(key(4), "dfa3")
	tree.Insert(key(6), "sd4")
	tree.Insert(key(5), "jcd4")
	tree.Insert(key(2), "bcd4")
	if tree.Size() != 6 {
		t.Error("Error size")
	}
	if tree.Empty() {
		t.Error("tree empty")
	}
	tree.Preorder()
}

func TestFind(t *testing.T) {

	tree := NewTree()

	tree.Insert(key(1), "123")
	tree.Insert(key(3), "234")
	tree.Insert(key(4), "dfa3")
	tree.Insert(key(6), "sd4")
	tree.Insert(key(5), "jcd4")
	tree.Insert(key(2), "bcd4")

	n := tree.FindIt(key(4))
	if n.Value != "dfa3" {
		t.Error("Error value")
	}
	n.Value = "bdsf"
	if n.Value != "bdsf" {
		t.Error("Error value modify")
	}
	value := tree.Find(key(5)).(string)
	if value != "jcd4" {
		t.Error("Error value after modifyed other node")
	}
}
func TestIterator(t *testing.T) {
	tree := NewTree()

	tree.Insert(key(1), "123")
	tree.Insert(key(3), "234")
	tree.Insert(key(4), "dfa3")
	tree.Insert(key(6), "sd4")
	tree.Insert(key(5), "jcd4")
	tree.Insert(key(2), "bcd4")

	it := tree.Iterator()

	for it != nil {
		it = it.Next()
	}

}

func TestDelete(t *testing.T) {
	tree := NewTree()

	tree.Insert(key(1), "123")
	tree.Insert(key(3), "234")
	tree.Insert(key(4), "dfa3")
	tree.Insert(key(6), "sd4")
	tree.Insert(key(5), "jcd4")
	tree.Insert(key(2), "bcd4")
	for i := 1; i <= 6; i++ {
		tree.Delete(key(i))
		if tree.Size() != 6-i {
			t.Error("Delete Error")
		}
	}
	tree.Insert(key(1), "bcd4")
	tree.Clear()
	tree.Preorder()
	if tree.Find(key(1)) != nil {
		t.Error("Can't clear")
	}
}

func TestDelete2(t *testing.T) {
	tree := NewTree()
	tree.Insert(key(4), "1qa")
	tree.Insert(key(2), "2ws")
	tree.Insert(key(3), "3ed")
	tree.Insert(key(1), "4rf")
	tree.Insert(key(8), "5tg")
	tree.Insert(key(5), "6yh")
	tree.Insert(key(7), "7uj")
	tree.Insert(key(9), "8ik")
	tree.Delete(key(1))
	tree.Delete(key(2))
}

func TestCeil(t *testing.T) {
	assert := assert.New(t)
	tree := NewTree()

	tree.Insert(key(4), "1qa")
	tree.Insert(key(2), "2ws")
	tree.Insert(key(3), "3ed")
	tree.Insert(key(1), "4rf")
	tree.Insert(key(8), "5tg")
	tree.Insert(key(5), "6yh")
	tree.Insert(key(7), "7uj")
	tree.Insert(key(9), "8ik")
	tree.Insert(key(12), "9ik")
	v := tree.Ceil(key(6))
	assert.Equal(v, "7uj")
	v = tree.Ceil(key(10))
	assert.Equal(v, "9ik")
	v = tree.Ceil(key(11))
	assert.Equal(v, "9ik")
	v = tree.Ceil(key(8))
	assert.Equal(v, "5tg")
}

func TestFloor(t *testing.T) {
	assert := assert.New(t)
	tree := NewTree()

	tree.Insert(key(4), "1qa")
	tree.Insert(key(2), "2ws")
	tree.Insert(key(3), "3ed")
	tree.Insert(key(1), "4rf")
	tree.Insert(key(8), "5tg")
	tree.Insert(key(5), "6yh")
	tree.Insert(key(7), "7uj")
	tree.Insert(key(9), "8ik")
	v := tree.Floor(key(10))
	assert.Equal(v, "8ik")
	v = tree.Floor(key(11))
	assert.Equal(v, "8ik")
	v = tree.Floor(key(15))
	assert.Equal(v, "8ik")
	v = tree.Floor(key(7))
	assert.Equal(v, "7uj")
	v = tree.Floor(key(6))
	assert.Equal(v, "6yh")
}
