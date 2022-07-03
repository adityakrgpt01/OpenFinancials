package controller

import "bitbucket.org/exotel/exotel_code/gollum/fsm/requests"


func AddUrl(reqParam ){

	req:=exractReqBody(reqParam)

	CreateShortedUrl(req.ActualUrl,req.id)

}


func exractReqBody(req string)Request{

	newReq:=Reques{}
	Json.Unmarshall(req,&newReq)

	return newReq

}