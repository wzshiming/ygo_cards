package cards

import ygo "github.com/wzshiming/ygo_core"

func vol(cardBag *ygo.CardVersion) {

	/*0*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 1356
		 调整: [异次元的战士]
		<異次元の戦士>
		[2010/09/08]

		●这张卡和怪兽进行战斗时，那只怪兽和这张卡从游戏中除外。
		◇诱发效果（进入连锁）
		◇必须发动
		◇不取对象
		◇效果发动时点同反转怪兽在伤害步骤发动反转效果的时点（具体请参照伤害步骤相关帖）
		◇效果处理时这张卡或和这张卡战斗的怪兽不在场上存在的场合，剩下的那只从游戏中除外
		◇效果处理时这张卡与进行战斗的怪兽都不在场上存在的场合，这个效果不适用
		 中文名: 异次元的战士
		 日文名: 異次元の戦士
		 英文名: D.D. Warrior
		 卡片种类: 效果怪兽
		 卡片密码: 37043180
		 使用限制: 无限制
		 种族: 战士
		 属性: 地
		 星级: 4
		 攻击力: 1200
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: VOL07，SD14，SD17，SJ2
		 效果: 这张卡和怪兽进行过战斗时，那只怪兽和这张卡从游戏中除外。
		}
		*/
		Id:       1356,
		Password: "37043180",
		Name:     "异次元的战士",             // "D.D. Warrior"  "異次元の戦士"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   4,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Attack:  1200,
		Defense: 1000,
		Initialize: func(ca *ygo.Card) bool {
			ca.AddEvent(ygo.Fought, func(c *ygo.Card) {
				ca.Dispatch(ygo.Removed)
				c.Dispatch(ygo.Removed, ca)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*1*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 159
		 调整: [传说之剑]
		<伝説の剣>
		[2010/09/08]

		●战士族怪兽才能装备。1只装备怪兽的攻击力及守备力上升300分。
		◇发动时选择场上1只符合条件的怪兽（取对象）
		◇效果处理时对象怪兽不是战士族的场合，这张卡送去墓地
		◇效果适用中对象怪兽变成战士族以外的种族的场合，这张卡破坏
		 中文名: 传说之剑
		 日文名: 伝説の剣
		 英文名: Legendary Sword
		 卡片种类: 装备魔法
		 卡片密码: 61854111
		 使用限制: 无限制
		 罕见度: 平卡N，银字R
		 卡包: LB，VOL01，DL02，Booster01
		 效果: 战士族才能装备。1只装备怪兽的攻击力和守备力上升300。
		}
		*/
		Id:       159,
		Password: "61854111",
		Name:     "传说之剑",            // "Legendary Sword"  "伝説の剣"
		Lc:       ygo.LC_EquipMagic, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterEquipMagic(func(c *ygo.Card) bool {
				return c.RaceIsWarrior()
			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() + 300)
				c.SetDefense(c.GetDefense() + 300)

			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() - 300)
				c.SetDefense(c.GetDefense() - 300)

			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*2*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 160
		 调整: [猛兽之齿]
		<猛獣の歯>
		[2010/09/10]

		●兽族才能装备。1只装备怪兽的攻击力与守备力上升300分。
		◇发动时选择1只符合条件的怪兽（取对象）
		◇效果处理时装备怪兽种族变为兽族以外的场合，这张卡送去墓地
		◇装备状态中对象怪兽变为兽族以外的场合，这张卡破坏
		 中文名: 猛兽之齿
		 日文名: 猛獣の歯
		 英文名: Beast Fangs
		 卡片种类: 装备魔法
		 卡片密码: 46009906
		 使用限制: 无限制
		 罕见度: 平卡N，银字R
		 卡包: LB，VOL01，DL02，Booster01
		 效果: 兽族才能装备。1只装备怪兽的攻击力和守备力上升300。
		}
		*/
		Id:       160,
		Password: "46009906",
		Name:     "猛兽之齿",            // "Beast Fangs"  "猛獣の歯"
		Lc:       ygo.LC_EquipMagic, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterEquipMagic(func(c *ygo.Card) bool {
				return c.RaceIsBeast()
			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() + 300)
				c.SetDefense(c.GetDefense() + 300)

			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() - 300)
				c.SetDefense(c.GetDefense() - 300)

			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*3*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 161
		 调整: [紫水晶]
		<紫水晶>
		[2010/09/10]

		●不死族才能装备。1只装备怪兽的攻击力与守备力上升300分。
		◇发动时选择1只符合条件的怪兽（取对象）
		◇效果处理时装备怪兽种族变为不死族以外的场合，这张卡送去墓地
		◇装备状态中对象怪兽变为不死族以外的场合，这张卡破坏
		 中文名: 紫水晶
		 日文名: 紫水晶
		 英文名: Violet Crystal
		 卡片种类: 装备魔法
		 卡片密码: 15052462
		 使用限制: 无限制
		 罕见度: 平卡N，银字R
		 卡包: LB，VOL01，DL02，Booster01
		 效果: 不死族才能装备。1只装备怪兽的攻击力和守备力上升300。
		}
		*/
		Id:       161,
		Password: "15052462",
		Name:     "紫水晶",             // "Violet Crystal"  "紫水晶"
		Lc:       ygo.LC_EquipMagic, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterEquipMagic(func(c *ygo.Card) bool {
				return c.RaceIsZombie()
			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() + 300)
				c.SetDefense(c.GetDefense() + 300)

			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() - 300)
				c.SetDefense(c.GetDefense() - 300)

			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*4*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 162
		 调整: [秘术之书]
		<秘術の書>
		[2010/09/08]

		●魔法师族怪兽才能装备。1只装备怪兽的攻击力及守备力上升300分。
		◇发动时选择场上1只符合条件的怪兽（取对象）
		◇效果处理时对象怪兽不是魔法师族的场合，这张卡送去墓地
		◇效果适用中对象怪兽变成魔法师族以外的种族的场合，这张卡破坏
		 中文名: 秘术之书
		 日文名: 秘術の書
		 英文名: Book of Secret Arts
		 卡片种类: 装备魔法
		 卡片密码: 91595718
		 使用限制: 无限制
		 罕见度: 平卡N，银字R
		 卡包: LB，EX-R(EX)，VOL01，DL02，Booster01，TP21
		 效果: 魔法师族怪兽才能装备。装备怪兽的攻击力·守备力上升300。
		}
		*/
		Id:       162,
		Password: "91595718",
		Name:     "秘术之书",            // "Book of Secret Arts"  "秘術の書"
		Lc:       ygo.LC_EquipMagic, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterEquipMagic(func(c *ygo.Card) bool {
				return c.RaceIsSpellCaster()
			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() + 300)
				c.SetDefense(c.GetDefense() + 300)

			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() - 300)
				c.SetDefense(c.GetDefense() - 300)

			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*5*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 163
		 调整: [波塞冬之力]
		<ポセイドンの力>
		[2010/09/10]

		●水族才能装备。1只装备怪兽的攻击力与守备力上升300分。
		◇发动时选择1只符合条件的怪兽（取对象）
		◇效果处理时装备怪兽种族变为水族以外的场合，这张卡送去墓地
		◇装备状态中对象怪兽变为水族以外的场合，这张卡破坏
		 中文名: 波塞冬之力
		 日文名: ポセイドンの力
		 英文名: Power of Kaishin
		 卡片种类: 装备魔法
		 卡片密码: 77027445
		 使用限制: 无限制
		 罕见度: 平卡N，银字R
		 卡包: LB，VOL01，DL02，Booster01
		 效果: 水族才能装备。1只装备怪兽的攻击力和守备力上升300。
		}
		*/
		Id:       163,
		Password: "77027445",
		Name:     "波塞冬之力",           // "Power of Kaishin"  "ポセイドンの力"
		Lc:       ygo.LC_EquipMagic, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterEquipMagic(func(c *ygo.Card) bool {
				return c.RaceIsAqua()
			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() + 300)
				c.SetDefense(c.GetDefense() + 300)

			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() - 300)
				c.SetDefense(c.GetDefense() - 300)

			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*6*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 171
		 调整: [黑洞]
		<ブラック·ホール>
		[14/04/27]

		●①：场上的怪兽全部破坏。
		◇不取对象
		 中文名: 黑洞
		 日文名: ブラック·ホール
		 英文名: Dark Hole
		 卡片种类: 通常魔法
		 卡片密码: 53129443
		 使用限制: 限制卡
		 罕见度: 金字UR，面闪SR，平卡N，黄金GR，爆闪PR
		 卡包: LB，BE01，EX-R(EX)，VOL01，DL02，YU，JY，KA，TU04，TU05，GS03，SD23，SDWA，ST12，DS13，ST14
		 效果: ①：场上的怪兽全部破坏。
		}
		*/
		Id:       171,
		Password: "53129443",
		Name:     "黑洞",                 // "Dark Hole"  "ブラック·ホール"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				cs := ygo.NewCards(pl.Mzone, tar.Mzone)
				cs.ForEach(func(c *ygo.Card) bool {
					c.Dispatch(ygo.Destroy, ca)
					return true
				})
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*7*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 173
		 调整: [红色药剂]
		<レッド·ポーション>
		[2010/09/08]

		●自己回复500基本分。
		◇「精灵之镜/精霊の鏡」可以对应这张卡的发动而发动
		 中文名: 红色药剂
		 日文名: レッド·ポーション
		 英文名: Red Medicine
		 卡片种类: 通常魔法
		 卡片密码: 38199696
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: LB，BE01，YSD01，VOL01，DL02
		 效果: 自己回复500基本分。
		}
		*/
		Id:       173,
		Password: "38199696",
		Name:     "红色药剂",               // "Red Medicine"  "レッド·ポーション"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				pl.ChangeHp(500)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*8*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 174
		 调整: [火星]
		<火の粉>
		[2010/09/08]

		●给予对方基本分200分伤害。
		◇「精灵之镜/精霊の鏡」可以对应这张卡的发动而发动
		 中文名: 火星
		 日文名: 火の粉
		 英文名: Sparks
		 卡片种类: 通常魔法
		 卡片密码: 76103675
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: LB，VOL01，DL02
		 效果: 给予对方基本分200分伤害。
		}
		*/
		Id:       174,
		Password: "76103675",
		Name:     "火星",                 // "Sparks"  "火の粉"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				tar.ChangeHp(-200)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*9*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 176
		 调整: [地割]
		<地割れ>
		[2010/09/08]

		●对方场上表侧表示存在的攻击力最低的1只怪兽破坏。
		◇在这张卡的效果处理时判断“对方场上表侧表示存在的攻击力最低的1只怪兽”（不取对象）
		◇效果处理时，若符合条件的怪兽有2只以上存在，这张卡的发动者在这时选择其中之一破坏
		 中文名: 地割
		 日文名: 地割れ
		 英文名: Fissure
		 卡片种类: 通常魔法
		 卡片密码: 66788016
		 使用限制: 无限制
		 罕见度: 平卡N，银字R，黄金GR，金字UR
		 卡包: LB，BE01，BE01，EX-R(EX)，YSD01，VOL01，DL02，SD16，GS02，YU，PE，SY2，YSD06，DB12，ST12，DS14
		 效果: 对方场上表侧表示存在的1只攻击力最低的怪兽破坏。
		}
		*/
		Id:       176,
		Password: "66788016",
		Name:     "地割",                 // "Fissure"  "地割れ"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {

				pl := ca.GetSummoner()
				tar := pl.GetTarget()

				i := -1
				tar.Mzone.ForEach(func(c *ygo.Card) bool {
					if c.IsFaceUp() && (i == -1 || i > c.GetAttack()) {
						i = c.GetAttack()
					}
					return true
				})

				cs := tar.Mzone.Find(func(c *ygo.Card) bool {
					if c.IsFaceUp() && i == c.GetAttack() {
						return true
					}
					return false
				})

				if cs.Len() == 1 {
					cs.EndPop().Dispatch(ygo.Destroy, ca)
				} else if cs.Len() > 1 {
					if c := pl.SelectForWarn(cs); c != nil {
						c.Dispatch(ygo.Destroy, ca)
					}
				}
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*10*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 177
		 调整: [落穴]
		<落とし穴>
		[14/04/27]

		●①：对方对攻击力1000以上的怪兽的召唤·反转召唤成功时，以那1只怪兽为对象才能发动。那只攻击力1000以上的怪兽破坏。
		◇取对象
		◇进行召唤·反转召唤的怪物在此卡发动前就从场上离开或变为里侧表示的场合（举例：对方召唤鸟兽族怪物，在召唤成功时行使优先权将其解放发动《神鸟攻击》），符合条件的对象不存在，此卡不能发动
		◇效果处理时对象怪兽不在场上表侧表示存在的场合（举例：对方连锁此卡发动《月之书》将对象变成里侧表示），此卡的效果不适用
		◇效果处理时对象怪兽的攻击力不足1000的场合（举例：对方连锁此卡发动《收缩》将对象攻击力变成900），此卡的效果不适用
		 中文名: 落穴
		 日文名: 落とし穴
		 英文名: Trap Hole
		 卡片种类: 通常陷阱
		 卡片密码: 04206964
		 使用限制: 无限制
		 罕见度: 银字R，平卡N，黄金GR
		 卡包: LB，BE01，EX-R(EX)，YSD02，VOL01，DL02，YSD03，DT04，YSD05，JY，SY2，YSD06，GS03，ST12，ST14
		 效果: ①：对方对攻击力1000以上的怪兽的召唤·反转召唤成功时，以那1只怪兽为对象才能发动。那只攻击力1000以上的怪兽破坏。

		}
		*/
		Id:       177,
		Password: "04206964",
		Name:     "落穴",                // "Trap Hole"  "落とし穴"
		Lc:       ygo.LC_OrdinaryTrap, // 通常陷阱

		Initialize: func(ca *ygo.Card) bool {

			e := func(c *ygo.Card) {
				if c.GetAttack() >= 1000 && c.GetSummoner() != ca.GetSummoner() {
					ca.PushChain(func() {
						c.Dispatch(ygo.Destroy, ca)
					})
				}
			}
			ca.RegisterOrdinaryTrap(ygo.Summon, e)
			ca.RegisterOrdinaryTrap(ygo.SummonFlip, e)
			return true
		}, // 初始
		IsValid: true,
	})

	/*11*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 178
		 调整: [融合]
		<融合>
		[2010/09/08]

		●从手卡·自己场上，把融合怪兽卡决定的融合素材怪兽送去墓地，从额外卡组特殊召唤1只融合怪兽。
		◇可以只用手卡的或只用自己场上的融合素材怪兽，也可以结合起来用
		◇「王宫的弹压/王宮の弾圧」只能对应这张卡的发动而发动，不能对应这张卡效果处理时的特殊召唤而发动
		◇效果处理时选择作为融合素材的怪兽送去墓地
		◇效果处理时，融合怪兽的融合素材不足时，这个效果不适用
		◇承上述，那个场合对方可以要求查看自己的手牌、场上覆盖的怪兽、额外卡组
		◇这张卡的发动被无效的场合，对方不能要求查看自己的手牌、场上覆盖的怪兽、额外卡组
		 中文名: 融合
		 日文名: 融合
		 英文名: Polymerization
		 卡片种类: 通常魔法
		 卡片密码: 24094653
		 使用限制: 无限制
		 罕见度: 平卡N，面闪SR
		 卡包: LB，DP01，BE01，VOL06，DL02，Starter Box，DPYG，DT07，DPKB，YU，JY，KA，PE，SY2，15AY，AT06，SD27
		 效果: 从手卡·自己场上把融合怪兽卡决定的融合素材怪兽送去墓地，那1只融合怪兽从额外卡组特殊召唤。
		}
		*/
		Id:       178,
		Password: "24094653",
		Name:     "融合",                 // "Polymerization"  "融合"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				if c := pl.SelectForWarn(pl.Extra); c != nil {
					if c.IsFusionMonster() {
						c.Dispatch(ygo.SummonFusion, ca)
					}
				}
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*12*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 184
		 中文名: 龙骑士 盖亚
		 日文名: 竜騎士ガイア
		 英文名: Gaia the Dragon Champion
		 卡片种类: 融合怪兽
		 卡片密码: 66889139
		 使用限制: 无限制
		 种族: 龙
		 属性: 风
		 星级: 7
		 攻击力: 2600
		 防御力: 2100
		 罕见度: 立体UTR，面闪SR，银字R，金字UR
		 卡包: PG，BE01，VOL03，DL02，309，Booster R2，15AY
		 效果: 融合：「暗黑骑士 盖亚」＋「诅咒之龙」。
		}
		*/
		Id:       184,
		Password: "66889139",
		Name:     "龙骑士 盖亚",             // "Gaia the Dragon Champion"  "竜騎士ガイア"
		Lc:       ygo.LC_FusionMonster, // 融合怪兽

		Level:   7,
		La:      ygo.LA_Wind,   // 风
		Lr:      ygo.LR_Dragon, // 龙
		Attack:  2600,
		Defense: 2100,

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterFusionMonster("暗黑骑士 盖亚", "诅咒之龙")
			return true
		}, // 初始
		IsValid: true,
	})

	/*13*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 188
		 调整: [猎卡死神]
		<カードを狩る死神>
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
		 日文名: カードを狩る死神
		 英文名: Reaper of the Cards
		 卡片种类: 效果怪兽
		 卡片密码: 33066139
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 5
		 攻击力: 1380
		 防御力: 1930
		 罕见度: 平卡N
		 卡包: PG，BE01，VOL03，DL02
		 效果: 反转：选择场上存在的1张陷阱卡破坏。选择的卡是盖放的场合，把那张卡翻开确认，是陷阱卡则破坏。魔法卡的场合回到原状。
		}
		*/
		Id:       188,
		Password: "33066139",
		Name:     "猎卡死神",               // "Reaper of the Cards"  "カードを狩る死神"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   5,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Attack:  1380,
		Defense: 1930,
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

	/*14*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 194
		 调整: [暗能量]
		<闇·エネルギー>
		[2010/09/08]

		●恶魔族怪兽才能装备。1只装备怪兽的攻击力及守备力上升300分。
		◇发动时选择场上1只符合条件的怪兽（取对象）
		◇效果处理时对象怪兽不是恶魔族的场合，这张卡送去墓地
		◇效果适用中对象怪兽变成恶魔族以外的种族的场合，这张卡破坏
		 中文名: 暗能量
		 日文名: 闇·エネルギー
		 英文名: Dark Energy
		 卡片种类: 装备魔法
		 卡片密码: 04614116
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: PG，EX-R(EX)，VOL02，DL02，Booster R2
		 效果: 恶魔族才能装备。1只装备怪兽的攻击力·守备力上升300。
		}
		*/
		Id:       194,
		Password: "04614116",
		Name:     "暗能量",             // "Dark Energy"  "闇·エネルギー"
		Lc:       ygo.LC_EquipMagic, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterEquipMagic(func(c *ygo.Card) bool {
				return c.RaceIsFiend()
			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() + 300)
				c.SetDefense(c.GetDefense() + 300)

			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() - 300)
				c.SetDefense(c.GetDefense() - 300)

			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*15*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 195
		 调整: [镭射炮机甲铠]
		<レーザー砲機甲鎧>
		[2010/09/08]

		（这张卡在规则上不当作「机甲/マシンナーズ」卡使用）
		●昆虫族怪兽才能装备。1只装备怪兽的攻击力及守备力上升300分。
		◇发动时选择场上1只符合条件的怪兽（取对象）
		◇效果处理时对象怪兽不是昆虫族的场合，这张卡送去墓地
		◇效果适用中对象怪兽变成昆虫族以外的种族的场合，这张卡破坏
		 中文名: 镭射炮机甲铠
		 日文名: レーザー砲機甲鎧
		 英文名: Laser Cannon Armor
		 卡片种类: 装备魔法
		 卡片密码: 77007920
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: PG，VOL02，DL02，TP19
		 效果: （这张卡在规则上不当作「机甲」卡使用）昆虫族才能装备。1只装备怪兽的攻击力·守备力上升300。
		}
		*/
		Id:       195,
		Password: "77007920",
		Name:     "镭射炮机甲铠",          // "Laser Cannon Armor"  "レーザー砲機甲鎧"
		Lc:       ygo.LC_EquipMagic, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterEquipMagic(func(c *ygo.Card) bool {
				return c.RaceIsInsect()
			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() + 300)
				c.SetDefense(c.GetDefense() + 300)

			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() - 300)
				c.SetDefense(c.GetDefense() - 300)

			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*16*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 196
		 调整: [魔菌]
		<魔菌>
		[2010/09/08]

		●植物族怪兽才能装备。1只装备怪兽的攻击力及守备力上升300分。
		◇发动时选择场上1只符合条件的怪兽（取对象）
		◇效果处理时对象怪兽不是植物族的场合，这张卡送去墓地
		◇效果适用中对象怪兽变成植物族以外的种族的场合，这张卡破坏
		 中文名: 魔菌
		 日文名: 魔菌
		 英文名: Vile Germs
		 卡片种类: 装备魔法
		 卡片密码: 39774685
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: PG，VOL02，DL02
		 效果: 植物族才能装备。1只装备怪兽的攻击力·守备力上升300。
		}
		*/
		Id:       196,
		Password: "39774685",
		Name:     "魔菌",              // "Vile Germs"  "魔菌"
		Lc:       ygo.LC_EquipMagic, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterEquipMagic(func(c *ygo.Card) bool {
				return c.RaceIsPlant()
			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() + 300)
				c.SetDefense(c.GetDefense() + 300)

			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() - 300)
				c.SetDefense(c.GetDefense() - 300)

			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*17*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 197
		 调整: [银之弓矢]
		<銀の弓矢>
		[2010/09/08]

		●天使族怪兽才能装备。1只装备怪兽的攻击力及守备力上升300分。
		◇发动时选择场上1只符合条件的怪兽（取对象）
		◇效果处理时对象怪兽不是天使族的场合，这张卡送去墓地
		◇效果适用中对象怪兽变成天使族以外的种族的场合，这张卡破坏
		 中文名: 银之弓矢
		 日文名: 銀の弓矢
		 英文名: Silver Bow and Arrow
		 卡片种类: 装备魔法
		 卡片密码: 01557499
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: PG，VOL03，DL02，TP15
		 效果: 天使族怪兽才能装备。装备怪兽的攻击力·守备力上升300。
		}
		*/
		Id:       197,
		Password: "01557499",
		Name:     "银之弓矢",            // "Silver Bow and Arrow"  "銀の弓矢"
		Lc:       ygo.LC_EquipMagic, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterEquipMagic(func(c *ygo.Card) bool {
				return c.RaceIsAngel()
			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() + 300)
				c.SetDefense(c.GetDefense() + 300)

			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() - 300)
				c.SetDefense(c.GetDefense() - 300)

			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*18*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 198
		 调整: [龙之秘宝]
		<ドラゴンの秘宝>
		[2010/09/08]

		●龙族怪兽才能装备。1只装备怪兽的攻击力及守备力上升300分。
		◇发动时选择场上1只符合条件的怪兽（取对象）
		◇效果处理时对象怪兽不是龙族的场合，这张卡送去墓地
		◇效果适用中对象怪兽变成龙族以外的种族的场合，这张卡破坏
		 中文名: 龙之秘宝
		 日文名: ドラゴンの秘宝
		 英文名: Dragon Treasure
		 卡片种类: 装备魔法
		 卡片密码: 01435851
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: PG，VOL03，DL02，Booster R2
		 效果: 龙族才能装备。1只装备怪兽的攻击力·守备力上升300。
		}
		*/
		Id:       198,
		Password: "01435851",
		Name:     "龙之秘宝",            // "Dragon Treasure"  "ドラゴンの秘宝"
		Lc:       ygo.LC_EquipMagic, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterEquipMagic(func(c *ygo.Card) bool {
				return c.RaceIsDragon()
			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() + 300)
				c.SetDefense(c.GetDefense() + 300)

			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() - 300)
				c.SetDefense(c.GetDefense() - 300)

			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*19*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 199
		 调整: [电击鞭]
		<電撃鞭>
		[2010/09/08]

		●雷族怪兽才能装备。1只装备怪兽的攻击力及守备力上升300分。
		◇发动时选择场上1只符合条件的怪兽（取对象）
		◇效果处理时对象怪兽不是雷族的场合，这张卡送去墓地
		◇效果适用中对象怪兽变成雷族以外的种族的场合，这张卡破坏
		 中文名: 电击鞭
		 日文名: 電撃鞭
		 英文名: Electro-Whip
		 卡片种类: 装备魔法
		 卡片密码: 37820550
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: PG，VOL03，DL02
		 效果: 雷族才能装备。1只装备怪兽的攻击力·守备力上升300。
		}
		*/
		Id:       199,
		Password: "37820550",
		Name:     "电击鞭",             // "Electro-Whip"  "電撃鞭"
		Lc:       ygo.LC_EquipMagic, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterEquipMagic(func(c *ygo.Card) bool {
				return c.RaceIsThunder()
			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() + 300)
				c.SetDefense(c.GetDefense() + 300)

			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() - 300)
				c.SetDefense(c.GetDefense() - 300)

			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*20*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 200
		 调整: [魔性之月]
		<魔性の月>
		[2010/09/08]

		●兽战士族怪兽才能装备。1只装备怪兽的攻击力及守备力上升300分。
		◇发动时选择场上1只符合条件的怪兽（取对象）
		◇效果处理时对象怪兽不是兽战士族的场合，这张卡送去墓地
		◇效果适用中对象怪兽变成兽战士族以外的种族的场合，这张卡破坏
		 中文名: 魔性之月
		 日文名: 魔性の月
		 英文名: Mystical Moon
		 卡片种类: 装备魔法
		 卡片密码: 36607978
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: PG，VOL03，DL02，15AY
		 效果: 兽战士族才能装备。1只装备怪兽的攻击力·守备力上升300。
		}
		*/
		Id:       200,
		Password: "36607978",
		Name:     "魔性之月",            // "Mystical Moon"  "魔性の月"
		Lc:       ygo.LC_EquipMagic, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterEquipMagic(func(c *ygo.Card) bool {
				return c.RaceIsBeastWarror()
			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() + 300)
				c.SetDefense(c.GetDefense() + 300)

			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() - 300)
				c.SetDefense(c.GetDefense() - 300)

			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*21*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 201
		 调整: [「守备」封禁]
		<『守備』封じ>
		[2010/09/08]

		●选择对方场上的1只守备表示怪兽，变为表侧攻击表示。
		◇发动时选择对方场上存在的1只守备表示怪兽（取对象）
		◇效果处理时对象卡不在场上守备表示存在的场合，这个效果不适用
		◇「等级限制B地区/レベル制限Ｂ地区」的效果适用中，这张卡把怪兽变为攻击表示后，由于「等级限制B地区/レベル制限Ｂ地区」的效果立刻变为守备表示
		 中文名: 「守备」封禁
		 日文名: 『守備』封じ
		 英文名: Stop Defense
		 卡片种类: 通常魔法
		 卡片密码: 63102017
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: PG，BE01，VOL03，DL02，Booster R2
		 效果: 选择1只对方场上的守备表示的怪兽变表侧攻击表示。
		}
		*/
		Id:       201,
		Password: "63102017",
		Name:     "「守备」封禁",             // "Stop Defense"  "『守備』封じ"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				if c := pl.SelectForWarn(tar.Mzone, func(c0 *ygo.Card) bool {
					return c0.IsDefense()
				}); c != nil {
					c.SetFaceUpAttack()
				}
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*22*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 202
		 调整: [机械改造工厂]
		<機械改造工場>
		[2010/09/08]

		●机械族怪兽才能装备。1只装备怪兽的攻击力及守备力上升300分。
		◇发动时选择场上1只符合条件的怪兽（取对象）
		◇效果处理时对象怪兽不是机械族的场合，这张卡送去墓地
		◇效果适用中对象怪兽变成机械族以外的种族的场合，这张卡破坏
		 中文名: 机械改造工厂
		 日文名: 機械改造工場
		 英文名: Machine Conversion Factory
		 卡片种类: 装备魔法
		 卡片密码: 25769732
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: PG，VOL02，DL02
		 效果: 只有机械族怪兽可以装备。装备怪兽的攻击力·守备力上升300。
		}
		*/
		Id:       202,
		Password: "25769732",
		Name:     "机械改造工厂",          // "Machine Conversion Factory"  "機械改造工場"
		Lc:       ygo.LC_EquipMagic, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterEquipMagic(func(c *ygo.Card) bool {
				return c.RaceIsMachine()
			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() + 300)
				c.SetDefense(c.GetDefense() + 300)

			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() - 300)
				c.SetDefense(c.GetDefense() - 300)

			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*23*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 203
		 调整: [体温上升]
		<体温の上昇>
		[2010/09/08]

		●恐龙族怪兽才能装备。1只装备怪兽的攻击力及守备力上升300分。
		◇发动时选择场上1只符合条件的怪兽（取对象）
		◇效果处理时对象怪兽不是恐龙族的场合，这张卡送去墓地
		◇效果适用中对象怪兽变成恐龙族以外的种族的场合，这张卡破坏
		 中文名: 体温上升
		 日文名: 体温の上昇
		 英文名: Raise Body Heat
		 卡片种类: 装备魔法
		 卡片密码: 51267887
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: PG，VOL02，DL02
		 效果: 只有恐龙族怪兽可以装备。装备怪兽的攻击力·守备力上升300。
		}
		*/
		Id:       203,
		Password: "51267887",
		Name:     "体温上升",            // "Raise Body Heat"  "体温の上昇"
		Lc:       ygo.LC_EquipMagic, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterEquipMagic(func(c *ygo.Card) bool {
				return c.RaceIsDinosaur()
			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() + 300)
				c.SetDefense(c.GetDefense() + 300)

			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() - 300)
				c.SetDefense(c.GetDefense() - 300)

			})
			return true

		}, // 初始
		IsValid: true,
	})

	/*24*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 204
		 调整: [随风翼]
		<フォロー·ウィンド>
		[2010/09/08]

		●鸟兽族怪兽才能装备。1只装备怪兽的攻击力及守备力上升300分。
		◇发动时选择场上1只符合条件的怪兽（取对象）
		◇效果处理时对象怪兽不是鸟兽族的场合，这张卡送去墓地
		◇效果适用中对象怪兽变成鸟兽族以外的种族的场合，这张卡破坏
		 中文名: 随风翼
		 日文名: フォロー·ウィンド
		 英文名: Follow Wind
		 卡片种类: 装备魔法
		 卡片密码: 98252586
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: PG，VOL03，DL02
		 效果: 只有鸟兽族怪兽可以装备。装备怪兽的攻击力·守备力上升300。
		}
		*/
		Id:       204,
		Password: "98252586",
		Name:     "随风翼",             // "Follow Wind"  "フォロー·ウィンド"
		Lc:       ygo.LC_EquipMagic, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterEquipMagic(func(c *ygo.Card) bool {
				return c.RaceIsWindBeast()
			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() + 300)
				c.SetDefense(c.GetDefense() + 300)

			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() - 300)
				c.SetDefense(c.GetDefense() - 300)

			})
			return true

		}, // 初始
		IsValid: true,
	})

	/*25*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 205
		 调整: [哥布林的秘药]
		<ゴブリンの秘薬>
		[2010/09/08]

		●自己回复600基本分。
		◇「精灵之镜/精霊の鏡」可以对应这张卡的发动而发动
		 中文名: 哥布林的秘药
		 日文名: ゴブリンの秘薬
		 英文名: Goblin's Secret Remedy
		 卡片种类: 通常魔法
		 卡片密码: 11868825
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: PG，VOL02，DL02
		 效果: 自己回复600基本分。
		}
		*/
		Id:       205,
		Password: "11868825",
		Name:     "哥布林的秘药",             // "Goblin's Secret Remedy"  "ゴブリンの秘薬"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				pl.ChangeHp(600)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*26*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 206
		 调整: [火刑]
		<火あぶりの刑>
		[2010/09/08]

		●给予对方基本分600分伤害。
		◇「精灵之镜/精霊の鏡」可以对应这张卡的发动而发动
		 中文名: 火刑
		 日文名: 火あぶりの刑
		 英文名: Final Flame
		 卡片种类: 通常魔法
		 卡片密码: 73134081
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: PG，VOL02，DL02
		 效果: 给予对方基本分600分的伤害。
		}
		*/
		Id:       206,
		Password: "73134081",
		Name:     "火刑",                 // "Final Flame"  "火あぶりの刑"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				tar.ChangeHp(-600)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*27*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 207
		 调整: [光之护封剑]
		<光の護封剣>
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
		 日文名: 光の護封剣
		 英文名: Swords of Revealing Light
		 卡片种类: 通常魔法
		 卡片密码: 72302403
		 使用限制: 无限制
		 罕见度: 面闪SR，黄金GR，平卡N，金字UR
		 卡包: PG，BE01，VOL02，DL02，SD01，SD06，SD07，SD11，GLD01，GS01，DPYG，SD18，YU，YSD06，DB12，ST12，ST13，15AY
		 效果: 对方场上的怪兽全部变成表侧表示。这张卡发动后，用对方回合计算的3回合内继续留在场上。只要这张卡在场上存在，对方场上的怪兽不能攻击宣言。
		}
		*/
		Id:       207,
		Password: "72302403",
		Name:     "光之护封剑",              // "Swords of Revealing Light"  "光の護封剣"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterUnordinaryMagic(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				tar.Mzone.ForEach(func(c *ygo.Card) bool {
					c.SetFaceUp()
					return true
				})
				i := 0

				ca.RegisterGlobalListen(ygo.BP, func(pl0 *ygo.Player) {
					if tar != pl0 {
						return
					}
					tar.Mzone.ForEach(func(c *ygo.Card) bool {
						c.SetNotCanAttack()
						return true
					})
				})

				ca.RegisterGlobalListen(ygo.RoundEnd, func(pl0 *ygo.Player) {
					if tar != pl0 {
						return
					}
					i++
					if i >= 3 {
						ca.Dispatch(ygo.Depleted)
					}
				})
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*28*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 208
		 中文名: 金属龙
		 日文名: メタル·ドラゴン
		 英文名: Metal Dragon
		 卡片种类: 融合怪兽
		 卡片密码: 09293977
		 使用限制: 无限制
		 种族: 机械
		 属性: 风
		 星级: 6
		 攻击力: 1850
		 防御力: 1700
		 罕见度: 平卡N
		 卡包: PG，VOL03，DL02
		 效果: 融合：「钢铁巨神像」＋「下级龙」。
		}
		*/
		Id:       208,
		Password: "09293977",
		Name:     "金属龙",                // "Metal Dragon"  "メタル·ドラゴン"
		Lc:       ygo.LC_FusionMonster, // 融合怪兽

		Level:   6,
		La:      ygo.LA_Wind,    // 风
		Lr:      ygo.LR_Machine, // 机械
		Attack:  1850,
		Defense: 1700,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterFusionMonster("钢铁巨神像", "下级龙")
			return true
		}, // 初始
		IsValid: true,
	})

	/*29*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2080
		 调整: [骸骨天使]
		<スケルエンジェル>
		[14/08/07]

		●反转：自己从卡组抽1张卡。
		◇诱发效果，强制发动，开连锁，不取对象
		◇这张卡被攻击在信息确认时变为表侧的场合，效果延迟到怪兽被破坏确定时之后发动（具体请参考伤害步骤的详细说明）
		◇效果处理时，自己卡组中的卡不足1张的场合，抽卡时决斗败北
		◇「神殿守卫者/神殿を守る者」在场上存在的场合这个效果也发动，效果处理时效果不适用
		 中文名: 骸骨天使
		 日文名: スケルエンジェル
		 英文名: Skelengel
		 卡片种类: 效果怪兽
		 卡片密码: 60694662
		 使用限制: 无限制
		 种族: 天使
		 属性: 光
		 星级: 2
		 攻击力: 900
		 防御力: 400
		 罕见度: 平卡N
		 卡包: YSD01，VOL03，YSD04，DB12，DC01
		 效果: 反转：从自己卡组抽1张卡。
		}
		*/
		Id:       2080,
		Password: "60694662",
		Name:     "骸骨天使",               // "Skelengel"  "スケルエンジェル"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   2,
		La:      ygo.LA_Light, // 光
		Lr:      ygo.LR_Angel, // 天使
		Attack:  900,
		Defense: 400,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterFlip(func() {
				pl := ca.GetSummoner()
				pl.ActionDraw(1)
			})
			return true
		}, // 初始

		IsValid: true,
	})

	/*30*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 211
		 调整: [青衣忍者]
		<青い忍者>
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
		 日文名: 青い忍者
		 英文名: Armed Ninja
		 卡片种类: 效果怪兽
		 卡片密码: 09076207
		 使用限制: 无限制
		 种族: 战士
		 属性: 地
		 星级: 1
		 攻击力: 300
		 防御力: 300
		 罕见度: 银字R，平卡N
		 卡包: PG，BE01，VOL03，DL02
		 效果: 反转：破坏场上的1张魔法卡。选择的卡是盖伏的场合，确认那张卡，如果是魔法卡就破坏。如果选择的卡确认是陷阱卡的场合变回原来的盖伏形式。
		}
		*/
		Id:       211,
		Password: "09076207",
		Name:     "青衣忍者",               // "Armed Ninja"  "青い忍者"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   1,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Attack:  300,
		Defense: 300,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterFlip(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				if c := pl.SelectForWarn(tar.Szone, pl.Szone); c != nil {
					c.SetFaceUp()
					if c.IsMagic() {
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

	/*31*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 212
		 调整: [食人虫]
		<人喰い虫>
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
		 日文名: 人喰い虫
		 英文名: Man-Eater Bug
		 卡片种类: 效果怪兽
		 卡片密码: 54652250
		 使用限制: 无限制
		 种族: 昆虫
		 属性: 地
		 星级: 2
		 攻击力: 450
		 防御力: 600
		 罕见度: 面闪SR，银字R
		 卡包: PG，BE01，EX-R(EX)，YSD02，VOL03，DL02，KA，SK2
		 效果: 反转：选择场上1只怪兽破坏。
		}
		*/
		Id:       212,
		Password: "54652250",
		Name:     "食人虫",                // "Man-Eater Bug"  "人喰い虫"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   2,
		La:      ygo.LA_Earth,  // 地
		Lr:      ygo.LR_Insect, // 昆虫
		Attack:  450,
		Defense: 600,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterFlip(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				if c := pl.SelectForWarn(tar.Mzone, pl.Mzone); c != nil {
					c.Dispatch(ygo.Destroy, ca)
				}
			})
			return true
		}, // 初始

		IsValid: true,
	})

	/*32*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 214
		 调整: [哈尼哈尼]
		<ハネハネ>
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
		 日文名: ハネハネ
		 英文名: Hane-Hane
		 卡片种类: 效果怪兽
		 卡片密码: 07089711
		 使用限制: 无限制
		 种族: 兽
		 属性: 地
		 星级: 2
		 攻击力: 450
		 防御力: 500
		 罕见度: 银字R，平卡N
		 卡包: PG，BE01，EX-R(EX)，VOL03，DL02
		 效果: 反转：选择场上1只怪兽回到原持有人手卡。
		}
		*/
		Id:       214,
		Password: "07089711",
		Name:     "哈尼哈尼",               // "Hane-Hane"  "ハネハネ"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   2,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Attack:  450,
		Defense: 500,

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterFlip(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				if c := pl.SelectForWarn(tar.Mzone, pl.Mzone); c != nil {
					c.ToHand()
				}
			})
			return true
		}, // 初始

		IsValid: true,
	})

	/*33*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 218
		 调整: [魔法除去]
		<魔法除去>
		[2010/09/08]

		●场上的1张魔法卡破坏。选择的卡是覆盖的场合，翻开那张卡确认，是魔法卡就破坏。陷阱卡的场合变回原样。
		◇发动时选择这张卡以外的在自己或对方场上的魔法及陷阱区域存在的1张表侧表示的魔法卡或里侧表示覆盖着的1张卡（取对象）
		◇可以选择作为装备的怪兽卡
		◇可以选择由于「纳祭之魔/サクリファイス」的效果作为装备的「阿匹卜之化神/アポピスの化神」等为对象
		◇效果处理时，对象卡不在场上存在的场合，这个效果不适用
		◇作为怪兽装备的在魔法及陷阱区域存在的里侧怪兽卡在这个效果处理时破坏
		 中文名: 魔法除去
		 日文名: 魔法除去
		 英文名: De-Spell
		 卡片种类: 通常魔法
		 卡片密码: 19159413
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: PG，EX-R(EX)，VOL02，DL02，Booster R2，YU
		 效果: 破坏1张场上的魔法卡。选择的卡盖伏的场合，确认那张卡，是魔法卡就破坏，是陷阱卡就变回原来的盖伏形式。
		}
		*/
		Id:       218,
		Password: "19159413",
		Name:     "魔法除去",               // "De-Spell"  "魔法除去"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				if c := pl.SelectForWarn(tar.Szone, pl.Szone); c != nil {
					c.SetFaceUp()
					if c.IsMagic() {
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

	/*34*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 219
		 调整: [死者苏生]
		<死者蘇生>
		[2010/09/08]

		●从自己或对方墓地选择1只怪兽发动。选择的怪兽在自己场上特殊召唤。
		◇发动时从自己或对方墓地选择1只可以被特殊召唤的怪兽（取对象）
		◇效果处理时对象怪兽不在墓地存在的场合，这个效果不适用
		◇效果处理时自己场上的怪兽区域没有空位的场合，这个效果不适用
		 中文名: 死者苏生
		 日文名: 死者蘇生
		 英文名: Monster Reborn
		 卡片种类: 通常魔法
		 卡片密码: 83764718
		 使用限制: 限制卡
		 罕见度: 面闪SR，金字UR，黄金GR，平卡N，银字R
		 卡包: PG，BE01，EX-R(EX)，VOL02，DL02，GS01，DPYG，YU，KA，PE，SY2，SDM，SD25，15AY
		 效果: 选择自己或者对方的墓地1只怪兽才能发动。选择的怪兽在自己场上特殊召唤。
		}
		*/
		Id:       219,
		Password: "83764718",
		Name:     "死者苏生",               // "Monster Reborn"  "死者蘇生"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				if c := pl.SelectForWarn(tar.Grave, pl.Grave, func(c0 *ygo.Card) bool {
					return c0.IsMonster()
				}); c != nil {
					if c.GetSummoner() != ca.GetSummoner() {
						pl := ca.GetSummoner()
						c.SetSummoner(pl)
					}
					c.Init()
					c.Dispatch(ygo.SummonSpecial, ca)
				}
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*35*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 220
		 调整: [强欲之壶]
		<強欲な壺>
		[2010/09/08]

		●从自己卡组抽2张卡。
		◇自己卡组中的卡不足2张时不能发动
		◇「精灵之镜/精霊の鏡」可以对应这张卡的发动而发动
		◇效果处理时，自己卡组中的卡不足2张的场合，抽卡时决斗败北
		 中文名: 强欲之壶
		 日文名: 強欲な壺
		 英文名: Pot of Greed
		 卡片种类: 通常魔法
		 卡片密码: 55144522
		 使用限制: 禁止卡
		 罕见度: 面闪SR，银字R，平卡N，立体UTR
		 卡包: PG，BE01，VOL03，DL02，SD01，SD02，SD03，SD04，DPKB，YU，JY，SY2，15AY
		 效果: 从自己的卡组抽2张卡。
		}
		*/
		Id:       220,
		Password: "55144522",
		Name:     "强欲之壶",               // "Pot of Greed"  "強欲な壺"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				ca.GetSummoner().ActionDraw(2)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*36*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 221
		 调整: [掘墓的食尸鬼]
		<墓掘りグール>
		[2010/09/08]

		●从对方墓地选择1张或2张的怪兽卡。被选择的卡从游戏中除外。
		◇发动时从对方墓地选择1张或2张的怪兽卡（取对象）
		◇效果处理时若其中1张对象不在墓地存在的场合，把剩余的对象卡从游戏中除外
		◇效果处理时若对象卡全部都不在墓地中存在，这个效果不适用
		 中文名: 掘墓的食尸鬼
		 日文名: 墓掘りグール
		 英文名: Gravedigger Ghoul
		 卡片种类: 通常魔法
		 卡片密码: 82542267
		 使用限制: 无限制
		 罕见度: 银字R
		 卡包: PG，VOL03，DL02
		 效果: 选择对方的墓地1张到2张的怪兽卡。选择的卡从游戏中除外。
		}
		*/
		Id:       221,
		Password: "82542267",
		Name:     "掘墓的食尸鬼",             // "Gravedigger Ghoul"  "墓掘りグール"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				for i := 0; i != 2; i++ {
					if c := pl.SelectForWarn(tar.Grave, func(c0 *ygo.Card) bool {
						return c0.IsMonster()
					}); c != nil {
						c.ToRemoved()
					}
				}
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*37*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2282
		 中文名: 炭烧战士
		 日文名: カルボナーラ戦士
		 英文名: Karbonala Warrior
		 卡片种类: 融合怪兽
		 卡片密码: 54541900
		 使用限制: 无限制
		 种族: 战士
		 属性: 地
		 星级: 4
		 攻击力: 1500
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: PG，VOL02，TP02
		 效果: 融合：「磁力战士1号」＋「磁力战士2号」。
		}
		*/
		Id:       2282,
		Password: "54541900",
		Name:     "炭烧战士",               // "Karbonala Warrior"  "カルボナーラ戦士"
		Lc:       ygo.LC_FusionMonster, // 融合怪兽

		Level:   4,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Attack:  1500,
		Defense: 1200,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterFusionMonster("磁力战士1号", "磁力战士2号")
			return true
		}, // 初始
		IsValid: true,
	})

	/*38*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2285
		 中文名: 混沌男巫
		 日文名: カオス·ウィザード
		 英文名: Kamionwizard
		 卡片种类: 融合怪兽
		 卡片密码: 41544074
		 使用限制: 无限制
		 种族: 魔法师
		 属性: 暗
		 星级: 4
		 攻击力: 1300
		 防御力: 1100
		 罕见度: 平卡N
		 卡包: VOL02，TP03
		 效果: 融合：「圣精灵」＋「黑魔族的窗帘」。
		}
		*/
		Id:       2285,
		Password: "41544074",
		Name:     "混沌男巫",               // "Kamionwizard"  "カオス·ウィザード"
		Lc:       ygo.LC_FusionMonster, // 融合怪兽

		Level:   4,
		La:      ygo.LA_Dark,        // 暗
		Lr:      ygo.LR_SpellCaster, // 魔法师
		Attack:  1300,
		Defense: 1100,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterFusionMonster("圣精灵", "黑魔族的窗帘")
			return true
		}, // 初始
		IsValid: true,
	})

	/*39*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2287
		 中文名: 奇迹鸟
		 日文名: マブラス
		 英文名: Mavelus
		 卡片种类: 融合怪兽
		 卡片密码: 59036972
		 使用限制: 无限制
		 种族: 鸟兽
		 属性: 风
		 星级: 4
		 攻击力: 1300
		 防御力: 900
		 罕见度: 平卡N
		 卡包: VOL02，TP17
		 效果: 融合：「大炮鸟」＋「邪炎之翼」。
		}
		*/
		Id:       2287,
		Password: "59036972",
		Name:     "奇迹鸟",                // "Mavelus"  "マブラス"
		Lc:       ygo.LC_FusionMonster, // 融合怪兽

		Level:   4,
		La:      ygo.LA_Wind,      // 风
		Lr:      ygo.LR_WindBeast, // 鸟兽
		Attack:  1300,
		Defense: 900,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterFusionMonster("大炮鸟", "邪炎之翼")
			return true
		}, // 初始
		IsValid: true,
	})

	/*40*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2290
		 中文名: 不死战士
		 日文名: アンデット·ウォーリアー
		 英文名: Zombie Warrior
		 卡片种类: 融合怪兽
		 卡片密码: 31339260
		 使用限制: 无限制
		 种族: 不死
		 属性: 暗
		 星级: 3
		 攻击力: 1200
		 防御力: 900
		 罕见度: 平卡N
		 卡包: VOL02，TP16
		 效果: 融合：「白骨」＋「格斗战士 阿提米特」。
		}
		*/
		Id:       2290,
		Password: "31339260",
		Name:     "不死战士",               // "Zombie Warrior"  "アンデット·ウォーリアー"
		Lc:       ygo.LC_FusionMonster, // 融合怪兽

		Level:   3,
		La:      ygo.LA_Dark,   // 暗
		Lr:      ygo.LR_Zombie, // 不死
		Attack:  1200,
		Defense: 900,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterFusionMonster("白骨", "格斗战士 阿提米特")
			return true
		}, // 初始
		IsValid: true,
	})

	/*41*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2291
		 中文名: 魔装骑士 德拉格尼斯
		 日文名: 魔装騎士ドラゴネス
		 英文名: Dragoness the Wicked Knight
		 卡片种类: 融合怪兽
		 卡片密码: 70681994
		 使用限制: 无限制
		 种族: 战士
		 属性: 风
		 星级: 3
		 攻击力: 1200
		 防御力: 900
		 罕见度: 平卡N
		 卡包: PG，VOL02，TP09
		 效果: 融合：「铠甲剑尾战士」＋「独眼盾龙」。
		}
		*/
		Id:       2291,
		Password: "70681994",
		Name:     "魔装骑士 德拉格尼斯",         // "Dragoness the Wicked Knight"  "魔装騎士ドラゴネス"
		Lc:       ygo.LC_FusionMonster, // 融合怪兽

		Level:   3,
		La:      ygo.LA_Wind,    // 风
		Lr:      ygo.LR_Warrior, // 战士
		Attack:  1200,
		Defense: 900,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterFusionMonster("神鱼", "舌鱼")
			return true
		}, // 初始
		IsValid: true,
	})

	/*42*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2328
		 中文名: 潜伏深海的鲨鱼
		 日文名: 深海に潜むサメ
		 英文名: Deepsea Shark
		 卡片种类: 融合怪兽
		 卡片密码: 28593363
		 使用限制: 无限制
		 种族: 鱼
		 属性: 水
		 星级: 5
		 攻击力: 1900
		 防御力: 1600
		 罕见度: 平卡N
		 卡包: RB，VOL04
		 效果: 融合：「神鱼」＋「舌鱼」。
		}
		*/
		Id:       2328,
		Password: "28593363",
		Name:     "潜伏深海的鲨鱼",            // "Deepsea Shark"  "深海に潜むサメ"
		Lc:       ygo.LC_FusionMonster, // 融合怪兽

		Level:   5,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_Fish,  // 鱼
		Attack:  1900,
		Defense: 1600,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterFusionMonster("神鱼", "舌鱼")
			return true
		}, // 初始
		IsValid: true,
	})

	/*43*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2329
		 中文名: 尖刺龙
		 日文名: プラグティカル
		 英文名: Pragtical
		 卡片种类: 融合怪兽
		 卡片密码: 33691040
		 使用限制: 无限制
		 种族: 恐龙
		 属性: 地
		 星级: 5
		 攻击力: 1900
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: VOL04
		 效果: 融合：「虎纹龙」＋「火焰毒蛇」。
		}
		*/
		Id:       2329,
		Password: "33691040",
		Name:     "尖刺龙",                // "Pragtical"  "プラグティカル"
		Lc:       ygo.LC_FusionMonster, // 融合怪兽

		Level:   5,
		La:      ygo.LA_Earth,    // 地
		Lr:      ygo.LR_Dinosaur, // 恐龙
		Attack:  1900,
		Defense: 1500,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterFusionMonster("虎纹龙", "火焰毒蛇")
			return true
		}, // 初始
		IsValid: true,
	})

	/*44*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 233
		 调整: [飞蛾幼虫]
		<ラーバモス>
		[2010/09/08]

		●这张卡不能通常召唤。装备有「进化之茧」的「幼虫宝宝」2回合后（计算自己回合）作为祭品来召唤。
		◇召唤规则
		◇通过正规方式出场后，可以用其他卡的效果从墓地或除外状态特殊召唤
		 中文名: 飞蛾幼虫
		 日文名: ラーバモス
		 英文名: Larvae Moth
		 卡片种类: 效果怪兽
		 卡片密码: 87756343
		 使用限制: 无限制
		 种族: 昆虫
		 属性: 地
		 星级: 2
		 攻击力: 500
		 防御力: 400
		 罕见度: 平卡N
		 卡包: RB，BE01，VOL05，DL02
		 效果: 这张卡通常召唤不能。需要装备了「进化之茧」2回合（数自己的回合数）的「幼虫宝宝」做祭品特殊召唤上场。
		}
		*/
		Id:       233,
		Password: "87756343",
		Name:     "飞蛾幼虫",               // "Larvae Moth"  "ラーバモス"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   2,
		La:      ygo.LA_Earth,  // 地
		Lr:      ygo.LR_Insect, // 昆虫
		Attack:  500,
		Defense: 400,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*45*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 235
		 调整: [鹰身女郎三姐妹]
		<ハーピィ·レディ三姉妹>
		[2010/09/08]

		●这张卡不能通常召唤。用「万华镜－华丽的分身－」的效果来特殊召唤。
		◇召唤规则
		◇通过正规方式出场后，可以从墓地或除外状态被其他卡的效果特殊召唤
		 中文名: 鹰身女郎三姐妹
		 日文名: ハーピィ·レディ三姉妹
		 英文名: Harpie Lady Sisters
		 卡片种类: 效果怪兽
		 卡片密码: 12206212
		 使用限制: 无限制
		 种族: 鸟兽
		 属性: 风
		 星级: 6
		 攻击力: 1950
		 防御力: 2100
		 罕见度: 银字R，平卡N
		 卡包: RB，BE01，VOL04，DL02，SD08
		 效果: 这张卡通常召唤不能。需要「万华镜-华丽的分身-」的效果特殊召唤。
		}
		*/
		Id:       235,
		Password: "12206212",
		Name:     "鹰身女郎三姐妹",            // "Harpie Lady Sisters"  "ハーピィ·レディ三姉妹"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   6,
		La:      ygo.LA_Wind,      // 风
		Lr:      ygo.LR_WindBeast, // 鸟兽
		Attack:  1950,
		Defense: 2100,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*46*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 237
		 调整: [进化之茧]
		<進化の繭>
		[14/06/04]

		◎此卡的裁定已经全面变更，旧裁定完全失效
		●这张卡可以从手卡当作装备卡使用给场上表侧表示存在的「幼虫宝宝」装备。
		◇起动效果，开连锁，取自己场上表侧表示的1只「幼虫宝宝/プチモス」为对象
		◇不能以对方场上的「幼虫宝宝/プチモス」为对象
		◇适用后此卡作为装备魔法卡存在于自己场上，持续取装备怪兽为对象
		●以这个效果装备了此卡的「幼虫宝宝/プチモス」的攻击力与守备力适用「进化之茧/進化の繭」的数值。
		◇在此卡作为装备魔法卡存在中，这个效果是装备魔法卡的效果，不开连锁
		 中文名: 进化之茧
		 日文名: 進化の繭
		 英文名: Cocoon of Evolution
		 卡片种类: 效果怪兽
		 卡片密码: 40240595
		 使用限制: 无限制
		 种族: 昆虫
		 属性: 地
		 星级: 3
		 攻击力: 0
		 防御力: 2000
		 罕见度: 平卡N
		 卡包: RB，BE01，VOL04，DL02
		 效果: 这张卡可以从手卡当作装备卡使用给场上表侧表示的「幼虫宝宝」装备。以这个效果装备了此卡的「幼虫宝宝」的攻击力与守备力适用「进化之茧」的数值。
		}
		*/
		Id:       237,
		Password: "40240595",
		Name:     "进化之茧",               // "Cocoon of Evolution"  "進化の繭"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   3,
		La:      ygo.LA_Earth,  // 地
		Lr:      ygo.LR_Insect, // 昆虫
		Attack:  0,
		Defense: 2000,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*47*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 240
		 调整: [暗之假面]
		<闇の仮面>
		[2010/09/08]

		●反转：从自己的墓地选择1张陷阱卡。被选择的卡加入自己的手卡。
		◇反转效果（进入连锁）
		◇必须发动
		◇发动时选择自己墓地中的1张陷阱卡（取对象）
		◇这张卡被攻击在信息确认时变为表侧的场合，效果延迟到怪兽被破坏确定时之后发动（具体请参考伤害步骤的详细说明）
		◇自己的墓地不存在能被选择的卡时，这个效果不发动
		◇效果处理时对象卡不在墓地中存在的场合，这个效果不适用
		 中文名: 暗之假面
		 日文名: 闇の仮面
		 英文名: Mask of Darkness
		 卡片种类: 效果怪兽
		 卡片密码: 28933734
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 2
		 攻击力: 900
		 防御力: 400
		 罕见度: 银字R，平卡N，黄金GR
		 卡包: RB，BE01，YSD02，VOL04，DL02，SD12，Booster R3，GS03
		 效果: 反转：选择自己墓地存在的1张陷阱卡加入手卡。
		}
		*/
		Id:       240,
		Password: "28933734",
		Name:     "暗之假面",               // "Mask of Darkness"  "闇の仮面"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   2,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Attack:  900,
		Defense: 400,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterFlip(func() {
				pl := ca.GetSummoner()
				if c := pl.SelectForWarn(pl.Grave, func(c0 *ygo.Card) bool {
					return c0.IsTrap()
				}); c != nil {
					c.ToHand()
				}
			})
			return true
		}, // 初始

		IsValid: true,
	})

	/*48*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 241
		 调整: [白衣怪盗]
		<白い泥棒>
		[2010/09/08]

		●这张卡给予对方玩家战斗伤害时，对方随机丢弃1张手卡。
		◇诱发效果（进入连锁）
		◇必须发动
		◇详细时点请参考伤害步骤详解
		 中文名: 白衣怪盗
		 日文名: 白い泥棒
		 英文名: White Magical Hat
		 卡片种类: 效果怪兽
		 卡片密码: 15150365
		 使用限制: 无限制
		 种族: 魔法师
		 属性: 光
		 星级: 3
		 攻击力: 1000
		 防御力: 700
		 罕见度: 平卡N，银字R
		 卡包: RB，BE01，VOL05，DL02，Booster R3
		 效果: 这张卡造成对方玩家基本分伤害的时候，对方随机丢弃1张卡。
		}
		*/
		Id:       241,
		Password: "15150365",
		Name:     "白衣怪盗",               // "White Magical Hat"  "白い泥棒"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   3,
		La:      ygo.LA_Light,       // 光
		Lr:      ygo.LR_SpellCaster, // 魔法师
		Attack:  1000,
		Defense: 700,
		Initialize: func(ca *ygo.Card) bool {
			ca.AddEvent(ygo.Deduct, func(tar *ygo.Player) {
				tar.Hand.Get(ygo.RandInt(tar.Hand.Len())).Dispatch(ygo.Discard, ca)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*49*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 242
		 调整: [眼球大王]
		<大王目玉>
		[2010/09/08]

		●反转：从自己的卡组上方看5张卡，可以把那些卡以喜欢的顺序交换位置。
		◇反转效果（进入连锁）
		◇必须发动
		◇这张卡被攻击在信息确认时变为表侧的场合，效果延迟到怪兽被破坏确定时之后发动（具体请参考伤害步骤的详细说明）
		◇效果处理时决定是否替换位置
		 中文名: 眼球大王
		 日文名: 大王目玉
		 英文名: Big Eye
		 卡片种类: 效果怪兽
		 卡片密码: 16768387
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 1200
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: RB，BE01，VOL05，DL02
		 效果: 反转：查看自己的卡组最上面的5张卡，把这5张卡按自己的意愿顺序放回卡组最上面。
		}
		*/
		Id:       242,
		Password: "16768387",
		Name:     "眼球大王",               // "Big Eye"  "大王目玉"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   4,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Attack:  1200,
		Defense: 1000,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterFlip(func() {
				pl := ca.GetSummoner()
				cs := ygo.NewCards()
				for i := 0; i != 5; i++ {
					if c := pl.Deck.EndPop(); c != nil {
						cs.EndPush(c)
					}
				}
				for cs.Len() != 0 {

					if c := pl.SelectForPopup(cs); c != nil {
						pl.Deck.EndPush(c)
						cs.PickedFor(c)
					}
				}

			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*50*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 243
		 调整: （这张卡在规则上也当作「恶魔/デーモン」卡使用）
		 中文名: 暗黑魔龙
		 日文名: ブラック·デーモンズ·ドラゴン
		 英文名: B. Skull Dragon
		 卡片种类: 融合怪兽
		 卡片密码: 11901678
		 使用限制: 无限制
		 种族: 龙
		 属性: 暗
		 星级: 9
		 攻击力: 3200
		 防御力: 2500
		 罕见度: 面闪SR，立体UTR，金字UR，爆闪PR
		 卡包: RB，BE01，MA，VOL05，DL02
		 效果: 融合：「恶魔召唤」＋「真红眼黑龙」。
		（这张卡在规则上也当作「恶魔」卡使用）
		}
		*/
		Id:       243,
		Password: "11901678",
		Name:     "暗黑魔龙",               // "B. Skull Dragon"  "ブラック·デーモンズ·ドラゴン"
		Lc:       ygo.LC_FusionMonster, // 融合怪兽

		Level:   9,
		La:      ygo.LA_Dark,                  // 暗
		Lr:      ygo.LR_Dragon | ygo.LR_Fiend, // 龙
		Attack:  3200,
		Defense: 2500,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterFusionMonster("恶魔召唤", "真红眼黑龙")
			return true
		}, // 初始
		IsValid: true,
	})

	/*51*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 244
		 调整: [假面魔道士]
		<仮面魔道士>
		[2010/09/08]

		●这张卡给予对方玩家战斗伤害时，自己抽1张卡。
		◇诱发效果（进入连锁）
		◇必须发动
		◇详细时点请参考伤害步骤详解
		 中文名: 假面魔道士
		 日文名: 仮面魔道士
		 英文名: Masked Sorcerer
		 卡片种类: 效果怪兽
		 卡片密码: 10189126
		 使用限制: 无限制
		 种族: 魔法师
		 属性: 暗
		 星级: 4
		 攻击力: 900
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: RB，BE01，VOL05，DL02
		 效果: 这张卡造成对方玩家基本分伤害的时候，自己抽1张卡。
		}
		*/
		Id:       244,
		Password: "10189126",
		Name:     "假面魔道士",              // "Masked Sorcerer"  "仮面魔道士"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   4,
		La:      ygo.LA_Dark,        // 暗
		Lr:      ygo.LR_SpellCaster, // 魔法师
		Attack:  900,
		Defense: 1400,
		Initialize: func(ca *ygo.Card) bool {
			ca.AddEvent(ygo.Deduct, func(tar *ygo.Player) {
				pl := ca.GetSummoner()
				pl.ActionDraw(1)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*52*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 245
		 中文名: 轰鸣的大海蛇
		 日文名: 轟きの大海蛇
		 英文名: Roaring Ocean Snake
		 卡片种类: 融合怪兽
		 卡片密码: 19066538
		 使用限制: 无限制
		 种族: 水
		 属性: 水
		 星级: 6
		 攻击力: 2100
		 防御力: 1800
		 罕见度: 平卡N
		 卡包: RB，VOL05，DL02
		 效果: 融合：「魔法灯」＋「兵主部」。
		}
		*/
		Id:       245,
		Password: "19066538",
		Name:     "轰鸣的大海蛇",             // "Roaring Ocean Snake"  "轟きの大海蛇"
		Lc:       ygo.LC_FusionMonster, // 融合怪兽

		Level:   6,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_None,  // 水
		Attack:  2100,
		Defense: 1800,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterFusionMonster("魔法灯", "兵主部")
			return true
		}, // 初始
		IsValid: true,
	})

	/*53*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 249
		 调整: [万华镜-华丽的分身-]
		<万華鏡－華麗なる分身－>
		[2010/09/08]

		●「鹰身女郎」表侧表示在场上有1只以上存在时可以发动。从手卡或卡组把1只「鹰身女郎」或「鹰身女郎三姐妹」特殊召唤。
		◇效果处理时「鹰身女郎/ハーピィ·レディ」不在场上表侧表示存在的场合，这个效果不适用。
		 中文名: 万华镜-华丽的分身-
		 日文名: 万華鏡－華麗なる分身－
		 英文名: Elegant Egotist
		 卡片种类: 通常魔法
		 卡片密码: 90219263
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: RB，BE01，VOL04，DL02，SD08
		 效果: 1只以上的「鹰身女郎」在场上表侧表示存在的时候发动。从手卡·卡组特殊召唤1只「鹰身女郎」或「鹰身女郎三姐妹」。
		}
		*/
		Id:       249,
		Password: "90219263",
		Name:     "万华镜-华丽的分身-",         // "Elegant Egotist"  "万華鏡－華麗なる分身－"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*54*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 250
		 调整: [雷魔神-桑迦]
		<雷魔神－サンガ>
		[2010/09/08]

		●对方回合的战斗伤害计算时才能发动。攻击这张卡的怪兽的攻击力变成0。只要这张卡表侧表示在场上存在这个效果只能使用1次。
		◇诱发即时效果（进入连锁）
		◇任意发动
		◇取对象
		◇详细时点请参考伤害步骤详解
		 中文名: 雷魔神-桑迦
		 日文名: 雷魔神－サンガ
		 英文名: Sanga of the Thunder
		 卡片种类: 效果怪兽
		 卡片密码: 25955164
		 使用限制: 无限制
		 种族: 雷
		 属性: 光
		 星级: 7
		 攻击力: 2600
		 防御力: 2200
		 罕见度: 平卡N，银字R
		 卡包: RB，BE01，VOL05，DL02
		 效果: 对方回合的战斗伤害计算时才能发动。攻击这张卡的怪兽的攻击力变为0。只要这张卡在表侧表示存在这个效果只能使用1次。
		}
		*/
		Id:       250,
		Password: "25955164",
		Name:     "雷魔神-桑迦",             // "Sanga of the Thunder"  "雷魔神－サンガ"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   7,
		La:      ygo.LA_Light,   // 光
		Lr:      ygo.LR_Thunder, // 雷
		Attack:  2600,
		Defense: 2200,
		Initialize: func(ca *ygo.Card) bool {
			ca.OnlyOnce(ygo.BearAttack, func(c *ygo.Card) {
				c.SetAttack(0)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*55*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 251
		 调整: [风魔神-修迦]
		<風魔神－ヒューガ>
		[2010/09/08]

		●对方回合的战斗伤害计算时才能发动。攻击这张卡的怪兽的攻击力变成0。只要这张卡表侧表示在场上存在这个效果只能使用1次。
		◇诱发即时效果（进入连锁）
		◇任意发动
		◇取对象
		◇详细时点请参考伤害步骤详解
		 中文名: 风魔神-修迦
		 日文名: 風魔神－ヒューガ
		 英文名: Kazejin
		 卡片种类: 效果怪兽
		 卡片密码: 62340868
		 使用限制: 无限制
		 种族: 魔法师
		 属性: 风
		 星级: 7
		 攻击力: 2400
		 防御力: 2200
		 罕见度: 平卡N，银字R
		 卡包: RB，BE01，VOL05，DL02
		 效果: 对方回合的战斗伤害计算时才能发动。攻击这张卡的怪兽的攻击力变为0。只要这张卡在表侧表示存在这个效果只能使用1次。
		}
		*/
		Id:       251,
		Password: "62340868",
		Name:     "风魔神-修迦",             // "Kazejin"  "風魔神－ヒューガ"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   7,
		La:      ygo.LA_Wind,        // 风
		Lr:      ygo.LR_SpellCaster, // 魔法师
		Attack:  2400,
		Defense: 2200,
		Initialize: func(ca *ygo.Card) bool {
			ca.OnlyOnce(ygo.BearAttack, func(c *ygo.Card) {
				c.SetAttack(0)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*56*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 252
		 调整: [水魔神-斯迦]
		<水魔神－スーガ>
		[2010/09/08]

		●对方回合的战斗伤害计算时才能发动。攻击这张卡的怪兽的攻击力变成0。只要这张卡表侧表示在场上存在这个效果只能使用1次。
		◇诱发即时效果（进入连锁）
		◇任意发动
		◇取对象
		◇详细时点请参考伤害步骤详解
		 中文名: 水魔神-斯迦
		 日文名: 水魔神－スーガ
		 英文名: Suijin
		 卡片种类: 效果怪兽
		 卡片密码: 98434877
		 使用限制: 无限制
		 种族: 水
		 属性: 水
		 星级: 7
		 攻击力: 2500
		 防御力: 2400
		 罕见度: 平卡N，银字R
		 卡包: RB，BE01，VOL05，DL02
		 效果: 对方回合的战斗伤害计算时才能发动。攻击这张卡的怪兽的攻击力变为0。只要这张卡在表侧表示存在这个效果只能使用1次。
		}
		*/
		Id:       252,
		Password: "98434877",
		Name:     "水魔神-斯迦",             // "Suijin"  "水魔神－スーガ"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   7,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_None,  // 水
		Attack:  2500,
		Defense: 2400,
		Initialize: func(ca *ygo.Card) bool {
			ca.OnlyOnce(ygo.BearAttack, func(c *ygo.Card) {
				c.SetAttack(0)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*57*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 253
		 调整: [魔法灯]
		<魔法のランプ>
		[2010/09/08]

		●这张卡可以直接攻击对方玩家。
		◇永续效果（不进入连锁）
		◇对方场上表侧表示存在有2只「冲锋陷阵的队长/切り込み隊長」的场合，这张卡可以直接攻击对方玩家
		 中文名: 魔法灯
		 日文名: 魔法のランプ
		 英文名: Mystic Lamp
		 卡片种类: 效果怪兽
		 卡片密码: 98049915
		 使用限制: 无限制
		 种族: 魔法师
		 属性: 暗
		 星级: 1
		 攻击力: 400
		 防御力: 300
		 罕见度: 平卡N
		 卡包: RB，BE01，VOL05，DL02
		 效果: 这张卡可以直接攻击对方玩家。
		}
		*/
		Id:       253,
		Password: "98049915",
		Name:     "魔法灯",                // "Mystic Lamp"  "魔法のランプ"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   1,
		La:      ygo.LA_Dark,        // 暗
		Lr:      ygo.LR_SpellCaster, // 魔法师
		Attack:  400,
		Defense: 300,
		Initialize: func(ca *ygo.Card) bool {
			ca.SetCanDirect()
			return true
		}, // 初始
		IsValid: true,
	})

	/*58*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 254
		 调整: [铁蝎]
		<鉄のサソリ>
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
		 日文名: 鉄のサソリ
		 英文名: Steel Scorpion
		 卡片种类: 效果怪兽
		 卡片密码: 13599884
		 使用限制: 无限制
		 种族: 机械
		 属性: 地
		 星级: 1
		 攻击力: 250
		 防御力: 300
		 罕见度: 平卡N
		 卡包: RB，VOL04，DL02
		 效果: 机械族以外的怪兽攻击这张卡的场合，那只怪兽（以对方的回合来数）第3个回合的回合结束时破坏。
		}
		*/
		Id:       254,
		Password: "13599884",
		Name:     "铁蝎",                 // "Steel Scorpion"  "鉄のサソリ"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   1,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Machine, // 机械
		Attack:  250,
		Defense: 300,
		Initialize: func(ca *ygo.Card) bool {
			ca.AddEvent(ygo.Fought, func(c *ygo.Card) {
				if c.RaceIsMachine() {
					return
				}
				i := 0
				c.RegisterGlobalListen(ygo.RoundEnd, func(pl0 *ygo.Player) {
					if c.GetSummoner() != pl0 {
						return
					}
					i++
					if i >= 3 {
						c.Dispatch(ygo.Destroy, ca)
					}
				})
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*59*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 255
		 调整: [百足虫]
		<レッグル>
		[2010/09/08]

		●这张卡可以直接攻击对方玩家。
		◇永续效果（不进入连锁）
		◇对方场上表侧表示存在有2只「冲锋陷阵的队长/切り込み隊長」的场合，这张卡可以直接攻击对方玩家
		 中文名: 百足虫
		 日文名: レッグル
		 英文名: Leghul
		 卡片种类: 效果怪兽
		 卡片密码: 12472242
		 使用限制: 无限制
		 种族: 昆虫
		 属性: 地
		 星级: 1
		 攻击力: 300
		 防御力: 350
		 罕见度: 平卡N
		 卡包: RB，VOL05，DL02
		 效果: 这张卡可以直接攻击对方玩家。
		}
		*/
		Id:       255,
		Password: "12472242",
		Name:     "百足虫",                // "Leghul"  "レッグル"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   1,
		La:      ygo.LA_Earth,  // 地
		Lr:      ygo.LR_Insect, // 昆虫
		Attack:  300,
		Defense: 350,
		Initialize: func(ca *ygo.Card) bool {
			ca.SetCanDirect()
			return true
		}, // 初始
		IsValid: true,
	})

	/*60*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 256
		 调整: [巨口]
		<ラージマウス>
		[2010/09/08]

		●这张卡可以直接攻击对方玩家。
		◇永续效果（不进入连锁）
		◇对方场上表侧表示存在有2只「冲锋陷阵的队长/切り込み隊長」的场合，这张卡可以直接攻击对方玩家
		 中文名: 巨口
		 日文名: ラージマウス
		 英文名: Ooguchi
		 卡片种类: 效果怪兽
		 卡片密码: 58861941
		 使用限制: 无限制
		 种族: 水
		 属性: 水
		 星级: 1
		 攻击力: 300
		 防御力: 250
		 罕见度: 平卡N
		 卡包: RB，VOL05，DL02
		 效果: 这张卡可以直接攻击对方玩家。
		}
		*/
		Id:       256,
		Password: "58861941",
		Name:     "巨口",                 // "Ooguchi"  "ラージマウス"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   1,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_None,  // 水
		Attack:  300,
		Defense: 250,
		Initialize: func(ca *ygo.Card) bool {
			ca.SetCanDirect()
			return true
		}, // 初始
		IsValid: true,
	})

	/*61*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 258
		 调整: [炸弹先生]
		<ミスター·ボンバー>
		[2010/09/08]

		●自己准备阶段时才能发动。表侧表示的这张卡作为祭品，选择2只攻击力1000以下的表侧表示怪兽破坏。
		◇诱发效果（进入连锁）
		◇任意发动
		◇符合条件的怪兽不足2只的场合不能发动
		◇发动时选择符合条件的2只怪兽（取对象）
		◇这张卡作为祭品是代价
		◇不能选择里侧表示的怪兽为对象
		★效果处理时，对象之一不在场上表侧表示存在的场合如何处理 调整中
		★效果处理时，对象中有攻击力高于1000的场合如何处理 调整中
		 中文名: 炸弹先生
		 日文名: ミスター·ボンバー
		 英文名: Blast Juggler
		 卡片种类: 效果怪兽
		 卡片密码: 70138455
		 使用限制: 无限制
		 种族: 机械
		 属性: 炎
		 星级: 3
		 攻击力: 800
		 防御力: 900
		 罕见度: 平卡N
		 卡包: RB，BE01，VOL05，DL02
		 效果: 在自己的准备阶段时发动。表侧表示的这张卡作为祭品，选择2只表侧表示的攻击力1000以下的怪兽破坏。
		}
		*/
		Id:       258,
		Password: "70138455",
		Name:     "炸弹先生",               // "Blast Juggler"  "ミスター·ボンバー"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   3,
		La:      ygo.LA_None,    // 炎
		Lr:      ygo.LR_Machine, // 机械
		Attack:  800,
		Defense: 900,
		Initialize: func(ca *ygo.Card) bool {
			e := func() {
				ca.RegisterIgnitionSelector(ygo.SP, func(pl0 *ygo.Player) bool {
					if pl := ca.GetSummoner(); pl == pl0 {
						tar := pl.GetTarget()
						css := ygo.NewCards(tar.Mzone, pl.Mzone, func(c0 *ygo.Card) bool {
							return c0.IsFaceUp() && (c0.GetAttack() <= 1000) && c0 != ca
						})
						if css.Len() > 1 {
							ca.PushChain(func() {
								ca.Dispatch(ygo.Cost)
								for i := 0; i != 2; i++ {
									if c := pl.SelectForWarn(css); c != nil {
										c.Dispatch(ygo.Destroy, ca)
									}
								}

							})
						}
					}
					return true
				})
			}
			ca.AddEvent(ygo.FaceUp, e)

			return true
		}, // 初始
		IsValid: true,
	})

	/*62*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 259
		 调整: [人造人7号]
		<人造人間７号>
		[2010/09/08]

		●这张卡可以直接攻击对方玩家。
		◇永续效果（不进入连锁）
		◇对方场上表侧表示存在有2只「冲锋陷阵的队长/切り込み隊長」的场合，这张卡可以直接攻击对方玩家
		 中文名: 人造人7号
		 日文名: 人造人間７号
		 英文名: Jinzo #7
		 卡片种类: 效果怪兽
		 卡片密码: 32809211
		 使用限制: 无限制
		 种族: 机械
		 属性: 暗
		 星级: 2
		 攻击力: 500
		 防御力: 400
		 罕见度: 平卡N
		 卡包: RB，BE01，VOL05，DL02
		 效果: 这张卡可以直接攻击对方玩家。
		}
		*/
		Id:       259,
		Password: "32809211",
		Name:     "人造人7号",              // "Jinzo #7"  "人造人間７号"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   2,
		La:      ygo.LA_Dark,    // 暗
		Lr:      ygo.LR_Machine, // 机械
		Attack:  500,
		Defense: 400,
		Initialize: func(ca *ygo.Card) bool {
			ca.SetCanDirect()
			return true
		}, // 初始
		IsValid: true,
	})

	/*63*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 260
		 调整: [神圣魔术师]
		<聖なる魔術師>
		[2010/09/08]

		●反转：从自己的墓地选择1张魔法卡。被选择的卡加入自己的手卡。
		◇反转效果（进入连锁）
		◇必须发动
		◇发动时选择自己墓地中的1张魔法卡（取对象）
		◇这张卡被攻击在信息确认时变为表侧的场合，效果延迟到怪兽被破坏确定时之后发动（具体请参考伤害步骤的详细说明）
		◇自己的墓地不存在能被选择的卡时，这个效果不发动
		◇效果处理时对象卡不在墓地中存在的场合，这个效果不适用
		 中文名: 神圣魔术师
		 日文名: 聖なる魔術師
		 英文名: Magician of Faith
		 卡片种类: 效果怪兽
		 卡片密码: 31560081
		 使用限制: 禁止卡
		 种族: 魔法师
		 属性: 光
		 星级: 1
		 攻击力: 300
		 防御力: 400
		 罕见度: 平卡N，银字R
		 卡包: RB，BE01，YSD01，VOL04，DL02，SD06，PE，SY2，SK2
		 效果: 反转：选择自己的墓地的1张魔法卡。选择的卡加入自己手卡。
		}
		*/
		Id:       260,
		Password: "31560081",
		Name:     "神圣魔术师",              // "Magician of Faith"  "聖なる魔術師"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   1,
		La:      ygo.LA_Light,       // 光
		Lr:      ygo.LR_SpellCaster, // 魔法师
		Attack:  300,
		Defense: 400,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterFlip(func() {
				pl := ca.GetSummoner()
				if c := pl.SelectForWarn(pl.Grave, func(c0 *ygo.Card) bool {
					return c0.IsMagic()
				}); c != nil {
					c.ToHand()
				}
			})
			return true
		}, // 初始

		IsValid: true,
	})

	/*64*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 262
		 调整: [彩虹花]
		<レインボー·フラワー>
		[2010/09/08]

		●这张卡可以直接攻击对方玩家。
		◇永续效果（不进入连锁）
		◇对方场上表侧表示存在有2只「冲锋陷阵的队长/切り込み隊長」的场合，这张卡可以直接攻击对方玩家
		 中文名: 彩虹花
		 日文名: レインボー·フラワー
		 英文名: Rainbow Flower
		 卡片种类: 效果怪兽
		 卡片密码: 21347810
		 使用限制: 无限制
		 种族: 植物
		 属性: 地
		 星级: 2
		 攻击力: 400
		 防御力: 500
		 罕见度: 平卡N
		 卡包: RB，VOL05，DL02
		 效果: 这张卡可以直接攻击对方玩家。
		}
		*/
		Id:       262,
		Password: "21347810",
		Name:     "彩虹花",                // "Rainbow Flower"  "レインボー·フラワー"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   2,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Plant, // 植物
		Attack:  400,
		Defense: 500,
		Initialize: func(ca *ygo.Card) bool {
			ca.SetCanDirect()
			return true
		}, // 初始
		IsValid: true,
	})

	/*65*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 264
		 调整: [电气蜥蜴]
		<でんきトカゲ>
		[2010/09/08]

		●不死族以外的怪兽被这张卡攻击的场合，那只怪兽在下个回合不能攻击宣言。
		◇诱发效果（进入连锁）
		◇必须发动
		◇不取对象
		◇在怪兽被战斗破坏送去墓地时发动
		◇那只怪兽在这个效果发动后一度不在场上表侧表示存在的场合，这个效果重置
		◇那只怪兽在这个效果发动后一度变为不死族怪兽的场合，这个效果重置
		 中文名: 电气蜥蜴
		 日文名: でんきトカゲ
		 英文名: Electric Lizard
		 卡片种类: 效果怪兽
		 卡片密码: 55875323
		 使用限制: 无限制
		 种族: 雷
		 属性: 地
		 星级: 3
		 攻击力: 850
		 防御力: 800
		 罕见度: 平卡N
		 卡包: RB，VOL04，DL02
		 效果: 不死族以外的怪兽攻击这张卡的场合，那只怪兽下一个回合不能攻击宣言。
		}
		*/
		Id:       264,
		Password: "55875323",
		Name:     "电气蜥蜴",               // "Electric Lizard"  "でんきトカゲ"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   3,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Thunder, // 雷
		Attack:  850,
		Defense: 800,
		Initialize: func(ca *ygo.Card) bool {
			ca.AddEvent(ygo.Fought, func(c *ygo.Card) {
				if c.RaceIsMachine() {
					return
				}
				c.SetSizeRoundNotCanAttack(1)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*66*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 266
		 调整: [女王的影武者]
		<女王の影武者>
		[2010/09/08]

		●这张卡可以直接攻击对方玩家。
		◇永续效果（不进入连锁）
		◇对方场上表侧表示存在有2只「冲锋陷阵的队长/切り込み隊長」的场合，这张卡可以直接攻击对方玩家
		 中文名: 女王的影武者
		 日文名: 女王の影武者
		 英文名: Queen's Double
		 卡片种类: 效果怪兽
		 卡片密码: 05901497
		 使用限制: 无限制
		 种族: 战士
		 属性: 地
		 星级: 1
		 攻击力: 350
		 防御力: 300
		 罕见度: 平卡N
		 卡包: RB，VOL05，DL02
		 效果: 这张卡可以直接攻击对方玩家。
		}
		*/
		Id:       266,
		Password: "05901497",
		Name:     "女王的影武者",             // "Queen's Double"  "女王の影武者"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   1,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Attack:  350,
		Defense: 300,
		Initialize: func(ca *ygo.Card) bool {
			ca.SetCanDirect()
			return true
		}, // 初始
		IsValid: true,
	})

	/*67*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 268
		 调整: [伪陷阱]
		<伪物のわな>
		[10/09/13]

		●对方使用破坏自己场上的陷阱卡的魔法·陷阱·效果怪兽的效果时可以发动。这张卡代替破坏，其他的陷阱卡不破坏。覆盖的卡要被破坏的场合，那些卡全部翻开确认。
		◇必须在自己场上有其他陷阱卡存在时才能发动
		◇发动时给对方确认覆盖的卡
		◇连锁这张卡的发动，发动「沙尘之大龙卷/砂尘の大竜巻」在场上覆盖陷阱卡的场合，之后不能确认那张覆盖的陷阱卡
		◇不能对应「元素英雄 天空侠/E·HERO エアーマン」等不确定的破坏效果发动这张卡
		◇「王宫的通告/王宫のお触れ」在场上存在时，发动这张卡的场合也确认覆盖卡
		◇这张卡在自己场上有其他陷阱卡存在时可以对应「暗之卡组破坏病毒/闇のデッキ破壊ウイルス」发动，那个场合自己场上的陷阱卡不破坏，手卡中的陷阱卡破坏
		◇这张卡对应「黑蔷薇龙/ブラック·ローズ·ドラゴン」的破坏效果发动的场合，自己场上的其他陷阱卡不破坏
		◇「旋风/サイクロン」以1张覆盖的陷阱卡为对象发动的场合这张卡对应发动，在发动时只确认那张对象卡
		◇对应这张卡的发动而发动「旋风/サイクロン」并以这张卡为对象的场合，效果处理时不能代替其他陷阱卡破坏
		◇可以对应「命运英雄 钻石人/D－HERO ダイヤモンドガイ」墓地发动的「大风暴/大岚」等发动
		◇这张卡可以对应「星尘龙/スターダスト·ドラゴン」的效果发动而发动
		◇「星尘龙/スターダスト·ドラゴン」不能对应这张卡的发动而发动
		★是否能对应魔法卡效果的发动而发动 调整中
		 中文名: 伪陷阱
		 日文名: 偽物のわな
		 英文名: Fake Trap
		 卡片种类: 通常陷阱
		 卡片密码: 03027001
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: RB，BE01，Booster R3，VOL05，DL02，TP02，JY
		 效果: 把自己场上的陷阱卡破坏的魔法·陷阱·效果怪兽的效果对方发动时才能发动，把这张卡作为代替破坏，其他的自己的陷阱卡不破坏。盖放的卡破坏的场合，那些卡全部翻开确认。
		}
		*/
		Id:       268,
		Password: "03027001",
		Name:     "伪陷阱",               // "Fake Trap"  "偽物のわな"
		Lc:       ygo.LC_OrdinaryTrap, // 通常陷阱

		//		Initialize: func(ca *ygo.Card) bool {
		//			ca.RegisterOrdinaryTrap(ygo.Destroy, func(c *ygo.Card) {
		//				pl := ca.GetSummoner()
		//				if c.GetSummoner() == pl && c != ca && c.IsTrap() {
		//					ca.PushChain(func() {

		//					})
		//				}
		//			})
		//			return true
		//		}, // 初始
		IsValid: false,
	})

	/*68*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 269
		 调整: [对死者的供奉]
		<死者への手向け>
		[2010/09/08]

		●丢弃1张手卡，选择场上存在的1只怪兽发动。被选择的怪兽破坏。
		◇发动时丢弃1张手卡（代价）
		◇发动时选择自己或对方场上存在的1只怪兽（取对象）
		◇效果处理时对象怪兽不在场上存在的场合，这个效果不适用
		 中文名: 对死者的供奉
		 日文名: 死者への手向け
		 英文名: Tribute to The Doomed
		 卡片种类: 通常魔法
		 卡片密码: 79759861
		 使用限制: 无限制
		 罕见度: 平卡N，银字R
		 卡包: RB，BE01，YSD01，VOL05，DL02，SD03，YSD04，KA，SY2
		 效果: 丢弃1张手卡，选择场上存在的1只怪兽发动。选择的怪兽破坏。
		}
		*/
		Id:       269,
		Password: "79759861",
		Name:     "对死者的供奉",             // "Tribute to The Doomed"  "死者への手向け"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterPay(func(s string) {
				if s != ygo.Onset {
					return
				}
				pl := ca.GetSummoner()
				if c := pl.SelectForWarn(pl.Hand, func(c0 *ygo.Card) bool {
					return c0 != ca
				}); c != nil {
					c.Dispatch(ygo.Cost, ca)
				} else {
					ca.StopOnce(s)
				}
			})
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				if c := pl.SelectForWarn(tar.Mzone, pl.Mzone); c != nil {
					c.Dispatch(ygo.Destroy, ca)
				}
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*69*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 270
		 调整: [魂之解放]
		<魂の解放>
		[2012/01/29]

		●选择双方墓地的卡合计最多5张，那些卡从游戏中除外。
		◇取对象效果。
		◇效果发动时选择双方墓地合计1~5张卡作为效果的对象。
		◇双方墓地任意一方有1张以上的卡存在的场合就可以将这张卡发动。
		 中文名: 魂之解放
		 日文名: 魂の解放
		 英文名: Soul Release
		 卡片种类: 通常魔法
		 卡片密码: 05758500
		 使用限制: 无限制
		 罕见度: 平卡N，银字R，黄金GR
		 卡包: RB，BE01，VOL05，DL02，SD14，GS04
		 效果: 选择双方墓地的卡合计最多5张，那些卡从游戏中除外。
		}
		*/
		Id:       270,
		Password: "05758500",
		Name:     "魂之解放",               // "Soul Release"  "魂の解放"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				for i := 0; i != 5; i++ {
					if c := pl.SelectForWarn(tar.Grave, pl.Grave); c != nil {
						c.ToRemoved()
					}
				}
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*70*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 271
		 调整: [开朗的殡葬师]
		<陽気な葬儀屋>
		[2010/09/08]

		●从自己的手卡把最多3张怪兽卡丢弃去墓地。
		◇效果处理时选择要丢弃的卡（不取对象）
		◇发动的场合，必须至少选择1张卡
		 中文名: 开朗的殡葬师
		 日文名: 陽気な葬儀屋
		 英文名: The Cheerful Coffin
		 卡片种类: 通常魔法
		 卡片密码: 41142615
		 使用限制: 无限制
		 罕见度: 平卡N，银字R
		 卡包: RB，BE01，VOL05，DL02
		 效果: 从自己的手卡中丢弃最多3张怪兽卡送去墓地。
		}
		*/
		Id:       271,
		Password: "41142615",
		Name:     "开朗的殡葬师",             // "The Cheerful Coffin"  "陽気な葬儀屋"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				for i := 0; i != 3; i++ {
					if c := pl.SelectForWarn(pl.Hand, func(c0 *ygo.Card) bool {
						return c0.IsMonster()
					}); c != nil {
						c.Dispatch(ygo.Discard, ca)
					}
				}
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*71*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 272
		 调整: [心变]
		<心変わり>
		[2010/09/08]

		●选择对方场上的1只怪兽。直到发动回合的结束阶段为止，得到被选择的卡的控制权。
		◇发动时选择对方场上的1只怪兽（取对象）
		◇效果处理时对象怪兽不在场上存在的场合，这个效果不适用
		◇对方在我方回合发动「血之代偿/血の代償」「活死人的呼声/リビングデッドの呼び声」等召唤、特殊召唤怪兽的场合，得到那些怪兽控制权后可以改变那些怪兽的表示形式
		◇这个效果得到控制权的怪兽，变为里侧表示的场合，结束阶段时也归还控制权
		◇「天使的手镜/天使の手鏡」对应这张卡发动的场合，也只能选择「心变/心変わり」控制者的对方的场上怪兽
		◇用这个效果得到控制权的对方怪兽，再用其他卡的效果把那只怪兽的控制权转移给对方，那个回合的结束阶段时那只怪兽不在需要归还控制权
		◇用这个效果得到控制权的对方怪兽，被「亚空间物质传送装置/亜空間物質転送装置」暂时从游戏中除外，结束阶段时也要归还控制权
		◇用这个效果得到控制权的对方怪兽，被「虫洞/ワーム·ホール」暂时从游戏中除外，在下次的准备阶段时那只怪兽回到场上后立刻回到对方场上
		★对应这张卡的发动，对方发动「扰乱三人组/おジャマトリオ」让自己场上没有空余的怪兽区域时如何处理 调整中
		◇结束阶段归还控制权时，若对方场上不存在空余的怪兽区域，那只怪兽送到持有者的墓地
		 中文名: 心变
		 日文名: 心変わり
		 英文名: Change of Heart
		 卡片种类: 通常魔法
		 卡片密码: 04031928
		 使用限制: 禁止卡
		 罕见度: 金字UR，面闪SR
		 卡包: RB，BE01，EX-R(EX)，VOL05，DL02，YU，PE
		 效果: 选择对方的场上的1只怪兽。直到发动回合的结束阶段得到选择的卡的控制权。
		}
		*/
		Id:       272,
		Password: "04031928",
		Name:     "心变",                 // "Change of Heart"  "心変わり"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				if c := pl.SelectForWarn(tar.Mzone); c != nil {
					t := c.GetSummoner()
					if c.SetSummoner(ca.GetSummoner()); c.IsInMzone() {
						c.ToMzone()
						pl.OnlyOnce(ygo.RoundEnd, func() {
							if c.SetSummoner(t); c.IsInMzone() {
								c.ToMzone()
							}
						}, ca, c)
					}
				}
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*72*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 410
		 调整: [野蛮人2号]
		<バーバリアン２号>
		[2010/09/08]

		●自己的场上每有1只表侧表示存在的「野蛮人１号」，这张卡的攻击力上升500分。
		◇永续效果（不进入连锁）
		 中文名: 野蛮人2号
		 日文名: バーバリアン２号
		 英文名: Swamp Battleguard
		 卡片种类: 效果怪兽
		 卡片密码: 40453765
		 使用限制: 无限制
		 种族: 战士
		 属性: 地
		 星级: 5
		 攻击力: 1800
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: ME，VOL07，DL04，AT06
		 效果: 每有1只自己场上表侧表示存在的「野蛮人1号」，这张卡的攻击力上升500。
		}
		*/
		Id:       410,
		Password: "40453765",
		Name:     "野蛮人2号",              // "Swamp Battleguard"  "バーバリアン２号"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   5,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Attack:  1800,
		Defense: 1500,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterMzoneAccessArea(func(c *ygo.Card) bool {
				return c.GetName() == "野蛮人1号"
			}, func(c *ygo.Card) {
				ca.SetAttack(ca.GetAttack() + 500)
			}, func(c *ygo.Card) {
				ca.SetAttack(ca.GetAttack() - 500)
			})
			return true

		}, // 初始
		IsValid: true,
	})

	/*73*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 414
		 调整: [壶魔人]
		<壺魔人>
		[2010/09/08]

		●反转：场上表侧表示存在的「龙族封印之壶」破坏。破坏的场合，场上表侧表示存在的龙族怪兽全部变为攻击表示。
		◇诱发效果，强制发动，开连锁，不取对象
		◇效果处理时场上表侧表示存在的「龙族封印之壶/ドラゴン族·封印の壺」全部破坏
		 中文名: 壶魔人
		 日文名: 壺魔人
		 英文名: Dragon Piper
		 卡片种类: 效果怪兽
		 卡片密码: 55763552
		 使用限制: 无限制
		 种族: 炎
		 属性: 炎
		 星级: 3
		 攻击力: 200
		 防御力: 1800
		 罕见度: 平卡N
		 卡包: ME，BE02，VOL06，DL04，PE
		 效果: 反转：场上表侧表示的「龙族封印之壶」破坏。破坏时，场上表侧表示存在的龙族怪兽全部攻击表示。
		}
		*/
		Id:       414,
		Password: "55763552",
		Name:     "壶魔人",                // "Dragon Piper"  "壺魔人"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   3,
		La:      ygo.LA_Fire, // 炎
		Lr:      ygo.LR_Pyro, // 炎
		Attack:  200,
		Defense: 1800,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterFlip(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				css := ygo.NewCards(pl.Mzone, tar.Mzone, func(c *ygo.Card) bool {
					return c.GetName() == "龙族封印之壶"
				})
				if css.Len() == 0 {
					return
				}
				css.ForEach(func(c *ygo.Card) bool {
					c.Dispatch(ygo.Destroy)
					return true
				})
				css2 := ygo.NewCards(pl.Mzone, tar.Mzone, func(c *ygo.Card) bool {
					return c.RaceIsDragon() && c.IsFaceUpDefense()
				})
				css2.ForEach(func(c *ygo.Card) bool {
					c.SetFaceUpAttack()
					return true
				})
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*74*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 416
		 调整: [三眼怪]
		<クリッター>
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
		 日文名: クリッター
		 英文名: Sangan
		 卡片种类: 效果怪兽
		 卡片密码: 26202165
		 使用限制: 禁止卡
		 种族: 恶魔
		 属性: 暗
		 星级: 3
		 攻击力: 1000
		 防御力: 600
		 罕见度: 银字R，黄金GR，平卡N，金字UR
		 卡包: ME，BE02，VOL06，DL04，GS01，GLD02，JY，KA，PE，SY2，SD21，TU06，DB12，ST12，DS13
		 效果: 这张卡从场上送去墓地时，从自己卡组把1只攻击力1500以下的怪兽加入手卡。
		}
		*/
		Id:       416,
		Password: "26202165",
		Name:     "三眼怪",                // "Sangan"  "クリッター"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   3,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Attack:  1000,
		Defense: 600,
		Initialize: func(ca *ygo.Card) bool {
			ca.AddEvent(ygo.InMzone, func() {
				ca.AddEvent(ygo.InGrave, func() {
					pl := ca.GetSummoner()
					cs := pl.Deck.Find(func(c *ygo.Card) bool {
						return c.IsMonster() && c.GetDefense() <= 1500
					})
					if c := pl.SelectForPopup(cs); c != nil {
						c.ToHand()
					}
				})
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*75*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 417
		 调整: [大飞蛾]
		<グレート·モス>
		[2010/09/08]

		●把装备了「进化之茧」4回合（计算自己回合）后的「幼虫宝宝」作为祭品可以特殊召唤。
		◇召唤规则
		◇不能通常召唤
		◇经过正规出场后，可以从墓地或除外状态用其他卡的效果来特殊召唤
		 中文名: 大飞蛾
		 日文名: グレート·モス
		 英文名: Great Moth
		 卡片种类: 效果怪兽
		 卡片密码: 14141448
		 使用限制: 无限制
		 种族: 昆虫
		 属性: 地
		 星级: 8
		 攻击力: 2600
		 防御力: 2500
		 罕见度: 平卡N
		 卡包: ME，BE02，VOL06，DL04
		 效果: 装备了「进化之茧」的「幼虫宝宝」4回合后（用自己的回合来数）作祭品来特殊召唤。
		}
		*/
		Id:       417,
		Password: "14141448",
		Name:     "大飞蛾",                // "Great Moth"  "グレート·モス"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   8,
		La:      ygo.LA_Earth,  // 地
		Lr:      ygo.LR_Insect, // 昆虫
		Attack:  2600,
		Defense: 2500,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*76*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 418
		 调整: [栗子球]
		<クリボー>
		[2010/09/08]

		●对方回合的战斗伤害计算时，这张卡从手卡丢弃发动。那个战斗发生对控制者的战斗伤害变成0。
		◇诱发即时效果（进入连锁）
		◇任意发动
		◇发动时把这张卡从手卡丢弃（代价）
		◇「死灵骑士/死霊騎士デスカリバー·ナイト」的攻击造成战斗伤害的场合发动这张卡的效果，「死灵骑士/死霊騎士デスカリバー·ナイト」解放把这张卡的效果的发动无效，由于是在受到伤害之前，这张卡的控制者不会受到那次战斗所计算的伤害。
		◇对方怪兽的攻击力比自己怪兽低的场合，这个效果也能发动
		◇攻击力为0的怪兽进行战斗的场合，不能发动这个效果
		 中文名: 栗子球
		 日文名: クリボー
		 英文名: Kuriboh
		 卡片种类: 效果怪兽
		 卡片密码: 40640057
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 1
		 攻击力: 300
		 防御力: 200
		 罕见度: 银字R，平卡N，金字UR
		 卡包: ME，BE02，YSD01，VOL07，DL04，DT01，DPYG，DTP01，YU，SY2，15AY
		 效果: 从手卡丢弃这张卡，自己受到的战斗伤害1次为0。这个效果只能在对方的战斗回合使用。
		}
		*/
		Id:       418,
		Password: "40640057",
		Name:     "栗子球",                // "Kuriboh"  "クリボー"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   1,
		La:      ygo.LA_Dark,  // 暗
		Lr:      ygo.LR_Fiend, // 恶魔
		Attack:  300,
		Defense: 200,
		Initialize: func(ca *ygo.Card) bool {
			ca.AddEvent(ygo.InHand, func() {
				ca.RegisterIgnitionSelector(ygo.Deduct, func(c *ygo.Card, tar *ygo.Player) {
					pl := ca.GetSummoner()
					if pl == tar {
						ca.PushChain(func() {
							ca.Dispatch(ygo.Cost)
							c.StopOnce(ygo.Deduct)
						})
					}
				})
			})
			ca.AddEvent(ygo.OutHand, func() {
				ca.UnregisterGlobalListen()
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*77*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 419
		 中文名: 千年龙
		 日文名: 千年竜
		 英文名: Thousand Dragon
		 卡片种类: 融合怪兽
		 卡片密码: 41462083
		 使用限制: 无限制
		 种族: 龙
		 属性: 风
		 星级: 7
		 攻击力: 2400
		 防御力: 2000
		 罕见度: 平卡N，金字UR
		 卡包: ME，BE02，LE02，VOL06，DL04，JY
		 效果: 融合：「时间魔术师」＋「宝贝龙」。
		}
		*/
		Id:       419,
		Password: "41462083",
		Name:     "千年龙",                // "Thousand Dragon"  "千年竜"
		Lc:       ygo.LC_FusionMonster, // 融合怪兽

		Level:   7,
		La:      ygo.LA_Wind,   // 风
		Lr:      ygo.LR_Dragon, // 龙
		Attack:  2400,
		Defense: 2000,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterFusionMonster("时间魔术师", "宝贝龙")
			return true
		}, // 初始
		IsValid: true,
	})

	/*78*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 423
		 调整: [弹射龟]
		<カタパルト·タートル>
		[14/08/16]

		●1回合1次，把自己场上1只怪兽解放才能发动。给与对方基本分解放的怪兽的攻击力一半数值的伤害。
		◇起动效果，开连锁，解放怪兽是效果发动COST
		◇计算攻击力的参照值是那只祭品怪兽在场上的攻击力
		◇可以把此卡自身解放发动这个效果
		 中文名: 弹射龟
		 日文名: カタパルト·タートル
		 英文名: Catapult Turtle
		 卡片种类: 效果怪兽
		 卡片密码: 95727991
		 使用限制: 无限制
		 种族: 水
		 属性: 水
		 星级: 5
		 攻击力: 1000
		 防御力: 2000
		 罕见度: 平卡N，银字R，面闪SR
		 卡包: ME，BE02，VOL07，DL04，DPYG，15AY
		 效果: 1回合1次，把自己场上1只怪兽解放才能发动。给与对方基本分解放的怪兽的攻击力一半数值的伤害。
		}
		*/
		Id:       423,
		Password: "95727991",
		Name:     "弹射龟",                // "Catapult Turtle"  "カタパルト·タートル"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   5,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_None,  // 水
		Attack:  1000,
		Defense: 2000,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*79*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 425
		 中文名: 牛头人马兽
		 日文名: ミノケンタウロス
		 英文名: Rabid Horseman
		 卡片种类: 融合怪兽
		 卡片密码: 94905343
		 使用限制: 无限制
		 种族: 兽战士
		 属性: 地
		 星级: 6
		 攻击力: 2000
		 防御力: 1700
		 罕见度: 平卡N
		 卡包: ME，BE02，VOL07，DL04
		 效果: 融合：「牛头人」＋「人马兽」。
		}
		*/
		Id:       425,
		Password: "94905343",
		Name:     "牛头人马兽",              // "Rabid Horseman"  "ミノケンタウロス"
		Lc:       ygo.LC_FusionMonster, // 融合怪兽

		Level:   6,
		La:      ygo.LA_Earth,       // 地
		Lr:      ygo.LR_BeastWarror, // 兽战士
		Attack:  2000,
		Defense: 1700,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterFusionMonster("牛头人", "人马兽")
			return true
		}, // 初始
		IsValid: true,
	})

	/*80*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 431
		 调整: [剑之女王]
		<剣の女王>
		[2010/09/08]

		●反转：对方场上每有1张魔法·陷阱卡，给予对方基本分500分的伤害。
		◇反转效果（进入连锁）
		◇必须发动
		◇这张卡被攻击在信息确认时变为表侧的场合，效果延迟到怪兽被破坏确定时之后发动（具体请参考伤害步骤的详细说明）
		◇效果处理时计算对方场上的魔法·陷阱卡数量
		◇效果处理时对方场上不存在魔法·陷阱卡的场合，这个效果不适用
		 中文名: 剑之女王
		 日文名: 剣の女王
		 英文名: Princess of Tsurugi
		 卡片种类: 效果怪兽
		 卡片密码: 51371017
		 使用限制: 无限制
		 种族: 战士
		 属性: 风
		 星级: 3
		 攻击力: 900
		 防御力: 700
		 罕见度: 平卡N，银字R
		 卡包: ME，BE02，YSD01，VOL07，DL04，Booster07
		 效果: 反转：对方场上每存在1张魔法·陷阱，对方受到500分的伤害。
		}
		*/
		Id:       431,
		Password: "51371017",
		Name:     "剑之女王",               // "Princess of Tsurugi"  "剣の女王"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   3,
		La:      ygo.LA_Wind,    // 风
		Lr:      ygo.LR_Warrior, // 战士
		Attack:  900,
		Defense: 700,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterFlip(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				i := tar.Szone.Len()
				if i != 0 {
					tar.ChangeHp(i * -500)
				}
			})
			return true
		}, // 初始

		IsValid: true,
	})

	/*81*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 432
		 调整: [火炎地狱]
		<火炎地獄>
		[2010/09/08]

		●给予对方基本分1000分的伤害，自己受到500分的伤害。
		◇同时发生伤害
		◇对应这张卡的发动，对方发动「地狱的冷枪/地獄の扉越し銃」「痛魂之咒术/痛魂の呪術」的场合，自己受到合计1500分的伤害
		◇对应这张卡的发动，对方发动「黑板擦的陷阱/黒板消しの罠」的场合，对方受到的1000分伤害无效
		 中文名: 火炎地狱
		 日文名: 火炎地獄
		 英文名: Tremendous Fire
		 卡片种类: 通常魔法
		 卡片密码: 46918794
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: ME，BE02，VOL07，DL04，Booster07，YSD03
		 效果: 给予对方基本分1000分伤害，自己受到500分伤害。
		}
		*/
		Id:       432,
		Password: "46918794",
		Name:     "火炎地狱",               // "Tremendous Fire"  "火炎地獄"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				tar.ChangeHp(-1000)
				pl.ChangeHp(-500)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*82*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 434
		 调整: [影之食尸鬼]
		<シャドウ·グール>
		[2010/09/08]

		●自己的墓地每存在1只怪兽，这张卡的攻击力上升100分。
		◇永续效果（不进入连锁）
		 中文名: 影之食尸鬼
		 日文名: シャドウ·グール
		 英文名: Shadow Ghoul
		 卡片种类: 效果怪兽
		 卡片密码: 30778711
		 使用限制: 无限制
		 种族: 不死
		 属性: 暗
		 星级: 5
		 攻击力: 1600
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: ME，BE02，VOL06，DL04
		 效果: 自己的墓地每存在1只怪兽，这张卡的攻击力上升100。
		}
		*/
		Id:       434,
		Password: "30778711",
		Name:     "影之食尸鬼",              // "Shadow Ghoul"  "シャドウ·グール"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   5,
		La:      ygo.LA_Dark,   // 暗
		Lr:      ygo.LR_Zombie, // 不死
		Attack:  1600,
		Defense: 1300,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterGraveAccessArea(func(c *ygo.Card) bool {
				return c.IsMonster()
			}, func(c *ygo.Card) {
				ca.SetAttack(ca.GetAttack() + 100)
			}, func(c *ygo.Card) {
				ca.SetAttack(ca.GetAttack() - 100)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*83*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 438
		 调整: [雷龙]
		<サンダー·ドラゴン>
		[2010/09/08]

		●可以从手卡把这张卡丢弃，从卡组把另外的最多2张「雷龙」加入手卡。那之后卡组洗切。这个效果只能在自己的主要阶段使用。
		◇启动效果（进入连锁）
		◇发动时把这张卡从手卡丢弃（代价）
		◇效果处理时至少选择1张加入手牌
		◇确认卡组中不存在可以检索的卡时，不能发动这个效果
		 中文名: 雷龙
		 日文名: サンダー·ドラゴン
		 英文名: Thunder Dragon
		 卡片种类: 效果怪兽
		 卡片密码: 31786629
		 使用限制: 无限制
		 种族: 雷
		 属性: 光
		 星级: 5
		 攻击力: 1600
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: ME，BE02，VOL07，DL04，Booster07
		 效果: 丢弃手卡的这张卡，从卡组选出最多2张的「雷龙」加入手卡。之后洗卡。这个效果只能在自己的主要阶段使用。
		}
		*/
		Id:       438,
		Password: "31786629",
		Name:     "雷龙",                 // "Thunder Dragon"  "サンダー·ドラゴン"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   5,
		La:      ygo.LA_Light,   // 光
		Lr:      ygo.LR_Thunder, // 雷
		Attack:  1600,
		Defense: 1500,
		Initialize: func(ca *ygo.Card) bool {
			ca.AddEvent(ygo.InHand, func() {
				ca.RegisterIgnitionSelector(ygo.MP, func(pl0 *ygo.Player) {
					pl := ca.GetSummoner()
					if pl == pl0 {
						css := ygo.NewCards(pl.Deck, func(c0 *ygo.Card) bool {
							return c0.GetId() == ca.GetId()
						})
						if css.Len() > 0 {
							ca.PushChain(func() {
								ca.Dispatch(ygo.Cost)
								for i := 0; i != 2; i++ {
									if c := pl.SelectForPopup(css); c != nil {
										c.ToHand()
									}
								}
							})
						}
					}
				})
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*84*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 440
		 调整: [雷仙人]
		<雷仙人>
		[2010/09/08]

		●反转：回复3000基本分。
		◇诱发效果，强制发动，开连锁，不取对象
		◇这张卡被攻击在信息确认时变为表侧的场合，效果延迟到怪兽被破坏确定时之后发动（具体请参考伤害步骤的详细说明）
		●这张卡从场上送去墓地时，失去5000基本分。
		◇诱发效果，强制发动，开连锁，不取对象
		◇基本分回复的效果还没有发动、此卡就被直接送去墓地的场合，这个效果不发动
		◇失去基本分不是伤害效果
		◇在墓地发动的效果
		 中文名: 雷仙人
		 日文名: 雷仙人
		 英文名: The Immortal of Thunder
		 卡片种类: 效果怪兽
		 卡片密码: 84926738
		 使用限制: 无限制
		 种族: 雷
		 属性: 光
		 星级: 4
		 攻击力: 1500
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: ME，BE02，VOL07，DL04
		 效果: 反转：基本分回复3000。这张卡从场上送去墓地的时候，基本分失去5000分。
		}
		*/
		Id:       440,
		Password: "84926738",
		Name:     "雷仙人",                // "The Immortal of Thunder"  "雷仙人"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   4,
		La:      ygo.LA_Light,   // 光
		Lr:      ygo.LR_Thunder, // 雷
		Attack:  1500,
		Defense: 1300,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterFlip(func() {
				pl := ca.GetSummoner()
				pl.ChangeHp(3000)
			})
			ca.AddEvent(ygo.OutMzone, func() {
				pl := ca.GetSummoner()
				e1 := func() {
					pl.ChangeHp(-5000)
				}
				e2 := func() {
					ca.RemoveEvent(ygo.InGrave, e1)
				}
				ca.OnlyOnce(ygo.InHand, e2)
				ca.OnlyOnce(ygo.InRemoved, e2)
				ca.OnlyOnce(ygo.InGrave, e1)
			})
			return true
		}, // 初始

		IsValid: true,
	})

	/*85*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 442
		 调整: [超级流星]
		<スーパースター>
		[2010/09/08]

		●只要这张卡在场上表侧表示存在，全部光属性怪兽的攻击力上升500分。暗属性怪兽的攻击力下降400分。
		◇永续效果（不进入连锁）
		 中文名: 超级流星
		 日文名: スーパースター
		 英文名: Hoshiningen
		 卡片种类: 效果怪兽
		 卡片密码: 67629977
		 使用限制: 无限制
		 种族: 天使
		 属性: 光
		 星级: 2
		 攻击力: 500
		 防御力: 700
		 罕见度: 平卡N，银字R
		 卡包: ME，BE02，VOL06，DL04，Booster06
		 效果: 只要这张卡场上表侧表示存在，全部的光属性怪兽攻击力上升500。暗属性的怪兽攻击力下降400。
		}
		*/
		Id:       442,
		Password: "67629977",
		Name:     "超级流星",               // "Hoshiningen"  "スーパースター"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   2,
		La:      ygo.LA_Light, // 光
		Lr:      ygo.LR_Angel, // 天使
		Attack:  500,
		Defense: 700,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterAllMzoneHalo(func(c *ygo.Card) bool {
				return c.AttrIsLight() || c.AttrIsDark()
			}, func(c *ygo.Card) {
				if c.AttrIsLight() {
					c.SetAttack(c.GetAttack() + 500)
				} else if c.AttrIsDark() {
					c.SetAttack(c.GetAttack() - 400)
				}
			}, func(c *ygo.Card) {
				if c.AttrIsLight() {
					c.SetAttack(c.GetAttack() - 500)
				} else if c.AttrIsDark() {
					c.SetAttack(c.GetAttack() + 400)
				}
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*86*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 443
		 中文名: 音乐家帝王
		 日文名: 音楽家の帝王
		 英文名: Musician King
		 卡片种类: 融合怪兽
		 卡片密码: 56907389
		 使用限制: 无限制
		 种族: 魔法师
		 属性: 光
		 星级: 5
		 攻击力: 1750
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: ME，VOL06，DL04，TP10
		 效果: 融合：「黑森林的魔女」＋「高等女祭司」。
		}
		*/
		Id:       443,
		Password: "56907389",
		Name:     "音乐家帝王",              // "Musician King"  "音楽家の帝王"
		Lc:       ygo.LC_FusionMonster, // 融合怪兽

		Level:   5,
		La:      ygo.LA_Light,       // 光
		Lr:      ygo.LR_SpellCaster, // 魔法师
		Attack:  1750,
		Defense: 1500,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterFusionMonster("黑森林的魔女", "高等女祭司")
			return true
		}, // 初始
		IsValid: true,
	})

	/*87*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 445
		 中文名: 机械恐龙
		 日文名: メカ·ザウルス
		 英文名: Cyber Saurus
		 卡片种类: 融合怪兽
		 卡片密码: 89112729
		 使用限制: 无限制
		 种族: 机械
		 属性: 地
		 星级: 5
		 攻击力: 1800
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: ME，VOL07，DL04，TP19
		 效果: 融合：「炸弹先生」＋「双头恐龙王」。
		}
		*/
		Id:       445,
		Password: "89112729",
		Name:     "机械恐龙",               // "Cyber Saurus"  "メカ·ザウルス"
		Lc:       ygo.LC_FusionMonster, // 融合怪兽

		Level:   5,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Machine, // 机械
		Attack:  1800,
		Defense: 1400,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterFusionMonster("炸弹先生", "双头恐龙王")
			return true
		}, // 初始
		IsValid: true,
	})

	/*88*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 446
		 调整: [加农炮兵]
		<キャノン·ソルジャー>
		[2010/09/08]

		●每把自己场上存在的1只怪兽作为祭品，给予对方基本分500分的伤害。
		◇启动效果（进入连锁）
		◇发动时把自己场上存在的1只怪兽作为祭品（代价）
		 中文名: 加农炮兵
		 日文名: キャノン·ソルジャー
		 英文名: Cannon Soldier
		 卡片种类: 效果怪兽
		 卡片密码: 11384280
		 使用限制: 无限制
		 种族: 机械
		 属性: 暗
		 星级: 4
		 攻击力: 1400
		 防御力: 1300
		 罕见度: 平卡N，银字R，面闪SR
		 卡包: ME，BE02，VOL06，DL04，SD10
		 效果: 每祭品1只自己场上存在的怪兽，对方受到500分的基本分伤害。
		}
		*/
		Id:       446,
		Password: "11384280",
		Name:     "加农炮兵",               // "Cannon Soldier"  "キャノン·ソルジャー"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   4,
		La:      ygo.LA_Dark,    // 暗
		Lr:      ygo.LR_Machine, // 机械
		Attack:  1400,
		Defense: 1300,
		//Initialize:    func(ca *ygo.Card) bool {}, // 初始
		IsValid: false,
	})

	/*89*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 447
		 调整: [姆卡姆卡]
		<ムカムカ>
		[2010/09/08]

		●只要这张卡在场上表侧表示存在，控制者每有1张手卡这张卡的攻击力和守备力上升300分。
		◇永续效果（不进入连锁）
		 中文名: 姆卡姆卡
		 日文名: ムカムカ
		 英文名: Muka Muka
		 卡片种类: 效果怪兽
		 卡片密码: 46657337
		 使用限制: 无限制
		 种族: 岩石
		 属性: 地
		 星级: 2
		 攻击力: 600
		 防御力: 300
		 罕见度: 平卡N
		 卡包: ME，BE02，VOL06，DL04
		 效果: 只要这张卡表侧表示在场上存在，控制者手上每有1张卡，这张卡的攻击力·守备力上升300。
		}
		*/
		Id:       447,
		Password: "46657337",
		Name:     "姆卡姆卡",               // "Muka Muka"  "ムカムカ"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   2,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Rock,  // 岩石
		Attack:  600,
		Defense: 300,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterHandAccessArea(func(c *ygo.Card) bool {
				return true
			}, func() {
				ca.SetAttack(ca.GetAttack() + 300)
				ca.SetDefense(ca.GetDefense() + 300)
			}, func() {
				ca.SetAttack(ca.GetAttack() - 300)
				ca.SetDefense(ca.GetDefense() - 300)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*90*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 449
		 调整: [海星小子]
		<スター·ボーイ>
		[2010/09/08]

		●只要这张卡在场上表侧表示存在，全部水属性怪兽的攻击力上升500分。炎属性怪兽的攻击力下降400分。
		◇永续效果（不进入连锁）
		 中文名: 海星小子
		 日文名: スター·ボーイ
		 英文名: Star Boy
		 卡片种类: 效果怪兽
		 卡片密码: 08201910
		 使用限制: 无限制
		 种族: 水
		 属性: 水
		 星级: 2
		 攻击力: 550
		 防御力: 500
		 罕见度: 平卡N
		 卡包: ME，BE02，VOL06，DL04，Booster06，SD04
		 效果: 只要这张卡在场上表侧表示存在，全部水属性的怪兽攻击力上升500。炎属性的怪兽攻击力下降400。
		}
		*/
		Id:       449,
		Password: "08201910",
		Name:     "海星小子",               // "Star Boy"  "スター·ボーイ"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   2,
		La:      ygo.LA_Water, // 水
		Lr:      ygo.LR_None,  // 水
		Attack:  550,
		Defense: 500,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterAllMzoneHalo(func(c *ygo.Card) bool {
				return c.AttrIsWater() || c.AttrIsFire()
			}, func(c *ygo.Card) {
				if c.AttrIsWater() {
					c.SetAttack(c.GetAttack() + 500)
				} else if c.AttrIsFire() {
					c.SetAttack(c.GetAttack() - 400)
				}
			}, func(c *ygo.Card) {
				if c.AttrIsWater() {
					c.SetAttack(c.GetAttack() - 500)
				} else if c.AttrIsFire() {
					c.SetAttack(c.GetAttack() + 400)
				}
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*91*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 450
		 调整: [金毛犼]
		<ミリス·レディエント>
		[2010/09/08]

		●只要这张卡在场上表侧表示存在，全部地属性怪兽的攻击力上升500分。风属性怪兽的攻击力下降400分。
		◇永续效果（不进入连锁）
		 中文名: 金毛犼
		 日文名: ミリス·レディエント
		 英文名: Milus Radiant
		 卡片种类: 效果怪兽
		 卡片密码: 07489323
		 使用限制: 无限制
		 种族: 兽
		 属性: 地
		 星级: 1
		 攻击力: 300
		 防御力: 250
		 罕见度: 平卡N
		 卡包: ME，BE02，VOL06，DL04，JY
		 效果: 只要这张卡在场上表侧表示存在，全部地属性的怪兽攻击力上升500。风属性的怪兽攻击力下降400。
		}
		*/
		Id:       450,
		Password: "07489323",
		Name:     "金毛犼",                // "Milus Radiant"  "ミリス·レディエント"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   1,
		La:      ygo.LA_Earth, // 地
		Lr:      ygo.LR_Beast, // 兽
		Attack:  300,
		Defense: 250,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterAllMzoneHalo(func(c *ygo.Card) bool {
				return c.AttrIsEarth() || c.AttrIsWind()
			}, func(c *ygo.Card) {
				if c.AttrIsEarth() {
					c.SetAttack(c.GetAttack() + 500)
				} else if c.AttrIsWind() {
					c.SetAttack(c.GetAttack() - 400)
				}
			}, func(c *ygo.Card) {
				if c.AttrIsEarth() {
					c.SetAttack(c.GetAttack() - 500)
				} else if c.AttrIsWind() {
					c.SetAttack(c.GetAttack() + 400)
				}
			})

			return true
		}, // 初始
		IsValid: true,
	})

	/*92*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 451
		 调整: [暗精灵]
		<ダーク·エルフ>
		[2010/09/08]

		●这张卡不支付1000基本分就不能攻击。
		◇进行攻击宣言的代价（不进入连锁）
		◇「技能抽取/スキルドレイン」的效果适用中，进行攻击宣言不需要支付基本分
		 中文名: 暗精灵
		 日文名: ダーク·エルフ
		 英文名: Dark Elf
		 卡片种类: 效果怪兽
		 卡片密码: 21417692
		 使用限制: 无限制
		 种族: 魔法师
		 属性: 暗
		 星级: 4
		 攻击力: 2000
		 防御力: 800
		 罕见度: 平卡N
		 卡包: ME，VOL07，DL04
		 效果: 这张卡不支付1000基本分不能攻击。
		}
		*/
		Id:       451,
		Password: "21417692",
		Name:     "暗精灵",                // "Dark Elf"  "ダーク·エルフ"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   4,
		La:      ygo.LA_Dark,        // 暗
		Lr:      ygo.LR_SpellCaster, // 魔法师
		Attack:  2000,
		Defense: 800,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterPay(func(s string) {
				if s != ygo.Declaration {
					return
				}
				pl := ca.GetSummoner()
				pl.ChangeHp(-1000)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*93*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 452
		 调整: [蘑菇怪人]
		<マタンゴ>
		[2010/09/08]

		●每次自己的准备阶段，给予控制者300分的伤害。
		◇诱发效果（进入连锁）
		◇必须发动
		●在自己的结束阶段支付500基本分，这张卡的控制权转移给对方。
		◇诱发效果（进入连锁）
		◇任意发动
		◇发动时支付500基本分（代价）
		★效果处理时这张卡变为里侧表示的场合如何处理 调整中
		★效果处理时对方怪兽区域不存在空位的场合如何处理 调整中
		 中文名: 蘑菇怪人
		 日文名: マタンゴ
		 英文名: Mushroom Man #2
		 卡片种类: 效果怪兽
		 卡片密码: 93900406
		 使用限制: 无限制
		 种族: 战士
		 属性: 地
		 星级: 3
		 攻击力: 1250
		 防御力: 800
		 罕见度: 平卡N
		 卡包: ME，VOL07，DL04
		 效果: 每次的自己准备阶段受到300分的伤害。可以在自己的结束阶段支付500分，把这张卡的控制权转移给对方。
		}
		*/
		Id:       452,
		Password: "93900406",
		Name:     "蘑菇怪人",               // "Mushroom Man #2"  "マタンゴ"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   3,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Attack:  1250,
		Defense: 800,
		Initialize: func(ca *ygo.Card) bool {
			ca.AddEvent(ygo.FaceUp, func() {

				var e func()
				e = func() {
					ca.RegisterGlobalListen(ygo.SP, func(pl0 *ygo.Player) {
						pl := ca.GetSummoner()
						if pl != pl0 {
							return
						}
						pl.ChangeHp(-300)
					})
					ca.RegisterIgnitionSelector(ygo.EP, func(pl0 *ygo.Player) {
						pl := ca.GetSummoner()
						if pl != pl0 {
							return
						}
						ca.PushChain(func() {
							pl.ChangeHp(-500)
							tar := pl.GetTarget()
							ca.SetSummoner(tar)
							ca.ToMzone()
							e()
						})
					})
				}
				e()
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*94*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 453
		 调整: [野蛮人1号]
		<バーバリアン１号>
		[2010/09/08]

		●自己的场上每有1只表侧表示存在的「野蛮人２号」，这张卡的攻击力上升500分。
		◇永续效果（不进入连锁）
		 中文名: 野蛮人1号
		 日文名: バーバリアン１号
		 英文名: Lava Battleguard
		 卡片种类: 效果怪兽
		 卡片密码: 20394040
		 使用限制: 无限制
		 种族: 战士
		 属性: 地
		 星级: 5
		 攻击力: 1550
		 防御力: 1800
		 罕见度: 平卡N
		 卡包: ME，VOL07，DL04，AT06
		 效果: 每有1只自己场上表侧表示存在的「野蛮人2号」，这张卡的攻击力上升500。
		}
		*/
		Id:       453,
		Password: "20394040",
		Name:     "野蛮人1号",              // "Lava Battleguard"  "バーバリアン１号"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   5,
		La:      ygo.LA_Earth,   // 地
		Lr:      ygo.LR_Warrior, // 战士
		Attack:  1550,
		Defense: 1800,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterMzoneAccessArea(func(c *ygo.Card) bool {
				return c.GetName() == "野蛮人2号"
			}, func(c *ygo.Card) {
				ca.SetAttack(ca.GetAttack() + 500)

			}, func(c *ygo.Card) {
				ca.SetAttack(ca.GetAttack() - 500)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*95*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 454
		 调整: [黑森林的魔女]
		<黒き森のウィッチ>
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
		 日文名: 黒き森のウィッチ
		 英文名: Witch of the Black Forest
		 卡片种类: 效果怪兽
		 卡片密码: 78010363
		 使用限制: 禁止卡
		 种族: 魔法师
		 属性: 暗
		 星级: 4
		 攻击力: 1100
		 防御力: 1200
		 罕见度: 银字R
		 卡包: ME，BE02，VOL06，DL04，JY，PE
		 效果: 这张卡从场上送去墓地时，选1只自己卡组的守备力1500以下的怪兽，互相确认后加入自己手卡。之后洗卡组。
		}
		*/
		Id:       454,
		Password: "78010363",
		Name:     "黑森林的魔女",             // "Witch of the Black Forest"  "黒き森のウィッチ"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   4,
		La:      ygo.LA_Dark,        // 暗
		Lr:      ygo.LR_SpellCaster, // 魔法师
		Attack:  1100,
		Defense: 1200,
		Initialize: func(ca *ygo.Card) bool {
			ca.AddEvent(ygo.InMzone, func() {
				ca.AddEvent(ygo.InGrave, func() {
					pl := ca.GetSummoner()
					cs := pl.Deck.Find(func(c *ygo.Card) bool {
						return c.IsMonster() && c.GetDefense() <= 1500
					})
					if c := pl.SelectForPopup(cs); c != nil {
						c.SetFaceUp()
						c.ToHand()
					}
				})
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*96*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 455
		 调整: [小奇美拉]
		<リトル·キメラ>
		[2010/09/08]

		●只要这张卡在场上表侧表示存在，全部炎属性怪兽的攻击力上升500分。水属性怪兽的攻击力下降400分。
		◇永续效果（不进入连锁）
		 中文名: 小奇美拉
		 日文名: リトル·キメラ
		 英文名: Little Chimera
		 卡片种类: 效果怪兽
		 卡片密码: 68658728
		 使用限制: 无限制
		 种族: 兽
		 属性: 炎
		 星级: 2
		 攻击力: 600
		 防御力: 550
		 罕见度: 平卡N
		 卡包: ME，BE02，VOL06，DL04，Booster06，SD03，SD24
		 效果: 只要这张卡在场上表侧表示存在，全部炎属性的怪兽攻击力上升500。水属性的怪兽攻击力下降400。
		}
		*/
		Id:       455,
		Password: "68658728",
		Name:     "小奇美拉",               // "Little Chimera"  "リトル·キメラ"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   2,
		La:      ygo.LA_None,  // 炎
		Lr:      ygo.LR_Beast, // 兽
		Attack:  600,
		Defense: 550,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterAllMzoneHalo(func(c *ygo.Card) bool {
				return c.AttrIsFire() || c.AttrIsWater()
			}, func(c *ygo.Card) {
				if c.AttrIsFire() {
					c.SetAttack(c.GetAttack() + 500)
				} else if c.AttrIsWater() {
					c.SetAttack(c.GetAttack() - 400)
				}
			}, func(c *ygo.Card) {
				if c.AttrIsFire() {
					c.SetAttack(c.GetAttack() - 500)
				} else if c.AttrIsWater() {
					c.SetAttack(c.GetAttack() + 400)
				}
			})

			return true
		}, // 初始
		IsValid: true,
	})

	/*97*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 456
		 调整: [刃蝇]
		<ブレードフライ>
		[2010/09/08]

		●只要这张卡在场上表侧表示存在，全部风属性怪兽的攻击力上升500分。地属性怪兽的攻击力下降400分。
		◇永续效果（不进入连锁）
		 中文名: 刃蝇
		 日文名: ブレードフライ
		 英文名: Bladefly
		 卡片种类: 效果怪兽
		 卡片密码: 28470714
		 使用限制: 无限制
		 种族: 昆虫
		 属性: 风
		 星级: 2
		 攻击力: 600
		 防御力: 700
		 罕见度: 平卡N
		 卡包: ME，BE02，VOL06，DL04，Booster06，SD08
		 效果: 只要这张卡在场上表侧表示存在，全部风属性的怪兽攻击力上升500。地属性的怪兽攻击力下降400。
		}
		*/
		Id:       456,
		Password: "28470714",
		Name:     "刃蝇",                 // "Bladefly"  "ブレードフライ"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   2,
		La:      ygo.LA_Wind,   // 风
		Lr:      ygo.LR_Insect, // 昆虫
		Attack:  600,
		Defense: 700,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterAllMzoneHalo(func(c *ygo.Card) bool {
				return c.AttrIsWind() || c.AttrIsEarth()
			}, func(c *ygo.Card) {
				if c.AttrIsWind() {
					c.SetAttack(c.GetAttack() + 500)
				} else if c.AttrIsEarth() {
					c.SetAttack(c.GetAttack() - 400)
				}
			}, func(c *ygo.Card) {
				if c.AttrIsWind() {
					c.SetAttack(c.GetAttack() - 500)
				} else if c.AttrIsEarth() {
					c.SetAttack(c.GetAttack() + 400)
				}
			})

			return true
		}, // 初始
		IsValid: true,
	})

	/*98*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 458
		 中文名: 双头雷龙
		 日文名: 双頭の雷龍
		 英文名: Twin-Headed Thunder Dragon
		 卡片种类: 融合怪兽
		 卡片密码: 54752875
		 使用限制: 无限制
		 种族: 雷
		 属性: 光
		 星级: 7
		 攻击力: 2800
		 防御力: 2100
		 罕见度: 平卡N
		 卡包: ME，BE02，VOL07，DL04
		 效果: 融合：「雷龙」＋「雷龙」。
		}
		*/
		Id:       458,
		Password: "54752875",
		Name:     "双头雷龙",               // "Twin-Headed Thunder Dragon"  "双頭の雷龍"
		Lc:       ygo.LC_FusionMonster, // 融合怪兽

		Level:   7,
		La:      ygo.LA_Light,   // 光
		Lr:      ygo.LR_Thunder, // 雷
		Attack:  2800,
		Defense: 2100,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterFusionMonster("雷龙", "雷龙")
			return true
		}, // 初始
		IsValid: true,
	})

	/*99*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 459
		 调整: [见习魔女]
		<見習い魔女>
		[2010/09/08]

		●只要这张卡在场上表侧表示存在，全部暗属性怪兽的攻击力上升500分。光属性怪兽的攻击力下降400分。
		◇永续效果（不进入连锁）
		 中文名: 见习魔女
		 日文名: 見習い魔女
		 英文名: Witch's Apprentice
		 卡片种类: 效果怪兽
		 卡片密码: 80741828
		 使用限制: 无限制
		 种族: 魔法师
		 属性: 暗
		 星级: 2
		 攻击力: 550
		 防御力: 500
		 罕见度: 平卡N
		 卡包: ME，BE02，VOL06，DL04，Booster06
		 效果: 只要这张卡在场上表侧表示存在，全部暗属性的怪兽攻击力上升500。光属性的怪兽攻击力下降400。
		}
		*/
		Id:       459,
		Password: "80741828",
		Name:     "见习魔女",               // "Witch's Apprentice"  "見習い魔女"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   2,
		La:      ygo.LA_Dark,        // 暗
		Lr:      ygo.LR_SpellCaster, // 魔法师
		Attack:  550,
		Defense: 500,
		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterAllMzoneHalo(func(c *ygo.Card) bool {
				return c.AttrIsDark() || c.AttrIsLight()
			}, func(c *ygo.Card) {
				if c.AttrIsDark() {
					c.SetAttack(c.GetAttack() + 500)
				} else if c.AttrIsLight() {
					c.SetAttack(c.GetAttack() - 400)
				}
			}, func(c *ygo.Card) {
				if c.AttrIsDark() {
					c.SetAttack(c.GetAttack() - 500)
				} else if c.AttrIsLight() {
					c.SetAttack(c.GetAttack() + 400)
				}
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*100*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 462
		 调整: [左轮手枪龙]
		<リボルバー·ドラゴン>
		[2010/09/08]

		●选择对方场上存在的1只怪兽发动。指3次硬币，那之内2次以上正面的场合，那只怪兽破坏。这个效果1回合只能使用1次。
		◇启动效果（进入连锁）
		◇发动时选择对方场上存在的1只怪兽（取对象）
		 中文名: 左轮手枪龙
		 日文名: リボルバー·ドラゴン
		 英文名: Barrel Dragon
		 卡片种类: 效果怪兽
		 卡片密码: 81480460
		 使用限制: 无限制
		 种族: 机械
		 属性: 暗
		 星级: 7
		 攻击力: 2600
		 防御力: 2200
		 罕见度: 立体UTR，金字UR，爆闪PR，银碎SER，面闪SR
		 卡包: ME，BE02，VOL07，VB05，DL04，302，DT03
		 效果: 掷3次硬币。在3次中有2次以上掷到正面的形式，破坏对方场上的1只怪兽。这个效果每个回合只能用1次，在自己的主要阶段使用。
		}
		*/
		Id:       462,
		Password: "81480460",
		Name:     "左轮手枪龙",              // "Barrel Dragon"  "リボルバー·ドラゴン"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   7,
		La:      ygo.LA_Dark,    // 暗
		Lr:      ygo.LR_Machine, // 机械
		Attack:  2600,
		Defense: 2200,
		Initialize: func(ca *ygo.Card) bool {
			r := 0
			var e func()
			e = func() {
				ca.RegisterIgnitionSelector(ygo.MP, func(pl0 *ygo.Player) {
					pl := ca.GetSummoner()
					if pl == pl0 && r != pl.GetRound() {
						tar := pl.GetTarget()
						ca.PushChain(func() {
							if ygo.RandInt(3) <= 1 {
								if c := pl.SelectForWarn(tar.Mzone); c != nil {
									c.Dispatch(ygo.Destroy, ca)
								}
							}
							r = pl.GetRound()
							e()
						})
					}
				})
			}
			ca.AddEvent(ygo.FaceUp, e)
			return true

		}, // 初始
		IsValid: true,
	})

	/*101*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 463
		 调整: [神之宣告]
		<神の宣告>
		[2010/09/08]

		●支付一半基本分发动。魔法·陷阱卡的发动，怪兽的召唤·反转召唤·特殊召唤的其中之1无效并破坏。
		◇发动时支付一半基本分（代价）
		◇不能对应「喧闹的邪恶灵/ポルターガイスト」「电脑增幅器/電脳増幅器」的发动而发动
		◇不能对应由于卡的效果特殊召唤的怪兽而发动（「死皇帝之陵墓/死皇帝の陵墓」「紧急同调/緊急同調」「血之代偿/血の代償」除外）
		◇基本分奇数的状态，发动这张卡的场合基本分以四舍五入计算
		◇不能对应魔法或陷阱卡效果的发动而发动
		 中文名: 神之宣告
		 日文名: 神の宣告
		 英文名: Solemn Judgment
		 卡片种类: 反击陷阱
		 卡片密码: 41420027
		 使用限制: 限制卡
		 罕见度: 平卡N，面闪SR，黄金GR
		 卡包: ME，BE02，VOL06，DL04，SD11，SD14，GLD02，GS02，SD20，SD28
		 效果: 把基本分支付一半发动。魔法·陷阱卡的发动，怪兽的召唤·反转召唤·特殊召唤的其中1个无效并破坏。
		}
		*/
		Id:       463,
		Password: "41420027",
		Name:     "神之宣告",              // "Solemn Judgment"  "神の宣告"
		Lc:       ygo.LC_ReactionTrap, // 反击陷阱

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterPay(func(s string) {
				if s != ygo.Trigger {
					return
				}
				pl := ca.GetSummoner()
				pl.ChangeHp(pl.GetHp() / -2)
			})
			e := func(c *ygo.Card) {
				ca.PushChain(func() {
					c.Dispatch(ygo.Destroy, ca)
				})
			}
			ca.RegisterOrdinaryTrap(ygo.UseMagic, e)
			ca.RegisterOrdinaryTrap(ygo.UseTrap, e)
			ca.RegisterOrdinaryTrap(ygo.Summon, e)
			ca.RegisterOrdinaryTrap(ygo.SummonFlip, e)
			ca.RegisterOrdinaryTrap(ygo.SummonSpecial, e)
			return true
		}, // 初始
		IsValid: true,
	})

	/*102*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 464
		 调整: [魔法干扰阵]
		<マジック·ジャマー>
		[14/04/27]

		●①：魔法卡发动时，丢弃1张手卡才能发动。那个发动无效并破坏。
		◇丢弃1张手卡是发动COST
		◇不能对应魔法卡的“效果的发动”（举例：此卡不能对应已经表侧表示存在于场上的[波动加农炮/波動キャノン]的效果的发动）
		 中文名: 魔法干扰阵
		 日文名: マジック·ジャマー
		 英文名: Magic Jammer
		 卡片种类: 反击陷阱
		 卡片密码: 77414722
		 使用限制: 无限制
		 罕见度: 金字UR，面闪SR，平卡N
		 卡包: ME，BE02，YSD02，VOL06，DL04，SD02，SD05，SD08，SD11，SD13，DT01，DTP01，YSD05，YU，SY2，SK2，ST14
		 效果: ①：魔法卡发动时，丢弃1张手卡才能发动。那个发动无效并破坏。
		}
		*/
		Id:       464,
		Password: "77414722",
		Name:     "魔法干扰阵",             // "Magic Jammer"  "マジック·ジャマー"
		Lc:       ygo.LC_ReactionTrap, // 反击陷阱

		Initialize: func(ca *ygo.Card) bool {

			ca.RegisterPay(func(s string) {
				if s != ygo.Trigger {
					return
				}
				pl := ca.GetSummoner()
				if c := pl.SelectForWarn(pl.Hand, func(c0 *ygo.Card) bool {
					return c0 != ca
				}); c != nil {
					c.Dispatch(ygo.Cost, ca)
				} else {
					ca.StopOnce(s)
				}

			})

			ca.RegisterOrdinaryTrap(ygo.UseMagic, func(c *ygo.Card) {
				ca.PushChain(func() {
					c.Dispatch(ygo.Destroy, ca)
				})
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*103*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 465
		 调整: [盗贼的七道具]
		<盗賊の七つ道具>
		[14/04/27]

		●①：陷阱卡发动时，支付1000基本分才能发动。那个发动无效并破坏。
		◇支付1000LP（基本分）是发动COST
		◇不能对应陷阱卡的“效果的发动”（举例：此卡不能对应已经表侧表示存在于场上的[血之代偿/血の代償]的效果的发动）
		 中文名: 盗贼的七道具
		 日文名: 盗賊の七つ道具
		 英文名: Seven Tools of the Bandit
		 卡片种类: 反击陷阱
		 卡片密码: 03819470
		 使用限制: 无限制
		 罕见度: 金字UR，面闪SR，平卡N
		 卡包: ME，BE02，VOL06，DL04，SD11，YSD06，TU05，ST12，ST14
		 效果: ①：陷阱卡发动时，支付1000基本分才能发动。那个发动无效并破坏。
		}
		*/
		Id:       465,
		Password: "03819470",
		Name:     "盗贼的七道具",            // "Seven Tools of the Bandit"  "盗賊の七つ道具"
		Lc:       ygo.LC_ReactionTrap, // 反击陷阱

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterPay(func(s string) {
				if s != ygo.Trigger {
					return
				}
				pl := ca.GetSummoner()
				pl.ChangeHp(-1000)
			})

			ca.RegisterOrdinaryTrap(ygo.UseTrap, func(c *ygo.Card) {
				ca.PushChain(func() {
					c.Dispatch(ygo.Destroy, ca)
				})
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*104*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 466
		 调整: [升天之角笛]
		<昇天の角笛>
		[2010/09/08]

		●自己场上的1只怪兽作为祭品。怪兽的召唤·反转召唤·特殊召唤无效并破坏。
		◇发动时把自己场上的1只怪兽作为祭品（代价）
		◇不能对应由于卡的效果特殊召唤的怪兽而发动（「死皇帝之陵墓/死皇帝の陵墓」「紧急同调/緊急同調」「血之代偿/血の代償」除外）
		 中文名: 升天之角笛
		 日文名: 昇天の角笛
		 英文名: Horn of Heaven
		 卡片种类: 反击陷阱
		 卡片密码: 98069388
		 使用限制: 无限制
		 罕见度: 面闪SR，银字R
		 卡包: ME，BE02，VOL06，DL04
		 效果: 自己场上的1只怪兽作为祭品。怪兽的召唤·反转召唤·特殊召唤无效并破坏。
		}
		*/
		Id:       466,
		Password: "98069388",
		Name:     "升天之角笛",             // "Horn of Heaven"  "昇天の角笛"
		Lc:       ygo.LC_ReactionTrap, // 反击陷阱

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterPay(func(s string) {
				if s != ygo.Trigger {
					return
				}
				pl := ca.GetSummoner()
				if c := pl.SelectForWarn(pl.Mzone); c != nil {
					c.Dispatch(ygo.Cost, ca)
				} else {
					ca.StopOnce(s)
				}
			})
			e := func(c *ygo.Card) {
				ca.PushChain(func() {
					c.Dispatch(ygo.Destroy, ca)
				})
			}
			ca.RegisterOrdinaryTrap(ygo.Summon, e)
			ca.RegisterOrdinaryTrap(ygo.SummonFlip, e)
			ca.RegisterOrdinaryTrap(ygo.SummonSpecial, e)
			return true
		}, // 初始
		IsValid: true,
	})

	/*105*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 467
		 调整: [右手持盾左手持剑]
		<右手に盾を左手に剣を>
		[2010/09/08]

		●直到结束阶段结束时为止，这张卡的发动时在场上的全部表侧表示怪兽的原本的攻击力与守备力交换。
		◇交换的数值是原本数值
		◇效果处理时在场上表侧表示存在的怪兽全部适用（不是叙述上的“发动时”）
		 中文名: 右手持盾左手持剑
		 日文名: 右手に盾を左手に剣を
		 英文名: Shield & Sword
		 卡片种类: 通常魔法
		 卡片密码: 52097679
		 使用限制: 无限制
		 罕见度: 银字R，平卡N
		 卡包: ME，BE02，VOL07，DL04，SD07，JY，SJ2
		 效果: 这个回合结束前，这张卡发动时全部场上存在的表侧表示的怪兽原来的攻击力和原来的守备力交替。
		}
		*/
		Id:       467,
		Password: "52097679",
		Name:     "右手持盾左手持剑",           // "Shield & Sword"  "右手に盾を左手に剣を"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				var e func(c *ygo.Card) bool
				e = func(c *ygo.Card) bool {
					if c.IsFaceUp() {
						c.SetAttack(c.GetAttack() - c.GetBaseAttack() + c.GetBaseDefense())
						c.SetDefense(c.GetDefense() - c.GetBaseDefense() + c.GetBaseAttack())
						pl.OnlyOnce(ygo.RoundEnd, func() {
							c.SetAttack(c.GetAttack() + c.GetBaseAttack() - c.GetBaseDefense())
							c.SetDefense(c.GetDefense() + c.GetBaseDefense() - c.GetBaseAttack())
						}, ca, c)
					}
					return true
				}

				pl.Mzone.ForEach(e)
				tar.Mzone.ForEach(e)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*106*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 468
		 调整: [执念之剑]
		<執念の剣>
		[2010/09/08]

		●装备怪兽的攻击力·守备力上升500分。
		◇不进入连锁
		●这张卡送去墓地时，这张卡回到卡组的最上方。
		◇进入连锁（1速）
		◇必须发动
		◇效果处理时这张卡不在墓地存在的场合，这个效果不适用
		 中文名: 执念之剑
		 日文名: 執念の剣
		 英文名: Sword of Deep-Seated
		 卡片种类: 装备魔法
		 卡片密码: 98495314
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: ME，VOL07，DL04，TP13
		 效果: 装备的怪兽攻击力·守备力上升500。这张卡送去墓地时，回到卡组最上面。
		}
		*/
		Id:       468,
		Password: "98495314",
		Name:     "执念之剑",            // "Sword of Deep-Seated"  "執念の剣"
		Lc:       ygo.LC_EquipMagic, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.AddEvent(ygo.InGrave, func() {
				ca.ToDeck()
			})
			ca.RegisterEquipMagic(func(c *ygo.Card) bool {
				return true
			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() + 500)
				c.SetDefense(c.GetDefense() + 500)

			}, func(c *ygo.Card) {
				c.SetAttack(c.GetAttack() - 500)
				c.SetDefense(c.GetDefense() - 500)

			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*107*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 469
		 调整: [「攻击」封禁]
		<『攻撃』封じ>
		[2010/09/08]

		●指定对方场上的1只表侧攻击表示怪兽变为表侧守备表示。
		◇发动时选择对方场上表侧攻击表示存在的1只怪兽（取对象）
		◇效果处理时那只怪兽已经是表侧守备表示的场合，这个效果不适用
		◇效果处理时那只怪兽不在场上表侧表示存在的场合，这个效果不适用
		◇「最终突击命令/最終突撃命令」效果适用中，这张卡把怪兽变为守备表示后立刻变为攻击表示
		 中文名: 「攻击」封禁
		 日文名: 『攻撃』封じ
		 英文名: Block Attack
		 卡片种类: 通常魔法
		 卡片密码: 25880422
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: ME，BE02，VOL07，DL04
		 效果: 指定的对方场上的1只表侧表示的怪兽转为守备表示。
		}
		*/
		Id:       469,
		Password: "25880422",
		Name:     "「攻击」封禁",             // "Block Attack"  "『攻撃』封じ"
		Lc:       ygo.LC_OrdinaryMagic, // 通常魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryMagic(func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				if c := pl.SelectForWarn(tar.Mzone, func(c0 *ygo.Card) bool {
					return c0.IsFaceUp()
				}); c != nil {
					c.SetFaceDefense()

				}
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*108*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 470
		 调整: [不幸的美少女]
		<薄幸の美少女>
		[2010/09/08]

		●这张卡被战斗破坏并送去墓地时，那个回合的战斗阶段结束。
		◇诱发效果（进入连锁）
		◇必须发动
		 中文名: 不幸的美少女
		 日文名: 薄幸の美少女
		 英文名: The Unhappy Maiden
		 卡片种类: 效果怪兽
		 卡片密码: 51275027
		 使用限制: 无限制
		 种族: 魔法师
		 属性: 光
		 星级: 1
		 攻击力: 0
		 防御力: 100
		 罕见度: 平卡N
		 卡包: ME，BE02，VOL07，DL04
		 效果: 这张卡被战斗破坏送去墓地，当时那个回合的战斗阶段结束。
		}
		*/
		Id:       470,
		Password: "51275027",
		Name:     "不幸的美少女",             // "The Unhappy Maiden"  "薄幸の美少女"
		Lc:       ygo.LC_EffectMonster, // 效果怪兽

		Level:   1,
		La:      ygo.LA_Light,       // 光
		Lr:      ygo.LR_SpellCaster, // 魔法师
		Attack:  0,
		Defense: 100,
		Initialize: func(ca *ygo.Card) bool {
			ca.AddEvent(ygo.DestroyBeBattle, func() {
				pl := ca.GetSummoner()
				tar := pl.GetTarget()
				tar.Mzone.ForEach(func(c *ygo.Card) bool {
					c.SetNotCanAttack()
					return true
				})
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*109*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 471
		 调整: [哥布林拦路贼]
		<追い剥ぎゴブリン>
		[2010/09/08]

		●自己场上的怪兽每次给予对方玩家战斗伤害，对方随机丢弃1张手卡。
		◇效果发动进入连锁（1速）
		◇不能在伤害步骤发动这张卡
		◇效果发动时点同「天空骑士 柏修斯/天空騎士パーシアス」（具体请参考关于伤害步骤的详解）
		◇「次元壁/ディメンション·ウォール」将战斗伤害转移给对方的场合，这个效果不发动
		 中文名: 哥布林拦路贼
		 日文名: 追い剥ぎゴブリン
		 英文名: Robbin' Goblin
		 卡片种类: 永续陷阱
		 卡片密码: 88279736
		 使用限制: 无限制
		 罕见度: 银字R，平卡N
		 卡包: ME，BE02，VOL07，DL04，SD07，GLD04
		 效果: 自己场上的怪兽每次造成对方玩家的战斗伤害时，对方随机丢弃1张手卡。
		}
		*/
		Id:       471,
		Password: "88279736",
		Name:     "哥布林拦路贼",            // "Robbin' Goblin"  "追い剥ぎゴブリン"
		Lc:       ygo.LC_SustainsTrap, // 永续陷阱

		Initialize: func(ca *ygo.Card) bool {
			use := false
			ca.RegisterUnordinaryTrap(ygo.Deduct, func(tar *ygo.Player, c *ygo.Card) {
				pl := ca.GetSummoner()
				e := func() {
					if c.GetSummoner() == pl && pl != tar {
						tar.Hand.Get(ygo.RandInt(tar.Hand.Len())).Dispatch(ygo.Discard, ca)
					}
				}
				if use {
					e()
				} else {
					if c.GetSummoner() == pl && pl != tar {
						ca.PushChain(func() {
							e()
							use = true
						})
					}
				}
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*110*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 472
		 调整: [细菌感染]
		<細菌感染>
		[2010/09/08]

		●机械族以外的怪兽才能装备。装备怪兽的攻击力在每次自己的准备阶段下降300分。
		◇发动时选择场上表侧表示存在的机械族以外的怪兽（取对象）
		◇效果处理时对象怪兽变为机械族的场合，这张卡送去墓地
		◇效果适用中对象怪兽变为机械族的场合，这张卡破坏
		◇「王宮の勅命/王宮の勅命」效果适用不影响这张卡的回合计算，在「王宮の勅命/王宮の勅命」效果失去后，装备怪兽降低那个回合数×300分的攻击力
		 中文名: 细菌感染
		 日文名: 細菌感染
		 英文名: Germ Infection
		 卡片种类: 装备魔法
		 卡片密码: 24668830
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: ME，VOL07，DL04，Booster07
		 效果: 机械族以外的怪兽装备可能。装备怪兽的攻击力在每次的自己的准备阶段攻击力下降300。
		}
		*/
		Id:       472,
		Password: "24668830",
		Name:     "细菌感染",            // "Germ Infection"  "細菌感染"
		Lc:       ygo.LC_EquipMagic, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			var e func()
			tar := ca.GetSummoner()
			ca.RegisterEquipMagic(func(c *ygo.Card) bool {
				return !c.RaceIsMachine()
			}, func(c *ygo.Card) {
				tar = c.GetSummoner()
				e = func() {
					c.SetAttack(c.GetAttack() - 300)
				}
				tar.AddEvent(ygo.SP, e, c)

			}, func(c *ygo.Card) {
				tar.RemoveEvent(ygo.SP, e, c)
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*111*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 473
		 调整: [麻药]
		<しびれ薬>
		[2010/09/08]

		●机械族以外的怪兽才能装备。装备怪兽不能攻击宣言。
		◇发动时选择场上表侧表示存在的机械族以外的怪兽（取对象）
		◇效果处理时对象怪兽变为机械族的场合，这张卡送去墓地
		◇效果适用中对象怪兽变为机械族的场合，这张卡破坏
		 中文名: 麻药
		 日文名: しびれ薬
		 英文名: Paralyzing Potion
		 卡片种类: 装备魔法
		 卡片密码: 50152549
		 使用限制: 无限制
		 罕见度: 平卡N
		 卡包: ME，VOL07，DL04，Booster07
		 效果: 机械族以外的怪兽装备可能。装备怪兽不能攻击宣言。
		}
		*/
		Id:       473,
		Password: "50152549",
		Name:     "麻药",              // "Paralyzing Potion"  "しびれ薬"
		Lc:       ygo.LC_EquipMagic, // 装备魔法

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterEquipMagic(func(c *ygo.Card) bool {
				return !c.RaceIsMachine()
			}, func(c *ygo.Card) {
				ca.RegisterGlobalListen(ygo.BP, func(pl0 *ygo.Player) {
					pl := c.GetSummoner()
					if pl != pl0 {
						return
					}
					c.SetNotCanAttack()
				})
			}, func(c *ygo.Card) {})
			return true
		}, // 初始
		IsValid: true,
	})

	/*112*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 474
		 调整: [神圣防护罩-反射镜力-]
		<聖なるバリア－ミラーフォース－>
		[14/04/27]

		●①：对方怪兽的攻击宣言时才能发动。对方场上的攻击表示怪兽全部破坏。
		◇不取对象
		◇对方场上不存在攻击表示怪兽的场合，这张卡不能发动（举例：「绝对防御将军/絶対防御将軍」以表侧守备表示攻击时）
		◇里侧攻击表示的怪兽也破坏
		◇效果处理时，对方场上不存在攻击表示怪兽的场合，这个效果不适用。
		 中文名: 神圣防护罩-反射镜力-
		 日文名: 聖なるバリア －ミラーフォース－
		 英文名: Mirror Force
		 卡片种类: 通常陷阱
		 卡片密码: 44095762
		 使用限制: 无限制
		 罕见度: 金字UR，面闪SR，爆闪PR，黄金GR，平卡N，银碎SER
		 卡包: ME，BE02，VOL07，DL04，GLD01，GS01，DPYG，SD19，YU，SY2，SDM，YSD06，DB12，ST13，15AY，ST14
		 效果: ①：对方怪兽的攻击宣言时才能发动。对方场上的攻击表示怪兽全部破坏。
		}
		*/
		Id:       474,
		Password: "44095762",
		Name:     "神圣防护罩-反射镜力",        // "Mirror Force"  "聖なるバリア －ミラーフォース－"
		Lc:       ygo.LC_OrdinaryTrap, // 通常陷阱

		Initialize: func(ca *ygo.Card) bool {
			ca.RegisterOrdinaryTrap(ygo.Declaration, func(c *ygo.Card) {
				if obj := c.GetSummoner(); obj == ca.GetSummoner() {
					tar := obj.GetTarget()
					css := ygo.NewCards(tar.Mzone, func(c0 *ygo.Card) bool {
						return c0.IsAttack()
					})
					if css.Len() > 0 {
						ca.PushChain(func() {
							css.ForEach(func(c0 *ygo.Card) bool {
								c0.Dispatch(ygo.Destroy, ca)
								return true
							})
						})
					}
				}
			})
			return true
		}, // 初始
		IsValid: true,
	})

	/*113*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 1221
		 中文名: 单摆刃拷问机械
		 日文名: 振り子刃の拷問機械
		 英文名: Pendulum Machine
		 卡片种类: 通常怪兽
		 卡片密码: 24433920
		 使用限制: 无限制
		 种族: 机械
		 属性: 暗
		 星级: 6
		 攻击力: 1750
		 防御力: 2000
		 罕见度: 金字UR，平卡N
		 卡包: LE02，VOL07，GLD04
		 效果: 描述：使用带有大摆子的刀刃将对方劈成两半！可怕的拷问机器呀！
		}
		*/
		Id:       1221,
		Password: "24433920",
		Name:     "单摆刃拷问机械",              // "Pendulum Machine"  "振り子刃の拷問機械"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    6,
		La:       ygo.LA_Dark,    // 暗
		Lr:       ygo.LR_Machine, // 机械
		Attack:   1750,
		Defense:  2000,
		IsValid:  true,
	})

	/*114*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 140
		 中文名: 独眼巨人
		 日文名: サイクロプス
		 英文名: Hitotsu-Me Giant
		 卡片种类: 通常怪兽
		 卡片密码: 76184692
		 使用限制: 无限制
		 种族: 兽战士
		 属性: 地
		 星级: 4
		 攻击力: 1200
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: LB，BE01，EX-R(EX)，VOL01，DL02，DPKB
		 效果: 描述：只长有一只眼睛的巨人。利用巨腕殴打敌人，值得小心。
		}
		*/
		Id:       140,
		Password: "76184692",
		Name:     "独眼巨人",                 // "Hitotsu-Me Giant"  "サイクロプス"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Earth,       // 地
		Lr:       ygo.LR_BeastWarror, // 兽战士
		Attack:   1200,
		Defense:  1000,
		IsValid:  true,
	})

	/*115*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 143
		 中文名: 黑魔术师
		 日文名: ブラック·マジシャン
		 英文名: Dark Magician
		 卡片种类: 通常怪兽
		 卡片密码: 46986414
		 使用限制: 无限制
		 种族: 魔法师
		 属性: 暗
		 星级: 7
		 攻击力: 2500
		 防御力: 2100
		 罕见度: 金字UR，爆闪PR，银字R，平卡N，立体UTR
		 卡包: LB，BE01，PP04，LN，EX-R(EX)，VOL01，DL02，SD06，DT01，WJMP，DPYG，DTP01，YU，SY2，LC01，15AY
		 效果: 描述：作为魔法师，攻击力·守备力都是最高级别。
		}
		*/
		Id:       143,
		Password: "46986414",
		Name:     "黑魔术师",                 // "Dark Magician"  "ブラック·マジシャン"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    7,
		La:       ygo.LA_Dark,        // 暗
		Lr:       ygo.LR_SpellCaster, // 魔法师
		Attack:   2500,
		Defense:  2100,
		IsValid:  true,
	})

	/*116*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 144
		 中文名: 暗黑骑士 盖亚
		 日文名: 暗黒騎士ガイア
		 英文名: Gaia The Fierce Knight
		 卡片种类: 通常怪兽
		 卡片密码: 06368038
		 使用限制: 无限制
		 种族: 战士
		 属性: 地
		 星级: 7
		 攻击力: 2300
		 防御力: 2100
		 罕见度: 金字UR，银字R，立体UTR，平卡N
		 卡包: LB，BE01，LE02，PH，EX-R(EX)，VOL01，DL02，Booster R1，15AY
		 效果: 描述：骑着风驰电掣般的马的骑士。当心突进攻击。
		}
		*/
		Id:       144,
		Password: "06368038",
		Name:     "暗黑骑士 盖亚",              // "Gaia The Fierce Knight"  "暗黒騎士ガイア"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    7,
		La:       ygo.LA_Earth,   // 地
		Lr:       ygo.LR_Warrior, // 战士
		Attack:   2300,
		Defense:  2100,
		IsValid:  true,
	})

	/*117*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 147
		 中文名: 猛犸的墓场
		 日文名: マンモスの墓場
		 英文名: Mammoth Graveyard
		 卡片种类: 通常怪兽
		 卡片密码: 40374923
		 使用限制: 无限制
		 种族: 恐龙
		 属性: 地
		 星级: 3
		 攻击力: 1200
		 防御力: 800
		 罕见度: 平卡N
		 卡包: LB，BE01，EX-R(EX)，VOL01，DL02，15AY
		 效果: 描述：守护着伙伴墓地的猛犸，对于任何胆敢践踏墓地的盗墓者进行毫不客气的攻击。
		}
		*/
		Id:       147,
		Password: "40374923",
		Name:     "猛犸的墓场",                // "Mammoth Graveyard"  "マンモスの墓場"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Earth,    // 地
		Lr:       ygo.LR_Dinosaur, // 恐龙
		Attack:   1200,
		Defense:  800,
		IsValid:  true,
	})

	/*118*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 148
		 中文名: 银牙狼
		 日文名: シルバー·フォング
		 英文名: Silver Fang
		 卡片种类: 通常怪兽
		 卡片密码: 90357090
		 使用限制: 无限制
		 种族: 兽
		 属性: 地
		 星级: 3
		 攻击力: 1200
		 防御力: 800
		 罕见度: 平卡N
		 卡包: LB，BE01，EX-R(EX)，VOL01，DL02，YU，AT05
		 效果: 描述：闪烁着银光的狼，看起来虽美丽，但性格却十分凶暴。
		}
		*/
		Id:       148,
		Password: "90357090",
		Name:     "银牙狼",                  // "Silver Fang"  "シルバー·フォング"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Earth, // 地
		Lr:       ygo.LR_Beast, // 兽
		Attack:   1200,
		Defense:  800,
		IsValid:  true,
	})

	/*119*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 181
		 调整: [14/04/27]
		此卡是通常怪物
		 中文名: 圣精灵
		 日文名: ホーリー·エルフ
		 英文名: Mystical Elf
		 卡片种类: 通常怪兽
		 卡片密码: 15025844
		 使用限制: 无限制
		 种族: 魔法师
		 属性: 光
		 星级: 4
		 攻击力: 800
		 防御力: 2000
		 罕见度: 平卡N
		 卡包: PG，BE01，EX-R(EX)，VOL02，DL02，SY2，ST13，ST14，15AY
		 效果: 描述：虽然是纤弱的精灵，但以神圣力量护身有很强的守备。
		}
		*/
		Id:       181,
		Password: "15025844",
		Name:     "圣精灵",                  // "Mystical Elf"  "ホーリー·エルフ"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Light,       // 光
		Lr:       ygo.LR_SpellCaster, // 魔法师
		Attack:   800,
		Defense:  2000,
		IsValid:  true,
	})

	/*120*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 182
		 中文名: 大炮鸟
		 日文名: タイホーン
		 英文名: Tyhone
		 卡片种类: 通常怪兽
		 卡片密码: 72842870
		 使用限制: 无限制
		 种族: 鸟兽
		 属性: 风
		 星级: 4
		 攻击力: 1200
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: PG，VOL02，DL02，TP17
		 效果: 描述：从口中打出炮弹攻击远处。在山上的炮击很强劲。
		}
		*/
		Id:       182,
		Password: "72842870",
		Name:     "大炮鸟",                  // "Tyhone"  "タイホーン"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Wind,      // 风
		Lr:       ygo.LR_WindBeast, // 鸟兽
		Attack:   1200,
		Defense:  1400,
		IsValid:  true,
	})

	/*121*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 183
		 中文名: 路易斯
		 日文名: ルイーズ
		 英文名: Beaver Warrior
		 卡片种类: 通常怪兽
		 卡片密码: 32452818
		 使用限制: 无限制
		 种族: 兽战士
		 属性: 地
		 星级: 4
		 攻击力: 1200
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: PG，BE01，EX-R(EX)，VOL03，DL02，YU，SY2，15AY
		 效果: 描述：身体虽小，但在草原上具备高强的守备力。
		}
		*/
		Id:       183,
		Password: "32452818",
		Name:     "路易斯",                  // "Beaver Warrior"  "ルイーズ"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Earth,       // 地
		Lr:       ygo.LR_BeastWarror, // 兽战士
		Attack:   1200,
		Defense:  1500,
		IsValid:  true,
	})

	/*122*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 185
		 中文名: 诅咒之龙
		 日文名: カース·オブ·ドラゴン
		 英文名: Curse of Dragon
		 卡片种类: 通常怪兽
		 卡片密码: 28279543
		 使用限制: 无限制
		 种族: 龙
		 属性: 暗
		 星级: 5
		 攻击力: 2000
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: PG，BE01，EX-R(EX)，VOL02，DL02，15AY
		 效果: 描述：邪恶的龙。使用暗之力的攻击很强力。
		}
		*/
		Id:       185,
		Password: "28279543",
		Name:     "诅咒之龙",                 // "Curse of Dragon"  "カース·オブ·ドラゴン"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Dark,   // 暗
		Lr:       ygo.LR_Dragon, // 龙
		Attack:   2000,
		Defense:  1500,
		IsValid:  true,
	})

	/*123*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 186
		 中文名: 岩石巨兵
		 日文名: 岩石の巨兵
		 英文名: Giant Soldier of Stone
		 卡片种类: 通常怪兽
		 卡片密码: 13039848
		 使用限制: 无限制
		 种族: 岩石
		 属性: 地
		 星级: 3
		 攻击力: 1300
		 防御力: 2000
		 罕见度: 平卡N
		 卡包: PG，BE01，EX-R(EX)，YSD02，VOL03，DL02，YU，SY2，15AY
		 效果: 描述：岩石的巨人兵。粗壮手腕的攻击让大地动摇。
		}
		*/
		Id:       186,
		Password: "13039848",
		Name:     "岩石巨兵",                 // "Giant Soldier of Stone"  "岩石の巨兵"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Earth, // 地
		Lr:       ygo.LR_Rock,  // 岩石
		Attack:   1300,
		Defense:  2000,
		IsValid:  true,
	})

	/*124*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 187
		 中文名: 荒野盗龙
		 日文名: ワイルド·ラプター
		 英文名: Uraby
		 卡片种类: 通常怪兽
		 卡片密码: 01784619
		 使用限制: 无限制
		 种族: 恐龙
		 属性: 地
		 星级: 4
		 攻击力: 1500
		 防御力: 800
		 罕见度: 平卡N
		 卡包: PG，BE01，EX-R(EX)，VOL02，DL02
		 效果: 描述：擅长奔跑的恐龙，利用尖锐的爪子进行攻击。
		}
		*/
		Id:       187,
		Password: "01784619",
		Name:     "荒野盗龙",                 // "Uraby"  "ワイルド·ラプター"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Earth,    // 地
		Lr:       ygo.LR_Dinosaur, // 恐龙
		Attack:   1500,
		Defense:  800,
		IsValid:  true,
	})

	/*125*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 189
		 中文名: 魔人 死亡撒旦
		 日文名: 魔人デスサタン
		 英文名: Witty Phantom
		 卡片种类: 通常怪兽
		 卡片密码: 36304921
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 1400
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: PG，EX-R(EX)，VOL03，DL02
		 效果: 描述：穿着融入黑暗的黑色燕尾服，操纵着死亡的恶魔。
		}
		*/
		Id:       189,
		Password: "36304921",
		Name:     "魔人 死亡撒旦",              // "Witty Phantom"  "魔人デスサタン"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Dark,  // 暗
		Lr:       ygo.LR_Fiend, // 恶魔
		Attack:   1400,
		Defense:  1300,
		IsValid:  true,
	})

	/*126*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 190
		 中文名: 竖琴精灵
		 日文名: ハープの精
		 英文名: Spirit of the Harp
		 卡片种类: 通常怪兽
		 卡片密码: 80770678
		 使用限制: 无限制
		 种族: 天使
		 属性: 光
		 星级: 4
		 攻击力: 800
		 防御力: 2000
		 罕见度: 平卡N
		 卡包: PG，VOL02，DL02，YSD06，ST12
		 效果: 描述：在天界弹奏竖琴的精灵。那音色使周围听众的心变得平静。
		}
		*/
		Id:       190,
		Password: "80770678",
		Name:     "竖琴精灵",                 // "Spirit of the Harp"  "ハープの精"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Light, // 光
		Lr:       ygo.LR_Angel, // 天使
		Attack:   800,
		Defense:  2000,
		IsValid:  true,
	})

	/*127*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 191
		 中文名: 魔人 泰拉
		 日文名: 魔人 テラ
		 英文名: Terra the Terrible
		 卡片种类: 通常怪兽
		 卡片密码: 63308047
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 1200
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: PG，EX-R(EX)，VOL02，DL02
		 效果: 描述：住在沼泽地的恶魔手下，看起来不怎么强，但绝不能对他轻心大意。
		}
		*/
		Id:       191,
		Password: "63308047",
		Name:     "魔人 泰拉",                // "Terra the Terrible"  "魔人 テラ"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Dark,  // 暗
		Lr:       ygo.LR_Fiend, // 恶魔
		Attack:   1200,
		Defense:  1300,
		IsValid:  true,
	})

	/*128*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 192
		 中文名: 恍惚的人鱼
		 日文名: 恍惚の人魚
		 英文名: Enchanting Mermaid
		 卡片种类: 通常怪兽
		 卡片密码: 75376965
		 使用限制: 无限制
		 种族: 鱼
		 属性: 水
		 星级: 3
		 攻击力: 1200
		 防御力: 900
		 罕见度: 平卡N
		 卡包: PG，VOL02，DL02，TP15
		 效果: 描述：诱惑海上航海者使之溺水的美丽人鱼。
		}
		*/
		Id:       192,
		Password: "75376965",
		Name:     "恍惚的人鱼",                // "Enchanting Mermaid"  "恍惚の人魚"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Water, // 水
		Lr:       ygo.LR_Fish,  // 鱼
		Attack:   1200,
		Defense:  900,
		IsValid:  true,
	})

	/*129*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 193
		 中文名: 炎之魔神
		 日文名: 炎の魔神
		 英文名: Fireyarou
		 卡片种类: 通常怪兽
		 卡片密码: 71407486
		 使用限制: 无限制
		 种族: 炎
		 属性: 炎
		 星级: 4
		 攻击力: 1300
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: PG，VOL03，DL02，TP23
		 效果: 描述：被吞没在火中的魔人。运用自如地操纵周围的火焰攻击。
		}
		*/
		Id:       193,
		Password: "71407486",
		Name:     "炎之魔神",                 // "Fireyarou"  "炎の魔神"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Fire, // 炎
		Lr:       ygo.LR_Pyro, // 炎
		Attack:   1300,
		Defense:  1000,
		IsValid:  true,
	})

	/*130*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 209
		 中文名: 尖刺海龙
		 日文名: スパイクシードラ
		 英文名: Spike Seadra
		 卡片种类: 通常怪兽
		 卡片密码: 85326399
		 使用限制: 无限制
		 种族: 海龙
		 属性: 水
		 星级: 5
		 攻击力: 1600
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: PG，VOL03，DL02
		 效果: 描述：用身体的刺攻击敌人，并能放出电流。
		}
		*/
		Id:       209,
		Password: "85326399",
		Name:     "尖刺海龙",                 // "Spike Seadra"  "スパイクシードラ"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Water,      // 水
		Lr:       ygo.LR_Seaserpent, // 海龙
		Attack:   1600,
		Defense:  1300,
		IsValid:  true,
	})

	/*131*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 210
		 中文名: 天空猎手
		 日文名: スカイ·ハンター
		 英文名: Skull Red Bird
		 卡片种类: 通常怪兽
		 卡片密码: 10202894
		 使用限制: 无限制
		 种族: 鸟兽
		 属性: 风
		 星级: 4
		 攻击力: 1550
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: PG，EX-R(EX)，VOL03，DL02
		 效果: 描述：在羽毛里藏着小刀，从空中降下发起攻击。
		}
		*/
		Id:       210,
		Password: "10202894",
		Name:     "天空猎手",                 // "Skull Red Bird"  "スカイ·ハンター"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Wind,      // 风
		Lr:       ygo.LR_WindBeast, // 鸟兽
		Attack:   1550,
		Defense:  1200,
		IsValid:  true,
	})

	/*132*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 213
		 中文名: 沙石怪
		 日文名: サンド·ストーン
		 英文名: Sand Stone
		 卡片种类: 通常怪兽
		 卡片密码: 73051941
		 使用限制: 无限制
		 种族: 岩石
		 属性: 地
		 星级: 5
		 攻击力: 1300
		 防御力: 1600
		 罕见度: 平卡N
		 卡包: PG，VOL03，DL02
		 效果: 描述：从地下突然出现在眼前，用触手做攻击。
		}
		*/
		Id:       213,
		Password: "73051941",
		Name:     "沙石怪",                  // "Sand Stone"  "サンド·ストーン"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Earth, // 地
		Lr:       ygo.LR_Rock,  // 岩石
		Attack:   1300,
		Defense:  1600,
		IsValid:  true,
	})

	/*133*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 215
		 中文名: 钢铁巨神像
		 日文名: 鋼鉄の巨神像
		 英文名: Steel Ogre Grotto #1
		 卡片种类: 通常怪兽
		 卡片密码: 29172562
		 使用限制: 无限制
		 种族: 机械
		 属性: 地
		 星级: 5
		 攻击力: 1400
		 防御力: 1800
		 罕见度: 平卡N
		 卡包: PG，VOL03，DL02
		 效果: 描述：据域是被机械帝国所祭奠着的钢铁的巨神像。
		}
		*/
		Id:       215,
		Password: "29172562",
		Name:     "钢铁巨神像",                // "Steel Ogre Grotto #1"  "鋼鉄の巨神像"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Earth,   // 地
		Lr:       ygo.LR_Machine, // 机械
		Attack:   1400,
		Defense:  1800,
		IsValid:  true,
	})

	/*134*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 216
		 中文名: 下级龙
		 日文名: レッサー·ドラゴン
		 英文名: Lesser Dragon
		 卡片种类: 通常怪兽
		 卡片密码: 55444629
		 使用限制: 无限制
		 种族: 龙
		 属性: 风
		 星级: 4
		 攻击力: 1200
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: PG，VOL03，DL02
		 效果: 描述：不怎么强，连吐息攻击都不会的低级龙。
		}
		*/
		Id:       216,
		Password: "55444629",
		Name:     "下级龙",                  // "Lesser Dragon"  "レッサー·ドラゴン"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Wind,   // 风
		Lr:       ygo.LR_Dragon, // 龙
		Attack:   1200,
		Defense:  1000,
		IsValid:  true,
	})

	/*135*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 217
		 中文名: 女夜魔骑士
		 日文名: サキュバス·ナイト
		 英文名: Succubus Knight
		 卡片种类: 通常怪兽
		 卡片密码: 55291359
		 使用限制: 无限制
		 种族: 战士
		 属性: 暗
		 星级: 5
		 攻击力: 1650
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: PG，VOL03，DL02
		 效果: 描述：咏唱起魔法，给予对方血的祭礼的恶魔魔法战士。
		}
		*/
		Id:       217,
		Password: "55291359",
		Name:     "女夜魔骑士",                // "Succubus Knight"  "サキュバス·ナイト"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Dark,    // 暗
		Lr:       ygo.LR_Warrior, // 战士
		Attack:   1650,
		Defense:  1300,
		IsValid:  true,
	})

	/*136*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 222
		 中文名: 被封印者的右足
		 日文名: 封印されし者の右足
		 英文名: Right Leg of the Forbidden One
		 卡片种类: 通常怪兽
		 卡片密码: 08124921
		 使用限制: 限制卡
		 种族: 魔法师
		 属性: 暗
		 星级: 1
		 攻击力: 200
		 防御力: 300
		 罕见度: 平卡N，银字R
		 卡包: PG，BE01，VOL04，DL02，15AY
		 效果: 描述：被封印者的右足。封印解开后将得到无限的力量。
		}
		*/
		Id:       222,
		Password: "08124921",
		Name:     "被封印者的右足",              // "Right Leg of the Forbidden One"  "封印されし者の右足"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    1,
		La:       ygo.LA_Dark,        // 暗
		Lr:       ygo.LR_SpellCaster, // 魔法师
		Attack:   200,
		Defense:  300,
		IsValid:  true,
	})

	/*137*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 223
		 中文名: 被封印者的左足
		 日文名: 封印されし者の左足
		 英文名: Left Leg of the Forbidden One
		 卡片种类: 通常怪兽
		 卡片密码: 44519536
		 使用限制: 限制卡
		 种族: 魔法师
		 属性: 暗
		 星级: 1
		 攻击力: 200
		 防御力: 300
		 罕见度: 银字R
		 卡包: PG，BE01，VOL03，DL02，15AY
		 效果: 描述：被封印者的左足。封印解开后将得到无限的力量。
		}
		*/
		Id:       223,
		Password: "44519536",
		Name:     "被封印者的左足",              // "Left Leg of the Forbidden One"  "封印されし者の左足"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    1,
		La:       ygo.LA_Dark,        // 暗
		Lr:       ygo.LR_SpellCaster, // 魔法师
		Attack:   200,
		Defense:  300,
		IsValid:  true,
	})

	/*138*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2257
		 中文名: 小天使
		 日文名: プチテンシ
		 英文名: Petit Angel
		 卡片种类: 通常怪兽
		 卡片密码: 38142739
		 使用限制: 无限制
		 种族: 天使
		 属性: 光
		 星级: 3
		 攻击力: 600
		 防御力: 900
		 罕见度: 平卡N
		 卡包: LB，VOL01，TP15
		 效果: 描述：依靠灵活的行动躲避攻击的小天使。
		}
		*/
		Id:       2257,
		Password: "38142739",
		Name:     "小天使",                  // "Petit Angel"  "プチテンシ"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Light, // 光
		Lr:       ygo.LR_Angel, // 天使
		Attack:   600,
		Defense:  900,
		IsValid:  true,
	})

	/*139*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2258
		 中文名: 暗黑灰羚
		 日文名: ダーク·グレイ
		 英文名: Dark Gray
		 卡片种类: 通常怪兽
		 卡片密码: 09159938
		 使用限制: 无限制
		 种族: 兽
		 属性: 地
		 星级: 3
		 攻击力: 800
		 防御力: 900
		 罕见度: 平卡N
		 卡包: LB，VOL01
		 效果: 描述：全身灰色的野兽，是非常少见的珍贵品种。
		}
		*/
		Id:       2258,
		Password: "09159938",
		Name:     "暗黑灰羚",                 // "Dark Gray"  "ダーク·グレイ"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Earth, // 地
		Lr:       ygo.LR_Beast, // 兽
		Attack:   800,
		Defense:  900,
		IsValid:  true,
	})

	/*140*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2259
		 中文名: 紫炎的影武者
		 日文名: 紫炎の影武者
		 英文名: Kagemusha of the Blue Flame
		 卡片种类: 通常怪兽
		 卡片密码: 15401633
		 使用限制: 无限制
		 种族: 战士
		 属性: 地
		 星级: 2
		 攻击力: 800
		 防御力: 400
		 罕见度: 平卡N
		 卡包: LB，VOL01，TP15
		 效果: 描述：侍奉着紫炎的影武者，持有锋利的名刀。
		}
		*/
		Id:       2259,
		Password: "15401633",
		Name:     "紫炎的影武者",               // "Kagemusha of the Blue Flame"  "紫炎の影武者"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    2,
		La:       ygo.LA_Earth,   // 地
		Lr:       ygo.LR_Warrior, // 战士
		Attack:   800,
		Defense:  400,
		IsValid:  true,
	})

	/*141*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2260
		 中文名: 鸭人
		 日文名: ドレイク
		 英文名: Kurama
		 卡片种类: 通常怪兽
		 卡片密码: 85705804
		 使用限制: 无限制
		 种族: 鸟兽
		 属性: 风
		 星级: 3
		 攻击力: 800
		 防御力: 800
		 罕见度: 平卡N
		 卡包: LB，VOL01，TP14
		 效果: 描述：尾巴很长的鸟，从空中发起攻击。
		}
		*/
		Id:       2260,
		Password: "85705804",
		Name:     "鸭人",                   // "Kurama"  "ドレイク"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Wind,      // 风
		Lr:       ygo.LR_WindBeast, // 鸟兽
		Attack:   800,
		Defense:  800,
		IsValid:  true,
	})

	/*142*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2261
		 中文名: 睡眠之子
		 日文名: 眠り子
		 英文名: Nemuriko
		 卡片种类: 通常怪兽
		 卡片密码: 90963488
		 使用限制: 无限制
		 种族: 魔法师
		 属性: 暗
		 星级: 3
		 攻击力: 800
		 防御力: 700
		 罕见度: 平卡N
		 卡包: LB，VOL01
		 效果: 描述：孩子被梦魔所操控永远沉睡着不再醒来。
		}
		*/
		Id:       2261,
		Password: "90963488",
		Name:     "睡眠之子",                 // "Nemuriko"  "眠り子"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Dark,        // 暗
		Lr:       ygo.LR_SpellCaster, // 魔法师
		Attack:   800,
		Defense:  700,
		IsValid:  true,
	})

	/*143*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2262
		 中文名: 死亡之足
		 日文名: デス·フット
		 英文名: The Drdek
		 卡片种类: 通常怪兽
		 卡片密码: 08944575
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 3
		 攻击力: 700
		 防御力: 800
		 罕见度: 平卡N
		 卡包: VOL01
		 效果: 描述：眼珠上长着脚的怪兽。高跳着用爪钩攻击。
		}
		*/
		Id:       2262,
		Password: "08944575",
		Name:     "死亡之足",                 // "The Drdek"  "デス·フット"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Dark,  // 暗
		Lr:       ygo.LR_Fiend, // 恶魔
		Attack:   700,
		Defense:  800,
		IsValid:  true,
	})

	/*144*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2263
		 中文名: 愤怒的海王
		 日文名: 怒りの海王
		 英文名: The Furious Sea King
		 卡片种类: 通常怪兽
		 卡片密码: 18710707
		 使用限制: 无限制
		 种族: 水
		 属性: 水
		 星级: 3
		 攻击力: 800
		 防御力: 700
		 罕见度: 平卡N
		 卡包: LB，VOL01
		 效果: 描述：伟大的海之王。唤起永不停息的大海啸将敌人吞没。
		}
		*/
		Id:       2263,
		Password: "18710707",
		Name:     "愤怒的海王",                // "The Furious Sea King"  "怒りの海王"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Water, // 水
		Lr:       ygo.LR_None,  // 水
		Attack:   800,
		Defense:  700,
		IsValid:  true,
	})

	/*145*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2264
		 中文名: 生命之沙漏
		 日文名: 命の砂時計
		 英文名: Hourglass of Life
		 卡片种类: 通常怪兽
		 卡片密码: 08783685
		 使用限制: 无限制
		 种族: 天使
		 属性: 光
		 星级: 2
		 攻击力: 700
		 防御力: 600
		 罕见度: 平卡N
		 卡包: VOL01，TP16
		 效果: 描述：掌管生命的天使。通过削减自己的寿命将力量赋予他人。
		}
		*/
		Id:       2264,
		Password: "08783685",
		Name:     "生命之沙漏",                // "Hourglass of Life"  "命の砂時計"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    2,
		La:       ygo.LA_Light, // 光
		Lr:       ygo.LR_Angel, // 天使
		Attack:   700,
		Defense:  600,
		IsValid:  true,
	})

	/*146*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 227
		 中文名: 小精怪
		 日文名: グレムリン
		 英文名: Feral Imp
		 卡片种类: 通常怪兽
		 卡片密码: 41392891
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 1300
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: RB，BE01，EX-R(EX)，VOL05，DL02，SY2，15AY
		 效果: 描述：喜欢恶作剧的小恶魔，从黑暗中袭来。警惕！
		}
		*/
		Id:       227,
		Password: "41392891",
		Name:     "小精怪",                  // "Feral Imp"  "グレムリン"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Dark,  // 暗
		Lr:       ygo.LR_Fiend, // 恶魔
		Attack:   1300,
		Defense:  1400,
		IsValid:  true,
	})

	/*147*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2279
		 中文名: 恶魔海狸
		 日文名: デーモン·ビーバー
		 英文名: Archfiend Marmot of Nefariousness
		 卡片种类: 通常怪兽
		 卡片密码: 75889523
		 使用限制: 无限制
		 种族: 兽
		 属性: 地
		 星级: 2
		 攻击力: 400
		 防御力: 600
		 罕见度: 平卡N
		 卡包: VOL01，TP06
		 效果: 描述：拥有恶魔角与翅膀的海狸。投掷橡实果进行攻击。
		}
		*/
		Id:       2279,
		Password: "75889523",
		Name:     "恶魔海狸",                 // "Archfiend Marmot of Nefariousness"  "デーモン·ビーバー"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    2,
		La:       ygo.LA_Earth, // 地
		Lr:       ygo.LR_Beast, // 兽
		Attack:   400,
		Defense:  600,
		IsValid:  true,
	})

	/*148*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 228
		 中文名: 守城的翼龙
		 日文名: 砦を守る翼竜
		 英文名: Winged Dragon, Guardian of the Fortress #1
		 卡片种类: 通常怪兽
		 卡片密码: 87796900
		 使用限制: 无限制
		 种族: 龙
		 属性: 风
		 星级: 4
		 攻击力: 1400
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: RB，BE01，EX-R(EX)，VOL04，DL02，SY2，15AY
		 效果: 描述：守护山寨的龙，从空中急速下降攻击敌人。
		}
		*/
		Id:       228,
		Password: "87796900",
		Name:     "守城的翼龙",                // "Winged Dragon, Guardian of the Fortress #1"  "砦を守る翼竜"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Wind,   // 风
		Lr:       ygo.LR_Dragon, // 龙
		Attack:   1400,
		Defense:  1200,
		IsValid:  true,
	})

	/*149*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2283
		 中文名: 蜘蛛男
		 日文名: 蜘蛛男
		 英文名: Kumootoko
		 卡片种类: 通常怪兽
		 卡片密码: 56283725
		 使用限制: 无限制
		 种族: 昆虫
		 属性: 地
		 星级: 3
		 攻击力: 700
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: PG，VOL02
		 效果: 描述：巨大的蜘蛛被赋予智慧的形态，通过吐丝将敌人的行动封锁。
		}
		*/
		Id:       2283,
		Password: "56283725",
		Name:     "蜘蛛男",                  // "Kumootoko"  "蜘蛛男"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Earth,  // 地
		Lr:       ygo.LR_Insect, // 昆虫
		Attack:   700,
		Defense:  1400,
		IsValid:  true,
	})

	/*150*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2284
		 中文名: 铠甲剑尾战士
		 日文名: アーメイル
		 英文名: Armaill
		 卡片种类: 通常怪兽
		 卡片密码: 53153481
		 使用限制: 无限制
		 种族: 战士
		 属性: 地
		 星级: 3
		 攻击力: 700
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: PG，VOL02，TP09
		 效果: 描述：有着剑状尾的怪异战士。用双手和尾巴进行三连攻击。
		}
		*/
		Id:       2284,
		Password: "53153481",
		Name:     "铠甲剑尾战士",               // "Armaill"  "アーメイル"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Earth,   // 地
		Lr:       ygo.LR_Warrior, // 战士
		Attack:   700,
		Defense:  1300,
		IsValid:  true,
	})

	/*151*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2286
		 中文名: 独眼盾龙
		 日文名: 一眼の盾竜
		 英文名: One-Eyed Shield Dragon
		 卡片种类: 通常怪兽
		 卡片密码: 33064647
		 使用限制: 无限制
		 种族: 龙
		 属性: 风
		 星级: 3
		 攻击力: 700
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: PG，VOL02
		 效果: 描述：身上的盾不仅能保护身体，还能用来突击。
		}
		*/
		Id:       2286,
		Password: "33064647",
		Name:     "独眼盾龙",                 // "One-Eyed Shield Dragon"  "一眼の盾竜"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Wind,   // 风
		Lr:       ygo.LR_Dragon, // 龙
		Attack:   700,
		Defense:  1300,
		IsValid:  true,
	})

	/*152*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2288
		 中文名: 岩石犰狳
		 日文名: ストーン·アルマジラー
		 英文名: Stone Armadiller
		 卡片种类: 通常怪兽
		 卡片密码: 63432835
		 使用限制: 无限制
		 种族: 岩石
		 属性: 地
		 星级: 3
		 攻击力: 800
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: VOL02
		 效果: 描述：身体被像石头般坚固的毛覆盖着、守备坚固。
		}
		*/
		Id:       2288,
		Password: "63432835",
		Name:     "岩石犰狳",                 // "Stone Armadiller"  "ストーン·アルマジラー"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Earth, // 地
		Lr:       ygo.LR_Rock,  // 岩石
		Attack:   800,
		Defense:  1200,
		IsValid:  true,
	})

	/*153*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2289
		 中文名: 坚硬铠甲
		 日文名: ハードアーマー
		 英文名: Hard Armor
		 卡片种类: 通常怪兽
		 卡片密码: 20060230
		 使用限制: 无限制
		 种族: 战士
		 属性: 地
		 星级: 3
		 攻击力: 300
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: PG，VOL02
		 效果: 描述：拥有心灵的铠甲，坚硬的身体装备在战士身上。
		}
		*/
		Id:       2289,
		Password: "20060230",
		Name:     "坚硬铠甲",                 // "Hard Armor"  "ハードアーマー"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Earth,   // 地
		Lr:       ygo.LR_Warrior, // 战士
		Attack:   300,
		Defense:  1200,
		IsValid:  true,
	})

	/*154*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 229
		 中文名: 恶魔召唤
		 日文名: デーモンの召喚
		 英文名: Summoned Skull
		 卡片种类: 通常怪兽
		 卡片密码: 70781052
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 6
		 攻击力: 2500
		 防御力: 1200
		 罕见度: 平卡N，金字UR，面闪SR，银字R，立体UTR，爆闪PR
		 卡包: RB，BE01，LE03，SC，EX-R(EX)，VOL04，DL02，Booster R3，YAP01，DPYG，DT09，15AY
		 效果: 描述：使用黑暗力量，迷惑人心的恶魔。在恶魔族中以相当强大的力量著称。
		}
		*/
		Id:       229,
		Password: "70781052",
		Name:     "恶魔召唤",                 // "Summoned Skull"  "デーモンの召喚"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    6,
		La:       ygo.LA_Dark,  // 暗
		Lr:       ygo.LR_Fiend, // 恶魔
		Attack:   2500,
		Defense:  1200,
		IsValid:  true,
	})

	/*155*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2292
		 中文名: 全息投影者
		 日文名: ホログラー
		 英文名: Holograh
		 卡片种类: 通常怪兽
		 卡片密码: 10859908
		 使用限制: 无限制
		 种族: 机械
		 属性: 地
		 星级: 3
		 攻击力: 1100
		 防御力: 700
		 罕见度: 平卡N
		 卡包: VOL02
		 效果: 描述：使人看见各种幻想、趁机进行攻击的机器。
		}
		*/
		Id:       2292,
		Password: "10859908",
		Name:     "全息投影者",                // "Holograh"  "ホログラー"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Earth,   // 地
		Lr:       ygo.LR_Machine, // 机械
		Attack:   1100,
		Defense:  700,
		IsValid:  true,
	})

	/*156*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 230
		 中文名: 岩窟魔人
		 日文名: 岩窟魔人オーガ·ロック
		 英文名: Rock Ogre Grotto #1
		 卡片种类: 通常怪兽
		 卡片密码: 68846917
		 使用限制: 无限制
		 种族: 岩石
		 属性: 地
		 星级: 3
		 攻击力: 800
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: RB，VOL04，DL02，JY
		 效果: 描述：岩石之躯成就高强的守备力。挥起巨腕时务必警惕。
		}
		*/
		Id:       230,
		Password: "68846917",
		Name:     "岩窟魔人",                 // "Rock Ogre Grotto #1"  "岩窟魔人オーガ·ロック"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Earth, // 地
		Lr:       ygo.LR_Rock,  // 岩石
		Attack:   800,
		Defense:  1200,
		IsValid:  true,
	})

	/*157*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2303
		 中文名: 暗黑拿破仑
		 日文名: D·ナポレオン
		 英文名: Meda Bat
		 卡片种类: 通常怪兽
		 卡片密码: 76211194
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 2
		 攻击力: 800
		 防御力: 400
		 罕见度: 平卡N
		 卡包: PG，VOL02，TP05
		 效果: 描述：恶人创造的眼珠恶魔。用黑暗炸弹进行爆破攻击。
		}
		*/
		Id:       2303,
		Password: "76211194",
		Name:     "暗黑拿破仑",                // "Meda Bat"  "D·ナポレオン"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    2,
		La:       ygo.LA_Dark,  // 暗
		Lr:       ygo.LR_Fiend, // 恶魔
		Attack:   800,
		Defense:  400,
		IsValid:  true,
	})

	/*158*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2305
		 中文名: 暗杀者
		 日文名: アサシン
		 英文名: Ansatsu
		 卡片种类: 通常怪兽
		 卡片密码: 48365709
		 使用限制: 无限制
		 种族: 战士
		 属性: 地
		 星级: 5
		 攻击力: 1700
		 防御力: 1200
		 罕见度: 平卡N，银字R
		 卡包: EX-R(EX)，VOL03，Booster02
		 效果: 描述：能在黑暗中悄然无声地靠近敌人，精通暗杀的战士。
		}
		*/
		Id:       2305,
		Password: "48365709",
		Name:     "暗杀者",                  // "Ansatsu"  "アサシン"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Earth,   // 地
		Lr:       ygo.LR_Warrior, // 战士
		Attack:   1700,
		Defense:  1200,
		IsValid:  true,
	})

	/*159*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2306
		 中文名: 卡库塔斯
		 日文名: カクタス
		 英文名: Akihiron
		 卡片种类: 通常怪兽
		 卡片密码: 36904469
		 使用限制: 无限制
		 种族: 水
		 属性: 水
		 星级: 5
		 攻击力: 1700
		 防御力: 1400
		 罕见度: 银字R
		 卡包: VOL03，Booster02
		 效果: 描述：潜在水中形状不明的怪兽。
		}
		*/
		Id:       2306,
		Password: "36904469",
		Name:     "卡库塔斯",                 // "Akihiron"  "カクタス"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Water, // 水
		Lr:       ygo.LR_None,  // 水
		Attack:   1700,
		Defense:  1400,
		IsValid:  true,
	})

	/*160*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2307
		 中文名: 机械巨兵
		 日文名: 機械の巨兵
		 英文名: Giant Mech-soldier
		 卡片种类: 通常怪兽
		 卡片密码: 72299832
		 使用限制: 无限制
		 种族: 机械
		 属性: 地
		 星级: 6
		 攻击力: 1750
		 防御力: 1900
		 罕见度: 银字R
		 卡包: VOL03，Booster02
		 效果: 描述：巨斧的一击可以割开大地。
		}
		*/
		Id:       2307,
		Password: "72299832",
		Name:     "机械巨兵",                 // "Giant Mech-soldier"  "機械の巨兵"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    6,
		La:       ygo.LA_Earth,   // 地
		Lr:       ygo.LR_Machine, // 机械
		Attack:   1750,
		Defense:  1900,
		IsValid:  true,
	})

	/*161*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2308
		 中文名: 神圣人偶
		 日文名: ホーリー·ドール
		 英文名: Rogue Doll
		 卡片种类: 通常怪兽
		 卡片密码: 91939608
		 使用限制: 无限制
		 种族: 魔法师
		 属性: 光
		 星级: 4
		 攻击力: 1600
		 防御力: 1000
		 罕见度: 银字R
		 卡包: EX-R(EX)，VOL03，Booster02，PE
		 效果: 描述：操纵神圣力量的人偶。在黑暗之中攻击力很强。
		}
		*/
		Id:       2308,
		Password: "91939608",
		Name:     "神圣人偶",                 // "Rogue Doll"  "ホーリー·ドール"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Light,       // 光
		Lr:       ygo.LR_SpellCaster, // 魔法师
		Attack:   1600,
		Defense:  1000,
		IsValid:  true,
	})

	/*162*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2309
		 中文名: 魔加农
		 日文名: マキャノン
		 英文名: Mabarrel
		 卡片种类: 通常怪兽
		 卡片密码: 98795934
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 5
		 攻击力: 1700
		 防御力: 1400
		 罕见度: 银字R，平卡N
		 卡包: VOL03，Booster02
		 效果: 描述：大炮状的恶魔。以飞快的速度发射眼球弹。
		}
		*/
		Id:       2309,
		Password: "98795934",
		Name:     "魔加农",                  // "Mabarrel"  "マキャノン"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Dark,  // 暗
		Lr:       ygo.LR_Fiend, // 恶魔
		Attack:   1700,
		Defense:  1400,
		IsValid:  true,
	})

	/*163*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 231
		 中文名: 铠蜥蜴
		 日文名: 鎧蜥蜴
		 英文名: Armored Lizard
		 卡片种类: 通常怪兽
		 卡片密码: 15480588
		 使用限制: 无限制
		 种族: 爬虫类
		 属性: 地
		 星级: 4
		 攻击力: 1500
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: RB，BE01，VOL04，DL02，JY，SJ2
		 效果: 描述：拥有坚硬躯体的蜥蜴。如果被咬到会立即毙命。
		}
		*/
		Id:       231,
		Password: "15480588",
		Name:     "铠蜥蜴",                  // "Armored Lizard"  "鎧蜥蜴"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Earth,   // 地
		Lr:       ygo.LR_Reptile, // 爬虫类
		Attack:   1500,
		Defense:  1200,
		IsValid:  true,
	})

	/*164*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2310
		 中文名: 甲化海星
		 日文名: アーマード·スターフィッシュ
		 英文名: Armored Starfish
		 卡片种类: 通常怪兽
		 卡片密码: 17535588
		 使用限制: 无限制
		 种族: 水
		 属性: 水
		 星级: 4
		 攻击力: 850
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: PG，VOL03
		 效果: 描述：表面坚固，守备力较高的青色海星。
		}
		*/
		Id:       2310,
		Password: "17535588",
		Name:     "甲化海星",                 // "Armored Starfish"  "アーマード·スターフィッシュ"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Water, // 水
		Lr:       ygo.LR_None,  // 水
		Attack:   850,
		Defense:  1400,
		IsValid:  true,
	})

	/*165*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2311
		 中文名: 鲜血吮吸者
		 日文名: 生き血をすするもの
		 英文名: Drooling Lizard
		 卡片种类: 通常怪兽
		 卡片密码: 16353197
		 使用限制: 无限制
		 种族: 爬虫类
		 属性: 地
		 星级: 3
		 攻击力: 900
		 防御力: 800
		 罕见度: 平卡N
		 卡包: PG，VOL03
		 效果: 描述：借着夜色袭击行人的人形吸血蛇。
		}
		*/
		Id:       2311,
		Password: "16353197",
		Name:     "鲜血吮吸者",                // "Drooling Lizard"  "生き血をすするもの"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Earth,   // 地
		Lr:       ygo.LR_Reptile, // 爬虫类
		Attack:   900,
		Defense:  800,
		IsValid:  true,
	})

	/*166*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 232
		 中文名: 杀人蜂
		 日文名: キラー·ビー
		 英文名: Killer Needle
		 卡片种类: 通常怪兽
		 卡片密码: 88979991
		 使用限制: 无限制
		 种族: 昆虫
		 属性: 风
		 星级: 4
		 攻击力: 1200
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: RB，VOL05，DL02
		 效果: 描述：巨大的黄蜂，攻击出乎意料的强，若被群体攻击则非常危险。
		}
		*/
		Id:       232,
		Password: "88979991",
		Name:     "杀人蜂",                  // "Killer Needle"  "キラー·ビー"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Wind,   // 风
		Lr:       ygo.LR_Insect, // 昆虫
		Attack:   1200,
		Defense:  1000,
		IsValid:  true,
	})

	/*167*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2330
		 中文名: 神鱼
		 日文名: 神魚
		 英文名: Bottom Dweller
		 卡片种类: 通常怪兽
		 卡片密码: 81386177
		 使用限制: 无限制
		 种族: 鱼
		 属性: 水
		 星级: 5
		 攻击力: 1650
		 防御力: 1700
		 罕见度: 平卡N
		 卡包: RB，VOL04
		 效果: 描述：在水中幽雅游动的鱼神，若它发怒会很危险。
		}
		*/
		Id:       2330,
		Password: "81386177",
		Name:     "神鱼",                   // "Bottom Dweller"  "神魚"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Water, // 水
		Lr:       ygo.LR_Fish,  // 鱼
		Attack:   1650,
		Defense:  1700,
		IsValid:  true,
	})

	/*168*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2331
		 中文名: 耳天使
		 日文名: エンゼル·イヤーズ
		 英文名: Ocubeam
		 卡片种类: 通常怪兽
		 卡片密码: 86088138
		 使用限制: 无限制
		 种族: 天使
		 属性: 光
		 星级: 5
		 攻击力: 1550
		 防御力: 1650
		 罕见度: 平卡N
		 卡包: RB，VOL04
		 效果: 描述：巨大的耳朵和眼睛能监视周围的一切，相貌可怕的天使。
		}
		*/
		Id:       2331,
		Password: "86088138",
		Name:     "耳天使",                  // "Ocubeam"  "エンゼル·イヤーズ"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Light, // 光
		Lr:       ygo.LR_Angel, // 天使
		Attack:   1550,
		Defense:  1650,
		IsValid:  true,
	})

	/*169*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2332
		 中文名: 死之沉默天使 多玛
		 日文名: 死の沈黙の天使 ドマ
		 英文名: Doma The Angel of Silence
		 卡片种类: 通常怪兽
		 卡片密码: 16972957
		 使用限制: 无限制
		 种族: 天使
		 属性: 暗
		 星级: 5
		 攻击力: 1600
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: RB，EX-R(EX)，VOL04
		 效果: 描述：司职死亡的天使，若被这家伙盯上，难逃一死。
		}
		*/
		Id:       2332,
		Password: "16972957",
		Name:     "死之沉默天使 多玛",            // "Doma The Angel of Silence"  "死の沈黙の天使 ドマ"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Dark,  // 暗
		Lr:       ygo.LR_Angel, // 天使
		Attack:   1600,
		Defense:  1400,
		IsValid:  true,
	})

	/*170*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2333
		 中文名: 猎手蜘蛛
		 日文名: ハンター·スパイダー
		 英文名: Hunter Spider
		 卡片种类: 通常怪兽
		 卡片密码: 80141480
		 使用限制: 无限制
		 种族: 昆虫
		 属性: 地
		 星级: 5
		 攻击力: 1600
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: RB，VOL04
		 效果: 描述：布置蜘蛛网狩猎，落入网中的东西将被吃掉。
		}
		*/
		Id:       2333,
		Password: "80141480",
		Name:     "猎手蜘蛛",                 // "Hunter Spider"  "ハンター·スパイダー"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Earth,  // 地
		Lr:       ygo.LR_Insect, // 昆虫
		Attack:   1600,
		Defense:  1400,
		IsValid:  true,
	})

	/*171*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2334
		 中文名: 莫林芬
		 日文名: モリンフェン
		 英文名: Morinphen
		 卡片种类: 通常怪兽
		 卡片密码: 55784832
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 5
		 攻击力: 1550
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: RB，VOL04
		 效果: 描述：以长臂钩爪为特征的长相奇妙的恶魔。
		}
		*/
		Id:       2334,
		Password: "55784832",
		Name:     "莫林芬",                  // "Morinphen"  "モリンフェン"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Dark,  // 暗
		Lr:       ygo.LR_Fiend, // 恶魔
		Attack:   1550,
		Defense:  1300,
		IsValid:  true,
	})

	/*172*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2335
		 中文名: 古代精灵
		 日文名: エンシェント·エルフ
		 英文名: Ancient Elf
		 卡片种类: 通常怪兽
		 卡片密码: 93221206
		 使用限制: 无限制
		 种族: 魔法师
		 属性: 光
		 星级: 4
		 攻击力: 1450
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: RB，EX-R(EX)，VOL04
		 效果: 描述：活了数千年的妖精，能够操纵精灵发动攻击。
		}
		*/
		Id:       2335,
		Password: "93221206",
		Name:     "古代精灵",                 // "Ancient Elf"  "エンシェント·エルフ"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Light,       // 光
		Lr:       ygo.LR_SpellCaster, // 魔法师
		Attack:   1450,
		Defense:  1200,
		IsValid:  true,
	})

	/*173*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2336
		 中文名: 黑影之鬼王
		 日文名: 黒い影の鬼王
		 英文名: Ogre of the Black Shadow
		 卡片种类: 通常怪兽
		 卡片密码: 45121025
		 使用限制: 无限制
		 种族: 兽战士
		 属性: 地
		 星级: 4
		 攻击力: 1200
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: EX-R(EX)，VOL04
		 效果: 描述：被黑影附体的恶鬼。用惊人的速度突击。
		}
		*/
		Id:       2336,
		Password: "45121025",
		Name:     "黑影之鬼王",                // "Ogre of the Black Shadow"  "黒い影の鬼王"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Earth,       // 地
		Lr:       ygo.LR_BeastWarror, // 兽战士
		Attack:   1200,
		Defense:  1400,
		IsValid:  true,
	})

	/*174*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2337
		 中文名: 绝对鸟
		 日文名: アブソリューター
		 英文名: Fiend Reflection#1
		 卡片种类: 通常怪兽
		 卡片密码: 68870276
		 使用限制: 无限制
		 种族: 鸟兽
		 属性: 风
		 星级: 4
		 攻击力: 1300
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: VOL04
		 效果: 描述：可以将对方吸入镜中世界。
		}
		*/
		Id:       2337,
		Password: "68870276",
		Name:     "绝对鸟",                  // "Fiend Reflection#1"  "アブソリューター"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Wind,      // 风
		Lr:       ygo.LR_WindBeast, // 鸟兽
		Attack:   1300,
		Defense:  1400,
		IsValid:  true,
	})

	/*175*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2338
		 中文名: 舌鱼
		 日文名: 舌魚
		 英文名: Tongyo
		 卡片种类: 通常怪兽
		 卡片密码: 69572024
		 使用限制: 无限制
		 种族: 鱼
		 属性: 水
		 星级: 4
		 攻击力: 1350
		 防御力: 800
		 罕见度: 平卡N
		 卡包: RB，VOL04
		 效果: 描述：用长舌头抓住其他的鱼，吸取能量。
		}
		*/
		Id:       2338,
		Password: "69572024",
		Name:     "舌鱼",                   // "Tongyo"  "舌魚"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Water, // 水
		Lr:       ygo.LR_Fish,  // 鱼
		Attack:   1350,
		Defense:  800,
		IsValid:  true,
	})

	/*176*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2339
		 中文名: 龙人
		 日文名: ドラゴヒューマン
		 英文名: D. Human
		 卡片种类: 通常怪兽
		 卡片密码: 81057959
		 使用限制: 无限制
		 种族: 战士
		 属性: 地
		 星级: 4
		 攻击力: 1300
		 防御力: 1100
		 罕见度: 平卡N
		 卡包: EX-R(EX)，VOL04
		 效果: 描述：挥动用龙牙打造的剑、拥有龙之力量的战士。
		}
		*/
		Id:       2339,
		Password: "81057959",
		Name:     "龙人",                   // "D. Human"  "ドラゴヒューマン"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Earth,   // 地
		Lr:       ygo.LR_Warrior, // 战士
		Attack:   1300,
		Defense:  1100,
		IsValid:  true,
	})

	/*177*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 234
		 中文名: 鹰身女郎
		 日文名: ハーピィ·レディ
		 英文名: Harpie Lady
		 卡片种类: 通常怪兽
		 卡片密码: 76812113
		 使用限制: 无限制
		 种族: 鸟兽
		 属性: 风
		 星级: 4
		 攻击力: 1300
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: RB，BE01，VOL04，DL02
		 效果: 描述：人长出羽毛的魔兽，表演华丽的舞蹈进行快攻。
		}
		*/
		Id:       234,
		Password: "76812113",
		Name:     "鹰身女郎",                 // "Harpie Lady"  "ハーピィ·レディ"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Wind,      // 风
		Lr:       ygo.LR_WindBeast, // 鸟兽
		Attack:   1300,
		Defense:  1400,
		IsValid:  true,
	})

	/*178*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2358
		 中文名: 乌鸦天狗
		 日文名: カラス天狗
		 英文名: Crow Goblin
		 卡片种类: 通常怪兽
		 卡片密码: 77998771
		 使用限制: 无限制
		 种族: 鸟兽
		 属性: 风
		 星级: 5
		 攻击力: 1850
		 防御力: 1600
		 罕见度: 银字R
		 卡包: VOL05，Booster05
		 效果: 描述：知道各种事的天狗。据说能使用神通之力。
		}
		*/
		Id:       2358,
		Password: "77998771",
		Name:     "乌鸦天狗",                 // "Crow Goblin"  "カラス天狗"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Wind,      // 风
		Lr:       ygo.LR_WindBeast, // 鸟兽
		Attack:   1850,
		Defense:  1600,
		IsValid:  true,
	})

	/*179*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2359
		 调整: （这张卡在规则上也当作「异虫」卡使用）
		 中文名: 迷宫的蠕虫
		 日文名: ダンジョン·ワーム
		 英文名: Dungeon Worm
		 卡片种类: 通常怪兽
		 卡片密码: 51228280
		 使用限制: 无限制
		 种族: 昆虫
		 属性: 地
		 星级: 5
		 攻击力: 1800
		 防御力: 1500
		 罕见度: 银字R
		 卡包: VOL05，Booster05
		 效果: 描述：（这张卡在规则上也当作「异虫」卡使用）潜在迷路者的地下，捕食经过此处迷路的生物。
		}
		*/
		Id:       2359,
		Password: "51228280",
		Name:     "迷宫的蠕虫",                // "Dungeon Worm"  "ダンジョン·ワーム"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Earth,  // 地
		Lr:       ygo.LR_Insect, // 昆虫
		Attack:   1800,
		Defense:  1500,
		IsValid:  true,
	})

	/*180*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 236
		 中文名: 魔物狩人
		 日文名: 魔物の狩人
		 英文名: Kojikocy
		 卡片种类: 通常怪兽
		 卡片密码: 01184620
		 使用限制: 无限制
		 种族: 战士
		 属性: 地
		 星级: 4
		 攻击力: 1500
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: RB，BE01，EX-R(EX)，VOL04，DL02，SJ2
		 效果: 描述：猎杀人类的猎手，具有连岩石都能击碎的力量。
		}
		*/
		Id:       236,
		Password: "01184620",
		Name:     "魔物狩人",                 // "Kojikocy"  "魔物の狩人"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Earth,   // 地
		Lr:       ygo.LR_Warrior, // 战士
		Attack:   1500,
		Defense:  1200,
		IsValid:  true,
	})

	/*181*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2360
		 中文名: 粗暴帝王
		 日文名: ルード·カイザー
		 英文名: Rude Kaiser
		 卡片种类: 通常怪兽
		 卡片密码: 26378150
		 使用限制: 无限制
		 种族: 兽战士
		 属性: 地
		 星级: 5
		 攻击力: 1800
		 防御力: 1600
		 罕见度: 银字R
		 卡包: EX-R(EX)，VOL05，Booster05
		 效果: 描述：双手所持的魔人斧的破坏力相当强。
		}
		*/
		Id:       2360,
		Password: "26378150",
		Name:     "粗暴帝王",                 // "Rude Kaiser"  "ルード·カイザー"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Earth,       // 地
		Lr:       ygo.LR_BeastWarror, // 兽战士
		Attack:   1800,
		Defense:  1600,
		IsValid:  true,
	})

	/*182*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2361
		 中文名: 暗黑兔
		 日文名: ダーク·ラビット
		 英文名: Dark Rabbit
		 卡片种类: 通常怪兽
		 卡片密码: 99261403
		 使用限制: 无限制
		 种族: 兽
		 属性: 暗
		 星级: 4
		 攻击力: 1100
		 防御力: 1500
		 罕见度: 平卡N，银字R
		 卡包: VOL05，SOVR(606)，PE，TP18
		 效果: 描述：美国卡通世界里面的兔子罗杰。经常敏捷而匆忙的行动。
		}
		*/
		Id:       2361,
		Password: "99261403",
		Name:     "暗黑兔",                  // "Dark Rabbit"  "ダーク·ラビット"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Dark,  // 暗
		Lr:       ygo.LR_Beast, // 兽
		Attack:   1100,
		Defense:  1500,
		IsValid:  true,
	})

	/*183*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2362
		 中文名: 响女
		 日文名: 響女
		 英文名: Hibikime
		 卡片种类: 通常怪兽
		 卡片密码: 64501875
		 使用限制: 无限制
		 种族: 战士
		 属性: 地
		 星级: 4
		 攻击力: 1450
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: RB，VOL05
		 效果: 描述：发出扰乱听觉的噪音，使对方陷入行动不能的状态。
		}
		*/
		Id:       2362,
		Password: "64501875",
		Name:     "响女",                   // "Hibikime"  "響女"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Earth,   // 地
		Lr:       ygo.LR_Warrior, // 战士
		Attack:   1450,
		Defense:  1000,
		IsValid:  true,
	})

	/*184*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2363
		 中文名: 魔法幽灵
		 日文名: マジカル·ゴースト
		 英文名: Magical Ghost
		 卡片种类: 通常怪兽
		 卡片密码: 46474915
		 使用限制: 无限制
		 种族: 不死
		 属性: 暗
		 星级: 4
		 攻击力: 1300
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: EX-R(EX)，VOL05
		 效果: 描述：向对方施展法术，使敌人陷入混乱和恐惧，从而胡乱攻击。
		}
		*/
		Id:       2363,
		Password: "46474915",
		Name:     "魔法幽灵",                 // "Magical Ghost"  "マジカル·ゴースト"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Dark,   // 暗
		Lr:       ygo.LR_Zombie, // 不死
		Attack:   1300,
		Defense:  1400,
		IsValid:  true,
	})

	/*185*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 238
		 中文名: 伏地龙
		 日文名: 地を這うドラゴン
		 英文名: Crawling Dragon
		 卡片种类: 通常怪兽
		 卡片密码: 67494157
		 使用限制: 无限制
		 种族: 龙
		 属性: 地
		 星级: 5
		 攻击力: 1600
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: RB，VOL04，DL02
		 效果: 描述：力量弱小无法飞天的龙，但攻击还是很强的。
		}
		*/
		Id:       238,
		Password: "67494157",
		Name:     "伏地龙",                  // "Crawling Dragon"  "地を這うドラゴン"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Earth,  // 地
		Lr:       ygo.LR_Dragon, // 龙
		Attack:   1600,
		Defense:  1400,
		IsValid:  true,
	})

	/*186*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2381
		 中文名: 海之龙王
		 日文名: 海の竜王
		 英文名: Sea King Dragon
		 卡片种类: 通常怪兽
		 卡片密码: 23659124
		 使用限制: 无限制
		 种族: 海龙
		 属性: 水
		 星级: 6
		 攻击力: 2000
		 防御力: 1700
		 罕见度: 平卡N
		 卡包: VOL06
		 效果: 描述：海之王。长有坚硬的贝壳，从口中吐出水泡攻击。
		}
		*/
		Id:       2381,
		Password: "23659124",
		Name:     "海之龙王",                 // "Sea King Dragon"  "海の竜王"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    6,
		La:       ygo.LA_Water,      // 水
		Lr:       ygo.LR_Seaserpent, // 海龙
		Attack:   2000,
		Defense:  1700,
		IsValid:  true,
	})

	/*187*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2382
		 中文名: 猫女郎
		 日文名: キャット·レディ
		 英文名: Nekogal #2
		 卡片种类: 通常怪兽
		 卡片密码: 43352213
		 使用限制: 无限制
		 种族: 兽战士
		 属性: 地
		 星级: 6
		 攻击力: 1900
		 防御力: 2000
		 罕见度: 平卡N
		 卡包: VOL06
		 效果: 描述：用敏捷的动作躲避攻击、用锋利的钩爪袭击对方。
		}
		*/
		Id:       2382,
		Password: "43352213",
		Name:     "猫女郎",                  // "Nekogal #2"  "キャット·レディ"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    6,
		La:       ygo.LA_Earth,       // 地
		Lr:       ygo.LR_BeastWarror, // 兽战士
		Attack:   1900,
		Defense:  2000,
		IsValid:  true,
	})

	/*188*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2383
		 中文名: 骰子犰狳
		 日文名: ダイス·アルマジロ
		 英文名: Dice Armadillo
		 卡片种类: 通常怪兽
		 卡片密码: 69893315
		 使用限制: 无限制
		 种族: 机械
		 属性: 地
		 星级: 5
		 攻击力: 1650
		 防御力: 1800
		 罕见度: 平卡N
		 卡包: VOL06
		 效果: 描述：卷起身子就像骰子一样的犰狳。
		}
		*/
		Id:       2383,
		Password: "69893315",
		Name:     "骰子犰狳",                 // "Dice Armadillo"  "ダイス·アルマジロ"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Earth,   // 地
		Lr:       ygo.LR_Machine, // 机械
		Attack:   1650,
		Defense:  1800,
		IsValid:  true,
	})

	/*189*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2384
		 中文名: 苍翼冠鸟
		 日文名: 冠を戴く蒼き翼
		 英文名: Blue-Winged Crown
		 卡片种类: 通常怪兽
		 卡片密码: 41396436
		 使用限制: 无限制
		 种族: 鸟兽
		 属性: 风
		 星级: 4
		 攻击力: 1600
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: ME，VOL06
		 效果: 描述：青白色的火鸟，头上的毛有如皇冠。
		}
		*/
		Id:       2384,
		Password: "41396436",
		Name:     "苍翼冠鸟",                 // "Blue-Winged Crown"  "冠を戴く蒼き翼"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Wind,      // 风
		Lr:       ygo.LR_WindBeast, // 鸟兽
		Attack:   1600,
		Defense:  1200,
		IsValid:  true,
	})

	/*190*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2385
		 中文名: 暴雨怪
		 日文名: スコール
		 英文名: Violent Rain
		 卡片种类: 通常怪兽
		 卡片密码: 94042337
		 使用限制: 无限制
		 种族: 水
		 属性: 水
		 星级: 4
		 攻击力: 1550
		 防御力: 800
		 罕见度: 平卡N
		 卡包: VOL06
		 效果: 描述：突然降下倾盆大雨的怪兽。
		}
		*/
		Id:       2385,
		Password: "94042337",
		Name:     "暴雨怪",                  // "Violent Rain"  "スコール"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Water, // 水
		Lr:       ygo.LR_None,  // 水
		Attack:   1550,
		Defense:  800,
		IsValid:  true,
	})

	/*191*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2386
		 中文名: 电击蜗牛
		 日文名: ボルト·エスカルゴ
		 英文名: Bolt Escargot
		 卡片种类: 通常怪兽
		 卡片密码: 12146024
		 使用限制: 无限制
		 种族: 雷
		 属性: 水
		 星级: 5
		 攻击力: 1400
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: VOL06
		 效果: 描述：吐出粘液然后用电击对方。
		}
		*/
		Id:       2386,
		Password: "12146024",
		Name:     "电击蜗牛",                 // "Bolt Escargot"  "ボルト·エスカルゴ"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Water,   // 水
		Lr:       ygo.LR_Thunder, // 雷
		Attack:   1400,
		Defense:  1500,
		IsValid:  true,
	})

	/*192*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2387
		 中文名: 小鬼
		 日文名: インプ
		 英文名: Horn Imp
		 卡片种类: 通常怪兽
		 卡片密码: 69669405
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 1300
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: VOL06
		 效果: 描述：住在黑暗的小鬼。攻击意外的强。要注意它头上的角。
		}
		*/
		Id:       2387,
		Password: "69669405",
		Name:     "小鬼",                   // "Horn Imp"  "インプ"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Dark,  // 暗
		Lr:       ygo.LR_Fiend, // 恶魔
		Attack:   1300,
		Defense:  1000,
		IsValid:  true,
	})

	/*193*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2388
		 中文名: 杰米亚之神
		 日文名: ゼミアの神
		 英文名: Lord of Zemia
		 卡片种类: 通常怪兽
		 卡片密码: 81618817
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 4
		 攻击力: 1300
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: VOL06
		 效果: 描述：善于欺骗对方、诱惑其加入破灭之路的邪恶之神。
		}
		*/
		Id:       2388,
		Password: "81618817",
		Name:     "杰米亚之神",                // "Lord of Zemia"  "ゼミアの神"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Dark,  // 暗
		Lr:       ygo.LR_Fiend, // 恶魔
		Attack:   1300,
		Defense:  1000,
		IsValid:  true,
	})

	/*194*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2389
		 中文名: 狮鹫兽卫
		 日文名: グリフォール
		 英文名: Griffore
		 卡片种类: 通常怪兽
		 卡片密码: 53829412
		 使用限制: 无限制
		 种族: 兽
		 属性: 地
		 星级: 4
		 攻击力: 1200
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: VOL06，15AY
		 效果: 描述：善于用坚固的身体守备。将对方的攻击反弹。
		}
		*/
		Id:       2389,
		Password: "53829412",
		Name:     "狮鹫兽卫",                 // "Griffore"  "グリフォール"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Earth, // 地
		Lr:       ygo.LR_Beast, // 兽
		Attack:   1200,
		Defense:  1500,
		IsValid:  true,
	})

	/*195*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 239
		 中文名: 铠武者僵尸
		 日文名: 鎧武者ゾンビ
		 英文名: Armored Zombie
		 卡片种类: 通常怪兽
		 卡片密码: 20277860
		 使用限制: 无限制
		 种族: 不死
		 属性: 暗
		 星级: 3
		 攻击力: 1500
		 防御力: 0
		 罕见度: 平卡N
		 卡包: RB，BE01，VOL05，DL02
		 效果: 描述：充满怨念而苏醒的武者，要警惕在黑暗中挥舞的太刀。
		}
		*/
		Id:       239,
		Password: "20277860",
		Name:     "铠武者僵尸",                // "Armored Zombie"  "鎧武者ゾンビ"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Dark,   // 暗
		Lr:       ygo.LR_Zombie, // 不死
		Attack:   1500,
		Defense:  0,
		IsValid:  true,
	})

	/*196*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2408
		 中文名: 王家守卫
		 日文名: ロイヤルガード
		 英文名: Royal Guard
		 卡片种类: 通常怪兽
		 卡片密码: 39239728
		 使用限制: 无限制
		 种族: 机械
		 属性: 地
		 星级: 6
		 攻击力: 1900
		 防御力: 2200
		 罕见度: 平卡N
		 卡包: VOL07
		 效果: 描述：为了守护王座而开发的，且拥有意志的机械部队。
		}
		*/
		Id:       2408,
		Password: "39239728",
		Name:     "王家守卫",                 // "Royal Guard"  "ロイヤルガード"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    6,
		La:       ygo.LA_Earth,   // 地
		Lr:       ygo.LR_Machine, // 机械
		Attack:   1900,
		Defense:  2200,
		IsValid:  true,
	})

	/*197*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2409
		 中文名: 铁球巨人
		 日文名: ギガント
		 英文名: Giganto
		 卡片种类: 通常怪兽
		 卡片密码: 33621868
		 使用限制: 无限制
		 种族: 机械
		 属性: 暗
		 星级: 5
		 攻击力: 1700
		 防御力: 1800
		 罕见度: 平卡N
		 卡包: VOL07
		 效果: 描述：振舞着带刺的铁球进行攻击。
		}
		*/
		Id:       2409,
		Password: "33621868",
		Name:     "铁球巨人",                 // "Giganto"  "ギガント"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Dark,    // 暗
		Lr:       ygo.LR_Machine, // 机械
		Attack:   1700,
		Defense:  1800,
		IsValid:  true,
	})

	/*198*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2410
		 中文名: 钢铁之心
		 日文名: アイアン·ハート
		 英文名: Ancient Tool
		 卡片种类: 通常怪兽
		 卡片密码: 49587396
		 使用限制: 无限制
		 种族: 机械
		 属性: 暗
		 星级: 5
		 攻击力: 1700
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: VOL07
		 效果: 描述：在古代文明遗迹里发现的机器。破坏是它的唯一目的。
		}
		*/
		Id:       2410,
		Password: "49587396",
		Name:     "钢铁之心",                 // "Ancient Tool"  "アイアン·ハート"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Dark,    // 暗
		Lr:       ygo.LR_Machine, // 机械
		Attack:   1700,
		Defense:  1400,
		IsValid:  true,
	})

	/*199*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2411
		 中文名: 暗黑奇美拉
		 日文名: ダーク·キメラ
		 英文名: Dark Chimera
		 卡片种类: 通常怪兽
		 卡片密码: 32344688
		 使用限制: 无限制
		 种族: 恶魔
		 属性: 暗
		 星级: 5
		 攻击力: 1610
		 防御力: 1460
		 罕见度: 平卡N
		 卡包: VOL07
		 效果: 描述：生息在魔界的怪兽。吐出暗炎攻击。
		}
		*/
		Id:       2411,
		Password: "32344688",
		Name:     "暗黑奇美拉",                // "Dark Chimera"  "ダーク·キメラ"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Dark,  // 暗
		Lr:       ygo.LR_Fiend, // 恶魔
		Attack:   1610,
		Defense:  1460,
		IsValid:  true,
	})

	/*200*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 2412
		 中文名: 撞击兽
		 日文名: クラッシュマン
		 英文名: Togex
		 卡片种类: 通常怪兽
		 卡片密码: 33878931
		 使用限制: 无限制
		 种族: 兽
		 属性: 地
		 星级: 5
		 攻击力: 1600
		 防御力: 1800
		 罕见度: 平卡N
		 卡包: VOL07
		 效果: 描述：动作意外地敏捷，把身体卷成圆形进行撞击。
		}
		*/
		Id:       2412,
		Password: "33878931",
		Name:     "撞击兽",                  // "Togex"  "クラッシュマン"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Earth, // 地
		Lr:       ygo.LR_Beast, // 兽
		Attack:   1600,
		Defense:  1800,
		IsValid:  true,
	})

	/*201*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 246
		 中文名: 水之舞女
		 日文名: 水の踊り子
		 英文名: Water Omotics
		 卡片种类: 通常怪兽
		 卡片密码: 02483611
		 使用限制: 无限制
		 种族: 水
		 属性: 水
		 星级: 4
		 攻击力: 1400
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: RB，VOL04，DL02
		 效果: 描述：从瓶中不断流淌的水，变成龙的形状发动攻击。
		}
		*/
		Id:       246,
		Password: "02483611",
		Name:     "水之舞女",                 // "Water Omotics"  "水の踊り子"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Water, // 水
		Lr:       ygo.LR_None,  // 水
		Attack:   1400,
		Defense:  1200,
		IsValid:  true,
	})

	/*202*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 247
		 中文名: 陆战型战斗艇
		 日文名: 陸戦型 バグロス
		 英文名: Ground Attacker Bugroth
		 卡片种类: 通常怪兽
		 卡片密码: 58314394
		 使用限制: 无限制
		 种族: 机械
		 属性: 地
		 星级: 4
		 攻击力: 1500
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: RB，VOL04，DL02
		 效果: 描述：陆地战斗机器人，总有一天能在海里使用。
		}
		*/
		Id:       247,
		Password: "58314394",
		Name:     "陆战型战斗艇",               // "Ground Attacker Bugroth"  "陸戦型 バグロス"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Earth,   // 地
		Lr:       ygo.LR_Machine, // 机械
		Attack:   1500,
		Defense:  1000,
		IsValid:  true,
	})

	/*203*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 248
		 中文名: 幼虫宝宝
		 日文名: プチモス
		 英文名: Petit Moth
		 卡片种类: 通常怪兽
		 卡片密码: 58192742
		 使用限制: 无限制
		 种族: 昆虫
		 属性: 地
		 星级: 1
		 攻击力: 300
		 防御力: 200
		 罕见度: 平卡N
		 卡包: RB，BE01，VOL04，DL02
		 效果: 描述：不清楚成长后会变成什么样的小青虫。
		}
		*/
		Id:       248,
		Password: "58192742",
		Name:     "幼虫宝宝",                 // "Petit Moth"  "プチモス"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    1,
		La:       ygo.LA_Earth,  // 地
		Lr:       ygo.LR_Insect, // 昆虫
		Attack:   300,
		Defense:  200,
		IsValid:  true,
	})

	/*204*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 257
		 中文名: 狮子王
		 日文名: レオグン
		 英文名: Leogun
		 卡片种类: 通常怪兽
		 卡片密码: 10538007
		 使用限制: 无限制
		 种族: 兽
		 属性: 地
		 星级: 5
		 攻击力: 1750
		 防御力: 1550
		 罕见度: 平卡N
		 卡包: RB，VOL05，DL02
		 效果: 描述：浑身百兽之王般漂亮的鬃毛，身体也很庞大。
		}
		*/
		Id:       257,
		Password: "10538007",
		Name:     "狮子王",                  // "Leogun"  "レオグン"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Earth, // 地
		Lr:       ygo.LR_Beast, // 兽
		Attack:   1750,
		Defense:  1550,
		IsValid:  true,
	})

	/*205*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 261
		 中文名: 破坏的石人
		 日文名: 破壊のゴーレム
		 英文名: Destroyer Golem
		 卡片种类: 通常怪兽
		 卡片密码: 73481154
		 使用限制: 无限制
		 种族: 岩石
		 属性: 地
		 星级: 4
		 攻击力: 1500
		 防御力: 1000
		 罕见度: 平卡N
		 卡包: RB，EX-R(EX)，VOL04，DL02
		 效果: 描述：巨大右臂为特征的石人，用右臂进行压迫攻击。
		}
		*/
		Id:       261,
		Password: "73481154",
		Name:     "破坏的石人",                // "Destroyer Golem"  "破壊のゴーレム"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Earth, // 地
		Lr:       ygo.LR_Rock,  // 岩石
		Attack:   1500,
		Defense:  1000,
		IsValid:  true,
	})

	/*206*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 263
		 中文名: 苍白兽
		 日文名: ペイルビースト
		 英文名: Pale Beast
		 卡片种类: 通常怪兽
		 卡片密码: 21263083
		 使用限制: 无限制
		 种族: 兽
		 属性: 地
		 星级: 4
		 攻击力: 1500
		 防御力: 1200
		 罕见度: 平卡N
		 卡包: RB，EX-R(EX)，VOL05，DL02
		 效果: 描述：青白色的皮肤，让人感觉恶心的不明怪兽。
		}
		*/
		Id:       263,
		Password: "21263083",
		Name:     "苍白兽",                  // "Pale Beast"  "ペイルビースト"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Earth, // 地
		Lr:       ygo.LR_Beast, // 兽
		Attack:   1500,
		Defense:  1200,
		IsValid:  true,
	})

	/*207*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 265
		 中文名: 古代的蜥蜴战士
		 日文名: 古代のトカゲ戦士
		 英文名: Ancient Lizard Warrior
		 卡片种类: 通常怪兽
		 卡片密码: 43230671
		 使用限制: 无限制
		 种族: 爬虫类
		 属性: 地
		 星级: 4
		 攻击力: 1400
		 防御力: 1100
		 罕见度: 平卡N
		 卡包: RB，VOL04，DL02
		 效果: 描述：远古的保持着昔日姿态的蜥蜴战士，意外的强悍。
		}
		*/
		Id:       265,
		Password: "43230671",
		Name:     "古代的蜥蜴战士",              // "Ancient Lizard Warrior"  "古代のトカゲ戦士"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Earth,   // 地
		Lr:       ygo.LR_Reptile, // 爬虫类
		Attack:   1400,
		Defense:  1100,
		IsValid:  true,
	})

	/*208*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 267
		 中文名: 兵主部
		 日文名: ひょうすべ
		 英文名: Hyosube
		 卡片种类: 通常怪兽
		 卡片密码: 02118022
		 使用限制: 无限制
		 种族: 水
		 属性: 水
		 星级: 4
		 攻击力: 1500
		 防御力: 900
		 罕见度: 平卡N
		 卡包: RB，VOL05，DL02
		 效果: 描述：河童的首领，攻击力意外的高，守备力偏低。
		}
		*/
		Id:       267,
		Password: "02118022",
		Name:     "兵主部",                  // "Hyosube"  "ひょうすべ"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Water, // 水
		Lr:       ygo.LR_None,  // 水
		Attack:   1500,
		Defense:  900,
		IsValid:  true,
	})

	/*209*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 408
		 中文名: 宝贝龙
		 日文名: ベビードラゴン
		 英文名: Baby Dragon
		 卡片种类: 通常怪兽
		 卡片密码: 88819587
		 使用限制: 无限制
		 种族: 龙
		 属性: 风
		 星级: 3
		 攻击力: 1200
		 防御力: 700
		 罕见度: 平卡N
		 卡包: ME，BE02，VOL06，DL04，JY，SJ2
		 效果: 描述：不可轻视的幼龙，说不定隐藏着什么力量。
		}
		*/
		Id:       408,
		Password: "88819587",
		Name:     "宝贝龙",                  // "Baby Dragon"  "ベビードラゴン"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Wind,   // 风
		Lr:       ygo.LR_Dragon, // 龙
		Attack:   1200,
		Defense:  700,
		IsValid:  true,
	})

	/*210*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 409
		 中文名: 暗黑之龙王
		 日文名: 暗黒の竜王
		 英文名: Blackland Fire Dragon
		 卡片种类: 通常怪兽
		 卡片密码: 87564352
		 使用限制: 无限制
		 种族: 龙
		 属性: 暗
		 星级: 4
		 攻击力: 1500
		 防御力: 800
		 罕见度: 平卡N
		 卡包: ME，BE02，VOL06，DL04
		 效果: 描述：生活在黑暗深处的龙，由于长期出活在暗处。视力不太好。
		}
		*/
		Id:       409,
		Password: "87564352",
		Name:     "暗黑之龙王",                // "Blackland Fire Dragon"  "暗黒の竜王"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Dark,   // 暗
		Lr:       ygo.LR_Dragon, // 龙
		Attack:   1500,
		Defense:  800,
		IsValid:  true,
	})

	/*211*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 411
		 中文名: 牛魔人
		 日文名: 牛魔人
		 英文名: Battle Steer
		 卡片种类: 通常怪兽
		 卡片密码: 18246479
		 使用限制: 无限制
		 种族: 兽战士
		 属性: 地
		 星级: 5
		 攻击力: 1800
		 防御力: 1300
		 罕见度: 平卡N
		 卡包: ME，BE02，VOL06，DL04
		 效果: 描述：住在森林里的牛魔人，用突起的角进行突袭。
		}
		*/
		Id:       411,
		Password: "18246479",
		Name:     "牛魔人",                  // "Battle Steer"  "牛魔人"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Earth,       // 地
		Lr:       ygo.LR_BeastWarror, // 兽战士
		Attack:   1800,
		Defense:  1300,
		IsValid:  true,
	})

	/*212*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 413
		 中文名: 暗道化师 扎奇
		 日文名: 闇·道化師のサギー
		 英文名: Saggi the Dark Clown
		 卡片种类: 通常怪兽
		 卡片密码: 66602787
		 使用限制: 无限制
		 种族: 魔法师
		 属性: 暗
		 星级: 3
		 攻击力: 600
		 防御力: 1500
		 罕见度: 平卡N
		 卡包: ME，BE02，VOL07，DL04，DPKB，KA，SK2
		 效果: 描述：来去无踪的道化师，以不可思议的动作躲避攻击。
		}
		*/
		Id:       413,
		Password: "66602787",
		Name:     "暗道化师 扎奇",              // "Saggi the Dark Clown"  "闇·道化師のサギー"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Dark,        // 暗
		Lr:       ygo.LR_SpellCaster, // 魔法师
		Attack:   600,
		Defense:  1500,
		IsValid:  true,
	})

	/*213*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 415
		 中文名: 无脸幻想师
		 日文名: 幻想師·ノー·フェイス
		 英文名: Illusionist Faceless Mage
		 卡片种类: 通常怪兽
		 卡片密码: 28546905
		 使用限制: 无限制
		 种族: 魔法师
		 属性: 暗
		 星级: 5
		 攻击力: 1200
		 防御力: 2200
		 罕见度: 平卡N
		 卡包: ME，BE02，VOL07，DL04，PE
		 效果: 描述：制造幻影，能轻易地躲避敌人的攻击。
		}
		*/
		Id:       415,
		Password: "28546905",
		Name:     "无脸幻想师",                // "Illusionist Faceless Mage"  "幻想師·ノー·フェイス"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    5,
		La:       ygo.LA_Dark,        // 暗
		Lr:       ygo.LR_SpellCaster, // 魔法师
		Attack:   1200,
		Defense:  2200,
		IsValid:  true,
	})

	/*214*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 424
		 中文名: 人马兽
		 日文名: ケンタウロス
		 英文名: Mystic Horseman
		 卡片种类: 通常怪兽
		 卡片密码: 68516705
		 使用限制: 无限制
		 种族: 兽
		 属性: 地
		 星级: 4
		 攻击力: 1300
		 防御力: 1550
		 罕见度: 平卡N
		 卡包: ME，BE02，EX-R(EX)，VOL07，DL04
		 效果: 描述：人与马化身的怪兽，奔跑的速度谁也追不上。
		}
		*/
		Id:       424,
		Password: "68516705",
		Name:     "人马兽",                  // "Mystic Horseman"  "ケンタウロス"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Earth, // 地
		Lr:       ygo.LR_Beast, // 兽
		Attack:   1300,
		Defense:  1550,
		IsValid:  true,
	})

	/*215*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 436
		 中文名: TM-1火箭炮蜘蛛
		 日文名: TM－１ランチャースパイダー
		 英文名: Launcher Spider
		 卡片种类: 通常怪兽
		 卡片密码: 87322377
		 使用限制: 无限制
		 种族: 机械
		 属性: 炎
		 星级: 7
		 攻击力: 2200
		 防御力: 2500
		 罕见度: 平卡N，金字UR
		 卡包: ME，BE02，LE02，VOL07，DL04
		 效果: 描述：火箭发射器乱射，将敌人爆杀的机器蜘蛛。
		}
		*/
		Id:       436,
		Password: "87322377",
		Name:     "TM-1火箭炮蜘蛛",            // "Launcher Spider"  "TM－１ランチャースパイダー"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    7,
		La:       ygo.LA_None,    // 炎
		Lr:       ygo.LR_Machine, // 机械
		Attack:   2200,
		Defense:  2500,
		IsValid:  true,
	})

	/*216*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 437
		 中文名: 高科技狼
		 日文名: ギガテック·ウルフ
		 英文名: Giga-Tech Wolf
		 卡片种类: 通常怪兽
		 卡片密码: 08471389
		 使用限制: 无限制
		 种族: 机械
		 属性: 炎
		 星级: 4
		 攻击力: 1200
		 防御力: 1400
		 罕见度: 平卡N
		 卡包: ME，VOL06，DL04，TP11
		 效果: 描述：全身由钢铁制成的狼，用锋利的牙齿撕咬敌人。
		}
		*/
		Id:       437,
		Password: "08471389",
		Name:     "高科技狼",                 // "Giga-Tech Wolf"  "ギガテック·ウルフ"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_None,    // 炎
		Lr:       ygo.LR_Machine, // 机械
		Attack:   1200,
		Defense:  1400,
		IsValid:  true,
	})

	/*217*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 439
		 中文名: 彩虹鱼
		 日文名: レインボー·フィッシュ
		 英文名: 7 Colored Fish
		 卡片种类: 通常怪兽
		 卡片密码: 23771716
		 使用限制: 无限制
		 种族: 鱼
		 属性: 水
		 星级: 4
		 攻击力: 1800
		 防御力: 800
		 罕见度: 平卡N
		 卡包: ME，VOL07，DL04，SD04，GLD01
		 效果: 描述：非常珍奇的七色鱼，要捕捉它非常困难。
		}
		*/
		Id:       439,
		Password: "23771716",
		Name:     "彩虹鱼",                  // "7 Colored Fish"  "レインボー·フィッシュ"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    4,
		La:       ygo.LA_Water, // 水
		Lr:       ygo.LR_Fish,  // 鱼
		Attack:   1800,
		Defense:  800,
		IsValid:  true,
	})

	/*218*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 457
		 中文名: 高等女祭司
		 日文名: ハイ·プリーステス
		 英文名: Lady of Faith
		 卡片种类: 通常怪兽
		 卡片密码: 17358176
		 使用限制: 无限制
		 种族: 魔法师
		 属性: 光
		 星级: 3
		 攻击力: 1100
		 防御力: 800
		 罕见度: 平卡N
		 卡包: ME，VOL06，DL04，TP10
		 效果: 描述：咏唱起闻所未闻的咒语，使所有人的心平静下来。
		}
		*/
		Id:       457,
		Password: "17358176",
		Name:     "高等女祭司",                // "Lady of Faith"  "ハイ·プリーステス"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    3,
		La:       ygo.LA_Light,       // 光
		Lr:       ygo.LR_SpellCaster, // 魔法师
		Attack:   1100,
		Defense:  800,
		IsValid:  true,
	})

	/*219*/
	cardBag.Register(&ygo.CardOriginal{
		/*{
		 id: 735
		 中文名: 真红眼黑龙
		 日文名: 真紅眼の黒竜
		 英文名: Red-Eyes B. Dragon
		 卡片种类: 通常怪兽
		 卡片密码: 74677422
		 使用限制: 无限制
		 种族: 龙
		 属性: 暗
		 星级: 7
		 攻击力: 2400
		 防御力: 2000
		 罕见度: 金字UR，爆闪PR，立体UTR，银字R，平卡N，面闪SR
		 卡包: PG，BE01，PP05，VOL03，DL02，301，SD01，DT01，YAP01，DTP01，JY，SJ2，LC01，SD22
		 效果: 描述：拥有真红之眼的黑龙。愤怒的黑炎会把映入其眼者全部烧成灰烬。
		}
		*/
		Id:       735,
		Password: "74677422",
		Name:     "真红眼黑龙",                // "Red-Eyes B. Dragon"  "真紅眼の黒竜"
		Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
		Level:    7,
		La:       ygo.LA_Dark,   // 暗
		Lr:       ygo.LR_Dragon, // 龙
		Attack:   2400,
		Defense:  2000,
		IsValid:  true,
	})

}
