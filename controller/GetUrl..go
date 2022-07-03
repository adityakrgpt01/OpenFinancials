package controller





func GetShortedUrl(shortedUrl string){
	return models.GetActualUrl(shortedUrl)
}