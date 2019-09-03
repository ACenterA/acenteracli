package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	errors "github.com/pkg/errors"
	"gopkg.in/yaml.v2"

	"github.com/alecthomas/chroma/quick"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	config "github.com/wallix/awless/config"
	global "github.com/wallix/awless/global"
	logger "github.com/wallix/awless/logger"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/context"
	"gopkg.in/h2non/gentleman.v2/plugins/auth"
)

// HTTP Client Errors
var (
	ErrCannotUnmarshal = errors.New("Unable to unmarshal response")
)

func indent(value string) string {
	trimmed := strings.TrimSuffix(value, "\x1b[0m")
	trimmed = strings.TrimRight(trimmed, "\n")
	return "  ╎ " + strings.Replace(trimmed, "\n", "\n  ╎ ", -1) + "\n"
}

type APIInternal struct {
	client *gentleman.Client
}

func API() *APIInternal {
	v := &APIInternal{
		client: Client,
	}
	return v
}
func (api *APIInternal) Path(url string) *gentleman.Client {
	if (strings.HasSuffix(APIPrefix, "/")) && (strings.HasPrefix(url, "/")) {
		url = strings.TrimPrefix(url, "/")
	}
	return api.client.Path(APIPrefix + url)
}
func APIAnon() *APIInternal {
	v := &APIInternal{
		client: AnonClient,
	}
	return v
}

func (api *APIInternal) Projects() *ProjectApi {
	return &ProjectApi{}
}

func (api *APIInternal) Get(url string) (map[string]interface{}, error) {
	req := API().Path(url).Get()
	resp, err := req.Do()
	if err != nil {
		// fmt.Println(err)
		fmt.Println(errors.Wrap(err, "Request failed"))
	}
	var decoded map[string]interface{}
	if resp.StatusCode < 400 {
		//fmt.Println(err)
		if err := UnmarshalResponse(resp, &decoded); err != nil {
			fmt.Println(errors.Wrap(err, "Unmarshalling response failed"))
		}
	} else {
		//fmt.Println(err)
		fmt.Println(errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String()))
	}
	if resp.StatusCode <= 208 && resp.StatusCode >= 200 {
		logger.Verbosef("Got decoded of %s", decoded)
		if _, ok := decoded["id"]; ok {
			// fmt.Printf("Username: %s, Id: %s\n", decoded["contactEmail"], decoded["accountId"])
			return decoded, nil
		}
		return decoded, errors.New("Could not decode response.")
	} else {
		//fmt.Println(resp.StatusCode)
		logger.Error("Invalid access token.")
	}
	return nil, errors.New(fmt.Sprintf("Invalid response from server %d", resp.StatusCode))
}

func (api *APIInternal) GetRawData(url string) (string, error) {
	req := API().Path(url).Get()
	resp, err := req.Do()
	if err != nil {
		// fmt.Println(err)
		fmt.Println(errors.Wrap(err, "Request failed"))
	}

	data := resp.Bytes()
	if len(data) == 0 {
		return "", errors.New("Empty data")
	}
	return string(data), nil
}

func (api *APIInternal) GetByKey(url string, key string) (map[string]interface{}, error) {
	tmpData, err := api.GetRawData(url)
	if err != nil {
		return nil, err
	}
	var tmpObj map[string]interface{}
	if err := json.Unmarshal([]byte(tmpData), &tmpObj); err != nil {
		return nil, err
	}
	if tmpItems, ok := tmpObj[key]; ok {
		return tmpItems.(map[string]interface{}), nil
	}
	return nil, err
}

/*
func (api *APIInternal) Get(url string) *gentleman.Request {
    return Client.Path(url).Get()
}
func (api *APIInternal) Post(url string) *gentleman.Request {
    return Client.Path(.url)Post(url)
}
func (api *APIInternal) Put(url string) *gentleman.Request {
    return Client.Put(url)
}
func (api *APIInternal) Delete(url string) *gentleman.Request {
    return Client.Delete(url)
}
*/

