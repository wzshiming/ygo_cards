package ygo_cards

import ygo "github.com/wzshiming/ygo_core"

func d_2(cardBag *ygo.CardVersion) {

	/* 0 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2083
		 调整:

		 [昼夜的大火事]<昼夜の大火事>
		 [2010/09/10]
		 ●给予对方基本分800分伤害。
		 ◇「精灵之镜/精霊の鏡」可以对应这张卡的发动而发动

		 中文名: 昼夜的大火事
		 卡片种类: 通常魔法
		 卡片密码: 19523799
		 罕见度: 平卡N
		 卡包: EX-R(EX)、YSD01、Booster02、Booster R1、YSD04
		 效果: 给予对方基本分800分伤害。

		*/
		Id:       2083,
		Password: "19523799",
		Name:     "昼夜的大火事",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 1 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2501
		 调整:

		 [未熟的密探]<未熟な密偵>
		 [2010/09/10]
		 ●指定对方的1张手卡，可以进行查看。
		 ◇「精灵之镜/精霊の鏡」可以对应这张卡的发动而发动
		 ◇效果处理时随机选择对方手卡中的1张卡，进行查看
		 ◇「仪式之钟/セレモニーベル」「真实之眼/真実の眼」等效果适用中时，效果处理中对方手卡以暂时不公开来进行随机查看

		 中文名: 未熟的密探
		 卡片种类: 通常魔法
		 卡片密码: 81820689
		 罕见度: 平卡N
		 卡包: EX-R(EX)、Booster02、Booster R1
		 效果: 指定对方的1张手卡进行观看。

		*/
		Id:       2501,
		Password: "81820689",
		Name:     "未熟的密探",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 2 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2531
		 调整:

		 [古代的望远镜]<古代の遠眼鏡>
		 [2010/09/10]
		 ●可以从对方的卡组最上方查看5张卡。那之后把卡放回原处。
		 ◇只能查看对方的主卡组
		 ◇查看后不能更改卡片的位置
		 ◇对方不能查看那些卡
		 ◇卡片放回卡组后不需要洗切

		 中文名: 古代的望远镜
		 卡片种类: 通常魔法
		 卡片密码: 17092736
		 罕见度: 平卡N
		 卡包: EX-R(EX)、Booster03、Booster R2
		 效果: 查看对方牌组最上面的5张卡，然后放回原处。

		*/
		Id:       2531,
		Password: "17092736",
		Name:     "古代的望远镜",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 3 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 476
		 调整:

		 [分担痛苦]<痛み分け>
		 [2010/09/10]
		 ●把自己场上的1只怪兽解放发动。对方必须把1只怪兽解放。
		 ◇发动时把自己场上的1只怪兽解放（代价）
		 ◇效果处理时对方选择1只怪兽解放（不取对象）
		 ◇这个效果是强制要求对方把1只怪兽解放，不是以魔法效果解放对方1只怪兽
		 ◇对方场上的「荷鲁斯之黑炎龙
		LV6/ホルスの黒炎竜
		ＬＶ６」可以被解放
		 ◇对方场上只有1只怪兽，自己不能利用「灵魂交错/クロス·ソウル」的效果把对方的怪兽解放来发动（发动时确定效果处理时不适用）

		 中文名: 分担痛苦
		 卡片种类: 通常魔法
		 卡片密码: 56830749
		 罕见度: 平卡N
		 卡包: ME、BE02、DL04、Booster07、DT02
		 效果: 自己场上的1只怪兽作祭品。对方选择对方场上1只怪做祭品。

		*/
		Id:       476,
		Password: "56830749",
		Name:     "分担痛苦",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 4 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 478
		 调整:

		 [大风暴]<大嵐>
		 [2010/09/10]
		 ●在场上存在的魔法·陷阱卡全部破坏。
		 ◇不取对象
		 ◇这张卡自身是在连锁处理结束后送去墓地，不是被破坏送去墓地
		 ◇同时复数张卡被破坏时，那些卡在墓地的放置顺序由控制者决定（破坏是同时破坏）

		 中文名: 大风暴
		 卡片种类: 通常魔法
		 卡片密码: 19613556
		 罕见度: 平卡N，银字R，黄金GR，面闪SR
		 卡包: ME、BE02、YSD01、YSD02、DL04、SD01、SD02、SD03、SD04、SD05、SD06、SD08、SD09、SD10、SD11、SD12、Booster07、GLD01、GS01、KA、PE、SK2、SD26、GDB1
		 效果: 场上存在的魔法·陷阱卡全部破坏。

		*/
		Id:       478,
		Password: "19613556",
		Name:     "大风暴",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 5 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 483
		 调整:

		 [天使的鲜血]<天使の生き血>
		 [2010/09/10]
		 ●自己回复800基本分。
		 ◇「精灵之镜/精霊の鏡」可以对应这张卡的发动而发动

		 中文名: 天使的鲜血
		 卡片种类: 通常魔法
		 卡片密码: 47852924
		 罕见度: 平卡N
		 卡包: BE02、DL04、Booster02、Booster Chronicle、Booster R1
		 效果: 自己回复800点的基本分。

		*/
		Id:       483,
		Password: "47852924",
		Name:     "天使的鲜血",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 6 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 484
		 调整:

		 [消除黑暗的光]<闇をかき消す光>
		 [2010/09/10]
		 ●对方场上里侧表示存在的怪兽全部变为表侧表示。
		 ◇里侧守备变为表侧守备，里侧攻击变为表侧攻击

		 中文名: 消除黑暗的光
		 卡片种类: 通常魔法
		 卡片密码: 45895206
		 罕见度: 平卡N
		 卡包: BE02、DL04、Booster02、Booster Chronicle、Booster R1
		 效果: 对方场上里侧表示的怪兽全部表侧表示。

		*/
		Id:       484,
		Password: "45895206",
		Name:     "消除黑暗的光",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 7 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 487
		 调整:

		 [蓝色药剂]<ブルー·ポーション>
		 [2010/09/10]
		 ●自己回复400基本分。
		 ◇「精灵之镜/精霊の鏡」可以对应这张卡的发动而发动

		 中文名: 蓝色药剂
		 卡片种类: 通常魔法
		 卡片密码: 20871001
		 罕见度: 平卡N
		 卡包: DL04、Booster01、Booster R1、Booster Chronicle
		 效果: 自己回复400点的基本分。

		*/
		Id:       487,
		Password: "20871001",
		Name:     "蓝色药剂",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 8 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 488
		 调整:

		 [雷鸣]<雷鳴>
		 [2010/09/10]
		 ●给予对方基本分300分伤害。
		 ◇「精灵之镜/精霊の鏡」可以对应这张卡的发动而发动
		 ◇这张卡的读音和「雷-鸣/Rai-Mei」相同，所以在使用「心灵崩坏/マインドクラッシュ」等宣言卡名时需要同时说明是宣言魔法卡的「雷鳴」还是效果怪兽的「Rai-Mei」

		 中文名: 雷鸣
		 卡片种类: 通常魔法
		 卡片密码: 56260110
		 罕见度: 平卡N
		 卡包: DL04、Booster01、Booster Chronicle、Booster R1、TP04
		 效果: 给与对方基本分300分伤害。

		*/
		Id:       488,
		Password: "56260110",
		Name:     "雷鸣",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 9 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 504
		 调整:

		 [天使的施舍]<天使の施し>
		 [2010/09/10]
		 ●从卡组抽3张卡，那之后从手卡丢弃2张卡。
		 ◇效果处理时先抽卡，之后再丢弃2张卡
		 ◇卡组中不足3张卡的场合，不能发动
		 ◇「精灵之镜/精霊の鏡」可以对应这张卡的发动而发动
		 ◇「死之卡组破坏病毒/死のデッキ破壊ウイルス」「魔之卡组破坏病毒/魔のデッキ破壊ウイルス」的效果在这个效果处理完丢卡后再适用

		 中文名: 天使的施舍
		 卡片种类: 通常魔法
		 卡片密码: 79571449
		 罕见度: 平卡N，银字R，面闪SR
		 卡包: BE02、DL04、Booster04、Booster Chronicle、Booster R2、YU、KA、SK2
		 效果: 从卡组抽3张卡，之后从手卡选2张丢弃。

		*/
		Id:       504,
		Password: "79571449",
		Name:     "天使的施舍",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 10 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 525
		 调整:

		 [魔女狩猎]<魔女狩り>
		 [2010/09/10]
		 ●在场上表侧表示存在的魔法师族怪兽全部破坏。
		 ◇不取对象

		 中文名: 魔女狩猎
		 卡片种类: 通常魔法
		 卡片密码: 90330453
		 罕见度: 平卡N
		 卡包: DL04、Booster06、Booster Chronicle、Booster R3
		 效果: 场上表侧表示的魔法师族的怪兽全部破坏。

		*/
		Id:       525,
		Password: "90330453",
		Name:     "魔女狩猎",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 11 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 526
		 调整:

		 [恶魔祓除]<悪魔払い>
		 [2010/09/10]
		 ●在场上表侧表示存在的恶魔族怪兽全部破坏。
		 ◇不取对象

		 中文名: 恶魔祓除
		 卡片种类: 通常魔法
		 卡片密码: 26725158
		 罕见度: 平卡N
		 卡包: DL04、Booster06、Booster Chronicle、Booster R3、GLD02
		 效果: 场上表侧表示的恶魔族的怪兽全部破坏。

		*/
		Id:       526,
		Password: "26725158",
		Name:     "恶魔祓除",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 12 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 528
		 调整:

		 [革命]<革命>
		 [2010/09/10]
		 ●给予对方基本分对方的手卡的数量×200分的伤害。
		 ◇「精灵之镜/精霊の鏡」可以对应这张卡的发动而发动
		 ◇对方没有手卡时不能发动
		 ◇效果处理时计算对方的手卡数量

		 中文名: 革命
		 卡片种类: 通常魔法
		 卡片密码: 99518961
		 罕见度: 平卡N
		 卡包: BE02、DL04、Booster06、Booster Chronicle、Booster R3、SY2
		 效果: 对方受到（对方的手卡数×200）的伤害。

		*/
		Id:       528,
		Password: "99518961",
		Name:     "革命",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 13 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 529
		 调整:

		 [融合贤者]<融合賢者>
		 [2010/09/10]
		 ●从自己的卡组把1张「融合」魔法卡加入手卡。那之后卡组洗切。
		 ◇遵守不能空发原则
		 ◇效果处理时卡组内没有符合条件的卡时这个效果不适用，对方可以确认自己的卡组（不取对象）

		 中文名: 融合贤者
		 卡片种类: 通常魔法
		 卡片密码: 26902560
		 罕见度: 平卡N
		 卡包: DP01、BE02、DL04、Booster06、Booster Chronicle、Booster R3、JY、PE、SY2
		 效果: 选1张自己的卡组的「融合」魔法卡加入手卡。之后洗切卡组。

		*/
		Id:       529,
		Password: "26902560",
		Name:     "融合贤者",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 14 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2499
		 调整:

		 [暗之破神剑]<闇の破神剣>
		 [2010/09/10]
		 ●暗属性怪兽的攻击力上升400分！守备力下降200分！
		 ◇发动时选择1只符合条件的怪兽（取对象）
		 ◇效果处理时装备怪兽属性变为暗属性以外的场合，这张卡送去墓地
		 ◇装备状态中对象怪兽变为暗属性以外的场合，这张卡破坏

		 中文名: 暗之破神剑
		 卡片种类: 装备魔法
		 卡片密码: 37120512
		 罕见度: 平卡N
		 卡包: EX-R(EX)、Booster02、Booster R1
		 效果: 暗属性的怪兽才能装备。装备的怪兽攻击力上升400，守备力下降200。

		*/
		Id:       2499,
		Password: "37120512",
		Name:     "暗之破神剑",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 15 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2500
		 调整:

		 [觉醒]<覚醒>
		 [2010/09/10]
		 ●地属性怪兽的攻击力上升400分！守备力下降200分！
		 ◇发动时选择1只符合条件的怪兽（取对象）
		 ◇效果处理时装备怪兽属性变为地属性以外的场合，这张卡送去墓地
		 ◇装备状态中对象怪兽变为地属性以外的场合，这张卡破坏

		 中文名: 觉醒
		 卡片种类: 装备魔法
		 卡片密码: 98374133
		 罕见度: 平卡N
		 卡包: EX-R(EX)、Booster02、Booster R1
		 效果: 地属性的怪兽才能装备。装备的怪兽攻击力上升400，守备力下降200。

		*/
		Id:       2500,
		Password: "98374133",
		Name:     "觉醒",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 16 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 475
		 调整:

		 [磁力指轮]<磁力の指輪>
		 [2010/09/10]
		 ●只能给自己场上的怪兽装备。装备怪兽的攻击力与守备力下降500分。对方只能攻击装备了这张卡的怪兽。
		 ◇发动时选择自己场上1只表侧表示的怪兽（取对象）
		 ◇装备怪兽的控制权转移到对方场上的场合，这些效果不适用（继续装备着）

		 中文名: 磁力指轮
		 卡片种类: 装备魔法
		 卡片密码: 20436034
		 罕见度: 平卡N，银字R
		 卡包: ME、BE02、DL04、Booster07
		 效果: 自己场上的怪兽才可以装备。装备的怪兽攻击力·守备力下降500。对方只能攻击装备了这张卡的怪兽。

		*/
		Id:       475,
		Password: "20436034",
		Name:     "磁力指轮",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 17 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 477
		 调整:

		 [兴奋剂]<ドーピング>
		 [2010/09/10]
		 ●装备怪兽的攻击力上升700分。并且，在每次自己的准备阶段，装备怪兽的攻击力下降200分。
		 ◇发动时选择场上1只表侧表示的怪兽（取对象）
		 ◇准备阶段时攻击力下降效果进入连锁（咒文速度1）

		 中文名: 兴奋剂
		 卡片种类: 装备魔法
		 卡片密码: 83225447
		 罕见度: 平卡N
		 卡包: ME、DL04、Booster07、TP14
		 效果: 装备怪兽的攻击力上升700。自己的每次的准备阶段装备怪兽的攻击力下降200。

		*/
		Id:       477,
		Password: "83225447",
		Name:     "兴奋剂",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 18 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 481
		 调整:

		 [精灵之光]<エルフの光>
		 [2010/09/10]
		 ●光属性怪兽才能装备。1只装备怪兽的攻击力上升400分。守备力下降200分。
		 ◇发动时选择1只符合条件的怪兽（取对象）
		 ◇效果处理时装备怪兽属性变为光属性以外的场合，这张卡送去墓地
		 ◇装备状态中对象怪兽变为光属性以外的场合，这张卡破坏

		 中文名: 精灵之光
		 卡片种类: 装备魔法
		 卡片密码: 39897277
		 罕见度: 平卡N
		 卡包: DL04、Booster02、Booster Chronicle、Booster R1、TP16
		 效果: 光属性的怪兽才能装备。装备的怪兽攻击力上升400，守备力下降200。

		*/
		Id:       481,
		Password: "39897277",
		Name:     "精灵之光",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 19 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 482
		 调整:

		 [钢甲壳]<はがねの甲羅>
		 [2010/09/10]
		 ●水属性怪兽才能装备。1只装备怪兽的攻击力上升400分。守备力下降200分。
		 ◇发动时选择1只符合条件的怪兽（取对象）
		 ◇效果处理时装备怪兽属性变为水属性以外的场合，这张卡送去墓地
		 ◇装备状态中对象怪兽变为水属性以外的场合，这张卡破坏

		 中文名: 钢甲壳
		 卡片种类: 装备魔法
		 卡片密码: 02370081
		 罕见度: 平卡N
		 卡包: DL04、Booster02、Booster Chronicle、Booster R1
		 效果: 水属性的怪兽才能装备。装备的怪兽攻击力上升400，守备力下降200。

		*/
		Id:       482,
		Password: "02370081",
		Name:     "钢甲壳",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 20 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 489
		 调整:

		 [灼热之枪]<灼熱の槍>
		 [2010/09/10]
		 ●炎属性怪兽才能装备。1只装备怪兽的攻击力上升400分，守备力下降200分。
		 ◇发动时选择1只符合条件的怪兽（取对象）
		 ◇效果处理时装备怪兽属性变为炎属性以外的场合，这张卡送去墓地
		 ◇装备状态中对象怪兽变为炎属性以外的场合，这张卡破坏

		 中文名: 灼热之枪
		 卡片种类: 装备魔法
		 卡片密码: 18937875
		 罕见度: 平卡N
		 卡包: DL04、Booster02、Booster Chronicle、Booster R1、TP13
		 效果: 炎属性的怪兽才能装备。装备的怪兽攻击力上升400，守备力下降200。

		*/
		Id:       489,
		Password: "18937875",
		Name:     "灼热之枪",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 21 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 490
		 调整:

		 [突风之扇]<突風の扇>
		 [2010/09/10]
		 ●风属性怪兽才能装备。1只装备怪兽的攻击力上升400分。守备力下降200分。
		 ◇发动时选择1只符合条件的怪兽（取对象）
		 ◇效果处理时装备怪兽属性变为风属性以外的场合，这张卡送去墓地
		 ◇装备状态中对象怪兽变为风属性以外的场合，这张卡破坏

		 中文名: 突风之扇
		 卡片种类: 装备魔法
		 卡片密码: 55321970
		 罕见度: 平卡N
		 卡包: DL04、Booster02、Booster Chronicle、Booster R1
		 效果: 风属性的怪兽才能装备。装备的怪兽攻击力上升400，守备力下降200。

		*/
		Id:       490,
		Password: "55321970",
		Name:     "突风之扇",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 22 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1358
		 调整:

		 [血之代偿]<血の代偿>
		 [11/02/09]
		 ●可以支付500基本分，把1只怪兽通常召唤。这个效果只能在自己的主要阶段及对方的战斗阶段时可以发动。
		 ◇效果发动时支付500基本分（代价）
		 ◇效果的发动进入连锁（咒文速度2）
		 ◇在卡的发动时可以同时发动效果，那个场合支付500基本分（卡的发动没有条件限制）
		 ◇效果处理时选择自己手卡的1只可以被通常召唤的怪兽进行通常召唤（不取对象）
		 ◇自己手卡没有能被通常召唤的怪兽的场合，不能发动这个效果★效果处理时，对方没有可以召唤的怪兽的场合，能否确认对方手卡？调整中
		 ◇这个通常召唤遵循普通的通常召唤规则（上级召唤在效果处理时需要解放怪兽）
		 ◇这张卡的发动的那组连锁中，这张卡不能再在那组连锁中发动效果（卡的发动尚处理完毕，所以不能再发动效果）★用这张卡的效果在连锁1进行通常召唤的场合，是否可以对应发动「反冲/キックバック」「升天之角笛/昇天の角笛」
		调整中
		 ◇手卡中有多少能被通常召唤的怪兽，就限定了能在一组连锁中发动的次数
		 ◇「王宫的通告/王宫のお触れ」发动中，可以反复发动这张卡来支付基本分
		 ◇可以让二重怪兽进行再度召唤
		 ◇手卡不存在怪兽的场合，可以让二重怪兽进行再度召唤★场上没卡，手卡只有1只二重怪兽，是否可以假设（连锁2）召唤二重怪兽，（连锁1）进行再度召唤来在同一组连锁中发动两次这张卡的效果
		调整中
		 ◇假设手卡只有「栗子球/クリボー」「恶魔召唤/デーモンの召唤」，自己场上只有这张卡，不能以假设（连锁2）上级召唤「恶魔召唤/デーモンの召唤」（以连锁1的「栗子球/クリボー」作为解放），（连锁1）召唤「栗子球/クリボー」
		 ◇与1回合1次正常的通常召唤没有冲突

		 中文名: 血之代偿
		 卡片种类: 永续陷阱
		 卡片密码: 80604091
		 罕见度: 平卡N，银字R
		 卡包: EX-R(EX)、Booster03、SD07、SD10、Booster R2、DT06、YU、JY、SJ2、DS14
		 效果: 可以支付500基本分，把1只怪兽通常召唤。这个效果在自己回合的主要阶段及对方回合的战斗阶段才能发动。

		*/
		Id:       1358,
		Password: "80604091",
		Name:     "血之代偿",
		Lc:       ygo.LC_TrapContinuous, // 永续陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 23 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 524
		 调整:

		 [王宫的通告]<王宮のお触れ>
		 [2010/09/10]
		 ●只要这张卡在场上表侧表示存在，这张卡以外的场上的陷阱卡的效果无效。
		 ◇只无效效果，不无效发动（不进入连锁）
		 ◇只要是在场上发动的陷阱卡，不论效果处理时在哪里，效果都无效化
		 ◇在场上以外发动的陷阱效果不会被无效

		 中文名: 王宫的通告
		 卡片种类: 永续陷阱
		 卡片密码: 51452091
		 罕见度: 平卡N，银字R，黄金GR，面闪SR
		 卡包: BE02、DL04、SD05、Booster05、Booster Chronicle、Booster R3、GS01、SY2、GDB1
		 效果: 只要这张卡在场上表侧表示存在，这张卡以外的场上的陷阱卡的效果无效。

		*/
		Id:       524,
		Password: "51452091",
		Name:     "王宫的通告",
		Lc:       ygo.LC_TrapContinuous, // 永续陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 24 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 527
		 调整:

		 [魔力之棘]<魔力の棘>
		 [2010/09/10]
		 ●对方的手卡被丢弃去墓地时，每丢弃1张卡给予对方500分伤害。
		 ◇伤害效果不进入连锁
		 ◇这张卡在场上有2张以上效果适用的场合，效果叠加
		 ◇原本自己的卡从对方手卡被丢弃、原本对方的卡从自己手卡被丢弃这个效果都不适用
		 ◇一次丢弃多张的场合，伤害数值为那个被丢弃数量×500
		 ◇从手卡送去墓地、被破坏都不在效果适用范围内
		 ◇对方场上「魔力之棘/魔力の棘」的效果适用中，我方回合自己发动「天使的施舍/天使の施し」从卡组抽3张卡，丢弃2张卡。「天使的施舍/天使の施し」效果处理完毕后自己凑齐了「艾克佐迪亚/エクゾディア」5张部件满足胜利条件，而由于「魔力之棘/魔力の棘」的效果自己受到1000分伤害基本分变为0对方也满足胜利条件的场合，以回合玩家优先来处理。（即实际先适用「被封印的艾克佐迪亚/封印されしエクゾディア」的效果自己获得决斗胜利）

		 中文名: 魔力之棘
		 卡片种类: 永续陷阱
		 卡片密码: 53119267
		 罕见度: 平卡N，银字R
		 卡包: BE02、DL04、Booster06、Booster Chronicle、Booster R3
		 效果: 对方的手卡丢弃去墓地时，每丢弃1张卡就给予对方基本分500分的伤害。

		*/
		Id:       527,
		Password: "53119267",
		Name:     "魔力之棘",
		Lc:       ygo.LC_TrapContinuous, // 永续陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 25 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1420
		 调整:

		 [自业自得]<自業自得>
		 [2010/09/10]
		 ●对方场上每存在1只怪兽，给予对方基本分500分伤害。
		 ◇对方场上不存在怪兽的场合，不能发动
		 ◇在效果处理时计算对方场上怪兽的数量

		 中文名: 自业自得
		 卡片种类: 通常陷阱
		 卡片密码: 24068492
		 罕见度: 平卡N
		 卡包: EX-R(EX)、Booster05、SD12、Booster R3
		 效果: 对方场上每存在1只怪兽对方受到500基本分伤害。

		*/
		Id:       1420,
		Password: "24068492",
		Name:     "自业自得",
		Lc:       ygo.LC_TrapNormal, // 通常陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 26 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2081
		 调整:

		 [城壁]<城壁>
		 [2010/09/10]
		 ●选择场上表侧表示存在的1只怪兽发动。被选择怪兽的守备力直到结束阶段时为止上升500分。
		 ◇发动时选择场上1只表侧表示存在的怪兽（取对象）
		 ◇可以在伤害步骤发动（具体请参考置顶内的伤害步骤详解）

		 中文名: 城壁
		 卡片种类: 通常陷阱
		 卡片密码: 44209392
		 罕见度: 平卡N
		 卡包: EX-R(EX)、YSD01、Booster03、Booster R2、TP11
		 效果: 表侧表示的1只怪兽守备力在回合结束前上升500。

		*/
		Id:       2081,
		Password: "44209392",
		Name:     "城壁",
		Lc:       ygo.LC_TrapNormal, // 通常陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 27 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2082
		 调整:

		 [援军]<援軍>
		 [2010/09/10]
		 ●直到结束阶段时为止，1只表侧表示存在的怪兽的攻击力上升500分。
		 ◇发动时选择场上1只表侧表示存在的怪兽（取对象）
		 ◇可以在伤害步骤发动（具体请参考置顶内的伤害步骤详解）

		 中文名: 援军
		 卡片种类: 通常陷阱
		 卡片密码: 17814387
		 罕见度: 平卡N，黄金GR
		 卡包: EX-R(EX)、YSD01、Booster03、Booster R2、DT01、GLD01、DTP01
		 效果: 表侧表示的1只怪兽攻击力在回合结束前上升500。

		*/
		Id:       2082,
		Password: "17814387",
		Name:     "援军",
		Lc:       ygo.LC_TrapNormal, // 通常陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 28 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2532
		 调整:

		 [天邪鬼的诅咒]<あまのじゃくの呪い>
		 [2010/09/10]
		 ●直到发动回合的结束阶段时为止，攻击力·守备力的上升·下降的效果变成相反。
		 ◇不能在伤害步骤中发动
		 ◇攻击力变成XXX不在效果适用范围内
		 ◇「冥王的咆哮/冥王の咆哮」的「下降」在这个效果适用范围内
		 ◇「突进/突進」「逆卷的炎之精灵/逆巻く炎の精霊」等已经适用的效果也受到这个效果影响

		 中文名: 天邪鬼的诅咒
		 卡片种类: 通常陷阱
		 卡片密码: 77622396
		 罕见度: 平卡N，银字R
		 卡包: EX-R(EX)、Booster03、Booster R2、TP01
		 效果: 直到发动回合的结束阶段时，攻击力·守备力的上升·下降的效果变成相反。

		*/
		Id:       2532,
		Password: "77622396",
		Name:     "天邪鬼的诅咒",
		Lc:       ygo.LC_TrapNormal, // 通常陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 29 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2556
		 调整:

		 [白洞]<ホワイト·ホール>
		 [2010/09/10]
		 ●对方使用「黑洞」时可以发动。自己的怪兽不会被那张「黑洞」的效果破坏。
		 ◇必须直接对应「黑洞/ブラック·ホール」的发动进行连锁发动

		 中文名: 白洞
		 卡片种类: 通常陷阱
		 卡片密码: 43487744
		 罕见度: 平卡N
		 卡包: Booster04、Booster Chronicle、Booster R2、TP15、TU04
		 效果: 对方使用了「黑洞」的时候，自己的怪兽不被破坏。发动后这张卡破坏。

		*/
		Id:       2556,
		Password: "43487744",
		Name:     "白洞",
		Lc:       ygo.LC_TrapNormal, // 通常陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 30 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2616
		 调整:

		 [狮鹫之翼]<グリフォンの翼>
		 [2010/09/10]
		 ●对方使用「鹰身女妖的羽毛扫」时，代替自己的魔法·陷阱卡对方的魔法·陷阱卡全部破坏。
		 ◇必须直接对应「鹰身女妖的羽毛扫/ハーピィの羽根帚」的发动而发动。
		 ◇「鹰身女妖的羽毛扫/ハーピィの羽根帚」效果处理时自己场上的魔法·陷阱卡不破坏，对方场上的魔法·陷阱卡代替我方的破坏。

		 中文名: 狮鹫之翼
		 卡片种类: 通常陷阱
		 卡片密码: 55608151
		 罕见度: 平卡N
		 卡包: ME、Booster07
		 效果: 当对方使用「鹰身女妖的羽毛扫」时，对方的魔法和陷阱卡，将代替自己的被全部破坏。

		*/
		Id:       2616,
		Password: "55608151",
		Name:     "狮鹫之翼",
		Lc:       ygo.LC_TrapNormal, // 通常陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 31 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 505
		 调整:

		 [来自墓场的呼声]<墓場からの呼び声>
		 [2010/09/10]
		 ●对方发动「死者苏生」时可以发动。那张「死者苏生」的效果无效。
		 ◇必须直接对应「死者苏生/死者蘇生」的发动进行连锁发动
		 ◇只无效效果，不无效发动

		 中文名: 来自墓场的呼声
		 卡片种类: 通常陷阱
		 卡片密码: 16970158
		 罕见度: 平卡N
		 卡包: DL04、Booster04、TP05
		 效果: 对方「死者苏生」发动时，发动效果无效。

		*/
		Id:       505,
		Password: "16970158",
		Name:     "来自墓场的呼声",
		Lc:       ygo.LC_TrapNormal, // 通常陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 32 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2527
		 调整:

		 中文名: 朱雀
		 卡片种类: 融合怪兽
		 卡片密码: 35752363
		 种族: 炎
		 属性: 炎
		 星级: 5
		 攻击力: 1900
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: Booster03
		 效果: 融合：「赤剑之莱蒙多斯」＋「炎之魔神」。

		*/
		Id:       2527,
		Password: "35752363",
		Name:     "朱雀",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 5,
		La:    ygo.LA_Fire, // 炎
		Lr:    ygo.LR_Fire, // 炎
		Atk:   1900,
		Def:   1500,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 33 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2528
		 调整:

		 中文名: 战场的死装束
		 卡片种类: 融合怪兽
		 卡片密码: 56413937
		 种族: 战士
		 属性: 地
		 星级: 6
		 攻击力: 1900
		 防御力: 1700
		 罕见度: 平卡N
		 卡包: Booster03、Booster Chronicle、Booster R2
		 效果: 融合：「音女」＋「斩首的美女」。

		*/
		Id:       2528,
		Password: "56413937",
		Name:     "战场的死装束",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 6,
		La:    ygo.LA_Earth,   // 地
		Lr:    ygo.LR_Warrior, // 战士
		Atk:   1900,
		Def:   1700,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 34 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2529
		 调整:

		 中文名: 死亡鸟
		 卡片种类: 融合怪兽
		 卡片密码: 08327462
		 种族: 鸟兽
		 属性: 风
		 星级: 6
		 攻击力: 1900
		 防御力: 1700
		 罕见度: 平卡N
		 卡包: Booster03
		 效果: 融合：「橐蜚」＋「骷髅寺院」。

		*/
		Id:       2529,
		Password: "08327462",
		Name:     "死亡鸟",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 6,
		La:    ygo.LA_Wind,        // 风
		Lr:    ygo.LR_WingedBeast, // 鸟兽
		Atk:   1900,
		Def:   1700,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 35 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2554
		 调整:

		 中文名: 水精龙
		 卡片种类: 融合怪兽
		 卡片密码: 86164529
		 种族: 海龙
		 属性: 水
		 星级: 6
		 攻击力: 2250
		 防御力: 1900
		 罕见度: 平卡N
		 卡包: Booster04、TP05
		 效果: 融合：「妖精龙」＋「海原的女战士」＋「区域吞噬者」。

		*/
		Id:       2554,
		Password: "86164529",
		Name:     "水精龙",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 6,
		La:    ygo.LA_Water,      // 水
		Lr:    ygo.LR_SeaSerpent, // 海龙
		Atk:   2250,
		Def:   1900,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 36 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2555
		 调整:

		 中文名: 黑色食人鲨
		 卡片种类: 融合怪兽
		 卡片密码: 80727036
		 种族: 鱼
		 属性: 水
		 星级: 5
		 攻击力: 2100
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: Booster04、TP04
		 效果: 融合：「深海切割手」＋「杀人污泥」＋「海原的女战士」。

		*/
		Id:       2555,
		Password: "80727036",
		Name:     "黑色食人鲨",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 5,
		La:    ygo.LA_Water, // 水
		Lr:    ygo.LR_Fish,  // 鱼
		Atk:   2100,
		Def:   1300,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 37 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2576
		 调整:

		 中文名: 腕龙
		 卡片种类: 融合怪兽
		 卡片密码: 16507828
		 种族: 恐龙
		 属性: 水
		 星级: 6
		 攻击力: 2200
		 防御力: 2000
		 罕见度: 平卡N
		 卡包: Booster05、Booster Chronicle、Booster R3
		 效果: 融合：「双头恐龙王」＋「贪尸龙」。

		*/
		Id:       2576,
		Password: "16507828",
		Name:     "腕龙",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 6,
		La:    ygo.LA_Water,    // 水
		Lr:    ygo.LR_Dinosaur, // 恐龙
		Atk:   2200,
		Def:   2000,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 38 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2577
		 调整:

		 中文名: 金色的魔象
		 卡片种类: 融合怪兽
		 卡片密码: 54622031
		 种族: 不死
		 属性: 暗
		 星级: 6
		 攻击力: 2200
		 防御力: 1800
		 罕见度: 平卡N
		 卡包: Booster05、Booster R3
		 效果: 融合：「美杜莎的亡灵」＋「龙僵尸」。

		*/
		Id:       2577,
		Password: "54622031",
		Name:     "金色的魔象",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 6,
		La:    ygo.LA_Dark,   // 暗
		Lr:    ygo.LR_Zombie, // 不死
		Atk:   2200,
		Def:   1800,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 39 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2578
		 调整:

		 中文名: 狩魂魔人
		 卡片种类: 融合怪兽
		 卡片密码: 72869010
		 种族: 恶魔
		 属性: 暗
		 星级: 6
		 攻击力: 2200
		 防御力: 1800
		 罕见度: 平卡N
		 卡包: Booster05、Booster Chronicle、Booster R3
		 效果: 融合：「神灯魔人」＋「来自异次元的侵略者」。

		*/
		Id:       2578,
		Password: "72869010",
		Name:     "狩魂魔人",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 6,
		La:    ygo.LA_Dark,  // 暗
		Lr:    ygo.LR_Fiend, // 恶魔
		Atk:   2200,
		Def:   1800,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 40 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2579
		 调整:

		 中文名: 海兽鱼
		 卡片种类: 融合怪兽
		 卡片密码: 29929832
		 种族: 鱼
		 属性: 水
		 星级: 5
		 攻击力: 1700
		 防御力: 1600
		 罕见度: 平卡N
		 卡包: Booster05
		 效果: 融合：「水之魔导师」＋「大肚海蛇」。

		*/
		Id:       2579,
		Password: "29929832",
		Name:     "海兽鱼",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 5,
		La:    ygo.LA_Water, // 水
		Lr:    ygo.LR_Fish,  // 鱼
		Atk:   1700,
		Def:   1600,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 41 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2596
		 调整:

		 中文名: 帝王龙
		 卡片种类: 融合怪兽
		 卡片密码: 94566432
		 种族: 龙
		 属性: 光
		 星级: 7
		 攻击力: 2300
		 防御力: 2000
		 罕见度: 平卡N
		 卡包: Booster06、Booster Chronicle、Booster R3
		 效果: 融合：「守城的翼龙」＋「妖精龙」。

		*/
		Id:       2596,
		Password: "94566432",
		Name:     "帝王龙",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 7,
		La:    ygo.LA_Light,  // 光
		Lr:    ygo.LR_Dragon, // 龙
		Atk:   2300,
		Def:   2000,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 42 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2597
		 调整:

		 中文名: 红阳鸟
		 卡片种类: 融合怪兽
		 卡片密码: 46696593
		 种族: 鸟兽
		 属性: 炎
		 星级: 6
		 攻击力: 2300
		 防御力: 1800
		 罕见度: 平卡N
		 卡包: Booster06、Booster Chronicle、Booster R3
		 效果: 融合：「圣鸟」＋「天空猎手」。

		*/
		Id:       2597,
		Password: "46696593",
		Name:     "红阳鸟",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 6,
		La:    ygo.LA_Fire,        // 炎
		Lr:    ygo.LR_WingedBeast, // 鸟兽
		Atk:   2300,
		Def:   1800,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 43 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2598
		 调整:

		 中文名: 沙之魔女
		 卡片种类: 融合怪兽
		 卡片密码: 32751480
		 种族: 岩石
		 属性: 地
		 星级: 6
		 攻击力: 2100
		 防御力: 1700
		 罕见度: 平卡N
		 卡包: Booster06、Booster R3
		 效果: 融合：「岩石巨兵」＋「古代精灵」。

		*/
		Id:       2598,
		Password: "32751480",
		Name:     "沙之魔女",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 6,
		La:    ygo.LA_Earth, // 地
		Lr:    ygo.LR_Rock,  // 岩石
		Atk:   2100,
		Def:   1700,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 44 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2599
		 调整:

		 中文名: 骸骨龙
		 卡片种类: 融合怪兽
		 卡片密码: 32355828
		 种族: 不死
		 属性: 暗
		 星级: 6
		 攻击力: 1700
		 防御力: 1900
		 罕见度: 平卡N
		 卡包: Booster06、Booster R3
		 效果: 融合：「美杜莎的亡灵」＋「暗黑之龙王」。

		*/
		Id:       2599,
		Password: "32355828",
		Name:     "骸骨龙",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 6,
		La:    ygo.LA_Dark,   // 暗
		Lr:    ygo.LR_Zombie, // 不死
		Atk:   1700,
		Def:   1900,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 45 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2600
		 调整:

		 中文名: 锹甲独角仙
		 卡片种类: 融合怪兽
		 卡片密码: 95144193
		 种族: 昆虫
		 属性: 地
		 星级: 6
		 攻击力: 1900
		 防御力: 1700
		 罕见度: 平卡N
		 卡包: Booster03、Booster Chronicle、Booster R2、TP20
		 效果: 融合：「锹甲阿尔法」＋「大力独角仙」。

		*/
		Id:       2600,
		Password: "95144193",
		Name:     "锹甲独角仙",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 6,
		La:    ygo.LA_Earth,  // 地
		Lr:    ygo.LR_Insect, // 昆虫
		Atk:   1900,
		Def:   1700,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 46 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2617
		 调整:

		 中文名: 骷髅祭司
		 卡片种类: 融合怪兽
		 卡片密码: 02504891
		 种族: 魔法师
		 属性: 暗
		 星级: 7
		 攻击力: 2650
		 防御力: 2250
		 罕见度: 平卡N
		 卡包: ME、Booster07
		 效果: 融合：「恶魔的智慧」＋「魔天老」。

		*/
		Id:       2617,
		Password: "02504891",
		Name:     "骷髅祭司",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 7,
		La:    ygo.LA_Dark,        // 暗
		Lr:    ygo.LR_Spellcaster, // 魔法师
		Atk:   2650,
		Def:   2250,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 47 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2618
		 调整:

		 中文名: 恶魔箱
		 卡片种类: 融合怪兽
		 卡片密码: 25655502
		 种族: 恶魔
		 属性: 暗
		 星级: 7
		 攻击力: 2300
		 防御力: 2000
		 罕见度: 平卡N
		 卡包: ME、Booster07
		 效果: 融合：「杀人小丑」＋「梦幻小丑」。

		*/
		Id:       2618,
		Password: "25655502",
		Name:     "恶魔箱",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 7,
		La:    ygo.LA_Dark,  // 暗
		Lr:    ygo.LR_Fiend, // 恶魔
		Atk:   2300,
		Def:   2000,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 48 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 522
		 调整:

		 中文名: 水陆两用战斗艇
		 卡片种类: 融合怪兽
		 卡片密码: 40173854
		 种族: 水
		 属性: 水
		 星级: 5
		 攻击力: 1850
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: DL04、Booster06、Booster Chronicle、Booster R3
		 效果: 融合：「陆战型战斗艇」＋「守卫海洋的战士」。

		*/
		Id:       522,
		Password: "40173854",
		Name:     "水陆两用战斗艇",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 5,
		La:    ygo.LA_Water, // 水
		Lr:    ygo.LR_Water, // 水
		Atk:   1850,
		Def:   1300,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 49 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1357
		 调整:

		 [陷阱大师]<トラップ·マスター>
		 [2010/09/10]
		 ●反转：场上的1张陷阱卡破坏。（对象卡里侧的场合，翻开为表侧若是陷阱卡就破坏。不是的场合变回原样。这张卡的效果不发动。）
		 ◇反转效果（进入连锁）
		 ◇强制发动
		 ◇发动时选择场上的1张表侧表示的陷阱卡或里侧表示的魔法或陷阱卡（取对象）
		 ◇场上没有可以作为对象的卡时，这个效果不发动
		 ◇选择的卡为里侧表示的场合，在效果处理时确认那张卡

		 中文名: 陷阱大师
		 卡片种类: 效果怪兽
		 卡片密码: 46461247
		 种族: 战士
		 属性: 地
		 星级: 3
		 攻击力: 500
		 防御力: 1100
		 罕见度: 平卡N
		 卡包: EX-R(EX)、Booster05、Booster R3
		 效果: 反转：场上1张陷阱卡破坏。里侧表示翻开确认后破坏。

		*/
		Id:       1357,
		Password: "46461247",
		Name:     "陷阱大师",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 3,
		La:    ygo.LA_Earth,   // 地
		Lr:    ygo.LR_Warrior, // 战士
		Atk:   500,
		Def:   1100,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 50 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1359
		 调整:

		 [谜之傀儡师]<謎の傀儡師>
		 [2010/09/10]
		 ●怪兽召唤·反转召唤成功时，自己回复500基本分。
		 ◇诱发效果（进入连锁）
		 ◇强制发动
		 ◇这张卡自身召唤、反转召唤成功时不发动
		 ◇二重怪兽再度召唤成功时也发动
		 ◇对方场上召唤、反转召唤怪兽成功时也发动★同一组连锁中召唤了2次怪兽的场合，连锁处理完毕后这张卡发动几次效果
		调整中

		 中文名: 谜之傀儡师
		 卡片种类: 效果怪兽
		 卡片密码: 54098121
		 种族: 战士
		 属性: 地
		 星级: 4
		 攻击力: 1000
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: EX-R(EX)、Booster05、Booster R3、TP07
		 效果: 怪兽召唤·反转召唤成功时，自己回复500基本分。

		*/
		Id:       1359,
		Password: "54098121",
		Name:     "谜之傀儡师",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Earth,   // 地
		Lr:    ygo.LR_Warrior, // 战士
		Atk:   1000,
		Def:   1500,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 51 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2530
		 调整:

		 [暗之艺术家]<闇の芸術家>
		 ●受到了光属性怪物的攻击时，此卡的守备力成为一半。
		 ◇诱发效果，强制发动★时点在伤判怪兽翻开后，具体未知
		 ◇
		 [天邪鬼的诅咒/あまのじゃくの呪い]効果不适用
		 ◇此卡没被战破的场合，此效果持续到伤判结束

		 中文名: 暗之艺术家
		 卡片种类: 效果怪兽
		 卡片密码: 72520073
		 种族: 恶魔
		 属性: 暗
		 星级: 3
		 攻击力: 600
		 防御力: 1400
		 罕见度: 金字UR
		 卡包: Booster03
		 效果: 受到了光属性怪兽的攻击时，这张卡的守备力成为一半。

		*/
		Id:       2530,
		Password: "72520073",
		Name:     "暗之艺术家",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 3,
		La:    ygo.LA_Dark,  // 暗
		Lr:    ygo.LR_Fiend, // 恶魔
		Atk:   600,
		Def:   1400,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 52 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2551
		 调整:

		 [机巧蜘蛛]<カラクリ蜘蛛>
		 [2010/09/10]
		 ●受到这张卡攻击的是暗属性怪兽的场合，那只怪兽破坏。伤害计算适用。
		 ◇诱发效果（进入连锁）
		 ◇强制发动
		 ◇不取对象
		 ◇在和反转效果怪兽发动反转效果相同时点发动
		 ◇由于卡的效果原本是暗属性的怪兽属性变为其他的怪兽受到攻击的场合，这个效果不适用

		 中文名: 机巧蜘蛛
		 卡片种类: 效果怪兽
		 卡片密码: 45688586
		 种族: 机械
		 属性: 地
		 星级: 2
		 攻击力: 400
		 防御力: 500
		 罕见度: 平卡N
		 卡包: Booster04、TP14、STBL(702)
		 效果: 受到「机巧蜘蛛」攻击的暗属性怪兽必定被破坏。（伤害的计算适用）

		*/
		Id:       2551,
		Password: "45688586",
		Name:     "机巧蜘蛛",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 2,
		La:    ygo.LA_Earth,   // 地
		Lr:    ygo.LR_Machine, // 机械
		Atk:   400,
		Def:   500,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 53 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2552
		 调整:

		 [异国的剑士]<異国の剣士>
		 [2010/09/10]
		 ●受到「异国的剑士」攻击的怪兽，5回合后破坏。
		 ◇诱发效果（进入连锁）
		 ◇强制发动
		 ◇不取对象
		 ◇在怪兽被战斗破坏并送去墓地时发动
		 ◇受到攻击的怪兽一度不在场上表侧表示存在，这个效果重置
		 ◇回合计算为受到攻击的怪兽的控制者的回合
		 ◇处理破坏在第5回合的结束阶段时

		 中文名: 异国的剑士
		 卡片种类: 效果怪兽
		 卡片密码: 85255550
		 种族: 战士
		 属性: 地
		 星级: 1
		 攻击力: 250
		 防御力: 250
		 罕见度: 平卡N
		 卡包: Booster04
		 效果: 受到「异国的剑士」攻击的怪兽，5回合后被破坏。

		*/
		Id:       2552,
		Password: "85255550",
		Name:     "异国的剑士",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 1,
		La:    ygo.LA_Earth,   // 地
		Lr:    ygo.LR_Warrior, // 战士
		Atk:   250,
		Def:   250,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 54 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2553
		 调整:

		 [区域吞噬者]<ゾーン·イーター>
		 [2010/09/10]
		 ●受到「区域吞噬者」攻击的怪兽，5回合后破坏。
		 ◇诱发效果（进入连锁）
		 ◇强制发动
		 ◇不取对象
		 ◇在怪兽被战斗破坏并送去墓地时发动
		 ◇受到攻击的怪兽一度不在场上表侧表示存在，这个效果重置
		 ◇回合计算为受到攻击的怪兽的控制者的回合
		 ◇处理破坏在第5回合的结束阶段时

		 中文名: 区域吞噬者
		 卡片种类: 效果怪兽
		 卡片密码: 86100785
		 种族: 水
		 属性: 水
		 星级: 1
		 攻击力: 250
		 防御力: 200
		 罕见度: 平卡N
		 卡包: Booster04
		 效果: 受到「区域吞噬者」攻击的怪兽，5回合后被破坏。

		*/
		Id:       2553,
		Password: "86100785",
		Name:     "区域吞噬者",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 1,
		La:    ygo.LA_Water, // 水
		Lr:    ygo.LR_Water, // 水
		Atk:   250,
		Def:   200,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 55 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2595
		 调整:

		 [狂风毒蛾]<ゲール·ドグラ>
		 [2010/09/10]
		 ●支付3000基本分发动。从自己的融合卡组把1只怪兽丢弃去墓地。
		 ◇发动时支付3000基本分（代价）
		 ◇效果解决时从自己的额外卡组把1只融合怪兽或同调怪兽丢弃去墓地（不取对象）

		 中文名: 狂风毒蛾
		 卡片种类: 效果怪兽
		 卡片密码: 16229315
		 种族: 昆虫
		 属性: 地
		 星级: 2
		 攻击力: 650
		 防御力: 600
		 罕见度: 平卡N，银字R
		 卡包: Booster06、TP04、ABPF(607)
		 效果: 支付3000点基本分，选择我方的额外卡组里的1只怪兽丢弃到墓地里。

		*/
		Id:       2595,
		Password: "16229315",
		Name:     "狂风毒蛾",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 2,
		La:    ygo.LA_Earth,  // 地
		Lr:    ygo.LR_Insect, // 昆虫
		Atk:   650,
		Def:   600,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 56 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 421
		 调整:

		 [暗晦之城]<闇晦ましの城>
		 [2010/09/10]
		 ●反转：在场上表侧表示存在的不死族怪兽的攻击力与守备力上升200分。并且，只要这张卡在场上表侧表示存在，每次自己的准备阶段上升200分。这个效果持续到自己的第4回合的准备阶段为止。
		 ◇反转效果（进入连锁）
		 ◇强制发动
		 ◇准备阶段攻守上升的对象为受到这个反转效果发动时效果的怪兽
		 ◇这张卡不在场上表侧表示存在的场合，这些攻守上升重置
		 ◇准备阶段的攻守提升进入连锁（是反转效果的一部分）

		 中文名: 暗晦之城
		 卡片种类: 效果怪兽
		 卡片密码: 00062121
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 920
		 防御力: 1930
		 罕见度: 平卡N
		 卡包: ME、DL04、Booster07
		 效果: 反转：场上表侧表示的全部不死族的怪兽攻击力·守备力加200。只要这张卡在场上表侧表示存在，自己的每个准备阶段不死族的攻守力都会加200。这个效果持续到自己的第4个准备阶段。

		*/
		Id:       421,
		Password: "00062121",
		Name:     "暗晦之城",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Dark,  // 暗
		Lr:    ygo.LR_Fiend, // 恶魔
		Atk:   920,
		Def:   1930,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 57 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 426
		 调整:

		 [杀人小丑]<マーダーサーカス>
		 [2010/09/10]
		 ●这张卡的表示形式从守备表示变为攻击表示时，对方场上的1只怪兽回到持有者的手卡。
		 ◇诱发效果（进入连锁）
		 ◇强制发动
		 ◇发动时选择对方场上1只怪兽（取对象）

		 中文名: 杀人小丑
		 卡片种类: 效果怪兽
		 卡片密码: 93889755
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 1350
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: ME、BE02、DL04、Booster07
		 效果: 这张卡的表示形式从守备表示变成攻击表示时，对方场上的1只怪兽回到持有者手卡。

		*/
		Id:       426,
		Password: "93889755",
		Name:     "杀人小丑",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Dark,  // 暗
		Lr:    ygo.LR_Fiend, // 恶魔
		Atk:   1350,
		Def:   1400,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 58 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 427
		 调整:
		 [幽灵王-南瓜王-]
		<ゴースト王－パンプキング－>

		 [2010/09/10]

		 ●只要「暗晦之城」在场上表侧表示存在，这张卡的攻击力与守备力上升100分。

		 ◇永续效果（不进入连锁）

		 ●并且，每次自己的准备阶段上升100分。这个效果持续到自己的第4回合的准备阶段为止。

		 ◇诱发效果（进入连锁）

		 ◇强制发动
		 中文名: 南瓜幽灵王
		 卡片种类: 效果怪兽
		 卡片密码: 29155212
		 种族: 不死
		 属性: 暗
		 星级: 6
		 攻击力: 1800
		 防御力: 2000
		 罕见度: {rarely }
		 卡包: ME，DL04，Booster07
		 效果: 只要「暗晦之城」在场上表侧表示存在，这张卡的攻击力·守备力上升100。每次自己的准备阶段这张卡的攻防加100。这个效果持续到自己的第4个准备阶段。

		*/
		Id:       427,
		Password: "29155212",
		Name:     "南瓜幽灵王",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 6,
		La:    ygo.LA_Dark,   // 暗
		Lr:    ygo.LR_Zombie, // 不死
		Atk:   1800,
		Def:   2000,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 59 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 428
		 调整:

		 [梦幻小丑]<ドリーム·ピエロ>
		 [2010/09/10]
		 ●这张卡的表示形式从攻击表示变为守备表示时，破坏对方场上的1只怪兽。
		 ◇诱发效果（进入连锁）
		 ◇强制发动
		 ◇发动时选择对方场上1只怪兽（取对象）
		 ◇变为里侧守备表示的场合不发动

		 中文名: 梦幻小丑
		 卡片种类: 效果怪兽
		 卡片密码: 13215230
		 种族: 战士
		 属性: 地
		 星级: 3
		 攻击力: 1200
		 防御力: 900
		 罕见度: 平卡N
		 卡包: ME、BE02、DL04、Booster07、PE
		 效果: 这张卡的表示形式从攻击表示变成守备表示时，破坏对方场上的1只怪兽。

		*/
		Id:       428,
		Password: "13215230",
		Name:     "梦幻小丑",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 3,
		La:    ygo.LA_Earth,   // 地
		Lr:    ygo.LR_Warrior, // 战士
		Atk:   1200,
		Def:   900,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 60 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 429
		 调整:

		 [恶魔的智慧]<悪魔の知恵>
		 [2010/09/10]
		 ●这张卡的表示形式从攻击表示变为守备表示时，洗切自己的卡组。
		 ◇诱发效果（进入连锁）
		 ◇强制发动
		 ◇变为里侧守备表示的场合不发动
		 ◇自己卡组只有1张或0张的场合，这个效果也发动

		 中文名: 恶魔的智慧
		 卡片种类: 效果怪兽
		 卡片密码: 28725004
		 种族: 恶魔
		 属性: 暗
		 星级: 3
		 攻击力: 1250
		 防御力: 800
		 罕见度: 平卡N
		 卡包: ME、DL04、Booster07
		 效果: 这张卡的表示形式从攻击表示变成守备表示时，洗自己的卡组。

		*/
		Id:       429,
		Password: "28725004",
		Name:     "恶魔的智慧",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 3,
		La:    ygo.LA_Dark,  // 暗
		Lr:    ygo.LR_Fiend, // 恶魔
		Atk:   1250,
		Def:   800,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 61 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 430
		 调整:

		 [艾尔的小剑士]<アイルの小剣士>
		 [2010/09/10]
		 ●每把自己的场上存在的这张卡以外的1只怪兽作为祭品，这张卡的攻击力直到回合结束时为止上升700分。
		 ◇启动效果（进入连锁）
		 ◇发动时把自己场上这张卡以外的1只（只能1只）怪兽解放（代价）
		 ◇1回合的发动次数不限（符合条件的前提下）
		 ◇回合结束时即为回合的结束阶段时

		 中文名: 艾尔的小剑士
		 卡片种类: 效果怪兽
		 卡片密码: 25109950
		 种族: 战士
		 属性: 水
		 星级: 3
		 攻击力: 800
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: ME、DL04、Booster07
		 效果: 每用自己场上存在的这张卡以外的1只怪兽做祭品，这张卡的攻击力在回合结束前加700。

		*/
		Id:       430,
		Password: "25109950",
		Name:     "艾尔的小剑士",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 3,
		La:    ygo.LA_Water,   // 水
		Lr:    ygo.LR_Warrior, // 战士
		Atk:   800,
		Def:   1300,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 62 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 433
		 调整:

		 [地雷蜘蛛]<地雷蜘蛛>
		 [2010/09/10]
		 ●这张卡的攻击宣言时，投掷硬币猜正反。猜中的场合就那样攻击。猜错的场合自己的基本分失去一半攻击。
		 ◇永续效果（不进入连锁）
		 ◇玩家在选择这张卡进行攻击时（选择攻击对象之前），这个效果适用并处理
		 ◇没猜中攻击也正常进行
		 ◇攻击宣言后攻击卷返再选择攻击对象进行攻击的场合，不需要再投掷硬币猜正反

		 中文名: 地雷蜘蛛
		 卡片种类: 效果怪兽
		 卡片密码: 94773007
		 种族: 昆虫
		 属性: 地
		 星级: 4
		 攻击力: 2200
		 防御力: 100
		 罕见度: 平卡N
		 卡包: ME、BE02、DL04、Booster07
		 效果: 这张卡攻击宣言时，猜硬币的正反。猜中的话就继续攻击。猜不中的话自己的基本分减半再攻击。

		*/
		Id:       433,
		Password: "94773007",
		Name:     "地雷蜘蛛",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Earth,  // 地
		Lr:    ygo.LR_Insect, // 昆虫
		Atk:   2200,
		Def:   100,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 63 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 441
		 调整:

		 [空中的昆虫兵]<空の昆虫兵>
		 [2010/09/10]
		 ●攻击风属性怪兽的场合，伤害步骤中这张卡的攻击力上升1000分。
		 ◇永续效果（不进入连锁）
		 ◇进入伤害步骤开始就适用

		 中文名: 空中的昆虫兵
		 卡片种类: 效果怪兽
		 卡片密码: 07019529
		 种族: 昆虫
		 属性: 风
		 星级: 3
		 攻击力: 1000
		 防御力: 800
		 罕见度: 平卡N
		 卡包: ME、DL04、Booster07
		 效果: 攻击风属性的怪兽时，伤害计算时攻击力上升1000。

		*/
		Id:       441,
		Password: "07019529",
		Name:     "空中的昆虫兵",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 3,
		La:    ygo.LA_Wind,   // 风
		Lr:    ygo.LR_Insect, // 昆虫
		Atk:   1000,
		Def:   800,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 64 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 444
		 调整:

		 [寄居龙]<ヤドカリュー>
		 [2010/09/10]
		 ●这张卡的表示形式从攻击表示变为表侧守备表示时，可以从自己的手卡把任意张数的卡放到卡组的最下方。
		 ◇诱发效果（进入连锁）
		 ◇任意发动
		 ◇自己没有手卡时不能发动
		 ◇效果处理时把任意张手卡放到卡组的最下方（这时必须至少有1张）
		 ◇顺序由这张卡的发动者决定
		 ◇原本持有者是对方的卡放到对方卡组的最下方

		 中文名: 寄居龙
		 卡片种类: 效果怪兽
		 卡片密码: 29380133
		 种族: 水
		 属性: 水
		 星级: 4
		 攻击力: 900
		 防御力: 1700
		 罕见度: 平卡N
		 卡包: ME、DL04、Booster07、TP01
		 效果: 这张卡的表示形式从攻击表示变成表侧守备表示时，可以从自己手卡把任意数量的卡回到卡组最下面。

		*/
		Id:       444,
		Password: "29380133",
		Name:     "寄居龙",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Water, // 水
		Lr:    ygo.LR_Water, // 水
		Atk:   900,
		Def:   1700,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 65 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 448
		 调整:

		 [恶魔烹调师]<悪魔の調理師>
		 [2010/09/10]
		 ●这张卡给予对方战斗伤害时，对方从卡组抽2张卡。
		 ◇诱发效果（进入连锁）
		 ◇强制发动

		 中文名: 恶魔烹调师
		 卡片种类: 效果怪兽
		 卡片密码: 71107816
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 1800
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: ME、BE02、DL04、Booster07
		 效果: 这张卡给予对方战斗伤害时，对方在卡组最上面抽2张卡。

		*/
		Id:       448,
		Password: "71107816",
		Name:     "恶魔烹调师",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Dark,  // 暗
		Lr:    ygo.LR_Fiend, // 恶魔
		Atk:   1800,
		Def:   1000,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 66 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 491
		 调整:

		 [邪恶蠕虫兽]<邪悪なるワーム·ビースト>
		 [2010/09/10]
		 ●自己回合的结束阶段时，表侧表示在场上存在的这张卡回到持有者的手卡。
		 ◇诱发效果（进入连锁）
		 ◇强制发动
		 ◇自己的结束阶段时「亚空间物质传送装置/亜空間物質転送装置」让这张卡回到场上的场合，这个效果也必须发动

		 中文名: 邪恶蠕虫兽
		 卡片种类: 效果怪兽
		 卡片密码: 06285791
		 种族: 兽
		 属性: 地
		 星级: 3
		 攻击力: 1400
		 防御力: 700
		 罕见度: 平卡N
		 卡包: BE02、EX-R(EX)、DL04、Booster04、Booster Chronicle、Booster R2、KA、SK2、GLD04
		 效果: 自己的回合的结束阶段时，场上表侧表示存在的这张卡回到持有者的手卡。

		*/
		Id:       491,
		Password: "06285791",
		Name:     "邪恶蠕虫兽",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 3,
		La:    ygo.LA_Earth, // 地
		Lr:    ygo.LR_Beast, // 兽
		Atk:   1400,
		Def:   700,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 67 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 494
		 调整:

		 [森之住人
		乌丹]<森の住人
		ウダン>
		 ●场上每存在一只打开表示的植物族怪兽，这张卡的攻击力上升100。
		 ◇永续效果
		 ◇数双方场上的植物族怪兽
		 ◇
		 [天邪鬼的诅咒/あまのじゃくの呪い]効果适用

		 中文名: 森之住人 乌丹
		 卡片种类: 效果怪兽
		 卡片密码: 42883273
		 种族: 战士
		 属性: 地
		 星级: 3
		 攻击力: 900
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: DL04、Booster03、Booster Chronicle、Booster R2
		 效果: 场上每存在1只表侧表示的植物族怪兽，这张卡的攻击力上升100。

		*/
		Id:       494,
		Password: "42883273",
		Name:     "森之住人 乌丹",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 3,
		La:    ygo.LA_Earth,   // 地
		Lr:    ygo.LR_Warrior, // 战士
		Atk:   900,
		Def:   1200,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 68 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 497
		 调整:

		 [蟑螂骑士]<コカローチ·ナイト>
		 [2010/09/10]
		 ●这张卡送去墓地时，这张卡回到卡组的最上方。
		 ◇诱发效果（进入连锁）
		 ◇强制发动
		 ◇取对象效果
		 ◇伤害步骤中也发动
		 ◇作为怪兽装备、被「升天之角笛/昇天の角笛」「神之宣告/神の宣告」无效召唤而破坏、结束阶段时手卡调整、里侧表示被送去墓地的场合，这个效果发动

		 中文名: 蟑螂骑士
		 卡片种类: 效果怪兽
		 卡片密码: 33413638
		 种族: 昆虫
		 属性: 地
		 星级: 3
		 攻击力: 800
		 防御力: 900
		 罕见度: 平卡N
		 卡包: DL04、Booster04、Booster Chronicle、Booster R2
		 效果: 这张卡送去墓地时，这张卡回到卡组最上面。

		*/
		Id:       497,
		Password: "33413638",
		Name:     "蟑螂骑士",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 3,
		La:    ygo.LA_Earth,  // 地
		Lr:    ygo.LR_Insect, // 昆虫
		Atk:   800,
		Def:   900,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 69 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 499
		 调整:

		 [巡逻机器人]<パトロール·ロボ>
		 ●只要这张卡在场上打开表示存在，可以在自己的准备流程看对方盖放的一张卡。(看完后盖回去，不诱发效果)
		 ◇诱发效果，任意发动，取对象
		 ◇一回合只能用一次
		 ◇里侧表示的怪/魔/陷都可以成为此效果的对象
		 ◇效果处理时对象变成表侧的场合，效果不发★效果处理时此卡不在场上表侧表示存在的场合，效果是否不发，调整中

		 中文名: 巡逻机器人
		 卡片种类: 效果怪兽
		 卡片密码: 76775123
		 种族: 机械
		 属性: 地
		 星级: 3
		 攻击力: 1100
		 防御力: 900
		 罕见度: 平卡N
		 卡包: DL04、Booster03、Booster Chronicle、Booster R2、TP02
		 效果: 只要这张卡在场上表侧表示存在，自己的准备阶段可以把对方场上盖放的1张卡确认。

		*/
		Id:       499,
		Password: "76775123",
		Name:     "巡逻机器人",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 3,
		La:    ygo.LA_Earth,   // 地
		Lr:    ygo.LR_Machine, // 机械
		Atk:   1100,
		Def:   900,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 70 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 501
		 调整:

		 [勇气之沙漏]<勇気の砂時計>
		 [2010/09/10]
		 ●召唤·反转召唤的场合，直到下次的自己回合的结束阶段为止，这张卡的原本的攻击力·原本的守备力减半。那之后，这张卡的原本的攻击力与原本的守备力变成两倍。
		 ◇诱发效果（进入连锁）
		 ◇强制发动
		 ◇改变的数值是原本数值
		 ◇这个效果适用后，给这张卡装备「巨大化/巨大化」的场合，以这个效果改变后的原本数值进行参照

		 中文名: 勇气之沙漏
		 卡片种类: 效果怪兽
		 卡片密码: 43530283
		 种族: 天使
		 属性: 光
		 星级: 4
		 攻击力: 1100
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: DL04、Booster03、Booster Chronicle、Booster R2、STOR(703)、TP16
		 效果: 召唤·反转召唤的场合，直到下次自己的回合的结束阶段这张卡的原来的攻击力·守备力减半，之后这张卡的原来的攻击力·守备力变成2倍。

		*/
		Id:       501,
		Password: "43530283",
		Name:     "勇气之沙漏",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Light, // 光
		Lr:    ygo.LR_Fairy, // 天使
		Atk:   1100,
		Def:   1200,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 71 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 509
		 调整:

		 [心眼之女神]<心眼の女神>
		 [2010/09/10]
		 ●这张卡可以代替1只融合素材怪兽。那时，其他的融合素材怪兽必须是正规物。
		 ◇无效果分类
		 ◇可以被「技能吸收/スキルドレイン」「星骸龙/デブリ·ドラゴン」「冥界的魔王
		哈·迪斯/冥界の魔王
		ハ·デス」等的效果无效
		 ◇只能代替融合怪兽上写全卡名的融合素材
		 ◇只能在进行融合召唤时代替融合素材怪兽
		 ◇不能从卡组、除外状态代替融合素材怪兽
		 ◇可以从手卡、场上（不论表里）、墓地代替融合素材怪兽

		 中文名: 心眼之女神
		 卡片种类: 效果怪兽
		 卡片密码: 53493204
		 种族: 天使
		 属性: 光
		 星级: 4
		 攻击力: 1200
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: BE02、DL04、Booster06、Booster Chronicle、Booster R3、JY、GLD04
		 效果: 这张卡可以代替融合怪兽素材的其中1只来融合。这个时候，其他的融合素材必须是指定的融合素材。

		*/
		Id:       509,
		Password: "53493204",
		Name:     "心眼之女神",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Light, // 光
		Lr:    ygo.LR_Fairy, // 天使
		Atk:   1200,
		Def:   1000,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 72 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 511
		 调整:

		 [沼地的魔兽王]<沼地の魔獣王>
		 [2010/09/10]
		 ●这张卡可以代替1只融合素材怪兽。那时，其他的融合素材怪兽必须是正规物。
		 ◇无效果分类
		 ◇可以被「技能吸收/スキルドレイン」「星骸龙/デブリ·ドラゴン」「冥界的魔王
		哈·迪斯/冥界の魔王
		ハ·デス」等的效果无效
		 ◇只能代替融合怪兽上写全卡名的融合素材
		 ◇只能在进行融合召唤时代替融合素材怪兽
		 ◇不能从卡组、除外状态代替融合素材怪兽
		 ◇可以从手卡、场上（不论表里）、墓地代替融合素材怪兽

		 中文名: 沼地的魔兽王
		 卡片种类: 效果怪兽
		 卡片密码: 99426834
		 种族: 水
		 属性: 水
		 星级: 4
		 攻击力: 1000
		 防御力: 1100
		 罕见度: 平卡N
		 卡包: DL04、Booster06、Booster Chronicle、Booster R3、GLD04
		 效果: 这张卡可以代替融合怪兽素材的其中1只来融合。这个时候，其他的融合素材必须是指定的融合素材。

		*/
		Id:       511,
		Password: "99426834",
		Name:     "沼地的魔兽王",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Water, // 水
		Lr:    ygo.LR_Water, // 水
		Atk:   1000,
		Def:   1100,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 73 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 512
		 调整:

		 [破坏神
		瓦沙克]<破壊神
		ヴァサーゴ>
		 [2010/09/10]
		 ●这张卡可以代替1只融合素材怪兽。那时，其他的融合素材怪兽必须是正规物。
		 ◇无效果分类
		 ◇可以被「技能吸收/スキルドレイン」「星骸龙/デブリ·ドラゴン」「冥界的魔王
		哈·迪斯/冥界の魔王
		ハ·デス」等的效果无效
		 ◇只能代替融合怪兽上写全卡名的融合素材
		 ◇只能在进行融合召唤时代替融合素材怪兽
		 ◇不能从卡组、除外状态代替融合素材怪兽
		 ◇可以从手卡、场上（不论表里）、墓地代替融合素材怪兽

		 中文名: 破坏神 瓦沙克
		 卡片种类: 效果怪兽
		 卡片密码: 50259460
		 种族: 恶魔
		 属性: 暗
		 星级: 3
		 攻击力: 1100
		 防御力: 900
		 罕见度: 平卡N
		 卡包: DL04、Booster06、Booster Chronicle、Booster R3、KA、GLD04
		 效果: 这张卡可以代替融合怪兽素材的其中1只来融合。这个时候，其他的融合素材必须是指定的融合素材。

		*/
		Id:       512,
		Password: "50259460",
		Name:     "破坏神 瓦沙克",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 3,
		La:    ygo.LA_Dark,  // 暗
		Lr:    ygo.LR_Fiend, // 恶魔
		Atk:   1100,
		Def:   900,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 74 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 513
		 调整:

		 [眼球怪]<モンスター·アイ>
		 [2010/09/10]
		 ●支付1000基本分。从自己墓地把1张「融合」魔法卡返回手卡。
		 ◇启动效果（进入连锁）
		 ◇取对象
		 ◇发动次数不限（符合条件的前提下，不得空发）

		 中文名: 眼球怪
		 卡片种类: 效果怪兽
		 卡片密码: 84133008
		 种族: 恶魔
		 属性: 暗
		 星级: 1
		 攻击力: 250
		 防御力: 350
		 罕见度: 平卡N
		 卡包: DL04、Booster06、Booster Chronicle、Booster R3、TP15
		 效果: 支付1000分。自己的墓地1张「融合」魔法卡回到手卡。

		*/
		Id:       513,
		Password: "84133008",
		Name:     "眼球怪",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 1,
		La:    ygo.LA_Dark,  // 暗
		Lr:    ygo.LR_Fiend, // 恶魔
		Atk:   250,
		Def:   350,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 75 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 514
		 调整:

		 [机械王]<機械王>
		 [2010/09/10]
		 ●场上每有1只表侧表示存在的机械族怪兽，这张卡的攻击力上升100分。
		 ◇永续效果（不进入连锁）
		 ◇这张卡自身在效果适用范围内

		 中文名: 机械王
		 卡片种类: 效果怪兽
		 卡片密码: 46700124
		 种族: 机械
		 属性: 地
		 星级: 6
		 攻击力: 2200
		 防御力: 2000
		 罕见度: 平卡N
		 卡包: BE02、DL04、Booster05、Booster Chronicle、Booster R3
		 效果: 每1只在场上表侧表示存在的机械族怪兽，这张卡的攻击力上升100。

		*/
		Id:       514,
		Password: "46700124",
		Name:     "机械王",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 6,
		La:    ygo.LA_Earth,   // 地
		Lr:    ygo.LR_Machine, // 机械
		Atk:   2200,
		Def:   2000,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 76 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 515
		 调整:
		 [恶魔科学怪人]
		<デビル·フランケン>

		 [15/02/02]
		（这张卡在规则上不当作「恶魔/デーモン」卡使用）

		 ●①：支付5000基本分才能发动。从额外卡组把1只融合怪兽攻击表示特殊召唤。

		 ◇起动效果，开连锁，不取对象

		 ◇支付5000基本分是效果发动COST

		 ◇效果处理时从自己的额外卡组把1只融合怪兽在自己场上表侧攻击表示特殊召唤，这个特殊召唤不是融合召唤、不解除苏生限制
		 中文名: 恶魔科学怪人
		 卡片种类: 效果怪兽
		 卡片密码: 69015963
		 种族: 机械
		 属性: 暗
		 星级: 2
		 攻击力: 700
		 防御力: 500
		 罕见度: {rarely }
		 卡包: BE02，DL04，Booster06，Booster Chronicle，Booster R3，KA，SK2，AT06
		 效果: （这张卡在规则上不当作「恶魔」卡使用）①：支付5000基本分才能发动。从额外卡组把1只融合怪兽攻击表示特殊召唤。

		*/
		Id:       515,
		Password: "69015963",
		Name:     "恶魔科学怪人",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 2,
		La:    ygo.LA_Dark,    // 暗
		Lr:    ygo.LR_Machine, // 机械
		Atk:   700,
		Def:   500,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 77 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 516
		 调整:

		 [针球]<ニードル·ボール>
		 [2010/09/10]
		 ●反转：可以支付2000基本分，给予对方的基本分1000分伤害。
		 ◇反转效果（进入连锁）
		 ◇强制发动
		 ◇发动时选择是否在效果处理时适用效果
		 ◇选择适用的场合在发动时支付2000基本分（代价）

		 中文名: 针球
		 卡片种类: 效果怪兽
		 卡片密码: 94230224
		 种族: 恶魔
		 属性: 暗
		 星级: 2
		 攻击力: 750
		 防御力: 700
		 罕见度: 平卡N
		 卡包: DL04、Booster05、Booster Chronicle、Booster R3、STOR(703)
		 效果: 反转：可以支付2000分，对方受到1000分的伤害。

		*/
		Id:       516,
		Password: "94230224",
		Name:     "针球",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 2,
		La:    ygo.LA_Dark,  // 暗
		Lr:    ygo.LR_Fiend, // 恶魔
		Atk:   750,
		Def:   700,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 78 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 517
		 调整:

		 [杀龙者]<竜殺者>
		 [2010/09/10]
		 ●这张卡在场上召唤·反转召唤时，破坏1只表侧表示的龙族怪兽。
		 ◇诱发效果（进入连锁）
		 ◇强制发动
		 ◇发动时选择场上表侧表示存在的1只龙族怪兽（取对象）
		 ◇效果处理时对象怪兽的种族变为龙族以外的场合，这个效果不适用

		 中文名: 杀龙者
		 卡片种类: 效果怪兽
		 卡片密码: 28563545
		 种族: 恶魔
		 属性: 暗
		 星级: 6
		 攻击力: 2000
		 防御力: 2100
		 罕见度: 平卡N
		 卡包: BE02、DL04、Booster05、Booster Chronicle、Booster R3
		 效果: 这张卡在场上召唤·反转召唤时，破坏表侧表示的1只龙族怪兽。

		*/
		Id:       517,
		Password: "28563545",
		Name:     "杀龙者",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 6,
		La:    ygo.LA_Dark,  // 暗
		Lr:    ygo.LR_Fiend, // 恶魔
		Atk:   2000,
		Def:   2100,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 79 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 518
		 调整:

		 [针虫]<ニードルワーム>
		 [2010/09/10]
		 ●反转：从对方的卡组把最上方的5张卡丢弃去墓地。
		 ◇反转效果（进入连锁）
		 ◇强制发动
		 ◇对方卡组不足5张卡的场合这个效果也发动
		 ◇效果处理时对方卡组的卡不足4张的场合，把那些卡剩余的卡全部丢弃去墓地

		 中文名: 针虫
		 卡片种类: 效果怪兽
		 卡片密码: 81843628
		 种族: 昆虫
		 属性: 地
		 星级: 2
		 攻击力: 750
		 防御力: 600
		 罕见度: 平卡N，银字R
		 卡包: BE02、DL04、Booster05、Booster Chronicle、Booster R3、GLD01、KA
		 效果: 反转：对方卡组最上面的5张卡送去墓地。

		*/
		Id:       518,
		Password: "81843628",
		Name:     "针虫",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 2,
		La:    ygo.LA_Earth,  // 地
		Lr:    ygo.LR_Insect, // 昆虫
		Atk:   750,
		Def:   600,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 80 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 519
		 调整:

		 [二重身]<ドッペルゲンガー>
		 [2010/09/10]
		 ●反转：必须选择场上覆盖的2张魔法或陷阱卡，破坏那些。
		 ◇反转效果（进入连锁）
		 ◇强制发动
		 ◇发动时选择场上2张覆盖的魔法或陷阱卡（取对象）
		 ◇场上可以作为对象的卡不足2张的场合，这个效果不发动
		 ◇效果处理时其中之一对象为表侧表示的场合，剩下的那张破坏
		 ◇效果处理时其中之一对象不在场上存在的场合，剩下的那张破坏

		 中文名: 二重身
		 卡片种类: 效果怪兽
		 卡片密码: 61831093
		 种族: 战士
		 属性: 暗
		 星级: 3
		 攻击力: 650
		 防御力: 900
		 罕见度: 平卡N
		 卡包: BE02、DL04、Booster06、Booster Chronicle、Booster R3、DT12、TU06
		 效果: 反转：选择场上盖放的2张魔法·陷阱卡破坏。

		*/
		Id:       519,
		Password: "61831093",
		Name:     "二重身",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 3,
		La:    ygo.LA_Dark,    // 暗
		Lr:    ygo.LR_Warrior, // 战士
		Atk:   650,
		Def:   900,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 81 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 520
		 调整:

		 [变形壶]<メタモルポット>
		 [2010/09/10]
		 ●反转：双方的手卡全部丢弃。那之后，各自从自己的卡组抽5张卡。
		 ◇反转效果（进入连锁）
		 ◇强制发动
		 ◇其中一方卡组的卡不足5张的场合也发动，效果处理时那一方决斗败北
		 ◇「神殿守卫者/神殿を守る者」在场上存在时，这张卡也可以反转召唤及发动效果，这个场合对方只丢卡不抽卡.
		 ◇双方手卡为0这个效果发动的场合，抽卡效果正常处理.

		 中文名: 变形壶
		 卡片种类: 效果怪兽
		 卡片密码: 33508719
		 种族: 岩石
		 属性: 地
		 星级: 2
		 攻击力: 700
		 防御力: 600
		 罕见度: 平卡N，银字R，黄金GR，面闪SR
		 卡包: BE02、DL04、Booster05、Booster Chronicle、Booster R3、SD13、GS02、DB12、GLD04、GDB1
		 效果: 反转：双方手卡全部丢弃。那之后，双方各自从自己卡组抽5张卡。

		*/
		Id:       520,
		Password: "33508719",
		Name:     "变形壶",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 2,
		La:    ygo.LA_Earth, // 地
		Lr:    ygo.LR_Rock,  // 岩石
		Atk:   700,
		Def:   600,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 82 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 521
		 调整:

		 [企鹅士兵]<ペンギン·ソルジャー>
		 [2010/09/10]
		 ●反转：可以让场上存在的最多2只怪兽回到持有者的手卡。
		 ◇反转效果（进入连锁）
		 ◇强制发动
		 ◇发动时选择0~2张符合条件的卡为对象（取对象）
		 ◇不能把被战斗破坏的卡作为对象
		 ◇发动时选择2只对象，效果处理时对象卡只剩下1只的场合，剩下的那只回到持有者手卡

		 中文名: 企鹅士兵
		 卡片种类: 效果怪兽
		 卡片密码: 93920745
		 种族: 水
		 属性: 水
		 星级: 2
		 攻击力: 750
		 防御力: 500
		 罕见度: 平卡N，银字R，面闪SR
		 卡包: BE02、DL04、Booster06、Booster Chronicle、Booster R3、YSD04、YU、JY、DT12、SD23
		 效果: 反转：可以让场上存在的最多2只怪兽回到持有者手卡。

		*/
		Id:       521,
		Password: "93920745",
		Name:     "企鹅士兵",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 2,
		La:    ygo.LA_Water, // 水
		Lr:    ygo.LR_Water, // 水
		Atk:   750,
		Def:   500,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 83 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 523
		 调整:

		 [幻想绵羊]<イリュージョン·シープ>
		 [2010/09/10]
		 ●这张卡可以代替1只融合素材怪兽。那时，其他的融合素材怪兽必须是正规物。
		 ◇无效果分类
		 ◇可以被「技能吸收/スキルドレイン」「星骸龙/デブリ·ドラゴン」「冥界的魔王
		哈·迪斯/冥界の魔王
		ハ·デス」等的效果无效
		 ◇只能代替融合怪兽上写全卡名的融合素材
		 ◇只能在进行融合召唤时代替融合素材怪兽
		 ◇不能从卡组、除外状态代替融合素材怪兽
		 ◇可以从手卡、场上（不论表里）、墓地代替融合素材怪兽

		 中文名: 幻想绵羊
		 卡片种类: 效果怪兽
		 卡片密码: 30451366
		 种族: 兽
		 属性: 地
		 星级: 3
		 攻击力: 1150
		 防御力: 900
		 罕见度: 平卡N
		 卡包: DL04、Booster06、Booster Chronicle、Booster R3、TP08
		 效果: 这张卡可以代替融合怪兽素材的其中1只来融合。这个时候，其他的融合素材必须是指定的融合素材。

		*/
		Id:       523,
		Password: "30451366",
		Name:     "幻想绵羊",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 3,
		La:    ygo.LA_Earth, // 地
		Lr:    ygo.LR_Beast, // 兽
		Atk:   1150,
		Def:   900,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 84 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1227
		 调整:

		 中文名: 灯之魔精
		 卡片种类: 通常怪兽
		 卡片密码: 97590747
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 1800
		 防御力: 1000
		 罕见度: 平卡N，金字UR
		 卡包: LE03、EX-R(EX)、Booster04、Booster R2、TP10、DPKB、KA、SK2
		 效果: 描述：听从呼唤他的主人的任何要求及命令的灯之精灵。

		*/
		Id:       1227,
		Password: "97590747",
		Name:     "灯之魔精",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     1800,
		Def:     1000,
		IsValid: true,
	})

	/* 85 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2445
		 调整:

		 中文名: 女夜魔战士
		 卡片种类: 通常怪兽
		 卡片密码: 78556320
		 种族: 战士
		 属性: 暗
		 星级: 3
		 攻击力: 900
		 防御力: 700
		 罕见度: 平卡N
		 卡包: Booster01、Booster R1
		 效果: 描述：侍奉黑暗的女战士。将对方进行血祭是她生存的意义。

		*/
		Id:       2445,
		Password: "78556320",
		Name:     "女夜魔战士",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Dark,    // 暗
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     900,
		Def:     700,
		IsValid: true,
	})

	/* 86 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2446
		 调整:

		 中文名: 气象控制员
		 卡片种类: 通常怪兽
		 卡片密码: 37243151
		 种族: 天使
		 属性: 光
		 星级: 2
		 攻击力: 600
		 防御力: 400
		 罕见度: 平卡N
		 卡包: Booster01
		 效果: 描述：可以自由操纵天气。山里的天气无常就是这家伙的杰作。

		*/
		Id:       2446,
		Password: "37243151",
		Name:     "气象控制员",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Light, // 光
		Lr:      ygo.LR_Fairy, // 天使
		Atk:     600,
		Def:     400,
		IsValid: true,
	})

	/* 87 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2447
		 调整:

		 中文名: 水元素
		 卡片种类: 通常怪兽
		 卡片密码: 03732747
		 种族: 水
		 属性: 水
		 星级: 3
		 攻击力: 900
		 防御力: 700
		 罕见度: 平卡N
		 卡包: Booster01、Booster R1
		 效果: 描述：住在水里的精灵。将四周用雾包围妨碍敌人的视线。

		*/
		Id:       2447,
		Password: "03732747",
		Name:     "水元素",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     900,
		Def:     700,
		IsValid: true,
	})

	/* 88 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2448
		 调整:

		 中文名: 石像怪
		 卡片种类: 通常怪兽
		 卡片密码: 15303296
		 种族: 恶魔
		 属性: 暗
		 星级: 3
		 攻击力: 1000
		 防御力: 500
		 罕见度: 平卡N
		 卡包: EX-R(EX)、Booster01、Booster R1
		 效果: 描述：使人误认为是石像，从而在黑暗之中攻击。逃跑速度也很快。

		*/
		Id:       2448,
		Password: "15303296",
		Name:     "石像怪",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     1000,
		Def:     500,
		IsValid: true,
	})

	/* 89 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2449
		 调整:

		 中文名: 格斗战士 阿提米特
		 卡片种类: 通常怪兽
		 卡片密码: 55550921
		 种族: 战士
		 属性: 地
		 星级: 3
		 攻击力: 700
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: Booster01、Booster R1、JY、TP16
		 效果: 描述：不使用任何武器，空手战斗的格斗战士。

		*/
		Id:       2449,
		Password: "55550921",
		Name:     "格斗战士 阿提米特",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     700,
		Def:     1000,
		IsValid: true,
	})

	/* 90 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2450
		 调整:

		 中文名: 风之番人 精
		 卡片种类: 通常怪兽
		 卡片密码: 97843505
		 种族: 魔法师
		 属性: 风
		 星级: 3
		 攻击力: 700
		 防御力: 900
		 罕见度: 平卡N
		 卡包: Booster01
		 效果: 描述：操纵风并产生龙卷风与飓风吹飞周围的东西。

		*/
		Id:       2450,
		Password: "97843505",
		Name:     "风之番人 精",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Wind,        // 风
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     700,
		Def:     900,
		IsValid: true,
	})

	/* 91 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2451
		 调整:

		 中文名: 格洛斯
		 卡片种类: 通常怪兽
		 卡片密码: 60589682
		 种族: 水
		 属性: 水
		 星级: 3
		 攻击力: 900
		 防御力: 700
		 罕见度: 平卡N
		 卡包: Booster01、Booster R1
		 效果: 描述：用像鞭子一般长的手臂进行攻击。稍远的地方也可以攻击的到。

		*/
		Id:       2451,
		Password: "60589682",
		Name:     "格洛斯",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     900,
		Def:     700,
		IsValid: true,
	})

	/* 92 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2452
		 调整:

		 中文名: 幽灵
		 卡片种类: 通常怪兽
		 卡片密码: 61201220
		 种族: 不死
		 属性: 暗
		 星级: 2
		 攻击力: 600
		 防御力: 800
		 罕见度: 平卡N
		 卡包: Booster01
		 效果: 描述：由这个世界上不能成佛的灵魂，慢慢聚集而成的怨灵。

		*/
		Id:       2452,
		Password: "61201220",
		Name:     "幽灵",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Dark,   // 暗
		Lr:      ygo.LR_Zombie, // 不死
		Atk:     600,
		Def:     800,
		IsValid: true,
	})

	/* 93 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2453
		 调整:

		 中文名: 萨塔那
		 卡片种类: 通常怪兽
		 卡片密码: 77603950
		 种族: 魔法师
		 属性: 暗
		 星级: 2
		 攻击力: 700
		 防御力: 600
		 罕见度: 平卡N
		 卡包: Booster01
		 效果: 描述：诅咒敌人。可以使人动弹不得的魔法师。

		*/
		Id:       2453,
		Password: "77603950",
		Name:     "萨塔那",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Dark,        // 暗
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     700,
		Def:     600,
		IsValid: true,
	})

	/* 94 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2454
		 调整:

		 中文名: 邪炎之翼
		 卡片种类: 通常怪兽
		 卡片密码: 92944626
		 种族: 炎
		 属性: 炎
		 星级: 2
		 攻击力: 700
		 防御力: 600
		 罕见度: 平卡N
		 卡包: Booster01、TP17
		 效果: 描述：原形是红黑色燃烧着的翅膀。从全身喷出火焰进行攻击。

		*/
		Id:       2454,
		Password: "92944626",
		Name:     "邪炎之翼",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Fire, // 炎
		Lr:      ygo.LR_Fire, // 炎
		Atk:     700,
		Def:     600,
		IsValid: true,
	})

	/* 95 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2455
		 调整:

		 中文名: 大嘴鸟
		 卡片种类: 通常怪兽
		 卡片密码: 97973387
		 种族: 鸟兽
		 属性: 风
		 星级: 2
		 攻击力: 600
		 防御力: 500
		 罕见度: 平卡N
		 卡包: Booster01
		 效果: 描述：嘴巴非常大、大声的叫吓跑胆小的对方。

		*/
		Id:       2455,
		Password: "97973387",
		Name:     "大嘴鸟",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Wind,        // 风
		Lr:      ygo.LR_WingedBeast, // 鸟兽
		Atk:     600,
		Def:     500,
		IsValid: true,
	})

	/* 96 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2456
		 调整:

		 中文名: 神圣之锁
		 卡片种类: 通常怪兽
		 卡片密码: 63515678
		 种族: 天使
		 属性: 光
		 星级: 2
		 攻击力: 700
		 防御力: 700
		 罕见度: 平卡N
		 卡包: Booster01
		 效果: 描述：据说可以用神圣的力量封锁行动的锁链。

		*/
		Id:       2456,
		Password: "63515678",
		Name:     "神圣之锁",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Light, // 光
		Lr:      ygo.LR_Fairy, // 天使
		Atk:     700,
		Def:     700,
		IsValid: true,
	})

	/* 97 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2457
		 调整:

		 中文名: 僵尸鬼灯
		 卡片种类: 通常怪兽
		 卡片密码: 63545455
		 种族: 不死
		 属性: 暗
		 星级: 2
		 攻击力: 500
		 防御力: 400
		 罕见度: 平卡N
		 卡包: Booster01
		 效果: 描述：把手腕像火箭一样发射攻击的不死怪兽。

		*/
		Id:       2457,
		Password: "63545455",
		Name:     "僵尸鬼灯",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Dark,   // 暗
		Lr:      ygo.LR_Zombie, // 不死
		Atk:     500,
		Def:     400,
		IsValid: true,
	})

	/* 98 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2458
		 调整:

		 中文名: 暗黑植物
		 卡片种类: 通常怪兽
		 卡片密码: 13193642
		 种族: 植物
		 属性: 暗
		 星级: 1
		 攻击力: 300
		 防御力: 400
		 罕见度: 平卡N
		 卡包: Booster01
		 效果: 描述：在被污染的土地以及暗之力的环境下生长的花。非常凶暴。

		*/
		Id:       2458,
		Password: "13193642",
		Name:     "暗黑植物",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   1,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Plant, // 植物
		Atk:     300,
		Def:     400,
		IsValid: true,
	})

	/* 99 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2459
		 调整:

		 中文名: 太古之壶
		 卡片种类: 通常怪兽
		 卡片密码: 81492226
		 种族: 岩石
		 属性: 地
		 星级: 1
		 攻击力: 400
		 防御力: 200
		 罕见度: 平卡N
		 卡包: Booster01
		 效果: 描述：非常脆弱的太古之壶。里面好像藏着什么东西。

		*/
		Id:       2459,
		Password: "81492226",
		Name:     "太古之壶",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   1,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Rock,  // 岩石
		Atk:     400,
		Def:     200,
		IsValid: true,
	})

	/* 100 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2460
		 调整:

		 中文名: 鸟嘴兽
		 卡片种类: 通常怪兽
		 卡片密码: 29948642
		 种族: 兽
		 属性: 地
		 星级: 2
		 攻击力: 500
		 防御力: 800
		 罕见度: 平卡N
		 卡包: Booster01
		 效果: 描述：卷起蛇一样长的身体，一边回转一边用嘴攻击。

		*/
		Id:       2460,
		Password: "29948642",
		Name:     "鸟嘴兽",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Atk:     500,
		Def:     800,
		IsValid: true,
	})

	/* 101 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2461
		 调整:

		 中文名: 招手的墓场
		 卡片种类: 通常怪兽
		 卡片密码: 27094595
		 种族: 不死
		 属性: 暗
		 星级: 3
		 攻击力: 700
		 防御力: 900
		 罕见度: 平卡N
		 卡包: Booster01
		 效果: 描述：给予死者进一步提升的力量，引诱生者步向死亡的墓场。

		*/
		Id:       2461,
		Password: "27094595",
		Name:     "招手的墓场",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Dark,   // 暗
		Lr:      ygo.LR_Zombie, // 不死
		Atk:     700,
		Def:     900,
		IsValid: true,
	})

	/* 102 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2462
		 调整:

		 中文名: 多隆
		 卡片种类: 通常怪兽
		 卡片密码: 00756652
		 种族: 战士
		 属性: 地
		 星级: 2
		 攻击力: 900
		 防御力: 500
		 罕见度: 平卡N
		 卡包: Booster01、Booster R1
		 效果: 描述：分身夹击的怪兽，要小心！

		*/
		Id:       2462,
		Password: "00756652",
		Name:     "多隆",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     900,
		Def:     500,
		IsValid: true,
	})

	/* 103 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2463
		 调整:

		 中文名: 溶化的赤影
		 卡片种类: 通常怪兽
		 卡片密码: 98898173
		 种族: 水
		 属性: 水
		 星级: 2
		 攻击力: 500
		 防御力: 700
		 罕见度: 平卡N
		 卡包: Booster01
		 效果: 描述：融化身体藏在脚下的影子里，从敌人的正下方发动攻击。

		*/
		Id:       2463,
		Password: "98898173",
		Name:     "溶化的赤影",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     500,
		Def:     700,
		IsValid: true,
	})

	/* 104 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2464
		 调整:

		 中文名: 梦魇蝎
		 卡片种类: 通常怪兽
		 卡片密码: 88643173
		 种族: 昆虫
		 属性: 地
		 星级: 3
		 攻击力: 900
		 防御力: 800
		 罕见度: 平卡N
		 卡包: Booster01、Booster R1
		 效果: 描述：造出恶梦，并在呻吟的时候用四条有毒的尾巴刺向敌人。

		*/
		Id:       2464,
		Password: "88643173",
		Name:     "梦魇蝎",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,  // 地
		Lr:      ygo.LR_Insect, // 昆虫
		Atk:     900,
		Def:     800,
		IsValid: true,
	})

	/* 105 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2465
		 调整:

		 中文名: 幸福恋人
		 卡片种类: 通常怪兽
		 卡片密码: 99030164
		 种族: 天使
		 属性: 光
		 星级: 2
		 攻击力: 800
		 防御力: 500
		 罕见度: 平卡N
		 卡包: Booster01、TP12
		 效果: 描述：从头上放出心形的光给予恋人以幸福的小天使。

		*/
		Id:       2465,
		Password: "99030164",
		Name:     "幸福恋人",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Light, // 光
		Lr:      ygo.LR_Fairy, // 天使
		Atk:     800,
		Def:     500,
		IsValid: true,
	})

	/* 106 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2466
		 调整:

		 中文名: 飓风怪
		 卡片种类: 通常怪兽
		 卡片密码: 15042735
		 种族: 魔法师
		 属性: 风
		 星级: 2
		 攻击力: 900
		 防御力: 200
		 罕见度: 平卡N
		 卡包: Booster01、Booster R1
		 效果: 描述：荒野的龙卷风。用风之刃切割对方。

		*/
		Id:       2466,
		Password: "15042735",
		Name:     "飓风怪",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Wind,        // 风
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     900,
		Def:     200,
		IsValid: true,
	})

	/* 107 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2467
		 调整:

		 中文名: 食人植物
		 卡片种类: 通常怪兽
		 卡片密码: 49127943
		 种族: 植物
		 属性: 地
		 星级: 2
		 攻击力: 800
		 防御力: 600
		 罕见度: 平卡N
		 卡包: Booster01
		 效果: 描述：食肉花。让人觉得是漂亮的花，然后猛地一口吃掉接近它的人。

		*/
		Id:       2467,
		Password: "49127943",
		Name:     "食人植物",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Plant, // 植物
		Atk:     800,
		Def:     600,
		IsValid: true,
	})

	/* 108 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2468
		 调整:

		 中文名: 火眼
		 卡片种类: 通常怪兽
		 卡片密码: 88435542
		 种族: 炎
		 属性: 炎
		 星级: 2
		 攻击力: 800
		 防御力: 600
		 罕见度: 平卡N
		 卡包: Booster01、Booster R1
		 效果: 描述：被火炎包围着的眼球。挥动翅膀就能引起火炎之风。

		*/
		Id:       2468,
		Password: "88435542",
		Name:     "火眼",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Fire, // 炎
		Lr:      ygo.LR_Fire, // 炎
		Atk:     800,
		Def:     600,
		IsValid: true,
	})

	/* 109 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2469
		 调整:

		 中文名: 法兰克斯
		 卡片种类: 通常怪兽
		 卡片密码: 75646173
		 种族: 兽
		 属性: 地
		 星级: 3
		 攻击力: 800
		 防御力: 900
		 罕见度: 平卡N
		 卡包: Booster01、Booster R1
		 效果: 描述：上下都有头的令人作呕的怪兽。从口中吐出激光。

		*/
		Id:       2469,
		Password: "75646173",
		Name:     "法兰克斯",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Atk:     800,
		Def:     900,
		IsValid: true,
	})

	/* 110 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2470
		 调整:

		 中文名: 地狱之门
		 卡片种类: 通常怪兽
		 卡片密码: 49258578
		 种族: 兽
		 属性: 暗
		 星级: 3
		 攻击力: 700
		 防御力: 800
		 罕见度: 平卡N
		 卡包: Booster01
		 效果: 描述：肚子上长着通向地狱的门，也可以被召唤的怪兽。

		*/
		Id:       2470,
		Password: "49258578",
		Name:     "地狱之门",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Beast, // 兽
		Atk:     700,
		Def:     800,
		IsValid: true,
	})

	/* 111 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2471
		 调整:

		 中文名: 神圣能量
		 卡片种类: 通常怪兽
		 卡片密码: 03985011
		 种族: 魔法师
		 属性: 光
		 星级: 2
		 攻击力: 600
		 防御力: 800
		 罕见度: 平卡N
		 卡包: Booster01
		 效果: 描述：虽然摇摇晃晃的，但被圣之力保护着。

		*/
		Id:       2471,
		Password: "03985011",
		Name:     "神圣能量",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Light,       // 光
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     600,
		Def:     800,
		IsValid: true,
	})

	/* 112 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2472
		 调整:

		 中文名: 魔头邪龙
		 卡片种类: 通常怪兽
		 卡片密码: 02957055
		 种族: 龙
		 属性: 风
		 星级: 3
		 攻击力: 900
		 防御力: 900
		 罕见度: 平卡N
		 卡包: Booster01、Booster R1
		 效果: 描述：长有两个头的龙。用两张嘴咬碎敌人。

		*/
		Id:       2472,
		Password: "02957055",
		Name:     "魔头邪龙",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Wind,   // 风
		Lr:      ygo.LR_Dragon, // 龙
		Atk:     900,
		Def:     900,
		IsValid: true,
	})

	/* 113 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2473
		 调整:

		 中文名: 未熟的恶魔
		 卡片种类: 通常怪兽
		 卡片密码: 64154377
		 种族: 恶魔
		 属性: 暗
		 星级: 2
		 攻击力: 500
		 防御力: 750
		 罕见度: 平卡N
		 卡包: Booster01
		 效果: 描述：还没成为完全体的丑恶魔。肚子上的洞会吸入任何东西。

		*/
		Id:       2473,
		Password: "64154377",
		Name:     "未熟的恶魔",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     500,
		Def:     750,
		IsValid: true,
	})

	/* 114 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2474
		 调整:

		 中文名: 午夜恶魔
		 卡片种类: 通常怪兽
		 卡片密码: 83678433
		 种族: 恶魔
		 属性: 暗
		 星级: 2
		 攻击力: 800
		 防御力: 600
		 罕见度: 平卡N
		 卡包: Booster01
		 效果: 描述：出现在深夜的鸟怪。据说必须祭生才能唤它出来。

		*/
		Id:       2474,
		Password: "83678433",
		Name:     "午夜恶魔",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     800,
		Def:     600,
		IsValid: true,
	})

	/* 115 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2475
		 调整:

		 中文名: 八俣龙绘卷
		 卡片种类: 通常怪兽
		 卡片密码: 76704943
		 种族: 龙
		 属性: 风
		 星级: 2
		 攻击力: 900
		 防御力: 300
		 罕见度: 平卡N
		 卡包: Booster01、Booster R1
		 效果: 描述：画卷龙实体他的产物。守备力相当低。

		*/
		Id:       2475,
		Password: "76704943",
		Name:     "八俣龙绘卷",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Wind,   // 风
		Lr:      ygo.LR_Dragon, // 龙
		Atk:     900,
		Def:     300,
		IsValid: true,
	})

	/* 116 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2476
		 调整:

		 中文名: 司暗的影子
		 卡片种类: 通常怪兽
		 卡片密码: 63125616
		 种族: 恶魔
		 属性: 暗
		 星级: 3
		 攻击力: 800
		 防御力: 700
		 罕见度: 平卡N
		 卡包: Booster01、Booster R1
		 效果: 描述：溶入黑暗之中的影子。用金网封住对方的行动。

		*/
		Id:       2476,
		Password: "63125616",
		Name:     "司暗的影子",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     800,
		Def:     700,
		IsValid: true,
	})

	/* 117 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2477
		 调整:

		 中文名: 邪恶老鼠
		 卡片种类: 通常怪兽
		 卡片密码: 56713552
		 种族: 兽
		 属性: 地
		 星级: 3
		 攻击力: 750
		 防御力: 800
		 罕见度: 平卡N
		 卡包: Booster02
		 效果: 描述：任何东西都咬的野老鼠，行仪极坏。

		*/
		Id:       2477,
		Password: "56713552",
		Name:     "邪恶老鼠",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Atk:     750,
		Def:     800,
		IsValid: true,
	})

	/* 118 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2478
		 调整:

		 中文名: 锯足锹形虫
		 卡片种类: 通常怪兽
		 卡片密码: 70924884
		 种族: 昆虫
		 属性: 地
		 星级: 3
		 攻击力: 950
		 防御力: 700
		 罕见度: 平卡N
		 卡包: Booster02、Booster R1
		 效果: 描述：大甲虫。除了头上的锯子，手腕也变成了锯子。

		*/
		Id:       2478,
		Password: "70924884",
		Name:     "锯足锹形虫",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,  // 地
		Lr:      ygo.LR_Insect, // 昆虫
		Atk:     950,
		Def:     700,
		IsValid: true,
	})

	/* 119 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2479
		 调整:

		 中文名: 威尔米
		 卡片种类: 通常怪兽
		 卡片密码: 92391084
		 种族: 兽
		 属性: 地
		 星级: 4
		 攻击力: 1000
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: Booster02、Booster R1
		 效果: 描述：相当凶暴的兔子。以锐利的钩爪血祭对方。

		*/
		Id:       2479,
		Password: "92391084",
		Name:     "威尔米",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Atk:     1000,
		Def:     1200,
		IsValid: true,
	})

	/* 120 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2480
		 调整:

		 中文名: 木灵小丑
		 卡片种类: 通常怪兽
		 卡片密码: 17511156
		 种族: 战士
		 属性: 地
		 星级: 3
		 攻击力: 800
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: Booster02、Booster R1
		 效果: 描述：作出令人讨厌的笑容的恶魔。以手中的镰刀熟练的回避着攻击。

		*/
		Id:       2480,
		Password: "17511156",
		Name:     "木灵小丑",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     800,
		Def:     1200,
		IsValid: true,
	})

	/* 121 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2481
		 调整:

		 中文名: 天使魔女
		 卡片种类: 通常怪兽
		 卡片密码: 37160778
		 种族: 魔法师
		 属性: 暗
		 星级: 3
		 攻击力: 800
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: Booster02、Booster R1
		 效果: 描述：背负着成为天使的命运，但却成为了向往中的魔女。

		*/
		Id:       2481,
		Password: "37160778",
		Name:     "天使魔女",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Dark,        // 暗
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     800,
		Def:     1000,
		IsValid: true,
	})

	/* 122 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2482
		 调整:

		 中文名: 电波英雄
		 卡片种类: 通常怪兽
		 卡片密码: 82065276
		 种族: 战士
		 属性: 地
		 星级: 3
		 攻击力: 1250
		 防御力: 700
		 罕见度: 平卡N
		 卡包: Booster02、Booster Chronicle、Booster R1、TU04
		 效果: 描述：从异世界来到这里，稀里糊涂的战士。

		*/
		Id:       2482,
		Password: "82065276",
		Name:     "电波英雄",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     1250,
		Def:     700,
		IsValid: true,
	})

	/* 123 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2483
		 调整:

		 中文名: 鱼战士
		 卡片种类: 通常怪兽
		 卡片密码: 69750536
		 种族: 鱼
		 属性: 水
		 星级: 4
		 攻击力: 1250
		 防御力: 900
		 罕见度: 平卡N
		 卡包: Booster02、Booster Chronicle、Booster R1
		 效果: 描述：拥有手和脚的鱼人兽，用尖锐的牙齿撕咬敌人。

		*/
		Id:       2483,
		Password: "69750536",
		Name:     "鱼战士",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Fish,  // 鱼
		Atk:     1250,
		Def:     900,
		IsValid: true,
	})

	/* 124 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2484
		 调整:

		 中文名: 龙虾怪
		 卡片种类: 通常怪兽
		 卡片密码: 10598400
		 种族: 水
		 属性: 水
		 星级: 2
		 攻击力: 600
		 防御力: 700
		 罕见度: 平卡N
		 卡包: Booster02
		 效果: 描述：由虾进化成的怪兽。用大钳子攻击对方的颈项。

		*/
		Id:       2484,
		Password: "10598400",
		Name:     "龙虾怪",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     600,
		Def:     700,
		IsValid: true,
	})

	/* 125 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2485
		 调整:

		 中文名: 开在深渊的花
		 卡片种类: 通常怪兽
		 卡片密码: 40387124
		 种族: 植物
		 属性: 地
		 星级: 2
		 攻击力: 750
		 防御力: 400
		 罕见度: 平卡N
		 卡包: Booster02
		 效果: 描述：默默生长在黑暗的深渊里的花，平常很难见到。

		*/
		Id:       2485,
		Password: "40387124",
		Name:     "开在深渊的花",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Plant, // 植物
		Atk:     750,
		Def:     400,
		IsValid: true,
	})

	/* 126 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2486
		 调整:

		 中文名: 岩石幽灵
		 卡片种类: 通常怪兽
		 卡片密码: 72269672
		 种族: 岩石
		 属性: 地
		 星级: 4
		 攻击力: 1200
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: Booster02、Booster R1
		 效果: 描述：激怒它的话，头上的火山可是会爆发的哦。

		*/
		Id:       2486,
		Password: "72269672",
		Name:     "岩石幽灵",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Rock,  // 岩石
		Atk:     1200,
		Def:     1000,
		IsValid: true,
	})

	/* 127 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2487
		 调整:

		 中文名: 狩魂者
		 卡片种类: 通常怪兽
		 卡片密码: 03606209
		 种族: 兽战士
		 属性: 地
		 星级: 4
		 攻击力: 1100
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: Booster02、Booster R1
		 效果: 描述：据说被其用剑斩杀之人的魂魄会离身体而去。

		*/
		Id:       2487,
		Password: "03606209",
		Name:     "狩魂者",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth,        // 地
		Lr:      ygo.LR_BeastWarrior, // 兽战士
		Atk:     1100,
		Def:     1000,
		IsValid: true,
	})

	/* 128 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2488
		 调整:

		 中文名: 死亡鲨
		 卡片种类: 通常怪兽
		 卡片密码: 34290067
		 种族: 不死
		 属性: 暗
		 星级: 3
		 攻击力: 1100
		 防御力: 700
		 罕见度: 平卡N
		 卡包: Booster02、Booster Chronicle、Booster R1
		 效果: 描述：在海里彷徨的鲨鱼，遇上它的人都会被施与诅咒。

		*/
		Id:       2488,
		Password: "34290067",
		Name:     "死亡鲨",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Dark,   // 暗
		Lr:      ygo.LR_Zombie, // 不死
		Atk:     1100,
		Def:     700,
		IsValid: true,
	})

	/* 129 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2489
		 调整:

		 中文名: 机械蝙蝠
		 卡片种类: 通常怪兽
		 卡片密码: 72076281
		 种族: 机械
		 属性: 风
		 星级: 1
		 攻击力: 300
		 防御力: 350
		 罕见度: 平卡N
		 卡包: Booster02
		 效果: 描述：投下左右翅膀里搭载的爆弹的机械蝙蝠。

		*/
		Id:       2489,
		Password: "72076281",
		Name:     "机械蝙蝠",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   1,
		La:      ygo.LA_Wind,    // 风
		Lr:      ygo.LR_Machine, // 机械
		Atk:     300,
		Def:     350,
		IsValid: true,
	})

	/* 130 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2490
		 调整:

		 中文名: 豌豆战士
		 卡片种类: 通常怪兽
		 卡片密码: 84990171
		 种族: 植物
		 属性: 地
		 星级: 4
		 攻击力: 1400
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: Booster02、Booster Chronicle、Booster R1
		 效果: 描述：使用剑和大豆攻击的植物战士，意外的强悍。

		*/
		Id:       2490,
		Password: "84990171",
		Name:     "豌豆战士",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Plant, // 植物
		Atk:     1400,
		Def:     1300,
		IsValid: true,
	})

	/* 131 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2491
		 调整:

		 中文名: 书之精灵
		 卡片种类: 通常怪兽
		 卡片密码: 68963107
		 种族: 魔法师
		 属性: 暗
		 星级: 2
		 攻击力: 650
		 防御力: 500
		 罕见度: 平卡N
		 卡包: Booster02
		 效果: 描述：书形的魔法师。书里写着各种各样的魔法。

		*/
		Id:       2491,
		Password: "68963107",
		Name:     "书之精灵",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Dark,        // 暗
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     650,
		Def:     500,
		IsValid: true,
	})

	/* 132 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2492
		 调整:

		 中文名: 小霸王龙
		 卡片种类: 通常怪兽
		 卡片密码: 42625254
		 种族: 恐龙
		 属性: 地
		 星级: 3
		 攻击力: 1100
		 防御力: 700
		 罕见度: 平卡N
		 卡包: Booster02、Booster R1
		 效果: 描述：恐龙的孩子。性格极其凶暴。

		*/
		Id:       2492,
		Password: "42625254",
		Name:     "小霸王龙",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,    // 地
		Lr:      ygo.LR_Dinosaur, // 恐龙
		Atk:     1100,
		Def:     700,
		IsValid: true,
	})

	/* 133 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2493
		 调整:

		 中文名: 骨鼠
		 卡片种类: 通常怪兽
		 卡片密码: 21239280
		 种族: 不死
		 属性: 暗
		 星级: 1
		 攻击力: 400
		 防御力: 300
		 罕见度: 平卡N
		 卡包: Booster02
		 效果: 描述：被猫杀死的老鼠。为了报仇雪恨化为不死形态。

		*/
		Id:       2493,
		Password: "21239280",
		Name:     "骨鼠",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   1,
		La:      ygo.LA_Dark,   // 暗
		Lr:      ygo.LR_Zombie, // 不死
		Atk:     400,
		Def:     300,
		IsValid: true,
	})

	/* 134 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2494
		 调整:

		 中文名: 电击企鹅
		 卡片种类: 通常怪兽
		 卡片密码: 48531733
		 种族: 雷
		 属性: 水
		 星级: 3
		 攻击力: 1100
		 防御力: 800
		 罕见度: 平卡N
		 卡包: Booster02、Booster R1
		 效果: 描述：用两腕的电击使对方麻痹，然后绞住颈项，进行攻击。

		*/
		Id:       2494,
		Password: "48531733",
		Name:     "电击企鹅",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Water,   // 水
		Lr:      ygo.LR_Thunder, // 雷
		Atk:     1100,
		Def:     800,
		IsValid: true,
	})

	/* 135 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2495
		 调整:

		 中文名: 白海豚
		 卡片种类: 通常怪兽
		 卡片密码: 92409659
		 种族: 鱼
		 属性: 水
		 星级: 2
		 攻击力: 500
		 防御力: 400
		 罕见度: 平卡N
		 卡包: Booster02
		 效果: 描述：头上有角的白色海豚。引起大浪攻击对方。

		*/
		Id:       2495,
		Password: "92409659",
		Name:     "白海豚",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Fish,  // 鱼
		Atk:     500,
		Def:     400,
		IsValid: true,
	})

	/* 136 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2496
		 调整:

		 中文名: 梅基拉斯之光
		 卡片种类: 通常怪兽
		 卡片密码: 23032273
		 种族: 恶魔
		 属性: 暗
		 星级: 3
		 攻击力: 900
		 防御力: 600
		 罕见度: 平卡N
		 卡包: Booster02、Booster R1
		 效果: 描述：从眼中放出恶光给予对方损伤。

		*/
		Id:       2496,
		Password: "23032273",
		Name:     "梅基拉斯之光",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     900,
		Def:     600,
		IsValid: true,
	})

	/* 137 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2497
		 调整:

		 中文名: 心锁妖精
		 卡片种类: 通常怪兽
		 卡片密码: 20541432
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 1050
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: Booster02、Booster R1
		 效果: 描述：封印心中的良知，使对方成为恶魔的手下。

		*/
		Id:       2497,
		Password: "20541432",
		Name:     "心锁妖精",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     1050,
		Def:     1200,
		IsValid: true,
	})

	/* 138 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2498
		 调整:

		 中文名: 翼龙
		 卡片种类: 通常怪兽
		 卡片密码: 57405307
		 种族: 鸟兽
		 属性: 风
		 星级: 4
		 攻击力: 1200
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: Booster02、Booster Chronicle、Booster R1
		 效果: 描述：挥动翅膀卷起强有力的龙卷风。

		*/
		Id:       2498,
		Password: "57405307",
		Name:     "翼龙",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Wind,        // 风
		Lr:      ygo.LR_WingedBeast, // 鸟兽
		Atk:     1200,
		Def:     1000,
		IsValid: true,
	})

	/* 139 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2502
		 调整:

		 中文名: 赤剑之莱蒙多斯
		 卡片种类: 通常怪兽
		 卡片密码: 62403074
		 种族: 战士
		 属性: 地
		 星级: 4
		 攻击力: 1200
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: Booster03
		 效果: 描述：拿着红色火炎之剑的战士。利用炎之束缚封住对方的行动。

		*/
		Id:       2502,
		Password: "62403074",
		Name:     "赤剑之莱蒙多斯",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     1200,
		Def:     1300,
		IsValid: true,
	})

	/* 140 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2503
		 调整:

		 中文名: 有生命的花瓶
		 卡片种类: 通常怪兽
		 卡片密码: 34320307
		 种族: 植物
		 属性: 地
		 星级: 3
		 攻击力: 900
		 防御力: 1100
		 罕见度: 平卡N
		 卡包: Booster03
		 效果: 描述：利用花散发花粉并咬向对方的活花瓶。

		*/
		Id:       2503,
		Password: "34320307",
		Name:     "有生命的花瓶",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Plant, // 植物
		Atk:     900,
		Def:     1100,
		IsValid: true,
	})

	/* 141 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2504
		 调整:

		 中文名: 食命者
		 卡片种类: 通常怪兽
		 卡片密码: 52367652
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 1200
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: Booster03
		 效果: 描述：嗜食所有生物的灵魂，将将其作为自己的能量。

		*/
		Id:       2504,
		Password: "52367652",
		Name:     "食命者",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     1200,
		Def:     1000,
		IsValid: true,
	})

	/* 142 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2505
		 调整:

		 中文名: 岩之战士
		 卡片种类: 通常怪兽
		 卡片密码: 46864967
		 种族: 岩石
		 属性: 地
		 星级: 4
		 攻击力: 1300
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: Booster03、Booster R2
		 效果: 描述：岩石战士。挥舞着非常重的石头剑。

		*/
		Id:       2505,
		Password: "46864967",
		Name:     "岩之战士",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Rock,  // 岩石
		Atk:     1300,
		Def:     1200,
		IsValid: true,
	})

	/* 143 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2506
		 调整:

		 中文名: 音女
		 卡片种类: 通常怪兽
		 卡片密码: 38942059
		 种族: 战士
		 属性: 地
		 星级: 3
		 攻击力: 1200
		 防御力: 900
		 罕见度: 平卡N
		 卡包: Booster03、Booster Chronicle、Booster R2
		 效果: 描述：擅长操纵声音的少女，能使音符变成镰刀攻击敌人。

		*/
		Id:       2506,
		Password: "38942059",
		Name:     "音女",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     1200,
		Def:     900,
		IsValid: true,
	})

	/* 144 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2507
		 调整:

		 中文名: 大白鲨
		 卡片种类: 通常怪兽
		 卡片密码: 13429800
		 种族: 鱼
		 属性: 水
		 星级: 4
		 攻击力: 1600
		 防御力: 800
		 罕见度: 平卡N
		 卡包: EX-R(EX)、Booster03、Booster R2
		 效果: 描述：巨大的白鲨，若被咬到绝对无法脱身。

		*/
		Id:       2507,
		Password: "13429800",
		Name:     "大白鲨",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Fish,  // 鱼
		Atk:     1600,
		Def:     800,
		IsValid: true,
	})

	/* 145 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2508
		 调整:

		 中文名: 锹甲阿尔法
		 卡片种类: 通常怪兽
		 卡片密码: 60802233
		 种族: 昆虫
		 属性: 地
		 星级: 4
		 攻击力: 1250
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: Booster03、Booster Chronicle、Booster R2、TP20
		 效果: 描述：凶暴的甲虫。它会瞄准对方的头瞬间斩落。

		*/
		Id:       2508,
		Password: "60802233",
		Name:     "锹甲阿尔法",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth,  // 地
		Lr:      ygo.LR_Insect, // 昆虫
		Atk:     1250,
		Def:     1000,
		IsValid: true,
	})

	/* 146 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2509
		 调整:

		 中文名: 钢铁魔人
		 卡片种类: 通常怪兽
		 卡片密码: 07526150
		 种族: 机械
		 属性: 地
		 星级: 4
		 攻击力: 900
		 防御力: 1600
		 罕见度: 平卡N
		 卡包: Booster03、Booster R2
		 效果: 描述：从与异次元相通的洞里出来的钢铁大魔人。

		*/
		Id:       2509,
		Password: "07526150",
		Name:     "钢铁魔人",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Machine, // 机械
		Atk:     900,
		Def:     1600,
		IsValid: true,
	})

	/* 147 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2510
		 调整:

		 中文名: 青玉眼怪
		 卡片种类: 通常怪兽
		 卡片密码: 55210709
		 种族: 兽
		 属性: 地
		 星级: 4
		 攻击力: 1300
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: Booster03
		 效果: 描述：蓝宝石眼睛的野兽。制造幻影，趁敌人混乱的时候进行攻击。

		*/
		Id:       2510,
		Password: "55210709",
		Name:     "青玉眼怪",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Atk:     1300,
		Def:     1300,
		IsValid: true,
	})

	/* 148 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2511
		 调整:

		 中文名: 斩首的美女
		 卡片种类: 通常怪兽
		 卡片密码: 16899564
		 种族: 战士
		 属性: 地
		 星级: 4
		 攻击力: 1600
		 防御力: 800
		 罕见度: 平卡N
		 卡包: Booster03、Booster Chronicle、Booster R2
		 效果: 描述：在那美丽容貌的背后，却是个用刀子使许多人身首异处的女子。

		*/
		Id:       2511,
		Password: "16899564",
		Name:     "斩首的美女",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     1600,
		Def:     800,
		IsValid: true,
	})

	/* 149 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2512
		 调整:

		 中文名: 海马兽
		 卡片种类: 通常怪兽
		 卡片密码: 47922711
		 种族: 兽
		 属性: 地
		 星级: 5
		 攻击力: 1350
		 防御力: 1600
		 罕见度: 平卡N
		 卡包: Booster03、Booster R2
		 效果: 描述：半马半鱼的怪兽。以风一般的速度在海中奔驰。

		*/
		Id:       2512,
		Password: "47922711",
		Name:     "海马兽",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Atk:     1350,
		Def:     1600,
		IsValid: true,
	})

	/* 150 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2513
		 调整:

		 中文名: 审判之手
		 卡片种类: 通常怪兽
		 卡片密码: 28003512
		 种族: 战士
		 属性: 地
		 星级: 3
		 攻击力: 1400
		 防御力: 700
		 罕见度: 平卡N
		 卡包: Booster03、Booster Chronicle、Booster R2
		 效果: 描述：用寄宿着神灵的手作最后的判决，给予敌人以猛烈的攻击。

		*/
		Id:       2513,
		Password: "28003512",
		Name:     "审判之手",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     1400,
		Def:     700,
		IsValid: true,
	})

	/* 151 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2514
		 调整:

		 中文名: 骷髅寺院
		 卡片种类: 通常怪兽
		 卡片密码: 00732302
		 种族: 不死
		 属性: 暗
		 星级: 4
		 攻击力: 900
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: Booster03
		 效果: 描述：全部由骷髅和骨头建立而成的寺院。会吸收接近的人的灵魂。

		*/
		Id:       2514,
		Password: "00732302",
		Name:     "骷髅寺院",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Dark,   // 暗
		Lr:      ygo.LR_Zombie, // 不死
		Atk:     900,
		Def:     1300,
		IsValid: true,
	})

	/* 152 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2515
		 调整:

		 中文名: 树精
		 卡片种类: 通常怪兽
		 卡片密码: 84916669
		 种族: 魔法师
		 属性: 地
		 星级: 4
		 攻击力: 1200
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: Booster03、TP07
		 效果: 描述：森之精灵。借助草木之力封住对方的行动。

		*/
		Id:       2515,
		Password: "84916669",
		Name:     "树精",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth,       // 地
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     1200,
		Def:     1400,
		IsValid: true,
	})

	/* 153 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2516
		 调整:

		 中文名: 狂战士
		 卡片种类: 通常怪兽
		 卡片密码: 47060154
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 1500
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: EX-R(EX)、Booster03、Booster R2
		 效果: 用狂暴的力量攻击，一旦暴走无人能挡。

		*/
		Id:       2516,
		Password: "47060154",
		Name:     "狂战士",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     1500,
		Def:     1000,
		IsValid: true,
	})

	/* 154 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2517
		 调整:

		 中文名: 复仇的河童
		 卡片种类: 通常怪兽
		 卡片密码: 48109103
		 种族: 水
		 属性: 水
		 星级: 3
		 攻击力: 1200
		 防御力: 900
		 罕见度: 平卡N
		 卡包: Booster03
		 效果: 描述：被同伴杀害，为了复仇而将心灵卖给邪恶的河童。

		*/
		Id:       2517,
		Password: "48109103",
		Name:     "复仇的河童",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     1200,
		Def:     900,
		IsValid: true,
	})

	/* 155 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2518
		 调整:

		 中文名: 大力独角仙
		 卡片种类: 通常怪兽
		 卡片密码: 52584282
		 种族: 昆虫
		 属性: 地
		 星级: 5
		 攻击力: 1500
		 防御力: 2000
		 罕见度: 平卡N
		 卡包: Booster03、Booster Chronicle、Booster R2、TP20
		 效果: 描述：巨大的甲虫，锐利的角与坚硬的身体非常强力。

		*/
		Id:       2518,
		Password: "52584282",
		Name:     "大力独角仙",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Earth,  // 地
		Lr:      ygo.LR_Insect, // 昆虫
		Atk:     1500,
		Def:     2000,
		IsValid: true,
	})

	/* 156 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2519
		 调整:

		 中文名: 书之精灵 飞鹰主教
		 卡片种类: 通常怪兽
		 卡片密码: 14037717
		 种族: 鸟兽
		 属性: 风
		 星级: 4
		 攻击力: 1400
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: Booster03、Booster Chronicle、Booster R2
		 效果: 描述：木之精灵，拥有非常高的智慧，发动多彩的攻击。

		*/
		Id:       2519,
		Password: "14037717",
		Name:     "书之精灵 飞鹰主教",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Wind,        // 风
		Lr:      ygo.LR_WingedBeast, // 鸟兽
		Atk:     1400,
		Def:     1200,
		IsValid: true,
	})

	/* 157 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2520
		 调整:

		 中文名: 魔界的机械兵
		 卡片种类: 通常怪兽
		 卡片密码: 75559356
		 种族: 机械
		 属性: 暗
		 星级: 4
		 攻击力: 1400
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: Booster03、Booster Chronicle、Booster R2
		 效果: 描述：暗之力制造的机器士兵，疯狂的将敌人破坏。

		*/
		Id:       2520,
		Password: "75559356",
		Name:     "魔界的机械兵",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Dark,    // 暗
		Lr:      ygo.LR_Machine, // 机械
		Atk:     1400,
		Def:     1200,
		IsValid: true,
	})

	/* 158 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2521
		 调整:

		 中文名: 缪斯天使
		 卡片种类: 通常怪兽
		 卡片密码: 69992868
		 种族: 天使
		 属性: 光
		 星级: 3
		 攻击力: 850
		 防御力: 900
		 罕见度: 平卡N
		 卡包: Booster03、TP16
		 效果: 描述：艺术家的天使。尤其琴技无人能出其右。

		*/
		Id:       2521,
		Password: "69992868",
		Name:     "缪斯天使",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Light, // 光
		Lr:      ygo.LR_Fairy, // 天使
		Atk:     850,
		Def:     900,
		IsValid: true,
	})

	/* 159 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2522
		 调整:

		 中文名: 蒙·拉瓦斯
		 卡片种类: 通常怪兽
		 卡片密码: 07225792
		 种族: 兽
		 属性: 地
		 星级: 4
		 攻击力: 1300
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: Booster03
		 效果: 描述：进化过的野兽。力量有相当的提高。

		*/
		Id:       2522,
		Password: "07225792",
		Name:     "蒙·拉瓦斯",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Atk:     1300,
		Def:     1400,
		IsValid: true,
	})

	/* 160 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2523
		 调整:

		 中文名: 椰树
		 卡片种类: 通常怪兽
		 卡片密码: 41061625
		 种族: 植物
		 属性: 地
		 星级: 2
		 攻击力: 800
		 防御力: 600
		 罕见度: 平卡N
		 卡包: Booster03
		 效果: 描述：有意志的椰子树。以落下的果实进行攻击。果实中的椰汁味道鲜美。

		*/
		Id:       2523,
		Password: "41061625",
		Name:     "椰树",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Plant, // 植物
		Atk:     800,
		Def:     600,
		IsValid: true,
	})

	/* 161 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2524
		 调整:

		 中文名: 尤蒙刚德
		 卡片种类: 通常怪兽
		 卡片密码: 17115745
		 种族: 爬虫类
		 属性: 地
		 星级: 3
		 攻击力: 1200
		 防御力: 900
		 罕见度: 平卡N
		 卡包: Booster03
		 效果: 描述：出现在神话世界的蛇。非常的长。

		*/
		Id:       2524,
		Password: "17115745",
		Name:     "尤蒙刚德",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Reptile, // 爬虫类
		Atk:     1200,
		Def:     900,
		IsValid: true,
	})

	/* 162 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2525
		 调整:

		 中文名: 月之魔法师
		 卡片种类: 通常怪兽
		 卡片密码: 75850803
		 种族: 魔法师
		 属性: 光
		 星级: 5
		 攻击力: 1200
		 防御力: 1700
		 罕见度: 平卡N
		 卡包: Booster03
		 效果: 描述：住在月亮上的魔法师。用月之魔力迷惑对方。

		*/
		Id:       2525,
		Password: "75850803",
		Name:     "月之魔法师",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Light,       // 光
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     1200,
		Def:     1700,
		IsValid: true,
	})

	/* 163 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2526
		 调整:

		 中文名: 狮子男巫
		 卡片种类: 通常怪兽
		 卡片密码: 04392470
		 种族: 魔法师
		 属性: 地
		 星级: 5
		 攻击力: 1350
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: Booster03、Booster R2
		 效果: 描述：披着黑披风的魔术师。其正体是只能说话的狮子。

		*/
		Id:       2526,
		Password: "04392470",
		Name:     "狮子男巫",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Earth,       // 地
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     1350,
		Def:     1200,
		IsValid: true,
	})

	/* 164 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2533
		 调整:

		 中文名: 龟鸟
		 卡片种类: 通常怪兽
		 卡片密码: 72929454
		 种族: 水
		 属性: 水
		 星级: 6
		 攻击力: 1900
		 防御力: 1700
		 罕见度: 平卡N
		 卡包: Booster04、Booster Chronicle、Booster R2
		 效果: 描述：主要栖息在水中，也能在空中飞翔的珍奇乌龟。

		*/
		Id:       2533,
		Password: "72929454",
		Name:     "龟鸟",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   6,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     1900,
		Def:     1700,
		IsValid: true,
	})

	/* 165 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2534
		 调整:

		 中文名: 电子鱼
		 卡片种类: 通常怪兽
		 卡片密码: 50176820
		 种族: 机械
		 属性: 水
		 星级: 5
		 攻击力: 1800
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: Booster04、Booster R2
		 效果: 描述：以背后附着的炮台发射闪光粒子加农炮。

		*/
		Id:       2534,
		Password: "50176820",
		Name:     "电子鱼",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Water,   // 水
		Lr:      ygo.LR_Machine, // 机械
		Atk:     1800,
		Def:     1500,
		IsValid: true,
	})

	/* 166 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2535
		 调整:

		 中文名: 水陆的帝王
		 卡片种类: 通常怪兽
		 卡片密码: 11250655
		 种族: 爬虫类
		 属性: 水
		 星级: 5
		 攻击力: 1800
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: Booster04、Booster R2
		 效果: 描述：可以用大嘴向四方喷火的爬虫怪。

		*/
		Id:       2535,
		Password: "11250655",
		Name:     "水陆的帝王",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Water,   // 水
		Lr:      ygo.LR_Reptile, // 爬虫类
		Atk:     1800,
		Def:     1500,
		IsValid: true,
	})

	/* 167 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2536
		 调整:

		 中文名: 红叶之女王
		 卡片种类: 通常怪兽
		 卡片密码: 04179849
		 种族: 植物
		 属性: 地
		 星级: 5
		 攻击力: 1800
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: Booster04、Booster06、Booster Chronicle、Booster R2
		 效果: 描述：生活在被红叶围绕的地方，绿树灵王的妃子。

		*/
		Id:       2536,
		Password: "04179849",
		Name:     "红叶之女王",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Plant, // 植物
		Atk:     1800,
		Def:     1500,
		IsValid: true,
	})

	/* 168 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2537
		 调整:

		 中文名: 战神 奥利安
		 卡片种类: 通常怪兽
		 卡片密码: 02971090
		 种族: 天使
		 属性: 光
		 星级: 5
		 攻击力: 1800
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: Booster04、Booster R2
		 效果: 描述：被誉为战神的天使。不过谁也没见过那场战事。

		*/
		Id:       2537,
		Password: "02971090",
		Name:     "战神 奥利安",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Light, // 光
		Lr:      ygo.LR_Fairy, // 天使
		Atk:     1800,
		Def:     1500,
		IsValid: true,
	})

	/* 169 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2538
		 调整:

		 中文名: 山之精灵
		 卡片种类: 通常怪兽
		 卡片密码: 34690519
		 种族: 魔法师
		 属性: 地
		 星级: 5
		 攻击力: 1300
		 防御力: 1800
		 罕见度: 平卡N
		 卡包: Booster04、Booster R2
		 效果: 描述：据说听了它的笛音的人会失去力量。

		*/
		Id:       2538,
		Password: "34690519",
		Name:     "山之精灵",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Earth,       // 地
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     1300,
		Def:     1800,
		IsValid: true,
	})

	/* 170 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2539
		 调整:

		 中文名: 诞生的天使
		 卡片种类: 通常怪兽
		 卡片密码: 42418084
		 种族: 天使
		 属性: 光
		 星级: 5
		 攻击力: 1400
		 防御力: 1700
		 罕见度: 平卡N
		 卡包: Booster04、Booster R2
		 效果: 描述：据说它能知道女性腹部是否存在着生命。

		*/
		Id:       2539,
		Password: "42418084",
		Name:     "诞生的天使",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Light, // 光
		Lr:      ygo.LR_Fairy, // 天使
		Atk:     1400,
		Def:     1700,
		IsValid: true,
	})

	/* 171 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2540
		 调整:

		 中文名: 巨大蚂蚁
		 卡片种类: 通常怪兽
		 卡片密码: 53606874
		 种族: 昆虫
		 属性: 地
		 星级: 4
		 攻击力: 1200
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: Booster04
		 效果: 描述：住在密林里的巨大蚂蚁。攻击和守备都意外的强。

		*/
		Id:       2540,
		Password: "53606874",
		Name:     "巨大蚂蚁",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth,  // 地
		Lr:      ygo.LR_Insect, // 昆虫
		Atk:     1200,
		Def:     1500,
		IsValid: true,
	})

	/* 172 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2541
		 调整:

		 中文名: 圣鸟
		 卡片种类: 通常怪兽
		 卡片密码: 75582395
		 种族: 鸟兽
		 属性: 风
		 星级: 4
		 攻击力: 1500
		 防御力: 1100
		 罕见度: 平卡N
		 卡包: Booster04、Booster Chronicle、Booster R2
		 效果: 描述：尾巴非常长的鸟，全身散发出神圣的光芒。

		*/
		Id:       2541,
		Password: "75582395",
		Name:     "圣鸟",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Wind,        // 风
		Lr:      ygo.LR_WingedBeast, // 鸟兽
		Atk:     1500,
		Def:     1100,
		IsValid: true,
	})

	/* 173 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2542
		 调整:

		 中文名: 恶魔乌贼
		 卡片种类: 通常怪兽
		 卡片密码: 77456781
		 种族: 水
		 属性: 水
		 星级: 4
		 攻击力: 1200
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: Booster04
		 效果: 描述：潜在海里的巨大章鱼。从海中突然出现并攻击。

		*/
		Id:       2542,
		Password: "77456781",
		Name:     "恶魔乌贼",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     1200,
		Def:     1400,
		IsValid: true,
	})

	/* 174 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2543
		 调整:

		 中文名: 海原的女战士
		 卡片种类: 通常怪兽
		 卡片密码: 17968114
		 种族: 鱼
		 属性: 水
		 星级: 4
		 攻击力: 1300
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: Booster04
		 效果: 描述：侍奉海神的美丽人鱼战士，守护着神圣的区域。

		*/
		Id:       2543,
		Password: "17968114",
		Name:     "海原的女战士",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Fish,  // 鱼
		Atk:     1300,
		Def:     1400,
		IsValid: true,
	})

	/* 175 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2544
		 调整:

		 中文名: 杀人污泥
		 卡片种类: 通常怪兽
		 卡片密码: 65623423
		 种族: 水
		 属性: 水
		 星级: 3
		 攻击力: 1300
		 防御力: 700
		 罕见度: 平卡N
		 卡包: Booster04
		 效果: 描述：史莱姆们的老大。尽管样子上没什么区别，但别小看它，否则可是会吃亏的哦。

		*/
		Id:       2544,
		Password: "65623423",
		Name:     "杀人污泥",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     1300,
		Def:     700,
		IsValid: true,
	})

	/* 176 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2545
		 调整:

		 中文名: 剪刀勇士
		 卡片种类: 通常怪兽
		 卡片密码: 74277583
		 种族: 机械
		 属性: 暗
		 星级: 4
		 攻击力: 1300
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: Booster04
		 效果: 描述：熟练地利用几个夹子切向对方。

		*/
		Id:       2545,
		Password: "74277583",
		Name:     "剪刀勇士",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Dark,    // 暗
		Lr:      ygo.LR_Machine, // 机械
		Atk:     1300,
		Def:     1000,
		IsValid: true,
	})

	/* 177 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2546
		 调整:

		 中文名: 深海切割手
		 卡片种类: 通常怪兽
		 卡片密码: 71746462
		 种族: 水
		 属性: 水
		 星级: 4
		 攻击力: 1100
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: Booster04
		 效果: 描述：以锐利的钩爪撕裂对方的残忍怪兽。

		*/
		Id:       2546,
		Password: "71746462",
		Name:     "深海切割手",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     1100,
		Def:     1300,
		IsValid: true,
	})

	/* 178 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2547
		 调整:

		 中文名: 机枪岩
		 卡片种类: 通常怪兽
		 卡片密码: 10476868
		 种族: 岩石
		 属性: 地
		 星级: 4
		 攻击力: 1000
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: Booster04
		 效果: 描述：一边用双肩的机关枪发射，一边用身体乱撞进行攻击。

		*/
		Id:       2547,
		Password: "10476868",
		Name:     "机枪岩",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Rock,  // 岩石
		Atk:     1000,
		Def:     1300,
		IsValid: true,
	})

	/* 179 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2548
		 调整:

		 中文名: 暗黑魔神 梦魇
		 卡片种类: 通常怪兽
		 卡片密码: 89494469
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 1300
		 防御力: 1100
		 罕见度: 平卡N
		 卡包: EX-R(EX)、Booster04
		 效果: 描述：据说这个恶魔潜住在梦中，当猎物睡觉时便夺去其生命。

		*/
		Id:       2548,
		Password: "89494469",
		Name:     "暗黑魔神 梦魇",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     1300,
		Def:     1100,
		IsValid: true,
	})

	/* 180 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2549
		 调整:

		 中文名: 戈耳工蛋
		 卡片种类: 通常怪兽
		 卡片密码: 11793047
		 种族: 恶魔
		 属性: 暗
		 星级: 3
		 攻击力: 300
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: Booster04
		 效果: 描述：蛇发女所产下的巨卵。其他一切资料不详。

		*/
		Id:       2549,
		Password: "11793047",
		Name:     "戈耳工蛋",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     300,
		Def:     1300,
		IsValid: true,
	})

	/* 181 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2550
		 调整:

		 中文名: 妖精龙
		 卡片种类: 通常怪兽
		 卡片密码: 20315854
		 种族: 龙
		 属性: 风
		 星级: 4
		 攻击力: 1100
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: Booster04、TP02
		 效果: 描述：在妖精之中出乎意料地强，非常美丽的龙之妖精。

		*/
		Id:       2550,
		Password: "20315854",
		Name:     "妖精龙",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Wind,   // 风
		Lr:      ygo.LR_Dragon, // 龙
		Atk:     1100,
		Def:     1200,
		IsValid: true,
	})

	/* 182 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2557
		 调整:

		 中文名: 飞鹰
		 卡片种类: 通常怪兽
		 卡片密码: 47319141
		 种族: 鸟兽
		 属性: 风
		 星级: 5
		 攻击力: 1800
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: Booster05、Booster Chronicle、Booster R3
		 效果: 描述：从高空寻找猎物，一旦看上绝不让其逃走。

		*/
		Id:       2557,
		Password: "47319141",
		Name:     "飞鹰",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Wind,        // 风
		Lr:      ygo.LR_WingedBeast, // 鸟兽
		Atk:     1800,
		Def:     1500,
		IsValid: true,
	})

	/* 183 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2558
		 调整:

		 中文名: 机器攻击者
		 卡片种类: 通常怪兽
		 卡片密码: 38116136
		 种族: 机械
		 属性: 地
		 星级: 5
		 攻击力: 1600
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: Booster05
		 效果: 描述：特攻用的机械。以突击之力打倒敌人。

		*/
		Id:       2558,
		Password: "38116136",
		Name:     "机器攻击者",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Machine, // 机械
		Atk:     1600,
		Def:     1300,
		IsValid: true,
	})

	/* 184 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2559
		 调整:

		 中文名: 贪尸龙
		 卡片种类: 通常怪兽
		 卡片密码: 38289717
		 种族: 恐龙
		 属性: 地
		 星级: 4
		 攻击力: 1600
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: Booster05、Booster Chronicle、Booster R3
		 效果: 描述：什么都能咬碎的恐龙，攻击力十分恐怖。

		*/
		Id:       2559,
		Password: "38289717",
		Name:     "贪尸龙",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth,    // 地
		Lr:      ygo.LR_Dinosaur, // 恐龙
		Atk:     1600,
		Def:     1200,
		IsValid: true,
	})

	/* 185 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2560
		 调整:

		 中文名: 彩虹人鱼
		 卡片种类: 通常怪兽
		 卡片密码: 29402771
		 种族: 鱼
		 属性: 水
		 星级: 5
		 攻击力: 1550
		 防御力: 1700
		 罕见度: 平卡N
		 卡包: Booster05
		 效果: 描述：空中出现彩虹桥时才会出现的珍奇美丽人鱼。

		*/
		Id:       2560,
		Password: "29402771",
		Name:     "彩虹人鱼",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Fish,  // 鱼
		Atk:     1550,
		Def:     1700,
		IsValid: true,
	})

	/* 186 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2561
		 调整:

		 中文名: 美杜莎的亡灵
		 卡片种类: 通常怪兽
		 卡片密码: 29491031
		 种族: 不死
		 属性: 暗
		 星级: 4
		 攻击力: 1500
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: Booster05、Booster R3
		 效果: 描述：头发满是毒蛇的怪兽。一旦被她的目光盯上，就会被石化。

		*/
		Id:       2561,
		Password: "29491031",
		Name:     "美杜莎的亡灵",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Dark,   // 暗
		Lr:      ygo.LR_Zombie, // 不死
		Atk:     1500,
		Def:     1200,
		IsValid: true,
	})

	/* 187 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2562
		 调整:

		 中文名: 水之魔导师
		 卡片种类: 通常怪兽
		 卡片密码: 93343894
		 种族: 水
		 属性: 水
		 星级: 4
		 攻击力: 1400
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: Booster05、Booster Chronicle、Booster R3
		 效果: 描述：用水裹住对方并包围着施以攻击。

		*/
		Id:       2562,
		Password: "93343894",
		Name:     "水之魔导师",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     1400,
		Def:     1000,
		IsValid: true,
	})

	/* 188 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2563
		 调整:

		 中文名: 大肚海蛇
		 卡片种类: 通常怪兽
		 卡片密码: 94022093
		 种族: 水
		 属性: 水
		 星级: 4
		 攻击力: 1350
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: Booster05、Booster R3
		 效果: 描述：样子很怪的海蛇。特征是大口和大牙。

		*/
		Id:       2563,
		Password: "94022093",
		Name:     "大肚海蛇",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     1350,
		Def:     1000,
		IsValid: true,
	})

	/* 189 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2564
		 调整:

		 中文名: 3万年的白龟
		 卡片种类: 通常怪兽
		 卡片密码: 11714098
		 种族: 水
		 属性: 水
		 星级: 5
		 攻击力: 1250
		 防御力: 2100
		 罕见度: 平卡N
		 卡包: Booster05
		 效果: 描述：生活了三万年的巨大乌龟。守备力高。

		*/
		Id:       2564,
		Password: "11714098",
		Name:     "3万年的白龟",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     1250,
		Def:     2100,
		IsValid: true,
	})

	/* 190 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2565
		 调整:

		 中文名: 狼
		 卡片种类: 通常怪兽
		 卡片密码: 49417509
		 种族: 兽
		 属性: 地
		 星级: 3
		 攻击力: 1200
		 防御力: 800
		 罕见度: 平卡N
		 卡包: Booster05
		 效果: 描述：比较难见到的狼。用灵敏的鼻子寻找猎物。

		*/
		Id:       2565,
		Password: "49417509",
		Name:     "狼",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Atk:     1200,
		Def:     800,
		IsValid: true,
	})

	/* 191 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2566
		 调整:

		 中文名: 鳄鱼人
		 卡片种类: 通常怪兽
		 卡片密码: 76512652
		 种族: 爬虫类
		 属性: 水
		 星级: 4
		 攻击力: 1100
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: Booster05
		 效果: 描述：有智慧而且狂暴化的鳄鱼。以坚硬的鳞片反弹对方的攻击。

		*/
		Id:       2566,
		Password: "76512652",
		Name:     "鳄鱼人",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Water,   // 水
		Lr:      ygo.LR_Reptile, // 爬虫类
		Atk:     1100,
		Def:     1200,
		IsValid: true,
	})

	/* 192 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2567
		 调整:

		 中文名: 水蛇
		 卡片种类: 通常怪兽
		 卡片密码: 12436646
		 种族: 水
		 属性: 水
		 星级: 3
		 攻击力: 1050
		 防御力: 900
		 罕见度: 平卡N
		 卡包: Booster05
		 效果: 描述：尾巴前端的光球可施展催眠术，使对方精神恍惚。

		*/
		Id:       2567,
		Password: "12436646",
		Name:     "水蛇",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     1050,
		Def:     900,
		IsValid: true,
	})

	/* 193 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2568
		 调整:

		 中文名: 气象精灵
		 卡片种类: 通常怪兽
		 卡片密码: 96643568
		 种族: 水
		 属性: 水
		 星级: 3
		 攻击力: 1000
		 防御力: 900
		 罕见度: 平卡N
		 卡包: Booster05
		 效果: 描述：操纵雨的精灵。召唤台风吹飞各种东西。

		*/
		Id:       2568,
		Password: "96643568",
		Name:     "气象精灵",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     1000,
		Def:     900,
		IsValid: true,
	})

	/* 194 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2569
		 调整:

		 中文名: 来自异次元的侵略者
		 卡片种类: 通常怪兽
		 卡片密码: 28450915
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 950
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: Booster05
		 效果: 描述：从银河系以外来到地球的外星人。

		*/
		Id:       2569,
		Password: "28450915",
		Name:     "来自异次元的侵略者",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     950,
		Def:     1400,
		IsValid: true,
	})

	/* 195 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2570
		 调整:

		 中文名: 鸟嘴蛇
		 卡片种类: 通常怪兽
		 卡片密码: 06103114
		 种族: 爬虫类
		 属性: 地
		 星级: 3
		 攻击力: 800
		 防御力: 900
		 罕见度: 平卡N
		 卡包: Booster05
		 效果: 描述：用长长的身体把对方缠起来，用大嘴进行攻击。

		*/
		Id:       2570,
		Password: "06103114",
		Name:     "鸟嘴蛇",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Reptile, // 爬虫类
		Atk:     800,
		Def:     900,
		IsValid: true,
	})

	/* 196 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2571
		 调整:

		 中文名: 龟狸
		 卡片种类: 通常怪兽
		 卡片密码: 17441953
		 种族: 水
		 属性: 水
		 星级: 3
		 攻击力: 700
		 防御力: 900
		 罕见度: 平卡N
		 卡包: Booster05
		 效果: 描述：背着乌龟甲壳的狸猫。诱惑敌人进行攻击。

		*/
		Id:       2571,
		Password: "17441953",
		Name:     "龟狸",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     700,
		Def:     900,
		IsValid: true,
	})

	/* 197 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2572
		 调整:

		 中文名: 屎壳郎
		 卡片种类: 通常怪兽
		 卡片密码: 32569498
		 种族: 昆虫
		 属性: 地
		 星级: 2
		 攻击力: 550
		 防御力: 400
		 罕见度: 平卡N
		 卡包: Booster05
		 效果: 描述：可以推动比自己身体大几倍的粪便，压倒对方。

		*/
		Id:       2572,
		Password: "32569498",
		Name:     "屎壳郎",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Earth,  // 地
		Lr:      ygo.LR_Insect, // 昆虫
		Atk:     550,
		Def:     400,
		IsValid: true,
	})

	/* 198 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2573
		 调整:

		 中文名: 藤蔓植物
		 卡片种类: 通常怪兽
		 卡片密码: 60715406
		 种族: 植物
		 属性: 水
		 星级: 2
		 攻击力: 500
		 防御力: 600
		 罕见度: 平卡N
		 卡包: Booster05
		 效果: 描述：当附近有会动的东西时，就会伸出蓝色的丝线攻击。

		*/
		Id:       2573,
		Password: "60715406",
		Name:     "藤蔓植物",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Plant, // 植物
		Atk:     500,
		Def:     600,
		IsValid: true,
	})

	/* 199 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2574
		 调整:

		 中文名: 戏法壶
		 卡片种类: 通常怪兽
		 卡片密码: 55567161
		 种族: 岩石
		 属性: 地
		 星级: 2
		 攻击力: 400
		 防御力: 400
		 罕见度: 平卡N
		 卡包: Booster05
		 效果: 描述：根据魔术师的命令行动的使魔。不是很强。

		*/
		Id:       2574,
		Password: "55567161",
		Name:     "戏法壶",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Rock,  // 岩石
		Atk:     400,
		Def:     400,
		IsValid: true,
	})

	/* 200 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2575
		 调整:

		 中文名: 魔界植物
		 卡片种类: 通常怪兽
		 卡片密码: 53830602
		 种族: 恶魔
		 属性: 暗
		 星级: 1
		 攻击力: 400
		 防御力: 300
		 罕见度: 平卡N
		 卡包: Booster05
		 效果: 描述：从割面发出强力的酸液融化接近的人。

		*/
		Id:       2575,
		Password: "53830602",
		Name:     "魔界植物",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   1,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     400,
		Def:     300,
		IsValid: true,
	})

	/* 201 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2580
		 调整:

		 中文名: 牛鬼
		 卡片种类: 通常怪兽
		 卡片密码: 48649353
		 种族: 恶魔
		 属性: 暗
		 星级: 6
		 攻击力: 2150
		 防御力: 1950
		 罕见度: 平卡N
		 卡包: Booster06、Booster Chronicle、Booster R3、TP11
		 效果: 描述：通过黑魔术苏醒的牛之恶魔。从壶中现身。

		*/
		Id:       2580,
		Password: "48649353",
		Name:     "牛鬼",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   6,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     2150,
		Def:     1950,
		IsValid: true,
	})

	/* 202 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2581
		 调整:

		 中文名: 加尔瓦斯
		 卡片种类: 通常怪兽
		 卡片密码: 69780745
		 种族: 兽
		 属性: 地
		 星级: 6
		 攻击力: 2000
		 防御力: 1700
		 罕见度: 平卡N
		 卡包: Booster06
		 效果: 描述：恶之化身。样子如同长着羽毛的狮子。

		*/
		Id:       2581,
		Password: "69780745",
		Name:     "加尔瓦斯",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   6,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Atk:     2000,
		Def:     1700,
		IsValid: true,
	})

	/* 203 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2582
		 调整:

		 中文名: 鹦鹉龙
		 卡片种类: 通常怪兽
		 卡片密码: 62762898
		 种族: 龙
		 属性: 风
		 星级: 5
		 攻击力: 2000
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: Booster06、Booster Chronicle、Booster R3、PE
		 效果: 描述：美国喜剧中的龙，不要被可爱的外表所迷惑。

		*/
		Id:       2582,
		Password: "62762898",
		Name:     "鹦鹉龙",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Wind,   // 风
		Lr:      ygo.LR_Dragon, // 龙
		Atk:     2000,
		Def:     1300,
		IsValid: true,
	})

	/* 204 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2583
		 调整:

		 中文名: 天空龙
		 卡片种类: 通常怪兽
		 卡片密码: 95288024
		 种族: 龙
		 属性: 风
		 星级: 6
		 攻击力: 1900
		 防御力: 1800
		 罕见度: 平卡N
		 卡包: Booster06、Booster Chronicle、Booster R3
		 效果: 描述：有着四枚羽翼的鸟的外形的龙，用如刀刃一般的羽毛进行攻击。

		*/
		Id:       2583,
		Password: "95288024",
		Name:     "天空龙",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   6,
		La:      ygo.LA_Wind,   // 风
		Lr:      ygo.LR_Dragon, // 龙
		Atk:     1900,
		Def:     1800,
		IsValid: true,
	})

	/* 205 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2584
		 调整:

		 中文名: 加鲁撒斯
		 卡片种类: 通常怪兽
		 卡片密码: 14977074
		 种族: 兽战士
		 属性: 炎
		 星级: 5
		 攻击力: 1800
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: Booster06、Booster Chronicle、Booster R3、JY
		 效果: 描述：长有龙头的兽战士，斧头具有相当大的威力。

		*/
		Id:       2584,
		Password: "14977074",
		Name:     "加鲁撒斯",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Fire,         // 炎
		Lr:      ygo.LR_BeastWarrior, // 兽战士
		Atk:     1800,
		Def:     1500,
		IsValid: true,
	})

	/* 206 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2585
		 调整:

		 中文名: 剑龙
		 卡片种类: 通常怪兽
		 卡片密码: 13069066
		 种族: 恐龙
		 属性: 地
		 星级: 6
		 攻击力: 1750
		 防御力: 2030
		 罕见度: 平卡N
		 卡包: Booster06、Booster R3
		 效果: 描述：全身都附有刀刺的恐龙。以突进为手段进行攻击。

		*/
		Id:       2585,
		Password: "13069066",
		Name:     "剑龙",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   6,
		La:      ygo.LA_Earth,    // 地
		Lr:      ygo.LR_Dinosaur, // 恐龙
		Atk:     1750,
		Def:     2030,
		IsValid: true,
	})

	/* 207 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2586
		 调整:

		 中文名: 猎头魔人
		 卡片种类: 通常怪兽
		 卡片密码: 70084224
		 种族: 恶魔
		 属性: 暗
		 星级: 6
		 攻击力: 1750
		 防御力: 1900
		 罕见度: 平卡N
		 卡包: Booster06
		 效果: 描述：挥舞着大镰刀的恶魔。从大眼睛中还会发出光柱。

		*/
		Id:       2586,
		Password: "70084224",
		Name:     "猎头魔人",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   6,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     1750,
		Def:     1900,
		IsValid: true,
	})

	/* 208 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2587
		 调整:

		 中文名: 风之精灵
		 卡片种类: 通常怪兽
		 卡片密码: 54615781
		 种族: 魔法师
		 属性: 风
		 星级: 5
		 攻击力: 1700
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: Booster06
		 效果: 描述：自在飞翔的风之精灵。心情不好的时候就会变成狂风。

		*/
		Id:       2587,
		Password: "54615781",
		Name:     "风之精灵",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Wind,        // 风
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     1700,
		Def:     1400,
		IsValid: true,
	})

	/* 209 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2588
		 调整:

		 中文名: 牛头人
		 卡片种类: 通常怪兽
		 卡片密码: 05053103
		 种族: 兽战士
		 属性: 地
		 星级: 4
		 攻击力: 1700
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: EX-R(EX)、Booster06、Booster R3、KA、SK2
		 效果: 描述：有着强大力量的牛怪，斧头一挥能砍倒一切东西。

		*/
		Id:       2588,
		Password: "05053103",
		Name:     "牛头人",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth,        // 地
		Lr:      ygo.LR_BeastWarrior, // 兽战士
		Atk:     1700,
		Def:     1000,
		IsValid: true,
	})

	/* 210 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2589
		 调整:

		 中文名: 岩石之精灵
		 卡片种类: 通常怪兽
		 卡片密码: 82818645
		 种族: 魔法师
		 属性: 地
		 星级: 5
		 攻击力: 1650
		 防御力: 1900
		 罕见度: 平卡N
		 卡包: Booster06
		 效果: 描述：指轮般的岩石精灵。攻击与守备力都很强。

		*/
		Id:       2589,
		Password: "82818645",
		Name:     "岩石之精灵",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Earth,       // 地
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     1650,
		Def:     1900,
		IsValid: true,
	})

	/* 211 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2590
		 调整:

		 中文名: 塞壬
		 卡片种类: 通常怪兽
		 卡片密码: 81686058
		 种族: 魔法师
		 属性: 光
		 星级: 5
		 攻击力: 1600
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: Booster06
		 效果: 描述：能操纵风并能引起强风吹飞所有东西。

		*/
		Id:       2590,
		Password: "81686058",
		Name:     "塞壬",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Light,       // 光
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     1600,
		Def:     1500,
		IsValid: true,
	})

	/* 212 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2591
		 调整:

		 中文名: 岩石龟
		 卡片种类: 通常怪兽
		 卡片密码: 09540040
		 种族: 水
		 属性: 水
		 星级: 6
		 攻击力: 1450
		 防御力: 2200
		 罕见度: 平卡N
		 卡包: Booster06
		 效果: 描述：全身由岩石做成的龟。特点是守备力非常高。

		*/
		Id:       2591,
		Password: "09540040",
		Name:     "岩石龟",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   6,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     1450,
		Def:     2200,
		IsValid: true,
	})

	/* 213 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2592
		 调整:

		 中文名: 守卫海洋的战士
		 卡片种类: 通常怪兽
		 卡片密码: 85448931
		 种族: 水
		 属性: 水
		 星级: 4
		 攻击力: 1300
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: Booster06、Booster R3
		 效果: 描述：彻底的攻击玷污大海的家伙、人鱼战士。

		*/
		Id:       2592,
		Password: "85448931",
		Name:     "守卫海洋的战士",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     1300,
		Def:     1000,
		IsValid: true,
	})

	/* 214 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2593
		 调整:

		 中文名: 龙魂之石像
		 卡片种类: 通常怪兽
		 卡片密码: 09197735
		 种族: 战士
		 属性: 地
		 星级: 3
		 攻击力: 1100
		 防御力: 900
		 罕见度: 平卡N
		 卡包: Booster06
		 效果: 描述：拥有龙的灵魂的石像战士。用自豪的剑劈开敌人。

		*/
		Id:       2593,
		Password: "09197735",
		Name:     "龙魂之石像",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     1100,
		Def:     900,
		IsValid: true,
	})

	/* 215 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2601
		 调整:

		 中文名: 金属守护者
		 卡片种类: 通常怪兽
		 卡片密码: 68339286
		 种族: 恶魔
		 属性: 暗
		 星级: 5
		 攻击力: 1150
		 防御力: 2150
		 罕见度: 平卡N
		 卡包: Booster07
		 效果: 描述：守护魔界宝物的恶魔在黑暗中守备相当强。

		*/
		Id:       2601,
		Password: "68339286",
		Name:     "金属守护者",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     1150,
		Def:     2150,
		IsValid: true,
	})

	/* 216 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2602
		 调整:

		 中文名: 古代魔导士
		 卡片种类: 通常怪兽
		 卡片密码: 36821538
		 种族: 魔法师
		 属性: 暗
		 星级: 4
		 攻击力: 1000
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: Booster07
		 效果: 描述：拿着许多手杖。各自分别使用多样的攻击。

		*/
		Id:       2602,
		Password: "36821538",
		Name:     "古代魔导士",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Dark,        // 暗
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     1000,
		Def:     1300,
		IsValid: true,
	})

	/* 217 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2603
		 调整:

		 中文名: 恶魔龙
		 卡片种类: 通常怪兽
		 卡片密码: 67724379
		 种族: 龙
		 属性: 暗
		 星级: 4
		 攻击力: 1500
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: EX-R(EX)、Booster07
		 效果: 描述：凶恶的龙。吐出邪恶的火焰使人变得邪恶。

		*/
		Id:       2603,
		Password: "67724379",
		Name:     "恶魔龙",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Dark,   // 暗
		Lr:      ygo.LR_Dragon, // 龙
		Atk:     1500,
		Def:     1200,
		IsValid: true,
	})

	/* 218 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2604
		 调整:

		 中文名: 碎块人
		 卡片种类: 通常怪兽
		 卡片密码: 34743446
		 种族: 机械
		 属性: 暗
		 星级: 4
		 攻击力: 850
		 防御力: 1800
		 罕见度: 平卡N
		 卡包: Booster07、TP20
		 效果: 描述：身体各部位都是武器、分裂袭击对方。

		*/
		Id:       2604,
		Password: "34743446",
		Name:     "碎块人",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Dark,    // 暗
		Lr:      ygo.LR_Machine, // 机械
		Atk:     850,
		Def:     1800,
		IsValid: true,
	})

	/* 219 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2605
		 调整:

		 中文名: 转职的魔镜
		 卡片种类: 通常怪兽
		 卡片密码: 55337339
		 种族: 恶魔
		 属性: 暗
		 星级: 3
		 攻击力: 800
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: Booster07
		 效果: 描述：恶魔镜。受到攻击也不会破裂、守备损伤。

		*/
		Id:       2605,
		Password: "55337339",
		Name:     "转职的魔镜",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     800,
		Def:     1300,
		IsValid: true,
	})

	/* 220 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2606
		 调整:

		 中文名: 魔天老
		 卡片种类: 通常怪兽
		 卡片密码: 42431843
		 种族: 恶魔
		 属性: 暗
		 星级: 3
		 攻击力: 1000
		 防御力: 700
		 罕见度: 平卡N
		 卡包: ME、Booster07
		 效果: 描述：被天界放逐的堕天使，在暗域中的战斗很出色。

		*/
		Id:       2606,
		Password: "42431843",
		Name:     "魔天老",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     1000,
		Def:     700,
		IsValid: true,
	})

	/* 221 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2607
		 调整:

		 中文名: 岩石龙
		 卡片种类: 通常怪兽
		 卡片密码: 68171737
		 种族: 岩石
		 属性: 地
		 星级: 7
		 攻击力: 2000
		 防御力: 2300
		 罕见度: 平卡N
		 卡包: Booster07
		 效果: 描述：完全由岩石组成的龙。岩石的攻击很强。

		*/
		Id:       2607,
		Password: "68171737",
		Name:     "岩石龙",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   7,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Rock,  // 岩石
		Atk:     2000,
		Def:     2300,
		IsValid: true,
	})

	/* 222 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2608
		 调整:

		 中文名: 千年石人
		 卡片种类: 通常怪兽
		 卡片密码: 47986555
		 种族: 岩石
		 属性: 地
		 星级: 6
		 攻击力: 2000
		 防御力: 2200
		 罕见度: 平卡N
		 卡包: Booster07
		 效果: 描述：千年间一直在守护着王家财宝的石像。

		*/
		Id:       2608,
		Password: "47986555",
		Name:     "千年石人",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   6,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Rock,  // 岩石
		Atk:     2000,
		Def:     2200,
		IsValid: true,
	})

	/* 223 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2609
		 调整:

		 中文名: 角龙
		 卡片种类: 通常怪兽
		 卡片密码: 75390004
		 种族: 恐龙
		 属性: 地
		 星级: 6
		 攻击力: 1800
		 防御力: 2000
		 罕见度: 平卡N
		 卡包: Booster07
		 效果: 描述：全身长满角的恐龙。突击攻击很强烈！

		*/
		Id:       2609,
		Password: "75390004",
		Name:     "角龙",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   6,
		La:      ygo.LA_Earth,    // 地
		Lr:      ygo.LR_Dinosaur, // 恐龙
		Atk:     1800,
		Def:     2000,
		IsValid: true,
	})

	/* 224 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2610
		 调整:

		 中文名: 美丽的魔物使
		 卡片种类: 通常怪兽
		 卡片密码: 29616941
		 种族: 战士
		 属性: 地
		 星级: 5
		 攻击力: 1750
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: Booster07
		 效果: 描述：平常难得一见的女性魔物使。鞭子一到手上性格就立即改变。

		*/
		Id:       2610,
		Password: "29616941",
		Name:     "美丽的魔物使",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     1750,
		Def:     1500,
		IsValid: true,
	})

	/* 225 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2611
		 调整:

		 中文名: 地狱的魔物使
		 卡片种类: 通常怪兽
		 卡片密码: 97612389
		 种族: 战士
		 属性: 地
		 星级: 5
		 攻击力: 1800
		 防御力: 1600
		 罕见度: 平卡N
		 卡包: Booster07
		 效果: 描述：能自由自在地操纵怪兽攻击的怪兽使。

		*/
		Id:       2611,
		Password: "97612389",
		Name:     "地狱的魔物使",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     1800,
		Def:     1600,
		IsValid: true,
	})

	/* 226 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2612
		 调整:

		 中文名: 蘑菇人
		 卡片种类: 通常怪兽
		 卡片密码: 14181608
		 种族: 植物
		 属性: 地
		 星级: 2
		 攻击力: 800
		 防御力: 600
		 罕见度: 平卡N
		 卡包: Booster07
		 效果: 描述：在潮湿的地方发挥力量！从蘑菇帽中吐丝攻击。

		*/
		Id:       2612,
		Password: "14181608",
		Name:     "蘑菇人",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Plant, // 植物
		Atk:     800,
		Def:     600,
		IsValid: true,
	})

	/* 227 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2613
		 调整:

		 中文名: 捕猎蛇
		 卡片种类: 通常怪兽
		 卡片密码: 02906250
		 种族: 爬虫类
		 属性: 水
		 星级: 4
		 攻击力: 1300
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: Booster07
		 效果: 描述：奸诈狡猾的蛇，注意其用粗长的身体束缚攻击！

		*/
		Id:       2613,
		Password: "02906250",
		Name:     "捕猎蛇",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Water,   // 水
		Lr:      ygo.LR_Reptile, // 爬虫类
		Atk:     1300,
		Def:     1200,
		IsValid: true,
	})

	/* 228 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2614
		 调整:

		 中文名: 蜥蜴骑士
		 卡片种类: 通常怪兽
		 卡片密码: 78402798
		 种族: 水
		 属性: 水
		 星级: 4
		 攻击力: 1150
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: Booster07
		 效果: 描述：守护着大海、被绿色的鳞片包住身体的蜥蜴战士。

		*/
		Id:       2614,
		Password: "78402798",
		Name:     "蜥蜴骑士",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     1150,
		Def:     1300,
		IsValid: true,
	})

	/* 229 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2615
		 调整:

		 中文名: 月之使者
		 卡片种类: 通常怪兽
		 卡片密码: 45909477
		 种族: 战士
		 属性: 光
		 星级: 4
		 攻击力: 1100
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: Booster07
		 效果: 描述：侍奉月之女神的战士。使用像新月一样的戈进行攻击！

		*/
		Id:       2615,
		Password: "45909477",
		Name:     "月之使者",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Light,   // 光
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     1100,
		Def:     1000,
		IsValid: true,
	})

	/* 230 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 420
		 调整:

		 中文名: 水母
		 卡片种类: 通常怪兽
		 卡片密码: 14851496
		 种族: 水
		 属性: 水
		 星级: 4
		 攻击力: 1200
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: ME、DL04、Booster07
		 效果: 描述：漂浮在海面上的海蜇，半透明的身体使人难以察觉。

		*/
		Id:       420,
		Password: "14851496",
		Name:     "水母",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     1200,
		Def:     1500,
		IsValid: true,
	})

	/* 231 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 422
		 调整:

		 中文名: 暗魔界的霸王
		 卡片种类: 通常怪兽
		 卡片密码: 69455834
		 种族: 恶魔
		 属性: 暗
		 星级: 5
		 攻击力: 2000
		 防御力: 1530
		 罕见度: 平卡N
		 卡包: ME、BE02、DL04、Booster07
		 效果: 描述：使用强大的暗之力量，破坏周围的一切。

		*/
		Id:       422,
		Password: "69455834",
		Name:     "暗魔界的霸王",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     2000,
		Def:     1530,
		IsValid: true,
	})

	/* 232 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 435
		 调整:

		 中文名: 强化石像怪
		 卡片种类: 通常怪兽
		 卡片密码: 24611934
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 1600
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: ME、BE02、EX-R(EX)、DL04、Booster07、KA、SK2
		 效果: 描述：取得暗黑之力强化而成的石像鬼，尖锐的爪子值得警惕。

		*/
		Id:       435,
		Password: "24611934",
		Name:     "强化石像怪",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     1600,
		Def:     1200,
		IsValid: true,
	})

	/* 233 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 479
		 调整:

		 中文名: 电气小子
		 卡片种类: 通常怪兽
		 卡片密码: 27324313
		 种族: 雷
		 属性: 光
		 星级: 3
		 攻击力: 1000
		 防御力: 500
		 罕见度: 平卡N
		 卡包: BE02、DL04、Booster01、Booster Chronicle、Booster R1、TU04
		 效果: 描述：雷属性攻击异常厉害，若轻视它则会遭到电击。

		*/
		Id:       479,
		Password: "27324313",
		Name:     "电气小子",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Light,   // 光
		Lr:      ygo.LR_Thunder, // 雷
		Atk:     1000,
		Def:     500,
		IsValid: true,
	})

	/* 234 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 480
		 调整:

		 中文名: 吸血跳蚤
		 卡片种类: 通常怪兽
		 卡片密码: 41762634
		 种族: 昆虫
		 属性: 地
		 星级: 4
		 攻击力: 1500
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: DL04、Booster02、Booster Chronicle、Booster R1
		 效果: 描述：攻击力高强的巨型吸血跳蚤，轻视它会非常危险。

		*/
		Id:       480,
		Password: "41762634",
		Name:     "吸血跳蚤",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth,  // 地
		Lr:      ygo.LR_Insect, // 昆虫
		Atk:     1500,
		Def:     1200,
		IsValid: true,
	})

	/* 235 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 485
		 调整:
		 中文名: 复活节岛的摩艾石像
		 卡片种类: 通常怪兽
		 卡片密码: 10262698
		 种族: 岩石
		 属性: 地
		 星级: 4
		 攻击力: 1100
		 防御力: 1400
		 罕见度: {rarely }
		 卡包: BE02，DL04，Booster02，Booster Chronicle，Booster R1
		 效果: 描述：存在于复活岛的石像，能够发射圆形的激光。

		*/
		Id:       485,
		Password: "10262698",
		Name:     "复活节岛的摩艾石像",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Rock,  // 岩石
		Atk:     1100,
		Def:     1400,
		IsValid: true,
	})

	/* 236 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 486
		 调整:

		 中文名: 友谊天使
		 卡片种类: 通常怪兽
		 卡片密码: 82085619
		 种族: 天使
		 属性: 光
		 星级: 4
		 攻击力: 1300
		 防御力: 1100
		 罕见度: 平卡N
		 卡包: BE02、DL04、Booster02、Booster Chronicle、Booster R1
		 效果: 描述：即便在决斗中吵架，显示出友情就能和好。

		*/
		Id:       486,
		Password: "82085619",
		Name:     "友谊天使",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Light, // 光
		Lr:      ygo.LR_Fairy, // 天使
		Atk:     1300,
		Def:     1100,
		IsValid: true,
	})

	/* 237 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 492
		 调整:

		 中文名: 老虎斧战士
		 卡片种类: 通常怪兽
		 卡片密码: 49791927
		 种族: 兽战士
		 属性: 地
		 星级: 4
		 攻击力: 1300
		 防御力: 1100
		 罕见度: 平卡N，金字UR
		 卡包: BE02、LE02、DL04、Booster03、Booster Chronicle、Booster R2、JY
		 效果: 描述：手持巨斧的兽战士。能够放出行动迅速的人偶，攻击强劲。

		*/
		Id:       492,
		Password: "49791927",
		Name:     "老虎斧战士",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth,        // 地
		Lr:      ygo.LR_BeastWarrior, // 兽战士
		Atk:     1300,
		Def:     1100,
		IsValid: true,
	})

	/* 238 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 493
		 调整:

		 中文名: 巨斧袭击者
		 卡片种类: 通常怪兽
		 卡片密码: 48305365
		 种族: 战士
		 属性: 地
		 星级: 4
		 攻击力: 1700
		 防御力: 1150
		 罕见度: 平卡N
		 卡包: BE02、DL04、Booster04、Booster Chronicle、Booster R2、YSD03、JY、SJ2
		 效果: 描述：持斧的战士。单手挥舞斧头的攻击相当强劲。

		*/
		Id:       493,
		Password: "48305365",
		Name:     "巨斧袭击者",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     1700,
		Def:     1150,
		IsValid: true,
	})

	/* 239 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 495
		 调整:

		 中文名: 机械猎手
		 卡片种类: 通常怪兽
		 卡片密码: 07359741
		 种族: 机械
		 属性: 暗
		 星级: 4
		 攻击力: 1850
		 防御力: 800
		 罕见度: 平卡N
		 卡包: BE02、DL04、SD10、Booster04、Booster Chronicle、Booster R2
		 效果: 描述：听从机械王的命令，在抓住目标前都会紧追不舍的猎捕。

		*/
		Id:       495,
		Password: "07359741",
		Name:     "机械猎手",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Dark,    // 暗
		Lr:      ygo.LR_Machine, // 机械
		Atk:     1850,
		Def:     800,
		IsValid: true,
	})

	/* 240 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 496
		 调整:

		 中文名: 海鳞蛇
		 卡片种类: 通常怪兽
		 卡片密码: 58831685
		 种族: 水
		 属性: 水
		 星级: 4
		 攻击力: 1800
		 防御力: 800
		 罕见度: 平卡N
		 卡包: DL04、Booster04、Booster Chronicle、Booster R2
		 效果: 描述：居住在海洋里的蛇型怪兽，它会去咬任何接近的物体。

		*/
		Id:       496,
		Password: "58831685",
		Name:     "海鳞蛇",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     1800,
		Def:     800,
		IsValid: true,
	})

	/* 241 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 498
		 调整:

		 中文名: 双生精灵
		 卡片种类: 通常怪兽
		 卡片密码: 69140098
		 种族: 魔法师
		 属性: 地
		 星级: 4
		 攻击力: 1900
		 防御力: 900
		 罕见度: 平卡N
		 卡包: BE02、DL04、SD06、Booster04、Booster Chronicle、Booster R2
		 效果: 描述：相互配合发动攻击的双胞胎妖精姐妹。

		*/
		Id:       498,
		Password: "69140098",
		Name:     "双生精灵",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth,       // 地
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     1900,
		Def:     900,
		IsValid: true,
	})

	/* 242 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 500
		 调整:

		 中文名: 橐蜚
		 卡片种类: 通常怪兽
		 卡片密码: 03170832
		 种族: 鸟兽
		 属性: 风
		 星级: 4
		 攻击力: 1450
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: DL04、Booster03、Booster Chronicle、Booster R2
		 效果: 描述：这鸟出现的时候，就是什么不幸之事发生前的征兆。

		*/
		Id:       500,
		Password: "03170832",
		Name:     "橐蜚",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Wind,        // 风
		Lr:      ygo.LR_WingedBeast, // 鸟兽
		Atk:     1450,
		Def:     1000,
		IsValid: true,
	})

	/* 243 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 502
		 调整:

		 中文名: 泉之妖精
		 卡片种类: 通常怪兽
		 卡片密码: 81563416
		 种族: 水
		 属性: 水
		 星级: 4
		 攻击力: 1600
		 防御力: 1100
		 罕见度: 平卡N
		 卡包: DL04、Booster04、Booster Chronicle、Booster R2
		 效果: 描述：守护泉水的精灵。对任何污浊泉水者杀无赦。

		*/
		Id:       502,
		Password: "81563416",
		Name:     "泉之妖精",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     1600,
		Def:     1100,
		IsValid: true,
	})

	/* 244 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 503
		 调整:

		 中文名: 月光少女
		 卡片种类: 通常怪兽
		 卡片密码: 79629370
		 种族: 魔法师
		 属性: 光
		 星级: 4
		 攻击力: 1500
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: DL04、Booster04、Booster Chronicle、Booster R2
		 效果: 描述：月亮加护的魔导士。神秘的魔法能使人看到幻觉。

		*/
		Id:       503,
		Password: "79629370",
		Name:     "月光少女",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Light,       // 光
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     1500,
		Def:     1300,
		IsValid: true,
	})

	/* 245 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 506
		 调整:

		 中文名: 双头恐龙王
		 卡片种类: 通常怪兽
		 卡片密码: 94119974
		 种族: 恐龙
		 属性: 地
		 星级: 4
		 攻击力: 1600
		 防御力: 1200
		 罕见度: 平卡N，金字UR，爆闪PR
		 卡包: BE02、DL04、Booster05、Booster Chronicle、Booster R3、TP19
		 效果: 描述：恐龙族中的强力怪兽，两只头同时攻击。

		*/
		Id:       506,
		Password: "94119974",
		Name:     "双头恐龙王",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth,    // 地
		Lr:      ygo.LR_Dinosaur, // 恐龙
		Atk:     1600,
		Def:     1200,
		IsValid: true,
	})

	/* 246 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 507
		 调整:

		 中文名: 海龙神
		 卡片种类: 通常怪兽
		 卡片密码: 76634149
		 种族: 海龙
		 属性: 水
		 星级: 5
		 攻击力: 1800
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: DL04、Booster06、Booster Chronicle、Booster R3
		 效果: 描述：被称为海洋之主的海龙。掀起海啸吞噬一切。

		*/
		Id:       507,
		Password: "76634149",
		Name:     "海龙神",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   5,
		La:      ygo.LA_Water,      // 水
		Lr:      ygo.LR_SeaSerpent, // 海龙
		Atk:     1800,
		Def:     1500,
		IsValid: true,
	})

	/* 247 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 508
		 调整:

		 中文名: 龙僵尸
		 卡片种类: 通常怪兽
		 卡片密码: 66672569
		 种族: 不死
		 属性: 暗
		 星级: 3
		 攻击力: 1600
		 防御力: 0
		 罕见度: 平卡N
		 卡包: EX-R(EX)、DL04、Booster05、Booster Chronicle、Booster R3
		 效果: 描述：被魔力唤醒的龙，吐出的气息能使物体腐烂。

		*/
		Id:       508,
		Password: "66672569",
		Name:     "龙僵尸",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Dark,   // 暗
		Lr:      ygo.LR_Zombie, // 不死
		Atk:     1600,
		Def:     0,
		IsValid: true,
	})

	/* 248 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 510
		 调整:

		 中文名: 神灯魔人
		 卡片种类: 通常怪兽
		 卡片密码: 99510761
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 1400
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: BE02、DL04、Booster05、Booster Chronicle、Booster R3
		 效果: 描述：从神灯里出现的魔灵，服从召唤者的意志。

		*/
		Id:       510,
		Password: "99510761",
		Name:     "神灯魔人",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     1400,
		Def:     1200,
		IsValid: true,
	})
}
