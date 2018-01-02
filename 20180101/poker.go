package main

import (
	"fmt"
)

/*
   PokerGroup 代表一组牌，每张牌上若干个姓。每张牌宽高固定
 */
type PokerGroup struct {
	size int            // 有几个set
	pokerSetList []*Set    // 每张牌上的姓，组成一个集合
}

func NewPokerGroup() *PokerGroup {
	return &PokerGroup{
		size: 0,
		pokerSetList: []*Set{},
	}
}

// 新加一张牌
func (pg *PokerGroup) AddPoker(lst []interface{}) {
	pg.size++
	pg.pokerSetList = append(pg.pokerSetList, NewSetFromList(lst))
}

func (pg *PokerGroup) Size() int {
	return pg.size
}

// 每行打印4张牌
func (pg *PokerGroup) ToString() (s string) {
	s = fmt.Sprintf("PokerGroup Size: %d\n", pg.Size())
	for pokerIndex:=0; pokerIndex<pg.Size(); pokerIndex++ {
		s += fmt.Sprintf("%2d : %s\n", pokerIndex+1, pg.pokerSetList[pokerIndex].ToLineString())
	}
	s += "\n"
	return
}
