import axios from "axios";
import router from "@/router";
import { useLoginStore } from "@/stores/login";

const service = axios.create({
  // baseURL: "http://localhost:3000",
  // baseURL: "http://192.168.10.8:3000",
  baseURL: "/api",
});

service.interceptors.request.use(
  function (config) {
    const token = localStorage.getItem("token");
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  function (error) {
    return Promise.reject(error);
  }
);

service.interceptors.response.use(
  function (response) {
    return response.data;
  },
  function (error) {
    const loginStore = useLoginStore();
    if (error.response.status === 401) {
      router.push({ name: "login" });
      loginStore.removeToken();
    }
    return Promise.reject(error);
  }
);

export default service;
