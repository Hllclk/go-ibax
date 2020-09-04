/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package api
}

func (m Mode) getAppContentHandler(w http.ResponseWriter, r *http.Request) {
	form := &appParamsForm{
		ecosystemForm: ecosystemForm{
			Validator: m.EcosysIDValidator,
		},
	}

	if err := parseForm(r, form); err != nil {
		errorResponse(w, err, http.StatusBadRequest)
		return
	}

	logger := getLogger(r)
	params := mux.Vars(r)

	bi := &model.BlockInterface{}
	p := &model.Page{}
	c := &model.Contract{}
	appID := converter.StrToInt64(params["appID"])
	ecosystemID := converter.StrToInt64(form.EcosystemPrefix)

	blocks, err := bi.GetByApp(appID, ecosystemID)
	if err != nil {
		logger.WithFields(log.Fields{"type": consts.DBError, "error": err}).Error("Getting block interfaces by appID")
		errorResponse(w, err)
		return
	}

	pages, err := p.GetByApp(appID, ecosystemID)
	if err != nil {
		logger.WithFields(log.Fields{"type": consts.DBError, "error": err}).Error("Getting pages by appID")
		errorResponse(w, err)
		return
	}

	contracts, err := c.GetByApp(appID, ecosystemID)
	if err != nil {
		logger.WithFields(log.Fields{"type": consts.DBError, "error": err}).Error("Getting pages by appID")
		errorResponse(w, err)
		return
	}

	jsonResponse(w, &appContentResult{
		Blocks:    blocks,
		Pages:     pages,
		Contracts: contracts,
	})
}