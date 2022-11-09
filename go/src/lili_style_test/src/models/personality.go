package models

//1. やると決めたことを最後までやりきる
//2. 無責任な人を見るとイライラする
//3. 「〇〇すべき」という言葉をよく使っている
//4. 厳しいしつけを受けてきた

//5. わたしはホメ上手だと思う
//6. 聞き役になることが多い
//7. 困っている人をみるとなんとかしてあげたくなる
//8. ボランティア活動などに参加するのが好き

//9. 結果を予測して準備する
//10.他の人はどうするだろう？と客観視する
//11.物事を分析して、事実に基づいて考える
//12.上手くいかない時でもあまりイライラしない

//13.欲しいものは手に入れないと気が済まない
//14.人のことを見ただけで、なんとなくこの人はこんな人かなと感じる
//15.将来のイメージを妄想することが好き
//16.みんなでワイワイ集まるのが好き

//17.嫌なことを嫌と言えないことが多い
//18.リーダー的な人の意見に従う方が楽
//19.人からの視線や評価がいつも気になる
//20.何か頼まれると先延ばしにしがち

//はい        ->2
//どちらでもない->1
//いいえ　　　　->0

type Personality struct {
	One 	  		  string
	Two 	  		  string
	Three 	  		  string
	Four      		  string
	Five      		  string
	Six       		  string
	Seven     		  string
	Eight     		  string
	Nine      		  string
	Ten       		  string
	Eleven    		  string
	Twelve    		  string
	Thirteen  		  string
	Fourteen  		  string
	Fifteen   		  string
	Sixteen   		  string
	Seventeen 		  string
	Eighteen  		  string
	Nineteen  		  string
	Twenty    		  string
	RigorRate 		  int
	AcceptabilityRate int
	LogicalityRate 	  int
	FreedomRate 	  int
	CoordinationRate  int
}
