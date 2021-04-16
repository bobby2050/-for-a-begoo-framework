<!DOCTYPE html>

<html>
<head>
    <title>考试答题</title>
    <meta charset='utf-8'/>

    <meta name="viewport" content="width=device-width,initial-scale=1.0,user-scalable=0"/>
    <style type="text/css">
        *, body {
            margin: 0px;
            padding: 0px;
        }

        body {
            background: #f5f6fb;
            max-width: 100%;
            height: 1px;
            font-size: 62.5%;
            font-family: "Microsoft YaHei", Arial;
            overflow-x: hidden;
            overflow-y: auto;
        }

        #header {
            float: left;
            width: 100%;
            height: 239px;
            background: #3f8eea url("/static/img/top.png") bottom no-repeat;
            background-size: contain;
        }

        .top {
            padding-top: 25px;
            float: left;
        }

        #avatar {
            width: 70px;
            border-radius: 90px;
            border: #fff 5px solid;
        }

        .top1 {
            width: 20%;
            padding-left: 15px;
        }

        .top2 {
            width: 20%;

            padding-left: 10px;
        }

        .top3 {
            padding-top: 45px;

            float: right;
        }

        .top2-font {
            font-size: 20px;
            color: #fff;
        }

        .top3 img {
            height: 45px;
        }

        .top2-font-nickname {
            padding-top: 10px;
            font-size: 18px;
        }

        .top2-font-number {
            padding-top: 10px;
            font-size: 18px;
            font-weight: lighter;
        }

        #body-block {
            float: left;
            width: 100%;
            height: auto;
            margin-top: -40px;
        }

        .item-block {
            /*float: left;*/
            width: 94%;
            height: 150px;
            background: #cccccc url("/static/img/blockImg.png") center no-repeat;
            background-size: contain;
            margin: 0 auto;
            border-radius: 8px;
            margin-bottom: 20px;
        }

        .item-block-no {
            /*float: left;*/
            width: 94%;
            height: 150px;
            background: #eeeeee;
            background-size: contain;
            margin: 0 auto;
            border-radius: 8px;
            margin-bottom: 20px;
        }

        #item-block-division {
            width: 100%;
            height: 70px;
            margin: 0 auto;
            /*margin-top: 20px;*/
            /*margin-bottom: 10px;*/
            color: #cccccc;
        }

        .item-block-division-center {
            margin-top: 30px;
            margin: 0 auto;
            width: 80%;

        }

        .item-block-division-line {
            border-bottom: 1px solid #ccc;
            float: left;
            width: 32%;
            height: 35px;
            margin: 0 auto;
        }

        .item-block-division-words {
            float: left;
            line-height: 70px;
            margin-left: 10px;
            margin-right: 10px;
            font-size: 14px;
            color: #999;
            /*background: lightgreen;*/
        }

        .item-block-num, .item-block-title, .item-block-total {
            text-align: center;
        }

        .item-block-num {
            padding-top: 10px;
            font-size: 30px;
        }

        .item-block-title {
            font-size: 24px;
        }

        .item-block-total {
            padding-top: 20px;
            font-size: 15px;
            color: #666;
        }

        .item-block-color {
            color: #ec682a;
        }

        .item-block-title a {
            color: #000;
            text-decoration: none

        }

        .top2-font-number-a {
            padding-top: 10px;
            font-size: 18px;
            font-weight: lighter;
            text-decoration: none

        }
    </style>
</head>

<body>
<div id="header">
    <div class="top top1"><img id="avatar" src="{{.AvatarUrl}}"/></div>
    <div class="top top2">
        <div class="top2-font top2-font-nickname"> {{.NickName}}</div>
        <div class="top2-font top2-font-number" id="mobile-num"></div>
    </div>
    <div class="top3"><img id="msQuestion" src="/static/img/ms.png"/></div>
</div>

<div id="body-block">
    <div id="add" lastQuestionId=""></div>
</div>


