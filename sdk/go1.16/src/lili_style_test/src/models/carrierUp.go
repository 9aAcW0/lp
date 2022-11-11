package models

//1. 残業したくない                         01
//2. ノルマがある仕事は嫌だ                   01
//3. 結婚や出産後は専業主婦やパートで働きたい    01
//4. 全国転勤のある会社は選びたくない           01
//5. 圧倒的に成長したい                      10
//6. 自立できる人になりたい                   10
//7. ヘッドハンティングされるような人になりたい   10
//8. 個の力で活躍できる自分になりたい           10
//9. 将来は起業したい                        10
//10.結果を出すなら１番をとりたい              10
//11.世の中の社会課題をビジネスで解決していきたい 10
//12.実力主義の会社で働きたい                 10
//13.出世・昇格したい                        10
//14.管理職は目指していない                   01
//15.将来、いくらぐらいの年収を目指したいですか？ 1~5
//11~25


type CarrierUp struct {
	One 	  string //n
	Two 	  string //n
	Three 	  string //n
	Four      string //n
	Five      string //y
	Six       string //y
	Seven     string //y
	Eight     string //y
	Nine      string //y
	Ten       string //y
	Eleven    string //y
	Twelve    string //y
	Thirteen  string //y
	Fourteen  string //n
	Fifteen   string //1~5
	Rate      int
}
