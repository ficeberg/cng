package regular

import (
	"errors"
	"fmt"
	"strings"

	"github.com/festum/cng"
	"github.com/xormsharp/xorm"
)

var regularList = map[int]string{
	1:  `一、乙`,
	2:  `二、十、丁、廠、七、卜、人、入、八、九、幾、兒、了、力、乃、刀、又`,
	3:  `三、於、幹、虧、士、工、土、才、寸、下、大、丈、與、萬、上、小、口、巾、山、千、乞、川、億、個、勺、久、凡、及、夕、丸、麼、廣、亡、門、義、之、屍、弓、己、已、子、衛、也、女、飛、刃、習、叉、馬、鄉`,
	4:  `豐、王、井、開、夫、天、無、元、專、雲、扎、藝、木、五、支、廳、不、太、犬、區、歷、尤、友、匹、車、巨、牙、屯、比、互、切、瓦、止、少、日、中、岡、貝、內、水、見、午、牛、手、毛、氣、升、長、仁、什、片、僕、化、仇、幣、仍、僅、斤、爪、反、介、父、從、今、凶、分、乏、公、倉、月、氏、勿、欠、風、丹、勻、烏、鳳、勾、文、六、方、火、為、鬥、憶、訂、計、戶、認、心、尺、引、醜、巴、孔、隊、辦、以、允、予、勸、雙、書、幻`,
	5:  `玉、刊、示、末、未、擊、打、巧、正、撲、扒、功、扔、去、甘、世、古、節、本、術、可、丙、左、厲、右、石、布、龍、平、滅、軋、東、卡、北、佔、業、舊、帥、歸、且、旦、目、葉、甲、申、叮、電、號、田、由、史、只、央、兄、叼、叫、另、叨、嘆、四、生、失、禾、丘、付、仗、代、仙、們、儀、白、仔、他、斥、瓜、乎、叢、令、用、甩、印、樂、句、匆、冊、犯、外、處、冬、鳥、務、包、飢、主、市、立、閃、蘭、半、汁、匯、頭、漢、寧、穴、它、討、寫、讓、禮、訓、必、議、訊、記、永、司、尼、民、出、遼、奶、奴、加、召、皮、邊、發、孕、聖、對、臺、矛、糾、母、幼、絲`,
	6:  `式、刑、動、扛、寺、吉、扣、考、託、老、執、鞏、圾、擴、掃、地、揚、場、耳、共、芒、亞、芝、朽、樸、機、權、過、臣、再、協、西、壓、厭、在、有、百、存、而、頁、匠、誇、奪、灰、達、列、死、成、夾、軌、邪、劃、邁、畢、至、此、貞、師、塵、尖、劣、光、當、早、吐、嚇、蟲、曲、團、同、吊、吃、因、吸、嗎、嶼、帆、歲、回、豈、剛、則、肉、網、年、朱、先、丟、舌、竹、遷、喬、偉、傳、乒、乓、休、伍、伏、優、伐、延、件、任、傷、價、份、華、仰、仿、夥、偽、自、血、向、似、後、行、舟、全、會、殺、合、兆、企、眾、爺、傘、創、肌、朵、雜、危、旬、旨、負、各、名、多、爭、色、壯、衝、冰、莊、慶、亦、劉、齊、交、次、衣、產、決、充、妄、閉、問、闖、羊、並、關、米、燈、州、汗、汙、江、池、湯、忙、興、宇、守、宅、字、安、講、軍、許、論、農、諷、設、訪、尋、那、迅、盡、導、異、孫、陣、陽、收、階、陰、防、奸、如、婦、好、她、媽、戲、羽、觀、歡、買、紅、纖、級、約、紀、馳、巡`,
	7:  `壽、弄、麥、形、進、戒、吞、遠、違、運、扶、撫、壇、技、壞、擾、拒、找、批、扯、址、走、抄、壩、貢、攻、赤、折、抓、扮、搶、孝、均、拋、投、墳、抗、坑、坊、抖、護、殼、志、扭、塊、聲、把、報、卻、劫、芽、花、芹、芬、蒼、芳、嚴、蘆、勞、克、蘇、杆、槓、杜、材、村、杏、極、李、楊、求、更、束、豆、兩、麗、醫、辰、勵、否、還、殲、來、連、步、堅、旱、盯、呈、時、吳、助、縣、裡、呆、園、曠、圍、呀、噸、足、郵、男、困、吵、串、員、聽、吩、吹、嗚、吧、吼、別、崗、帳、財、針、釘、告、我、亂、利、禿、秀、私、每、兵、估、體、何、但、伸、作、伯、伶、傭、低、你、住、位、伴、身、皁、佛、近、徹、役、返、餘、希、坐、谷、妥、含、鄰、岔、肝、肚、腸、龜、免、狂、猶、角、刪、條、卵、島、迎、飯、飲、系、言、凍、狀、畝、況、床、庫、療、應、冷、這、序、辛、棄、冶、忘、閒、間、悶、判、灶、燦、弟、汪、沙、汽、沃、泛、溝、沒、沈、沉、懷、憂、快、完、宋、宏、牢、究、窮、災、良、證、啟、評、補、初、社、識、訴、診、詞、譯、君、靈、即、層、尿、尾、遲、局、改、張、忌、際、陸、阿、陳、阻、附、妙、妖、妨、努、忍、勁、雞、驅、純、紗、納、綱、駁、縱、紛、紙、紋、紡、驢、紐`,
	8:  `奉、玩、環、武、青、責、現、表、規、抹、攏、拔、揀、擔、坦、押、抽、拐、拖、拍、者、頂、拆、擁、抵、拘、勢、抱、垃、拉、攔、拌、幸、招、坡、披、撥、擇、擡、其、取、苦、若、茂、蘋、苗、英、範、直、茄、莖、茅、林、枝、杯、櫃、析、板、鬆、槍、構、傑、述、枕、喪、或、畫、臥、事、刺、棗、雨、賣、礦、碼、廁、奔、奇、奮、態、歐、壟、妻、轟、頃、轉、斬、輪、軟、到、非、叔、肯、齒、些、虎、虜、腎、賢、尚、旺、具、果、味、昆、國、昌、暢、明、易、昂、典、固、忠、咐、呼、鳴、詠、呢、岸、巖、帖、羅、幟、嶺、凱、敗、販、購、圖、釣、制、知、垂、牧、物、乖、刮、稈、和、季、委、佳、侍、供、使、例、版、侄、偵、側、憑、僑、佩、貨、依、的、迫、質、欣、徵、往、爬、彼、徑、所、舍、金、命、斧、爸、採、受、乳、貪、念、貧、膚、肺、肢、腫、脹、朋、股、肥、服、脅、周、昏、魚、兔、狐、忽、狗、備、飾、飽、飼、變、京、享、店、夜、廟、府、底、劑、郊、廢、淨、盲、放、刻、育、閘、鬧、鄭、券、卷、單、炒、炊、炕、炎、爐、沫、淺、法、洩、河、沾、淚、油、泊、沿、泡、注、瀉、泳、泥、沸、波、潑、澤、治、怖、性、怕、憐、怪、學、寶、宗、定、宜、審、宙、官、空、簾、實、試、郎、詩、肩、房、誠、襯、衫、視、話、誕、詢、該、詳、建、肅、錄、隸、居、屆、刷、屈、弦、承、孟、孤、陝、降、限、妹、姑、姐、姓、始、駕、參、艱、線、練、組、細、駛、織、終、駐、駝、紹、經、貫`,
	9:  `奏、春、幫、珍、玻、毒、型、掛、封、持、項、垮、挎、城、撓、政、赴、趙、擋、挺、括、拴、拾、挑、指、墊、掙、擠、拼、挖、按、揮、挪、某、甚、革、薦、巷、帶、草、繭、茶、荒、茫、蕩、榮、故、胡、南、藥、標、枯、柄、棟、相、查、柏、柳、柱、柿、欄、樹、要、鹹、威、歪、研、磚、釐、厚、砌、砍、面、耐、耍、牽、殘、殃、輕、鴉、皆、背、戰、點、臨、覽、豎、省、削、嘗、是、盼、眨、哄、顯、啞、冒、映、星、昨、畏、趴、胃、貴、界、虹、蝦、蟻、思、螞、雖、品、咽、罵、譁、咱、響、哈、咬、咳、哪、炭、峽、罰、賤、貼、骨、鈔、鍾、鋼、鑰、鉤、卸、缸、拜、看、矩、怎、牲、選、適、秒、香、種、秋、科、重、復、竿、段、便、倆、貸、順、修、保、促、侮、儉、俗、俘、信、皇、泉、鬼、侵、追、俊、盾、待、律、很、須、敘、劍、逃、食、盆、膽、勝、胞、胖、脈、勉、狹、獅、獨、狡、獄、狠、貿、怨、急、饒、蝕、餃、餅、彎、將、獎、哀、亭、亮、度、跡、庭、瘡、瘋、疫、疤、姿、親、音、帝、施、聞、閥、閣、差、養、美、姜、叛、送、類、迷、前、首、逆、總、煉、炸、炮、爛、剃、潔、洪、灑、澆、濁、洞、測、洗、活、派、洽、染、濟、洋、洲、渾、濃、津、恆、恢、恰、惱、恨、舉、覺、宣、室、宮、憲、突、穿、竊、客、冠、語、扁、襖、祖、神、祝、誤、誘、說、誦、墾、退、既、屋、晝、費、陡、眉、孩、除、險、院、娃、姥、姨、姻、嬌、怒、架、賀、盈、勇、怠、柔、壘、綁、絨、結、繞、驕、繪、給、絡、駱、絕、絞、統`,
	10: `耕、耗、豔、泰、珠、班、素、蠶、頑、盞、匪、撈、栽、捕、振、載、趕、起、鹽、捎、捏、埋、捉、捆、捐、損、都、哲、逝、撿、換、挽、熱、恐、壺、挨、恥、耽、恭、蓮、莫、荷、獲、晉、惡、真、框、桂、檔、桐、株、橋、桃、格、校、核、樣、根、索、哥、速、逗、慄、配、翅、辱、脣、夏、礎、破、原、套、逐、烈、殊、顧、轎、較、頓、斃、致、柴、桌、慮、監、緊、黨、晒、眠、曉、鴨、晃、晌、暈、蚊、哨、哭、恩、喚、啊、唉、罷、峰、圓、賊、賄、錢、鉗、鑽、鐵、鈴、鉛、缺、氧、特、犧、造、乘、敵、秤、租、積、秧、秩、稱、祕、透、筆、笑、筍、債、借、值、倚、傾、倒、倘、俱、倡、候、俯、倍、倦、健、臭、射、躬、息、徒、徐、艦、艙、般、航、途、拿、爹、愛、頌、翁、脆、脂、胸、胳、髒、膠、腦、狸、狼、逢、留、皺、餓、戀、槳、漿、衰、高、席、準、座、脊、症、病、疾、疼、疲、效、離、唐、資、涼、站、剖、競、部、旁、旅、畜、閱、羞、瓶、拳、粉、料、益、兼、烤、烘、煩、燒、燭、煙、遞、濤、浙、澇、酒、涉、消、浩、海、塗、浴、浮、流、潤、浪、浸、漲、燙、湧、悟、悄、悔、悅、害、寬、家、宵、宴、賓、窄、容、宰、案、請、朗、諸、讀、扇、襪、袖、袍、被、祥、課、誰、調、冤、諒、談、誼、剝、懇、展、劇、屑、弱、陵、陶、陷、陪、娛、娘、通、能、難、預、桑、絹、繡、驗、繼`,
	11: `球、理、捧、堵、描、域、掩、捷、排、掉、堆、推、掀、授、教、掏、掠、培、接、控、探、據、掘、職、基、著、勒、黃、萌、蘿、菌、菜、萄、菊、萍、菠、營、械、夢、梢、梅、檢、梳、梯、桶、救、副、票、戚、爽、聾、襲、盛、雪、輔、輛、虛、雀、堂、常、匙、晨、睜、眯、眼、懸、野、啦、晚、啄、距、躍、略、蛇、累、唱、患、唯、崖、嶄、崇、圈、銅、鏟、銀、甜、梨、犁、移、笨、籠、笛、符、第、敏、做、袋、悠、償、偶、偷、您、售、停、偏、假、得、銜、盤、船、斜、盒、鴿、悉、欲、彩、領、腳、脖、臉、脫、象、夠、猜、豬、獵、貓、猛、餡、館、湊、減、毫、麻、癢、痕、廊、康、庸、鹿、盜、章、竟、商、族、旋、望、率、著、蓋、粘、粗、粒、斷、剪、獸、清、添、淋、淹、渠、漸、混、漁、淘、液、淡、深、婆、樑、滲、情、惜、慚、悼、懼、惕、驚、慘、慣、寇、寄、宿、窯、密、謀、謊、禍、謎、逮、敢、屠、彈、隨、蛋、隆、隱、婚、嬸、頸、績、緒、續、騎、繩、維、綿、綢、綠`,
	12: `琴、斑、替、款、堪、搭、塔、越、趁、趨、超、提、堤、博、揭、喜、插、揪、搜、煮、援、裁、擱、摟、攪、握、揉、斯、期、欺、聯、散、惹、葬、葛、董、葡、敬、蔥、落、朝、辜、葵、棒、棋、植、森、椅、椒、棵、棍、棉、棚、棕、惠、惑、逼、廚、廈、硬、確、雁、殖、裂、雄、暫、雅、輩、悲、紫、輝、敞、賞、掌、晴、暑、最、量、噴、晶、喇、遇、喊、景、踐、跌、跑、遺、蛙、蛛、蜓、喝、喂、喘、喉、幅、帽、賭、賠、黑、鑄、鋪、鏈、銷、鎖、鋤、鍋、鏽、鋒、銳、短、智、毯、鵝、剩、稍、程、稀、稅、筐、等、築、策、篩、筒、答、筋、箏、傲、傅、牌、堡、集、焦、傍、儲、奧、街、懲、御、循、艇、舒、番、釋、禽、臘、脾、腔、魯、猾、猴、然、饞、裝、蠻、就、痛、童、闊、善、羨、普、糞、尊、道、曾、焰、港、湖、渣、溼、溫、渴、滑、灣、渡、遊、滋、溉、憤、慌、惰、愧、愉、慨、割、寒、富、竄、窩、窗、遍、裕、褲、裙、謝、謠、謙、屬、屢、強、粥、疏、隔、隙、絮、嫂、登、緞、緩、編、騙、緣`,
	13: `瑞、魂、肆、攝、摸、填、搏、塌、鼓、擺、攜、搬、搖、搞、塘、攤、蒜、勤、鵲、藍、墓、幕、蓬、蓄、蒙、蒸、獻、禁、楚、想、槐、榆、樓、概、賴、酬、感、礙、碑、碎、碰、碗、碌、雷、零、霧、雹、輸、督、齡、鑑、睛、睡、睬、鄙、愚、暖、盟、歇、暗、照、跨、跳、跪、路、跟、遣、蛾、蜂、嗓、置、罪、罩、錯、錫、鑼、錘、錦、鍵、鋸、矮、辭、稠、愁、籌、籤、簡、毀、舅、鼠、催、傻、像、躲、微、愈、遙、腰、腥、腹、騰、腿、觸、解、醬、痰、廉、新、韻、意、糧、數、煎、塑、慈、煤、煌、滿、漠、源、濾、濫、滔、溪、溜、滾、濱、粱、灘、慎、譽、塞、謹、福、群、殿、闢、障、嫌、嫁、疊、縫、纏`,
	14: `靜、碧、璃、牆、撇、嘉、摧、截、誓、境、摘、摔、聚、蔽、慕、暮、蔑、模、榴、榜、榨、歌、遭、酷、釀、酸、磁、願、需、弊、裳、顆、嗽、蜻、蠟、蠅、蜘、賺、鍬、鍛、舞、穩、算、籮、管、僚、鼻、魄、貌、膜、膊、膀、鮮、疑、饅、裹、敲、豪、膏、遮、腐、瘦、辣、竭、端、旗、精、歉、熄、熔、漆、漂、漫、滴、演、漏、慢、寨、賽、察、蜜、譜、嫩、翠、熊、凳、騾、縮`,
	15: `慧、撕、撒、趣、趟、撐、播、撞、撤、增、聰、鞋、蕉、蔬、橫、槽、櫻、橡、飄、醋、醉、震、黴、瞞、題、暴、瞎、影、踢、踏、踩、蹤、蝶、蝴、囑、墨、鎮、靠、稻、黎、稿、稼、箱、箭、篇、僵、躺、僻、德、艘、膝、膛、熟、摩、顏、毅、糊、遵、潛、潮、懂、額、慰、劈`,
	16: `操、燕、薯、薪、薄、顛、橘、整、融、醒、餐、嘴、蹄、器、贈、默、鏡、贊、籃、邀、衡、膨、雕、磨、凝、辨、辯、糖、糕、燃、澡、激、懶、壁、避、繳`,
	17: `戴、擦、鞠、藏、霜、霞、瞧、蹈、螺、穗、繁、辮、贏、糟、糠、燥、臂、翼、驟`,
	18: `鞭、覆、蹦、鐮、翻、鷹`,
	19: `警、攀、蹲、顫、瓣、爆、疆`,
	20: `壤、耀、躁、嚼、嚷、籍、魔、灌`,
	21: `蠢、霸、露`,
	22: `囊`,
	23: `罐`,
}

