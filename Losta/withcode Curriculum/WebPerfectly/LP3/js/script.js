'use strict'
//スムーススクロール
$('#pc-nav a,#g-nav a').click(function () {
	let hrefElement = $(this).attr('href');
	let headerHight = $("#header").outerHeight(true);
	let position = Math.round($(hrefElement).offset().top - headerHight);
	$('body,html').animate({scrollTop: position}, 500);
	return false;
});


$('.openbtn').click(function () {
	$(this).toggleClass('active');
	$('#pc-nav').toggleClass('panelactive');
});

$('#pc-nav a').click(function () {
	$('.openbtn').removeClass('active');
	$('#pc-nav').removeClass('panelactive');
})

//スクロールした際の動きを関数でまとめる
function PageTopAnime() {
		var scroll = $(window).scrollTop();
		if (scroll >= 200){//上から200pxスクロールしたら
			$('#page-top').removeClass('RightMove');//#page-topについているRightMoveというクラス名を除く
			$('#page-top').addClass('LeftMove');//#page-topについているLeftMoveというクラス名を付与
		}else{
			if(
				$('#page-top').hasClass('LeftMove')){//すでに#page-topにLeftMoveというクラス名がついていたら
				$('#page-top').removeClass('LeftMove');//LeftMoveというクラス名を除き
				$('#page-top').addClass('RightMove');//RightMoveというクラス名を#page-topに付与
			}
		}
}

// #page-topをクリックした際の設定
$('#page-top').click(function () {
    $('body,html').animate({
        scrollTop: 0//ページトップまでスクロール
    }, 500);//ページトップスクロールの速さ。数字が大きいほど遅くなる
    return false;//リンク自体の無効化
});



/*===========================================================*/
/*9-2-2	任意の場所をクリックすると隠れていた内容が開き、先に開いていた内容が閉じる*/
/*===========================================================*/
//アコーディオンをクリックした時の動作
$('.title').on('click', function() {//タイトル要素をクリックしたら
	$('.box').slideUp(500);//クラス名.boxがついたすべてのアコーディオンを閉じる
    
	var findElm = $(this).next(".box");//タイトル直後のアコーディオンを行うエリアを取得
    
	if($(this).hasClass('close')){//タイトル要素にクラス名closeがあれば
		$(this).removeClass('close');//クラス名を除去    
	}else{//それ以外は
		$('.close').removeClass('close'); //クラス名closeを全て除去した後
		$(this).addClass('close');//クリックしたタイトルにクラス名closeを付与し
		$(findElm).slideDown(500);//アコーディオンを開く
	}
});

function fadeAnime(){
$('.fadeUpTrigger').each(function(){
		var elemPos = $(this).offset().top;
		var scroll = $(window).scrollTop();
		var windowHeight = $(window).height();
		if (scroll >= elemPos - windowHeight){
		$(this).addClass('fadeUp');
		}else{
		$(this).removeClass('fadeUp');
		}
		});
	
	// ぱたっ
$('.flipLeftTrigger').each(function(){ //flipLeftTriggerというクラス名が
		var elemPos = $(this).offset().top;//要素より、50px上の
		var scroll = $(window).scrollTop();
		var windowHeight = $(window).height();
		if (scroll >= elemPos - windowHeight){
		$(this).addClass('flipLeft');// 画面内に入ったらflipLeftというクラス名を追記
		}else{
		$(this).removeClass('flipLeft');// 画面外に出たらflipLeftというクラス名を外す
		}
		});
    
}


/*===========================================================*/
/* 関数をまとめる*/
/*===========================================================*/


// 画面をスクロールをしたら動かしたい場合の記述
$(window).scroll(function () {
    fadeAnime();//スクロールに連動した動きの関数を呼ぶ
});

// ページが読み込まれたらすぐに動かしたい場合の記述
$(window).on('load',function(){

    $("#splash-logo").delay(1200).fadeOut('slow');//ロゴを1.2秒でフェードアウトする記述

     //=====ここからローディングエリア（splashエリア）を1.5秒でフェードアウトした後に動かしたいJSをまとめる  
	$("#splash").delay(1500).fadeOut('slow',function(){//4-2-1 背景色が伸びる（下から上）が動作した後に下記アニメーションを実行
        $('body').addClass('appear');//4-2-1 背景色が伸びる（下から上）
        PageTopAnime();// 8-1-3	ページの指定の高さを超えたら右から出現する関数を呼ぶ

    /* 9-2-2	任意の場所をクリックすると隠れていた内容が開き、先に開いていた内容が閉じる*/
	$(".open").each(function(Element){	//openクラスを取得
		var Title =$(Element).children('.title');	//openクラスの子要素のtitleクラスを取得
		$(Title).addClass('close');				///タイトルにクラス名closeを付与し
		var Box =$(Element).children('.box');	//openクラスの子要素boxクラスを取得
		$(Box).slideDown(500);					//アコーディオンを開く
	});
        
    });
    //=====ここまでローディングエリア（splashエリア）を1.5秒でフェードアウトした後に動かしたいJSをまとめる
    
    /*===========================================================*/
    /* 4-2-1 背景色が伸びる（下から上） */
    /*===========================================================*/

    //=====ここから背景が伸びた後に動かしたいJSをまとめる   
    $('.splashbg').on('animationend', function() {        
        fadeAnime();//スクロールに連動した動きの関数を呼ぶ
    });  
    //=====ここまで背景が伸びた後に動かしたいJSをまとめる
	
});// ここまでページが読み込まれたらすぐに動かしたい場合の記述
