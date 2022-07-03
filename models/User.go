package models

var(
	Id=0
)

type User struct{
	UserId int
	Name string
	UrlList []string
}


var UserTable map[int]User

func (User)CreateUser(name string){
    
	newUser:=User{
		UserId: Id,
		Name: name,
	}

	UserTable[Id]=newUser
	Id++

}


func (User) AddUrl(url string,id int){

userRecord:=UserTable[id]
userRecord.UrlList = append(userRecord.UrlList, url)
UserTable[id]=userRecord
}

