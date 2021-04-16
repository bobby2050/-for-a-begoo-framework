<!DOCTYPE html>

<html>
<head>
    <title>用户注册</title>
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
            margin: 0 auto;
            width: 100%;
            height: 239px;
            background: #3f8eea url("/static/img/top.png") bottom no-repeat;
            background-size: contain;
        }


        .top3 img {
            height: 45px;
        }

        .top2-font-nickname {

            display: block;
            margin: 0 auto;
            padding-top: 40px;
            font-size: 18px;
            width: 90px;
            height: 90px;
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
            height: 370px;
            background: #fff;
            background-size: contain;
            margin: 0 auto;
            border-radius: 8px;
            margin-bottom: 20px;

        }

        .item-block-title a {
            color: #000;
            text-decoration: none

        }

        .input-title {
            padding-top: 30px;
            /*padding-left: 5px;*/
            font-size: 18px;
            padding-bottom: 30px;
            float: left;

        }

        .input-text {
            height: 30px;

            border: none;
            /*width: 80%;*/
            padding-bottom: 20px;
            float: left;
            font-size: 16px;

        }

        input {
            background: none;
            outline: none;
            border: none;
        }



        .input-text1 {
            clear: both;
            /*background: fuchsia;*/
            width: 100%;
            height: 30px;
            margin-bottom: 2px;
            border-bottom: 1px solid #ccc;
            /*float: left;*/
        }

        .input-text2 {
            clear: both;
            background:url('/static/img/ok.png') no-repeat right bottom;
            background-size: contain;
            width: 80%;
            margin-left: 10%;
            height: 60px;
            margin-bottom: 2px;

            /*float: left;*/
        }

        .m-item-block {
            width: 90%;
            height: 380px;
            margin: 0 auto;
            /*background: fuchsia;*/
        }
    </style>
</head>

<body>
<div id="header">
    <div class="top2-font-nickname"><img width="90" src="/static/img/h.png"/></div>
</div>

<div id="body-block">

    <div class="item-block">
        <div class="m-item-block">
            <div class="input-title">请完善您的注册信息</div>
            <div class="input-text1"><input class="input-text" id="username" name="username"  placeholder="请输入您的姓名"/></div>
            <div class="input-text1"><input class="input-text" id="position" name="position"  placeholder="请输入您所在的岗位"/></div>
            <div class="input-text1"><input class="input-text" id="mobileNum" name="mobileNum"  placeholder="请输入您的手机号"/></div>
            <div class="input-text1"><input class="input-text" id="reCode" placeholder="请输入短信验证码"/>
                <span id="sendBtn" style="color: #e8662c;float: right;">发送验证码</span>
            </div>
            <div class="input-text2" id="btn">

            </div>
        </div>
    </div>

</div>


<script src="/static/js/jquery.min.js"></script>
<script type="text/javascript" src="https://res.wx.qq.com/open/js/jweixin-1.3.0.js"></script>

<script type="application/javascript">
    //获取url中的参数
    function getUrlParam(name) {
        var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
        var r = window.location.search.substr(1).match(reg);  //匹配目标参数
        if (r != null) return unescape(r[2]);
        return null; //返回参数值
    }

    $(function () {
        var ServerUrl = "https://cw.bestjan.com/"
        // var ServerUrl = "http://127.0.0.1:8080/"

        var avatarUrl = getUrlParam("avatarUrl");
        var nickName = getUrlParam("nickName");
        var category = getUrlParam("category");

        var userId = getUrlParam("userId");
        $("#btn").click(function () {
            var username = $("#username").val()
            if (username == "") {
                $("#username").focus()
                return false
            }
            var position = $("#position").val()
            if (position == "") {
                $("#position").focus()
                return false
            }
            var mobileNum = $("#mobileNum").val()
            if (mobileNum == "") {
                $("#mobileNum").focus()
                return false
            }

            var reCode = $("#reCode").val()
            if (reCode == "") {
                $("#reCode").focus()
                return false
            }

            $.ajax({
                type: "POST",
                url: ServerUrl + "submitUserInfo",
                data: {
                    username: username,
                    position: position,
                    mobileNum: mobileNum,
                    reCode: reCode,
                    userId: userId,
                },
                //请求成功
                success: function (re) {
                    console.log("444")
                    console.log(re)
                    if (re == "succ") {
                        var redict_url = ServerUrl + "page?avatarUrl=" + avatarUrl + "&userId=" + userId + "&nickName=" + escape(nickName) + "&category=" + category;
                        // window.location.href = redict_url
                        wx.miniProgram.navigateTo({url: '/pages/index/index'})
                    } else {
                        $("#reCode").focus()
                    }
                }
            })

            return false
        })

        $("#sendBtn").click(function () {
            var mobileNum = $("#mobileNum").val()
            if (mobileNum == "") {
                $("#mobileNum").focus()
                return false
            } else {
                $.ajax({
                    type: "POST",
                    url: ServerUrl + "sendSms",
                    data: {
                        mobile: mobileNum
                    },
                    //请求成功
                    success: function (re) {
                        console.log("444")
                    }
                })
            }
            console.log(mobileNum)
            return false
        })

    })

</script>
</body>
</html>
