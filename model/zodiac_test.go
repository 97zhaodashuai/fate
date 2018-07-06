package model_test

import (
	"testing"

	"github.com/godcong/fate/model"
)

var zodiacList = map[model.ZodiacType][]string{}

func TestNewZodiac(t *testing.T) {
	//z.Best = "冖,口,广,门,宀,山,爫,亠,木,艹,月,钅"
	//z.Create()
}

//属鼠：子午冲，忌讳包含午马的汉字；子未相害，忌讳包含未羊的汉字；卯刑子，忌讳包含卯兔的汉字。
//　　属牛：丑未冲，忌讳包含未羊的汉字；丑午相害，忌讳包含午马的汉字；未刑丑，忌讳包含未羊的汉字。
//　　属虎：寅申冲，忌讳包含申猴的汉字；寅巳相害，忌讳包含巳蛇的汉字；申刑寅，忌讳包含申猴的汉字。
//　　属兔：卯酉冲，忌讳包含酉鸡的汉字；卯辰相害，忌讳包含辰龙的汉字；子刑卯，忌讳包含子鼠的汉字。
//　　属龙：辰戌冲，忌讳包含戌狗的汉字；卯辰相害，忌讳包含卯兔的汉字；辰自刑，忌讳包含辰龙的汉字。
//　　属蛇：巳亥冲，忌讳包含亥猪的汉字；寅巳相害，忌讳包含寅虎的汉字；寅刑巳，忌讳包含寅虎的汉字。
//　　属马：子午冲，忌讳包含子鼠的汉字；丑午相害，忌讳包含丑牛的汉字；午自刑，忌讳包含午马的汉字。
//　　属羊：丑未冲，忌讳包含丑牛的汉字；子未相害，忌讳包含子鼠的汉字；戌刑未，忌讳包含戌狗的汉字。
//　　属猴：寅申冲，忌讳包含寅虎的汉字；申亥相害，忌讳包含亥猪的汉字；巳刑申，忌讳包含巳蛇的汉字。
//　　属鸡：卯酉冲，忌讳包含卯兔的汉字；酉戌相害，忌讳包含戌狗的汉字；酉自刑，忌讳包含酉鸡的汉字。
//　　属狗：辰戌冲，忌讳包含辰龙的汉字；酉戌相害，忌讳包含酉鸡的汉字；丑刑戌，忌讳包含丑牛的汉字。
//　　属猪：巳亥冲，忌讳包含巳蛇的汉字；申亥相害，忌讳包含申猴的汉字；亥自刑，忌讳包含亥猪的汉字。
//　　在起名时，一般考虑到以上的忌讳就够了。
//
//子鼠
//肖鼠的人起名時，鄭博士認為要考慮屬鼠者的喜用與忌用的部首字：
//
//喜用：牛辰艸水木禾田金玉等部首字
//忌用：午馬犭絲旁土火 等部首字。
//1根據六合三合相生之理
//喜用：
//
//申字根：紳坤伸
//辰字根：震振農濃
//亥字根：該家豪豫
//丑字根：紐鈕牟生
//2不合生肖：午(馬)未(羊)
//忌用：
//
//午字根：許馮馬駱
//未字根：妹美洋翔
//3鼠五行屬木
//喜用字如下：
//
//洞穴：口冖宀冂
//住家：門戶廣攵
//柵欄：冊聿豐豐
//草地：艸⑹鼠是夜行動物，在日間無法發揮所長。

