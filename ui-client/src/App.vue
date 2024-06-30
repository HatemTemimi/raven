<script setup>
import {onMounted, ref} from "vue";
import axios from "axios";

const proxies = ref([])
const target = ref('')
const loading = ref(true)

const fetchProxies = async () => {
  await axios.get('http://localhost:8080/api/fetch')
      .then((response)=>{
        proxies.value = response.data.slice(0, 33)
        loading.value = false

      })
}

const testProxies = async ()=>{
  proxies.value = []
  loading.value = true
  const response = await axios.get('http://localhost:8080/api/fetch/'+target.value)
      .then(()=>{
        proxies.value = response.data.slice(0, 33)
        loading.value = false
      })
}


onMounted(async ()=>{
  await fetchProxies()
})

</script>

<template>
  <div class="container" style="width: 100vw; height: 100vh; display: flex; justify-content: center ">
    <div  style="width:80%; height: auto; display: flex; flex-direction: column; gap: 1vh">
      <div style="height: 6%; display: flex; justify-content: center; align-items: center; margin-top: 1vh" class="bordered">
        Raven Proxy Lister & Tester
      </div>
      <div  class="bordered">
        <div style="padding: 4px; height: auto" class="bordered">
          <div style="text-align: center">
            <div style="display: flex; justify-content: center; ">
              <input v-model="target" style="width:25%" placeholder="type the target url you want to test against.."/>
              <button @click="testProxies">test</button>
            </div>
          </div>
          <div style="margin-top: 2vh">
            <div style="text-align: center" v-if="loading">loading..</div>
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
