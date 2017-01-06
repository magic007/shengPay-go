# shengPay for Go

[![Build Status](https://travis-ci.org/silenceper/wechat.svg?branch=master)]()
[![Go Report Card](https://goreportcard.com/badge/github.com/silenceper/wechat)]()

**shengPay for Go**，简单、易用。

* * *

> 快速开始：
> 
>  win系统执行`shengPayDemo.exe` ，Uinx* 执行shengPayDemo，输入订单号，查看演示支付数据。

* * *

### 所有调用都在`shengPay.go`，具体步骤安装一下流程

<table>
<thead>
<tr>
<th>*</th>
<th style="text-align:center">步骤一</th>
<th style="text-align:right">操作者</th>
</tr>
</thead>
<tbody>
<tr>
<td>1</td>
<td style="text-align:center">生成支付链接</td>
<td style="text-align:right">系统</td>
</tr>
<tr>
<td>2</td>
<td style="text-align:center">生成二维码</td>
<td style="text-align:right">系统</td>
</tr>
<tr>
<td>3</td>
<td style="text-align:center">扫描二维码</td>
<td style="text-align:right">用户</td>
</tr>
<tr>
<td>4</td>
<td style="text-align:center">回调</td>
<td style="text-align:right">盛付通</td>
</tr>
<tr>
<td>5</td>
<td style="text-align:center">处理回调</td>
<td style="text-align:right">系统</td>
</tr>
<tr>
<td>6</td>
<td style="text-align:center">提供订单验证查询</td>
<td style="text-align:right">系统</td>
</tr>
</tbody>
</table>

## PS.目录文件

    .
    ├── README<span class="hljs-class">.md</span>
    ├── shengPayDemo
    ├── shengPayDemo<span class="hljs-class">.exe</span>
    ├── shengPayDemo<span class="hljs-class">.go</span>                     <span class="hljs-comment">//Demo 文件</span>
    ├── shengPay<span class="hljs-class">.go</span>                         <span class="hljs-comment">//支付请求文件</span>
    ├── shengPayPaymentService<span class="hljs-class">.go</span>           <span class="hljs-comment">//支付解析</span>
    ├── shengPayQuery<span class="hljs-class">.go</span>                   <span class="hljs-comment">//查询解析</span>
    └── shengPayReceiveOrderService<span class="hljs-class">.go</span>     <span class="hljs-comment">//收单解析</span>
    `</pre>

    ## 一.生成支付链接

    ### 1.代码--请求

    <pre>`<span class="hljs-literal">result</span>, err := shengPay.<span class="hljs-type">Pay</span>(orderNO, money, productName)
    <span class="hljs-keyword">if</span> err != <span class="hljs-keyword">nil</span> {
        fmt.<span class="hljs-type">Println</span>(<span class="hljs-string">"错误"</span>, err)
    }

    fmt.<span class="hljs-type">Printf</span>(<span class="hljs-string">"返回结果:%+v\n"</span>,<span class="hljs-literal">result</span>.<span class="hljs-type">Return</span>)

    fmt.<span class="hljs-type">Println</span>(<span class="hljs-string">"订单号："</span>,orderNO,<span class="hljs-string">",支付地址"</span>, <span class="hljs-literal">result</span>.<span class="hljs-type">Return</span>.<span class="hljs-type">Extension</span>.<span class="hljs-type">Ext3</span>)
    `</pre>

    ### 2.代码--返回

    <pre>`<span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span>result.Return<span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span>
    BankForm:
    Extension:0xc0420ed230
    Header:0xc04213eb40
    OrderAmount:0.01
    OrderNo:magictest
    PaidAmount:0.00
    PaidFee:
    PayableAmount:
    PayableFee:0.00
    PaymentNo:CP20170106162648688660
    PaymentStatus:00
    PaymentTime:
    ReceivableFee:0.00
    ReceivedFee:
    ReturnInfo:0xc04213dc40
    SessionId:76a21298-0a02-4406-865d-05353e3e9147
    Signature:0xc04213dd00
    TransNo:C20170106162648729693
    TransStatus:00

    <span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span>result.Return.Extension.Ext3<span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span><span class="hljs-keyword">*</span>
    订单号： magictest ,支付地址 https://qr.alipay.com/bax00808vyn5m6fvvwx44098
    `</pre>> 请把支付地址用手机浏览器打开，或在线生成二维码,用手机支付宝APP扫描

    * * *

    ## 二.生成二维码

    > 支付宝，微信都支持把返回的`result.Return.Extension.Ext3`结果生成二维码，用APP扫描支付，其中支付宝结果还支持浏览器打开调用_支付宝APP支付_。

    ## 三.处理回调

    ### 1.盛付通回调,异步回调结果

    <pre>`<span class="hljs-label">Charset:</span>UTF-<span class="hljs-number">8</span>
    <span class="hljs-label">ErrorCode:</span>
    <span class="hljs-label">ErrorMsg:</span>
    <span class="hljs-label">Ext1:</span>
    <span class="hljs-label">Ext2:</span>
    <span class="hljs-label">InstCode:</span>SFT
    <span class="hljs-label">MerchantNo:</span><span class="hljs-number">10170080</span>
    <span class="hljs-label">MsgSender:</span>SFT
    <span class="hljs-label">Name:</span>REP_B2CPAYMENT
    <span class="hljs-label">OrderAmount:</span><span class="hljs-number">0.01</span>
    <span class="hljs-label">OrderNo:</span>b7fcb8a440f6f477f466c24a4c9e2eb2
    <span class="hljs-label">PayChannel:</span>wp
    <span class="hljs-label">PayableFee:</span><span class="hljs-number">0.00</span>
    <span class="hljs-label">PaymentNo:</span>CP20170105151804971807
    <span class="hljs-label">ReceivableFee:</span><span class="hljs-number">0.00</span>
    <span class="hljs-label">SendTime:</span><span class="hljs-number">20170105151840</span>
    <span class="hljs-label">SignMsg:</span>C6B099EF637EF71E984C1DA5E160AF9B
    <span class="hljs-label">SignType:</span>MD5
    <span class="hljs-label">TraceNo:</span>b7533cdc-<span class="hljs-number">3071</span>-<span class="hljs-number">459</span>d-a199-<span class="hljs-number">97</span>c2fb01d860
    <span class="hljs-label">TransAmount:</span><span class="hljs-number">0.01</span>
    <span class="hljs-label">TransNo:</span>C20170105151758868679
    <span class="hljs-label">TransStatus:</span><span class="hljs-number">01</span>
    <span class="hljs-label">TransTime:</span><span class="hljs-number">20170105151827</span>
    <span class="hljs-label">TransType:</span>PT314
    <span class="hljs-label">Version:</span>V4<span class="hljs-number">.1</span><span class="hljs-number">.2</span><span class="hljs-number">.1</span><span class="hljs-number">.1</span>
    <span class="hljs-label">paymentDetail:</span>
    `</pre>

    ### 2.验证支付是否成功，修改本系统订单状态，标注支付成功

    <pre>`r := c.<span class="hljs-type">Ctx</span>.<span class="hljs-type">Request</span>
    b, _ := ioutil.<span class="hljs-type">ReadAll</span>(r.<span class="hljs-type">Body</span>)
    defer r.<span class="hljs-type">Body</span>.<span class="hljs-type">Close</span>()
    <span class="hljs-keyword">var</span> form url.<span class="hljs-type">Values</span>
    <span class="hljs-keyword">if</span> len(b) == <span class="hljs-number">0</span> {
        r.<span class="hljs-type">ParseForm</span>()
        form = r.<span class="hljs-type">Form</span>    
    } <span class="hljs-keyword">else</span> {
        form, _ = url.<span class="hljs-type">ParseQuery</span>(<span class="hljs-type">string</span>(b))
        fmt.<span class="hljs-type">Println</span>(<span class="hljs-string">"form"</span>, form)    
    }
    notifyParams := map[<span class="hljs-type">string</span>]<span class="hljs-type">string</span>{}
    <span class="hljs-keyword">for</span> k, v := <span class="hljs-type">range</span> form {
        notifyParams[k] = v[<span class="hljs-number">0</span>]
    }

    orderId = notifyParams[<span class="hljs-string">"OrderNo"</span>]
    <span class="hljs-literal">result</span>, err := shengPay.<span class="hljs-type">Query</span>(orderId)

    <span class="hljs-keyword">if</span> err != <span class="hljs-keyword">nil</span> {
        fmt.<span class="hljs-type">Println</span>(<span class="hljs-string">"错误信息:"</span>, err)
        <span class="hljs-keyword">return</span>
    }

    fmt.<span class="hljs-type">Printf</span>(<span class="hljs-string">"%+v\n Extension:%+v"</span>, <span class="hljs-literal">result</span>.<span class="hljs-type">Return</span>.<span class="hljs-type">OrderNo</span>, err)

    <span class="hljs-keyword">if</span> <span class="hljs-literal">result</span>.<span class="hljs-type">Return</span>.<span class="hljs-type">OrderNo</span> == <span class="hljs-string">""</span> {    
        c.<span class="hljs-type">Ctx</span>.<span class="hljs-type">Output</span>.<span class="hljs-type">Body</span>([]byte(<span class="hljs-string">"error"</span>))
        <span class="hljs-keyword">return</span>
    }

    //标注支付状态
    `</pre>

    ## 四.提供订单验证查询

    <pre>`<span class="hljs-literal">result</span>, err := shengPay.<span class="hljs-type">Query</span>(orderId)

    <span class="hljs-keyword">if</span> err != <span class="hljs-keyword">nil</span> {
        fmt.<span class="hljs-type">Println</span>(<span class="hljs-string">"错误信息:"</span>, err)
        <span class="hljs-keyword">return</span>
    }

* * *

> 目前完成以上几步，更多接口请阅读代码。