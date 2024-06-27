<template>
  <v-navigation-drawer v-model="drawer" location="right" temporary>
    <v-toolbar>
      <v-toolbar-title>设置</v-toolbar-title>

      <v-spacer></v-spacer>

      <v-btn icon="mdi-close" @click="drawer = !drawer"></v-btn>
    </v-toolbar>
    <v-container>
      <v-text-field
        v-model="updateDelay"
        label="更新延时"
        clearable
        @blur="setUpdateDelay"
      ></v-text-field>
    </v-container>
  </v-navigation-drawer>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { getDelay, setDelay } from "@/api/config";
import { storeToRefs } from "pinia";
import { useDrawerStore } from "@/stores/drawer";
const drawerStore = useDrawerStore();
const { drawer } = storeToRefs(drawerStore);

// 设置抽屉
const updateDelay = ref("");
onMounted(async () => {
  const { data } = await getDelay();
  updateDelay.value = data.update_delay;
});
async function setUpdateDelay() {
  await setDelay({ update_delay: updateDelay.value });
}
</script>

<style lang="scss" scoped></style>
