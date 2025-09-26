wrk.method = "POST"
wrk.headers["Content-Type"] = "application/json"
wrk.headers["Accept"] = "application/json, application/problem+json"

-- 你要发送的 JSON 数据
wrk.body = [[
{
  "app_id": "string",
  "device_info": {
    "property1": "string",
    "property2": "string"
  },
  "event_type": "string",
  "payload": {
    "property1": null,
    "property2": null
  },
  "timestamp": 0,
  "user_id": "string"
}
]]

