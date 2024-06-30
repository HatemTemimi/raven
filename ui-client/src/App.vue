<script setup>

import {onMounted, ref} from "vue";
import axios from "axios";

const proxies = ref([])

onMounted(async ()=>{
  const response = await axios.get('http://localhost:8080/api/fetch')
  console.log(response)
  proxies.value = response.data
})

</script>

<template>
  <div class="container" style="width: 100vw; height: 100vh; display: flex; justify-content: center ">
    <div  style="width:80%; height: auto; display: flex; flex-direction: column;">
      <div style="text-align: center; height: 5%" class="bordered">
        Raven Proxy Lister & Tester
      </div>
      <div style="height: auto" class="bordered">
        <div style="padding: 4px; height: auto" class="bordered">
          <div style="text-align: center">
            the proxy table
          </div>
          <div v-for="proxy of proxies">
            <div style="display: flex; justify-content: space-evenly; flex-direction: row">
              <div style="max-width: 30%">
                  {{ proxy.Ip }}
              </div>

                <div style="max-width: 30%">
                {{ proxy.Port }}
              </div>

              <div style="max-width: 30%">
                {{ proxy.Status }}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}

.bordered {
  border-width: 2px;
  border-color: black;
  border-style: solid;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
