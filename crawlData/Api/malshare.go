package api

import (
	"crawl/database/repository"
	"crawl/model"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type MalshareDailyHandler struct {
	MalshareDailyRepository repository.MalshareDailyRepository
}

func getByPage(start string, end string, limit string, page string, arr []model.MalshareDaily) ([]model.MalshareDaily, error) {
	st, errST := strconv.Atoi(start)
	ed, errED := strconv.Atoi(end)
	lm, errLM := strconv.Atoi(limit)
	pg, errPG := strconv.Atoi(page)
	if errST != nil {
		st = 0
	}
	if errED != nil {
		ed = len(arr)
	}
	if errLM != nil {
		lm = len(arr)
	}
	if errPG != nil {
		pg = 1
	}
	divideResult := len(arr) / lm
	surplus := len(arr) % lm
	if surplus == 0 && pg > divideResult || surplus != 0 && pg > divideResult+1 {
		return nil, errors.New("don't have record")
	}
	if surplus != 0 && pg == divideResult+1 {
		return arr[lm*(pg-1) : ed], nil
	}
	return arr[st:ed], nil

}
func (h *MalshareDailyHandler) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	query := r.URL.Query()
	start := query.Get("start")
	end := query.Get("end")
	limit := query.Get("limit")
	page := query.Get("page")
	MalshareDaily, err := h.MalshareDailyRepository.All()
	if err != nil {
		WriteJSON(w, HTTP_ERROR_CODE_READ_FAILED, ResponseBody{
			Message: "unable to get User from server",
		})
	}
	results, err := getByPage(start, end, limit, page, MalshareDaily)
	if err != nil {
		WriteJSON(w, http.StatusNotFound, ResponseBody{
			Message: err.Error(),
			Code:    http.StatusNotFound,
		})
		return
	}
	WriteJSON(w, http.StatusOK, results)
}
func (h *MalshareDailyHandler) GetByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	MalshareDaily, err := h.MalshareDailyRepository.FindByID(id)
	if err != nil {
		WriteJSON(w, HTTP_ERROR_CODE_READ_FAILED, ResponseBody{
			Message: "unable to get User by id : " + id,
		})
	}
	WriteJSON(w, http.StatusOK, MalshareDaily)
}

func (h *MalshareDailyHandler) GetBySha256(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	Sha256 := p.ByName("Sha256")
	MalshareDaily, err := h.MalshareDailyRepository.FindByMd5(Sha256)
	if err != nil {
		WriteJSON(w, HTTP_ERROR_CODE_READ_FAILED, ResponseBody{
			Message: "unable to get User by id : " + Sha256,
		})
	}
	WriteJSON(w, http.StatusOK, MalshareDaily)
}

func (h *MalshareDailyHandler) GetBySha1(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	Sha1 := p.ByName("Md5")
	MalshareDaily, err := h.MalshareDailyRepository.FindByMd5(Sha1)
	if err != nil {
		WriteJSON(w, HTTP_ERROR_CODE_READ_FAILED, ResponseBody{
			Message: "unable to get User by id : " + Sha1,
		})
	}
	WriteJSON(w, http.StatusOK, MalshareDaily)
}

func (h *MalshareDailyHandler) GetByBase64(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	Base64 := p.ByName("Base64")
	MalshareDaily, err := h.MalshareDailyRepository.FindByMd5(Base64)
	if err != nil {
		WriteJSON(w, HTTP_ERROR_CODE_READ_FAILED, ResponseBody{
			Message: "unable to get User by id : " + Base64,
		})
	}
	WriteJSON(w, http.StatusOK, MalshareDaily)
}
func (h *MalshareDailyHandler) GetByMd5(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	Md5 := p.ByName("Md5")
	MalshareDaily, err := h.MalshareDailyRepository.FindByMd5(Md5)
	if err != nil {
		WriteJSON(w, HTTP_ERROR_CODE_READ_FAILED, ResponseBody{
			Message: "unable to get User by id : " + Md5,
		})
	}
	WriteJSON(w, http.StatusOK, MalshareDaily)
}
func (h *MalshareDailyHandler) RemoveByMd5(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	Md5 := p.ByName("Md5")
	if err := h.MalshareDailyRepository.RemoveByMd5(Md5); err != nil {
		WriteJSON(w, HTTP_ERROR_CODE_DELETE_FAILED, ResponseBody{
			Message: err.Error(),
			Code:    HTTP_ERROR_CODE_DELETE_FAILED,
		})
		return
	}
	WriteJSON(w, http.StatusOK, ResponseBody{
		Message: "Delete succesfully",
	})
}
func (h *MalshareDailyHandler) RemoveBySha256(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	Sha256 := p.ByName("Sha256")
	if err := h.MalshareDailyRepository.RemoveBySha256(Sha256); err != nil {
		WriteJSON(w, HTTP_ERROR_CODE_DELETE_FAILED, ResponseBody{
			Message: err.Error(),
			Code:    HTTP_ERROR_CODE_DELETE_FAILED,
		})
		return
	}
	WriteJSON(w, http.StatusOK, ResponseBody{
		Message: "Delete succesfully",
	})
}

func (h *MalshareDailyHandler) RemoveBySha1(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	Sha1 := p.ByName("Sha1")
	if err := h.MalshareDailyRepository.RemoveByMd5(Sha1); err != nil {
		WriteJSON(w, HTTP_ERROR_CODE_DELETE_FAILED, ResponseBody{
			Message: err.Error(),
			Code:    HTTP_ERROR_CODE_DELETE_FAILED,
		})
		return
	}
	WriteJSON(w, http.StatusOK, ResponseBody{
		Message: "Delete succesfully",
	})
}

func (h *MalshareDailyHandler) RemoveByBase64(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	Base64 := p.ByName("Md5")
	if err := h.MalshareDailyRepository.RemoveByBase64(Base64); err != nil {
		WriteJSON(w, HTTP_ERROR_CODE_DELETE_FAILED, ResponseBody{
			Message: err.Error(),
			Code:    HTTP_ERROR_CODE_DELETE_FAILED,
		})
		return
	}
	WriteJSON(w, http.StatusOK, ResponseBody{
		Message: "Delete succesfully",
	})
}
