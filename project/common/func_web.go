package comm

import (
	"LuckyDraw/project/models"
	"net"
	"net/http"
	"net/url"
	"strconv"
)

func ClientIP(request *http.Request) string {
	host, _, _ := net.SplitHostPort(request.RemoteAddr)
	return host
}

func Redirect(writer http.ResponseWriter, url string) {
	writer.Header().Add("Location", url)
	writer.WriteHeader(http.StatusFound)
}

func GetLoginUser(request *http.Request) *models.ObjLoginuser {
	c, err := request.Cookie("login_user")
	if err != nil {
		return nil
	}
	params, err := url.ParseQuery(c.Value)
	if err != nil {
		return nil
	}
	uid, err := strconv.Atoi(params.Get("uid"))
	if err != nil || uid < 1 {
		return nil
	}
	loginuser := &models.ObjLoginuser{}
	loginuser.Uid = uid
	loginuser.Username = params.Get("username")
	loginuser.Ip = ClientIP(request)
	loginuser.Sign = params.Get("sign")

	return loginuser
}

func SetLoginuser(writer http.ResponseWriter, loginuser *models.ObjLoginuser) {
	params := url.Values{}
	params.Add("uid", strconv.Itoa(loginuser.Uid))
	c := &http.Cookie{
		Name:  "login_user",
		Value: params.Encode(),
		Path:  "/",
	}
	http.SetCookie(writer, c)
}
