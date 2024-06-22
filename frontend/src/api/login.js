import service from "@/utils/axios";

// 用户登录
export function userlogin(data) {
  return service({
    method: "POST",
    url: "/login",
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
    data,
  });
}

// 是否初始化
export function isSetup() {
  return service({
    method: "GET",
    url: "/user",
  });
}

// 初始化用户
export function setupUser(data) {
  return service({
    method: "POST",
    url: "/user",
    data,
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
  });
}
