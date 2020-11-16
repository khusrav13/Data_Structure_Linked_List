package main

import (
	"fmt"
	"strings"
)

type Node struct {
	previous *Node
	name     string
	priority int64
	next     *Node
}

type ListNode struct {
	name         string
	firstEl       *Node
	lastEl       *Node
	currentQueue *Node
	size         int
}

////Check if our list is empty or not
func (p *ListNode) IsEmpty() bool {
	if p.firstEl == nil {
		return true
	}
	return false
}

func (p *ListNode) SettingsOfUser() *Node {
	p.currentQueue = p.firstEl
	return p.currentQueue
}

func (p *ListNode) addUser(name string, priority int64) error {
	s := &Node{
		name:     name,
		priority: priority,
	}
	if p.firstEl == nil {
		p.firstEl = s
	} else {
		currentNode := p.lastEl
		currentNode.next = s
		s.previous = p.lastEl
	}
	p.lastEl = s
	return nil
}

func createTableOfUser(name string) *ListNode {
	return &ListNode{
		name: name,
	}
}


func (p *ListNode) showAllUsers() error {
	currentNode := p.firstEl
	if currentNode == nil {
		fmt.Println("Cannot find any user")
		return nil
	}
	fmt.Printf("%+v\n", *currentNode)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", *currentNode)
	}
	return nil
}

///NEXT USER
func (p *ListNode) nextUser() *Node {
	p.currentQueue = p.currentQueue.next
	return p.currentQueue
}

///PREVIOUS USER
func (p *ListNode) previousUser() *Node {
	p.currentQueue = p.currentQueue.previous
	return p.currentQueue
}

////Size of Users (length)
func (p *ListNode) sizeOfUser() int {

	size := 1
	last := p.firstEl
	for {
		if last == nil || last.next == nil {
			break
		}
		last = last.next
		size++
	}
	return size
}


func main() {
	dashes := strings.Repeat("-", 65)
	todayUsers := "TodayQueue"
	Users := createTableOfUser(todayUsers)
	fmt.Println("The table has been created!")
	fmt.Println()

	fmt.Print(dashes, "\n\n")
	_ = Users.addUser("Ansor", 1)
	_ = Users.addUser("Siyovush", 2)
	_ = Users.addUser("Komron", 3)
	_ = Users.addUser("Alisher", 4)
	_ = Users.addUser("khurshed", 5)
	fmt.Println("User names and priority of the users:")
	_ = Users.showAllUsers()
	fmt.Println(dashes, "\n\n")

	Users.SettingsOfUser()
	fmt.Printf("Our first user is: %s and his/her queue is %d\n", Users.currentQueue.name, Users.currentQueue.priority)
	fmt.Println()

	Users.nextUser()
	fmt.Println("Next user is...")
	fmt.Printf("Username: %s and his/her queue is %d\n", Users.currentQueue.name, Users.currentQueue.priority)
	fmt.Println()
	Users.nextUser()
	fmt.Println("Next user is...")
	fmt.Printf("Username: %s and his/her queue is %d\n", Users.currentQueue.name, Users.currentQueue.priority)
	fmt.Println()

	fmt.Println("Komron wants to know who is behind him.")
	fmt.Println()

	Users.previousUser()
	fmt.Println("Previous user is...")
	fmt.Printf("Username: %s and his/her queue is %d\n", Users.currentQueue.name, Users.currentQueue.priority)
	fmt.Println()
	//Users.previousUser()
	//fmt.Println("Previous user is....")
	//fmt.Printf("Username: %s and his/her queue is %d\n", Users.currentQueue.name, Users.currentQueue.priority)

	fmt.Print(dashes, "\n\n")

	fmt.Println("Number of users", Users.sizeOfUser())
	fmt.Print(dashes, "\n\n")

	fmt.Println("Is our list empty??", Users.IsEmpty())
	fmt.Print(dashes, "\n\n")

}
