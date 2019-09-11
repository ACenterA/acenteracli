package cli

import (
	"encoding/json"
	"fmt"

	errors "github.com/pkg/errors"
	logger "github.com/wallix/awless/logger"
)

/*	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"gopkg.in/yaml.v2"
        errors "github.com/pkg/errors"

	"github.com/alecthomas/chroma/quick"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/auth"
	config "github.com/wallix/awless/config"
	global "github.com/wallix/awless/global"
	logger "github.com/wallix/awless/logger"
	"gopkg.in/h2non/gentleman.v2/context"
*/

/*
	url := "/customer/v1/websites/me"
		customer/v1/websites/logout
		req := cli.API().Path(url).Get()
		// cli.HandleBefore(handlerPath, params, req)
		resp, err := req.Do()
		if err != nil {
			fmt.Println(errors.Wrap(err, "Request failed"))
		}
		var decoded map[string]interface{}
		if resp.StatusCode < 400 {
			if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
				fmt.Println(errors.Wrap(err, "Unmarshalling response failed"))
			}
		} else {
			fmt.Println(errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String()))
		}
		if resp.StatusCode <= 208 && resp.StatusCode >= 200 {
			logger.Verbosef("Got decoded of %s", decoded)
			if _, ok := decoded["contactEmail"]; ok {
				fmt.Printf("Username: %s, Id: %s\n", decoded["contactEmail"], decoded["accountId"])
			}
		} else {
			logger.Error("Invalid access token.")
		}
*/

type AccountApi struct {
	Id   string `json:"projectId"`
	Name string `json:"name"`
}

func (api *APIInternal) Account() *AccountApi {
	return &AccountApi{}
}

// POST TO sites/v1/databases/create using projectId as param

func (api *AccountApi) GetAccount() (map[string]ProjectApi, error) {
	var resp map[string]ProjectApi
	itm, err := API().Get("/customer/v1/websites/me")
	if err != nil {
		return nil, err
	}
	jsonData, err := json.Marshal(itm)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(jsonData, &resp); err != nil {
		return nil, err
	}
	return resp, err
}

func (api *AccountApi) Logout() error {

	url := "customer/v1/websites/logout"
	req := API().Path(url).Post()

	// Method to be used
	// req.Method("POST")
	res, err := req.Do()
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return errors.Wrap(err, fmt.Sprintf("Request error: %s\n", err))
	}
	if !res.Ok {
		// fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return errors.New(fmt.Sprintf("Invalid server response: %d\n", res.StatusCode))
	}

	/*
		fmt.Printf("Status: %d\n", res.StatusCode)
		fmt.Printf("Body: %s", res.String())
	*/
	if !(res.StatusCode >= 200 && res.StatusCode <= 204) {
		logger.Verbosef("Body: %s", res.String())
		return errors.New(fmt.Sprintf("Invalid server response: %d\n", res.StatusCode))
	}
	return nil
}

/*
func (api *APIInternal) Project() (map[string]interface{}, err) {
	url := "/customer/v1/websites/me"
	req := cli.API().Path(url).Get()
	resp, err := req.Do()
	if err != nil {
		fmt.Println(errors.Wrap(err, "Request failed"))
	}
	var decoded map[string]interface{}
	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			fmt.Println(errors.Wrap(err, "Unmarshalling response failed"))
		}
	} else {
		fmt.Println(errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String()))
	}
	if resp.StatusCode <= 208 && resp.StatusCode >= 200 {
		logger.Verbosef("Got decoded of %s", decoded)
		if _, ok := decoded["project"]; ok {
			fmt.Printf("Username: %s, Id: %s\n", decoded["contactEmail"], decoded["accountId"])
		}
		return decoded
	} else {
		logger.Error("Invalid access token.")
	}
}
*/
