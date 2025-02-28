// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/farms": {
            "get": {
                "description": "Get all farms on the grid, It has pagination",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GridProxy"
                ],
                "summary": "Show farms on the grid",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Max result per page",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Min number of free ips in the farm",
                        "name": "free_ips",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Pricing policy id",
                        "name": "pricing_policy_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "farm version",
                        "name": "version",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "farm id",
                        "name": "farm_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "twin id associated with the farm",
                        "name": "twin_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "farm name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "farm stellar_address",
                        "name": "stellar_address",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/db.Farm"
                            }
                        }
                    }
                }
            }
        },
        "/gateways": {
            "get": {
                "description": "Get all nodes on the grid, It has pagination",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GridProxy"
                ],
                "summary": "Show nodes on the grid",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Max result per page",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Min free reservable mru in bytes",
                        "name": "free_mru",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Min free reservable hru in bytes",
                        "name": "free_hru",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Min free reservable sru in bytes",
                        "name": "free_sru",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Min number of free ips in the farm of the node",
                        "name": "free_ips",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Node status filter, up/down.",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Node city filter",
                        "name": "city",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Node country filter",
                        "name": "country",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Get nodes for specific farm",
                        "name": "farm_name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Set to true to filter nodes with ipv4",
                        "name": "ipv4",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Set to true to filter nodes with ipv6",
                        "name": "ipv6",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Set to true to filter nodes with domain",
                        "name": "domain",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "List of farms separated by comma to fetch nodes from (e.g. '1,2,3')",
                        "name": "farm_ids",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/explorer.node"
                            }
                        }
                    }
                }
            }
        },
        "/gateways/{node_id}": {
            "get": {
                "description": "Get all details for specific node hardware, capacity, DMI, hypervisor",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GridProxy"
                ],
                "summary": "Show the details for specific node",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Node ID",
                        "name": "node_id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/explorer.node"
                        }
                    }
                }
            }
        },
        "/nodes": {
            "get": {
                "description": "Get all nodes on the grid, It has pagination",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GridProxy"
                ],
                "summary": "Show nodes on the grid",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Max result per page",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Min free reservable mru in bytes",
                        "name": "free_mru",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Min free reservable hru in bytes",
                        "name": "free_hru",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Min free reservable sru in bytes",
                        "name": "free_sru",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Min number of free ips in the farm of the node",
                        "name": "free_ips",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Node status filter, up/down.",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Node city filter",
                        "name": "city",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Node country filter",
                        "name": "country",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Get nodes for specific farm",
                        "name": "farm_name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Set to true to filter nodes with ipv4",
                        "name": "ipv4",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Set to true to filter nodes with ipv6",
                        "name": "ipv6",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Set to true to filter nodes with domain",
                        "name": "domain",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "List of farms separated by comma to fetch nodes from (e.g. '1,2,3')",
                        "name": "farm_ids",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/explorer.node"
                            }
                        }
                    }
                }
            }
        },
        "/nodes/{node_id}": {
            "get": {
                "description": "Get all details for specific node hardware, capacity, DMI, hypervisor",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GridProxy"
                ],
                "summary": "Show the details for specific node",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Node ID",
                        "name": "node_id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/explorer.node"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "ping the server to check if it running",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ping"
                ],
                "summary": "ping the server",
                "responses": {
                    "200": {
                        "description": "pong",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/stats": {
            "get": {
                "description": "Get statistics about the grid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GridProxy"
                ],
                "summary": "Show stats about the grid",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/db.Counters"
                            }
                        }
                    }
                }
            }
        },
        "/twin/{twin_id}": {
            "post": {
                "description": "submit the message",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RMB"
                ],
                "summary": "submit the message",
                "parameters": [
                    {
                        "description": "rmb.Message",
                        "name": "msg",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rmbproxy.Message"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "twin id",
                        "name": "twin_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rmbproxy.MessageIdentifier"
                        }
                    }
                }
            }
        },
        "/twin/{twin_id}/{retqueue}": {
            "get": {
                "description": "Get the message result",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RMB"
                ],
                "summary": "Get the message result",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "twin id",
                        "name": "twin_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "message retqueue",
                        "name": "retqueue",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/rmbproxy.Message"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "db.Counters": {
            "type": "object",
            "properties": {
                "accessNodes": {
                    "type": "integer"
                },
                "contracts": {
                    "type": "integer"
                },
                "countries": {
                    "type": "integer"
                },
                "farms": {
                    "type": "integer"
                },
                "gateways": {
                    "type": "integer"
                },
                "nodes": {
                    "type": "integer"
                },
                "publicIps": {
                    "type": "integer"
                },
                "totalCru": {
                    "type": "integer"
                },
                "totalHru": {
                    "type": "integer"
                },
                "totalMru": {
                    "type": "integer"
                },
                "totalSru": {
                    "type": "integer"
                },
                "twins": {
                    "type": "integer"
                }
            }
        },
        "db.Farm": {
            "type": "object",
            "properties": {
                "farmId": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "pricingPolicyId": {
                    "type": "integer"
                },
                "publicIps": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/db.PublicIP"
                    }
                },
                "stellarAddress": {
                    "type": "string"
                },
                "twinId": {
                    "type": "integer"
                },
                "version": {
                    "type": "integer"
                }
            }
        },
        "db.PublicConfig": {
            "type": "object",
            "properties": {
                "domain": {
                    "type": "string"
                },
                "gw4": {
                    "type": "string"
                },
                "gw6": {
                    "type": "string"
                },
                "ipv4": {
                    "type": "string"
                },
                "ipv6": {
                    "type": "string"
                }
            }
        },
        "db.PublicIP": {
            "type": "object",
            "properties": {
                "contractId": {
                    "type": "integer"
                },
                "farmId": {
                    "type": "string"
                },
                "gateway": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "ip": {
                    "type": "string"
                }
            }
        },
        "explorer.location": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                }
            }
        },
        "explorer.node": {
            "type": "object",
            "properties": {
                "certificationType": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "created": {
                    "type": "integer"
                },
                "farmId": {
                    "type": "integer"
                },
                "farmingPolicyId": {
                    "type": "integer"
                },
                "gridVersion": {
                    "type": "integer"
                },
                "hypervisor": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "location": {
                    "$ref": "#/definitions/explorer.location"
                },
                "nodeId": {
                    "type": "integer"
                },
                "proxyUpdatedAt": {
                    "type": "integer"
                },
                "publicConfig": {
                    "$ref": "#/definitions/db.PublicConfig"
                },
                "status": {
                    "description": "added node status field for up or down",
                    "type": "string"
                },
                "total_resources": {
                    "$ref": "#/definitions/gridtypes.Capacity"
                },
                "twinId": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "uptime": {
                    "type": "integer"
                },
                "used_resources": {
                    "$ref": "#/definitions/gridtypes.Capacity"
                },
                "version": {
                    "type": "integer"
                },
                "zosVersion": {
                    "type": "string"
                }
            }
        },
        "gridtypes.Capacity": {
            "type": "object",
            "properties": {
                "cru": {
                    "type": "integer"
                },
                "hru": {
                    "type": "integer"
                },
                "ipv4u": {
                    "type": "integer"
                },
                "mru": {
                    "type": "integer"
                },
                "sru": {
                    "type": "integer"
                }
            }
        },
        "rmbproxy.Message": {
            "type": "object",
            "properties": {
                "cmd": {
                    "type": "string",
                    "example": "zos.statistics.get"
                },
                "dat": {
                    "type": "string"
                },
                "dst": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "err": {
                    "type": "string"
                },
                "exp": {
                    "type": "integer",
                    "example": 0
                },
                "now": {
                    "type": "integer",
                    "example": 1631078674
                },
                "ret": {
                    "type": "string"
                },
                "shm": {
                    "type": "string"
                },
                "src": {
                    "type": "integer",
                    "example": 1
                },
                "try": {
                    "type": "integer",
                    "example": 2
                },
                "uid": {
                    "type": "string"
                },
                "ver": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "rmbproxy.MessageIdentifier": {
            "type": "object",
            "properties": {
                "retqueue": {
                    "type": "string",
                    "example": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Grid Proxy Server API",
	Description: "grid proxy server has the main methods to list farms, nodes, node details in the grid.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
