package models

var(id=0)
	




type Url struct{
	UrlId  int
	UserId int
	ActualUrl string
	ShortedUrl string
	TTL   int
	CreationTime  string
}


var UrlTable map[int]Url

func (Url)CreateShortedUrl(ActualUrl string,userId int){

	ShortedUrl:=
	newUrl:=Url{
		UrlId: id,
		UserId: userId,
		ActualUrl: ActualUrl,
		ShortedUrl: ShortedUrl,
	}
	UrlTable[id]=newUrl
}


func(Url)GetActualUrl(shortedUrl string)string{

	for _,val:=range UrlTable{
		if val.ShortedUrl==shortedUrl{
			return val.ActualUrl
		}
	}

	return "Not Found"
}