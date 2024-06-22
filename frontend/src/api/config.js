import service from "@/utils/axios";

// 获取配置
export function getDelay() {
  return service({
    method: "GET",
    url: "/option",
  });
}

// 更新配置
export function setDelay(data) {
  return service({
    method: "POST",
    url: "/option",
    data,
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
  });
}

// 停止sing-box
export function stopSingBox() {
  return service({
    method: "PUT",
    url: "/singbox",
  });
}
