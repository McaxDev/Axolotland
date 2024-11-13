## 开发文档
### HTTP协议注释
* 有一些规则是不断重复说明的，在这一部分列出。
#### 请求头带JWT
* 请求头加上`Authorization: Bearer 234f62.236f3245.345vy3`
#### 服务器ID 枚举类型
* 互通服是`paper`，生电服是`fabric`，基岩服是`bedrock`，饥荒服是`dst`，星露谷物语服是`stardew`，泰拉瑞亚服是`terraria`
### HTTP协议
* 请求体都是application/json类型，而且对象的键，大小写皆可。例如，如果文档写的键是`Username`，那么`username`或`UserName`或`USERNAME`皆可，但`user_name`就不行。
* 响应体也是application/json类型，但对象的键是大小写严格的。例如，如果文档写`userId`，那么`userID`或`userid`或`UserId`或`user_id`都不行。
* 响应体默认是此结构，后文只解释`data`部分：
```json
{
    "message": "请求成功",
    "data": 可能是任何类型
}
```
#### 注册 POST
* 路径：`/account/signup`
* 请求体：
```json
{
    "Username": "Nerakolo",
    "Password": "9aosy3094j82j4ct9",
    "Email": "nerakolo@outlook.com",
    "EmailCode": "114514",
    "CaptchaID": "j20934yt0909234t8",
    "CaptchaValue": "114514"
}
```
* 响应体data部分：
```json
{
    "token": "Bearer 432f5.342f52.234vf523"
}
```
#### 登录 POST
* 路径：`/account/login`
* 请求体：
```json
{
    "Username": "Nerakolo",
    "Password": "9aosy3094j82j4ct9",
    "CaptchaID": "j20934yt0909234t8",
    "CaptchaValue": "114514"
}
```
* 响应体data部分：
```json
{
    "token": "Bearer 432f5.342f52.234vf523"
}
```
#### 获取用户信息 GET
* 路径：`/account/get/userinfo`
* 请求头带JWT
* 响应体：
```json
{
    "userId": 123,
    "username": "Nerakolo",
    "avatar": "filename.jpg",
    "profile": "Ciallo!",
    "admin": true,
    "money": 100,
    "email": "nerakolo@outlook.com",
    "telephone": "12312341234",
    "bedrockName": "Nerakolox",
    "javaName": "Nerakolo",
    "dstName": "nerakolo"
}
```
#### 获取Captcha验证码 GET
* 路径：`/verify/captcha`
* 响应头：`X-Captcha-Id: jc8u9wty8jcw90t35`、`Content-Type: image/png`
#### 获取邮件验证码 GET
* 路径：`/verify/email/nerakolo@outlook.com`
#### 获取手机验证码 GET
* 路径：`/verify/sms/12312341234`
#### 获取特定玩家的统计信息 GET
* 路径：`/dashboard/player/Nerakolo`
* 查询字符串参数：`server=paper`，见上文“服务器ID”
#### 获取特定统计信息的排行榜 GET
* 路径：`/dashboard/:stat`，例如`/dashboard/break`，`:stat`包括`mined`,`picked_up`,`crafted`,`broken`,`play_time`,`deaths`,`mob_kills`,`damage_dealt`,`drop`
* 查询字符串参数：`server=paper`，见上文“服务器ID”
* 响应体data部分：
```json
[
    {
        "Score": 100,
        "Member": "Nerakolo"
    },
    {
        "Score": 90,
        "Member": "Bestcb233"
    }
]
```
