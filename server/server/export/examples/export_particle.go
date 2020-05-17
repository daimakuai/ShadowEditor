// Copyright 2017-2020 The ShadowEditor Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
//
// For more information, please visit: https://github.com/tengge1/ShadowEditor
// You can also visit: https://gitee.com/tengge1/ShadowEditor

package examples

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/tengge1/shadoweditor/server/helper"
	"github.com/tengge1/shadoweditor/server/server"
)

func exportParticle(path string) {
	port := server.Config.Server.Port

	result, _ := helper.Get(fmt.Sprintf("http://%v/api/Particle/List", port))

	dirName := filepath.Join(path, "api", "Particle")
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		os.MkdirAll(dirName, 0755)
	}

	// get list
	fileName := filepath.Join(path, "api", "Particle", "List")
	ioutil.WriteFile(fileName, []byte(result), 0755)

	// export particles
	var obj map[string]interface{}
	helper.FromJSON([]byte(result), &obj)

	if array, ok := obj["Data"].([]interface{}); ok {
		for _, i := range array {
			id := i.(map[string]interface{})["ID"].(string)
			result, _ = helper.Get(fmt.Sprintf("http://%v/api/Particle/Get?ID=%v", port, id))
			fileName = fmt.Sprintf("%v/api/Particle/Particle_%v", path, id)
			ioutil.WriteFile(fileName, []byte(result), 0755)
		}
	}

	// other apis
	apiList := []string{
		"/api/Particle/Edit",
		"/api/Particle/Save",
		"/api/Particle/Delete",
	}

	data, _ := helper.ToJSON(map[string]interface{}{
		"Code": 200,
		"Msg":  "Execute successfully!",
	})

	for _, i := range apiList {
		fileName = filepath.Join(path, i)
		ioutil.WriteFile(fileName, []byte(data), 0755)
	}
}
