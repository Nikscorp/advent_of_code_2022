package days

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"time"
)

func Day20() int {
	reader := bufio.NewReader(os.Stdin)
	lst := list.New()
	toList := make([]*list.Element, 0)
	for {
		var num int
		_, err := fmt.Fscanf(reader, "%d\n", &num)
		if err != nil {
			break
		}
		num *= 811589153
		cur := lst.PushBack(num)
		toList = append(toList, cur)
	}

	for k := 0; k < 10; k++ {
		fmt.Println(k, time.Now())
		for i, tl := range toList {
			v := tl.Value.(int)
			bkp := v
			if v%(len(toList)-1) == 0 {
				continue
			}
			if v > 0 {
				dst := tl.Next()
				if dst == nil {
					dst = lst.Front()
				}
				lst.Remove(tl)
				v = v % (len(toList) - 1)

				for i := 0; i < v-1; i++ {
					if dst.Next() == nil {
						dst = lst.Front()
					} else {
						dst = dst.Next()
					}
				}

				newTL := lst.PushBack(bkp)
				lst.MoveAfter(newTL, dst)
				toList[i] = newTL
				continue
			}
			if v < 0 {
				v = -v
				dst := tl.Prev()
				if dst == nil {
					dst = lst.Back()
				}
				lst.Remove(tl)
				v = v % (len(toList) - 1)

				for i := 0; i < v-1; i++ {
					if dst.Prev() == nil {
						dst = lst.Back()
					} else {
						dst = dst.Prev()
					}
				}
				newTL := lst.PushBack(bkp)
				lst.MoveBefore(newTL, dst)
				toList[i] = newTL
				continue
			}
		}
	}

	cur := lst.Front()
	for cur != nil {
		if cur.Value == 0 {
			break
		}
		cur = cur.Next()
	}
	res := 0
	for i := 0; i <= 3000; i++ {
		if i == 1000 || i == 2000 || i == 3000 {
			fmt.Println(cur.Value.(int))
			res += cur.Value.(int)
		}
		if cur.Next() == nil {
			cur = lst.Front()
		} else {
			cur = cur.Next()
		}
	}

	return res
}
