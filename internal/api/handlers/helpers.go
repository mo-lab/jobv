package handlers

import (
	"net/http"
	"net/url"
	"strconv"
)

func CheckPaginationParams(q url.Values, w http.ResponseWriter) (int, int, bool) {
	pageStr := q.Get("page")
	limitStr := q.Get("limit")
	if pageStr == "" {
		pageStr = "1"
	}
	if limitStr == "" {
		limitStr = "10"
	}
	pg, err := strconv.Atoi(pageStr)
	if err != nil {
		http.Error(w, "iternal server error", http.StatusInternalServerError)
		return 0, 0, true
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		http.Error(w, "iternal server error", http.StatusInternalServerError)
		return 0, 0, true
	}
	return pg, limit, false
}
