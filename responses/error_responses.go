package responses

import (
	"github.com/lucaiatropulus/social/internal/utils"
	"go.uber.org/zap"
	"net/http"
)

func InternalServerError(w http.ResponseWriter, r *http.Request, err error, logger *zap.SugaredLogger) {
	logger.Errorw("Internal Server Error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	utils.WriteJSONError(w, http.StatusInternalServerError, "The server encountered a problem")
}

func BadRequest(w http.ResponseWriter, r *http.Request, err error, reason string, logger *zap.SugaredLogger) {
	logger.Warnf("Bad Request", "method", r.Method, "path", r.URL.Path, "reason", reason, "error", err.Error())
	utils.WriteJSONError(w, http.StatusBadRequest, reason)
}

func NotFound(w http.ResponseWriter, r *http.Request, err error, logger *zap.SugaredLogger) {
	logger.Warnf("Resource Not Found", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	utils.WriteJSONError(w, http.StatusNotFound, "We were unable to find the requested resource")
}

func Conflict(w http.ResponseWriter, r *http.Request, err error, logger *zap.SugaredLogger) {
	logger.Errorw("Conflict Error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	utils.WriteJSONError(w, http.StatusConflict, err.Error())
}

func Unauthorized(w http.ResponseWriter, r *http.Request, logger *zap.SugaredLogger) {
	logger.Warnf("Unauthorized Error", "method", r.Method, "path", r.URL.Path)
	utils.WriteJSONError(w, http.StatusUnauthorized, "Your credentials are invalid")
}

func Forbidden(w http.ResponseWriter, r *http.Request, logger *zap.SugaredLogger) {
	logger.Warnf("Forbidden Error", "method", r.Method, "path", r.URL.Path)
	utils.WriteJSONError(w, http.StatusForbidden, "You are not allowed to make this action")
}

func RateLimitExceededResponse(w http.ResponseWriter, r *http.Request, retryAfter string, logger *zap.SugaredLogger) {
	logger.Warnw("Rate limit exceeded", "method", r.Method, "path", r.URL.Path)
	w.Header().Set("Retry-After", retryAfter)
	utils.WriteJSONError(w, http.StatusTooManyRequests, "Rate limit exceeded, retry after: "+retryAfter)
}
