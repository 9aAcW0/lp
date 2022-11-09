package models

//1. 大切な人に何ができるかを考える	                                    10
//2. 自分の周りには嫌な人が多い	                                        01
//3. 今まで人に裏切られたことがある	                                    01
//4. うまくいかないことがあったときは自分のどの行動が悪かったか考える	        10
//5. 自分は周りの人よりも、上だと思う	                                    01
//6. 困っているときに助けてくれた人には恩返ししたい	                        10
//7. パートナーや恋人を選ぶときは「自分がこの人をどれだけ幸せにできるか」を考える	10
//8. 誰かに応援してもらえている	                                        10
//9. 自分に自信がない	                                                01
//10.メリット・デメリットを常に考えている	                                01
//11.誰かの役に立ちたい	                                            10
//12.自分は幸せだと思う	                                            10
//47~58
type AltruismRate int

type Altruism struct {
	One 	  string //y
	Two 	  string //n
	Three 	  string //n
	Four      string //y
	Five      string //n
	Six       string //y
	Seven     string //y
	Eight     string //y
	Nine      string //n
	Ten       string //n
	Eleven    string //y
	Twelve    string //y
	Rate      int
}

