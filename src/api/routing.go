package api

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
)

type GetRes struct {
  Result string
}


/**
 * Query the API for the registry contract address.
 *
 * @param  api    Full base URI of the api
 * @return        (contract address, error)
 */
func GetRegistry(api string) (string, error) {
  var result = new(GetRes)
  res, err := http.Get(api+"/Registry")
  if err != nil {
    return "", fmt.Errorf("Could not get registry address: ", err)
  } else {
    body, err2 := ioutil.ReadAll(res.Body)
    if err2 != nil {
      return "", fmt.Errorf("Could not read response body: ", err2)
    } else {
      err3 := json.Unmarshal(body, &result)
      if err3 != nil {
        return "", fmt.Errorf("Could not unmarshal response body: ", err3)
      }
    }
  }
  return result.Result, nil
}

/**
 * Query the API for the USDX token contract address.
 *
 * @param  api    Full base URI of the api
 * @return        (contract address, error)
 */
func GetUSDX(api string) (string, error) {
  var result = new(GetRes)
  res, err := http.Get(api+"/USDX")
  if err != nil {
    return "", fmt.Errorf("Could not get registry address: ", err)
  } else {
    body, err2 := ioutil.ReadAll(res.Body)
    if err2 != nil {
      return "", fmt.Errorf("Could not read response body: ", err2)
    } else {
      err3 := json.Unmarshal(body, &result)
      if err3 != nil {
        return "", fmt.Errorf("Could not unmarshal response body: ", err3)
      }
    }
  }
  return result.Result, nil
}
