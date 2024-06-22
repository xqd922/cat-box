import { defineStore } from "pinia";
import { ref } from "vue";

export const useLoginStore = defineStore("login", () => {
  let token = ref(localStorage.getItem("token"));

  function setToken(resToken) {
    token.value = resToken;
    localStorage.setItem("token", resToken);
  }

  function removeToken() {
    token.value = null;
    localStorage.removeItem("token");
  }
  return { token, setToken, removeToken };
});
