## 开发文档
### HTTP协议
* 请求体都是Application/json类型，而且对象的键，大小写皆可。例如，如果文档写的键是`Username`，那么`username`或`UserName`或`USERNAME`皆可，但`user_name`就不行。
* 响应体默认全是下面结构，下面的接口只说明`data`部分的结构：
```json
{
    "message": "请求成功",
    "data": 可能是任何类型
}
```
#### 注册
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

