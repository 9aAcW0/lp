package models

//1. 目標達成できない時は、とても悔しい	      10
//2. 環境のせいで、達成できないことが多い	      01
//3. 難題が出てきた時、とっさに「できない」と思う 01
//4. 達成したら、すぐに次の目標を作りたい	      10
//5. 目標達成するために、とにかく誰よりも行動する 10
//6. 「できない」ことは断るべきだ	          01
//7. 諦めたくなったら、諦めればいい	          01
//8. 環境のせいで、達成できないことが多い	      01
//9. 報告しづらいことは隠せば問題ない	          01
//10.自分なりの結果を出せば良い	              01
//59~68

type Commit struct {
	One 	  string //y
	Two 	  string //n
	Three 	  string //n
	Four      string //y
	Five      string //y
	Six       string //n
	Seven     string //n
	Eight     string //n
	Nine      string //n
	Ten       string //n
	Rate 	  int
}

