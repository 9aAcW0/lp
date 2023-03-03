package models

//1. 決められた目標は達成する	                         10
//2. 目標を達成するためのノウハウを習得し工夫するようにしている 10
//3. 1年後のキャリア・ライフプランが描けている	             10
//4. 10年後のキャリア・ライフプランが描けている	         10
//5. チームで活動するとき、リーダーになることが多い	         10
//6. いつも新しいチャレンジをしている	                     10
//7. 授業や講座で学んだことは図解で記録する	             10
//8. モノゴトを多面的にみる傾向が強い	                     10
//9. 目の前で起こっていることの背景や原因を考える	         10
//10.経験したことのない分野にも挑戦している	             10
//11.常に新たなことを学ぶのが好きだ	                     10
//12.失敗したことは忘れて次のチャレンジをする	             01
//13.目標達成するために、まずは戦略をたてる	             10
//68~80

type BusinessStance struct {
	One 	  string //y
	Two 	  string //y
	Three 	  string //y
	Four      string //y
	Five      string //y
	Six       string //y
	Seven     string //y
	Eight     string //y
	Nine      string //y
	Ten       string //y
	Eleven    string //y
	Twelve    string //n
	Thirteen  string //y
	Rate      int
}

