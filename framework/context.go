package framework

import (
	"context"
	"net/http"
	"sync"
	"time"
)

type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	ctx            context.Context
	hasTimeout     bool
	writerMux      *sync.Mutex
	handlers       []Controller
	index          int
	params         map[string]string
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
		ctx:            r.Context(),
		writerMux:      &sync.Mutex{},
		index:          -1,
	}
}

// #region base function

func (ctx *Context) WriterMux() *sync.Mutex {
	return ctx.writerMux
}

func (ctx *Context) GetRequest() *http.Request {
	return ctx.request
}

func (ctx *Context) GetResponse() http.ResponseWriter {
	return ctx.responseWriter
}

func (ctx *Context) SetHasTimeout() {
	ctx.hasTimeout = true
}

func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}

func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}
func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.BaseContext().Deadline()
}
func (ctx *Context) Err() error {
	return ctx.BaseContext().Err()
}

func (ctx *Context) Value(key interface{}) interface{} {
	return ctx.BaseContext().Value(key)
}

// func (ctx *Context) QueryInt(key string, def int) int {
// 	params := ctx.QueryAll()
// 	if vals, ok := params[key]; ok {
// 		len := len(vals)
// 		if len > 0 {
// 			intval, err := strconv.Atoi(vals[len-1])
// 			if err != nil {
// 				return def
// 			}
// 			return intval
// 		}
// 	}
// 	return def
// }

// func (ctx *Context) QueryString(key string, def string) string {
// 	params := ctx.QueryAll()
// 	if vals, ok := params[key]; ok {
// 		len := len(vals)
// 		if len > 0 {
// 			return vals[len-1]
// 		}
// 	}
// 	return def
// }

func (ctx *Context) QueryArray(key string, def []string) []string {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		return vals
	}
	return def
}

// func (ctx *Context) QueryAll() map[string][]string {
// 	if ctx.request != nil {
// 		return map[string][]string(ctx.request.URL.Query())
// 	}
// 	return map[string][]string{}
// }

// #endregion

// #region form post
// func (ctx *Context) FormInt(key string, def int) int {
// 	params := ctx.FormAll()
// 	if vals, ok := params[key]; ok {
// 		len := len(vals)
// 		if len > 0 {
// 			intval, err := strconv.Atoi(vals[len-1])
// 			if err != nil {
// 				return def
// 			}
// 			return intval
// 		}
// 	}
// 	return def
// }

func (ctx *Context) FormString(key string, def string) string {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		len := len(vals)
		if len > 0 {
			return vals[len-1]
		}
	}
	return def
}

func (ctx *Context) FormArray(key string, def []string) []string {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		return vals
	}
	return def
}

// func (ctx *Context) FormAll() map[string][]string {
// 	if ctx.request != nil {
// 		return map[string][]string(ctx.request.PostForm)
// 	}
// 	return map[string][]string{}
// }

// #endregion

// #region application/json post

// func (ctx *Context) BindJson(obj interface{}) error {
// 	if ctx.request != nil {
// 		body, err := ioutil.ReadAll(ctx.request.Body)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

// 		err = json.Unmarshal(body, obj)
// 		if err != nil {
// 			return err
// 		}
// 	} else {
// 		return errors.New("ctx.request empty")
// 	}
// 	return nil
// }

// #endregion

// #region response

// func (ctx *Context) Json(status int, obj interface{}) error {
// 	if ctx.HasTimeout() {
// 		return nil
// 	}
// 	ctx.responseWriter.Header().Set("Content-Type", "application/json")
// 	ctx.responseWriter.WriteHeader(status)
// 	byt, err := json.Marshal(obj)
// 	if err != nil {
// 		ctx.responseWriter.WriteHeader(500)
// 		return err
// 	}
// 	ctx.responseWriter.Write(byt)
// 	return nil
// }

func (ctx *Context) HTML(status int, obj interface{}, template string) error {
	return nil
}

// func (ctx *Context) Text(status int, obj string) error {
// 	return nil
// }

func (ctx *Context) SetHandlers(handlers []Controller) {
	ctx.handlers = handlers
}

func (ctx *Context) SetParams(params map[string]string) {
	ctx.params = params
}

func (ctx *Context) Next() error {
	ctx.index++
	if ctx.index < len(ctx.handlers) {
		if err := ctx.handlers[ctx.index](ctx); err != nil {
			return err
		}
	}
	return nil
}

// func Fool(request *http.Request, response http.ResponseWriter) {
// 	obj := map[string]interface{}{
// 		"data": nil,
// 	}

// 	response.Header().Set("Content-Type", "application/json")

// 	foo := request.PostFormValue("foo")
// 	if foo == "" {
// 		foo = "10"
// 	}

// 	fooInt, err := strconv.Atoi(foo)
// 	if err != nil {
// 		response.WriteHeader(500)
// 		return
// 	}

// 	obj["data"] = fooInt
// 	byt, err := json.Marshal(obj)
// 	if err != nil {
// 		response.WriteHeader(500)
// 		return
// 	}

// 	response.WriteHeader(200)
// 	response.Write(byt)
// 	return
// }

// func Fool2(ctx *framework.Context) error {

// 	obj := map[string]interface{}{
// 		"data": nil,
// 	}

// 	fooInt := ctx.FromInt("foo", 10)
// 	obj["data"] = fooInt
// 	return ctx.Json(http.StatusOK, obj)
// }

// func Fool3(ctx *framework.Context) error {
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379",
// 		Password: "",
// 		DB:       0,
// 	})

// 	return rdb.Set(ctx, "key", "value", 0).Err()
// }
