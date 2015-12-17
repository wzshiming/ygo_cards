package ygo_cards

import ygo "github.com/wzshiming/ygo_core"

func d_1(cardBag *ygo.CardVersion) {

	/* 0 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 159
		 调整:

		 [传说之剑]<伝説の剣>
		 [2010/09/08]
		 ●战士族怪兽才能装备。1只装备怪兽的攻击力及守备力上升300分。
		 ◇发动时选择场上1只符合条件的怪兽（取对象）
		 ◇效果处理时对象怪兽不是战士族的场合，这张卡送去墓地
		 ◇效果适用中对象怪兽变成战士族以外的种族的场合，这张卡破坏

		 中文名: 传说之剑
		 卡片种类: 装备魔法
		 卡片密码: 61854111
		 罕见度: 平卡N，银字R
		 卡包: LB、VOL01、DL02、Booster01
		 效果: 战士族才能装备。1只装备怪兽的攻击力和守备力上升300。

		*/
		Id:       159,
		Password: "61854111",
		Name:     "传说之剑",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellEquipEffect1(func(c *ygo.Card) bool {
				return c.RaceIsWarrior()
			}, func(c *ygo.Card) {
				c.AddAtk(300)
				c.AddDef(300)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 1 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 160
		 调整:

		 [猛兽之齿]<猛獣の歯>
		 [2010/09/10]
		 ●兽族才能装备。1只装备怪兽的攻击力与守备力上升300分。
		 ◇发动时选择1只符合条件的怪兽（取对象）
		 ◇效果处理时装备怪兽种族变为兽族以外的场合，这张卡送去墓地
		 ◇装备状态中对象怪兽变为兽族以外的场合，这张卡破坏

		 中文名: 猛兽之齿
		 卡片种类: 装备魔法
		 卡片密码: 46009906
		 罕见度: 平卡N，银字R
		 卡包: LB、VOL01、DL02、Booster01
		 效果: 兽族才能装备。1只装备怪兽的攻击力和守备力上升300。

		*/
		Id:       160,
		Password: "46009906",
		Name:     "猛兽之齿",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellEquipEffect1(func(c *ygo.Card) bool {
				return c.RaceIsBeast()
			}, func(c *ygo.Card) {
				c.AddAtk(300)
				c.AddDef(300)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 2 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 161
		 调整:

		 [紫水晶]<紫水晶>
		 [2010/09/10]
		 ●不死族才能装备。1只装备怪兽的攻击力与守备力上升300分。
		 ◇发动时选择1只符合条件的怪兽（取对象）
		 ◇效果处理时装备怪兽种族变为不死族以外的场合，这张卡送去墓地
		 ◇装备状态中对象怪兽变为不死族以外的场合，这张卡破坏

		 中文名: 紫水晶
		 卡片种类: 装备魔法
		 卡片密码: 15052462
		 罕见度: 平卡N，银字R
		 卡包: LB、VOL01、DL02、Booster01
		 效果: 不死族才能装备。1只装备怪兽的攻击力和守备力上升300。

		*/
		Id:       161,
		Password: "15052462",
		Name:     "紫水晶",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellEquipEffect1(func(c *ygo.Card) bool {
				return c.RaceIsZombie()
			}, func(c *ygo.Card) {
				c.AddAtk(300)
				c.AddDef(300)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 3 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 162
		 调整:

		 [秘术之书]<秘術の書>
		 [2010/09/08]
		 ●魔法师族怪兽才能装备。1只装备怪兽的攻击力及守备力上升300分。
		 ◇发动时选择场上1只符合条件的怪兽（取对象）
		 ◇效果处理时对象怪兽不是魔法师族的场合，这张卡送去墓地
		 ◇效果适用中对象怪兽变成魔法师族以外的种族的场合，这张卡破坏

		 中文名: 秘术之书
		 卡片种类: 装备魔法
		 卡片密码: 91595718
		 罕见度: 平卡N，银字R
		 卡包: LB、EX-R(EX)、VOL01、DL02、Booster01
		 效果: 魔法师族才能装备。1只装备怪兽的攻击力和守备力上升300。

		*/
		Id:       162,
		Password: "91595718",
		Name:     "秘术之书",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellEquipEffect1(func(c *ygo.Card) bool {
				return c.RaceIsSpellcaster()
			}, func(c *ygo.Card) {
				c.AddAtk(300)
				c.AddDef(300)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 4 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 163
		 调整:

		 [波塞冬之力]<ポセイドンの力>
		 [2010/09/10]
		 ●水族才能装备。1只装备怪兽的攻击力与守备力上升300分。
		 ◇发动时选择1只符合条件的怪兽（取对象）
		 ◇效果处理时装备怪兽种族变为水族以外的场合，这张卡送去墓地
		 ◇装备状态中对象怪兽变为水族以外的场合，这张卡破坏

		 中文名: 波塞冬之力
		 卡片种类: 装备魔法
		 卡片密码: 77027445
		 罕见度: 平卡N，银字R
		 卡包: LB、VOL01、DL02、Booster01
		 效果: 水族才能装备。1只装备怪兽的攻击力和守备力上升300。

		*/
		Id:       163,
		Password: "77027445",
		Name:     "波塞冬之力",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellEquipEffect1(func(c *ygo.Card) bool {
				return c.RaceIsWater()
			}, func(c *ygo.Card) {
				c.AddAtk(300)
				c.AddDef(300)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 5 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 194
		 调整:

		 [暗能量]<闇·エネルギー>
		 [2010/09/08]
		 ●恶魔族怪兽才能装备。1只装备怪兽的攻击力及守备力上升300分。
		 ◇发动时选择场上1只符合条件的怪兽（取对象）
		 ◇效果处理时对象怪兽不是恶魔族的场合，这张卡送去墓地
		 ◇效果适用中对象怪兽变成恶魔族以外的种族的场合，这张卡破坏

		 中文名: 暗能量
		 卡片种类: 装备魔法
		 卡片密码: 04614116
		 罕见度: 平卡N
		 卡包: PG、EX-R(EX)、VOL02、DL02、Booster R2
		 效果: 恶魔族才能装备。1只装备怪兽的攻击力·守备力上升300。

		*/
		Id:       194,
		Password: "04614116",
		Name:     "暗能量",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellEquipEffect1(func(c *ygo.Card) bool {
				return c.RaceIsFiend()
			}, func(c *ygo.Card) {
				c.AddAtk(300)
				c.AddDef(300)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 6 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 195
		 调整:

		 [镭射炮机甲铠]<レーザー砲機甲鎧>
		 [2010/09/08]
		 ●昆虫族怪兽才能装备。1只装备怪兽的攻击力及守备力上升300分。
		 ◇发动时选择场上1只符合条件的怪兽（取对象）
		 ◇效果处理时对象怪兽不是昆虫族的场合，这张卡送去墓地
		 ◇效果适用中对象怪兽变成昆虫族以外的种族的场合，这张卡破坏

		 中文名: 镭射炮机甲铠
		 卡片种类: 装备魔法
		 卡片密码: 77007920
		 罕见度: 平卡N
		 卡包: PG、VOL02、DL02、TP19
		 效果: 昆虫族才能装备。1只装备怪兽的攻击力·守备力上升300。

		*/
		Id:       195,
		Password: "77007920",
		Name:     "镭射炮机甲铠",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellEquipEffect1(func(c *ygo.Card) bool {
				return c.RaceIsInsect()
			}, func(c *ygo.Card) {
				c.AddAtk(300)
				c.AddDef(300)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 7 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 196
		 调整:

		 [魔菌]<魔菌>
		 [2010/09/08]
		 ●植物族怪兽才能装备。1只装备怪兽的攻击力及守备力上升300分。
		 ◇发动时选择场上1只符合条件的怪兽（取对象）
		 ◇效果处理时对象怪兽不是植物族的场合，这张卡送去墓地
		 ◇效果适用中对象怪兽变成植物族以外的种族的场合，这张卡破坏

		 中文名: 魔菌
		 卡片种类: 装备魔法
		 卡片密码: 39774685
		 罕见度: 平卡N
		 卡包: PG、VOL02、DL02
		 效果: 植物族才能装备。1只装备怪兽的攻击力·守备力上升300。

		*/
		Id:       196,
		Password: "39774685",
		Name:     "魔菌",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellEquipEffect1(func(c *ygo.Card) bool {
				return c.RaceIsPlant()
			}, func(c *ygo.Card) {
				c.AddAtk(300)
				c.AddDef(300)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 8 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 197
		 调整:

		 [银之弓矢]<銀の弓矢>
		 [2010/09/08]
		 ●天使族怪兽才能装备。1只装备怪兽的攻击力及守备力上升300分。
		 ◇发动时选择场上1只符合条件的怪兽（取对象）
		 ◇效果处理时对象怪兽不是天使族的场合，这张卡送去墓地
		 ◇效果适用中对象怪兽变成天使族以外的种族的场合，这张卡破坏

		 中文名: 银之弓矢
		 卡片种类: 装备魔法
		 卡片密码: 01557499
		 罕见度: 平卡N
		 卡包: PG、VOL03、DL02、TP15
		 效果: 天使族怪兽才能装备。装备怪兽的攻击力·守备力上升300。

		*/
		Id:       197,
		Password: "01557499",
		Name:     "银之弓矢",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellEquipEffect1(func(c *ygo.Card) bool {
				return c.RaceIsFairy()
			}, func(c *ygo.Card) {
				c.AddAtk(300)
				c.AddDef(300)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 9 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 198
		 调整:

		 [龙之秘宝]<ドラゴンの秘宝>
		 [2010/09/08]
		 ●龙族怪兽才能装备。1只装备怪兽的攻击力及守备力上升300分。
		 ◇发动时选择场上1只符合条件的怪兽（取对象）
		 ◇效果处理时对象怪兽不是龙族的场合，这张卡送去墓地
		 ◇效果适用中对象怪兽变成龙族以外的种族的场合，这张卡破坏

		 中文名: 龙之秘宝
		 卡片种类: 装备魔法
		 卡片密码: 01435851
		 罕见度: 平卡N
		 卡包: PG、VOL03、DL02、Booster R2
		 效果: 龙族才能装备。1只装备怪兽的攻击力·守备力上升300。

		*/
		Id:       198,
		Password: "01435851",
		Name:     "龙之秘宝",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellEquipEffect1(func(c *ygo.Card) bool {
				return c.RaceIsDragon()
			}, func(c *ygo.Card) {
				c.AddAtk(300)
				c.AddDef(300)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 10 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 199
		 调整:

		 [电击鞭]<電撃鞭>
		 [2010/09/08]
		 ●雷族怪兽才能装备。1只装备怪兽的攻击力及守备力上升300分。
		 ◇发动时选择场上1只符合条件的怪兽（取对象）
		 ◇效果处理时对象怪兽不是雷族的场合，这张卡送去墓地
		 ◇效果适用中对象怪兽变成雷族以外的种族的场合，这张卡破坏

		 中文名: 电击鞭
		 卡片种类: 装备魔法
		 卡片密码: 37820550
		 罕见度: 平卡N
		 卡包: PG、VOL03、DL02
		 效果: 雷族才能装备。1只装备怪兽的攻击力·守备力上升300。

		*/
		Id:       199,
		Password: "37820550",
		Name:     "电击鞭",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellEquipEffect1(func(c *ygo.Card) bool {
				return c.RaceIsThunder()
			}, func(c *ygo.Card) {
				c.AddAtk(300)
				c.AddDef(300)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 11 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 200
		 调整:

		 [魔性之月]<魔性の月>
		 [2010/09/08]
		 ●兽战士族怪兽才能装备。1只装备怪兽的攻击力及守备力上升300分。
		 ◇发动时选择场上1只符合条件的怪兽（取对象）
		 ◇效果处理时对象怪兽不是兽战士族的场合，这张卡送去墓地
		 ◇效果适用中对象怪兽变成兽战士族以外的种族的场合，这张卡破坏

		 中文名: 魔性之月
		 卡片种类: 装备魔法
		 卡片密码: 36607978
		 罕见度: 平卡N
		 卡包: PG、VOL03、DL02
		 效果: 兽战士族才能装备。1只装备怪兽的攻击力·守备力上升300。

		*/
		Id:       200,
		Password: "36607978",
		Name:     "魔性之月",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellEquipEffect1(func(c *ygo.Card) bool {
				return c.RaceIsBeastWarrior()
			}, func(c *ygo.Card) {
				c.AddAtk(300)
				c.AddDef(300)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 12 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 202
		 调整:

		 [机械改造工厂]<機械改造工場>
		 [2010/09/08]
		 ●机械族怪兽才能装备。1只装备怪兽的攻击力及守备力上升300分。
		 ◇发动时选择场上1只符合条件的怪兽（取对象）
		 ◇效果处理时对象怪兽不是机械族的场合，这张卡送去墓地
		 ◇效果适用中对象怪兽变成机械族以外的种族的场合，这张卡破坏

		 中文名: 机械改造工厂
		 卡片种类: 装备魔法
		 卡片密码: 25769732
		 罕见度: 平卡N
		 卡包: PG、VOL02、DL02
		 效果: 只有机械族怪兽可以装备。装备怪兽的攻击力·守备力上升300。

		*/
		Id:       202,
		Password: "25769732",
		Name:     "机械改造工厂",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellEquipEffect1(func(c *ygo.Card) bool {
				return c.RaceIsMachine()
			}, func(c *ygo.Card) {
				c.AddAtk(300)
				c.AddDef(300)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 13 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 203
		 调整:

		 [体温上升]<体温の上昇>
		 [2010/09/08]
		 ●恐龙族怪兽才能装备。1只装备怪兽的攻击力及守备力上升300分。
		 ◇发动时选择场上1只符合条件的怪兽（取对象）
		 ◇效果处理时对象怪兽不是恐龙族的场合，这张卡送去墓地
		 ◇效果适用中对象怪兽变成恐龙族以外的种族的场合，这张卡破坏

		 中文名: 体温上升
		 卡片种类: 装备魔法
		 卡片密码: 51267887
		 罕见度: 平卡N
		 卡包: PG、VOL02、DL02
		 效果: 只有恐龙族怪兽可以装备。装备怪兽的攻击力·守备力上升300。

		*/
		Id:       203,
		Password: "51267887",
		Name:     "体温上升",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellEquipEffect1(func(c *ygo.Card) bool {
				return c.RaceIsDinosaur()
			}, func(c *ygo.Card) {
				c.AddAtk(300)
				c.AddDef(300)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 14 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 204
		 调整:

		 [随风翼]<フォロー·ウィンド>
		 [2010/09/08]
		 ●鸟兽族怪兽才能装备。1只装备怪兽的攻击力及守备力上升300分。
		 ◇发动时选择场上1只符合条件的怪兽（取对象）
		 ◇效果处理时对象怪兽不是鸟兽族的场合，这张卡送去墓地
		 ◇效果适用中对象怪兽变成鸟兽族以外的种族的场合，这张卡破坏

		 中文名: 随风翼
		 卡片种类: 装备魔法
		 卡片密码: 98252586
		 罕见度: 平卡N
		 卡包: PG、VOL03、DL02
		 效果: 只有鸟兽族怪兽可以装备。装备怪兽的攻击力·守备力上升300。

		*/
		Id:       204,
		Password: "98252586",
		Name:     "随风翼",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellEquipEffect1(func(c *ygo.Card) bool {
				return c.RaceIsWingedBeast()
			}, func(c *ygo.Card) {
				c.AddAtk(300)
				c.AddDef(300)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 15 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 468
		 调整:

		 [执念之剑]<執念の剣>
		 [2010/09/08]
		 ●装备怪兽的攻击力·守备力上升500分。
		 ◇不进入连锁
		 ●这张卡送去墓地时，这张卡回到卡组的最上方。
		 ◇进入连锁（1速）
		 ◇必须发动
		 ◇效果处理时这张卡不在墓地存在的场合，这个效果不适用

		 中文名: 执念之剑
		 卡片种类: 装备魔法
		 卡片密码: 98495314
		 罕见度: 平卡N
		 卡包: ME、VOL07、DL04、TP13
		 效果: 装备的怪兽攻击力·守备力上升500。这张卡送去墓地时，回到卡组最上面。

		*/
		Id:       468,
		Password: "98495314",
		Name:     "执念之剑",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellEquipEffect1(func(c *ygo.Card) bool {
				return true
			}, func(c *ygo.Card) {
				c.AddAtk(500)
				c.AddDef(500)
			})
			ca.AddEvent(ygo.InGrave, ca.ToDeckTop)
			return true
		}, // 初始
		IsValid: true,
	})

	/* 16 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 472
		 调整:

		 [细菌感染]<細菌感染>
		 [2010/09/08]
		 ●机械族以外的怪兽才能装备。装备怪兽的攻击力在每次自己的准备阶段下降300分。
		 ◇发动时选择场上表侧表示存在的机械族以外的怪兽（取对象）
		 ◇效果处理时对象怪兽变为机械族的场合，这张卡送去墓地
		 ◇效果适用中对象怪兽变为机械族的场合，这张卡破坏
		 ◇「王宮の勅命/王宮の勅命」效果适用不影响这张卡的回合计算，在「王宮の勅命/王宮の勅命」效果失去后，装备怪兽降低那个回合数×300分的攻击力

		 中文名: 细菌感染
		 卡片种类: 装备魔法
		 卡片密码: 24668830
		 罕见度: 平卡N
		 卡包: ME、VOL07、DL04、Booster07
		 效果: 机械族以外的怪兽装备可能。装备怪兽的攻击力在每次的自己的准备阶段攻击力下降300。

		*/
		Id:       472,
		Password: "24668830",
		Name:     "细菌感染",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.Counter = 0
			ca.RegisterSpellEquipEffect2(func(c *ygo.Card) bool {
				return !c.RaceIsMachine()
			}, func(c *ygo.Card) {
				c.AddAtk(ca.Counter * -300)
				c.AddDef(ca.Counter * -300)
			}, ygo.SP, func(c *ygo.Card) {
				p := c.GetSummoner()
				if !p.IsCurrent() {
					return
				}
				ca.Counter++
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 17 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 473
		 调整:

		 [麻药]<しびれ薬>
		 [2010/09/08]
		 ●机械族以外的怪兽才能装备。装备怪兽不能攻击宣言。
		 ◇发动时选择场上表侧表示存在的机械族以外的怪兽（取对象）
		 ◇效果处理时对象怪兽变为机械族的场合，这张卡送去墓地
		 ◇效果适用中对象怪兽变为机械族的场合，这张卡破坏

		 中文名: 麻药
		 卡片种类: 装备魔法
		 卡片密码: 50152549
		 罕见度: 平卡N
		 卡包: ME、VOL07、DL04、Booster07
		 效果: 机械族以外的怪兽装备可能。装备怪兽不能攻击宣言。

		*/
		Id:       473,
		Password: "50152549",
		Name:     "麻药",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.Counter = 0
			ca.RegisterSpellEquipEffect2(func(c *ygo.Card) bool {
				return !c.RaceIsMachine()
			}, func(c *ygo.Card) {

			}, ygo.BP, func(c *ygo.Card) {
				p := c.GetSummoner()
				if !p.IsCurrent() {
					return
				}
				c.SetNotCanAttack()
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 18 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 171
		 调整:

		 [黑洞]<ブラック·ホール>
		 [2010/09/08]
		 ●场上存在的全部怪兽破坏。
		 ◇不取对象

		 中文名: 黑洞
		 卡片种类: 通常魔法
		 卡片密码: 53129443
		 罕见度: 平卡N，黄金GR，面闪SR，金字UR
		 卡包: LB、BE01、EX-R(EX)、VOL01、DL02、YU、JY、KA、TU04、TU05、GS03、SD23、DS13、GDB1
		 效果: 场上存在的怪兽全部破坏。

		*/
		Id:       171,
		Password: "53129443",
		Name:     "黑洞",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellNormalPuah(ygo.LO_DestroyAllMonster, func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				ygo.NewCards(pl.Mzone(), tar.Mzone()).ForEach(func(c *ygo.Card) bool {
					c.Destroy(ca)
					return true
				})
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 19 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 173
		 调整:

		 [红色药剂]<レッド·ポーション>
		 [2010/09/08]
		 ●自己回复500基本分。
		 ◇「精灵之镜/精霊の鏡」可以对应这张卡的发动而发动

		 中文名: 红色药剂
		 卡片种类: 通常魔法
		 卡片密码: 38199696
		 罕见度: 平卡N
		 卡包: LB、BE01、YSD01、VOL01、DL02
		 效果: 自己回复500基本分。

		*/
		Id:       173,
		Password: "38199696",
		Name:     "红色药剂",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellNormalPuah(ygo.LO_Reply, func() {
				pl := ca.GetSummoner()
				pl.ChangeLp(500)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 20 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 174
		 调整:

		 [火星]<火の粉>
		 [2010/09/08]
		 ●给予对方基本分200分伤害。
		 ◇「精灵之镜/精霊の鏡」可以对应这张卡的发动而发动

		 中文名: 火星
		 卡片种类: 通常魔法
		 卡片密码: 76103675
		 罕见度: 平卡N
		 卡包: LB、VOL01、DL02
		 效果: 给予对方基本分200分伤害。

		*/
		Id:       174,
		Password: "76103675",
		Name:     "火星",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellNormalPuah(ygo.LO_Hurt, func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				tar.ChangeLp(-200)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 21 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 176
		 调整:

		 [地割]<地割れ>
		 [2010/09/08]
		 ●对方场上表侧表示存在的攻击力最低的1只怪兽破坏。
		 ◇在这张卡的效果处理时判断“对方场上表侧表示存在的攻击力最低的1只怪兽”（不取对象）
		 ◇效果处理时，若符合条件的怪兽有2只以上存在，这张卡的发动者在这时选择其中之一破坏

		 中文名: 地割
		 卡片种类: 通常魔法
		 卡片密码: 66788016
		 罕见度: 平卡N，银字R，黄金GR
		 卡包: LB、BE01、BE01、EX-R(EX)、YSD01、VOL01、DL02、SD16、GS02、YU、PE、SY2、YSD06、DB12、ST12、DS14、GDB1
		 效果: 对方场上表侧表示存在的1只攻击力最低的怪兽破坏。

		*/
		Id:       176,
		Password: "66788016",
		Name:     "地割",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellNormal(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				i := -1
				css := ygo.NewCards(tar.Mzone(), func(c *ygo.Card) bool {
					if i == -1 || i > c.GetAtk() {
						i = c.GetAtk()
					}
					return c.IsFaceUp()
				})
				if css.Len() == 0 {
					return
				}
				ca.PushSpell(ygo.LO_DestroyAMonster, func() {
					if c := pl.SelectRequiredShor(ygo.LO_Destroy, css, func(c *ygo.Card) bool {
						return c.GetAtk() == i
					}); c != nil {
						c.Discard(ca)
					}
				})
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 22 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 178
		 调整:

		 [融合]<融合>
		 [2010/09/08]
		 ●从手卡·自己场上，把融合怪兽卡决定的融合素材怪兽送去墓地，从额外卡组特殊召唤1只融合怪兽。
		 ◇可以只用手卡的或只用自己场上的融合素材怪兽，也可以结合起来用
		 ◇「王宫的弹压/王宮の弾圧」只能对应这张卡的发动而发动，不能对应这张卡效果处理时的特殊召唤而发动
		 ◇效果处理时选择作为融合素材的怪兽送去墓地
		 ◇效果处理时，融合怪兽的融合素材不足时，这个效果不适用
		 ◇承上述，那个场合对方可以要求查看自己的手牌、场上覆盖的怪兽、额外卡组
		 ◇这张卡的发动被无效的场合，对方不能要求查看自己的手牌、场上覆盖的怪兽、额外卡组

		 中文名: 融合
		 卡片种类: 通常魔法
		 卡片密码: 24094653
		 罕见度: 平卡N，面闪SR
		 卡包: LB、DP01、BE01、VOL06、DL02、Starter Box、DPYG、DT07、DPKB、YU、JY、KA、PE、SY2、SD27
		 效果: 从手卡·自己场上把融合怪兽卡决定的融合素材怪兽送去墓地，那1只融合怪兽从额外卡组特殊召唤。

		*/
		Id:       178,
		Password: "24094653",
		Name:     "融合",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 23 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 201
		 调整:

		 [「守备」封禁]<『守備』封じ>
		 [2010/09/08]
		 ●选择对方场上的1只守备表示怪兽，变为表侧攻击表示。
		 ◇发动时选择对方场上存在的1只守备表示怪兽（取对象）
		 ◇效果处理时对象卡不在场上守备表示存在的场合，这个效果不适用
		 ◇「等级限制B地区/レベル制限Ｂ地区」的效果适用中，这张卡把怪兽变为攻击表示后，由于「等级限制B地区/レベル制限Ｂ地区」的效果立刻变为守备表示

		 中文名: 「守备」封禁
		 卡片种类: 通常魔法
		 卡片密码: 63102017
		 罕见度: 平卡N
		 卡包: PG、BE01、VOL03、DL02、Booster R2
		 效果: 选择1只对方场上的守备表示的怪兽变表侧攻击表示。

		*/
		Id:       201,
		Password: "63102017",
		Name:     "「守备」封禁",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellNormal(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()

				css := ygo.NewCards(tar.Mzone(), func(c *ygo.Card) bool {
					return c.IsDefense()
				})
				ca.PushSpell(ygo.LO_SetAttack, func() {
					if c := pl.SelectRequiredShor(ygo.LO_SetAttack, css); c != nil {
						c.SetFaceUpAttack()
					}
				})
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 24 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 205
		 调整:

		 [哥布林的秘药]<ゴブリンの秘薬>
		 [2010/09/08]
		 ●自己回复600基本分。
		 ◇「精灵之镜/精霊の鏡」可以对应这张卡的发动而发动

		 中文名: 哥布林的秘药
		 卡片种类: 通常魔法
		 卡片密码: 11868825
		 罕见度: 平卡N
		 卡包: PG、VOL02、DL02
		 效果: 自己回复600基本分。

		*/
		Id:       205,
		Password: "11868825",
		Name:     "哥布林的秘药",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellNormalPuah(ygo.LO_Reply, func() {
				pl := ca.GetSummoner()
				pl.ChangeLp(600)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 25 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 206
		 调整:

		 [火刑]<火あぶりの刑>
		 [2010/09/08]
		 ●给予对方基本分600分伤害。
		 ◇「精灵之镜/精霊の鏡」可以对应这张卡的发动而发动

		 中文名: 火刑
		 卡片种类: 通常魔法
		 卡片密码: 73134081
		 罕见度: 平卡N
		 卡包: PG、VOL02、DL02
		 效果: 给予对方基本分600分的伤害。

		*/
		Id:       206,
		Password: "73134081",
		Name:     "火刑",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellNormalPuah(ygo.LO_Hurt, func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				tar.ChangeLp(-600)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 26 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 207
		 调整:

		 [光之护封剑]<光の護封剣>
		 [2010/09/08]
		 ●对方场上存在的怪兽全部变为表侧表示。这张卡发动后，计算对方的回合3回合内在场上停留。只要这张卡在场上存在，对方场上存在的怪兽不能攻击宣言。
		 ◇对方场上不存在怪兽的场合这张卡也能发动
		 ◇发动被无效的场合，这张卡不能在场上停留
		 ◇对应这张卡的发动，发动「王宮の勅命/王宮の勅命」的场合，这张卡依然在场上正常停留（使用规则）
		 ◇承上述「王宮の勅命/王宮の勅命」在这张卡离开前效果失去的场合，对方怪兽不能攻击的效果适用
		 ◇「命运之火钟/運命の火時計」可以让这张卡的回合计算推进
		 ◇「命运之火钟/運命の火時計」在第3回合时让这张卡的回合计算推进的场合，在「命运之火钟/運命の火時計」适用时这张卡破坏
		 ◇「魔法防护器/マジック·ガードナー」的效果不能防止这张卡在3回合结束后的破坏（规则破坏不能被代替）
		 ◇效果处理时，对方场上不存在里侧表示怪兽的场合，怪兽变为表侧的效果不适用
		 ◇对应这张卡的发动，发动「旋风/サイクロン」并以这张卡为对象的场合，怪兽变为表侧的效果正常适用

		 中文名: 光之护封剑
		 卡片种类: 通常魔法
		 卡片密码: 72302403
		 罕见度: 平卡N，黄金GR，面闪SR
		 卡包: PG、BE01、VOL02、DL02、SD01、SD06、SD07、SD11、GLD01、GS01、DPYG、SD18、YU、YSD06、DB12、ST12、DS13、GDB1
		 效果: 对方场上存在的怪兽全部变成表侧表示。这张卡发动后，用对方回合计算的3回合内继续留在场上。只要这张卡在场上存在，对方场上存在的怪兽不能攻击宣言。

		*/
		Id:       207,
		Password: "72302403",
		Name:     "光之护封剑",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 27 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 218
		 调整:

		 [魔法除去]<魔法除去>
		 [2010/09/08]
		 ●场上的1张魔法卡破坏。选择的卡是覆盖的场合，翻开那张卡确认，是魔法卡就破坏。陷阱卡的场合变回原样。
		 ◇发动时选择这张卡以外的在自己或对方场上的魔法及陷阱区域存在的1张表侧表示的魔法卡或里侧表示覆盖着的1张卡（取对象）
		 ◇可以选择作为装备的怪兽卡
		 ◇可以选择由于「纳祭之魔/サクリファイス」的效果作为装备的「阿匹卜之化神/アポピスの化神」等为对象
		 ◇效果处理时，对象卡不在场上存在的场合，这个效果不适用
		 ◇作为怪兽装备的在魔法及陷阱区域存在的里侧怪兽卡在这个效果处理时破坏

		 中文名: 魔法除去
		 卡片种类: 通常魔法
		 卡片密码: 19159413
		 罕见度: 平卡N
		 卡包: PG、EX-R(EX)、VOL02、DL02、Booster R2、YU
		 效果: 破坏1张场上的魔法卡。选择的卡盖伏的场合，确认那张卡，是魔法卡就破坏，是陷阱卡就变回原来的盖伏形式。

		*/
		Id:       218,
		Password: "19159413",
		Name:     "魔法除去",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 28 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 219
		 调整:

		 [死者苏生]<死者蘇生>
		 [2010/09/08]
		 ●从自己或对方墓地选择1只怪兽发动。选择的怪兽在自己场上特殊召唤。
		 ◇发动时从自己或对方墓地选择1只可以被特殊召唤的怪兽（取对象）
		 ◇效果处理时对象怪兽不在墓地存在的场合，这个效果不适用
		 ◇效果处理时自己场上的怪兽区域没有空位的场合，这个效果不适用

		 中文名: 死者苏生
		 卡片种类: 通常魔法
		 卡片密码: 83764718
		 罕见度: 平卡N，银字R，黄金GR，面闪SR，金字UR
		 卡包: PG、BE01、EX-R(EX)、VOL02、DL02、GS01、DPYG、YU、KA、PE、SY2、SDM、SD25、GDB1
		 效果: 选择自己的或对方的墓地1只怪兽。选择的怪兽特殊召唤到自己场上。

		*/
		Id:       219,
		Password: "83764718",
		Name:     "死者苏生",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 29 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 220
		 调整:

		 [强欲之壶]<強欲な壺>
		 [2010/09/08]
		 ●从自己卡组抽2张卡。
		 ◇自己卡组中的卡不足2张时不能发动
		 ◇「精灵之镜/精霊の鏡」可以对应这张卡的发动而发动
		 ◇效果处理时，自己卡组中的卡不足2张的场合，抽卡时决斗败北

		 中文名: 强欲之壶
		 卡片种类: 通常魔法
		 卡片密码: 55144522
		 罕见度: 平卡N，银字R，面闪SR，立体UTR
		 卡包: PG、BE01、VOL03、DL02、SD01、SD02、SD03、SD04、DPKB、YU、JY、SY2
		 效果: 从自己的卡组抽2张卡。

		*/
		Id:       220,
		Password: "55144522",
		Name:     "强欲之壶",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 30 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 221
		 调整:

		 [掘墓的食尸鬼]<墓掘りグール>
		 [2010/09/08]
		 ●从对方墓地选择1张或2张的怪兽卡。被选择的卡从游戏中除外。
		 ◇发动时从对方墓地选择1张或2张的怪兽卡（取对象）
		 ◇效果处理时若其中1张对象不在墓地存在的场合，把剩余的对象卡从游戏中除外
		 ◇效果处理时若对象卡全部都不在墓地中存在，这个效果不适用

		 中文名: 掘墓的食尸鬼
		 卡片种类: 通常魔法
		 卡片密码: 82542267
		 罕见度: 银字R
		 卡包: PG、VOL03、DL02
		 效果: 选择对方的墓地1张到2张的怪兽卡。选择的卡从游戏中除外。

		*/
		Id:       221,
		Password: "82542267",
		Name:     "掘墓的食尸鬼",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 31 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2351
		 调整:

		 [尖刺神的杀虫剂]<トゲトゲ神の殺虫剤>
		 [2010/09/08]
		 ●场上的昆虫族怪兽全部破坏。
		 ◇不取对象
		 ◇效果处理时在场上表侧表示存在的昆虫族怪兽全部破坏

		 中文名: 尖刺神的杀虫剂
		 卡片种类: 通常魔法
		 卡片密码: 94716515
		 罕见度: 平卡N，银字R
		 卡包: VOL04、Booster04、GLD04
		 效果: 场上表侧表示存在的昆虫族怪兽全部破坏。

		*/
		Id:       2351,
		Password: "94716515",
		Name:     "尖刺神的杀虫剂",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 32 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2352
		 调整:

		 [永远的渴水]<永遠の渇水>
		 [2010/09/08]
		 ●场上的鱼族怪兽全部破坏。
		 ◇不取对象
		 ◇效果处理时在场上表侧表示存在的鱼族怪兽全部破坏

		 中文名: 永远的渴水
		 卡片种类: 通常魔法
		 卡片密码: 56606928
		 罕见度: 平卡N，银字R
		 卡包: VOL04、Booster04、GLD04
		 效果: 场上表侧表示存在的鱼族怪兽全部破坏。

		*/
		Id:       2352,
		Password: "56606928",
		Name:     "永远的渴水",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 33 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2353
		 调整:

		 [酸性风暴]<酸の嵐>
		 [2010/09/08]
		 ●场上的机械族怪兽全部破坏。
		 ◇不取对象
		 ◇效果处理时在场上表侧表示存在的机械族怪兽全部破坏

		 中文名: 酸性风暴
		 卡片种类: 通常魔法
		 卡片密码: 21323861
		 罕见度: 平卡N，银字R
		 卡包: VOL04、Booster04、TP01、DT03
		 效果: 场上表侧表示存在的机械族怪兽全部破坏。

		*/
		Id:       2353,
		Password: "21323861",
		Name:     "酸性风暴",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 34 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2354
		 调整:

		 [神之息吹]<神の息吹>
		 [2010/09/08]
		 ●场上的岩石族怪兽全部破坏。
		 ◇不取对象
		 ◇效果处理时在场上表侧表示存在的岩石族怪兽全部破坏

		 中文名: 神之息吹
		 卡片种类: 通常魔法
		 卡片密码: 20101223
		 罕见度: 平卡N，银字R
		 卡包: VOL04、Booster04、TP03
		 效果: 场上表侧表示存在的岩石族怪兽全部破坏。

		*/
		Id:       2354,
		Password: "20101223",
		Name:     "神之息吹",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 35 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2355
		 调整:

		 [战士抹杀]<戦士抹殺>
		 [2010/09/08]
		 ●场上的战士族怪兽全部破坏。
		 ◇不取对象
		 ◇效果处理时在场上表侧表示存在的战士族怪兽全部破坏

		 中文名: 战士抹杀
		 卡片种类: 通常魔法
		 卡片密码: 90873992
		 罕见度: 平卡N，银字R
		 卡包: VOL04、Booster04、TP02、GLD02
		 效果: 场上表侧表示存在的战士族怪兽全部破坏。

		*/
		Id:       2355,
		Password: "90873992",
		Name:     "战士抹杀",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 36 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 249
		 调整:

		 [万华镜-华丽的分身-]<万華鏡－華麗なる分身－>
		 [2010/09/08]
		 ●「鹰身女郎」表侧表示在场上有1只以上存在时可以发动。从手卡或卡组把1只「鹰身女郎」或「鹰身女郎三姐妹」特殊召唤。
		 ◇效果处理时「鹰身女郎/ハーピィ·レディ」不在场上表侧表示存在的场合，这个效果不适用。

		 中文名: 万华镜-华丽的分身-
		 卡片种类: 通常魔法
		 卡片密码: 90219263
		 罕见度: 平卡N
		 卡包: RB、BE01、VOL04、DL02、SD08
		 效果: 1只以上的「鹰身女郎」在场上表侧表示存在的时候发动。从手卡·卡组特殊召唤1只「鹰身女郎」或「鹰身女郎三姐妹」。

		*/
		Id:       249,
		Password: "90219263",
		Name:     "万华镜-华丽的分身-",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 37 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 269
		 调整:

		 [对死者的供奉]<死者への手向け>
		 [2010/09/08]
		 ●丢弃1张手卡，选择场上存在的1只怪兽发动。被选择的怪兽破坏。
		 ◇发动时丢弃1张手卡（代价）
		 ◇发动时选择自己或对方场上存在的1只怪兽（取对象）
		 ◇效果处理时对象怪兽不在场上存在的场合，这个效果不适用

		 中文名: 对死者的供奉
		 卡片种类: 通常魔法
		 卡片密码: 79759861
		 罕见度: 平卡N，银字R
		 卡包: RB、BE01、YSD01、VOL05、DL02、SD03、YSD04、KA、SY2
		 效果: 丢弃1张手卡，选择场上存在的1只怪兽发动。选择的怪兽破坏。

		*/
		Id:       269,
		Password: "79759861",
		Name:     "对死者的供奉",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 38 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 270
		 调整:

		 [魂之解放]<魂の解放>
		 [2012/01/29]
		 ●选择双方墓地的卡合计最多5张，那些卡从游戏中除外。
		 ◇取对象效果。
		 ◇效果发动时选择双方墓地合计1~5张卡作为效果的对象。
		 ◇双方墓地任意一方有1张以上的卡存在的场合就可以将这张卡发动。

		 中文名: 魂之解放
		 卡片种类: 通常魔法
		 卡片密码: 05758500
		 罕见度: 平卡N，银字R，黄金GR
		 卡包: RB、BE01、VOL05、DL02、SD14、GS04、GDB1
		 效果: 选择双方墓地的卡合计最多5张，那些卡从游戏中除外。

		*/
		Id:       270,
		Password: "05758500",
		Name:     "魂之解放",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 39 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 271
		 调整:

		 [开朗的殡葬师]<陽気な葬儀屋>
		 [2010/09/08]
		 ●从自己的手卡把最多3张怪兽卡丢弃去墓地。
		 ◇效果处理时选择要丢弃的卡（不取对象）
		 ◇发动的场合，必须至少选择1张卡

		 中文名: 开朗的殡葬师
		 卡片种类: 通常魔法
		 卡片密码: 41142615
		 罕见度: 平卡N，银字R
		 卡包: RB、BE01、VOL05、DL02
		 效果: 从自己的手卡中丢弃最多3张怪兽卡送去墓地。

		*/
		Id:       271,
		Password: "41142615",
		Name:     "开朗的殡葬师",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 40 */
	cardBag.Register(&ygo.CardOriginal{
		/*
			 id: 272
			 调整:

			 [心变]<心変わり>
			 [2010/09/08]
			 ●选择对方场上的1只怪兽。直到发动回合的结束阶段为止，得到被选择的卡的控制权。
			 ◇发动时选择对方场上的1只怪兽（取对象）
			 ◇效果处理时对象怪兽不在场上存在的场合，这个效果不适用
			 ◇对方在我方回合发动「血之代偿/血の代償」「活死人的呼声/リビングデッドの呼び声」等召唤、特殊召唤怪兽的场合，得到那些怪兽控制权后可以改变那些怪兽的表示形式
			 ◇这个效果得到控制权的怪兽，变为里侧表示的场合，结束阶段时也归还控制权
			 ◇「天使的手镜/天使の手鏡」对应这张卡发动的场合，也只能选择「心变/心変わり」控制者的对方的场上怪兽
			 ◇用这个效果得到控制权的对方怪兽，再用其他卡的效果把那只怪兽的控制权转移给对方，那个回合的结束阶段时那只怪兽不在需要归还控制权
			 ◇用这个效果得到控制权的对方怪兽，被「亚空间物质传送装置/亜空間物質転送装置」暂时从游戏中除外，结束阶段时也要归还控制权
			 ◇用这个效果得到控制权的对方怪兽，被「虫洞/ワーム·ホール」暂时从游戏中除外，在下次的准备阶段时那只怪兽回到场上后立刻回到对方场上★对应这张卡的发动，对方发动「扰乱三人组/おジャマトリオ」让自己场上没有空余的怪兽区域时如何处理
			调整中
			 ◇结束阶段归还控制权时，若对方场上不存在空余的怪兽区域，那只怪兽送到持有者的墓地

			 中文名: 心变
			 卡片种类: 通常魔法
			 卡片密码: 04031928
			 罕见度: 面闪SR，金字UR
			 卡包: RB、BE01、EX-R(EX)、VOL05、DL02、YU、PE
			 效果: 选择对方的场上的1只怪兽。直到发动回合的结束阶段得到选择的卡的控制权。

		*/
		Id:       272,
		Password: "04031928",
		Name:     "心变",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 41 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 432
		 调整:

		 [火炎地狱]<火炎地獄>
		 [2010/09/08]
		 ●给予对方基本分1000分的伤害，自己受到500分的伤害。
		 ◇同时发生伤害
		 ◇对应这张卡的发动，对方发动「地狱的冷枪/地獄の扉越し銃」「痛魂之咒术/痛魂の呪術」的场合，自己受到合计1500分的伤害
		 ◇对应这张卡的发动，对方发动「黑板擦的陷阱/黒板消しの罠」的场合，对方受到的1000分伤害无效

		 中文名: 火炎地狱
		 卡片种类: 通常魔法
		 卡片密码: 46918794
		 罕见度: 平卡N
		 卡包: ME、BE02、VOL07、DL04、Booster07、YSD03
		 效果: 给予对方基本分1000分伤害，自己受到500分伤害。

		*/
		Id:       432,
		Password: "46918794",
		Name:     "火炎地狱",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 42 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 467
		 调整:

		 [右手持盾左手持剑]<右手に盾を左手に剣を>
		 [2010/09/08]
		 ●直到结束阶段结束时为止，这张卡的发动时在场上的全部表侧表示怪兽的原本的攻击力与守备力交换。
		 ◇交换的数值是原本数值
		 ◇效果处理时在场上表侧表示存在的怪兽全部适用（不是叙述上的“发动时”）

		 中文名: 右手持盾左手持剑
		 卡片种类: 通常魔法
		 卡片密码: 52097679
		 罕见度: 平卡N，银字R
		 卡包: ME、BE02、VOL07、DL04、SD07、JY、SJ2
		 效果: 这个回合结束前，这张卡发动时全部场上存在的表侧表示的怪兽原来的攻击力和原来的守备力交替。

		*/
		Id:       467,
		Password: "52097679",
		Name:     "右手持盾左手持剑",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 43 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 469
		 调整:

		 [「攻击」封禁]<『攻撃』封じ>
		 [2010/09/08]
		 ●指定对方场上的1只表侧攻击表示怪兽变为表侧守备表示。
		 ◇发动时选择对方场上表侧攻击表示存在的1只怪兽（取对象）
		 ◇效果处理时那只怪兽已经是表侧守备表示的场合，这个效果不适用
		 ◇效果处理时那只怪兽不在场上表侧表示存在的场合，这个效果不适用
		 ◇「最终突击命令/最終突撃命令」效果适用中，这张卡把怪兽变为守备表示后立刻变为攻击表示

		 中文名: 「攻击」封禁
		 卡片种类: 通常魔法
		 卡片密码: 25880422
		 罕见度: 平卡N
		 卡包: ME、BE02、VOL07、DL04
		 效果: 指定的对方场上的1只表侧表示的怪兽转为守备表示。

		*/
		Id:       469,
		Password: "25880422",
		Name:     "「攻击」封禁",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 44 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 177
		 调整:

		 [落穴]<落とし穴>
		 [2010/09/08]
		 ●对方把攻击力1000以上的怪兽召唤·反转召唤时可以发动。那1只攻击力1000以上的怪兽破坏。
		 ◇发动时以那只被召唤·反转召唤的怪兽为对象（取对象）
		 ◇由于是取对象效果，所以例如「流氓佣兵部队/ならず者傭兵部隊」等在召唤成功时的这张卡发动前从场上离开或变为里侧表示的场合，符合条件的对象不存在，这张卡不能发动
		 ◇效果处理时对象怪兽不在场上表侧表示存在的场合，这个效果不适用
		 ◇效果处理时对象怪兽的攻击力比1000低的场合，这个效果不适用

		 中文名: 落穴
		 卡片种类: 通常陷阱
		 卡片密码: 04206964
		 罕见度: 平卡N，银字R，黄金GR
		 卡包: LB、BE01、EX-R(EX)、YSD02、VOL01、DL02、YSD03、DT04、YSD05、JY、SY2、YSD06、GS03、ST12、GDB1
		 效果: 对方对攻击力1000以上的怪兽的召唤·反转召唤成功时才能发动。那1只攻击力1000以上的怪兽破坏。

		*/
		Id:       177,
		Password: "04206964",
		Name:     "落穴",
		Lc:       ygo.LC_TrapNormal, // 通常陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 45 */
	cardBag.Register(&ygo.CardOriginal{
		/*
			 id: 2378
			 调整:

			 [避雷针]<避雷針>
			 [2010/09/08]
			 ●对方使用「雷击」时，把对方怪兽全部破坏来代替自己怪兽。发动后这张卡破坏。
			 ◇对方场上没有怪兽的场合，这张卡也能发动★这个效果是否为“让对方的「雷击/サンダー·ボルト」无效，把对方场上的怪兽全部破坏的效果”
			调整中★这个效果是否是把「雷击/サンダー·ボルト」的效果变为“破坏自己场上的全部怪兽”
			调整中
			 ◇「星光大道/スターライト·ロード」是否能对应这张卡的发动而发动
			调整中

			 中文名: 避雷针
			 卡片种类: 通常陷阱
			 卡片密码: 42364257
			 罕见度: 银字R
			 卡包: VOL05、Booster05
			 效果: 对方使用了「雷击」的时候，破坏对方全部怪兽代替自己的怪兽。发动后这张卡破坏。

		*/
		Id:       2378,
		Password: "42364257",
		Name:     "避雷针",
		Lc:       ygo.LC_TrapNormal, // 通常陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 46 */
	cardBag.Register(&ygo.CardOriginal{
		/*
			 id: 268
			 调整:

			 [伪陷阱]<伪物のわな>
			 [10/09/13]
			 ●对方使用破坏自己场上的陷阱卡的魔法·陷阱·效果怪兽的效果时可以发动。这张卡代替破坏，其他的陷阱卡不破坏。覆盖的卡要被破坏的场合，那些卡全部翻开确认。
			 ◇必须在自己场上有其他陷阱卡存在时才能发动
			 ◇发动时给对方确认覆盖的卡
			 ◇连锁这张卡的发动，发动「沙尘之大龙卷/砂尘の大竜巻」在场上覆盖陷阱卡的场合，之后不能确认那张覆盖的陷阱卡

			 ◇不能对应「元素英雄
			天空侠/E·HERO
			エアーマン」等不确定的破坏效果发动这张卡
			 ◇「王宫的通告/王宫のお触れ」在场上存在时，发动这张卡的场合也确认覆盖卡
			 ◇这张卡在自己场上有其他陷阱卡存在时可以对应「暗之卡组破坏病毒/闇のデッキ破壊ウイルス」发动，那个场合自己场上的陷阱卡不破坏，手卡中的陷阱卡破坏
			 ◇这张卡对应「黑蔷薇龙/ブラック·ローズ·ドラゴン」的破坏效果发动的场合，自己场上的其他陷阱卡不破坏
			 ◇「旋风/サイクロン」以1张覆盖的陷阱卡为对象发动的场合这张卡对应发动，在发动时只确认那张对象卡
			 ◇对应这张卡的发动而发动「旋风/サイクロン」并以这张卡为对象的场合，效果处理时不能代替其他陷阱卡破坏
			 ◇可以对应「命运英雄
			钻石人/D－HERO
			ダイヤモンドガイ」墓地发动的「大风暴/大岚」等发动
			 ◇这张卡可以对应「星尘龙/スターダスト·ドラゴン」的效果发动而发动
			 ◇「星尘龙/スターダスト·ドラゴン」不能对应这张卡的发动而发动★是否能对应魔法卡效果的发动而发动
			调整中

			 中文名: 伪陷阱
			 卡片种类: 通常陷阱
			 卡片密码: 03027001
			 罕见度: 平卡N
			 卡包: RB、BE01、Booster R3、VOL05、DL02、TP02、JY
			 效果: 把自己场上的陷阱卡破坏的魔法·陷阱·效果怪兽的效果对方发动时才能发动，把这张卡作为代替破坏，其他的自己的陷阱卡不破坏。盖放的卡破坏的场合，那些卡全部翻开确认。

		*/
		Id:       268,
		Password: "03027001",
		Name:     "伪陷阱",
		Lc:       ygo.LC_TrapNormal, // 通常陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 47 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 474
		 调整:

		 [神圣防护罩-反射镜力-]<聖なるバリア－ミラーフォース－>
		 [2010/09/08]
		 ●对方怪兽的攻击宣言时可以发动。对方场上存在的攻击表示怪兽全部破坏。
		 ◇不取对象
		 ◇对方场上不存在攻击表示怪兽的场合，这张卡不能发动（「绝对防御将军/絶対防御将軍」）
		 ◇里侧攻击表示的怪兽也破坏
		 ◇效果处理时，对方场上不存在攻击表示怪兽的场合，这个效果不适用。

		 中文名: 神圣防护罩-反射镜力-
		 卡片种类: 通常陷阱
		 卡片密码: 44095762
		 罕见度: 平卡N，黄金GR，面闪SR，金字UR，爆闪PR，银碎SER
		 卡包: ME、BE02、VOL07、DL04、GLD01、GS01、DPYG、SD19、YU、SY2、SDM、YSD06、DB12、DS13、GDB1
		 效果: 对方怪兽的攻击宣言时才能发动。对方场上存在的攻击表示怪兽全部破坏。

		*/
		Id:       474,
		Password: "44095762",
		Name:     "神圣防护罩-反射镜力-",
		Lc:       ygo.LC_TrapNormal, // 通常陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 48 */
	cardBag.Register(&ygo.CardOriginal{
		/*
			 id: 2377
			 调整:

			 [来自黑暗的呼声]<闇からの呼び声>
			 [2010/09/08]
			 ●用「死者苏生」的效果特殊召唤的怪兽全部送去墓地。只要这张卡在场上存在，双方不能使用「死者苏生/死者蘇生」。
			 ◇这张卡对应「死者苏生/死者蘇生」的发动而发动的场合，在「死者苏生/死者蘇生」效果处理时，「死者苏生/死者蘇生」的效果不适用
			 ◇这张卡的永续效果适用时，之前由于「死者苏生/死者蘇生」的效果特殊召唤的怪兽就送去墓地
			 ◇这张卡效果适用中一度被无效化，期间用「死者苏生/死者蘇生」特殊召唤了怪兽，这张卡效果再次适用时那些怪兽送去墓地★用「死者苏生/死者蘇生」的效果特殊召唤的怪兽一度不在场上表侧表示存在后再次出现的场合，是否会被送去墓地
			调整中

			 中文名: 来自黑暗的呼声
			 卡片种类: 永续陷阱
			 卡片密码: 78637313
			 罕见度: 平卡N
			 卡包: VOL05、Booster05、Booster Chronicle、Booster R2、TP09
			 效果: 用「死者苏生」苏生的怪兽全部送到墓地。只要场上有这张卡「死者苏生」就不能使用。

		*/
		Id:       2377,
		Password: "78637313",
		Name:     "来自黑暗的呼声",
		Lc:       ygo.LC_TrapContinuous, // 永续陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 49 */
	cardBag.Register(&ygo.CardOriginal{
		/*
			 id: 471
			 调整:

			 [哥布林拦路贼]<追い剥ぎゴブリン>
			 [2010/09/08]
			 ●自己场上的怪兽每次给予对方玩家战斗伤害，对方随机丢弃1张手卡。
			 ◇效果发动进入连锁（1速）
			 ◇不能在伤害步骤发动这张卡
			 ◇效果发动时点同「天空骑士
			柏修斯/天空騎士パーシアス」（具体请参考关于伤害步骤的详解）
			 ◇「次元壁/ディメンション·ウォール」将战斗伤害转移给对方的场合，这个效果不发动

			 中文名: 哥布林拦路贼
			 卡片种类: 永续陷阱
			 卡片密码: 88279736
			 罕见度: 平卡N，银字R
			 卡包: ME、BE02、VOL07、DL04、SD07、GLD04
			 效果: 自己场上的怪兽每次造成对方玩家的战斗伤害时，对方随机丢弃1张手卡。

		*/
		Id:       471,
		Password: "88279736",
		Name:     "哥布林拦路贼",
		Lc:       ygo.LC_TrapContinuous, // 永续陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 50 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 463
		 调整:

		 [神之宣告]<神の宣告>
		 [2010/09/08]
		 ●支付一半基本分发动。魔法·陷阱卡的发动，怪兽的召唤·反转召唤·特殊召唤的其中之1无效并破坏。
		 ◇发动时支付一半基本分（代价）
		 ◇不能对应「喧闹的邪恶灵/ポルターガイスト」「电脑增幅器/電脳増幅器」的发动而发动
		 ◇不能对应由于卡的效果特殊召唤的怪兽而发动（「死皇帝之陵墓/死皇帝の陵墓」「紧急同调/緊急同調」「血之代偿/血の代償」除外）
		 ◇基本分奇数的状态，发动这张卡的场合基本分以四舍五入计算
		 ◇不能对应魔法或陷阱卡效果的发动而发动

		 中文名: 神之宣告
		 卡片种类: 反击陷阱
		 卡片密码: 41420027
		 罕见度: 平卡N，黄金GR，面闪SR
		 卡包: ME、BE02、VOL06、DL04、SD11、SD14、GLD02、GS02、SD20、GDB1
		 效果: 把基本分支付一半发动。魔法·陷阱卡的发动，怪兽的召唤·反转召唤·特殊召唤的其中1个无效并破坏。

		*/
		Id:       463,
		Password: "41420027",
		Name:     "神之宣告",
		Lc:       ygo.LC_TrapCounter, // 反击陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 51 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 464
		 调整:

		 [魔法干扰阵]<マジック·ジャマー>
		 [2010/09/08]
		 ●丢弃1张手卡发动。魔法卡的发动无效并破坏。
		 ◇发动时丢弃1张手牌（代价）
		 ◇不能对应魔法卡效果的发动而发动

		 中文名: 魔法干扰阵
		 卡片种类: 反击陷阱
		 卡片密码: 77414722
		 罕见度: 平卡N，面闪SR
		 卡包: ME、BE02、YSD02、VOL06、DL04、SD02、SD05、SD08、SD11、SD13、DT01、DTP01、YSD05、YU、SY2、SK2
		 效果: 丢弃1张手卡发动。魔法卡的发动无效并破坏。

		*/
		Id:       464,
		Password: "77414722",
		Name:     "魔法干扰阵",
		Lc:       ygo.LC_TrapCounter, // 反击陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 52 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 465
		 调整:

		 [盗贼的七道具]<盗賊の七つ道具>
		 [2010/09/08]
		 ●支付1000基本分。陷阱卡的发动无效并破坏。
		 ◇发动时支付1000基本分（代价）
		 ◇不能对应陷阱卡效果的发动而发动

		 中文名: 盗贼的七道具
		 卡片种类: 反击陷阱
		 卡片密码: 03819470
		 罕见度: 平卡N，面闪SR
		 卡包: ME、BE02、VOL06、DL04、SD11、YSD06、TU05、ST12
		 效果: 支付1000基本分发动。陷阱卡的发动无效并破坏。

		*/
		Id:       465,
		Password: "03819470",
		Name:     "盗贼的七道具",
		Lc:       ygo.LC_TrapCounter, // 反击陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 53 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 466
		 调整:

		 [升天之角笛]<昇天の角笛>
		 [2010/09/08]
		 ●自己场上的1只怪兽作为祭品。怪兽的召唤·反转召唤·特殊召唤无效并破坏。
		 ◇发动时把自己场上的1只怪兽作为祭品（代价）
		 ◇不能对应由于卡的效果特殊召唤的怪兽而发动（「死皇帝之陵墓/死皇帝の陵墓」「紧急同调/緊急同調」「血之代偿/血の代償」除外）

		 中文名: 升天之角笛
		 卡片种类: 反击陷阱
		 卡片密码: 98069388
		 罕见度: 银字R，面闪SR
		 卡包: ME、BE02、VOL06、DL04
		 效果: 自己场上的1只怪兽作为祭品。怪兽的召唤·反转召唤·特殊召唤无效并破坏。

		*/
		Id:       466,
		Password: "98069388",
		Name:     "升天之角笛",
		Lc:       ygo.LC_TrapCounter, // 反击陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 54 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 184
		 调整:

		 中文名: 龙骑士 盖亚
		 卡片种类: 融合怪兽
		 卡片密码: 66889139
		 种族: 龙
		 属性: 风
		 星级: 7
		 攻击力: 2600
		 防御力: 2100
		 罕见度: 银字R，面闪SR，立体UTR
		 卡包: PG、BE01、VOL03、DL02、309、Booster R2
		 效果: 融合：「暗黑骑士 盖亚」＋「诅咒之龙」。

		*/
		Id:       184,
		Password: "66889139",
		Name:     "龙骑士 盖亚",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 7,
		La:    ygo.LA_Wind,   // 风
		Lr:    ygo.LR_Dragon, // 龙
		Atk:   2600,
		Def:   2100,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 55 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 208
		 调整:

		 中文名: 金属龙
		 卡片种类: 融合怪兽
		 卡片密码: 09293977
		 种族: 机械
		 属性: 风
		 星级: 6
		 攻击力: 1850
		 防御力: 1700
		 罕见度: 平卡N
		 卡包: PG、VOL03、DL02
		 效果: 融合：「钢铁巨神像」＋「下级龙」。

		*/
		Id:       208,
		Password: "09293977",
		Name:     "金属龙",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 6,
		La:    ygo.LA_Wind,    // 风
		Lr:    ygo.LR_Machine, // 机械
		Atk:   1850,
		Def:   1700,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 56 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2282
		 调整:

		 中文名: 炭烧战士
		 卡片种类: 融合怪兽
		 卡片密码: 54541900
		 种族: 战士
		 属性: 地
		 星级: 4
		 攻击力: 1500
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: PG、VOL02、TP02
		 效果: 融合：「磁力战士1号」＋「磁力战士2号」。

		*/
		Id:       2282,
		Password: "54541900",
		Name:     "炭烧战士",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 4,
		La:    ygo.LA_Earth,   // 地
		Lr:    ygo.LR_Warrior, // 战士
		Atk:   1500,
		Def:   1200,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 57 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2285
		 调整:

		 中文名: 混沌男巫
		 卡片种类: 融合怪兽
		 卡片密码: 41544074
		 种族: 魔法师
		 属性: 暗
		 星级: 4
		 攻击力: 1300
		 防御力: 1100
		 罕见度: 平卡N
		 卡包: VOL02、TP03
		 效果: 融合：「圣精灵」＋「黑魔族的窗帘」。

		*/
		Id:       2285,
		Password: "41544074",
		Name:     "混沌男巫",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 4,
		La:    ygo.LA_Dark,        // 暗
		Lr:    ygo.LR_Spellcaster, // 魔法师
		Atk:   1300,
		Def:   1100,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 58 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2287
		 调整:

		 中文名: 奇迹鸟
		 卡片种类: 融合怪兽
		 卡片密码: 59036972
		 种族: 鸟兽
		 属性: 风
		 星级: 4
		 攻击力: 1300
		 防御力: 900
		 罕见度: 平卡N
		 卡包: VOL02、TP17
		 效果: 融合：「大炮鸟」＋「邪炎之翼」。

		*/
		Id:       2287,
		Password: "59036972",
		Name:     "奇迹鸟",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 4,
		La:    ygo.LA_Wind,        // 风
		Lr:    ygo.LR_WingedBeast, // 鸟兽
		Atk:   1300,
		Def:   900,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 59 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2290
		 调整:

		 中文名: 不死战士
		 卡片种类: 融合怪兽
		 卡片密码: 31339260
		 种族: 不死
		 属性: 暗
		 星级: 3
		 攻击力: 1200
		 防御力: 900
		 罕见度: 平卡N
		 卡包: VOL02、TP16
		 效果: 融合：「白骨」＋「格斗战士 阿提米特」。

		*/
		Id:       2290,
		Password: "31339260",
		Name:     "不死战士",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 3,
		La:    ygo.LA_Dark,   // 暗
		Lr:    ygo.LR_Zombie, // 不死
		Atk:   1200,
		Def:   900,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 60 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2291
		 调整:

		 中文名: 魔装骑士 德拉格尼斯
		 卡片种类: 融合怪兽
		 卡片密码: 70681994
		 种族: 战士
		 属性: 风
		 星级: 3
		 攻击力: 1200
		 防御力: 900
		 罕见度: 平卡N
		 卡包: PG、VOL02、TP09
		 效果: 融合：「铠甲剑尾战士」＋「独眼盾龙」。

		*/
		Id:       2291,
		Password: "70681994",
		Name:     "魔装骑士 德拉格尼斯",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 3,
		La:    ygo.LA_Wind,    // 风
		Lr:    ygo.LR_Warrior, // 战士
		Atk:   1200,
		Def:   900,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 61 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2325
		 调整:

		 中文名: 巴洛克斯
		 卡片种类: 融合怪兽
		 卡片密码: 06840573
		 种族: 恶魔
		 属性: 暗
		 星级: 5
		 攻击力: 1380
		 防御力: 1530
		 罕见度: 平卡N
		 卡包: VOL03
		 效果: 融合：「杀人熊猫」＋「石像怪」。

		*/
		Id:       2325,
		Password: "06840573",
		Name:     "巴洛克斯",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 5,
		La:    ygo.LA_Dark,  // 暗
		Lr:    ygo.LR_Fiend, // 恶魔
		Atk:   1380,
		Def:   1530,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 62 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2326
		 调整:

		 中文名: 花狼
		 卡片种类: 融合怪兽
		 卡片密码: 95952802
		 种族: 兽
		 属性: 地
		 星级: 5
		 攻击力: 1800
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: PG、VOL03
		 效果: 融合：「银牙狼」＋「魔界之棘」。

		*/
		Id:       2326,
		Password: "95952802",
		Name:     "花狼",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 5,
		La:    ygo.LA_Earth, // 地
		Lr:    ygo.LR_Beast, // 兽
		Atk:   1800,
		Def:   1400,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 63 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2327
		 调整:

		 中文名: 珍鱼
		 卡片种类: 融合怪兽
		 卡片密码: 80516007
		 种族: 鱼
		 属性: 水
		 星级: 4
		 攻击力: 1500
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: VOL03、TP15
		 效果: 融合：「融合体」＋「恍惚的人鱼」。

		*/
		Id:       2327,
		Password: "80516007",
		Name:     "珍鱼",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 4,
		La:    ygo.LA_Water, // 水
		Lr:    ygo.LR_Fish,  // 鱼
		Atk:   1500,
		Def:   1200,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 64 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2328
		 调整:

		 中文名: 潜伏深海的鲨鱼
		 卡片种类: 融合怪兽
		 卡片密码: 28593363
		 种族: 鱼
		 属性: 水
		 星级: 5
		 攻击力: 1900
		 防御力: 1600
		 罕见度: 平卡N
		 卡包: RB、VOL04
		 效果: 融合：「神鱼」＋「舌鱼」。

		*/
		Id:       2328,
		Password: "28593363",
		Name:     "潜伏深海的鲨鱼",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 5,
		La:    ygo.LA_Water, // 水
		Lr:    ygo.LR_Fish,  // 鱼
		Atk:   1900,
		Def:   1600,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 65 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2329
		 调整:

		 中文名: 尖刺龙
		 卡片种类: 融合怪兽
		 卡片密码: 33691040
		 种族: 恐龙
		 属性: 地
		 星级: 5
		 攻击力: 1900
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: VOL04
		 效果: 融合：「虎纹龙」＋「火焰毒蛇」。

		*/
		Id:       2329,
		Password: "33691040",
		Name:     "尖刺龙",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 5,
		La:    ygo.LA_Earth,    // 地
		Lr:    ygo.LR_Dinosaur, // 恐龙
		Atk:   1900,
		Def:   1500,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 66 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2357
		 调整:

		 中文名: 雷神之怒
		 卡片种类: 融合怪兽
		 卡片密码: 09653271
		 种族: 雷
		 属性: 风
		 星级: 5
		 攻击力: 1900
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: RB、VOL04、TP12
		 效果: 融合：「耳天使」＋「大雷电球」。

		*/
		Id:       2357,
		Password: "09653271",
		Name:     "雷神之怒",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 5,
		La:    ygo.LA_Wind,    // 风
		Lr:    ygo.LR_Thunder, // 雷
		Atk:   1900,
		Def:   1400,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 67 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2379
		 调整:

		 中文名: 下达裁决的女帝
		 卡片种类: 融合怪兽
		 卡片密码: 15237615
		 种族: 战士
		 属性: 地
		 星级: 6
		 攻击力: 2100
		 防御力: 1700
		 罕见度: 平卡N
		 卡包: RB、VOL05
		 效果: 融合：「女王的影武者」＋「响女」。

		*/
		Id:       2379,
		Password: "15237615",
		Name:     "下达裁决的女帝",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 6,
		La:    ygo.LA_Earth,   // 地
		Lr:    ygo.LR_Warrior, // 战士
		Atk:   2100,
		Def:   1700,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 68 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2380
		 调整:

		 中文名: 栖身蔷薇的恶灵
		 卡片种类: 融合怪兽
		 卡片密码: 32485271
		 种族: 植物
		 属性: 暗
		 星级: 6
		 攻击力: 2000
		 防御力: 1800
		 罕见度: 平卡N
		 卡包: VOL05
		 效果: 融合：「小精怪」＋「蛇椰树」。

		*/
		Id:       2380,
		Password: "32485271",
		Name:     "栖身蔷薇的恶灵",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 6,
		La:    ygo.LA_Dark,  // 暗
		Lr:    ygo.LR_Plant, // 植物
		Atk:   2000,
		Def:   1800,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 69 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2404
		 调整:

		 中文名: 裁决之鹰
		 卡片种类: 融合怪兽
		 卡片密码: 74703140
		 种族: 鸟兽
		 属性: 风
		 星级: 6
		 攻击力: 2100
		 防御力: 1800
		 罕见度: 平卡N
		 卡包: ME、VOL06
		 效果: 融合：「苍翼冠鸟」＋「雏鸡」。

		*/
		Id:       2404,
		Password: "74703140",
		Name:     "裁决之鹰",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 6,
		La:    ygo.LA_Wind,        // 风
		Lr:    ygo.LR_WingedBeast, // 鸟兽
		Atk:   2100,
		Def:   1800,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 70 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2405
		 调整:

		 中文名: 魔导骑士 基尔提亚
		 卡片种类: 融合怪兽
		 卡片密码: 51828629
		 种族: 战士
		 属性: 光
		 星级: 5
		 攻击力: 1850
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: ME、VOL06、TP05
		 效果: 融合：「冥界的番人」＋「王座守护者」。

		*/
		Id:       2405,
		Password: "51828629",
		Name:     "魔导骑士 基尔提亚",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 5,
		La:    ygo.LA_Light,   // 光
		Lr:    ygo.LR_Warrior, // 战士
		Atk:   1850,
		Def:   1500,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 71 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2428
		 调整:

		 中文名: 迷宫的魔战车
		 卡片种类: 融合怪兽
		 卡片密码: 99551425
		 种族: 机械
		 属性: 暗
		 星级: 7
		 攻击力: 2400
		 防御力: 2400
		 罕见度: 平卡N
		 卡包: ME、VOL07、TP11
		 效果: 融合：「高科技狼」＋「加农炮兵」。

		*/
		Id:       2428,
		Password: "99551425",
		Name:     "迷宫的魔战车",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 7,
		La:    ygo.LA_Dark,    // 暗
		Lr:    ygo.LR_Machine, // 机械
		Atk:   2400,
		Def:   2400,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 72 */
	cardBag.Register(&ygo.CardOriginal{
		/*
			 id: 243
			 调整:
			（这张卡的日文原卡名带有「恶魔/デーモン」）

			 中文名: 暗黑魔龙
			 卡片种类: 融合怪兽
			 卡片密码: 11901678
			 种族: 龙
			 属性: 暗
			 星级: 9
			 攻击力: 3200
			 防御力: 2500
			 罕见度: 面闪SR，金字UR，爆闪PR，立体UTR
			 卡包: RB、BE01、MA、VOL05、DL02
			 效果: 融合：「恶魔召唤」＋「真红眼黑龙」。

		*/
		Id:       243,
		Password: "11901678",
		Name:     "暗黑魔龙",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 9,
		La:    ygo.LA_Dark,   // 暗
		Lr:    ygo.LR_Dragon, // 龙
		Atk:   3200,
		Def:   2500,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 73 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 245
		 调整:

		 中文名: 轰鸣的大海蛇
		 卡片种类: 融合怪兽
		 卡片密码: 19066538
		 种族: 水
		 属性: 水
		 星级: 6
		 攻击力: 2100
		 防御力: 1800
		 罕见度: 平卡N
		 卡包: RB、VOL05、DL02
		 效果: 融合：「魔法灯」＋「兵主部」。

		*/
		Id:       245,
		Password: "19066538",
		Name:     "轰鸣的大海蛇",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 6,
		La:    ygo.LA_Water, // 水
		Lr:    ygo.LR_Water, // 水
		Atk:   2100,
		Def:   1800,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 74 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 419
		 调整:

		 中文名: 千年龙
		 卡片种类: 融合怪兽
		 卡片密码: 41462083
		 种族: 龙
		 属性: 风
		 星级: 7
		 攻击力: 2400
		 防御力: 2000
		 罕见度: 平卡N，金字UR
		 卡包: ME、BE02、LE02、VOL06、DL04、JY
		 效果: 融合：「时间魔术师」＋「宝贝龙」。

		*/
		Id:       419,
		Password: "41462083",
		Name:     "千年龙",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 7,
		La:    ygo.LA_Wind,   // 风
		Lr:    ygo.LR_Dragon, // 龙
		Atk:   2400,
		Def:   2000,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 75 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 425
		 调整:

		 中文名: 牛头人马兽
		 卡片种类: 融合怪兽
		 卡片密码: 94905343
		 种族: 兽战士
		 属性: 地
		 星级: 6
		 攻击力: 2000
		 防御力: 1700
		 罕见度: 平卡N
		 卡包: ME、BE02、VOL07、DL04
		 效果: 融合：「牛头人」＋「人马兽」。

		*/
		Id:       425,
		Password: "94905343",
		Name:     "牛头人马兽",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 6,
		La:    ygo.LA_Earth,        // 地
		Lr:    ygo.LR_BeastWarrior, // 兽战士
		Atk:   2000,
		Def:   1700,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 76 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 443
		 调整:

		 中文名: 音乐家帝王
		 卡片种类: 融合怪兽
		 卡片密码: 56907389
		 种族: 魔法师
		 属性: 光
		 星级: 5
		 攻击力: 1750
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: ME、VOL06、DL04、TP10
		 效果: 融合：「黑森林的魔女」＋「高等女祭司」。

		*/
		Id:       443,
		Password: "56907389",
		Name:     "音乐家帝王",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 5,
		La:    ygo.LA_Light,       // 光
		Lr:    ygo.LR_Spellcaster, // 魔法师
		Atk:   1750,
		Def:   1500,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 77 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 445
		 调整:

		 中文名: 机械恐龙
		 卡片种类: 融合怪兽
		 卡片密码: 89112729
		 种族: 机械
		 属性: 地
		 星级: 5
		 攻击力: 1800
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: ME、VOL07、DL04、TP19
		 效果: 融合：「炸弹先生」＋「双头恐龙王」。

		*/
		Id:       445,
		Password: "89112729",
		Name:     "机械恐龙",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 5,
		La:    ygo.LA_Earth,   // 地
		Lr:    ygo.LR_Machine, // 机械
		Atk:   1800,
		Def:   1400,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 78 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 458
		 调整:

		 中文名: 双头雷龙
		 卡片种类: 融合怪兽
		 卡片密码: 54752875
		 种族: 雷
		 属性: 光
		 星级: 7
		 攻击力: 2800
		 防御力: 2100
		 罕见度: 平卡N
		 卡包: ME、BE02、VOL07、DL04
		 效果: 融合：「雷龙」＋「雷龙」。

		*/
		Id:       458,
		Password: "54752875",
		Name:     "双头雷龙",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 7,
		La:    ygo.LA_Light,   // 光
		Lr:    ygo.LR_Thunder, // 雷
		Atk:   2800,
		Def:   2100,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 79 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1356
		 调整:

		 [异次元的战士]<異次元の戦士>
		 [2010/09/08]
		 ●这张卡和怪兽进行战斗时，那只怪兽和这张卡从游戏中除外。
		 ◇诱发效果（进入连锁）
		 ◇必须发动
		 ◇不取对象
		 ◇效果发动时点同反转怪兽在伤害步骤发动反转效果的时点（具体请参照伤害步骤相关帖）
		 ◇效果处理时这张卡或和这张卡战斗的怪兽不在场上存在的场合，剩下的那只从游戏中除外
		 ◇效果处理时这张卡与进行战斗的怪兽都不在场上存在的场合，这个效果不适用

		 中文名: 异次元的战士
		 卡片种类: 效果怪兽
		 卡片密码: 37043180
		 种族: 战士
		 属性: 地
		 星级: 4
		 攻击力: 1200
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: VOL07、SD14、SD17、SJ2
		 效果: 这张卡和怪兽进行过战斗时，那只怪兽和这张卡从游戏中除外。

		*/
		Id:       1356,
		Password: "37043180",
		Name:     "异次元的战士",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Earth,   // 地
		Lr:    ygo.LR_Warrior, // 战士
		Atk:   1200,
		Def:   1000,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 80 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 188
		 调整:

		 [猎卡死神]<カードを狩る死神>
		 [2010/09/08]
		 ●反转：场上的1张陷阱卡破坏。选择的卡是覆盖的场合，翻开那张卡确认，陷阱卡就破坏。魔法卡的场合变回原样。
		 ◇反转效果（进入连锁）
		 ◇必须发动
		 ◇发动时选择这张卡以外的在自己或对方场上的魔法及陷阱区域存在的1张表侧表示的陷阱卡或里侧表示覆盖着的1张卡（取对象）
		 ◇这张卡被攻击在信息确认时变为表侧的场合，效果延迟到怪兽被破坏确定时之后发动（具体请参考伤害步骤的详细说明）
		 ◇场上不存在能被选择的卡时，这个效果不发动
		 ◇可以选择陷阱怪兽在魔法及陷阱区域的所在
		 ◇不能选择由于「纳祭之魔/サクリファイス」的效果作为装备的「阿匹卜之化神/アポピスの化神」等为对象
		 ◇效果处理时，对象卡不在场上存在的场合，这个效果不适用

		 中文名: 猎卡死神
		 卡片种类: 效果怪兽
		 卡片密码: 33066139
		 种族: 恶魔
		 属性: 暗
		 星级: 5
		 攻击力: 1380
		 防御力: 1930
		 罕见度: 平卡N
		 卡包: PG、BE01、VOL03、DL02
		 效果: 反转：选择场上存在的1张陷阱卡破坏。选择的卡是盖放的场合，把那张卡翻开确认，是陷阱卡则破坏。魔法卡的场合回到原状。

		*/
		Id:       188,
		Password: "33066139",
		Name:     "猎卡死神",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 5,
		La:    ygo.LA_Dark,  // 暗
		Lr:    ygo.LR_Fiend, // 恶魔
		Atk:   1380,
		Def:   1930,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 81 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2080
		 调整:

		 [骸骨天使]<スケルエンジェル>
		 [2010/09/08]
		 ●反转：自己从卡组抽1张卡。
		 ◇反转效果（进入连锁）
		 ◇必须发动
		 ◇这张卡被攻击在信息确认时变为表侧的场合，效果延迟到怪兽被破坏确定时之后发动（具体请参考伤害步骤的详细说明）
		 ◇效果处理时，自己卡组中的卡不足1张的场合，抽卡时决斗败北
		 ◇「神殿守卫者/神殿を守る者」在场上存在的场合这个效果也发动，效果处理时效果不适用

		 中文名: 骸骨天使
		 卡片种类: 效果怪兽
		 卡片密码: 60694662
		 种族: 天使
		 属性: 光
		 星级: 2
		 攻击力: 900
		 防御力: 400
		 罕见度: 平卡N
		 卡包: YSD01、VOL03、YSD04、DB12
		 效果: 反转：从自己卡组抽1张卡。

		*/
		Id:       2080,
		Password: "60694662",
		Name:     "骸骨天使",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 2,
		La:    ygo.LA_Light, // 光
		Lr:    ygo.LR_Fairy, // 天使
		Atk:   900,
		Def:   400,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 82 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 211
		 调整:

		 [青衣忍者]<青い忍者>
		 [2010/09/08]
		 ●反转：场上的1张魔法卡破坏。选择的卡是覆盖的场合，翻开那张卡确认，魔法卡就破坏。陷阱卡的场合变回原样。
		 ◇反转效果（进入连锁）
		 ◇必须发动
		 ◇发动时选择这张卡以外的在自己或对方场上的魔法及陷阱区域存在的1张表侧表示的魔法卡或里侧表示覆盖着的1张卡（取对象）
		 ◇这张卡被攻击在信息确认时变为表侧的场合，效果延迟到怪兽被破坏确定时之后发动（具体请参考伤害步骤的详细说明）
		 ◇场上不存在能被选择的卡时，这个效果不发动
		 ◇可以选择作为装备的怪兽卡
		 ◇可以选择由于「纳祭之魔/サクリファイス」的效果作为装备的「阿匹卜之化神/アポピスの化神」等为对象
		 ◇效果处理时，对象卡不在场上存在的场合，这个效果不适用
		 ◇作为怪兽装备的在魔法及陷阱区域存在的里侧怪兽卡在这个效果处理时破坏

		 中文名: 青衣忍者
		 卡片种类: 效果怪兽
		 卡片密码: 09076207
		 种族: 战士
		 属性: 地
		 星级: 1
		 攻击力: 300
		 防御力: 300
		 罕见度: 平卡N，银字R
		 卡包: PG、BE01、VOL03、DL02
		 效果: 反转：破坏场上的1张魔法卡。选择的卡是盖伏的场合，确认那张卡，如果是魔法卡就破坏。如果选择的卡确认是陷阱卡的场合变回原来的盖伏形式。

		*/
		Id:       211,
		Password: "09076207",
		Name:     "青衣忍者",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 1,
		La:    ygo.LA_Earth,   // 地
		Lr:    ygo.LR_Warrior, // 战士
		Atk:   300,
		Def:   300,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 83 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 212
		 调整:

		 [食人虫]<人喰い虫>
		 [2010/09/08]
		 ●反转：选择场上的1只怪兽破坏。
		 ◇反转效果（进入连锁）
		 ◇必须发动
		 ◇发动时选择自己或对方场上存在的1只怪兽（取对象）
		 ◇对方场上不存在能被选择的怪兽时，这个效果不发动
		 ◇不能选择被战斗破坏确定的卡
		 ◇这张卡被攻击在信息确认时变为表侧的场合，效果延迟到怪兽被破坏确定时之后发动（具体请参考伤害步骤的详细说明）
		 ◇效果处理时对象怪兽不在场上存在的场合，这个效果不适用

		 中文名: 食人虫
		 卡片种类: 效果怪兽
		 卡片密码: 54652250
		 种族: 昆虫
		 属性: 地
		 星级: 2
		 攻击力: 450
		 防御力: 600
		 罕见度: 银字R，面闪SR
		 卡包: PG、BE01、EX-R(EX)、YSD02、VOL03、DL02、KA、SK2
		 效果: 反转：选择场上1只怪兽破坏。

		*/
		Id:       212,
		Password: "54652250",
		Name:     "食人虫",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 2,
		La:    ygo.LA_Earth,  // 地
		Lr:    ygo.LR_Insect, // 昆虫
		Atk:   450,
		Def:   600,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 84 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 214
		 调整:

		 [哈尼哈尼]<ハネハネ>
		 [2010/09/08]
		 ●反转：选择场上的1只怪兽，回到持有者的手卡。
		 ◇反转效果（进入连锁）
		 ◇必须发动
		 ◇发动时选择自己或对方场上存在的1只怪兽（取对象）
		 ◇对方场上不存在能被选择的怪兽时，这个效果不发动
		 ◇不能选择被战斗破坏确定的卡
		 ◇这张卡被攻击在信息确认时变为表侧的场合，效果延迟到怪兽被破坏确定时之后发动（具体请参考伤害步骤的详细说明）
		 ◇效果处理时对象怪兽不在场上存在的场合，这个效果不适用
		 ◇选择的是怪兽衍生物的场合，效果适用时那只衍生物消灭
		 ◇选择的是融合怪兽、同调怪兽的场合，效果适用时那只怪兽回到额外卡组

		 中文名: 哈尼哈尼
		 卡片种类: 效果怪兽
		 卡片密码: 07089711
		 种族: 兽
		 属性: 地
		 星级: 2
		 攻击力: 450
		 防御力: 500
		 罕见度: 平卡N，银字R
		 卡包: PG、BE01、EX-R(EX)、VOL03、DL02
		 效果: 反转：选择场上1只怪兽回到原持有人手卡。

		*/
		Id:       214,
		Password: "07089711",
		Name:     "哈尼哈尼",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 2,
		La:    ygo.LA_Earth, // 地
		Lr:    ygo.LR_Beast, // 兽
		Atk:   450,
		Def:   500,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 85 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 233
		 调整:

		 [飞蛾幼虫]<ラーバモス>
		 [2010/09/08]
		 ●这张卡不能通常召唤。装备有「进化之茧」的「幼虫宝宝」2回合后（计算自己回合）作为祭品来召唤。
		 ◇召唤规则
		 ◇通过正规方式出场后，可以用其他卡的效果从墓地或除外状态特殊召唤

		 中文名: 飞蛾幼虫
		 卡片种类: 效果怪兽
		 卡片密码: 87756343
		 种族: 昆虫
		 属性: 地
		 星级: 2
		 攻击力: 500
		 防御力: 400
		 罕见度: 平卡N
		 卡包: RB、BE01、VOL05、DL02
		 效果: 这张卡通常召唤不能。需要装备了「进化之茧」2回合（数自己的回合数）的「幼虫宝宝」做祭品特殊召唤上场。

		*/
		Id:       233,
		Password: "87756343",
		Name:     "飞蛾幼虫",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 2,
		La:    ygo.LA_Earth,  // 地
		Lr:    ygo.LR_Insect, // 昆虫
		Atk:   500,
		Def:   400,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 86 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 235
		 调整:

		 [鹰身女郎三姐妹]<ハーピィ·レディ三姉妹>
		 [2010/09/08]
		 ●这张卡不能通常召唤。用「万华镜－华丽的分身－」的效果来特殊召唤。
		 ◇召唤规则
		 ◇通过正规方式出场后，可以从墓地或除外状态被其他卡的效果特殊召唤

		 中文名: 鹰身女郎三姐妹
		 卡片种类: 效果怪兽
		 卡片密码: 12206212
		 种族: 鸟兽
		 属性: 风
		 星级: 6
		 攻击力: 1950
		 防御力: 2100
		 罕见度: 平卡N，银字R
		 卡包: RB、BE01、VOL04、DL02、SD08
		 效果: 这张卡通常召唤不能。需要「万华镜-华丽的分身-」的效果特殊召唤。

		*/
		Id:       235,
		Password: "12206212",
		Name:     "鹰身女郎三姐妹",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 6,
		La:    ygo.LA_Wind,        // 风
		Lr:    ygo.LR_WingedBeast, // 鸟兽
		Atk:   1950,
		Def:   2100,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 87 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 237
		 调整:

		 [进化之茧]<進化の繭>
		 [2010/09/08]
		 ●可以从手卡作为装备给场上表侧表示存在的「幼虫宝宝」装备。
		 ◇无召唤条件
		 ◇只能在自己的主要阶段进行
		 ◇在自己的主要阶段可以通过「血之代偿/血の代償」得到通常召唤机会来进行上述操作
		 ◇占用通常召唤次数，但不作为通常召唤
		 ◇不能给对方场上的「幼虫宝宝/プチモス」装备
		 ◇1只「幼虫宝宝/プチモス」只能装备1张「进化之茧/進化の繭」
		 ◇「激流葬/激流葬」「升天之角笛/昇天の角笛」「神之宣告/神の宣告」「魔法干扰阵/マジック·ジャマー」等不能对应这个操作发动
		 ◇适用后作为装备魔法卡
		 ●装备的场合，「幼虫宝宝/プチモス」的攻击力与守备力适用「进化之茧/進化の繭」的数值。
		 ◇装备魔法效果

		 中文名: 进化之茧
		 卡片种类: 效果怪兽
		 卡片密码: 40240595
		 种族: 昆虫
		 属性: 地
		 星级: 3
		 攻击力: 0
		 防御力: 2000
		 罕见度: 平卡N
		 卡包: RB、BE01、VOL04、DL02
		 效果: 这张卡可以在手上当作装备卡使用装备在场上表侧表示的「幼虫宝宝」上。装备的场合，「幼虫宝宝」的攻防数值使用「进化之茧」的数值。

		*/
		Id:       237,
		Password: "40240595",
		Name:     "进化之茧",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 3,
		La:    ygo.LA_Earth,  // 地
		Lr:    ygo.LR_Insect, // 昆虫
		Atk:   0,
		Def:   2000,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 88 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 240
		 调整:

		 [暗之假面]<闇の仮面>
		 [2010/09/08]
		 ●反转：从自己的墓地选择1张陷阱卡。被选择的卡加入自己的手卡。
		 ◇反转效果（进入连锁）
		 ◇必须发动
		 ◇发动时选择自己墓地中的1张陷阱卡（取对象）
		 ◇这张卡被攻击在信息确认时变为表侧的场合，效果延迟到怪兽被破坏确定时之后发动（具体请参考伤害步骤的详细说明）
		 ◇自己的墓地不存在能被选择的卡时，这个效果不发动
		 ◇效果处理时对象卡不在墓地中存在的场合，这个效果不适用

		 中文名: 暗之假面
		 卡片种类: 效果怪兽
		 卡片密码: 28933734
		 种族: 恶魔
		 属性: 暗
		 星级: 2
		 攻击力: 900
		 防御力: 400
		 罕见度: 平卡N，银字R，黄金GR
		 卡包: RB、BE01、YSD02、VOL04、DL02、SD12、Booster R3、GS03、GDB1
		 效果: 反转：选择自己墓地存在的1张陷阱卡加入手卡。

		*/
		Id:       240,
		Password: "28933734",
		Name:     "暗之假面",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 2,
		La:    ygo.LA_Dark,  // 暗
		Lr:    ygo.LR_Fiend, // 恶魔
		Atk:   900,
		Def:   400,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 89 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 241
		 调整:

		 [白衣怪盗]<白い泥棒>
		 [2010/09/08]
		 ●这张卡给予对方玩家战斗伤害时，对方随机丢弃1张手卡。
		 ◇诱发效果（进入连锁）
		 ◇必须发动
		 ◇详细时点请参考伤害步骤详解

		 中文名: 白衣怪盗
		 卡片种类: 效果怪兽
		 卡片密码: 15150365
		 种族: 魔法师
		 属性: 光
		 星级: 3
		 攻击力: 1000
		 防御力: 700
		 罕见度: 平卡N，银字R
		 卡包: RB、BE01、VOL05、DL02、Booster R3
		 效果: 这张卡造成对方玩家基本分伤害的时候，对方随机丢弃1张卡。

		*/
		Id:       241,
		Password: "15150365",
		Name:     "白衣怪盗",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 3,
		La:    ygo.LA_Light,       // 光
		Lr:    ygo.LR_Spellcaster, // 魔法师
		Atk:   1000,
		Def:   700,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 90 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 242
		 调整:

		 [眼球大王]<大王目玉>
		 [2010/09/08]
		 ●反转：从自己的卡组上方看5张卡，可以把那些卡以喜欢的顺序交换位置。
		 ◇反转效果（进入连锁）
		 ◇必须发动
		 ◇这张卡被攻击在信息确认时变为表侧的场合，效果延迟到怪兽被破坏确定时之后发动（具体请参考伤害步骤的详细说明）
		 ◇效果处理时决定是否替换位置

		 中文名: 眼球大王
		 卡片种类: 效果怪兽
		 卡片密码: 16768387
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 1200
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: RB、BE01、VOL05、DL02
		 效果: 反转：查看自己的卡组最上面的5张卡，把这5张卡按自己的意愿顺序放回卡组最上面。

		*/
		Id:       242,
		Password: "16768387",
		Name:     "眼球大王",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Dark,  // 暗
		Lr:    ygo.LR_Fiend, // 恶魔
		Atk:   1200,
		Def:   1000,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 91 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 244
		 调整:

		 [假面魔道士]<仮面魔道士>
		 [2010/09/08]
		 ●这张卡给予对方玩家战斗伤害时，自己抽1张卡。
		 ◇诱发效果（进入连锁）
		 ◇必须发动
		 ◇详细时点请参考伤害步骤详解

		 中文名: 假面魔道士
		 卡片种类: 效果怪兽
		 卡片密码: 10189126
		 种族: 魔法师
		 属性: 暗
		 星级: 4
		 攻击力: 900
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: RB、BE01、VOL05、DL02
		 效果: 这张卡造成对方玩家基本分伤害的时候，自己抽1张卡。

		*/
		Id:       244,
		Password: "10189126",
		Name:     "假面魔道士",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Dark,        // 暗
		Lr:    ygo.LR_Spellcaster, // 魔法师
		Atk:   900,
		Def:   1400,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 92 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 250
		 调整:

		 [雷魔神-桑迦]<雷魔神－サンガ>
		 [2010/09/08]
		 ●对方回合的战斗伤害计算时才能发动。攻击这张卡的怪兽的攻击力变成0。只要这张卡表侧表示在场上存在这个效果只能使用1次。
		 ◇诱发即时效果（进入连锁）
		 ◇任意发动
		 ◇取对象
		 ◇详细时点请参考伤害步骤详解

		 中文名: 雷魔神-桑迦
		 卡片种类: 效果怪兽
		 卡片密码: 25955164
		 种族: 雷
		 属性: 光
		 星级: 7
		 攻击力: 2600
		 防御力: 2200
		 罕见度: 平卡N，银字R
		 卡包: RB、BE01、VOL05、DL02
		 效果: 对方回合的战斗伤害计算时才能发动。攻击这张卡的怪兽的攻击力变为0。只要这张卡在表侧表示存在这个效果只能使用1次。

		*/
		Id:       250,
		Password: "25955164",
		Name:     "雷魔神-桑迦",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 7,
		La:    ygo.LA_Light,   // 光
		Lr:    ygo.LR_Thunder, // 雷
		Atk:   2600,
		Def:   2200,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 93 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 251
		 调整:

		 [风魔神-修迦]<風魔神－ヒューガ>
		 [2010/09/08]
		 ●对方回合的战斗伤害计算时才能发动。攻击这张卡的怪兽的攻击力变成0。只要这张卡表侧表示在场上存在这个效果只能使用1次。
		 ◇诱发即时效果（进入连锁）
		 ◇任意发动
		 ◇取对象
		 ◇详细时点请参考伤害步骤详解

		 中文名: 风魔神-修迦
		 卡片种类: 效果怪兽
		 卡片密码: 62340868
		 种族: 魔法师
		 属性: 风
		 星级: 7
		 攻击力: 2400
		 防御力: 2200
		 罕见度: 平卡N，银字R
		 卡包: RB、BE01、VOL05、DL02
		 效果: 对方回合的战斗伤害计算时才能发动。攻击这张卡的怪兽的攻击力变为0。只要这张卡在表侧表示存在这个效果只能使用1次。

		*/
		Id:       251,
		Password: "62340868",
		Name:     "风魔神-修迦",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 7,
		La:    ygo.LA_Wind,        // 风
		Lr:    ygo.LR_Spellcaster, // 魔法师
		Atk:   2400,
		Def:   2200,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 94 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 252
		 调整:

		 [水魔神-斯迦]<水魔神－スーガ>
		 [2010/09/08]
		 ●对方回合的战斗伤害计算时才能发动。攻击这张卡的怪兽的攻击力变成0。只要这张卡表侧表示在场上存在这个效果只能使用1次。
		 ◇诱发即时效果（进入连锁）
		 ◇任意发动
		 ◇取对象
		 ◇详细时点请参考伤害步骤详解

		 中文名: 水魔神-斯迦
		 卡片种类: 效果怪兽
		 卡片密码: 98434877
		 种族: 水
		 属性: 水
		 星级: 7
		 攻击力: 2500
		 防御力: 2400
		 罕见度: 平卡N，银字R
		 卡包: RB、BE01、VOL05、DL02
		 效果: 对方回合的战斗伤害计算时才能发动。攻击这张卡的怪兽的攻击力变为0。只要这张卡在表侧表示存在这个效果只能使用1次。

		*/
		Id:       252,
		Password: "98434877",
		Name:     "水魔神-斯迦",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 7,
		La:    ygo.LA_Water, // 水
		Lr:    ygo.LR_Water, // 水
		Atk:   2500,
		Def:   2400,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 95 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 253
		 调整:

		 [魔法灯]<魔法のランプ>
		 [2010/09/08]
		 ●这张卡可以直接攻击对方玩家。
		 ◇永续效果（不进入连锁）
		 ◇对方场上表侧表示存在有2只「冲锋陷阵的队长/切り込み隊長」的场合，这张卡可以直接攻击对方玩家

		 中文名: 魔法灯
		 卡片种类: 效果怪兽
		 卡片密码: 98049915
		 种族: 魔法师
		 属性: 暗
		 星级: 1
		 攻击力: 400
		 防御力: 300
		 罕见度: 平卡N
		 卡包: RB、BE01、VOL05、DL02
		 效果: 这张卡可以直接攻击对方玩家。

		*/
		Id:       253,
		Password: "98049915",
		Name:     "魔法灯",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 1,
		La:    ygo.LA_Dark,        // 暗
		Lr:    ygo.LR_Spellcaster, // 魔法师
		Atk:   400,
		Def:   300,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 96 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 254
		 调整:

		 [铁蝎]<鉄のサソリ>
		 [2010/09/08]
		 ●机械族以外的怪兽被这张卡攻击的场合，那只怪兽在（计算对方回合）第3回合的结束阶段破坏。
		 ◇诱发效果（进入连锁）
		 ◇必须发动
		 ◇不取对象
		 ◇在怪兽被战斗破坏送去墓地时发动
		 ◇那只怪兽在这个效果发动后一度不在场上表侧表示存在的场合，这个效果重置
		 ◇那只怪兽在这个效果发动后一度变为机械族怪兽的场合，这个效果重置
		 ◇同1只怪兽被2张「铁蝎/鉄のサソリ」攻击过的场合，先适用的「铁蝎/鉄のサソリ」的效果不会被重置

		 中文名: 铁蝎
		 卡片种类: 效果怪兽
		 卡片密码: 13599884
		 种族: 机械
		 属性: 地
		 星级: 1
		 攻击力: 250
		 防御力: 300
		 罕见度: 平卡N
		 卡包: RB、VOL04、DL02
		 效果: 机械族以外的怪兽攻击这张卡的场合，那只怪兽（以对方的回合来数）第3个回合的回合结束时破坏。

		*/
		Id:       254,
		Password: "13599884",
		Name:     "铁蝎",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 1,
		La:    ygo.LA_Earth,   // 地
		Lr:    ygo.LR_Machine, // 机械
		Atk:   250,
		Def:   300,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 97 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 255
		 调整:

		 [百足虫]<レッグル>
		 [2010/09/08]
		 ●这张卡可以直接攻击对方玩家。
		 ◇永续效果（不进入连锁）
		 ◇对方场上表侧表示存在有2只「冲锋陷阵的队长/切り込み隊長」的场合，这张卡可以直接攻击对方玩家

		 中文名: 百足虫
		 卡片种类: 效果怪兽
		 卡片密码: 12472242
		 种族: 昆虫
		 属性: 地
		 星级: 1
		 攻击力: 300
		 防御力: 350
		 罕见度: 平卡N
		 卡包: RB、VOL05、DL02
		 效果: 这张卡可以直接攻击对方玩家。

		*/
		Id:       255,
		Password: "12472242",
		Name:     "百足虫",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 1,
		La:    ygo.LA_Earth,  // 地
		Lr:    ygo.LR_Insect, // 昆虫
		Atk:   300,
		Def:   350,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 98 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 256
		 调整:

		 [巨口]<ラージマウス>
		 [2010/09/08]
		 ●这张卡可以直接攻击对方玩家。
		 ◇永续效果（不进入连锁）
		 ◇对方场上表侧表示存在有2只「冲锋陷阵的队长/切り込み隊長」的场合，这张卡可以直接攻击对方玩家

		 中文名: 巨口
		 卡片种类: 效果怪兽
		 卡片密码: 58861941
		 种族: 水
		 属性: 水
		 星级: 1
		 攻击力: 300
		 防御力: 250
		 罕见度: 平卡N
		 卡包: RB、VOL05、DL02
		 效果: 这张卡可以直接攻击对方玩家。

		*/
		Id:       256,
		Password: "58861941",
		Name:     "巨口",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 1,
		La:    ygo.LA_Water, // 水
		Lr:    ygo.LR_Water, // 水
		Atk:   300,
		Def:   250,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 99 */
	cardBag.Register(&ygo.CardOriginal{
		/*
			 id: 258
			 调整:

			 [炸弹先生]<ミスター·ボンバー>
			 [2010/09/08]
			 ●自己准备阶段时才能发动。表侧表示的这张卡作为祭品，选择2只攻击力1000以下的表侧表示怪兽破坏。
			 ◇诱发效果（进入连锁）
			 ◇任意发动
			 ◇符合条件的怪兽不足2只的场合不能发动
			 ◇发动时选择符合条件的2只怪兽（取对象）
			 ◇这张卡作为祭品是代价
			 ◇不能选择里侧表示的怪兽为对象★效果处理时，对象之一不在场上表侧表示存在的场合如何处理
			调整中★效果处理时，对象中有攻击力高于1000的场合如何处理
			调整中

			 中文名: 炸弹先生
			 卡片种类: 效果怪兽
			 卡片密码: 70138455
			 种族: 机械
			 属性: 炎
			 星级: 3
			 攻击力: 800
			 防御力: 900
			 罕见度: 平卡N
			 卡包: RB、BE01、VOL05、DL02
			 效果: 在自己的准备阶段时发动。表侧表示的这张卡作为祭品，选择2只表侧表示的攻击力1000以下的怪兽破坏。

		*/
		Id:       258,
		Password: "70138455",
		Name:     "炸弹先生",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 3,
		La:    ygo.LA_Fire,    // 炎
		Lr:    ygo.LR_Machine, // 机械
		Atk:   800,
		Def:   900,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 100 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 259
		 调整:

		 [人造人7号]<人造人間７号>
		 [2010/09/08]
		 ●这张卡可以直接攻击对方玩家。
		 ◇永续效果（不进入连锁）
		 ◇对方场上表侧表示存在有2只「冲锋陷阵的队长/切り込み隊長」的场合，这张卡可以直接攻击对方玩家

		 中文名: 人造人7号
		 卡片种类: 效果怪兽
		 卡片密码: 32809211
		 种族: 机械
		 属性: 暗
		 星级: 2
		 攻击力: 500
		 防御力: 400
		 罕见度: 平卡N
		 卡包: RB、BE01、VOL05、DL02
		 效果: 这张卡可以直接攻击对方玩家。

		*/
		Id:       259,
		Password: "32809211",
		Name:     "人造人7号",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 2,
		La:    ygo.LA_Dark,    // 暗
		Lr:    ygo.LR_Machine, // 机械
		Atk:   500,
		Def:   400,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 101 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 260
		 调整:

		 [神圣魔术师]<聖なる魔術師>
		 [2010/09/08]
		 ●反转：从自己的墓地选择1张魔法卡。被选择的卡加入自己的手卡。
		 ◇反转效果（进入连锁）
		 ◇必须发动
		 ◇发动时选择自己墓地中的1张魔法卡（取对象）
		 ◇这张卡被攻击在信息确认时变为表侧的场合，效果延迟到怪兽被破坏确定时之后发动（具体请参考伤害步骤的详细说明）
		 ◇自己的墓地不存在能被选择的卡时，这个效果不发动
		 ◇效果处理时对象卡不在墓地中存在的场合，这个效果不适用

		 中文名: 神圣魔术师
		 卡片种类: 效果怪兽
		 卡片密码: 31560081
		 种族: 魔法师
		 属性: 光
		 星级: 1
		 攻击力: 300
		 防御力: 400
		 罕见度: 平卡N，银字R
		 卡包: RB、BE01、YSD01、VOL04、DL02、SD06、PE、SY2、SK2
		 效果: 反转：选择自己的墓地的1张魔法卡。选择的卡加入自己手卡。

		*/
		Id:       260,
		Password: "31560081",
		Name:     "神圣魔术师",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 1,
		La:    ygo.LA_Light,       // 光
		Lr:    ygo.LR_Spellcaster, // 魔法师
		Atk:   300,
		Def:   400,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 102 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 262
		 调整:

		 [彩虹花]<レインボー·フラワー>
		 [2010/09/08]
		 ●这张卡可以直接攻击对方玩家。
		 ◇永续效果（不进入连锁）
		 ◇对方场上表侧表示存在有2只「冲锋陷阵的队长/切り込み隊長」的场合，这张卡可以直接攻击对方玩家

		 中文名: 彩虹花
		 卡片种类: 效果怪兽
		 卡片密码: 21347810
		 种族: 植物
		 属性: 地
		 星级: 2
		 攻击力: 400
		 防御力: 500
		 罕见度: 平卡N
		 卡包: RB、VOL05、DL02
		 效果: 这张卡可以直接攻击对方玩家。

		*/
		Id:       262,
		Password: "21347810",
		Name:     "彩虹花",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 2,
		La:    ygo.LA_Earth, // 地
		Lr:    ygo.LR_Plant, // 植物
		Atk:   400,
		Def:   500,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 103 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 264
		 调整:

		 [电气蜥蜴]<でんきトカゲ>
		 [2010/09/08]
		 ●不死族以外的怪兽被这张卡攻击的场合，那只怪兽在下个回合不能攻击宣言。
		 ◇诱发效果（进入连锁）
		 ◇必须发动
		 ◇不取对象
		 ◇在怪兽被战斗破坏送去墓地时发动
		 ◇那只怪兽在这个效果发动后一度不在场上表侧表示存在的场合，这个效果重置
		 ◇那只怪兽在这个效果发动后一度变为不死族怪兽的场合，这个效果重置

		 中文名: 电气蜥蜴
		 卡片种类: 效果怪兽
		 卡片密码: 55875323
		 种族: 雷
		 属性: 地
		 星级: 3
		 攻击力: 850
		 防御力: 800
		 罕见度: 平卡N
		 卡包: RB、VOL04、DL02
		 效果: 不死族以外的怪兽攻击这张卡的场合，那只怪兽下一个回合不能攻击宣言。

		*/
		Id:       264,
		Password: "55875323",
		Name:     "电气蜥蜴",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 3,
		La:    ygo.LA_Earth,   // 地
		Lr:    ygo.LR_Thunder, // 雷
		Atk:   850,
		Def:   800,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 104 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 266
		 调整:

		 [女王的影武者]<女王の影武者>
		 [2010/09/08]
		 ●这张卡可以直接攻击对方玩家。
		 ◇永续效果（不进入连锁）
		 ◇对方场上表侧表示存在有2只「冲锋陷阵的队长/切り込み隊長」的场合，这张卡可以直接攻击对方玩家

		 中文名: 女王的影武者
		 卡片种类: 效果怪兽
		 卡片密码: 05901497
		 种族: 战士
		 属性: 地
		 星级: 1
		 攻击力: 350
		 防御力: 300
		 罕见度: 平卡N
		 卡包: RB、VOL05、DL02
		 效果: 这张卡可以直接攻击对方玩家。

		*/
		Id:       266,
		Password: "05901497",
		Name:     "女王的影武者",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 1,
		La:    ygo.LA_Earth,   // 地
		Lr:    ygo.LR_Warrior, // 战士
		Atk:   350,
		Def:   300,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 105 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 410
		 调整:

		 [野蛮人2号]<バーバリアン２号>
		 [2010/09/08]
		 ●自己的场上每有1只表侧表示存在的「野蛮人１号」，这张卡的攻击力上升500分。
		 ◇永续效果（不进入连锁）

		 中文名: 野蛮人2号
		 卡片种类: 效果怪兽
		 卡片密码: 40453765
		 种族: 战士
		 属性: 地
		 星级: 5
		 攻击力: 1800
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: ME、VOL07、DL04
		 效果: 每有1只自己场上表侧表示存在的「野蛮人1号」，这张卡的攻击力上升500。

		*/
		Id:       410,
		Password: "40453765",
		Name:     "野蛮人2号",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 5,
		La:    ygo.LA_Earth,   // 地
		Lr:    ygo.LR_Warrior, // 战士
		Atk:   1800,
		Def:   1500,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 106 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 414
		 调整:

		 [壶魔人]<壺魔人>
		 [2010/09/08]
		 ●反转：场上表侧表示存在的「龙族封印之壶」破坏。破坏的场合，场上表侧表示存在的龙族怪兽全部变为攻击表示。
		 ◇反转效果（进入连锁）
		 ◇必须发动
		 ◇效果处理时场上表侧表示存在的「龙族封印之壶/ドラゴン族·封印の壺」全部破坏

		 中文名: 壶魔人
		 卡片种类: 效果怪兽
		 卡片密码: 55763552
		 种族: 炎
		 属性: 炎
		 星级: 3
		 攻击力: 200
		 防御力: 1800
		 罕见度: 平卡N
		 卡包: ME、BE02、VOL06、DL04、PE
		 效果: 反转：场上表侧表示的「龙族封印之壶」破坏。破坏时，场上表侧表示存在的龙族怪兽全部攻击表示。

		*/
		Id:       414,
		Password: "55763552",
		Name:     "壶魔人",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 3,
		La:    ygo.LA_Fire, // 炎
		Lr:    ygo.LR_Fire, // 炎
		Atk:   200,
		Def:   1800,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 107 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 416
		 调整:

		 [三眼怪]<クリッター>
		 [2010/09/08]
		 ●这张卡从场上送去墓地时，从自己的卡组把1只攻击力1500以下的怪兽加入手卡。
		 ◇诱发效果（进入连锁）
		 ◇必须发动
		 ◇效果处理时从卡组中选择合适的卡加入手卡，并且让对方确认，之后卡组规则洗切（洗切行为不占用时点）
		 ◇效果处理时，不能选择攻击力记载为？的怪兽
		 ◇这张卡在对方场上被送去墓地的场合，在自己的墓地发动效果，以自己的效果来处理
		 ◇在一组连锁处理完毕前，这张卡被送去墓地的场合，在那组连锁处理完毕后另开连锁发动这个效果
		 ◇作为装备卡在场上的这张卡被送去墓地的场合，这个效果也发动
		 ◇这张卡的效果发动后，这张卡从墓地离开的场合这个效果也适用
		 ◇被战斗破坏的这张卡，在伤害步骤中的怪兽送去墓地时点发动效果（具体请参考置顶中的伤害步骤详解）

		 中文名: 三眼怪
		 卡片种类: 效果怪兽
		 卡片密码: 26202165
		 种族: 恶魔
		 属性: 暗
		 星级: 3
		 攻击力: 1000
		 防御力: 600
		 罕见度: 平卡N，银字R，黄金GR，金字UR
		 卡包: ME、BE02、VOL06、DL04、GS01、GLD02、JY、KA、PE、SY2、SD21、TU06、DB12、ST12、DS13、GDB1
		 效果: 这张卡从场上送去墓地时，从自己卡组把1只攻击力1500以下的怪兽加入手卡。

		*/
		Id:       416,
		Password: "26202165",
		Name:     "三眼怪",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 3,
		La:    ygo.LA_Dark,  // 暗
		Lr:    ygo.LR_Fiend, // 恶魔
		Atk:   1000,
		Def:   600,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 108 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 417
		 调整:

		 [大飞蛾]<グレート·モス>
		 [2010/09/08]
		 ●把装备了「进化之茧」4回合（计算自己回合）后的「幼虫宝宝」作为祭品可以特殊召唤。
		 ◇召唤规则
		 ◇不能通常召唤
		 ◇经过正规出场后，可以从墓地或除外状态用其他卡的效果来特殊召唤

		 中文名: 大飞蛾
		 卡片种类: 效果怪兽
		 卡片密码: 14141448
		 种族: 昆虫
		 属性: 地
		 星级: 8
		 攻击力: 2600
		 防御力: 2500
		 罕见度: 平卡N
		 卡包: ME、BE02、VOL06、DL04
		 效果: 装备了「进化之茧」的「幼虫宝宝」4回合后（用自己的回合来数）作祭品来特殊召唤。

		*/
		Id:       417,
		Password: "14141448",
		Name:     "大飞蛾",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 8,
		La:    ygo.LA_Earth,  // 地
		Lr:    ygo.LR_Insect, // 昆虫
		Atk:   2600,
		Def:   2500,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 109 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 418
		 调整:

		 中文名: 栗子球
		 卡片种类: 效果怪兽
		 卡片密码: 40640057
		 种族: 恶魔
		 属性: 暗
		 星级: 1
		 攻击力: 300
		 防御力: 200
		 罕见度: 平卡N，银字R
		 卡包: ME、BE02、YSD01、VOL07、DL04、DT01、DPYG、DTP01、YU、SY2
		 效果: 从手卡丢弃这张卡，自己受到的战斗伤害1次为0。这个效果只能在对方的战斗回合使用。

		*/
		Id:       418,
		Password: "40640057",
		Name:     "栗子球",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 1,
		La:    ygo.LA_Dark,  // 暗
		Lr:    ygo.LR_Fiend, // 恶魔
		Atk:   300,
		Def:   200,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 110 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 423
		 调整:

		 [弹射龟]<カタパルト·タートル>
		 [2010/09/08]
		 ●自己场上存在的1只怪兽作为祭品。给予对方玩家那只怪兽攻击力的一半的伤害。
		 ◇启动效果（进入连锁）
		 ◇发动时把自己场上存在的1只怪兽作为祭品（代价）
		 ◇计算攻击力的参照值是那只祭品怪兽在场上的攻击力
		 ◇可以把自己作为祭品发动这个效果

		 中文名: 弹射龟
		 卡片种类: 效果怪兽
		 卡片密码: 95727991
		 种族: 水
		 属性: 水
		 星级: 5
		 攻击力: 1000
		 防御力: 2000
		 罕见度: 平卡N，银字R，面闪SR
		 卡包: ME、BE02、VOL07、DL04、DPYG
		 效果: 自己场上存在的1只怪兽做祭品。对方受到这只怪兽攻击力一半的伤害。

		*/
		Id:       423,
		Password: "95727991",
		Name:     "弹射龟",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 5,
		La:    ygo.LA_Water, // 水
		Lr:    ygo.LR_Water, // 水
		Atk:   1000,
		Def:   2000,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 111 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 431
		 调整:

		 [剑之女王]<剣の女王>
		 [2010/09/08]
		 ●反转：对方场上每有1张魔法·陷阱卡，给予对方基本分500分的伤害。
		 ◇反转效果（进入连锁）
		 ◇必须发动
		 ◇这张卡被攻击在信息确认时变为表侧的场合，效果延迟到怪兽被破坏确定时之后发动（具体请参考伤害步骤的详细说明）
		 ◇效果处理时计算对方场上的魔法·陷阱卡数量
		 ◇效果处理时对方场上不存在魔法·陷阱卡的场合，这个效果不适用

		 中文名: 剑之女王
		 卡片种类: 效果怪兽
		 卡片密码: 51371017
		 种族: 战士
		 属性: 风
		 星级: 3
		 攻击力: 900
		 防御力: 700
		 罕见度: 平卡N，银字R
		 卡包: ME、BE02、YSD01、VOL07、DL04、Booster07
		 效果: 反转：对方场上每存在1张魔法·陷阱，对方受到500分的伤害。

		*/
		Id:       431,
		Password: "51371017",
		Name:     "剑之女王",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 3,
		La:    ygo.LA_Wind,    // 风
		Lr:    ygo.LR_Warrior, // 战士
		Atk:   900,
		Def:   700,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 112 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 434
		 调整:

		 [影之食尸鬼]<シャドウ·グール>
		 [2010/09/08]
		 ●自己的墓地每存在1只怪兽，这张卡的攻击力上升100分。
		 ◇永续效果（不进入连锁）

		 中文名: 影之食尸鬼
		 卡片种类: 效果怪兽
		 卡片密码: 30778711
		 种族: 不死
		 属性: 暗
		 星级: 5
		 攻击力: 1600
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: ME、BE02、VOL06、DL04
		 效果: 自己的墓地每存在1只怪兽，这张卡的攻击力上升100。

		*/
		Id:       434,
		Password: "30778711",
		Name:     "影之食尸鬼",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 5,
		La:    ygo.LA_Dark,   // 暗
		Lr:    ygo.LR_Zombie, // 不死
		Atk:   1600,
		Def:   1300,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 113 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 438
		 调整:

		 [雷龙]<サンダー·ドラゴン>
		 [2010/09/08]
		 ●可以从手卡把这张卡丢弃，从卡组把另外的最多2张「雷龙」加入手卡。那之后卡组洗切。这个效果只能在自己的主要阶段使用。
		 ◇启动效果（进入连锁）
		 ◇发动时把这张卡从手卡丢弃（代价）
		 ◇效果处理时至少选择1张加入手牌
		 ◇确认卡组中不存在可以检索的卡时，不能发动这个效果

		 中文名: 雷龙
		 卡片种类: 效果怪兽
		 卡片密码: 31786629
		 种族: 雷
		 属性: 光
		 星级: 5
		 攻击力: 1600
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: ME、BE02、VOL07、DL04、Booster07
		 效果: 丢弃手卡的这张卡，从卡组选出最多2张的「雷龙」加入手卡。之后洗卡。这个效果只能在自己的主要阶段使用。

		*/
		Id:       438,
		Password: "31786629",
		Name:     "雷龙",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 5,
		La:    ygo.LA_Light,   // 光
		Lr:    ygo.LR_Thunder, // 雷
		Atk:   1600,
		Def:   1500,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 114 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 440
		 调整:

		 [雷仙人]<雷仙人>
		 [2010/09/08]
		 ●反转：回复3000基本分。
		 ◇反转效果（进入连锁）
		 ◇必须发动
		 ◇这张卡被攻击在信息确认时变为表侧的场合，效果延迟到怪兽被破坏确定时之后发动（具体请参考伤害步骤的详细说明）
		 ●这张卡从场上送去墓地时，失去5000基本分。
		 ◇诱发效果（进入连锁）
		 ◇必须发动
		 ◇反转效果没有发动，这张卡被直接送去墓地的场合，这个效果不发动
		 ◇失去基本分不是伤害效果
		 ◇在墓地发动的效果

		 中文名: 雷仙人
		 卡片种类: 效果怪兽
		 卡片密码: 84926738
		 种族: 雷
		 属性: 光
		 星级: 4
		 攻击力: 1500
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: ME、BE02、VOL07、DL04
		 效果: 反转：基本分回复3000。这张卡从场上送去墓地的时候，基本分失去5000分。

		*/
		Id:       440,
		Password: "84926738",
		Name:     "雷仙人",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Light,   // 光
		Lr:    ygo.LR_Thunder, // 雷
		Atk:   1500,
		Def:   1300,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 115 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 442
		 调整:

		 [超级流星]<スーパースター>
		 [2010/09/08]
		 ●只要这张卡在场上表侧表示存在，全部光属性怪兽的攻击力上升500分。暗属性怪兽的攻击力下降400分。
		 ◇永续效果（不进入连锁）

		 中文名: 超级流星
		 卡片种类: 效果怪兽
		 卡片密码: 67629977
		 种族: 天使
		 属性: 光
		 星级: 2
		 攻击力: 500
		 防御力: 700
		 罕见度: 平卡N，银字R
		 卡包: ME、BE02、VOL06、DL04、Booster06
		 效果: 只要这张卡场上表侧表示存在，全部的光属性怪兽攻击力上升500。暗属性的怪兽攻击力下降400。

		*/
		Id:       442,
		Password: "67629977",
		Name:     "超级流星",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 2,
		La:    ygo.LA_Light, // 光
		Lr:    ygo.LR_Fairy, // 天使
		Atk:   500,
		Def:   700,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 116 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 446
		 调整:

		 [加农炮兵]<キャノン·ソルジャー>
		 [2010/09/08]
		 ●每把自己场上存在的1只怪兽作为祭品，给予对方基本分500分的伤害。
		 ◇启动效果（进入连锁）
		 ◇发动时把自己场上存在的1只怪兽作为祭品（代价）

		 中文名: 加农炮兵
		 卡片种类: 效果怪兽
		 卡片密码: 11384280
		 种族: 机械
		 属性: 暗
		 星级: 4
		 攻击力: 1400
		 防御力: 1300
		 罕见度: 平卡N，银字R，面闪SR
		 卡包: ME、BE02、VOL06、DL04、SD10
		 效果: 每祭品1只自己场上存在的怪兽，对方受到500分的基本分伤害。

		*/
		Id:       446,
		Password: "11384280",
		Name:     "加农炮兵",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Dark,    // 暗
		Lr:    ygo.LR_Machine, // 机械
		Atk:   1400,
		Def:   1300,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 117 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 447
		 调整:

		 [姆卡姆卡]<ムカムカ>
		 [2010/09/08]
		 ●只要这张卡在场上表侧表示存在，控制者每有1张手卡这张卡的攻击力和守备力上升300分。
		 ◇永续效果（不进入连锁）

		 中文名: 姆卡姆卡
		 卡片种类: 效果怪兽
		 卡片密码: 46657337
		 种族: 岩石
		 属性: 地
		 星级: 2
		 攻击力: 600
		 防御力: 300
		 罕见度: 平卡N
		 卡包: ME、BE02、VOL06、DL04
		 效果: 只要这张卡表侧表示在场上存在，控制者手上每有1张卡，这张卡的攻击力·守备力上升300。

		*/
		Id:       447,
		Password: "46657337",
		Name:     "姆卡姆卡",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 2,
		La:    ygo.LA_Earth, // 地
		Lr:    ygo.LR_Rock,  // 岩石
		Atk:   600,
		Def:   300,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 118 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 449
		 调整:

		 [海星小子]<スター·ボーイ>
		 [2010/09/08]
		 ●只要这张卡在场上表侧表示存在，全部水属性怪兽的攻击力上升500分。炎属性怪兽的攻击力下降400分。
		 ◇永续效果（不进入连锁）

		 中文名: 海星小子
		 卡片种类: 效果怪兽
		 卡片密码: 08201910
		 种族: 水
		 属性: 水
		 星级: 2
		 攻击力: 550
		 防御力: 500
		 罕见度: 平卡N
		 卡包: ME、BE02、VOL06、DL04、Booster06、SD04
		 效果: 只要这张卡在场上表侧表示存在，全部水属性的怪兽攻击力上升500。炎属性的怪兽攻击力下降400。

		*/
		Id:       449,
		Password: "08201910",
		Name:     "海星小子",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 2,
		La:    ygo.LA_Water, // 水
		Lr:    ygo.LR_Water, // 水
		Atk:   550,
		Def:   500,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 119 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 450
		 调整:

		 [金毛犼]<ミリス·レディエント>
		 [2010/09/08]
		 ●只要这张卡在场上表侧表示存在，全部地属性怪兽的攻击力上升500分。风属性怪兽的攻击力下降400分。
		 ◇永续效果（不进入连锁）

		 中文名: 金毛犼
		 卡片种类: 效果怪兽
		 卡片密码: 07489323
		 种族: 兽
		 属性: 地
		 星级: 1
		 攻击力: 300
		 防御力: 250
		 罕见度: 平卡N
		 卡包: ME、BE02、VOL06、DL04、JY
		 效果: 只要这张卡在场上表侧表示存在，全部地属性的怪兽攻击力上升500。风属性的怪兽攻击力下降400。

		*/
		Id:       450,
		Password: "07489323",
		Name:     "金毛犼",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 1,
		La:    ygo.LA_Earth, // 地
		Lr:    ygo.LR_Beast, // 兽
		Atk:   300,
		Def:   250,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 120 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 451
		 调整:

		 [暗精灵]<ダーク·エルフ>
		 [2010/09/08]
		 ●这张卡不支付1000基本分就不能攻击。
		 ◇进行攻击宣言的代价（不进入连锁）
		 ◇「技能吸收/スキルドレイン」的效果适用中，进行攻击宣言不需要支付基本分

		 中文名: 暗精灵
		 卡片种类: 效果怪兽
		 卡片密码: 21417692
		 种族: 魔法师
		 属性: 暗
		 星级: 4
		 攻击力: 2000
		 防御力: 800
		 罕见度: 平卡N
		 卡包: ME、VOL07、DL04
		 效果: 这张卡不支付1000基本分不能攻击。

		*/
		Id:       451,
		Password: "21417692",
		Name:     "暗精灵",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Dark,        // 暗
		Lr:    ygo.LR_Spellcaster, // 魔法师
		Atk:   2000,
		Def:   800,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 121 */
	cardBag.Register(&ygo.CardOriginal{
		/*
			 id: 452
			 调整:

			 [蘑菇怪人]<マタンゴ>
			 [2010/09/08]
			 ●每次自己的准备阶段，给予控制者300分的伤害。
			 ◇诱发效果（进入连锁）
			 ◇必须发动
			 ●在自己的结束阶段支付500基本分，这张卡的控制权转移给对方。
			 ◇诱发效果（进入连锁）
			 ◇任意发动
			 ◇发动时支付500基本分（代价）★效果处理时这张卡变为里侧表示的场合如何处理
			调整中★效果处理时对方怪兽区域不存在空位的场合如何处理
			调整中

			 中文名: 蘑菇怪人
			 卡片种类: 效果怪兽
			 卡片密码: 93900406
			 种族: 战士
			 属性: 地
			 星级: 3
			 攻击力: 1250
			 防御力: 800
			 罕见度: 平卡N
			 卡包: ME、VOL07、DL04
			 效果: 每次的自己准备阶段受到300分的伤害。可以在自己的结束阶段支付500分，把这张卡的控制权转移给对方。

		*/
		Id:       452,
		Password: "93900406",
		Name:     "蘑菇怪人",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 3,
		La:    ygo.LA_Earth,   // 地
		Lr:    ygo.LR_Warrior, // 战士
		Atk:   1250,
		Def:   800,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 122 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 453
		 调整:

		 [野蛮人1号]<バーバリアン１号>
		 [2010/09/08]
		 ●自己的场上每有1只表侧表示存在的「野蛮人２号」，这张卡的攻击力上升500分。
		 ◇永续效果（不进入连锁）

		 中文名: 野蛮人1号
		 卡片种类: 效果怪兽
		 卡片密码: 20394040
		 种族: 战士
		 属性: 地
		 星级: 5
		 攻击力: 1550
		 防御力: 1800
		 罕见度: 平卡N
		 卡包: ME、VOL07、DL04
		 效果: 每有1只自己场上表侧表示存在的「野蛮人2号」，这张卡的攻击力上升500。

		*/
		Id:       453,
		Password: "20394040",
		Name:     "野蛮人1号",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 5,
		La:    ygo.LA_Earth,   // 地
		Lr:    ygo.LR_Warrior, // 战士
		Atk:   1550,
		Def:   1800,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 123 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 454
		 调整:

		 [黑森林的魔女]<黒き森のウィッチ>
		 [2010/09/08]
		 ●这张卡从场上送去墓地时，从自己的卡组选择1只守备力1500以下的怪兽，互相确认后加入手卡。那之后卡组洗切。
		 ◇诱发效果（进入连锁）
		 ◇效果处理时从卡组中选择合适的卡加入手卡，并且让对方确认，之后卡组规则洗切（洗切行为不占用时点）
		 ◇效果处理时，不能选择守备力记载为？的怪兽
		 ◇这张卡在对方场上被送去墓地的场合，在自己的墓地发动效果，以自己的效果来处理
		 ◇在一组连锁处理完毕前，这张卡被送去墓地的场合，在那组连锁处理完毕后另开连锁发动这个效果
		 ◇作为装备卡在场上的这张卡被送去墓地的场合，这个效果也发动
		 ◇这张卡的效果发动后，这张卡从墓地离开的场合这个效果也适用
		 ◇被战斗破坏的这张卡，在伤害步骤中的怪兽送去墓地时点发动效果（具体请参考置顶中的伤害步骤详解）

		 中文名: 黑森林的魔女
		 卡片种类: 效果怪兽
		 卡片密码: 78010363
		 种族: 魔法师
		 属性: 暗
		 星级: 4
		 攻击力: 1100
		 防御力: 1200
		 罕见度: 银字R
		 卡包: ME、BE02、VOL06、DL04、JY、PE
		 效果: 这张卡从场上送去墓地时，选1只自己卡组的守备力1500以下的怪兽，互相确认后加入自己手卡。之后洗卡组。

		*/
		Id:       454,
		Password: "78010363",
		Name:     "黑森林的魔女",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Dark,        // 暗
		Lr:    ygo.LR_Spellcaster, // 魔法师
		Atk:   1100,
		Def:   1200,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 124 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 455
		 调整:

		 [小奇美拉]<リトル·キメラ>
		 [2010/09/08]
		 ●只要这张卡在场上表侧表示存在，全部炎属性怪兽的攻击力上升500分。水属性怪兽的攻击力下降400分。
		 ◇永续效果（不进入连锁）

		 中文名: 小奇美拉
		 卡片种类: 效果怪兽
		 卡片密码: 68658728
		 种族: 兽
		 属性: 炎
		 星级: 2
		 攻击力: 600
		 防御力: 550
		 罕见度: 平卡N
		 卡包: ME、BE02、VOL06、DL04、Booster06、SD03、SD24
		 效果: 只要这张卡在场上表侧表示存在，全部炎属性的怪兽攻击力上升500。水属性的怪兽攻击力下降400。

		*/
		Id:       455,
		Password: "68658728",
		Name:     "小奇美拉",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 2,
		La:    ygo.LA_Fire,  // 炎
		Lr:    ygo.LR_Beast, // 兽
		Atk:   600,
		Def:   550,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 125 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 456
		 调整:

		 [刃蝇]<ブレードフライ>
		 [2010/09/08]
		 ●只要这张卡在场上表侧表示存在，全部风属性怪兽的攻击力上升500分。地属性怪兽的攻击力下降400分。
		 ◇永续效果（不进入连锁）

		 中文名: 刃蝇
		 卡片种类: 效果怪兽
		 卡片密码: 28470714
		 种族: 昆虫
		 属性: 风
		 星级: 2
		 攻击力: 600
		 防御力: 700
		 罕见度: 平卡N
		 卡包: ME、BE02、VOL06、DL04、Booster06、SD08
		 效果: 只要这张卡在场上表侧表示存在，全部风属性的怪兽攻击力上升500。地属性的怪兽攻击力下降400。

		*/
		Id:       456,
		Password: "28470714",
		Name:     "刃蝇",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 2,
		La:    ygo.LA_Wind,   // 风
		Lr:    ygo.LR_Insect, // 昆虫
		Atk:   600,
		Def:   700,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 126 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 459
		 调整:

		 [见习魔女]<見習い魔女>
		 [2010/09/08]
		 ●只要这张卡在场上表侧表示存在，全部暗属性怪兽的攻击力上升500分。光属性怪兽的攻击力下降400分。
		 ◇永续效果（不进入连锁）

		 中文名: 见习魔女
		 卡片种类: 效果怪兽
		 卡片密码: 80741828
		 种族: 魔法师
		 属性: 暗
		 星级: 2
		 攻击力: 550
		 防御力: 500
		 罕见度: 平卡N
		 卡包: ME、BE02、VOL06、DL04、Booster06
		 效果: 只要这张卡在场上表侧表示存在，全部暗属性的怪兽攻击力上升500。光属性的怪兽攻击力下降400。

		*/
		Id:       459,
		Password: "80741828",
		Name:     "见习魔女",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 2,
		La:    ygo.LA_Dark,        // 暗
		Lr:    ygo.LR_Spellcaster, // 魔法师
		Atk:   550,
		Def:   500,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 127 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 462
		 调整:

		 [左轮手枪龙]<リボルバー·ドラゴン>
		 [2010/09/08]
		 ●选择对方场上存在的1只怪兽发动。指3次硬币，那之内2次以上正面的场合，那只怪兽破坏。这个效果1回合只能使用1次。
		 ◇启动效果（进入连锁）
		 ◇发动时选择对方场上存在的1只怪兽（取对象）

		 中文名: 左轮手枪龙
		 卡片种类: 效果怪兽
		 卡片密码: 81480460
		 种族: 机械
		 属性: 暗
		 星级: 7
		 攻击力: 2600
		 防御力: 2200
		 罕见度: 面闪SR，金字UR，爆闪PR，银碎SER，立体UTR
		 卡包: ME、BE02、VOL07、VB05、DL04、302、DT03
		 效果: 掷3次硬币。在3次中有2次以上掷到正面的形式，破坏对方场上的1只怪兽。这个效果每个回合只能用1次，在自己的主要阶段使用。

		*/
		Id:       462,
		Password: "81480460",
		Name:     "左轮手枪龙",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 7,
		La:    ygo.LA_Dark,    // 暗
		Lr:    ygo.LR_Machine, // 机械
		Atk:   2600,
		Def:   2200,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 128 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 470
		 调整:

		 [不幸的美少女]<薄幸の美少女>
		 [2010/09/08]
		 ●这张卡被战斗破坏并送去墓地时，那个回合的战斗阶段结束。
		 ◇诱发效果（进入连锁）
		 ◇必须发动

		 中文名: 不幸的美少女
		 卡片种类: 效果怪兽
		 卡片密码: 51275027
		 种族: 魔法师
		 属性: 光
		 星级: 1
		 攻击力: 0
		 防御力: 100
		 罕见度: 平卡N
		 卡包: ME、BE02、VOL07、DL04
		 效果: 这张卡被战斗破坏送去墓地，当时那个回合的战斗阶段结束。

		*/
		Id:       470,
		Password: "51275027",
		Name:     "不幸的美少女",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 1,
		La:    ygo.LA_Light,       // 光
		Lr:    ygo.LR_Spellcaster, // 魔法师
		Atk:   0,
		Def:   100,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 129 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1221
		 调整:

		 中文名: 单摆刃拷问机械
		 卡片种类: 通常怪兽
		 卡片密码: 24433920
		 种族: 机械
		 属性: 暗
		 星级: 6
		 攻击力: 1750
		 防御力: 2000
		 罕见度: 平卡N，金字UR
		 卡包: LE02、VOL07、GLD04
		 效果: 描述：使用带有大摆子的刀刃将对方劈成两半！可怕的拷问机器呀！

		*/
		Id:       1221,
		Password: "24433920",
		Name:     "单摆刃拷问机械",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   6,
		La:      ygo.LA_Dark,    // 暗
		Lr:      ygo.LR_Machine, // 机械
		Atk:     1750,
		Def:     2000,
		IsValid: true,
	})

	/* 130 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 140
		 调整:

		 中文名: 独眼巨人
		 卡片种类: 通常怪兽
		 卡片密码: 76184692
		 种族: 兽战士
		 属性: 地
		 星级: 4
		 攻击力: 1200
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: LB、BE01、EX-R(EX)、VOL01、DL02、DPKB
		 效果: 描述：只长有一只眼睛的巨人。利用巨腕殴打敌人，值得小心。

		*/
		Id:       140,
		Password: "76184692",
		Name:     "独眼巨人",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth,        // 地
		Lr:      ygo.LR_BeastWarrior, // 兽战士
		Atk:     1200,
		Def:     1000,
		IsValid: true,
	})

	/* 131 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 143
		 调整:

		 中文名: 黑魔术师
		 卡片种类: 通常怪兽
		 卡片密码: 46986414
		 种族: 魔法师
		 属性: 暗
		 星级: 7
		 攻击力: 2500
		 防御力: 2100
		 罕见度: 平卡N，银字R，金字UR，爆闪PR，立体UTR
		 卡包: LB、BE01、PP04、LN、EX-R(EX)、VOL01、DL02、SD06、DT01、WJMP、DPYG、DTP01、YU、SY2、LC01
		 效果: 描述：作为魔法师，攻击力·守备力都是最高级别。

		*/
		Id:       143,
		Password: "46986414",
		Name:     "黑魔术师",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   7,
		La:      ygo.LA_Dark,        // 暗
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     2500,
		Def:     2100,
		IsValid: true,
	})

	/* 132 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 144
		 调整:

		 中文名: 暗黑骑士 盖亚
		 卡片种类: 通常怪兽
		 卡片密码: 06368038
		 种族: 战士
		 属性: 地
		 星级: 7
		 攻击力: 2300
		 防御力: 2100
		 罕见度: 银字R，金字UR，立体UTR
		 卡包: LB、BE01、LE02、PH、EX-R(EX)、VOL01、DL02、Booster R1
		 效果: 描述：骑着风驰电掣般的马的骑士。当心突进攻击。

		*/
		Id:       144,
		Password: "06368038",
		Name:     "暗黑骑士 盖亚",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   7,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     2300,
		Def:     2100,
		IsValid: true,
	})

	/* 133 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 147
		 调整:

		 中文名: 猛犸的墓场
		 卡片种类: 通常怪兽
		 卡片密码: 40374923
		 种族: 恐龙
		 属性: 地
		 星级: 3
		 攻击力: 1200
		 防御力: 800
		 罕见度: 平卡N
		 卡包: LB、BE01、EX-R(EX)、VOL01、DL02
		 效果: 描述：守护着伙伴墓地的猛犸，对于任何胆敢践踏墓地的盗墓者进行毫不客气的攻击。

		*/
		Id:       147,
		Password: "40374923",
		Name:     "猛犸的墓场",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,    // 地
		Lr:      ygo.LR_Dinosaur, // 恐龙
		Atk:     1200,
		Def:     800,
		IsValid: true,
	})

	/* 134 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 148
		 调整:

		 中文名: 银牙狼
		 卡片种类: 通常怪兽
		 卡片密码: 90357090
		 种族: 兽
		 属性: 地
		 星级: 3
		 攻击力: 1200
		 防御力: 800
		 罕见度: 平卡N
		 卡包: LB、BE01、EX-R(EX)、VOL01、DL02、YU
		 效果: 描述：闪烁着银光的狼，看起来虽美丽，但性格却十分凶暴。

		*/
		Id:       148,
		Password: "90357090",
		Name:     "银牙狼",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Atk:     1200,
		Def:     800,
		IsValid: true,
	})

	/* 135 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 181
		 调整:

		 中文名: 圣精灵
		 卡片种类: 通常怪兽
		 卡片密码: 15025844
		 种族: 魔法师
		 属性: 光
		 星级: 4
		 攻击力: 800
		 防御力: 2000
		 罕见度: 平卡N
		 卡包: PG、BE01、EX-R(EX)、VOL02、DL02、SY2
		 效果: 描述：虽是纤弱的精灵，以神圣力量守护自身，守备非常高。

		*/
		Id:       181,
		Password: "15025844",
		Name:     "圣精灵",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Light,       // 光
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     800,
		Def:     2000,
		IsValid: true,
	})

	/* 136 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 182
		 调整:

		 中文名: 大炮鸟
		 卡片种类: 通常怪兽
		 卡片密码: 72842870
		 种族: 鸟兽
		 属性: 风
		 星级: 4
		 攻击力: 1200
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: PG、VOL02、DL02、TP17
		 效果: 描述：从口中打出炮弹攻击远处。在山上的炮击很强劲。

		*/
		Id:       182,
		Password: "72842870",
		Name:     "大炮鸟",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Wind,        // 风
		Lr:      ygo.LR_WingedBeast, // 鸟兽
		Atk:     1200,
		Def:     1400,
		IsValid: true,
	})

	/* 137 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 183
		 调整:

		 中文名: 路易斯
		 卡片种类: 通常怪兽
		 卡片密码: 32452818
		 种族: 兽战士
		 属性: 地
		 星级: 4
		 攻击力: 1200
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: PG、BE01、EX-R(EX)、VOL03、DL02、YU、SY2
		 效果: 描述：身体虽小，但在草原上具备高强的守备力。

		*/
		Id:       183,
		Password: "32452818",
		Name:     "路易斯",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth,        // 地
		Lr:      ygo.LR_BeastWarrior, // 兽战士
		Atk:     1200,
		Def:     1500,
		IsValid: true,
	})

	/* 138 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 185
		 调整:

		 中文名: 诅咒之龙
		 卡片种类: 通常怪兽
		 卡片密码: 28279543
		 种族: 龙
		 属性: 暗
		 星级: 5
		 攻击力: 2000
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: PG、BE01、EX-R(EX)、VOL02、DL02
		 效果: 描述：邪恶的龙。使用暗之力的攻击很强力。

		*/
		Id:       185,
		Password: "28279543",
		Name:     "诅咒之龙",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Dark,   // 暗
		Lr:      ygo.LR_Dragon, // 龙
		Atk:     2000,
		Def:     1500,
		IsValid: true,
	})

	/* 139 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 186
		 调整:

		 中文名: 岩石巨兵
		 卡片种类: 通常怪兽
		 卡片密码: 13039848
		 种族: 岩石
		 属性: 地
		 星级: 3
		 攻击力: 1300
		 防御力: 2000
		 罕见度: 平卡N
		 卡包: PG、BE01、EX-R(EX)、YSD02、VOL03、DL02、YU、SY2
		 效果: 描述：岩石的巨人兵。粗壮手腕的攻击让大地动摇。

		*/
		Id:       186,
		Password: "13039848",
		Name:     "岩石巨兵",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Rock,  // 岩石
		Atk:     1300,
		Def:     2000,
		IsValid: true,
	})

	/* 140 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 187
		 调整:

		 中文名: 荒野盗龙
		 卡片种类: 通常怪兽
		 卡片密码: 01784619
		 种族: 恐龙
		 属性: 地
		 星级: 4
		 攻击力: 1500
		 防御力: 800
		 罕见度: 平卡N
		 卡包: PG、BE01、EX-R(EX)、VOL02、DL02
		 效果: 描述：擅长奔跑的恐龙，利用尖锐的爪子进行攻击。

		*/
		Id:       187,
		Password: "01784619",
		Name:     "荒野盗龙",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth,    // 地
		Lr:      ygo.LR_Dinosaur, // 恐龙
		Atk:     1500,
		Def:     800,
		IsValid: true,
	})

	/* 141 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 189
		 调整:

		 中文名: 魔人 死亡撒旦
		 卡片种类: 通常怪兽
		 卡片密码: 36304921
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 1400
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: PG、EX-R(EX)、VOL03、DL02
		 效果: 描述：穿着融入黑暗的黑色燕尾服，操纵着死亡的恶魔。

		*/
		Id:       189,
		Password: "36304921",
		Name:     "魔人 死亡撒旦",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     1400,
		Def:     1300,
		IsValid: true,
	})

	/* 142 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 190
		 调整:

		 中文名: 竖琴精灵
		 卡片种类: 通常怪兽
		 卡片密码: 80770678
		 种族: 天使
		 属性: 光
		 星级: 4
		 攻击力: 800
		 防御力: 2000
		 罕见度: 平卡N
		 卡包: PG、VOL02、DL02、YSD06、ST12
		 效果: 描述：在天界弹奏竖琴的精灵。那音色使周围听众的心变得平静。

		*/
		Id:       190,
		Password: "80770678",
		Name:     "竖琴精灵",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Light, // 光
		Lr:      ygo.LR_Fairy, // 天使
		Atk:     800,
		Def:     2000,
		IsValid: true,
	})

	/* 143 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 191
		 调整:

		 中文名: 魔人 泰拉
		 卡片种类: 通常怪兽
		 卡片密码: 63308047
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 1200
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: PG、EX-R(EX)、VOL02、DL02
		 效果: 描述：住在沼泽地的恶魔手下，看起来不怎么强，但绝不能对他轻心大意。

		*/
		Id:       191,
		Password: "63308047",
		Name:     "魔人 泰拉",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     1200,
		Def:     1300,
		IsValid: true,
	})

	/* 144 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 192
		 调整:

		 中文名: 恍惚的人鱼
		 卡片种类: 通常怪兽
		 卡片密码: 75376965
		 种族: 鱼
		 属性: 水
		 星级: 3
		 攻击力: 1200
		 防御力: 900
		 罕见度: 平卡N
		 卡包: PG、VOL02、DL02、TP15
		 效果: 描述：诱惑海上航海者使之溺水的美丽人鱼。

		*/
		Id:       192,
		Password: "75376965",
		Name:     "恍惚的人鱼",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Fish,  // 鱼
		Atk:     1200,
		Def:     900,
		IsValid: true,
	})

	/* 145 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 193
		 调整:

		 中文名: 炎之魔神
		 卡片种类: 通常怪兽
		 卡片密码: 71407486
		 种族: 炎
		 属性: 炎
		 星级: 4
		 攻击力: 1300
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: PG、VOL03、DL02
		 效果: 描述：被火焰包围的魔人，能自由操纵火焰。

		*/
		Id:       193,
		Password: "71407486",
		Name:     "炎之魔神",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Fire, // 炎
		Lr:      ygo.LR_Fire, // 炎
		Atk:     1300,
		Def:     1000,
		IsValid: true,
	})

	/* 146 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 209
		 调整:

		 中文名: 尖刺海龙
		 卡片种类: 通常怪兽
		 卡片密码: 85326399
		 种族: 海龙
		 属性: 水
		 星级: 5
		 攻击力: 1600
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: PG、VOL03、DL02
		 效果: 描述：用身体的刺攻击敌人，并能放出电流。

		*/
		Id:       209,
		Password: "85326399",
		Name:     "尖刺海龙",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Water,      // 水
		Lr:      ygo.LR_SeaSerpent, // 海龙
		Atk:     1600,
		Def:     1300,
		IsValid: true,
	})

	/* 147 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 210
		 调整:

		 中文名: 天空猎手
		 卡片种类: 通常怪兽
		 卡片密码: 10202894
		 种族: 鸟兽
		 属性: 风
		 星级: 4
		 攻击力: 1550
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: PG、EX-R(EX)、VOL03、DL02
		 效果: 描述：在羽毛里藏着小刀，从空中降下发起攻击。

		*/
		Id:       210,
		Password: "10202894",
		Name:     "天空猎手",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Wind,        // 风
		Lr:      ygo.LR_WingedBeast, // 鸟兽
		Atk:     1550,
		Def:     1200,
		IsValid: true,
	})

	/* 148 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 213
		 调整:

		 中文名: 沙石怪
		 卡片种类: 通常怪兽
		 卡片密码: 73051941
		 种族: 岩石
		 属性: 地
		 星级: 5
		 攻击力: 1300
		 防御力: 1600
		 罕见度: 平卡N
		 卡包: PG、VOL03、DL02
		 效果: 描述：从地下突然出现在眼前，用触手做攻击。

		*/
		Id:       213,
		Password: "73051941",
		Name:     "沙石怪",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Rock,  // 岩石
		Atk:     1300,
		Def:     1600,
		IsValid: true,
	})

	/* 149 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 215
		 调整:

		 中文名: 钢铁巨神像
		 卡片种类: 通常怪兽
		 卡片密码: 29172562
		 种族: 机械
		 属性: 地
		 星级: 5
		 攻击力: 1400
		 防御力: 1800
		 罕见度: 平卡N
		 卡包: PG、VOL03、DL02
		 效果: 描述：据域是被机械帝国所祭奠着的钢铁的巨神像。

		*/
		Id:       215,
		Password: "29172562",
		Name:     "钢铁巨神像",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Machine, // 机械
		Atk:     1400,
		Def:     1800,
		IsValid: true,
	})

	/* 150 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 216
		 调整:

		 中文名: 下级龙
		 卡片种类: 通常怪兽
		 卡片密码: 55444629
		 种族: 龙
		 属性: 风
		 星级: 4
		 攻击力: 1200
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: PG、VOL03、DL02
		 效果: 描述：不怎么强，连吐息攻击都不会的低级龙。

		*/
		Id:       216,
		Password: "55444629",
		Name:     "下级龙",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Wind,   // 风
		Lr:      ygo.LR_Dragon, // 龙
		Atk:     1200,
		Def:     1000,
		IsValid: true,
	})

	/* 151 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 217
		 调整:

		 中文名: 女夜魔骑士
		 卡片种类: 通常怪兽
		 卡片密码: 55291359
		 种族: 战士
		 属性: 暗
		 星级: 5
		 攻击力: 1650
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: PG、VOL03、DL02
		 效果: 描述：咏唱起魔法，给予对方血的祭礼的恶魔魔法战士。

		*/
		Id:       217,
		Password: "55291359",
		Name:     "女夜魔骑士",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Dark,    // 暗
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     1650,
		Def:     1300,
		IsValid: true,
	})

	/* 152 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 222
		 调整:

		 中文名: 被封印者的右足
		 卡片种类: 通常怪兽
		 卡片密码: 08124921
		 种族: 魔法师
		 属性: 暗
		 星级: 1
		 攻击力: 200
		 防御力: 300
		 罕见度: 平卡N，银字R
		 卡包: PG、BE01、VOL04、DL02
		 效果: 描述：被封印者的右足。封印解开后将得到无限的力量。

		*/
		Id:       222,
		Password: "08124921",
		Name:     "被封印者的右足",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   1,
		La:      ygo.LA_Dark,        // 暗
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     200,
		Def:     300,
		IsValid: true,
	})

	/* 153 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 223
		 调整:

		 中文名: 被封印者的左足
		 卡片种类: 通常怪兽
		 卡片密码: 44519536
		 种族: 魔法师
		 属性: 暗
		 星级: 1
		 攻击力: 200
		 防御力: 300
		 罕见度: 银字R
		 卡包: PG、BE01、VOL03、DL02
		 效果: 描述：被封印者的左足。封印解开后将得到无限的力量。

		*/
		Id:       223,
		Password: "44519536",
		Name:     "被封印者的左足",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   1,
		La:      ygo.LA_Dark,        // 暗
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     200,
		Def:     300,
		IsValid: true,
	})

	/* 154 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2257
		 调整:

		 中文名: 小天使
		 卡片种类: 通常怪兽
		 卡片密码: 38142739
		 种族: 天使
		 属性: 光
		 星级: 3
		 攻击力: 600
		 防御力: 900
		 罕见度: 平卡N
		 卡包: LB、VOL01、TP15
		 效果: 描述：依靠灵活的行动躲避攻击的小天使。

		*/
		Id:       2257,
		Password: "38142739",
		Name:     "小天使",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Light, // 光
		Lr:      ygo.LR_Fairy, // 天使
		Atk:     600,
		Def:     900,
		IsValid: true,
	})

	/* 155 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2258
		 调整:

		 中文名: 暗黑灰羚
		 卡片种类: 通常怪兽
		 卡片密码: 09159938
		 种族: 兽
		 属性: 地
		 星级: 3
		 攻击力: 800
		 防御力: 900
		 罕见度: 平卡N
		 卡包: LB、VOL01
		 效果: 描述：全身灰色的野兽，是非常少见的珍贵品种。

		*/
		Id:       2258,
		Password: "09159938",
		Name:     "暗黑灰羚",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Atk:     800,
		Def:     900,
		IsValid: true,
	})

	/* 156 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2259
		 调整:

		 中文名: 紫炎的影武者
		 卡片种类: 通常怪兽
		 卡片密码: 15401633
		 种族: 战士
		 属性: 地
		 星级: 2
		 攻击力: 800
		 防御力: 400
		 罕见度: 平卡N
		 卡包: LB、VOL01、TP15
		 效果: 描述：侍奉着紫炎的影武者，持有锋利的名刀。

		*/
		Id:       2259,
		Password: "15401633",
		Name:     "紫炎的影武者",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     800,
		Def:     400,
		IsValid: true,
	})

	/* 157 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2260
		 调整:

		 中文名: 鸭人
		 卡片种类: 通常怪兽
		 卡片密码: 85705804
		 种族: 鸟兽
		 属性: 风
		 星级: 3
		 攻击力: 800
		 防御力: 800
		 罕见度: 平卡N
		 卡包: LB、VOL01、TP14
		 效果: 描述：尾巴很长的鸟，从空中发起攻击。

		*/
		Id:       2260,
		Password: "85705804",
		Name:     "鸭人",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Wind,        // 风
		Lr:      ygo.LR_WingedBeast, // 鸟兽
		Atk:     800,
		Def:     800,
		IsValid: true,
	})

	/* 158 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2261
		 调整:

		 中文名: 睡眠之子
		 卡片种类: 通常怪兽
		 卡片密码: 90963488
		 种族: 魔法师
		 属性: 暗
		 星级: 3
		 攻击力: 800
		 防御力: 700
		 罕见度: 平卡N
		 卡包: LB、VOL01
		 效果: 描述：孩子被梦魔所操控永远沉睡着不再醒来。

		*/
		Id:       2261,
		Password: "90963488",
		Name:     "睡眠之子",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Dark,        // 暗
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     800,
		Def:     700,
		IsValid: true,
	})

	/* 159 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2262
		 调整:

		 中文名: 死亡之足
		 卡片种类: 通常怪兽
		 卡片密码: 08944575
		 种族: 恶魔
		 属性: 暗
		 星级: 3
		 攻击力: 700
		 防御力: 800
		 罕见度: 平卡N
		 卡包: VOL01
		 效果: 描述：眼珠上长着脚的怪兽。高跳着用爪钩攻击。

		*/
		Id:       2262,
		Password: "08944575",
		Name:     "死亡之足",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     700,
		Def:     800,
		IsValid: true,
	})

	/* 160 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2263
		 调整:

		 中文名: 愤怒的海王
		 卡片种类: 通常怪兽
		 卡片密码: 18710707
		 种族: 水
		 属性: 水
		 星级: 3
		 攻击力: 800
		 防御力: 700
		 罕见度: 平卡N
		 卡包: LB、VOL01
		 效果: 描述：伟大的海之王。唤起永不停息的大海啸将敌人吞没。

		*/
		Id:       2263,
		Password: "18710707",
		Name:     "愤怒的海王",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     800,
		Def:     700,
		IsValid: true,
	})

	/* 161 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2264
		 调整:

		 中文名: 生命之沙漏
		 卡片种类: 通常怪兽
		 卡片密码: 08783685
		 种族: 天使
		 属性: 光
		 星级: 2
		 攻击力: 700
		 防御力: 600
		 罕见度: 平卡N
		 卡包: VOL01、TP16
		 效果: 描述：掌管生命的天使。通过削减自己的寿命将力量赋予他人。

		*/
		Id:       2264,
		Password: "08783685",
		Name:     "生命之沙漏",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Light, // 光
		Lr:      ygo.LR_Fairy, // 天使
		Atk:     700,
		Def:     600,
		IsValid: true,
	})

	/* 162 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2265
		 调整:

		 中文名: 恶魔之镜
		 卡片种类: 通常怪兽
		 卡片密码: 15150371
		 种族: 恶魔
		 属性: 暗
		 星级: 2
		 攻击力: 700
		 防御力: 600
		 罕见度: 平卡N
		 卡包: VOL01
		 效果: 描述：对映在镜子上的东西施以催眠术，从而避开攻击的恶魔之镜。

		*/
		Id:       2265,
		Password: "15150371",
		Name:     "恶魔之镜",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     700,
		Def:     600,
		IsValid: true,
	})

	/* 163 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2266
		 调整:

		 中文名: 火焰恶魔
		 卡片种类: 通常怪兽
		 卡片密码: 53581214
		 种族: 不死
		 属性: 暗
		 星级: 2
		 攻击力: 700
		 防御力: 500
		 罕见度: 平卡N
		 卡包: VOL01
		 效果: 描述：手持炎之箭的死神，被其箭击中的话就会变成火人。

		*/
		Id:       2266,
		Password: "53581214",
		Name:     "火焰恶魔",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Dark,   // 暗
		Lr:      ygo.LR_Zombie, // 不死
		Atk:     700,
		Def:     500,
		IsValid: true,
	})

	/* 164 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2267
		 调整:

		 中文名: 小龙
		 卡片种类: 通常怪兽
		 卡片密码: 75356564
		 种族: 龙
		 属性: 风
		 星级: 2
		 攻击力: 600
		 防御力: 700
		 罕见度: 平卡N
		 卡包: LB、VOL01、TP01
		 效果: 描述：体型细小的龙。使上整个细小身体来进行攻击。

		*/
		Id:       2267,
		Password: "75356564",
		Name:     "小龙",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Wind,   // 风
		Lr:      ygo.LR_Dragon, // 龙
		Atk:     600,
		Def:     700,
		IsValid: true,
	})

	/* 165 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2268
		 调整:

		 中文名: 巴比伦
		 卡片种类: 通常怪兽
		 卡片密码: 53832650
		 种族: 兽
		 属性: 地
		 星级: 2
		 攻击力: 700
		 防御力: 600
		 罕见度: 平卡N
		 卡包: VOL01
		 效果: 描述：长有独眼的巨大怪兽。从眼睛之中发出光线攻击。

		*/
		Id:       2268,
		Password: "53832650",
		Name:     "巴比伦",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Atk:     700,
		Def:     600,
		IsValid: true,
	})

	/* 166 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2269
		 调整:

		 中文名: 火炎草
		 卡片种类: 通常怪兽
		 卡片密码: 53293545
		 种族: 植物
		 属性: 地
		 星级: 2
		 攻击力: 700
		 防御力: 600
		 罕见度: 平卡N
		 卡包: LB、VOL01
		 效果: 描述：在火山附近生存的草。从花内吹出来火焰攻击。

		*/
		Id:       2269,
		Password: "53293545",
		Name:     "火炎草",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Plant, // 植物
		Atk:     700,
		Def:     600,
		IsValid: true,
	})

	/* 167 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 227
		 调整:

		 中文名: 小精怪
		 卡片种类: 通常怪兽
		 卡片密码: 41392891
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 1300
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: RB、BE01、EX-R(EX)、VOL05、DL02、SY2
		 效果: 描述：喜欢恶作剧的小恶魔，从黑暗中袭来。警惕！

		*/
		Id:       227,
		Password: "41392891",
		Name:     "小精怪",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     1300,
		Def:     1400,
		IsValid: true,
	})

	/* 168 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2270
		 调整:

		 中文名: 暗黑杀手
		 卡片种类: 通常怪兽
		 卡片密码: 39175982
		 种族: 昆虫
		 属性: 地
		 星级: 2
		 攻击力: 700
		 防御力: 700
		 罕见度: 平卡N
		 卡包: VOL01
		 效果: 描述：挥动如镰刀般发达的手臂进行攻击。

		*/
		Id:       2270,
		Password: "39175982",
		Name:     "暗黑杀手",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Earth,  // 地
		Lr:      ygo.LR_Insect, // 昆虫
		Atk:     700,
		Def:     700,
		IsValid: true,
	})

	/* 169 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2271
		 调整:

		 中文名: 黑暗随从者
		 卡片种类: 通常怪兽
		 卡片密码: 15507080
		 种族: 魔法师
		 属性: 暗
		 星级: 2
		 攻击力: 700
		 防御力: 500
		 罕见度: 平卡N
		 卡包: VOL01
		 效果: 描述：崇拜黑暗的魔法师。呼唤魔手将敌人拖入暗阴之中。

		*/
		Id:       2271,
		Password: "15507080",
		Name:     "黑暗随从者",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Dark,        // 暗
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     700,
		Def:     500,
		IsValid: true,
	})

	/* 170 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2272
		 调整:

		 中文名: 海星葵
		 卡片种类: 通常怪兽
		 卡片密码: 46718686
		 种族: 水
		 属性: 水
		 星级: 2
		 攻击力: 600
		 防御力: 700
		 罕见度: 平卡N
		 卡包: VOL01
		 效果: 描述：被污染的水源侵蚀而狂暴化的海星。从口中吐出溶解液攻击敌人。

		*/
		Id:       2272,
		Password: "46718686",
		Name:     "海星葵",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     600,
		Def:     700,
		IsValid: true,
	})

	/* 171 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2273
		 调整:

		 中文名: 雷电小子
		 卡片种类: 通常怪兽
		 卡片密码: 15510988
		 种族: 雷
		 属性: 风
		 星级: 2
		 攻击力: 700
		 防御力: 600
		 罕见度: 平卡N
		 卡包: VOL01
		 效果: 描述：在它的体内平时储蓄着雷电。因此当它哭泣时非常危险。

		*/
		Id:       2273,
		Password: "15510988",
		Name:     "雷电小子",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Wind,    // 风
		Lr:      ygo.LR_Thunder, // 雷
		Atk:     700,
		Def:     600,
		IsValid: true,
	})

	/* 172 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2274
		 调整:

		 中文名: 复印者
		 卡片种类: 通常怪兽
		 卡片密码: 64511793
		 种族: 战士
		 属性: 地
		 星级: 2
		 攻击力: 600
		 防御力: 500
		 罕见度: 平卡N
		 卡包: VOL01
		 效果: 描述：变身成为各种敌人的模样，从而欺骗对方进行战斗的战士。

		*/
		Id:       2274,
		Password: "64511793",
		Name:     "复印者",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     600,
		Def:     500,
		IsValid: true,
	})

	/* 173 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2275
		 调整:

		 中文名: 雷云怪
		 卡片种类: 通常怪兽
		 卡片密码: 09430387
		 种族: 雷
		 属性: 风
		 星级: 2
		 攻击力: 600
		 防御力: 600
		 罕见度: 平卡N
		 卡包: VOL01
		 效果: 描述：带电的云形怪兽。下着能溶掉一切的危险之雨。

		*/
		Id:       2275,
		Password: "09430387",
		Name:     "雷云怪",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Wind,    // 风
		Lr:      ygo.LR_Thunder, // 雷
		Atk:     600,
		Def:     600,
		IsValid: true,
	})

	/* 174 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2276
		 调整:

		 中文名: 命运之蜡烛
		 卡片种类: 通常怪兽
		 卡片密码: 47695416
		 种族: 恶魔
		 属性: 暗
		 星级: 2
		 攻击力: 600
		 防御力: 600
		 罕见度: 平卡N
		 卡包: VOL01
		 效果: 描述：指尖的火焰消失时，将决定对方的命运。

		*/
		Id:       2276,
		Password: "47695416",
		Name:     "命运之蜡烛",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     600,
		Def:     600,
		IsValid: true,
	})

	/* 175 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2277
		 调整:

		 中文名: 黑魔族的窗帘
		 卡片种类: 通常怪兽
		 卡片密码: 22026707
		 种族: 魔法师
		 属性: 暗
		 星级: 2
		 攻击力: 600
		 防御力: 500
		 罕见度: 平卡N
		 卡包: VOL01
		 效果: 描述：魔术师制作的窗帘。据说可以提升魔法师的力量。

		*/
		Id:       2277,
		Password: "22026707",
		Name:     "黑魔族的窗帘",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Dark,        // 暗
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     600,
		Def:     500,
		IsValid: true,
	})

	/* 176 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2278
		 调整:

		 中文名: 死者之腕
		 卡片种类: 通常怪兽
		 卡片密码: 52800428
		 种族: 不死
		 属性: 暗
		 星级: 2
		 攻击力: 600
		 防御力: 600
		 罕见度: 平卡N
		 卡包: VOL01
		 效果: 描述：从混沌沼泽中伸出手腕，将生者拖入其中。

		*/
		Id:       2278,
		Password: "52800428",
		Name:     "死者之腕",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Dark,   // 暗
		Lr:      ygo.LR_Zombie, // 不死
		Atk:     600,
		Def:     600,
		IsValid: true,
	})

	/* 177 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2279
		 调整:

		 中文名: 恶魔海狸
		 卡片种类: 通常怪兽
		 卡片密码: 75889523
		 种族: 兽
		 属性: 地
		 星级: 2
		 攻击力: 400
		 防御力: 600
		 罕见度: 平卡N
		 卡包: VOL01、TP06
		 效果: 描述：拥有恶魔角与翅膀的海狸。投掷橡实果进行攻击。

		*/
		Id:       2279,
		Password: "75889523",
		Name:     "恶魔海狸",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Atk:     400,
		Def:     600,
		IsValid: true,
	})

	/* 178 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 228
		 调整:

		 中文名: 守城的翼龙
		 卡片种类: 通常怪兽
		 卡片密码: 87796900
		 种族: 龙
		 属性: 风
		 星级: 4
		 攻击力: 1400
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: RB、BE01、EX-R(EX)、VOL04、DL02、SY2
		 效果: 描述：守护山寨的龙，从空中急速下降攻击敌人。

		*/
		Id:       228,
		Password: "87796900",
		Name:     "守城的翼龙",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Wind,   // 风
		Lr:      ygo.LR_Dragon, // 龙
		Atk:     1400,
		Def:     1200,
		IsValid: true,
	})

	/* 179 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2280
		 调整:

		 中文名: 小恐龙
		 卡片种类: 通常怪兽
		 卡片密码: 46457856
		 种族: 恐龙
		 属性: 地
		 星级: 2
		 攻击力: 500
		 防御力: 400
		 罕见度: 平卡N
		 卡包: VOL01
		 效果: 描述：体形虽然小，但性格凶暴。经常和同伴相争。

		*/
		Id:       2280,
		Password: "46457856",
		Name:     "小恐龙",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Earth,    // 地
		Lr:      ygo.LR_Dinosaur, // 恐龙
		Atk:     500,
		Def:     400,
		IsValid: true,
	})

	/* 180 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2281
		 调整:

		 中文名: 埴轮
		 卡片种类: 通常怪兽
		 卡片密码: 84285623
		 种族: 岩石
		 属性: 地
		 星级: 2
		 攻击力: 500
		 防御力: 500
		 罕见度: 平卡N
		 卡包: VOL01
		 效果: 描述：古代王墓中守护宝物的土人偶。

		*/
		Id:       2281,
		Password: "84285623",
		Name:     "埴轮",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Rock,  // 岩石
		Atk:     500,
		Def:     500,
		IsValid: true,
	})

	/* 181 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2283
		 调整:

		 中文名: 蜘蛛男
		 卡片种类: 通常怪兽
		 卡片密码: 56283725
		 种族: 昆虫
		 属性: 地
		 星级: 3
		 攻击力: 700
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: PG、VOL02
		 效果: 描述：巨大的蜘蛛被赋予智慧的形态，通过吐丝将敌人的行动封锁。

		*/
		Id:       2283,
		Password: "56283725",
		Name:     "蜘蛛男",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,  // 地
		Lr:      ygo.LR_Insect, // 昆虫
		Atk:     700,
		Def:     1400,
		IsValid: true,
	})

	/* 182 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2284
		 调整:

		 中文名: 铠甲剑尾战士
		 卡片种类: 通常怪兽
		 卡片密码: 53153481
		 种族: 战士
		 属性: 地
		 星级: 3
		 攻击力: 700
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: PG、VOL02、TP09
		 效果: 描述：有着剑状尾的怪异战士。用双手和尾巴进行三连攻击。

		*/
		Id:       2284,
		Password: "53153481",
		Name:     "铠甲剑尾战士",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     700,
		Def:     1300,
		IsValid: true,
	})

	/* 183 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2286
		 调整:

		 中文名: 独眼盾龙
		 卡片种类: 通常怪兽
		 卡片密码: 33064647
		 种族: 龙
		 属性: 风
		 星级: 3
		 攻击力: 700
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: PG、VOL02
		 效果: 描述：身上的盾不仅能保护身体，还能用来突击。

		*/
		Id:       2286,
		Password: "33064647",
		Name:     "独眼盾龙",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Wind,   // 风
		Lr:      ygo.LR_Dragon, // 龙
		Atk:     700,
		Def:     1300,
		IsValid: true,
	})

	/* 184 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2288
		 调整:

		 中文名: 岩石犰狳
		 卡片种类: 通常怪兽
		 卡片密码: 63432835
		 种族: 岩石
		 属性: 地
		 星级: 3
		 攻击力: 800
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: VOL02
		 效果: 描述：身体被像石头般坚固的毛覆盖着、守备坚固。

		*/
		Id:       2288,
		Password: "63432835",
		Name:     "岩石犰狳",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Rock,  // 岩石
		Atk:     800,
		Def:     1200,
		IsValid: true,
	})

	/* 185 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2289
		 调整:

		 中文名: 坚硬铠甲
		 卡片种类: 通常怪兽
		 卡片密码: 20060230
		 种族: 战士
		 属性: 地
		 星级: 3
		 攻击力: 300
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: PG、VOL02
		 效果: 描述：拥有心灵的铠甲，坚硬的身体装备在战士身上。

		*/
		Id:       2289,
		Password: "20060230",
		Name:     "坚硬铠甲",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     300,
		Def:     1200,
		IsValid: true,
	})

	/* 186 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 229
		 调整:

		 中文名: 恶魔召唤
		 卡片种类: 通常怪兽
		 卡片密码: 70781052
		 种族: 恶魔
		 属性: 暗
		 星级: 6
		 攻击力: 2500
		 防御力: 1200
		 罕见度: 平卡N，银字R，面闪SR，金字UR，爆闪PR，立体UTR
		 卡包: RB、BE01、LE03、SC、EX-R(EX)、VOL04、DL02、Booster R3、YAP01、DPYG、DT09
		 效果: 描述：使用黑暗力量，迷惑人心的恶魔。在恶魔族中以相当强大的力量著称。

		*/
		Id:       229,
		Password: "70781052",
		Name:     "恶魔召唤",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   6,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     2500,
		Def:     1200,
		IsValid: true,
	})

	/* 187 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2292
		 调整:

		 中文名: 全息投影者
		 卡片种类: 通常怪兽
		 卡片密码: 10859908
		 种族: 机械
		 属性: 地
		 星级: 3
		 攻击力: 1100
		 防御力: 700
		 罕见度: 平卡N
		 卡包: VOL02
		 效果: 描述：使人看见各种幻想、趁机进行攻击的机器。

		*/
		Id:       2292,
		Password: "10859908",
		Name:     "全息投影者",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Machine, // 机械
		Atk:     1100,
		Def:     700,
		IsValid: true,
	})

	/* 188 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2293
		 调整:

		 中文名: 孤寂
		 卡片种类: 通常怪兽
		 卡片密码: 84794011
		 种族: 兽战士
		 属性: 地
		 星级: 3
		 攻击力: 1050
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: VOL02
		 效果: 描述：下半身是鹿，持有吸收灵魂的大镰刀的兽 战士。

		*/
		Id:       2293,
		Password: "84794011",
		Name:     "孤寂",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,        // 地
		Lr:      ygo.LR_BeastWarrior, // 兽战士
		Atk:     1050,
		Def:     1000,
		IsValid: true,
	})

	/* 189 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2294
		 调整:

		 中文名: 拉瓦斯
		 卡片种类: 通常怪兽
		 卡片密码: 94675535
		 种族: 兽
		 属性: 地
		 星级: 3
		 攻击力: 800
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: PG、VOL02
		 效果: 描述：行动速度飞快的鸟怪，用细长的手腕将敌人绞杀。

		*/
		Id:       2294,
		Password: "94675535",
		Name:     "拉瓦斯",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Atk:     800,
		Def:     1000,
		IsValid: true,
	})

	/* 190 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2295
		 调整:

		 中文名: 利爪杀手
		 卡片种类: 通常怪兽
		 卡片密码: 41218256
		 种族: 恶魔
		 属性: 暗
		 星级: 3
		 攻击力: 1000
		 防御力: 800
		 罕见度: 平卡N
		 卡包: EX-R(EX)、VOL02
		 效果: 描述：手腕能自由的伸缩，尖利的爪子能将敌人刺穿。

		*/
		Id:       2295,
		Password: "41218256",
		Name:     "利爪杀手",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     1000,
		Def:     800,
		IsValid: true,
	})

	/* 191 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2296
		 调整:

		 中文名: 森之尸
		 卡片种类: 通常怪兽
		 卡片密码: 17733394
		 种族: 不死
		 属性: 暗
		 星级: 3
		 攻击力: 1000
		 防御力: 900
		 罕见度: 平卡N
		 卡包: VOL02
		 效果: 描述：森林守护神死后，被恶魔复活的尸体。

		*/
		Id:       2296,
		Password: "17733394",
		Name:     "森之尸",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Dark,   // 暗
		Lr:      ygo.LR_Zombie, // 不死
		Atk:     1000,
		Def:     900,
		IsValid: true,
	})

	/* 192 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2297
		 调整:

		 中文名: 阴暗处的协力者
		 卡片种类: 通常怪兽
		 卡片密码: 41422426
		 种族: 战士
		 属性: 地
		 星级: 3
		 攻击力: 1000
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: VOL02
		 效果: 描述：从暗处协助他人的可爱小人。

		*/
		Id:       2297,
		Password: "41422426",
		Name:     "阴暗处的协力者",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     1000,
		Def:     1000,
		IsValid: true,
	})

	/* 193 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2298
		 调整:

		 中文名: 磁力战士1号
		 卡片种类: 通常怪兽
		 卡片密码: 56342351
		 种族: 战士
		 属性: 地
		 星级: 3
		 攻击力: 1000
		 防御力: 500
		 罕见度: 平卡N
		 卡包: PG、VOL02
		 效果: 描述：擅长配合的战士，发动强大的磁力，谁也逃不掉。

		*/
		Id:       2298,
		Password: "56342351",
		Name:     "磁力战士1号",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     1000,
		Def:     500,
		IsValid: true,
	})

	/* 194 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2299
		 调整:

		 中文名: 磁力战士2号
		 卡片种类: 通常怪兽
		 卡片密码: 92731455
		 种族: 战士
		 属性: 地
		 星级: 3
		 攻击力: 500
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: PG、VOL02
		 效果: 描述：擅长配合的战士，表面涂有电磁的铠甲非常坚固。

		*/
		Id:       2299,
		Password: "92731455",
		Name:     "磁力战士2号",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     500,
		Def:     1000,
		IsValid: true,
	})

	/* 195 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 230
		 调整:

		 中文名: 岩窟魔人
		 卡片种类: 通常怪兽
		 卡片密码: 68846917
		 种族: 岩石
		 属性: 地
		 星级: 3
		 攻击力: 800
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: RB、VOL04、DL02、JY
		 效果: 描述：岩石之躯成就高强的守备力。挥起巨腕时务必警惕。

		*/
		Id:       230,
		Password: "68846917",
		Name:     "岩窟魔人",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Rock,  // 岩石
		Atk:     800,
		Def:     1200,
		IsValid: true,
	})

	/* 196 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2300
		 调整:

		 中文名: 泥浆兽
		 卡片种类: 通常怪兽
		 卡片密码: 24194033
		 种族: 水
		 属性: 水
		 星级: 3
		 攻击力: 900
		 防御力: 800
		 罕见度: 平卡N
		 卡包: VOL02
		 效果: 描述：粘呼呼的令人感到恶心的怪兽。吐出剧毒的瓦斯进行攻击。

		*/
		Id:       2300,
		Password: "24194033",
		Name:     "泥浆兽",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     900,
		Def:     800,
		IsValid: true,
	})

	/* 197 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2301
		 调整:

		 中文名: 青眼银僵尸
		 卡片种类: 通常怪兽
		 卡片密码: 35282433
		 种族: 不死
		 属性: 暗
		 星级: 3
		 攻击力: 900
		 防御力: 700
		 罕见度: 平卡N
		 卡包: VOL02
		 效果: 描述：据说可以从眼睛里射出怪异的光线，据说能让对方变成僵尸。

		*/
		Id:       2301,
		Password: "35282433",
		Name:     "青眼银僵尸",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Dark,   // 暗
		Lr:      ygo.LR_Zombie, // 不死
		Atk:     900,
		Def:     700,
		IsValid: true,
	})

	/* 198 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2302
		 调整:

		 中文名: 食人花
		 卡片种类: 通常怪兽
		 卡片密码: 93553943
		 种族: 植物
		 属性: 地
		 星级: 2
		 攻击力: 800
		 防御力: 600
		 罕见度: 平卡N
		 卡包: PG、VOL02
		 效果: 描述：吃人的人面花，用有毒的触手进行攻击。

		*/
		Id:       2302,
		Password: "93553943",
		Name:     "食人花",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Plant, // 植物
		Atk:     800,
		Def:     600,
		IsValid: true,
	})

	/* 199 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2303
		 调整:

		 中文名: 暗黑拿破仑
		 卡片种类: 通常怪兽
		 卡片密码: 76211194
		 种族: 恶魔
		 属性: 暗
		 星级: 2
		 攻击力: 800
		 防御力: 400
		 罕见度: 平卡N
		 卡包: PG、VOL02、TP05
		 效果: 描述：恶人创造的眼珠恶魔。用黑暗炸弹进行爆破攻击。

		*/
		Id:       2303,
		Password: "76211194",
		Name:     "暗黑拿破仑",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     800,
		Def:     400,
		IsValid: true,
	})

	/* 200 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2304
		 调整:

		 中文名: 魅惑的怪盗
		 卡片种类: 通常怪兽
		 卡片密码: 24348204
		 种族: 魔法师
		 属性: 暗
		 星级: 2
		 攻击力: 700
		 防御力: 700
		 罕见度: 平卡N
		 卡包: VOL02
		 效果: 描述：披着黑色披风装模作样的怪盗。挥着手杖迷惑对方。

		*/
		Id:       2304,
		Password: "24348204",
		Name:     "魅惑的怪盗",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Dark,        // 暗
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     700,
		Def:     700,
		IsValid: true,
	})

	/* 201 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2305
		 调整:

		 中文名: 暗杀者
		 卡片种类: 通常怪兽
		 卡片密码: 48365709
		 种族: 战士
		 属性: 地
		 星级: 5
		 攻击力: 1700
		 防御力: 1200
		 罕见度: 平卡N，银字R
		 卡包: EX-R(EX)、VOL03、Booster02
		 效果: 描述：能在黑暗中悄然无声地靠近敌人，精通暗杀的战士。

		*/
		Id:       2305,
		Password: "48365709",
		Name:     "暗杀者",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     1700,
		Def:     1200,
		IsValid: true,
	})

	/* 202 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2306
		 调整:

		 中文名: 卡库塔斯
		 卡片种类: 通常怪兽
		 卡片密码: 36904469
		 种族: 水
		 属性: 水
		 星级: 5
		 攻击力: 1700
		 防御力: 1400
		 罕见度: 银字R
		 卡包: VOL03、Booster02
		 效果: 描述：潜在水中形状不明的怪兽。

		*/
		Id:       2306,
		Password: "36904469",
		Name:     "卡库塔斯",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     1700,
		Def:     1400,
		IsValid: true,
	})

	/* 203 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2307
		 调整:

		 中文名: 机械巨兵
		 卡片种类: 通常怪兽
		 卡片密码: 72299832
		 种族: 机械
		 属性: 地
		 星级: 6
		 攻击力: 1750
		 防御力: 1900
		 罕见度: 银字R
		 卡包: VOL03、Booster02
		 效果: 描述：巨斧的一击可以割开大地。

		*/
		Id:       2307,
		Password: "72299832",
		Name:     "机械巨兵",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   6,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Machine, // 机械
		Atk:     1750,
		Def:     1900,
		IsValid: true,
	})

	/* 204 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2308
		 调整:

		 中文名: 神圣人偶
		 卡片种类: 通常怪兽
		 卡片密码: 91939608
		 种族: 魔法师
		 属性: 光
		 星级: 4
		 攻击力: 1600
		 防御力: 1000
		 罕见度: 银字R
		 卡包: EX-R(EX)、VOL03、Booster02、PE
		 效果: 描述：操纵神圣力量的人偶。在黑暗之中攻击力很强。

		*/
		Id:       2308,
		Password: "91939608",
		Name:     "神圣人偶",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Light,       // 光
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     1600,
		Def:     1000,
		IsValid: true,
	})

	/* 205 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2309
		 调整:

		 中文名: 魔加农
		 卡片种类: 通常怪兽
		 卡片密码: 98795934
		 种族: 恶魔
		 属性: 暗
		 星级: 5
		 攻击力: 1700
		 防御力: 1400
		 罕见度: 平卡N，银字R
		 卡包: VOL03、Booster02
		 效果: 描述：大炮状的恶魔。以飞快的速度发射眼球弹。

		*/
		Id:       2309,
		Password: "98795934",
		Name:     "魔加农",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     1700,
		Def:     1400,
		IsValid: true,
	})

	/* 206 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 231
		 调整:

		 中文名: 铠蜥蜴
		 卡片种类: 通常怪兽
		 卡片密码: 15480588
		 种族: 爬虫类
		 属性: 地
		 星级: 4
		 攻击力: 1500
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: RB、BE01、VOL04、DL02、JY、SJ2
		 效果: 描述：拥有坚硬躯体的蜥蜴。如果被咬到会立即毙命。

		*/
		Id:       231,
		Password: "15480588",
		Name:     "铠蜥蜴",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Reptile, // 爬虫类
		Atk:     1500,
		Def:     1200,
		IsValid: true,
	})

	/* 207 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2310
		 调整:

		 中文名: 甲化海星
		 卡片种类: 通常怪兽
		 卡片密码: 17535588
		 种族: 水
		 属性: 水
		 星级: 4
		 攻击力: 850
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: PG、VOL03
		 效果: 描述：表面坚固，守备力较高的青色海星。

		*/
		Id:       2310,
		Password: "17535588",
		Name:     "甲化海星",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     850,
		Def:     1400,
		IsValid: true,
	})

	/* 208 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2311
		 调整:

		 中文名: 鲜血吮吸者
		 卡片种类: 通常怪兽
		 卡片密码: 16353197
		 种族: 爬虫类
		 属性: 地
		 星级: 3
		 攻击力: 900
		 防御力: 800
		 罕见度: 平卡N
		 卡包: PG、VOL03
		 效果: 描述：借着夜色袭击行人的人形吸血蛇。

		*/
		Id:       2311,
		Password: "16353197",
		Name:     "鲜血吮吸者",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Reptile, // 爬虫类
		Atk:     900,
		Def:     800,
		IsValid: true,
	})

	/* 209 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2312
		 调整:

		 中文名: 艾尔丁
		 卡片种类: 通常怪兽
		 卡片密码: 06367785
		 种族: 魔法师
		 属性: 光
		 星级: 3
		 攻击力: 950
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: VOL03
		 效果: 描述：用手中的杖，使出各种魔法攻击。

		*/
		Id:       2312,
		Password: "06367785",
		Name:     "艾尔丁",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Light,       // 光
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     950,
		Def:     1000,
		IsValid: true,
	})

	/* 210 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2313
		 调整:

		 中文名: 蟹蛛
		 卡片种类: 通常怪兽
		 卡片密码: 34536276
		 种族: 昆虫
		 属性: 地
		 星级: 2
		 攻击力: 600
		 防御力: 800
		 罕见度: 平卡N
		 卡包: VOL03
		 效果: 描述：长有钳子的蜘蛛。用蜘蛛丝封住对方的动作用螃蟹钳子杀死对方。

		*/
		Id:       2313,
		Password: "34536276",
		Name:     "蟹蛛",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Earth,  // 地
		Lr:      ygo.LR_Insect, // 昆虫
		Atk:     600,
		Def:     800,
		IsValid: true,
	})

	/* 211 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2314
		 调整:

		 中文名: 螳螂杀手
		 卡片种类: 通常怪兽
		 卡片密码: 68928540
		 种族: 昆虫
		 属性: 地
		 星级: 4
		 攻击力: 1150
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: VOL03
		 效果: 描述：用两把镰刀袭击对方的一种人型螳螂怪兽。

		*/
		Id:       2314,
		Password: "68928540",
		Name:     "螳螂杀手",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth,  // 地
		Lr:      ygo.LR_Insect, // 昆虫
		Atk:     1150,
		Def:     1400,
		IsValid: true,
	})

	/* 212 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2315
		 调整:

		 中文名: 恐龙人
		 卡片种类: 通常怪兽
		 卡片密码: 89904598
		 种族: 恐龙
		 属性: 地
		 星级: 3
		 攻击力: 1000
		 防御力: 850
		 罕见度: 平卡N
		 卡包: VOL03
		 效果: 描述：人形的恐龙。虽然智慧很高，但攻击力却不怎么强劲。

		*/
		Id:       2315,
		Password: "89904598",
		Name:     "恐龙人",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,    // 地
		Lr:      ygo.LR_Dinosaur, // 恐龙
		Atk:     1000,
		Def:     850,
		IsValid: true,
	})

	/* 213 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2316
		 调整:

		 中文名: 杀人熊猫
		 卡片种类: 通常怪兽
		 卡片密码: 98818516
		 种族: 兽
		 属性: 地
		 星级: 4
		 攻击力: 1200
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: PG、VOL03
		 效果: 描述：经常手持一根粗大的竹子四处找茬，性格非常凶暴。

		*/
		Id:       2316,
		Password: "98818516",
		Name:     "杀人熊猫",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Atk:     1200,
		Def:     1000,
		IsValid: true,
	})

	/* 214 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2317
		 调整:

		 中文名: 彷徨的亡者
		 卡片种类: 通常怪兽
		 卡片密码: 93788854
		 种族: 不死
		 属性: 暗
		 星级: 2
		 攻击力: 800
		 防御力: 600
		 罕见度: 平卡N
		 卡包: VOL03、TP18
		 效果: 描述：因不能成佛，而只能漫无目的徘徊的怪兽。

		*/
		Id:       2317,
		Password: "93788854",
		Name:     "彷徨的亡者",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Dark,   // 暗
		Lr:      ygo.LR_Zombie, // 不死
		Atk:     800,
		Def:     600,
		IsValid: true,
	})

	/* 215 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2318
		 调整:

		 中文名: 地雷兽
		 卡片种类: 通常怪兽
		 卡片密码: 45042329
		 种族: 雷
		 属性: 地
		 星级: 4
		 攻击力: 1200
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: PG、VOL03
		 效果: 描述：强大的电磁波不停的向周围发射。

		*/
		Id:       2318,
		Password: "45042329",
		Name:     "地雷兽",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Thunder, // 雷
		Atk:     1200,
		Def:     1300,
		IsValid: true,
	})

	/* 216 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2319
		 调整:

		 中文名: 死亡潜行者
		 卡片种类: 通常怪兽
		 卡片密码: 54844990
		 种族: 战士
		 属性: 暗
		 星级: 3
		 攻击力: 900
		 防御力: 800
		 罕见度: 平卡N
		 卡包: VOL03
		 效果: 描述：行动敏捷、用钳子捕捉对方，带有剧毒刺针的蝎子战士。

		*/
		Id:       2319,
		Password: "54844990",
		Name:     "死亡潜行者",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Dark,    // 暗
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     900,
		Def:     800,
		IsValid: true,
	})

	/* 217 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 232
		 调整:

		 中文名: 杀人蜂
		 卡片种类: 通常怪兽
		 卡片密码: 88979991
		 种族: 昆虫
		 属性: 风
		 星级: 4
		 攻击力: 1200
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: RB、VOL05、DL02
		 效果: 描述：巨大的黄蜂，攻击出乎意料的强，若被群体攻击则非常危险。

		*/
		Id:       232,
		Password: "88979991",
		Name:     "杀人蜂",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Wind,   // 风
		Lr:      ygo.LR_Insect, // 昆虫
		Atk:     1200,
		Def:     1000,
		IsValid: true,
	})

	/* 218 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2320
		 调整:

		 中文名: 奈耳
		 卡片种类: 通常怪兽
		 卡片密码: 33178416
		 种族: 鱼
		 属性: 水
		 星级: 5
		 攻击力: 1400
		 防御力: 1600
		 罕见度: 平卡N
		 卡包: PG、VOL03
		 效果: 描述：全身长着针的鱼，从腹部下面发射导弹。

		*/
		Id:       2320,
		Password: "33178416",
		Name:     "奈耳",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Fish,  // 鱼
		Atk:     1400,
		Def:     1600,
		IsValid: true,
	})

	/* 219 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2321
		 调整:

		 中文名: 睡狮子
		 卡片种类: 通常怪兽
		 卡片密码: 40200834
		 种族: 兽
		 属性: 地
		 星级: 4
		 攻击力: 700
		 防御力: 1700
		 罕见度: 平卡N
		 卡包: VOL03
		 效果: 描述：平时沉睡着的猛兽。一旦苏醒就使人招架不住。

		*/
		Id:       2321,
		Password: "40200834",
		Name:     "睡狮子",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Atk:     700,
		Def:     1700,
		IsValid: true,
	})

	/* 220 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2322
		 调整:

		 中文名: 封印之锁
		 卡片种类: 通常怪兽
		 卡片密码: 08058240
		 种族: 天使
		 属性: 光
		 星级: 3
		 攻击力: 1000
		 防御力: 1100
		 罕见度: 平卡N
		 卡包: VOL03
		 效果: 描述：咯吱咯吱的勒住对方、拥有封印的力量。

		*/
		Id:       2322,
		Password: "08058240",
		Name:     "封印之锁",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Light, // 光
		Lr:      ygo.LR_Fairy, // 天使
		Atk:     1000,
		Def:     1100,
		IsValid: true,
	})

	/* 221 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2323
		 调整:

		 中文名: 魔界之棘
		 卡片种类: 通常怪兽
		 卡片密码: 43500484
		 种族: 植物
		 属性: 地
		 星级: 3
		 攻击力: 1200
		 防御力: 900
		 罕见度: 平卡N
		 卡包: PG、VOL03
		 效果: 描述：栖息在魔界的有刺灌木，将硬闯的人牢牢缠住。

		*/
		Id:       2323,
		Password: "43500484",
		Name:     "魔界之棘",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Plant, // 植物
		Atk:     1200,
		Def:     900,
		IsValid: true,
	})

	/* 222 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2324
		 调整:

		 中文名: 阁楼上的妖怪
		 卡片种类: 通常怪兽
		 卡片密码: 17238333
		 种族: 恶魔
		 属性: 暗
		 星级: 2
		 攻击力: 550
		 防御力: 400
		 罕见度: 平卡N
		 卡包: VOL03
		 效果: 描述：各家屋顶都有的冤魂。从不做坏事。

		*/
		Id:       2324,
		Password: "17238333",
		Name:     "阁楼上的妖怪",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     550,
		Def:     400,
		IsValid: true,
	})

	/* 223 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2330
		 调整:

		 中文名: 神鱼
		 卡片种类: 通常怪兽
		 卡片密码: 81386177
		 种族: 鱼
		 属性: 水
		 星级: 5
		 攻击力: 1650
		 防御力: 1700
		 罕见度: 平卡N
		 卡包: RB、VOL04
		 效果: 描述：在水中幽雅游动的鱼神，若它发怒会很危险。

		*/
		Id:       2330,
		Password: "81386177",
		Name:     "神鱼",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Fish,  // 鱼
		Atk:     1650,
		Def:     1700,
		IsValid: true,
	})

	/* 224 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2331
		 调整:

		 中文名: 耳天使
		 卡片种类: 通常怪兽
		 卡片密码: 86088138
		 种族: 天使
		 属性: 光
		 星级: 5
		 攻击力: 1550
		 防御力: 1650
		 罕见度: 平卡N
		 卡包: RB、VOL04
		 效果: 描述：巨大的耳朵和眼睛能监视周围的一切，相貌可怕的天使。

		*/
		Id:       2331,
		Password: "86088138",
		Name:     "耳天使",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Light, // 光
		Lr:      ygo.LR_Fairy, // 天使
		Atk:     1550,
		Def:     1650,
		IsValid: true,
	})

	/* 225 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2332
		 调整:

		 中文名: 死之沉默天使 多玛
		 卡片种类: 通常怪兽
		 卡片密码: 16972957
		 种族: 天使
		 属性: 暗
		 星级: 5
		 攻击力: 1600
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: RB、EX-R(EX)、VOL04
		 效果: 描述：司职死亡的天使，若被这家伙盯上，难逃一死。

		*/
		Id:       2332,
		Password: "16972957",
		Name:     "死之沉默天使 多玛",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fairy, // 天使
		Atk:     1600,
		Def:     1400,
		IsValid: true,
	})

	/* 226 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2333
		 调整:

		 中文名: 猎手蜘蛛
		 卡片种类: 通常怪兽
		 卡片密码: 80141480
		 种族: 昆虫
		 属性: 地
		 星级: 5
		 攻击力: 1600
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: RB、VOL04
		 效果: 描述：布置蜘蛛网狩猎，落入网中的东西将被吃掉。

		*/
		Id:       2333,
		Password: "80141480",
		Name:     "猎手蜘蛛",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Earth,  // 地
		Lr:      ygo.LR_Insect, // 昆虫
		Atk:     1600,
		Def:     1400,
		IsValid: true,
	})

	/* 227 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2334
		 调整:

		 中文名: 莫林芬
		 卡片种类: 通常怪兽
		 卡片密码: 55784832
		 种族: 恶魔
		 属性: 暗
		 星级: 5
		 攻击力: 1550
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: RB、VOL04
		 效果: 描述：以长臂钩爪为特征的长相奇妙的恶魔。

		*/
		Id:       2334,
		Password: "55784832",
		Name:     "莫林芬",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     1550,
		Def:     1300,
		IsValid: true,
	})

	/* 228 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2335
		 调整:

		 中文名: 古代精灵
		 卡片种类: 通常怪兽
		 卡片密码: 93221206
		 种族: 魔法师
		 属性: 光
		 星级: 4
		 攻击力: 1450
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: RB、EX-R(EX)、VOL04
		 效果: 描述：活了数千年的妖精，能够操纵精灵发动攻击。

		*/
		Id:       2335,
		Password: "93221206",
		Name:     "古代精灵",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Light,       // 光
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     1450,
		Def:     1200,
		IsValid: true,
	})

	/* 229 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2336
		 调整:

		 中文名: 黑影之鬼王
		 卡片种类: 通常怪兽
		 卡片密码: 45121025
		 种族: 兽战士
		 属性: 地
		 星级: 4
		 攻击力: 1200
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: EX-R(EX)、VOL04
		 效果: 描述：被黑影附体的恶鬼。用惊人的速度突击。

		*/
		Id:       2336,
		Password: "45121025",
		Name:     "黑影之鬼王",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth,        // 地
		Lr:      ygo.LR_BeastWarrior, // 兽战士
		Atk:     1200,
		Def:     1400,
		IsValid: true,
	})

	/* 230 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2337
		 调整:

		 中文名: 绝对鸟
		 卡片种类: 通常怪兽
		 卡片密码: 68870276
		 种族: 鸟兽
		 属性: 风
		 星级: 4
		 攻击力: 1300
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: VOL04
		 效果: 描述：可以将对方吸入镜中世界。

		*/
		Id:       2337,
		Password: "68870276",
		Name:     "绝对鸟",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Wind,        // 风
		Lr:      ygo.LR_WingedBeast, // 鸟兽
		Atk:     1300,
		Def:     1400,
		IsValid: true,
	})

	/* 231 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2338
		 调整:

		 中文名: 舌鱼
		 卡片种类: 通常怪兽
		 卡片密码: 69572024
		 种族: 鱼
		 属性: 水
		 星级: 4
		 攻击力: 1350
		 防御力: 800
		 罕见度: 平卡N
		 卡包: RB、VOL04
		 效果: 描述：用长舌头抓住其他的鱼，吸取能量。

		*/
		Id:       2338,
		Password: "69572024",
		Name:     "舌鱼",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Fish,  // 鱼
		Atk:     1350,
		Def:     800,
		IsValid: true,
	})

	/* 232 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2339
		 调整:

		 中文名: 龙人
		 卡片种类: 通常怪兽
		 卡片密码: 81057959
		 种族: 战士
		 属性: 地
		 星级: 4
		 攻击力: 1300
		 防御力: 1100
		 罕见度: 平卡N
		 卡包: EX-R(EX)、VOL04
		 效果: 描述：挥动用龙牙打造的剑、拥有龙之力量的战士。

		*/
		Id:       2339,
		Password: "81057959",
		Name:     "龙人",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     1300,
		Def:     1100,
		IsValid: true,
	})

	/* 233 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 234
		 调整:

		 中文名: 鹰身女郎
		 卡片种类: 通常怪兽
		 卡片密码: 76812113
		 种族: 鸟兽
		 属性: 风
		 星级: 4
		 攻击力: 1300
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: RB、BE01、VOL04、DL02
		 效果: 描述：人长出羽毛的魔兽，表演华丽的舞蹈进行快攻。

		*/
		Id:       234,
		Password: "76812113",
		Name:     "鹰身女郎",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Wind,        // 风
		Lr:      ygo.LR_WingedBeast, // 鸟兽
		Atk:     1300,
		Def:     1400,
		IsValid: true,
	})

	/* 234 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2340
		 调整:

		 中文名: 虎纹龙
		 卡片种类: 通常怪兽
		 卡片密码: 42348802
		 种族: 恐龙
		 属性: 地
		 星级: 3
		 攻击力: 1300
		 防御力: 800
		 罕见度: 平卡N
		 卡包: VOL04
		 效果: 描述：虎型恐龙。跑过旷野的速度相当快。

		*/
		Id:       2340,
		Password: "42348802",
		Name:     "虎纹龙",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,    // 地
		Lr:      ygo.LR_Dinosaur, // 恐龙
		Atk:     1300,
		Def:     800,
		IsValid: true,
	})

	/* 235 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2341
		 调整:

		 中文名: 大嘴
		 卡片种类: 通常怪兽
		 卡片密码: 55691901
		 种族: 兽
		 属性: 地
		 星级: 4
		 攻击力: 1250
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: VOL04
		 效果: 描述：拥有可以吞下一切的大嘴。

		*/
		Id:       2341,
		Password: "55691901",
		Name:     "大嘴",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Atk:     1250,
		Def:     1300,
		IsValid: true,
	})

	/* 236 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2342
		 调整:

		 中文名: 水之少女
		 卡片种类: 通常怪兽
		 卡片密码: 55014050
		 种族: 水
		 属性: 水
		 星级: 4
		 攻击力: 1250
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: VOL04
		 效果: 描述：用水像冰箭般攻击的美丽姐姐。

		*/
		Id:       2342,
		Password: "55014050",
		Name:     "水之少女",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     1250,
		Def:     1000,
		IsValid: true,
	})

	/* 237 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2343
		 调整:

		 中文名: 死神骷髅
		 卡片种类: 通常怪兽
		 卡片密码: 25882881
		 种族: 不死
		 属性: 暗
		 星级: 3
		 攻击力: 900
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: VOL04
		 效果: 描述：使用地狱一击夺取灵魂的死神。

		*/
		Id:       2343,
		Password: "25882881",
		Name:     "死神骷髅",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Dark,   // 暗
		Lr:      ygo.LR_Zombie, // 不死
		Atk:     900,
		Def:     1200,
		IsValid: true,
	})

	/* 238 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2344
		 调整:

		 中文名: 暗之暗杀者
		 卡片种类: 通常怪兽
		 卡片密码: 41949033
		 种族: 不死
		 属性: 暗
		 星级: 4
		 攻击力: 1200
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: EX-R(EX)、VOL04
		 效果: 描述：持有精神之剑、君临魔界的暗杀者。

		*/
		Id:       2344,
		Password: "41949033",
		Name:     "暗之暗杀者",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Dark,   // 暗
		Lr:      ygo.LR_Zombie, // 不死
		Atk:     1200,
		Def:     1200,
		IsValid: true,
	})

	/* 239 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2345
		 调整:

		 中文名: 阴阳师 道
		 卡片种类: 通常怪兽
		 卡片密码: 46247516
		 种族: 魔法师
		 属性: 地
		 星级: 3
		 攻击力: 1200
		 防御力: 900
		 罕见度: 平卡N
		 卡包: VOL04
		 效果: 描述：侵蚀阴与阳的力量、产生出邪恶力量的魔导师。

		*/
		Id:       2345,
		Password: "46247516",
		Name:     "阴阳师 道",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,       // 地
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     1200,
		Def:     900,
		IsValid: true,
	})

	/* 240 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2346
		 调整:

		 中文名: 猫妖精
		 卡片种类: 通常怪兽
		 卡片密码: 01761063
		 种族: 兽
		 属性: 地
		 星级: 3
		 攻击力: 1100
		 防御力: 900
		 罕见度: 平卡N
		 卡包: VOL04
		 效果: 描述：猫精灵。与可爱的身姿相反、敏捷的攻击对方。

		*/
		Id:       2346,
		Password: "01761063",
		Name:     "猫妖精",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Atk:     1100,
		Def:     900,
		IsValid: true,
	})

	/* 241 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2347
		 调整:

		 中文名: 枪管百合
		 卡片种类: 通常怪兽
		 卡片密码: 67841515
		 种族: 植物
		 属性: 地
		 星级: 3
		 攻击力: 1100
		 防御力: 600
		 罕见度: 平卡N
		 卡包: VOL04
		 效果: 描述：所谓的枪管百合。就是射出花粉攻击。

		*/
		Id:       2347,
		Password: "67841515",
		Name:     "枪管百合",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Plant, // 植物
		Atk:     1100,
		Def:     600,
		IsValid: true,
	})

	/* 242 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2348
		 调整:

		 中文名: 大雷电球
		 卡片种类: 通常怪兽
		 卡片密码: 21817254
		 种族: 雷
		 属性: 风
		 星级: 2
		 攻击力: 750
		 防御力: 600
		 罕见度: 平卡N
		 卡包: RB、VOL04、TP12
		 效果: 描述：在地止旋转，向周围发出电击的球。

		*/
		Id:       2348,
		Password: "21817254",
		Name:     "大雷电球",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Wind,    // 风
		Lr:      ygo.LR_Thunder, // 雷
		Atk:     750,
		Def:     600,
		IsValid: true,
	})

	/* 243 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2349
		 调整:

		 中文名: 火焰毒蛇
		 卡片种类: 通常怪兽
		 卡片密码: 02830619
		 种族: 炎
		 属性: 地
		 星级: 2
		 攻击力: 400
		 防御力: 450
		 罕见度: 平卡N
		 卡包: VOL04
		 效果: 描述：敏捷的行动、从口中吐出火焰的蛇。

		*/
		Id:       2349,
		Password: "02830619",
		Name:     "火焰毒蛇",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Fire,  // 炎
		Atk:     400,
		Def:     450,
		IsValid: true,
	})

	/* 244 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2350
		 调整:

		 中文名: 青虫
		 卡片种类: 通常怪兽
		 卡片密码: 81179446
		 种族: 昆虫
		 属性: 地
		 星级: 1
		 攻击力: 250
		 防御力: 300
		 罕见度: 平卡N
		 卡包: VOL04
		 效果: 描述：吐丝攻击。不知道以后会长成什么样子的虫。

		*/
		Id:       2350,
		Password: "81179446",
		Name:     "青虫",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   1,
		La:      ygo.LA_Earth,  // 地
		Lr:      ygo.LR_Insect, // 昆虫
		Atk:     250,
		Def:     300,
		IsValid: true,
	})

	/* 245 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2356
		 调整:

		 中文名: 特伦特
		 卡片种类: 通常怪兽
		 卡片密码: 78780140
		 种族: 植物
		 属性: 地
		 星级: 5
		 攻击力: 1500
		 防御力: 1800
		 罕见度: 平卡N
		 卡包: RB、VOL04
		 效果: 描述：还在逐渐成长的大树，是森林的守护神。

		*/
		Id:       2356,
		Password: "78780140",
		Name:     "特伦特",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Plant, // 植物
		Atk:     1500,
		Def:     1800,
		IsValid: true,
	})

	/* 246 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2358
		 调整:

		 中文名: 乌鸦天狗
		 卡片种类: 通常怪兽
		 卡片密码: 77998771
		 种族: 鸟兽
		 属性: 风
		 星级: 5
		 攻击力: 1850
		 防御力: 1600
		 罕见度: 银字R
		 卡包: VOL05、Booster05
		 效果: 描述：知道各种事的天狗。据说能使用神通之力。

		*/
		Id:       2358,
		Password: "77998771",
		Name:     "乌鸦天狗",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Wind,        // 风
		Lr:      ygo.LR_WingedBeast, // 鸟兽
		Atk:     1850,
		Def:     1600,
		IsValid: true,
	})

	/* 247 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2359
		 调整:

		 中文名: 迷宫的蠕虫
		 卡片种类: 通常怪兽
		 卡片密码: 51228280
		 种族: 昆虫
		 属性: 地
		 星级: 5
		 攻击力: 1800
		 防御力: 1500
		 罕见度: 银字R
		 卡包: VOL05、Booster05
		 效果: 描述：潜在迷路者的地下，捕食经过此处迷路的生物。

		*/
		Id:       2359,
		Password: "51228280",
		Name:     "迷宫的蠕虫",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Earth,  // 地
		Lr:      ygo.LR_Insect, // 昆虫
		Atk:     1800,
		Def:     1500,
		IsValid: true,
	})

	/* 248 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 236
		 调整:

		 中文名: 魔物狩人
		 卡片种类: 通常怪兽
		 卡片密码: 01184620
		 种族: 战士
		 属性: 地
		 星级: 4
		 攻击力: 1500
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: RB、BE01、EX-R(EX)、VOL04、DL02、SJ2
		 效果: 描述：猎杀人类的猎手，具有连岩石都能击碎的力量。

		*/
		Id:       236,
		Password: "01184620",
		Name:     "魔物狩人",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     1500,
		Def:     1200,
		IsValid: true,
	})

	/* 249 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2360
		 调整:

		 中文名: 粗暴帝王
		 卡片种类: 通常怪兽
		 卡片密码: 26378150
		 种族: 兽战士
		 属性: 地
		 星级: 5
		 攻击力: 1800
		 防御力: 1600
		 罕见度: 银字R
		 卡包: EX-R(EX)、VOL05、Booster05
		 效果: 描述：双手所持的魔人斧的破坏力相当强。

		*/
		Id:       2360,
		Password: "26378150",
		Name:     "粗暴帝王",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Earth,        // 地
		Lr:      ygo.LR_BeastWarrior, // 兽战士
		Atk:     1800,
		Def:     1600,
		IsValid: true,
	})

	/* 250 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2361
		 调整:

		 中文名: 暗黑兔
		 卡片种类: 通常怪兽
		 卡片密码: 99261403
		 种族: 兽
		 属性: 暗
		 星级: 4
		 攻击力: 1100
		 防御力: 1500
		 罕见度: 平卡N，银字R
		 卡包: VOL05、SOVR(606)、PE、TP18
		 效果: 描述：美国卡通世界里面的兔子罗杰。经常敏捷而匆忙的行动。

		*/
		Id:       2361,
		Password: "99261403",
		Name:     "暗黑兔",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Beast, // 兽
		Atk:     1100,
		Def:     1500,
		IsValid: true,
	})

	/* 251 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2362
		 调整:

		 中文名: 响女
		 卡片种类: 通常怪兽
		 卡片密码: 64501875
		 种族: 战士
		 属性: 地
		 星级: 4
		 攻击力: 1450
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: RB、VOL05
		 效果: 描述：发出扰乱听觉的噪音，使对方陷入行动不能的状态。

		*/
		Id:       2362,
		Password: "64501875",
		Name:     "响女",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     1450,
		Def:     1000,
		IsValid: true,
	})

	/* 252 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2363
		 调整:

		 中文名: 魔法幽灵
		 卡片种类: 通常怪兽
		 卡片密码: 46474915
		 种族: 不死
		 属性: 暗
		 星级: 4
		 攻击力: 1300
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: EX-R(EX)、VOL05
		 效果: 描述：向对方施展法术，使敌人陷入混乱和恐惧，从而胡乱攻击。

		*/
		Id:       2363,
		Password: "46474915",
		Name:     "魔法幽灵",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Dark,   // 暗
		Lr:      ygo.LR_Zombie, // 不死
		Atk:     1300,
		Def:     1400,
		IsValid: true,
	})

	/* 253 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2364
		 调整:

		 中文名: 圆盘魔术师
		 卡片种类: 通常怪兽
		 卡片密码: 76446915
		 种族: 机械
		 属性: 暗
		 星级: 4
		 攻击力: 1350
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: RB、VOL05
		 效果: 描述：将自己封印在铁盘里，攻击时才露出本体。

		*/
		Id:       2364,
		Password: "76446915",
		Name:     "圆盘魔术师",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Dark,    // 暗
		Lr:      ygo.LR_Machine, // 机械
		Atk:     1350,
		Def:     1000,
		IsValid: true,
	})

	/* 254 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2365
		 调整:

		 中文名: 生化植物
		 卡片种类: 通常怪兽
		 卡片密码: 07670542
		 种族: 恶魔
		 属性: 暗
		 星级: 3
		 攻击力: 600
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: VOL05
		 效果: 描述：地下研究所的试验大失败所产生的怪兽。

		*/
		Id:       2365,
		Password: "07670542",
		Name:     "生化植物",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     600,
		Def:     1300,
		IsValid: true,
	})

	/* 255 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2366
		 调整:

		 中文名: 泥浆潜居者
		 卡片种类: 通常怪兽
		 卡片密码: 18180762
		 种族: 岩石
		 属性: 地
		 星级: 4
		 攻击力: 1200
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: VOL05
		 效果: 描述：脚下粘呼呼的溶化时、是这家伙出现的前兆。

		*/
		Id:       2366,
		Password: "18180762",
		Name:     "泥浆潜居者",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Rock,  // 岩石
		Atk:     1200,
		Def:     1300,
		IsValid: true,
	})

	/* 256 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2367
		 调整:

		 中文名: 切割机器人
		 卡片种类: 通常怪兽
		 卡片密码: 10315429
		 种族: 机械
		 属性: 暗
		 星级: 4
		 攻击力: 1000
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: VOL05
		 效果: 描述：方形的身体内部隐藏着无数刀刃、能切碎一切的机械。

		*/
		Id:       2367,
		Password: "10315429",
		Name:     "切割机器人",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Dark,    // 暗
		Lr:      ygo.LR_Machine, // 机械
		Atk:     1000,
		Def:     1300,
		IsValid: true,
	})

	/* 257 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2368
		 调整:

		 中文名: 翼蛋精灵
		 卡片种类: 通常怪兽
		 卡片密码: 98582704
		 种族: 天使
		 属性: 光
		 星级: 3
		 攻击力: 500
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: VOL05
		 效果: 描述：被蛋壳包住的天使。用巨大翅膀守备攻击。

		*/
		Id:       2368,
		Password: "98582704",
		Name:     "翼蛋精灵",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Light, // 光
		Lr:      ygo.LR_Fairy, // 天使
		Atk:     500,
		Def:     1300,
		IsValid: true,
	})

	/* 258 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2369
		 调整:

		 中文名: 冰
		 卡片种类: 通常怪兽
		 卡片密码: 38982356
		 种族: 战士
		 属性: 水
		 星级: 3
		 攻击力: 800
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: VOL05
		 效果: 描述：全身用冰做的战士。将一切与其解触者冰冻起来。

		*/
		Id:       2369,
		Password: "38982356",
		Name:     "冰",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Water,   // 水
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     800,
		Def:     1200,
		IsValid: true,
	})

	/* 259 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2370
		 调整:

		 中文名: 主人与能手
		 卡片种类: 通常怪兽
		 卡片密码: 75499502
		 种族: 兽
		 属性: 地
		 星级: 4
		 攻击力: 1200
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: EX-R(EX)、VOL05
		 效果: 描述：怪兽使高手、忠心于主人的怪兽。组合起来相当完美。

		*/
		Id:       2370,
		Password: "75499502",
		Name:     "主人与能手",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Atk:     1200,
		Def:     1000,
		IsValid: true,
	})

	/* 260 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2371
		 调整:

		 中文名: 蛇椰树
		 卡片种类: 通常怪兽
		 卡片密码: 29802344
		 种族: 植物
		 属性: 地
		 星级: 4
		 攻击力: 1000
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: VOL05
		 效果: 描述：许多的蛇聚集在一起。一旦敌人靠近就分散进行攻击。

		*/
		Id:       2371,
		Password: "29802344",
		Name:     "蛇椰树",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Plant, // 植物
		Atk:     1000,
		Def:     1200,
		IsValid: true,
	})

	/* 261 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2372
		 调整:

		 中文名: 铠鼠
		 卡片种类: 通常怪兽
		 卡片密码: 16246527
		 种族: 兽
		 属性: 地
		 星级: 3
		 攻击力: 950
		 防御力: 1100
		 罕见度: 平卡N
		 卡包: VOL05
		 效果: 描述：可以用像铠甲般坚固的毛守护身体的老鼠。

		*/
		Id:       2372,
		Password: "16246527",
		Name:     "铠鼠",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Atk:     950,
		Def:     1100,
		IsValid: true,
	})

	/* 262 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2373
		 调整:

		 中文名: 恶之无名战士
		 卡片种类: 通常怪兽
		 卡片密码: 97360116
		 种族: 战士
		 属性: 暗
		 星级: 3
		 攻击力: 1000
		 防御力: 500
		 罕见度: 平卡N
		 卡包: EX-R(EX)、VOL05
		 效果: 描述：用敏捷的动作产生真空、切碎对方的战士。

		*/
		Id:       2373,
		Password: "97360116",
		Name:     "恶之无名战士",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Dark,    // 暗
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     1000,
		Def:     500,
		IsValid: true,
	})

	/* 263 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2374
		 调整:

		 中文名: 蛤蟆仙人
		 卡片种类: 通常怪兽
		 卡片密码: 62671448
		 种族: 水
		 属性: 水
		 星级: 3
		 攻击力: 1000
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: VOL05
		 效果: 描述：活了几千年的青蛙仙人。用蝌蚪进行攻击。

		*/
		Id:       2374,
		Password: "62671448",
		Name:     "蛤蟆仙人",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     1000,
		Def:     1000,
		IsValid: true,
	})

	/* 264 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2375
		 调整:

		 中文名: 暗影
		 卡片种类: 通常怪兽
		 卡片密码: 40196604
		 种族: 恶魔
		 属性: 风
		 星级: 3
		 攻击力: 1000
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: VOL05
		 效果: 描述：从水晶中发射强烈的光线进行攻击。

		*/
		Id:       2375,
		Password: "40196604",
		Name:     "暗影",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Wind,  // 风
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     1000,
		Def:     1000,
		IsValid: true,
	})

	/* 265 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2376
		 调整:

		 中文名: 酸液爬虫
		 卡片种类: 通常怪兽
		 卡片密码: 77568553
		 种族: 昆虫
		 属性: 地
		 星级: 3
		 攻击力: 900
		 防御力: 700
		 罕见度: 平卡N
		 卡包: VOL05
		 效果: 描述：巨大的青虫。吐出浓烈的酸液、可以溶解任何东西。

		*/
		Id:       2376,
		Password: "77568553",
		Name:     "酸液爬虫",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,  // 地
		Lr:      ygo.LR_Insect, // 昆虫
		Atk:     900,
		Def:     700,
		IsValid: true,
	})

	/* 266 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 238
		 调整:

		 中文名: 伏地龙
		 卡片种类: 通常怪兽
		 卡片密码: 67494157
		 种族: 龙
		 属性: 地
		 星级: 5
		 攻击力: 1600
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: RB、VOL04、DL02
		 效果: 描述：力量弱小无法飞天的龙，但攻击还是很强的。

		*/
		Id:       238,
		Password: "67494157",
		Name:     "伏地龙",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Earth,  // 地
		Lr:      ygo.LR_Dragon, // 龙
		Atk:     1600,
		Def:     1400,
		IsValid: true,
	})

	/* 267 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2381
		 调整:

		 中文名: 海之龙王
		 卡片种类: 通常怪兽
		 卡片密码: 23659124
		 种族: 海龙
		 属性: 水
		 星级: 6
		 攻击力: 2000
		 防御力: 1700
		 罕见度: 平卡N
		 卡包: VOL06
		 效果: 描述：海之王。长有坚硬的贝壳，从口中吐出水泡攻击。

		*/
		Id:       2381,
		Password: "23659124",
		Name:     "海之龙王",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   6,
		La:      ygo.LA_Water,      // 水
		Lr:      ygo.LR_SeaSerpent, // 海龙
		Atk:     2000,
		Def:     1700,
		IsValid: true,
	})

	/* 268 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2382
		 调整:

		 中文名: 猫女郎
		 卡片种类: 通常怪兽
		 卡片密码: 43352213
		 种族: 兽战士
		 属性: 地
		 星级: 6
		 攻击力: 1900
		 防御力: 2000
		 罕见度: 平卡N
		 卡包: VOL06
		 效果: 描述：用敏捷的动作躲避攻击、用锋利的钩爪袭击对方。

		*/
		Id:       2382,
		Password: "43352213",
		Name:     "猫女郎",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   6,
		La:      ygo.LA_Earth,        // 地
		Lr:      ygo.LR_BeastWarrior, // 兽战士
		Atk:     1900,
		Def:     2000,
		IsValid: true,
	})

	/* 269 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2383
		 调整:

		 中文名: 骰子犰狳
		 卡片种类: 通常怪兽
		 卡片密码: 69893315
		 种族: 机械
		 属性: 地
		 星级: 5
		 攻击力: 1650
		 防御力: 1800
		 罕见度: 平卡N
		 卡包: VOL06
		 效果: 描述：卷起身子就像骰子一样的犰狳。

		*/
		Id:       2383,
		Password: "69893315",
		Name:     "骰子犰狳",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Machine, // 机械
		Atk:     1650,
		Def:     1800,
		IsValid: true,
	})

	/* 270 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2384
		 调整:

		 中文名: 苍翼冠鸟
		 卡片种类: 通常怪兽
		 卡片密码: 41396436
		 种族: 鸟兽
		 属性: 风
		 星级: 4
		 攻击力: 1600
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: ME、VOL06
		 效果: 描述：青白色的火鸟，头上的毛有如皇冠。

		*/
		Id:       2384,
		Password: "41396436",
		Name:     "苍翼冠鸟",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Wind,        // 风
		Lr:      ygo.LR_WingedBeast, // 鸟兽
		Atk:     1600,
		Def:     1200,
		IsValid: true,
	})

	/* 271 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2385
		 调整:

		 中文名: 暴雨怪
		 卡片种类: 通常怪兽
		 卡片密码: 94042337
		 种族: 水
		 属性: 水
		 星级: 4
		 攻击力: 1550
		 防御力: 800
		 罕见度: 平卡N
		 卡包: VOL06
		 效果: 描述：突然降下倾盆大雨的怪兽。

		*/
		Id:       2385,
		Password: "94042337",
		Name:     "暴雨怪",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     1550,
		Def:     800,
		IsValid: true,
	})

	/* 272 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2386
		 调整:

		 中文名: 电击蜗牛
		 卡片种类: 通常怪兽
		 卡片密码: 12146024
		 种族: 雷
		 属性: 水
		 星级: 5
		 攻击力: 1400
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: VOL06
		 效果: 描述：吐出粘液然后用电击对方。

		*/
		Id:       2386,
		Password: "12146024",
		Name:     "电击蜗牛",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Water,   // 水
		Lr:      ygo.LR_Thunder, // 雷
		Atk:     1400,
		Def:     1500,
		IsValid: true,
	})

	/* 273 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2387
		 调整:

		 中文名: 小鬼
		 卡片种类: 通常怪兽
		 卡片密码: 69669405
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 1300
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: VOL06
		 效果: 描述：住在黑暗的小鬼。攻击意外的强。要注意它头上的角。

		*/
		Id:       2387,
		Password: "69669405",
		Name:     "小鬼",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     1300,
		Def:     1000,
		IsValid: true,
	})

	/* 274 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2388
		 调整:

		 中文名: 杰米亚之神
		 卡片种类: 通常怪兽
		 卡片密码: 81618817
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 1300
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: VOL06
		 效果: 描述：善于欺骗对方、诱惑其加入破灭之路的邪恶之神。

		*/
		Id:       2388,
		Password: "81618817",
		Name:     "杰米亚之神",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     1300,
		Def:     1000,
		IsValid: true,
	})

	/* 275 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2389
		 调整:

		 中文名: 狮鹫兽卫
		 卡片种类: 通常怪兽
		 卡片密码: 53829412
		 种族: 兽
		 属性: 地
		 星级: 4
		 攻击力: 1200
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: VOL06
		 效果: 描述：善于用坚固的身体守备。将对方的攻击反弹。

		*/
		Id:       2389,
		Password: "53829412",
		Name:     "狮鹫兽卫",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Atk:     1200,
		Def:     1500,
		IsValid: true,
	})

	/* 276 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 239
		 调整:

		 中文名: 铠武者僵尸
		 卡片种类: 通常怪兽
		 卡片密码: 20277860
		 种族: 不死
		 属性: 暗
		 星级: 3
		 攻击力: 1500
		 防御力: 0
		 罕见度: 平卡N
		 卡包: RB、BE01、VOL05、DL02
		 效果: 描述：充满怨念而苏醒的武者，要警惕在黑暗中挥舞的太刀。

		*/
		Id:       239,
		Password: "20277860",
		Name:     "铠武者僵尸",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Dark,   // 暗
		Lr:      ygo.LR_Zombie, // 不死
		Atk:     1500,
		Def:     0,
		IsValid: true,
	})

	/* 277 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2390
		 调整:

		 中文名: 蟑螂球
		 卡片种类: 通常怪兽
		 卡片密码: 15367030
		 种族: 昆虫
		 属性: 地
		 星级: 4
		 攻击力: 1200
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: VOL06
		 效果: 描述：圆形的小矮人。咕噜咕噜的滚动攻击。守备意外的高。

		*/
		Id:       2390,
		Password: "15367030",
		Name:     "蟑螂球",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth,  // 地
		Lr:      ygo.LR_Insect, // 昆虫
		Atk:     1200,
		Def:     1400,
		IsValid: true,
	})

	/* 278 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2391
		 调整:

		 中文名: 犀牛
		 卡片种类: 通常怪兽
		 卡片密码: 80813021
		 种族: 兽
		 属性: 地
		 星级: 3
		 攻击力: 1200
		 防御力: 600
		 罕见度: 平卡N
		 卡包: VOL06
		 效果: 描述：虽然守备不像外表般高强。但是角的攻击力还是很强。

		*/
		Id:       2391,
		Password: "80813021",
		Name:     "犀牛",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Atk:     1200,
		Def:     600,
		IsValid: true,
	})

	/* 279 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2392
		 调整:

		 中文名: 冰水
		 卡片种类: 通常怪兽
		 卡片密码: 20848593
		 种族: 水
		 属性: 水
		 星级: 3
		 攻击力: 1150
		 防御力: 900
		 罕见度: 平卡N
		 卡包: VOL06
		 效果: 描述：攻击性美人鱼。使用生长在身体表面的荆刺进行攻击。

		*/
		Id:       2392,
		Password: "20848593",
		Name:     "冰水",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     1150,
		Def:     900,
		IsValid: true,
	})

	/* 280 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2393
		 调整:

		 中文名: 冻土带的大蝎
		 卡片种类: 通常怪兽
		 卡片密码: 41403766
		 种族: 昆虫
		 属性: 地
		 星级: 3
		 攻击力: 1100
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: VOL06
		 效果: 描述：不是在沙漠而是分布在冰土带上的罕见的深蓝色的蝎子。

		*/
		Id:       2393,
		Password: "41403766",
		Name:     "冻土带的大蝎",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,  // 地
		Lr:      ygo.LR_Insect, // 昆虫
		Atk:     1100,
		Def:     1000,
		IsValid: true,
	})

	/* 281 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2394
		 调整:

		 中文名: 冥界的番人
		 卡片种类: 通常怪兽
		 卡片密码: 89272878
		 种族: 战士
		 属性: 地
		 星级: 4
		 攻击力: 1000
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: ME、VOL06
		 效果: 描述：守护着冥界入口的战士，对侵略者豪不留情的斩杀。

		*/
		Id:       2394,
		Password: "89272878",
		Name:     "冥界的番人",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     1000,
		Def:     1200,
		IsValid: true,
	})

	/* 282 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2395
		 调整:

		 中文名: 铁铲粉碎机
		 卡片种类: 通常怪兽
		 卡片密码: 71950093
		 种族: 机械
		 属性: 地
		 星级: 3
		 攻击力: 900
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: VOL06
		 效果: 描述：善于破坏一切。要注意其双手的大铲。

		*/
		Id:       2395,
		Password: "71950093",
		Name:     "铁铲粉碎机",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Machine, // 机械
		Atk:     900,
		Def:     1200,
		IsValid: true,
	})

	/* 283 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2396
		 调整:

		 中文名: 雏鸡
		 卡片种类: 通常怪兽
		 卡片密码: 07805359
		 种族: 鸟兽
		 属性: 地
		 星级: 3
		 攻击力: 900
		 防御力: 800
		 罕见度: 平卡N
		 卡包: ME、VOL06
		 效果: 描述：把敌人整个吞下，并消化作为能量来补充自己。

		*/
		Id:       2396,
		Password: "07805359",
		Name:     "雏鸡",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,       // 地
		Lr:      ygo.LR_WingedBeast, // 鸟兽
		Atk:     900,
		Def:     800,
		IsValid: true,
	})

	/* 284 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2397
		 调整:

		 中文名: 笑花
		 卡片种类: 通常怪兽
		 卡片密码: 42591472
		 种族: 植物
		 属性: 地
		 星级: 2
		 攻击力: 900
		 防御力: 500
		 罕见度: 平卡N
		 卡包: VOL06
		 效果: 描述：总是在笑着的花。如果持续听它的笑声的话，头脑就会变得混乱。

		*/
		Id:       2397,
		Password: "42591472",
		Name:     "笑花",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Plant, // 植物
		Atk:     900,
		Def:     500,
		IsValid: true,
	})

	/* 285 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2398
		 调整:

		 中文名: 强盗鼠
		 卡片种类: 通常怪兽
		 卡片密码: 06297941
		 种族: 兽
		 属性: 地
		 星级: 3
		 攻击力: 850
		 防御力: 800
		 罕见度: 平卡N
		 卡包: VOL06
		 效果: 描述：奸诈狡猾的老鼠。使用左手的大钩爪进行攻击。

		*/
		Id:       2398,
		Password: "06297941",
		Name:     "强盗鼠",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Atk:     850,
		Def:     800,
		IsValid: true,
	})

	/* 286 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2399
		 调整:

		 中文名: 王座守护者
		 卡片种类: 通常怪兽
		 卡片密码: 10071456
		 种族: 战士
		 属性: 地
		 星级: 4
		 攻击力: 800
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: ME、VOL06
		 效果: 描述：国王离开时，保卫王座不受外敌攻击的王妃。守备力高。

		*/
		Id:       2399,
		Password: "10071456",
		Name:     "王座守护者",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     800,
		Def:     1500,
		IsValid: true,
	})

	/* 287 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2400
		 调整:

		 中文名: 棱镜人
		 卡片种类: 通常怪兽
		 卡片密码: 80234301
		 种族: 岩石
		 属性: 光
		 星级: 3
		 攻击力: 800
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: VOL06
		 效果: 描述：晶莹剔透的水晶凝固体。收集光线而射出激光。

		*/
		Id:       2400,
		Password: "80234301",
		Name:     "棱镜人",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Light, // 光
		Lr:      ygo.LR_Rock,  // 岩石
		Atk:     800,
		Def:     1000,
		IsValid: true,
	})

	/* 288 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2401
		 调整:

		 中文名: 突击兵
		 卡片种类: 通常怪兽
		 卡片密码: 06400512
		 种族: 机械
		 属性: 暗
		 星级: 2
		 攻击力: 750
		 防御力: 700
		 罕见度: 平卡N
		 卡包: VOL06
		 效果: 描述：装备火箭发射器与反坦克火箭炮的实战部队。

		*/
		Id:       2401,
		Password: "06400512",
		Name:     "突击兵",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Dark,    // 暗
		Lr:      ygo.LR_Machine, // 机械
		Atk:     750,
		Def:     700,
		IsValid: true,
	})

	/* 289 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2402
		 调整:

		 中文名: 温柔天使
		 卡片种类: 通常怪兽
		 卡片密码: 57935140
		 种族: 天使
		 属性: 光
		 星级: 3
		 攻击力: 700
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: VOL06
		 效果: 描述：为恋人们献上祝福的可爱天使。

		*/
		Id:       2402,
		Password: "57935140",
		Name:     "温柔天使",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Light, // 光
		Lr:      ygo.LR_Fairy, // 天使
		Atk:     700,
		Def:     1400,
		IsValid: true,
	})

	/* 290 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2403
		 调整:

		 中文名: 甲壳蟹
		 卡片种类: 通常怪兽
		 卡片密码: 84103702
		 种族: 水
		 属性: 水
		 星级: 3
		 攻击力: 650
		 防御力: 900
		 罕见度: 平卡N
		 卡包: VOL06
		 效果: 描述：螃蟹怪兽。用双手的大钳子切碎对方。

		*/
		Id:       2403,
		Password: "84103702",
		Name:     "甲壳蟹",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     650,
		Def:     900,
		IsValid: true,
	})

	/* 291 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2406
		 调整:

		 中文名: 火焰地狱犬
		 卡片种类: 通常怪兽
		 卡片密码: 60862676
		 种族: 炎
		 属性: 炎
		 星级: 6
		 攻击力: 2100
		 防御力: 1800
		 罕见度: 平卡N
		 卡包: ME、VOL06
		 效果: 描述：全身被火包围的魔兽，给予对方地狱之炎的处刑。

		*/
		Id:       2406,
		Password: "60862676",
		Name:     "火焰地狱犬",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   6,
		La:      ygo.LA_Fire, // 炎
		Lr:      ygo.LR_Fire, // 炎
		Atk:     2100,
		Def:     1800,
		IsValid: true,
	})

	/* 292 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2407
		 调整:

		 中文名: 空气吸收者
		 卡片种类: 通常怪兽
		 卡片密码: 08353769
		 种族: 恶魔
		 属性: 风
		 星级: 6
		 攻击力: 2100
		 防御力: 1600
		 罕见度: 平卡N
		 卡包: VOL06
		 效果: 描述：吸收周围的空气，使对方窒息的妖怪。

		*/
		Id:       2407,
		Password: "08353769",
		Name:     "空气吸收者",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   6,
		La:      ygo.LA_Wind,  // 风
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     2100,
		Def:     1600,
		IsValid: true,
	})

	/* 293 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2408
		 调整:

		 中文名: 王家守卫
		 卡片种类: 通常怪兽
		 卡片密码: 39239728
		 种族: 机械
		 属性: 地
		 星级: 6
		 攻击力: 1900
		 防御力: 2200
		 罕见度: 平卡N
		 卡包: VOL07
		 效果: 描述：为了守护王座而开发的，且拥有意志的机械部队。

		*/
		Id:       2408,
		Password: "39239728",
		Name:     "王家守卫",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   6,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Machine, // 机械
		Atk:     1900,
		Def:     2200,
		IsValid: true,
	})

	/* 294 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2409
		 调整:

		 中文名: 铁球巨人
		 卡片种类: 通常怪兽
		 卡片密码: 33621868
		 种族: 机械
		 属性: 暗
		 星级: 5
		 攻击力: 1700
		 防御力: 1800
		 罕见度: 平卡N
		 卡包: VOL07
		 效果: 描述：振舞着带刺的铁球进行攻击。

		*/
		Id:       2409,
		Password: "33621868",
		Name:     "铁球巨人",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Dark,    // 暗
		Lr:      ygo.LR_Machine, // 机械
		Atk:     1700,
		Def:     1800,
		IsValid: true,
	})

	/* 295 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2410
		 调整:

		 中文名: 钢铁之心
		 卡片种类: 通常怪兽
		 卡片密码: 49587396
		 种族: 机械
		 属性: 暗
		 星级: 5
		 攻击力: 1700
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: VOL07
		 效果: 描述：在古代文明遗迹里发现的机器。破坏是它的唯一目的。

		*/
		Id:       2410,
		Password: "49587396",
		Name:     "钢铁之心",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Dark,    // 暗
		Lr:      ygo.LR_Machine, // 机械
		Atk:     1700,
		Def:     1400,
		IsValid: true,
	})

	/* 296 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2411
		 调整:

		 中文名: 暗黑奇美拉
		 卡片种类: 通常怪兽
		 卡片密码: 32344688
		 种族: 恶魔
		 属性: 暗
		 星级: 5
		 攻击力: 1610
		 防御力: 1460
		 罕见度: 平卡N
		 卡包: VOL07
		 效果: 描述：生息在魔界的怪兽。吐出暗炎攻击。

		*/
		Id:       2411,
		Password: "32344688",
		Name:     "暗黑奇美拉",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     1610,
		Def:     1460,
		IsValid: true,
	})

	/* 297 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2412
		 调整:

		 中文名: 撞击兽
		 卡片种类: 通常怪兽
		 卡片密码: 33878931
		 种族: 兽
		 属性: 地
		 星级: 5
		 攻击力: 1600
		 防御力: 1800
		 罕见度: 平卡N
		 卡包: VOL07
		 效果: 描述：动作意外地敏捷，把身体卷成圆形进行撞击。

		*/
		Id:       2412,
		Password: "33878931",
		Name:     "撞击兽",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Atk:     1600,
		Def:     1800,
		IsValid: true,
	})

	/* 298 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2413
		 调整:

		 中文名: 贪食的食尸鬼
		 卡片种类: 通常怪兽
		 卡片密码: 95265975
		 种族: 不死
		 属性: 暗
		 星级: 4
		 攻击力: 1600
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: VOL07
		 效果: 描述：无论怎么吃也不会饱的怪兽。

		*/
		Id:       2413,
		Password: "95265975",
		Name:     "贪食的食尸鬼",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Dark,   // 暗
		Lr:      ygo.LR_Zombie, // 不死
		Atk:     1600,
		Def:     1200,
		IsValid: true,
	})

	/* 299 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2414
		 调整:

		 中文名: 狂鱼
		 卡片种类: 通常怪兽
		 卡片密码: 53713014
		 种族: 鱼
		 属性: 水
		 星级: 4
		 攻击力: 1600
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: VOL07
		 效果: 描述：利用尖锐的头刺撞过来攻击敌人。

		*/
		Id:       2414,
		Password: "53713014",
		Name:     "狂鱼",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Fish,  // 鱼
		Atk:     1600,
		Def:     1200,
		IsValid: true,
	})

	/* 300 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2415
		 调整:

		 中文名: 看门人
		 卡片种类: 通常怪兽
		 卡片密码: 19737320
		 种族: 机械
		 属性: 暗
		 星级: 5
		 攻击力: 1500
		 防御力: 1800
		 罕见度: 平卡N
		 卡包: VOL07
		 效果: 描述：为守备王家墓地入口而造的机器。非常难以被破坏。

		*/
		Id:       2415,
		Password: "19737320",
		Name:     "看门人",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Dark,    // 暗
		Lr:      ygo.LR_Machine, // 机械
		Atk:     1500,
		Def:     1800,
		IsValid: true,
	})

	/* 301 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2416
		 调整:

		 中文名: 机械士兵
		 卡片种类: 通常怪兽
		 卡片密码: 44865098
		 种族: 机械
		 属性: 暗
		 星级: 5
		 攻击力: 1500
		 防御力: 1700
		 罕见度: 平卡N
		 卡包: VOL07
		 效果: 描述：护卫机械王的兵队。圆圆的身体到处转动着进行巡逻。

		*/
		Id:       2416,
		Password: "44865098",
		Name:     "机械士兵",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Dark,    // 暗
		Lr:      ygo.LR_Machine, // 机械
		Atk:     1500,
		Def:     1700,
		IsValid: true,
	})

	/* 302 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2417
		 调整:

		 中文名: 铠武者 斩鬼
		 卡片种类: 通常怪兽
		 卡片密码: 30090452
		 种族: 战士
		 属性: 地
		 星级: 5
		 攻击力: 1500
		 防御力: 1700
		 罕见度: 平卡N
		 卡包: VOL07
		 效果: 描述：喜好单挑，能抓住一瞬间的机会发动反击。

		*/
		Id:       2417,
		Password: "30090452",
		Name:     "铠武者 斩鬼",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     1500,
		Def:     1700,
		IsValid: true,
	})

	/* 303 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2418
		 调整:

		 中文名: 杀人机器
		 卡片种类: 通常怪兽
		 卡片密码: 73911410
		 种族: 机械
		 属性: 暗
		 星级: 5
		 攻击力: 1450
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: VOL07
		 效果: 描述：挥舞着大弯刀的狂暴杀人机器。

		*/
		Id:       2418,
		Password: "73911410",
		Name:     "杀人机器",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Dark,    // 暗
		Lr:      ygo.LR_Machine, // 机械
		Atk:     1450,
		Def:     1500,
		IsValid: true,
	})

	/* 304 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2419
		 调整:

		 中文名: 被诅咒的魔剑
		 卡片种类: 通常怪兽
		 卡片密码: 22855882
		 种族: 战士
		 属性: 暗
		 星级: 4
		 攻击力: 1400
		 防御力: 800
		 罕见度: 平卡N
		 卡包: VOL07
		 效果: 描述：据说战胜施加在身上的诅咒的人就可以以此获得力量。

		*/
		Id:       2419,
		Password: "22855882",
		Name:     "被诅咒的魔剑",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Dark,    // 暗
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     1400,
		Def:     800,
		IsValid: true,
	})

	/* 305 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2420
		 调整:

		 中文名: 杀人小丑僵尸
		 卡片种类: 通常怪兽
		 卡片密码: 92667214
		 种族: 不死
		 属性: 暗
		 星级: 2
		 攻击力: 1350
		 防御力: 0
		 罕见度: 平卡N
		 卡包: VOL07
		 效果: 描述：使用黑暗力量重生的小丑，利用摇晃的舞蹈将对方引向死亡。

		*/
		Id:       2420,
		Password: "92667214",
		Name:     "杀人小丑僵尸",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Dark,   // 暗
		Lr:      ygo.LR_Zombie, // 不死
		Atk:     1350,
		Def:     0,
		IsValid: true,
	})

	/* 306 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2421
		 调整:

		 中文名: 东方英雄
		 卡片种类: 通常怪兽
		 卡片密码: 89987208
		 种族: 战士
		 属性: 地
		 星级: 3
		 攻击力: 1100
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: VOL07
		 效果: 描述：传说来自遥远的东方国度的武士。手持的刀相当的锋利。

		*/
		Id:       2421,
		Password: "89987208",
		Name:     "东方英雄",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     1100,
		Def:     1000,
		IsValid: true,
	})

	/* 307 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2422
		 调整:

		 中文名: 温蒂妮
		 卡片种类: 通常怪兽
		 卡片密码: 66836598
		 种族: 水
		 属性: 水
		 星级: 3
		 攻击力: 1100
		 防御力: 700
		 罕见度: 平卡N
		 卡包: VOL07
		 效果: 描述：在水中摇摇晃晃漂浮的妖精。好像可以召唤水龙。

		*/
		Id:       2422,
		Password: "66836598",
		Name:     "温蒂妮",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     1100,
		Def:     700,
		IsValid: true,
	})

	/* 308 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2423
		 调整:

		 中文名: 菊石骑士
		 卡片种类: 通常怪兽
		 卡片密码: 36151751
		 种族: 水
		 属性: 水
		 星级: 4
		 攻击力: 1000
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: VOL07
		 效果: 描述：从很久以前就守护着大海的骑士。

		*/
		Id:       2423,
		Password: "36151751",
		Name:     "菊石骑士",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     1000,
		Def:     1200,
		IsValid: true,
	})

	/* 309 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2424
		 调整:

		 中文名: 红A
		 卡片种类: 通常怪兽
		 卡片密码: 38035986
		 种族: 魔法师
		 属性: 暗
		 星级: 3
		 攻击力: 1000
		 防御力: 800
		 罕见度: 平卡N
		 卡包: VOL07
		 效果: 描述：施放死亡诅咒的魔法师。听到咒文的敌人将变得神智不清。

		*/
		Id:       2424,
		Password: "38035986",
		Name:     "红A",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Dark,        // 暗
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     1000,
		Def:     800,
		IsValid: true,
	})

	/* 310 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2425
		 调整:

		 中文名: 卡通鳄鱼
		 卡片种类: 通常怪兽
		 卡片密码: 59383041
		 种族: 爬虫类
		 属性: 水
		 星级: 4
		 攻击力: 800
		 防御力: 1600
		 罕见度: 平卡N
		 卡包: VOL07、PE
		 效果: 描述：从美国卡通世界中出现的鳄鱼型怪兽。

		*/
		Id:       2425,
		Password: "59383041",
		Name:     "卡通鳄鱼",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Water,   // 水
		Lr:      ygo.LR_Reptile, // 爬虫类
		Atk:     800,
		Def:     1600,
		IsValid: true,
	})

	/* 311 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2426
		 调整:

		 中文名: 阿罗妮
		 卡片种类: 通常怪兽
		 卡片密码: 14708569
		 种族: 植物
		 属性: 地
		 星级: 3
		 攻击力: 800
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: VOL07
		 效果: 描述：花中的女性散发着有毒的花粉。任何人都不能靠近她。

		*/
		Id:       2426,
		Password: "14708569",
		Name:     "阿罗妮",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Plant, // 植物
		Atk:     800,
		Def:     1000,
		IsValid: true,
	})

	/* 312 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2427
		 调整:

		 中文名: 硬毛鼠
		 卡片种类: 通常怪兽
		 卡片密码: 00549481
		 种族: 兽
		 属性: 地
		 星级: 4
		 攻击力: 500
		 防御力: 2000
		 罕见度: 平卡N
		 卡包: ME、VOL07
		 效果: 描述：毛像硬皮一样聚集在一起，守备力非常的高。

		*/
		Id:       2427,
		Password: "00549481",
		Name:     "硬毛鼠",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Atk:     500,
		Def:     2000,
		IsValid: true,
	})

	/* 313 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2429
		 调整:

		 中文名: 密林的黑龙王
		 卡片种类: 通常怪兽
		 卡片密码: 89832901
		 种族: 龙
		 属性: 地
		 星级: 6
		 攻击力: 2100
		 防御力: 1800
		 罕见度: 平卡N
		 卡包: VOL07
		 效果: 描述：栖息在密林深处的漆黑龙王。咯吱咯吱的只知道啃食树木。

		*/
		Id:       2429,
		Password: "89832901",
		Name:     "密林的黑龙王",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   6,
		La:      ygo.LA_Earth,  // 地
		Lr:      ygo.LR_Dragon, // 龙
		Atk:     2100,
		Def:     1800,
		IsValid: true,
	})

	/* 314 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2430
		 调整:

		 中文名: 巨大怪鸟
		 卡片种类: 通常怪兽
		 卡片密码: 35712107
		 种族: 鸟兽
		 属性: 风
		 星级: 6
		 攻击力: 2000
		 防御力: 1900
		 罕见度: 平卡N
		 卡包: VOL07
		 效果: 描述：很少见的大鸟。从空中急速下除袭击。

		*/
		Id:       2430,
		Password: "35712107",
		Name:     "巨大怪鸟",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   6,
		La:      ygo.LA_Wind,        // 风
		Lr:      ygo.LR_WingedBeast, // 鸟兽
		Atk:     2000,
		Def:     1900,
		IsValid: true,
	})

	/* 315 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 246
		 调整:

		 中文名: 水之舞女
		 卡片种类: 通常怪兽
		 卡片密码: 02483611
		 种族: 水
		 属性: 水
		 星级: 4
		 攻击力: 1400
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: RB、VOL04、DL02
		 效果: 描述：从瓶中不断流淌的水，变成龙的形状发动攻击。

		*/
		Id:       246,
		Password: "02483611",
		Name:     "水之舞女",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     1400,
		Def:     1200,
		IsValid: true,
	})

	/* 316 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 247
		 调整:

		 中文名: 陆战型战斗艇
		 卡片种类: 通常怪兽
		 卡片密码: 58314394
		 种族: 机械
		 属性: 地
		 星级: 4
		 攻击力: 1500
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: RB、VOL04、DL02
		 效果: 描述：陆地战斗机器人，总有一天能在海里使用。

		*/
		Id:       247,
		Password: "58314394",
		Name:     "陆战型战斗艇",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Machine, // 机械
		Atk:     1500,
		Def:     1000,
		IsValid: true,
	})

	/* 317 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 248
		 调整:

		 中文名: 幼虫宝宝
		 卡片种类: 通常怪兽
		 卡片密码: 58192742
		 种族: 昆虫
		 属性: 地
		 星级: 1
		 攻击力: 300
		 防御力: 200
		 罕见度: 平卡N
		 卡包: RB、BE01、VOL04、DL02
		 效果: 描述：不清楚成长后会变成什么样的小青虫。

		*/
		Id:       248,
		Password: "58192742",
		Name:     "幼虫宝宝",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   1,
		La:      ygo.LA_Earth,  // 地
		Lr:      ygo.LR_Insect, // 昆虫
		Atk:     300,
		Def:     200,
		IsValid: true,
	})

	/* 318 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 257
		 调整:

		 中文名: 狮子王
		 卡片种类: 通常怪兽
		 卡片密码: 10538007
		 种族: 兽
		 属性: 地
		 星级: 5
		 攻击力: 1750
		 防御力: 1550
		 罕见度: 平卡N
		 卡包: RB、VOL05、DL02
		 效果: 描述：浑身百兽之王般漂亮的鬃毛，身体也很庞大。

		*/
		Id:       257,
		Password: "10538007",
		Name:     "狮子王",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Atk:     1750,
		Def:     1550,
		IsValid: true,
	})

	/* 319 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 261
		 调整:

		 中文名: 破坏的石人
		 卡片种类: 通常怪兽
		 卡片密码: 73481154
		 种族: 岩石
		 属性: 地
		 星级: 4
		 攻击力: 1500
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: RB、EX-R(EX)、VOL04、DL02
		 效果: 描述：巨大右臂为特征的石人，用右臂进行压迫攻击。

		*/
		Id:       261,
		Password: "73481154",
		Name:     "破坏的石人",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Rock,  // 岩石
		Atk:     1500,
		Def:     1000,
		IsValid: true,
	})

	/* 320 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 263
		 调整:

		 中文名: 苍白兽
		 卡片种类: 通常怪兽
		 卡片密码: 21263083
		 种族: 兽
		 属性: 地
		 星级: 4
		 攻击力: 1500
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: RB、EX-R(EX)、VOL05、DL02
		 效果: 描述：青白色的皮肤，让人感觉恶心的不明怪兽。

		*/
		Id:       263,
		Password: "21263083",
		Name:     "苍白兽",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Atk:     1500,
		Def:     1200,
		IsValid: true,
	})

	/* 321 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 265
		 调整:

		 中文名: 古代的蜥蜴战士
		 卡片种类: 通常怪兽
		 卡片密码: 43230671
		 种族: 爬虫类
		 属性: 地
		 星级: 4
		 攻击力: 1400
		 防御力: 1100
		 罕见度: 平卡N
		 卡包: RB、VOL04、DL02
		 效果: 描述：远古的保持着昔日姿态的蜥蜴战士，意外的强悍。

		*/
		Id:       265,
		Password: "43230671",
		Name:     "古代的蜥蜴战士",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Reptile, // 爬虫类
		Atk:     1400,
		Def:     1100,
		IsValid: true,
	})

	/* 322 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 267
		 调整:

		 中文名: 兵主部
		 卡片种类: 通常怪兽
		 卡片密码: 02118022
		 种族: 水
		 属性: 水
		 星级: 4
		 攻击力: 1500
		 防御力: 900
		 罕见度: 平卡N
		 卡包: RB、VOL05、DL02
		 效果: 描述：河童的首领，攻击力意外的高，守备力偏低。

		*/
		Id:       267,
		Password: "02118022",
		Name:     "兵主部",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     1500,
		Def:     900,
		IsValid: true,
	})

	/* 323 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 408
		 调整:

		 中文名: 宝贝龙
		 卡片种类: 通常怪兽
		 卡片密码: 88819587
		 种族: 龙
		 属性: 风
		 星级: 3
		 攻击力: 1200
		 防御力: 700
		 罕见度: 平卡N
		 卡包: ME、BE02、VOL06、DL04、JY、SJ2
		 效果: 描述：不可轻视的幼龙，说不定隐藏着什么力量。

		*/
		Id:       408,
		Password: "88819587",
		Name:     "宝贝龙",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Wind,   // 风
		Lr:      ygo.LR_Dragon, // 龙
		Atk:     1200,
		Def:     700,
		IsValid: true,
	})

	/* 324 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 409
		 调整:

		 中文名: 暗黑之龙王
		 卡片种类: 通常怪兽
		 卡片密码: 87564352
		 种族: 龙
		 属性: 暗
		 星级: 4
		 攻击力: 1500
		 防御力: 800
		 罕见度: 平卡N
		 卡包: ME、BE02、VOL06、DL04
		 效果: 描述：生活在黑暗深处的龙，由于长期出活在暗处。视力不太好。

		*/
		Id:       409,
		Password: "87564352",
		Name:     "暗黑之龙王",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Dark,   // 暗
		Lr:      ygo.LR_Dragon, // 龙
		Atk:     1500,
		Def:     800,
		IsValid: true,
	})

	/* 325 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 411
		 调整:

		 中文名: 牛魔人
		 卡片种类: 通常怪兽
		 卡片密码: 18246479
		 种族: 兽战士
		 属性: 地
		 星级: 5
		 攻击力: 1800
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: ME、BE02、VOL06、DL04
		 效果: 描述：住在森林里的牛魔人，用突起的角进行突袭。

		*/
		Id:       411,
		Password: "18246479",
		Name:     "牛魔人",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Earth,        // 地
		Lr:      ygo.LR_BeastWarrior, // 兽战士
		Atk:     1800,
		Def:     1300,
		IsValid: true,
	})

	/* 326 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 413
		 调整:

		 中文名: 暗道化师 扎奇
		 卡片种类: 通常怪兽
		 卡片密码: 66602787
		 种族: 魔法师
		 属性: 暗
		 星级: 3
		 攻击力: 600
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: ME、BE02、VOL07、DL04、DPKB、KA、SK2
		 效果: 描述：来去无踪的道化师，以不可思议的动作躲避攻击。

		*/
		Id:       413,
		Password: "66602787",
		Name:     "暗道化师 扎奇",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Dark,        // 暗
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     600,
		Def:     1500,
		IsValid: true,
	})

	/* 327 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 415
		 调整:

		 中文名: 无脸幻想师
		 卡片种类: 通常怪兽
		 卡片密码: 28546905
		 种族: 魔法师
		 属性: 暗
		 星级: 5
		 攻击力: 1200
		 防御力: 2200
		 罕见度: 平卡N
		 卡包: ME、BE02、VOL07、DL04、PE
		 效果: 描述：制造幻影，能轻易地躲避敌人的攻击。

		*/
		Id:       415,
		Password: "28546905",
		Name:     "无脸幻想师",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Dark,        // 暗
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     1200,
		Def:     2200,
		IsValid: true,
	})

	/* 328 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 424
		 调整:

		 中文名: 人马兽
		 卡片种类: 通常怪兽
		 卡片密码: 68516705
		 种族: 兽
		 属性: 地
		 星级: 4
		 攻击力: 1300
		 防御力: 1550
		 罕见度: 平卡N
		 卡包: ME、BE02、EX-R(EX)、VOL07、DL04
		 效果: 描述：人与马化身的怪兽，奔跑的速度谁也追不上。

		*/
		Id:       424,
		Password: "68516705",
		Name:     "人马兽",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Atk:     1300,
		Def:     1550,
		IsValid: true,
	})

	/* 329 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 436
		 调整:

		 中文名: TM-1火箭炮蜘蛛
		 卡片种类: 通常怪兽
		 卡片密码: 87322377
		 种族: 机械
		 属性: 炎
		 星级: 7
		 攻击力: 2200
		 防御力: 2500
		 罕见度: 平卡N，金字UR
		 卡包: ME、BE02、LE02、VOL07、DL04
		 效果: 描述：火箭发射器乱射，将敌人爆杀的机器蜘蛛。

		*/
		Id:       436,
		Password: "87322377",
		Name:     "TM-1火箭炮蜘蛛",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   7,
		La:      ygo.LA_Fire,    // 炎
		Lr:      ygo.LR_Machine, // 机械
		Atk:     2200,
		Def:     2500,
		IsValid: true,
	})

	/* 330 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 437
		 调整:

		 中文名: 高科技狼
		 卡片种类: 通常怪兽
		 卡片密码: 08471389
		 种族: 机械
		 属性: 炎
		 星级: 4
		 攻击力: 1200
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: ME、VOL06、DL04、TP11
		 效果: 描述：全身由钢铁制成的狼，用锋利的牙齿撕咬敌人。

		*/
		Id:       437,
		Password: "08471389",
		Name:     "高科技狼",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Fire,    // 炎
		Lr:      ygo.LR_Machine, // 机械
		Atk:     1200,
		Def:     1400,
		IsValid: true,
	})

	/* 331 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 439
		 调整:

		 中文名: 彩虹鱼
		 卡片种类: 通常怪兽
		 卡片密码: 23771716
		 种族: 鱼
		 属性: 水
		 星级: 4
		 攻击力: 1800
		 防御力: 800
		 罕见度: 平卡N
		 卡包: ME、VOL07、DL04、SD04、GLD01
		 效果: 描述：非常珍奇的七色鱼，要捕捉它非常困难。

		*/
		Id:       439,
		Password: "23771716",
		Name:     "彩虹鱼",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Fish,  // 鱼
		Atk:     1800,
		Def:     800,
		IsValid: true,
	})

	/* 332 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 457
		 调整:

		 中文名: 高等女祭司
		 卡片种类: 通常怪兽
		 卡片密码: 17358176
		 种族: 魔法师
		 属性: 光
		 星级: 3
		 攻击力: 1100
		 防御力: 800
		 罕见度: 平卡N
		 卡包: ME、VOL06、DL04、TP10
		 效果: 描述：咏唱起闻所未闻的咒语，使所有人的心平静下来。

		*/
		Id:       457,
		Password: "17358176",
		Name:     "高等女祭司",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Light,       // 光
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     1100,
		Def:     800,
		IsValid: true,
	})

	/* 333 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 735
		 调整:

		 中文名: 真红眼黑龙
		 卡片种类: 通常怪兽
		 卡片密码: 74677422
		 种族: 龙
		 属性: 暗
		 星级: 7
		 攻击力: 2400
		 防御力: 2000
		 罕见度: 平卡N，银字R，面闪SR，金字UR，爆闪PR，立体UTR
		 卡包: PG、BE01、PP05、VOL03、DL02、301、SD01、DT01、YAP01、DTP01、JY、SJ2、LC01、SD22
		 效果: 描述：拥有真红之眼的黑龙。愤怒的黑炎会把映入其眼者全部烧成灰烬。

		*/
		Id:       735,
		Password: "74677422",
		Name:     "真红眼黑龙",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   7,
		La:      ygo.LA_Dark,   // 暗
		Lr:      ygo.LR_Dragon, // 龙
		Atk:     2400,
		Def:     2000,
		IsValid: true,
	})
}
