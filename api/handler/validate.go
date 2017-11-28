package handler

import (
	"github.com/rancher/norman/parse"
	"github.com/rancher/norman/parse/builder"
	"github.com/rancher/norman/types"
)

func ParseAndValidateBody(apiContext *types.APIContext) (map[string]interface{}, error) {
	data, err := parse.Body(apiContext.Request)
	if err != nil {
		return nil, err
	}

	b := builder.NewBuilder(apiContext)

	data, err = b.Construct(apiContext.Schema, data, builder.Create)
	if err != nil {
		return nil, err
	}

	if apiContext.Schema.Validator != nil {
		if err := apiContext.Schema.Validator(apiContext, data); err != nil {
			return nil, err
		}
	}

	return data, nil
}
