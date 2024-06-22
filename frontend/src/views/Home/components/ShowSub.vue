<template>
  <v-row class="mt-3">
    <v-col v-for="item in subData" :key="item.id" cols="12" md="6" lg="6">
      <v-card
        elevation="2"
        :class="{ active: item.active }"
        @click="swOne(item.id)"
      >
        <v-card-item>
          <v-card-title>{{ item.name }}</v-card-title>
          <v-card-subtitle class="my-3">
            {{ item.link }}
          </v-card-subtitle>
        </v-card-item>

        <v-card-actions>
          <v-spacer></v-spacer>
          <EditSub :item="item" />
        </v-card-actions>
      </v-card>
    </v-col>
  </v-row>
</template>

<script setup>
import EditSub from "./EditSub.vue";
import { ref, onBeforeMount, onUnmounted } from "vue";
import { get_sub, sw_sub } from "@/api/home";
import emitter from "@/utils/emitter";

let subData = ref([]);

async function fetchData() {
  let res = await get_sub();
  subData.value = res.data;
}
onBeforeMount(async () => {
  fetchData();
});

emitter.on("reloadData", () => {
  fetchData();
});

onUnmounted(() => {
  emitter.off("reloadData");
});

async function swOne(id) {
  await sw_sub(id);
  fetchData();
}
</script>

<style lang="css" scoped>
.active {
  border-left: 5px solid rgb(24, 103, 192);
}
</style>
