<!DOCTYPE html>

<html>
<head>
    <title>考试试题</title>
    <meta charset='utf-8'/>

    <meta name="viewport" content="width=device-width,initial-scale=1.0,user-scalable=0"/>
    <style type="text/css">
        *, body {
            margin: 0px;
            padding: 0px;
        }

        body {
            background: #f3f6fb;
            max-width: 100%;
            height: 1px;
            font-size: 62.5%;
            font-family: "Microsoft YaHei", Arial;
            overflow-x: hidden;
            overflow-y: auto;
        }

        .header {
            /*text-align: center;*/
            background: #3f8de9;
            width: 100%;
            height: 180px;
            float: left;
        }

        #header-title {
            color: #fff;
            text-align: center;
            padding-top: 60px;
            font-size: 30px;
        }

        #header-sec-title {
            margin-left: 30px;
            margin-top: 20px;
            color: #82c6ed;
            /*font-size: 26px;*/
        }

        .content {
            height: auto;
            background: #fff;

            width: 95%;
            margin-top: 10px;

            padding: 20px 11px;
            float: left;
        }

        .question {

            color: #000;
            font-weight: bold;
            margin-bottom: 40px;
            font-size: 16px;
        }

        .optionList {
            width: 100%;
            height: 40px;
            margin-top: 15px;
            /*padding-bottom: 10px;*/
            border-bottom: 1px solid #ccc;
        }

        .optionWords {
            width: 90%;
            float: left;
            height: 20px;
            line-height: 20px;
            color: #666;
            font-size: 14px;

        }

        .optionBtn {
            height: 30px;
            width: 10%;
            float: right;

        }

        .btnImg {
            height: 30px;
            width: 30px;
            float: right;
            line-height: 30px;
        }

        .btnBlock {

            width: 100%;
            height: 80px;
            display: block;
        }

        .btn {
            width: 200px;
            height: 40px;
            color: #ec6627;
            font-size: 20px;
            margin: 0 auto;
            padding-top: 13px;
        }

        .btnWords {

            height: 40px;
            border-radius: 30px;
            border: 1px solid #ccc;
            padding-left: 55px;
            padding-top: 10px;
            font-size: 20px;
        }

        .answer-title {
            font-size: 18px;
            color: #ec6627;
        }

        .answer-content {
            margin-top: 10px;
            color: #999;
            line-height: 20px;
            font-size: 14px;
        }

        .next {
            float: left;
            width: 100%;
            height: 125px;
        }

        .next-box {
            height: 80px;
            /*background: #3f8de9;*/
        }

        .box {

            height: 80px;
            line-height: 80px;
            margin-top: 60px;
            margin: 20px 20px;
            background: #e8662c;
            border-radius: 80px;
            color: #fff;
            text-align: center;
            font-size: 40px;
        }

        .ad-block, .ad-block-copyright {
            width: 100%;
            height: auto;
            text-align: center;
            float: left;
            clear: both;
        }

        .ad-block img {
            width: 100%;
        }

        .answer {
            display: none;
        }

        #fullBox {
            position: absolute;
            top: 0;
            right: 0;
            bottom: 0;
            left: 0;

            width: 100%;
            height: 100%;
            background: #666666;
        }

        #fullBoxPanl {
            z-index: 11;
            width: 80%;
            height: 200px;
            position: fixed;
            top: 0;
            right: 0;
            left: 0;
            bottom: 0;
            margin: auto;
            background: #ffffff;
            border-radius: 15px;
        }

        #fullBoxPanlh1 h1 {
            font-size: 24px;
            text-align: center;
            margin-top: 30px;
        }

        #fullBoxPanlContent {
            text-align: center;
            margin-top: 20px;
            margin-bottom: 30px;
            text-indent: 20px;
            color: #666666;
            font-size: 16px;
        }

        #fullBoxPanlBlockBtn {
            border-top: #bbb 1px solid;
        }

        .fullBoxBoxPanlBtn {
            line-height: 64px;
            font-size: 24px;
            /*margin-top: 20px;*/
            width: 49%;
            height: 64px;
            float: left;
            text-align: center;
        }
    </style>
</head>

<body>