//丑牛
//肖牛年生人，起名宜有“氵”字，清爽享福，上下敦睦；
//
//有“亻”字“木”字，義利分明，操守廉正；有“月”字，孤勞不順；
//有“火”字，不利健康或忌車怕水；有“田”字“車”字“馬”字，勞苦一生；
//有“石”字“山”字，易孤獨，不利家庭，晚婚遲得子大吉；
//有“血”字“糸”字“刀”字“力”字“幾”字，多不順，忌車怕水。
//肖牛年生人，到歲次肖羊年肖龍年肖馬年要小心注意，到歲次肖鼠年蛇年一帆風順，成功隆昌。鄭博士建議根據以下特征改名起名：
//
//1根據六合三合相生之理
//喜用：
//
//子字根：孟季學存淳
//巳字根：巴融進建張凱
//酉字根：鄭鴻翁斐茜鈞
//可用：亥字根：該核
//
//2不合生肖：馬(午)未(羊)戍(狗)
//忌用：
//
//午字根：許馬駿騏駱驊馮
//未字根：朱妹茉美祥群
//戍字根：成國獻然狄猛
//3牛五行屬土
//喜用字如下：
//
//洞穴：口冖宀
//平地：平原田甫
//牛在水里是一種享受，因此可取個水字旁的字。
//
//寅虎
//肖虎年生人，起名宜有“山”字，雄霸山林，智勇雙全，福壽興家；
//
//有“玉”字，英俊才人，多才巧智；
//有“金”字“木”字“衣”字“氵”字，溫和賢淑，名利雙收，環境良好；
//有“月”字“犭”字“馬”字，義利分明，操守廉正，克己助人；
//有“日”字“火”字，性剛果斷，幼年不順或憂心勞神；
//有“田”字“口”字“兒”字，不利家庭，晚婚遲得子大吉；
//有“糸”字“石”字“刀”字“力”字“血”字“弓”字“父”字“足”字，多不順，忌車怕水或不利健康。
//肖虎年生人，到歲次肖猴年肖豬年肖蛇年要小心注意，到歲次肖馬年肖狗年一帆風順，成功隆昌。鄭博士建議根據以下特征改名起名：
//
//1根據六合三合相生之理
//喜用：
//
//午字根：許杵駿騏騰
//戍字根：成威然獻猛
//卯字根：柳昴菟勉
//2不合生肖：巳(蛇)申(猴)
//忌用：
//
//巳字根：選虹進建強凡
//申字根：申袁侯蔡
//3虎五行屬木
//喜用字如下：
//
//山林：山林
//洞穴：宀冖
//忌用字如下：
//
//平地：平原田甫谷
//柵欄：冊聿豐豐
//
//卯兔
//肖兔年生人，起名宜有“月”字，清秀多才，溫和廉正，安富尊榮；
//
//有“亻”字“禾”字“木”字，貴人明現，精誠公正；
//有“入”字“宀”字，重義信用，環境良好；
//有“金”字“白”字“玉”字“豆”字，勤儉勵業，成功隆昌，富貴增榮；
//有“犭”字，良善積德，子孫興旺；有“馬”字“酉”字，多不順，不利健康；
//有“石”字“力”字“刀”字，不利家庭，晚婚遲得子大吉；
//有“皮”字“氵”字大兇；兔怕水，有“川”字更兇，忌車怕水。
//肖兔年生人，到歲次肖雞年肖馬年肖龍年要小心注意，到歲次肖羊年肖豬年肖狗年一帆風順，成功隆昌。   鄭博士建議根據以下特征改名起名：
//
//1根據六合三合相生之理
//喜用：
//
//亥字根：核孩
//未字根：朱妹善美
//寅字根：演虔虛
//2不合生肖：酉(雞)辰(龍)
//忌用：
//
//酉字根：鴻
//辰字根：振農龔龐云
//3兔五行屬木
//喜用：水木火
//忌用：金土
//4兔是非肉食動物，吃的是五谷類草類食物
//喜用：豆米禾麥梁艸
//忌用：忄(肉)月心
//5兔子自然的生長環境是在戶外
//喜用字如下：
//
//洞穴：口冖宀
//居地：田甫山丘谷
//
//辰龍
//肖龍年生人，起名宜有“氵”字，大吉，有沖天之勢，成功隆昌，富貴增榮，一生享福祿；
//
//有“金”字“玉”字“白”字“赤”字，精明公正，學識淵博，福壽興家；
//有“月”字，溫和賢淑，克己助人，良善積德，子孫鼎盛；
//有“魚”字“酉”字“亻”字，勤儉建業，家聲克振，貴人明現；
//有“土”字“田”字“禾”字“衣”字，多不順，不利家庭，晚婚遲得子大吉；
//有“土”字“忄”字“日”字，性剛果斷或憂心勞神；
//有“石”字“艸”字清雅平凡，貴人明現，易因情誤事；
//有“力”字“刀”字，不利家庭；
//有“糸”字“犭”字，奔波勞苦；
//有“火”字，無自立之地，忌車怕水，不利健康。
//肖龍年生人，到歲次肖狗年肖牛年肖兔年要小心注意，到歲次肖鼠年肖猴年肖雞年一帆風順，成功隆昌。鄭博士建議根據以下特征改名起名：
//
//1根據六合三合相生之理相合生肖：子(鼠)申(猴)酉(雞)
//喜用：
//
//子字根：孫李學孟淳郭
//申字根：紳侯神
//酉字根：鄭鴻茜
//2不合生肖－－虎(寅)兔(卯)戍(狗)未(羊)
//忌用：
//
//寅字根：彪演
//卯字根：昴柳卿茆逸勉
//戍字根：成獻然狄猛
//未字根：茉妹味美善羚
//3龍五行屬土
//
//喜用：水火
//
//忌用：土金⑹龍是神化動物，不食人間煙火。
//忌用：豆米禾麥梁田甫忄(肉)月心
//4龍宜有「水」字，沖天之勢。困在洞穴等地，無法發揮
//喜用： 氵水雨云天
//
//忌用字如下：
//
//洞穴：口宀冖冂
//住家：門戶廣攵
//柵欄：冊聿豐豐
//5龍為帝王象征，若使用蛇小龍暗示文字則有降位的意念
//喜用：君帝皇主長王大天京首顫
//忌用：士丞相公臣少小卒工力民巳蟲毛它辶廴弓幾
//肖蛇年生人，起名就宜
//
//有“艸”字，大吉，一生享福祿，富貴增榮；
//有“蟲”字“魚”字，智勇雙全，精誠溫和；
//有“木”字“禾”字“田”字“山”字，重義信用，學識淵博，成功隆昌，名利永在；
//有“金”字“玉”字，多才巧智，克己助人，良善積德；
//有“月”字“土”字，操守廉正，一門鼎盛；有“忄”字，性剛或憂心勞神；
//有“石”字“刀”字“血”字“弓”字，不利家庭，晚婚遲得子大吉，忌車怕水；
//有“火”字“亻”字“糸”字，不利健康。
//
//巳蛇
//肖蛇年生人，到次肖豬年肖猴年肖虎年要小心注意，到歲次肖牛年肖雞年一帆風順，成功隆昌。鄭博士建議根據以下特征改名起名：
//
//1根據六合三合相生之理相合生肖：丑(牛)午(馬)未(羊)酉(雞)
//喜用：
//
//酉字根：鄭鴻翔
//丑字根：鈕特
//可用：
//
//午字根：許駿駱騏
//未字根：朱美善羚
//2不合生肖：寅虎亥豬
//忌用：
//
//寅字根：演彪
//亥字根：核該駭
//3蛇五行屬火
//喜用：金木土火
//忌用：水
//4蛇是肉食動物，不吃五谷類。
//喜用：忄(肉)月心
//忌用：豆米禾麥梁
//5蛇應有藏身之處，例如：在草叢中活動，不然很容易被捕食或被發現。
//喜用字如下：
//
//洞穴：口宀冖
//住家：門戶廣攵
//柵欄：冊聿豐豐
//平地：艸平原田甫
//忌用字如下：山艮丘屯
//
//6蛇如果用龍字，小龍變成大龍。如果用蛇形文字沒有好處。
//喜用：辰字根：辰龍鹿言京貝云雨君民
//忌用：辶廴弓幾之
//
//午馬
//肖馬年生人，起名宜有“艸”字“金”字，學識淵博，安尊榮，享福終世；
//
//有“玉”字“木”字“禾”字，貴人明現，多才巧智，成功隆昌；
//
//有“蟲”字“豆”字“米”字，福祿雙收，名利永在；
//
//有“亻”字“月”字，英俊才人，智勇雙全，清雅榮貴；
//
//有“土”字，義利分明，溫和觀淑，克己助人，重義信用；
//
//有“田”字“火”字“氵”字，憂心勞神或性剛；
//
//有“車”字“石”字“力”字“酉”字“馬”字，不利家庭，婚遲得子大吉，或不利健康。
//
//肖馬年生人，到處次肖鼠年肖兔年肖牛年要小心注意，到歲次肖虎年肖狗年，肖羊年一帆風順隆昌。鄭博士建議根據以下特征來起名改名：
//
//1根據六合三合相生之理相合生肖：寅(虎)已(蛇)未(羊)戍(狗)
//喜用：
//
//寅字根：虔彪演
//戍字根：成威然獻狄
//可用：
//
//巳字根：融陀尾
//未字根：朱姝祥群
//2不合生肖：子(鼠)丑(牛)酉(雞)
//忌用：
//
//子字根：孫李孺郭游
//丑字根：紐特生
//酉字根：鴻茜
//3馬五行屬火
//喜用：木火土
//忌用：金水
//4馬是非肉食動物，吃的是五谷類草類食物。
//喜用：豆米禾麥梁田甫
//忌用：忄(肉)月心
//5馬在草原賓士，有所發揮所長。馬若下田耕種，必屬劣馬。
//喜用字如下：
//
//草原：艸若芷蘋
//忌用：田甫
//
//未羊
//肖羊年生人，起名宜
//
//有“金”字“白”字“玉”字“艸”字，學識淵博，操守廉正，重義信用，富貴增榮；
//有“月”字“田”字“豆”字“米”字，勤儉建業，名利雙收，安享清福；
//有“馬”字“禾”字“木”字“亻”字“魚”字，英俊才人，多才巧智，溫和賢淑，克己助人；
//有“忄”字“犭”字“糸”字，憂心勞神或不利家庭；
//有“車”字“氵”字“山”字“日”字“火”字，不利家庭或健康，忌車怕水。
//肖羊年生人，到歲次肖牛羊，肖狗年，肖鼠年要小心注意，到歲次肖馬年肖兔年肖豬年一帆風順，成功隆昌。鄭博士建議根據以下特征改名起名：
//
//1根據六合三合相生之理相合生肖：卯(兔)亥(豬)已(蛇)午(馬)
//喜用：
//
//卯字根：卿菟
//亥字根：該核
//可用：
//
//巳字根：過逸迪
//午字根：許駿
//2不合生肖：丑(牛)辰(龍)子(鼠)
//忌用：
//
//丑字根：紐鈕
//辰字根：龐振龔
//子字根：孫李學存孝游
//3羊五行屬土
//喜用：木火土
//忌用：金水
//4羊是非肉食動物，吃的是五谷類草類食物。
//喜用：艸豆米禾麥梁竹
//忌用：忄(肉)月心
//5羊在的大草原與平地上生存，同時也能攀登高山，但也常有羊在山里最后前無出路后無退路的窘境。
//喜用字如下：
//
//洞穴：口宀冖
//平地：艸原田甫谷
//忌用字如下：山地：山丘屯艮
//
//申猴
//肖猴年生人，起名宜
//
//有“木”字“禾”字，清貴享福，成功發達；
//有“金”字“玉”字“豆”字“米”字，英俊佳人，多才賢淑，福祿雙收；
//有“田”字“山”字“月”字，操守廉正，名利雙收，一門鼎盛；有“氵”字“亻”字，風流樂天，上下敦睦，智勇雙全；
//有“火”字“石”字，性剛果斷或不利家庭；有“口”字“人”字“冖”字，忌車怕水，不利家庭；
//有“糸”字“刀”字“力”字“皮”字“犭”字，多不順，不利健康；有“山”字，安富尊榮，福壽興家。
//肖猴年生人，到歲次肖虎年肖蛇年肖豬年要小心注意，肖鼠年肖龍年一帆風順，成功隆昌。鄭博士建議根據以下特征改名起名：
//
//1根據六合三合相生之理相合生肖：子(鼠)辰(龍)酉(雞)戍(狗)
//喜用：
//
//子字根：學孺郭孝李
//辰字根：振宸震龐瓏
//可用：
//
//酉字根：酒鴻
//戍字根：成誠默猛
//2不合生肖：寅(虎)亥(豬)
//忌用：
//
//寅字根；演彪
//亥字根：核該
//3猴五行屬金
//喜用：土水
//忌用：火木
//4猴子吃的是水果五谷類或草類植物，不吃肉類。
//
//喜用：豆米禾麥梁
//
//忌用：忄(肉)月心
//
//5猴子棲息森林或山上。
//
//喜用字如下：
//
//山地：山丘屯艮
//洞穴：宀冖
//住家：門戶廣攵
//
//酉雞
//肖雞年生人，起名宜
//
//有“米”字“豆”字“蟲”字，福壽興家，富貴清吉；
//有“木”字“禾”字“玉”字“田”字，福祿雙收，名利永在；
//有“月”字“人”字“冖”字，棲宿安閑，多才巧智，環境良好；
//有“山”字“艸”字“日”字“金”字，智勇雙全，清雅榮貴；
//有“石”字“犭”字“刀”字“力”字“日”字“酉”字“血”字“弓”字“糸”字“車”字“馬”字等，幼年不順或性剛果斷，不利健康或忌車怕水。
//肖雞年生人，到歲次肖兔年肖鼠年肖狗年要小心注意，到歲次肖蛇年肖牛年肖龍年一帆風順，成功隆昌。鄭博士建議根據以下特征改名起名：
//
//1根據六合三合相生之理相合生肖：巳(蛇)丑(牛)申(猴)
//喜用：
//
//巳字根：蟲它毛之弓幾廴辶
//丑字根：紐鈕牟特
//申字根：坤紳
//2不合生肖：卯(兔)戍(狗)
//忌用：
//
//卯字根：昴柳勉逸育朋柔林
//戍字根：成然慶狄
//3雞五行屬金
//喜用：土
//可用：金水
//忌用：火木
//4雞主要吃谷類食物，不吃肉
//喜用：豆米禾麥梁
//忌用：忄(肉)月心,雞主要在田中棲息或被人類飼養。
//喜用字如下：
//
//洞穴：宀冖口
//住家：門戶廣攵
//柵欄：冊聿豐豐曲
//居地：田甫山丘屯谷
//
//戌狗
//狗是最忠于人的動物，所以肖狗之人喜有“亻”字“人”“入”之字形，意味有其飼主，并終于主人，終于事業，終于錢財。
//
//狗喜披彩衣，有虎風之味，增加其威勢，提升地味之感。如字型是“系”“巾 ”“衣”狗喜有“冖” “宀”“人”字型，意味家庭內的狗，比較好命，有主人有房子住。
//
//肖狗之人喜三合之字根，狗為戌，“寅午戌”為三合。鄭博士舉例：寅字型如虎處虛獻。“午”字型如：瑪竹篤駿駐騏騫驊。三合的力量對人的幫扶大，人緣貴人運都好。
//
//肖狗之人喜有“心”“忄”“月”之字型，因狗喜食肉。
//
//肖狗之人還喜有“小”“少”“士”“臣”之字型，一般而言，小狗比大狗可愛，狗不為君帝將帥寧為辰士。
//
//肖狗之人不宜見到“木”“辰” “酉” “兆”“鳥” “羽”“兌”“西”“金”“未”“羊” “丑”“牛” “水 ”“氵” “子”“北”“亥”“呂”“雨”“云”“佳”“田” “日”“禾”“米”“豆”“梁”等字。
//
//肖狗年生人，起名宜有：
//
//“豆“字“米”字“魚”字，福祿雙收，名利永在，富貴清潔；
//“氵”字“金”字“玉”字，智勇雙全，精明公正，克己助人，溫和賢淑；
//“月”字“木”字“禾”字，子孫興旺，環境良好；
//“亻”字“山”字“土”字“艸”字，英俊才人，重義信用；
//“糸”字“石”字“刀”字“力”字“血”字“弓”字“兒”字“皮”字“父”字等，不利健康或忌車怕水，不利家庭。
//
//亥豬
//肖豬年生人，到歲次肖蛇年肖虎年肖猴年要小心注意，不過鄭博士認為到歲次肖兔年，肖羊年一帆順，成功隆昌。鄭博士建議根據以下特征改名起名：
//
//1相合生肖：卯(兔)未(羊)子(鼠)丑(牛)
//喜用：
//
//卯字根：昴柳卿勉
//未字根：味妹茉
//可用：
//
//子字根：李孫存厚敦
//丑字根：鈕生特
//2不合生肖：巳(蛇)申(猴)
//忌用：
//
//巳字根：蟲它之弓幾陀芝弟凡
//申字根：坤暢侯
//3豬五行屬水
//喜用：金木水
//忌用：火土
//4豬不吃肉，只吃五谷類
//喜用：豆米禾麥梁
//忌用：忄(肉)月心
//5豬的自然棲息地
//喜用：洞穴：口冖宀
//山地：山艮恒屯
//平地：原田甫谷
//鄭偉建博士點評：
//屬相與名字的關系是共性的問題，因此，也是僅供朋友們參考而已，如果生搬硬套恐怕就形而上學了。
//一個真正屬于自己的名字，一定是符合自己命理喜忌的范疇，這樣的名字才能風生水起好運來。
//如果你的名字不盡人意或者你有了可愛的小寶寶，那么，鄭博士建議你一定要認真對待起名問題，因為取一個好名字關系到人的一生。