<script src="/static/js/jquery.min.js"></script>
<script type="application/javascript">
    //获取url中的参数
    function getUrlParam(name) {
        var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
        var r = window.location.search.substr(1).match(reg);  //匹配目标参数
        if (r != null) return unescape(r[2]);
        return null; //返回参数值
    }

    $(function () {
        // var ServerUrl = "https://cw.bestjan.com/"
        var ServerUrl = "http://127.0.0.1:8080/"
        // 全局 1:可用 2:不可用
        var isUse = 1

        $.ajax({
            type: "GET",
            url: ServerUrl + "getData?userId=" + getUrlParam("userId") + "&category=" + getUrlParam("category"),
            //请求成功
            success: function (re) {
                //console.log(re)
                $("#add").attr("lastQuestionId", re.LastQuestionId)
                if (re.MobileNum == "") {
                    isUse = 2
                    var url = ServerUrl + "register?userId=" + getUrlParam("userId") + "&avatarUrl=" + getUrlParam("avatarUrl") + "&nickName=" + escape(getUrlParam("nickName")) + "&category=" + getUrlParam("category");
                    console.log(url)

                    $("#mobile-num").html("<a class='top2-font top2-font-number-a' href='"+url+"' >点击注册</a>")
                } else {

                    $("#mobile-num").html(re.MobileNum)
                }

                var s = "";
                var pp = "";
                pp += '<div id="item-block-division">'
                pp += '<div class="item-block-division-center">'
                pp += '<div class="item-block-division-line"></div> <div class="item-block-division-words">暂无考试权限</div> <div class="item-block-division-line"></div>'
                pp += '</div>'
                pp += '</div>'

                var q = ""; // 前
                var e = ""; // 后
                var ti = 0
                for (var p in re.ReCategoryTao) {
                    console.log(ti)

                    if (ti == 0) {

                        q += '<div class="item-block">'
                        q += '<div class="item-block-num">第 <span class="item-block-color">' + re.ReCategoryTao[p].Ranking + '</span> 套</div>';
                        q += '<div class="item-block-title"><a href="getQuestion?categroy=' + re.ReCategoryTao[p].Id + '&userId=' + getUrlParam('userId') + '&per=1&lastQuestionId=' + re.LastQuestionId + '&type=1&category2=' + getUrlParam("category")+'">' + re.ReCategoryTao[p].Name  + '</a></div>';
                        q += '<div class="item-block-total">共' + re.ReCategoryTao[p].TotalQuestion + '题</div>';
                        q += '</div>';
                    } else {
                        if (re.IsAuth == 1){
                            e += '<div class="item-block">'
                            e += '<div class="item-block-num">第 <span class="item-block-color">' + re.ReCategoryTao[p].Ranking + '</span> 套</div>';
                            e += '<div class="item-block-title"><a href="getQuestion?categroy=' + re.ReCategoryTao[p].Id + '&userId=' + getUrlParam('userId') + '&per=1&lastQuestionId=' + re.LastQuestionId + '&type=1&category2=' + getUrlParam("category")+'">' + re.ReCategoryTao[p].Name + '</a></div>';
                            e += '<div class="item-block-total">共' + re.ReCategoryTao[p].TotalQuestion + '题</div>';
                            e += '</div>';
                        } else {
                            e += '<div class="item-block-no">'
                            e += '<div class="item-block-num">第 <span class="item-block-color">' + re.ReCategoryTao[p].Ranking + '</span> 套</div>';
                            e += '<div class="item-block-title">' + re.ReCategoryTao[p].Name + '</div>';
                            e += '<div class="item-block-total">共' + re.ReCategoryTao[p].TotalQuestion + '题</div>';
                            e += '</div>';
                        }
                    }
                    ti++
                }
                var end = "";
                if (e != "") {
                    if (re.IsAuth == 2) {
                        end = pp + e
                    } else {
                        end = e
                    }
                }
                s = q + end
                //console.log(s)
                $("#add").after(s)
            }
        });

        $('#msQuestion').click(function () {
            var type = 2
            var url = ServerUrl + 'getQuestion?userId=' + getUrlParam("userId") + '&type=' + type + "&category=" + getUrlParam("category")
            window.location.href = url
        })

    })

</script>
</body>
</html>
