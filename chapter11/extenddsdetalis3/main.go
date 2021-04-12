package main

import "fmt"

type Goods struct {
	Name string
	Price float64
}

func (g *Goods) Show() {
	fmt.Printf("商品名：%v\n商品价格：%v\n",g.Name,g.Price)
}

type Brand struct {
	Name string
	Address string
}

func (b *Brand) Show() {
	fmt.Printf("品牌名：%v\n品牌产地：%v\n",b.Name,b.Address)
}

type TV struct {
	Goods
	Brand
}

type TV1 struct {
	*Goods
	*Brand
}

func main()  {
	// 第一种写法给结构体赋值，这种写法依赖于结构体字段顺序
	tv := TV{ Goods{"电视1",100.0},Brand{"长虹","中国"} }
	fmt.Println(tv)
	tv.Goods.Show()
	tv.Brand.Show()
	fmt.Println()
	
	// 第二种写法给结构体赋值，好处是不依赖于结构体中的字段顺序，这种写法是指定结构体字段名
	tv2 := TV{
		Goods{
			Name:"电视2",
			Price : 200.0,
		},
		Brand{
			Name:"海尔",
			Address:"山东",
		},
	}
	
	fmt.Println(tv2)
	tv2.Goods.Show()
	tv2.Brand.Show()
	fmt.Println()
	
	// 第三种写法直接取地址运算
	tv3 := TV1{ &Goods{"电视3",100.0},&Brand{"松下","日本"} }
	fmt.Println(*tv3.Goods,*tv3.Brand)
	tv3.Goods.Show()
	tv3.Brand.Show()
	fmt.Println()
	
	// 第四种写法直接取地址运算
	tv4 := TV1{
		&Goods{
			Name:"电视4",
			Price : 400.0,
		},
		&Brand{
			Name:"飞利浦",
			Address:"美国",
		},
	}
	fmt.Println(*tv4.Goods,*tv4.Brand)
	tv4.Goods.Show()
	tv4.Brand.Show()
}