// getBody returns and wraps the request/response body in a new reader, which
// is useful for logging purposes.
func getBody(r io.ReadCloser) (string, io.ReadCloser, error) {
	newReader := r

	body := ""
	if r != nil {
		data, err := ioutil.ReadAll(r)
		if err != nil {
			return "", nil, err
		}

		if len(data) > 0 {
			body = "\n" + string(data) + "\n"
			newReader = ioutil.NopCloser(bytes.NewReader(data))
		}
	}

	return body, newReader, nil
}

// UserAgentMiddleware sets the user-agent header on requests.
func UserAgentMiddleware(c *gentleman.Client) {
	ua, _ := config.Get("app-name")
	if ua == nil {
		ua = "ACenterA"
	}
	c.UseRequest(func(ctx *context.Context, h context.Handler) {
		ctx.Request.Header.Set("User-Agent", fmt.Sprintf("%s-cli-%s", ua, config.Version))
		h.Next(ctx)
	})
}

func PathMiddleware(c *gentleman.Client) {
	HostPath := global.API_ENDPOINT
	c.UseRequest(func(ctx *context.Context, h context.Handler) {
		// fmt.Println("Ned to append path to xisting path?" + HostPath)
		if strings.HasSuffix(HostPath, ".com/") {
		} else {
			// fmt.Println("ctx.Request.URL is ...", ctx.Request.URL)
			// fmt.Println(ctx.Request.URL)
		}
		h.Next(ctx)
	})
}

//AuthorizationMiddleware

func AuthorizationMiddleware(c *gentleman.Client) {
	c.UseRequest(func(ctx *context.Context, h context.Handler) {
		token := config.GetToken()
		if !(token == "") {
			if isTokenExpired(token) {
				token = ""
			} else {
			}
		}
		if len(token) <= 0 {
			// Ok we ua = "ACenterA"
			c := APIAnon().Path("/customer/v1/websites/login")
			c.Use(auth.Basic(config.GetUsername(), config.GetPasswordPlainText()))
			req := c.Post()
			req = req.AddHeader("Content-Type", "application/json")
			resp, err := req.Do()
			if err != nil {
				// fmt.Println(errors.Wrap(err, "Request failed"))
			}
			var decoded map[string]interface{}
			if resp.StatusCode < 400 {
				if err := UnmarshalResponse(resp, &decoded); err != nil {
					// fmt.Println(errors.Wrap(err, "Unmarshalling response failed"))
				}
			} else {
				logger.ExtraVerbosef("HTTP %d: %s", resp.StatusCode, resp.String())
			}
			if _, ok := decoded["accessToken"]; ok {
				config.Set("_token", decoded["accessToken"].(string))
				token = decoded["accessToken"].(string)
				config.Set("user.id", decoded["id"].(string))
			}
		} else {
			// Need to validate token ?
		}
		// alwasy cleanit up
		if token == "" {
			h.Error(ctx, errors.New("Invalid Username / Password. Please login and try again."))
		} else {
			ctx.Request.Header.Del("Authorization")
			// if (ctx.Request.Header.Get("Authorization") == "") {
			ctx.Request.Header.Set("Authorization", token)
			// }
			h.Next(ctx)
		}
	})
}

