package api
import (
	"fmt"
	"strconv"
	"time"
	"net/http"
	"crypto/rand"
	"github.com/gin-gonic/gin"
)



var usinformation = gin.H{
	"lala":gin.H{"email":"mclovin@gmail.com","token":""},
	"york":gin.H{"email":"york@gmail.com","token":""},

}

func status(c* gin.Context){
	user:=c.MustGet(gin.AuthUserKey).(string)
	dt := time. Now()
	fmt.Println(maptokens)
	if _,usok:=maptokens[user];usok{
		c.JSON(http.StatusOK,gin.H{"message": "Hi username, the DPIP System is Up and Running",
			"time": dt})
	}else{
		c.AbortWithStatus(401)
	}
}


func uploadfile(c*  gin.Context){


	_,header,err :=c.Request.FormFile("image")
	if err!=nil{
		return
	}
	size:= strconv.Itoa(int(header.Size))
	c.JSON(http.StatusOK,gin.H{"status":"SUCCESS","Filename":header.Filename,"filesize":size+" bytes"})



}

func logoff(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)

	if _, usok := maptokens[user]; usok {

		delete(maptokens,user)
		c.AbortWithStatus(401)
		c.JSON(http.StatusOK,gin.H{"message": "Bye username, your token has been revoked" })
		return
	}else{
		c.AbortWithStatus(401)

	}

}

func generatetoken() string {
	b := make([]byte, 15)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

//it utilizes the user password to save
var maptokens =make(map[string]string)

func login(c *gin.Context){
	user:=c.MustGet(gin.AuthUserKey).(string)
	token:=generatetoken()

	maptokens[user]=token

	if _,usok:=usinformation[user];usok{
		c.JSON(http.StatusOK,gin.H{"message": "Hi username, welcome to the DPIP System",
			"token": maptokens[user]})
	}else{
		c.AbortWithStatus(401)
	}
}

func main(){
	r:= gin.Default()
	r.Use()
	authorization:=r.Group("/",gin.BasicAuth(gin.Accounts{"user":"mclovin","york":"mribs"}))
	authorization.GET("/login", login)
	authorization.GET("/logout", logoff)
	authorization.GET("/status", status)
	authorization.GET("/upload", uploadfile)
	r.Run(":8080")


}


