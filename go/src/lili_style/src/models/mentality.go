package models

//1. 思った通りにならなかったとき、しばらく考え込む	                     01
//2. 自分のいい面、得意なことを聞かれたら10個言える	                     10
//3. トラブルが起きた時、まず人に頼る	                                 01
//4. 人と比較することが多い	                                         01
//5. 昔から「マイペースだよね」とよく言われる	                         10
//6. 過去を振り返ったときに、結果を出してきたことがない	                 01
//7. 何かにつけて「無理」「できない」などの言葉が出がち	                 01
//8. 電車やエスカレーターの乗り降りなどでノロノロした人にイライラする	     01
//9. 人の顔色が気になる	                                         01
//10. チームで活動することが楽しい、好き	                             10
//11.朝起きて鏡を見たとき、嫌なところに目が行く	                         01
//12.アイデア力、発想力は高いと思う	                                 10
//13.職場や学校で注意や指摘を受けると落ち込む	                         01
//14.やるぞ！と決めても本当にこれでいいのか考えてしまい、なかなか行動できない 01
//15.自分のペースを乱されるとイラっとしてしまう	                         01
//16.友達は多い方だ	                                             10
//17.出かける前の洋服選びにかなり時間をかける	                         01
//18.新しいことにチャレンジしたいが、自分には難しいと感じる	             01
//19.人から言われた何気ない一言が頭から消えないことがある	                 01
//20.落ち込んだときの自分のモチベーションのあげ方は分かっている	         10
//21.周りとの人間関係は上手くいっている方だ	                         01
//22.直接「ありがとう」と言われない仕事は嫌だ	                         01
//26~47


type Mentality struct {
	One 	  string //n
	Two 	  string //y
	Three 	  string //n
	Four      string //n
	Five      string //y
	Six       string //n
	Seven     string //n
	Eight     string //n
	Nine      string //n
	Ten       string //y
	Eleven    string //n
	Twelve    string //y
	Thirteen  string //n
	Fourteen  string //n
	Fifteen   string //n
	Sixteen   string //y
	Seventeen string //n
	Eighteen  string //n
	Nineteen  string //n
	Twenty    string //y
	TwentyOne string //n
	TwentyTwo string //n
	Rate      int
}

