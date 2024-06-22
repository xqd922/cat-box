<template>
  <v-main>
    <v-container>
      <v-sheet
        class="mx-auto d-flex flex-column justify-center mb-10"
        max-width="390"
      >
        <v-img
          :width="100"
          cover
          class="mx-auto mt-6 mb-3"
          src="@/assets/logo.svg"
        ></v-img>
        <h2 class="text-h4 text-center font-weight-black title-color">
          sub-box
        </h2>
      </v-sheet>
      <v-card class="mx-auto" elevation="0" max-width="390" rounded="lg">
        <v-card-text>
          <v-form v-model="valid">
            <v-text-field
              v-model="username"
              :rules="inputRules"
              placeholder="用户"
              prepend-inner-icon="mdi-account-circle-outline"
              variant="outlined"
            ></v-text-field>

            <v-text-field
              v-model="password"
              :rules="inputRules"
              :append-inner-icon="visible ? 'mdi-eye-off' : 'mdi-eye'"
              :type="visible ? 'text' : 'password'"
              placeholder="密码"
              prepend-inner-icon="mdi-lock-outline"
              variant="outlined"
              @click:append-inner="visible = !visible"
            ></v-text-field>

            <v-btn
              color="primary"
              :disabled="!valid"
              block
              rounded
              @click="submit"
            >
              {{ switchText }}
            </v-btn>
          </v-form>
        </v-card-text>
      </v-card>
    </v-container>
  </v-main>
</template>

<script setup>
import { ref, reactive, computed, onBeforeMount } from "vue";
import { userlogin, isSetup, setupUser } from "@/api/login";
import { useLoginStore } from "@/stores/login";
import { useRouter } from "vue-router";

const router = useRouter();

let valid = ref(true);
let visible = ref(false);
let isRegistered = ref(true);
let username = ref("");
let password = ref("");
let inputRules = reactive([(value) => !!value]);
let loginStore = useLoginStore();

onBeforeMount(async () => {
  let res = await isSetup();
  isRegistered.value = res.data.is_registered;
});

let switchText = computed(() => {
  if (isRegistered.value) {
    return "登录";
  } else {
    return "注册";
  }
});

async function submit() {
  if (isRegistered.value) {
    let res = await userlogin({
      username: username.value,
      password: password.value,
    });
    loginStore.setToken(res.data.token);
    router.push({ name: "home" });
  } else {
    await setupUser({
      username: username.value,
      password: password.value,
    });
    isRegistered.value = true;
  }
}
</script>

<style lang="css" scoped>
.title-color {
  color: #546e7a;
}
</style>