type regular struct {
	db      cng.Database
	total   int
	fixed   int
	unfixed int
}

// Run ...
func (r *regular) Run() {
	e := setAllRegular(r.db, false)
	if e != nil {
		panic(e)
	}
	for _, strVal := range regularList {
		strVal = strings.TrimSpace(strVal)
		strVal = strings.Trim(strVal, `
`)
		strVal = strings.Trim(strVal, "\n")
		list := strings.Split(strVal, "、")
		if len(list) == 0 {
			continue
		}
		for _, ch := range list {
			r.total++
			fmt.Printf("character %s is fixing\n", ch)
			if fixRegular(r.db, ch) {
				r.fixed++
			} else {
				r.unfixed++
			}
		}
	}
	fmt.Printf("fix regular finished(total:%d,fixed:%d,unfixed:%d)\n", r.total, r.fixed, r.unfixed)
}

// Regular ...
type Regular interface {
	Run()
}

// New ...
func New(database cng.Database) Regular {
	return &regular{
		db: database,
	}
}

func setAllRegular(db cng.Database, regular bool) (e error) {
	engine, b := db.Database().(*xorm.Engine)
	if !b {
		return errors.New("wrong type")
	}
	if !regular {
		_, e = engine.Exec("UPDATE `character` set regular = 0")
	} else {
		_, e = engine.Exec("UPDATE `character` set regular = 1")
	}
	if e != nil {
		return e
	}
	return nil
}

func fixRegular(db cng.Database, ch string) bool {
	char, err := db.GetCharacter(cng.Char(ch))
	if err != nil {
		fmt.Printf("failed get char(%s) with error (%v)\n", ch, err)
		return false
	}
	char.Regular = true
	engine, b := db.Database().(*xorm.Engine)
	if !b {
		fmt.Println("failed to engine")
		return false
	}
	update, e := engine.Where("hash = ?", char.Hash).UseBool("regular").Update(char)
	if e != nil {
		fmt.Printf("failed update char(%s) with error (%v)\n", ch, e)
		return false
	}
	if update == 0 {
		return false
	}
	return true
}
