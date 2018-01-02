package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	commonList = []rune("王李张刘陈杨黄赵吴周徐孙马朱胡郭何高林郑谢罗梁宋唐许韩冯邓曹彭曾肖田董袁潘于蒋蔡余杜叶程苏魏吕丁任沈姚卢姜崔钟谭陆汪范金石廖贾夏韦付方白邹孟熊秦邱江尹薛闫段雷侯龙史陶黎贺顾毛郝龚邵万钱严覃武戴莫孔向汤草发放受改卿取印即变文盛笔又盖第新目荣南古华十题颜可升都艺抒如部要术特未臆朴和阳至屈最阅有合属后难名朝展期版自同各宪撇家胸狂能官需被稿体作得但律态怀缘佼使约开式思连进还这应流楷运羲庄过迄迅楚意追浮迹感")
	rareList = []rune("块前骨战出起成处或夺毕准大止义们之正畅此泛令以法代仿指元也先略仍今介具其习关仅书挥汉他全八公商尖雕行权小将雅品就然表来尤而者导肥韵音端章竖状皆的密肃寄呵瘦富词识诞话四率现语询因国读粗说常长柳点规素见间染始写横由甲模画节再用贯醒祖奔生欲没种上欧演与不欢七一主为丰据情称捷中捺壮真历族字等省整原日无早故化北时学到源篆确造刷制承功力别划顺动那列创则初分切所速才通溯手认已己记论工年并广唯易键喜明果显黑是系构")
)


func convSlice(fromSlice []rune) (toSlice []interface{}) {
	toSlice = make([]interface{}, len(fromSlice))
	for index, value := range fromSlice {
		toSlice[index] = string(value)
	}
	return
}

// 打乱元素次序
func shuffleSlice(rawSlice []interface{}) (s []interface{}) {
	s = make([]interface{}, len(rawSlice))
	rand.Seed(int64(time.Now().Nanosecond()))
	for rawIndex, index := range rand.Perm(len(rawSlice)) {
		s[index] = rawSlice[rawIndex]
	}
	return
}


func main() {

	commonSlice := shuffleSlice(convSlice(commonList))
	rareSlice := shuffleSlice(convSlice(rareList))

	 //用于确定是否为常用姓
	commonSet := NewSetFromList(commonSlice)

	// 新建 PokerGroup 20 个 poker
	pg1 := NewPokerGroup()
	pg1ValueList := make([][]interface{}, 20)

	for pokerIndex:=0; pokerIndex<20; pokerIndex++ { // FIXME 这里写的有问题

		valueList := make([]interface{}, 20)
		for index, value := range commonSlice[pokerIndex*10:pokerIndex*10+10] {
			valueList[index] = value
		}

		for index, value := range rareSlice[pokerIndex*10:pokerIndex*10+10] {
			valueList[index+10] = value
		}

		pg1ValueList[pokerIndex] = valueList
		pg1.AddPoker(valueList)
	}

	// 新建一个 PokerGroup 10 个 poker
	pg2 := NewPokerGroup()

	// 从每个里边取出2个，包含一个常用的姓，一个不常用的姓
	for pokerIndex:=0; pokerIndex<10; pokerIndex++ {
		eleList := []interface{}{}
		for i:=0; i<20; i++ {
			eleList = append(eleList, pg1ValueList[i][pokerIndex])
			eleList = append(eleList, pg1ValueList[i][pokerIndex+10])
		}
		pg2.AddPoker(eleList)
	}

	// 从这里选出来一个
	fmt.Println(pg1.ToString())
	fmt.Println("选择你的姓所在的行号：")
	input1 := 1
	fmt.Scanln(&input1)

	// 从这里再选出来一个
	fmt.Println(pg2.ToString())
	fmt.Println("选择你的姓所在的行号：")
	input2 := 2
	fmt.Scanln(&input2)

	// 做一个交集
	inSet := IntersectSet(
		pg1.pokerSetList[input1-1],
		pg2.pokerSetList[input2-1],
	)

	for value := range inSet.IteratorChan() {
		if commonSet.Contains(value) {
			fmt.Println("\n\n你的姓是: ", value)
			break
		}
	}

}

