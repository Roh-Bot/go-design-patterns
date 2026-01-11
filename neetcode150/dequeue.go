package neetcode150

// Dequeue Helper Implementation
// A lightweight deque using slice operations to store indices.
type Dequeue []int

func (d *Dequeue) PopEnd()        { *d = (*d)[:len(*d)-1] } // remove last element
func (d *Dequeue) PopStart()      { *d = (*d)[1:] }         // remove first element
func (d *Dequeue) PushEnd(v int)  { *d = append(*d, v) }    // append new index at end
func (d *Dequeue) PeekEnd() int   { return (*d)[len(*d)-1] }
func (d *Dequeue) PeekStart() int { return (*d)[0] }
func (d *Dequeue) Length() int    { return len(*d) }
