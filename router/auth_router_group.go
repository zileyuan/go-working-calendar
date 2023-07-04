package router

import (
	"path"

	"github.com/gin-gonic/gin"
)

// AuthRouterGroup 可以验证Token的路由组
type AuthRouterGroup struct {
	*gin.RouterGroup
}

// Get 不需要验证Token
func (r *AuthRouterGroup) Get(relativePath string, handlers ...gin.HandlerFunc) {
	r.GET(relativePath, handlers...)
	r.AppendToAuth(relativePath, false)
}

// Post 不需要验证Token
func (r *AuthRouterGroup) Post(relativePath string, handlers ...gin.HandlerFunc) {
	r.POST(relativePath, handlers...)
	r.AppendToAuth(relativePath, false)
}

// Put 不需要验证Token
func (r *AuthRouterGroup) Put(relativePath string, handlers ...gin.HandlerFunc) {
	r.PUT(relativePath, handlers...)
	r.AppendToAuth(relativePath, false)
}

// Options 不需要验证Token
func (r *AuthRouterGroup) Options(relativePath string, handlers ...gin.HandlerFunc) {
	r.OPTIONS(relativePath, handlers...)
	r.AppendToAuth(relativePath, false)
}

// Patch 不需要验证Token
func (r *AuthRouterGroup) Patch(relativePath string, handlers ...gin.HandlerFunc) {
	r.PATCH(relativePath, handlers...)
	r.AppendToAuth(relativePath, false)
}

// Head 不需要验证Token
func (r *AuthRouterGroup) Head(relativePath string, handlers ...gin.HandlerFunc) {
	r.HEAD(relativePath, handlers...)
	r.AppendToAuth(relativePath, false)
}

// Delete 不需要验证Token
func (r *AuthRouterGroup) Delete(relativePath string, handlers ...gin.HandlerFunc) {
	r.DELETE(relativePath, handlers...)
	r.AppendToAuth(relativePath, false)
}

// AuthGet 需要验证Token
func (r *AuthRouterGroup) AuthGet(relativePath string, handlers ...gin.HandlerFunc) {
	r.GET(relativePath, handlers...)
	r.AppendToAuth(relativePath, true)
}

// AuthPost 需要验证Token
func (r *AuthRouterGroup) AuthPost(relativePath string, handlers ...gin.HandlerFunc) {
	r.POST(relativePath, handlers...)
	r.AppendToAuth(relativePath, true)
}

// AuthPut 需要验证Token
func (r *AuthRouterGroup) AuthPut(relativePath string, handlers ...gin.HandlerFunc) {
	r.PUT(relativePath, handlers...)
	r.AppendToAuth(relativePath, true)
}

// AuthOptions 需要验证Token
func (r *AuthRouterGroup) AuthOptions(relativePath string, handlers ...gin.HandlerFunc) {
	r.OPTIONS(relativePath, handlers...)
	r.AppendToAuth(relativePath, true)
}

// AuthPatch 需要验证Token
func (r *AuthRouterGroup) AuthPatch(relativePath string, handlers ...gin.HandlerFunc) {
	r.PATCH(relativePath, handlers...)
	r.AppendToAuth(relativePath, true)
}

// AuthHead 需要验证Token
func (r *AuthRouterGroup) AuthHead(relativePath string, handlers ...gin.HandlerFunc) {
	r.HEAD(relativePath, handlers...)
	r.AppendToAuth(relativePath, true)
}

// AuthDelete 需要验证Token
func (r *AuthRouterGroup) AuthDelete(relativePath string, handlers ...gin.HandlerFunc) {
	r.DELETE(relativePath, handlers...)
	r.AppendToAuth(relativePath, true)
}

// AuthAny 需要验证Token
func (r *AuthRouterGroup) AuthAny(relativePath string, handlers ...gin.HandlerFunc) {
	r.Any(relativePath, handlers...)
	r.AppendToAuth(relativePath, true)
}

// AppendToAuth 加入路由到状态表
func (r *AuthRouterGroup) AppendToAuth(relativePath string, value bool) {
	uri := r.calculateAbsolutePath(relativePath)
	AuthTokenRoutes[uri] = value
}

func (r *AuthRouterGroup) calculateAbsolutePath(relativePath string) string {
	return r.joinPaths(r.BasePath(), relativePath)
}

func (r *AuthRouterGroup) joinPaths(absolutePath, relativePath string) string {
	if relativePath == "" {
		return absolutePath
	}
	finalPath := path.Join(absolutePath, relativePath)
	appendSlash := r.lastChar(relativePath) == '/' && r.lastChar(finalPath) != '/'
	if appendSlash {
		return finalPath + "/"
	}
	return finalPath
}

func (r *AuthRouterGroup) lastChar(str string) uint8 {
	if str == "" {
		panic("The length of the string can't be 0")
	}
	return str[len(str)-1]
}
