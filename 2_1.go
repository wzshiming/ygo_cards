package ygo_cards

import ygo "github.com/wzshiming/ygo_core"

func d2_1(cardBag *ygo.CardVersion) {

	/* 0 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1140
		 调整:

		 中文名: 幻想的仪式
		 卡片种类: 仪式魔法
		 卡片密码: 41426869
		 罕见度: 平卡N，金字UR
		 卡包: BE02、VB02、PE
		 效果: 「纳祭之魔」降临必要。场上和手卡，合计1星以上的卡做祭品。

		*/
		Id:       1140,
		Password: "41426869",
		Name:     "幻想的仪式",
		Lc:       ygo.LC_SpellRitual, // 仪式魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 1 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1269
		 调整:

		 中文名: 混沌-黑魔术的仪式
		 卡片种类: 仪式魔法
		 卡片密码: 76792184
		 罕见度: 平卡N，面闪SR，金字UR，斜碎SCR
		 卡包: PP03、TP10、YU
		 效果: 召唤「黑混沌之魔术师」必须。场上和手卡加起来总共8颗星以上的怪兽作祭品。

		*/
		Id:       1269,
		Password: "76792184",
		Name:     "混沌-黑魔术的仪式",
		Lc:       ygo.LC_SpellRitual, // 仪式魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 2 */
	cardBag.Register(&ygo.CardOriginal{
		/*
			 id: 1142
			 调整:

			 [噩梦之铁栏]<悪夢の鉄檻>
			 ●全部怪兽在2回合（数对方的回合）以内不能攻击。
			 ◇可以进入战斗阶段。
			 ◇「荷鲁斯之黑炎龙LV6/ホルスの黒炎竜
			ＬＶ６」等怪兽可以进入进行攻击宣言。
			 ●2回合以后这张卡破坏。
			 ◇「魔法干扰阵/マジック·ジャマー」等使发动无效并破坏的场合，「噩梦之铁栏/悪夢の鉄檻」被破坏不会留在场上。
			 ◇「光与暗之龙/光と闇の竜」使发动无效的场合，「噩梦之铁栏/悪夢の鉄檻」不会留在场上。
			 ◇「旋风/サイクロン」等使之破壊的场合，「噩梦之铁栏/悪夢の鉄檻」不会留在场上。
			 ◇「魔人
			暗黑护持师/魔人
			ダーク·バルター」使効果无效化的场合，「噩梦之铁栏/悪夢の鉄檻」会留在场上。
			 ◇「王宫的敕命/王宮の勅命」使效果无效化的场合，「噩梦之铁栏/悪夢の鉄檻」会留在场上。
			 ◇「暗之取引/闇の取引」使「噩梦之铁栏/悪夢の鉄檻」的效果变更的场合，「噩梦之铁栏/悪夢の鉄檻」不会留在场上。
			 ◇「连续魔法/連続魔法」使「噩梦之铁栏/悪夢の鉄檻」的效果发动的场合，「/連続魔法」不会留在场上。
			 ◇「命运英雄
			钻石人/Ｄ－ＨＥＲＯダイヤモンドガイ」的効果发动的场合，墓地的「噩梦之铁栏/悪夢の鉄檻」不会放到场上。
			 ◇「二重魔法/二重魔法」的效果发动的场合，墓地的「噩梦之铁栏/悪夢の鉄檻」会放到场上并留在场上。

			 中文名: 噩梦之铁栏
			 卡片种类: 通常魔法
			 卡片密码: 58775978
			 罕见度: 平卡N，金字UR
			 卡包: SD06、SD12、VB03、PE、SDM
			 效果: 全部怪兽在2回合（数对方的回合）以内不能攻击。2回合以后这张卡破坏。

		*/
		Id:       1142,
		Password: "58775978",
		Name:     "噩梦之铁栏",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellUnnormal(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				i := 0

				ca.RegisterGlobalListen(ygo.BP, func(pl0 *ygo.Player) {
					if tar != pl0 {
						return
					}
					tar.Mzone().ForEach(func(c *ygo.Card) bool {
						c.SetNotCanAttack()
						return true
					})
				})

				ca.RegisterGlobalListen(ygo.RoundEnd, func(pl0 *ygo.Player) {
					if tar != pl0 {
						return
					}
					i++
					if i >= 2 {
						ca.Dispatch(ygo.Depleted)
					}
				})
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 3 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1169
		 调整:

		 中文名: 催眠术
		 卡片种类: 通常魔法
		 卡片密码: 48642904
		 罕见度: 平卡N，金字UR
		 卡包: SD09、KA、SK2
		 效果: 这张卡使用后，下次的对方的回合，对方不能改变怪兽的表示形式。

		*/
		Id:       1169,
		Password: "48642904",
		Name:     "催眠术",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 4 */
	cardBag.Register(&ygo.CardOriginal{
		/*
			 id: 124
			 调整:

			 [抹杀之使徒]<抹殺の使徒>
			 ●破坏一只覆盖表示的怪兽，并从游戏中除外。
			 ◇取对象
			 ◇对象被破坏的时候是里侧，被除外后是表侧
			 ◇效果处理时对象变成表侧的场合，效果不发
			 ●破坏的怪兽是翻转效果的场合，双方互相确认对方卡组，和破坏的怪兽同名的卡全部从游戏中除外。之后洗切卡组。
			 ◇破坏的不是反转效果怪兽的场合，无需确认
			 ◇即使公开情报已经排除了一方卡组里还有同名卡的可能，也要确认
			 ◇满足条件的场合，必须确认，不能因任何理由不给对方看卡组或者不看对方的卡组
			 ◇对象是
			 [寄生虫
			帕拉赛德/寄生虫パラサイド]的场合，因
			 [寄生虫
			帕拉赛德/寄生虫パラサイド]的效果洗入对手卡组的
			 [寄生虫
			帕拉赛德/寄生虫パラサイド]也要被除外

			 中文名: 抹杀之使徒
			 卡片种类: 通常魔法
			 卡片密码: 71044499
			 罕见度: 平卡N，银字R，黄金GR，面闪SR
			 卡包: BE01、CA、DL01、SD01、SD02、SD03、SD06、SD08、SD11、SD14、SD12、DT05、DTP01、TU03、YU、JY、KA、SK2、GS03、GDB1
			 效果: 选择场上里侧表示存在的1只怪兽破坏，从游戏中除外。那是反转效果怪兽的场合，把双方卡组确认，同名卡全部从游戏中除外。

		*/
		Id:       124,
		Password: "71044499",
		Name:     "抹杀之使徒",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 5 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 125
		 调整:

		 [扑灭之使徒]<撲滅の使徒>
		 ●破坏一张覆盖表示的魔法或陷阱卡，并从游戏中除外。
		 ◇取对象
		 ◇对象被破坏的时候是里侧，被除外后是表侧
		 ◇效果处理时对象变成表侧的场合，效果不发
		 ●破坏的卡是陷阱卡的场合，双方互相确认对方卡组，和破坏的陷阱卡同名的卡全部从游戏中除外。之后洗切卡组。
		 ◇破坏的是魔法卡的场合，无需确认
		 ◇即使公开情报已经排除了一方卡组里还有同名卡的可能，也要确认
		 ◇满足条件的场合，必须确认，不能因任何理由不给对方看卡组或者不看对方的卡组

		 中文名: 扑灭之使徒
		 卡片种类: 通常魔法
		 卡片密码: 17449108
		 罕见度: 平卡N，银字R
		 卡包: BE01、CA、DL01、DT01、KA
		 效果: 选择场上盖放的1张魔法·陷阱卡破坏，从游戏中除外。那是陷阱卡的场合，把双方卡组确认，同名卡全部从游戏中除外。

		*/
		Id:       125,
		Password: "17449108",
		Name:     "扑灭之使徒",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellNormal(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				if c := pl.SelectForWarnShort(ygo.LO_Removed, 0, pl.Szone(), tar.Szone(), func(c0 *ygo.Card) bool {
					return c0.IsFaceDown()
				}); c != nil {
					c.Removed(ca)
					if c.IsTrap() {
						ygo.NewCards(pl.Deck(), tar.Deck(), func(c0 *ygo.Card) bool {
							return c0.GetName() == c.GetName()
						}).ForEach(func(c0 *ygo.Card) bool {
							c0.Removed(ca)
							return true
						})
					}
				}
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 6 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 126
		 调整:

		 [过浅的墓穴]<浅すぎた墓穴>
		 [10/12/09]
		 ●自己和对方在各自的墓地选择一只怪兽，以防守表示的形式在场上放置。
		 ◇取对象
		 ◇一方墓地里没有满足条件的怪兽的场合，此卡不能发动
		 ◇一方有5只怪兽的场合，此卡不能发动
		 ◇效果处理时有一方场上没有空怪兽区的场合，那方不特殊召唤，另一方特殊召唤
		 ◇一方没有空怪兽区但有不能使用的怪兽区的场合，可以发动此卡，那方不特殊召唤
		 ◇“以防守表示的形式在场上放置”等价于“里侧守备表示特殊召唤”
		 ◇对象有以“从墓地特殊召唤成功”为诱发条件的诱发效果的场合，那效果不能发动
		 ◇特殊召唤成功时点不能发动
		 [奈落的落穴/奈落の落とし穴]之类需要判定攻击力的卡
		 ◇满足
		 [生还的宝札/生还の宝札]的效果的发动条件
		 ◇
		 [神圣光辉/圣なる辉き]效果适用中,可以发动怪兽表侧守备特殊召唤。

		 中文名: 过浅的墓穴
		 卡片种类: 通常魔法
		 卡片密码: 43434803
		 罕见度: 平卡N
		 卡包: BE01、CA、DL01、SD13、DT06、KA、DS13
		 效果: 双方玩家选择各自墓地存在的1只怪兽，在各自的场上里侧守备表示盖放。

		*/
		Id:       126,
		Password: "43434803",
		Name:     "过浅的墓穴",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 7 */
	cardBag.Register(&ygo.CardOriginal{
		/*
			 id: 1272
			 调整:

			 [死之魔术箱]<死のマジック·ボックス>
			 [09/09/20]
			 ●场上的对方的怪兽一只破坏，自己场上的一只怪兽的控制权转移给对方。
			 ◇取对象效果。
			 ◇发动这张卡时在双方场上各选择一只怪兽作为效果的对象。
			 ◇自己场上没有能够转移控制权的怪兽存在的场合，不能发动。
			 ◇自己场上只有
			 [荷鲁斯之黑炎龙LV6/ホルスの黒炎竜
			ＬＶ６]存在的场合可以发动。这个场合选择的怪兽被破坏。
			 [荷鲁斯之黑炎龙LV6/ホルスの黒炎竜
			ＬＶ６]的控制权不会被转移。
			 ◇效果处理时任意一个对象不在场上表侧表示存在的场合，效果不发。

			 中文名: 死之魔术箱
			 卡片种类: 通常魔法
			 卡片密码: 25774450
			 罕见度: 平卡N，金字UR
			 卡包: PP04、SD06、DT04、YU、SY2
			 效果: 场上的对方的怪兽1只破坏，自己场上的1只怪兽的控制权转移给对方。

		*/
		Id:       1272,
		Password: "25774450",
		Name:     "死之魔术箱",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 8 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1294
		 调整:

		 中文名: 呼龙笛
		 卡片种类: 通常魔法
		 卡片密码: 43973174
		 罕见度: 平卡N，金字UR
		 卡包: EX-R(EX)、TP07、DPKB、KA、SK2、SD22
		 效果: 场上有「龙之支配者」存在的场合，从手卡把最多2只龙族怪兽特殊召唤。

		*/
		Id:       1294,
		Password: "43973174",
		Name:     "呼龙笛",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 9 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1295
		 调整:

		 [灵魂交错]<クロス·ソウル>
		 [10/07/14]
		 ◇只能在主阶段1发动
		 ◇可以在同1个主阶段1发动复数
		 ◇可在先攻第一回合发动
		 ◇取对象,不能进入战斗阶段是契约
		 ◇可以连锁发动
		 [天使的手镜/天使の手镜]
		 ◇不能无意义的单纯解放那只怪兽
		 ◇除可用于上级召唤外,还可用于卡的效果COST而解放和仪式魔法发动所需要的解放
		 ◇不受魔法卡效果影响的怪兽也不能进行攻击宣言。
		 ◇不能把效果持有『不受魔法卡效果影响』的对方怪兽解放
		 ◇不能选择对方场上表侧表示存在的「召唤僧/召唤僧サモンプリースト」为对象
		 ◇「牲祭封印之假面/生贽封じの仮面」效果适用时，这张卡可以发动

		 中文名: 灵魂交错
		 卡片种类: 通常魔法
		 卡片密码: 68005187
		 罕见度: 平卡N，银碎SER
		 卡包: EX-R(EX)、SD14、DT01、DTP01、KA、PE、SK2、GLD04、SD25、DS13
		 效果: 选择对方场上的1只怪兽发动。这个回合自己的怪兽被解放的场合，作为代替1只自己的怪兽必须把被选择的对方怪兽解放。这张卡发动的回合，自己不能进行战斗阶段。

		*/
		Id:       1295,
		Password: "68005187",
		Name:     "灵魂交错",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 10 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1296
		 调整:

		 中文名: 手札抹杀
		 卡片种类: 通常魔法
		 卡片密码: 72892473
		 罕见度: 平卡N，银碎SER
		 卡包: EX-R(EX)、SD08、SD15、SD19、YU、SD21
		 效果: 双方手卡全部丢弃，各自从自己卡组抽出丢弃数量的卡。

		*/
		Id:       1296,
		Password: "72892473",
		Name:     "手札抹杀",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellNormal(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				h1 := pl.Hand()
				s1 := h1.Len()
				for h1.Len() != 0 {
					h1.EndPop().Discard(ca)
				}
				h2 := tar.Hand()
				s2 := h2.Len()
				for h2.Len() != 0 {
					h2.EndPop().Discard(ca)
				}
				pl.DrawCard(s1)
				tar.DrawCard(s2)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 11 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 172
		 调整:

		 中文名: 雷击
		 卡片种类: 通常魔法
		 卡片密码: 12580477
		 罕见度: 面闪SR，金字UR
		 卡包: LB、BE01、DL02、Starter Box、PE
		 效果: 对方场上存在的怪兽全部破坏。

		*/
		Id:       172,
		Password: "12580477",
		Name:     "雷击",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellNormal(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()

				tar.Mzone().ForEach(func(c0 *ygo.Card) bool {
					c0.Destroy(ca)
					return true
				})
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 12 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 175
		 调整:

		 中文名: 火球
		 卡片种类: 通常魔法
		 卡片密码: 46130346
		 罕见度: 平卡N
		 卡包: LB、BE01、DL02、Starter Box
		 效果: 给予对方基本分500分伤害。

		*/
		Id:       175,
		Password: "46130346",
		Name:     "火球",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellNormal(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				tar.ChangeLp(-500)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 13 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 179
		 调整:

		 中文名: 陷阱拆除
		 卡片种类: 通常魔法
		 卡片密码: 51482758
		 罕见度: 平卡N
		 卡包: LB、EX-R(EX)、DL02、Starter Box、SY2
		 效果: 选择场上表侧表示存在的1张陷阱卡破坏。

		*/
		Id:       179,
		Password: "51482758",
		Name:     "陷阱拆除",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellNormal(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				if c := pl.SelectForWarnShort(ygo.LO_Destroy, 0, tar.Szone(), func(c0 *ygo.Card) bool {
					return c0.IsTrap() && c0.IsFaceUp()
				}); c != nil {
					c.Destroy(ca)
				}
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 14 */
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

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellNormal(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				if c := pl.SelectForWarnShort(ygo.LO_Destroy, 0, tar.Hand()); c != nil {
					c.PeekFor(pl)
				}
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 15 */
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

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellNormal(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				cs := ygo.NewCards()
				for i := 0; i != 5; i++ {
					if c := tar.Deck().EndPop(); c != nil {
						cs.EndPush(c)
					}
				}

				tar.SelectForPopup(ygo.LO_PutBack, cs)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 16 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 288
		 调整:

		 [虫洞]<ワーム·ホール>
		 ●选自己场上的一只怪兽，在自己的下次的准备流程之前除外。
		 ◇取对象
		 ◇可选择里侧表示的怪兽，此时对象以里侧表示除外
		 ◇下个准备阶段被跳过的场合，对象不能因此卡的效果回来★取里侧的
		 [巨盾守卫者/ビッグ·シールド·ガードナー]为对象的场合，是否会被无效，调整中注：此时对手并不知道里侧的是
		 [巨盾守卫者/ビッグ·シールド·ガードナー]，所以不处理看上去也不会有破绽
		 ●除外的时候，被选择的怪兽的怪兽区的位置不能使用。
		 ◇对象不被除外的场合，此效果不适用

		 中文名: 虫洞
		 卡片种类: 通常魔法
		 卡片密码: 22959079
		 罕见度: 平卡N
		 卡包: BE01、TB 、DL03、PE
		 效果: 选自己场上的1只怪兽，在自己的下次的准备阶段之前除外。除外的时候，被选择的怪兽的怪兽区的位置不能使用。

		*/
		Id:       288,
		Password: "22959079",
		Name:     "虫洞",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 17 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 38
		 调整:

		 [飓风]<ハリケーン>
		 [2010/10/01]
		 ●场上存在的魔法·陷阱卡全部回到持有者的手卡。
		 ◇这张卡自身不能回到手卡
		 ◇这张卡发动的那组连锁中，确定使用后要送去墓地的魔法、陷阱卡不能回到手卡

		 中文名: 飓风
		 卡片种类: 通常魔法
		 卡片密码: 42703248
		 罕见度: 平卡N，银字R，黄金GR
		 卡包: BE01、MR、YSD02、DL01、SD02、SD05、SD12、SD16、YSD04、GLD02、GS02、YU、JY、PE、SJ2、GDB1
		 效果: 场上存在的魔法·陷阱卡全部回到持有者手卡。

		*/
		Id:       38,
		Password: "42703248",
		Name:     "飓风",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellNormal(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()

				ygo.NewCards(pl.Szone(), tar.Szone(), func(c0 *ygo.Card) bool {
					return c0 != ca
				}).ForEach(func(c0 *ygo.Card) bool {
					c0.ToHand()
					return true
				})
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 18 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 541
		 调整:

		 中文名: 卡通目录
		 卡片种类: 通常魔法
		 卡片密码: 89997728
		 罕见度: 平卡N，黄金GR，平罕NR
		 卡包: BE02、DL04、PE、GLD04
		 效果: 从卡组选1张名字有「卡通」字眼的卡加入手卡。

		*/
		Id:       541,
		Password: "89997728",
		Name:     "卡通目录",
		Lc:       ygo.LC_SpellNormal, // 通常魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 19 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1147
		 调整:

		 中文名: 增殖
		 卡片种类: 速攻魔法
		 卡片密码: 40703222
		 罕见度: 平卡N，面闪SR，金字UR
		 卡包: CRMS(603)、YU、TP15、其他
		 效果: 表侧表示的1只「栗子球」作祭品。自己的场上空出来的怪兽区域全部放置守备表示的「栗子球衍生物」（恶魔族·暗·1星·攻300/守200）。不能作祭品召唤的祭品。

		*/
		Id:       1147,
		Password: "40703222",
		Name:     "增殖",
		Lc:       ygo.LC_SpellQuickPlay, // 速攻魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 20 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1184
		 调整:

		 中文名: 天使的骰子
		 卡片种类: 速攻魔法
		 卡片密码: 74137509
		 罕见度: 平卡N，金字UR，银碎SER
		 卡包: TP05、JY、SJ2、GLD04、其他
		 效果: 掷1个骰子。自己控制的全部怪兽的攻击力守备力在回合结束前上升投掷出的数目×100。

		*/
		Id:       1184,
		Password: "74137509",
		Name:     "天使的骰子",
		Lc:       ygo.LC_SpellQuickPlay, // 速攻魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellQuickPlay(func() {
				pl := ca.GetSummoner()
				oi := pl.Coins(6) + 1
				pl.Mzone().ForEach(func(c *ygo.Card) bool {
					if c.IsFaceUp() {
						c.SetAtk(c.GetAtk() + oi*100)
						c.SetDef(c.GetDef() + oi*100)

						pl.GetCurrent().OnlyOnce(ygo.RoundEnd, func() {
							c.SetAtk(c.GetAtk() - oi*100)
							c.SetDef(c.GetDef() - oi*100)
						}, ca, c)
					}
					return true
				})
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 21 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1226
		 调整:

		 [替罪羊]<スケープ·ゴート>
		 ●这张卡发动的场合，这个回合内不能召唤，反转召唤，特殊召唤。
		 ◇誓约效果。
		 ◇发动「扰乱衍生物/おジャマトークン」等卡后，这个回合不能再发动这张卡。
		 ◇在对方场上将「熔岩魔神/溶岩魔神ラヴァ·ゴーレム」等特殊召唤的场合，这个回合不能再发动这张卡。
		 ◇可以将怪兽放置（SET）。
		 ●放置4只「羊衍生物/羊トークン」（地/1星/0/0）到自己的场上。（不能用来做召唤的祭品）
		 ◇自己场上怪兽区域不满4个位置的场合不能发动。
		 ◇效果处理时自己场上怪兽区域不满4个位置的场合效果不发（但誓约效果仍然适用）。

		 中文名: 替罪羊
		 卡片种类: 速攻魔法
		 卡片密码: 73915051
		 罕见度: 平卡N，黄金GR，金字UR
		 卡包: LE03、YSD02、GS01、JY、SJ2、GDB1
		 效果: 这张卡发动的回合，自己不能召唤·反转召唤·特殊召唤。在自己场上把4只「羊衍生物」（兽族·地·1星·攻/守0）守备表示特殊召唤。这衍生物不能为上级召唤而解放。

		*/
		Id:       1226,
		Password: "73915051",
		Name:     "替罪羊",
		Lc:       ygo.LC_SpellQuickPlay, // 速攻魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 22 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 37
		 调整:

		 [旋风]<サイクロン>
		 [2010/10/01]
		 ●场上的1张魔法或陷阱卡破坏。
		 ◇发动时选择场上的1张魔法或陷阱卡（取对象）
		 ◇不能选择这张卡自身

		 中文名: 旋风
		 卡片种类: 速攻魔法
		 卡片密码: 05318639
		 罕见度: 平卡N，银字R，黄金GR，面闪SR
		 卡包: DP04、BE01、MR、DL01、SD01、SD02、SD03、SD04、SD05、SD06、SD07、SD08、SD09、SD10、SD11、SD12、DT02、GS01、GLD03、YU、SY2、SJ2、SDM、YSD06、DB12、ST12、SD26、DS13、GDB1
		 效果: 选择场上存在的1张魔法·陷阱卡破坏。

		*/
		Id:       37,
		Password: "05318639",
		Name:     "旋风",
		Lc:       ygo.LC_SpellQuickPlay, // 速攻魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellQuickPlay(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				if c := pl.SelectForWarn(ygo.LO_Destroy, pl.Szone(), tar.Szone(), func(c0 *ygo.Card) bool {
					return c0 != ca
				}); c != nil {
					c.Destroy(ca)
				}
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 23 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 401
		 调整:

		 [融合解除]<融合解除>
		 [11/05/15]
		 ●场上表侧表示存在的１只融合怪兽回到额外卡组。
		 ◇发动时选择场上表侧表示存在的1只融合怪兽（取对象）
		 ◇效果处理时选择的怪兽不在场上表侧表示的场合，这张卡效果不适用
		 ●这个时候，回到融合卡组的这只怪兽的融合召唤使用过的１组怪兽在自己墓地齐集，可以再把这１组在自己场上特殊召唤。★暂调整为不取墓地里的怪兽为对象★
		 [终结之阿努比斯/エンド·オブ·アヌビス]的效果暂调整为不会影响这张卡★
		 [王家长眠之谷/王家の眠る谷－ネクロバレ－]的效果是否会影响这张卡，调整中
		 ◇融合怪兽回去后，选择是否特殊召唤素材，选“是”的场合，必须召唤1组，不能只召唤其中1只
		 ◇融合素材中有1只不能特殊召唤的场合，所有素材都不能特殊召唤
		 ◇效果处理时怪兽区的空位比融合素材少的场合，不能进行特殊召唤
		 ◇代替的素材也可以因此效果特殊召唤
		 ◇素材在墓地里名字不对的场合，不能进行特殊召唤例：用
		 [原型电子龙/プロト·サイバー·ドラゴン]当
		 [电子龙/サイバー·ドラゴン]进行融合的场合，不能用此卡特殊召唤
		 [原型电子龙/プロト·サイバー·ドラゴン]，
		 ◇选择对方场上的融合怪兽的场合，除了这只融合怪兽是此卡发动者融合然后将控制权转移给对手的场合以外，不能进行特殊召唤
		 ◇素材之一进了对手墓地的场合，不能进行特殊召唤
		 ◇就算因此效果特殊召唤复数怪兽，一张
		 [生还的宝札/生还の宝札]也只能抽1张卡
		 ◇用
		 [暴君龙/タイラント·ドラゴン]作
		 [五神龙/F·G·D]的素材的场合，不能进行特殊召唤
		 ◇用原本不是龙族的怪兽作
		 [五神龙/F·G·D]的素材的场合，不能进行特殊召唤
		 ◇非正规出场的融合怪兽，因此卡效果返回融合卡组的场合，就算那样的素材在墓地里存在，也不能特殊召唤
		 ◇对象是
		 [VWXYZ-神龙强击炮/VWXYZ－ドラゴン·カタパルトキャノン]
		 [VW-强击虎/VW－タイガー·カタパル]
		 [XYZ-神龙炮/XYZ－ドラゴン·キャノン]
		 [XY-钢龙炮/XY－ドラゴン·キャノン]
		 [XZ-机甲炮/XZ－キャタピラー·キャノン]
		 [YZ-机甲龙/YZ－キャタピラー·ドラゴン]且所用的“素材”在墓地的场合，也不能把那一组怪兽特殊召唤
		 ◇融合呪印生物因自身效果进行特殊召唤的融合怪兽不当做融合召唤，不能把那一组怪兽特殊召唤
		 ◇融合怪兽被「亚空间物质传送装置/亜空间物质転送装置」等效果暂时除外又回到场上之后，这个效果可以把那一组怪兽特殊召唤
		 ◇融合召唤过后送去墓地或从游戏中除外再用卡的效果特殊召唤的融合怪兽被这张卡回到额外卡组的场合不能把之前的融合召唤时使用过的那一组怪兽特殊召唤
		 ◇选择的怪兽的融合召唤所使用的怪兽曾经离开过墓地的场合，这个效果不能将这些怪兽进行特殊召唤

		 中文名: 融合解除
		 卡片种类: 速攻魔法
		 卡片密码: 95286165
		 罕见度: 平卡N
		 卡包: DP04、BE02、LN、DL03、DT08、PE、SY2
		 效果: 场上的1只融合怪兽回到融合卡组。这个时候，回到融合卡组的这只怪兽的融合召唤时使用的融合素材怪兽1组从自己的墓地拿起，这样的1组特殊召唤到自己的场上。

		*/
		Id:       401,
		Password: "95286165",
		Name:     "融合解除",
		Lc:       ygo.LC_SpellQuickPlay, // 速攻魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 24 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1151
		 调整:

		 中文名: 光辉城堡
		 卡片种类: 装备魔法
		 卡片密码: 82878489
		 罕见度: 银字R，金字UR
		 卡包: SOVR(606)、PE、其他
		 效果: 光属性的怪兽攻击力上升700。

		*/
		Id:       1151,
		Password: "82878489",
		Name:     "光辉城堡",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellEquip(func(c *ygo.Card) bool {
				return c.AttrIsLight()
			}, func(c *ygo.Card) {
				c.SetAtk(c.GetAtk() + 700)

			}, func(c *ygo.Card) {
				c.SetAtk(c.GetAtk() - 700)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 25 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 127
		 调整:

		 [过早的埋葬]<早すぎた埋葬>
		 ●付出800分。选择自己的墓地的一只怪兽攻击表示特殊召唤到场上，并装备上这张卡。这张卡破坏时，装备的怪兽破坏。
		 ◇取对象
		 ◇此卡因破坏以外的方式离场的场合，装备怪兽不破坏
		 ◇“装备上这张卡”这个动作不占时点

		 中文名: 过早的埋葬
		 卡片种类: 装备魔法
		 卡片密码: 70828912
		 罕见度: 平卡N，面闪SR
		 卡包: BE01、CA、YSD01、YSD02、DL01、SD01、SD03、SD04、SD06、SD07、SD11、SD13、YU、KA、SJ2
		 效果: 支付800基本分，选择自己墓地存在的1只怪兽发动。选择的怪兽表侧攻击表示特殊召唤，并装备这张卡。这张卡破坏时，装备怪兽破坏。

		*/
		Id:       127,
		Password: "70828912",
		Name:     "过早的埋葬",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 26 */
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

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellEquip(func(c *ygo.Card) bool {
				return c.AttrIsDark()
			}, func(c *ygo.Card) {
				c.SetAtk(c.GetAtk() + 400)
				c.SetDef(c.GetDef() - 200)

			}, func(c *ygo.Card) {
				c.SetAtk(c.GetAtk() - 400)
				c.SetDef(c.GetDef() + 200)

			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 27 */
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

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellEquip(func(c *ygo.Card) bool {
				return c.AttrIsEarth()
			}, func(c *ygo.Card) {
				c.SetAtk(c.GetAtk() + 400)
				c.SetDef(c.GetDef() - 200)

			}, func(c *ygo.Card) {
				c.SetAtk(c.GetAtk() - 400)
				c.SetDef(c.GetDef() + 200)

			})
			return true
		}, // 初始
		IsValid: true,
	})

	/* 28 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 26
		 调整:

		 [强夺]<強奪>
		 [2010/10/01]
		 ●得到装备了这张卡的对方怪兽的控制权。每次对方的准备阶段，对方恢复1000基本分。
		 ◇发动时选择对方场上表侧表示存在的1只怪兽（取对象）
		 ◇装备怪兽不在自己场上存在，回复效果也发动
		 ◇回复效果是进入连锁的魔法效果（咒文速度1）
		 ◇回复效果强制发动

		 中文名: 强夺
		 卡片种类: 装备魔法
		 卡片密码: 45986603
		 罕见度: 平卡N，面闪SR，金字UR
		 卡包: BE01、MR、DL01、SD01、SD02、SD03、SD04、SD05、KA、PE
		 效果: 得到这张卡装备的1只对方怪兽的控制权。在每次对方的准备阶段对方回复1000基本分。

		*/
		Id:       26,
		Password: "45986603",
		Name:     "强夺",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 29 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 282
		 调整:

		 [流星直击]<メテオ·ストライク>
		 ●攻击防守表示的怪兽时，装备这张卡的怪兽超过被攻击的防守怪兽的防守力，对方的基本分受到这个差值的伤害。
		 ◇“对方”指此卡控制者的对手，给对方场上怪兽装备此卡，装备怪兽攻击防守表示的怪兽时，装备这张卡的怪兽超过被攻击的防守怪兽的防守力的场合，受伤害的也是对手
		 ◇此效果不能重复

		 中文名: 流星直击
		 卡片种类: 装备魔法
		 卡片密码: 97687912
		 罕见度: 平卡N，银字R
		 卡包: BE01、TB 、DL03、DT02、KA
		 效果: 攻击守备表示的怪兽时，装备这张卡的怪兽超过被攻击的守备怪兽的守备力，对方的基本分受到这个差值的战斗伤害。

		*/
		Id:       282,
		Password: "97687912",
		Name:     "流星直击",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 30 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 3
		 调整:

		 [恶魔之斧]<デーモンの斧>
		 [2010/10/01]
		 ●装备怪兽的攻击力上升1000分。这张卡从场上送去墓地时，可以把自己场上存在的1只怪兽解放回到卡组最上方。
		 ◇发动时选择场上1只表侧表示怪兽（取对象）
		 ◇自身回收效果是进入连锁的魔法效果（咒文速度1）
		 ◇自身回收效果任意发动
		 ◇发动的场合把自己场上存在的1只怪兽解放（代价）
		 ◇覆盖在场上的这张卡被送去墓地的场合也可以发动
		 ◇「魔法干扰阵/マジック·ジャマー」「神之宣告/神の宣告」等把这张卡发动无效并破坏的场合，自身回收效果不能发动
		 ◇「魔术礼帽/マジカルシルクハット」的效果把这张卡作为怪兽时从场上被送去墓地的场合自身回收也可以效果发动
		 ◇自身回收效果可以在伤害步骤中发动

		 中文名: 恶魔之斧
		 卡片种类: 装备魔法
		 卡片密码: 40619825
		 罕见度: 平卡N，银字R，面闪SR
		 卡包: BE01、MR、DL01、DT04、JY
		 效果: 装备怪兽的攻击力上升1000。这张卡从场上送去墓地时，可以把自己场上存在的1只怪兽解放让这张卡回到卡组最上面。

		*/
		Id:       3,
		Password: "40619825",
		Name:     "恶魔之斧",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 31 */
	cardBag.Register(&ygo.CardOriginal{
		/*
			 id: 403
			 调整:

			 [疫病病毒
			黑尘]<疫病ウィルス
			ブラックダスト>
			 ●装备这张卡的怪兽不能攻击。★是否可以通过
			 [善变的裁缝师/移り気な仕立屋]的效果阻止攻击，调整中
			 ●装备怪兽的控制者的第2次回合结束时，装备怪兽破坏。这个效果成功的场合，这张卡回到主人的手卡。★是否入连锁，调整中
			 ◇此卡效果被无效的场合，也数回合
			 ◇此卡效果被无效后，装备怪兽的控制者的第2次回合结束时，此卡残留在场上，此后此卡效果再次有效的场合，只有“装备这张卡的怪兽不能攻击”这个效果适用
			 ◇通过
			 [善变的裁缝师/移り気な仕立屋]变更装备对象的场合，回合重数
			 ◇装备怪兽控制权转移的场合，回合重数★此卡装备给不受魔法效果的怪兽的场合，是否数回合，调整中★因装备怪兽同时装备
			 [明镜止水之心/明鏡止水の心]而没破到怪兽的场合，怎么处理，调整中

			 中文名: 疫病病毒 黑尘
			 卡片种类: 装备魔法
			 卡片密码: 69954399
			 罕见度: 平卡N
			 卡包: BE02、LN、DL03、KA
			 效果: 装备这张卡的怪兽不能攻击。装备怪兽的控制者的第2次回合结束时，装备怪兽破坏。这个效果成功的场合，这张卡回到持有者的手卡。

		*/
		Id:       403,
		Password: "69954399",
		Name:     "疫病病毒 黑尘",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 32 */
	cardBag.Register(&ygo.CardOriginal{
		/*
			 id: 46
			 调整:

			 [巨大化]<巨大化>
			 ●自己的基本分比对方低的场合，装备的怪兽的攻击力变成原本的攻击力的2倍。自己的基本分比对方的基本分高的场合，装备的怪兽的攻击力变成原本攻击力的一半。
			 ◇不可叠加
			 ◇双方LP相等的场合，此卡没有效果
			 ◇变的不是原本攻击力
			 ◇“自己”指此卡控制者，不是装备怪兽控制者
			 ◇
			 [天邪鬼的诅咒/あまのじゃくの呪い]効果不适用
			 ◇
			 [合成魔兽
			聚变鬼/合成魔獣
			ガーゼット]不受此卡效果影响
			 ◇以TOKEN为对象的场合，以召TOKEN的卡上记述的原本攻击力为准，但此卡对
			 [物理分身/物理分身]
			 [克隆复制/クローン複製]产生的TOKEN无效
			 ◇以攻击力增减效果适用的
			 [勇气之沙漏/勇気の砂時計]
			 [任性的女神/きまぐれの女神]为对象的场合，此卡效果适用，那些卡的效果无效

			 中文名: 巨大化
			 卡片种类: 装备魔法
			 卡片密码: 22046459
			 罕见度: 平卡N，面闪SR
			 卡包: BE01、PS、DL01、SD09、SD12、DT07、KA、SK2、SD26
			 效果: 自己基本分比对方低的场合，装备怪兽的攻击力变成原本攻击力2倍的数值。自己基本分比对方高的场合，装备怪兽的攻击力变成原本攻击力一半的数值。

		*/
		Id:       46,
		Password: "22046459",
		Name:     "巨大化",
		Lc:       ygo.LC_SpellEquip, // 装备魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 33 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 165
		 调整:

		 中文名: 森
		 卡片种类: 场地魔法
		 卡片密码: 87430998
		 罕见度: 平卡N
		 卡包: LB、DL02、Starter Box
		 效果: 全部昆虫族·植物族·兽族·兽战士族怪兽的攻击力和守备力上升200。

		*/
		Id:       165,
		Password: "87430998",
		Name:     "森",
		Lc:       ygo.LC_SpellField, // 场地魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellField(ca.EffectMzoneHalo(func(c *ygo.Card) bool {
				return c.RaceIsInsect() || c.RaceIsPlant() || c.RaceIsBeast() || c.RaceIsBeastWarrior()
			}, func(c *ygo.Card) {
				c.SetAtk(c.GetAtk() + 200)
				c.SetDef(c.GetDef() + 200)
			}, func(c *ygo.Card) {
				c.SetAtk(c.GetAtk() - 200)
				c.SetDef(c.GetDef() - 200)
			}))
			return true
		}, // 初始
		IsValid: true,
	})

	/* 34 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 166
		 调整:

		 中文名: 荒野
		 卡片种类: 场地魔法
		 卡片密码: 23424603
		 罕见度: 平卡N
		 卡包: LB、DL02、Starter Box
		 效果: 全部恐龙族·不死族·岩石族怪兽的攻击力和守备力上升200。

		*/
		Id:       166,
		Password: "23424603",
		Name:     "荒野",
		Lc:       ygo.LC_SpellField, // 场地魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellField(ca.EffectMzoneHalo(func(c *ygo.Card) bool {
				return c.RaceIsDinosaur() || c.RaceIsZombie() || c.RaceIsRock()
			}, func(c *ygo.Card) {
				c.SetAtk(c.GetAtk() + 200)
				c.SetDef(c.GetDef() + 200)
			}, func(c *ygo.Card) {
				c.SetAtk(c.GetAtk() - 200)
				c.SetDef(c.GetDef() - 200)
			}))
			return true
		}, // 初始
		IsValid: true,
	})

	/* 35 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 167
		 调整:

		 中文名: 山
		 卡片种类: 场地魔法
		 卡片密码: 50913601
		 罕见度: 平卡N
		 卡包: LB、DL02、Starter Box、TP12
		 效果: 场上表侧表示存在的龙族·鸟兽族·雷族怪兽的攻击力·守备力上升200。

		*/
		Id:       167,
		Password: "50913601",
		Name:     "山",
		Lc:       ygo.LC_SpellField, // 场地魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellField(ca.EffectMzoneHalo(func(c *ygo.Card) bool {
				return c.RaceIsDragon() || c.RaceIsWingedBeast() || c.RaceIsThunder()
			}, func(c *ygo.Card) {
				c.SetAtk(c.GetAtk() + 200)
				c.SetDef(c.GetDef() + 200)
			}, func(c *ygo.Card) {
				c.SetAtk(c.GetAtk() - 200)
				c.SetDef(c.GetDef() - 200)
			}))
			return true
		}, // 初始
		IsValid: true,
	})

	/* 36 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 168
		 调整:

		 中文名: 草原
		 卡片种类: 场地魔法
		 卡片密码: 86318356
		 罕见度: 平卡N
		 卡包: LB、EX-R(EX)、DL02、Starter Box
		 效果: 全部战士族·兽战士族怪兽的攻击力·守备力上升200。

		*/
		Id:       168,
		Password: "86318356",
		Name:     "草原",
		Lc:       ygo.LC_SpellField, // 场地魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellField(ca.EffectMzoneHalo(func(c *ygo.Card) bool {
				return c.RaceIsWarrior() || c.RaceIsWingedBeast()
			}, func(c *ygo.Card) {
				c.SetAtk(c.GetAtk() + 200)
				c.SetDef(c.GetDef() + 200)
			}, func(c *ygo.Card) {
				c.SetAtk(c.GetAtk() - 200)
				c.SetDef(c.GetDef() - 200)
			}))
			return true
		}, // 初始
		IsValid: true,
	})

	/* 37 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 169
		 调整:

		 中文名: 海
		 卡片种类: 场地魔法
		 卡片密码: 22702055
		 罕见度: 平卡N
		 卡包: LB、BE01、DL02、Starter Box
		 效果: 场上表侧表示存在的鱼族·海龙族·雷族·水族怪兽的攻击力·守备力上升200。场上表侧表示存在的机械族·炎族怪兽的攻击力·守备力下降200。

		*/
		Id:       169,
		Password: "22702055",
		Name:     "海",
		Lc:       ygo.LC_SpellField, // 场地魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellField(ca.EffectMzoneHalo(func(c *ygo.Card) bool {
				return c.RaceIsFish() || c.RaceIsSeaSerpent() || c.RaceIsThunder() || c.RaceIsWater() || c.RaceIsMachine() || c.RaceIsFire()
			}, func(c *ygo.Card) {
				if c.RaceIsMachine() || c.RaceIsFire() {
					c.SetAtk(c.GetAtk() - 200)
					c.SetDef(c.GetDef() - 200)
				} else {
					c.SetAtk(c.GetAtk() + 200)
					c.SetDef(c.GetDef() + 200)
				}
			}, func(c *ygo.Card) {
				if c.RaceIsMachine() || c.RaceIsFire() {
					c.SetAtk(c.GetAtk() + 200)
					c.SetDef(c.GetDef() + 200)
				} else {
					c.SetAtk(c.GetAtk() - 200)
					c.SetDef(c.GetDef() - 200)
				}
			}))
			return true
		}, // 初始
		IsValid: true,
	})

	/* 38 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 170
		 调整:

		 中文名: 暗
		 卡片种类: 场地魔法
		 卡片密码: 59197169
		 罕见度: 平卡N
		 卡包: LB、EX-R(EX)、DL02、Starter Box、TP10
		 效果: 场上表侧表示存在的恶魔族·魔法师族怪兽的攻击力·守备力上升200。场上表侧表示存在的天使族怪兽的攻击力·守备力下降200。

		*/
		Id:       170,
		Password: "59197169",
		Name:     "暗",
		Lc:       ygo.LC_SpellField, // 场地魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellField(ca.EffectMzoneHalo(func(c *ygo.Card) bool {
				return c.RaceIsFiend() || c.RaceIsSpellcaster() || c.RaceIsFairy()
			}, func(c *ygo.Card) {
				if c.RaceIsFairy() {
					c.SetAtk(c.GetAtk() - 200)
					c.SetDef(c.GetDef() - 200)
				} else {
					c.SetAtk(c.GetAtk() + 200)
					c.SetDef(c.GetDef() + 200)
				}
			}, func(c *ygo.Card) {
				if c.RaceIsFairy() {
					c.SetAtk(c.GetAtk() + 200)
					c.SetDef(c.GetDef() + 200)
				} else {
					c.SetAtk(c.GetAtk() - 200)
					c.SetDef(c.GetDef() - 200)
				}
			}))
			return true
		}, // 初始
		IsValid: true,
	})

	/* 39 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 81
		 调整:

		 中文名: 大地力量
		 卡片种类: 场地魔法
		 卡片密码: 56594520
		 罕见度: 平卡N
		 卡包: BE01、PS、YSD01、DL01、YU、JY
		 效果: 场上表侧表示存在的地属性怪兽攻击力上升500，守备力下降400。

		*/
		Id:       81,
		Password: "56594520",
		Name:     "大地力量",
		Lc:       ygo.LC_SpellField, // 场地魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterSpellField(ca.EffectMzoneHalo(func(c *ygo.Card) bool {
				return c.AttrIsEarth()
			}, func(c *ygo.Card) {
				c.SetAtk(c.GetAtk() + 500)
				c.SetDef(c.GetDef() - 400)

			}, func(c *ygo.Card) {
				c.SetAtk(c.GetAtk() - 500)
				c.SetDef(c.GetDef() + 400)
			}))
			return true
		}, // 初始
		IsValid: true,
	})

	/* 40 */
	cardBag.Register(&ygo.CardOriginal{
		/*
			 id: 407
			 调整:

			 [怨灵的湿地带]<怨霊の湿地帯>
			 ●全场怪兽召唤，反转召唤，特殊召唤的那个回合不能攻击。
			 ◇同一回合此卡发动前召唤，反转召唤，特殊召唤的怪兽也不能攻击★“不能攻击”是否等价于“不能进行攻击宣言”，调整中
			例：
			场上有表侧表示的此卡，
			 [静寂虫/静寂虫]效果适用，当回合召唤的怪兽攻击时，用
			 [月之书/月の書]的效果将
			 [静寂虫/静寂虫]变成里侧表示，此时此卡效果适用，攻击是否无效，调整中

			 中文名: 怨灵的湿地带
			 卡片种类: 永续魔法
			 卡片密码: 95220856
			 罕见度: 平卡N
			 卡包: BE02、LN、DL03、PE
			 效果: 全场怪兽召唤·反转召唤·特殊召唤的那个回合不能攻击。

		*/
		Id:       407,
		Password: "95220856",
		Name:     "怨灵的湿地带",
		Lc:       ygo.LC_SpellContinuous, // 永续魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 41 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 61
		 调整:

		 [卡通世界]<トゥーン·ワールド>
		 ●这张卡付出1000分才可以发动。
		 ◇COST
		 ◇效果处理事不会发生任何事
		 ◇此卡在场上正面表示的时候没有效果
		 ◇此卡的效果说明有另外一个版本，那个版本的卡也能用，但效果说明以此版本为准

		 中文名: 卡通世界
		 卡片种类: 永续魔法
		 卡片密码: 15259703
		 罕见度: 平卡N，银字R
		 卡包: BE01、PS、DL01、PE、GLD04
		 效果: 支付1000基本分发动。

		*/
		Id:       61,
		Password: "15259703",
		Name:     "卡通世界",
		Lc:       ygo.LC_SpellContinuous, // 永续魔法

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 42 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 100
		 调整:

		 [沙尘之大龙卷]<砂塵の大竜巻>
		 ●对方的场上的一张魔法卡或陷阱卡破坏。
		 ◇取对象
		 ●破坏后，可以把自己的手卡的一张魔法卡或陷阱卡覆盖在场上。
		 ◇任意发动
		 ◇因此效果覆盖的速攻魔法/陷阱不能立刻发动

		 中文名: 沙尘之大龙卷
		 卡片种类: 通常陷阱
		 卡片密码: 60082869
		 罕见度: 平卡N，银字R，黄金GR
		 卡包: BE01、CA、YSD01、YSD02、DL01、SD02、SD03、SD04、SD08、SD14、DT01、YSD04、DTP01、YSD05、JY、SY2、SK2、YSD06、GS03、DB12、ST12
		 效果: 选择对方场上存在的1张魔法·陷阱卡破坏。那之后，可以从自己手卡把1张魔法·陷阱卡盖放。

		*/
		Id:       100,
		Password: "60082869",
		Name:     "沙尘之大龙卷",
		Lc:       ygo.LC_TrapNormal, // 通常陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 43 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1154
		 调整:

		 [死之卡组破坏病毒]<死のデッキ破壊ウイルス>
		 [2011/03/01]
		 ●一只攻击力1000以下的暗属性怪兽做祭品。
		 ◇将暗属性怪兽做祭品是COST。
		 ◇可以将里侧表示的暗属性怪兽做祭品
		 ◇可以将里侧表示攻击力记载为「？」的暗属性怪兽做祭品。
		 ◇可以使用「灵魂交错/クロス·ソウル」将对方场上表侧表示的怪兽作为祭品。
		 ◇不可以使用「灵魂交错/クロス·ソウル」将对方场上里侧表示的怪兽作为祭品。
		 ●发动时对方的场上，手卡，发动后3回合之内抽的卡1500攻击力以上的怪兽全部破坏。
		 ◇在手卡中或抽到攻击力记载为「？」的怪兽不会被这个效果破坏。
		 ◇在魔法·陷阱区域的「玩具魔术师/トイ·マジシャン」不会被这个效果破坏。
		 ◇这个效果适用后，「人造人-念力震慑者/人造人間－サイコ·ショッカー」在场上表侧表示存在的场合「抽的卡1500攻击力以上的怪兽全部破坏」的效果仍然有效。
		 ◇发动后确认放置（SET）的怪兽并符合条件破坏的场合，视为以表侧表示的状态被破坏。
		 ◇因为「暗之诱惑/闇の誘惑」抽卡的场合，在处理完「暗之诱惑/闇の誘惑」的效果之后（抽卡→除外），再处理这个效果。

		 中文名: 死之卡组破坏病毒
		 卡片种类: 通常陷阱
		 卡片密码: 57728570
		 罕见度: 平卡N，银字R，黄金GR，金字UR，爆闪PR，立体UTR
		 卡包: PP05、SD12、GLD01、GS01、TU01、DPKB、KA、SK2、GDB1
		 效果: 把自己场上存在的1只攻击力1000以下的暗属性怪兽解放发动。对方场上存在的怪兽、对方的手卡、用对方回合计算的3回合内对方抽到的卡全部确认，攻击力1500以上的怪兽破坏。

		*/
		Id:       1154,
		Password: "57728570",
		Name:     "死之卡组破坏病毒",
		Lc:       ygo.LC_TrapNormal, // 通常陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 44 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 117
		 调整:

		 [补充要员]<補充要員>
		 ●自己的墓地的怪兽卡5张以上存在时才能发动。
		 ◇发动条件
		 ◇“怪兽卡”不单指“效果怪兽以外的攻击力1500以下的怪兽卡”
		 ●3张自己的墓地效果怪兽以外的攻击力1500以下的怪兽卡加入手卡。
		 ◇取对象
		 ◇可选择1~3张
		 ◇“效果怪兽以外的怪兽”的判定同
		 [绝对魔法禁止区域/絶対魔法禁止区域]

		 中文名: 补充要员
		 卡片种类: 通常陷阱
		 卡片密码: 36280194
		 罕见度: 平卡N
		 卡包: BE01、CA、DL01、YSD03、YU、SJ2
		 效果: 自己墓地有怪兽5只以上存在的场合才能发动。选择自己墓地存在的最多3只效果怪兽以外的攻击力1500以下的怪兽加入手卡。

		*/
		Id:       117,
		Password: "36280194",
		Name:     "补充要员",
		Lc:       ygo.LC_TrapNormal, // 通常陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 45 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1171
		 调整:

		 [破坏轮]<破壊輪>
		 [2011/03/01]
		 ●场上的表侧表示的一只怪兽破坏。互相受到被破坏怪兽的攻击力数值的伤害。
		 ◇取对象效果。
		 ◇那个怪兽没有被破坏的场合，伤害效果不适用。

		 中文名: 破坏轮
		 卡片种类: 通常陷阱
		 卡片密码: 83555666
		 罕见度: 金字UR，立体UTR
		 卡包: DPKB、KA、SK2、其他
		 效果: 场上的表侧表示的1只怪兽破坏。互相受到被破坏怪兽的攻击力数值的伤害。

		*/
		Id:       1171,
		Password: "83555666",
		Name:     "破坏轮",
		Lc:       ygo.LC_TrapNormal, // 通常陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 46 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1181
		 调整:

		 [魔术臂盾]
		 [15/04/13]
		 ●自己场上有怪兽存在的场合，对方怪兽的攻击宣言时才能发动。对方场上表侧表示存在的1只攻击怪兽以外的怪兽的控制权直到战斗阶段结束时得到，由那只怪兽受到攻击进行伤害计算。
		 ◇陷阱卡的效果，取对象
		 ◇不能取不会转移控制权的怪兽为对象
		 ◇「洗脑解除/洗脳解除」适用中这张卡可以发动且正常处理，伤害计算完毕后那只怪兽的控制权还给对手
		 ◇可以选择「地缚神/地缚神」怪兽等“对方不能选择这张卡作为攻击对象”的怪兽为对象并正常处理
		 ◇在效果处理中进行伤害计算，不能发动其他卡的效果（不开连锁的效果可以适用，这包括得到控制权的那只怪兽的不开连锁的效果）
		 ◇攻击怪兽的攻击力高于攻击表示的对象怪兽的攻击力的场合，超出的部分正常给我方造成战斗伤害
		 ◇攻击怪兽具有穿刺效果而对象怪兽是守备表示的场合，攻击力超出守备力的部分正常给我方造成穿刺战斗伤害
		 ◇连锁2以上发动的这张卡的效果导致有怪兽被战斗破坏的场合，战斗破坏确定的怪兽暂时留在场上，在这个连锁处理完毕后进入伤害步骤结束时，战斗破坏确定的怪兽送去墓地（其具有以此触发的效果的场合，可以发动）。

		 中文名: 魔术臂盾
		 卡片种类: 通常陷阱
		 卡片密码: 96008713
		 罕见度: 平卡N，金字UR
		 卡包: D09、SD15、DT03、GLD03、JY、SJ2、其他
		 效果: 自己场上有怪兽存在的场合，对方怪兽的攻击宣言时才能发动。对方场上表侧表示存在的1只攻击怪兽以外的怪兽的控制权直到战斗阶段结束时得到，由那只怪兽受到攻击。

		*/
		Id:       1181,
		Password: "96008713",
		Name:     "魔术臂盾",
		Lc:       ygo.LC_TrapNormal, // 通常陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 47 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1185
		 调整:

		 中文名: 恶魔的骰子
		 卡片种类: 通常陷阱
		 卡片密码: 00126218
		 罕见度: 平卡N，金字UR，银碎SER
		 卡包: TP05、JY、SJ2、GLD04、其他
		 效果: 掷1个骰子。对方控制的全部怪兽的攻击力守备力在回合结束前下降投掷出的数目×100。

		*/
		Id:       1185,
		Password: "00126218",
		Name:     "恶魔的骰子",
		Lc:       ygo.LC_TrapNormal, // 通常陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 48 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 123
		 调整:

		 [魔法礼帽]<マジカルシルクハット>
		 ●从卡组选择怪兽卡以外的2张卡出来和一只场上的自己的怪兽，之后卡组洗切。
		 ◇取怪兽为对象
		 ◇可以选择里侧表示的怪兽
		 ◇不能选择不能变为覆盖的TOKEN作对象
		 ◇自己场上没有2个空怪兽区时无法发动
		 ◇不能选择因另一张此卡的效果覆盖在场上的魔/陷卡
		 ◇效果处理时对象不在场上的场合，效果不发
		 ◇效果处理时取出2张卡
		 ◇必须且只能取出2张卡
		 ◇取出的2张卡需要确认
		 ●选择的卡洗切，覆盖防守表示放置在场上。
		 ◇这个动作视为特殊召唤
		 ●从卡组选出的2张卡当作怪兽卡使用（攻防0）并且在战斗流程结束时破坏。
		 ◇那些魔/陷卡没有等级/种族/属性的设定，当作普通怪兽
		 ◇只在场上算作怪兽，在场上以外的地方不算
		 ◇那些魔/陷卡可以通过
		 [血之代价/血の代償]的效果解放，用来进行上级召唤
		 ◇因此效果将
		 [死灵佐玛/死霊ゾーマ]当作怪兽使用并且被战破的场合，满足
		 [死灵佐玛/死霊ゾーマ]伤血效果的发动条件
		 ◇因此效果出现在怪兽区的魔/陷卡不可用作同调召唤的素材★因此效果造成的破坏是否入连锁，调整中★
		 [次元的裂缝/次元の裂け目]效果适用中，因此效果出现在怪兽区的魔/陷卡破坏的场合是除外还是送墓地，调整中★因此效果出现在怪兽区的表侧的
		 [扰乱魔术/おジャマジック]可否作为
		 [扰乱骑士/おジャマ·ナイト]的融合素材，调整中★因此效果出现在怪兽区的表侧的魔/陷卡因其他卡的效果装备
		 [同调英雄/シンクロ·ヒーロー]的场合，星数怎么算，调整中
		 ●这个效果只能在对方的战斗流程使用。
		 ◇可在战斗阶段结束宣言时发动此卡
		 ◇可在攻击宣言到战斗步骤发动，此时会造成攻击卷回

		 中文名: 魔术礼帽
		 卡片种类: 通常陷阱
		 卡片密码: 81210420
		 罕见度: 银字R
		 卡包: BE01、CA、DL01、DPYG、YU
		 效果: 对方的战斗阶段时才能发动。从自己卡组选择怪兽以外的2张卡。那2张当作怪兽（攻／守0），和自己场上存在的1只怪兽合在一起洗切并里侧守备表示盖放。从卡组选择特殊召唤的2张卡在战斗阶段结束时破坏。

		*/
		Id:       123,
		Password: "81210420",
		Name:     "魔术礼帽",
		Lc:       ygo.LC_TrapNormal, // 通常陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 49 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1273
		 调整:

		 [魔法筒]<魔法の筒>
		 [09/09/20]
		 ●对方一只怪兽的攻击无效，对方基本分受到那只怪兽的攻击力数值的伤害。
		 ◇取对象效果。
		 ◇不能对应
		 [魔法筒/魔法の筒]的效果发动另一张
		 [魔法筒/魔法の筒]。
		 ◇效果处理时对象怪兽不在场上表侧表示存在的场合，效果不发。

		 中文名: 魔法筒
		 卡片种类: 通常陷阱
		 卡片密码: 62279055
		 罕见度: 平卡N，黄金GR，金字UR
		 卡包: PP04、YSD01、YSD02、SD06、SD12、SD16、GS02、YU、SY2、DB12、ST12、GDB1
		 效果: 对方怪兽的攻击宣言时才能发动。把1只对方怪兽的攻击无效，给予对方基本分那只怪兽的攻击力数值的伤害。

		*/
		Id:       1273,
		Password: "62279055",
		Name:     "魔法筒",
		Lc:       ygo.LC_TrapNormal, // 通常陷阱

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterTrapNormal(ygo.Declaration, func(c *ygo.Card) {
				c.StopOnce(ygo.Declaration)
				c.GetSummoner().ChangeLp(-c.GetAtk())
			})
			return true
		}, // 初始
		IsValid: false,
	})

	/* 50 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 180
		 调整:

		 [夹击]<はさみ撃ち>
		 [10/05/13]
		 ◇发动时选择自己场上的2只怪兽和对方场上的1只怪兽（取对象）
		 ◇自己场上只有1怪兽存在或者对方场上没有怪兽存在的场合，这张卡不能发动
		 ◇效果处理时选择的怪兽不在场上存在的场合，破坏效果不适用
		 ◇选择了不会被卡的效果破坏的卡为这个效果对象的场合，这些卡不会被破坏，选择的其他卡破坏

		 中文名: 夹击
		 卡片种类: 通常陷阱
		 卡片密码: 83887306
		 罕见度: 平卡N
		 卡包: LB、EX-R(EX)、DL02、Starter Box、TP06
		 效果: 选择自己场上存在的2只怪兽和对方场上存在的1只怪兽破坏。

		*/
		Id:       180,
		Password: "83887306",
		Name:     "夹击",
		Lc:       ygo.LC_TrapNormal, // 通常陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 51 */
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

	/* 52 */
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

	/* 53 */
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

	/* 54 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 286
		 调整:

		 [换位]<シフトチェンジ>
		 ●对方的魔法，陷阱，战斗对自己场上的一只怪兽的指定的时候才可以发动。
		 ◇发动条件，对应魔/陷的场合连锁其发动，对应战斗的场合发动时点同
		 [神圣防护罩-反射镜力-/聖なるバリア－ミラーフォース－]
		 ◇怪兽直接攻击的场合，此卡不能发动★可否连锁
		 [天使的手镜/天使の手鏡]发动此卡，调整中
		 ●其他的自己的场上的怪兽作为对象代替。
		 ◇取要作为新对象的怪兽为对象
		 ◇不能取不符合条件的卡为对象★效果处理时对象不符合条件的场合，是此卡不发还是被此卡对应的魔/陷不发，调整中

		 中文名: 换位
		 卡片种类: 通常陷阱
		 卡片密码: 59560625
		 罕见度: 平卡N，银字R
		 卡包: BE01、TB 、DL03、YU、SY2
		 效果: 对方的魔法·陷阱，战斗对自己场上的1只怪兽的指定的时候才可以发动。其他的自己的场上的怪兽作为对象代替。

		*/
		Id:       286,
		Password: "59560625",
		Name:     "换位",
		Lc:       ygo.LC_TrapNormal, // 通常陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 55 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 31
		 调整:

		 [天使的手镜]<天使の手鏡>
		 [2010/10/01]
		 ●以场上的1只怪兽为对象的对方的魔法，转移到别的正确对象。
		 ◇不能在伤害步骤中发动
		 ◇必须紧接着对应的卡的发动而发动

		 中文名: 天使的手镜
		 卡片种类: 通常陷阱
		 卡片密码: 17653779
		 罕见度: 平卡N
		 卡包: BE01、MR、DL01、JY
		 效果: 把以场上1只怪兽为对象发动的对方的魔法，转移给其他正确的对象。

		*/
		Id:       31,
		Password: "17653779",
		Name:     "天使的手镜",
		Lc:       ygo.LC_TrapNormal, // 通常陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 56 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 355
		 调整:

		 中文名: 强欲之瓶
		 卡片种类: 通常陷阱
		 卡片密码: 83968380
		 罕见度: 平卡N，银字R
		 卡包: BE01、SM、YSD01、YSD02、DL03、SD03、PE、SJ2
		 效果: 从自己的卡组抽1张卡。

		*/
		Id:       355,
		Password: "83968380",
		Name:     "强欲之瓶",
		Lc:       ygo.LC_TrapNormal, // 通常陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 57 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 98
		 调整:

		 [圣精灵的祝福]<ホーリー·エルフの祝福>
		 ●每一只场上存在的怪兽，恢复300分。
		 ◇效果处理时判定数目
		 ◇数且只数出现在怪兽区的卡

		 中文名: 圣精灵的祝福
		 卡片种类: 通常陷阱
		 卡片密码: 98299011
		 罕见度: 平卡N
		 卡包: BE01、CA、DL01、KA、SK2
		 效果: 自己回复场上存在的怪兽数量×300的数值的基本分。

		*/
		Id:       98,
		Password: "98299011",
		Name:     "圣精灵的祝福",
		Lc:       ygo.LC_TrapNormal, // 通常陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 58 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 99
		 调整:

		 [盗墓者]<墓荒らし>
		 ●选择对方墓地的一张魔法卡，在回合结束前加入自己的手卡并可以使用。
		 ◇取对象★不使用的场合，送墓地时点是手牌调整之后还是结束阶段中，调整中
		 ◇那张卡覆盖在场上的场合，也要送墓地
		 ◇那张卡因其他效果洗回卡组的场合，不用送墓地
		 ●使用那张魔法卡的场合，受到2000分的伤害。
		 ◇不入连锁
		 ◇“使用”包括“发动”与“作为COST”
		 ◇因效果舍入墓地的场合，不受伤害★发动的场合，暂调整为在魔法卡效果处理完后造成伤害★发动被无效的场合，是否受伤害，调整中★不发的场合，是否受伤害，调整中★作为COST的场合，暂调整为在支付完后造成伤害★是否占时点，调整中

		 中文名: 盗墓者
		 卡片种类: 通常陷阱
		 卡片密码: 61705417
		 罕见度: 平卡N，银字R
		 卡包: BE01、CA、DL01、JY、SJ2
		 效果: 选择对方墓地的1张魔法卡，直到回合结束时可以作为自己的手卡使用。使用那张魔法卡的场合，受到2000分伤害。

		*/
		Id:       99,
		Password: "61705417",
		Name:     "盗墓者",
		Lc:       ygo.LC_TrapNormal, // 通常陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 59 */
	cardBag.Register(&ygo.CardOriginal{
		/*
			 id: 101
			 调整:

			 [活死人的呼声]
			 [14/04/27]
			 ●①：以自己墓地1只怪兽为对象才能把这张卡发动。那只怪兽攻击表示特殊召唤。这张卡从场上离开时那只怪兽破坏。那只怪兽破坏时这张卡破坏。
			 ◇特殊召唤效果取对象。
			 ◇此卡的“卡的发动”在伤害步骤不能进行。
			 ◇把此卡自身破坏或把对象怪兽破坏的处理不开连锁，该效果不取对象。此卡特殊召唤的
			 [异次元龙/異次元竜
			トワイライトゾーンドラゴン]不会被该效果破坏。
			 ◇
			 [宫廷的规矩/宮廷のしきたり]或
			 [白龙忍者/白竜の忍者]效果适用中，此卡不能把自身破坏。
			 ◇此卡的效果特殊召唤的怪兽变成里侧表示或被
			 [亚空间物质传送装置/亚空间物质传送装置]暂时除外再返回场上，此卡与那只怪兽的关系即告消失。
			 ◇此卡特殊召唤怪兽后，此卡的效果被
			 [王宫的通告/王宮のお触れ]无效，然后此卡离场（或此卡和王宫的通告同时离场），则此卡特殊召唤的那只怪兽不会被破坏。
			 ◇此卡特殊召唤怪兽后，此卡的效果被
			 [王宫的通告/王宮のお触れ]无效，然后王宫的通告离场，则此卡的效果恢复，再让此卡离场的话此卡特殊召唤的怪兽会被此卡的效果破坏。
			 ◇此卡特殊召唤
			 [人造人-念力震慑者/人造人間－サイコ·ショッカー]，此卡的效果被其无效，此卡离场不会导致念力震慑者被破坏。
			 ◇此卡特殊召唤
			 [人造人-念力震慑者/人造人間－サイコ·ショッカー]，此卡的效果被其无效，然后念力震慑者被战斗或其他效果破坏，则此卡的效果恢复，此卡被自己的效果破坏。
			 ◇
			 [神禽王
			亚力克特/神禽王アレクトール]把此卡的效果无效化，然后此卡离场，此卡所特殊召唤的怪兽不会被此卡的效果破坏。

			 中文名: 活死人的呼声
			 卡片种类: 永续陷阱
			 卡片密码: 97077563
			 罕见度: 平卡N，黄金GR，面闪SR
			 卡包: DP04、BE01、CA、YSD02、DL01、SD01、SD03、SD04、SD05、SD06、SD08、SD13、GS02、JY、SK2、DB12、SD22、ST12、SD25、SD26、DS13、GDB1
			 效果: 选择自己墓地1只怪兽，表侧攻击表示特殊召唤。这张卡从场上离开时，那只怪兽破坏。那只怪兽破坏时这张卡破坏。

		*/
		Id:       101,
		Password: "97077563",
		Name:     "活死人的呼声",
		Lc:       ygo.LC_TrapContinuous, // 永续陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 60 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 120
		 调整:

		 [神圣光辉]<聖なる輝き>
		 ●怪兽不能覆盖防守表示放置。
		 ◇不能翻开已经在场上存在的里侧守备表示怪兽
		 ◇此卡效果适用中，包含让怪兽变成里侧守备表示的效果可以发动，效果处理时不发
		 ●覆盖防守表示放置的场合，变成打开防守表示召唤。
		 ◇里侧守备表示特殊召唤的场合，变成表侧守备表示特殊召唤
		 ◇表侧守备表示通常召唤除表现形式不同外，与普通的召唤没有区别

		 中文名: 神圣光辉
		 卡片种类: 永续陷阱
		 卡片密码: 62867251
		 罕见度: 平卡N
		 卡包: BE01、CA、DL01、DT03、KA、SK2
		 效果: 只要这张卡在场上存在，不能把怪兽盖放。此外，把怪兽盖放的场合必须变成表侧守备表示。

		*/
		Id:       120,
		Password: "62867251",
		Name:     "神圣光辉",
		Lc:       ygo.LC_TrapContinuous, // 永续陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 61 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 121
		 调整:

		 [正正堂堂]<正々堂々>
		 ●玩家互相在自己的回合永续公开手卡。
		 ◇此卡效果适用中，只有回合玩家需要公开手牌
		 ◇此卡效果适用中，回合玩家必须公开手牌

		 中文名: 正正堂堂
		 卡片种类: 永续陷阱
		 卡片密码: 08951260
		 罕见度: 平卡N
		 卡包: BE01、CA、DL01、PE
		 效果: 双方玩家必须在各自的回合把手卡全部持续公开。

		*/
		Id:       121,
		Password: "08951260",
		Name:     "正正堂堂",
		Lc:       ygo.LC_TrapContinuous, // 永续陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 62 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1229
		 调整:

		 [暗之咒缚]<闇の呪縛>
		 ●指定对方的一只表侧表示的怪兽，指定的怪兽攻击力下降700，不能攻击和改变表示形式。
		 ◇取对象效果。
		 ◇控制权转移后这个效果仍然适用。
		 ●指定的怪兽不在场上存在时，这张卡破坏。
		 ◇对象怪兽被「月之书/月の書」等卡变成里侧表示的场合，这张卡不会破坏，将以无意义的状态留场。
		 ◇(接上)此时不能再选择其他的怪兽作为这个效果的对象。

		 中文名: 暗之咒缚
		 卡片种类: 永续陷阱
		 卡片密码: 29267084
		 罕见度: 平卡N，金字UR
		 卡包: LE03、DT05、TP07、KA、SK2
		 效果: 指定对方的1只表侧表示的怪兽，指定的怪兽攻击力下降700，不能攻击和改变表示形式。指定的怪兽不在场上存在时，这张卡破坏。

		*/
		Id:       1229,
		Password: "29267084",
		Name:     "暗之咒缚",
		Lc:       ygo.LC_TrapContinuous, // 永续陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 63 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 164
		 调整:

		 中文名: 龙族封印之壶
		 卡片种类: 永续陷阱
		 卡片密码: 50045299
		 罕见度: 平卡N
		 卡包: LB、BE01、EX-R(EX)、DL02、Starter Box、PE
		 效果: 只要这张卡在场上存在，场上表侧表示存在的龙族怪兽全部变成守备表示，不能把表示形式变更。

		*/
		Id:       164,
		Password: "50045299",
		Name:     "龙族封印之壶",
		Lc:       ygo.LC_TrapContinuous, // 永续陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 64 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 542
		 调整:

		 中文名: 卡通防御
		 卡片种类: 永续陷阱
		 卡片密码: 43509019
		 罕见度: 平卡N，平罕NR
		 卡包: BE02、DL04、PE
		 效果: 自己场上表侧表示存在的4星以下的卡通怪兽受到攻击宣言时，这个攻击作为对玩家的直接攻击处理。

		*/
		Id:       542,
		Password: "43509019",
		Name:     "卡通防御",
		Lc:       ygo.LC_TrapContinuous, // 永续陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 65 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 554
		 调整:

		 [生命吸收装置]<生命吸収装置>
		 ●每次自己的准备流程，恢复之前的那个自己的回合付出的基本分的一半。
		 ◇相当于永续效果，不入连锁
		 ◇无论上个自己的回合支付多少次，此效果只回复一次
		 ◇复数此卡效果适用的场合，回复复数次★此卡发动前支付的LP是否计算在内，调整中

		 中文名: 生命吸收装置
		 卡片种类: 永续陷阱
		 卡片密码: 14318794
		 罕见度: 平卡N
		 卡包: SC、DL05、PE
		 效果: 每次自己的准备阶段，回复之前的那个自己的回合支付的基本分的一半。

		*/
		Id:       554,
		Password: "14318794",
		Name:     "生命吸收装置",
		Lc:       ygo.LC_TrapContinuous, // 永续陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 66 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 628
		 调整:

		 [不吉的占卜]<不吉な占い>
		 ●自己的准备流程时，随机选择对方的一张手卡。猜那张选择的卡的种类（怪兽/魔法/陷阱）。猜中的场合，给与对方700分的伤害。这个效果一个回合只能使用一次。★是否入连锁，调整中
		 ◇就算对方手牌处于公开状态，也会暂时不公开，洗切后随机丢弃
		 ◇猜错的场合，什么事都不会发生

		 中文名: 不吉的占卜
		 卡片种类: 永续陷阱
		 卡片密码: 56995655
		 罕见度: 平卡N
		 卡包: BE02、MA、DL05、PE
		 效果: 自己的准备阶段时，随机选择对方的1张手卡。猜那张选择的卡的种类（怪兽·魔法·陷阱）。猜中的场合，给予对方700分的伤害。这个效果1个回合只能使用1次。

		*/
		Id:       628,
		Password: "56995655",
		Name:     "不吉的占卜",
		Lc:       ygo.LC_TrapContinuous, // 永续陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 67 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 7
		 调整:

		 [六芒星之咒缚]<六芒星の呪縛>
		 [2010/10/01]
		 ●选择对方场上存在的1只怪兽发动。被选择的对方怪兽不能攻击，也不能改变表示形式。被选择的对方怪兽破坏时，这张卡破坏。
		 ◇发动时选择对方场上存在的1只怪兽（取对象）
		 ◇被选择怪兽一度从场上离开后再返回的场合，和这张卡的关系切断
		 ◇被选择怪兽由于其他卡的效果从表侧变为里侧或里侧变为表侧的场合，关系不会切断
		 ◇被选择怪兽以破坏以外方式从场上离开的场合，这张卡无意义留场
		 ◇「王宫的通告/王宮のお触れ」的效果适用后，「王宫的通告/王宮のお触れ」从场上离开的场合，这张卡无意义留场

		 中文名: 六芒星之咒缚
		 卡片种类: 永续陷阱
		 卡片密码: 18807108
		 罕见度: 平卡N，银字R，金字UR，爆闪PR
		 卡包: BE01、MR、DL01、YSD03、DPYG、YU、SY2
		 效果: 选择对方场上存在的1只怪兽发动。选择的怪兽不能攻击，也不能把表示形式变更。选择的怪兽破坏时，这张卡破坏。

		*/
		Id:       7,
		Password: "18807108",
		Name:     "六芒星之咒缚",
		Lc:       ygo.LC_TrapContinuous, // 永续陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 68 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 97
		 调整:

		 中文名: 真实之眼
		 卡片种类: 永续陷阱
		 卡片密码: 34694160
		 罕见度: 平卡N
		 卡包: BE01、CA、DL01、YU、PE、SY2
		 效果: 只要这张卡在场上表侧表示存在，对方把手卡全部持续公开。对方的准备阶段时对方手卡有魔法卡存在的场合，对方回复1000基本分。

		*/
		Id:       97,
		Password: "34694160",
		Name:     "真实之眼",
		Lc:       ygo.LC_TrapContinuous, // 永续陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 69 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1180
		 调整:

		 中文名: 攻击无力化
		 卡片种类: 反击陷阱
		 卡片密码: 14315573
		 罕见度: 平卡N，银字R，金字UR
		 卡包: DP01、YSD02、SD09、SD11、DT01、DTP01、KA、SK2、DB12
		 效果: 对方怪兽的攻击宣言时才能发动。把1只对方怪兽的攻击无效，战斗阶段结束。

		*/
		Id:       1180,
		Password: "14315573",
		Name:     "攻击无力化",
		Lc:       ygo.LC_TrapCounter, // 反击陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 70 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 385
		 调整:

		 中文名: 力场
		 卡片种类: 反击陷阱
		 卡片密码: 70344351
		 罕见度: 银字R
		 卡包: BE02、LN、DL03、KA、PE
		 效果: 以场上的1只怪兽为对象的魔法的发动无效并破坏那张卡。

		*/
		Id:       385,
		Password: "70344351",
		Name:     "力场",
		Lc:       ygo.LC_TrapCounter, // 反击陷阱

		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 71 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1005
		 调整:

		 中文名: 黑混沌之魔术师
		 卡片种类: 仪式怪兽
		 卡片密码: 30208479
		 种族: 魔法师
		 属性: 暗
		 星级: 8
		 攻击力: 2800
		 防御力: 2600
		 罕见度: 面闪SR，斜碎SCR，立体UTR
		 卡包: 306、PP03、TP10、YU
		 效果: 仪式：通过仪式魔法卡「混沌-黑魔术的仪式」特殊召唤。特殊召唤时，必须以场上和/或手卡中合计8颗星以上的怪兽作为祭品。

		*/
		Id:       1005,
		Password: "30208479",
		Name:     "黑混沌之魔术师",
		Lc:       ygo.LC_MonsterRitual, // 仪式怪兽

		Level: 8,
		La:    ygo.LA_Dark,        // 暗
		Lr:    ygo.LR_Spellcaster, // 魔法师
		Atk:   2800,
		Def:   2600,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 72 */
	cardBag.Register(&ygo.CardOriginal{
		/*
			 id: 19
			 调整:

			 [纳祭之魔]<サクリファイス>
			 [10/12/09]
			 ●指定对方一只怪兽，作为这张卡的装备卡装备(这个效果一个回合只能使用一次，同时装备的怪兽只能一只)。
			 ◇起动效果（进入连锁）
			 ◇发动时选择对方场上存在的1只怪兽（取对象）★效果处理时选择的怪兽控制权转移到自己场上的场合，这个效果如何处理？调整中
			 ◇不能选择不能转移控制权的怪兽
			 ◇可以选择对方场上里侧表示的怪兽为这个效果的对象
			 ◇效果处理时此卡不在场上正面表示的场合，将对象送入墓地
			 ◇因效果处理时此卡不在场上正面表示而将对象送入墓地的场合，那张怪兽是从对手场上怪兽区送入墓地
			 ◇
			 [士气高扬/士気高扬]的效果适用
			 ◇可用这个效果装备TOKEN
			 ◇那只装备怪兽不能随便取下来
			 ◇装备怪兽被视作“装备魔法”但没有效果
			 ◇那只装备怪兽是覆盖表示的场合,可以成为
			 [二重身/ドッペルゲンガー]
			 [扑灭之使徒/扑灭の使徒]的对象
			 ◇
			 [洗脑解除/洗脳解除]之类归还控制权的效果不会影响那只装备怪兽
			 ◇装备怪兽的正确装备对象只有此卡，不包括同名卡
			 ◇此卡的控制权转移给对手的场合，那只装备怪兽也转移
			 ◇那只装备怪兽是正面表示的场合，可以放置
			 [魔法防护器/マジック·ガードナー]的指示物
			 ◇这张卡因这个效果装备了「龙骑兵团-重标枪龙/ドラグニティ－ピルム」的场合，这张卡可以进行直接攻击
			 ◇这张卡因这个效果装备了「球体时限炸弹/スフィア·ボム球体时限爆弾」的场合，「球体时限炸弹/スフィア·ボム球体时限爆弾」的效果不适用★这张卡因这个效果装备了「武器手套/アームズ·エイド」的场合，这张卡战斗破坏怪兽送去墓地时，『给与对方基本分破坏怪兽的攻击力数值的伤害』的效果能否发动？调整中
			 ◇“一回合只有一次”会因暂时离场或者变成里侧表示而重置注：此效果俗称“吸怪”
			 ●这张卡的攻击力和防守力使用装备怪兽的数值。
			 ◇永续效果
			 ◇表侧表示的装备怪兽按卡上面写的攻击/守备力算，里侧表示的装备怪兽按攻击/守备力都是0算
			 ◇TOKEN/陷阱怪兽的场合,按卡上面写的算
			 ◇攻击/守备力是?的怪兽按0算
			 ◇
			 [天邪鬼的诅咒/あまのじゃくの呪い]効果不适用
			 ●战斗造成这张卡破坏时，装备的怪兽代替破坏。
			 ◇永续效果
			 ◇这个破坏按效果破坏处理
			 ◇变成机械族的这张卡因这个效果装备了「强化支援机械·重装武器/强化支援メカ·ヘビーウェポン」同盟怪兽的场合，这张卡被战斗破坏时，双方受到同样伤害。★
			 [黑蝎-拆除陷阱的克里夫/黒蠍－罠はずしのクリフ]在攻击吸收了怪兽的此卡并且此卡确定被战破的场合，可否选择代破的那只装备怪兽为对象；可以的场合，是否不能代破而使此卡战破，调整中★那只装备怪兽是
			 [有翼幻兽
			奇美拉/有翼幻獣キマイラ]的场合,调整中
			 ◇此卡装备的怪兽有
			 [魔法防护器/マジック·ガードナー]的指示物时，此卡被战破的场合，可以去掉那个指示物来代替破坏
			 ●超过的伤害数值对方也承受。
			 ◇诱发效果，仅当此卡因自身效果装备着怪兽的场合才能发动
			 ◇这张卡有被战斗破坏的事实（即使使用了装备怪兽进行代替破坏）才能发动
			 ◇这个伤害属于效果伤害。
			 ◇双方基本分只有1000时，攻击力为1000的有怪兽装备的这张卡向对方场上攻击表示攻击力为3000的怪兽进行攻击的场合，（在没有其他卡的效果适用而不会改变这个事实的场合）双方基本分同时归零，平局。
			 ◇自己场上存在的「仪式之槛/仪式の槛」效果适用中时，自己不受战斗伤害，不适用。
			 ◇同
			 [黑蝎-拆除陷阱的克里夫/黒蠍－罠はずしのクリフ]
			 ◇可以连锁
			 [地狱的冷枪/地狱の扉越し铳]
			 [黑板擦的陷阱/黒板消しの罠]注：虽然卡面上没写，但此卡不能通常召唤，遵守苏生限制

			 中文名: 纳祭之魔
			 卡片种类: 仪式怪兽
			 卡片密码: 64631466
			 种族: 魔法师
			 属性: 暗
			 星级: 1
			 攻击力: 0
			 防御力: 0
			 罕见度: 面闪SR，金字UR，爆闪PR
			 卡包: BE01、MR、DL01、PE
			 效果: 仪式：「幻想的仪式」降临。1回合1次，可以选择对方场上存在的1只怪兽，当作装备卡使用只有1只给这张卡装备。这张卡的攻击力·守备力变成这张卡的效果装备的怪兽的各自数值。因这个效果有怪兽装备的场合，给予对方基本分和自己受到的战斗伤害相同的伤害。此外，这张卡被战斗破坏的场合，作为代替把这张卡的效果装备的怪兽破坏。

		*/
		Id:       19,
		Password: "64631466",
		Name:     "纳祭之魔",
		Lc:       ygo.LC_MonsterRitual, // 仪式怪兽

		Level: 1,
		La:    ygo.LA_Dark,        // 暗
		Lr:    ygo.LR_Spellcaster, // 魔法师
		Atk:   0,
		Def:   0,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 73 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1118
		 调整:

		 [漆黑之豹战士]<漆黑の豹戦士パンサーウォリアー>
		 [2011/04/28]
		 ●每次用自己场上一只怪兽作祭品才能攻击。
		 ◇永续效果（不进入连锁）
		 ◇一只怪兽作祭品是COST。

		 中文名: 漆黑之豹战士
		 卡片种类: 效果怪兽
		 卡片密码: 42035044
		 种族: 兽战士
		 属性: 地
		 星级: 4
		 攻击力: 2000
		 防御力: 1600
		 罕见度: 平卡N，金字UR
		 卡包: LE03、DT04、JY
		 效果: 每次用自己场上1只怪兽作祭品才能攻击。

		*/
		Id:       1118,
		Password: "42035044",
		Name:     "漆黑之豹战士",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Earth,        // 地
		Lr:    ygo.LR_BeastWarrior, // 兽战士
		Atk:   2000,
		Def:   1600,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 74 */
	cardBag.Register(&ygo.CardOriginal{
		/*
			 id: 1120
			 调整:

			 [巴风特]<バフォメット>
			 [2011/04/28]
			 ●这张卡召唤成功时，在卡组里选一张「幻兽王
			卡杰」加入手卡。
			 ◇诱发效果（进入连锁）。
			 ◇强制发动。

			 中文名: 巴风特
			 卡片种类: 效果怪兽
			 卡片密码: 77207191
			 种族: 恶魔
			 属性: 暗
			 星级: 5
			 攻击力: 1400
			 防御力: 1800
			 罕见度: 银字R，金字UR
			 卡包: ABPF(607)、YU、其他
			 效果: 这张卡召唤、反转召唤成功时，可以在卡组里选1张「幻兽王 加泽尔」加入手卡。

		*/
		Id:       1120,
		Password: "77207191",
		Name:     "巴风特",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 5,
		La:    ygo.LA_Dark,  // 暗
		Lr:    ygo.LR_Fiend, // 恶魔
		Atk:   1400,
		Def:   1800,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 75 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1139
		 调整:

		 [邪眼幻想师]<ダーク·アイズ·イリュージョニスト>
		 ●反转效果：这张卡场上存在的时候，指定的一只怪兽永续不能攻击。
		 ◇进入连锁。
		 ◇取对象效果。

		 中文名: 邪眼幻想师
		 卡片种类: 效果怪兽
		 卡片密码: 38247752
		 种族: 魔法师
		 属性: 暗
		 星级: 2
		 攻击力: 0
		 防御力: 1400
		 罕见度: 银字R，金字UR
		 卡包: PTDN(507)、VB02、PE
		 效果: 反转：这张卡场上存在的时候，指定的1只怪兽永续不能攻击。

		*/
		Id:       1139,
		Password: "38247752",
		Name:     "邪眼幻想师",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 2,
		La:    ygo.LA_Dark,        // 暗
		Lr:    ygo.LR_Spellcaster, // 魔法师
		Atk:   0,
		Def:   1400,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 76 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1141
		 调整:

		 [火箭战士]<ロケット戦士>
		 ●这张卡的效果自己的战斗流程有效。这张卡受到的战斗伤害为0。
		 ◇永续效果（不进入连锁）。
		 ●受到这张卡的攻击的怪兽回合结束前攻击力下降500。
		 ◇诱发效果（进入连锁）。
		 ◇强制发动。
		 ◇在伤害计算后发动。

		 中文名: 火箭战士
		 卡片种类: 效果怪兽
		 卡片密码: 30860696
		 种族: 战士
		 属性: 光
		 星级: 4
		 攻击力: 1500
		 防御力: 1300
		 罕见度: 平卡N，金字UR
		 卡包: TP06、VB03、JY
		 效果: 自己的战斗阶段时，这张卡不会被战斗破坏，这张卡的战斗发生的对自己的战斗伤害变成0。这张卡向怪兽攻击的场合，伤害计算后攻击对象怪兽的攻击力直到结束阶段时下降500。

		*/
		Id:       1141,
		Password: "30860696",
		Name:     "火箭战士",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Light,   // 光
		Lr:    ygo.LR_Warrior, // 战士
		Atk:   1500,
		Def:   1300,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 77 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1176
		 调整:

		 中文名: 磁石战士 电磁武神
		 卡片种类: 效果怪兽
		 卡片密码: 75347539
		 种族: 岩石
		 属性: 地
		 星级: 8
		 攻击力: 3500
		 防御力: 3850
		 罕见度: 金字UR
		 卡包: TP08、YU、其他
		 效果: 这张卡不能通常召唤。从自己的手卡·场上把「磁石战士α」「磁石战士β」「磁石战士γ」各1只解放的场合可以特殊召唤。并且，可以把自己场上存在的这张卡解放，选择自己墓地存在的「磁石战士α」「磁石战士β」「磁石战士γ」各1只特殊召唤。

		*/
		Id:       1176,
		Password: "75347539",
		Name:     "磁石战士 电磁武神",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 8,
		La:    ygo.LA_Earth, // 地
		Lr:    ygo.LR_Rock,  // 岩石
		Atk:   3500,
		Def:   3850,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 78 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1225
		 调整:

		 [模仿的幻想师]<ものマネ幻想師>
		 ●这张卡召唤，反转召唤，特殊召唤成功时，选择对方的1只怪兽，拥有那只怪兽一样的原来的攻击力和守备力。
		 ◇诱发效果（进入连锁）。
		 ◇强制发动。
		 ◇取对象效果。
		 ◇获得那个怪兽的原本攻击力和守备力不视为这张卡的原本攻击力和守备力。

		 中文名: 模仿的幻想师
		 卡片种类: 效果怪兽
		 卡片密码: 26376390
		 种族: 魔法师
		 属性: 光
		 星级: 1
		 攻击力: 0
		 防御力: 0
		 罕见度: 平卡N，金字UR
		 卡包: LE03、YSD03、JY、DT11
		 效果: 这张卡召唤·反转召唤·特殊召唤成功时，选择对方场上表侧表示存在的1只怪兽发动。这张卡的攻击力·守备力变成和选择的怪兽的原本的攻击力·守备力相同数值。

		*/
		Id:       1225,
		Password: "26376390",
		Name:     "模仿的幻想师",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 1,
		La:    ygo.LA_Light,       // 光
		Lr:    ygo.LR_Spellcaster, // 魔法师
		Atk:   0,
		Def:   0,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 79 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1228
		 调整:

		 [魔法神灯]<マジック·ランプ>
		 ●里侧守备表示的这张卡受到攻击时，可以选对方的1只怪兽代替那个攻击。
		 ◇诱发效果（进入连锁）。
		 ◇任意发动。
		 ◇自己场上的这张卡发动效果指定对方场上里侧守备表示的这张卡时，那张「魔法神灯/マジック·ランプ」的效果不会发动。
		 ●只要这张卡在场上存在，可以从手卡特殊召唤「灯之魔精/ランプの魔精·ラ·ジーン」。
		 ◇起动效果（进入连锁）。
		 ◇只能在自己的主要流程发动。
		 ◇一回合可以复数次发动这个效果。
		 ◇效果处理时从手卡特殊召唤一张「灯之魔精/ランプの魔精·ラ·ジーン」。

		 中文名: 魔法神灯
		 卡片种类: 效果怪兽
		 卡片密码: 54912977
		 种族: 魔法师
		 属性: 风
		 星级: 3
		 攻击力: 900
		 防御力: 1400
		 罕见度: 平卡N，金字UR
		 卡包: LE03、TP10、DPKB、KA、SK2
		 效果: 里侧守备表示的这张卡受到攻击时，选择对方的攻击怪兽以外的一只怪兽发动，攻击怪兽和选择的怪兽进行战斗，进行伤害计算。只要这张卡在场上存在，可以从手卡特殊召唤「灯之魔精」。

		*/
		Id:       1228,
		Password: "54912977",
		Name:     "魔法神灯",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 3,
		La:    ygo.LA_Wind,        // 风
		Lr:    ygo.LR_Spellcaster, // 魔法师
		Atk:   900,
		Def:   1400,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 80 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1293
		 调整:

		 中文名: 龙之支配者
		 卡片种类: 效果怪兽
		 卡片密码: 17985575
		 种族: 魔法师
		 属性: 暗
		 星级: 4
		 攻击力: 1200
		 防御力: 1100
		 罕见度: 平卡N，金字UR
		 卡包: EX-R(EX)、TP07、DPKB、KA、SK2、SD22
		 效果: 只要这张卡在场上表侧表示存在，场上的龙族怪兽不能成为魔法·陷阱·效果怪兽的效果的对象。

		*/
		Id:       1293,
		Password: "17985575",
		Name:     "龙之支配者",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Dark,        // 暗
		Lr:    ygo.LR_Spellcaster, // 魔法师
		Atk:   1200,
		Def:   1100,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 81 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 130
		 调整:

		 [混沌壶]<カオスポッド>
		 [2012/01/29]
		 ●反转：场上的怪兽全部加入持有者卡组洗切。那之后，双方玩家直到和加入各自卡组的数量相同数量的怪兽出现为止把卡组翻开，从那之中把4星以下的怪兽全部里侧守备表示特殊召唤。那以外的翻开的卡全部丢弃去墓地。
		 ◇反转效果（进入连锁）。
		 ◇强制发动。
		 ◇不取对象效果。
		 ◇效果处理时控制权在对方而持有者是自己的怪兽回到自己的卡组时，计入自己回到卡组的怪兽数量。
		 ◇「混沌壶/カオスポッド」被战斗破坏确定的场合，这张「混沌壶/カオスポッド」不会回到卡组也不会被计入回到卡组的怪兽数量。
		 ◇融合、同调、XYZ怪兽因为这个效果回到额外卡组的场合不会被计入回到卡组的怪兽数量。
		 ◇陷阱怪兽因为这个效果回到卡组的场合不会被计入回到卡组的怪兽数量。
		 ◇衍生物因为这个效果直接被消灭不会被计入回到卡组的怪兽数量。
		 ◇翻到了不能特殊召唤或5星以上的怪兽的场合将这些怪兽送去墓地。
		 ◇回到自己卡组的怪兽数量比自己场上可用的怪兽区域要多时将可以特殊召唤的怪兽全部特殊后，将剩余的满足特殊召唤条件的怪兽送去墓地。
		 ◇因为这个效果以里侧守备表示特殊召唤的怪兽在那个回合中不能将表示形式变更。

		 中文名: 混沌壶
		 卡片种类: 效果怪兽
		 卡片密码: 79106360
		 种族: 岩石
		 属性: 地
		 星级: 3
		 攻击力: 800
		 防御力: 700
		 罕见度: 平卡N，银字R，黄金GR
		 卡包: BE01、CA、DL01、KA、GS04、GDB1
		 效果: 反转：场上的怪兽全部加入持有者卡组洗切。那之后，双方玩家直到和加入各自卡组的数量相同数量的怪兽出现为止把卡组翻开，从那之中把4星以下的怪兽全部里侧守备表示特殊召唤。那以外的翻开的卡全部丢弃去墓地。

		*/
		Id:       130,
		Password: "79106360",
		Name:     "混沌壶",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 3,
		La:    ygo.LA_Earth, // 地
		Lr:    ygo.LR_Rock,  // 岩石
		Atk:   800,
		Def:   700,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 82 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 136
		 调整:

		 [破龙剑士]<バスター·ブレイダー>
		 [2010/08/28]
		 ●这张卡的攻击力，在对方场上及对方的墓地每存在1只龙族怪兽就上升500分。
		 ◇永续效果（不进入连锁）

		 中文名: 破龙剑士
		 卡片种类: 效果怪兽
		 卡片密码: 78193831
		 种族: 战士
		 属性: 地
		 星级: 7
		 攻击力: 2600
		 防御力: 2300
		 罕见度: 平卡N，面闪SR，金字UR，爆闪PR，立体UTR
		 卡包: BE01、CA、DL01、303、DT01、YAP01、DPYG、DTP01、YU、SY2
		 效果: 这张卡的攻击力上升对方场上以及对方墓地存在的龙族怪兽每1只500。

		*/
		Id:       136,
		Password: "78193831",
		Name:     "破龙剑士",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 7,
		La:    ygo.LA_Earth,   // 地
		Lr:    ygo.LR_Warrior, // 战士
		Atk:   2600,
		Def:   2300,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 83 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 15
		 调整:

		 [恶魔侦察者]<悪魔の偵察者>
		 [2010/10/01]
		 ●反转：对方抽3张卡。抽的卡双方确认，那之后有魔法卡的场合，那些魔法卡全部丢弃去墓地。
		 ◇反转效果（进入连锁）
		 ◇强制发动

		 中文名: 恶魔侦察者
		 卡片种类: 效果怪兽
		 卡片密码: 81863068
		 种族: 恶魔
		 属性: 暗
		 星级: 2
		 攻击力: 650
		 防御力: 500
		 罕见度: 平卡N
		 卡包: BE01、MR、DL01、KA
		 效果: 反转：对方从卡组抽3张卡。这个效果抽到的卡给双方确认，从那之中把魔法卡全部丢弃去墓地。

		*/
		Id:       15,
		Password: "81863068",
		Name:     "恶魔侦察者",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 2,
		La:    ygo.LA_Dark,  // 暗
		Lr:    ygo.LR_Fiend, // 恶魔
		Atk:   650,
		Def:   500,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 84 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 226
		 调整:

		 [被封印的艾克佐迪亚]<封印されしエクゾディア>
		 [09/09/20]
		 ●包括这张卡在内，
		 [被封印者的左足·右足·左腕·右腕/封印されし者の左足·右足·左腕·右腕]全在手卡的时候，决定胜利。
		 ◇无种类效果（不进入连锁）。
		 ◇不能对应这个效果发动
		 [天罚/天罰]。
		 ◇当上记卡片都集齐在手时，将它们展示。
		 ◇在对方的回合也适用。

		 中文名: 被封印的艾克佐迪亚
		 卡片种类: 效果怪兽
		 卡片密码: 33396948
		 种族: 魔法师
		 属性: 暗
		 星级: 3
		 攻击力: 1000
		 防御力: 1000
		 罕见度: 平卡N，银字R，黄金GR，金字UR，立体UTR
		 卡包: PG、BE01、PP01、DL02、307、GS01、GDB1
		 效果: 这张卡和「被封印者的右足」「被封印者的左足」「被封印者的右腕」「被封印者的左腕」在手卡全部齐集时，决斗胜利。

		*/
		Id:       226,
		Password: "33396948",
		Name:     "被封印的艾克佐迪亚",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 3,
		La:    ygo.LA_Dark,        // 暗
		Lr:    ygo.LR_Spellcaster, // 魔法师
		Atk:   1000,
		Def:   1000,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 85 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 300
		 调整:

		 [诱发暗黑长眠的路西法]<暗黒の眠りを誘うルシファー>
		 ●召唤，反转召唤的时候，选对方场上的一只怪兽。只要这张卡在场上打开表示存在，被选择的怪兽不能攻击。
		 ◇诱发效果，强制发动，取对象
		 ◇可以选择里侧表示的怪兽
		 ◇怪兽攻击宣言后到战斗步骤这段时间发动此效果的场合，攻击无效

		 中文名: 诱发暗黑长眠的路西法
		 卡片种类: 效果怪兽
		 卡片密码: 52675689
		 种族: 魔法师
		 属性: 暗
		 星级: 5
		 攻击力: 1500
		 防御力: 1800
		 罕见度: 平卡N
		 卡包: TB 、DL03、KA、SK2
		 效果: 召唤·反转召唤的时候，选对方场上的1只怪兽。只要这张卡在场上表侧表示存在，被选择的怪兽不能攻击。

		*/
		Id:       300,
		Password: "52675689",
		Name:     "诱发暗黑长眠的路西法",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 5,
		La:    ygo.LA_Dark,        // 暗
		Lr:    ygo.LR_Spellcaster, // 魔法师
		Atk:   1500,
		Def:   1800,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 86 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 303
		 调整:

		 [隼骑士]<隼の騎士>
		 ●在每一次的战斗流程可以2回攻击。参见
		 [六武众-二藏/六武衆－ニサシ]的对应内容

		 中文名: 隼骑士
		 卡片种类: 效果怪兽
		 卡片密码: 21015833
		 种族: 战士
		 属性: 地
		 星级: 3
		 攻击力: 1000
		 防御力: 700
		 罕见度: 平卡N，银字R
		 卡包: BE01、TB 、DL03、YSD05、JY
		 效果: 这张卡在同1次的战斗阶段中可以作2次攻击。

		*/
		Id:       303,
		Password: "21015833",
		Name:     "隼骑士",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 3,
		La:    ygo.LA_Earth,   // 地
		Lr:    ygo.LR_Warrior, // 战士
		Atk:   1000,
		Def:   700,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 87 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 309
		 调整:

		 [哥布林突击部队]<ゴブリン突撃部隊>
		 ●这张卡攻击的场合，战斗流程结束时改为防守表示。
		 ◇永续效果
		 ◇时点同剑斗兽回卡组换斗
		 ◇攻击被无效的场合，此效果不适用
		 ◇该时点被跳过的场合，此效果不适用★是否入连锁，调整中注：同类卡都不入连锁
		 ●直到下次自己的回合结束这张卡都不能改变表示形式。
		 ◇永续效果
		 ◇与此卡的表示形式无关
		 ◇攻击被无效的场合，此效果不适用

		 中文名: 哥布林突击部队
		 卡片种类: 效果怪兽
		 卡片密码: 78658564
		 种族: 战士
		 属性: 地
		 星级: 4
		 攻击力: 2300
		 防御力: 0
		 罕见度: 平卡N，面闪SR
		 卡包: BE01、TB 、DL03、SD05、YSD04、YSD05、JY
		 效果: 这张卡攻击的场合，战斗阶段结束时变成守备表示，直到下次的自己回合的结束阶段时不能把表示形式改变。

		*/
		Id:       309,
		Password: "78658564",
		Name:     "哥布林突击部队",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Earth,   // 地
		Lr:    ygo.LR_Warrior, // 战士
		Atk:   2300,
		Def:   0,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 88 */
	cardBag.Register(&ygo.CardOriginal{
		/*
			 id: 311
			 调整:

			 [铁骑士
			基亚·弗里德]<鉄の騎士
			ギア·フリード>
			 ●这张卡装备装备卡的时候，那张装备卡破坏。
			 ◇诱发效果，强制发动，入连锁★同一连锁中复数装备卡装备给此卡的场合，此效果进行几次连锁，调整中

			 中文名: 铁骑士 基亚·弗里德
			 卡片种类: 效果怪兽
			 卡片密码: 00423705
			 种族: 战士
			 属性: 地
			 星级: 4
			 攻击力: 1800
			 防御力: 1600
			 罕见度: 平卡N，银字R
			 卡包: BE01、TB 、DL03、SD05、JY、SJ2
			 效果: 这张卡装备装备卡的时候，那张装备卡破坏。

		*/
		Id:       311,
		Password: "00423705",
		Name:     "铁骑士 基亚·弗里德",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Earth,   // 地
		Lr:    ygo.LR_Warrior, // 战士
		Atk:   1800,
		Def:   1600,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 89 */
	cardBag.Register(&ygo.CardOriginal{
		/*
			 id: 367
			 调整:

			 [幻想召唤师]<幻想召唤师>
			 [10/10/27]
			 ●反转：把这张卡以外的1只怪兽解放，把1只融合怪兽从额外卡组特殊召唤。这个效果特殊召唤的融合怪兽在结束阶段时破坏。
			 ◇反转效果（进入连锁）
			 ◇强制发动
			 ◇把这张卡以外的1只怪兽解放是效果（不取对象）
			 ◇效果处理时选择自己额外卡组中的一只融合怪兽（不取对象）
			 ◇「虚无魔人/虚无魔人」等适用中不能发动★自己场上只有这张卡盖放时，这张卡反转的场合效果是否发动？调整中
			 ◇自己的额外卡组为0，或者自己额外卡组中只有同调怪兽存在的场合，这个效果也发动，效果处理时给对方确认自己的额外卡组『没有可以特殊召唤的融合怪兽存在』这个事实
			 ◇这个特殊召唤不属于融合召唤
			 ◇不能把自身不能融合召唤的融合怪兽（「剑斗兽
			凯撒/剣闘獣ガイザレス」「嵌合要塞龙/キメラテック·フォートレス·ドラゴン」等）用这个效果特殊召唤
			 ◇复数张这张卡同时反转的场合，可以相互解放
			 ◇『解放』『特殊召唤』同时处理
			 ◇效果处理时这张卡变成里侧表示的场合也不能把这张卡自身因这个效果而解放
			 ◇效果处理时这张卡不在场上表侧表示存在的场合，『解放』『特殊召唤』仍适用
			 ◇效果处理时没有把那只融合怪兽特殊召唤的场合，『解放』仍适用
			 ◇效果处理时没有把怪兽解放的场合，『特殊召唤』不适用
			 ◇对方场上有怪兽表侧表示存在，自己对应自己场上这张卡的反转效果的发动连锁发动「敌人操纵器/エネミーコントローラー」将这张卡解放而获得对方场上表侧表示怪兽的控制权的场合，这张卡效果处理时可以将获得控制权的那只怪兽解放发动效果
			 ●这个效果特殊召唤的融合怪兽在结束阶段时破坏。
			 ◇不进入连锁
			 ◇这个效果特殊召唤的怪兽之后变过里侧或离过场的场合不适用
			 ◇无论这张卡如何，那只怪兽回合结束阶段时仍会被这个效果所破坏
			 ◇因此效果特殊召唤「死亡魔龙/デス·デーモン·ドラゴン」的场合，「死亡魔龙/デス·デーモン·ドラゴン」不会被破坏★「王宫的号令/王宫の号令」适用的场合那只融合怪兽是否会因这个效果而破坏？调整中

			 中文名: 幻想召唤师
			 卡片种类: 效果怪兽
			 卡片密码: 14644902
			 种族: 魔法师
			 属性: 光
			 星级: 3
			 攻击力: 800
			 防御力: 900
			 罕见度: 平卡N
			 卡包: LN、DL03、GLD01、TP08、KA、PE、SK2
			 效果: 反转：这张卡以外的1只怪兽做祭品。1只融合怪兽从融合卡组特殊召唤。此融合怪兽在回合结束时破坏。

		*/
		Id:       367,
		Password: "14644902",
		Name:     "幻想召唤师",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 3,
		La:    ygo.LA_Light,       // 光
		Lr:    ygo.LR_Spellcaster, // 魔法师
		Atk:   800,
		Def:   900,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 90 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 375
		 调整:

		 [骏足之迅猛龙]<俊足のギラザウルス>
		 ●这张卡可以正常召唤也可以通过特殊召唤来使用。
		 ◇召唤规则
		 ◇可以通常召唤
		 ◇是否能以这个方式特殊召唤与能否发动下面的诱发效果无关
		 ●通过特殊召唤使用的场合，对方可以在对方的墓地选一只怪兽特殊召唤上场。
		 ◇诱发效果，任意发动，取对象
		 ◇发动与否的决定权在对手
		 ◇对手墓地里没有可以特殊召唤的怪兽的场合，此效果不能发动

		 中文名: 骏足之迅猛龙
		 卡片种类: 效果怪兽
		 卡片密码: 45894482
		 种族: 恐龙
		 属性: 地
		 星级: 3
		 攻击力: 1400
		 防御力: 400
		 罕见度: 平卡N
		 卡包: BE02、LN、DL03、SD09、KA
		 效果: 这张卡可以通常召唤也可以通过特殊召唤来使用。通过特殊召唤使用的场合，对方可以在对方的墓地选1只怪兽特殊召唤上场。

		*/
		Id:       375,
		Password: "45894482",
		Name:     "骏足之迅猛龙",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 3,
		La:    ygo.LA_Earth,    // 地
		Lr:    ygo.LR_Dinosaur, // 恐龙
		Atk:   1400,
		Def:   400,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 91 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 412
		 调整:

		 [时间魔术师]<時の魔術師>
		 [09/09/20]

		 ●1回合1次，自己的主要阶段时才能发动。进行1次投掷硬币，对里表作猜测。
		 ◇起动效果（进入连锁）。
		 ◇不能对应这个效果发动
		 [地狱的冷枪/地狱の扉越し铳]。
		 ◇可以连锁这个效果发动
		 [星尘龙/スターダスト·ドラゴン]
		 [我身作盾/我が身を盾に]的效果。
		 ●猜中的场合，对方场上存在的怪兽全部破坏。
		 ◇不取对象效果。
		 ●猜错的场合，自己场上存在的怪兽全部破坏，自己受到破坏怪兽的攻击力合计数值一半的伤害。
		 ◇『破坏』『伤害』同时处理。

		 中文名: 时间魔术师
		 卡片种类: 效果怪兽
		 卡片密码: 71625222
		 种族: 魔法师
		 属性: 光
		 星级: 2
		 攻击力: 500
		 防御力: 400
		 罕见度: 平卡N，面闪SR，银碎SER
		 卡包: ME、BE02、DL04、PP01、DT05、JY、SJ2
		 效果: 掷1个硬币猜正反。猜中则把对方场上的怪兽全部破坏。猜错则把自己场上的怪兽全部破坏。之后受到由于这个效果被破坏的自己怪兽全部攻击力合计一半伤害。这个效果1回合1次在自己的主要阶段可以使用。

		*/
		Id:       412,
		Password: "71625222",
		Name:     "时间魔术师",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 2,
		La:    ygo.LA_Light,       // 光
		Lr:    ygo.LR_Spellcaster, // 魔法师
		Atk:   500,
		Def:   400,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 92 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 531
		 调整:

		 中文名: 疾风之暗黑骑士 盖亚
		 卡片种类: 效果怪兽
		 卡片密码: 16589042
		 种族: 战士
		 属性: 暗
		 星级: 7
		 攻击力: 2300
		 防御力: 2100
		 罕见度: 平卡N，银字R，金字UR
		 卡包: BE02、DL04、SD05、DT02、DTP01、YU
		 效果: 自己的手卡只有这1张卡的场合，这张卡可以不用解放来召唤。

		*/
		Id:       531,
		Password: "16589042",
		Name:     "疾风之暗黑骑士 盖亚",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 7,
		La:    ygo.LA_Dark,    // 暗
		Lr:    ygo.LR_Warrior, // 战士
		Atk:   2300,
		Def:   2100,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 93 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 532
		 调整:

		 中文名: 翻弄敌人的精灵剑士
		 卡片种类: 效果怪兽
		 卡片密码: 52077741
		 种族: 战士
		 属性: 地
		 星级: 4
		 攻击力: 1400
		 防御力: 1200
		 罕见度: 平卡N，银字R，金字UR
		 卡包: BE02、DL04、SD05、YU、SY2
		 效果: 这张卡不会被攻击力1900以上的怪兽战斗破坏。（伤害计算适用）

		*/
		Id:       532,
		Password: "52077741",
		Name:     "翻弄敌人的精灵剑士",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Earth,   // 地
		Lr:    ygo.LR_Warrior, // 战士
		Atk:   1400,
		Def:   1200,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 94 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 533
		 调整:

		 中文名: 太阳之战士
		 卡片种类: 效果怪兽
		 卡片密码: 57482479
		 种族: 战士
		 属性: 光
		 星级: 5
		 攻击力: 2100
		 防御力: 1400
		 罕见度: 平卡N，银字R，金字UR
		 卡包: BE02、DL04、JY
		 效果: 和暗属性的怪兽战斗的场合，伤害计算阶段这张卡的攻击力上升500。

		*/
		Id:       533,
		Password: "57482479",
		Name:     "太阳之战士",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 5,
		La:    ygo.LA_Light,   // 光
		Lr:    ygo.LR_Warrior, // 战士
		Atk:   2100,
		Def:   1400,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 95 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 534
		 调整:

		 中文名: 指挥骑士
		 卡片种类: 效果怪兽
		 卡片密码: 10375182
		 种族: 战士
		 属性: 炎
		 星级: 4
		 攻击力: 1200
		 防御力: 1900
		 罕见度: 平卡N，面闪SR，金字UR
		 卡包: BE02、DL04、SD05、JY
		 效果: 只要自己场上有这张卡以外的怪兽存在，对方不能选择这张卡为攻击对象。只要这张卡在场上存在，自己的战士族怪兽攻击力上升400。

		*/
		Id:       534,
		Password: "10375182",
		Name:     "指挥骑士",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Fire,    // 炎
		Lr:    ygo.LR_Warrior, // 战士
		Atk:   1200,
		Def:   1900,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 96 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 535
		 调整:

		 中文名: 帝王海马
		 卡片种类: 效果怪兽
		 卡片密码: 17444133
		 种族: 海龙
		 属性: 光
		 星级: 4
		 攻击力: 1700
		 防御力: 1650
		 罕见度: 平卡N，银字R，面闪SR
		 卡包: BE02、DL04、SD13、DPKB、KA、SK2、SD25
		 效果: 光属性的怪兽祭品召唤的场合，这只怪兽1只可以算作2只份量的祭品。

		*/
		Id:       535,
		Password: "17444133",
		Name:     "帝王海马",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Light,      // 光
		Lr:    ygo.LR_SeaSerpent, // 海龙
		Atk:   1700,
		Def:   1650,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 97 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 536
		 调整:

		 中文名: 吸血鬼领主
		 卡片种类: 效果怪兽
		 卡片密码: 53839837
		 种族: 不死
		 属性: 暗
		 星级: 5
		 攻击力: 2000
		 防御力: 1500
		 罕见度: 平卡N，银字R，金字UR，爆闪PR
		 卡包: BE02、DL04、SD02、DPKB、KA
		 效果: 这张卡每次给予对方战斗伤害，卡的种类（怪兽·魔法·陷阱）宣言。对方在卡组选择1张宣言的种类的卡送去墓地。这张卡被对方的卡的效果破坏送去墓地时，这张卡在下1次自己的准备阶段特殊召唤上场。

		*/
		Id:       536,
		Password: "53839837",
		Name:     "吸血鬼领主",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 5,
		La:    ygo.LA_Dark,   // 暗
		Lr:    ygo.LR_Zombie, // 不死
		Atk:   2000,
		Def:   1500,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 98 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 537
		 调整:

		 中文名: 卡通哥布林突击部队
		 卡片种类: 效果怪兽
		 卡片密码: 15270885
		 种族: 战士
		 属性: 地
		 星级: 4
		 攻击力: 2300
		 防御力: 0
		 罕见度: 平卡N，银字R
		 卡包: BE02、DL04、PE
		 效果: 效果·卡通：这张卡召唤·反转召唤·特殊召唤的回合不能攻击。场上的「卡通世界」被破坏时这张卡也破坏。自己场上有「卡通世界」且对方不控制卡通的场合，这张卡可以直接攻击对方玩家。这张卡攻击的场合在战斗阶段结束时守备表示，在下次的自己的回合结束前不能改变这张卡的表示形式。

		*/
		Id:       537,
		Password: "15270885",
		Name:     "卡通哥布林突击部队",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Earth,   // 地
		Lr:    ygo.LR_Warrior, // 战士
		Atk:   2300,
		Def:   0,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 99 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 538
		 调整:

		 中文名: 卡通加农炮兵
		 卡片种类: 效果怪兽
		 卡片密码: 79875176
		 种族: 机械
		 属性: 暗
		 星级: 4
		 攻击力: 1400
		 防御力: 1300
		 罕见度: 平卡N，银字R
		 卡包: BE02、DL04、PE
		 效果: 效果·卡通：这张卡召唤·反转召唤·特殊召唤的回合不能攻击。场上的「卡通世界」被破坏时这张卡也破坏。自己场上有「卡通世界」且对方不控制卡通的场合，这张卡可以直接攻击对方玩家。每祭品1只自己场上存在的怪兽，对方受到500分的基本分伤害。

		*/
		Id:       538,
		Password: "79875176",
		Name:     "卡通加农炮兵",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Dark,    // 暗
		Lr:    ygo.LR_Machine, // 机械
		Atk:   1400,
		Def:   1300,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 100 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 539
		 调整:

		 中文名: 卡通双生精灵
		 卡片种类: 效果怪兽
		 卡片密码: 42386471
		 种族: 魔法师
		 属性: 地
		 星级: 4
		 攻击力: 1900
		 防御力: 900
		 罕见度: 平卡N，银字R
		 卡包: BE02、DL04、PE
		 效果: 效果·卡通：这张卡召唤·反转召唤·特殊召唤的回合不能攻击。场上的「卡通世界」被破坏时这张卡也破坏。自己场上有「卡通世界」且对方不控制卡通的场合，这张卡可以直接攻击对方玩家。这张卡给予对方玩家战斗伤害时，对方随机丢弃1张手卡。

		*/
		Id:       539,
		Password: "42386471",
		Name:     "卡通双生精灵",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Earth,       // 地
		Lr:    ygo.LR_Spellcaster, // 魔法师
		Atk:   1900,
		Def:   900,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 101 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 540
		 调整:

		 中文名: 卡通假面魔道士
		 卡片种类: 效果怪兽
		 卡片密码: 16392422
		 种族: 魔法师
		 属性: 暗
		 星级: 4
		 攻击力: 900
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: BE02、DL04、PE
		 效果: 效果·卡通：这张卡召唤·反转召唤·特殊召唤的回合不能攻击。场上的「卡通世界」被破坏时这张卡也破坏。自己场上有「卡通世界」且对方不控制卡通的场合，这张卡可以直接攻击对方玩家。这张卡造成对方伤害时，这张卡的持有者抽1张卡。

		*/
		Id:       540,
		Password: "16392422",
		Name:     "卡通假面魔道士",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Dark,        // 暗
		Lr:    ygo.LR_Spellcaster, // 魔法师
		Atk:   900,
		Def:   1400,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 102 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 56
		 调整:

		 [卡通蛋龙]<トゥーン·ドラゴン·エッガー>
		 ●自己场上没有
		 [卡通世界/トゥーン·ワールド]不能召唤。
		 ●召唤的回合不能攻击。
		 ●不付出500分不能攻击。
		 ●
		 [卡通世界/トゥーン·ワールド]被破坏时这张卡破坏。
		 ●对方可以控制的怪兽没有卡通怪兽的场合，这张卡可以直接攻击对方。
		 ●对方可控制的怪兽有卡通怪兽的场合，必须要选择对方的卡通怪兽攻击。参见
		 [青眼卡通白龙/ブルーアイズ·トゥーン·ドラゴン]的对应内容

		 中文名: 卡通蛋龙
		 卡片种类: 效果怪兽
		 卡片密码: 38369349
		 种族: 龙
		 属性: 炎
		 星级: 7
		 攻击力: 2200
		 防御力: 2600
		 罕见度: 平卡N
		 卡包: BE01、PS、DL01、PE
		 效果: 效果·卡通：这张卡不能通常召唤。自己场上有「卡通世界」存在的场合才能特殊召唤（5星以上需要解放）。这张卡在特殊召唤的回合不能攻击。这张卡若不支付500基本分则不能攻击宣言。对方场上不存在卡通怪兽的场合，这张卡可以直接攻击对方玩家。存在的场合，必须选择卡通怪兽作为攻击对象。场上的「卡通世界」被破坏时，这张卡破坏。

		*/
		Id:       56,
		Password: "38369349",
		Name:     "卡通蛋龙",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 7,
		La:    ygo.LA_Fire,   // 炎
		Lr:    ygo.LR_Dragon, // 龙
		Atk:   2200,
		Def:   2600,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 103 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 57
		 调整:

		 [卡通美人鱼]<トゥーン·マーメイド>
		 ●自己场上没有
		 [卡通世界/トゥーン·ワールド]不能召唤。
		 ●召唤的回合不能攻击。
		 ●不付出500分不能攻击。
		 ●
		 [卡通世界/トゥーン·ワールド]被破坏时这张卡破坏。
		 ●对方可以控制的怪兽没有卡通怪兽的场合，这张卡可以直接攻击对方。
		 ●对方可控制的怪兽有卡通怪兽的场合，必须要选择对方的卡通怪兽攻击。参见
		 [青眼卡通白龙/ブルーアイズ·トゥーン·ドラゴン]的对应内容

		 中文名: 卡通人鱼
		 卡片种类: 效果怪兽
		 卡片密码: 65458948
		 种族: 水
		 属性: 水
		 星级: 4
		 攻击力: 1400
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: BE01、PS、DL01、PE
		 效果: 效果·卡通：这张卡不能通常召唤。自己场上有「卡通世界」存在的场合才能特殊召唤（5星以上需要解放）。这张卡在特殊召唤的回合不能攻击。这张卡若不支付500基本分则不能攻击宣言。对方场上不存在卡通怪兽的场合，这张卡可以直接攻击对方玩家。存在的场合，必须选择卡通怪兽作为攻击对象。场上的「卡通世界」被破坏时，这张卡破坏。

		*/
		Id:       57,
		Password: "65458948",
		Name:     "卡通人鱼",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Water, // 水
		Lr:    ygo.LR_Water, // 水
		Atk:   1400,
		Def:   1500,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 104 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 58
		 调整:

		 [卡通恶魔]<トゥーン·デーモン>
		 ●自己场上没有
		 [卡通世界/トゥーン·ワールド]不能召唤。
		 ●召唤的回合不能攻击。
		 ●不付出500分不能攻击。
		 ●
		 [卡通世界/トゥーン·ワールド]被破坏时这张卡破坏。
		 ●对方可以控制的怪兽没有卡通怪兽的场合，这张卡可以直接攻击对方。
		 ●对方可控制的怪兽有卡通怪兽的场合，必须要选择对方的卡通怪兽攻击。参见
		 [青眼卡通白龙/ブルーアイズ·トゥーン·ドラゴン]的对应内容

		 中文名: 卡通恶魔
		 卡片种类: 效果怪兽
		 卡片密码: 91842653
		 种族: 恶魔
		 属性: 暗
		 星级: 6
		 攻击力: 2500
		 防御力: 1200
		 罕见度: 银字R，面闪SR
		 卡包: BE01、PS、DL01、PE
		 效果: 效果·卡通：这张卡不能通常召唤。自己场上有「卡通世界」存在的场合才能特殊召唤（5星以上需要解放）。这张卡在特殊召唤的回合不能攻击。这张卡若不支付500基本分则不能攻击宣言。对方场上不存在卡通怪兽的场合，这张卡可以直接攻击对方玩家。存在的场合，必须选择卡通怪兽作为攻击对象。场上的「卡通世界」被破坏时，这张卡破坏。

		*/
		Id:       58,
		Password: "91842653",
		Name:     "卡通恶魔",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 6,
		La:    ygo.LA_Dark,  // 暗
		Lr:    ygo.LR_Fiend, // 恶魔
		Atk:   2500,
		Def:   1200,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 105 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 59
		 调整:

		 [定时炸弹]<タイム·ボマー>
		 ●翻转效果：自己的准备流程的这张卡做祭品。全部自己的怪兽破坏，之后被破坏的怪兽总攻击力的一半的基本分数值造成对方的伤害。
		 ◇确定发动是反转效果
		 ◇暂时离场会造成“反转过”这一事实重置★此效果是否入连锁，调整中★此效果适用后
		 [王宫的号令/王宮の号令]效果适用的场合，下面的诱发效果是否发动，调整中★此卡在准备阶段因其他卡的效果反转的场合，假设反转效果入连锁，下面的诱发效果是随着反转效果一起处理还是新开连锁，调整中
		 ◇确定后发动时是诱发效果，强制发动，“自己的准备流程的把这张卡解放”是COST★反转后的这张卡转移控制权后，在准备阶段解放这张卡发动效果的场合，怎么处理，调整中
		 ◇攻击力在场上判定★破坏里侧怪兽的场合，里侧怪兽的攻击力按0算还是按原本攻击力算，没说

		 中文名: 定时炸弹
		 卡片种类: 效果怪兽
		 卡片密码: 90020065
		 种族: 炎
		 属性: 炎
		 星级: 2
		 攻击力: 200
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: PS、DL01、PE
		 效果: 反转：自己的准备阶段把这张卡作为祭品。全部自己的怪兽破坏，给予对方那个总攻击力一半数值的伤害。

		*/
		Id:       59,
		Password: "90020065",
		Name:     "定时炸弹",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 2,
		La:    ygo.LA_Fire, // 炎
		Lr:    ygo.LR_Fire, // 炎
		Atk:   200,
		Def:   1000,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 106 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 62
		 调整:

		 中文名: 电子壶
		 卡片种类: 效果怪兽
		 卡片密码: 34124316
		 种族: 岩石
		 属性: 暗
		 星级: 3
		 攻击力: 900
		 防御力: 900
		 罕见度: 平卡N，银字R，立体UTR
		 卡包: BE01、PS、DL01、DPKB、YU
		 效果: 反转：场上存在的怪兽全部破坏。那之后，双方从卡组上面把5张卡翻开，把那之中的4星以下的怪兽全部表侧攻击表示或者里侧守备表示特殊召唤。那以外的卡全部加入手卡。

		*/
		Id:       62,
		Password: "34124316",
		Name:     "电子壶",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 3,
		La:    ygo.LA_Dark, // 暗
		Lr:    ygo.LR_Rock, // 岩石
		Atk:   900,
		Def:   900,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 107 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 64
		 调整:

		 [巨大老鼠]<巨大ネズミ>
		 ●这张卡被战斗破坏送去墓地的时候，可以从卡组选一只攻击力1500以下的地属性怪兽打开攻击表示特殊召唤到自己场上。之后卡组洗切。
		 ◇诱发效果，任意发动
		 ◇不能选攻击力为？的怪兽
		 ◇时点在“战破确定的怪物送入墓地”注：此卡是遗言众之一

		 中文名: 巨大老鼠
		 卡片种类: 效果怪兽
		 卡片密码: 97017120
		 种族: 兽
		 属性: 地
		 星级: 4
		 攻击力: 1400
		 防御力: 1450
		 罕见度: 平卡N，银字R
		 卡包: BE01、PS、YSD02、DL01、SD07、YU、SY2、YSD06、ST12、DS14
		 效果: 这张卡被战斗破坏送去墓地时，可以从自己卡组把1只攻击力1500以下的地属性怪兽在自己场上表侧攻击表示特殊召唤。

		*/
		Id:       64,
		Password: "97017120",
		Name:     "巨大老鼠",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Earth, // 地
		Lr:    ygo.LR_Beast, // 兽
		Atk:   1400,
		Def:   1450,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 108 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 65
		 调整:

		 [千手神]<センジュ·ゴッド>
		 ●这张卡召唤，反转召唤时，从卡组选择一张仪式怪兽卡加入手卡。参见
		 [万手神/マンジュ·ゴッド]的对应内容

		 中文名: 千手神
		 卡片种类: 效果怪兽
		 卡片密码: 23401839
		 种族: 天使
		 属性: 光
		 星级: 4
		 攻击力: 1400
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: BE01、PS、DL01、PE、DT10
		 效果: 这张卡召唤·反转召唤成功时，可以从自己卡组把1只仪式怪兽加入手卡。

		*/
		Id:       65,
		Password: "23401839",
		Name:     "千手神",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Light, // 光
		Lr:    ygo.LR_Fairy, // 天使
		Atk:   1400,
		Def:   1000,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 109 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 68
		 调整:

		 [空手道家]<カラテマン>
		 ●一回合一次这张卡的原本攻击力变成2倍。
		 ◇起动效果
		 ◇给此效果适用的此卡装备
		 [巨大化/巨大化]的场合，攻击力按
		 [巨大化/巨大化]的效果算
		 ◇此效果适用后，
		 [技能吸收/スキルドレイン]效果适用的场合，攻击力变回1000
		 ●这个效果使用的场合，在回合结束时这张卡破坏。
		 ◇此时此卡不在场上表侧表示存在的场合，此效果不适用★
		 [技能吸收/スキルドレイン]效果适用的场合，是否破坏，调整中

		 中文名: 空手道家
		 卡片种类: 效果怪兽
		 卡片密码: 23289281
		 种族: 战士
		 属性: 地
		 星级: 3
		 攻击力: 1000
		 防御力: 1000
		 罕见度: 平卡N，银字R
		 卡包: BE01、PS、DL01、JY
		 效果: 1回合1次，自己的主要阶段时才能发动。这张卡的攻击力直到结束阶段时变成原本攻击力2倍的数值。这个效果使用的场合，这张卡在结束阶段时破坏。

		*/
		Id:       68,
		Password: "23289281",
		Name:     "空手道家",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 3,
		La:    ygo.LA_Earth,   // 地
		Lr:    ygo.LR_Warrior, // 战士
		Atk:   1000,
		Def:   1000,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 110 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 78
		 调整:

		 [音速鸟]<ソニックバード>
		 ●这张卡召唤，反转召唤的时候，从卡组选一张仪式魔法卡加入手卡。参见
		 [万手神/マンジュ·ゴッド]的对应内容

		 中文名: 音速鸟
		 卡片种类: 效果怪兽
		 卡片密码: 57617178
		 种族: 鸟兽
		 属性: 风
		 星级: 4
		 攻击力: 1400
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: BE01、PS、DL01、GLD01、PE、DT10
		 效果: 这张卡召唤·反转召唤成功时，可以从自己卡组把1张仪式魔法卡加入手卡。

		*/
		Id:       78,
		Password: "57617178",
		Name:     "音速鸟",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Wind,        // 风
		Lr:    ygo.LR_WingedBeast, // 鸟兽
		Atk:   1400,
		Def:   1000,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 111 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 79
		 调整:

		 [杀人番茄]<キラー·トマト>
		 ●这张卡被战斗破坏送去墓地的时候，可以从卡组选一只攻击力1500以下的暗属性怪兽打开攻击表示特殊召唤到自己场上。之后卡组洗切。参见
		 [巨大老鼠/巨大ネズミ]的对应内容

		 中文名: 杀人番茄
		 卡片种类: 效果怪兽
		 卡片密码: 83011277
		 种族: 植物
		 属性: 暗
		 星级: 4
		 攻击力: 1400
		 防御力: 1100
		 罕见度: 平卡N，银字R
		 卡包: BE01、PS、DL01、SD12、YSD03、PE、DT11
		 效果: 这张卡被战斗破坏送去墓地时，可以从自己卡组把1只攻击力1500以下的暗属性怪兽在自己场上表侧攻击表示特殊召唤。

		*/
		Id:       79,
		Password: "83011277",
		Name:     "杀人番茄",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 4,
		La:    ygo.LA_Dark,  // 暗
		Lr:    ygo.LR_Plant, // 植物
		Atk:   1400,
		Def:   1100,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 112 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 88
		 调整:

		 [青眼卡通白龙]<ブルーアイズ·トゥーン·ドラゴン>
		 ●自己场上没有
		 [卡通世界/トゥーン·ワールド]不能召唤。
		 ◇召唤规则
		 ◇不能通常召唤
		 ◇遵守苏生限制
		 ◇从手牌里特殊召唤的场合需要解放同等星数怪兽上级召唤所需要解放的怪兽数目那么多只怪兽
		 ◇正规出场后的此卡苏生的场合，需要自己场上有
		 [卡通世界/トゥーン·ワールド]
		 ●召唤的回合不能攻击。
		 ◇永续效果
		 ●不付出500分不能攻击。
		 ◇永续效果
		 ◇攻击宣言的COST
		 ●
		 [卡通世界/トゥーン·ワールド]被破坏时这张卡破坏。
		 ◇永续效果
		 ◇无论被破坏的
		 [卡通世界/トゥーン·ワールド]的控制者是谁，此卡都要破坏
		 ◇无论被破坏的
		 [卡通世界/トゥーン·ワールド]是表侧还是里侧，此卡都要破坏
		 ◇
		 [卡通世界/トゥーン·ワールド]因破坏以外的效果离场的场合，此卡不用破坏
		 ●对方可以控制的怪兽没有卡通怪兽的场合，这张卡可以直接攻击对方。
		 ◇永续效果
		 ◇可以攻击对方控制的其他怪兽
		 ●对方可控制的怪兽有卡通怪兽的场合，必须要选择对方的卡通怪兽攻击。
		 ◇永续效果
		 ◇并不意味着此时不进行战斗就不能结束战斗阶段

		 中文名: 青眼卡通龙
		 卡片种类: 效果怪兽
		 卡片密码: 53183600
		 种族: 龙
		 属性: 光
		 星级: 8
		 攻击力: 3000
		 防御力: 2500
		 罕见度: 面闪SR，金字UR，爆闪PR
		 卡包: BE01、PS、DL01、PE
		 效果: 效果·卡通：这张卡不能通常召唤。自己场上有「卡通世界」存在的场合才能特殊召唤（5星以上需要解放）。这张卡在特殊召唤的回合不能攻击。这张卡若不支付500基本分则不能攻击宣言。对方场上不存在卡通怪兽的场合，这张卡可以直接攻击对方玩家。存在的场合，必须选择卡通怪兽作为攻击对象。场上的「卡通世界」被破坏时，这张卡破坏。

		*/
		Id:       88,
		Password: "53183600",
		Name:     "青眼卡通龙",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 8,
		La:    ygo.LA_Light,  // 光
		Lr:    ygo.LR_Dragon, // 龙
		Atk:   3000,
		Def:   2500,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 113 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 91
		 调整:

		 [人造人-念力震慑者]<人造人間－サイコ·ショッカー>
		 ●只要这张卡在场上打开表示存在，陷阱卡不能发动，全场的陷阱卡的效果无效。
		 ◇永续效果
		 ◇不会无效在墓地里发动的陷阱（例：
		 [诅咒之棺/呪われた棺]）
		 ◇不会无效已经适用的非留场的陷阱的效果
		 ◇连锁陷阱卡的发动用
		 [活死人的呼声/リビングデッドの呼び声]特殊召唤此卡的场合，那张陷阱的效果无效，发动不无效
		 ◇用
		 [活死人的呼声/リビングデッドの呼び声]特殊召唤此卡并且此卡效果适用的场合，此卡破坏则
		 [活死人的呼声/リビングデッドの呼び声]破坏，
		 [活死人的呼声/リビングデッドの呼び声]离场此卡不破坏

		 中文名: 人造人-念力震慑者
		 卡片种类: 效果怪兽
		 卡片密码: 77585513
		 种族: 机械
		 属性: 暗
		 星级: 6
		 攻击力: 2400
		 防御力: 1500
		 罕见度: 平卡N，黄金GR，金字UR，爆闪PR，立体UTR
		 卡包: BE01、CA、DL01、308、GLD01、YAP01、GS01、JY、SJ2、GDB1
		 效果: 只要这张卡在场上表侧表示存在，双方不能把陷阱卡发动，场上的陷阱卡的效果无效化。

		*/
		Id:       91,
		Password: "77585513",
		Name:     "人造人-念力震慑者",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 6,
		La:    ygo.LA_Dark,    // 暗
		Lr:    ygo.LR_Machine, // 机械
		Atk:   2400,
		Def:   1500,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 114 */
	cardBag.Register(&ygo.CardOriginal{
		/*
			 id: 948
			 调整:

			 [暗黑魔族
			基尔法恶魔]<暗黒魔族ギルファー·デーモン>
			 ●此卡被送进墓地时可以发动效果。此卡可以变成一张攻击力下降500点的装备卡，装备在场上一只怪物身上。
			 ◇诱发效果（进入连锁）。
			 ◇任意发动。
			 ◇取对象效果。

			 中文名: 暗黑魔族 基尔法恶魔
			 卡片种类: 效果怪兽
			 卡片密码: 50287060
			 种族: 恶魔
			 属性: 暗
			 星级: 6
			 攻击力: 2200
			 防御力: 2500
			 罕见度: 平卡N，面闪SR，金字UR，立体UTR
			 卡包: WJ、305、DT06、YU
			 效果: 这张卡被送去墓地时可以发动效果。这张卡可以变成1张攻击力下降500点的装备卡，装备在场上1只怪兽身上。

		*/
		Id:       948,
		Password: "50287060",
		Name:     "暗黑魔族 基尔法恶魔",
		Lc:       ygo.LC_MonsterEffect, // 效果怪兽

		Level: 6,
		La:    ygo.LA_Dark,  // 暗
		Lr:    ygo.LR_Fiend, // 恶魔
		Atk:   2200,
		Def:   2500,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 115 */
	cardBag.Register(&ygo.CardOriginal{
		/*
			 id: 1121
			 调整:

			 [有翼幻兽
			奇美拉]<有翼幻獣キマイラ>
			 [2011/04/28]
			 ●这张卡被破坏时，可以选墓地的一张「巴风特」或「幻兽王
			卡杰」特殊召唤上场。
			 ◇诱发效果（进入连锁）。
			 ◇任意发动。
			 ◇取对象效果。

			 中文名: 有翼幻兽 奇美拉
			 卡片种类: 融合怪兽
			 卡片密码: 04796100
			 种族: 兽
			 属性: 风
			 星级: 6
			 攻击力: 2100
			 防御力: 1800
			 罕见度: 平卡N，银字R
			 卡包: ABPF(607)、YU、其他
			 效果: 融合：「幻兽王 卡杰」＋「巴风特」。这张卡被破坏时，可以选墓地的1张「巴风特」或「幻兽王 卡杰」特殊召唤上场。

		*/
		Id:       1121,
		Password: "04796100",
		Name:     "有翼幻兽 奇美拉",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 6,
		La:    ygo.LA_Wind,  // 风
		Lr:    ygo.LR_Beast, // 兽
		Atk:   2100,
		Def:   1800,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 116 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1182
		 调整:

		 [乘龙的飞龙战士]<ドラゴンに兼るワイバーン>
		 [2011/03/01]
		 ●对方场上的怪兽只有地，水，炎属性的时候，可以直接攻击对方玩家。
		 ◇永续效果（不进入连锁）。
		 ◇对方场上有里侧表示怪兽存在的场合，这个效果适用

		 中文名: 乘龙的翼龙战士
		 卡片种类: 融合怪兽
		 卡片密码: 03366982
		 种族: 龙
		 属性: 风
		 星级: 5
		 攻击力: 1700
		 防御力: 1500
		 罕见度: 平卡N，金字UR
		 卡包: TP07、DT08、JY、其他
		 效果: 融合：「宝贝龙」＋「翼龙战士」。对方场上的怪兽只有地，水，炎的时候，可以直接攻击对方玩家。

		*/
		Id:       1182,
		Password: "03366982",
		Name:     "乘龙的翼龙战士",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 5,
		La:    ygo.LA_Wind,   // 风
		Lr:    ygo.LR_Dragon, // 龙
		Atk:   1700,
		Def:   1500,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 117 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1262
		 调整:

		 中文名: 青眼究极龙
		 卡片种类: 融合怪兽
		 卡片密码: 23995346
		 种族: 龙
		 属性: 光
		 星级: 12
		 攻击力: 4500
		 防御力: 3800
		 罕见度: 黄金GR，面闪SR，金字UR，斜碎SCR，立体UTR
		 卡包: PP03、GLD01、DPKB、KA
		 效果: 融合：「青眼白龙」＋「青眼白龙」＋「青眼白龙」。

		*/
		Id:       1262,
		Password: "23995346",
		Name:     "青眼究极龙",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 12,
		La:    ygo.LA_Light,  // 光
		Lr:    ygo.LR_Dragon, // 龙
		Atk:   4500,
		Def:   3800,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterMonsterFusion("青眼白龙", "青眼白龙", "青眼白龙")
			return true
		}, // 初始
		IsValid: true,
	})

	/* 118 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 141
		 调整:

		 中文名: 炎之剑士
		 卡片种类: 融合怪兽
		 卡片密码: 45231177
		 种族: 战士
		 属性: 炎
		 星级: 5
		 攻击力: 1800
		 防御力: 1600
		 罕见度: 银字R，金碎USR
		 卡包: LB、BE01、DL02、Starter Box
		 效果: 融合：「火焰操纵者」＋「传说的剑豪 正树」。

		*/
		Id:       141,
		Password: "45231177",
		Name:     "炎之剑士",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 5,
		La:    ygo.LA_Fire,    // 炎
		Lr:    ygo.LR_Warrior, // 战士
		Atk:   1800,
		Def:   1600,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterMonsterFusion("火焰操纵者", "传说的剑豪 正树")
			return true
		}, // 初始
		IsValid: true,
	})

	/* 119 */
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
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterMonsterFusion("银牙狼", "魔界之棘")
			return true
		}, // 初始
		IsValid: true,
	})

	/* 120 */
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
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterMonsterFusion("耳天使", "大雷电球")
			return true
		}, // 初始
		IsValid: true,
	})

	/* 121 */
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
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterMonsterFusion("女王的影武者", "响女")
			return true
		}, // 初始
		IsValid: true,
	})

	/* 122 */
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
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterMonsterFusion("苍翼冠鸟", "雏鸡")
			return true
		}, // 初始
		IsValid: true,
	})

	/* 123 */
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
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterMonsterFusion("冥界的番人", "王座守护者")
			return true
		}, // 初始
		IsValid: true,
	})

	/* 124 */
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
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterMonsterFusion("高科技狼", "加农炮兵")
			return true
		}, // 初始
		IsValid: true,
	})

	/* 125 */
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
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterMonsterFusion("音女", "斩首的美女")
			return true
		}, // 初始
		IsValid: true,
	})

	/* 126 */
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
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterMonsterFusion("双头恐龙王", "贪尸龙")
			return true
		}, // 初始
		IsValid: true,
	})

	/* 127 */
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
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterMonsterFusion("美杜莎的亡灵", "龙僵尸")
			return true
		}, // 初始
		IsValid: true,
	})

	/* 128 */
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
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterMonsterFusion("神灯魔人", "来自异次元的侵略者")
			return true
		}, // 初始
		IsValid: true,
	})

	/* 129 */
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
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterMonsterFusion("守城的翼龙", "妖精龙")
			return true
		}, // 初始
		IsValid: true,
	})

	/* 130 */
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
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterMonsterFusion("圣鸟", "天空猎手")
			return true
		}, // 初始
		IsValid: true,
	})

	/* 131 */
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
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterMonsterFusion("岩石巨兵", "古代精灵")
			return true
		}, // 初始
		IsValid: true,
	})

	/* 132 */
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
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterMonsterFusion("美杜莎的亡灵", "暗黑之龙王")
			return true
		}, // 初始
		IsValid: true,
	})

	/* 133 */
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
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterMonsterFusion("锹甲阿尔法", "大力独角仙")
			return true
		}, // 初始
		IsValid: true,
	})

	/* 134 */
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
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterMonsterFusion("恶魔的智慧", "魔天老")
			return true
		}, // 初始
		IsValid: true,
	})

	/* 135 */
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
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterMonsterFusion("杀人小丑", "梦幻小丑")
			return true
		}, // 初始
		IsValid: true,
	})

	/* 136 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2643
		 调整:

		 中文名: 火焰幽灵
		 卡片种类: 融合怪兽
		 卡片密码: 58528964
		 种族: 不死
		 属性: 暗
		 星级: 3
		 攻击力: 1000
		 防御力: 800
		 罕见度: 平卡N
		 卡包: LB、Starter Box、TP08
		 效果: 融合：「白骨」＋「岩浆人」。

		*/
		Id:       2643,
		Password: "58528964",
		Name:     "火焰幽灵",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 3,
		La:    ygo.LA_Dark,   // 暗
		Lr:    ygo.LR_Zombie, // 不死
		Atk:   1000,
		Def:   800,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterMonsterFusion("白骨", "岩浆人")
			return true
		}, // 初始
		IsValid: true,
	})

	/* 137 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2644
		 调整:

		 中文名: 暗黑火炎龙
		 卡片种类: 融合怪兽
		 卡片密码: 17881964
		 种族: 龙
		 属性: 暗
		 星级: 4
		 攻击力: 1500
		 防御力: 1250
		 罕见度: 平卡N
		 卡包: LB、Starter Box、TP01
		 效果: 融合：「火炎草」＋「小龙」。

		*/
		Id:       2644,
		Password: "17881964",
		Name:     "暗黑火炎龙",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 4,
		La:    ygo.LA_Dark,   // 暗
		Lr:    ygo.LR_Dragon, // 龙
		Atk:   1500,
		Def:   1250,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterMonsterFusion("火炎草", "小龙")
			return true
		}, // 初始
		IsValid: true,
	})

	/* 138 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2645
		 调整:

		 中文名: 融合体
		 卡片种类: 融合怪兽
		 卡片密码: 01641882
		 种族: 兽
		 属性: 地
		 星级: 3
		 攻击力: 900
		 防御力: 700
		 罕见度: 平卡N
		 卡包: LB、Starter Box、TP15
		 效果: 融合：「小天使」＋「催眠羊」。

		*/
		Id:       2645,
		Password: "01641882",
		Name:     "融合体",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 3,
		La:    ygo.LA_Earth, // 地
		Lr:    ygo.LR_Beast, // 兽
		Atk:   900,
		Def:   700,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterMonsterFusion("小天使", "催眠羊")
			return true
		}, // 初始
		IsValid: true,
	})

	/* 139 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2646
		 调整:

		 中文名: 炎之骑士 基拉
		 卡片种类: 融合怪兽
		 卡片密码: 37421579
		 种族: 炎
		 属性: 炎
		 星级: 3
		 攻击力: 1100
		 防御力: 800
		 罕见度: 平卡N，金字UR
		 卡包: LB、Starter Box、TP13
		 效果: 融合：「怪兽蛋」＋「史汀」

		*/
		Id:       2646,
		Password: "37421579",
		Name:     "炎之骑士 基拉",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 3,
		La:    ygo.LA_Fire, // 炎
		Lr:    ygo.LR_Fire, // 炎
		Atk:   1100,
		Def:   800,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterMonsterFusion("怪兽蛋", "史汀")
			return true
		}, // 初始
		IsValid: true,
	})

	/* 140 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 302
		 调整:

		 中文名: 千眼纳祭神
		 卡片种类: 融合怪兽
		 卡片密码: 63519819
		 种族: 魔法师
		 属性: 暗
		 星级: 1
		 攻击力: 0
		 防御力: 0
		 罕见度: 金字UR，爆闪PR
		 卡包: BE01、TB 、DL03、PE
		 效果: 融合：「纳祭之魔」＋「千眼邪教神」。只要这张卡在场上存在，其他怪兽的表示形式不能改变，不能攻击。指定对方的1只怪兽，作为装备卡装备在这张卡上。这个效果1个回合只能用1次，同时装备的怪兽最多1只。装备怪兽的数值就是这张卡的攻击力守备力。战斗时要这张卡破坏的时候，装备怪兽代替破坏。

		*/
		Id:       302,
		Password: "63519819",
		Name:     "千眼纳祭神",
		Lc:       ygo.LC_MonsterFusion, // 融合怪兽

		Level: 1,
		La:    ygo.LA_Dark,        // 暗
		Lr:    ygo.LR_Spellcaster, // 魔法师
		Atk:   0,
		Def:   0,
		//Initialize: func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/* 141 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1163
		 调整:

		 中文名: 磁石战士α
		 卡片种类: 通常怪兽
		 卡片密码: 99785935
		 种族: 岩石
		 属性: 地
		 星级: 4
		 攻击力: 1400
		 防御力: 1700
		 罕见度: 平卡N，金字UR
		 卡包: TP05、YU、其他
		 效果: 描述：α、β、γ能变形合体。

		*/
		Id:       1163,
		Password: "99785935",
		Name:     "磁石战士α",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Rock,  // 岩石
		Atk:     1400,
		Def:     1700,
		IsValid: true,
	})

	/* 142 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1164
		 调整:

		 中文名: 磁石战士β
		 卡片种类: 通常怪兽
		 卡片密码: 39256679
		 种族: 岩石
		 属性: 地
		 星级: 4
		 攻击力: 1700
		 防御力: 1600
		 罕见度: 平卡N，金字UR
		 卡包: TP06、YU、其他
		 效果: 描述：α、β、γ能变形合体。

		*/
		Id:       1164,
		Password: "39256679",
		Name:     "磁石战士β",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Rock,  // 岩石
		Atk:     1700,
		Def:     1600,
		IsValid: true,
	})

	/* 143 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1170
		 调整:

		 中文名: 磁石战士γ
		 卡片种类: 通常怪兽
		 卡片密码: 11549357
		 种族: 岩石
		 属性: 地
		 星级: 4
		 攻击力: 1500
		 防御力: 1800
		 罕见度: 平卡N，金字UR
		 卡包: TP07、YU、其他
		 效果: 描述：α、β、γ能变形合体。

		*/
		Id:       1170,
		Password: "11549357",
		Name:     "磁石战士γ",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Rock,  // 岩石
		Atk:     1500,
		Def:     1800,
		IsValid: true,
	})

	/* 144 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1220
		 调整:

		 中文名: 翼龙战士
		 卡片种类: 通常怪兽
		 卡片密码: 64428736
		 种族: 兽
		 属性: 地
		 星级: 4
		 攻击力: 1500
		 防御力: 1200
		 罕见度: 平卡N，金字UR
		 卡包: LE02、TP08、JY、SJ2
		 效果: 描述：剑技出色的蜥蜴人，音速般地挥舞着剑。

		*/
		Id:       1220,
		Password: "64428736",
		Name:     "翼龙战士",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Atk:     1500,
		Def:     1200,
		IsValid: true,
	})

	/* 145 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 1240
		 调整:

		 中文名: 血腥魔兽人
		 卡片种类: 通常怪兽
		 卡片密码: 14898066
		 种族: 兽战士
		 属性: 暗
		 星级: 4
		 攻击力: 1900
		 防御力: 1200
		 罕见度: 平卡N，银字R，金字UR
		 卡包: LE05、DT05、DPKB、KA、SK2、其他
		 效果: 描述：以做尽坏事而引以为荣的魔兽人，手上的斧头沾满鲜血。

		*/
		Id:       1240,
		Password: "14898066",
		Name:     "血腥魔兽人",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Dark,         // 暗
		Lr:      ygo.LR_BeastWarrior, // 兽战士
		Atk:     1900,
		Def:     1200,
		IsValid: true,
	})

	/* 146 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 139
		 调整:

		 中文名: 青眼白龙
		 卡片种类: 通常怪兽
		 卡片密码: 89631139
		 种族: 龙
		 属性: 光
		 星级: 8
		 攻击力: 3000
		 防御力: 2500
		 罕见度: 平卡N，面闪SR，金字UR，爆闪PR，立体UTR
		 卡包: LB、BE01、SM、EX-R(EX)、DL02、Starter Box、DT01、YAP01、DTP01、DPKB、KA、SK2、MFC03、LC01、SD22、SD25
		 效果: 描述：以高攻击力著称的传说之龙。任何对手都能粉碎，其破坏力不可估量。

		*/
		Id:       139,
		Password: "89631139",
		Name:     "青眼白龙",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   8,
		La:      ygo.LA_Light,  // 光
		Lr:      ygo.LR_Dragon, // 龙
		Atk:     3000,
		Def:     2500,
		IsValid: true,
	})

	/* 147 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 142
		 调整:

		 中文名: 白骨
		 卡片种类: 通常怪兽
		 卡片密码: 32274490
		 种族: 不死
		 属性: 暗
		 星级: 1
		 攻击力: 300
		 防御力: 200
		 罕见度: 平卡N
		 卡包: LB、BE01、DL02、Starter Box
		 效果: 描述：到处出没的骸骨妖怪。攻击虽弱，聚集起来非同小可。

		*/
		Id:       142,
		Password: "32274490",
		Name:     "白骨",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   1,
		La:      ygo.LA_Dark,   // 暗
		Lr:      ygo.LR_Zombie, // 不死
		Atk:     300,
		Def:     200,
		IsValid: true,
	})

	/* 148 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 145
		 调整:

		 中文名: 精灵剑士
		 卡片种类: 通常怪兽
		 卡片密码: 91152256
		 种族: 战士
		 属性: 地
		 星级: 4
		 攻击力: 1400
		 防御力: 1200
		 罕见度: 平卡N，金字UR，金碎USR
		 卡包: LB、BE01、LE02、EX-R(EX)、DL02、Starter Box、YAP01
		 效果: 描述：学会了剑术的精灵。以快速的攻击作弄敌人。

		*/
		Id:       145,
		Password: "91152256",
		Name:     "精灵剑士",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     1400,
		Def:     1200,
		IsValid: true,
	})

	/* 149 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 146
		 调整:

		 中文名: 昆虫人
		 卡片种类: 通常怪兽
		 卡片密码: 89091579
		 种族: 昆虫
		 属性: 地
		 星级: 2
		 攻击力: 500
		 防御力: 700
		 罕见度: 平卡N
		 卡包: LB、DL02、Starter Box、TP19
		 效果: 描述：群居生活的昆虫。森林是他们的乐园。

		*/
		Id:       146,
		Password: "89091579",
		Name:     "昆虫人",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Earth,  // 地
		Lr:      ygo.LR_Insect, // 昆虫
		Atk:     500,
		Def:     700,
		IsValid: true,
	})

	/* 150 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 153
		 调整:

		 中文名: 幻象鸟
		 卡片种类: 通常怪兽
		 卡片密码: 02863439
		 种族: 鸟兽
		 属性: 光
		 星级: 4
		 攻击力: 1100
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: LB、DL02、Starter Box
		 效果: 描述：能够从手里的镜子中召唤出同伴的鸟怪。

		*/
		Id:       153,
		Password: "02863439",
		Name:     "幻象鸟",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Light,       // 光
		Lr:      ygo.LR_WingedBeast, // 鸟兽
		Atk:     1100,
		Def:     1400,
		IsValid: true,
	})

	/* 151 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 155
		 调整:

		 中文名: 双口的暗之支配者
		 卡片种类: 通常怪兽
		 卡片密码: 57305373
		 种族: 恐龙
		 属性: 地
		 星级: 3
		 攻击力: 900
		 防御力: 700
		 罕见度: 平卡N
		 卡包: LB、DL02、Starter Box
		 效果: 描述：同时长有两个嘴巴的恐龙，角能蓄电并在背部放电。

		*/
		Id:       155,
		Password: "57305373",
		Name:     "双口的暗之支配者",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,    // 地
		Lr:      ygo.LR_Dinosaur, // 恐龙
		Atk:     900,
		Def:     700,
		IsValid: true,
	})

	/* 152 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 156
		 调整:

		 中文名: 北风与太阳
		 卡片种类: 通常怪兽
		 卡片密码: 85309439
		 种族: 天使
		 属性: 光
		 星级: 3
		 攻击力: 1000
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: LB、DL02、Starter Box
		 效果: 描述：交情深厚的北风和太阳，能发动飓风与热光线的攻击。

		*/
		Id:       156,
		Password: "85309439",
		Name:     "北风与太阳",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Light, // 光
		Lr:      ygo.LR_Fairy, // 天使
		Atk:     1000,
		Def:     1000,
		IsValid: true,
	})

	/* 153 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 157
		 调整:

		 中文名: 烟之王
		 卡片种类: 通常怪兽
		 卡片密码: 84686841
		 种族: 恶魔
		 属性: 暗
		 星级: 3
		 攻击力: 1000
		 防御力: 900
		 罕见度: 平卡N
		 卡包: LB、DL02、Starter Box
		 效果: 描述：潜伏在烟雾中的恶魔，四周被浓烟包围，毫无踪迹可循。

		*/
		Id:       157,
		Password: "84686841",
		Name:     "烟之王",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Atk:     1000,
		Def:     900,
		IsValid: true,
	})

	/* 154 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 158
		 调整:

		 中文名: 传说的剑豪 正树
		 卡片种类: 通常怪兽
		 卡片密码: 44287299
		 种族: 战士
		 属性: 地
		 星级: 4
		 攻击力: 1100
		 防御力: 1100
		 罕见度: 平卡N
		 卡包: LB、BE01、DL02、Starter Box
		 效果: 描述：据说是完成了百人斩的传说中的剑豪。

		*/
		Id:       158,
		Password: "44287299",
		Name:     "传说的剑豪 正树",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     1100,
		Def:     1100,
		IsValid: true,
	})

	/* 155 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 20
		 调整:

		 中文名: 拉弓的人鱼
		 卡片种类: 通常怪兽
		 卡片密码: 65570596
		 种族: 水
		 属性: 水
		 星级: 4
		 攻击力: 1400
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: BE01、MR、DL01、PE
		 效果: 描述：平常躲在贝壳里，对擅自接近的人射箭的美人鱼。

		*/
		Id:       20,
		Password: "65570596",
		Name:     "拉弓的人鱼",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     1400,
		Def:     1500,
		IsValid: true,
	})

	/* 156 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 224
		 调整:

		 中文名: 被封印者的右腕
		 卡片种类: 通常怪兽
		 卡片密码: 70903634
		 种族: 魔法师
		 属性: 暗
		 星级: 1
		 攻击力: 200
		 防御力: 300
		 罕见度: 银字R，金字UR
		 卡包: PG、BE01、DL02
		 效果: 描述：被封印者的右腕。封印解开后将得到无限的力量。

		*/
		Id:       224,
		Password: "70903634",
		Name:     "被封印者的右腕",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   1,
		La:      ygo.LA_Dark,        // 暗
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     200,
		Def:     300,
		IsValid: true,
	})

	/* 157 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 225
		 调整:

		 中文名: 被封印者的左腕
		 卡片种类: 通常怪兽
		 卡片密码: 07902349
		 种族: 魔法师
		 属性: 暗
		 星级: 1
		 攻击力: 200
		 防御力: 300
		 罕见度: 银字R，金字UR
		 卡包: PG、BE01、DL02
		 效果: 描述：被封印者的左腕。封印解开后将得到无限的力量。

		*/
		Id:       225,
		Password: "07902349",
		Name:     "被封印者的左腕",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   1,
		La:      ygo.LA_Dark,        // 暗
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     200,
		Def:     300,
		IsValid: true,
	})

	/* 158 */
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

	/* 159 */
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

	/* 160 */
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

	/* 161 */
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

	/* 162 */
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

	/* 163 */
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

	/* 164 */
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

	/* 165 */
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

	/* 166 */
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

	/* 167 */
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

	/* 168 */
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

	/* 169 */
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

	/* 170 */
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

	/* 171 */
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

	/* 172 */
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

	/* 173 */
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

	/* 174 */
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

	/* 175 */
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

	/* 176 */
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

	/* 177 */
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

	/* 178 */
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

	/* 179 */
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

	/* 180 */
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

	/* 181 */
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

	/* 182 */
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

	/* 183 */
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

	/* 184 */
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

	/* 185 */
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

	/* 186 */
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

	/* 187 */
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

	/* 188 */
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

	/* 189 */
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

	/* 190 */
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

	/* 191 */
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

	/* 192 */
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

	/* 193 */
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

	/* 194 */
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

	/* 195 */
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

	/* 196 */
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

	/* 197 */
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

	/* 198 */
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

	/* 199 */
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

	/* 200 */
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

	/* 201 */
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

	/* 202 */
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

	/* 203 */
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

	/* 204 */
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

	/* 205 */
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

	/* 206 */
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

	/* 207 */
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

	/* 208 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2619
		 调整:

		 中文名: 龟虎
		 卡片种类: 通常怪兽
		 卡片密码: 37313348
		 种族: 水
		 属性: 水
		 星级: 4
		 攻击力: 1000
		 防御力: 1500
		 罕见度: 平卡N，金碎USR
		 卡包: LB、Starter Box、Booster R1、TP03
		 效果: 描述：拥有甲壳的老虎。用坚硬的甲壳保护身体，尖锐的牙齿发动攻击。

		*/
		Id:       2619,
		Password: "37313348",
		Name:     "龟虎",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Water, // 水
		Atk:     1000,
		Def:     1500,
		IsValid: true,
	})

	/* 209 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2621
		 调整:

		 中文名: 怪兽蛋
		 卡片种类: 通常怪兽
		 卡片密码: 36121917
		 种族: 战士
		 属性: 地
		 星级: 3
		 攻击力: 600
		 防御力: 900
		 罕见度: 平卡N
		 卡包: LB、Starter Box、TP13
		 效果: 描述：蛋壳保卫全身的迷之战士，用壳飞射攻击。

		*/
		Id:       2621,
		Password: "36121917",
		Name:     "怪兽蛋",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     600,
		Def:     900,
		IsValid: true,
	})

	/* 210 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2632
		 调整:

		 中文名: 催眠羊
		 卡片种类: 通常怪兽
		 卡片密码: 83464209
		 种族: 兽
		 属性: 地
		 星级: 3
		 攻击力: 800
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: LB、Starter Box、TP15
		 效果: 描述：尾巴很长的绵羊，用尾巴使出催眠术，让敌人入睡。

		*/
		Id:       2632,
		Password: "83464209",
		Name:     "催眠羊",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Atk:     800,
		Def:     1000,
		IsValid: true,
	})

	/* 211 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2634
		 调整:

		 中文名: 绿树之灵王
		 卡片种类: 通常怪兽
		 卡片密码: 22910685
		 种族: 植物
		 属性: 地
		 星级: 3
		 攻击力: 500
		 防御力: 1600
		 罕见度: 平卡N
		 卡包: LB、Starter Box
		 效果: 描述：居住在被绿树包围的地方，统治森林的年轻国王。

		*/
		Id:       2634,
		Password: "22910685",
		Name:     "绿树之灵王",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Plant, // 植物
		Atk:     500,
		Def:     1600,
		IsValid: true,
	})

	/* 212 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2637
		 调整:

		 中文名: 根潭鱼人
		 卡片种类: 通常怪兽
		 卡片密码: 39004808
		 种族: 鱼
		 属性: 水
		 星级: 3
		 攻击力: 900
		 防御力: 800
		 罕见度: 平卡N
		 卡包: LB、Starter Box
		 效果: 描述：潜伏在海里的半鱼人，掀起黑暗浪潮袭击。

		*/
		Id:       2637,
		Password: "39004808",
		Name:     "根潭鱼人",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Fish,  // 鱼
		Atk:     900,
		Def:     800,
		IsValid: true,
	})

	/* 213 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2640
		 调整:

		 中文名: 史汀
		 卡片种类: 通常怪兽
		 卡片密码: 96851799
		 种族: 炎
		 属性: 炎
		 星级: 2
		 攻击力: 600
		 防御力: 500
		 罕见度: 平卡N
		 卡包: LB、Starter Box、TP13
		 效果: 描述：非常炽热的火焰形成的块，并用那身体冲撞敌人。

		*/
		Id:       2640,
		Password: "96851799",
		Name:     "史汀",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   2,
		La:      ygo.LA_Fire, // 炎
		Lr:      ygo.LR_Fire, // 炎
		Atk:     600,
		Def:     500,
		IsValid: true,
	})

	/* 214 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2641
		 调整:

		 中文名: 岩浆人
		 卡片种类: 通常怪兽
		 卡片密码: 40826495
		 种族: 岩石
		 属性: 地
		 星级: 3
		 攻击力: 900
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: LB、Starter Box
		 效果: 描述：岩浆里出生的怪兽，灼热的身体使接近的物体融化。

		*/
		Id:       2641,
		Password: "40826495",
		Name:     "岩浆人",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Rock,  // 岩石
		Atk:     900,
		Def:     1000,
		IsValid: true,
	})

	/* 215 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 2770
		 调整:

		 中文名: 齿轮战士
		 卡片种类: 通常怪兽
		 卡片密码: 86281779
		 种族: 机械
		 属性: 炎
		 星级: 6
		 攻击力: 1800
		 防御力: 2000
		 罕见度: 平卡N
		 卡包: SM、KA、SK2
		 效果: 描述：为战斗而生的机器人战士。用不会生锈的金属制造而成。

		*/
		Id:       2770,
		Password: "86281779",
		Name:     "齿轮战士",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   6,
		La:      ygo.LA_Fire,    // 炎
		Lr:      ygo.LR_Machine, // 机械
		Atk:     1800,
		Def:     2000,
		IsValid: true,
	})

	/* 216 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 301
		 调整:

		 中文名: 千眼邪教神
		 卡片种类: 通常怪兽
		 卡片密码: 27125110
		 种族: 魔法师
		 属性: 暗
		 星级: 1
		 攻击力: 0
		 防御力: 0
		 罕见度: 平卡N
		 卡包: BE01、TB 、DL03、PE
		 效果: 描述：能够操纵人心的邪神，千只邪眼能够将人的负面的心看透并不断增强。

		*/
		Id:       301,
		Password: "27125110",
		Name:     "千眼邪教神",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   1,
		La:      ygo.LA_Dark,        // 暗
		Lr:      ygo.LR_Spellcaster, // 魔法师
		Atk:     0,
		Def:     0,
		IsValid: true,
	})

	/* 217 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 313
		 调整:

		 中文名: 地星剑士
		 卡片种类: 通常怪兽
		 卡片密码: 03573512
		 种族: 战士
		 属性: 地
		 星级: 3
		 攻击力: 500
		 防御力: 1200
		 罕见度: 平卡N，面闪SR
		 卡包: BE01、SM、DL03、JY、SJ2
		 效果: 描述：虽然剑术不太成熟，但却有不可思议力量的妖精剑士。

		*/
		Id:       313,
		Password: "03573512",
		Name:     "地星剑士",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   3,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Atk:     500,
		Def:     1200,
		IsValid: true,
	})

	/* 218 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 460
		 调整:

		 中文名: 幻兽王 加泽尔
		 卡片种类: 通常怪兽
		 卡片密码: 05818798
		 种族: 兽
		 属性: 地
		 星级: 4
		 攻击力: 1500
		 防御力: 1200
		 罕见度: 平卡N，金字UR
		 卡包: ME、BE02、LE02、YSD01、DL04、YU、SY2
		 效果: 描述：跑动速度快的惊人，身姿犹如幻影一般的野兽。

		*/
		Id:       460,
		Password: "05818798",
		Name:     "幻兽王 加泽尔",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   4,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Atk:     1500,
		Def:     1200,
		IsValid: true,
	})

	/* 219 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 461
		 调整:

		 中文名: 甘尼许巨象
		 卡片种类: 通常怪兽
		 卡片密码: 49888191
		 种族: 兽战士
		 属性: 地
		 星级: 7
		 攻击力: 2400
		 防御力: 2000
		 罕见度: 平卡N，金字UR
		 卡包: ME、DL04、VJ
		 效果: 描述：拥有可怕的力量，因体形臃肿，行走时会发生地震。

		*/
		Id:       461,
		Password: "49888191",
		Name:     "甘尼许巨象",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   7,
		La:      ygo.LA_Earth,        // 地
		Lr:      ygo.LR_BeastWarrior, // 兽战士
		Atk:     2400,
		Def:     2000,
		IsValid: true,
	})

	/* 220 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 55
		 调整:

		 中文名: 蛋龙
		 卡片种类: 通常怪兽
		 卡片密码: 02964201
		 种族: 龙
		 属性: 炎
		 星级: 7
		 攻击力: 2200
		 防御力: 2600
		 罕见度: 平卡N
		 卡包: PS、DL01、PE
		 效果: 描述：被蛋壳包围的龙。如果认为只是个小鬼的话会吃苦头的哦。

		*/
		Id:       55,
		Password: "02964201",
		Name:     "蛋龙",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   7,
		La:      ygo.LA_Fire,   // 炎
		Lr:      ygo.LR_Dragon, // 龙
		Atk:     2200,
		Def:     2600,
		IsValid: true,
	})

	/* 221 */
	cardBag.Register(&ygo.CardOriginal{
		/*
		 id: 89
		 调整:

		 中文名: 铁腕巨人
		 卡片种类: 通常怪兽
		 卡片密码: 90908427
		 种族: 机械
		 属性: 地
		 星级: 6
		 攻击力: 1900
		 防御力: 2200
		 罕见度: 平卡N
		 卡包: CA、DL01、KA、SK2
		 效果: 描述：用钢铁做成的机器人偶，拥有恐怖的怪力。

		*/
		Id:       89,
		Password: "90908427",
		Name:     "铁腕巨人",
		Lc:       ygo.LC_MonsterNormal, // 通常怪兽

		Level:   6,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Machine, // 机械
		Atk:     1900,
		Def:     2200,
		IsValid: true,
	})
}
