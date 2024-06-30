<script setup>
import {onMounted, ref} from "vue";
import axios from "axios";

const proxies = ref([])
const target = ref('')
const loading = ref(true)

const fetchProxies = async () => {
  loading.value = true
  await axios.get('http://localhost:8080/api/fetch')
      .then((response)=>{
        proxies.value = response.data.slice(0, 29)
        loading.value = false
      })
}

const refresh = ()=>{
  proxies.value = []
  fetchProxies()
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
    <div style="width:80%; height: auto; display: flex; flex-direction: column; gap: 1vh">
      <div style="height: 6%; display: flex; justify-content: center; align-items: center; margin-top: 1vh"
           class="bordered">
        Raven Proxy Lister & Tester
      </div>
      <div class="bordered">
        <div style=" height: auto">
          <div class="tools" style="margin-top: 2vh; display: flex; justify-content: space-between;">
            <button @click="refresh" style="margin-left: 10px">Refresh</button>
            <div class="search" style="width: 80%; display: flex; justify-content: center">
              <input v-model="target" style="width:25%" placeholder="type the target url you want to test against.."/>
              <button style="margin-left: 0.5rem" @click="testProxies">Test</button>
            </div>
            <button style="margin-right: 10px">Copy to Clipboard</button>
          </div>
          <div class="content" style="padding: 10px">
            <div style="margin-top: 2vh; margin-bottom: 1vh">
              <div style="text-align: center" v-if="loading">
                hold on a second..
              </div>
              <div v-if="!loading" style=" width:100%; display: flex; justify-content: space-between; flex-direction: row;"
                   class="bordered">
                <div style="margin-left: 5px">IP:Port</div>
                <div>Status</div>
                <div style="margin-right: 5px">Speed</div>
              </div>
              <div class="bordered" style="margin-top: 0.5vh; padding: 0.2rem">
                <div v-for="proxy of proxies">
                  <div style="display: flex; justify-content: space-between; flex-direction: row;">
                    <div>
                      {{ proxy.Ip }}: {{ proxy.Port }}
                    </div>
                    <div>
                      {{ proxy.Status }}
                    </div>
                    <div>
                      {{ proxy.Speed }}
                    </div>
                  </div>
                  <div class="line"></div>
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
.line {
  flex: 1;
  border-bottom: 1px dashed black;
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