<div class="header">
    <h1 id="header-title">共计错题<span id="Total"></span>道 </h1>

    <input type="hidden" id="userId" value="{{.Data.UserId}}"/>
    <input type="hidden" id="categoryTaoId" value=""/>
</div>


<input type="hidden" id="Total" value="{{.Data.Total}}"/>
<div class="next">
    <div class="next-box">
        <div class="box" isClick="1" lastQuestionId="1" tmpCurrentPer="1">
            下一页
        </div>
    </div>
</div>


<div class="ad-block">
    <img src="/static/img/ad-bottom.png"/>
</div>

<div class="ad-block-copyright">
    <img src="/static/img/ad-sc.png"/>
</div>


<div id="fullBox" style="display: none">
    <div id="fullBoxPanl">
        <div id="fullBoxPanlh1"><h1>继续答题</h1></div>
        <div id="fullBoxPanlContent">您上次答题到第<span id="fullBoxPanlContentSpan"></span>题，是否继续？</div>
        <div id="fullBoxPanlBlockBtn">
            <div id="startFirstQuestion" class="fullBoxBoxPanlBtn">取消</div>
            <div id="startGoOnQuestion" style="border-left:1px #aaa solid" class="fullBoxBoxPanlBtn">继续</div>
        </div>
    </div>
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

        var userIdT = getUrlParam("userId")
        var userId = $("#userId").val(userIdT)

        $.ajax({
            type: "GET",
            url: ServerUrl + "getQuestionFenye?category=" + getUrlParam('category') + "&userId=" + getUrlParam("userId") + "&per=1&lastQuestionId=0&type=2",
            //请求成功
            success: function (result) {
                console.log("1" + result);
                $("#Total").html(result.Total)
                $("#categoryTaoId").val(result.CategoryTaoId)


                if (result.ReQuestion && result.ReQuestion.length > 0) {

                    var str = "";
                    for (var p in result.ReQuestion) {

                        str += '<div class="content" questionId="' + result.ReQuestion[p].Id + '" optionType="' + result.ReQuestion[p].OptionType + '" tmpclick="">'
                        str += '<h3 class="question"> ' + result.ReQuestion[p].Num + '.' + result.ReQuestion[p].Title + '</h3>'


                        var op = '';

                        for (var o in result.ReQuestion[p].Answer) {
                            if (result.ReQuestion[p].Answer[o].OrderNum == 0) {
                                op += '<div class="optionList"><span class="optionWords">A.  ' + result.ReQuestion[p].Answer[o].OptionName + '</span>' +
                                    '<div class="optionBtn"><img option="' + result.ReQuestion[p].Answer[o].OptionAnswer + '" tmpoption="" class="btnImg" src="/static/img/default.png"></div>' +
                                    '   </div>'
                            }
                            if (result.ReQuestion[p].Answer[o].OrderNum == 1) {
                                op += '<div class="optionList"><span class="optionWords">B.  ' + result.ReQuestion[p].Answer[o].OptionName + '</span>' +
                                    '<div class="optionBtn"><img option="' + result.ReQuestion[p].Answer[o].OptionAnswer + '" tmpoption="" class="btnImg" src="/static/img/default.png"></div>' +
                                    '   </div>'
                            }
                            if (result.ReQuestion[p].Answer[o].OrderNum == 2) {
                                op += '<div class="optionList"><span class="optionWords">C.  ' + result.ReQuestion[p].Answer[o].OptionName + '</span>' +
                                    '<div class="optionBtn"><img option="' + result.ReQuestion[p].Answer[o].OptionAnswer + '" tmpoption="" class="btnImg" src="/static/img/default.png"></div>' +
                                    '   </div>'
                            }
                            if (result.ReQuestion[p].Answer[o].OrderNum == 3) {
                                op += '<div class="optionList"><span class="optionWords">D.  ' + result.ReQuestion[p].Answer[o].OptionName + '</span>' +
                                    '<div class="optionBtn"><img option="' + result.ReQuestion[p].Answer[o].OptionAnswer + '" tmpoption="" class="btnImg" src="/static/img/default.png"></div>' +
                                    '   </div>'
                            }
                        }

                        str += op

                        str += '<div class="btnBlock" style="display: none">' +
                            '<div class="btn">' +
                            '<div class="btnWords">答案解析</div>' +
                            '</div>' +
                            '</div>'

                        str += '<div class="answer">' +
                            '        <h4 class="answer-title">答案解析</h4>' +
                            '        <p class="answer-content"> ' + result.ReQuestion[p].AnswerAnalysis + ' </p>' +
                            '    </div>'

                        str += '</div>'

                    }
                    var currentper = $('.box').attr("tmpcurrentper")
                    var total = result.Total
                    if (total == currentper) {
                        $('.box').html('结束')
                        var currentper = $('.box').attr("tmpcurrentper")


                        $(".box").attr('isclick', 2)


                    }
                    // console.log(   getUrlParam("per") )

                    // if(result.Total == 1 && length(result.ReQuestion) == 10){
                    //     $('.box').html('结束')
                    // }

                    $(".header").after(str)
                } else {
                    $('.box').html('结束')
                }
            },

        })
        // $("#fullBoxPanlContentSpan").html(re.Num)
        // $("#fullBox").show()


        // 图标按钮
        $("body").on('click', '.optionBtn', function () {
            // console.log($(this).parent().parent().attr("tmpClick")  )
            if ($(this).parent().parent().attr("tmpClick") != 1) {
                // 多选
                var optiontype = $(this).parent().parent().attr("optiontype")

                if (optiontype == 1) {
                    // 单选 效果处理
                    $(this).parent().siblings("div.optionList").each(function () {
                        $(this).children("div.optionBtn").children("img.btnImg").attr("src", "/static/img/default.png")
                        $(this).children("div.optionBtn").children("img.btnImg").attr("tmpOption", 0)
                    })

                    $(this).children("img.btnImg").attr("tmpOption", 1)
                    $(this).children("img.btnImg").attr('src', '/static/img/right.png');

                    $(this).parent().siblings("div.btnBlock").show();
                } else {
                    // 多选 效果处理
                    if ($(this).children("img.btnImg").attr("tmpOption") == 1) {
                        $(this).children("img.btnImg").attr("tmpOption", "").attr('src', '/static/img/default.png');
                        $(this).parent().siblings("div.btnBlock").show();
                    } else {
                        $(this).children("img.btnImg").attr("tmpOption", 1).attr('src', '/static/img/right.png');
                        $(this).parent().siblings("div.btnBlock").show();
                    }

                }
            }
            return false;
        });

        // 答案解析按钮
        $("body").on('click', ".btnWords", function (e) {
            $(this).parent().parent().hide()
            $(this).parent().parent().siblings("div.answer").removeClass("answer")

            // 答案
            var answer = 2;

            var obj = $(this).parent().parent().parent().children(".optionList")
            // 批错对
            obj.each(function () {
                var option = $(this).children("div.optionBtn").children("img.btnImg").attr("option")
                var tmpoption = $(this).children("div.optionBtn").children("img.btnImg").attr("tmpoption")

                // 错误
                if (option == "0" && tmpoption == 1) {
                    $(this).children("div.optionBtn").children("img.btnImg").attr("src", "/static/img/wrong.png")
                    answer = 1;
                }
                if (option == "1" && tmpoption == "") {
                    $(this).children("div.optionBtn").children("img.btnImg").attr("src", "/static/img/wrong.png")
                    answer = 1;
                }
                //  正确
                if (option == "1") {
                    $(this).children("div.optionBtn").children("img.btnImg").attr("src", "/static/img/right.png")
                }
            })

            var questionId = $(this).parent().parent().parent().attr("questionId");
            var userId = $("#userId").val();
            console.log("answer: " + answer)
            if (answer != 1) {
                console.log("ddddd")
                var categoryTaoId = $("#categoryTaoId").val()
                $.ajax({
                    type: "POST",
                    url: ServerUrl + "reSubmitQuestion",
                    data: {userId: userId, questionId: questionId, ok: 2, categroyTaoId: categoryTaoId},
                    async: false,
                    method: 'post',
                    //请求成功
                    success: function (result) {
                        //alert(result)
                    }
                })

            }


            var flag = 2; // 1错误 2正确
            if (answer == 1) {
                flag = 1
            }


            $(this).parent().parent().parent().attr("tmpClick", 1)
            return false;

        })

        //  下一页面
        $(".box").click(function () {
            console.log($(".box").attr('isclick'))
            if ($(".box").attr('isclick') == "2") {
                return false
            }

            var lastQuestionId = $('.box').attr('lastQuestionId') // 题的Id
            console.log("lastQuestionId: " + lastQuestionId)

            var current = Number($(this).attr("tmpCurrentPer"));
            var next = current + 1

            var total = $('#Total').val();

            // var Total = Number(total) + 1


            $(".content").remove()
            var userId = $("#userId").val();



            $.ajax({
                type: "GET",
                url: ServerUrl + "getQuestionFenye?category=" + getUrlParam("category")+ "&userId=" + userId + "&per=" + next + "&lastQuestionId=" + lastQuestionId + "&type=2",
                //请求成功
                success: function (result) {
                    console.log(result);

                    var currentper = $('.box').attr("tmpcurrentper")

                    if (result.TotalPage == currentper) {
                        $('.box').html('结束')
                        $(".box").attr('isclick', 2)
                    }

                    if (result.ReQuestion) {

                        var str = "";
                        for (var p in result.ReQuestion) {
                            // console.log(result.ReQuestion[p].Title) //  标题
                            // console.log(result.ReQuestion[p].AnswerAnalysis) //  答案解析


                            str += '<div class="content" questionId="' + result.ReQuestion[p].Id + '" optionType="' + result.ReQuestion[p].OptionType + '" tmpclick="">'
                            str += '<h3 class="question"> ' + result.ReQuestion[p].Num + '.' + result.ReQuestion[p].Title + '</h3>'


                            var op = '';

                            for (var o in result.ReQuestion[p].Answer) {
                                if (result.ReQuestion[p].Answer[o].OrderNum == 0) {
                                    op += '<div class="optionList"><span class="optionWords">A.  ' + result.ReQuestion[p].Answer[o].OptionName + '</span>' +
                                        '<div class="optionBtn"><img option="' + result.ReQuestion[p].Answer[o].OptionAnswer + '" tmpoption="" class="btnImg" src="/static/img/default.png"></div>' +
                                        '   </div>'
                                }
                                if (result.ReQuestion[p].Answer[o].OrderNum == 1) {
                                    op += '<div class="optionList"><span class="optionWords">B.  ' + result.ReQuestion[p].Answer[o].OptionName + '</span>' +
                                        '<div class="optionBtn"><img option="' + result.ReQuestion[p].Answer[o].OptionAnswer + '" tmpoption="" class="btnImg" src="/static/img/default.png"></div>' +
                                        '   </div>'
                                }
                                if (result.ReQuestion[p].Answer[o].OrderNum == 2) {
                                    op += '<div class="optionList"><span class="optionWords">C.  ' + result.ReQuestion[p].Answer[o].OptionName + '</span>' +
                                        '<div class="optionBtn"><img option="' + result.ReQuestion[p].Answer[o].OptionAnswer + '" tmpoption="" class="btnImg" src="/static/img/default.png"></div>' +
                                        '   </div>'
                                }
                                if (result.ReQuestion[p].Answer[o].OrderNum == 3) {
                                    op += '<div class="optionList"><span class="optionWords">D.  ' + result.ReQuestion[p].Answer[o].OptionName + '</span>' +
                                        '<div class="optionBtn"><img option="' + result.ReQuestion[p].Answer[o].OptionAnswer + '" tmpoption="" class="btnImg" src="/static/img/default.png"></div>' +
                                        '   </div>'
                                }
                            }

                            str += op

                            str += '<div class="btnBlock" style="display: none">' +
                                '<div class="btn">' +
                                '<div class="btnWords">答案解析</div>' +
                                '</div>' +
                                '</div>'

                            str += '<div class="answer">' +
                                '        <h4 class="answer-title">答案解析</h4>' +
                                '        <p class="answer-content"> ' + result.ReQuestion[p].AnswerAnalysis + ' </p>' +
                                '    </div>'

                            str += '</div>'

                        }
                        // console.log('str:' + str)
                        $(".header").after(str)
                        // console.log("next: " + next)
                        // console.log("total: " + total)

                        // 结束按钮
                        // if (Number(next) == Number(total)) {
                        //     $('.box').html('结束')
                        // }
                    }
                },

            })
            $(this).attr("tmpCurrentPer", next)


            return false;
        })
    });
</script>
</body>
</html>
