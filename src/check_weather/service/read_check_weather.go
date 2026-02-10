package service

import (
	"context"
	"encoding/json"
	"net/http"

	"gitlab.com/wit-id/service-hub-test/toolkit/config"
	"gitlab.com/wit-id/service-hub-test/toolkit/log"
	"github.com/pkg/errors"
	"io"
	"gitlab.com/wit-id/service-hub-test/src/repository/payload"
)

func (s *CheckWeatherService) GetCheckWeather(ctx context.Context, client http.Client, cfg config.KVStore, queryparam payload.CheckWeatherPayload) (bodyRes json.RawMessage, err error) {
	var newReq *http.Request
	newReq, err = http.NewRequestWithContext(ctx, http.MethodGet,
		cfg.GetString("check-weather-api.base-url"), nil)

	if err != nil {
		log.FromCtx(ctx).Error(err, "invalid weather api base url")
		return nil, errors.Wrap(err, "failed create weather api request")
	}

	query := newReq.URL.Query()
	query.Add("q", queryparam.Longitude+","+queryparam.Latitude)
	query.Add("key", cfg.GetString("check-weather-api.api-key"))
	newReq.URL.RawQuery = query.Encode()

	// log.FromCtx(context.Background()).Info("Sending request to check weather api with query: %s", newReq.URL.RawQuery)

	var response *http.Response
	maxRetries := cfg.GetInt("check-weather-api.max-retry")

	for i := 0; i < maxRetries; i++ {
		response, err = client.Do(newReq)
		if err != nil {
			log.FromCtx(context.Background()).Error(err, "failed send request to check weather api")
			err = errors.Wrap(err, "")
			return
		}

		defer response.Body.Close()

		if response.StatusCode == http.StatusOK {
			break
		}

		// Log and retry if status code is not 200
		log.FromCtx(ctx).Info("Request failed with status code %d, retrying (%d/%d)", response.StatusCode, i+1, maxRetries)
		if i == maxRetries-1 {
			err = errors.New("maximum retries reached")
			log.FromCtx(ctx).Error(err, "failed request to check weather api after retries")
			return
		}
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to read response body")
		err = errors.Wrap(err, "failed to read response body")
		return
	}

	bodyRes = json.RawMessage(body)
	log.FromCtx(ctx).Info("weather api response: %s", string(body))

	return
}