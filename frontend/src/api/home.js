import service from "@/utils/axios";

// 获取订阅
export function get_sub() {
  return service({
    method: "GET",
    url: "/subscribe",
  });
}

// 添加订阅
export function add_sub(data) {
  return service({
    method: "POST",
    url: "/subscribe",
    data,
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
  });
}

// 删除订阅
export function rm_sub(id) {
  return service({
    method: "DELETE",
    url: `/subscribe/${id}`,
  });
}

// 修改订阅
export function rw_sub(id, data) {
  return service({
    method: "PUT",
    url: `/subscribe/${id}`,
    data,
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
  });
}

// 更新订阅
export function up_sub(id, queryParams = {}) {
  return service({
    method: "PATCH",
    url: `/subscribe/${id}`,
    params: queryParams,
  });
}

// 自动更新
export function auto_sub(id, data) {
  return service({
    method: "PATCH",
    url: `/subscribe/${id}`,
    params: { sw_update: "1" },
    data,
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
  });
}

// 切换订阅
export function sw_sub(id) {
  return service({
    method: "PATCH",
    url: `/subscribe/${id}`,
  });
}
