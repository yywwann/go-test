//package main
//
//import (
//	"crypto/sha256"
//	"encoding/base64"
//	"encoding/hex"
//	"encoding/json"
//	"fmt"
//	"gitlab.33.cn/btrade/auto_trade_tools/reqtypes"
//	"gitlab.33.cn/btrade/auto_trade_tools/util"
//	"time"
//)
//
//var (
//	url            = "https://rtcsms.cn-north-1.myhuaweicloud.com:10743/sms/batchSendSms/v1"
//	appKey         = "qc6kcqz5VHwUESlONP98dZO2BKX5"
//	appSecret      = "TSPsjKqJDke8CEO1D29PVDNxfk9d"
//	sender         = "8821092731831"
//	templateId     = "882e4abe269e45aeae9f0c5d31f12d8f"
//	signature      = "%E5%BE%AE%E6%9D%B0%E7%A7%91%E6%8A%80"
//	receiver       = "15158496687"
//	statusCallBack = ""
//	templateParam  = "%5B%22123456%22%5D"
//)
//
//const (
//	authorization = "WSSE realm=\"SDP\",profile=\"UsernameToken\",type=\"Appkey\""
//)
//
//func buildWSSEHeader(appKey, appSecret string) string {
//	sdf := "2021-09-30T16:48:00Z"
//	nonce := "b67cd25268f74ccb957c80c6f8ec35b1"
//	msg := []byte(nonce + sdf + appSecret)
//	digest := sha256.Sum256(msg)
//	hexDigest := hex.EncodeToString(digest[:])
//
//	base64HexDigest := base64.StdEncoding.EncodeToString([]byte(hexDigest))
//
//	return fmt.Sprintf("UsernameToken Username=\"%s\",PasswordDigest=\"%s\",Nonce=\"%s\",Created=\"%s\"",
//		appKey, base64HexDigest, nonce, sdf)
//}
//
//func main() {
//	headers := map[string]string{}
//	headers["Authorization"] = authorization
//	headers["Content-Type"] = "application/x-www-form-urlencoded"
//	headers["X-WSSE"] = buildWSSEHeader(appKey, appSecret)
//	params := fmt.Sprintf("form=%s&to=%s&templateId=%s&templateParas=%s&statusCallback=%s&signature=%s",
//		sender, receiver, templateId, templateParam, statusCallBack, signature)
//
//	fmt.Println(buildWSSEHeader(appKey, appSecret))
//	fmt.Println(params)
//
//	bytes, err := util.HttpReq(&reqtypes.HttpParams{
//		Method:    "POST",
//		ReqUrl:    url,
//		HeaderMap: headers,
//		Timeout:   10 * time.Second,
//		StrParams: params,
//	})
//	if err != nil {
//		fmt.Println("err 1 = ", err)
//		return
//	}
//
//	var res map[string]interface{}
//	err = json.Unmarshal(bytes, &res)
//	if err != nil {
//		fmt.Println("err 2 = ", err)
//		return
//	}
//
//	fmt.Println(res)
//}
