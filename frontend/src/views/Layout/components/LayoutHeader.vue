<template>
  <v-app-bar>
    <template v-slot:prepend>
      <v-img
        src="@/assets/logo.svg"
        width="35"
        transition="scale-transition"
      ></v-img>
    </template>
    <v-app-bar-title>sub-box</v-app-bar-title>
    <template v-slot:append>
      <v-btn
        variant="outlined"
        rounded
        color="primary"
        prepend-icon="mdi-vpn"
        class="mr-2"
        link
        :href="toProxyManager"
      >
        代理面板
      </v-btn>
      <v-btn icon="mdi-cog-outline" @click="drawer = !drawer"></v-btn>
      <v-btn icon="mdi-logout" @click="signOut"> </v-btn>
    </template>
  </v-app-bar>
</template>

<script setup>
import { ref } from "vue";
import { storeToRefs } from "pinia";
import { useDrawerStore } from "@/stores/drawer";
const drawerStore = useDrawerStore();
const { drawer } = storeToRefs(drawerStore);

// 退出登录
import { useRouter } from "vue-router";
import { useLoginStore } from "@/stores/login";
const router = useRouter();
const loginStore = useLoginStore();
function signOut() {
  loginStore.removeToken();
  router.push({ name: "login" });
}

// 代理面板地址
const toProxyManager = ref(`http://${window.location.hostname}:9090/ui/`);
</script>

<style lang="scss" scoped></style>
