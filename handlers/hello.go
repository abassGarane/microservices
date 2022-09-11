package handlers

import "net/http"

type HelloHandler struct{
}


func (h *HelloHandler) ServeHTTP( w http.ResponseWriter, r *http.Request)  {
  
}
