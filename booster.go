package cards

import ygo "github.com/wzshiming/ygo_core"

func booster(cardBag *ygo.CardVersion) {

	/*0*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 1357
		 调整: [陷阱大师]
		<トラップ·マスター>
		[14/08/07]

		●反转：场上1张陷阱卡破坏。里侧表示翻开确认后破坏。
		◇诱发效果，强制发动，开连锁，发动时选择场上的1张表侧表示的陷阱卡或里侧表示的魔法或陷阱卡（取对象）
		◇场上没有可以作为对象的卡时，这个效果仍然会发动（此时不取对象，效果处理时不处理）
		◇选择的卡为里侧表示的场合，在效果处理时确认那张卡，是陷阱则破坏，不是则原状放回
		 中文名: 陷阱大师
		 日文名: トラップ·マスター
		 英文名: Trap Master
		 卡片种类: 效果怪兽
		 卡片密码: 46461247
		 使用限制: 无限制
		 种族: 战士
		 属性: 地
		 星级: 3
		 攻击力: 500
		 防御力: 1100
		 罕见度: 平卡N
		 卡包: EX-R(EX)，Booster05，Booster R3
		 效果: 反转：场上1张陷阱卡破坏。里侧表示翻开确认后破坏。
		}
		*/
		Id:       1357,
		Password: "46461247",
		Name:     "陷阱大师",               // "Trap Master"  "トラップ·マスター"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   3,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Attack:  500,
		Defense: 1100,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterFlip(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				if c := pl.SelectForWarn(tar.Szone, pl.Szone); c != nil {
					c.SetFaceUp()
					if c.IsTrap() {
						c.Dispatch(ygo.Destroy, ca)
					} else {
						c.SetFaceDown()
					}
				}
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*1*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 1358
		 调整: [血之代偿]
		<血の代偿>
		[11/02/09]

		●可以支付500基本分，把1只怪兽通常召唤。这个效果只能在自己的主要阶段及对方的战斗阶段时可以发动。
		◇效果发动时支付500基本分（代价）
		◇效果的发动进入连锁（咒文速度2）
		◇在卡的发动时可以同时发动效果，那个场合支付500基本分（卡的发动没有条件限制）
		◇效果处理时选择自己手卡的1只可以被通常召唤的怪兽进行通常召唤（不取对象）
		◇自己手卡没有能被通常召唤的怪兽的场合，不能发动这个效果
		★效果处理时，对方没有可以召唤的怪兽的场合，能否确认对方手卡？调整中
		◇这个通常召唤遵循普通的通常召唤规则（上级召唤在效果处理时需要解放怪兽）
		◇这张卡的发动的那组连锁中，这张卡不能再在那组连锁中发动效果（卡的发动尚处理完毕，所以不能再发动效果）
		★用这张卡的效果在连锁1进行通常召唤的场合，是否可以对应发动「反冲/キックバック」「升天之角笛/昇天の角笛」 调整中
		◇手卡中有多少能被通常召唤的怪兽，就限定了能在一组连锁中发动的次数
		◇「王宫的通告/王宫のお触れ」发动中，可以反复发动这张卡来支付基本分
		◇可以让二重怪兽进行再度召唤
		◇手卡不存在怪兽的场合，可以让二重怪兽进行再度召唤
		★场上没卡，手卡只有1只二重怪兽，是否可以假设（连锁2）召唤二重怪兽，（连锁1）进行再度召唤来在同一组连锁中发动两次这张卡的效果 调整中
		◇假设手卡只有「栗子球/クリボー」「恶魔召唤/デーモンの召唤」，自己场上只有这张卡，不能以假设（连锁2）上级召唤「恶魔召唤/デーモンの召唤」（以连锁1的「栗子球/クリボー」作为解放），（连锁1）召唤「栗子球/クリボー」
		◇与1回合1次正常的通常召唤没有冲突
		 中文名: 血之代偿
		 日文名: 血の代償
		 英文名: Ultimate Offering
		 卡片种类: 永续陷阱
		 卡片密码: 80604091
		 使用限制: 禁止卡
		 罕见度: 平卡N，银字R，金字UR
		 卡包: EX-R(EX)，Booster03，SD07，SD10，Booster R2，DT06，YU，JY，SJ2，DS14
		 效果: 可以支付500基本分，把1只怪兽通常召唤。这个效果在自己回合的主要阶段及对方回合的战斗阶段才能发动。
		}
		*/
		Id:       1358,
		Password: "80604091",
		Name:     "血之代偿",              // "Ultimate Offering"  "血の代償"
		Lc:       ygo.LC_SustainsTrap, // 永续陷阱

		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*2*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 1359
		 调整: [谜之傀儡师]
		<謎の傀儡師>
		[2010/09/10]

		●怪兽召唤·反转召唤成功时，自己回复500基本分。
		◇诱发效果（进入连锁）
		◇强制发动
		◇这张卡自身召唤、反转召唤成功时不发动
		◇二重怪兽再度召唤成功时也发动
		◇对方场上召唤、反转召唤怪兽成功时也发动
		★同一组连锁中召唤了2次怪兽的场合，连锁处理完毕后这张卡发动几次效果 调整中
		 中文名: 谜之傀儡师
		 日文名: 謎の傀儡師
		 英文名: Mysterious Puppeteer
		 卡片种类: 效果怪兽
		 卡片密码: 54098121
		 使用限制: 无限制
		 种族: 战士
		 属性: 地
		 星级: 4
		 攻击力: 1000
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: EX-R(EX)，Booster05，Booster R3，TP07
		 效果: 怪兽召唤·反转召唤成功时，自己回复500基本分。
		}
		*/
		Id:       1359,
		Password: "54098121",
		Name:     "谜之傀儡师",              // "Mysterious Puppeteer"  "謎の傀儡師"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   4,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Attack:  1000,
		Defense: 1500,
		Initialize: func(ca *ygo.Card) bool {
			e := func() {
				pl := ca.GetSummoner()
				pl.ChangeHp(500)
			}
			ca.AddEvent(ygo.Summon, e)
			ca.AddEvent(ygo.SummonFlip, e)
			return true
		}, // 初始

		IsValid: true,
	})

	/*3*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 1420
		 调整: [自业自得]
		<自業自得>
		[2010/09/10]

		●对方场上每存在1只怪兽，给予对方基本分500分伤害。
		◇对方场上不存在怪兽的场合，不能发动
		◇在效果处理时计算对方场上怪兽的数量
		 中文名: 自业自得
		 日文名: 自業自得
		 英文名: Just Desserts
		 卡片种类: 通常陷阱
		 卡片密码: 24068492
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: EX-R(EX)，Booster05，SD12，Booster R3
		 效果: 对方场上每存在1只怪兽对方受到500基本分伤害。
		}
		*/
		Id:       1420,
		Password: "24068492",
		Name:     "自业自得",              // "Just Desserts"  "自業自得"
		Lc:       ygo.LC_OrdinaryTrap, // 通常陷阱

		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*4*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2081
		 调整: [城壁]
		<城壁>
		[2010/09/10]

		●选择场上表侧表示存在的1只怪兽发动。被选择怪兽的守备力直到结束阶段时为止上升500分。
		◇发动时选择场上1只表侧表示存在的怪兽（取对象）
		◇可以在伤害步骤发动（具体请参考置顶内的伤害步骤详解）
		 中文名: 城壁
		 日文名: 城壁
		 英文名: Castle Walls
		 卡片种类: 通常陷阱
		 卡片密码: 44209392
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: EX-R(EX)，YSD01，Booster03，Booster R2，TP11
		 效果: 表侧表示的1只怪兽守备力在回合结束前上升500。
		}
		*/
		Id:       2081,
		Password: "44209392",
		Name:     "城壁",                // "Castle Walls"  "城壁"
		Lc:       ygo.LC_OrdinaryTrap, // 通常陷阱

		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*5*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2082
		 调整: [援军]
		<援軍>
		[2010/09/10]

		●直到结束阶段时为止，1只表侧表示存在的怪兽的攻击力上升500分。
		◇发动时选择场上1只表侧表示存在的怪兽（取对象）
		◇可以在伤害步骤发动（具体请参考置顶内的伤害步骤详解）
		 中文名: 援军
		 日文名: 援軍
		 英文名: Reinforcements
		 卡片种类: 通常陷阱
		 卡片密码: 17814387
		 使用限制: 无限制
		 罕见度: 平卡N，黄金GR
		 卡包: EX-R(EX)，YSD01，Booster03，Booster R2，DT01，GLD01，DTP01
		 效果: 表侧表示的1只怪兽攻击力在回合结束前上升500。
		}
		*/
		Id:       2082,
		Password: "17814387",
		Name:     "援军",                // "Reinforcements"  "援軍"
		Lc:       ygo.LC_OrdinaryTrap, // 通常陷阱

		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*6*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2083
		 调整: [昼夜的大火事]
		<昼夜の大火事>
		[2010/09/10]

		●给予对方基本分800分伤害。
		◇「精灵之镜/精霊の鏡」可以对应这张卡的发动而发动
		 中文名: 昼夜的大火事
		 日文名: 昼夜の大火事
		 英文名: Ookazi
		 卡片种类: 通常魔法
		 卡片密码: 19523799
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: EX-R(EX)，YSD01，Booster02，Booster R1，YSD04
		 效果: 给予对方基本分800分伤害。
		}
		*/
		Id:       2083,
		Password: "19523799",
		Name:     "昼夜的大火事",             // "Ookazi"  "昼夜の大火事"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				tar.ChangeHp(-800)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*7*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2351
		 调整: [尖刺神的杀虫剂]
		<トゲトゲ神の殺虫剤>
		[2010/09/08]

		●场上的昆虫族怪兽全部破坏。
		◇不取对象
		◇效果处理时在场上表侧表示存在的昆虫族怪兽全部破坏
		 中文名: 尖刺神的杀虫剂
		 日文名: トゲトゲ神の殺虫剤
		 英文名: Eradicating Aerosol
		 卡片种类: 通常魔法
		 卡片密码: 94716515
		 使用限制: 无限制
		 罕见度: 平卡N，银字R
		 卡包: VOL04，Booster04，GLD04
		 效果: 场上表侧表示存在的昆虫族怪兽全部破坏。
		}
		*/
		Id:       2351,
		Password: "94716515",
		Name:     "尖刺神的杀虫剂",            // "Eradicating Aerosol"  "トゲトゲ神の殺虫剤"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				cs := ygo.NewCards(&pl.Mzone.Cards, &tar.Mzone.Cards)
				cs.ForEach(func(c *ygo.Card) bool {
					if c.IsFaceUp() && c.RaceIsInsect() {
						c.Dispatch(ygo.Destroy, ca)
					}
					return true
				})
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*8*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2352
		 调整: [永远的渴水]
		<永遠の渇水>
		[2010/09/08]

		●场上的鱼族怪兽全部破坏。
		◇不取对象
		◇效果处理时在场上表侧表示存在的鱼族怪兽全部破坏
		 中文名: 永远的渴水
		 日文名: 永遠の渇水
		 英文名: Eternal Drought
		 卡片种类: 通常魔法
		 卡片密码: 56606928
		 使用限制: 无限制
		 罕见度: 平卡N，银字R
		 卡包: VOL04，Booster04，GLD04
		 效果: 场上表侧表示存在的鱼族怪兽全部破坏。
		}
		*/
		Id:       2352,
		Password: "56606928",
		Name:     "永远的渴水",              // "Eternal Drought"  "永遠の渇水"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				cs := ygo.NewCards(&pl.Mzone.Cards, &tar.Mzone.Cards)
				cs.ForEach(func(c *ygo.Card) bool {
					if c.IsFaceUp() && c.RaceIsFish() {
						c.Dispatch(ygo.Destroy, ca)
					}
					return true
				})
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*9*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2353
		 调整: [酸性风暴]
		<酸の嵐>
		[2010/09/08]

		●场上的机械族怪兽全部破坏。
		◇不取对象
		◇效果处理时在场上表侧表示存在的机械族怪兽全部破坏
		 中文名: 酸性风暴
		 日文名: 酸の嵐
		 英文名: Acid Rain
		 卡片种类: 通常魔法
		 卡片密码: 21323861
		 使用限制: 无限制
		 罕见度: 平卡N，银字R
		 卡包: VOL04，Booster04，TP01，DT03
		 效果: 场上表侧表示存在的机械族怪兽全部破坏。
		}
		*/
		Id:       2353,
		Password: "21323861",
		Name:     "酸性风暴",               // "Acid Rain"  "酸の嵐"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				cs := ygo.NewCards(&pl.Mzone.Cards, &tar.Mzone.Cards)
				cs.ForEach(func(c *ygo.Card) bool {
					if c.IsFaceUp() && c.RaceIsMachine() {
						c.Dispatch(ygo.Destroy, ca)
					}
					return true
				})
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*10*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2354
		 调整: [神之息吹]
		<神の息吹>
		[2010/09/08]

		●场上的岩石族怪兽全部破坏。
		◇不取对象
		◇效果处理时在场上表侧表示存在的岩石族怪兽全部破坏
		 中文名: 神之息吹
		 日文名: 神の息吹
		 英文名: Breath of Light
		 卡片种类: 通常魔法
		 卡片密码: 20101223
		 使用限制: 无限制
		 罕见度: 平卡N，银字R
		 卡包: VOL04，Booster04，TP03
		 效果: 场上表侧表示存在的岩石族怪兽全部破坏。
		}
		*/
		Id:       2354,
		Password: "20101223",
		Name:     "神之息吹",               // "Breath of Light"  "神の息吹"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				cs := ygo.NewCards(&pl.Mzone.Cards, &tar.Mzone.Cards)
				cs.ForEach(func(c *ygo.Card) bool {
					if c.IsFaceUp() && c.RaceIsRock() {
						c.Dispatch(ygo.Destroy, ca)
					}
					return true
				})
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*11*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2355
		 调整: [战士抹杀]
		<戦士抹殺>
		[2010/09/08]

		●场上的战士族怪兽全部破坏。
		◇不取对象
		◇效果处理时在场上表侧表示存在的战士族怪兽全部破坏
		 中文名: 战士抹杀
		 日文名: 戦士抹殺
		 英文名: Warrior Elimination
		 卡片种类: 通常魔法
		 卡片密码: 90873992
		 使用限制: 无限制
		 罕见度: 平卡N，银字R
		 卡包: VOL04，Booster04，TP02，GLD02
		 效果: 场上表侧表示存在的战士族怪兽全部破坏。
		}
		*/
		Id:       2355,
		Password: "90873992",
		Name:     "战士抹杀",               // "Warrior Elimination"  "戦士抹殺"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				cs := ygo.NewCards(&pl.Mzone.Cards, &tar.Mzone.Cards)
				cs.ForEach(func(c *ygo.Card) bool {
					if c.IsFaceUp() && c.RaceIsWarrior() {
						c.Dispatch(ygo.Destroy, ca)
					}
					return true
				})
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*12*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2377
		 调整: [来自黑暗的呼声]
		<闇からの呼び声>
		[2010/09/08]

		●用「死者苏生」的效果特殊召唤的怪兽全部送去墓地。只要这张卡在场上存在，双方不能使用「死者苏生/死者蘇生」。
		◇这张卡对应「死者苏生/死者蘇生」的发动而发动的场合，在「死者苏生/死者蘇生」效果处理时，「死者苏生/死者蘇生」的效果不适用
		◇这张卡的永续效果适用时，之前由于「死者苏生/死者蘇生」的效果特殊召唤的怪兽就送去墓地
		◇这张卡效果适用中一度被无效化，期间用「死者苏生/死者蘇生」特殊召唤了怪兽，这张卡效果再次适用时那些怪兽送去墓地
		★用「死者苏生/死者蘇生」的效果特殊召唤的怪兽一度不在场上表侧表示存在后再次出现的场合，是否会被送去墓地 调整中
		 中文名: 来自黑暗的呼声
		 日文名: 闇からの呼び声
		 英文名: Call of Darkness
		 卡片种类: 永续陷阱
		 卡片密码: 78637313
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: VOL05，Booster05，TP09
		 效果: 用「死者苏生」苏生的怪兽全部送到墓地。只要场上有这张卡「死者苏生」就不能使用。
		}
		*/
		Id:       2377,
		Password: "78637313",
		Name:     "来自黑暗的呼声",           // "Call of Darkness"  "闇からの呼び声"
		Lc:       ygo.LC_SustainsTrap, // 永续陷阱

		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*13*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2378
		 调整: [避雷针]
		<避雷針>
		[14/09/14]

		●对方使用「雷击」时，把对方怪兽全部破坏来代替自己怪兽。发动后这张卡破坏。
		◇对方场上没有怪兽的场合，这张卡也能发动
		◇这个效果是让对方的「雷击/サンダー·ボルト」无效、把对方场上的怪兽全部破坏的效果
		◇这个效果不会把「雷击/サンダー·ボルト」的效果变为“破坏自己场上的全部怪兽”
		◇此卡要破坏2只以上对方怪兽时，「星光大道/スターライト·ロード」能对应这张卡的发动而发动
		 中文名: 避雷针
		 日文名: 避雷針
		 英文名: Anti Raigeki
		 卡片种类: 通常陷阱
		 卡片密码: 42364257
		 使用限制: 无限制
		 罕见度: 银字R
		 卡包: VOL05，Booster05
		 效果: 对方使用了「雷击」的时候，破坏对方全部怪兽代替自己的怪兽。发动后这张卡破坏。
		}
		*/
		Id:       2378,
		Password: "42364257",
		Name:     "避雷针",               // "Anti Raigeki"  "避雷針"
		Lc:       ygo.LC_OrdinaryTrap, // 通常陷阱

		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*14*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2554
		 中文名: 水精龙
		 日文名: アクア·ドラゴン
		 英文名: Aqua Dragon
		 卡片种类: 融合怪兽
		 卡片密码: 86164529
		 使用限制: 无限制
		 种族: 海龙
		 属性: 水
		 星级: 6
		 攻击力: 2250
		 防御力: 1900
		 罕见度: 平卡N
		 卡包: Booster04，TP05
		 效果: 融合：「妖精龙」＋「海原的女战士」＋「区域吞噬者」。
		}
		*/
		Id:       2554,
		Password: "86164529",
		Name:     "水精龙",                // "Aqua Dragon"  "アクア·ドラゴン"
		Lc:       ygo.LC_FusionMonster, // 融合怪兽

		Level:   6,
		La:      ygo.LA_Water,      // 水
		Lr:      ygo.LR_Seaserpent, // 海龙
		Attack:  2250,
		Defense: 1900,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterFusionMonster("妖精龙", "海原的女战士", "区域吞噬者")
			return true
		}, // 初始
		IsValid: true,
	})

	/*15*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 421
		 调整: [暗晦之城]
		<闇晦ましの城>
		[2010/09/10]

		●反转：在场上表侧表示存在的不死族怪兽的攻击力与守备力上升200分。并且，只要这张卡在场上表侧表示存在，每次自己的准备阶段上升200分。这个效果持续到自己的第4回合的准备阶段为止。
		◇诱发效果，强制发动，开连锁，不取对象
		◇准备阶段攻守上升的对象为受到这个反转效果发动时效果的怪兽
		◇这张卡不在场上表侧表示存在的场合，这些攻守上升重置
		◇准备阶段的攻守提升进入连锁（是反转效果的一部分）
		 中文名: 暗晦之城
		 日文名: 闇晦ましの城
		 英文名: Castle of Dark Illusions
		 卡片种类: 效果怪兽
		 卡片密码: 00062121
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 920
		 防御力: 1930
		 罕见度: 平卡N
		 卡包: ME，DL04，Booster07
		 效果: 反转：场上表侧表示的全部不死族的怪兽攻击力·守备力加200。只要这张卡在场上表侧表示存在，自己的每个准备阶段不死族的攻守力都会加200。这个效果持续到自己的第4个准备阶段。
		}
		*/
		Id:       421,
		Password: "00062121",
		Name:     "暗晦之城",               // "Castle of Dark Illusions"  "闇晦ましの城"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   4,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Attack:  920,
		Defense: 1930,
		//		Initialize: func(ca *ygo.Card) bool {
		//			i := 1
		//			e := ca.EffectAllMzoneHalo(func(c *ygo.Card) bool {
		//				if c.RaceIsZombie() {
		//					c.SetAttack(c.GetAttack() + i*200)
		//					c.SetDefense(c.GetDefense() + i)*200)
		//					ca.OnlyOnce(ygo.Disabled, func() {
		//						c.SetAttack(c.GetAttack() - i*200)
		//						c.SetDefense(c.GetDefense() - i*200)
		//					}, c)
		//				}
		//				return true
		//			})
		//			ca.RegisterFlip(func() {
		//				e()
		//				ca.RegisterGlobalListen(ygo.SP, func(pl0 *ygo.Player) {
		//					if i >= 5 {
		//						return
		//					}
		//					i++
		//					pl := ca.GetSummoner()
		//					tar := pl.GetTarget()
		//					if pl == pl0 {
		//						cs := ygo.NewCards(tar.Mzone, pl.Mzone)
		//						cs.ForEach(func(c *ygo.Card) bool {
		//							if c.RaceIsZombie() {
		//								c.SetAttack(c.GetAttack() + 200)
		//								c.SetDefense(c.GetDefense() + 200)
		//							}
		//							return true
		//						})
		//					}
		//				})
		//			})
		//			return true
		//		}, // 初始
		IsValid: false,
	})

	/*16*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 426
		 调整: [杀人小丑]
		<マーダーサーカス>
		[2010/09/10]

		●这张卡的表示形式从守备表示变为攻击表示时，对方场上的1只怪兽回到持有者的手卡。
		◇诱发效果（进入连锁）
		◇强制发动
		◇发动时选择对方场上1只怪兽（取对象）
		 中文名: 杀人小丑
		 日文名: マーダーサーカス
		 英文名: Crass Clown
		 卡片种类: 效果怪兽
		 卡片密码: 93889755
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 1350
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: ME，BE02，DL04，Booster07
		 效果: 这张卡的表示形式从守备表示变成攻击表示时，对方场上的1只怪兽回到持有者手卡。
		}
		*/
		Id:       426,
		Password: "93889755",
		Name:     "杀人小丑",               // "Crass Clown"  "マーダーサーカス"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   4,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Attack:  1350,
		Defense: 1400,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*17*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 427
		 调整: [幽灵王-南瓜王-]
		<ゴースト王－パンプキング－>
		[2010/09/10]

		●只要「暗晦之城」在场上表侧表示存在，这张卡的攻击力与守备力上升100分。
		◇永续效果（不进入连锁）
		●并且，每次自己的准备阶段上升100分。这个效果持续到自己的第4回合的准备阶段为止。
		◇诱发效果（进入连锁）
		◇强制发动
		 中文名: 幽灵王-南瓜王-
		 日文名: ゴースト王－パンプキング－
		 英文名: Pumpking the King of Ghosts
		 卡片种类: 效果怪兽
		 卡片密码: 29155212
		 使用限制: 无限制
		 种族: 不死
		 属性: 暗
		 星级: 6
		 攻击力: 1800
		 防御力: 2000
		 罕见度: 平卡N
		 卡包: ME，DL04，Booster07
		 效果: 只要「暗晦之城」在场上表侧表示存在，这张卡的攻击力·守备力上升100。每次自己的准备阶段这张卡的攻防加100。这个效果持续到自己的第4个准备阶段。
		}
		*/
		Id:       427,
		Password: "29155212",
		Name:     "幽灵王-南瓜王-",           // "Pumpking the King of Ghosts"  "ゴースト王－パンプキング－"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   6,
		La:      ygo.LA_Dark,   // 暗
		Lr:      ygo.LR_Zombie, // 不死
		Attack:  1800,
		Defense: 2000,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*18*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 428
		 调整: [梦幻小丑]
		<ドリーム·ピエロ>
		[2010/09/10]

		●这张卡的表示形式从攻击表示变为守备表示时，破坏对方场上的1只怪兽。
		◇诱发效果（进入连锁）
		◇强制发动
		◇发动时选择对方场上1只怪兽（取对象）
		◇变为里侧守备表示的场合不发动
		 中文名: 梦幻小丑
		 日文名: ドリーム·ピエロ
		 英文名: Dream Clown
		 卡片种类: 效果怪兽
		 卡片密码: 13215230
		 使用限制: 无限制
		 种族: 战士
		 属性: 地
		 星级: 3
		 攻击力: 1200
		 防御力: 900
		 罕见度: 平卡N
		 卡包: ME，BE02，DL04，Booster07，PE
		 效果: 这张卡的表示形式从攻击表示变成守备表示时，破坏对方场上的1只怪兽。
		}
		*/
		Id:       428,
		Password: "13215230",
		Name:     "梦幻小丑",               // "Dream Clown"  "ドリーム·ピエロ"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   3,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Attack:  1200,
		Defense: 900,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*19*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 429
		 调整: [恶魔的智慧]
		<悪魔の知恵>
		[2010/09/10]

		（这张卡在规则上不当作「恶魔/デーモン」卡使用）
		●这张卡的表示形式从攻击表示变为守备表示时，洗切自己的卡组。
		◇诱发效果（进入连锁）
		◇强制发动
		◇变为里侧守备表示的场合不发动
		◇自己卡组只有1张或0张的场合，这个效果也发动
		 中文名: 恶魔的智慧
		 日文名: 悪魔の知恵
		 英文名: Tainted Wisdom
		 卡片种类: 效果怪兽
		 卡片密码: 28725004
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 3
		 攻击力: 1250
		 防御力: 800
		 罕见度: 平卡N
		 卡包: ME，DL04，Booster07
		 效果: （这张卡在规则上不当作「恶魔」卡使用）这张卡的表示形式从攻击表示变成守备表示时，洗自己的卡组。
		}
		*/
		Id:       429,
		Password: "28725004",
		Name:     "恶魔的智慧",              // "Tainted Wisdom"  "悪魔の知恵"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   3,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Attack:  1250,
		Defense: 800,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*20*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 430
		 调整: [艾尔的小剑士]
		<アイルの小剣士>
		[2010/09/10]

		●每把自己的场上存在的这张卡以外的1只怪兽作为祭品，这张卡的攻击力直到回合结束时为止上升700分。
		◇启动效果（进入连锁）
		◇发动时把自己场上这张卡以外的1只（只能1只）怪兽解放（代价）
		◇1回合的发动次数不限（符合条件的前提下）
		◇回合结束时即为回合的结束阶段时
		 中文名: 艾尔的小剑士
		 日文名: アイルの小剣士
		 英文名: The Little Swordsman of Aile
		 卡片种类: 效果怪兽
		 卡片密码: 25109950
		 使用限制: 无限制
		 种族: 战士
		 属性: 水
		 星级: 3
		 攻击力: 800
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: ME，DL04，Booster07
		 效果: 每用自己场上存在的这张卡以外的1只怪兽做祭品，这张卡的攻击力在回合结束前加700。
		}
		*/
		Id:       430,
		Password: "25109950",
		Name:     "艾尔的小剑士",             // "The Little Swordsman of Aile"  "アイルの小剣士"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   3,
		La:      ygo.LA_Water,   // 水
		Lr:      ygo.LR_Warrior, // 战士
		Attack:  800,
		Defense: 1300,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*21*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 433
		 调整: [地雷蜘蛛]
		<地雷蜘蛛>
		[2010/09/10]

		●这张卡的攻击宣言时，投掷硬币猜正反。猜中的场合就那样攻击。猜错的场合自己的基本分失去一半攻击。
		◇永续效果（不进入连锁）
		◇玩家在选择这张卡进行攻击时（选择攻击对象之前），这个效果适用并处理
		◇没猜中攻击也正常进行
		◇攻击宣言后攻击卷返再选择攻击对象进行攻击的场合，不需要再投掷硬币猜正反
		 中文名: 地雷蜘蛛
		 日文名: 地雷蜘蛛
		 英文名: Jirai Gumo
		 卡片种类: 效果怪兽
		 卡片密码: 94773007
		 使用限制: 无限制
		 种族: 昆虫
		 属性: 地
		 星级: 4
		 攻击力: 2200
		 防御力: 100
		 罕见度: 平卡N
		 卡包: ME，BE02，DL04，Booster07
		 效果: 这张卡攻击宣言时，猜硬币的正反。猜中的话就继续攻击。猜不中的话自己的基本分减半再攻击。
		}
		*/
		Id:       433,
		Password: "94773007",
		Name:     "地雷蜘蛛",               // "Jirai Gumo"  "地雷蜘蛛"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   4,
		La:      ygo.LA_Earth,  // 地
		Lr:      ygo.LR_Insect, // 昆虫
		Attack:  2200,
		Defense: 100,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*22*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 441
		 调整: [空中的昆虫兵]
		<空の昆虫兵>
		[2010/09/10]

		●攻击风属性怪兽的场合，伤害步骤中这张卡的攻击力上升1000分。
		◇永续效果（不进入连锁）
		◇进入伤害步骤开始就适用
		 中文名: 空中的昆虫兵
		 日文名: 空の昆虫兵
		 英文名: Insect Soldiers of the Sky
		 卡片种类: 效果怪兽
		 卡片密码: 07019529
		 使用限制: 无限制
		 种族: 昆虫
		 属性: 风
		 星级: 3
		 攻击力: 1000
		 防御力: 800
		 罕见度: 平卡N
		 卡包: ME，DL04，Booster07
		 效果: 攻击风属性的怪兽时，伤害计算时攻击力上升1000。
		}
		*/
		Id:       441,
		Password: "07019529",
		Name:     "空中的昆虫兵",             // "Insect Soldiers of the Sky"  "空の昆虫兵"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   3,
		La:      ygo.LA_Wind,   // 风
		Lr:      ygo.LR_Insect, // 昆虫
		Attack:  1000,
		Defense: 800,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*23*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 444
		 调整: [寄居龙]
		<ヤドカリュー>
		[2010/09/10]

		●这张卡的表示形式从攻击表示变为表侧守备表示时，可以从自己的手卡把任意张数的卡放到卡组的最下方。
		◇诱发效果（进入连锁）
		◇任意发动
		◇自己没有手卡时不能发动
		◇效果处理时把任意张手卡放到卡组的最下方（这时必须至少有1张）
		◇顺序由这张卡的发动者决定
		◇原本持有者是对方的卡放到对方卡组的最下方
		 中文名: 寄居龙
		 日文名: ヤドカリュー
		 英文名: Yado Karu
		 卡片种类: 效果怪兽
		 卡片密码: 29380133
		 使用限制: 无限制
		 种族: 水
		 属性: 水
		 星级: 4
		 攻击力: 900
		 防御力: 1700
		 罕见度: 平卡N
		 卡包: ME，DL04，Booster07，TP01
		 效果: 这张卡的表示形式从攻击表示变成表侧守备表示时，可以从自己手卡把任意数量的卡回到卡组最下面。
		}
		*/
		Id:       444,
		Password: "29380133",
		Name:     "寄居龙",                // "Yado Karu"  "ヤドカリュー"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   4,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_None,  // 水
		Attack:  900,
		Defense: 1700,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*24*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 448
		 调整: [恶魔烹调师]
		<悪魔の調理師>
		[2010/09/10]

		（这张卡在规则上不当作「恶魔/デーモン」卡使用）
		●这张卡给予对方战斗伤害时，对方从卡组抽2张卡。
		◇诱发效果（进入连锁）
		◇强制发动
		 中文名: 恶魔烹调师
		 日文名: 悪魔の調理師
		 英文名: The Bistro Butcher
		 卡片种类: 效果怪兽
		 卡片密码: 71107816
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 1800
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: ME，BE02，DL04，Booster07
		 效果: （这张卡在规则上不当作「恶魔」卡使用）这张卡给予对方战斗伤害时，对方在卡组最上面抽2张卡。
		}
		*/
		Id:       448,
		Password: "71107816",
		Name:     "恶魔烹调师",              // "The Bistro Butcher"  "悪魔の調理師"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   4,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Attack:  1800,
		Defense: 1000,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*25*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 475
		 调整: [磁力指轮]
		<磁力の指輪>
		[2010/09/10]

		●只能给自己场上的怪兽装备。装备怪兽的攻击力与守备力下降500分。对方只能攻击装备了这张卡的怪兽。
		◇发动时选择自己场上1只表侧表示的怪兽（取对象）
		◇装备怪兽的控制权转移到对方场上的场合，这些效果不适用（继续装备着）
		 中文名: 磁力指轮
		 日文名: 磁力の指輪
		 英文名: Ring of Magnetism
		 卡片种类: 装备魔法
		 卡片密码: 20436034
		 使用限制: 无限制
		 罕见度: 银字R，平卡N
		 卡包: ME，BE02，DL04，Booster07
		 效果: 自己场上的怪兽才可以装备。装备的怪兽攻击力·守备力下降500。对方只能攻击装备了这张卡的怪兽。
		}
		*/
		Id:       475,
		Password: "20436034",
		Name:     "磁力指轮",            // "Ring of Magnetism"  "磁力の指輪"
		Lc:       ygo.LC_EquipMagic, // 装备魔法

		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*26*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 476
		 调整: [分担痛苦]
		<痛み分け>
		[2010/09/10]

		●把自己场上的1只怪兽解放发动。对方必须把1只怪兽解放。
		◇发动时把自己场上的1只怪兽解放（代价）
		◇效果处理时对方选择1只怪兽解放（不取对象）
		◇这个效果是强制要求对方把1只怪兽解放，不是以魔法效果解放对方1只怪兽
		◇对方场上的「荷鲁斯之黑炎龙 LV6/ホルスの黒炎竜 ＬＶ６」可以被解放
		◇对方场上只有1只怪兽，自己不能利用「灵魂交错/クロス·ソウル」的效果把对方的怪兽解放来发动（发动时确定效果处理时不适用）
		 中文名: 分担痛苦
		 日文名: 痛み分け
		 英文名: Share the Pain
		 卡片种类: 通常魔法
		 卡片密码: 56830749
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: ME，BE02，DL04，Booster07，DT02
		 效果: 自己场上的1只怪兽作祭品。对方选择对方场上1只怪做祭品。
		}
		*/
		Id:       476,
		Password: "56830749",
		Name:     "分担痛苦",               // "Share the Pain"  "痛み分け"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*27*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 477
		 调整: [兴奋剂]
		<ドーピング>
		[2010/09/10]

		●装备怪兽的攻击力上升700分。并且，在每次自己的准备阶段，装备怪兽的攻击力下降200分。
		◇发动时选择场上1只表侧表示的怪兽（取对象）
		◇准备阶段时攻击力下降效果进入连锁（咒文速度1）
		 中文名: 兴奋剂
		 日文名: ドーピング
		 英文名: Stim-Pack
		 卡片种类: 装备魔法
		 卡片密码: 83225447
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: ME，DL04，Booster07，TP14
		 效果: 装备怪兽的攻击力上升700。自己的每次的准备阶段装备怪兽的攻击力下降200。
		}
		*/
		Id:       477,
		Password: "83225447",
		Name:     "兴奋剂",             // "Stim-Pack"  "ドーピング"
		Lc:       ygo.LC_EquipMagic, // 装备魔法

		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*28*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 478
		 调整: [大风暴]
		<大嵐>
		[2010/09/10]

		●在场上存在的魔法·陷阱卡全部破坏。
		◇不取对象
		◇这张卡自身是在连锁处理结束后送去墓地，不是被破坏送去墓地
		◇同时复数张卡被破坏时，那些卡在墓地的放置顺序由控制者决定（破坏是同时破坏）
		 中文名: 大风暴
		 日文名: 大嵐
		 英文名: Heavy Storm
		 卡片种类: 通常魔法
		 卡片密码: 19613556
		 使用限制: 禁止卡
		 罕见度: 银字R，黄金GR，平卡N，面闪SR
		 卡包: ME，BE02，YSD01，YSD02，DL04，SD01，SD02，SD03，SD04，SD05，SD06，SD08，SD09，SD10，SD11，SD12，Booster07，GLD01，GS01，KA，PE，SK2，SD26
		 效果: 场上存在的魔法·陷阱卡全部破坏。
		}
		*/
		Id:       478,
		Password: "19613556",
		Name:     "大风暴",                // "Heavy Storm"  "大嵐"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				cs := ygo.NewCards(&pl.Szone.Cards, &tar.Szone.Cards)
				cs.ForEach(func(c *ygo.Card) bool {
					c.Dispatch(ygo.Destroy, ca)
					return true
				})
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*29*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 481
		 调整: [精灵之光]
		<エルフの光>
		[2010/09/10]

		●光属性怪兽才能装备。1只装备怪兽的攻击力上升400分。守备力下降200分。
		◇发动时选择1只符合条件的怪兽（取对象）
		◇效果处理时装备怪兽属性变为光属性以外的场合，这张卡送去墓地
		◇装备状态中对象怪兽变为光属性以外的场合，这张卡破坏
		 中文名: 精灵之光
		 日文名: エルフの光
		 英文名: Elf's Light
		 卡片种类: 装备魔法
		 卡片密码: 39897277
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: DL04，Booster02，Booster Chronicle，Booster R1，TP16
		 效果: 光属性的怪兽才能装备。装备的怪兽攻击力上升400，守备力下降200。
		}
		*/
		Id:       481,
		Password: "39897277",
		Name:     "精灵之光",            // "Elf's Light"  "エルフの光"
		Lc:       ygo.LC_EquipMagic, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterEquipMagic(func(c *ygo.Card) bool {
				return c.AttrIsLight()
			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() + 400)
				c.SetDefense(c.GetDefense() - 200)

			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() - 400)
				c.SetDefense(c.GetDefense() + 200)

			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*30*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 482
		 调整: [钢甲壳]
		<はがねの甲羅>
		[2010/09/10]

		●水属性怪兽才能装备。1只装备怪兽的攻击力上升400分。守备力下降200分。
		◇发动时选择1只符合条件的怪兽（取对象）
		◇效果处理时装备怪兽属性变为水属性以外的场合，这张卡送去墓地
		◇装备状态中对象怪兽变为水属性以外的场合，这张卡破坏
		 中文名: 钢甲壳
		 日文名: はがねの甲羅
		 英文名: Steel Shell
		 卡片种类: 装备魔法
		 卡片密码: 02370081
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: DL04，Booster02，Booster Chronicle，Booster R1，TP22
		 效果: 水属性怪兽才能装备。装备怪兽的攻击力上升400，守备力下降200。
		}
		*/
		Id:       482,
		Password: "02370081",
		Name:     "钢甲壳",             // "Steel Shell"  "はがねの甲羅"
		Lc:       ygo.LC_EquipMagic, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterEquipMagic(func(c *ygo.Card) bool {
				return c.AttrIsWater()
			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() + 400)
				c.SetDefense(c.GetDefense() - 200)

			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() - 400)
				c.SetDefense(c.GetDefense() + 200)

			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*31*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 483
		 调整: [天使的鲜血]
		<天使の生き血>
		[2010/09/10]

		●自己回复800基本分。
		◇「精灵之镜/精霊の鏡」可以对应这张卡的发动而发动
		 中文名: 天使的鲜血
		 日文名: 天使の生き血
		 英文名: Soul of the Pure
		 卡片种类: 通常魔法
		 卡片密码: 47852924
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: BE02，DL04，Booster02，Booster Chronicle，Booster R1
		 效果: 自己回复800点的基本分。
		}
		*/
		Id:       483,
		Password: "47852924",
		Name:     "天使的鲜血",              // "Soul of the Pure"  "天使の生き血"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				pl.ChangeHp(800)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*32*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 484
		 调整: [消除黑暗的光]
		<闇をかき消す光>
		[2010/09/10]

		●对方场上里侧表示存在的怪兽全部变为表侧表示。
		◇里侧守备变为表侧守备，里侧攻击变为表侧攻击
		 中文名: 消除黑暗的光
		 日文名: 闇をかき消す光
		 英文名: Dark-Piercing Light
		 卡片种类: 通常魔法
		 卡片密码: 45895206
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: BE02，DL04，Booster02，Booster Chronicle，Booster R1
		 效果: 对方场上里侧表示的怪兽全部表侧表示。
		}
		*/
		Id:       484,
		Password: "45895206",
		Name:     "消除黑暗的光",             // "Dark-Piercing Light"  "闇をかき消す光"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				tar.Mzone.ForEach(func(c *ygo.Card) bool {
					c.SetFaceUp()
					return true
				})
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*33*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 487
		 调整: [蓝色药剂]
		<ブルー·ポーション>
		[2010/09/10]

		●自己回复400基本分。
		◇「精灵之镜/精霊の鏡」可以对应这张卡的发动而发动
		 中文名: 蓝色药剂
		 日文名: ブルー·ポーション
		 英文名: Blue Medicine
		 卡片种类: 通常魔法
		 卡片密码: 20871001
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: DL04，Booster01，Booster R1，Booster Chronicle
		 效果: 自己回复400点的基本分。
		}
		*/
		Id:       487,
		Password: "20871001",
		Name:     "蓝色药剂",               // "Blue Medicine"  "ブルー·ポーション"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				pl.ChangeHp(400)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*34*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 488
		 调整: [雷鸣]
		<雷鳴>
		[2010/09/10]

		●给予对方基本分300分伤害。
		◇「精灵之镜/精霊の鏡」可以对应这张卡的发动而发动
		◇这张卡的读音和「雷-鸣/Rai-Mei」相同，所以在使用「心灵崩坏/マインドクラッシュ」等宣言卡名时需要同时说明是宣言魔法卡的「雷鳴」还是效果怪兽的「Rai-Mei」
		 中文名: 雷鸣
		 日文名: 雷鳴
		 英文名: Raimei
		 卡片种类: 通常魔法
		 卡片密码: 56260110
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: DL04，Booster01，Booster Chronicle，Booster R1，TP04
		 效果: 给与对方基本分300分伤害。
		}
		*/
		Id:       488,
		Password: "56260110",
		Name:     "雷鸣",                 // "Raimei"  "雷鳴"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				tar.ChangeHp(-300)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*35*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 489
		 调整: [灼热之枪]
		<灼熱の槍>
		[2010/09/10]

		●炎属性怪兽才能装备。1只装备怪兽的攻击力上升400分，守备力下降200分。
		◇发动时选择1只符合条件的怪兽（取对象）
		◇效果处理时装备怪兽属性变为炎属性以外的场合，这张卡送去墓地
		◇装备状态中对象怪兽变为炎属性以外的场合，这张卡破坏
		 中文名: 灼热之枪
		 日文名: 灼熱の槍
		 英文名: Burning Spear
		 卡片种类: 装备魔法
		 卡片密码: 18937875
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: DL04，Booster02，Booster Chronicle，Booster R1，TP13
		 效果: 炎属性的怪兽才能装备。装备的怪兽攻击力上升400，守备力下降200。
		}
		*/
		Id:       489,
		Password: "18937875",
		Name:     "灼热之枪",            // "Burning Spear"  "灼熱の槍"
		Lc:       ygo.LC_EquipMagic, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterEquipMagic(func(c *ygo.Card) bool {
				return c.AttrIsFire()
			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() + 400)
				c.SetDefense(c.GetDefense() - 200)

			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() - 400)
				c.SetDefense(c.GetDefense() + 200)

			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*36*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 490
		 调整: [突风之扇]
		<突風の扇>
		[2010/09/10]

		●风属性怪兽才能装备。1只装备怪兽的攻击力上升400分。守备力下降200分。
		◇发动时选择1只符合条件的怪兽（取对象）
		◇效果处理时装备怪兽属性变为风属性以外的场合，这张卡送去墓地
		◇装备状态中对象怪兽变为风属性以外的场合，这张卡破坏
		 中文名: 突风之扇
		 日文名: 突風の扇
		 英文名: Gust Fan
		 卡片种类: 装备魔法
		 卡片密码: 55321970
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: DL04，Booster02，Booster Chronicle，Booster R1
		 效果: 风属性的怪兽才能装备。装备的怪兽攻击力上升400，守备力下降200。
		}
		*/
		Id:       490,
		Password: "55321970",
		Name:     "突风之扇",            // "Gust Fan"  "突風の扇"
		Lc:       ygo.LC_EquipMagic, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterEquipMagic(func(c *ygo.Card) bool {
				return c.AttrIsWind()
			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() + 400)
				c.SetDefense(c.GetDefense() - 200)

			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() - 400)
				c.SetDefense(c.GetDefense() + 200)

			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*37*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 491
		 调整: [邪恶蠕虫兽]
		<邪悪なるワーム·ビースト>
		[2010/09/10]

		（这张卡在规则上也当作「异虫」卡使用）
		●自己回合的结束阶段时，表侧表示在场上存在的这张卡回到持有者的手卡。
		◇诱发效果（进入连锁）
		◇强制发动
		◇自己的结束阶段时「亚空间物质传送装置/亜空間物質転送装置」让这张卡回到场上的场合，这个效果也必须发动
		 中文名: 邪恶蠕虫兽
		 日文名: 邪悪なるワーム·ビースト
		 英文名: The Wicked Worm Beast
		 卡片种类: 效果怪兽
		 卡片密码: 06285791
		 使用限制: 无限制
		 种族: 兽
		 属性: 地
		 星级: 3
		 攻击力: 1400
		 防御力: 700
		 罕见度: 平卡N
		 卡包: BE02，EX-R(EX)，DL04，Booster04，Booster Chronicle，Booster R2，KA，SK2，GLD04
		 效果: （这张卡在规则上也当作「异虫」卡使用）自己的回合的结束阶段时，场上表侧表示存在的这张卡回到持有者的手卡。
		}
		*/
		Id:       491,
		Password: "06285791",
		Name:     "邪恶蠕虫兽",              // "The Wicked Worm Beast"  "邪悪なるワーム·ビースト"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   3,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Attack:  1400,
		Defense: 700,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*38*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 494
		 调整: [森之住人 乌丹]<森の住人 ウダン>

		●场上每存在一只打开表示的植物族怪兽，这张卡的攻击力上升100。
		◇永续效果
		◇数双方场上的植物族怪兽
		◇[天邪鬼的诅咒/あまのじゃくの呪い]効果适用
		 中文名: 森之住人 乌丹
		 日文名: 森の住人 ウダン
		 英文名: Wodan the Resident of the Forest
		 卡片种类: 效果怪兽
		 卡片密码: 42883273
		 使用限制: 无限制
		 种族: 战士
		 属性: 地
		 星级: 3
		 攻击力: 900
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: DL04，Booster03，Booster Chronicle，Booster R2
		 效果: 场上每存在1只表侧表示的植物族怪兽，这张卡的攻击力上升100。
		}
		*/
		Id:       494,
		Password: "42883273",
		Name:     "森之住人 乌丹",            // "Wodan the Resident of the Forest"  "森の住人 ウダン"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   3,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Attack:  900,
		Defense: 1200,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*39*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 497
		 调整: [蟑螂骑士]
		<コカローチ·ナイト>
		[2010/09/10]

		●这张卡送去墓地时，这张卡回到卡组的最上方。
		◇诱发效果（进入连锁）
		◇强制发动
		◇取对象效果
		◇伤害步骤中也发动
		◇作为怪兽装备、被「升天之角笛/昇天の角笛」「神之宣告/神の宣告」无效召唤而破坏、结束阶段时手卡调整、里侧表示被送去墓地的场合，这个效果发动
		 中文名: 蟑螂骑士
		 日文名: コカローチ·ナイト
		 英文名: Cockroach Knight
		 卡片种类: 效果怪兽
		 卡片密码: 33413638
		 使用限制: 无限制
		 种族: 昆虫
		 属性: 地
		 星级: 3
		 攻击力: 800
		 防御力: 900
		 罕见度: 平卡N
		 卡包: DL04，Booster04，Booster Chronicle，Booster R2
		 效果: 这张卡送去墓地时，这张卡回到卡组最上面。
		}
		*/
		Id:       497,
		Password: "33413638",
		Name:     "蟑螂骑士",               // "Cockroach Knight"  "コカローチ·ナイト"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   3,
		La:      ygo.LA_Earth,  // 地
		Lr:      ygo.LR_Insect, // 昆虫
		Attack:  800,
		Defense: 900,
		Initialize: func(ca *ygo.Card) bool {
			ca.AddEvent(ygo.InGrave, func() {
				ca.ToDeck()
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*40*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 499
		 调整: [巡逻机器人]<パトロール·ロボ>

		●只要这张卡在场上打开表示存在，可以在自己的准备流程看对方盖放的一张卡。(看完后盖回去，不诱发效果)
		◇诱发效果，任意发动，取对象
		◇一回合只能用一次
		◇里侧表示的怪/魔/陷都可以成为此效果的对象
		◇效果处理时对象变成表侧的场合，效果不发
		★效果处理时此卡不在场上表侧表示存在的场合，效果是否不发，调整中
		 中文名: 巡逻机器人
		 日文名: パトロール·ロボ
		 英文名: Patrol Robo
		 卡片种类: 效果怪兽
		 卡片密码: 76775123
		 使用限制: 无限制
		 种族: 机械
		 属性: 地
		 星级: 3
		 攻击力: 1100
		 防御力: 900
		 罕见度: 平卡N
		 卡包: DL04，Booster03，Booster Chronicle，Booster R2，TP02
		 效果: 只要这张卡在场上表侧表示存在，自己的准备阶段可以把对方场上盖放的1张卡确认。
		}
		*/
		Id:       499,
		Password: "76775123",
		Name:     "巡逻机器人",              // "Patrol Robo"  "パトロール·ロボ"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   3,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Machine, // 机械
		Attack:  1100,
		Defense: 900,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*41*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 501
		 调整: [勇气之沙漏]
		<勇気の砂時計>
		[2010/09/10]

		●召唤·反转召唤的场合，直到下次的自己回合的结束阶段为止，这张卡的原本的攻击力·原本的守备力减半。那之后，这张卡的原本的攻击力与原本的守备力变成两倍。
		◇诱发效果（进入连锁）
		◇强制发动
		◇改变的数值是原本数值
		◇这个效果适用后，给这张卡装备「巨大化/巨大化」的场合，以这个效果改变后的原本数值进行参照
		 中文名: 勇气之沙漏
		 日文名: 勇気の砂時計
		 英文名: Hourglass of Courage
		 卡片种类: 效果怪兽
		 卡片密码: 43530283
		 使用限制: 无限制
		 种族: 天使
		 属性: 光
		 星级: 4
		 攻击力: 1100
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: DL04，Booster03，Booster Chronicle，Booster R2，STOR(703)，TP16
		 效果: 召唤·反转召唤的场合，直到下次自己的回合的结束阶段这张卡的原来的攻击力·守备力减半，之后这张卡的原来的攻击力·守备力变成2倍。
		}
		*/
		Id:       501,
		Password: "43530283",
		Name:     "勇气之沙漏",              // "Hourglass of Courage"  "勇気の砂時計"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   4,
		La:      ygo.LA_Light, // 光
		Lr:      ygo.LR_Angel, // 天使
		Attack:  1100,
		Defense: 1200,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*42*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 504
		 调整: [天使的施舍]
		<天使の施し>
		[2010/09/10]

		●从卡组抽3张卡，那之后从手卡丢弃2张卡。
		◇效果处理时先抽卡，之后再丢弃2张卡
		◇卡组中不足3张卡的场合，不能发动
		◇「精灵之镜/精霊の鏡」可以对应这张卡的发动而发动
		◇「死之卡组破坏病毒/死のデッキ破壊ウイルス」「魔之卡组破坏病毒/魔のデッキ破壊ウイルス」的效果在这个效果处理完丢卡后再适用
		 中文名: 天使的施舍
		 日文名: 天使の施し
		 英文名: Graceful Charity
		 卡片种类: 通常魔法
		 卡片密码: 79571449
		 使用限制: 禁止卡
		 罕见度: 面闪SR，平卡N，银字R
		 卡包: BE02，DL04，Booster04，Booster Chronicle，Booster R2，YU，KA，SK2
		 效果: 从卡组抽3张卡，之后从手卡选2张丢弃。
		}
		*/
		Id:       504,
		Password: "79571449",
		Name:     "天使的施舍",              // "Graceful Charity"  "天使の施し"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				ca.GetSummoner().ActionDraw(3)
				for i := 0; i != 2; i++ {
					c := pl.SelectForWarn(pl.Hand, func(c0 *ygo.Card) bool {
						return c0.IsMonster()
					})
					if c == nil {
						c = pl.Hand.EndPop()
					}
					c.Dispatch(ygo.Discard, ca)
				}
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*43*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 505
		 调整: [来自墓场的呼声]
		<墓場からの呼び声>
		[2010/09/10]

		●对方发动「死者苏生」时可以发动。那张「死者苏生」的效果无效。
		◇必须直接对应「死者苏生/死者蘇生」的发动进行连锁发动
		◇只无效效果，不无效发动
		 中文名: 来自墓场的呼声
		 日文名: 墓場からの呼び声
		 英文名: Call of the Grave
		 卡片种类: 通常陷阱
		 卡片密码: 16970158
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: DL04，Booster04，TP05，Booster R2，Booster Chronicle
		 效果: 对方「死者苏生」发动时，发动效果无效。
		}
		*/
		Id:       505,
		Password: "16970158",
		Name:     "来自墓场的呼声",           // "Call of the Grave"  "墓場からの呼び声"
		Lc:       ygo.LC_OrdinaryTrap, // 通常陷阱

		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*44*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 509
		 调整: [心眼之女神]
		<心眼の女神>
		[2010/09/10]

		●这张卡可以代替1只融合素材怪兽。那时，其他的融合素材怪兽必须是正规物。
		◇无效果分类
		◇可以被「技能抽取/スキルドレイン」「星骸龙/デブリ·ドラゴン」「冥界的魔王 哈·迪斯/冥界の魔王 ハ·デス」等的效果无效
		◇只能代替融合怪兽上写全卡名的融合素材
		◇只能在进行融合召唤时代替融合素材怪兽
		◇不能从卡组、除外状态代替融合素材怪兽
		◇可以从手卡、场上（不论表里）、墓地代替融合素材怪兽
		 中文名: 心眼之女神
		 日文名: 心眼の女神
		 英文名: Goddess with the Third Eye
		 卡片种类: 效果怪兽
		 卡片密码: 53493204
		 使用限制: 无限制
		 种族: 天使
		 属性: 光
		 星级: 4
		 攻击力: 1200
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: BE02，DL04，Booster06，Booster Chronicle，Booster R3，JY，GLD04
		 效果: 这张卡可以代替融合怪兽素材的其中1只来融合。这个时候，其他的融合素材必须是指定的融合素材。
		}
		*/
		Id:       509,
		Password: "53493204",
		Name:     "心眼之女神",              // "Goddess with the Third Eye"  "心眼の女神"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   4,
		La:      ygo.LA_Light, // 光
		Lr:      ygo.LR_Angel, // 天使
		Attack:  1200,
		Defense: 1000,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*45*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 511
		 调整: [沼地的魔兽王]
		<沼地の魔獣王>
		[2010/09/10]

		●这张卡可以代替1只融合素材怪兽。那时，其他的融合素材怪兽必须是正规物。
		◇无效果分类
		◇可以被「技能抽取/スキルドレイン」「星骸龙/デブリ·ドラゴン」「冥界的魔王 哈·迪斯/冥界の魔王 ハ·デス」等的效果无效
		◇只能代替融合怪兽上写全卡名的融合素材
		◇只能在进行融合召唤时代替融合素材怪兽
		◇不能从卡组、除外状态代替融合素材怪兽
		◇可以从手卡、场上（不论表里）、墓地代替融合素材怪兽
		 中文名: 沼地的魔兽王
		 日文名: 沼地の魔獣王
		 英文名: Beastking of the Swamps
		 卡片种类: 效果怪兽
		 卡片密码: 99426834
		 使用限制: 无限制
		 种族: 水
		 属性: 水
		 星级: 4
		 攻击力: 1000
		 防御力: 1100
		 罕见度: 平卡N
		 卡包: DL04，Booster06，Booster Chronicle，Booster R3，GLD04
		 效果: 这张卡可以代替融合怪兽素材的其中1只来融合。这个时候，其他的融合素材必须是指定的融合素材。
		}
		*/
		Id:       511,
		Password: "99426834",
		Name:     "沼地的魔兽王",             // "Beastking of the Swamps"  "沼地の魔獣王"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   4,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_None,  // 水
		Attack:  1000,
		Defense: 1100,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*46*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 512
		 调整: [破坏神 瓦沙克]
		<破壊神 ヴァサーゴ>
		[2010/09/10]

		●这张卡可以代替1只融合素材怪兽。那时，其他的融合素材怪兽必须是正规物。
		◇无效果分类
		◇可以被「技能抽取/スキルドレイン」「星骸龙/デブリ·ドラゴン」「冥界的魔王 哈·迪斯/冥界の魔王 ハ·デス」等的效果无效
		◇只能代替融合怪兽上写全卡名的融合素材
		◇只能在进行融合召唤时代替融合素材怪兽
		◇不能从卡组、除外状态代替融合素材怪兽
		◇可以从手卡、场上（不论表里）、墓地代替融合素材怪兽
		 中文名: 破坏神 瓦沙克
		 日文名: 破壊神 ヴァサーゴ
		 英文名: Versago the Destroyer
		 卡片种类: 效果怪兽
		 卡片密码: 50259460
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 3
		 攻击力: 1100
		 防御力: 900
		 罕见度: 平卡N
		 卡包: DL04，Booster06，Booster Chronicle，Booster R3，KA，GLD04
		 效果: 这张卡可以代替融合怪兽素材的其中1只来融合。这个时候，其他的融合素材必须是指定的融合素材。
		}
		*/
		Id:       512,
		Password: "50259460",
		Name:     "破坏神 瓦沙克",            // "Versago the Destroyer"  "破壊神 ヴァサーゴ"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   3,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Attack:  1100,
		Defense: 900,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*47*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 513
		 调整: [眼球怪]
		<モンスター·アイ>
		[2010/09/10]

		●支付1000基本分。从自己墓地把1张「融合」魔法卡返回手卡。
		◇启动效果（进入连锁）
		◇取对象
		◇发动次数不限（符合条件的前提下，不得空发）
		 中文名: 眼球怪
		 日文名: モンスター·アイ
		 英文名: Monster Eye
		 卡片种类: 效果怪兽
		 卡片密码: 84133008
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 1
		 攻击力: 250
		 防御力: 350
		 罕见度: 平卡N
		 卡包: DL04，Booster06，Booster Chronicle，Booster R3，TP15
		 效果: 支付1000分。自己的墓地1张「融合」魔法卡回到手卡。
		}
		*/
		Id:       513,
		Password: "84133008",
		Name:     "眼球怪",                // "Monster Eye"  "モンスター·アイ"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   1,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Attack:  250,
		Defense: 350,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*48*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 514
		 调整: [机械王]
		<機械王>
		[2010/09/10]

		●场上每有1只表侧表示存在的机械族怪兽，这张卡的攻击力上升100分。
		◇永续效果（不进入连锁）
		◇这张卡自身在效果适用范围内
		 中文名: 机械王
		 日文名: 機械王
		 英文名: Machine King
		 卡片种类: 效果怪兽
		 卡片密码: 46700124
		 使用限制: 无限制
		 种族: 机械
		 属性: 地
		 星级: 6
		 攻击力: 2200
		 防御力: 2000
		 罕见度: 平卡N
		 卡包: BE02，DL04，Booster05，Booster Chronicle，Booster R3
		 效果: 每1只在场上表侧表示存在的机械族怪兽，这张卡的攻击力上升100。
		}
		*/
		Id:       514,
		Password: "46700124",
		Name:     "机械王",                // "Machine King"  "機械王"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   6,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Machine, // 机械
		Attack:  2200,
		Defense: 2000,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*49*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 515
		 调整: [恶魔科学怪人]
		<デビル·フランケン>
		[15/02/02]

		（这张卡在规则上不当作「恶魔/デーモン」卡使用）
		●①：支付5000基本分才能发动。从额外卡组把1只融合怪兽攻击表示特殊召唤。
		◇起动效果，开连锁，不取对象
		◇支付5000基本分是效果发动COST
		◇效果处理时从自己的额外卡组把1只融合怪兽在自己场上表侧攻击表示特殊召唤，这个特殊召唤不是融合召唤、不解除苏生限制
		 中文名: 恶魔科学怪人
		 日文名: デビル·フランケン
		 英文名: Cyber-Stein
		 卡片种类: 效果怪兽
		 卡片密码: 69015963
		 使用限制: 无限制
		 种族: 机械
		 属性: 暗
		 星级: 2
		 攻击力: 700
		 防御力: 500
		 罕见度: 平卡N，银字R
		 卡包: BE02，DL04，Booster06，Booster Chronicle，Booster R3，KA，SK2，AT06
		 效果: （这张卡在规则上不当作「恶魔」卡使用）①：支付5000基本分才能发动。从额外卡组把1只融合怪兽攻击表示特殊召唤。
		}
		*/
		Id:       515,
		Password: "69015963",
		Name:     "恶魔科学怪人",             // "Cyber-Stein"  "デビル·フランケン"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   2,
		La:      ygo.LA_Dark,    // 暗
		Lr:      ygo.LR_Machine, // 机械
		Attack:  700,
		Defense: 500,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*50*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 516
		 调整: [针球]
		<ニードル·ボール>
		[14/08/07]

		●反转：可以支付2000基本分，给予对方的基本分1000分伤害。
		◇诱发效果，任意发动，开连锁，不取对象
		◇支付2000基本分是效果发动COST
		 中文名: 针球
		 日文名: ニードル·ボール
		 英文名: Needle Ball
		 卡片种类: 效果怪兽
		 卡片密码: 94230224
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 2
		 攻击力: 750
		 防御力: 700
		 罕见度: 平卡N
		 卡包: DL04，Booster05，Booster Chronicle，Booster R3，STOR(703)
		 效果: 反转：可以支付2000分，对方受到1000分的伤害。
		}
		*/
		Id:       516,
		Password: "94230224",
		Name:     "针球",                 // "Needle Ball"  "ニードル·ボール"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   2,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Attack:  750,
		Defense: 700,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*51*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 517
		 调整: [杀龙者]
		<竜殺者>
		[2010/09/10]

		●这张卡在场上召唤·反转召唤时，破坏1只表侧表示的龙族怪兽。
		◇诱发效果（进入连锁）
		◇强制发动
		◇发动时选择场上表侧表示存在的1只龙族怪兽（取对象）
		◇效果处理时对象怪兽的种族变为龙族以外的场合，这个效果不适用
		 中文名: 杀龙者
		 日文名: 竜殺者
		 英文名: Dragon Seeker
		 卡片种类: 效果怪兽
		 卡片密码: 28563545
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 6
		 攻击力: 2000
		 防御力: 2100
		 罕见度: 平卡N
		 卡包: BE02，DL04，Booster05，Booster Chronicle，Booster R3
		 效果: 这张卡在场上召唤·反转召唤时，破坏表侧表示的1只龙族怪兽。
		}
		*/
		Id:       517,
		Password: "28563545",
		Name:     "杀龙者",                // "Dragon Seeker"  "竜殺者"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   6,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Attack:  2000,
		Defense: 2100,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*52*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 518
		 调整: [针虫]
		<ニードルワーム>
		[14/08/07]

		（这张卡在规则上也当作「异虫」卡使用）
		●反转：从对方的卡组把最上方的5张卡丢弃去墓地。
		◇诱发效果，强制发动，开连锁，不取对象
		◇对方卡组不足5张卡的场合这个效果也发动
		◇效果处理时对方卡组的卡不足4张的场合，把那些卡剩余的卡全部丢弃去墓地
		 中文名: 针虫
		 日文名: ニードルワーム
		 英文名: Needle Worm
		 卡片种类: 效果怪兽
		 卡片密码: 81843628
		 使用限制: 无限制
		 种族: 昆虫
		 属性: 地
		 星级: 2
		 攻击力: 750
		 防御力: 600
		 罕见度: 平卡N，银字R
		 卡包: BE02，DL04，Booster05，Booster Chronicle，Booster R3，GLD01，KA
		 效果: （这张卡在规则上也当作「异虫」卡使用）反转：对方卡组最上面的5张卡送去墓地。
		}
		*/
		Id:       518,
		Password: "81843628",
		Name:     "针虫",                 // "Needle Worm"  "ニードルワーム"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   2,
		La:      ygo.LA_Earth,  // 地
		Lr:      ygo.LR_Insect, // 昆虫
		Attack:  750,
		Defense: 600,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*53*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 519
		 调整: [二重身]
		<ドッペルゲンガー>
		[2010/09/10]

		●反转：必须选择场上覆盖的2张魔法或陷阱卡，破坏那些。
		◇诱发效果，强制发动，开连锁，发动时选择场上2张覆盖的魔法或陷阱卡为对象
		◇场上可以作为对象的卡不足2张的场合，这个效果不发动
		◇效果处理时其中之一对象为表侧表示的场合，剩下的那张破坏
		◇效果处理时其中之一对象不在场上存在的场合，剩下的那张破坏
		 中文名: 二重身
		 日文名: ドッペルゲンガー
		 英文名: Greenkappa
		 卡片种类: 效果怪兽
		 卡片密码: 61831093
		 使用限制: 无限制
		 种族: 战士
		 属性: 暗
		 星级: 3
		 攻击力: 650
		 防御力: 900
		 罕见度: 平卡N
		 卡包: BE02，DL04，Booster06，Booster Chronicle，Booster R3，DT12，TU06
		 效果: 反转：选择场上盖放的2张魔法·陷阱卡破坏。
		}
		*/
		Id:       519,
		Password: "61831093",
		Name:     "二重身",                // "Greenkappa"  "ドッペルゲンガー"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   3,
		La:      ygo.LA_Dark,    // 暗
		Lr:      ygo.LR_Warrior, // 战士
		Attack:  650,
		Defense: 900,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*54*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 520
		 调整: [变形壶]
		<メタモルポット>
		[14/08/07]

		●反转：双方的手卡全部丢弃。那之后，各自从自己的卡组抽5张卡。
		◇诱发效果，强制发动，开连锁，不取对象
		◇其中一方卡组的卡不足5张的场合也发动，效果处理时那一方决斗败北
		◇「神殿守卫者/神殿を守る者」在场上存在时，这张卡也可以反转召唤及发动效果，这个场合对方只丢卡不抽卡
		◇双方手卡为0这个效果发动的场合，抽卡效果正常处理
		 中文名: 变形壶
		 日文名: メタモルポット
		 英文名: Morphing Jar
		 卡片种类: 效果怪兽
		 卡片密码: 33508719
		 使用限制: 限制卡
		 种族: 岩石
		 属性: 地
		 星级: 2
		 攻击力: 700
		 防御力: 600
		 罕见度: 银字R，平卡N，面闪SR，黄金GR
		 卡包: BE02，DL04，Booster05，Booster Chronicle，Booster R3，SD13，GS02，DB12，GLD04
		 效果: 反转：双方手卡全部丢弃。那之后，双方各自从自己卡组抽5张卡。
		}
		*/
		Id:       520,
		Password: "33508719",
		Name:     "变形壶",                // "Morphing Jar"  "メタモルポット"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   2,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Rock,  // 岩石
		Attack:  700,
		Defense: 600,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*55*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 521
		 调整: [企鹅士兵]
		<ペンギン·ソルジャー>
		[14/08/07]

		●反转：可以让场上存在的最多2只怪兽回到持有者的手卡。
		◇诱发效果，任意发动，开连锁，发动时取1~2张符合条件的卡为对象
		◇里侧表示的此卡因战斗而反转、要发动这个效果时，不能把已经被战斗破坏的怪兽作为对象
		◇发动时选择2只对象，效果处理时对象卡只剩下1只的场合，剩下的那只回到持有者手卡
		 中文名: 企鹅士兵
		 日文名: ペンギン·ソルジャー
		 英文名: Penguin Soldier
		 卡片种类: 效果怪兽
		 卡片密码: 93920745
		 使用限制: 无限制
		 种族: 水
		 属性: 水
		 星级: 2
		 攻击力: 750
		 防御力: 500
		 罕见度: 面闪SR，平卡N，银字R
		 卡包: BE02，DL04，Booster06，Booster Chronicle，Booster R3，YSD04，YU，JY，DT12，SD23
		 效果: 反转：可以选择场上最多2只怪兽回到持有者手卡。
		}
		*/
		Id:       521,
		Password: "93920745",
		Name:     "企鹅士兵",               // "Penguin Soldier"  "ペンギン·ソルジャー"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   2,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_None,  // 水
		Attack:  750,
		Defense: 500,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterFlip(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				for i := 0; i != 2; i++ {
					if c := pl.SelectForWarn(tar.Mzone, pl.Mzone); c != nil {
						c.ToHand()
					}
				}
			})
			return true
		}, // 初始

		IsValid: true,
	})

	/*56*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 522
		 中文名: 水陆两用战斗艇
		 日文名: 水陸両用バグロス
		 英文名: Amphibious Bugroth
		 卡片种类: 融合怪兽
		 卡片密码: 40173854
		 使用限制: 无限制
		 种族: 水
		 属性: 水
		 星级: 5
		 攻击力: 1850
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: DL04，Booster06，Booster Chronicle，Booster R3
		 效果: 融合：「陆战型战斗艇」＋「守卫海洋的战士」。
		}
		*/
		Id:       522,
		Password: "40173854",
		Name:     "水陆两用战斗艇",            // "Amphibious Bugroth"  "水陸両用バグロス"
		Lc:       ygo.LC_FusionMonster, // 融合怪兽

		Level:   5,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_None,  // 水
		Attack:  1850,
		Defense: 1300,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterFusionMonster("陆战型战斗艇", "守卫海洋的战士")
			return true
		}, // 初始
		IsValid: true,
	})

	/*57*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 523
		 调整: [幻想绵羊]
		<イリュージョン·シープ>
		[2010/09/10]

		●这张卡可以代替1只融合素材怪兽。那时，其他的融合素材怪兽必须是正规物。
		◇无效果分类
		◇可以被「技能抽取/スキルドレイン」「星骸龙/デブリ·ドラゴン」「冥界的魔王 哈·迪斯/冥界の魔王 ハ·デス」等的效果无效
		◇只能代替融合怪兽上写全卡名的融合素材
		◇只能在进行融合召唤时代替融合素材怪兽
		◇不能从卡组、除外状态代替融合素材怪兽
		◇可以从手卡、场上（不论表里）、墓地代替融合素材怪兽
		 中文名: 幻想绵羊
		 日文名: イリュージョン·シープ
		 英文名: Mystical Sheep #1
		 卡片种类: 效果怪兽
		 卡片密码: 30451366
		 使用限制: 无限制
		 种族: 兽
		 属性: 地
		 星级: 3
		 攻击力: 1150
		 防御力: 900
		 罕见度: 平卡N
		 卡包: DL04，Booster06，Booster Chronicle，Booster R3，TP08
		 效果: 这张卡可以代替融合怪兽素材的其中1只来融合。这个时候，其他的融合素材必须是指定的融合素材。
		}
		*/
		Id:       523,
		Password: "30451366",
		Name:     "幻想绵羊",               // "Mystical Sheep #1"  "イリュージョン·シープ"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   3,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Attack:  1150,
		Defense: 900,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*58*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 524
		 调整: [王宫的通告]
		<王宮のお触れ>
		[2010/09/10]

		●只要这张卡在场上表侧表示存在，这张卡以外的场上的陷阱卡的效果无效。
		◇只无效效果，不无效发动（不进入连锁）
		◇只要是在场上发动的陷阱卡，不论效果处理时在哪里，效果都无效化
		◇在场上以外发动的陷阱效果不会被无效
		 中文名: 王宫的通告
		 日文名: 王宮のお触れ
		 英文名: Royal Decree
		 卡片种类: 永续陷阱
		 卡片密码: 51452091
		 使用限制: 无限制
		 罕见度: 面闪SR，银字R，黄金GR，平卡N，平爆NPR
		 卡包: BE02，DL04，SD05，Booster05，Booster Chronicle，Booster R3，GS01，SY2，SPTR
		 效果: 只要这张卡在场上表侧表示存在，这张卡以外的场上的陷阱卡的效果无效。
		}
		*/
		Id:       524,
		Password: "51452091",
		Name:     "王宫的通告",             // "Royal Decree"  "王宮のお触れ"
		Lc:       ygo.LC_SustainsTrap, // 永续陷阱

		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*59*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 525
		 调整: [魔女狩猎]
		<魔女狩り>
		[2010/09/10]

		●在场上表侧表示存在的魔法师族怪兽全部破坏。
		◇不取对象
		 中文名: 魔女狩猎
		 日文名: 魔女狩り
		 英文名: Last Day of Witch
		 卡片种类: 通常魔法
		 卡片密码: 90330453
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: DL04，Booster06，Booster Chronicle，Booster R3
		 效果: 场上表侧表示的魔法师族的怪兽全部破坏。
		}
		*/
		Id:       525,
		Password: "90330453",
		Name:     "魔女狩猎",               // "Last Day of Witch"  "魔女狩り"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				cs := ygo.NewCards(&pl.Mzone.Cards, &tar.Mzone.Cards)
				cs.ForEach(func(c *ygo.Card) bool {
					if c.IsFaceUp() && c.RaceIsSpellCaster() {
						c.Dispatch(ygo.Destroy, ca)
					}
					return true
				})
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*60*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 526
		 调整: [恶魔祓除]
		<悪魔払い>
		[2010/09/10]

		（这张卡在规则上不当作「恶魔/デーモン」卡使用）
		●在场上表侧表示存在的恶魔族怪兽全部破坏。
		◇不取对象
		 中文名: 恶魔祓除
		 日文名: 悪魔払い
		 英文名: Exile of the Wicked
		 卡片种类: 通常魔法
		 卡片密码: 26725158
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: DL04，Booster06，Booster Chronicle，Booster R3，GLD02，TP23
		 效果: （这张卡在规则上不当作「恶魔」卡使用）场上的恶魔族怪兽全部破坏。
		}
		*/
		Id:       526,
		Password: "26725158",
		Name:     "恶魔祓除",               // "Exile of the Wicked"  "悪魔払い"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				cs := ygo.NewCards(&pl.Mzone.Cards, &tar.Mzone.Cards)
				cs.ForEach(func(c *ygo.Card) bool {
					if c.IsFaceUp() && c.RaceIsFiend() {
						c.Dispatch(ygo.Destroy, ca)
					}
					return true
				})
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*61*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 527
		 调整: [魔力之棘]
		<魔力の棘>
		[2010/09/10]

		●对方的手卡被丢弃去墓地时，每丢弃1张卡给予对方500分伤害。
		◇伤害效果不进入连锁
		◇这张卡在场上有2张以上效果适用的场合，效果叠加
		◇原本自己的卡从对方手卡被丢弃、原本对方的卡从自己手卡被丢弃这个效果都不适用
		◇一次丢弃多张的场合，伤害数值为那个被丢弃数量×500
		◇从手卡送去墓地、被破坏都不在效果适用范围内
		◇对方场上「魔力之棘/魔力の棘」的效果适用中，我方回合自己发动「天使的施舍/天使の施し」从卡组抽3张卡，丢弃2张卡。「天使的施舍/天使の施し」效果处理完毕后自己凑齐了「艾克佐迪亚/エクゾディア」5张部件满足胜利条件，而由于「魔力之棘/魔力の棘」的效果自己受到1000分伤害基本分变为0对方也满足胜利条件的场合，以回合玩家优先来处理。（即实际先适用「被封印的艾克佐迪亚/封印されしエクゾディア」的效果自己获得决斗胜利）
		 中文名: 魔力之棘
		 日文名: 魔力の棘
		 英文名: Magical Thorn
		 卡片种类: 永续陷阱
		 卡片密码: 53119267
		 使用限制: 无限制
		 罕见度: 银字R，平卡N
		 卡包: BE02，DL04，Booster06，Booster Chronicle，Booster R3
		 效果: 对方的手卡丢弃去墓地时，每丢弃1张卡就给予对方基本分500分的伤害。
		}
		*/
		Id:       527,
		Password: "53119267",
		Name:     "魔力之棘",              // "Magical Thorn"  "魔力の棘"
		Lc:       ygo.LC_SustainsTrap, // 永续陷阱

		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*62*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 528
		 调整: [革命]
		<革命>
		[2010/09/10]

		●给予对方基本分对方的手卡的数量×200分的伤害。
		◇「精灵之镜/精霊の鏡」可以对应这张卡的发动而发动
		◇对方没有手卡时不能发动
		◇效果处理时计算对方的手卡数量
		 中文名: 革命
		 日文名: 革命
		 英文名: Restructer Revolution
		 卡片种类: 通常魔法
		 卡片密码: 99518961
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: BE02，DL04，Booster06，Booster Chronicle，Booster R3，SY2
		 效果: 对方受到（对方的手卡数×200）的伤害。
		}
		*/
		Id:       528,
		Password: "99518961",
		Name:     "革命",                 // "Restructer Revolution"  "革命"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				tar.ChangeHp(-200 * tar.Hand.Len())
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*63*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 529
		 调整: [融合贤者]
		<融合賢者>
		[2010/09/10]

		●从自己的卡组把1张「融合」魔法卡加入手卡。那之后卡组洗切。
		◇遵守不能空发原则
		◇效果处理时卡组内没有符合条件的卡时这个效果不适用，对方可以确认自己的卡组（不取对象）
		 中文名: 融合贤者
		 日文名: 融合賢者
		 英文名: Fusion Sage
		 卡片种类: 通常魔法
		 卡片密码: 26902560
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: DP01，BE02，DL04，Booster06，Booster Chronicle，Booster R3，JY，PE，SY2
		 效果: 选1张自己的卡组的「融合」魔法卡加入手卡。之后洗切卡组。
		}
		*/
		Id:       529,
		Password: "26902560",
		Name:     "融合贤者",               // "Fusion Sage"  "融合賢者"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*64*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 530
		 调整: [绝对防御将军]<絶対防御将軍>

		●这张卡召唤，反转召唤成功时自动变成守备表示。
		◇诱发效果（进入连锁）。
		◇强制发动。
		●这张卡可以在防守表示的状态攻击（要计算战斗伤害）。
		◇永续效果（不进入连锁）。
		◇攻击表示的这张卡攻击宣言时，「附锁链的回旋镖/鎖付きブーメラン」「伊塔库亚的暴风/イタクァの暴風」等发动的场合，不能阻止这张卡的攻击行为。
		◇会被「炸裂装甲/炸裂装甲」的效果破坏。
		 中文名: 绝对防御将军
		 日文名: 絶対防御将軍
		 英文名: Total Defense Shogun
		 卡片种类: 效果怪兽
		 卡片密码: 75372290
		 使用限制: 无限制
		 种族: 战士
		 属性: 暗
		 星级: 6
		 攻击力: 1550
		 防御力: 2500
		 罕见度: 面闪SR，金字UR，爆闪PR，银碎SER
		 卡包: BE02，DL04，VJ，Booster Chronicle
		 效果: 这张卡召唤·反转召唤成功的场合守备表示。这张卡可以在守备表示时攻击。守备表示攻击时用攻击力的数值计算伤害。
		}
		*/
		Id:       530,
		Password: "75372290",
		Name:     "绝对防御将军",             // "Total Defense Shogun"  "絶対防御将軍"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   6,
		La:      ygo.LA_Dark,    // 暗
		Lr:      ygo.LR_Warrior, // 战士
		Attack:  1550,
		Defense: 2500,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*65*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 1227
		 中文名: 灯之魔精
		 日文名: ランプの魔精·ラ·ジーン
		 英文名: La Jinn the Mystical Genie of the Lamp
		 卡片种类: 通常怪兽
		 卡片密码: 97590747
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 1800
		 防御力: 1000
		 罕见度: 平卡N，金字UR
		 卡包: LE03，EX-R(EX)，Booster04，Booster R2，TP10，DPKB，KA，SK2
		 效果: 描述：听从呼唤他的主人的任何要求及命令的灯之精灵。
		}
		*/
		Id:       1227,
		Password: "97590747",
		Name:     "灯之魔精",                 // "La Jinn the Mystical Genie of the Lamp"  "ランプの魔精·ラ·ジーン"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Dark,  // 暗
		Lr:       ygo.LR_Fiend, // 恶魔
		Attack:   1800,
		Defense:  1000,
		IsValid:  true,
	})

	/*66*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 149
		 中文名: 地狱的裁判
		 日文名: 地獄の裁判
		 英文名: Trial of Nightmare
		 卡片种类: 通常怪兽
		 卡片密码: 77827521
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 1300
		 防御力: 900
		 罕见度: 平卡N，金碎USR
		 卡包: LB，DL02，Starter Box，Booster R1
		 效果: 描述：将敌人封入棺材里。然后由地狱使者下达判决。
		}
		*/
		Id:       149,
		Password: "77827521",
		Name:     "地狱的裁判",                // "Trial of Nightmare"  "地獄の裁判"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Dark,  // 暗
		Lr:       ygo.LR_Fiend, // 恶魔
		Attack:   1300,
		Defense:  900,
		IsValid:  true,
	})

	/*67*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 150
		 中文名: 第13人的埋葬者
		 日文名: １３人目の埋葬者
		 英文名: The 13th Grave
		 卡片种类: 通常怪兽
		 卡片密码: 00032864
		 使用限制: 无限制
		 种族: 不死
		 属性: 暗
		 星级: 3
		 攻击力: 1200
		 防御力: 900
		 罕见度: 平卡N，金碎USR
		 卡包: LB，DL02，Starter Box，Booster R1
		 效果: 描述：原本没有任何人的一零三号墓地突然出现的丧尸。
		}
		*/
		Id:       150,
		Password: "00032864",
		Name:     "第13人的埋葬者",             // "The 13th Grave"  "１３人目の埋葬者"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Dark,   // 暗
		Lr:       ygo.LR_Zombie, // 不死
		Attack:   1200,
		Defense:  900,
		IsValid:  true,
	})

	/*68*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 152
		 中文名: 深渊的冥王
		 日文名: 深淵の冥王
		 英文名: Dark King of the Abyss
		 卡片种类: 通常怪兽
		 卡片密码: 53375573
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 3
		 攻击力: 1200
		 防御力: 800
		 罕见度: 平卡N，金碎USR
		 卡包: LB，BE01，DL02，Starter Box，Booster R1
		 效果: 描述：冥界之王，听说以前有着支配所有黑暗力量的能力。
		}
		*/
		Id:       152,
		Password: "53375573",
		Name:     "深渊的冥王",                // "Dark King of the Abyss"  "深淵の冥王"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Dark,  // 暗
		Lr:       ygo.LR_Fiend, // 恶魔
		Attack:   1200,
		Defense:  800,
		IsValid:  true,
	})

	/*69*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 154
		 中文名: 水魔道士
		 日文名: アクア·マドール
		 英文名: Aqua Madoor
		 卡片种类: 通常怪兽
		 卡片密码: 85639257
		 使用限制: 无限制
		 种族: 魔法师
		 属性: 水
		 星级: 4
		 攻击力: 1200
		 防御力: 2000
		 罕见度: 金碎USR，平罕NR，平卡N
		 卡包: LB，BE01，DL02，Starter Box，Booster R1
		 效果: 描述：操纵水的魔法师。造出厚重的水墙压垮敌人。
		}
		*/
		Id:       154,
		Password: "85639257",
		Name:     "水魔道士",                 // "Aqua Madoor"  "アクア·マドール"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Water,       // 水
		Lr:       ygo.LR_SpellCaster, // 魔法师
		Attack:   1200,
		Defense:  2000,
		IsValid:  true,
	})

	/*70*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2445
		 中文名: 女夜魔战士
		 日文名: ヴィシュワ·ランディー
		 英文名: Vishwar Randi
		 卡片种类: 通常怪兽
		 卡片密码: 78556320
		 使用限制: 无限制
		 种族: 战士
		 属性: 暗
		 星级: 3
		 攻击力: 900
		 防御力: 700
		 罕见度: 平卡N
		 卡包: Booster01，Booster R1
		 效果: 描述：侍奉黑暗的女战士。将对方进行血祭是她生存的意义。
		}
		*/
		Id:       2445,
		Password: "78556320",
		Name:     "女夜魔战士",                // "Vishwar Randi"  "ヴィシュワ·ランディー"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Dark,    // 暗
		Lr:       ygo.LR_Warrior, // 战士
		Attack:   900,
		Defense:  700,
		IsValid:  true,
	})

	/*71*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2446
		 中文名: 气象控制员
		 日文名: ウェザー·コントロール
		 英文名: Weather Control
		 卡片种类: 通常怪兽
		 卡片密码: 37243151
		 使用限制: 无限制
		 种族: 天使
		 属性: 光
		 星级: 2
		 攻击力: 600
		 防御力: 400
		 罕见度: 平卡N
		 卡包: Booster01
		 效果: 描述：可以自由操纵天气。山里的天气无常就是这家伙的杰作。
		}
		*/
		Id:       2446,
		Password: "37243151",
		Name:     "气象控制员",                // "Weather Control"  "ウェザー·コントロール"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    2,
		La:       ygo.LA_Light, // 光
		Lr:       ygo.LR_Angel, // 天使
		Attack:   600,
		Defense:  400,
		IsValid:  true,
	})

	/*72*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2447
		 中文名: 水元素
		 日文名: ウォーター·エレメント
		 英文名: Water Element
		 卡片种类: 通常怪兽
		 卡片密码: 03732747
		 使用限制: 无限制
		 种族: 水
		 属性: 水
		 星级: 3
		 攻击力: 900
		 防御力: 700
		 罕见度: 平卡N
		 卡包: Booster01，Booster R1
		 效果: 描述：住在水里的精灵。将四周用雾包围妨碍敌人的视线。
		}
		*/
		Id:       2447,
		Password: "03732747",
		Name:     "水元素",                  // "Water Element"  "ウォーター·エレメント"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Water, // 水
		Lr:       ygo.LR_None,  // 水
		Attack:   900,
		Defense:  700,
		IsValid:  true,
	})

	/*73*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2448
		 中文名: 石像怪
		 日文名: ガーゴイル
		 英文名: Ryu-Kishin
		 卡片种类: 通常怪兽
		 卡片密码: 15303296
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 3
		 攻击力: 1000
		 防御力: 500
		 罕见度: 平卡N
		 卡包: EX-R(EX)，Booster01，Booster R1
		 效果: 描述：使人误认为是石像，从而在黑暗之中攻击。逃跑速度也很快。
		}
		*/
		Id:       2448,
		Password: "15303296",
		Name:     "石像怪",                  // "Ryu-Kishin"  "ガーゴイル"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Dark,  // 暗
		Lr:       ygo.LR_Fiend, // 恶魔
		Attack:   1000,
		Defense:  500,
		IsValid:  true,
	})

	/*74*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2449
		 中文名: 格斗战士 阿提米特
		 日文名: 格闘戦士アルティメーター
		 英文名: Battle Warrior
		 卡片种类: 通常怪兽
		 卡片密码: 55550921
		 使用限制: 无限制
		 种族: 战士
		 属性: 地
		 星级: 3
		 攻击力: 700
		 防御力: 1000
		 罕见度: 平卡N，面闪SR
		 卡包: Booster01，Booster R1，JY，TP16，NUMH
		 效果: 描述：不使用任何武器，空手战斗的格斗战士。
		}
		*/
		Id:       2449,
		Password: "55550921",
		Name:     "格斗战士 阿提米特",            // "Battle Warrior"  "格闘戦士アルティメーター"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Earth,   // 地
		Lr:       ygo.LR_Warrior, // 战士
		Attack:   700,
		Defense:  1000,
		IsValid:  true,
	})

	/*75*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2450
		 中文名: 风之番人 精
		 日文名: 風の番人 ジン
		 英文名: Djinn the Watcher of the Wind
		 卡片种类: 通常怪兽
		 卡片密码: 97843505
		 使用限制: 无限制
		 种族: 魔法师
		 属性: 风
		 星级: 3
		 攻击力: 700
		 防御力: 900
		 罕见度: 平卡N
		 卡包: Booster01
		 效果: 描述：操纵风并产生龙卷风与飓风吹飞周围的东西。
		}
		*/
		Id:       2450,
		Password: "97843505",
		Name:     "风之番人 精",               // "Djinn the Watcher of the Wind"  "風の番人 ジン"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Wind,        // 风
		Lr:       ygo.LR_SpellCaster, // 魔法师
		Attack:   700,
		Defense:  900,
		IsValid:  true,
	})

	/*76*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2451
		 中文名: 格洛斯
		 日文名: グロス
		 英文名: Twin Long Rods #1
		 卡片种类: 通常怪兽
		 卡片密码: 60589682
		 使用限制: 无限制
		 种族: 水
		 属性: 水
		 星级: 3
		 攻击力: 900
		 防御力: 700
		 罕见度: 平卡N
		 卡包: Booster01，Booster R1
		 效果: 描述：用像鞭子一般长的手臂进行攻击。稍远的地方也可以攻击的到。
		}
		*/
		Id:       2451,
		Password: "60589682",
		Name:     "格洛斯",                  // "Twin Long Rods #1"  "グロス"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Water, // 水
		Lr:       ygo.LR_None,  // 水
		Attack:   900,
		Defense:  700,
		IsValid:  true,
	})

	/*77*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2452
		 中文名: 幽灵
		 日文名: ゴースト
		 英文名: Phantom Ghost
		 卡片种类: 通常怪兽
		 卡片密码: 61201220
		 使用限制: 无限制
		 种族: 不死
		 属性: 暗
		 星级: 2
		 攻击力: 600
		 防御力: 800
		 罕见度: 平卡N
		 卡包: Booster01
		 效果: 描述：由这个世界上不能成佛的灵魂，慢慢聚集而成的怨灵。
		}
		*/
		Id:       2452,
		Password: "61201220",
		Name:     "幽灵",                   // "Phantom Ghost"  "ゴースト"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    2,
		La:       ygo.LA_Dark,   // 暗
		Lr:       ygo.LR_Zombie, // 不死
		Attack:   600,
		Defense:  800,
		IsValid:  true,
	})

	/*78*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2453
		 中文名: 萨塔那
		 日文名: サターナ
		 英文名: Phantom Dewan
		 卡片种类: 通常怪兽
		 卡片密码: 77603950
		 使用限制: 无限制
		 种族: 魔法师
		 属性: 暗
		 星级: 2
		 攻击力: 700
		 防御力: 600
		 罕见度: 平卡N
		 卡包: Booster01
		 效果: 描述：诅咒敌人。可以使人动弹不得的魔法师。
		}
		*/
		Id:       2453,
		Password: "77603950",
		Name:     "萨塔那",                  // "Phantom Dewan"  "サターナ"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    2,
		La:       ygo.LA_Dark,        // 暗
		Lr:       ygo.LR_SpellCaster, // 魔法师
		Attack:   700,
		Defense:  600,
		IsValid:  true,
	})

	/*79*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2454
		 中文名: 邪炎之翼
		 日文名: 邪炎の翼
		 英文名: Wings of Wicked Flame
		 卡片种类: 通常怪兽
		 卡片密码: 92944626
		 使用限制: 无限制
		 种族: 炎
		 属性: 炎
		 星级: 2
		 攻击力: 700
		 防御力: 600
		 罕见度: 平卡N
		 卡包: Booster01，TP17
		 效果: 描述：原形是红黑色燃烧着的翅膀。从全身喷出火焰进行攻击。
		}
		*/
		Id:       2454,
		Password: "92944626",
		Name:     "邪炎之翼",                 // "Wings of Wicked Flame"  "邪炎の翼"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    2,
		La:       ygo.LA_Fire, // 炎
		Lr:       ygo.LR_Pyro, // 炎
		Attack:   700,
		Defense:  600,
		IsValid:  true,
	})

	/*80*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2455
		 中文名: 大嘴鸟
		 日文名: スピック
		 英文名: Droll Bird
		 卡片种类: 通常怪兽
		 卡片密码: 97973387
		 使用限制: 无限制
		 种族: 鸟兽
		 属性: 风
		 星级: 2
		 攻击力: 600
		 防御力: 500
		 罕见度: 平卡N
		 卡包: Booster01
		 效果: 描述：嘴巴非常大、大声的叫吓跑胆小的对方。
		}
		*/
		Id:       2455,
		Password: "97973387",
		Name:     "大嘴鸟",                  // "Droll Bird"  "スピック"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    2,
		La:       ygo.LA_Wind,      // 风
		Lr:       ygo.LR_WindBeast, // 鸟兽
		Attack:   600,
		Defense:  500,
		IsValid:  true,
	})

	/*81*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2456
		 中文名: 神圣之锁
		 日文名: 聖なる鎖
		 英文名: Mystical Capture Chain
		 卡片种类: 通常怪兽
		 卡片密码: 63515678
		 使用限制: 无限制
		 种族: 天使
		 属性: 光
		 星级: 2
		 攻击力: 700
		 防御力: 700
		 罕见度: 平卡N
		 卡包: Booster01
		 效果: 描述：据说可以用神圣的力量封锁行动的锁链。
		}
		*/
		Id:       2456,
		Password: "63515678",
		Name:     "神圣之锁",                 // "Mystical Capture Chain"  "聖なる鎖"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    2,
		La:       ygo.LA_Light, // 光
		Lr:       ygo.LR_Angel, // 天使
		Attack:   700,
		Defense:  700,
		IsValid:  true,
	})

	/*82*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2457
		 中文名: 僵尸鬼灯
		 日文名: ゾンビランプ
		 英文名: Mech Mole Zombie
		 卡片种类: 通常怪兽
		 卡片密码: 63545455
		 使用限制: 无限制
		 种族: 不死
		 属性: 暗
		 星级: 2
		 攻击力: 500
		 防御力: 400
		 罕见度: 平卡N
		 卡包: Booster01
		 效果: 描述：把手腕像火箭一样发射攻击的不死怪兽。
		}
		*/
		Id:       2457,
		Password: "63545455",
		Name:     "僵尸鬼灯",                 // "Mech Mole Zombie"  "ゾンビランプ"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    2,
		La:       ygo.LA_Dark,   // 暗
		Lr:       ygo.LR_Zombie, // 不死
		Attack:   500,
		Defense:  400,
		IsValid:  true,
	})

	/*83*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2458
		 中文名: 暗黑植物
		 日文名: ダーク·プラント
		 英文名: Dark Plant
		 卡片种类: 通常怪兽
		 卡片密码: 13193642
		 使用限制: 无限制
		 种族: 植物
		 属性: 暗
		 星级: 1
		 攻击力: 300
		 防御力: 400
		 罕见度: 平卡N
		 卡包: Booster01
		 效果: 描述：在被污染的土地以及暗之力的环境下生长的花。非常凶暴。
		}
		*/
		Id:       2458,
		Password: "13193642",
		Name:     "暗黑植物",                 // "Dark Plant"  "ダーク·プラント"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    1,
		La:       ygo.LA_Dark,  // 暗
		Lr:       ygo.LR_Plant, // 植物
		Attack:   300,
		Defense:  400,
		IsValid:  true,
	})

	/*84*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2459
		 中文名: 太古之壶
		 日文名: 太古の壺
		 英文名: Ancient Jar
		 卡片种类: 通常怪兽
		 卡片密码: 81492226
		 使用限制: 无限制
		 种族: 岩石
		 属性: 地
		 星级: 1
		 攻击力: 400
		 防御力: 200
		 罕见度: 平卡N
		 卡包: Booster01
		 效果: 描述：非常脆弱的太古之壶。里面好像藏着什么东西。
		}
		*/
		Id:       2459,
		Password: "81492226",
		Name:     "太古之壶",                 // "Ancient Jar"  "太古の壺"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    1,
		La:       ygo.LA_Earth, // 地
		Lr:       ygo.LR_Rock,  // 岩石
		Attack:   400,
		Defense:  200,
		IsValid:  true,
	})

	/*85*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2477
		 中文名: 邪恶老鼠
		 日文名: イビル·ラット
		 英文名: Obese Marmot of Nefariousness
		 卡片种类: 通常怪兽
		 卡片密码: 56713552
		 使用限制: 无限制
		 种族: 兽
		 属性: 地
		 星级: 3
		 攻击力: 750
		 防御力: 800
		 罕见度: 平卡N
		 卡包: Booster02
		 效果: 描述：任何东西都咬的野老鼠，行仪极坏。
		}
		*/
		Id:       2477,
		Password: "56713552",
		Name:     "邪恶老鼠",                 // "Obese Marmot of Nefariousness"  "イビル·ラット"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Earth, // 地
		Lr:       ygo.LR_Beast, // 兽
		Attack:   750,
		Defense:  800,
		IsValid:  true,
	})

	/*86*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2478
		 中文名: 锯足锹形虫
		 日文名: インセクション
		 英文名: Alinsection
		 卡片种类: 通常怪兽
		 卡片密码: 70924884
		 使用限制: 无限制
		 种族: 昆虫
		 属性: 地
		 星级: 3
		 攻击力: 950
		 防御力: 700
		 罕见度: 平卡N
		 卡包: Booster02，Booster R1
		 效果: 描述：大甲虫。除了头上的锯子，手腕也变成了锯子。
		}
		*/
		Id:       2478,
		Password: "70924884",
		Name:     "锯足锹形虫",                // "Alinsection"  "インセクション"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Earth,  // 地
		Lr:       ygo.LR_Insect, // 昆虫
		Attack:   950,
		Defense:  700,
		IsValid:  true,
	})

	/*87*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2479
		 中文名: 威尔米
		 日文名: ウィルミー
		 英文名: Wilmee
		 卡片种类: 通常怪兽
		 卡片密码: 92391084
		 使用限制: 无限制
		 种族: 兽
		 属性: 地
		 星级: 4
		 攻击力: 1000
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: Booster02，Booster R1
		 效果: 描述：相当凶暴的兔子。以锐利的钩爪血祭对方。
		}
		*/
		Id:       2479,
		Password: "92391084",
		Name:     "威尔米",                  // "Wilmee"  "ウィルミー"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Earth, // 地
		Lr:       ygo.LR_Beast, // 兽
		Attack:   1000,
		Defense:  1200,
		IsValid:  true,
	})

	/*88*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2480
		 中文名: 木灵小丑
		 日文名: ウッド·ジョーカー
		 英文名: Wood Clown
		 卡片种类: 通常怪兽
		 卡片密码: 17511156
		 使用限制: 无限制
		 种族: 战士
		 属性: 地
		 星级: 3
		 攻击力: 800
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: Booster02，Booster R1
		 效果: 描述：作出令人讨厌的笑容的恶魔。以手中的镰刀熟练的回避着攻击。
		}
		*/
		Id:       2480,
		Password: "17511156",
		Name:     "木灵小丑",                 // "Wood Clown"  "ウッド·ジョーカー"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Earth,   // 地
		Lr:       ygo.LR_Warrior, // 战士
		Attack:   800,
		Defense:  1200,
		IsValid:  true,
	})

	/*89*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2481
		 中文名: 天使魔女
		 日文名: エンジェル·魔女
		 英文名: Angelwitch
		 卡片种类: 通常怪兽
		 卡片密码: 37160778
		 使用限制: 无限制
		 种族: 魔法师
		 属性: 暗
		 星级: 3
		 攻击力: 800
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: Booster02，Booster R1
		 效果: 描述：背负着成为天使的命运，但却成为了向往中的魔女。
		}
		*/
		Id:       2481,
		Password: "37160778",
		Name:     "天使魔女",                 // "Angelwitch"  "エンジェル·魔女"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Dark,        // 暗
		Lr:       ygo.LR_SpellCaster, // 魔法师
		Attack:   800,
		Defense:  1000,
		IsValid:  true,
	})

	/*90*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2482
		 中文名: 电波英雄
		 日文名: オシロ·ヒーロー
		 英文名: Oscillo Hero
		 卡片种类: 通常怪兽
		 卡片密码: 82065276
		 使用限制: 无限制
		 种族: 战士
		 属性: 地
		 星级: 3
		 攻击力: 1250
		 防御力: 700
		 罕见度: 平卡N
		 卡包: Booster02，Booster Chronicle，Booster R1，TU04
		 效果: 描述：从异世界来到这里，稀里糊涂的战士。
		}
		*/
		Id:       2482,
		Password: "82065276",
		Name:     "电波英雄",                 // "Oscillo Hero"  "オシロ·ヒーロー"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Earth,   // 地
		Lr:       ygo.LR_Warrior, // 战士
		Attack:   1250,
		Defense:  700,
		IsValid:  true,
	})

	/*91*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2483
		 中文名: 鱼战士
		 日文名: 魚ギョ戦士
		 英文名: Wow Warrior
		 卡片种类: 通常怪兽
		 卡片密码: 69750536
		 使用限制: 无限制
		 种族: 鱼
		 属性: 水
		 星级: 4
		 攻击力: 1250
		 防御力: 900
		 罕见度: 平卡N
		 卡包: Booster02，Booster Chronicle，Booster R1
		 效果: 描述：拥有手和脚的鱼人兽，用尖锐的牙齿撕咬敌人。
		}
		*/
		Id:       2483,
		Password: "69750536",
		Name:     "鱼战士",                  // "Wow Warrior"  "魚ギョ戦士"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Water, // 水
		Lr:       ygo.LR_Fish,  // 鱼
		Attack:   1250,
		Defense:  900,
		IsValid:  true,
	})

	/*92*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2484
		 中文名: 龙虾怪
		 日文名: ザリガン
		 英文名: Zarigun
		 卡片种类: 通常怪兽
		 卡片密码: 10598400
		 使用限制: 无限制
		 种族: 水
		 属性: 水
		 星级: 2
		 攻击力: 600
		 防御力: 700
		 罕见度: 平卡N
		 卡包: Booster02
		 效果: 描述：由虾进化成的怪兽。用大钳子攻击对方的颈项。
		}
		*/
		Id:       2484,
		Password: "10598400",
		Name:     "龙虾怪",                  // "Zarigun"  "ザリガン"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    2,
		La:       ygo.LA_Water, // 水
		Lr:       ygo.LR_None,  // 水
		Attack:   600,
		Defense:  700,
		IsValid:  true,
	})

	/*93*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2502
		 中文名: 赤剑之莱蒙多斯
		 日文名: 赤き剣のライムンドス
		 英文名: Rhaimundos of the Red Sword
		 卡片种类: 通常怪兽
		 卡片密码: 62403074
		 使用限制: 无限制
		 种族: 战士
		 属性: 地
		 星级: 4
		 攻击力: 1200
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: Booster03，TP23
		 效果: 描述：持有赤红炎剑的战士。利用火焰束缚来封住行动。
		}
		*/
		Id:       2502,
		Password: "62403074",
		Name:     "赤剑之莱蒙多斯",              // "Rhaimundos of the Red Sword"  "赤き剣のライムンドス"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Earth,   // 地
		Lr:       ygo.LR_Warrior, // 战士
		Attack:   1200,
		Defense:  1300,
		IsValid:  true,
	})

	/*94*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2503
		 中文名: 有生命的花瓶
		 日文名: 命ある花瓶
		 英文名: Living Vase
		 卡片种类: 通常怪兽
		 卡片密码: 34320307
		 使用限制: 无限制
		 种族: 植物
		 属性: 地
		 星级: 3
		 攻击力: 900
		 防御力: 1100
		 罕见度: 平卡N
		 卡包: Booster03
		 效果: 描述：利用花散发花粉并咬向对方的活花瓶。
		}
		*/
		Id:       2503,
		Password: "34320307",
		Name:     "有生命的花瓶",               // "Living Vase"  "命ある花瓶"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Earth, // 地
		Lr:       ygo.LR_Plant, // 植物
		Attack:   900,
		Defense:  1100,
		IsValid:  true,
	})

	/*95*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2504
		 中文名: 食命者
		 日文名: 命を食する者
		 英文名: That Which Feeds on Life
		 卡片种类: 通常怪兽
		 卡片密码: 52367652
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 1200
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: Booster03
		 效果: 描述：嗜食所有生物的灵魂，将将其作为自己的能量。
		}
		*/
		Id:       2504,
		Password: "52367652",
		Name:     "食命者",                  // "That Which Feeds on Life"  "命を食する者"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Dark,  // 暗
		Lr:       ygo.LR_Fiend, // 恶魔
		Attack:   1200,
		Defense:  1000,
		IsValid:  true,
	})

	/*96*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2505
		 中文名: 岩之战士
		 日文名: 岩の戦士
		 英文名: Minomushi Warrior
		 卡片种类: 通常怪兽
		 卡片密码: 46864967
		 使用限制: 无限制
		 种族: 岩石
		 属性: 地
		 星级: 4
		 攻击力: 1300
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: Booster03，Booster R2
		 效果: 描述：岩石战士。挥舞着非常重的石头剑。
		}
		*/
		Id:       2505,
		Password: "46864967",
		Name:     "岩之战士",                 // "Minomushi Warrior"  "岩の戦士"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Earth, // 地
		Lr:       ygo.LR_Rock,  // 岩石
		Attack:   1300,
		Defense:  1200,
		IsValid:  true,
	})

	/*97*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2506
		 中文名: 音女
		 日文名: 音女
		 英文名: Sonic Maid
		 卡片种类: 通常怪兽
		 卡片密码: 38942059
		 使用限制: 无限制
		 种族: 战士
		 属性: 地
		 星级: 3
		 攻击力: 1200
		 防御力: 900
		 罕见度: 平卡N
		 卡包: Booster03，Booster Chronicle，Booster R2
		 效果: 描述：擅长操纵声音的少女，能使音符变成镰刀攻击敌人。
		}
		*/
		Id:       2506,
		Password: "38942059",
		Name:     "音女",                   // "Sonic Maid"  "音女"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Earth,   // 地
		Lr:       ygo.LR_Warrior, // 战士
		Attack:   1200,
		Defense:  900,
		IsValid:  true,
	})

	/*98*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2507
		 中文名: 大白鲨
		 日文名: グレート·ホワイト
		 英文名: Great White
		 卡片种类: 通常怪兽
		 卡片密码: 13429800
		 使用限制: 无限制
		 种族: 鱼
		 属性: 水
		 星级: 4
		 攻击力: 1600
		 防御力: 800
		 罕见度: 平卡N
		 卡包: EX-R(EX)，Booster03，Booster R2
		 效果: 描述：巨大的白鲨，若被咬到绝对无法脱身。
		}
		*/
		Id:       2507,
		Password: "13429800",
		Name:     "大白鲨",                  // "Great White"  "グレート·ホワイト"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Water, // 水
		Lr:       ygo.LR_Fish,  // 鱼
		Attack:   1600,
		Defense:  800,
		IsValid:  true,
	})

	/*99*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2508
		 中文名: 锹甲阿尔法
		 日文名: クワガタ·アルファ
		 英文名: Kuwagata α
		 卡片种类: 通常怪兽
		 卡片密码: 60802233
		 使用限制: 无限制
		 种族: 昆虫
		 属性: 地
		 星级: 4
		 攻击力: 1250
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: Booster03，Booster Chronicle，Booster R2，TP20
		 效果: 描述：凶暴的甲虫。它会瞄准对方的头瞬间斩落。
		}
		*/
		Id:       2508,
		Password: "60802233",
		Name:     "锹甲阿尔法",                // "Kuwagata α"  "クワガタ·アルファ"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Earth,  // 地
		Lr:       ygo.LR_Insect, // 昆虫
		Attack:   1250,
		Defense:  1000,
		IsValid:  true,
	})

	/*100*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2509
		 中文名: 钢铁魔人
		 日文名: ゴルゴイル
		 英文名: Golgoil
		 卡片种类: 通常怪兽
		 卡片密码: 07526150
		 使用限制: 无限制
		 种族: 机械
		 属性: 地
		 星级: 4
		 攻击力: 900
		 防御力: 1600
		 罕见度: 平卡N
		 卡包: Booster03
		 效果: 描述：从与异次元相通的洞里出来的钢铁大魔人。
		}
		*/
		Id:       2509,
		Password: "07526150",
		Name:     "钢铁魔人",                 // "Golgoil"  "ゴルゴイル"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Earth,   // 地
		Lr:       ygo.LR_Machine, // 机械
		Attack:   900,
		Defense:  1600,
		IsValid:  true,
	})

	/*101*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2510
		 中文名: 青玉眼怪
		 日文名: サファイヤ·リサーク
		 英文名: Lisark
		 卡片种类: 通常怪兽
		 卡片密码: 55210709
		 使用限制: 无限制
		 种族: 兽
		 属性: 地
		 星级: 4
		 攻击力: 1300
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: Booster03，Booster R2
		 效果: 描述：蓝宝石眼睛的野兽。制造幻影，趁敌人混乱的时候进行攻击。
		}
		*/
		Id:       2510,
		Password: "55210709",
		Name:     "青玉眼怪",                 // "Lisark"  "サファイヤ·リサーク"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Earth, // 地
		Lr:       ygo.LR_Beast, // 兽
		Attack:   1300,
		Defense:  1300,
		IsValid:  true,
	})

	/*102*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2511
		 中文名: 斩首的美女
		 日文名: 斬首の美女
		 英文名: Beautiful Headhuntress
		 卡片种类: 通常怪兽
		 卡片密码: 16899564
		 使用限制: 无限制
		 种族: 战士
		 属性: 地
		 星级: 4
		 攻击力: 1600
		 防御力: 800
		 罕见度: 平卡N
		 卡包: Booster03，Booster Chronicle，Booster R2
		 效果: 描述：在那美丽容貌的背后，却是个用刀子使许多人身首异处的女子。
		}
		*/
		Id:       2511,
		Password: "16899564",
		Name:     "斩首的美女",                // "Beautiful Headhuntress"  "斬首の美女"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Earth,   // 地
		Lr:       ygo.LR_Warrior, // 战士
		Attack:   1600,
		Defense:  800,
		IsValid:  true,
	})

	/*103*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2512
		 中文名: 海马兽
		 日文名: シーホース
		 英文名: Tatsunootoshigo
		 卡片种类: 通常怪兽
		 卡片密码: 47922711
		 使用限制: 无限制
		 种族: 兽
		 属性: 地
		 星级: 5
		 攻击力: 1350
		 防御力: 1600
		 罕见度: 平卡N
		 卡包: Booster03，Booster R2
		 效果: 描述：半马半鱼的怪兽。以风一般的速度在海中奔驰。
		}
		*/
		Id:       2512,
		Password: "47922711",
		Name:     "海马兽",                  // "Tatsunootoshigo"  "シーホース"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Earth, // 地
		Lr:       ygo.LR_Beast, // 兽
		Attack:   1350,
		Defense:  1600,
		IsValid:  true,
	})

	/*104*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2513
		 中文名: 审判之手
		 日文名: ジャジメント·ザ·ハンド
		 英文名: The Judgement Hand
		 卡片种类: 通常怪兽
		 卡片密码: 28003512
		 使用限制: 无限制
		 种族: 战士
		 属性: 地
		 星级: 3
		 攻击力: 1400
		 防御力: 700
		 罕见度: 平卡N
		 卡包: Booster03，Booster Chronicle，Booster R2
		 效果: 描述：用寄宿着神灵的手作最后的判决，给予敌人以猛烈的攻击。
		}
		*/
		Id:       2513,
		Password: "28003512",
		Name:     "审判之手",                 // "The Judgement Hand"  "ジャジメント·ザ·ハンド"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Earth,   // 地
		Lr:       ygo.LR_Warrior, // 战士
		Attack:   1400,
		Defense:  700,
		IsValid:  true,
	})

	/*105*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2514
		 中文名: 骷髅寺院
		 日文名: 髑髏の寺院
		 英文名: Temple of Skulls
		 卡片种类: 通常怪兽
		 卡片密码: 00732302
		 使用限制: 无限制
		 种族: 不死
		 属性: 暗
		 星级: 4
		 攻击力: 900
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: Booster03
		 效果: 描述：全部由骷髅和骨头建立而成的寺院。会吸收接近的人的灵魂。
		}
		*/
		Id:       2514,
		Password: "00732302",
		Name:     "骷髅寺院",                 // "Temple of Skulls"  "髑髏の寺院"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Dark,   // 暗
		Lr:       ygo.LR_Zombie, // 不死
		Attack:   900,
		Defense:  1300,
		IsValid:  true,
	})

	/*106*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2515
		 中文名: 树精
		 日文名: ドリアード
		 英文名: Dryad
		 卡片种类: 通常怪兽
		 卡片密码: 84916669
		 使用限制: 无限制
		 种族: 魔法师
		 属性: 地
		 星级: 4
		 攻击力: 1200
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: Booster03，TP07
		 效果: 描述：森之精灵。借助草木之力封住对方的行动。
		}
		*/
		Id:       2515,
		Password: "84916669",
		Name:     "树精",                   // "Dryad"  "ドリアード"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Earth,       // 地
		Lr:       ygo.LR_SpellCaster, // 魔法师
		Attack:   1200,
		Defense:  1400,
		IsValid:  true,
	})

	/*107*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2516
		 中文名: 狂战士
		 日文名: バーサーカー
		 英文名: Mystic Clown
		 卡片种类: 通常怪兽
		 卡片密码: 47060154
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 1500
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: EX-R(EX)，Booster03，Booster R2
		 效果: 用狂暴的力量攻击，一旦暴走无人能挡。
		}
		*/
		Id:       2516,
		Password: "47060154",
		Name:     "狂战士",                  // "Mystic Clown"  "バーサーカー"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Dark,  // 暗
		Lr:       ygo.LR_Fiend, // 恶魔
		Attack:   1500,
		Defense:  1000,
		IsValid:  true,
	})

	/*108*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2517
		 中文名: 复仇的河童
		 日文名: 復讐のカッパ
		 英文名: Kappa Avenger
		 卡片种类: 通常怪兽
		 卡片密码: 48109103
		 使用限制: 无限制
		 种族: 水
		 属性: 水
		 星级: 3
		 攻击力: 1200
		 防御力: 900
		 罕见度: 平卡N
		 卡包: Booster03
		 效果: 描述：被同伴杀害，为了复仇而将心灵卖给邪恶的河童。
		}
		*/
		Id:       2517,
		Password: "48109103",
		Name:     "复仇的河童",                // "Kappa Avenger"  "復讐のカッパ"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Water, // 水
		Lr:       ygo.LR_None,  // 水
		Attack:   1200,
		Defense:  900,
		IsValid:  true,
	})

	/*109*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2533
		 中文名: 龟鸟
		 日文名: タートル·バード
		 英文名: Turtle Bird
		 卡片种类: 通常怪兽
		 卡片密码: 72929454
		 使用限制: 无限制
		 种族: 水
		 属性: 水
		 星级: 6
		 攻击力: 1900
		 防御力: 1700
		 罕见度: 平卡N
		 卡包: Booster04，Booster Chronicle，Booster R2
		 效果: 描述：主要栖息在水中，也能在空中飞翔的珍奇乌龟。
		}
		*/
		Id:       2533,
		Password: "72929454",
		Name:     "龟鸟",                   // "Turtle Bird"  "タートル·バード"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    6,
		La:       ygo.LA_Water, // 水
		Lr:       ygo.LR_None,  // 水
		Attack:   1900,
		Defense:  1700,
		IsValid:  true,
	})

	/*110*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2534
		 中文名: 电子鱼
		 日文名: サイボーグ·バス
		 英文名: Mech Bass
		 卡片种类: 通常怪兽
		 卡片密码: 50176820
		 使用限制: 无限制
		 种族: 机械
		 属性: 水
		 星级: 5
		 攻击力: 1800
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: Booster04，Booster R2
		 效果: 描述：以背后附着的炮台发射闪光粒子加农炮。
		}
		*/
		Id:       2534,
		Password: "50176820",
		Name:     "电子鱼",                  // "Mech Bass"  "サイボーグ·バス"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Water,   // 水
		Lr:       ygo.LR_Machine, // 机械
		Attack:   1800,
		Defense:  1500,
		IsValid:  true,
	})

	/*111*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2535
		 中文名: 水陆的帝王
		 日文名: 水陸の帝王
		 英文名: Emperor of the Land and Sea
		 卡片种类: 通常怪兽
		 卡片密码: 11250655
		 使用限制: 无限制
		 种族: 爬虫类
		 属性: 水
		 星级: 5
		 攻击力: 1800
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: Booster04，Booster R2
		 效果: 描述：可以用大嘴向四方喷火的爬虫怪。
		}
		*/
		Id:       2535,
		Password: "11250655",
		Name:     "水陆的帝王",                // "Emperor of the Land and Sea"  "水陸の帝王"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Water,   // 水
		Lr:       ygo.LR_Reptile, // 爬虫类
		Attack:   1800,
		Defense:  1500,
		IsValid:  true,
	})

	/*112*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2536
		 中文名: 红叶之女王
		 日文名: 紅葉の女王
		 英文名: Queen of Autumn Leaves
		 卡片种类: 通常怪兽
		 卡片密码: 04179849
		 使用限制: 无限制
		 种族: 植物
		 属性: 地
		 星级: 5
		 攻击力: 1800
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: Booster04，Booster06，Booster Chronicle，Booster R2
		 效果: 描述：生活在被红叶围绕的地方，绿树灵王的妃子。
		}
		*/
		Id:       2536,
		Password: "04179849",
		Name:     "红叶之女王",                // "Queen of Autumn Leaves"  "紅葉の女王"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Earth, // 地
		Lr:       ygo.LR_Plant, // 植物
		Attack:   1800,
		Defense:  1500,
		IsValid:  true,
	})

	/*113*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2537
		 中文名: 战神 奥利安
		 日文名: 戦いの神 オリオン
		 英文名: Orion the Battle Kami
		 卡片种类: 通常怪兽
		 卡片密码: 02971090
		 使用限制: 无限制
		 种族: 天使
		 属性: 光
		 星级: 5
		 攻击力: 1800
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: Booster04，Booster R2
		 效果: 描述：被誉为战神的天使。不过谁也没见过那场战事。
		}
		*/
		Id:       2537,
		Password: "02971090",
		Name:     "战神 奥利安",               // "Orion the Battle Kami"  "戦いの神 オリオン"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Light, // 光
		Lr:       ygo.LR_Angel, // 天使
		Attack:   1800,
		Defense:  1500,
		IsValid:  true,
	})

	/*114*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2538
		 中文名: 山之精灵
		 日文名: 山の精霊
		 英文名: Spirit of the Mountain
		 卡片种类: 通常怪兽
		 卡片密码: 34690519
		 使用限制: 无限制
		 种族: 魔法师
		 属性: 地
		 星级: 5
		 攻击力: 1300
		 防御力: 1800
		 罕见度: 平卡N
		 卡包: Booster04，Booster R2
		 效果: 描述：据说听了它的笛音的人会失去力量。
		}
		*/
		Id:       2538,
		Password: "34690519",
		Name:     "山之精灵",                 // "Spirit of the Mountain"  "山の精霊"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Earth,       // 地
		Lr:       ygo.LR_SpellCaster, // 魔法师
		Attack:   1300,
		Defense:  1800,
		IsValid:  true,
	})

	/*115*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2539
		 中文名: 诞生的天使
		 日文名: 誕生の天使
		 英文名: Winged Egg of New Life
		 卡片种类: 通常怪兽
		 卡片密码: 42418084
		 使用限制: 无限制
		 种族: 天使
		 属性: 光
		 星级: 5
		 攻击力: 1400
		 防御力: 1700
		 罕见度: 平卡N
		 卡包: Booster04，Booster R2
		 效果: 描述：据说它能知道女性腹部是否存在着生命。
		}
		*/
		Id:       2539,
		Password: "42418084",
		Name:     "诞生的天使",                // "Winged Egg of New Life"  "誕生の天使"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Light, // 光
		Lr:       ygo.LR_Angel, // 天使
		Attack:   1400,
		Defense:  1700,
		IsValid:  true,
	})

	/*116*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2557
		 中文名: 飞鹰
		 日文名: ウイング·イーグル
		 英文名: Wing Eagle
		 卡片种类: 通常怪兽
		 卡片密码: 47319141
		 使用限制: 无限制
		 种族: 鸟兽
		 属性: 风
		 星级: 5
		 攻击力: 1800
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: Booster05，Booster Chronicle，Booster R3
		 效果: 描述：从高空寻找猎物，一旦看上绝不让其逃走。
		}
		*/
		Id:       2557,
		Password: "47319141",
		Name:     "飞鹰",                   // "Wing Eagle"  "ウイング·イーグル"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Wind,      // 风
		Lr:       ygo.LR_WindBeast, // 鸟兽
		Attack:   1800,
		Defense:  1500,
		IsValid:  true,
	})

	/*117*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2558
		 中文名: 机器攻击者
		 日文名: マシン·アタッカー
		 英文名: Machine Attacker
		 卡片种类: 通常怪兽
		 卡片密码: 38116136
		 使用限制: 无限制
		 种族: 机械
		 属性: 地
		 星级: 5
		 攻击力: 1600
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: Booster05
		 效果: 描述：特攻用的机械。以突击之力打倒敌人。
		}
		*/
		Id:       2558,
		Password: "38116136",
		Name:     "机器攻击者",                // "Machine Attacker"  "マシン·アタッカー"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Earth,   // 地
		Lr:       ygo.LR_Machine, // 机械
		Attack:   1600,
		Defense:  1300,
		IsValid:  true,
	})

	/*118*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2559
		 中文名: 贪尸龙
		 日文名: 屍を貪る竜
		 英文名: Crawling Dragon #2
		 卡片种类: 通常怪兽
		 卡片密码: 38289717
		 使用限制: 无限制
		 种族: 恐龙
		 属性: 地
		 星级: 4
		 攻击力: 1600
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: Booster05，Booster Chronicle，Booster R3
		 效果: 描述：什么都能咬碎的恐龙，攻击力十分恐怖。
		}
		*/
		Id:       2559,
		Password: "38289717",
		Name:     "贪尸龙",                  // "Crawling Dragon #2"  "屍を貪る竜"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Earth,    // 地
		Lr:       ygo.LR_Dinosaur, // 恐龙
		Attack:   1600,
		Defense:  1200,
		IsValid:  true,
	})

	/*119*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2560
		 中文名: 彩虹人鱼
		 日文名: レインボー·マリン·マーメイド
		 英文名: Rainbow Marine Mermaid
		 卡片种类: 通常怪兽
		 卡片密码: 29402771
		 使用限制: 无限制
		 种族: 鱼
		 属性: 水
		 星级: 5
		 攻击力: 1550
		 防御力: 1700
		 罕见度: 平卡N
		 卡包: Booster05
		 效果: 描述：空中出现彩虹桥时才会出现的珍奇美丽人鱼。
		}
		*/
		Id:       2560,
		Password: "29402771",
		Name:     "彩虹人鱼",                 // "Rainbow Marine Mermaid"  "レインボー·マリン·マーメイド"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Water, // 水
		Lr:       ygo.LR_Fish,  // 鱼
		Attack:   1550,
		Defense:  1700,
		IsValid:  true,
	})

	/*120*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2561
		 中文名: 美杜莎的亡灵
		 日文名: メデューサの亡霊
		 英文名: The Snake Hair
		 卡片种类: 通常怪兽
		 卡片密码: 29491031
		 使用限制: 无限制
		 种族: 不死
		 属性: 暗
		 星级: 4
		 攻击力: 1500
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: Booster05，Booster R3
		 效果: 描述：头发满是毒蛇的怪兽。一旦被她的目光盯上，就会被石化。
		}
		*/
		Id:       2561,
		Password: "29491031",
		Name:     "美杜莎的亡灵",               // "The Snake Hair"  "メデューサの亡霊"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Dark,   // 暗
		Lr:       ygo.LR_Zombie, // 不死
		Attack:   1500,
		Defense:  1200,
		IsValid:  true,
	})

	/*121*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2580
		 中文名: 牛鬼
		 日文名: 牛鬼
		 英文名: Ushi Oni
		 卡片种类: 通常怪兽
		 卡片密码: 48649353
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 6
		 攻击力: 2150
		 防御力: 1950
		 罕见度: 平卡N
		 卡包: Booster06，Booster Chronicle，Booster R3，TP11
		 效果: 描述：通过黑魔术苏醒的牛之恶魔。从壶中现身。
		}
		*/
		Id:       2580,
		Password: "48649353",
		Name:     "牛鬼",                   // "Ushi Oni"  "牛鬼"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    6,
		La:       ygo.LA_Dark,  // 暗
		Lr:       ygo.LR_Fiend, // 恶魔
		Attack:   2150,
		Defense:  1950,
		IsValid:  true,
	})

	/*122*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2581
		 中文名: 加尔瓦斯
		 日文名: ガルヴァス
		 英文名: Garvas
		 卡片种类: 通常怪兽
		 卡片密码: 69780745
		 使用限制: 无限制
		 种族: 兽
		 属性: 地
		 星级: 6
		 攻击力: 2000
		 防御力: 1700
		 罕见度: 平卡N
		 卡包: Booster06
		 效果: 描述：恶之化身。样子如同长着羽毛的狮子。
		}
		*/
		Id:       2581,
		Password: "69780745",
		Name:     "加尔瓦斯",                 // "Garvas"  "ガルヴァス"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    6,
		La:       ygo.LA_Earth, // 地
		Lr:       ygo.LR_Beast, // 兽
		Attack:   2000,
		Defense:  1700,
		IsValid:  true,
	})

	/*123*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 420
		 中文名: 水母
		 日文名: 海月－ジェリーフィッシュ－
		 英文名: Jellyfish
		 卡片种类: 通常怪兽
		 卡片密码: 14851496
		 使用限制: 无限制
		 种族: 水
		 属性: 水
		 星级: 4
		 攻击力: 1200
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: ME，DL04，Booster07
		 效果: 描述：漂浮在海面上的海蜇，半透明的身体使人难以察觉。
		}
		*/
		Id:       420,
		Password: "14851496",
		Name:     "水母",                   // "Jellyfish"  "海月－ジェリーフィッシュ－"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Water, // 水
		Lr:       ygo.LR_None,  // 水
		Attack:   1200,
		Defense:  1500,
		IsValid:  true,
	})

	/*124*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 422
		 中文名: 暗魔界的霸王
		 日文名: 闇魔界の覇王
		 英文名: King of Yamimakai
		 卡片种类: 通常怪兽
		 卡片密码: 69455834
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 5
		 攻击力: 2000
		 防御力: 1530
		 罕见度: 平卡N
		 卡包: ME，BE02，DL04，Booster07
		 效果: 描述：使用强大的暗之力量，破坏周围的一切。
		}
		*/
		Id:       422,
		Password: "69455834",
		Name:     "暗魔界的霸王",               // "King of Yamimakai"  "闇魔界の覇王"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Dark,  // 暗
		Lr:       ygo.LR_Fiend, // 恶魔
		Attack:   2000,
		Defense:  1530,
		IsValid:  true,
	})

	/*125*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 435
		 中文名: 强化石像怪
		 日文名: ガーゴイル·パワード
		 英文名: Ryu-Kishin Powered
		 卡片种类: 通常怪兽
		 卡片密码: 24611934
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 1600
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: ME，BE02，EX-R(EX)，DL04，Booster07，KA，SK2
		 效果: 描述：取得暗黑之力强化而成的石像鬼，尖锐的爪子值得警惕。
		}
		*/
		Id:       435,
		Password: "24611934",
		Name:     "强化石像怪",                // "Ryu-Kishin Powered"  "ガーゴイル·パワード"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Dark,  // 暗
		Lr:       ygo.LR_Fiend, // 恶魔
		Attack:   1600,
		Defense:  1200,
		IsValid:  true,
	})

	/*126*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 479
		 中文名: 电气小子
		 日文名: エレキッズ
		 英文名: Wattkid
		 卡片种类: 通常怪兽
		 卡片密码: 27324313
		 使用限制: 无限制
		 种族: 雷
		 属性: 光
		 星级: 3
		 攻击力: 1000
		 防御力: 500
		 罕见度: 平卡N
		 卡包: BE02，DL04，Booster01，Booster Chronicle，Booster R1，TU04
		 效果: 描述：雷属性攻击异常厉害，若轻视它则会遭到电击。
		}
		*/
		Id:       479,
		Password: "27324313",
		Name:     "电气小子",                 // "Wattkid"  "エレキッズ"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Light,   // 光
		Lr:       ygo.LR_Thunder, // 雷
		Attack:   1000,
		Defense:  500,
		IsValid:  true,
	})

	/*127*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 480
		 中文名: 吸血跳蚤
		 日文名: 吸血ノミ
		 英文名: Giant Flea
		 卡片种类: 通常怪兽
		 卡片密码: 41762634
		 使用限制: 无限制
		 种族: 昆虫
		 属性: 地
		 星级: 4
		 攻击力: 1500
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: DL04，Booster02，Booster Chronicle，Booster R1
		 效果: 描述：攻击力高强的巨型吸血跳蚤，轻视它会非常危险。
		}
		*/
		Id:       480,
		Password: "41762634",
		Name:     "吸血跳蚤",                 // "Giant Flea"  "吸血ノミ"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Earth,  // 地
		Lr:       ygo.LR_Insect, // 昆虫
		Attack:   1500,
		Defense:  1200,
		IsValid:  true,
	})

	/*128*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 485
		 中文名: 复活节岛的摩艾石像
		 日文名: イースター島のモアイ
		 英文名: The Statue of Easter Island
		 卡片种类: 通常怪兽
		 卡片密码: 10262698
		 使用限制: 无限制
		 种族: 岩石
		 属性: 地
		 星级: 4
		 攻击力: 1100
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: BE02，DL04，Booster02，Booster Chronicle，Booster R1
		 效果: 描述：存在于复活岛的石像，能够发射圆形的激光。
		}
		*/
		Id:       485,
		Password: "10262698",
		Name:     "复活节岛的摩艾石像",            // "The Statue of Easter Island"  "イースター島のモアイ"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Earth, // 地
		Lr:       ygo.LR_Rock,  // 岩石
		Attack:   1100,
		Defense:  1400,
		IsValid:  true,
	})

	/*129*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 486
		 中文名: 友谊天使
		 日文名: フレンドシップ
		 英文名: Shining Friendship
		 卡片种类: 通常怪兽
		 卡片密码: 82085619
		 使用限制: 无限制
		 种族: 天使
		 属性: 光
		 星级: 4
		 攻击力: 1300
		 防御力: 1100
		 罕见度: 平卡N
		 卡包: BE02，DL04，Booster02，Booster Chronicle，Booster R1
		 效果: 描述：即便在决斗中吵架，显示出友情就能和好。
		}
		*/
		Id:       486,
		Password: "82085619",
		Name:     "友谊天使",                 // "Shining Friendship"  "フレンドシップ"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Light, // 光
		Lr:       ygo.LR_Angel, // 天使
		Attack:   1300,
		Defense:  1100,
		IsValid:  true,
	})

	/*130*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 492
		 中文名: 老虎斧战士
		 日文名: タイガー·アックス
		 英文名: Tiger Axe
		 卡片种类: 通常怪兽
		 卡片密码: 49791927
		 使用限制: 无限制
		 种族: 兽战士
		 属性: 地
		 星级: 4
		 攻击力: 1300
		 防御力: 1100
		 罕见度: 平卡N，金字UR
		 卡包: BE02，LE02，DL04，Booster03，Booster Chronicle，Booster R2，JY
		 效果: 描述：手持巨斧的兽战士。能够放出行动迅速的人偶，攻击强劲。
		}
		*/
		Id:       492,
		Password: "49791927",
		Name:     "老虎斧战士",                // "Tiger Axe"  "タイガー·アックス"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Earth,       // 地
		Lr:       ygo.LR_BeastWarror, // 兽战士
		Attack:   1300,
		Defense:  1100,
		IsValid:  true,
	})

	/*131*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 493
		 中文名: 巨斧袭击者
		 日文名: アックス·レイダー
		 英文名: Axe Raider
		 卡片种类: 通常怪兽
		 卡片密码: 48305365
		 使用限制: 无限制
		 种族: 战士
		 属性: 地
		 星级: 4
		 攻击力: 1700
		 防御力: 1150
		 罕见度: 平卡N
		 卡包: BE02，DL04，Booster04，Booster Chronicle，Booster R2，YSD03，JY，SJ2
		 效果: 描述：持斧的战士。单手挥舞斧头的攻击相当强劲。
		}
		*/
		Id:       493,
		Password: "48305365",
		Name:     "巨斧袭击者",                // "Axe Raider"  "アックス·レイダー"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Earth,   // 地
		Lr:       ygo.LR_Warrior, // 战士
		Attack:   1700,
		Defense:  1150,
		IsValid:  true,
	})

	/*132*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 495
		 中文名: 机械猎手
		 日文名: メカ·ハンター
		 英文名: Mechanicalchaser
		 卡片种类: 通常怪兽
		 卡片密码: 07359741
		 使用限制: 无限制
		 种族: 机械
		 属性: 暗
		 星级: 4
		 攻击力: 1850
		 防御力: 800
		 罕见度: 平卡N
		 卡包: BE02，DL04，SD10，Booster04，Booster Chronicle，Booster R2
		 效果: 描述：听从机械王的命令，在抓住目标前都会紧追不舍的猎捕。
		}
		*/
		Id:       495,
		Password: "07359741",
		Name:     "机械猎手",                 // "Mechanicalchaser"  "メカ·ハンター"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Dark,    // 暗
		Lr:       ygo.LR_Machine, // 机械
		Attack:   1850,
		Defense:  800,
		IsValid:  true,
	})

	/*133*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 496
		 中文名: 海鳞蛇
		 日文名: シーザリオン
		 英文名: Giant Red Seasnake
		 卡片种类: 通常怪兽
		 卡片密码: 58831685
		 使用限制: 无限制
		 种族: 水
		 属性: 水
		 星级: 4
		 攻击力: 1800
		 防御力: 800
		 罕见度: 平卡N
		 卡包: DL04，Booster04，Booster Chronicle，Booster R2
		 效果: 描述：居住在海洋里的蛇型怪兽，它会去咬任何接近的物体。
		}
		*/
		Id:       496,
		Password: "58831685",
		Name:     "海鳞蛇",                  // "Giant Red Seasnake"  "シーザリオン"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Water, // 水
		Lr:       ygo.LR_None,  // 水
		Attack:   1800,
		Defense:  800,
		IsValid:  true,
	})

	/*134*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 498
		 中文名: 双生精灵
		 日文名: ヂェミナイ·エルフ
		 英文名: Gemini Elf
		 卡片种类: 通常怪兽
		 卡片密码: 69140098
		 使用限制: 无限制
		 种族: 魔法师
		 属性: 地
		 星级: 4
		 攻击力: 1900
		 防御力: 900
		 罕见度: 平卡N
		 卡包: BE02，DL04，SD06，Booster04，Booster Chronicle，Booster R2
		 效果: 描述：相互配合发动攻击的双胞胎妖精姐妹。
		}
		*/
		Id:       498,
		Password: "69140098",
		Name:     "双生精灵",                 // "Gemini Elf"  "ヂェミナイ·エルフ"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Earth,       // 地
		Lr:       ygo.LR_SpellCaster, // 魔法师
		Attack:   1900,
		Defense:  900,
		IsValid:  true,
	})

	/*135*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 500
		 中文名: 橐蜚
		 日文名: タクヒ
		 英文名: Takuhee
		 卡片种类: 通常怪兽
		 卡片密码: 03170832
		 使用限制: 无限制
		 种族: 鸟兽
		 属性: 风
		 星级: 4
		 攻击力: 1450
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: DL04，Booster03，Booster Chronicle，Booster R2
		 效果: 描述：这鸟出现的时候，就是什么不幸之事发生前的征兆。
		}
		*/
		Id:       500,
		Password: "03170832",
		Name:     "橐蜚",                   // "Takuhee"  "タクヒ"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Wind,      // 风
		Lr:       ygo.LR_WindBeast, // 鸟兽
		Attack:   1450,
		Defense:  1000,
		IsValid:  true,
	})

	/*136*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 502
		 中文名: 泉之妖精
		 日文名: 泉の妖精
		 英文名: Fairy of the Fountain
		 卡片种类: 通常怪兽
		 卡片密码: 81563416
		 使用限制: 无限制
		 种族: 水
		 属性: 水
		 星级: 4
		 攻击力: 1600
		 防御力: 1100
		 罕见度: 平卡N
		 卡包: DL04，Booster04，Booster Chronicle，Booster R2
		 效果: 描述：守护泉水的精灵。对任何污浊泉水者杀无赦。
		}
		*/
		Id:       502,
		Password: "81563416",
		Name:     "泉之妖精",                 // "Fairy of the Fountain"  "泉の妖精"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Water, // 水
		Lr:       ygo.LR_None,  // 水
		Attack:   1600,
		Defense:  1100,
		IsValid:  true,
	})

	/*137*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 503
		 中文名: 月光少女
		 日文名: 月明かりの乙女
		 英文名: Maiden of the Moonlight
		 卡片种类: 通常怪兽
		 卡片密码: 79629370
		 使用限制: 无限制
		 种族: 魔法师
		 属性: 光
		 星级: 4
		 攻击力: 1500
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: DL04，Booster04，Booster Chronicle，Booster R2
		 效果: 描述：月亮加护的魔导士。神秘的魔法能使人看到幻觉。
		}
		*/
		Id:       503,
		Password: "79629370",
		Name:     "月光少女",                 // "Maiden of the Moonlight"  "月明かりの乙女"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Light,       // 光
		Lr:       ygo.LR_SpellCaster, // 魔法师
		Attack:   1500,
		Defense:  1300,
		IsValid:  true,
	})

	/*138*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 506
		 中文名: 双头恐龙王
		 日文名: 二頭を持つキング·レックス
		 英文名: Two-Headed King Rex
		 卡片种类: 通常怪兽
		 卡片密码: 94119974
		 使用限制: 无限制
		 种族: 恐龙
		 属性: 地
		 星级: 4
		 攻击力: 1600
		 防御力: 1200
		 罕见度: 平卡N，金字UR，爆闪PR
		 卡包: BE02，DL04，Booster05，Booster Chronicle，Booster R3，TP19
		 效果: 描述：恐龙族中的强力怪兽，两只头同时攻击。
		}
		*/
		Id:       506,
		Password: "94119974",
		Name:     "双头恐龙王",                // "Two-Headed King Rex"  "二頭を持つキング·レックス"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Earth,    // 地
		Lr:       ygo.LR_Dinosaur, // 恐龙
		Attack:   1600,
		Defense:  1200,
		IsValid:  true,
	})

	/*139*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 507
		 中文名: 海龙神
		 日文名: 海竜神
		 英文名: Kairyu-Shin
		 卡片种类: 通常怪兽
		 卡片密码: 76634149
		 使用限制: 无限制
		 种族: 海龙
		 属性: 水
		 星级: 5
		 攻击力: 1800
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: DL04，Booster06，Booster Chronicle，Booster R3
		 效果: 描述：被称为海洋之主的海龙。掀起海啸吞噬一切。
		}
		*/
		Id:       507,
		Password: "76634149",
		Name:     "海龙神",                  // "Kairyu-Shin"  "海竜神"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Water,      // 水
		Lr:       ygo.LR_Seaserpent, // 海龙
		Attack:   1800,
		Defense:  1500,
		IsValid:  true,
	})

	/*140*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 508
		 中文名: 龙僵尸
		 日文名: ドラゴン·ゾンビ
		 英文名: Dragon Zombie
		 卡片种类: 通常怪兽
		 卡片密码: 66672569
		 使用限制: 无限制
		 种族: 不死
		 属性: 暗
		 星级: 3
		 攻击力: 1600
		 防御力: 0
		 罕见度: 平卡N
		 卡包: EX-R(EX)，DL04，Booster05，Booster Chronicle，Booster R3
		 效果: 描述：被魔力唤醒的龙，吐出的气息能使物体腐烂。
		}
		*/
		Id:       508,
		Password: "66672569",
		Name:     "龙僵尸",                  // "Dragon Zombie"  "ドラゴン·ゾンビ"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Dark,   // 暗
		Lr:       ygo.LR_Zombie, // 不死
		Attack:   1600,
		Defense:  0,
		IsValid:  true,
	})

	/*141*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 510
		 中文名: 神灯魔人
		 日文名: ランプの魔人
		 英文名: Lord of the Lamp
		 卡片种类: 通常怪兽
		 卡片密码: 99510761
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 1400
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: BE02，DL04，Booster05，Booster Chronicle，Booster R3
		 效果: 描述：从神灯里出现的魔灵，服从召唤者的意志。
		}
		*/
		Id:       510,
		Password: "99510761",
		Name:     "神灯魔人",                 // "Lord of the Lamp"  "ランプの魔人"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Dark,  // 暗
		Lr:       ygo.LR_Fiend, // 恶魔
		Attack:   1400,
		Defense:  1200,
		IsValid:  true,
	})

}
