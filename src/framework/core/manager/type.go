/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package manager

import (
	"configcenter/src/framework/core/types"
	"github.com/tidwall/gjson"
)

const (
	EventHost           = "host"
	EventHostIdentifier = "hostidentifier"
	EventSet            = "set"
	EventModule         = "module"
	EventBusiness       = "biz"
	EventPlat           = "plat"

	EventModuleTransfer = "moduletransfer"
	EventActionCreate   = "create"
	EventActionDelete   = "delete"
	EventActionUpdate   = "update"
)

// Action the http action
type Action struct {
	Method      string
	Path        string
	HandlerFunc func(data *gjson.Result) (types.MapStr, error)
}

// FrameworkContext definition the framework context
type FrameworkContext interface {
	RegisterEvent(eventType types.EventType, eventFunc types.EventCallbackFunc) types.EventKey
	UnRegisterEvent(eventKey types.EventKey)
}