// LogMiddleware adds verbose log info to HTTP requests.
func LogMiddleware(c *gentleman.Client, useColor bool) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	c.UseRequest(func(ctx *context.Context, h context.Handler) {
		l := log.With().Str("request-id", fmt.Sprintf("%x", rnd.Uint64())).Logger()
		ctx.Set("log", &l)

		h.Next(ctx)
	})

	c.UseHandler("before dial", func(ctx *context.Context, h context.Handler) {
		ctx.Set("start", time.Now())

		log := ctx.Get("log").(*zerolog.Logger)

		// Make the request body available to downstream processors through the
		// request context as `request-body`.
		body, newReader, err := getBody(ctx.Request.Body)
		if err != nil {
			h.Error(ctx, err)
			return
		}
		ctx.Set("request-body", body)
		ctx.Request.Body = newReader

		if global.Verbose != "" {
			headers := ""
			for key, val := range ctx.Request.Header {
				headers += key + ": " + val[0] + "\n"
			}

			if body != "" {
				body = "\n" + body
			}

			http := fmt.Sprintf("%s %s %s\n%s%s", ctx.Request.Method, ctx.Request.URL, ctx.Request.Proto, headers, body)

			if useColor {
				sb := strings.Builder{}
				if err := quick.Highlight(&sb, http, "http", "terminal256", "cli-dark"); err != nil {
					h.Error(ctx, err)
				}
				http = sb.String()
			}

			log.Debug().Msgf("Making request:\n%s", indent(http))
		}

		h.Next(ctx)
	})

	c.UseResponse(func(ctx *context.Context, h context.Handler) {
		l := ctx.Get("log").(*zerolog.Logger)

		if global.Verbose != "" {
			headers := ""
			for key, val := range ctx.Response.Header {
				headers += key + ": " + val[0] + "\n"
			}

			body, newReader, err := getBody(ctx.Response.Body)
			if err != nil {
				h.Error(ctx, err)
				return
			}
			ctx.Response.Body = newReader

			http := fmt.Sprintf("%s %s\n%s\n%s", ctx.Response.Proto, ctx.Response.Status, headers, body)

			if useColor {
				sb := strings.Builder{}
				if err := quick.Highlight(&sb, http, "http", "terminal256", "cli-dark"); err != nil {
					h.Error(ctx, err)
				}
				http = sb.String()
			}

			l.Debug().Msgf("Got response in %s:\n%s", time.Since(ctx.Get("start").(time.Time)), indent(http))
		}

		h.Next(ctx)
	})
}

// UnmarshalRequest body into a given structure `s`. Supports both JSON and
// YAML depending on the request's content-type header.
func UnmarshalRequest(ctx *context.Context, s interface{}) error {
	return unmarshalBody(ctx.Request.Header, []byte(ctx.GetString("request-body")), s)
}

// UnmarshalResponse into a given structure `s`. Supports both JSON and
// YAML depending on the response's content-type header.
func UnmarshalResponse(resp *gentleman.Response, s interface{}) error {
	data := resp.Bytes()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("HTTP %d:\n%s", resp.StatusCode, string(data))
	}

	return unmarshalBody(resp.Header, data, s)
}

func unmarshalBody(headers http.Header, data []byte, s interface{}) error {
	if len(data) == 0 {
		return nil
	}

	ct := headers.Get("content-type")
	if strings.Contains(ct, "json") || strings.Contains(ct, "javascript") {
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}
	} else if strings.Contains(ct, "yaml") {
		if err := yaml.Unmarshal(data, &s); err != nil {
			return err
		}
	} else {
		// return fmt.Errorf("Not sure how to unmarshal %s", ct)
		logger.Errorf("Not sure how to unmarshal %s", ct)
		return nil
	}

	return nil
}

func isTokenExpired(tokenString string) bool {
	// fmt.Printf("The token is:\n%s\n", tokenString)
	var p jwt.Parser

	token, _, err := p.ParseUnverified(tokenString, &jwt.StandardClaims{})

	/*
	   for _, p := range b {
	           fmt.Printf("%s\n", p)
	   }
	*/
	if err = token.Claims.Valid(); err != nil {
		//handle invalid token
		// fmt.Println("INVALID TOKEN????")
	}

	switch err.(type) {
	case *jwt.ValidationError:
		vErr := err.(*jwt.ValidationError)
		switch vErr.Errors {
		case jwt.ValidationErrorExpired:
			return true
		default:
			return false
		}

	default:
		return false
	}
}
